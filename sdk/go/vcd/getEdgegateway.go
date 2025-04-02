// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/ergSey/pulumi-vcd/sdk/go/vcd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VMware Cloud Director NSX-V edge gateway data source, directly connected to one or more external networks. This can be used to reference
// edge gateways for Org VDC networks to connect.
//
// Supported in provider *v2.5+*
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
//			mygw, err := vcd.LookupEdgegateway(ctx, &vcd.LookupEdgegatewayArgs{
//				Name: pulumi.StringRef("mygw"),
//				Org:  pulumi.StringRef("myorg"),
//				Vdc:  pulumi.StringRef("myvdc"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("edgeGatewayId", mygw.Id)
//			// Get the name of the default gateway from the data source
//			// and use it to establish a second data source
//			externalNetwork1, err := vcd.LookupExternalNetwork(ctx, &vcd.LookupExternalNetworkArgs{
//				Name: mygw.ExternalNetworks.Name,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("gateway", externalNetwork1.IpScopes[0].Gateway)
//			ctx.Export("netmask", externalNetwork1.IpScopes[0].Netmask)
//			ctx.Export("DNS", externalNetwork1.IpScopes[0].Dns1)
//			ctx.Export("externalIp", externalNetwork1.IpScopes[0].StaticIpPools[0].StartAddress)
//			return nil
//		})
//	}
//
// ```
//
// ## Filter arguments
//
// (Supported in provider *v2.9+*)
//
// * `nameRegex` - (Optional) matches the name using a regular expression.
//
// See [Filters reference](https://www.terraform.io/providers/vmware/vcd/latest/docs/guides/data_source_filters) for details and examples.
func LookupEdgegateway(ctx *pulumi.Context, args *LookupEdgegatewayArgs, opts ...pulumi.InvokeOption) (*LookupEdgegatewayResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupEdgegatewayResult
	err := ctx.Invoke("vcd:index/getEdgegateway:getEdgegateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getEdgegateway.
type LookupEdgegatewayArgs struct {
	// Retrieves the data source using one or more filter parameters
	Filter *GetEdgegatewayFilter `pulumi:"filter"`
	// A unique name for the edge gateway (optional when `filter` is used)
	Name *string `pulumi:"name"`
	// The name of organization to which the VDC belongs. Optional if defined at provider level.
	Org *string `pulumi:"org"`
	// The name of VDC that owns the edge gateway. Optional if defined at provider level.
	Vdc *string `pulumi:"vdc"`
}

// A collection of values returned by getEdgegateway.
type LookupEdgegatewayResult struct {
	Configuration               string                          `pulumi:"configuration"`
	DefaultExternalNetworkIp    string                          `pulumi:"defaultExternalNetworkIp"`
	Description                 string                          `pulumi:"description"`
	DistributedRouting          bool                            `pulumi:"distributedRouting"`
	ExternalNetworkIps          []string                        `pulumi:"externalNetworkIps"`
	ExternalNetworks            []GetEdgegatewayExternalNetwork `pulumi:"externalNetworks"`
	Filter                      *GetEdgegatewayFilter           `pulumi:"filter"`
	FipsModeEnabled             bool                            `pulumi:"fipsModeEnabled"`
	FwDefaultRuleAction         string                          `pulumi:"fwDefaultRuleAction"`
	FwDefaultRuleLoggingEnabled bool                            `pulumi:"fwDefaultRuleLoggingEnabled"`
	FwEnabled                   bool                            `pulumi:"fwEnabled"`
	HaEnabled                   bool                            `pulumi:"haEnabled"`
	// The provider-assigned unique ID for this managed resource.
	Id                         string  `pulumi:"id"`
	LbAccelerationEnabled      bool    `pulumi:"lbAccelerationEnabled"`
	LbEnabled                  bool    `pulumi:"lbEnabled"`
	LbLoggingEnabled           bool    `pulumi:"lbLoggingEnabled"`
	LbLoglevel                 string  `pulumi:"lbLoglevel"`
	Name                       *string `pulumi:"name"`
	Org                        *string `pulumi:"org"`
	UseDefaultRouteForDnsRelay bool    `pulumi:"useDefaultRouteForDnsRelay"`
	Vdc                        *string `pulumi:"vdc"`
}

func LookupEdgegatewayOutput(ctx *pulumi.Context, args LookupEdgegatewayOutputArgs, opts ...pulumi.InvokeOption) LookupEdgegatewayResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupEdgegatewayResultOutput, error) {
			args := v.(LookupEdgegatewayArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vcd:index/getEdgegateway:getEdgegateway", args, LookupEdgegatewayResultOutput{}, options).(LookupEdgegatewayResultOutput), nil
		}).(LookupEdgegatewayResultOutput)
}

