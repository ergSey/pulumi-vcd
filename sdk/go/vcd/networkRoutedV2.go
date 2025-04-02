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

type NetworkRoutedV2 struct {
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
	// The ID of the Edge Gateway (NSX-V or NSX-T)
	EdgeGatewayId pulumi.StringOutput `pulumi:"edgeGatewayId"`
	// The gateway for this network (e.g. 192.168.1.1, 2002:0:0:1234:abcd:ffff:c0a7:121)
	Gateway pulumi.StringOutput `pulumi:"gateway"`
	// Set to `true` if network should allow guest VLAN tagging.
	// Default `false`.
	GuestVlanAllowed pulumi.BoolPtrOutput `pulumi:"guestVlanAllowed"`
	// An interface for the network. One of `internal` (default),
	// `subinterface`, `distributed`, `nonDistributed` (requires the Edge Gateway to support distributed
	// networks). NSX-T supports only `internal` and `nonDistributed` (*v3.14+*, requires Edge Gateway
	// to have [non-distributed routing
	// enabled](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_edgegateway#non_distributed_routing_enabled)).
	InterfaceType pulumi.StringPtrOutput `pulumi:"interfaceType"`
	// Use `metadataEntry` instead. Key value map of metadata to assign to this network. **Not supported** if the owner edge gateway belongs to a VDC Group.
	//
	// Deprecated: Use metadataEntry instead
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// A set of metadata entries to assign. See Metadata section for details.
	MetadataEntries NetworkRoutedV2MetadataEntryArrayOutput `pulumi:"metadataEntries"`
	// A unique name for the network
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful when
	// connected as sysadmin working across different organisations
	Org pulumi.StringPtrOutput `pulumi:"org"`
	// ID of VDC or VDC Group
	OwnerId pulumi.StringOutput `pulumi:"ownerId"`
	// The prefix length for the new network (e.g. 24 for netmask 255.255.255.0).
	PrefixLength pulumi.IntOutput `pulumi:"prefixLength"`
	// Enables route advertising for
	// this network. **Note** This requires Edge Gateway to use IP Spaces and IP Space *must have* route
	// advertisement
	// enabled.
	RouteAdvertisementEnabled pulumi.BoolPtrOutput `pulumi:"routeAdvertisementEnabled"`
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
	SecondaryStaticIpPools NetworkRoutedV2SecondaryStaticIpPoolArrayOutput `pulumi:"secondaryStaticIpPools"`
	// A range of IPs permitted to be used as static IPs for
	// virtual machines; see IP Pools below for details.
	StaticIpPools NetworkRoutedV2StaticIpPoolArrayOutput `pulumi:"staticIpPools"`
	// The name of VDC to use. *v3.6+* inherits parent VDC or VDC Group
	// from `edgeGatewayId`)
	//
	// Deprecated: 'vdc' is deprecated and ineffective. Routed networks will inherit VDC setting from parent Edge Gateway
	Vdc pulumi.StringOutput `pulumi:"vdc"`
}

