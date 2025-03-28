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

type VdcGroup struct {
	pulumi.CustomResourceState

	// Whether this security policy is enabled. `dfwEnabled` must be `true`.
	DefaultPolicyStatus pulumi.BoolOutput `pulumi:"defaultPolicyStatus"`
	// VDC group description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Whether Distributed Firewall is enabled for this VDC group.
	DfwEnabled pulumi.BoolOutput `pulumi:"dfwEnabled"`
	// More detailed error message when VDC group has error status
	ErrorMessage pulumi.StringOutput `pulumi:"errorMessage"`
	// When `true`, will request VCD to force VDC Group deletion. It
	// should clean up child components. Default `false` (VCD may fail removing VDC Group if there are
	// child components remaining). **Note:** when setting it to `true` for existing resource, it will
	// cause a plan change (update), but this will not alter the resource in any way.
	ForceDelete pulumi.BoolPtrOutput `pulumi:"forceDelete"`
	// Status whether local egress is enabled for a universal router belonging to a universal VDC group.
	LocalEgress pulumi.BoolOutput `pulumi:"localEgress"`
	// The name for VDC group
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of used network pool.
	NetworkPoolId pulumi.StringOutput `pulumi:"networkPoolId"`
	// The network provider’s universal id that is backing the universal network pool.
	NetworkPoolUniversalId pulumi.StringOutput `pulumi:"networkPoolUniversalId"`
	// Defines the networking provider backing the VDC group.
	NetworkProviderType pulumi.StringOutput `pulumi:"networkProviderType"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org pulumi.StringPtrOutput `pulumi:"org"`
	// A list of blocks providing organization VDCs that are participating in this group details. See Participating Org VDCs below for details.
	ParticipatingOrgVdcs VdcGroupParticipatingOrgVdcArrayOutput `pulumi:"participatingOrgVdcs"`
	// The list of organization VDCs that are participating in this group. **Note**: `startingVdcId` isn't automatically included in this list.
	ParticipatingVdcIds pulumi.StringArrayOutput `pulumi:"participatingVdcIds"`
	// Marks whether default firewall rule should be
	// removed after activating. Both `dfwEnabled` and `defaultPolicyStatus` must be true. **Note.**
	// This is mainly useful when using
	// [`NsxtDistributedFirewallRule`](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_distributed_firewall_rule)
	// resource as it cannot remove the default rule.
	RemoveDefaultFirewallRule pulumi.BoolPtrOutput `pulumi:"removeDefaultFirewallRule"`
	// With selecting a starting VDC you will be able to create a group in which this VDC can participate. **Note**: `startingVdcId` must be included in `participatingVdcIds` to participate in this group.
	StartingVdcId pulumi.StringOutput `pulumi:"startingVdcId"`
	// The status that the group can be in (e.g. 'SAVING', 'SAVED', 'CONFIGURING', 'REALIZED', 'REALIZATION_FAILED', 'DELETING', 'DELETE_FAILED', 'OBJECT_NOT_FOUND', 'UNCONFIGURED').
	Status pulumi.StringOutput `pulumi:"status"`
	// Defines the group as LOCAL or UNIVERSAL.
	Type pulumi.StringOutput `pulumi:"type"`
	// True means that a VDC group router has been created.
	UniversalNetworkingEnabled pulumi.BoolOutput `pulumi:"universalNetworkingEnabled"`
}

// NewVdcGroup registers a new resource with the given unique name, arguments, and options.
func NewVdcGroup(ctx *pulumi.Context,
	name string, args *VdcGroupArgs, opts ...pulumi.ResourceOption) (*VdcGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ParticipatingVdcIds == nil {
		return nil, errors.New("invalid value for required argument 'ParticipatingVdcIds'")
	}
	if args.StartingVdcId == nil {
		return nil, errors.New("invalid value for required argument 'StartingVdcId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VdcGroup
	err := ctx.RegisterResource("vcd:index/vdcGroup:VdcGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVdcGroup gets an existing VdcGroup resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVdcGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VdcGroupState, opts ...pulumi.ResourceOption) (*VdcGroup, error) {
	var resource VdcGroup
	err := ctx.ReadResource("vcd:index/vdcGroup:VdcGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VdcGroup resources.
type vdcGroupState struct {
	// Whether this security policy is enabled. `dfwEnabled` must be `true`.
	DefaultPolicyStatus *bool `pulumi:"defaultPolicyStatus"`
	// VDC group description
	Description *string `pulumi:"description"`
	// Whether Distributed Firewall is enabled for this VDC group.
	DfwEnabled *bool `pulumi:"dfwEnabled"`
	// More detailed error message when VDC group has error status
	ErrorMessage *string `pulumi:"errorMessage"`
	// When `true`, will request VCD to force VDC Group deletion. It
	// should clean up child components. Default `false` (VCD may fail removing VDC Group if there are
	// child components remaining). **Note:** when setting it to `true` for existing resource, it will
	// cause a plan change (update), but this will not alter the resource in any way.
	ForceDelete *bool `pulumi:"forceDelete"`
	// Status whether local egress is enabled for a universal router belonging to a universal VDC group.
	LocalEgress *bool `pulumi:"localEgress"`
	// The name for VDC group
	Name *string `pulumi:"name"`
	// ID of used network pool.
	NetworkPoolId *string `pulumi:"networkPoolId"`
	// The network provider’s universal id that is backing the universal network pool.
	NetworkPoolUniversalId *string `pulumi:"networkPoolUniversalId"`
	// Defines the networking provider backing the VDC group.
	NetworkProviderType *string `pulumi:"networkProviderType"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `pulumi:"org"`
	// A list of blocks providing organization VDCs that are participating in this group details. See Participating Org VDCs below for details.
	ParticipatingOrgVdcs []VdcGroupParticipatingOrgVdc `pulumi:"participatingOrgVdcs"`
	// The list of organization VDCs that are participating in this group. **Note**: `startingVdcId` isn't automatically included in this list.
	ParticipatingVdcIds []string `pulumi:"participatingVdcIds"`
	// Marks whether default firewall rule should be
	// removed after activating. Both `dfwEnabled` and `defaultPolicyStatus` must be true. **Note.**
	// This is mainly useful when using
	// [`NsxtDistributedFirewallRule`](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_distributed_firewall_rule)
	// resource as it cannot remove the default rule.
	RemoveDefaultFirewallRule *bool `pulumi:"removeDefaultFirewallRule"`
	// With selecting a starting VDC you will be able to create a group in which this VDC can participate. **Note**: `startingVdcId` must be included in `participatingVdcIds` to participate in this group.
	StartingVdcId *string `pulumi:"startingVdcId"`
	// The status that the group can be in (e.g. 'SAVING', 'SAVED', 'CONFIGURING', 'REALIZED', 'REALIZATION_FAILED', 'DELETING', 'DELETE_FAILED', 'OBJECT_NOT_FOUND', 'UNCONFIGURED').
	Status *string `pulumi:"status"`
	// Defines the group as LOCAL or UNIVERSAL.
	Type *string `pulumi:"type"`
	// True means that a VDC group router has been created.
	UniversalNetworkingEnabled *bool `pulumi:"universalNetworkingEnabled"`
}

