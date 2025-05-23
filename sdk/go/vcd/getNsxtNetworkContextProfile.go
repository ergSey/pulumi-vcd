// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/ergSey/pulumi-vcd/sdk/go/vcd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a data source for NSX-T Network Context Profile lookup to later be used in Distributed
// Firewall.
//
// ## Example Usage
//
// ### SYSTEM Scope Network Context Profile Lookup In A VDC Group)
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
//			existing, err := vcd.LookupVdcGroup(ctx, &vcd.LookupVdcGroupArgs{
//				Org:  pulumi.StringRef("my-org"),
//				Name: pulumi.StringRef("main-vdc-group"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vcd.GetNsxtNetworkContextProfile(ctx, &vcd.GetNsxtNetworkContextProfileArgs{
//				ContextId: existing.Id,
//				Name:      "CTRXICA",
//				Scope:     pulumi.StringRef("SYSTEM"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ### SYSTEM Profile Lookup In An NSX-T Manager)
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
//				Name: "first-nsxt-manager",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vcd.GetNsxtNetworkContextProfile(ctx, &vcd.GetNsxtNetworkContextProfileArgs{
//				ContextId: main.Id,
//				Name:      "CTRXICA",
//				Scope:     pulumi.StringRef("SYSTEM"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetNsxtNetworkContextProfile(ctx *pulumi.Context, args *GetNsxtNetworkContextProfileArgs, opts ...pulumi.InvokeOption) (*GetNsxtNetworkContextProfileResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetNsxtNetworkContextProfileResult
	err := ctx.Invoke("vcd:index/getNsxtNetworkContextProfile:getNsxtNetworkContextProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNsxtNetworkContextProfile.
type GetNsxtNetworkContextProfileArgs struct {
	// Context ID specifies the context for Network Context Profile look up.
	// This ID can be one of `VDC Group ID` (data source `VdcGroup`), `Org VDC ID` (data source
	// `OrgVdc`), or `NSX-T Manager ID` (data source `getNsxtManager`)
	ContextId string `pulumi:"contextId"`
	// Name of Network Context Profile
	Name string `pulumi:"name"`
	// Can be one of `SYSTEM`, `TENANT`, `PROVIDER`. (default `SYSTEM`)
	Scope *string `pulumi:"scope"`
}

// A collection of values returned by getNsxtNetworkContextProfile.
type GetNsxtNetworkContextProfileResult struct {
	ContextId string `pulumi:"contextId"`
	// The provider-assigned unique ID for this managed resource.
	Id    string  `pulumi:"id"`
	Name  string  `pulumi:"name"`
	Scope *string `pulumi:"scope"`
}

func GetNsxtNetworkContextProfileOutput(ctx *pulumi.Context, args GetNsxtNetworkContextProfileOutputArgs, opts ...pulumi.InvokeOption) GetNsxtNetworkContextProfileResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetNsxtNetworkContextProfileResultOutput, error) {
			args := v.(GetNsxtNetworkContextProfileArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vcd:index/getNsxtNetworkContextProfile:getNsxtNetworkContextProfile", args, GetNsxtNetworkContextProfileResultOutput{}, options).(GetNsxtNetworkContextProfileResultOutput), nil
		}).(GetNsxtNetworkContextProfileResultOutput)
}

// A collection of arguments for invoking getNsxtNetworkContextProfile.
type GetNsxtNetworkContextProfileOutputArgs struct {
	// Context ID specifies the context for Network Context Profile look up.
	// This ID can be one of `VDC Group ID` (data source `VdcGroup`), `Org VDC ID` (data source
	// `OrgVdc`), or `NSX-T Manager ID` (data source `getNsxtManager`)
	ContextId pulumi.StringInput `pulumi:"contextId"`
	// Name of Network Context Profile
	Name pulumi.StringInput `pulumi:"name"`
	// Can be one of `SYSTEM`, `TENANT`, `PROVIDER`. (default `SYSTEM`)
	Scope pulumi.StringPtrInput `pulumi:"scope"`
}

func (GetNsxtNetworkContextProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNsxtNetworkContextProfileArgs)(nil)).Elem()
}

// A collection of values returned by getNsxtNetworkContextProfile.
type GetNsxtNetworkContextProfileResultOutput struct{ *pulumi.OutputState }

func (GetNsxtNetworkContextProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetNsxtNetworkContextProfileResult)(nil)).Elem()
}

func (o GetNsxtNetworkContextProfileResultOutput) ToGetNsxtNetworkContextProfileResultOutput() GetNsxtNetworkContextProfileResultOutput {
	return o
}

func (o GetNsxtNetworkContextProfileResultOutput) ToGetNsxtNetworkContextProfileResultOutputWithContext(ctx context.Context) GetNsxtNetworkContextProfileResultOutput {
	return o
}

func (o GetNsxtNetworkContextProfileResultOutput) ContextId() pulumi.StringOutput {
	return o.ApplyT(func(v GetNsxtNetworkContextProfileResult) string { return v.ContextId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetNsxtNetworkContextProfileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetNsxtNetworkContextProfileResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetNsxtNetworkContextProfileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetNsxtNetworkContextProfileResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetNsxtNetworkContextProfileResultOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetNsxtNetworkContextProfileResult) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetNsxtNetworkContextProfileResultOutput{})
}
