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

type RdeInterface struct {
	pulumi.CustomResourceState

	// The name of the RDE Interface.
	Name pulumi.StringOutput `pulumi:"name"`
	// A unique namespace associated with the RDE Interface. Only alphanumeric characters, underscores and hyphens allowed.
	Nss pulumi.StringOutput `pulumi:"nss"`
	// Specifies if the RDE Interface can be only read.
	Readonly pulumi.BoolOutput `pulumi:"readonly"`
	// The vendor of the RDE Interface. Only alphanumeric characters, underscores and hyphens allowed.
	Vendor pulumi.StringOutput `pulumi:"vendor"`
	// The version of the RDE Interface. Must follow [semantic versioning](https://semver.org/) syntax.
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewRdeInterface registers a new resource with the given unique name, arguments, and options.
func NewRdeInterface(ctx *pulumi.Context,
	name string, args *RdeInterfaceArgs, opts ...pulumi.ResourceOption) (*RdeInterface, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Nss == nil {
		return nil, errors.New("invalid value for required argument 'Nss'")
	}
	if args.Vendor == nil {
		return nil, errors.New("invalid value for required argument 'Vendor'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource RdeInterface
	err := ctx.RegisterResource("vcd:index/rdeInterface:RdeInterface", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRdeInterface gets an existing RdeInterface resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRdeInterface(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RdeInterfaceState, opts ...pulumi.ResourceOption) (*RdeInterface, error) {
	var resource RdeInterface
	err := ctx.ReadResource("vcd:index/rdeInterface:RdeInterface", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RdeInterface resources.
type rdeInterfaceState struct {
	// The name of the RDE Interface.
	Name *string `pulumi:"name"`
	// A unique namespace associated with the RDE Interface. Only alphanumeric characters, underscores and hyphens allowed.
	Nss *string `pulumi:"nss"`
	// Specifies if the RDE Interface can be only read.
	Readonly *bool `pulumi:"readonly"`
	// The vendor of the RDE Interface. Only alphanumeric characters, underscores and hyphens allowed.
	Vendor *string `pulumi:"vendor"`
	// The version of the RDE Interface. Must follow [semantic versioning](https://semver.org/) syntax.
	Version *string `pulumi:"version"`
}

type RdeInterfaceState struct {
	// The name of the RDE Interface.
	Name pulumi.StringPtrInput
	// A unique namespace associated with the RDE Interface. Only alphanumeric characters, underscores and hyphens allowed.
	Nss pulumi.StringPtrInput
	// Specifies if the RDE Interface can be only read.
	Readonly pulumi.BoolPtrInput
	// The vendor of the RDE Interface. Only alphanumeric characters, underscores and hyphens allowed.
	Vendor pulumi.StringPtrInput
	// The version of the RDE Interface. Must follow [semantic versioning](https://semver.org/) syntax.
	Version pulumi.StringPtrInput
}

func (RdeInterfaceState) ElementType() reflect.Type {
	return reflect.TypeOf((*rdeInterfaceState)(nil)).Elem()
}

type rdeInterfaceArgs struct {
	// The name of the RDE Interface.
	Name *string `pulumi:"name"`
	// A unique namespace associated with the RDE Interface. Only alphanumeric characters, underscores and hyphens allowed.
	Nss string `pulumi:"nss"`
	// The vendor of the RDE Interface. Only alphanumeric characters, underscores and hyphens allowed.
	Vendor string `pulumi:"vendor"`
	// The version of the RDE Interface. Must follow [semantic versioning](https://semver.org/) syntax.
	Version string `pulumi:"version"`
}

// The set of arguments for constructing a RdeInterface resource.
type RdeInterfaceArgs struct {
	// The name of the RDE Interface.
	Name pulumi.StringPtrInput
	// A unique namespace associated with the RDE Interface. Only alphanumeric characters, underscores and hyphens allowed.
	Nss pulumi.StringInput
	// The vendor of the RDE Interface. Only alphanumeric characters, underscores and hyphens allowed.
	Vendor pulumi.StringInput
	// The version of the RDE Interface. Must follow [semantic versioning](https://semver.org/) syntax.
	Version pulumi.StringInput
}

func (RdeInterfaceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*rdeInterfaceArgs)(nil)).Elem()
}

type RdeInterfaceInput interface {
	pulumi.Input

	ToRdeInterfaceOutput() RdeInterfaceOutput
	ToRdeInterfaceOutputWithContext(ctx context.Context) RdeInterfaceOutput
}

func (*RdeInterface) ElementType() reflect.Type {
	return reflect.TypeOf((**RdeInterface)(nil)).Elem()
}

func (i *RdeInterface) ToRdeInterfaceOutput() RdeInterfaceOutput {
	return i.ToRdeInterfaceOutputWithContext(context.Background())
}

func (i *RdeInterface) ToRdeInterfaceOutputWithContext(ctx context.Context) RdeInterfaceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdeInterfaceOutput)
}

// RdeInterfaceArrayInput is an input type that accepts RdeInterfaceArray and RdeInterfaceArrayOutput values.
// You can construct a concrete instance of `RdeInterfaceArrayInput` via:
//
//	RdeInterfaceArray{ RdeInterfaceArgs{...} }
type RdeInterfaceArrayInput interface {
	pulumi.Input

	ToRdeInterfaceArrayOutput() RdeInterfaceArrayOutput
	ToRdeInterfaceArrayOutputWithContext(context.Context) RdeInterfaceArrayOutput
}

type RdeInterfaceArray []RdeInterfaceInput

func (RdeInterfaceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RdeInterface)(nil)).Elem()
}

func (i RdeInterfaceArray) ToRdeInterfaceArrayOutput() RdeInterfaceArrayOutput {
	return i.ToRdeInterfaceArrayOutputWithContext(context.Background())
}

func (i RdeInterfaceArray) ToRdeInterfaceArrayOutputWithContext(ctx context.Context) RdeInterfaceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdeInterfaceArrayOutput)
}