// NewNetworkRoutedV2 registers a new resource with the given unique name, arguments, and options.
func NewNetworkRoutedV2(ctx *pulumi.Context,
	name string, args *NetworkRoutedV2Args, opts ...pulumi.ResourceOption) (*NetworkRoutedV2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EdgeGatewayId == nil {
		return nil, errors.New("invalid value for required argument 'EdgeGatewayId'")
	}
	if args.Gateway == nil {
		return nil, errors.New("invalid value for required argument 'Gateway'")
	}
	if args.PrefixLength == nil {
		return nil, errors.New("invalid value for required argument 'PrefixLength'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NetworkRoutedV2
	err := ctx.RegisterResource("vcd:index/networkRoutedV2:NetworkRoutedV2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNetworkRoutedV2 gets an existing NetworkRoutedV2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNetworkRoutedV2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NetworkRoutedV2State, opts ...pulumi.ResourceOption) (*NetworkRoutedV2, error) {
	var resource NetworkRoutedV2
	err := ctx.ReadResource("vcd:index/networkRoutedV2:NetworkRoutedV2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NetworkRoutedV2 resources.
type networkRoutedV2State struct {
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
	// The ID of the Edge Gateway (NSX-V or NSX-T)
	EdgeGatewayId *string `pulumi:"edgeGatewayId"`
	// The gateway for this network (e.g. 192.168.1.1, 2002:0:0:1234:abcd:ffff:c0a7:121)
	Gateway *string `pulumi:"gateway"`
	// Set to `true` if network should allow guest VLAN tagging.
	// Default `false`.
	GuestVlanAllowed *bool `pulumi:"guestVlanAllowed"`
	// An interface for the network. One of `internal` (default),
	// `subinterface`, `distributed`, `nonDistributed` (requires the Edge Gateway to support distributed
	// networks). NSX-T supports only `internal` and `nonDistributed` (*v3.14+*, requires Edge Gateway
	// to have [non-distributed routing
	// enabled](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_edgegateway#non_distributed_routing_enabled)).
	InterfaceType *string `pulumi:"interfaceType"`
	// Use `metadataEntry` instead. Key value map of metadata to assign to this network. **Not supported** if the owner edge gateway belongs to a VDC Group.
	//
	// Deprecated: Use metadataEntry instead
	Metadata map[string]string `pulumi:"metadata"`
	// A set of metadata entries to assign. See Metadata section for details.
	MetadataEntries []NetworkRoutedV2MetadataEntry `pulumi:"metadataEntries"`
	// A unique name for the network
	Name *string `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful when
	// connected as sysadmin working across different organisations
	Org *string `pulumi:"org"`
	// ID of VDC or VDC Group
	OwnerId *string `pulumi:"ownerId"`
	// The prefix length for the new network (e.g. 24 for netmask 255.255.255.0).
	PrefixLength *int `pulumi:"prefixLength"`
	// Enables route advertising for
	// this network. **Note** This requires Edge Gateway to use IP Spaces and IP Space *must have* route
	// advertisement
	// enabled.
	RouteAdvertisementEnabled *bool `pulumi:"routeAdvertisementEnabled"`
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
	SecondaryStaticIpPools []NetworkRoutedV2SecondaryStaticIpPool `pulumi:"secondaryStaticIpPools"`
	// A range of IPs permitted to be used as static IPs for
	// virtual machines; see IP Pools below for details.
	StaticIpPools []NetworkRoutedV2StaticIpPool `pulumi:"staticIpPools"`
	// The name of VDC to use. *v3.6+* inherits parent VDC or VDC Group
	// from `edgeGatewayId`)
	//
	// Deprecated: 'vdc' is deprecated and ineffective. Routed networks will inherit VDC setting from parent Edge Gateway
	Vdc *string `pulumi:"vdc"`
}

type NetworkRoutedV2State struct {
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
	// The ID of the Edge Gateway (NSX-V or NSX-T)
	EdgeGatewayId pulumi.StringPtrInput
	// The gateway for this network (e.g. 192.168.1.1, 2002:0:0:1234:abcd:ffff:c0a7:121)
	Gateway pulumi.StringPtrInput
	// Set to `true` if network should allow guest VLAN tagging.
	// Default `false`.
	GuestVlanAllowed pulumi.BoolPtrInput
	// An interface for the network. One of `internal` (default),
	// `subinterface`, `distributed`, `nonDistributed` (requires the Edge Gateway to support distributed
	// networks). NSX-T supports only `internal` and `nonDistributed` (*v3.14+*, requires Edge Gateway
	// to have [non-distributed routing
	// enabled](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_edgegateway#non_distributed_routing_enabled)).
	InterfaceType pulumi.StringPtrInput
	// Use `metadataEntry` instead. Key value map of metadata to assign to this network. **Not supported** if the owner edge gateway belongs to a VDC Group.
	//
	// Deprecated: Use metadataEntry instead
	Metadata pulumi.StringMapInput
	// A set of metadata entries to assign. See Metadata section for details.
	MetadataEntries NetworkRoutedV2MetadataEntryArrayInput
	// A unique name for the network
	Name pulumi.StringPtrInput
	// The name of organization to use, optional if defined at provider level. Useful when
	// connected as sysadmin working across different organisations
	Org pulumi.StringPtrInput
	// ID of VDC or VDC Group
	OwnerId pulumi.StringPtrInput
	// The prefix length for the new network (e.g. 24 for netmask 255.255.255.0).
	PrefixLength pulumi.IntPtrInput
	// Enables route advertising for
	// this network. **Note** This requires Edge Gateway to use IP Spaces and IP Space *must have* route
	// advertisement
	// enabled.
	RouteAdvertisementEnabled pulumi.BoolPtrInput
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
	SecondaryStaticIpPools NetworkRoutedV2SecondaryStaticIpPoolArrayInput
	// A range of IPs permitted to be used as static IPs for
	// virtual machines; see IP Pools below for details.
	StaticIpPools NetworkRoutedV2StaticIpPoolArrayInput
	// The name of VDC to use. *v3.6+* inherits parent VDC or VDC Group
	// from `edgeGatewayId`)
	//
	// Deprecated: 'vdc' is deprecated and ineffective. Routed networks will inherit VDC setting from parent Edge Gateway
	Vdc pulumi.StringPtrInput
}

func (NetworkRoutedV2State) ElementType() reflect.Type {
	return reflect.TypeOf((*networkRoutedV2State)(nil)).Elem()
}

type networkRoutedV2Args struct {
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
	// The ID of the Edge Gateway (NSX-V or NSX-T)
	EdgeGatewayId string `pulumi:"edgeGatewayId"`
	// The gateway for this network (e.g. 192.168.1.1, 2002:0:0:1234:abcd:ffff:c0a7:121)
	Gateway string `pulumi:"gateway"`
	// Set to `true` if network should allow guest VLAN tagging.
	// Default `false`.
	GuestVlanAllowed *bool `pulumi:"guestVlanAllowed"`
	// An interface for the network. One of `internal` (default),
	// `subinterface`, `distributed`, `nonDistributed` (requires the Edge Gateway to support distributed
	// networks). NSX-T supports only `internal` and `nonDistributed` (*v3.14+*, requires Edge Gateway
	// to have [non-distributed routing
	// enabled](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_edgegateway#non_distributed_routing_enabled)).
	InterfaceType *string `pulumi:"interfaceType"`
	// Use `metadataEntry` instead. Key value map of metadata to assign to this network. **Not supported** if the owner edge gateway belongs to a VDC Group.
	//
	// Deprecated: Use metadataEntry instead
	Metadata map[string]string `pulumi:"metadata"`
	// A set of metadata entries to assign. See Metadata section for details.
	MetadataEntries []NetworkRoutedV2MetadataEntry `pulumi:"metadataEntries"`
	// A unique name for the network
	Name *string `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful when
	// connected as sysadmin working across different organisations
	Org *string `pulumi:"org"`
	// The prefix length for the new network (e.g. 24 for netmask 255.255.255.0).
	PrefixLength int `pulumi:"prefixLength"`
	// Enables route advertising for
	// this network. **Note** This requires Edge Gateway to use IP Spaces and IP Space *must have* route
	// advertisement
	// enabled.
	RouteAdvertisementEnabled *bool `pulumi:"routeAdvertisementEnabled"`
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
	SecondaryStaticIpPools []NetworkRoutedV2SecondaryStaticIpPool `pulumi:"secondaryStaticIpPools"`
	// A range of IPs permitted to be used as static IPs for
	// virtual machines; see IP Pools below for details.
	StaticIpPools []NetworkRoutedV2StaticIpPool `pulumi:"staticIpPools"`
	// The name of VDC to use. *v3.6+* inherits parent VDC or VDC Group
	// from `edgeGatewayId`)
	//
	// Deprecated: 'vdc' is deprecated and ineffective. Routed networks will inherit VDC setting from parent Edge Gateway
	Vdc *string `pulumi:"vdc"`
}

// The set of arguments for constructing a NetworkRoutedV2 resource.
type NetworkRoutedV2Args struct {
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
	// The ID of the Edge Gateway (NSX-V or NSX-T)
	EdgeGatewayId pulumi.StringInput
	// The gateway for this network (e.g. 192.168.1.1, 2002:0:0:1234:abcd:ffff:c0a7:121)
	Gateway pulumi.StringInput
	// Set to `true` if network should allow guest VLAN tagging.
	// Default `false`.
	GuestVlanAllowed pulumi.BoolPtrInput
	// An interface for the network. One of `internal` (default),
	// `subinterface`, `distributed`, `nonDistributed` (requires the Edge Gateway to support distributed
	// networks). NSX-T supports only `internal` and `nonDistributed` (*v3.14+*, requires Edge Gateway
	// to have [non-distributed routing
	// enabled](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_edgegateway#non_distributed_routing_enabled)).
	InterfaceType pulumi.StringPtrInput
	// Use `metadataEntry` instead. Key value map of metadata to assign to this network. **Not supported** if the owner edge gateway belongs to a VDC Group.
	//
	// Deprecated: Use metadataEntry instead
	Metadata pulumi.StringMapInput
	// A set of metadata entries to assign. See Metadata section for details.
	MetadataEntries NetworkRoutedV2MetadataEntryArrayInput
	// A unique name for the network
	Name pulumi.StringPtrInput
	// The name of organization to use, optional if defined at provider level. Useful when
	// connected as sysadmin working across different organisations
	Org pulumi.StringPtrInput
	// The prefix length for the new network (e.g. 24 for netmask 255.255.255.0).
	PrefixLength pulumi.IntInput
	// Enables route advertising for
	// this network. **Note** This requires Edge Gateway to use IP Spaces and IP Space *must have* route
	// advertisement
	// enabled.
	RouteAdvertisementEnabled pulumi.BoolPtrInput
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
	SecondaryStaticIpPools NetworkRoutedV2SecondaryStaticIpPoolArrayInput
	// A range of IPs permitted to be used as static IPs for
	// virtual machines; see IP Pools below for details.
	StaticIpPools NetworkRoutedV2StaticIpPoolArrayInput
	// The name of VDC to use. *v3.6+* inherits parent VDC or VDC Group
	// from `edgeGatewayId`)
	//
	// Deprecated: 'vdc' is deprecated and ineffective. Routed networks will inherit VDC setting from parent Edge Gateway
	Vdc pulumi.StringPtrInput
}

func (NetworkRoutedV2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*networkRoutedV2Args)(nil)).Elem()
}

type NetworkRoutedV2Input interface {
	pulumi.Input

	ToNetworkRoutedV2Output() NetworkRoutedV2Output
	ToNetworkRoutedV2OutputWithContext(ctx context.Context) NetworkRoutedV2Output
}

func (*NetworkRoutedV2) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRoutedV2)(nil)).Elem()
}

