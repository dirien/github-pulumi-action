package main

import (
	"github.com/pulumi/pulumi-civo/sdk/v2/go/civo"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {

		instance, err := civo.NewInstance(ctx, "test2334", &civo.InstanceArgs{
			Hostname:         pulumi.String("test2334"),
			Region:           pulumi.String("FRA1"),
			Size:             pulumi.String("g3.large"),
			PublicIpRequired: pulumi.String("create"),
			InitialUser:      pulumi.String("root"),
			DiskImage:        pulumi.String("19c4c893-b452-4bcc-a3f2-42f8204a36ac"),
		})
		if err != nil {
			return err
		}

		ctx.Export("ip!", instance.PublicIp)

		return nil
	})
}
