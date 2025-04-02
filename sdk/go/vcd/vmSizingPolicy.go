// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/ergSey/pulumi-vcd/sdk/go/vcd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VmSizingPolicy struct {
	pulumi.CustomResourceState

	// Configures cpu policy; see Cpu below for details.
	Cpu VmSizingPolicyCpuPtrOutput `pulumi:"cpu"`
	// description of VM sizing policy.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Configures memory policy; see Memory below for details.
	//
	// > **Note:**
	// Previously, it was incorrectly stated that the `org` argument was required. In fact, it is not, and it has been deprecated in the resource schema.
	// To preserve compatibility until the next release, though, the parameter is still parsed, but ignored.
	//
	// <a id="cpu"></a>
	Memory VmSizingPolicyMemoryPtrOutput `pulumi:"memory"`
	// The name of VM sizing policy.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of organization to use - Deprecated and unneeded: will be ignored if used
	//
	// Deprecated: Unneeded property, which was included by mistake
	Org pulumi.StringPtrOutput `pulumi:"org"`
}

// NewVmSizingPolicy registers a new resource with the given unique name, arguments, and options.
func NewVmSizingPolicy(ctx *pulumi.Context,
	name string, args *VmSizingPolicyArgs, opts ...pulumi.ResourceOption) (*VmSizingPolicy, error) {
	if args == nil {
		args = &VmSizingPolicyArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VmSizingPolicy
	err := ctx.RegisterResource("vcd:index/vmSizingPolicy:VmSizingPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVmSizingPolicy gets an existing VmSizingPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVmSizingPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VmSizingPolicyState, opts ...pulumi.ResourceOption) (*VmSizingPolicy, error) {
	var resource VmSizingPolicy
	err := ctx.ReadResource("vcd:index/vmSizingPolicy:VmSizingPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VmSizingPolicy resources.
type vmSizingPolicyState struct {
	// Configures cpu policy; see Cpu below for details.
	Cpu *VmSizingPolicyCpu `pulumi:"cpu"`
	// description of VM sizing policy.
	Description *string `pulumi:"description"`
	// Configures memory policy; see Memory below for details.
	//
	// > **Note:**
	// Previously, it was incorrectly stated that the `org` argument was required. In fact, it is not, and it has been deprecated in the resource schema.
	// To preserve compatibility until the next release, though, the parameter is still parsed, but ignored.
	//
	// <a id="cpu"></a>
	Memory *VmSizingPolicyMemory `pulumi:"memory"`
	// The name of VM sizing policy.
	Name *string `pulumi:"name"`
	// The name of organization to use - Deprecated and unneeded: will be ignored if used
	//
	// Deprecated: Unneeded property, which was included by mistake
	Org *string `pulumi:"org"`
}

type VmSizingPolicyState struct {
	// Configures cpu policy; see Cpu below for details.
	Cpu VmSizingPolicyCpuPtrInput
	// description of VM sizing policy.
	Description pulumi.StringPtrInput
	// Configures memory policy; see Memory below for details.
	//
	// > **Note:**
	// Previously, it was incorrectly stated that the `org` argument was required. In fact, it is not, and it has been deprecated in the resource schema.
	// To preserve compatibility until the next release, though, the parameter is still parsed, but ignored.
	//
	// <a id="cpu"></a>
	Memory VmSizingPolicyMemoryPtrInput
	// The name of VM sizing policy.
	Name pulumi.StringPtrInput
	// The name of organization to use - Deprecated and unneeded: will be ignored if used
	//
	// Deprecated: Unneeded property, which was included by mistake
	Org pulumi.StringPtrInput
}

func (VmSizingPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*vmSizingPolicyState)(nil)).Elem()
}

type vmSizingPolicyArgs struct {
	// Configures cpu policy; see Cpu below for details.
	Cpu *VmSizingPolicyCpu `pulumi:"cpu"`
	// description of VM sizing policy.
	Description *string `pulumi:"description"`
	// Configures memory policy; see Memory below for details.
	//
	// > **Note:**
	// Previously, it was incorrectly stated that the `org` argument was required. In fact, it is not, and it has been deprecated in the resource schema.
	// To preserve compatibility until the next release, though, the parameter is still parsed, but ignored.
	//
	// <a id="cpu"></a>
	Memory *VmSizingPolicyMemory `pulumi:"memory"`
	// The name of VM sizing policy.
	Name *string `pulumi:"name"`
	// The name of organization to use - Deprecated and unneeded: will be ignored if used
	//
	// Deprecated: Unneeded property, which was included by mistake
	Org *string `pulumi:"org"`
}

// The set of arguments for constructing a VmSizingPolicy resource.
type VmSizingPolicyArgs struct {
	// Configures cpu policy; see Cpu below for details.
	Cpu VmSizingPolicyCpuPtrInput
	// description of VM sizing policy.
	Description pulumi.StringPtrInput
	// Configures memory policy; see Memory below for details.
	//
	// > **Note:**
	// Previously, it was incorrectly stated that the `org` argument was required. In fact, it is not, and it has been deprecated in the resource schema.
	// To preserve compatibility until the next release, though, the parameter is still parsed, but ignored.
	//
	// <a id="cpu"></a>
	Memory VmSizingPolicyMemoryPtrInput
	// The name of VM sizing policy.
	Name pulumi.StringPtrInput
	// The name of organization to use - Deprecated and unneeded: will be ignored if used
	//
	// Deprecated: Unneeded property, which was included by mistake
	Org pulumi.StringPtrInput
}

func (VmSizingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vmSizingPolicyArgs)(nil)).Elem()
}

