// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"errors"
	"github.com/ergSey/pulumi-vcd/sdk/go/vcd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type NsxtAlbVirtualServiceHttpReqRules struct {
	pulumi.CustomResourceState

	// One or more rule blocks with matching criteria and actions.
	//
	// <a id="rule-block"></a>
	Rules NsxtAlbVirtualServiceHttpReqRulesRuleArrayOutput `pulumi:"rules"`
	// An ID of existing ALB Virtual Service.
	VirtualServiceId pulumi.StringOutput `pulumi:"virtualServiceId"`
}

// NewNsxtAlbVirtualServiceHttpReqRules registers a new resource with the given unique name, arguments, and options.
func NewNsxtAlbVirtualServiceHttpReqRules(ctx *pulumi.Context,
	name string, args *NsxtAlbVirtualServiceHttpReqRulesArgs, opts ...pulumi.ResourceOption) (*NsxtAlbVirtualServiceHttpReqRules, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Rules == nil {
		return nil, errors.New("invalid value for required argument 'Rules'")
	}
	if args.VirtualServiceId == nil {
		return nil, errors.New("invalid value for required argument 'VirtualServiceId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NsxtAlbVirtualServiceHttpReqRules
	err := ctx.RegisterResource("vcd:index/nsxtAlbVirtualServiceHttpReqRules:NsxtAlbVirtualServiceHttpReqRules", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNsxtAlbVirtualServiceHttpReqRules gets an existing NsxtAlbVirtualServiceHttpReqRules resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNsxtAlbVirtualServiceHttpReqRules(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NsxtAlbVirtualServiceHttpReqRulesState, opts ...pulumi.ResourceOption) (*NsxtAlbVirtualServiceHttpReqRules, error) {
	var resource NsxtAlbVirtualServiceHttpReqRules
	err := ctx.ReadResource("vcd:index/nsxtAlbVirtualServiceHttpReqRules:NsxtAlbVirtualServiceHttpReqRules", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NsxtAlbVirtualServiceHttpReqRules resources.
type nsxtAlbVirtualServiceHttpReqRulesState struct {
	// One or more rule blocks with matching criteria and actions.
	//
	// <a id="rule-block"></a>
	Rules []NsxtAlbVirtualServiceHttpReqRulesRule `pulumi:"rules"`
	// An ID of existing ALB Virtual Service.
	VirtualServiceId *string `pulumi:"virtualServiceId"`
}

type NsxtAlbVirtualServiceHttpReqRulesState struct {
	// One or more rule blocks with matching criteria and actions.
	//
	// <a id="rule-block"></a>
	Rules NsxtAlbVirtualServiceHttpReqRulesRuleArrayInput
	// An ID of existing ALB Virtual Service.
	VirtualServiceId pulumi.StringPtrInput
}

func (NsxtAlbVirtualServiceHttpReqRulesState) ElementType() reflect.Type {
	return reflect.TypeOf((*nsxtAlbVirtualServiceHttpReqRulesState)(nil)).Elem()
}

type nsxtAlbVirtualServiceHttpReqRulesArgs struct {
	// One or more rule blocks with matching criteria and actions.
	//
	// <a id="rule-block"></a>
	Rules []NsxtAlbVirtualServiceHttpReqRulesRule `pulumi:"rules"`
	// An ID of existing ALB Virtual Service.
	VirtualServiceId string `pulumi:"virtualServiceId"`
}

// The set of arguments for constructing a NsxtAlbVirtualServiceHttpReqRules resource.
type NsxtAlbVirtualServiceHttpReqRulesArgs struct {
	// One or more rule blocks with matching criteria and actions.
	//
	// <a id="rule-block"></a>
	Rules NsxtAlbVirtualServiceHttpReqRulesRuleArrayInput
	// An ID of existing ALB Virtual Service.
	VirtualServiceId pulumi.StringInput
}

func (NsxtAlbVirtualServiceHttpReqRulesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nsxtAlbVirtualServiceHttpReqRulesArgs)(nil)).Elem()
}

type NsxtAlbVirtualServiceHttpReqRulesInput interface {
	pulumi.Input

	ToNsxtAlbVirtualServiceHttpReqRulesOutput() NsxtAlbVirtualServiceHttpReqRulesOutput
	ToNsxtAlbVirtualServiceHttpReqRulesOutputWithContext(ctx context.Context) NsxtAlbVirtualServiceHttpReqRulesOutput
}

func (*NsxtAlbVirtualServiceHttpReqRules) ElementType() reflect.Type {
	return reflect.TypeOf((**NsxtAlbVirtualServiceHttpReqRules)(nil)).Elem()
}

func (i *NsxtAlbVirtualServiceHttpReqRules) ToNsxtAlbVirtualServiceHttpReqRulesOutput() NsxtAlbVirtualServiceHttpReqRulesOutput {
	return i.ToNsxtAlbVirtualServiceHttpReqRulesOutputWithContext(context.Background())
}

func (i *NsxtAlbVirtualServiceHttpReqRules) ToNsxtAlbVirtualServiceHttpReqRulesOutputWithContext(ctx context.Context) NsxtAlbVirtualServiceHttpReqRulesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxtAlbVirtualServiceHttpReqRulesOutput)
}

// NsxtAlbVirtualServiceHttpReqRulesArrayInput is an input type that accepts NsxtAlbVirtualServiceHttpReqRulesArray and NsxtAlbVirtualServiceHttpReqRulesArrayOutput values.
// You can construct a concrete instance of `NsxtAlbVirtualServiceHttpReqRulesArrayInput` via:
//
//	NsxtAlbVirtualServiceHttpReqRulesArray{ NsxtAlbVirtualServiceHttpReqRulesArgs{...} }
type NsxtAlbVirtualServiceHttpReqRulesArrayInput interface {
	pulumi.Input

	ToNsxtAlbVirtualServiceHttpReqRulesArrayOutput() NsxtAlbVirtualServiceHttpReqRulesArrayOutput
	ToNsxtAlbVirtualServiceHttpReqRulesArrayOutputWithContext(context.Context) NsxtAlbVirtualServiceHttpReqRulesArrayOutput
}

type NsxtAlbVirtualServiceHttpReqRulesArray []NsxtAlbVirtualServiceHttpReqRulesInput

func (NsxtAlbVirtualServiceHttpReqRulesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NsxtAlbVirtualServiceHttpReqRules)(nil)).Elem()
}

