// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/ergSey/pulumi-vcd/sdk/go/vcd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VMware Cloud Director NSX-T Segment Security Profile data source. This can be used to read NSX-T Segment Profile definitions.
//
// Supported in provider *v3.11+*.
//
// ## Example Usage
//
// ### Segment Security Profile)
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
//			nsxt, err := vcd.GetNsxtManager(ctx, &vcd.GetNsxtManagerArgs{
//				Name: "nsxManager1",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vcd.GetNsxtSegmentSecurityProfile(ctx, &vcd.GetNsxtSegmentSecurityProfileArgs{
//				Name:          "segment-security-profile-0",
//				NsxtManagerId: pulumi.StringRef(nsxt.Id),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetNsxtSegmentSecurityProfile(ctx *pulumi.Context, args *GetNsxtSegmentSecurityProfileArgs, opts ...pulumi.InvokeOption) (*GetNsxtSegmentSecurityProfileResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetNsxtSegmentSecurityProfileResult
	err := ctx.Invoke("vcd:index/getNsxtSegmentSecurityProfile:getNsxtSegmentSecurityProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNsxtSegmentSecurityProfile.
type GetNsxtSegmentSecurityProfileArgs struct {
	// Pre-defined list of allowed MAC addresses to be excluded from BPDU filtering.
	BpduFilterAllowLists []string `pulumi:"bpduFilterAllowLists"`
	// The name of Segment Profile
	Name string `pulumi:"name"`
	// Segment Profile search context. Use when searching by NSX-T manager
	NsxtManagerId *string `pulumi:"nsxtManagerId"`
	// Segment Profile search context. Use when searching by VDC group
	//
	// > Note: only one of `nsxtManagerId`, `vdcId`, `vdcGroupId` can be used
	VdcGroupId *string `pulumi:"vdcGroupId"`
	// Segment Profile search context. Use when searching by VDC
	VdcId *string `pulumi:"vdcId"`
}

// A collection of values returned by getNsxtSegmentSecurityProfile.
type GetNsxtSegmentSecurityProfileResult struct {
	// Pre-defined list of allowed MAC addresses to be excluded from BPDU filtering.
	BpduFilterAllowLists []string `pulumi:"bpduFilterAllowLists"`
	// Description of Segment Security Profile
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Defines whether BPDU filter is enabled.
	IsBpduFilterEnabled bool `pulumi:"isBpduFilterEnabled"`
	// Defines whether DHCP Client block IPv4 is enabled. This filters DHCP Client IPv4 traffic.
	IsDhcpV4ClientBlockEnabled bool `pulumi:"isDhcpV4ClientBlockEnabled"`
	// Defines whether DHCP Server block IPv4 is enabled. This filters DHCP Server IPv4 traffic.
	IsDhcpV4ServerBlockEnabled bool `pulumi:"isDhcpV4ServerBlockEnabled"`
	// Defines whether DHCP Client block IPv6 is enabled. This filters DHCP Client IPv6 traffic.
	IsDhcpV6ClientBlockEnabled bool `pulumi:"isDhcpV6ClientBlockEnabled"`
	// Defines whether DHCP Server block IPv6 is enabled. This filters DHCP Server IPv6 traffic.
	IsDhcpV6ServerBlockEnabled bool `pulumi:"isDhcpV6ServerBlockEnabled"`
	// Defines whether non IP traffic block is enabled. If true, it blocks all traffic except IP/(G)ARP/BPDU.
	IsNonIpTrafficBlockEnabled bool `pulumi:"isNonIpTrafficBlockEnabled"`
	// Defines whether Router Advertisement Guard is enabled. This filters DHCP Server IPv6 traffic.
	IsRaGuardEnabled bool `pulumi:"isRaGuardEnabled"`
	// Defines whether Rate Limiting is enabled.
	IsRateLimittingEnabled bool    `pulumi:"isRateLimittingEnabled"`
	Name                   string  `pulumi:"name"`
	NsxtManagerId          *string `pulumi:"nsxtManagerId"`
	// Incoming broadcast traffic limit in packets per second.
	RxBroadcastLimit int `pulumi:"rxBroadcastLimit"`
	// Incoming multicast traffic limit in packets per second.
	RxMulticastLimit int `pulumi:"rxMulticastLimit"`
	// Outgoing broadcast traffic limit in packets per second.
	TxBroadcastLimit int `pulumi:"txBroadcastLimit"`
	// Outgoing multicast traffic limit in packets per second.
	TxMulticastLimit int     `pulumi:"txMulticastLimit"`
	VdcGroupId       *string `pulumi:"vdcGroupId"`
	VdcId            *string `pulumi:"vdcId"`
}

func GetNsxtSegmentSecurityProfileOutput(ctx *pulumi.Context, args GetNsxtSegmentSecurityProfileOutputArgs, opts ...pulumi.InvokeOption) GetNsxtSegmentSecurityProfileResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetNsxtSegmentSecurityProfileResultOutput, error) {
			args := v.(GetNsxtSegmentSecurityProfileArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vcd:index/getNsxtSegmentSecurityProfile:getNsxtSegmentSecurityProfile", args, GetNsxtSegmentSecurityProfileResultOutput{}, options).(GetNsxtSegmentSecurityProfileResultOutput), nil
		}).(GetNsxtSegmentSecurityProfileResultOutput)
}

