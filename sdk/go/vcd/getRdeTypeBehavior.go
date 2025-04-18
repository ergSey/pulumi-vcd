// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/ergSey/pulumi-vcd/sdk/go/vcd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides the capability of reading RDE Type Behaviors in VMware Cloud Director, which override an existing [RDE Interface
// Behavior](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/rde_interface_behavior).
//
// Supported in provider *v3.10+*. Requires System Administrator privileges.
func LookupRdeTypeBehavior(ctx *pulumi.Context, args *LookupRdeTypeBehaviorArgs, opts ...pulumi.InvokeOption) (*LookupRdeTypeBehaviorResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupRdeTypeBehaviorResult
	err := ctx.Invoke("vcd:index/getRdeTypeBehavior:getRdeTypeBehavior", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getRdeTypeBehavior.
type LookupRdeTypeBehaviorArgs struct {
	BehaviorId string `pulumi:"behaviorId"`
	// The ID of the [RDE Type](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/rde_type) that owns the Behavior override
	RdeTypeId string `pulumi:"rdeTypeId"`
}

// A collection of values returned by getRdeTypeBehavior.
type LookupRdeTypeBehaviorResult struct {
	BehaviorId  string `pulumi:"behaviorId"`
	Description string `pulumi:"description"`
	// Deprecated: This argument cannot consider complex execution structures. For that purpose, use 'execution_json' instead
	Execution     map[string]string `pulumi:"execution"`
	ExecutionJson string            `pulumi:"executionJson"`
	// The provider-assigned unique ID for this managed resource.
	Id        string `pulumi:"id"`
	Name      string `pulumi:"name"`
	RdeTypeId string `pulumi:"rdeTypeId"`
	Ref       string `pulumi:"ref"`
}

func LookupRdeTypeBehaviorOutput(ctx *pulumi.Context, args LookupRdeTypeBehaviorOutputArgs, opts ...pulumi.InvokeOption) LookupRdeTypeBehaviorResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupRdeTypeBehaviorResultOutput, error) {
			args := v.(LookupRdeTypeBehaviorArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vcd:index/getRdeTypeBehavior:getRdeTypeBehavior", args, LookupRdeTypeBehaviorResultOutput{}, options).(LookupRdeTypeBehaviorResultOutput), nil
		}).(LookupRdeTypeBehaviorResultOutput)
}

// A collection of arguments for invoking getRdeTypeBehavior.
type LookupRdeTypeBehaviorOutputArgs struct {
	BehaviorId pulumi.StringInput `pulumi:"behaviorId"`
	// The ID of the [RDE Type](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/rde_type) that owns the Behavior override
	RdeTypeId pulumi.StringInput `pulumi:"rdeTypeId"`
}

func (LookupRdeTypeBehaviorOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRdeTypeBehaviorArgs)(nil)).Elem()
}

// A collection of values returned by getRdeTypeBehavior.
type LookupRdeTypeBehaviorResultOutput struct{ *pulumi.OutputState }

func (LookupRdeTypeBehaviorResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRdeTypeBehaviorResult)(nil)).Elem()
}

func (o LookupRdeTypeBehaviorResultOutput) ToLookupRdeTypeBehaviorResultOutput() LookupRdeTypeBehaviorResultOutput {
	return o
}

func (o LookupRdeTypeBehaviorResultOutput) ToLookupRdeTypeBehaviorResultOutputWithContext(ctx context.Context) LookupRdeTypeBehaviorResultOutput {
	return o
}

func (o LookupRdeTypeBehaviorResultOutput) BehaviorId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRdeTypeBehaviorResult) string { return v.BehaviorId }).(pulumi.StringOutput)
}

func (o LookupRdeTypeBehaviorResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRdeTypeBehaviorResult) string { return v.Description }).(pulumi.StringOutput)
}

// Deprecated: This argument cannot consider complex execution structures. For that purpose, use 'execution_json' instead
func (o LookupRdeTypeBehaviorResultOutput) Execution() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRdeTypeBehaviorResult) map[string]string { return v.Execution }).(pulumi.StringMapOutput)
}

func (o LookupRdeTypeBehaviorResultOutput) ExecutionJson() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRdeTypeBehaviorResult) string { return v.ExecutionJson }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupRdeTypeBehaviorResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRdeTypeBehaviorResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRdeTypeBehaviorResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRdeTypeBehaviorResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRdeTypeBehaviorResultOutput) RdeTypeId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRdeTypeBehaviorResult) string { return v.RdeTypeId }).(pulumi.StringOutput)
}

func (o LookupRdeTypeBehaviorResultOutput) Ref() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRdeTypeBehaviorResult) string { return v.Ref }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRdeTypeBehaviorResultOutput{})
}
