// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/ergSey/pulumi-vcd/sdk/go/vcd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a data source for available NSX-T Tier-0 routers.
//
// Supported in provider *v3.0+*
//
// > **Note:** This resource uses new VMware Cloud Director
// [OpenAPI](https://code.vmware.com/docs/11982/getting-started-with-vmware-cloud-director-openapi) and
// requires at least VCD *10.1.1+* and NSX-T *3.0+*.
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
//			main, err := vcd.GetNsxtManager(ctx, &vcd.GetNsxtManagerArgs{
//				Name: "nsxt-manager-one",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vcd.GetNsxtTier0Router(ctx, &vcd.GetNsxtTier0RouterArgs{
//				Name:          "nsxt-tier0-router",
//				NsxtManagerId: main.Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetNsxtTier0Router(ctx *pulumi.Context, args *GetNsxtTier0RouterArgs, opts ...pulumi.InvokeOption) (*GetNsxtTier0RouterResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetNsxtTier0RouterResult
	err := ctx.Invoke("vcd:index/getNsxtTier0Router:getNsxtTier0Router", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNsxtTier0Router.
type GetNsxtTier0RouterArgs struct {
	// NSX-T Tier-0 router name. **Note**. Tier-0 router name must be unique inside NSX-T manager because
	// API does not allow to filter by other fields.
	Name string `pulumi:"name"`
	// NSX-T manager should be referenced.
	NsxtManagerId string `pulumi:"nsxtManagerId"`
}

// A collection of values returned by getNsxtTier0Router.
type GetNsxtTier0RouterResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// Boolean value reflecting if Tier-0 router is already consumed by external network.
	IsAssigned    bool   `pulumi:"isAssigned"`
	Name          string `pulumi:"name"`
	NsxtManagerId string `pulumi:"nsxtManagerId"`
}

func GetNsxtTier0RouterOutput(ctx *pulumi.Context, args GetNsxtTier0RouterOutputArgs, opts ...pulumi.InvokeOption) GetNsxtTier0RouterResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetNsxtTier0RouterResultOutput, error) {
			args := v.(GetNsxtTier0RouterArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vcd:index/getNsxtTier0Router:getNsxtTier0Router", args, GetNsxtTier0RouterResultOutput{}, options).(GetNsxtTier0RouterResultOutput), nil
		}).(GetNsxtTier0RouterResultOutput)
}

// A collection of arguments for invoking getNsxtTier0Router.
type GetNsxtTier0RouterOutputArgs struct {
	// NSX-T Tier-0 router name. **Note**. Tier-0 router name must be unique inside NSX-T manager because
	// API does not allow to filter by other fields.
	Name pulumi.StringInput `pulumi:"name"`
	// NSX-T manager should be referenced.
	NsxtManagerId pulumi.StringInput `pulumi:"nsxtManagerId"`
}

func (GetNsxtTier0RouterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNsxtTier0RouterArgs)(nil)).Elem()
}

// A collection of values returned by getNsxtTier0Router.
type GetNsxtTier0RouterResultOutput struct{ *pulumi.OutputState }

func (GetNsxtTier0RouterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNsxtTier0RouterResult)(nil)).Elem()
}

func (o GetNsxtTier0RouterResultOutput) ToGetNsxtTier0RouterResultOutput() GetNsxtTier0RouterResultOutput {
	return o
}

func (o GetNsxtTier0RouterResultOutput) ToGetNsxtTier0RouterResultOutputWithContext(ctx context.Context) GetNsxtTier0RouterResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetNsxtTier0RouterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetNsxtTier0RouterResult) string { return v.Id }).(pulumi.StringOutput)
}

// Boolean value reflecting if Tier-0 router is already consumed by external network.
func (o GetNsxtTier0RouterResultOutput) IsAssigned() pulumi.BoolOutput {
	return o.ApplyT(func(v GetNsxtTier0RouterResult) bool { return v.IsAssigned }).(pulumi.BoolOutput)
}

func (o GetNsxtTier0RouterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetNsxtTier0RouterResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetNsxtTier0RouterResultOutput) NsxtManagerId() pulumi.StringOutput {
	return o.ApplyT(func(v GetNsxtTier0RouterResult) string { return v.NsxtManagerId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetNsxtTier0RouterResultOutput{})
}
