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

type ProviderVdc struct {
	pulumi.CustomResourceState

	// Set of virtual hardware versions supported by this Provider VDC.
	Capabilities pulumi.StringArrayOutput `pulumi:"capabilities"`
	// An indicator of CPU and memory capacity. See Compute Capacity below for details.
	ComputeCapacities ProviderVdcComputeCapacityArrayOutput `pulumi:"computeCapacities"`
	// Represents the compute fault domain for this Provider VDC. This value is a tenant-facing tag that is shown to tenants when viewing fault domains of the child Organization VDCs (for example, a VDC Group).
	ComputeProviderScope pulumi.StringOutput `pulumi:"computeProviderScope"`
	// Description of the Provider VDC.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Set of IDs of External Networks.
	ExternalNetworkIds pulumi.StringArrayOutput `pulumi:"externalNetworkIds"`
	// The highest virtual hardware version supported by this Provider VDC. This value cannot be changed to a lower version, and can only be updated when adding a new resource pool.
	HighestSupportedHardwareVersion pulumi.StringOutput `pulumi:"highestSupportedHardwareVersion"`
	// Set containing all the hosts which are connected to VC server.
	HostIds pulumi.StringArrayOutput `pulumi:"hostIds"`
	// True if this Provider VDC is enabled and can provide resources to organization VDCs. A Provider VDC is always enabled on creation.
	IsEnabled pulumi.BoolPtrOutput `pulumi:"isEnabled"`
	// A set of metadata entries assigned to the Provider VDC. See Metadata section for details.
	MetadataEntries ProviderVdcMetadataEntryArrayOutput `pulumi:"metadataEntries"`
	// Provider VDC name
	Name pulumi.StringOutput `pulumi:"name"`
	// Set IDs of the Network Pools used by this Provider VDC.
	NetworkPoolIds pulumi.StringArrayOutput `pulumi:"networkPoolIds"`
	// ID of the registered NSX-T Manager that backs networking operations for this Provider VDC.
	NsxtManagerId pulumi.StringPtrOutput `pulumi:"nsxtManagerId"`
	// Set of IDs of the Resource Pools backing this provider VDC. (Note: only one resource pool can be set at creation).
	ResourcePoolIds pulumi.StringArrayOutput `pulumi:"resourcePoolIds"`
	// Status of the Provider VDC: -1 (creation failed), 0 (not ready), 1 (ready), 2 (unknown) or 3 (unrecognized).
	Status pulumi.IntOutput `pulumi:"status"`
	// Set of IDs of the vSphere datastores backing this provider VDC
	StorageContainerIds pulumi.StringArrayOutput `pulumi:"storageContainerIds"`
	// Set of IDs to the Storage Profiles available to this Provider VDC.
	StorageProfileIds pulumi.StringArrayOutput `pulumi:"storageProfileIds"`
	// Set of Storage Profile names used to create this provider VDC.
	StorageProfileNames pulumi.StringArrayOutput `pulumi:"storageProfileNames"`
	// ID of the universal network reference.
	UniversalNetworkPoolId pulumi.StringOutput `pulumi:"universalNetworkPoolId"`
	// ID of the vCenter Server that provides the Resource Pools and Datastores.
	VcenterId pulumi.StringOutput `pulumi:"vcenterId"`
}

