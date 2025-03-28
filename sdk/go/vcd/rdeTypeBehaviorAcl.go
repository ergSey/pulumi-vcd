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

type RdeTypeBehaviorAcl struct {
	pulumi.CustomResourceState

	// Set of Access Level IDs to associate to the Behavior defined in `behaviorId` argument
	AccessLevelIds pulumi.StringArrayOutput `pulumi:"accessLevelIds"`
	// The ID of either a [RDE Type Behavior](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/rde_type_behavior)
	// or a [RDE Interface Behavior](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/rde_interface_behavior)
	BehaviorId pulumi.StringOutput `pulumi:"behaviorId"`
	// The ID of the RDE Type
	RdeTypeId pulumi.StringOutput `pulumi:"rdeTypeId"`
}

// NewRdeTypeBehaviorAcl registers a new resource with the given unique name, arguments, and options.
func NewRdeTypeBehaviorAcl(ctx *pulumi.Context,
	name string, args *RdeTypeBehaviorAclArgs, opts ...pulumi.ResourceOption) (*RdeTypeBehaviorAcl, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccessLevelIds == nil {
		return nil, errors.New("invalid value for required argument 'AccessLevelIds'")
	}
	if args.BehaviorId == nil {
		return nil, errors.New("invalid value for required argument 'BehaviorId'")
	}
	if args.RdeTypeId == nil {
		return nil, errors.New("invalid value for required argument 'RdeTypeId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RdeTypeBehaviorAcl
	err := ctx.RegisterResource("vcd:index/rdeTypeBehaviorAcl:RdeTypeBehaviorAcl", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRdeTypeBehaviorAcl gets an existing RdeTypeBehaviorAcl resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRdeTypeBehaviorAcl(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RdeTypeBehaviorAclState, opts ...pulumi.ResourceOption) (*RdeTypeBehaviorAcl, error) {
	var resource RdeTypeBehaviorAcl
	err := ctx.ReadResource("vcd:index/rdeTypeBehaviorAcl:RdeTypeBehaviorAcl", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RdeTypeBehaviorAcl resources.
type rdeTypeBehaviorAclState struct {
	// Set of Access Level IDs to associate to the Behavior defined in `behaviorId` argument
	AccessLevelIds []string `pulumi:"accessLevelIds"`
	// The ID of either a [RDE Type Behavior](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/rde_type_behavior)
	// or a [RDE Interface Behavior](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/rde_interface_behavior)
	BehaviorId *string `pulumi:"behaviorId"`
	// The ID of the RDE Type
	RdeTypeId *string `pulumi:"rdeTypeId"`
}

type RdeTypeBehaviorAclState struct {
	// Set of Access Level IDs to associate to the Behavior defined in `behaviorId` argument
	AccessLevelIds pulumi.StringArrayInput
	// The ID of either a [RDE Type Behavior](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/rde_type_behavior)
	// or a [RDE Interface Behavior](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/rde_interface_behavior)
	BehaviorId pulumi.StringPtrInput
	// The ID of the RDE Type
	RdeTypeId pulumi.StringPtrInput
}

func (RdeTypeBehaviorAclState) ElementType() reflect.Type {
	return reflect.TypeOf((*rdeTypeBehaviorAclState)(nil)).Elem()
}

type rdeTypeBehaviorAclArgs struct {
	// Set of Access Level IDs to associate to the Behavior defined in `behaviorId` argument
	AccessLevelIds []string `pulumi:"accessLevelIds"`
	// The ID of either a [RDE Type Behavior](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/rde_type_behavior)
	// or a [RDE Interface Behavior](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/rde_interface_behavior)
	BehaviorId string `pulumi:"behaviorId"`
	// The ID of the RDE Type
	RdeTypeId string `pulumi:"rdeTypeId"`
}

// The set of arguments for constructing a RdeTypeBehaviorAcl resource.
type RdeTypeBehaviorAclArgs struct {
	// Set of Access Level IDs to associate to the Behavior defined in `behaviorId` argument
	AccessLevelIds pulumi.StringArrayInput
	// The ID of either a [RDE Type Behavior](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/rde_type_behavior)
	// or a [RDE Interface Behavior](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/rde_interface_behavior)
	BehaviorId pulumi.StringInput
	// The ID of the RDE Type
	RdeTypeId pulumi.StringInput
}

func (RdeTypeBehaviorAclArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rdeTypeBehaviorAclArgs)(nil)).Elem()
}

type RdeTypeBehaviorAclInput interface {
	pulumi.Input

	ToRdeTypeBehaviorAclOutput() RdeTypeBehaviorAclOutput
	ToRdeTypeBehaviorAclOutputWithContext(ctx context.Context) RdeTypeBehaviorAclOutput
}

func (*RdeTypeBehaviorAcl) ElementType() reflect.Type {
	return reflect.TypeOf((**RdeTypeBehaviorAcl)(nil)).Elem()
}

func (i *RdeTypeBehaviorAcl) ToRdeTypeBehaviorAclOutput() RdeTypeBehaviorAclOutput {
	return i.ToRdeTypeBehaviorAclOutputWithContext(context.Background())
}

func (i *RdeTypeBehaviorAcl) ToRdeTypeBehaviorAclOutputWithContext(ctx context.Context) RdeTypeBehaviorAclOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdeTypeBehaviorAclOutput)
}

// RdeTypeBehaviorAclArrayInput is an input type that accepts RdeTypeBehaviorAclArray and RdeTypeBehaviorAclArrayOutput values.
// You can construct a concrete instance of `RdeTypeBehaviorAclArrayInput` via:
//
//	RdeTypeBehaviorAclArray{ RdeTypeBehaviorAclArgs{...} }
type RdeTypeBehaviorAclArrayInput interface {
	pulumi.Input

	ToRdeTypeBehaviorAclArrayOutput() RdeTypeBehaviorAclArrayOutput
	ToRdeTypeBehaviorAclArrayOutputWithContext(context.Context) RdeTypeBehaviorAclArrayOutput
}

type RdeTypeBehaviorAclArray []RdeTypeBehaviorAclInput

func (RdeTypeBehaviorAclArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RdeTypeBehaviorAcl)(nil)).Elem()
}

