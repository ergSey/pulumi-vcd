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

type NsxtNetworkImported struct {
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
	// ID of Distributed Virtual Port Group used by this network
	DvpgId pulumi.StringOutput `pulumi:"dvpgId"`
	// Unique name of an existing Distributed Virtual Port Group (DVPG).
	// **Note** it will never be refreshed because API does not allow reading this name after it is
	// consumed. Instead ID will be stored in `dvpgId` attribute.
	//
	// > One of `nsxtLogicalSwitchName` or `dvpgName` must be provided.
	DvpgName pulumi.StringPtrOutput `pulumi:"dvpgName"`
	// The gateway for this network (e.g. 192.168.1.1, 2002:0:0:1234:abcd:ffff:c0a7:121)
	Gateway pulumi.StringOutput `pulumi:"gateway"`
	// A unique name for the network
	Name pulumi.StringOutput `pulumi:"name"`
	// ID of NSX-T logical switch used by this network
	NsxtLogicalSwitchId pulumi.StringOutput `pulumi:"nsxtLogicalSwitchId"`
	// Unique name of an existing NSX-T segment.
	// **Note** it will never be refreshed because API does not allow reading this name after it is
	// consumed. Instead ID will be stored in `nsxtLogicalSwitchId` attribute.
	//
	// This resource **will fail** if multiple segments with the same name are available. One can rename
	// them in NSX-T manager to make them unique.
	NsxtLogicalSwitchName pulumi.StringPtrOutput `pulumi:"nsxtLogicalSwitchName"`
	// The name of organization to use, optional if defined at provider level. Useful when
	// connected as sysadmin working across different organisations
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
	SecondaryStaticIpPools NsxtNetworkImportedSecondaryStaticIpPoolArrayOutput `pulumi:"secondaryStaticIpPools"`
	// A range of IPs permitted to be used as static IPs for
	// virtual machines; see IP Pools below for details.
	StaticIpPools NsxtNetworkImportedStaticIpPoolArrayOutput `pulumi:"staticIpPools"`
	// The name of VDC to use. **Deprecated**  in favor of new field
	// `ownerId` which supports VDC and VDC Group IDs.
	//
	// Deprecated: This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
	Vdc pulumi.StringOutput `pulumi:"vdc"`
}

