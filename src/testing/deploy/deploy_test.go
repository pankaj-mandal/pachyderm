//go:build k8s

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/pachyderm/pachyderm/v2/src/auth"
	"github.com/pachyderm/pachyderm/v2/src/client"
	"github.com/pachyderm/pachyderm/v2/src/identity"
	"github.com/pachyderm/pachyderm/v2/src/internal/errors"
	"github.com/pachyderm/pachyderm/v2/src/internal/minikubetestenv"
	"github.com/pachyderm/pachyderm/v2/src/internal/require"
	"github.com/pachyderm/pachyderm/v2/src/internal/testutil"
)

var valueOverrides map[string]string

func TestInstallAndUpgradeEnterpriseWithEnv(t *testing.T) {
	t.Parallel()
	ns, portOffset := minikubetestenv.ClaimCluster(t)
	k := testutil.GetKubeClient(t)
	opts := &minikubetestenv.DeployOpts{
		AuthUser:   auth.RootUser,
		Enterprise: true,
		PortOffset: portOffset,
	}
	valueOverrides = map[string]string{"pachd.logLevel": "debug"}
	opts.ValueOverrides = valueOverrides
	// Test Install
	minikubetestenv.PutNamespace(t, ns)
	c := minikubetestenv.InstallRelease(t, context.Background(), ns, k, opts)
	whoami, err := c.AuthAPIClient.WhoAmI(c.Ctx(), &auth.WhoAmIRequest{})
	require.NoError(t, err)
	require.Equal(t, auth.RootUser, whoami.Username)
	c.SetAuthToken("")
	mockIDPLogin(t, c)
	// Test Upgrade
	// opts.CleanupAfter = true
	// set new root token via env
	opts.AuthUser = ""
	token := "new-root-token"
	opts.ValueOverrides = map[string]string{"pachd.rootToken": token, "pachd.logLevel": "debug"}
	// add config file with trusted peers & new clients
	opts.ValuesFiles = []string{createAdditionalClientsFile(t), createTrustedPeersFile(t)}
	// apply upgrade
	c = minikubetestenv.UpgradeRelease(t, context.Background(), ns, k, opts)
	// c.SetAuthToken(token)
	// whoami, err = c.AuthAPIClient.WhoAmI(c.Ctx(), &auth.WhoAmIRequest{})
	// require.NoError(t, err)
	// require.Equal(t, auth.RootUser, whoami.Username)
	// // old token should no longer work
	// c.SetAuthToken(testutil.RootToken)
	// _, err = c.AuthAPIClient.WhoAmI(c.Ctx(), &auth.WhoAmIRequest{})
	// require.YesError(t, err)
	// time.Sleep(1800 * time.Millisecond)
	c.SetAuthToken("")
	mockIDPLogin(t, c)
	// assert new trusted peer and client
	resp, err := c.IdentityAPIClient.GetOIDCClient(c.Ctx(), &identity.GetOIDCClientRequest{Id: "pachd"})
	require.NoError(t, err)
	require.EqualOneOf(t, resp.Client.TrustedPeers, "example-app")
	_, err = c.IdentityAPIClient.GetOIDCClient(c.Ctx(), &identity.GetOIDCClientRequest{Id: "example-app"})
	require.NoError(t, err)
}

func TestEnterpriseServerMember(t *testing.T) {
	t.Parallel()
	ns, portOffset := minikubetestenv.ClaimCluster(t)
	k := testutil.GetKubeClient(t)
	minikubetestenv.PutNamespace(t, "enterprise")
	ec := minikubetestenv.InstallRelease(t, context.Background(), "enterprise", k, &minikubetestenv.DeployOpts{
		AuthUser:         auth.RootUser,
		EnterpriseServer: true,
		CleanupAfter:     true,
		ValueOverrides:   valueOverrides,
	})
	whoami, err := ec.AuthAPIClient.WhoAmI(ec.Ctx(), &auth.WhoAmIRequest{})
	require.NoError(t, err)
	require.Equal(t, auth.RootUser, whoami.Username)
	mockIDPLogin(t, ec)
	minikubetestenv.PutNamespace(t, ns)
	c := minikubetestenv.InstallRelease(t, context.Background(), ns, k, &minikubetestenv.DeployOpts{
		AuthUser:         auth.RootUser,
		EnterpriseMember: true,
		Enterprise:       true,
		PortOffset:       portOffset,
		CleanupAfter:     true,
		ValueOverrides:   valueOverrides,
	})
	whoami, err = c.AuthAPIClient.WhoAmI(c.Ctx(), &auth.WhoAmIRequest{})
	require.NoError(t, err)
	require.Equal(t, auth.RootUser, whoami.Username)
	c.SetAuthToken("")
	loginInfo, err := c.GetOIDCLogin(c.Ctx(), &auth.GetOIDCLoginRequest{})
	require.NoError(t, err)
	require.True(t, strings.Contains(loginInfo.LoginURL, ":31658"))
	mockIDPLogin(t, c)
}

