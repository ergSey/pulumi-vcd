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

type NetworkIsolatedV2 struct {
	pulumi.CustomResourceState

	// An optional description of the network
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// First DNS server to use.
	Dns1 pulumi.StringPtrOutput `pulumi:"dns1"`
	// Second DNS server to use.
	Dns2 pulumi.StringPtrOutput `pulumi:"dns2"`
	// A FQDN for the virtual machines on this network
	DnsSuffix pulumi.StringPtrOutput `pulumi:"dnsSuffix"`
	// Enables Dual-Stack mode so that one can configure one
	// IPv4 and one IPv6 networks. **Note** In such case *IPv4* addresses must be used in `gateway`,
	// `prefixLength` and `staticIpPool` while *IPv6* addresses in `secondaryGateway`,
	// `secondaryPrefixLength` and `secondaryStaticIpPool` fields.
	DualStackEnabled pulumi.BoolPtrOutput `pulumi:"dualStackEnabled"`
	// The gateway for this network (e.g. 192.168.1.1, 2002:0:0:1234:abcd:ffff:c0a7:121)
	Gateway pulumi.StringOutput `pulumi:"gateway"`
	// Set to `true` if network should allow guest VLAN tagging.
	// Default `false`.
	GuestVlanAllowed pulumi.BoolPtrOutput `pulumi:"guestVlanAllowed"`
	// **NSX-V only.** Defines if this network is shared between multiple VDCs
	// in the Org.  Defaults to `false`.
	IsShared pulumi.BoolOutput `pulumi:"isShared"`
	// Use `metadataEntry` instead. Key value map of metadata to assign to this network. **Not supported** if the network belongs to a VDC Group.
	//
	// Deprecated: Use metadataEntry instead
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// A set of metadata entries to assign. See Metadata section for details.
	MetadataEntries NetworkIsolatedV2MetadataEntryArrayOutput `pulumi:"metadataEntries"`
	// A unique name for the network
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful
	// when connected as sysadmin working across different organisations
	Org pulumi.StringPtrOutput `pulumi:"org"`
	// VDC or VDC Group ID. Always takes precedence over `vdc` fields (in resource
	// and inherited from provider configuration)
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// The prefix length for the new network (e.g. 24 for netmask 255.255.255.0).
	PrefixLength pulumi.IntOutput `pulumi:"prefixLength"`
	// IPv6 gateway *when Dual-Stack mode is enabled*
	SecondaryGateway pulumi.StringPtrOutput `pulumi:"secondaryGateway"`
	// IPv6 prefix length *when Dual-Stack mode is
	// enabled*
	SecondaryPrefixLength pulumi.StringPtrOutput `pulumi:"secondaryPrefixLength"`
	// One or more IPv6 static
	// pools *when Dual-Stack mode is enabled*
	//
	// > When using IPv6, VCD API will expand IP Addresses if they are specified using *double colon*
	// notation and it will cause inconsistent plan. (e.g. `2002::1234:abcd:ffff:c0a6:121` will be
	// converted to `2002:0:0:1234:abcd:ffff:c0a6:121`)
	//
	// <a id="ip-pools"></a>
	SecondaryStaticIpPools NetworkIsolatedV2SecondaryStaticIpPoolArrayOutput `pulumi:"secondaryStaticIpPools"`
	// A range of IPs permitted to be used as static IPs for
	// virtual machines; see IP Pools below for details.
	StaticIpPools NetworkIsolatedV2StaticIpPoolArrayOutput `pulumi:"staticIpPools"`
	// The name of VDC to use. **Deprecated**  in favor of new field
	// `ownerId` which supports VDC and VDC Group IDs.
	//
	// Deprecated: This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
	Vdc pulumi.StringOutput `pulumi:"vdc"`
}