// NewNsxtNetworkImported registers a new resource with the given unique name, arguments, and options.
func NewNsxtNetworkImported(ctx *pulumi.Context,
	name string, args *NsxtNetworkImportedArgs, opts ...pulumi.ResourceOption) (*NsxtNetworkImported, error) {
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
	var resource NsxtNetworkImported
	err := ctx.RegisterResource("vcd:index/nsxtNetworkImported:NsxtNetworkImported", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNsxtNetworkImported gets an existing NsxtNetworkImported resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNsxtNetworkImported(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NsxtNetworkImportedState, opts ...pulumi.ResourceOption) (*NsxtNetworkImported, error) {
	var resource NsxtNetworkImported
	err := ctx.ReadResource("vcd:index/nsxtNetworkImported:NsxtNetworkImported", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NsxtNetworkImported resources.
type nsxtNetworkImportedState struct {
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
	// ID of Distributed Virtual Port Group used by this network
	DvpgId *string `pulumi:"dvpgId"`
	// Unique name of an existing Distributed Virtual Port Group (DVPG).
	// **Note** it will never be refreshed because API does not allow reading this name after it is
	// consumed. Instead ID will be stored in `dvpgId` attribute.
	//
	// > One of `nsxtLogicalSwitchName` or `dvpgName` must be provided.
	DvpgName *string `pulumi:"dvpgName"`
	// The gateway for this network (e.g. 192.168.1.1, 2002:0:0:1234:abcd:ffff:c0a7:121)
	Gateway *string `pulumi:"gateway"`
	// A unique name for the network
	Name *string `pulumi:"name"`
	// ID of NSX-T logical switch used by this network
	NsxtLogicalSwitchId *string `pulumi:"nsxtLogicalSwitchId"`
	// Unique name of an existing NSX-T segment.
	// **Note** it will never be refreshed because API does not allow reading this name after it is
	// consumed. Instead ID will be stored in `nsxtLogicalSwitchId` attribute.
	//
	// This resource **will fail** if multiple segments with the same name are available. One can rename
	// them in NSX-T manager to make them unique.
	NsxtLogicalSwitchName *string `pulumi:"nsxtLogicalSwitchName"`
	// The name of organization to use, optional if defined at provider level. Useful when
	// connected as sysadmin working across different organisations
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
	SecondaryStaticIpPools []NsxtNetworkImportedSecondaryStaticIpPool `pulumi:"secondaryStaticIpPools"`
	// A range of IPs permitted to be used as static IPs for
	// virtual machines; see IP Pools below for details.
	StaticIpPools []NsxtNetworkImportedStaticIpPool `pulumi:"staticIpPools"`
	// The name of VDC to use. **Deprecated**  in favor of new field
	// `ownerId` which supports VDC and VDC Group IDs.
	//
	// Deprecated: This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
	Vdc *string `pulumi:"vdc"`
}

type NsxtNetworkImportedState struct {
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
	// ID of Distributed Virtual Port Group used by this network
	DvpgId pulumi.StringPtrInput
	// Unique name of an existing Distributed Virtual Port Group (DVPG).
	// **Note** it will never be refreshed because API does not allow reading this name after it is
	// consumed. Instead ID will be stored in `dvpgId` attribute.
	//
	// > One of `nsxtLogicalSwitchName` or `dvpgName` must be provided.
	DvpgName pulumi.StringPtrInput
	// The gateway for this network (e.g. 192.168.1.1, 2002:0:0:1234:abcd:ffff:c0a7:121)
	Gateway pulumi.StringPtrInput
	// A unique name for the network
	Name pulumi.StringPtrInput
	// ID of NSX-T logical switch used by this network
	NsxtLogicalSwitchId pulumi.StringPtrInput
	// Unique name of an existing NSX-T segment.
	// **Note** it will never be refreshed because API does not allow reading this name after it is
	// consumed. Instead ID will be stored in `nsxtLogicalSwitchId` attribute.
	//
	// This resource **will fail** if multiple segments with the same name are available. One can rename
	// them in NSX-T manager to make them unique.
	NsxtLogicalSwitchName pulumi.StringPtrInput
	// The name of organization to use, optional if defined at provider level. Useful when
	// connected as sysadmin working across different organisations
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
	SecondaryStaticIpPools NsxtNetworkImportedSecondaryStaticIpPoolArrayInput
	// A range of IPs permitted to be used as static IPs for
	// virtual machines; see IP Pools below for details.
	StaticIpPools NsxtNetworkImportedStaticIpPoolArrayInput
	// The name of VDC to use. **Deprecated**  in favor of new field
	// `ownerId` which supports VDC and VDC Group IDs.
	//
	// Deprecated: This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
	Vdc pulumi.StringPtrInput
}

func (NsxtNetworkImportedState) ElementType() reflect.Type {
	return reflect.TypeOf((*nsxtNetworkImportedState)(nil)).Elem()
}

type nsxtNetworkImportedArgs struct {
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
	// Unique name of an existing Distributed Virtual Port Group (DVPG).
	// **Note** it will never be refreshed because API does not allow reading this name after it is
	// consumed. Instead ID will be stored in `dvpgId` attribute.
	//
	// > One of `nsxtLogicalSwitchName` or `dvpgName` must be provided.
	DvpgName *string `pulumi:"dvpgName"`
	// The gateway for this network (e.g. 192.168.1.1, 2002:0:0:1234:abcd:ffff:c0a7:121)
	Gateway string `pulumi:"gateway"`
	// A unique name for the network
	Name *string `pulumi:"name"`
	// Unique name of an existing NSX-T segment.
	// **Note** it will never be refreshed because API does not allow reading this name after it is
	// consumed. Instead ID will be stored in `nsxtLogicalSwitchId` attribute.
	//
	// This resource **will fail** if multiple segments with the same name are available. One can rename
	// them in NSX-T manager to make them unique.
	NsxtLogicalSwitchName *string `pulumi:"nsxtLogicalSwitchName"`
	// The name of organization to use, optional if defined at provider level. Useful when
	// connected as sysadmin working across different organisations
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
	SecondaryStaticIpPools []NsxtNetworkImportedSecondaryStaticIpPool `pulumi:"secondaryStaticIpPools"`
	// A range of IPs permitted to be used as static IPs for
	// virtual machines; see IP Pools below for details.
	StaticIpPools []NsxtNetworkImportedStaticIpPool `pulumi:"staticIpPools"`
	// The name of VDC to use. **Deprecated**  in favor of new field
	// `ownerId` which supports VDC and VDC Group IDs.
	//
	// Deprecated: This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
	Vdc *string `pulumi:"vdc"`
}

// The set of arguments for constructing a NsxtNetworkImported resource.
type NsxtNetworkImportedArgs struct {
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
	// Unique name of an existing Distributed Virtual Port Group (DVPG).
	// **Note** it will never be refreshed because API does not allow reading this name after it is
	// consumed. Instead ID will be stored in `dvpgId` attribute.
	//
	// > One of `nsxtLogicalSwitchName` or `dvpgName` must be provided.
	DvpgName pulumi.StringPtrInput
	// The gateway for this network (e.g. 192.168.1.1, 2002:0:0:1234:abcd:ffff:c0a7:121)
	Gateway pulumi.StringInput
	// A unique name for the network
	Name pulumi.StringPtrInput
	// Unique name of an existing NSX-T segment.
	// **Note** it will never be refreshed because API does not allow reading this name after it is
	// consumed. Instead ID will be stored in `nsxtLogicalSwitchId` attribute.
	//
	// This resource **will fail** if multiple segments with the same name are available. One can rename
	// them in NSX-T manager to make them unique.
	NsxtLogicalSwitchName pulumi.StringPtrInput
	// The name of organization to use, optional if defined at provider level. Useful when
	// connected as sysadmin working across different organisations
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
	SecondaryStaticIpPools NsxtNetworkImportedSecondaryStaticIpPoolArrayInput
	// A range of IPs permitted to be used as static IPs for
	// virtual machines; see IP Pools below for details.
	StaticIpPools NsxtNetworkImportedStaticIpPoolArrayInput
	// The name of VDC to use. **Deprecated**  in favor of new field
	// `ownerId` which supports VDC and VDC Group IDs.
	//
	// Deprecated: This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
	Vdc pulumi.StringPtrInput
}

func (NsxtNetworkImportedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nsxtNetworkImportedArgs)(nil)).Elem()
}

type NsxtNetworkImportedInput interface {
	pulumi.Input

	ToNsxtNetworkImportedOutput() NsxtNetworkImportedOutput
	ToNsxtNetworkImportedOutputWithContext(ctx context.Context) NsxtNetworkImportedOutput
}

func (*NsxtNetworkImported) ElementType() reflect.Type {
	return reflect.TypeOf((**NsxtNetworkImported)(nil)).Elem()
}

func (i *NsxtNetworkImported) ToNsxtNetworkImportedOutput() NsxtNetworkImportedOutput {
	return i.ToNsxtNetworkImportedOutputWithContext(context.Background())
}

func (i *NsxtNetworkImported) ToNsxtNetworkImportedOutputWithContext(ctx context.Context) NsxtNetworkImportedOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxtNetworkImportedOutput)
}

