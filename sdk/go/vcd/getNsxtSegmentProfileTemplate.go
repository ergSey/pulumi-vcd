// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/ergSey/pulumi-vcd/sdk/go/vcd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a data source to read NSX-T Segment Profile Templates.
//
// Supported in provider *v3.11+* and VCD 10.4.0+ with NSX-T. Requires System Administrator privileges.
//
// ## Example Usage
//
// ### Complete Example With All Segment Profiles)
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
//			_, err := vcd.LookupNsxtSegmentProfileTemplate(ctx, &vcd.LookupNsxtSegmentProfileTemplateArgs{
//				Name: "my-segment-profile-template-name",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupNsxtSegmentProfileTemplate(ctx *pulumi.Context, args *LookupNsxtSegmentProfileTemplateArgs, opts ...pulumi.InvokeOption) (*LookupNsxtSegmentProfileTemplateResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNsxtSegmentProfileTemplateResult
	err := ctx.Invoke("vcd:index/getNsxtSegmentProfileTemplate:getNsxtSegmentProfileTemplate", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNsxtSegmentProfileTemplate.
type LookupNsxtSegmentProfileTemplateArgs struct {
	// Name of existing Segment Profile Template
	Name string `pulumi:"name"`
}

// A collection of values returned by getNsxtSegmentProfileTemplate.
type LookupNsxtSegmentProfileTemplateResult struct {
	Description string `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id                       string `pulumi:"id"`
	IpDiscoveryProfileId     string `pulumi:"ipDiscoveryProfileId"`
	MacDiscoveryProfileId    string `pulumi:"macDiscoveryProfileId"`
	Name                     string `pulumi:"name"`
	NsxtManagerId            string `pulumi:"nsxtManagerId"`
	QosProfileId             string `pulumi:"qosProfileId"`
	SegmentSecurityProfileId string `pulumi:"segmentSecurityProfileId"`
	SpoofGuardProfileId      string `pulumi:"spoofGuardProfileId"`
}

func LookupNsxtSegmentProfileTemplateOutput(ctx *pulumi.Context, args LookupNsxtSegmentProfileTemplateOutputArgs, opts ...pulumi.InvokeOption) LookupNsxtSegmentProfileTemplateResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupNsxtSegmentProfileTemplateResultOutput, error) {
			args := v.(LookupNsxtSegmentProfileTemplateArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vcd:index/getNsxtSegmentProfileTemplate:getNsxtSegmentProfileTemplate", args, LookupNsxtSegmentProfileTemplateResultOutput{}, options).(LookupNsxtSegmentProfileTemplateResultOutput), nil
		}).(LookupNsxtSegmentProfileTemplateResultOutput)
}

// A collection of arguments for invoking getNsxtSegmentProfileTemplate.
type LookupNsxtSegmentProfileTemplateOutputArgs struct {
	// Name of existing Segment Profile Template
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupNsxtSegmentProfileTemplateOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxtSegmentProfileTemplateArgs)(nil)).Elem()
}

// A collection of values returned by getNsxtSegmentProfileTemplate.
type LookupNsxtSegmentProfileTemplateResultOutput struct{ *pulumi.OutputState }

func (LookupNsxtSegmentProfileTemplateResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxtSegmentProfileTemplateResult)(nil)).Elem()
}

func (o LookupNsxtSegmentProfileTemplateResultOutput) ToLookupNsxtSegmentProfileTemplateResultOutput() LookupNsxtSegmentProfileTemplateResultOutput {
	return o
}

func (o LookupNsxtSegmentProfileTemplateResultOutput) ToLookupNsxtSegmentProfileTemplateResultOutputWithContext(ctx context.Context) LookupNsxtSegmentProfileTemplateResultOutput {
	return o
}

func (o LookupNsxtSegmentProfileTemplateResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtSegmentProfileTemplateResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNsxtSegmentProfileTemplateResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtSegmentProfileTemplateResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNsxtSegmentProfileTemplateResultOutput) IpDiscoveryProfileId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtSegmentProfileTemplateResult) string { return v.IpDiscoveryProfileId }).(pulumi.StringOutput)
}

func (o LookupNsxtSegmentProfileTemplateResultOutput) MacDiscoveryProfileId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtSegmentProfileTemplateResult) string { return v.MacDiscoveryProfileId }).(pulumi.StringOutput)
}

func (o LookupNsxtSegmentProfileTemplateResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtSegmentProfileTemplateResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNsxtSegmentProfileTemplateResultOutput) NsxtManagerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtSegmentProfileTemplateResult) string { return v.NsxtManagerId }).(pulumi.StringOutput)
}

func (o LookupNsxtSegmentProfileTemplateResultOutput) QosProfileId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtSegmentProfileTemplateResult) string { return v.QosProfileId }).(pulumi.StringOutput)
}

func (o LookupNsxtSegmentProfileTemplateResultOutput) SegmentSecurityProfileId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtSegmentProfileTemplateResult) string { return v.SegmentSecurityProfileId }).(pulumi.StringOutput)
}

func (o LookupNsxtSegmentProfileTemplateResultOutput) SpoofGuardProfileId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtSegmentProfileTemplateResult) string { return v.SpoofGuardProfileId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNsxtSegmentProfileTemplateResultOutput{})
}