// NewNetworkIsolatedV2 registers a new resource with the given unique name, arguments, and options.
func NewNetworkIsolatedV2(ctx *pulumi.Context,
	name string, args *NetworkIsolatedV2Args, opts ...pulumi.ResourceOption) (*NetworkIsolatedV2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Gateway == nil {
		return nil, errors.New("invalid value for required argument 'Gateway'")
	}
	if args.PrefixLength == nil {
		return nil, errors.New("invalid value for required argument 'PrefixLength'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NetworkIsolatedV2
	err := ctx.RegisterResource("vcd:index/networkIsolatedV2:NetworkIsolatedV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkIsolatedV2 gets an existing NetworkIsolatedV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkIsolatedV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkIsolatedV2State, opts ...pulumi.ResourceOption) (*NetworkIsolatedV2, error) {
	var resource NetworkIsolatedV2
	err := ctx.ReadResource("vcd:index/networkIsolatedV2:NetworkIsolatedV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkIsolatedV2 resources.
type networkIsolatedV2State struct {
	// An optional description of the network
	Description *string `pulumi:"description"`
	// First DNS server to use.
	Dns1 *string `pulumi:"dns1"`
	// Second DNS server to use.
	Dns2 *string `pulumi:"dns2"`
	// A FQDN for the virtual machines on this network
	DnsSuffix *string `pulumi:"dnsSuffix"`
	// Enables Dual-Stack mode so that one can configure one
	// IPv4 and one IPv6 networks. **Note** In such case *IPv4* addresses must be used in `gateway`,
	// `prefixLength` and `staticIpPool` while *IPv6* addresses in `secondaryGateway`,
	// `secondaryPrefixLength` and `secondaryStaticIpPool` fields.
	DualStackEnabled *bool `pulumi:"dualStackEnabled"`
	// The gateway for this network (e.g. 192.168.1.1, 2002:0:0:1234:abcd:ffff:c0a7:121)
	Gateway *string `pulumi:"gateway"`
	// Set to `true` if network should allow guest VLAN tagging.
	// Default `false`.
	GuestVlanAllowed *bool `pulumi:"guestVlanAllowed"`
	// **NSX-V only.** Defines if this network is shared between multiple VDCs
	// in the Org.  Defaults to `false`.
	IsShared *bool `pulumi:"isShared"`
	// Use `metadataEntry` instead. Key value map of metadata to assign to this network. **Not supported** if the network belongs to a VDC Group.
	//
	// Deprecated: Use metadataEntry instead
	Metadata map[string]string `pulumi:"metadata"`
	// A set of metadata entries to assign. See Metadata section for details.
	MetadataEntries []NetworkIsolatedV2MetadataEntry `pulumi:"metadataEntries"`
	// A unique name for the network
	Name *string `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful
	// when connected as sysadmin working across different organisations
	Org *string `pulumi:"org"`
	// VDC or VDC Group ID. Always takes precedence over `vdc` fields (in resource
	// and inherited from provider configuration)
	OwnerId *string `pulumi:"ownerId"`
	// The prefix length for the new network (e.g. 24 for netmask 255.255.255.0).
	PrefixLength *int `pulumi:"prefixLength"`
	// IPv6 gateway *when Dual-Stack mode is enabled*
	SecondaryGateway *string `pulumi:"secondaryGateway"`
	// IPv6 prefix length *when Dual-Stack mode is
	// enabled*
	SecondaryPrefixLength *string `pulumi:"secondaryPrefixLength"`
	// One or more IPv6 static
	// pools *when Dual-Stack mode is enabled*
	//
	// > When using IPv6, VCD API will expand IP Addresses if they are specified using *double colon*
	// notation and it will cause inconsistent plan. (e.g. `2002::1234:abcd:ffff:c0a6:121` will be
	// converted to `2002:0:0:1234:abcd:ffff:c0a6:121`)
	//
	// <a id="ip-pools"></a>
	SecondaryStaticIpPools []NetworkIsolatedV2SecondaryStaticIpPool `pulumi:"secondaryStaticIpPools"`
	// A range of IPs permitted to be used as static IPs for
	// virtual machines; see IP Pools below for details.
	StaticIpPools []NetworkIsolatedV2StaticIpPool `pulumi:"staticIpPools"`
	// The name of VDC to use. **Deprecated**  in favor of new field
	// `ownerId` which supports VDC and VDC Group IDs.
	//
	// Deprecated: This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
	Vdc *string `pulumi:"vdc"`
}

type NetworkIsolatedV2State struct {
	// An optional description of the network
	Description pulumi.StringPtrInput
	// First DNS server to use.
	Dns1 pulumi.StringPtrInput
	// Second DNS server to use.
	Dns2 pulumi.StringPtrInput
	// A FQDN for the virtual machines on this network
	DnsSuffix pulumi.StringPtrInput
	// Enables Dual-Stack mode so that one can configure one
	// IPv4 and one IPv6 networks. **Note** In such case *IPv4* addresses must be used in `gateway`,
	// `prefixLength` and `staticIpPool` while *IPv6* addresses in `secondaryGateway`,
	// `secondaryPrefixLength` and `secondaryStaticIpPool` fields.
	DualStackEnabled pulumi.BoolPtrInput
	// The gateway for this network (e.g. 192.168.1.1, 2002:0:0:1234:abcd:ffff:c0a7:121)
	Gateway pulumi.StringPtrInput
	// Set to `true` if network should allow guest VLAN tagging.
	// Default `false`.
	GuestVlanAllowed pulumi.BoolPtrInput
	// **NSX-V only.** Defines if this network is shared between multiple VDCs
	// in the Org.  Defaults to `false`.
	IsShared pulumi.BoolPtrInput
	// Use `metadataEntry` instead. Key value map of metadata to assign to this network. **Not supported** if the network belongs to a VDC Group.
	//
	// Deprecated: Use metadataEntry instead
	Metadata pulumi.StringMapInput
	// A set of metadata entries to assign. See Metadata section for details.
	MetadataEntries NetworkIsolatedV2MetadataEntryArrayInput
	// A unique name for the network
	Name pulumi.StringPtrInput
	// The name of organization to use, optional if defined at provider level. Useful
	// when connected as sysadmin working across different organisations
	Org pulumi.StringPtrInput
	// VDC or VDC Group ID. Always takes precedence over `vdc` fields (in resource
	// and inherited from provider configuration)
	OwnerId pulumi.StringPtrInput
	// The prefix length for the new network (e.g. 24 for netmask 255.255.255.0).
	PrefixLength pulumi.IntPtrInput
	// IPv6 gateway *when Dual-Stack mode is enabled*
	SecondaryGateway pulumi.StringPtrInput
	// IPv6 prefix length *when Dual-Stack mode is
	// enabled*
	SecondaryPrefixLength pulumi.StringPtrInput
	// One or more IPv6 static
	// pools *when Dual-Stack mode is enabled*
	//
	// > When using IPv6, VCD API will expand IP Addresses if they are specified using *double colon*
	// notation and it will cause inconsistent plan. (e.g. `2002::1234:abcd:ffff:c0a6:121` will be
	// converted to `2002:0:0:1234:abcd:ffff:c0a6:121`)
	//
	// <a id="ip-pools"></a>
	SecondaryStaticIpPools NetworkIsolatedV2SecondaryStaticIpPoolArrayInput
	// A range of IPs permitted to be used as static IPs for
	// virtual machines; see IP Pools below for details.
	StaticIpPools NetworkIsolatedV2StaticIpPoolArrayInput
	// The name of VDC to use. **Deprecated**  in favor of new field
	// `ownerId` which supports VDC and VDC Group IDs.
	//
	// Deprecated: This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
	Vdc pulumi.StringPtrInput
}

func (NetworkIsolatedV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*networkIsolatedV2State)(nil)).Elem()
}

type networkIsolatedV2Args struct {
	// An optional description of the network
	Description *string `pulumi:"description"`
	// First DNS server to use.
	Dns1 *string `pulumi:"dns1"`
	// Second DNS server to use.
	Dns2 *string `pulumi:"dns2"`
	// A FQDN for the virtual machines on this network
	DnsSuffix *string `pulumi:"dnsSuffix"`
	// Enables Dual-Stack mode so that one can configure one
	// IPv4 and one IPv6 networks. **Note** In such case *IPv4* addresses must be used in `gateway`,
	// `prefixLength` and `staticIpPool` while *IPv6* addresses in `secondaryGateway`,
	// `secondaryPrefixLength` and `secondaryStaticIpPool` fields.
	DualStackEnabled *bool `pulumi:"dualStackEnabled"`
	// The gateway for this network (e.g. 192.168.1.1, 2002:0:0:1234:abcd:ffff:c0a7:121)
	Gateway string `pulumi:"gateway"`
	// Set to `true` if network should allow guest VLAN tagging.
	// Default `false`.
	GuestVlanAllowed *bool `pulumi:"guestVlanAllowed"`
	// **NSX-V only.** Defines if this network is shared between multiple VDCs
	// in the Org.  Defaults to `false`.
	IsShared *bool `pulumi:"isShared"`
	// Use `metadataEntry` instead. Key value map of metadata to assign to this network. **Not supported** if the network belongs to a VDC Group.
	//
	// Deprecated: Use metadataEntry instead
	Metadata map[string]string `pulumi:"metadata"`
	// A set of metadata entries to assign. See Metadata section for details.
	MetadataEntries []NetworkIsolatedV2MetadataEntry `pulumi:"metadataEntries"`
	// A unique name for the network
	Name *string `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful
	// when connected as sysadmin working across different organisations
	Org *string `pulumi:"org"`
	// VDC or VDC Group ID. Always takes precedence over `vdc` fields (in resource
	// and inherited from provider configuration)
	OwnerId *string `pulumi:"ownerId"`
	// The prefix length for the new network (e.g. 24 for netmask 255.255.255.0).
	PrefixLength int `pulumi:"prefixLength"`
	// IPv6 gateway *when Dual-Stack mode is enabled*
	SecondaryGateway *string `pulumi:"secondaryGateway"`
	// IPv6 prefix length *when Dual-Stack mode is
	// enabled*
	SecondaryPrefixLength *string `pulumi:"secondaryPrefixLength"`
	// One or more IPv6 static
	// pools *when Dual-Stack mode is enabled*
	//
	// > When using IPv6, VCD API will expand IP Addresses if they are specified using *double colon*
	// notation and it will cause inconsistent plan. (e.g. `2002::1234:abcd:ffff:c0a6:121` will be
	// converted to `2002:0:0:1234:abcd:ffff:c0a6:121`)
	//
	// <a id="ip-pools"></a>
	SecondaryStaticIpPools []NetworkIsolatedV2SecondaryStaticIpPool `pulumi:"secondaryStaticIpPools"`
	// A range of IPs permitted to be used as static IPs for
	// virtual machines; see IP Pools below for details.
	StaticIpPools []NetworkIsolatedV2StaticIpPool `pulumi:"staticIpPools"`
	// The name of VDC to use. **Deprecated**  in favor of new field
	// `ownerId` which supports VDC and VDC Group IDs.
	//
	// Deprecated: This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
	Vdc *string `pulumi:"vdc"`
}

// The set of arguments for constructing a NetworkIsolatedV2 resource.
type NetworkIsolatedV2Args struct {
	// An optional description of the network
	Description pulumi.StringPtrInput
	// First DNS server to use.
	Dns1 pulumi.StringPtrInput
	// Second DNS server to use.
	Dns2 pulumi.StringPtrInput
	// A FQDN for the virtual machines on this network
	DnsSuffix pulumi.StringPtrInput
	// Enables Dual-Stack mode so that one can configure one
	// IPv4 and one IPv6 networks. **Note** In such case *IPv4* addresses must be used in `gateway`,
	// `prefixLength` and `staticIpPool` while *IPv6* addresses in `secondaryGateway`,
	// `secondaryPrefixLength` and `secondaryStaticIpPool` fields.
	DualStackEnabled pulumi.BoolPtrInput
	// The gateway for this network (e.g. 192.168.1.1, 2002:0:0:1234:abcd:ffff:c0a7:121)
	Gateway pulumi.StringInput
	// Set to `true` if network should allow guest VLAN tagging.
	// Default `false`.
	GuestVlanAllowed pulumi.BoolPtrInput
	// **NSX-V only.** Defines if this network is shared between multiple VDCs
	// in the Org.  Defaults to `false`.
	IsShared pulumi.BoolPtrInput
	// Use `metadataEntry` instead. Key value map of metadata to assign to this network. **Not supported** if the network belongs to a VDC Group.
	//
	// Deprecated: Use metadataEntry instead
	Metadata pulumi.StringMapInput
	// A set of metadata entries to assign. See Metadata section for details.
	MetadataEntries NetworkIsolatedV2MetadataEntryArrayInput
	// A unique name for the network
	Name pulumi.StringPtrInput
	// The name of organization to use, optional if defined at provider level. Useful
	// when connected as sysadmin working across different organisations
	Org pulumi.StringPtrInput
	// VDC or VDC Group ID. Always takes precedence over `vdc` fields (in resource
	// and inherited from provider configuration)
	OwnerId pulumi.StringPtrInput
	// The prefix length for the new network (e.g. 24 for netmask 255.255.255.0).
	PrefixLength pulumi.IntInput
	// IPv6 gateway *when Dual-Stack mode is enabled*
	SecondaryGateway pulumi.StringPtrInput
	// IPv6 prefix length *when Dual-Stack mode is
	// enabled*
	SecondaryPrefixLength pulumi.StringPtrInput
	// One or more IPv6 static
	// pools *when Dual-Stack mode is enabled*
	//
	// > When using IPv6, VCD API will expand IP Addresses if they are specified using *double colon*
	// notation and it will cause inconsistent plan. (e.g. `2002::1234:abcd:ffff:c0a6:121` will be
	// converted to `2002:0:0:1234:abcd:ffff:c0a6:121`)
	//
	// <a id="ip-pools"></a>
	SecondaryStaticIpPools NetworkIsolatedV2SecondaryStaticIpPoolArrayInput
	// A range of IPs permitted to be used as static IPs for
	// virtual machines; see IP Pools below for details.
	StaticIpPools NetworkIsolatedV2StaticIpPoolArrayInput
	// The name of VDC to use. **Deprecated**  in favor of new field
	// `ownerId` which supports VDC and VDC Group IDs.
	//
	// Deprecated: This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
	Vdc pulumi.StringPtrInput
}

func (NetworkIsolatedV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*networkIsolatedV2Args)(nil)).Elem()
}

type NetworkIsolatedV2Input interface {
	pulumi.Input

	ToNetworkIsolatedV2Output() NetworkIsolatedV2Output
	ToNetworkIsolatedV2OutputWithContext(ctx context.Context) NetworkIsolatedV2Output
}

func (*NetworkIsolatedV2) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkIsolatedV2)(nil)).Elem()
}