func (i *NetworkRoutedV2) ToNetworkRoutedV2Output() NetworkRoutedV2Output {
	return i.ToNetworkRoutedV2OutputWithContext(context.Background())
}

func (i *NetworkRoutedV2) ToNetworkRoutedV2OutputWithContext(ctx context.Context) NetworkRoutedV2Output {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRoutedV2Output)
}

// NetworkRoutedV2ArrayInput is an input type that accepts NetworkRoutedV2Array and NetworkRoutedV2ArrayOutput values.
// You can construct a concrete instance of `NetworkRoutedV2ArrayInput` via:
//
//	NetworkRoutedV2Array{ NetworkRoutedV2Args{...} }
type NetworkRoutedV2ArrayInput interface {
	pulumi.Input

	ToNetworkRoutedV2ArrayOutput() NetworkRoutedV2ArrayOutput
	ToNetworkRoutedV2ArrayOutputWithContext(context.Context) NetworkRoutedV2ArrayOutput
}

type NetworkRoutedV2Array []NetworkRoutedV2Input

func (NetworkRoutedV2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkRoutedV2)(nil)).Elem()
}

func (i NetworkRoutedV2Array) ToNetworkRoutedV2ArrayOutput() NetworkRoutedV2ArrayOutput {
	return i.ToNetworkRoutedV2ArrayOutputWithContext(context.Background())
}

