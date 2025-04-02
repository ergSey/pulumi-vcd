// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/ergSey/pulumi-vcd/sdk/go/vcd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a data source to read Network Profile for NSX-T VDCs.
//
// Supported in provider *v3.11+* and VCD 10.4.0+ with NSX-T.
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
//			_, err := vcd.LookupOrgVdcNsxtNetworkProfile(ctx, &vcd.LookupOrgVdcNsxtNetworkProfileArgs{
//				Org: pulumi.StringRef("my-org"),
//				Vdc: pulumi.StringRef("my-vdc"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupOrgVdcNsxtNetworkProfile(ctx *pulumi.Context, args *LookupOrgVdcNsxtNetworkProfileArgs, opts ...pulumi.InvokeOption) (*LookupOrgVdcNsxtNetworkProfileResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupOrgVdcNsxtNetworkProfileResult
	err := ctx.Invoke("vcd:index/getOrgVdcNsxtNetworkProfile:getOrgVdcNsxtNetworkProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrgVdcNsxtNetworkProfile.
type LookupOrgVdcNsxtNetworkProfileArgs struct {
	// The name of organization to use, optional if defined at provider level
	Org *string `pulumi:"org"`
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `pulumi:"vdc"`
}

// A collection of values returned by getOrgVdcNsxtNetworkProfile.
type LookupOrgVdcNsxtNetworkProfileResult struct {
	// An ID of NSX-T Edge Cluster which should provide vApp
	// Networking Services or DHCP for Isolated Networks. This value **might be unavailable** in data
	// source if user has insufficient rights.
	EdgeClusterId string `pulumi:"edgeClusterId"`
	// The provider-assigned unique ID for this managed resource.
	Id  string  `pulumi:"id"`
	Org *string `pulumi:"org"`
	// Default Segment Profile ID for all vApp
	// networks withing this VDC
	VappNetworksDefaultSegmentProfileTemplateId string  `pulumi:"vappNetworksDefaultSegmentProfileTemplateId"`
	Vdc                                         *string `pulumi:"vdc"`
	// Default Segment Profile ID for all Org VDC
	// networks withing this VDC
	VdcNetworksDefaultSegmentProfileTemplateId string `pulumi:"vdcNetworksDefaultSegmentProfileTemplateId"`
}

func LookupOrgVdcNsxtNetworkProfileOutput(ctx *pulumi.Context, args LookupOrgVdcNsxtNetworkProfileOutputArgs, opts ...pulumi.InvokeOption) LookupOrgVdcNsxtNetworkProfileResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupOrgVdcNsxtNetworkProfileResultOutput, error) {
			args := v.(LookupOrgVdcNsxtNetworkProfileArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vcd:index/getOrgVdcNsxtNetworkProfile:getOrgVdcNsxtNetworkProfile", args, LookupOrgVdcNsxtNetworkProfileResultOutput{}, options).(LookupOrgVdcNsxtNetworkProfileResultOutput), nil
		}).(LookupOrgVdcNsxtNetworkProfileResultOutput)
}

// A collection of arguments for invoking getOrgVdcNsxtNetworkProfile.
type LookupOrgVdcNsxtNetworkProfileOutputArgs struct {
	// The name of organization to use, optional if defined at provider level
	Org pulumi.StringPtrInput `pulumi:"org"`
	// The name of VDC to use, optional if defined at provider level
	Vdc pulumi.StringPtrInput `pulumi:"vdc"`
}

func (LookupOrgVdcNsxtNetworkProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrgVdcNsxtNetworkProfileArgs)(nil)).Elem()
}

// A collection of values returned by getOrgVdcNsxtNetworkProfile.
type LookupOrgVdcNsxtNetworkProfileResultOutput struct{ *pulumi.OutputState }

func (LookupOrgVdcNsxtNetworkProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrgVdcNsxtNetworkProfileResult)(nil)).Elem()
}

func (o LookupOrgVdcNsxtNetworkProfileResultOutput) ToLookupOrgVdcNsxtNetworkProfileResultOutput() LookupOrgVdcNsxtNetworkProfileResultOutput {
	return o
}

func (o LookupOrgVdcNsxtNetworkProfileResultOutput) ToLookupOrgVdcNsxtNetworkProfileResultOutputWithContext(ctx context.Context) LookupOrgVdcNsxtNetworkProfileResultOutput {
	return o
}

// An ID of NSX-T Edge Cluster which should provide vApp
// Networking Services or DHCP for Isolated Networks. This value **might be unavailable** in data
// source if user has insufficient rights.
func (o LookupOrgVdcNsxtNetworkProfileResultOutput) EdgeClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgVdcNsxtNetworkProfileResult) string { return v.EdgeClusterId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupOrgVdcNsxtNetworkProfileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgVdcNsxtNetworkProfileResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupOrgVdcNsxtNetworkProfileResultOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOrgVdcNsxtNetworkProfileResult) *string { return v.Org }).(pulumi.StringPtrOutput)
}

// Default Segment Profile ID for all vApp
// networks withing this VDC
func (o LookupOrgVdcNsxtNetworkProfileResultOutput) VappNetworksDefaultSegmentProfileTemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgVdcNsxtNetworkProfileResult) string {
		return v.VappNetworksDefaultSegmentProfileTemplateId
	}).(pulumi.StringOutput)
}

func (o LookupOrgVdcNsxtNetworkProfileResultOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupOrgVdcNsxtNetworkProfileResult) *string { return v.Vdc }).(pulumi.StringPtrOutput)
}

// Default Segment Profile ID for all Org VDC
// networks withing this VDC
func (o LookupOrgVdcNsxtNetworkProfileResultOutput) VdcNetworksDefaultSegmentProfileTemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgVdcNsxtNetworkProfileResult) string {
		return v.VdcNetworksDefaultSegmentProfileTemplateId
	}).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrgVdcNsxtNetworkProfileResultOutput{})
}