// A collection of arguments for invoking getEdgegateway.
type LookupEdgegatewayOutputArgs struct {
	// Retrieves the data source using one or more filter parameters
	Filter GetEdgegatewayFilterPtrInput `pulumi:"filter"`
	// A unique name for the edge gateway (optional when `filter` is used)
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The name of organization to which the VDC belongs. Optional if defined at provider level.
	Org pulumi.StringPtrInput `pulumi:"org"`
	// The name of VDC that owns the edge gateway. Optional if defined at provider level.
	Vdc pulumi.StringPtrInput `pulumi:"vdc"`
}

func (LookupEdgegatewayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEdgegatewayArgs)(nil)).Elem()
}

// A collection of values returned by getEdgegateway.
type LookupEdgegatewayResultOutput struct{ *pulumi.OutputState }

func (LookupEdgegatewayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupEdgegatewayResult)(nil)).Elem()
}

func (o LookupEdgegatewayResultOutput) ToLookupEdgegatewayResultOutput() LookupEdgegatewayResultOutput {
	return o
}

func (o LookupEdgegatewayResultOutput) ToLookupEdgegatewayResultOutputWithContext(ctx context.Context) LookupEdgegatewayResultOutput {
	return o
}

func (o LookupEdgegatewayResultOutput) Configuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEdgegatewayResult) string { return v.Configuration }).(pulumi.StringOutput)
}

func (o LookupEdgegatewayResultOutput) DefaultExternalNetworkIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEdgegatewayResult) string { return v.DefaultExternalNetworkIp }).(pulumi.StringOutput)
}

func (o LookupEdgegatewayResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEdgegatewayResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupEdgegatewayResultOutput) DistributedRouting() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupEdgegatewayResult) bool { return v.DistributedRouting }).(pulumi.BoolOutput)
}

func (o LookupEdgegatewayResultOutput) ExternalNetworkIps() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupEdgegatewayResult) []string { return v.ExternalNetworkIps }).(pulumi.StringArrayOutput)
}

func (o LookupEdgegatewayResultOutput) ExternalNetworks() GetEdgegatewayExternalNetworkArrayOutput {
	return o.ApplyT(func(v LookupEdgegatewayResult) []GetEdgegatewayExternalNetwork { return v.ExternalNetworks }).(GetEdgegatewayExternalNetworkArrayOutput)
}

func (o LookupEdgegatewayResultOutput) Filter() GetEdgegatewayFilterPtrOutput {
	return o.ApplyT(func(v LookupEdgegatewayResult) *GetEdgegatewayFilter { return v.Filter }).(GetEdgegatewayFilterPtrOutput)
}

func (o LookupEdgegatewayResultOutput) FipsModeEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupEdgegatewayResult) bool { return v.FipsModeEnabled }).(pulumi.BoolOutput)
}

func (o LookupEdgegatewayResultOutput) FwDefaultRuleAction() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEdgegatewayResult) string { return v.FwDefaultRuleAction }).(pulumi.StringOutput)
}

func (o LookupEdgegatewayResultOutput) FwDefaultRuleLoggingEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupEdgegatewayResult) bool { return v.FwDefaultRuleLoggingEnabled }).(pulumi.BoolOutput)
}

func (o LookupEdgegatewayResultOutput) FwEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupEdgegatewayResult) bool { return v.FwEnabled }).(pulumi.BoolOutput)
}

func (o LookupEdgegatewayResultOutput) HaEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupEdgegatewayResult) bool { return v.HaEnabled }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupEdgegatewayResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEdgegatewayResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupEdgegatewayResultOutput) LbAccelerationEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupEdgegatewayResult) bool { return v.LbAccelerationEnabled }).(pulumi.BoolOutput)
}

func (o LookupEdgegatewayResultOutput) LbEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupEdgegatewayResult) bool { return v.LbEnabled }).(pulumi.BoolOutput)
}

func (o LookupEdgegatewayResultOutput) LbLoggingEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupEdgegatewayResult) bool { return v.LbLoggingEnabled }).(pulumi.BoolOutput)
}

func (o LookupEdgegatewayResultOutput) LbLoglevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupEdgegatewayResult) string { return v.LbLoglevel }).(pulumi.StringOutput)
}

func (o LookupEdgegatewayResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEdgegatewayResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupEdgegatewayResultOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEdgegatewayResult) *string { return v.Org }).(pulumi.StringPtrOutput)
}

func (o LookupEdgegatewayResultOutput) UseDefaultRouteForDnsRelay() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupEdgegatewayResult) bool { return v.UseDefaultRouteForDnsRelay }).(pulumi.BoolOutput)
}

func (o LookupEdgegatewayResultOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupEdgegatewayResult) *string { return v.Vdc }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupEdgegatewayResultOutput{})
}