func (i NetworkRoutedV2Array) ToNetworkRoutedV2ArrayOutputWithContext(ctx context.Context) NetworkRoutedV2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRoutedV2ArrayOutput)
}

// NetworkRoutedV2MapInput is an input type that accepts NetworkRoutedV2Map and NetworkRoutedV2MapOutput values.
// You can construct a concrete instance of `NetworkRoutedV2MapInput` via:
//
//	NetworkRoutedV2Map{ "key": NetworkRoutedV2Args{...} }
type NetworkRoutedV2MapInput interface {
	pulumi.Input

	ToNetworkRoutedV2MapOutput() NetworkRoutedV2MapOutput
	ToNetworkRoutedV2MapOutputWithContext(context.Context) NetworkRoutedV2MapOutput
}

type NetworkRoutedV2Map map[string]NetworkRoutedV2Input

func (NetworkRoutedV2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkRoutedV2)(nil)).Elem()
}

func (i NetworkRoutedV2Map) ToNetworkRoutedV2MapOutput() NetworkRoutedV2MapOutput {
	return i.ToNetworkRoutedV2MapOutputWithContext(context.Background())
}

func (i NetworkRoutedV2Map) ToNetworkRoutedV2MapOutputWithContext(ctx context.Context) NetworkRoutedV2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRoutedV2MapOutput)
}

type NetworkRoutedV2Output struct{ *pulumi.OutputState }

func (NetworkRoutedV2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRoutedV2)(nil)).Elem()
}

func (o NetworkRoutedV2Output) ToNetworkRoutedV2Output() NetworkRoutedV2Output {
	return o
}

