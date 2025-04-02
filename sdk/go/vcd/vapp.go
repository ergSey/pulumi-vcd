// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/ergSey/pulumi-vcd/sdk/go/vcd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Vapp struct {
	pulumi.CustomResourceState

	// An optional description for the vApp, up to 256 characters.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// Key value map of vApp guest properties
	GuestProperties pulumi.StringMapOutput `pulumi:"guestProperties"`
	// (Computed) The vApp Hyper Reference.
	Href pulumi.StringOutput `pulumi:"href"`
	// (Computed; *v3.11+*; *VCD 10.5.1+*) A map that contains read-only metadata that is automatically added by VCD (10.5.1+) and provides
	// details on the origin of the vApp (e.g. `vapp.origin.id`, `vapp.origin.name`, `vapp.origin.type`).
	InheritedMetadata pulumi.StringMapOutput `pulumi:"inheritedMetadata"`
	// the information about the vApp lease. It includes the fields below. When this section is
	// included, both fields are mandatory. If lease values are higher than the ones allowed for the whole Org, the values
	// are **silently** reduced to the highest value allowed.
	Lease VappLeaseOutput `pulumi:"lease"`
	// Use `metadataEntry` instead. Key value map of metadata to assign to this vApp. Key and value can be any string. (Since *v2.2+* metadata is added directly to vApp instead of first VM in vApp)
	//
	// Deprecated: Use metadataEntry instead
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// A set of metadata entries to assign. See Metadata section for details.
	MetadataEntries VappMetadataEntryArrayOutput `pulumi:"metadataEntries"`
	// A unique name for the vApp
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
	Org pulumi.StringPtrOutput `pulumi:"org"`
	// A boolean value stating if this vApp should be powered on. Default is `false`. Works only on update when vApp already has VMs.
	PowerOn pulumi.BoolPtrOutput `pulumi:"powerOn"`
	// (Computed; *v2.5+*) The vApp status as a numeric code.
	Status pulumi.IntOutput `pulumi:"status"`
	// (Computed; *v2.5+*) The vApp status as text.
	StatusText pulumi.StringOutput `pulumi:"statusText"`
	// (*v3.13.0+*) A list of vApp network names included in this vApp
	VappNetworkNames pulumi.StringArrayOutput `pulumi:"vappNetworkNames"`
	// (*v3.13.0+*) A list of vApp Org network names included in this vApp
	VappOrgNetworkNames pulumi.StringArrayOutput `pulumi:"vappOrgNetworkNames"`
	// The name of VDC to use, optional if defined at provider level
	Vdc pulumi.StringPtrOutput `pulumi:"vdc"`
	// (*v3.13.0+*) A list of VM names included in this vApp
	VmNames pulumi.StringArrayOutput `pulumi:"vmNames"`
}

