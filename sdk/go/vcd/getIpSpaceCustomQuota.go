// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/ergSey/pulumi-vcd/sdk/go/vcd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a data source to read Custom Quotas for a given Org in a particular IP Space.
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
//			_, err := vcd.LookupIpSpaceCustomQuota(ctx, &vcd.LookupIpSpaceCustomQuotaArgs{
//				OrgId:     org1VcdOrg.Id,
//				IpSpaceId: space1.Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupIpSpaceCustomQuota(ctx *pulumi.Context, args *LookupIpSpaceCustomQuotaArgs, opts ...pulumi.InvokeOption) (*LookupIpSpaceCustomQuotaResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupIpSpaceCustomQuotaResult
	err := ctx.Invoke("vcd:index/getIpSpaceCustomQuota:getIpSpaceCustomQuota", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getIpSpaceCustomQuota.
type LookupIpSpaceCustomQuotaArgs struct {
	// IP Space ID to read Custom Quotas
	IpSpaceId string `pulumi:"ipSpaceId"`
	// Organization ID, for which the Custom Quota should be read
	OrgId string `pulumi:"orgId"`
}

// A collection of values returned by getIpSpaceCustomQuota.
type LookupIpSpaceCustomQuotaResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id             string                               `pulumi:"id"`
	IpPrefixQuotas []GetIpSpaceCustomQuotaIpPrefixQuota `pulumi:"ipPrefixQuotas"`
	IpRangeQuota   string                               `pulumi:"ipRangeQuota"`
	IpSpaceId      string                               `pulumi:"ipSpaceId"`
	OrgId          string                               `pulumi:"orgId"`
}

func LookupIpSpaceCustomQuotaOutput(ctx *pulumi.Context, args LookupIpSpaceCustomQuotaOutputArgs, opts ...pulumi.InvokeOption) LookupIpSpaceCustomQuotaResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupIpSpaceCustomQuotaResultOutput, error) {
			args := v.(LookupIpSpaceCustomQuotaArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vcd:index/getIpSpaceCustomQuota:getIpSpaceCustomQuota", args, LookupIpSpaceCustomQuotaResultOutput{}, options).(LookupIpSpaceCustomQuotaResultOutput), nil
		}).(LookupIpSpaceCustomQuotaResultOutput)
}

// A collection of arguments for invoking getIpSpaceCustomQuota.
type LookupIpSpaceCustomQuotaOutputArgs struct {
	// IP Space ID to read Custom Quotas
	IpSpaceId pulumi.StringInput `pulumi:"ipSpaceId"`
	// Organization ID, for which the Custom Quota should be read
	OrgId pulumi.StringInput `pulumi:"orgId"`
}

func (LookupIpSpaceCustomQuotaOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpSpaceCustomQuotaArgs)(nil)).Elem()
}

// A collection of values returned by getIpSpaceCustomQuota.
type LookupIpSpaceCustomQuotaResultOutput struct{ *pulumi.OutputState }

func (LookupIpSpaceCustomQuotaResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIpSpaceCustomQuotaResult)(nil)).Elem()
}

func (o LookupIpSpaceCustomQuotaResultOutput) ToLookupIpSpaceCustomQuotaResultOutput() LookupIpSpaceCustomQuotaResultOutput {
	return o
}

func (o LookupIpSpaceCustomQuotaResultOutput) ToLookupIpSpaceCustomQuotaResultOutputWithContext(ctx context.Context) LookupIpSpaceCustomQuotaResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupIpSpaceCustomQuotaResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpSpaceCustomQuotaResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIpSpaceCustomQuotaResultOutput) IpPrefixQuotas() GetIpSpaceCustomQuotaIpPrefixQuotaArrayOutput {
	return o.ApplyT(func(v LookupIpSpaceCustomQuotaResult) []GetIpSpaceCustomQuotaIpPrefixQuota { return v.IpPrefixQuotas }).(GetIpSpaceCustomQuotaIpPrefixQuotaArrayOutput)
}

func (o LookupIpSpaceCustomQuotaResultOutput) IpRangeQuota() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpSpaceCustomQuotaResult) string { return v.IpRangeQuota }).(pulumi.StringOutput)
}

func (o LookupIpSpaceCustomQuotaResultOutput) IpSpaceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpSpaceCustomQuotaResult) string { return v.IpSpaceId }).(pulumi.StringOutput)
}

func (o LookupIpSpaceCustomQuotaResultOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIpSpaceCustomQuotaResult) string { return v.OrgId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIpSpaceCustomQuotaResultOutput{})
}
