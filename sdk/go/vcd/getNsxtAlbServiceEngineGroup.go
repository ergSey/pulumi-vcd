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
// Provides a data source to read ALB Service Engine Groups. A Service Engine Group is an isolation domain that also
// defines shared service engine properties, such as size, network access, and failover. Resources in a service engine
// group can be used for different virtual services, depending on your tenant needs. These resources cannot be shared
// between different service engine groups.
//
// > Only `System Administrator` can use this data source.
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
//			_, err := vcd.LookupNsxtAlbServiceEngineGroup(ctx, &vcd.LookupNsxtAlbServiceEngineGroupArgs{
//				Name:          "configured-service-engine-group",
//				SyncOnRefresh: pulumi.BoolRef(false),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupNsxtAlbServiceEngineGroup(ctx *pulumi.Context, args *LookupNsxtAlbServiceEngineGroupArgs, opts ...pulumi.InvokeOption) (*LookupNsxtAlbServiceEngineGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNsxtAlbServiceEngineGroupResult
	err := ctx.Invoke("vcd:index/getNsxtAlbServiceEngineGroup:getNsxtAlbServiceEngineGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNsxtAlbServiceEngineGroup.
type LookupNsxtAlbServiceEngineGroupArgs struct {
	// Name of Service Engine Group.
	Name          string `pulumi:"name"`
	Overallocated *bool  `pulumi:"overallocated"`
	SyncOnRefresh *bool  `pulumi:"syncOnRefresh"`
}

// A collection of values returned by getNsxtAlbServiceEngineGroup.
type LookupNsxtAlbServiceEngineGroupResult struct {
	AlbCloudId              string `pulumi:"albCloudId"`
	DeployedVirtualServices int    `pulumi:"deployedVirtualServices"`
	Description             string `pulumi:"description"`
	HaMode                  string `pulumi:"haMode"`
	// The provider-assigned unique ID for this managed resource.
	Id                      string `pulumi:"id"`
	MaxVirtualServices      int    `pulumi:"maxVirtualServices"`
	Name                    string `pulumi:"name"`
	Overallocated           *bool  `pulumi:"overallocated"`
	ReservationModel        string `pulumi:"reservationModel"`
	ReservedVirtualServices int    `pulumi:"reservedVirtualServices"`
	SupportedFeatureSet     string `pulumi:"supportedFeatureSet"`
	SyncOnRefresh           *bool  `pulumi:"syncOnRefresh"`
}

func LookupNsxtAlbServiceEngineGroupOutput(ctx *pulumi.Context, args LookupNsxtAlbServiceEngineGroupOutputArgs, opts ...pulumi.InvokeOption) LookupNsxtAlbServiceEngineGroupResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupNsxtAlbServiceEngineGroupResultOutput, error) {
			args := v.(LookupNsxtAlbServiceEngineGroupArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vcd:index/getNsxtAlbServiceEngineGroup:getNsxtAlbServiceEngineGroup", args, LookupNsxtAlbServiceEngineGroupResultOutput{}, options).(LookupNsxtAlbServiceEngineGroupResultOutput), nil
		}).(LookupNsxtAlbServiceEngineGroupResultOutput)
}

// A collection of arguments for invoking getNsxtAlbServiceEngineGroup.
type LookupNsxtAlbServiceEngineGroupOutputArgs struct {
	// Name of Service Engine Group.
	Name          pulumi.StringInput  `pulumi:"name"`
	Overallocated pulumi.BoolPtrInput `pulumi:"overallocated"`
	SyncOnRefresh pulumi.BoolPtrInput `pulumi:"syncOnRefresh"`
}

func (LookupNsxtAlbServiceEngineGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxtAlbServiceEngineGroupArgs)(nil)).Elem()
}

// A collection of values returned by getNsxtAlbServiceEngineGroup.
type LookupNsxtAlbServiceEngineGroupResultOutput struct{ *pulumi.OutputState }

func (LookupNsxtAlbServiceEngineGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxtAlbServiceEngineGroupResult)(nil)).Elem()
}

func (o LookupNsxtAlbServiceEngineGroupResultOutput) ToLookupNsxtAlbServiceEngineGroupResultOutput() LookupNsxtAlbServiceEngineGroupResultOutput {
	return o
}

func (o LookupNsxtAlbServiceEngineGroupResultOutput) ToLookupNsxtAlbServiceEngineGroupResultOutputWithContext(ctx context.Context) LookupNsxtAlbServiceEngineGroupResultOutput {
	return o
}

func (o LookupNsxtAlbServiceEngineGroupResultOutput) AlbCloudId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAlbServiceEngineGroupResult) string { return v.AlbCloudId }).(pulumi.StringOutput)
}

func (o LookupNsxtAlbServiceEngineGroupResultOutput) DeployedVirtualServices() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNsxtAlbServiceEngineGroupResult) int { return v.DeployedVirtualServices }).(pulumi.IntOutput)
}

func (o LookupNsxtAlbServiceEngineGroupResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAlbServiceEngineGroupResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupNsxtAlbServiceEngineGroupResultOutput) HaMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAlbServiceEngineGroupResult) string { return v.HaMode }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNsxtAlbServiceEngineGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAlbServiceEngineGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNsxtAlbServiceEngineGroupResultOutput) MaxVirtualServices() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNsxtAlbServiceEngineGroupResult) int { return v.MaxVirtualServices }).(pulumi.IntOutput)
}

func (o LookupNsxtAlbServiceEngineGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAlbServiceEngineGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNsxtAlbServiceEngineGroupResultOutput) Overallocated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNsxtAlbServiceEngineGroupResult) *bool { return v.Overallocated }).(pulumi.BoolPtrOutput)
}

func (o LookupNsxtAlbServiceEngineGroupResultOutput) ReservationModel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAlbServiceEngineGroupResult) string { return v.ReservationModel }).(pulumi.StringOutput)
}

func (o LookupNsxtAlbServiceEngineGroupResultOutput) ReservedVirtualServices() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNsxtAlbServiceEngineGroupResult) int { return v.ReservedVirtualServices }).(pulumi.IntOutput)
}

func (o LookupNsxtAlbServiceEngineGroupResultOutput) SupportedFeatureSet() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAlbServiceEngineGroupResult) string { return v.SupportedFeatureSet }).(pulumi.StringOutput)
}

func (o LookupNsxtAlbServiceEngineGroupResultOutput) SyncOnRefresh() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupNsxtAlbServiceEngineGroupResult) *bool { return v.SyncOnRefresh }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNsxtAlbServiceEngineGroupResultOutput{})
}
