// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/ergSey/pulumi-vcd/sdk/go/vcd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Experimental in provider *3.11*.
//
// Provides a data source to read vGPU policies in VMware Cloud Director.
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
//			tf_policy_name, err := vcd.LookupVmVgpuPolicy(ctx, &vcd.LookupVmVgpuPolicyArgs{
//				Name: "my-rule",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("policyId", tf_policy_name.Id)
//			return nil
//		})
//	}
//
// ```
func LookupVmVgpuPolicy(ctx *pulumi.Context, args *LookupVmVgpuPolicyArgs, opts ...pulumi.InvokeOption) (*LookupVmVgpuPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVmVgpuPolicyResult
	err := ctx.Invoke("vcd:index/getVmVgpuPolicy:getVmVgpuPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVmVgpuPolicy.
type LookupVmVgpuPolicyArgs struct {
	// The name of the VM vGPU policy.
	//
	// All arguments defined in [`VmVgpuPolicy`](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vm_vgpu_policy#argument-reference) are supported.
	Name string `pulumi:"name"`
}

// A collection of values returned by getVmVgpuPolicy.
type LookupVmVgpuPolicyResult struct {
	Cpus        []GetVmVgpuPolicyCpus `pulumi:"cpus"`
	Description string                `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id                string                            `pulumi:"id"`
	Memories          []GetVmVgpuPolicyMemory           `pulumi:"memories"`
	Name              string                            `pulumi:"name"`
	ProviderVdcScopes []GetVmVgpuPolicyProviderVdcScope `pulumi:"providerVdcScopes"`
	VgpuProfiles      []GetVmVgpuPolicyVgpuProfile      `pulumi:"vgpuProfiles"`
}

func LookupVmVgpuPolicyOutput(ctx *pulumi.Context, args LookupVmVgpuPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupVmVgpuPolicyResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupVmVgpuPolicyResultOutput, error) {
			args := v.(LookupVmVgpuPolicyArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vcd:index/getVmVgpuPolicy:getVmVgpuPolicy", args, LookupVmVgpuPolicyResultOutput{}, options).(LookupVmVgpuPolicyResultOutput), nil
		}).(LookupVmVgpuPolicyResultOutput)
}

// A collection of arguments for invoking getVmVgpuPolicy.
type LookupVmVgpuPolicyOutputArgs struct {
	// The name of the VM vGPU policy.
	//
	// All arguments defined in [`VmVgpuPolicy`](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vm_vgpu_policy#argument-reference) are supported.
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupVmVgpuPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVmVgpuPolicyArgs)(nil)).Elem()
}

// A collection of values returned by getVmVgpuPolicy.
type LookupVmVgpuPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupVmVgpuPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVmVgpuPolicyResult)(nil)).Elem()
}

func (o LookupVmVgpuPolicyResultOutput) ToLookupVmVgpuPolicyResultOutput() LookupVmVgpuPolicyResultOutput {
	return o
}

func (o LookupVmVgpuPolicyResultOutput) ToLookupVmVgpuPolicyResultOutputWithContext(ctx context.Context) LookupVmVgpuPolicyResultOutput {
	return o
}

func (o LookupVmVgpuPolicyResultOutput) Cpus() GetVmVgpuPolicyCpusArrayOutput {
	return o.ApplyT(func(v LookupVmVgpuPolicyResult) []GetVmVgpuPolicyCpus { return v.Cpus }).(GetVmVgpuPolicyCpusArrayOutput)
}

func (o LookupVmVgpuPolicyResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVmVgpuPolicyResult) string { return v.Description }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVmVgpuPolicyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVmVgpuPolicyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVmVgpuPolicyResultOutput) Memories() GetVmVgpuPolicyMemoryArrayOutput {
	return o.ApplyT(func(v LookupVmVgpuPolicyResult) []GetVmVgpuPolicyMemory { return v.Memories }).(GetVmVgpuPolicyMemoryArrayOutput)
}

func (o LookupVmVgpuPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVmVgpuPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVmVgpuPolicyResultOutput) ProviderVdcScopes() GetVmVgpuPolicyProviderVdcScopeArrayOutput {
	return o.ApplyT(func(v LookupVmVgpuPolicyResult) []GetVmVgpuPolicyProviderVdcScope { return v.ProviderVdcScopes }).(GetVmVgpuPolicyProviderVdcScopeArrayOutput)
}

func (o LookupVmVgpuPolicyResultOutput) VgpuProfiles() GetVmVgpuPolicyVgpuProfileArrayOutput {
	return o.ApplyT(func(v LookupVmVgpuPolicyResult) []GetVmVgpuPolicyVgpuProfile { return v.VgpuProfiles }).(GetVmVgpuPolicyVgpuProfileArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVmVgpuPolicyResultOutput{})
}