func (o NetworkRoutedV2Output) ToNetworkRoutedV2OutputWithContext(ctx context.Context) NetworkRoutedV2Output {
	return o
}

// An optional description of the network
func (o NetworkRoutedV2Output) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkRoutedV2) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// First DNS server to use.
func (o NetworkRoutedV2Output) Dns1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkRoutedV2) pulumi.StringPtrOutput { return v.Dns1 }).(pulumi.StringPtrOutput)
}

// Second DNS server to use.
func (o NetworkRoutedV2Output) Dns2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkRoutedV2) pulumi.StringPtrOutput { return v.Dns2 }).(pulumi.StringPtrOutput)
}

// A FQDN for the virtual machines on this network
func (o NetworkRoutedV2Output) DnsSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkRoutedV2) pulumi.StringPtrOutput { return v.DnsSuffix }).(pulumi.StringPtrOutput)
}

// Enables Dual-Stack mode so that one can configure one
// IPv4 and one IPv6 networks. **Note** In such case *IPv4* addresses must be used in `gateway`,
// `prefixLength` and `staticIpPool` while *IPv6* addresses in `secondaryGateway`,
// `secondaryPrefixLength` and `secondaryStaticIpPool` fields.
func (o NetworkRoutedV2Output) DualStackEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NetworkRoutedV2) pulumi.BoolPtrOutput { return v.DualStackEnabled }).(pulumi.BoolPtrOutput)
}

// The ID of the Edge Gateway (NSX-V or NSX-T)
func (o NetworkRoutedV2Output) EdgeGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkRoutedV2) pulumi.StringOutput { return v.EdgeGatewayId }).(pulumi.StringOutput)
}

// The gateway for this network (e.g. 192.168.1.1, 2002:0:0:1234:abcd:ffff:c0a7:121)
func (o NetworkRoutedV2Output) Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkRoutedV2) pulumi.StringOutput { return v.Gateway }).(pulumi.StringOutput)
}

// Set to `true` if network should allow guest VLAN tagging.
// Default `false`.
func (o NetworkRoutedV2Output) GuestVlanAllowed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NetworkRoutedV2) pulumi.BoolPtrOutput { return v.GuestVlanAllowed }).(pulumi.BoolPtrOutput)
}

// An interface for the network. One of `internal` (default),
// `subinterface`, `distributed`, `nonDistributed` (requires the Edge Gateway to support distributed
// networks). NSX-T supports only `internal` and `nonDistributed` (*v3.14+*, requires Edge Gateway
// to have [non-distributed routing
// enabled](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_edgegateway#non_distributed_routing_enabled)).
func (o NetworkRoutedV2Output) InterfaceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkRoutedV2) pulumi.StringPtrOutput { return v.InterfaceType }).(pulumi.StringPtrOutput)
}

// Use `metadataEntry` instead. Key value map of metadata to assign to this network. **Not supported** if the owner edge gateway belongs to a VDC Group.
//
// Deprecated: Use metadataEntry instead
func (o NetworkRoutedV2Output) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *NetworkRoutedV2) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

// A set of metadata entries to assign. See Metadata section for details.
func (o NetworkRoutedV2Output) MetadataEntries() NetworkRoutedV2MetadataEntryArrayOutput {
	return o.ApplyT(func(v *NetworkRoutedV2) NetworkRoutedV2MetadataEntryArrayOutput { return v.MetadataEntries }).(NetworkRoutedV2MetadataEntryArrayOutput)
}

// A unique name for the network
func (o NetworkRoutedV2Output) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkRoutedV2) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of organization to use, optional if defined at provider level. Useful when
// connected as sysadmin working across different organisations
func (o NetworkRoutedV2Output) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkRoutedV2) pulumi.StringPtrOutput { return v.Org }).(pulumi.StringPtrOutput)
}

// ID of VDC or VDC Group
func (o NetworkRoutedV2Output) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkRoutedV2) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

// The prefix length for the new network (e.g. 24 for netmask 255.255.255.0).
func (o NetworkRoutedV2Output) PrefixLength() pulumi.IntOutput {
	return o.ApplyT(func(v *NetworkRoutedV2) pulumi.IntOutput { return v.PrefixLength }).(pulumi.IntOutput)
}

// Enables route advertising for
// this network. **Note** This requires Edge Gateway to use IP Spaces and IP Space *must have* route
// advertisement
// enabled.
func (o NetworkRoutedV2Output) RouteAdvertisementEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NetworkRoutedV2) pulumi.BoolPtrOutput { return v.RouteAdvertisementEnabled }).(pulumi.BoolPtrOutput)
}