// NewVapp registers a new resource with the given unique name, arguments, and options.
func NewVapp(ctx *pulumi.Context,
	name string, args *VappArgs, opts ...pulumi.ResourceOption) (*Vapp, error) {
	if args == nil {
		args = &VappArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Vapp
	err := ctx.RegisterResource("vcd:index/vapp:Vapp", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVapp gets an existing Vapp resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVapp(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VappState, opts ...pulumi.ResourceOption) (*Vapp, error) {
	var resource Vapp
	err := ctx.ReadResource("vcd:index/vapp:Vapp", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Vapp resources.
type vappState struct {
	// An optional description for the vApp, up to 256 characters.
	Description *string `pulumi:"description"`
	// Key value map of vApp guest properties
	GuestProperties map[string]string `pulumi:"guestProperties"`
	// (Computed) The vApp Hyper Reference.
	Href *string `pulumi:"href"`
	// (Computed; *v3.11+*; *VCD 10.5.1+*) A map that contains read-only metadata that is automatically added by VCD (10.5.1+) and provides
	// details on the origin of the vApp (e.g. `vapp.origin.id`, `vapp.origin.name`, `vapp.origin.type`).
	InheritedMetadata map[string]string `pulumi:"inheritedMetadata"`
	// the information about the vApp lease. It includes the fields below. When this section is
	// included, both fields are mandatory. If lease values are higher than the ones allowed for the whole Org, the values
	// are **silently** reduced to the highest value allowed.
	Lease *VappLease `pulumi:"lease"`
	// Use `metadataEntry` instead. Key value map of metadata to assign to this vApp. Key and value can be any string. (Since *v2.2+* metadata is added directly to vApp instead of first VM in vApp)
	//
	// Deprecated: Use metadataEntry instead
	Metadata map[string]string `pulumi:"metadata"`
	// A set of metadata entries to assign. See Metadata section for details.
	MetadataEntries []VappMetadataEntry `pulumi:"metadataEntries"`
	// A unique name for the vApp
	Name *string `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
	Org *string `pulumi:"org"`
	// A boolean value stating if this vApp should be powered on. Default is `false`. Works only on update when vApp already has VMs.
	PowerOn *bool `pulumi:"powerOn"`
	// (Computed; *v2.5+*) The vApp status as a numeric code.
	Status *int `pulumi:"status"`
	// (Computed; *v2.5+*) The vApp status as text.
	StatusText *string `pulumi:"statusText"`
	// (*v3.13.0+*) A list of vApp network names included in this vApp
	VappNetworkNames []string `pulumi:"vappNetworkNames"`
	// (*v3.13.0+*) A list of vApp Org network names included in this vApp
	VappOrgNetworkNames []string `pulumi:"vappOrgNetworkNames"`
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `pulumi:"vdc"`
	// (*v3.13.0+*) A list of VM names included in this vApp
	VmNames []string `pulumi:"vmNames"`
}

type VappState struct {
	// An optional description for the vApp, up to 256 characters.
	Description pulumi.StringPtrInput
	// Key value map of vApp guest properties
	GuestProperties pulumi.StringMapInput
	// (Computed) The vApp Hyper Reference.
	Href pulumi.StringPtrInput
	// (Computed; *v3.11+*; *VCD 10.5.1+*) A map that contains read-only metadata that is automatically added by VCD (10.5.1+) and provides
	// details on the origin of the vApp (e.g. `vapp.origin.id`, `vapp.origin.name`, `vapp.origin.type`).
	InheritedMetadata pulumi.StringMapInput
	// the information about the vApp lease. It includes the fields below. When this section is
	// included, both fields are mandatory. If lease values are higher than the ones allowed for the whole Org, the values
	// are **silently** reduced to the highest value allowed.
	Lease VappLeasePtrInput
	// Use `metadataEntry` instead. Key value map of metadata to assign to this vApp. Key and value can be any string. (Since *v2.2+* metadata is added directly to vApp instead of first VM in vApp)
	//
	// Deprecated: Use metadataEntry instead
	Metadata pulumi.StringMapInput
	// A set of metadata entries to assign. See Metadata section for details.
	MetadataEntries VappMetadataEntryArrayInput
	// A unique name for the vApp
	Name pulumi.StringPtrInput
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
	Org pulumi.StringPtrInput
	// A boolean value stating if this vApp should be powered on. Default is `false`. Works only on update when vApp already has VMs.
	PowerOn pulumi.BoolPtrInput
	// (Computed; *v2.5+*) The vApp status as a numeric code.
	Status pulumi.IntPtrInput
	// (Computed; *v2.5+*) The vApp status as text.
	StatusText pulumi.StringPtrInput
	// (*v3.13.0+*) A list of vApp network names included in this vApp
	VappNetworkNames pulumi.StringArrayInput
	// (*v3.13.0+*) A list of vApp Org network names included in this vApp
	VappOrgNetworkNames pulumi.StringArrayInput
	// The name of VDC to use, optional if defined at provider level
	Vdc pulumi.StringPtrInput
	// (*v3.13.0+*) A list of VM names included in this vApp
	VmNames pulumi.StringArrayInput
}

func (VappState) ElementType() reflect.Type {
	return reflect.TypeOf((*vappState)(nil)).Elem()
}

type vappArgs struct {
	// An optional description for the vApp, up to 256 characters.
	Description *string `pulumi:"description"`
	// Key value map of vApp guest properties
	GuestProperties map[string]string `pulumi:"guestProperties"`
	// the information about the vApp lease. It includes the fields below. When this section is
	// included, both fields are mandatory. If lease values are higher than the ones allowed for the whole Org, the values
	// are **silently** reduced to the highest value allowed.
	Lease *VappLease `pulumi:"lease"`
	// Use `metadataEntry` instead. Key value map of metadata to assign to this vApp. Key and value can be any string. (Since *v2.2+* metadata is added directly to vApp instead of first VM in vApp)
	//
	// Deprecated: Use metadataEntry instead
	Metadata map[string]string `pulumi:"metadata"`
	// A set of metadata entries to assign. See Metadata section for details.
	MetadataEntries []VappMetadataEntry `pulumi:"metadataEntries"`
	// A unique name for the vApp
	Name *string `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
	Org *string `pulumi:"org"`
	// A boolean value stating if this vApp should be powered on. Default is `false`. Works only on update when vApp already has VMs.
	PowerOn *bool `pulumi:"powerOn"`
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `pulumi:"vdc"`
}

// The set of arguments for constructing a Vapp resource.
type VappArgs struct {
	// An optional description for the vApp, up to 256 characters.
	Description pulumi.StringPtrInput
	// Key value map of vApp guest properties
	GuestProperties pulumi.StringMapInput
	// the information about the vApp lease. It includes the fields below. When this section is
	// included, both fields are mandatory. If lease values are higher than the ones allowed for the whole Org, the values
	// are **silently** reduced to the highest value allowed.
	Lease VappLeasePtrInput
	// Use `metadataEntry` instead. Key value map of metadata to assign to this vApp. Key and value can be any string. (Since *v2.2+* metadata is added directly to vApp instead of first VM in vApp)
	//
	// Deprecated: Use metadataEntry instead
	Metadata pulumi.StringMapInput
	// A set of metadata entries to assign. See Metadata section for details.
	MetadataEntries VappMetadataEntryArrayInput
	// A unique name for the vApp
	Name pulumi.StringPtrInput
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
	Org pulumi.StringPtrInput
	// A boolean value stating if this vApp should be powered on. Default is `false`. Works only on update when vApp already has VMs.
	PowerOn pulumi.BoolPtrInput
	// The name of VDC to use, optional if defined at provider level
	Vdc pulumi.StringPtrInput
}

func (VappArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vappArgs)(nil)).Elem()
}

type VappInput interface {
	pulumi.Input

	ToVappOutput() VappOutput
	ToVappOutputWithContext(ctx context.Context) VappOutput
}

func (*Vapp) ElementType() reflect.Type {
	return reflect.TypeOf((**Vapp)(nil)).Elem()
}

func (i *Vapp) ToVappOutput() VappOutput {
	return i.ToVappOutputWithContext(context.Background())
}

func (i *Vapp) ToVappOutputWithContext(ctx context.Context) VappOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VappOutput)
}

// VappArrayInput is an input type that accepts VappArray and VappArrayOutput values.
// You can construct a concrete instance of `VappArrayInput` via:
//
//	VappArray{ VappArgs{...} }
type VappArrayInput interface {
	pulumi.Input

	ToVappArrayOutput() VappArrayOutput
	ToVappArrayOutputWithContext(context.Context) VappArrayOutput
}

type VappArray []VappInput

func (VappArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vapp)(nil)).Elem()
}

