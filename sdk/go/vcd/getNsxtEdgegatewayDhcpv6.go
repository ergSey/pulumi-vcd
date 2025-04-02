// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/ergSey/pulumi-vcd/sdk/go/vcd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a data source to read DHCPv6 configuration for NSX-T Edge Gateways.
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
//			_, err := vcd.LookupNsxtEdgegatewayDhcpv6(ctx, &vcd.LookupNsxtEdgegatewayDhcpv6Args{
//				Org:           pulumi.StringRef("datacloud"),
//				EdgeGatewayId: testing_in_vdcVcdNsxtEdgegateway.Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupNsxtEdgegatewayDhcpv6(ctx *pulumi.Context, args *LookupNsxtEdgegatewayDhcpv6Args, opts ...pulumi.InvokeOption) (*LookupNsxtEdgegatewayDhcpv6Result, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNsxtEdgegatewayDhcpv6Result
	err := ctx.Invoke("vcd:index/getNsxtEdgegatewayDhcpv6:getNsxtEdgegatewayDhcpv6", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNsxtEdgegatewayDhcpv6.
type LookupNsxtEdgegatewayDhcpv6Args struct {
	// An ID of NSX-T Edge Gateway. Can be looked up using
	// [NsxtEdgegateway](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_edgegateway) data source
	EdgeGatewayId string `pulumi:"edgeGatewayId"`
	// The name of organization to which the edge gateway belongs. Optional if defined at provider level.
	Org *string `pulumi:"org"`
}

// A collection of values returned by getNsxtEdgegatewayDhcpv6.
type LookupNsxtEdgegatewayDhcpv6Result struct {
	DnsServers    []string `pulumi:"dnsServers"`
	DomainNames   []string `pulumi:"domainNames"`
	EdgeGatewayId string   `pulumi:"edgeGatewayId"`
	// The provider-assigned unique ID for this managed resource.
	Id   string  `pulumi:"id"`
	Mode string  `pulumi:"mode"`
	Org  *string `pulumi:"org"`
}

func LookupNsxtEdgegatewayDhcpv6Output(ctx *pulumi.Context, args LookupNsxtEdgegatewayDhcpv6OutputArgs, opts ...pulumi.InvokeOption) LookupNsxtEdgegatewayDhcpv6ResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupNsxtEdgegatewayDhcpv6ResultOutput, error) {
			args := v.(LookupNsxtEdgegatewayDhcpv6Args)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vcd:index/getNsxtEdgegatewayDhcpv6:getNsxtEdgegatewayDhcpv6", args, LookupNsxtEdgegatewayDhcpv6ResultOutput{}, options).(LookupNsxtEdgegatewayDhcpv6ResultOutput), nil
		}).(LookupNsxtEdgegatewayDhcpv6ResultOutput)
}

// A collection of arguments for invoking getNsxtEdgegatewayDhcpv6.
type LookupNsxtEdgegatewayDhcpv6OutputArgs struct {
	// An ID of NSX-T Edge Gateway. Can be looked up using
	// [NsxtEdgegateway](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_edgegateway) data source
	EdgeGatewayId pulumi.StringInput `pulumi:"edgeGatewayId"`
	// The name of organization to which the edge gateway belongs. Optional if defined at provider level.
	Org pulumi.StringPtrInput `pulumi:"org"`
}

func (LookupNsxtEdgegatewayDhcpv6OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxtEdgegatewayDhcpv6Args)(nil)).Elem()
}

// A collection of values returned by getNsxtEdgegatewayDhcpv6.
type LookupNsxtEdgegatewayDhcpv6ResultOutput struct{ *pulumi.OutputState }

func (LookupNsxtEdgegatewayDhcpv6ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxtEdgegatewayDhcpv6Result)(nil)).Elem()
}

func (o LookupNsxtEdgegatewayDhcpv6ResultOutput) ToLookupNsxtEdgegatewayDhcpv6ResultOutput() LookupNsxtEdgegatewayDhcpv6ResultOutput {
	return o
}

func (o LookupNsxtEdgegatewayDhcpv6ResultOutput) ToLookupNsxtEdgegatewayDhcpv6ResultOutputWithContext(ctx context.Context) LookupNsxtEdgegatewayDhcpv6ResultOutput {
	return o
}

func (o LookupNsxtEdgegatewayDhcpv6ResultOutput) DnsServers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNsxtEdgegatewayDhcpv6Result) []string { return v.DnsServers }).(pulumi.StringArrayOutput)
}

func (o LookupNsxtEdgegatewayDhcpv6ResultOutput) DomainNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNsxtEdgegatewayDhcpv6Result) []string { return v.DomainNames }).(pulumi.StringArrayOutput)
}

func (o LookupNsxtEdgegatewayDhcpv6ResultOutput) EdgeGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtEdgegatewayDhcpv6Result) string { return v.EdgeGatewayId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNsxtEdgegatewayDhcpv6ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtEdgegatewayDhcpv6Result) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNsxtEdgegatewayDhcpv6ResultOutput) Mode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtEdgegatewayDhcpv6Result) string { return v.Mode }).(pulumi.StringOutput)
}

func (o LookupNsxtEdgegatewayDhcpv6ResultOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNsxtEdgegatewayDhcpv6Result) *string { return v.Org }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNsxtEdgegatewayDhcpv6ResultOutput{})
}
