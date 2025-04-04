// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/ergSey/pulumi-vcd/sdk/go/vcd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Supported in provider *v3.13+* with Data Solution Extension.
//
// Provides a data source to read Data Solution publishing settings for a particular tenant.
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
//			tenant_org, err := vcd.LookupOrg(ctx, &vcd.LookupOrgArgs{
//				Name: "tenant_org",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vcd.LookupDseSolutionPublish(ctx, &vcd.LookupDseSolutionPublishArgs{
//				DataSolutionId: mongodb_communityVcdDseRegistryConfiguration.Id,
//				OrgId:          tenant_org.Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupDseSolutionPublish(ctx *pulumi.Context, args *LookupDseSolutionPublishArgs, opts ...pulumi.InvokeOption) (*LookupDseSolutionPublishResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDseSolutionPublishResult
	err := ctx.Invoke("vcd:index/getDseSolutionPublish:getDseSolutionPublish", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDseSolutionPublish.
type LookupDseSolutionPublishArgs struct {
	// ID of Data Solution
	DataSolutionId string `pulumi:"dataSolutionId"`
	// Organization ID
	OrgId string `pulumi:"orgId"`
}

// A collection of values returned by getDseSolutionPublish.
type LookupDseSolutionPublishResult struct {
	ConfluentLicenseType string `pulumi:"confluentLicenseType"`
	DataSolutionId       string `pulumi:"dataSolutionId"`
	DsOrgConfigId        string `pulumi:"dsOrgConfigId"`
	DsoAclId             string `pulumi:"dsoAclId"`
	// The provider-assigned unique ID for this managed resource.
	Id             string   `pulumi:"id"`
	OrgId          string   `pulumi:"orgId"`
	TemplateAclIds []string `pulumi:"templateAclIds"`
}

func LookupDseSolutionPublishOutput(ctx *pulumi.Context, args LookupDseSolutionPublishOutputArgs, opts ...pulumi.InvokeOption) LookupDseSolutionPublishResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDseSolutionPublishResultOutput, error) {
			args := v.(LookupDseSolutionPublishArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vcd:index/getDseSolutionPublish:getDseSolutionPublish", args, LookupDseSolutionPublishResultOutput{}, options).(LookupDseSolutionPublishResultOutput), nil
		}).(LookupDseSolutionPublishResultOutput)
}

// A collection of arguments for invoking getDseSolutionPublish.
type LookupDseSolutionPublishOutputArgs struct {
	// ID of Data Solution
	DataSolutionId pulumi.StringInput `pulumi:"dataSolutionId"`
	// Organization ID
	OrgId pulumi.StringInput `pulumi:"orgId"`
}

func (LookupDseSolutionPublishOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDseSolutionPublishArgs)(nil)).Elem()
}

// A collection of values returned by getDseSolutionPublish.
type LookupDseSolutionPublishResultOutput struct{ *pulumi.OutputState }

func (LookupDseSolutionPublishResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDseSolutionPublishResult)(nil)).Elem()
}

func (o LookupDseSolutionPublishResultOutput) ToLookupDseSolutionPublishResultOutput() LookupDseSolutionPublishResultOutput {
	return o
}

func (o LookupDseSolutionPublishResultOutput) ToLookupDseSolutionPublishResultOutputWithContext(ctx context.Context) LookupDseSolutionPublishResultOutput {
	return o
}

func (o LookupDseSolutionPublishResultOutput) ConfluentLicenseType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDseSolutionPublishResult) string { return v.ConfluentLicenseType }).(pulumi.StringOutput)
}

func (o LookupDseSolutionPublishResultOutput) DataSolutionId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDseSolutionPublishResult) string { return v.DataSolutionId }).(pulumi.StringOutput)
}

func (o LookupDseSolutionPublishResultOutput) DsOrgConfigId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDseSolutionPublishResult) string { return v.DsOrgConfigId }).(pulumi.StringOutput)
}

func (o LookupDseSolutionPublishResultOutput) DsoAclId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDseSolutionPublishResult) string { return v.DsoAclId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDseSolutionPublishResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDseSolutionPublishResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDseSolutionPublishResultOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDseSolutionPublishResult) string { return v.OrgId }).(pulumi.StringOutput)
}

func (o LookupDseSolutionPublishResultOutput) TemplateAclIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDseSolutionPublishResult) []string { return v.TemplateAclIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDseSolutionPublishResultOutput{})
}
