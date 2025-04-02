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

type NetworkPool struct {
	pulumi.CustomResourceState

	// The components used by the network pool. See Backing below for details
	Backing NetworkPoolBackingOutput `pulumi:"backing"`
	// Define how the backing components are considered. It should be one of the following:
	// * `use-explicit-name` (Default) The backing components must be named explicitly;
	// * `use-when-only-one` The automatically selected backing component will be used if there is only one available;
	// * `use-first-available` Use the first available backing component.
	BackingSelectionConstraint pulumi.StringPtrOutput `pulumi:"backingSelectionConstraint"`
	// Description of the network pool
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Unique name of network pool
	Name pulumi.StringOutput `pulumi:"name"`
	// Id of the network provider (either vCenter or NSX-T manager)
	NetworkProviderId pulumi.StringOutput `pulumi:"networkProviderId"`
	// Name of the network provider
	NetworkProviderName pulumi.StringOutput `pulumi:"networkProviderName"`
	// Type of network provider
	NetworkProviderType pulumi.StringOutput `pulumi:"networkProviderType"`
	// Whether the network pool is in promiscuous mode
	PromiscuousMode pulumi.BoolOutput `pulumi:"promiscuousMode"`
	// Status of the network pool
	Status pulumi.StringOutput `pulumi:"status"`
	// Total number of backings
	TotalBackingsCount pulumi.IntOutput `pulumi:"totalBackingsCount"`
	// Type of the network pool (one of `GENEVE`, `VLAN`, `PORTGROUP_BACKED`)
	Type pulumi.StringOutput `pulumi:"type"`
	// Number of used backings
	UsedBackingsCount pulumi.IntOutput `pulumi:"usedBackingsCount"`
}

// NewNetworkPool registers a new resource with the given unique name, arguments, and options.
func NewNetworkPool(ctx *pulumi.Context,
	name string, args *NetworkPoolArgs, opts ...pulumi.ResourceOption) (*NetworkPool, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NetworkProviderId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkProviderId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NetworkPool
	err := ctx.RegisterResource("vcd:index/networkPool:NetworkPool", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkPool gets an existing NetworkPool resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkPool(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkPoolState, opts ...pulumi.ResourceOption) (*NetworkPool, error) {
	var resource NetworkPool
	err := ctx.ReadResource("vcd:index/networkPool:NetworkPool", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkPool resources.
type networkPoolState struct {
	// The components used by the network pool. See Backing below for details
	Backing *NetworkPoolBacking `pulumi:"backing"`
	// Define how the backing components are considered. It should be one of the following:
	// * `use-explicit-name` (Default) The backing components must be named explicitly;
	// * `use-when-only-one` The automatically selected backing component will be used if there is only one available;
	// * `use-first-available` Use the first available backing component.
	BackingSelectionConstraint *string `pulumi:"backingSelectionConstraint"`
	// Description of the network pool
	Description *string `pulumi:"description"`
	// Unique name of network pool
	Name *string `pulumi:"name"`
	// Id of the network provider (either vCenter or NSX-T manager)
	NetworkProviderId *string `pulumi:"networkProviderId"`
	// Name of the network provider
	NetworkProviderName *string `pulumi:"networkProviderName"`
	// Type of network provider
	NetworkProviderType *string `pulumi:"networkProviderType"`
	// Whether the network pool is in promiscuous mode
	PromiscuousMode *bool `pulumi:"promiscuousMode"`
	// Status of the network pool
	Status *string `pulumi:"status"`
	// Total number of backings
	TotalBackingsCount *int `pulumi:"totalBackingsCount"`
	// Type of the network pool (one of `GENEVE`, `VLAN`, `PORTGROUP_BACKED`)
	Type *string `pulumi:"type"`
	// Number of used backings
	UsedBackingsCount *int `pulumi:"usedBackingsCount"`
}

type NetworkPoolState struct {
	// The components used by the network pool. See Backing below for details
	Backing NetworkPoolBackingPtrInput
	// Define how the backing components are considered. It should be one of the following:
	// * `use-explicit-name` (Default) The backing components must be named explicitly;
	// * `use-when-only-one` The automatically selected backing component will be used if there is only one available;
	// * `use-first-available` Use the first available backing component.
	BackingSelectionConstraint pulumi.StringPtrInput
	// Description of the network pool
	Description pulumi.StringPtrInput
	// Unique name of network pool
	Name pulumi.StringPtrInput
	// Id of the network provider (either vCenter or NSX-T manager)
	NetworkProviderId pulumi.StringPtrInput
	// Name of the network provider
	NetworkProviderName pulumi.StringPtrInput
	// Type of network provider
	NetworkProviderType pulumi.StringPtrInput
	// Whether the network pool is in promiscuous mode
	PromiscuousMode pulumi.BoolPtrInput
	// Status of the network pool
	Status pulumi.StringPtrInput
	// Total number of backings
	TotalBackingsCount pulumi.IntPtrInput
	// Type of the network pool (one of `GENEVE`, `VLAN`, `PORTGROUP_BACKED`)
	Type pulumi.StringPtrInput
	// Number of used backings
	UsedBackingsCount pulumi.IntPtrInput
}

func (NetworkPoolState) ElementType() reflect.Type {
	return reflect.TypeOf((*networkPoolState)(nil)).Elem()
}

type networkPoolArgs struct {
	// The components used by the network pool. See Backing below for details
	Backing *NetworkPoolBacking `pulumi:"backing"`
	// Define how the backing components are considered. It should be one of the following:
	// * `use-explicit-name` (Default) The backing components must be named explicitly;
	// * `use-when-only-one` The automatically selected backing component will be used if there is only one available;
	// * `use-first-available` Use the first available backing component.
	BackingSelectionConstraint *string `pulumi:"backingSelectionConstraint"`
	// Description of the network pool
	Description *string `pulumi:"description"`
	// Unique name of network pool
	Name *string `pulumi:"name"`
	// Id of the network provider (either vCenter or NSX-T manager)
	NetworkProviderId string `pulumi:"networkProviderId"`
	// Type of the network pool (one of `GENEVE`, `VLAN`, `PORTGROUP_BACKED`)
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a NetworkPool resource.
type NetworkPoolArgs struct {
	// The components used by the network pool. See Backing below for details
	Backing NetworkPoolBackingPtrInput
	// Define how the backing components are considered. It should be one of the following:
	// * `use-explicit-name` (Default) The backing components must be named explicitly;
	// * `use-when-only-one` The automatically selected backing component will be used if there is only one available;
	// * `use-first-available` Use the first available backing component.
	BackingSelectionConstraint pulumi.StringPtrInput
	// Description of the network pool
	Description pulumi.StringPtrInput
	// Unique name of network pool
	Name pulumi.StringPtrInput
	// Id of the network provider (either vCenter or NSX-T manager)
	NetworkProviderId pulumi.StringInput
	// Type of the network pool (one of `GENEVE`, `VLAN`, `PORTGROUP_BACKED`)
	Type pulumi.StringInput
}

func (NetworkPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*networkPoolArgs)(nil)).Elem()
}

type NetworkPoolInput interface {
	pulumi.Input

	ToNetworkPoolOutput() NetworkPoolOutput
	ToNetworkPoolOutputWithContext(ctx context.Context) NetworkPoolOutput
}

func (*NetworkPool) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkPool)(nil)).Elem()
}