// NewProviderVdc registers a new resource with the given unique name, arguments, and options.
func NewProviderVdc(ctx *pulumi.Context,
	name string, args *ProviderVdcArgs, opts ...pulumi.ResourceOption) (*ProviderVdc, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.HighestSupportedHardwareVersion == nil {
		return nil, errors.New("invalid value for required argument 'HighestSupportedHardwareVersion'")
	}
	if args.ResourcePoolIds == nil {
		return nil, errors.New("invalid value for required argument 'ResourcePoolIds'")
	}
	if args.StorageProfileNames == nil {
		return nil, errors.New("invalid value for required argument 'StorageProfileNames'")
	}
	if args.VcenterId == nil {
		return nil, errors.New("invalid value for required argument 'VcenterId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ProviderVdc
	err := ctx.RegisterResource("vcd:index/providerVdc:ProviderVdc", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetProviderVdc gets an existing ProviderVdc resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetProviderVdc(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProviderVdcState, opts ...pulumi.ResourceOption) (*ProviderVdc, error) {
	var resource ProviderVdc
	err := ctx.ReadResource("vcd:index/providerVdc:ProviderVdc", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ProviderVdc resources.
type providerVdcState struct {
	// Set of virtual hardware versions supported by this Provider VDC.
	Capabilities []string `pulumi:"capabilities"`
	// An indicator of CPU and memory capacity. See Compute Capacity below for details.
	ComputeCapacities []ProviderVdcComputeCapacity `pulumi:"computeCapacities"`
	// Represents the compute fault domain for this Provider VDC. This value is a tenant-facing tag that is shown to tenants when viewing fault domains of the child Organization VDCs (for example, a VDC Group).
	ComputeProviderScope *string `pulumi:"computeProviderScope"`
	// Description of the Provider VDC.
	Description *string `pulumi:"description"`
	// Set of IDs of External Networks.
	ExternalNetworkIds []string `pulumi:"externalNetworkIds"`
	// The highest virtual hardware version supported by this Provider VDC. This value cannot be changed to a lower version, and can only be updated when adding a new resource pool.
	HighestSupportedHardwareVersion *string `pulumi:"highestSupportedHardwareVersion"`
	// Set containing all the hosts which are connected to VC server.
	HostIds []string `pulumi:"hostIds"`
	// True if this Provider VDC is enabled and can provide resources to organization VDCs. A Provider VDC is always enabled on creation.
	IsEnabled *bool `pulumi:"isEnabled"`
	// A set of metadata entries assigned to the Provider VDC. See Metadata section for details.
	MetadataEntries []ProviderVdcMetadataEntry `pulumi:"metadataEntries"`
	// Provider VDC name
	Name *string `pulumi:"name"`
	// Set IDs of the Network Pools used by this Provider VDC.
	NetworkPoolIds []string `pulumi:"networkPoolIds"`
	// ID of the registered NSX-T Manager that backs networking operations for this Provider VDC.
	NsxtManagerId *string `pulumi:"nsxtManagerId"`
	// Set of IDs of the Resource Pools backing this provider VDC. (Note: only one resource pool can be set at creation).
	ResourcePoolIds []string `pulumi:"resourcePoolIds"`
	// Status of the Provider VDC: -1 (creation failed), 0 (not ready), 1 (ready), 2 (unknown) or 3 (unrecognized).
	Status *int `pulumi:"status"`
	// Set of IDs of the vSphere datastores backing this provider VDC
	StorageContainerIds []string `pulumi:"storageContainerIds"`
	// Set of IDs to the Storage Profiles available to this Provider VDC.
	StorageProfileIds []string `pulumi:"storageProfileIds"`
	// Set of Storage Profile names used to create this provider VDC.
	StorageProfileNames []string `pulumi:"storageProfileNames"`
	// ID of the universal network reference.
	UniversalNetworkPoolId *string `pulumi:"universalNetworkPoolId"`
	// ID of the vCenter Server that provides the Resource Pools and Datastores.
	VcenterId *string `pulumi:"vcenterId"`
}

type ProviderVdcState struct {
	// Set of virtual hardware versions supported by this Provider VDC.
	Capabilities pulumi.StringArrayInput
	// An indicator of CPU and memory capacity. See Compute Capacity below for details.
	ComputeCapacities ProviderVdcComputeCapacityArrayInput
	// Represents the compute fault domain for this Provider VDC. This value is a tenant-facing tag that is shown to tenants when viewing fault domains of the child Organization VDCs (for example, a VDC Group).
	ComputeProviderScope pulumi.StringPtrInput
	// Description of the Provider VDC.
	Description pulumi.StringPtrInput
	// Set of IDs of External Networks.
	ExternalNetworkIds pulumi.StringArrayInput
	// The highest virtual hardware version supported by this Provider VDC. This value cannot be changed to a lower version, and can only be updated when adding a new resource pool.
	HighestSupportedHardwareVersion pulumi.StringPtrInput
	// Set containing all the hosts which are connected to VC server.
	HostIds pulumi.StringArrayInput
	// True if this Provider VDC is enabled and can provide resources to organization VDCs. A Provider VDC is always enabled on creation.
	IsEnabled pulumi.BoolPtrInput
	// A set of metadata entries assigned to the Provider VDC. See Metadata section for details.
	MetadataEntries ProviderVdcMetadataEntryArrayInput
	// Provider VDC name
	Name pulumi.StringPtrInput
	// Set IDs of the Network Pools used by this Provider VDC.
	NetworkPoolIds pulumi.StringArrayInput
	// ID of the registered NSX-T Manager that backs networking operations for this Provider VDC.
	NsxtManagerId pulumi.StringPtrInput
	// Set of IDs of the Resource Pools backing this provider VDC. (Note: only one resource pool can be set at creation).
	ResourcePoolIds pulumi.StringArrayInput
	// Status of the Provider VDC: -1 (creation failed), 0 (not ready), 1 (ready), 2 (unknown) or 3 (unrecognized).
	Status pulumi.IntPtrInput
	// Set of IDs of the vSphere datastores backing this provider VDC
	StorageContainerIds pulumi.StringArrayInput
	// Set of IDs to the Storage Profiles available to this Provider VDC.
	StorageProfileIds pulumi.StringArrayInput
	// Set of Storage Profile names used to create this provider VDC.
	StorageProfileNames pulumi.StringArrayInput
	// ID of the universal network reference.
	UniversalNetworkPoolId pulumi.StringPtrInput
	// ID of the vCenter Server that provides the Resource Pools and Datastores.
	VcenterId pulumi.StringPtrInput
}

func (ProviderVdcState) ElementType() reflect.Type {
	return reflect.TypeOf((*providerVdcState)(nil)).Elem()
}

type providerVdcArgs struct {
	// Description of the Provider VDC.
	Description *string `pulumi:"description"`
	// The highest virtual hardware version supported by this Provider VDC. This value cannot be changed to a lower version, and can only be updated when adding a new resource pool.
	HighestSupportedHardwareVersion string `pulumi:"highestSupportedHardwareVersion"`
	// True if this Provider VDC is enabled and can provide resources to organization VDCs. A Provider VDC is always enabled on creation.
	IsEnabled *bool `pulumi:"isEnabled"`
	// A set of metadata entries assigned to the Provider VDC. See Metadata section for details.
	MetadataEntries []ProviderVdcMetadataEntry `pulumi:"metadataEntries"`
	// Provider VDC name
	Name *string `pulumi:"name"`
	// Set IDs of the Network Pools used by this Provider VDC.
	NetworkPoolIds []string `pulumi:"networkPoolIds"`
	// ID of the registered NSX-T Manager that backs networking operations for this Provider VDC.
	NsxtManagerId *string `pulumi:"nsxtManagerId"`
	// Set of IDs of the Resource Pools backing this provider VDC. (Note: only one resource pool can be set at creation).
	ResourcePoolIds []string `pulumi:"resourcePoolIds"`
	// Set of Storage Profile names used to create this provider VDC.
	StorageProfileNames []string `pulumi:"storageProfileNames"`
	// ID of the vCenter Server that provides the Resource Pools and Datastores.
	VcenterId string `pulumi:"vcenterId"`
}

// The set of arguments for constructing a ProviderVdc resource.
type ProviderVdcArgs struct {
	// Description of the Provider VDC.
	Description pulumi.StringPtrInput
	// The highest virtual hardware version supported by this Provider VDC. This value cannot be changed to a lower version, and can only be updated when adding a new resource pool.
	HighestSupportedHardwareVersion pulumi.StringInput
	// True if this Provider VDC is enabled and can provide resources to organization VDCs. A Provider VDC is always enabled on creation.
	IsEnabled pulumi.BoolPtrInput
	// A set of metadata entries assigned to the Provider VDC. See Metadata section for details.
	MetadataEntries ProviderVdcMetadataEntryArrayInput
	// Provider VDC name
	Name pulumi.StringPtrInput
	// Set IDs of the Network Pools used by this Provider VDC.
	NetworkPoolIds pulumi.StringArrayInput
	// ID of the registered NSX-T Manager that backs networking operations for this Provider VDC.
	NsxtManagerId pulumi.StringPtrInput
	// Set of IDs of the Resource Pools backing this provider VDC. (Note: only one resource pool can be set at creation).
	ResourcePoolIds pulumi.StringArrayInput
	// Set of Storage Profile names used to create this provider VDC.
	StorageProfileNames pulumi.StringArrayInput
	// ID of the vCenter Server that provides the Resource Pools and Datastores.
	VcenterId pulumi.StringInput
}

func (ProviderVdcArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerVdcArgs)(nil)).Elem()
}

type ProviderVdcInput interface {
	pulumi.Input

	ToProviderVdcOutput() ProviderVdcOutput
	ToProviderVdcOutputWithContext(ctx context.Context) ProviderVdcOutput
}

func (*ProviderVdc) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderVdc)(nil)).Elem()
}

func (i *ProviderVdc) ToProviderVdcOutput() ProviderVdcOutput {
	return i.ToProviderVdcOutputWithContext(context.Background())
}

func (i *ProviderVdc) ToProviderVdcOutputWithContext(ctx context.Context) ProviderVdcOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderVdcOutput)
}

