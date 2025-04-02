// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/ergSey/pulumi-vcd/sdk/go/vcd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Supported in provider *v3.14+* and VCD 10.5+ with NSX-T and ALB.
//
// Provides a data source to read ALB Service Engine Groups policies for HTTP requests. HTTP request
// rules modify requests before they are either forwarded to the application, used as a basis for
// content switching, or discarded.
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
//			_, err := vcd.LookupNsxtAlbVirtualServiceHttpReqRules(ctx, &vcd.LookupNsxtAlbVirtualServiceHttpReqRulesArgs{
//				VirtualServiceId: test.Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupNsxtAlbVirtualServiceHttpReqRules(ctx *pulumi.Context, args *LookupNsxtAlbVirtualServiceHttpReqRulesArgs, opts ...pulumi.InvokeOption) (*LookupNsxtAlbVirtualServiceHttpReqRulesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNsxtAlbVirtualServiceHttpReqRulesResult
	err := ctx.Invoke("vcd:index/getNsxtAlbVirtualServiceHttpReqRules:getNsxtAlbVirtualServiceHttpReqRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNsxtAlbVirtualServiceHttpReqRules.
type LookupNsxtAlbVirtualServiceHttpReqRulesArgs struct {
	// An ID of existing ALB Virtual Service.
	VirtualServiceId string `pulumi:"virtualServiceId"`
}

// A collection of values returned by getNsxtAlbVirtualServiceHttpReqRules.
type LookupNsxtAlbVirtualServiceHttpReqRulesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id               string                                     `pulumi:"id"`
	Rules            []GetNsxtAlbVirtualServiceHttpReqRulesRule `pulumi:"rules"`
	VirtualServiceId string                                     `pulumi:"virtualServiceId"`
}

func LookupNsxtAlbVirtualServiceHttpReqRulesOutput(ctx *pulumi.Context, args LookupNsxtAlbVirtualServiceHttpReqRulesOutputArgs, opts ...pulumi.InvokeOption) LookupNsxtAlbVirtualServiceHttpReqRulesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupNsxtAlbVirtualServiceHttpReqRulesResultOutput, error) {
			args := v.(LookupNsxtAlbVirtualServiceHttpReqRulesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vcd:index/getNsxtAlbVirtualServiceHttpReqRules:getNsxtAlbVirtualServiceHttpReqRules", args, LookupNsxtAlbVirtualServiceHttpReqRulesResultOutput{}, options).(LookupNsxtAlbVirtualServiceHttpReqRulesResultOutput), nil
		}).(LookupNsxtAlbVirtualServiceHttpReqRulesResultOutput)
}

// A collection of arguments for invoking getNsxtAlbVirtualServiceHttpReqRules.
type LookupNsxtAlbVirtualServiceHttpReqRulesOutputArgs struct {
	// An ID of existing ALB Virtual Service.
	VirtualServiceId pulumi.StringInput `pulumi:"virtualServiceId"`
}

func (LookupNsxtAlbVirtualServiceHttpReqRulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxtAlbVirtualServiceHttpReqRulesArgs)(nil)).Elem()
}

// A collection of values returned by getNsxtAlbVirtualServiceHttpReqRules.
type LookupNsxtAlbVirtualServiceHttpReqRulesResultOutput struct{ *pulumi.OutputState }

func (LookupNsxtAlbVirtualServiceHttpReqRulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxtAlbVirtualServiceHttpReqRulesResult)(nil)).Elem()
}

func (o LookupNsxtAlbVirtualServiceHttpReqRulesResultOutput) ToLookupNsxtAlbVirtualServiceHttpReqRulesResultOutput() LookupNsxtAlbVirtualServiceHttpReqRulesResultOutput {
	return o
}

func (o LookupNsxtAlbVirtualServiceHttpReqRulesResultOutput) ToLookupNsxtAlbVirtualServiceHttpReqRulesResultOutputWithContext(ctx context.Context) LookupNsxtAlbVirtualServiceHttpReqRulesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNsxtAlbVirtualServiceHttpReqRulesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAlbVirtualServiceHttpReqRulesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNsxtAlbVirtualServiceHttpReqRulesResultOutput) Rules() GetNsxtAlbVirtualServiceHttpReqRulesRuleArrayOutput {
	return o.ApplyT(func(v LookupNsxtAlbVirtualServiceHttpReqRulesResult) []GetNsxtAlbVirtualServiceHttpReqRulesRule {
		return v.Rules
	}).(GetNsxtAlbVirtualServiceHttpReqRulesRuleArrayOutput)
}

func (o LookupNsxtAlbVirtualServiceHttpReqRulesResultOutput) VirtualServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAlbVirtualServiceHttpReqRulesResult) string { return v.VirtualServiceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNsxtAlbVirtualServiceHttpReqRulesResultOutput{})
}