type VdcGroupState struct {
	// Whether this security policy is enabled. `dfwEnabled` must be `true`.
	DefaultPolicyStatus pulumi.BoolPtrInput
	// VDC group description
	Description pulumi.StringPtrInput
	// Whether Distributed Firewall is enabled for this VDC group.
	DfwEnabled pulumi.BoolPtrInput
	// More detailed error message when VDC group has error status
	ErrorMessage pulumi.StringPtrInput
	// When `true`, will request VCD to force VDC Group deletion. It
	// should clean up child components. Default `false` (VCD may fail removing VDC Group if there are
	// child components remaining). **Note:** when setting it to `true` for existing resource, it will
	// cause a plan change (update), but this will not alter the resource in any way.
	ForceDelete pulumi.BoolPtrInput
	// Status whether local egress is enabled for a universal router belonging to a universal VDC group.
	LocalEgress pulumi.BoolPtrInput
	// The name for VDC group
	Name pulumi.StringPtrInput
	// ID of used network pool.
	NetworkPoolId pulumi.StringPtrInput
	// The network provider’s universal id that is backing the universal network pool.
	NetworkPoolUniversalId pulumi.StringPtrInput
	// Defines the networking provider backing the VDC group.
	NetworkProviderType pulumi.StringPtrInput
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org pulumi.StringPtrInput
	// A list of blocks providing organization VDCs that are participating in this group details. See Participating Org VDCs below for details.
	ParticipatingOrgVdcs VdcGroupParticipatingOrgVdcArrayInput
	// The list of organization VDCs that are participating in this group. **Note**: `startingVdcId` isn't automatically included in this list.
	ParticipatingVdcIds pulumi.StringArrayInput
	// Marks whether default firewall rule should be
	// removed after activating. Both `dfwEnabled` and `defaultPolicyStatus` must be true. **Note.**
	// This is mainly useful when using
	// [`NsxtDistributedFirewallRule`](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_distributed_firewall_rule)
	// resource as it cannot remove the default rule.
	RemoveDefaultFirewallRule pulumi.BoolPtrInput
	// With selecting a starting VDC you will be able to create a group in which this VDC can participate. **Note**: `startingVdcId` must be included in `participatingVdcIds` to participate in this group.
	StartingVdcId pulumi.StringPtrInput
	// The status that the group can be in (e.g. 'SAVING', 'SAVED', 'CONFIGURING', 'REALIZED', 'REALIZATION_FAILED', 'DELETING', 'DELETE_FAILED', 'OBJECT_NOT_FOUND', 'UNCONFIGURED').
	Status pulumi.StringPtrInput
	// Defines the group as LOCAL or UNIVERSAL.
	Type pulumi.StringPtrInput
	// True means that a VDC group router has been created.
	UniversalNetworkingEnabled pulumi.BoolPtrInput
}