// ProviderVdcArrayInput is an input type that accepts ProviderVdcArray and ProviderVdcArrayOutput values.
// You can construct a concrete instance of `ProviderVdcArrayInput` via:
//
//	ProviderVdcArray{ ProviderVdcArgs{...} }
type ProviderVdcArrayInput interface {
	pulumi.Input

	ToProviderVdcArrayOutput() ProviderVdcArrayOutput
	ToProviderVdcArrayOutputWithContext(context.Context) ProviderVdcArrayOutput
}

type ProviderVdcArray []ProviderVdcInput

func (ProviderVdcArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProviderVdc)(nil)).Elem()
}

func (i ProviderVdcArray) ToProviderVdcArrayOutput() ProviderVdcArrayOutput {
	return i.ToProviderVdcArrayOutputWithContext(context.Background())
}

func (i ProviderVdcArray) ToProviderVdcArrayOutputWithContext(ctx context.Context) ProviderVdcArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderVdcArrayOutput)
}

// ProviderVdcMapInput is an input type that accepts ProviderVdcMap and ProviderVdcMapOutput values.
// You can construct a concrete instance of `ProviderVdcMapInput` via:
//
//	ProviderVdcMap{ "key": ProviderVdcArgs{...} }
type ProviderVdcMapInput interface {
	pulumi.Input

	ToProviderVdcMapOutput() ProviderVdcMapOutput
	ToProviderVdcMapOutputWithContext(context.Context) ProviderVdcMapOutput
}

