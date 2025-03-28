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
// Provides a data source to read ALB Service Engine Groups policies for HTTP responses. HTTP response
// rules can be used to to evaluate and modify the response and response attributes that the
// application returns.
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
//			_, err := vcd.LookupNsxtAlbVirtualServiceHttpRespRules(ctx, &vcd.LookupNsxtAlbVirtualServiceHttpRespRulesArgs{
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
func LookupNsxtAlbVirtualServiceHttpRespRules(ctx *pulumi.Context, args *LookupNsxtAlbVirtualServiceHttpRespRulesArgs, opts ...pulumi.InvokeOption) (*LookupNsxtAlbVirtualServiceHttpRespRulesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNsxtAlbVirtualServiceHttpRespRulesResult
	err := ctx.Invoke("vcd:index/getNsxtAlbVirtualServiceHttpRespRules:getNsxtAlbVirtualServiceHttpRespRules", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNsxtAlbVirtualServiceHttpRespRules.
type LookupNsxtAlbVirtualServiceHttpRespRulesArgs struct {
	// An ID of existing ALB Virtual Service.
	VirtualServiceId string `pulumi:"virtualServiceId"`
}

// A collection of values returned by getNsxtAlbVirtualServiceHttpRespRules.
type LookupNsxtAlbVirtualServiceHttpRespRulesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id               string                                      `pulumi:"id"`
	Rules            []GetNsxtAlbVirtualServiceHttpRespRulesRule `pulumi:"rules"`
	VirtualServiceId string                                      `pulumi:"virtualServiceId"`
}

func LookupNsxtAlbVirtualServiceHttpRespRulesOutput(ctx *pulumi.Context, args LookupNsxtAlbVirtualServiceHttpRespRulesOutputArgs, opts ...pulumi.InvokeOption) LookupNsxtAlbVirtualServiceHttpRespRulesResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupNsxtAlbVirtualServiceHttpRespRulesResultOutput, error) {
			args := v.(LookupNsxtAlbVirtualServiceHttpRespRulesArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vcd:index/getNsxtAlbVirtualServiceHttpRespRules:getNsxtAlbVirtualServiceHttpRespRules", args, LookupNsxtAlbVirtualServiceHttpRespRulesResultOutput{}, options).(LookupNsxtAlbVirtualServiceHttpRespRulesResultOutput), nil
		}).(LookupNsxtAlbVirtualServiceHttpRespRulesResultOutput)
}

// A collection of arguments for invoking getNsxtAlbVirtualServiceHttpRespRules.
type LookupNsxtAlbVirtualServiceHttpRespRulesOutputArgs struct {
	// An ID of existing ALB Virtual Service.
	VirtualServiceId pulumi.StringInput `pulumi:"virtualServiceId"`
}

func (LookupNsxtAlbVirtualServiceHttpRespRulesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxtAlbVirtualServiceHttpRespRulesArgs)(nil)).Elem()
}

// A collection of values returned by getNsxtAlbVirtualServiceHttpRespRules.
type LookupNsxtAlbVirtualServiceHttpRespRulesResultOutput struct{ *pulumi.OutputState }

func (LookupNsxtAlbVirtualServiceHttpRespRulesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxtAlbVirtualServiceHttpRespRulesResult)(nil)).Elem()
}

func (o LookupNsxtAlbVirtualServiceHttpRespRulesResultOutput) ToLookupNsxtAlbVirtualServiceHttpRespRulesResultOutput() LookupNsxtAlbVirtualServiceHttpRespRulesResultOutput {
	return o
}

func (o LookupNsxtAlbVirtualServiceHttpRespRulesResultOutput) ToLookupNsxtAlbVirtualServiceHttpRespRulesResultOutputWithContext(ctx context.Context) LookupNsxtAlbVirtualServiceHttpRespRulesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNsxtAlbVirtualServiceHttpRespRulesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAlbVirtualServiceHttpRespRulesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNsxtAlbVirtualServiceHttpRespRulesResultOutput) Rules() GetNsxtAlbVirtualServiceHttpRespRulesRuleArrayOutput {
	return o.ApplyT(func(v LookupNsxtAlbVirtualServiceHttpRespRulesResult) []GetNsxtAlbVirtualServiceHttpRespRulesRule {
		return v.Rules
	}).(GetNsxtAlbVirtualServiceHttpRespRulesRuleArrayOutput)
}

func (o LookupNsxtAlbVirtualServiceHttpRespRulesResultOutput) VirtualServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtAlbVirtualServiceHttpRespRulesResult) string { return v.VirtualServiceId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNsxtAlbVirtualServiceHttpRespRulesResultOutput{})
}