func (i RdeTypeBehaviorAclArray) ToRdeTypeBehaviorAclArrayOutput() RdeTypeBehaviorAclArrayOutput {
	return i.ToRdeTypeBehaviorAclArrayOutputWithContext(context.Background())
}

func (i RdeTypeBehaviorAclArray) ToRdeTypeBehaviorAclArrayOutputWithContext(ctx context.Context) RdeTypeBehaviorAclArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdeTypeBehaviorAclArrayOutput)
}

// RdeTypeBehaviorAclMapInput is an input type that accepts RdeTypeBehaviorAclMap and RdeTypeBehaviorAclMapOutput values.
// You can construct a concrete instance of `RdeTypeBehaviorAclMapInput` via:
//
//	RdeTypeBehaviorAclMap{ "key": RdeTypeBehaviorAclArgs{...} }
type RdeTypeBehaviorAclMapInput interface {
	pulumi.Input

	ToRdeTypeBehaviorAclMapOutput() RdeTypeBehaviorAclMapOutput
	ToRdeTypeBehaviorAclMapOutputWithContext(context.Context) RdeTypeBehaviorAclMapOutput
}

type RdeTypeBehaviorAclMap map[string]RdeTypeBehaviorAclInput

func (RdeTypeBehaviorAclMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RdeTypeBehaviorAcl)(nil)).Elem()
}

func (i RdeTypeBehaviorAclMap) ToRdeTypeBehaviorAclMapOutput() RdeTypeBehaviorAclMapOutput {
	return i.ToRdeTypeBehaviorAclMapOutputWithContext(context.Background())
}

func (i RdeTypeBehaviorAclMap) ToRdeTypeBehaviorAclMapOutputWithContext(ctx context.Context) RdeTypeBehaviorAclMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdeTypeBehaviorAclMapOutput)
}