func (i *NetworkIsolatedV2) ToNetworkIsolatedV2Output() NetworkIsolatedV2Output {
	return i.ToNetworkIsolatedV2OutputWithContext(context.Background())
}

func (i *NetworkIsolatedV2) ToNetworkIsolatedV2OutputWithContext(ctx context.Context) NetworkIsolatedV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkIsolatedV2Output)
}

// NetworkIsolatedV2ArrayInput is an input type that accepts NetworkIsolatedV2Array and NetworkIsolatedV2ArrayOutput values.
// You can construct a concrete instance of `NetworkIsolatedV2ArrayInput` via:
//
//	NetworkIsolatedV2Array{ NetworkIsolatedV2Args{...} }
type NetworkIsolatedV2ArrayInput interface {
	pulumi.Input

	ToNetworkIsolatedV2ArrayOutput() NetworkIsolatedV2ArrayOutput
	ToNetworkIsolatedV2ArrayOutputWithContext(context.Context) NetworkIsolatedV2ArrayOutput
}

type NetworkIsolatedV2Array []NetworkIsolatedV2Input

func (NetworkIsolatedV2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkIsolatedV2)(nil)).Elem()
}

func (i NetworkIsolatedV2Array) ToNetworkIsolatedV2ArrayOutput() NetworkIsolatedV2ArrayOutput {
	return i.ToNetworkIsolatedV2ArrayOutputWithContext(context.Background())
}