func (i VappArray) ToVappArrayOutput() VappArrayOutput {
	return i.ToVappArrayOutputWithContext(context.Background())
}

func (i VappArray) ToVappArrayOutputWithContext(ctx context.Context) VappArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VappArrayOutput)
}

// VappMapInput is an input type that accepts VappMap and VappMapOutput values.
// You can construct a concrete instance of `VappMapInput` via:
//
//	VappMap{ "key": VappArgs{...} }
type VappMapInput interface {
	pulumi.Input

	ToVappMapOutput() VappMapOutput
	ToVappMapOutputWithContext(context.Context) VappMapOutput
}

type VappMap map[string]VappInput

func (VappMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vapp)(nil)).Elem()
}

func (i VappMap) ToVappMapOutput() VappMapOutput {
	return i.ToVappMapOutputWithContext(context.Background())
}

func (i VappMap) ToVappMapOutputWithContext(ctx context.Context) VappMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VappMapOutput)
}

type VappOutput struct{ *pulumi.OutputState }

func (VappOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Vapp)(nil)).Elem()
}

func (o VappOutput) ToVappOutput() VappOutput {
	return o
}

func (o VappOutput) ToVappOutputWithContext(ctx context.Context) VappOutput {
	return o
}

// An optional description for the vApp, up to 256 characters.
func (o VappOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vapp) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// Key value map of vApp guest properties
func (o VappOutput) GuestProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Vapp) pulumi.StringMapOutput { return v.GuestProperties }).(pulumi.StringMapOutput)
}

// (Computed) The vApp Hyper Reference.
func (o VappOutput) Href() pulumi.StringOutput {
	return o.ApplyT(func(v *Vapp) pulumi.StringOutput { return v.Href }).(pulumi.StringOutput)
}

// (Computed; *v3.11+*; *VCD 10.5.1+*) A map that contains read-only metadata that is automatically added by VCD (10.5.1+) and provides
// details on the origin of the vApp (e.g. `vapp.origin.id`, `vapp.origin.name`, `vapp.origin.type`).
func (o VappOutput) InheritedMetadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Vapp) pulumi.StringMapOutput { return v.InheritedMetadata }).(pulumi.StringMapOutput)
}