// NsxtNetworkImportedArrayInput is an input type that accepts NsxtNetworkImportedArray and NsxtNetworkImportedArrayOutput values.
// You can construct a concrete instance of `NsxtNetworkImportedArrayInput` via:
//
//	NsxtNetworkImportedArray{ NsxtNetworkImportedArgs{...} }
type NsxtNetworkImportedArrayInput interface {
	pulumi.Input

	ToNsxtNetworkImportedArrayOutput() NsxtNetworkImportedArrayOutput
	ToNsxtNetworkImportedArrayOutputWithContext(context.Context) NsxtNetworkImportedArrayOutput
}

type NsxtNetworkImportedArray []NsxtNetworkImportedInput

func (NsxtNetworkImportedArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NsxtNetworkImported)(nil)).Elem()
}

func (i NsxtNetworkImportedArray) ToNsxtNetworkImportedArrayOutput() NsxtNetworkImportedArrayOutput {
	return i.ToNsxtNetworkImportedArrayOutputWithContext(context.Background())
}

func (i NsxtNetworkImportedArray) ToNsxtNetworkImportedArrayOutputWithContext(ctx context.Context) NsxtNetworkImportedArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxtNetworkImportedArrayOutput)
}

// NsxtNetworkImportedMapInput is an input type that accepts NsxtNetworkImportedMap and NsxtNetworkImportedMapOutput values.
// You can construct a concrete instance of `NsxtNetworkImportedMapInput` via:
//
//	NsxtNetworkImportedMap{ "key": NsxtNetworkImportedArgs{...} }
type NsxtNetworkImportedMapInput interface {
	pulumi.Input

	ToNsxtNetworkImportedMapOutput() NsxtNetworkImportedMapOutput
	ToNsxtNetworkImportedMapOutputWithContext(context.Context) NsxtNetworkImportedMapOutput
}

type NsxtNetworkImportedMap map[string]NsxtNetworkImportedInput

func (NsxtNetworkImportedMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NsxtNetworkImported)(nil)).Elem()
}

func (i NsxtNetworkImportedMap) ToNsxtNetworkImportedMapOutput() NsxtNetworkImportedMapOutput {
	return i.ToNsxtNetworkImportedMapOutputWithContext(context.Background())
}