func (i NetworkIsolatedV2Array) ToNetworkIsolatedV2ArrayOutputWithContext(ctx context.Context) NetworkIsolatedV2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkIsolatedV2ArrayOutput)
}

// NetworkIsolatedV2MapInput is an input type that accepts NetworkIsolatedV2Map and NetworkIsolatedV2MapOutput values.
// You can construct a concrete instance of `NetworkIsolatedV2MapInput` via:
//
//	NetworkIsolatedV2Map{ "key": NetworkIsolatedV2Args{...} }
type NetworkIsolatedV2MapInput interface {
	pulumi.Input

	ToNetworkIsolatedV2MapOutput() NetworkIsolatedV2MapOutput
	ToNetworkIsolatedV2MapOutputWithContext(context.Context) NetworkIsolatedV2MapOutput
}

type NetworkIsolatedV2Map map[string]NetworkIsolatedV2Input

func (NetworkIsolatedV2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkIsolatedV2)(nil)).Elem()
}

func (i NetworkIsolatedV2Map) ToNetworkIsolatedV2MapOutput() NetworkIsolatedV2MapOutput {
	return i.ToNetworkIsolatedV2MapOutputWithContext(context.Background())
}

func (i NetworkIsolatedV2Map) ToNetworkIsolatedV2MapOutputWithContext(ctx context.Context) NetworkIsolatedV2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkIsolatedV2MapOutput)
}

