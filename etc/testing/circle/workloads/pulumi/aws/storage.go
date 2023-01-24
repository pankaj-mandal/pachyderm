package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func DeployBucket(ctx *pulumi.Context) error {
	bucketName := fmt.Sprintf("s3-%s-bucket", ctx.Stack())
	_, err := s3.NewBucket(ctx, bucketName, &s3.BucketArgs{
		Bucket:       pulumi.String(bucketName),
		Acl:          pulumi.String("public-read-write"),
		ForceDestroy: pulumi.Bool(true),
		Tags: pulumi.StringMap{
			"Project": pulumi.String("Feature Testing"),
			"Service": pulumi.String("CI"),
			"Owner":   pulumi.String("pachyderm-ci"),
			"Team":    pulumi.String("Core"),
		},
	})

	if err != nil {
		return err
	}

	return nil
}