func mockIDPLogin(t testing.TB, c *client.APIClient) {
	// login using mock IDP admin
	hc := &http.Client{}
	c.SetAuthToken("")
	loginInfo, err := c.GetOIDCLogin(c.Ctx(), &auth.GetOIDCLoginRequest{})
	require.NoError(t, err)
	state := loginInfo.State

	// hc.CheckRedirect = func(req *http.Request, via []*http.Request) error {
	// 	return fmt.Errorf("Not redirecting to %s", req.URL)
	// }
	require.NoErrorWithinTRetry(t, 2*time.Minute, func() error {
		// Get the initial URL from the grpc, which should point to the dex login page
		resp, err := hc.Get(loginInfo.LoginURL)
		respBytes, _ := ioutil.ReadAll(resp.Body)
		respString := string(respBytes)
		require.NoError(t, err)
		fmt.Printf("LOGIN INFO RESP: %s - code: %d\n", resp.Request.URL.String(), resp.StatusCode)
		if resp.StatusCode != 200 {
			return errors.EnsureStack(fmt.Errorf("Recieved an invalid response retrieving DEX url! Code: %d, response url: %s \n", resp.StatusCode, respString))
		}

		vals := make(url.Values)
		vals.Add("login", "admin")
		vals.Add("password", "password")

		loginResp, err := hc.PostForm(resp.Request.URL.String(), vals)
		if err != nil {
			return err
		}
		// require.NoError(t, err)
		loginRespBytes, _ := ioutil.ReadAll(loginResp.Body)
		loginRespString := string(loginRespBytes)
		fmt.Printf("LOGIN RESP: %s - code: %d\n", loginRespString, loginResp.StatusCode)
		if loginResp.StatusCode != 200 || !strings.Contains(loginRespString, "You are now logged in") { // TODO instead of message - don't redirect on >=400 code
			return errors.EnsureStack(fmt.Errorf("Recieved an invalid response logging into DEX! Code: %d, response body: %s \n", loginResp.StatusCode, loginRespString))
		}
		return nil
	})

	authResp, err := c.AuthAPIClient.Authenticate(c.Ctx(), &auth.AuthenticateRequest{OIDCState: state})
	require.NoError(t, err)
	c.SetAuthToken(authResp.PachToken)
	whoami, err := c.AuthAPIClient.WhoAmI(c.Ctx(), &auth.WhoAmIRequest{})
	require.NoError(t, err)
	require.Equal(t, "user:"+testutil.DexMockConnectorEmail, whoami.Username)
}

func createTrustedPeersFile(t testing.TB) string {
	data := []byte(`pachd:
  additionalTrustedPeers:
    - example-app
`)
	tf, err := os.CreateTemp("", "pachyderm-trusted-peers-*.yaml")
	require.NoError(t, err)
	_, err = tf.Write(data)
	require.NoError(t, err)
	return tf.Name()
}

func createAdditionalClientsFile(t testing.TB) string {
	data := []byte(`oidc:
  additionalClients:
    - id: example-app
      secret: example-app-secret
      name: 'Example App'
      redirectURIs:
      - 'http://127.0.0.1:5555/callback'
`)
	tf, err := os.CreateTemp("", "pachyderm-additional-clients-*.yaml")
	require.NoError(t, err)
	_, err = tf.Write(data)
	require.NoError(t, err)
	return tf.Name()
}