type NetworkIsolatedV2Output struct{ *pulumi.OutputState }

func (NetworkIsolatedV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkIsolatedV2)(nil)).Elem()
}

func (o NetworkIsolatedV2Output) ToNetworkIsolatedV2Output() NetworkIsolatedV2Output {
	return o
}

func (o NetworkIsolatedV2Output) ToNetworkIsolatedV2OutputWithContext(ctx context.Context) NetworkIsolatedV2Output {
	return o
}

// An optional description of the network
func (o NetworkIsolatedV2Output) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkIsolatedV2) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// First DNS server to use.
func (o NetworkIsolatedV2Output) Dns1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkIsolatedV2) pulumi.StringPtrOutput { return v.Dns1 }).(pulumi.StringPtrOutput)
}

// Second DNS server to use.
func (o NetworkIsolatedV2Output) Dns2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkIsolatedV2) pulumi.StringPtrOutput { return v.Dns2 }).(pulumi.StringPtrOutput)
}

// A FQDN for the virtual machines on this network
func (o NetworkIsolatedV2Output) DnsSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkIsolatedV2) pulumi.StringPtrOutput { return v.DnsSuffix }).(pulumi.StringPtrOutput)
}

// Enables Dual-Stack mode so that one can configure one
// IPv4 and one IPv6 networks. **Note** In such case *IPv4* addresses must be used in `gateway`,
// `prefixLength` and `staticIpPool` while *IPv6* addresses in `secondaryGateway`,
// `secondaryPrefixLength` and `secondaryStaticIpPool` fields.
func (o NetworkIsolatedV2Output) DualStackEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NetworkIsolatedV2) pulumi.BoolPtrOutput { return v.DualStackEnabled }).(pulumi.BoolPtrOutput)
}