func (i NsxtNetworkImportedMap) ToNsxtNetworkImportedMapOutputWithContext(ctx context.Context) NsxtNetworkImportedMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxtNetworkImportedMapOutput)
}

type NsxtNetworkImportedOutput struct{ *pulumi.OutputState }

func (NsxtNetworkImportedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NsxtNetworkImported)(nil)).Elem()
}

func (o NsxtNetworkImportedOutput) ToNsxtNetworkImportedOutput() NsxtNetworkImportedOutput {
	return o
}

func (o NsxtNetworkImportedOutput) ToNsxtNetworkImportedOutputWithContext(ctx context.Context) NsxtNetworkImportedOutput {
	return o
}

// An optional description of the network
func (o NsxtNetworkImportedOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsxtNetworkImported) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// First DNS server to use.
func (o NsxtNetworkImportedOutput) Dns1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsxtNetworkImported) pulumi.StringPtrOutput { return v.Dns1 }).(pulumi.StringPtrOutput)
}

// Second DNS server to use.
func (o NsxtNetworkImportedOutput) Dns2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsxtNetworkImported) pulumi.StringPtrOutput { return v.Dns2 }).(pulumi.StringPtrOutput)
}

// A FQDN for the virtual machines on this network
func (o NsxtNetworkImportedOutput) DnsSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsxtNetworkImported) pulumi.StringPtrOutput { return v.DnsSuffix }).(pulumi.StringPtrOutput)
}

// Enables Dual-Stack mode so that one can configure one
// IPv4 and one IPv6 networks. **Note** In such case *IPv4* addresses must be used in `gateway`,
// `prefixLength` and `staticIpPool` while *IPv6* addresses in `secondaryGateway`,
// `secondaryPrefixLength` and `secondaryStaticIpPool` fields.
func (o NsxtNetworkImportedOutput) DualStackEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NsxtNetworkImported) pulumi.BoolPtrOutput { return v.DualStackEnabled }).(pulumi.BoolPtrOutput)
}

// ID of Distributed Virtual Port Group used by this network
func (o NsxtNetworkImportedOutput) DvpgId() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxtNetworkImported) pulumi.StringOutput { return v.DvpgId }).(pulumi.StringOutput)
}

// Unique name of an existing Distributed Virtual Port Group (DVPG).
// **Note** it will never be refreshed because API does not allow reading this name after it is
// consumed. Instead ID will be stored in `dvpgId` attribute.
//
// > One of `nsxtLogicalSwitchName` or `dvpgName` must be provided.
func (o NsxtNetworkImportedOutput) DvpgName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsxtNetworkImported) pulumi.StringPtrOutput { return v.DvpgName }).(pulumi.StringPtrOutput)
}

// The gateway for this network (e.g. 192.168.1.1, 2002:0:0:1234:abcd:ffff:c0a7:121)
func (o NsxtNetworkImportedOutput) Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxtNetworkImported) pulumi.StringOutput { return v.Gateway }).(pulumi.StringOutput)
}

// A unique name for the network
func (o NsxtNetworkImportedOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxtNetworkImported) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ID of NSX-T logical switch used by this network
func (o NsxtNetworkImportedOutput) NsxtLogicalSwitchId() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxtNetworkImported) pulumi.StringOutput { return v.NsxtLogicalSwitchId }).(pulumi.StringOutput)
}

// Unique name of an existing NSX-T segment.
// **Note** it will never be refreshed because API does not allow reading this name after it is
// consumed. Instead ID will be stored in `nsxtLogicalSwitchId` attribute.
//
// This resource **will fail** if multiple segments with the same name are available. One can rename
// them in NSX-T manager to make them unique.
func (o NsxtNetworkImportedOutput) NsxtLogicalSwitchName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsxtNetworkImported) pulumi.StringPtrOutput { return v.NsxtLogicalSwitchName }).(pulumi.StringPtrOutput)
}

// The name of organization to use, optional if defined at provider level. Useful when
// connected as sysadmin working across different organisations
func (o NsxtNetworkImportedOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsxtNetworkImported) pulumi.StringPtrOutput { return v.Org }).(pulumi.StringPtrOutput)
}

// VDC or VDC Group ID. Always takes precedence over `vdc` fields (in resource
// and inherited from provider configuration)
func (o NsxtNetworkImportedOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxtNetworkImported) pulumi.StringOutput { return v.OwnerId }).(pulumi.StringOutput)
}

// The prefix length for the new network (e.g. 24 for netmask 255.255.255.0).
func (o NsxtNetworkImportedOutput) PrefixLength() pulumi.IntOutput {
	return o.ApplyT(func(v *NsxtNetworkImported) pulumi.IntOutput { return v.PrefixLength }).(pulumi.IntOutput)
}

