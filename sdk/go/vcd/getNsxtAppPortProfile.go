// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/ergSey/pulumi-vcd/sdk/go/vcd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Supported in provider *v3.3+* and VCD 10.1+ with NSX-T backed VDCs.
//
// Provides a data source to read NSX-T Application Port Profiles. Application Port Profiles include a
// combination of a protocol and a port, or a group of ports, that is used for Firewall and NAT
// services on the Edge Gateway.
//
// ## Example Usage
//
// ### 1 (Find An Application Port Profile Defined By Provider)
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
//			_, err := vcd.LookupNsxtAppPortProfile(ctx, &vcd.LookupNsxtAppPortProfileArgs{
//				Org:       pulumi.StringRef("System"),
//				ContextId: pulumi.StringRef(first.Id),
//				Name:      "WINS",
//				Scope:     "PROVIDER",
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
// ### 2 (Find An Application Port Profile Defined By Tenant In A VDC Group)
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
//			g1, err := vcd.LookupVdcGroup(ctx, &vcd.LookupVdcGroupArgs{
//				Org:  pulumi.StringRef("myOrg"),
//				Name: pulumi.StringRef("myVDC"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vcd.LookupNsxtAppPortProfile(ctx, &vcd.LookupNsxtAppPortProfileArgs{
//				Org:       pulumi.StringRef("my-org"),
//				ContextId: pulumi.StringRef(g1.Id),
//				Name:      "SSH-custom",
//				Scope:     "TENANT",
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
// ### 3 (Find A System Defined Application Port Profile)
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
//			vdc1, err := vcd.LookupOrgVdc(ctx, &vcd.LookupOrgVdcArgs{
//				Org:  pulumi.StringRef("myOrg"),
//				Name: "myVDC",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vcd.LookupNsxtAppPortProfile(ctx, &vcd.LookupNsxtAppPortProfileArgs{
//				ContextId: pulumi.StringRef(vdc1.Id),
//				Scope:     "SYSTEM",
//				Name:      "SSH",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupNsxtAppPortProfile(ctx *pulumi.Context, args *LookupNsxtAppPortProfileArgs, opts ...pulumi.InvokeOption) (*LookupNsxtAppPortProfileResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNsxtAppPortProfileResult
	err := ctx.Invoke("vcd:index/getNsxtAppPortProfile:getNsxtAppPortProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNsxtAppPortProfile.
type LookupNsxtAppPortProfileArgs struct {
	// ID of NSX-T Manager, VDC or VDC Group. Replaces deprecated field `vdc`. Required if using more than one NSX-T Manager.
	ContextId *string `pulumi:"contextId"`
	// Unique name of existing Security Group.
	Name string `pulumi:"name"`
	// Deprecated: Deprecated in favor of 'context_id'
	NsxtManagerId *string `pulumi:"nsxtManagerId"`
	// The name of organization to use, optional if defined at provider level. Useful
	// when connected as sysadmin working across different organisations.
	Org *string `pulumi:"org"`
	// `SYSTEM`, `PROVIDER`, or `TENANT`.
	Scope string `pulumi:"scope"`
	// The name of VDC to use, optional if defined at provider level.
	// Deprecated and replaced by `contextId`
	//
	// Deprecated: Deprecated in favor of 'context_id'
	Vdc *string `pulumi:"vdc"`
}

// A collection of values returned by getNsxtAppPortProfile.
type LookupNsxtAppPortProfileResult struct {
	AppPorts    []GetNsxtAppPortProfileAppPort `pulumi:"appPorts"`
	ContextId   string                         `pulumi:"contextId"`
	Description string                         `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
	// Deprecated: Deprecated in favor of 'context_id'
	NsxtManagerId *string `pulumi:"nsxtManagerId"`
	Org           *string `pulumi:"org"`
	Scope         string  `pulumi:"scope"`
	// Deprecated: Deprecated in favor of 'context_id'
	Vdc *string `pulumi:"vdc"`
}

func LookupNsxtAppPortProfileOutput(ctx *pulumi.Context, args LookupNsxtAppPortProfileOutputArgs, opts ...pulumi.InvokeOption) LookupNsxtAppPortProfileResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupNsxtAppPortProfileResultOutput, error) {
			args := v.(LookupNsxtAppPortProfileArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vcd:index/getNsxtAppPortProfile:getNsxtAppPortProfile", args, LookupNsxtAppPortProfileResultOutput{}, options).(LookupNsxtAppPortProfileResultOutput), nil
		}).(LookupNsxtAppPortProfileResultOutput)
}

// A collection of arguments for invoking getNsxtAppPortProfile.
type LookupNsxtAppPortProfileOutputArgs struct {
	// ID of NSX-T Manager, VDC or VDC Group. Replaces deprecated field `vdc`. Required if using more than one NSX-T Manager.
	ContextId pulumi.StringPtrInput `pulumi:"contextId"`
	// Unique name of existing Security Group.
	Name pulumi.StringInput `pulumi:"name"`
	// Deprecated: Deprecated in favor of 'context_id'
	NsxtManagerId pulumi.StringPtrInput `pulumi:"nsxtManagerId"`
	// The name of organization to use, optional if defined at provider level. Useful
	// when connected as sysadmin working across different organisations.
	Org pulumi.StringPtrInput `pulumi:"org"`
	// `SYSTEM`, `PROVIDER`, or `TENANT`.
	Scope pulumi.StringInput `pulumi:"scope"`
	// The name of VDC to use, optional if defined at provider level.
	// Deprecated and replaced by `contextId`
	//
	// Deprecated: Deprecated in favor of 'context_id'
	Vdc pulumi.StringPtrInput `pulumi:"vdc"`
}

func (LookupNsxtAppPortProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxtAppPortProfileArgs)(nil)).Elem()
}

// A collection of values returned by getNsxtAppPortProfile.
type LookupNsxtAppPortProfileResultOutput struct{ *pulumi.OutputState }

func (LookupNsxtAppPortProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxtAppPortProfileResult)(nil)).Elem()
}

func (o LookupNsxtAppPortProfileResultOutput) ToLookupNsxtAppPortProfileResultOutput() LookupNsxtAppPortProfileResultOutput {
	return o
}

func (o LookupNsxtAppPortProfileResultOutput) ToLookupNsxtAppPortProfileResultOutputWithContext(ctx context.Context) LookupNsxtAppPortProfileResultOutput {
	return o
}

func (o LookupNsxtAppPortProfileResultOutput) AppPorts() GetNsxtAppPortProfileAppPortArrayOutput {
	return o.ApplyT(func(v LookupNsxtAppPortProfileResult) []GetNsxtAppPortProfileAppPort { return v.AppPorts }).(GetNsxtAppPortProfileAppPortArrayOutput)
}

func (o LookupNsxtAppPortProfileResultOutput) ContextId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAppPortProfileResult) string { return v.ContextId }).(pulumi.StringOutput)
}

func (o LookupNsxtAppPortProfileResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAppPortProfileResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNsxtAppPortProfileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAppPortProfileResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNsxtAppPortProfileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAppPortProfileResult) string { return v.Name }).(pulumi.StringOutput)
}

// Deprecated: Deprecated in favor of 'context_id'
func (o LookupNsxtAppPortProfileResultOutput) NsxtManagerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNsxtAppPortProfileResult) *string { return v.NsxtManagerId }).(pulumi.StringPtrOutput)
}

func (o LookupNsxtAppPortProfileResultOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNsxtAppPortProfileResult) *string { return v.Org }).(pulumi.StringPtrOutput)
}

func (o LookupNsxtAppPortProfileResultOutput) Scope() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAppPortProfileResult) string { return v.Scope }).(pulumi.StringOutput)
}

// Deprecated: Deprecated in favor of 'context_id'
func (o LookupNsxtAppPortProfileResultOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNsxtAppPortProfileResult) *string { return v.Vdc }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNsxtAppPortProfileResultOutput{})
}
