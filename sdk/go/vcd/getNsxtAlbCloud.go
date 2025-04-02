// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/ergSey/pulumi-vcd/sdk/go/vcd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Supported in provider *v3.4+* and VCD 10.2+ with NSX-T and ALB.
//
// Provides a data source to manage ALB Clouds for Providers. An NSX-T Cloud is a service provider-level construct that
// consists of an NSX-T Manager and an NSX-T Data Center transport zone.
//
// > Only `System Administrator` can use this data source.
//
// > VCD 10.3.0 has a caching bug which prevents listing importable clouds immediately (retrieved using
// [`getNsxtAlbImportableCloud`](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_alb_importable_cloud)) after ALB
// Controller is created. This data should be available 15 minutes after the Controller is created.
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
//			_, err := vcd.LookupNsxtAlbCloud(ctx, &vcd.LookupNsxtAlbCloudArgs{
//				Name: "cloud-one",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupNsxtAlbCloud(ctx *pulumi.Context, args *LookupNsxtAlbCloudArgs, opts ...pulumi.InvokeOption) (*LookupNsxtAlbCloudResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNsxtAlbCloudResult
	err := ctx.Invoke("vcd:index/getNsxtAlbCloud:getNsxtAlbCloud", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNsxtAlbCloud.
type LookupNsxtAlbCloudArgs struct {
	// Name of ALB Cloud
	Name string `pulumi:"name"`
}

// A collection of values returned by getNsxtAlbCloud.
type LookupNsxtAlbCloudResult struct {
	ControllerId  string `pulumi:"controllerId"`
	Description   string `pulumi:"description"`
	HealthMessage string `pulumi:"healthMessage"`
	HealthStatus  string `pulumi:"healthStatus"`
	// The provider-assigned unique ID for this managed resource.
	Id                string `pulumi:"id"`
	ImportableCloudId string `pulumi:"importableCloudId"`
	Name              string `pulumi:"name"`
	NetworkPoolId     string `pulumi:"networkPoolId"`
	NetworkPoolName   string `pulumi:"networkPoolName"`
}

func LookupNsxtAlbCloudOutput(ctx *pulumi.Context, args LookupNsxtAlbCloudOutputArgs, opts ...pulumi.InvokeOption) LookupNsxtAlbCloudResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupNsxtAlbCloudResultOutput, error) {
			args := v.(LookupNsxtAlbCloudArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vcd:index/getNsxtAlbCloud:getNsxtAlbCloud", args, LookupNsxtAlbCloudResultOutput{}, options).(LookupNsxtAlbCloudResultOutput), nil
		}).(LookupNsxtAlbCloudResultOutput)
}

// A collection of arguments for invoking getNsxtAlbCloud.
type LookupNsxtAlbCloudOutputArgs struct {
	// Name of ALB Cloud
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupNsxtAlbCloudOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxtAlbCloudArgs)(nil)).Elem()
}

// A collection of values returned by getNsxtAlbCloud.
type LookupNsxtAlbCloudResultOutput struct{ *pulumi.OutputState }

func (LookupNsxtAlbCloudResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxtAlbCloudResult)(nil)).Elem()
}

func (o LookupNsxtAlbCloudResultOutput) ToLookupNsxtAlbCloudResultOutput() LookupNsxtAlbCloudResultOutput {
	return o
}

func (o LookupNsxtAlbCloudResultOutput) ToLookupNsxtAlbCloudResultOutputWithContext(ctx context.Context) LookupNsxtAlbCloudResultOutput {
	return o
}

func (o LookupNsxtAlbCloudResultOutput) ControllerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAlbCloudResult) string { return v.ControllerId }).(pulumi.StringOutput)
}

func (o LookupNsxtAlbCloudResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAlbCloudResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupNsxtAlbCloudResultOutput) HealthMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAlbCloudResult) string { return v.HealthMessage }).(pulumi.StringOutput)
}

func (o LookupNsxtAlbCloudResultOutput) HealthStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAlbCloudResult) string { return v.HealthStatus }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNsxtAlbCloudResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAlbCloudResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNsxtAlbCloudResultOutput) ImportableCloudId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAlbCloudResult) string { return v.ImportableCloudId }).(pulumi.StringOutput)
}

func (o LookupNsxtAlbCloudResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAlbCloudResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNsxtAlbCloudResultOutput) NetworkPoolId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAlbCloudResult) string { return v.NetworkPoolId }).(pulumi.StringOutput)
}

func (o LookupNsxtAlbCloudResultOutput) NetworkPoolName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAlbCloudResult) string { return v.NetworkPoolName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNsxtAlbCloudResultOutput{})
}