func (i NsxtAlbVirtualServiceHttpReqRulesArray) ToNsxtAlbVirtualServiceHttpReqRulesArrayOutput() NsxtAlbVirtualServiceHttpReqRulesArrayOutput {
	return i.ToNsxtAlbVirtualServiceHttpReqRulesArrayOutputWithContext(context.Background())
}

func (i NsxtAlbVirtualServiceHttpReqRulesArray) ToNsxtAlbVirtualServiceHttpReqRulesArrayOutputWithContext(ctx context.Context) NsxtAlbVirtualServiceHttpReqRulesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxtAlbVirtualServiceHttpReqRulesArrayOutput)
}

// NsxtAlbVirtualServiceHttpReqRulesMapInput is an input type that accepts NsxtAlbVirtualServiceHttpReqRulesMap and NsxtAlbVirtualServiceHttpReqRulesMapOutput values.
// You can construct a concrete instance of `NsxtAlbVirtualServiceHttpReqRulesMapInput` via:
//
//	NsxtAlbVirtualServiceHttpReqRulesMap{ "key": NsxtAlbVirtualServiceHttpReqRulesArgs{...} }
type NsxtAlbVirtualServiceHttpReqRulesMapInput interface {
	pulumi.Input

	ToNsxtAlbVirtualServiceHttpReqRulesMapOutput() NsxtAlbVirtualServiceHttpReqRulesMapOutput
	ToNsxtAlbVirtualServiceHttpReqRulesMapOutputWithContext(context.Context) NsxtAlbVirtualServiceHttpReqRulesMapOutput
}

type NsxtAlbVirtualServiceHttpReqRulesMap map[string]NsxtAlbVirtualServiceHttpReqRulesInput

func (NsxtAlbVirtualServiceHttpReqRulesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NsxtAlbVirtualServiceHttpReqRules)(nil)).Elem()
}

func (i NsxtAlbVirtualServiceHttpReqRulesMap) ToNsxtAlbVirtualServiceHttpReqRulesMapOutput() NsxtAlbVirtualServiceHttpReqRulesMapOutput {
	return i.ToNsxtAlbVirtualServiceHttpReqRulesMapOutputWithContext(context.Background())
}

func (i NsxtAlbVirtualServiceHttpReqRulesMap) ToNsxtAlbVirtualServiceHttpReqRulesMapOutputWithContext(ctx context.Context) NsxtAlbVirtualServiceHttpReqRulesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxtAlbVirtualServiceHttpReqRulesMapOutput)
}

type NsxtAlbVirtualServiceHttpReqRulesOutput struct{ *pulumi.OutputState }

func (NsxtAlbVirtualServiceHttpReqRulesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NsxtAlbVirtualServiceHttpReqRules)(nil)).Elem()
}