type VmSizingPolicyInput interface {
	pulumi.Input

	ToVmSizingPolicyOutput() VmSizingPolicyOutput
	ToVmSizingPolicyOutputWithContext(ctx context.Context) VmSizingPolicyOutput
}

func (*VmSizingPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**VmSizingPolicy)(nil)).Elem()
}

func (i *VmSizingPolicy) ToVmSizingPolicyOutput() VmSizingPolicyOutput {
	return i.ToVmSizingPolicyOutputWithContext(context.Background())
}

func (i *VmSizingPolicy) ToVmSizingPolicyOutputWithContext(ctx context.Context) VmSizingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmSizingPolicyOutput)
}

// VmSizingPolicyArrayInput is an input type that accepts VmSizingPolicyArray and VmSizingPolicyArrayOutput values.
// You can construct a concrete instance of `VmSizingPolicyArrayInput` via:
//
//	VmSizingPolicyArray{ VmSizingPolicyArgs{...} }
type VmSizingPolicyArrayInput interface {
	pulumi.Input

	ToVmSizingPolicyArrayOutput() VmSizingPolicyArrayOutput
	ToVmSizingPolicyArrayOutputWithContext(context.Context) VmSizingPolicyArrayOutput
}

type VmSizingPolicyArray []VmSizingPolicyInput

func (VmSizingPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VmSizingPolicy)(nil)).Elem()
}

func (i VmSizingPolicyArray) ToVmSizingPolicyArrayOutput() VmSizingPolicyArrayOutput {
	return i.ToVmSizingPolicyArrayOutputWithContext(context.Background())
}

func (i VmSizingPolicyArray) ToVmSizingPolicyArrayOutputWithContext(ctx context.Context) VmSizingPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmSizingPolicyArrayOutput)
}

// VmSizingPolicyMapInput is an input type that accepts VmSizingPolicyMap and VmSizingPolicyMapOutput values.
// You can construct a concrete instance of `VmSizingPolicyMapInput` via:
//
//	VmSizingPolicyMap{ "key": VmSizingPolicyArgs{...} }
type VmSizingPolicyMapInput interface {
	pulumi.Input

	ToVmSizingPolicyMapOutput() VmSizingPolicyMapOutput
	ToVmSizingPolicyMapOutputWithContext(context.Context) VmSizingPolicyMapOutput
}

type VmSizingPolicyMap map[string]VmSizingPolicyInput

func (VmSizingPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VmSizingPolicy)(nil)).Elem()
}