type ProviderVdcMap map[string]ProviderVdcInput

func (ProviderVdcMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProviderVdc)(nil)).Elem()
}

func (i ProviderVdcMap) ToProviderVdcMapOutput() ProviderVdcMapOutput {
	return i.ToProviderVdcMapOutputWithContext(context.Background())
}

func (i ProviderVdcMap) ToProviderVdcMapOutputWithContext(ctx context.Context) ProviderVdcMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderVdcMapOutput)
}

type ProviderVdcOutput struct{ *pulumi.OutputState }

func (ProviderVdcOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProviderVdc)(nil)).Elem()
}

func (o ProviderVdcOutput) ToProviderVdcOutput() ProviderVdcOutput {
	return o
}

func (o ProviderVdcOutput) ToProviderVdcOutputWithContext(ctx context.Context) ProviderVdcOutput {
	return o
}

// Set of virtual hardware versions supported by this Provider VDC.
func (o ProviderVdcOutput) Capabilities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProviderVdc) pulumi.StringArrayOutput { return v.Capabilities }).(pulumi.StringArrayOutput)
}

// An indicator of CPU and memory capacity. See Compute Capacity below for details.
func (o ProviderVdcOutput) ComputeCapacities() ProviderVdcComputeCapacityArrayOutput {
	return o.ApplyT(func(v *ProviderVdc) ProviderVdcComputeCapacityArrayOutput { return v.ComputeCapacities }).(ProviderVdcComputeCapacityArrayOutput)
}

// Represents the compute fault domain for this Provider VDC. This value is a tenant-facing tag that is shown to tenants when viewing fault domains of the child Organization VDCs (for example, a VDC Group).
func (o ProviderVdcOutput) ComputeProviderScope() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderVdc) pulumi.StringOutput { return v.ComputeProviderScope }).(pulumi.StringOutput)
}

// Description of the Provider VDC.
func (o ProviderVdcOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderVdc) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Set of IDs of External Networks.
func (o ProviderVdcOutput) ExternalNetworkIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProviderVdc) pulumi.StringArrayOutput { return v.ExternalNetworkIds }).(pulumi.StringArrayOutput)
}

// The highest virtual hardware version supported by this Provider VDC. This value cannot be changed to a lower version, and can only be updated when adding a new resource pool.
func (o ProviderVdcOutput) HighestSupportedHardwareVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderVdc) pulumi.StringOutput { return v.HighestSupportedHardwareVersion }).(pulumi.StringOutput)
}

// Set containing all the hosts which are connected to VC server.
func (o ProviderVdcOutput) HostIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProviderVdc) pulumi.StringArrayOutput { return v.HostIds }).(pulumi.StringArrayOutput)
}

