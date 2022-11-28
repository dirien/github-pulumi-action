package main

import (
	"context"
	"fmt"
	"github.com/pulumi/pulumi-civo/sdk/v2/go/civo"
	"github.com/pulumi/pulumi/sdk/v3/go/auto"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"os"
)

func main() {
	/*pulumi.Run(func(ctx *pulumi.Context) error {

		instance, err := civo.NewInstance(ctx, "test", &civo.InstanceArgs{
			Hostname:         pulumi.String("test"),
			Region:           pulumi.String("FRA1"),
			Size:             pulumi.String("g3.large"),
			PublicIpRequired: pulumi.String("create"),
			InitialUser:      pulumi.String("root"),
			DiskImage:        pulumi.String("19c4c893-b452-4bcc-a3f2-42f8204a36ac"),
		})
		if err != nil {
			return err
		}

		ctx.Export("ip", instance.PublicIp)

		return nil
	})
	*/

	deployFunc := func(ctx *pulumi.Context) error {
		instance, err := civo.NewInstance(ctx, "test", &civo.InstanceArgs{
			Hostname:         pulumi.String("test"),
			Region:           pulumi.String("FRA1"),
			Size:             pulumi.String("g3.large"),
			PublicIpRequired: pulumi.String("create"),
			InitialUser:      pulumi.String("root"),
			DiskImage:        pulumi.String("19c4c893-b452-4bcc-a3f2-42f8204a36ac"),
		})
		if err != nil {
			return err
		}
		ctx.Export("ip", instance.PublicIp)

		return nil
	}
	ctx := context.Background()

	stackInlineSource, err := auto.UpsertStackInlineSource(ctx, "dev", "github-pulumi-action", deployFunc)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Installing the CIVO plugin")

	err = stackInlineSource.Workspace().InstallPlugin(ctx, "civo", "v2.3.0")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Successfully installed the CIVO plugin")

	stackInlineSource.SetConfig(ctx, "civo:token", auto.ConfigValue{Value: os.Getenv("CIVO_TOKEN")})

	preview, err := stackInlineSource.Preview(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(preview.StdOut)
}