// A collection of arguments for invoking getNsxtSegmentSecurityProfile.
type GetNsxtSegmentSecurityProfileOutputArgs struct {
	// Pre-defined list of allowed MAC addresses to be excluded from BPDU filtering.
	BpduFilterAllowLists pulumi.StringArrayInput `pulumi:"bpduFilterAllowLists"`
	// The name of Segment Profile
	Name pulumi.StringInput `pulumi:"name"`
	// Segment Profile search context. Use when searching by NSX-T manager
	NsxtManagerId pulumi.StringPtrInput `pulumi:"nsxtManagerId"`
	// Segment Profile search context. Use when searching by VDC group
	//
	// > Note: only one of `nsxtManagerId`, `vdcId`, `vdcGroupId` can be used
	VdcGroupId pulumi.StringPtrInput `pulumi:"vdcGroupId"`
	// Segment Profile search context. Use when searching by VDC
	VdcId pulumi.StringPtrInput `pulumi:"vdcId"`
}

func (GetNsxtSegmentSecurityProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNsxtSegmentSecurityProfileArgs)(nil)).Elem()
}

// A collection of values returned by getNsxtSegmentSecurityProfile.
type GetNsxtSegmentSecurityProfileResultOutput struct{ *pulumi.OutputState }

func (GetNsxtSegmentSecurityProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNsxtSegmentSecurityProfileResult)(nil)).Elem()
}

func (o GetNsxtSegmentSecurityProfileResultOutput) ToGetNsxtSegmentSecurityProfileResultOutput() GetNsxtSegmentSecurityProfileResultOutput {
	return o
}

func (o GetNsxtSegmentSecurityProfileResultOutput) ToGetNsxtSegmentSecurityProfileResultOutputWithContext(ctx context.Context) GetNsxtSegmentSecurityProfileResultOutput {
	return o
}

// Pre-defined list of allowed MAC addresses to be excluded from BPDU filtering.
func (o GetNsxtSegmentSecurityProfileResultOutput) BpduFilterAllowLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetNsxtSegmentSecurityProfileResult) []string { return v.BpduFilterAllowLists }).(pulumi.StringArrayOutput)
}

// Description of Segment Security Profile
func (o GetNsxtSegmentSecurityProfileResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetNsxtSegmentSecurityProfileResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetNsxtSegmentSecurityProfileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetNsxtSegmentSecurityProfileResult) string { return v.Id }).(pulumi.StringOutput)
}

// Defines whether BPDU filter is enabled.
func (o GetNsxtSegmentSecurityProfileResultOutput) IsBpduFilterEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetNsxtSegmentSecurityProfileResult) bool { return v.IsBpduFilterEnabled }).(pulumi.BoolOutput)
}

// Defines whether DHCP Client block IPv4 is enabled. This filters DHCP Client IPv4 traffic.
func (o GetNsxtSegmentSecurityProfileResultOutput) IsDhcpV4ClientBlockEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetNsxtSegmentSecurityProfileResult) bool { return v.IsDhcpV4ClientBlockEnabled }).(pulumi.BoolOutput)
}