func (i *NetworkPool) ToNetworkPoolOutput() NetworkPoolOutput {
	return i.ToNetworkPoolOutputWithContext(context.Background())
}

func (i *NetworkPool) ToNetworkPoolOutputWithContext(ctx context.Context) NetworkPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPoolOutput)
}

// NetworkPoolArrayInput is an input type that accepts NetworkPoolArray and NetworkPoolArrayOutput values.
// You can construct a concrete instance of `NetworkPoolArrayInput` via:
//
//	NetworkPoolArray{ NetworkPoolArgs{...} }
type NetworkPoolArrayInput interface {
	pulumi.Input

	ToNetworkPoolArrayOutput() NetworkPoolArrayOutput
	ToNetworkPoolArrayOutputWithContext(context.Context) NetworkPoolArrayOutput
}

type NetworkPoolArray []NetworkPoolInput

func (NetworkPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkPool)(nil)).Elem()
}

func (i NetworkPoolArray) ToNetworkPoolArrayOutput() NetworkPoolArrayOutput {
	return i.ToNetworkPoolArrayOutputWithContext(context.Background())
}

func (i NetworkPoolArray) ToNetworkPoolArrayOutputWithContext(ctx context.Context) NetworkPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPoolArrayOutput)
}

// NetworkPoolMapInput is an input type that accepts NetworkPoolMap and NetworkPoolMapOutput values.
// You can construct a concrete instance of `NetworkPoolMapInput` via:
//
//	NetworkPoolMap{ "key": NetworkPoolArgs{...} }
type NetworkPoolMapInput interface {
	pulumi.Input

	ToNetworkPoolMapOutput() NetworkPoolMapOutput
	ToNetworkPoolMapOutputWithContext(context.Context) NetworkPoolMapOutput
}

type NetworkPoolMap map[string]NetworkPoolInput

func (NetworkPoolMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkPool)(nil)).Elem()
}

func (i NetworkPoolMap) ToNetworkPoolMapOutput() NetworkPoolMapOutput {
	return i.ToNetworkPoolMapOutputWithContext(context.Background())
}