// True if this Provider VDC is enabled and can provide resources to organization VDCs. A Provider VDC is always enabled on creation.
func (o ProviderVdcOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ProviderVdc) pulumi.BoolPtrOutput { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

// A set of metadata entries assigned to the Provider VDC. See Metadata section for details.
func (o ProviderVdcOutput) MetadataEntries() ProviderVdcMetadataEntryArrayOutput {
	return o.ApplyT(func(v *ProviderVdc) ProviderVdcMetadataEntryArrayOutput { return v.MetadataEntries }).(ProviderVdcMetadataEntryArrayOutput)
}

// Provider VDC name
func (o ProviderVdcOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderVdc) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Set IDs of the Network Pools used by this Provider VDC.
func (o ProviderVdcOutput) NetworkPoolIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProviderVdc) pulumi.StringArrayOutput { return v.NetworkPoolIds }).(pulumi.StringArrayOutput)
}

// ID of the registered NSX-T Manager that backs networking operations for this Provider VDC.
func (o ProviderVdcOutput) NsxtManagerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProviderVdc) pulumi.StringPtrOutput { return v.NsxtManagerId }).(pulumi.StringPtrOutput)
}

// Set of IDs of the Resource Pools backing this provider VDC. (Note: only one resource pool can be set at creation).
func (o ProviderVdcOutput) ResourcePoolIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProviderVdc) pulumi.StringArrayOutput { return v.ResourcePoolIds }).(pulumi.StringArrayOutput)
}

// Status of the Provider VDC: -1 (creation failed), 0 (not ready), 1 (ready), 2 (unknown) or 3 (unrecognized).
func (o ProviderVdcOutput) Status() pulumi.IntOutput {
	return o.ApplyT(func(v *ProviderVdc) pulumi.IntOutput { return v.Status }).(pulumi.IntOutput)
}

// Set of IDs of the vSphere datastores backing this provider VDC
func (o ProviderVdcOutput) StorageContainerIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProviderVdc) pulumi.StringArrayOutput { return v.StorageContainerIds }).(pulumi.StringArrayOutput)
}

// Set of IDs to the Storage Profiles available to this Provider VDC.
func (o ProviderVdcOutput) StorageProfileIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProviderVdc) pulumi.StringArrayOutput { return v.StorageProfileIds }).(pulumi.StringArrayOutput)
}

// Set of Storage Profile names used to create this provider VDC.
func (o ProviderVdcOutput) StorageProfileNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ProviderVdc) pulumi.StringArrayOutput { return v.StorageProfileNames }).(pulumi.StringArrayOutput)
}

// ID of the universal network reference.
func (o ProviderVdcOutput) UniversalNetworkPoolId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderVdc) pulumi.StringOutput { return v.UniversalNetworkPoolId }).(pulumi.StringOutput)
}

// ID of the vCenter Server that provides the Resource Pools and Datastores.
func (o ProviderVdcOutput) VcenterId() pulumi.StringOutput {
	return o.ApplyT(func(v *ProviderVdc) pulumi.StringOutput { return v.VcenterId }).(pulumi.StringOutput)
}

type ProviderVdcArrayOutput struct{ *pulumi.OutputState }

func (ProviderVdcArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ProviderVdc)(nil)).Elem()
}

func (o ProviderVdcArrayOutput) ToProviderVdcArrayOutput() ProviderVdcArrayOutput {
	return o
}

func (o ProviderVdcArrayOutput) ToProviderVdcArrayOutputWithContext(ctx context.Context) ProviderVdcArrayOutput {
	return o
}

func (o ProviderVdcArrayOutput) Index(i pulumi.IntInput) ProviderVdcOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ProviderVdc {
		return vs[0].([]*ProviderVdc)[vs[1].(int)]
	}).(ProviderVdcOutput)
}

type ProviderVdcMapOutput struct{ *pulumi.OutputState }

func (ProviderVdcMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ProviderVdc)(nil)).Elem()
}

func (o ProviderVdcMapOutput) ToProviderVdcMapOutput() ProviderVdcMapOutput {
	return o
}

func (o ProviderVdcMapOutput) ToProviderVdcMapOutputWithContext(ctx context.Context) ProviderVdcMapOutput {
	return o
}

func (o ProviderVdcMapOutput) MapIndex(k pulumi.StringInput) ProviderVdcOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ProviderVdc {
		return vs[0].(map[string]*ProviderVdc)[vs[1].(string)]
	}).(ProviderVdcOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderVdcInput)(nil)).Elem(), &ProviderVdc{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderVdcArrayInput)(nil)).Elem(), ProviderVdcArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderVdcMapInput)(nil)).Elem(), ProviderVdcMap{})
	pulumi.RegisterOutputType(ProviderVdcOutput{})
	pulumi.RegisterOutputType(ProviderVdcArrayOutput{})
	pulumi.RegisterOutputType(ProviderVdcMapOutput{})
}