// the information about the vApp lease. It includes the fields below. When this section is
// included, both fields are mandatory. If lease values are higher than the ones allowed for the whole Org, the values
// are **silently** reduced to the highest value allowed.
func (o VappOutput) Lease() VappLeaseOutput {
	return o.ApplyT(func(v *Vapp) VappLeaseOutput { return v.Lease }).(VappLeaseOutput)
}

// Use `metadataEntry` instead. Key value map of metadata to assign to this vApp. Key and value can be any string. (Since *v2.2+* metadata is added directly to vApp instead of first VM in vApp)
//
// Deprecated: Use metadataEntry instead
func (o VappOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Vapp) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

// A set of metadata entries to assign. See Metadata section for details.
func (o VappOutput) MetadataEntries() VappMetadataEntryArrayOutput {
	return o.ApplyT(func(v *Vapp) VappMetadataEntryArrayOutput { return v.MetadataEntries }).(VappMetadataEntryArrayOutput)
}

// A unique name for the vApp
func (o VappOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Vapp) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
func (o VappOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vapp) pulumi.StringPtrOutput { return v.Org }).(pulumi.StringPtrOutput)
}

// A boolean value stating if this vApp should be powered on. Default is `false`. Works only on update when vApp already has VMs.
func (o VappOutput) PowerOn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Vapp) pulumi.BoolPtrOutput { return v.PowerOn }).(pulumi.BoolPtrOutput)
}

// (Computed; *v2.5+*) The vApp status as a numeric code.
func (o VappOutput) Status() pulumi.IntOutput {
	return o.ApplyT(func(v *Vapp) pulumi.IntOutput { return v.Status }).(pulumi.IntOutput)
}

// (Computed; *v2.5+*) The vApp status as text.
func (o VappOutput) StatusText() pulumi.StringOutput {
	return o.ApplyT(func(v *Vapp) pulumi.StringOutput { return v.StatusText }).(pulumi.StringOutput)
}

// (*v3.13.0+*) A list of vApp network names included in this vApp
func (o VappOutput) VappNetworkNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Vapp) pulumi.StringArrayOutput { return v.VappNetworkNames }).(pulumi.StringArrayOutput)
}

// (*v3.13.0+*) A list of vApp Org network names included in this vApp
func (o VappOutput) VappOrgNetworkNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Vapp) pulumi.StringArrayOutput { return v.VappOrgNetworkNames }).(pulumi.StringArrayOutput)
}

// The name of VDC to use, optional if defined at provider level
func (o VappOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Vapp) pulumi.StringPtrOutput { return v.Vdc }).(pulumi.StringPtrOutput)
}

// (*v3.13.0+*) A list of VM names included in this vApp
func (o VappOutput) VmNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Vapp) pulumi.StringArrayOutput { return v.VmNames }).(pulumi.StringArrayOutput)
}

type VappArrayOutput struct{ *pulumi.OutputState }

func (VappArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Vapp)(nil)).Elem()
}

func (o VappArrayOutput) ToVappArrayOutput() VappArrayOutput {
	return o
}

func (o VappArrayOutput) ToVappArrayOutputWithContext(ctx context.Context) VappArrayOutput {
	return o
}

func (o VappArrayOutput) Index(i pulumi.IntInput) VappOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Vapp {
		return vs[0].([]*Vapp)[vs[1].(int)]
	}).(VappOutput)
}

type VappMapOutput struct{ *pulumi.OutputState }

func (VappMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Vapp)(nil)).Elem()
}

func (o VappMapOutput) ToVappMapOutput() VappMapOutput {
	return o
}

func (o VappMapOutput) ToVappMapOutputWithContext(ctx context.Context) VappMapOutput {
	return o
}

func (o VappMapOutput) MapIndex(k pulumi.StringInput) VappOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Vapp {
		return vs[0].(map[string]*Vapp)[vs[1].(string)]
	}).(VappOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VappInput)(nil)).Elem(), &Vapp{})
	pulumi.RegisterInputType(reflect.TypeOf((*VappArrayInput)(nil)).Elem(), VappArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VappMapInput)(nil)).Elem(), VappMap{})
	pulumi.RegisterOutputType(VappOutput{})
	pulumi.RegisterOutputType(VappArrayOutput{})
	pulumi.RegisterOutputType(VappMapOutput{})
}