type RdeTypeBehaviorAclOutput struct{ *pulumi.OutputState }

func (RdeTypeBehaviorAclOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RdeTypeBehaviorAcl)(nil)).Elem()
}

func (o RdeTypeBehaviorAclOutput) ToRdeTypeBehaviorAclOutput() RdeTypeBehaviorAclOutput {
	return o
}

func (o RdeTypeBehaviorAclOutput) ToRdeTypeBehaviorAclOutputWithContext(ctx context.Context) RdeTypeBehaviorAclOutput {
	return o
}

// Set of Access Level IDs to associate to the Behavior defined in `behaviorId` argument
func (o RdeTypeBehaviorAclOutput) AccessLevelIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RdeTypeBehaviorAcl) pulumi.StringArrayOutput { return v.AccessLevelIds }).(pulumi.StringArrayOutput)
}

// The ID of either a [RDE Type Behavior](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/rde_type_behavior)
// or a [RDE Interface Behavior](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/rde_interface_behavior)
func (o RdeTypeBehaviorAclOutput) BehaviorId() pulumi.StringOutput {
	return o.ApplyT(func(v *RdeTypeBehaviorAcl) pulumi.StringOutput { return v.BehaviorId }).(pulumi.StringOutput)
}

// The ID of the RDE Type
func (o RdeTypeBehaviorAclOutput) RdeTypeId() pulumi.StringOutput {
	return o.ApplyT(func(v *RdeTypeBehaviorAcl) pulumi.StringOutput { return v.RdeTypeId }).(pulumi.StringOutput)
}

type RdeTypeBehaviorAclArrayOutput struct{ *pulumi.OutputState }

func (RdeTypeBehaviorAclArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RdeTypeBehaviorAcl)(nil)).Elem()
}

func (o RdeTypeBehaviorAclArrayOutput) ToRdeTypeBehaviorAclArrayOutput() RdeTypeBehaviorAclArrayOutput {
	return o
}

func (o RdeTypeBehaviorAclArrayOutput) ToRdeTypeBehaviorAclArrayOutputWithContext(ctx context.Context) RdeTypeBehaviorAclArrayOutput {
	return o
}

func (o RdeTypeBehaviorAclArrayOutput) Index(i pulumi.IntInput) RdeTypeBehaviorAclOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RdeTypeBehaviorAcl {
		return vs[0].([]*RdeTypeBehaviorAcl)[vs[1].(int)]
	}).(RdeTypeBehaviorAclOutput)
}

type RdeTypeBehaviorAclMapOutput struct{ *pulumi.OutputState }

func (RdeTypeBehaviorAclMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RdeTypeBehaviorAcl)(nil)).Elem()
}

func (o RdeTypeBehaviorAclMapOutput) ToRdeTypeBehaviorAclMapOutput() RdeTypeBehaviorAclMapOutput {
	return o
}

func (o RdeTypeBehaviorAclMapOutput) ToRdeTypeBehaviorAclMapOutputWithContext(ctx context.Context) RdeTypeBehaviorAclMapOutput {
	return o
}

func (o RdeTypeBehaviorAclMapOutput) MapIndex(k pulumi.StringInput) RdeTypeBehaviorAclOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RdeTypeBehaviorAcl {
		return vs[0].(map[string]*RdeTypeBehaviorAcl)[vs[1].(string)]
	}).(RdeTypeBehaviorAclOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RdeTypeBehaviorAclInput)(nil)).Elem(), &RdeTypeBehaviorAcl{})
	pulumi.RegisterInputType(reflect.TypeOf((*RdeTypeBehaviorAclArrayInput)(nil)).Elem(), RdeTypeBehaviorAclArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RdeTypeBehaviorAclMapInput)(nil)).Elem(), RdeTypeBehaviorAclMap{})
	pulumi.RegisterOutputType(RdeTypeBehaviorAclOutput{})
	pulumi.RegisterOutputType(RdeTypeBehaviorAclArrayOutput{})
	pulumi.RegisterOutputType(RdeTypeBehaviorAclMapOutput{})
}