func (i VmSizingPolicyMap) ToVmSizingPolicyMapOutput() VmSizingPolicyMapOutput {
	return i.ToVmSizingPolicyMapOutputWithContext(context.Background())
}

func (i VmSizingPolicyMap) ToVmSizingPolicyMapOutputWithContext(ctx context.Context) VmSizingPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmSizingPolicyMapOutput)
}

type VmSizingPolicyOutput struct{ *pulumi.OutputState }

func (VmSizingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VmSizingPolicy)(nil)).Elem()
}

func (o VmSizingPolicyOutput) ToVmSizingPolicyOutput() VmSizingPolicyOutput {
	return o
}

func (o VmSizingPolicyOutput) ToVmSizingPolicyOutputWithContext(ctx context.Context) VmSizingPolicyOutput {
	return o
}

// Configures cpu policy; see Cpu below for details.
func (o VmSizingPolicyOutput) Cpu() VmSizingPolicyCpuPtrOutput {
	return o.ApplyT(func(v *VmSizingPolicy) VmSizingPolicyCpuPtrOutput { return v.Cpu }).(VmSizingPolicyCpuPtrOutput)
}

// description of VM sizing policy.
func (o VmSizingPolicyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VmSizingPolicy) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Configures memory policy; see Memory below for details.
//
// > **Note:**
// Previously, it was incorrectly stated that the `org` argument was required. In fact, it is not, and it has been deprecated in the resource schema.
// To preserve compatibility until the next release, though, the parameter is still parsed, but ignored.
//
// <a id="cpu"></a>
func (o VmSizingPolicyOutput) Memory() VmSizingPolicyMemoryPtrOutput {
	return o.ApplyT(func(v *VmSizingPolicy) VmSizingPolicyMemoryPtrOutput { return v.Memory }).(VmSizingPolicyMemoryPtrOutput)
}

// The name of VM sizing policy.
func (o VmSizingPolicyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VmSizingPolicy) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of organization to use - Deprecated and unneeded: will be ignored if used
//
// Deprecated: Unneeded property, which was included by mistake
func (o VmSizingPolicyOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VmSizingPolicy) pulumi.StringPtrOutput { return v.Org }).(pulumi.StringPtrOutput)
}

type VmSizingPolicyArrayOutput struct{ *pulumi.OutputState }

func (VmSizingPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VmSizingPolicy)(nil)).Elem()
}

func (o VmSizingPolicyArrayOutput) ToVmSizingPolicyArrayOutput() VmSizingPolicyArrayOutput {
	return o
}

func (o VmSizingPolicyArrayOutput) ToVmSizingPolicyArrayOutputWithContext(ctx context.Context) VmSizingPolicyArrayOutput {
	return o
}

func (o VmSizingPolicyArrayOutput) Index(i pulumi.IntInput) VmSizingPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VmSizingPolicy {
		return vs[0].([]*VmSizingPolicy)[vs[1].(int)]
	}).(VmSizingPolicyOutput)
}

type VmSizingPolicyMapOutput struct{ *pulumi.OutputState }

func (VmSizingPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VmSizingPolicy)(nil)).Elem()
}

func (o VmSizingPolicyMapOutput) ToVmSizingPolicyMapOutput() VmSizingPolicyMapOutput {
	return o
}

func (o VmSizingPolicyMapOutput) ToVmSizingPolicyMapOutputWithContext(ctx context.Context) VmSizingPolicyMapOutput {
	return o
}

func (o VmSizingPolicyMapOutput) MapIndex(k pulumi.StringInput) VmSizingPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VmSizingPolicy {
		return vs[0].(map[string]*VmSizingPolicy)[vs[1].(string)]
	}).(VmSizingPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VmSizingPolicyInput)(nil)).Elem(), &VmSizingPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*VmSizingPolicyArrayInput)(nil)).Elem(), VmSizingPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VmSizingPolicyMapInput)(nil)).Elem(), VmSizingPolicyMap{})
	pulumi.RegisterOutputType(VmSizingPolicyOutput{})
	pulumi.RegisterOutputType(VmSizingPolicyArrayOutput{})
	pulumi.RegisterOutputType(VmSizingPolicyMapOutput{})
}
