// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/ergSey/pulumi-vcd/sdk/go/vcd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Supported in provider *v3.5+* and VCD 10.2+ with NSX-T and ALB.
//
// Provides a data source to read ALB General Settings for particular NSX-T Edge Gateway.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ergSey/pulumi-vcd/sdk/go/vcd"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			existing, err := vcd.LookupNsxtEdgegateway(ctx, &vcd.LookupNsxtEdgegatewayArgs{
//				Org:  pulumi.StringRef("my-org"),
//				Vdc:  pulumi.StringRef("nsxt-vdc"),
//				Name: "nsxt-gw",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vcd.LookupNsxtAlbSettings(ctx, &vcd.LookupNsxtAlbSettingsArgs{
//				Org:           pulumi.StringRef("my-org"),
//				EdgeGatewayId: existing.Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupNsxtAlbSettings(ctx *pulumi.Context, args *LookupNsxtAlbSettingsArgs, opts ...pulumi.InvokeOption) (*LookupNsxtAlbSettingsResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNsxtAlbSettingsResult
	err := ctx.Invoke("vcd:index/getNsxtAlbSettings:getNsxtAlbSettings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNsxtAlbSettings.
type LookupNsxtAlbSettingsArgs struct {
	// An ID of NSX-T Edge Gateway. Can be looked up using
	// [NsxtEdgegateway](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_edgegateway) data source
	EdgeGatewayId string `pulumi:"edgeGatewayId"`
	// The name of organization to which the edge gateway belongs. Optional if defined at provider level.
	Org                         *string `pulumi:"org"`
	ServiceNetworkSpecification *string `pulumi:"serviceNetworkSpecification"`
	// Deprecated: Edge Gateway will be looked up based on 'edge_gateway_id' field
	Vdc *string `pulumi:"vdc"`
}

// A collection of values returned by getNsxtAlbSettings.
type LookupNsxtAlbSettingsResult struct {
	EdgeGatewayId string `pulumi:"edgeGatewayId"`
	// The provider-assigned unique ID for this managed resource.
	Id                              string  `pulumi:"id"`
	Ipv6ServiceNetworkSpecification string  `pulumi:"ipv6ServiceNetworkSpecification"`
	IsActive                        bool    `pulumi:"isActive"`
	IsTransparentModeEnabled        bool    `pulumi:"isTransparentModeEnabled"`
	Org                             *string `pulumi:"org"`
	ServiceNetworkSpecification     string  `pulumi:"serviceNetworkSpecification"`
	SupportedFeatureSet             string  `pulumi:"supportedFeatureSet"`
	// Deprecated: Edge Gateway will be looked up based on 'edge_gateway_id' field
	Vdc *string `pulumi:"vdc"`
}

func LookupNsxtAlbSettingsOutput(ctx *pulumi.Context, args LookupNsxtAlbSettingsOutputArgs, opts ...pulumi.InvokeOption) LookupNsxtAlbSettingsResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupNsxtAlbSettingsResultOutput, error) {
			args := v.(LookupNsxtAlbSettingsArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vcd:index/getNsxtAlbSettings:getNsxtAlbSettings", args, LookupNsxtAlbSettingsResultOutput{}, options).(LookupNsxtAlbSettingsResultOutput), nil
		}).(LookupNsxtAlbSettingsResultOutput)
}

// A collection of arguments for invoking getNsxtAlbSettings.
type LookupNsxtAlbSettingsOutputArgs struct {
	// An ID of NSX-T Edge Gateway. Can be looked up using
	// [NsxtEdgegateway](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_edgegateway) data source
	EdgeGatewayId pulumi.StringInput `pulumi:"edgeGatewayId"`
	// The name of organization to which the edge gateway belongs. Optional if defined at provider level.
	Org                         pulumi.StringPtrInput `pulumi:"org"`
	ServiceNetworkSpecification pulumi.StringPtrInput `pulumi:"serviceNetworkSpecification"`
	// Deprecated: Edge Gateway will be looked up based on 'edge_gateway_id' field
	Vdc pulumi.StringPtrInput `pulumi:"vdc"`
}

func (LookupNsxtAlbSettingsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxtAlbSettingsArgs)(nil)).Elem()
}

// A collection of values returned by getNsxtAlbSettings.
type LookupNsxtAlbSettingsResultOutput struct{ *pulumi.OutputState }

func (LookupNsxtAlbSettingsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxtAlbSettingsResult)(nil)).Elem()
}

func (o LookupNsxtAlbSettingsResultOutput) ToLookupNsxtAlbSettingsResultOutput() LookupNsxtAlbSettingsResultOutput {
	return o
}

func (o LookupNsxtAlbSettingsResultOutput) ToLookupNsxtAlbSettingsResultOutputWithContext(ctx context.Context) LookupNsxtAlbSettingsResultOutput {
	return o
}

func (o LookupNsxtAlbSettingsResultOutput) EdgeGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAlbSettingsResult) string { return v.EdgeGatewayId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNsxtAlbSettingsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAlbSettingsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNsxtAlbSettingsResultOutput) Ipv6ServiceNetworkSpecification() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAlbSettingsResult) string { return v.Ipv6ServiceNetworkSpecification }).(pulumi.StringOutput)
}

func (o LookupNsxtAlbSettingsResultOutput) IsActive() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNsxtAlbSettingsResult) bool { return v.IsActive }).(pulumi.BoolOutput)
}

func (o LookupNsxtAlbSettingsResultOutput) IsTransparentModeEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNsxtAlbSettingsResult) bool { return v.IsTransparentModeEnabled }).(pulumi.BoolOutput)
}

func (o LookupNsxtAlbSettingsResultOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNsxtAlbSettingsResult) *string { return v.Org }).(pulumi.StringPtrOutput)
}

func (o LookupNsxtAlbSettingsResultOutput) ServiceNetworkSpecification() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAlbSettingsResult) string { return v.ServiceNetworkSpecification }).(pulumi.StringOutput)
}

func (o LookupNsxtAlbSettingsResultOutput) SupportedFeatureSet() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAlbSettingsResult) string { return v.SupportedFeatureSet }).(pulumi.StringOutput)
}

// Deprecated: Edge Gateway will be looked up based on 'edge_gateway_id' field
func (o LookupNsxtAlbSettingsResultOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNsxtAlbSettingsResult) *string { return v.Vdc }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNsxtAlbSettingsResultOutput{})
}