func (o NsxtAlbVirtualServiceHttpReqRulesOutput) ToNsxtAlbVirtualServiceHttpReqRulesOutput() NsxtAlbVirtualServiceHttpReqRulesOutput {
	return o
}

func (o NsxtAlbVirtualServiceHttpReqRulesOutput) ToNsxtAlbVirtualServiceHttpReqRulesOutputWithContext(ctx context.Context) NsxtAlbVirtualServiceHttpReqRulesOutput {
	return o
}

// One or more rule blocks with matching criteria and actions.
//
// <a id="rule-block"></a>
func (o NsxtAlbVirtualServiceHttpReqRulesOutput) Rules() NsxtAlbVirtualServiceHttpReqRulesRuleArrayOutput {
	return o.ApplyT(func(v *NsxtAlbVirtualServiceHttpReqRules) NsxtAlbVirtualServiceHttpReqRulesRuleArrayOutput {
		return v.Rules
	}).(NsxtAlbVirtualServiceHttpReqRulesRuleArrayOutput)
}

// An ID of existing ALB Virtual Service.
func (o NsxtAlbVirtualServiceHttpReqRulesOutput) VirtualServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxtAlbVirtualServiceHttpReqRules) pulumi.StringOutput { return v.VirtualServiceId }).(pulumi.StringOutput)
}

type NsxtAlbVirtualServiceHttpReqRulesArrayOutput struct{ *pulumi.OutputState }

func (NsxtAlbVirtualServiceHttpReqRulesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NsxtAlbVirtualServiceHttpReqRules)(nil)).Elem()
}

func (o NsxtAlbVirtualServiceHttpReqRulesArrayOutput) ToNsxtAlbVirtualServiceHttpReqRulesArrayOutput() NsxtAlbVirtualServiceHttpReqRulesArrayOutput {
	return o
}

func (o NsxtAlbVirtualServiceHttpReqRulesArrayOutput) ToNsxtAlbVirtualServiceHttpReqRulesArrayOutputWithContext(ctx context.Context) NsxtAlbVirtualServiceHttpReqRulesArrayOutput {
	return o
}

func (o NsxtAlbVirtualServiceHttpReqRulesArrayOutput) Index(i pulumi.IntInput) NsxtAlbVirtualServiceHttpReqRulesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NsxtAlbVirtualServiceHttpReqRules {
		return vs[0].([]*NsxtAlbVirtualServiceHttpReqRules)[vs[1].(int)]
	}).(NsxtAlbVirtualServiceHttpReqRulesOutput)
}

type NsxtAlbVirtualServiceHttpReqRulesMapOutput struct{ *pulumi.OutputState }

func (NsxtAlbVirtualServiceHttpReqRulesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NsxtAlbVirtualServiceHttpReqRules)(nil)).Elem()
}

func (o NsxtAlbVirtualServiceHttpReqRulesMapOutput) ToNsxtAlbVirtualServiceHttpReqRulesMapOutput() NsxtAlbVirtualServiceHttpReqRulesMapOutput {
	return o
}

func (o NsxtAlbVirtualServiceHttpReqRulesMapOutput) ToNsxtAlbVirtualServiceHttpReqRulesMapOutputWithContext(ctx context.Context) NsxtAlbVirtualServiceHttpReqRulesMapOutput {
	return o
}

func (o NsxtAlbVirtualServiceHttpReqRulesMapOutput) MapIndex(k pulumi.StringInput) NsxtAlbVirtualServiceHttpReqRulesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NsxtAlbVirtualServiceHttpReqRules {
		return vs[0].(map[string]*NsxtAlbVirtualServiceHttpReqRules)[vs[1].(string)]
	}).(NsxtAlbVirtualServiceHttpReqRulesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NsxtAlbVirtualServiceHttpReqRulesInput)(nil)).Elem(), &NsxtAlbVirtualServiceHttpReqRules{})
	pulumi.RegisterInputType(reflect.TypeOf((*NsxtAlbVirtualServiceHttpReqRulesArrayInput)(nil)).Elem(), NsxtAlbVirtualServiceHttpReqRulesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NsxtAlbVirtualServiceHttpReqRulesMapInput)(nil)).Elem(), NsxtAlbVirtualServiceHttpReqRulesMap{})
	pulumi.RegisterOutputType(NsxtAlbVirtualServiceHttpReqRulesOutput{})
	pulumi.RegisterOutputType(NsxtAlbVirtualServiceHttpReqRulesArrayOutput{})
	pulumi.RegisterOutputType(NsxtAlbVirtualServiceHttpReqRulesMapOutput{})
}