// RdeInterfaceMapInput is an input type that accepts RdeInterfaceMap and RdeInterfaceMapOutput values.
// You can construct a concrete instance of `RdeInterfaceMapInput` via:
//
//	RdeInterfaceMap{ "key": RdeInterfaceArgs{...} }
type RdeInterfaceMapInput interface {
	pulumi.Input

	ToRdeInterfaceMapOutput() RdeInterfaceMapOutput
	ToRdeInterfaceMapOutputWithContext(context.Context) RdeInterfaceMapOutput
}

type RdeInterfaceMap map[string]RdeInterfaceInput

func (RdeInterfaceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RdeInterface)(nil)).Elem()
}

func (i RdeInterfaceMap) ToRdeInterfaceMapOutput() RdeInterfaceMapOutput {
	return i.ToRdeInterfaceMapOutputWithContext(context.Background())
}

func (i RdeInterfaceMap) ToRdeInterfaceMapOutputWithContext(ctx context.Context) RdeInterfaceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RdeInterfaceMapOutput)
}

type RdeInterfaceOutput struct{ *pulumi.OutputState }

func (RdeInterfaceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RdeInterface)(nil)).Elem()
}

func (o RdeInterfaceOutput) ToRdeInterfaceOutput() RdeInterfaceOutput {
	return o
}

func (o RdeInterfaceOutput) ToRdeInterfaceOutputWithContext(ctx context.Context) RdeInterfaceOutput {
	return o
}

// The name of the RDE Interface.
func (o RdeInterfaceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RdeInterface) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// A unique namespace associated with the RDE Interface. Only alphanumeric characters, underscores and hyphens allowed.
func (o RdeInterfaceOutput) Nss() pulumi.StringOutput {
	return o.ApplyT(func(v *RdeInterface) pulumi.StringOutput { return v.Nss }).(pulumi.StringOutput)
}

// Specifies if the RDE Interface can be only read.
func (o RdeInterfaceOutput) Readonly() pulumi.BoolOutput {
	return o.ApplyT(func(v *RdeInterface) pulumi.BoolOutput { return v.Readonly }).(pulumi.BoolOutput)
}

// The vendor of the RDE Interface. Only alphanumeric characters, underscores and hyphens allowed.
func (o RdeInterfaceOutput) Vendor() pulumi.StringOutput {
	return o.ApplyT(func(v *RdeInterface) pulumi.StringOutput { return v.Vendor }).(pulumi.StringOutput)
}

// The version of the RDE Interface. Must follow [semantic versioning](https://semver.org/) syntax.
func (o RdeInterfaceOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *RdeInterface) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

type RdeInterfaceArrayOutput struct{ *pulumi.OutputState }

func (RdeInterfaceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*RdeInterface)(nil)).Elem()
}

func (o RdeInterfaceArrayOutput) ToRdeInterfaceArrayOutput() RdeInterfaceArrayOutput {
	return o
}

func (o RdeInterfaceArrayOutput) ToRdeInterfaceArrayOutputWithContext(ctx context.Context) RdeInterfaceArrayOutput {
	return o
}

func (o RdeInterfaceArrayOutput) Index(i pulumi.IntInput) RdeInterfaceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *RdeInterface {
		return vs[0].([]*RdeInterface)[vs[1].(int)]
	}).(RdeInterfaceOutput)
}

type RdeInterfaceMapOutput struct{ *pulumi.OutputState }

func (RdeInterfaceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*RdeInterface)(nil)).Elem()
}

func (o RdeInterfaceMapOutput) ToRdeInterfaceMapOutput() RdeInterfaceMapOutput {
	return o
}

func (o RdeInterfaceMapOutput) ToRdeInterfaceMapOutputWithContext(ctx context.Context) RdeInterfaceMapOutput {
	return o
}

func (o RdeInterfaceMapOutput) MapIndex(k pulumi.StringInput) RdeInterfaceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *RdeInterface {
		return vs[0].(map[string]*RdeInterface)[vs[1].(string)]
	}).(RdeInterfaceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RdeInterfaceInput)(nil)).Elem(), &RdeInterface{})
	pulumi.RegisterInputType(reflect.TypeOf((*RdeInterfaceArrayInput)(nil)).Elem(), RdeInterfaceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*RdeInterfaceMapInput)(nil)).Elem(), RdeInterfaceMap{})
	pulumi.RegisterOutputType(RdeInterfaceOutput{})
	pulumi.RegisterOutputType(RdeInterfaceArrayOutput{})
	pulumi.RegisterOutputType(RdeInterfaceMapOutput{})
}