// IPv6 gateway *when Dual-Stack mode is enabled*
func (o NsxtNetworkImportedOutput) SecondaryGateway() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsxtNetworkImported) pulumi.StringPtrOutput { return v.SecondaryGateway }).(pulumi.StringPtrOutput)
}

// IPv6 prefix length *when Dual-Stack mode is
// enabled*
func (o NsxtNetworkImportedOutput) SecondaryPrefixLength() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsxtNetworkImported) pulumi.StringPtrOutput { return v.SecondaryPrefixLength }).(pulumi.StringPtrOutput)
}

// One or more IPv6 static
// pools *when Dual-Stack mode is enabled*
//
// > When using IPv6, VCD API will expand IP Addresses if they are specified using *double colon*
// notation and it will cause inconsistent plan. (e.g. `2002::1234:abcd:ffff:c0a6:121` will be
// converted to `2002:0:0:1234:abcd:ffff:c0a6:121`)
//
// <a id="ip-pools"></a>
func (o NsxtNetworkImportedOutput) SecondaryStaticIpPools() NsxtNetworkImportedSecondaryStaticIpPoolArrayOutput {
	return o.ApplyT(func(v *NsxtNetworkImported) NsxtNetworkImportedSecondaryStaticIpPoolArrayOutput {
		return v.SecondaryStaticIpPools
	}).(NsxtNetworkImportedSecondaryStaticIpPoolArrayOutput)
}

// A range of IPs permitted to be used as static IPs for
// virtual machines; see IP Pools below for details.
func (o NsxtNetworkImportedOutput) StaticIpPools() NsxtNetworkImportedStaticIpPoolArrayOutput {
	return o.ApplyT(func(v *NsxtNetworkImported) NsxtNetworkImportedStaticIpPoolArrayOutput { return v.StaticIpPools }).(NsxtNetworkImportedStaticIpPoolArrayOutput)
}

// The name of VDC to use. **Deprecated**  in favor of new field
// `ownerId` which supports VDC and VDC Group IDs.
//
// Deprecated: This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
func (o NsxtNetworkImportedOutput) Vdc() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxtNetworkImported) pulumi.StringOutput { return v.Vdc }).(pulumi.StringOutput)
}

type NsxtNetworkImportedArrayOutput struct{ *pulumi.OutputState }

func (NsxtNetworkImportedArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NsxtNetworkImported)(nil)).Elem()
}

func (o NsxtNetworkImportedArrayOutput) ToNsxtNetworkImportedArrayOutput() NsxtNetworkImportedArrayOutput {
	return o
}

func (o NsxtNetworkImportedArrayOutput) ToNsxtNetworkImportedArrayOutputWithContext(ctx context.Context) NsxtNetworkImportedArrayOutput {
	return o
}

func (o NsxtNetworkImportedArrayOutput) Index(i pulumi.IntInput) NsxtNetworkImportedOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NsxtNetworkImported {
		return vs[0].([]*NsxtNetworkImported)[vs[1].(int)]
	}).(NsxtNetworkImportedOutput)
}

type NsxtNetworkImportedMapOutput struct{ *pulumi.OutputState }

func (NsxtNetworkImportedMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NsxtNetworkImported)(nil)).Elem()
}

func (o NsxtNetworkImportedMapOutput) ToNsxtNetworkImportedMapOutput() NsxtNetworkImportedMapOutput {
	return o
}

func (o NsxtNetworkImportedMapOutput) ToNsxtNetworkImportedMapOutputWithContext(ctx context.Context) NsxtNetworkImportedMapOutput {
	return o
}

func (o NsxtNetworkImportedMapOutput) MapIndex(k pulumi.StringInput) NsxtNetworkImportedOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NsxtNetworkImported {
		return vs[0].(map[string]*NsxtNetworkImported)[vs[1].(string)]
	}).(NsxtNetworkImportedOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NsxtNetworkImportedInput)(nil)).Elem(), &NsxtNetworkImported{})
	pulumi.RegisterInputType(reflect.TypeOf((*NsxtNetworkImportedArrayInput)(nil)).Elem(), NsxtNetworkImportedArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NsxtNetworkImportedMapInput)(nil)).Elem(), NsxtNetworkImportedMap{})
	pulumi.RegisterOutputType(NsxtNetworkImportedOutput{})
	pulumi.RegisterOutputType(NsxtNetworkImportedArrayOutput{})
	pulumi.RegisterOutputType(NsxtNetworkImportedMapOutput{})
}