func (VdcGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*vdcGroupState)(nil)).Elem()
}

type vdcGroupArgs struct {
	// Whether this security policy is enabled. `dfwEnabled` must be `true`.
	DefaultPolicyStatus *bool `pulumi:"defaultPolicyStatus"`
	// VDC group description
	Description *string `pulumi:"description"`
	// Whether Distributed Firewall is enabled for this VDC group.
	DfwEnabled *bool `pulumi:"dfwEnabled"`
	// When `true`, will request VCD to force VDC Group deletion. It
	// should clean up child components. Default `false` (VCD may fail removing VDC Group if there are
	// child components remaining). **Note:** when setting it to `true` for existing resource, it will
	// cause a plan change (update), but this will not alter the resource in any way.
	ForceDelete *bool `pulumi:"forceDelete"`
	// The name for VDC group
	Name *string `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org *string `pulumi:"org"`
	// The list of organization VDCs that are participating in this group. **Note**: `startingVdcId` isn't automatically included in this list.
	ParticipatingVdcIds []string `pulumi:"participatingVdcIds"`
	// Marks whether default firewall rule should be
	// removed after activating. Both `dfwEnabled` and `defaultPolicyStatus` must be true. **Note.**
	// This is mainly useful when using
	// [`NsxtDistributedFirewallRule`](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_distributed_firewall_rule)
	// resource as it cannot remove the default rule.
	RemoveDefaultFirewallRule *bool `pulumi:"removeDefaultFirewallRule"`
	// With selecting a starting VDC you will be able to create a group in which this VDC can participate. **Note**: `startingVdcId` must be included in `participatingVdcIds` to participate in this group.
	StartingVdcId string `pulumi:"startingVdcId"`
}

// The set of arguments for constructing a VdcGroup resource.
type VdcGroupArgs struct {
	// Whether this security policy is enabled. `dfwEnabled` must be `true`.
	DefaultPolicyStatus pulumi.BoolPtrInput
	// VDC group description
	Description pulumi.StringPtrInput
	// Whether Distributed Firewall is enabled for this VDC group.
	DfwEnabled pulumi.BoolPtrInput
	// When `true`, will request VCD to force VDC Group deletion. It
	// should clean up child components. Default `false` (VCD may fail removing VDC Group if there are
	// child components remaining). **Note:** when setting it to `true` for existing resource, it will
	// cause a plan change (update), but this will not alter the resource in any way.
	ForceDelete pulumi.BoolPtrInput
	// The name for VDC group
	Name pulumi.StringPtrInput
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
	Org pulumi.StringPtrInput
	// The list of organization VDCs that are participating in this group. **Note**: `startingVdcId` isn't automatically included in this list.
	ParticipatingVdcIds pulumi.StringArrayInput
	// Marks whether default firewall rule should be
	// removed after activating. Both `dfwEnabled` and `defaultPolicyStatus` must be true. **Note.**
	// This is mainly useful when using
	// [`NsxtDistributedFirewallRule`](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_distributed_firewall_rule)
	// resource as it cannot remove the default rule.
	RemoveDefaultFirewallRule pulumi.BoolPtrInput
	// With selecting a starting VDC you will be able to create a group in which this VDC can participate. **Note**: `startingVdcId` must be included in `participatingVdcIds` to participate in this group.
	StartingVdcId pulumi.StringInput
}

func (VdcGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vdcGroupArgs)(nil)).Elem()
}

type VdcGroupInput interface {
	pulumi.Input

	ToVdcGroupOutput() VdcGroupOutput
	ToVdcGroupOutputWithContext(ctx context.Context) VdcGroupOutput
}

func (*VdcGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**VdcGroup)(nil)).Elem()
}

func (i *VdcGroup) ToVdcGroupOutput() VdcGroupOutput {
	return i.ToVdcGroupOutputWithContext(context.Background())
}

func (i *VdcGroup) ToVdcGroupOutputWithContext(ctx context.Context) VdcGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VdcGroupOutput)
}

// VdcGroupArrayInput is an input type that accepts VdcGroupArray and VdcGroupArrayOutput values.
// You can construct a concrete instance of `VdcGroupArrayInput` via:
//
//	VdcGroupArray{ VdcGroupArgs{...} }
type VdcGroupArrayInput interface {
	pulumi.Input

	ToVdcGroupArrayOutput() VdcGroupArrayOutput
	ToVdcGroupArrayOutputWithContext(context.Context) VdcGroupArrayOutput
}

type VdcGroupArray []VdcGroupInput

func (VdcGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VdcGroup)(nil)).Elem()
}

func (i VdcGroupArray) ToVdcGroupArrayOutput() VdcGroupArrayOutput {
	return i.ToVdcGroupArrayOutputWithContext(context.Background())
}

func (i VdcGroupArray) ToVdcGroupArrayOutputWithContext(ctx context.Context) VdcGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VdcGroupArrayOutput)
}

// VdcGroupMapInput is an input type that accepts VdcGroupMap and VdcGroupMapOutput values.
// You can construct a concrete instance of `VdcGroupMapInput` via:
//
//	VdcGroupMap{ "key": VdcGroupArgs{...} }
type VdcGroupMapInput interface {
	pulumi.Input

	ToVdcGroupMapOutput() VdcGroupMapOutput
	ToVdcGroupMapOutputWithContext(context.Context) VdcGroupMapOutput
}

type VdcGroupMap map[string]VdcGroupInput

func (VdcGroupMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VdcGroup)(nil)).Elem()
}

func (i VdcGroupMap) ToVdcGroupMapOutput() VdcGroupMapOutput {
	return i.ToVdcGroupMapOutputWithContext(context.Background())
}

func (i VdcGroupMap) ToVdcGroupMapOutputWithContext(ctx context.Context) VdcGroupMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VdcGroupMapOutput)
}

type VdcGroupOutput struct{ *pulumi.OutputState }

func (VdcGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VdcGroup)(nil)).Elem()
}

func (o VdcGroupOutput) ToVdcGroupOutput() VdcGroupOutput {
	return o
}

func (o VdcGroupOutput) ToVdcGroupOutputWithContext(ctx context.Context) VdcGroupOutput {
	return o
}

// Whether this security policy is enabled. `dfwEnabled` must be `true`.
func (o VdcGroupOutput) DefaultPolicyStatus() pulumi.BoolOutput {
	return o.ApplyT(func(v *VdcGroup) pulumi.BoolOutput { return v.DefaultPolicyStatus }).(pulumi.BoolOutput)
}

// VDC group description
func (o VdcGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VdcGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Whether Distributed Firewall is enabled for this VDC group.
func (o VdcGroupOutput) DfwEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *VdcGroup) pulumi.BoolOutput { return v.DfwEnabled }).(pulumi.BoolOutput)
}

// More detailed error message when VDC group has error status
func (o VdcGroupOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *VdcGroup) pulumi.StringOutput { return v.ErrorMessage }).(pulumi.StringOutput)
}

// When `true`, will request VCD to force VDC Group deletion. It
// should clean up child components. Default `false` (VCD may fail removing VDC Group if there are
// child components remaining). **Note:** when setting it to `true` for existing resource, it will
// cause a plan change (update), but this will not alter the resource in any way.
func (o VdcGroupOutput) ForceDelete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VdcGroup) pulumi.BoolPtrOutput { return v.ForceDelete }).(pulumi.BoolPtrOutput)
}

// Status whether local egress is enabled for a universal router belonging to a universal VDC group.
func (o VdcGroupOutput) LocalEgress() pulumi.BoolOutput {
	return o.ApplyT(func(v *VdcGroup) pulumi.BoolOutput { return v.LocalEgress }).(pulumi.BoolOutput)
}

// The name for VDC group
func (o VdcGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VdcGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID of used network pool.
func (o VdcGroupOutput) NetworkPoolId() pulumi.StringOutput {
	return o.ApplyT(func(v *VdcGroup) pulumi.StringOutput { return v.NetworkPoolId }).(pulumi.StringOutput)
}

// The network provider’s universal id that is backing the universal network pool.
func (o VdcGroupOutput) NetworkPoolUniversalId() pulumi.StringOutput {
	return o.ApplyT(func(v *VdcGroup) pulumi.StringOutput { return v.NetworkPoolUniversalId }).(pulumi.StringOutput)
}

// Defines the networking provider backing the VDC group.
func (o VdcGroupOutput) NetworkProviderType() pulumi.StringOutput {
	return o.ApplyT(func(v *VdcGroup) pulumi.StringOutput { return v.NetworkProviderType }).(pulumi.StringOutput)
}

// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
func (o VdcGroupOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VdcGroup) pulumi.StringPtrOutput { return v.Org }).(pulumi.StringPtrOutput)
}

// A list of blocks providing organization VDCs that are participating in this group details. See Participating Org VDCs below for details.
func (o VdcGroupOutput) ParticipatingOrgVdcs() VdcGroupParticipatingOrgVdcArrayOutput {
	return o.ApplyT(func(v *VdcGroup) VdcGroupParticipatingOrgVdcArrayOutput { return v.ParticipatingOrgVdcs }).(VdcGroupParticipatingOrgVdcArrayOutput)
}

// The list of organization VDCs that are participating in this group. **Note**: `startingVdcId` isn't automatically included in this list.
func (o VdcGroupOutput) ParticipatingVdcIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VdcGroup) pulumi.StringArrayOutput { return v.ParticipatingVdcIds }).(pulumi.StringArrayOutput)
}

// Marks whether default firewall rule should be
// removed after activating. Both `dfwEnabled` and `defaultPolicyStatus` must be true. **Note.**
// This is mainly useful when using
// [`NsxtDistributedFirewallRule`](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_distributed_firewall_rule)
// resource as it cannot remove the default rule.
func (o VdcGroupOutput) RemoveDefaultFirewallRule() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VdcGroup) pulumi.BoolPtrOutput { return v.RemoveDefaultFirewallRule }).(pulumi.BoolPtrOutput)
}

// With selecting a starting VDC you will be able to create a group in which this VDC can participate. **Note**: `startingVdcId` must be included in `participatingVdcIds` to participate in this group.
func (o VdcGroupOutput) StartingVdcId() pulumi.StringOutput {
	return o.ApplyT(func(v *VdcGroup) pulumi.StringOutput { return v.StartingVdcId }).(pulumi.StringOutput)
}

// The status that the group can be in (e.g. 'SAVING', 'SAVED', 'CONFIGURING', 'REALIZED', 'REALIZATION_FAILED', 'DELETING', 'DELETE_FAILED', 'OBJECT_NOT_FOUND', 'UNCONFIGURED').
func (o VdcGroupOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *VdcGroup) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

// Defines the group as LOCAL or UNIVERSAL.
func (o VdcGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VdcGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// True means that a VDC group router has been created.
func (o VdcGroupOutput) UniversalNetworkingEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *VdcGroup) pulumi.BoolOutput { return v.UniversalNetworkingEnabled }).(pulumi.BoolOutput)
}

type VdcGroupArrayOutput struct{ *pulumi.OutputState }

func (VdcGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VdcGroup)(nil)).Elem()
}

func (o VdcGroupArrayOutput) ToVdcGroupArrayOutput() VdcGroupArrayOutput {
	return o
}

func (o VdcGroupArrayOutput) ToVdcGroupArrayOutputWithContext(ctx context.Context) VdcGroupArrayOutput {
	return o
}

func (o VdcGroupArrayOutput) Index(i pulumi.IntInput) VdcGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VdcGroup {
		return vs[0].([]*VdcGroup)[vs[1].(int)]
	}).(VdcGroupOutput)
}

type VdcGroupMapOutput struct{ *pulumi.OutputState }

func (VdcGroupMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VdcGroup)(nil)).Elem()
}

func (o VdcGroupMapOutput) ToVdcGroupMapOutput() VdcGroupMapOutput {
	return o
}

func (o VdcGroupMapOutput) ToVdcGroupMapOutputWithContext(ctx context.Context) VdcGroupMapOutput {
	return o
}

func (o VdcGroupMapOutput) MapIndex(k pulumi.StringInput) VdcGroupOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VdcGroup {
		return vs[0].(map[string]*VdcGroup)[vs[1].(string)]
	}).(VdcGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VdcGroupInput)(nil)).Elem(), &VdcGroup{})
	pulumi.RegisterInputType(reflect.TypeOf((*VdcGroupArrayInput)(nil)).Elem(), VdcGroupArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VdcGroupMapInput)(nil)).Elem(), VdcGroupMap{})
	pulumi.RegisterOutputType(VdcGroupOutput{})
	pulumi.RegisterOutputType(VdcGroupArrayOutput{})
	pulumi.RegisterOutputType(VdcGroupMapOutput{})
}