// IPv6 gateway *when Dual-Stack mode is enabled*
func (o NetworkRoutedV2Output) SecondaryGateway() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkRoutedV2) pulumi.StringPtrOutput { return v.SecondaryGateway }).(pulumi.StringPtrOutput)
}

// IPv6 prefix length *when Dual-Stack mode is
// enabled*
func (o NetworkRoutedV2Output) SecondaryPrefixLength() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkRoutedV2) pulumi.StringPtrOutput { return v.SecondaryPrefixLength }).(pulumi.StringPtrOutput)
}

// One or more IPv6 static
// pools *when Dual-Stack mode is enabled*
//
// > When using IPv6, VCD API will expand IP Addresses if they are specified using *double colon*
// notation and it will cause inconsistent plan. (e.g. `2002::1234:abcd:ffff:c0a6:121` will be
// converted to `2002:0:0:1234:abcd:ffff:c0a6:121`)
//
// <a id="ip-pools"></a>
func (o NetworkRoutedV2Output) SecondaryStaticIpPools() NetworkRoutedV2SecondaryStaticIpPoolArrayOutput {
	return o.ApplyT(func(v *NetworkRoutedV2) NetworkRoutedV2SecondaryStaticIpPoolArrayOutput {
		return v.SecondaryStaticIpPools
	}).(NetworkRoutedV2SecondaryStaticIpPoolArrayOutput)
}

// A range of IPs permitted to be used as static IPs for
// virtual machines; see IP Pools below for details.
func (o NetworkRoutedV2Output) StaticIpPools() NetworkRoutedV2StaticIpPoolArrayOutput {
	return o.ApplyT(func(v *NetworkRoutedV2) NetworkRoutedV2StaticIpPoolArrayOutput { return v.StaticIpPools }).(NetworkRoutedV2StaticIpPoolArrayOutput)
}

// The name of VDC to use. *v3.6+* inherits parent VDC or VDC Group
// from `edgeGatewayId`)
//
// Deprecated: 'vdc' is deprecated and ineffective. Routed networks will inherit VDC setting from parent Edge Gateway
func (o NetworkRoutedV2Output) Vdc() pulumi.StringOutput {
	return o.ApplyT(func(v *NetworkRoutedV2) pulumi.StringOutput { return v.Vdc }).(pulumi.StringOutput)
}

type NetworkRoutedV2ArrayOutput struct{ *pulumi.OutputState }

func (NetworkRoutedV2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NetworkRoutedV2)(nil)).Elem()
}

func (o NetworkRoutedV2ArrayOutput) ToNetworkRoutedV2ArrayOutput() NetworkRoutedV2ArrayOutput {
	return o
}

func (o NetworkRoutedV2ArrayOutput) ToNetworkRoutedV2ArrayOutputWithContext(ctx context.Context) NetworkRoutedV2ArrayOutput {
	return o
}

func (o NetworkRoutedV2ArrayOutput) Index(i pulumi.IntInput) NetworkRoutedV2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NetworkRoutedV2 {
		return vs[0].([]*NetworkRoutedV2)[vs[1].(int)]
	}).(NetworkRoutedV2Output)
}

type NetworkRoutedV2MapOutput struct{ *pulumi.OutputState }

func (NetworkRoutedV2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NetworkRoutedV2)(nil)).Elem()
}

func (o NetworkRoutedV2MapOutput) ToNetworkRoutedV2MapOutput() NetworkRoutedV2MapOutput {
	return o
}

func (o NetworkRoutedV2MapOutput) ToNetworkRoutedV2MapOutputWithContext(ctx context.Context) NetworkRoutedV2MapOutput {
	return o
}

func (o NetworkRoutedV2MapOutput) MapIndex(k pulumi.StringInput) NetworkRoutedV2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NetworkRoutedV2 {
		return vs[0].(map[string]*NetworkRoutedV2)[vs[1].(string)]
	}).(NetworkRoutedV2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkRoutedV2Input)(nil)).Elem(), &NetworkRoutedV2{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkRoutedV2ArrayInput)(nil)).Elem(), NetworkRoutedV2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*NetworkRoutedV2MapInput)(nil)).Elem(), NetworkRoutedV2Map{})
	pulumi.RegisterOutputType(NetworkRoutedV2Output{})
	pulumi.RegisterOutputType(NetworkRoutedV2ArrayOutput{})
	pulumi.RegisterOutputType(NetworkRoutedV2MapOutput{})
}