func (i NetworkPoolMap) ToNetworkPoolMapOutputWithContext(ctx context.Context) NetworkPoolMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkPoolMapOutput)
}

type NetworkPoolOutput struct{ *pulumi.OutputState }

func (NetworkPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkPool)(nil)).Elem()
}

func (o NetworkPoolOutput) ToNetworkPoolOutput() NetworkPoolOutput {
	return o
}

func (o NetworkPoolOutput) ToNetworkPoolOutputWithContext(ctx context.Context) NetworkPoolOutput {
	return o
}

// The components used by the network pool. See Backing below for details
func (o NetworkPoolOutput) Backing() NetworkPoolBackingOutput {
	return o.ApplyT(func(v *NetworkPool) NetworkPoolBackingOutput { return v.Backing }).(NetworkPoolBackingOutput)
}

// Define how the backing components are considered. It should be one of the following:
// * `use-explicit-name` (Default) The backing components must be named explicitly;
// * `use-when-only-one` The automatically selected backing component will be used if there is only one available;
// * `use-first-available` Use the first available backing component.
func (o NetworkPoolOutput) BackingSelectionConstraint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkPool) pulumi.StringPtrOutput { return v.BackingSelectionConstraint }).(pulumi.StringPtrOutput)
}

// Description of the network pool
func (o NetworkPoolOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkPool) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Unique name of network pool
func (o NetworkPoolOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkPool) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Id of the network provider (either vCenter or NSX-T manager)
func (o NetworkPoolOutput) NetworkProviderId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkPool) pulumi.StringOutput { return v.NetworkProviderId }).(pulumi.StringOutput)
}

// Name of the network provider
func (o NetworkPoolOutput) NetworkProviderName() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkPool) pulumi.StringOutput { return v.NetworkProviderName }).(pulumi.StringOutput)
}

// Type of network provider
func (o NetworkPoolOutput) NetworkProviderType() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkPool) pulumi.StringOutput { return v.NetworkProviderType }).(pulumi.StringOutput)
}

// Whether the network pool is in promiscuous mode
func (o NetworkPoolOutput) PromiscuousMode() pulumi.BoolOutput {
	return o.ApplyT(func(v *NetworkPool) pulumi.BoolOutput { return v.PromiscuousMode }).(pulumi.BoolOutput)
}

// Status of the network pool
func (o NetworkPoolOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkPool) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Total number of backings
func (o NetworkPoolOutput) TotalBackingsCount() pulumi.IntOutput {
	return o.ApplyT(func(v *NetworkPool) pulumi.IntOutput { return v.TotalBackingsCount }).(pulumi.IntOutput)
}

// Type of the network pool (one of `GENEVE`, `VLAN`, `PORTGROUP_BACKED`)
func (o NetworkPoolOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkPool) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Number of used backings
func (o NetworkPoolOutput) UsedBackingsCount() pulumi.IntOutput {
	return o.ApplyT(func(v *NetworkPool) pulumi.IntOutput { return v.UsedBackingsCount }).(pulumi.IntOutput)
}

type NetworkPoolArrayOutput struct{ *pulumi.OutputState }

func (NetworkPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkPool)(nil)).Elem()
}

func (o NetworkPoolArrayOutput) ToNetworkPoolArrayOutput() NetworkPoolArrayOutput {
	return o
}

func (o NetworkPoolArrayOutput) ToNetworkPoolArrayOutputWithContext(ctx context.Context) NetworkPoolArrayOutput {
	return o
}

func (o NetworkPoolArrayOutput) Index(i pulumi.IntInput) NetworkPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NetworkPool {
		return vs[0].([]*NetworkPool)[vs[1].(int)]
	}).(NetworkPoolOutput)
}

type NetworkPoolMapOutput struct{ *pulumi.OutputState }

func (NetworkPoolMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkPool)(nil)).Elem()
}

func (o NetworkPoolMapOutput) ToNetworkPoolMapOutput() NetworkPoolMapOutput {
	return o
}

func (o NetworkPoolMapOutput) ToNetworkPoolMapOutputWithContext(ctx context.Context) NetworkPoolMapOutput {
	return o
}

func (o NetworkPoolMapOutput) MapIndex(k pulumi.StringInput) NetworkPoolOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NetworkPool {
		return vs[0].(map[string]*NetworkPool)[vs[1].(string)]
	}).(NetworkPoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkPoolInput)(nil)).Elem(), &NetworkPool{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkPoolArrayInput)(nil)).Elem(), NetworkPoolArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkPoolMapInput)(nil)).Elem(), NetworkPoolMap{})
	pulumi.RegisterOutputType(NetworkPoolOutput{})
	pulumi.RegisterOutputType(NetworkPoolArrayOutput{})
	pulumi.RegisterOutputType(NetworkPoolMapOutput{})
}
