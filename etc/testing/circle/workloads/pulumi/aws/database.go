package main

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/rds"
	postgresql "github.com/pulumi/pulumi-postgresql/sdk/v3/go/postgresql"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

func DeployRDS(ctx *pulumi.Context) error {
	cfg := config.New(ctx, "")
	rdsAllocatedStorage, err := cfg.TryInt("rdsAllocatedStorage")
	if err != nil {
		rdsAllocatedStorage = 20
	}
	rdsInstanceClass, err := cfg.Try("rdsInstanceClass")
	if err != nil {
		rdsInstanceClass = "db.t2.medium"
	}
	rdsDiskType, err := cfg.Try("rdsDiskType")
	if err != nil {
		rdsDiskType = "gp2"
	}
	rdsDiskIOPs, err := cfg.TryInt("rdsDiskIOPs")
	if err != nil {
		rdsDiskIOPs = 1000
	}

	rdsInstanceArgs := &rds.InstanceArgs{
		AllocatedStorage:   pulumi.Int(rdsAllocatedStorage),
		Engine:             pulumi.String("postgres"),
		InstanceClass:      pulumi.String(rdsInstanceClass),
		DbName:             pulumi.String("pachyderm"),
		Password:           cfg.RequireSecret("rdsPGDBPassword"),
		SkipFinalSnapshot:  pulumi.Bool(true),
		StorageType:        pulumi.String(rdsDiskType),
		Username:           pulumi.String("postgres"),
		PubliclyAccessible: pulumi.Bool(true),
		Tags: pulumi.StringMap{
			"Project": pulumi.String("Feature Testing"),
			"Service": pulumi.String("CI"),
			"Owner":   pulumi.String("pachyderm-ci"),
			"Team":    pulumi.String("Core"),
		},
	}

	if rdsDiskType == "io1" {
		rdsInstanceArgs.Iops = pulumi.Int(rdsDiskIOPs)
	}

	rdsInstanceName := fmt.Sprintf("rds-%s-instance", ctx.Stack())
	r, err := rds.NewInstance(ctx, rdsInstanceName, rdsInstanceArgs)

	if err != nil {
		return err
	}

	rdsProviderName := fmt.Sprintf("%s-postgresql", ctx.Stack())
	postgresProvider, err := postgresql.NewProvider(ctx, rdsProviderName, &postgresql.ProviderArgs{
		Host:     r.Address,
		Port:     r.Port,
		Username: r.Username,
		Password: r.Password,
	})

	if err != nil {
		return err
	}

	_, err = postgresql.NewDatabase(ctx, "dex", &postgresql.DatabaseArgs{}, pulumi.Provider(postgresProvider))

	if err != nil {
		return err
	}

	return nil
}