// The gateway for this network (e.g. 192.168.1.1, 2002:0:0:1234:abcd:ffff:c0a7:121)
func (o NetworkIsolatedV2Output) Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkIsolatedV2) pulumi.StringOutput { return v.Gateway }).(pulumi.StringOutput)
}

// Set to `true` if network should allow guest VLAN tagging.
// Default `false`.
func (o NetworkIsolatedV2Output) GuestVlanAllowed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NetworkIsolatedV2) pulumi.BoolPtrOutput { return v.GuestVlanAllowed }).(pulumi.BoolPtrOutput)
}

// **NSX-V only.** Defines if this network is shared between multiple VDCs
// in the Org.  Defaults to `false`.
func (o NetworkIsolatedV2Output) IsShared() pulumi.BoolOutput {
	return o.ApplyT(func(v *NetworkIsolatedV2) pulumi.BoolOutput { return v.IsShared }).(pulumi.BoolOutput)
}

// Use `metadataEntry` instead. Key value map of metadata to assign to this network. **Not supported** if the network belongs to a VDC Group.
//
// Deprecated: Use metadataEntry instead
func (o NetworkIsolatedV2Output) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkIsolatedV2) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

// A set of metadata entries to assign. See Metadata section for details.
func (o NetworkIsolatedV2Output) MetadataEntries() NetworkIsolatedV2MetadataEntryArrayOutput {
	return o.ApplyT(func(v *NetworkIsolatedV2) NetworkIsolatedV2MetadataEntryArrayOutput { return v.MetadataEntries }).(NetworkIsolatedV2MetadataEntryArrayOutput)
}

// A unique name for the network
func (o NetworkIsolatedV2Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkIsolatedV2) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of organization to use, optional if defined at provider level. Useful
// when connected as sysadmin working across different organisations
func (o NetworkIsolatedV2Output) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkIsolatedV2) pulumi.StringPtrOutput { return v.Org }).(pulumi.StringPtrOutput)
}

// VDC or VDC Group ID. Always takes precedence over `vdc` fields (in resource
// and inherited from provider configuration)
func (o NetworkIsolatedV2Output) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkIsolatedV2) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

// The prefix length for the new network (e.g. 24 for netmask 255.255.255.0).
func (o NetworkIsolatedV2Output) PrefixLength() pulumi.IntOutput {
	return o.ApplyT(func(v *NetworkIsolatedV2) pulumi.IntOutput { return v.PrefixLength }).(pulumi.IntOutput)
}