// Defines whether DHCP Server block IPv4 is enabled. This filters DHCP Server IPv4 traffic.
func (o GetNsxtSegmentSecurityProfileResultOutput) IsDhcpV4ServerBlockEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetNsxtSegmentSecurityProfileResult) bool { return v.IsDhcpV4ServerBlockEnabled }).(pulumi.BoolOutput)
}

// Defines whether DHCP Client block IPv6 is enabled. This filters DHCP Client IPv6 traffic.
func (o GetNsxtSegmentSecurityProfileResultOutput) IsDhcpV6ClientBlockEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetNsxtSegmentSecurityProfileResult) bool { return v.IsDhcpV6ClientBlockEnabled }).(pulumi.BoolOutput)
}

// Defines whether DHCP Server block IPv6 is enabled. This filters DHCP Server IPv6 traffic.
func (o GetNsxtSegmentSecurityProfileResultOutput) IsDhcpV6ServerBlockEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetNsxtSegmentSecurityProfileResult) bool { return v.IsDhcpV6ServerBlockEnabled }).(pulumi.BoolOutput)
}

// Defines whether non IP traffic block is enabled. If true, it blocks all traffic except IP/(G)ARP/BPDU.
func (o GetNsxtSegmentSecurityProfileResultOutput) IsNonIpTrafficBlockEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetNsxtSegmentSecurityProfileResult) bool { return v.IsNonIpTrafficBlockEnabled }).(pulumi.BoolOutput)
}

// Defines whether Router Advertisement Guard is enabled. This filters DHCP Server IPv6 traffic.
func (o GetNsxtSegmentSecurityProfileResultOutput) IsRaGuardEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetNsxtSegmentSecurityProfileResult) bool { return v.IsRaGuardEnabled }).(pulumi.BoolOutput)
}

// Defines whether Rate Limiting is enabled.
func (o GetNsxtSegmentSecurityProfileResultOutput) IsRateLimittingEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GetNsxtSegmentSecurityProfileResult) bool { return v.IsRateLimittingEnabled }).(pulumi.BoolOutput)
}

func (o GetNsxtSegmentSecurityProfileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetNsxtSegmentSecurityProfileResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetNsxtSegmentSecurityProfileResultOutput) NsxtManagerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNsxtSegmentSecurityProfileResult) *string { return v.NsxtManagerId }).(pulumi.StringPtrOutput)
}

// Incoming broadcast traffic limit in packets per second.
func (o GetNsxtSegmentSecurityProfileResultOutput) RxBroadcastLimit() pulumi.IntOutput {
	return o.ApplyT(func(v GetNsxtSegmentSecurityProfileResult) int { return v.RxBroadcastLimit }).(pulumi.IntOutput)
}

// Incoming multicast traffic limit in packets per second.
func (o GetNsxtSegmentSecurityProfileResultOutput) RxMulticastLimit() pulumi.IntOutput {
	return o.ApplyT(func(v GetNsxtSegmentSecurityProfileResult) int { return v.RxMulticastLimit }).(pulumi.IntOutput)
}

// Outgoing broadcast traffic limit in packets per second.
func (o GetNsxtSegmentSecurityProfileResultOutput) TxBroadcastLimit() pulumi.IntOutput {
	return o.ApplyT(func(v GetNsxtSegmentSecurityProfileResult) int { return v.TxBroadcastLimit }).(pulumi.IntOutput)
}

// Outgoing multicast traffic limit in packets per second.
func (o GetNsxtSegmentSecurityProfileResultOutput) TxMulticastLimit() pulumi.IntOutput {
	return o.ApplyT(func(v GetNsxtSegmentSecurityProfileResult) int { return v.TxMulticastLimit }).(pulumi.IntOutput)
}

func (o GetNsxtSegmentSecurityProfileResultOutput) VdcGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNsxtSegmentSecurityProfileResult) *string { return v.VdcGroupId }).(pulumi.StringPtrOutput)
}

func (o GetNsxtSegmentSecurityProfileResultOutput) VdcId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNsxtSegmentSecurityProfileResult) *string { return v.VdcId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetNsxtSegmentSecurityProfileResultOutput{})
}