// IPv6 gateway *when Dual-Stack mode is enabled*
func (o NetworkIsolatedV2Output) SecondaryGateway() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkIsolatedV2) pulumi.StringPtrOutput { return v.SecondaryGateway }).(pulumi.StringPtrOutput)
}

// IPv6 prefix length *when Dual-Stack mode is
// enabled*
func (o NetworkIsolatedV2Output) SecondaryPrefixLength() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkIsolatedV2) pulumi.StringPtrOutput { return v.SecondaryPrefixLength }).(pulumi.StringPtrOutput)
}

// One or more IPv6 static
// pools *when Dual-Stack mode is enabled*
//
// > When using IPv6, VCD API will expand IP Addresses if they are specified using *double colon*
// notation and it will cause inconsistent plan. (e.g. `2002::1234:abcd:ffff:c0a6:121` will be
// converted to `2002:0:0:1234:abcd:ffff:c0a6:121`)
//
// <a id="ip-pools"></a>
func (o NetworkIsolatedV2Output) SecondaryStaticIpPools() NetworkIsolatedV2SecondaryStaticIpPoolArrayOutput {
	return o.ApplyT(func(v *NetworkIsolatedV2) NetworkIsolatedV2SecondaryStaticIpPoolArrayOutput {
		return v.SecondaryStaticIpPools
	}).(NetworkIsolatedV2SecondaryStaticIpPoolArrayOutput)
}

// A range of IPs permitted to be used as static IPs for
// virtual machines; see IP Pools below for details.
func (o NetworkIsolatedV2Output) StaticIpPools() NetworkIsolatedV2StaticIpPoolArrayOutput {
	return o.ApplyT(func(v *NetworkIsolatedV2) NetworkIsolatedV2StaticIpPoolArrayOutput { return v.StaticIpPools }).(NetworkIsolatedV2StaticIpPoolArrayOutput)
}

// The name of VDC to use. **Deprecated**  in favor of new field
// `ownerId` which supports VDC and VDC Group IDs.
//
// Deprecated: This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
func (o NetworkIsolatedV2Output) Vdc() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkIsolatedV2) pulumi.StringOutput { return v.Vdc }).(pulumi.StringOutput)
}

type NetworkIsolatedV2ArrayOutput struct{ *pulumi.OutputState }

func (NetworkIsolatedV2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkIsolatedV2)(nil)).Elem()
}

func (o NetworkIsolatedV2ArrayOutput) ToNetworkIsolatedV2ArrayOutput() NetworkIsolatedV2ArrayOutput {
	return o
}

func (o NetworkIsolatedV2ArrayOutput) ToNetworkIsolatedV2ArrayOutputWithContext(ctx context.Context) NetworkIsolatedV2ArrayOutput {
	return o
}

func (o NetworkIsolatedV2ArrayOutput) Index(i pulumi.IntInput) NetworkIsolatedV2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NetworkIsolatedV2 {
		return vs[0].([]*NetworkIsolatedV2)[vs[1].(int)]
	}).(NetworkIsolatedV2Output)
}

type NetworkIsolatedV2MapOutput struct{ *pulumi.OutputState }

func (NetworkIsolatedV2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkIsolatedV2)(nil)).Elem()
}

func (o NetworkIsolatedV2MapOutput) ToNetworkIsolatedV2MapOutput() NetworkIsolatedV2MapOutput {
	return o
}

func (o NetworkIsolatedV2MapOutput) ToNetworkIsolatedV2MapOutputWithContext(ctx context.Context) NetworkIsolatedV2MapOutput {
	return o
}

func (o NetworkIsolatedV2MapOutput) MapIndex(k pulumi.StringInput) NetworkIsolatedV2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NetworkIsolatedV2 {
		return vs[0].(map[string]*NetworkIsolatedV2)[vs[1].(string)]
	}).(NetworkIsolatedV2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkIsolatedV2Input)(nil)).Elem(), &NetworkIsolatedV2{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkIsolatedV2ArrayInput)(nil)).Elem(), NetworkIsolatedV2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkIsolatedV2MapInput)(nil)).Elem(), NetworkIsolatedV2Map{})
	pulumi.RegisterOutputType(NetworkIsolatedV2Output{})
	pulumi.RegisterOutputType(NetworkIsolatedV2ArrayOutput{})
	pulumi.RegisterOutputType(NetworkIsolatedV2MapOutput{})
}
