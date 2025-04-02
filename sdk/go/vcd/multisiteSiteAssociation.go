// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/ergSey/pulumi-vcd/sdk/go/vcd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MultisiteSiteAssociation struct {
	pulumi.CustomResourceState

	// URL of the associated site
	AssociatedSiteHref pulumi.StringOutput `pulumi:"associatedSiteHref"`
	// ID of the remote site associated with the current one.
	AssociatedSiteId pulumi.StringOutput `pulumi:"associatedSiteId"`
	// The name of the associated site.
	AssociatedSiteName pulumi.StringOutput `pulumi:"associatedSiteName"`
	// Data produced from another site, needed to associate to this site from another one.
	// (Used instead of `associatedDataFile`)
	AssociationData pulumi.StringPtrOutput `pulumi:"associationData"`
	// Name of the file containing the data used to associate to this site from another one.
	// (Used instead of `associatedData`). This file can be created (by the other site administrator) using the data source `getMultisiteSiteData`.
	AssociationDataFile pulumi.StringPtrOutput `pulumi:"associationDataFile"`
	// How many minutes we wait for the association to be complete. (0 = no check)
	// This property is only used during update, and should not be used until both sides of the association have been completed.
	ConnectionTimeoutMins pulumi.IntPtrOutput `pulumi:"connectionTimeoutMins"`
	// The status of the association (one of `ASYMMETRIC`, `ACTIVE`, `UNREACHABLE`, `ERROR`)
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewMultisiteSiteAssociation registers a new resource with the given unique name, arguments, and options.
func NewMultisiteSiteAssociation(ctx *pulumi.Context,
	name string, args *MultisiteSiteAssociationArgs, opts ...pulumi.ResourceOption) (*MultisiteSiteAssociation, error) {
	if args == nil {
		args = &MultisiteSiteAssociationArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MultisiteSiteAssociation
	err := ctx.RegisterResource("vcd:index/multisiteSiteAssociation:MultisiteSiteAssociation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMultisiteSiteAssociation gets an existing MultisiteSiteAssociation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMultisiteSiteAssociation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MultisiteSiteAssociationState, opts ...pulumi.ResourceOption) (*MultisiteSiteAssociation, error) {
	var resource MultisiteSiteAssociation
	err := ctx.ReadResource("vcd:index/multisiteSiteAssociation:MultisiteSiteAssociation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MultisiteSiteAssociation resources.
type multisiteSiteAssociationState struct {
	// URL of the associated site
	AssociatedSiteHref *string `pulumi:"associatedSiteHref"`
	// ID of the remote site associated with the current one.
	AssociatedSiteId *string `pulumi:"associatedSiteId"`
	// The name of the associated site.
	AssociatedSiteName *string `pulumi:"associatedSiteName"`
	// Data produced from another site, needed to associate to this site from another one.
	// (Used instead of `associatedDataFile`)
	AssociationData *string `pulumi:"associationData"`
	// Name of the file containing the data used to associate to this site from another one.
	// (Used instead of `associatedData`). This file can be created (by the other site administrator) using the data source `getMultisiteSiteData`.
	AssociationDataFile *string `pulumi:"associationDataFile"`
	// How many minutes we wait for the association to be complete. (0 = no check)
	// This property is only used during update, and should not be used until both sides of the association have been completed.
	ConnectionTimeoutMins *int `pulumi:"connectionTimeoutMins"`
	// The status of the association (one of `ASYMMETRIC`, `ACTIVE`, `UNREACHABLE`, `ERROR`)
	Status *string `pulumi:"status"`
}

type MultisiteSiteAssociationState struct {
	// URL of the associated site
	AssociatedSiteHref pulumi.StringPtrInput
	// ID of the remote site associated with the current one.
	AssociatedSiteId pulumi.StringPtrInput
	// The name of the associated site.
	AssociatedSiteName pulumi.StringPtrInput
	// Data produced from another site, needed to associate to this site from another one.
	// (Used instead of `associatedDataFile`)
	AssociationData pulumi.StringPtrInput
	// Name of the file containing the data used to associate to this site from another one.
	// (Used instead of `associatedData`). This file can be created (by the other site administrator) using the data source `getMultisiteSiteData`.
	AssociationDataFile pulumi.StringPtrInput
	// How many minutes we wait for the association to be complete. (0 = no check)
	// This property is only used during update, and should not be used until both sides of the association have been completed.
	ConnectionTimeoutMins pulumi.IntPtrInput
	// The status of the association (one of `ASYMMETRIC`, `ACTIVE`, `UNREACHABLE`, `ERROR`)
	Status pulumi.StringPtrInput
}

func (MultisiteSiteAssociationState) ElementType() reflect.Type {
	return reflect.TypeOf((*multisiteSiteAssociationState)(nil)).Elem()
}

type multisiteSiteAssociationArgs struct {
	// Data produced from another site, needed to associate to this site from another one.
	// (Used instead of `associatedDataFile`)
	AssociationData *string `pulumi:"associationData"`
	// Name of the file containing the data used to associate to this site from another one.
	// (Used instead of `associatedData`). This file can be created (by the other site administrator) using the data source `getMultisiteSiteData`.
	AssociationDataFile *string `pulumi:"associationDataFile"`
	// How many minutes we wait for the association to be complete. (0 = no check)
	// This property is only used during update, and should not be used until both sides of the association have been completed.
	ConnectionTimeoutMins *int `pulumi:"connectionTimeoutMins"`
}

// The set of arguments for constructing a MultisiteSiteAssociation resource.
type MultisiteSiteAssociationArgs struct {
	// Data produced from another site, needed to associate to this site from another one.
	// (Used instead of `associatedDataFile`)
	AssociationData pulumi.StringPtrInput
	// Name of the file containing the data used to associate to this site from another one.
	// (Used instead of `associatedData`). This file can be created (by the other site administrator) using the data source `getMultisiteSiteData`.
	AssociationDataFile pulumi.StringPtrInput
	// How many minutes we wait for the association to be complete. (0 = no check)
	// This property is only used during update, and should not be used until both sides of the association have been completed.
	ConnectionTimeoutMins pulumi.IntPtrInput
}

func (MultisiteSiteAssociationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*multisiteSiteAssociationArgs)(nil)).Elem()
}

type MultisiteSiteAssociationInput interface {
	pulumi.Input

	ToMultisiteSiteAssociationOutput() MultisiteSiteAssociationOutput
	ToMultisiteSiteAssociationOutputWithContext(ctx context.Context) MultisiteSiteAssociationOutput
}

func (*MultisiteSiteAssociation) ElementType() reflect.Type {
	return reflect.TypeOf((**MultisiteSiteAssociation)(nil)).Elem()
}

func (i *MultisiteSiteAssociation) ToMultisiteSiteAssociationOutput() MultisiteSiteAssociationOutput {
	return i.ToMultisiteSiteAssociationOutputWithContext(context.Background())
}

func (i *MultisiteSiteAssociation) ToMultisiteSiteAssociationOutputWithContext(ctx context.Context) MultisiteSiteAssociationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MultisiteSiteAssociationOutput)
}

// MultisiteSiteAssociationArrayInput is an input type that accepts MultisiteSiteAssociationArray and MultisiteSiteAssociationArrayOutput values.
// You can construct a concrete instance of `MultisiteSiteAssociationArrayInput` via:
//
//	MultisiteSiteAssociationArray{ MultisiteSiteAssociationArgs{...} }
type MultisiteSiteAssociationArrayInput interface {
	pulumi.Input

	ToMultisiteSiteAssociationArrayOutput() MultisiteSiteAssociationArrayOutput
	ToMultisiteSiteAssociationArrayOutputWithContext(context.Context) MultisiteSiteAssociationArrayOutput
}

type MultisiteSiteAssociationArray []MultisiteSiteAssociationInput

func (MultisiteSiteAssociationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MultisiteSiteAssociation)(nil)).Elem()
}

func (i MultisiteSiteAssociationArray) ToMultisiteSiteAssociationArrayOutput() MultisiteSiteAssociationArrayOutput {
	return i.ToMultisiteSiteAssociationArrayOutputWithContext(context.Background())
}

func (i MultisiteSiteAssociationArray) ToMultisiteSiteAssociationArrayOutputWithContext(ctx context.Context) MultisiteSiteAssociationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MultisiteSiteAssociationArrayOutput)
}

// MultisiteSiteAssociationMapInput is an input type that accepts MultisiteSiteAssociationMap and MultisiteSiteAssociationMapOutput values.
// You can construct a concrete instance of `MultisiteSiteAssociationMapInput` via:
//
//	MultisiteSiteAssociationMap{ "key": MultisiteSiteAssociationArgs{...} }
type MultisiteSiteAssociationMapInput interface {
	pulumi.Input

	ToMultisiteSiteAssociationMapOutput() MultisiteSiteAssociationMapOutput
	ToMultisiteSiteAssociationMapOutputWithContext(context.Context) MultisiteSiteAssociationMapOutput
}

type MultisiteSiteAssociationMap map[string]MultisiteSiteAssociationInput

func (MultisiteSiteAssociationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MultisiteSiteAssociation)(nil)).Elem()
}

func (i MultisiteSiteAssociationMap) ToMultisiteSiteAssociationMapOutput() MultisiteSiteAssociationMapOutput {
	return i.ToMultisiteSiteAssociationMapOutputWithContext(context.Background())
}

func (i MultisiteSiteAssociationMap) ToMultisiteSiteAssociationMapOutputWithContext(ctx context.Context) MultisiteSiteAssociationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MultisiteSiteAssociationMapOutput)
}

type MultisiteSiteAssociationOutput struct{ *pulumi.OutputState }

func (MultisiteSiteAssociationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MultisiteSiteAssociation)(nil)).Elem()
}

func (o MultisiteSiteAssociationOutput) ToMultisiteSiteAssociationOutput() MultisiteSiteAssociationOutput {
	return o
}

func (o MultisiteSiteAssociationOutput) ToMultisiteSiteAssociationOutputWithContext(ctx context.Context) MultisiteSiteAssociationOutput {
	return o
}

// URL of the associated site
func (o MultisiteSiteAssociationOutput) AssociatedSiteHref() pulumi.StringOutput {
	return o.ApplyT(func(v *MultisiteSiteAssociation) pulumi.StringOutput { return v.AssociatedSiteHref }).(pulumi.StringOutput)
}

// ID of the remote site associated with the current one.
func (o MultisiteSiteAssociationOutput) AssociatedSiteId() pulumi.StringOutput {
	return o.ApplyT(func(v *MultisiteSiteAssociation) pulumi.StringOutput { return v.AssociatedSiteId }).(pulumi.StringOutput)
}

// The name of the associated site.
func (o MultisiteSiteAssociationOutput) AssociatedSiteName() pulumi.StringOutput {
	return o.ApplyT(func(v *MultisiteSiteAssociation) pulumi.StringOutput { return v.AssociatedSiteName }).(pulumi.StringOutput)
}

// Data produced from another site, needed to associate to this site from another one.
// (Used instead of `associatedDataFile`)
func (o MultisiteSiteAssociationOutput) AssociationData() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MultisiteSiteAssociation) pulumi.StringPtrOutput { return v.AssociationData }).(pulumi.StringPtrOutput)
}

// Name of the file containing the data used to associate to this site from another one.
// (Used instead of `associatedData`). This file can be created (by the other site administrator) using the data source `getMultisiteSiteData`.
func (o MultisiteSiteAssociationOutput) AssociationDataFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MultisiteSiteAssociation) pulumi.StringPtrOutput { return v.AssociationDataFile }).(pulumi.StringPtrOutput)
}

// How many minutes we wait for the association to be complete. (0 = no check)
// This property is only used during update, and should not be used until both sides of the association have been completed.
func (o MultisiteSiteAssociationOutput) ConnectionTimeoutMins() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MultisiteSiteAssociation) pulumi.IntPtrOutput { return v.ConnectionTimeoutMins }).(pulumi.IntPtrOutput)
}

// The status of the association (one of `ASYMMETRIC`, `ACTIVE`, `UNREACHABLE`, `ERROR`)
func (o MultisiteSiteAssociationOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v *MultisiteSiteAssociation) pulumi.StringOutput { return v.Status }).(pulumi.StringOutput)
}

type MultisiteSiteAssociationArrayOutput struct{ *pulumi.OutputState }

func (MultisiteSiteAssociationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MultisiteSiteAssociation)(nil)).Elem()
}

func (o MultisiteSiteAssociationArrayOutput) ToMultisiteSiteAssociationArrayOutput() MultisiteSiteAssociationArrayOutput {
	return o
}

func (o MultisiteSiteAssociationArrayOutput) ToMultisiteSiteAssociationArrayOutputWithContext(ctx context.Context) MultisiteSiteAssociationArrayOutput {
	return o
}

func (o MultisiteSiteAssociationArrayOutput) Index(i pulumi.IntInput) MultisiteSiteAssociationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MultisiteSiteAssociation {
		return vs[0].([]*MultisiteSiteAssociation)[vs[1].(int)]
	}).(MultisiteSiteAssociationOutput)
}

type MultisiteSiteAssociationMapOutput struct{ *pulumi.OutputState }

func (MultisiteSiteAssociationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MultisiteSiteAssociation)(nil)).Elem()
}

func (o MultisiteSiteAssociationMapOutput) ToMultisiteSiteAssociationMapOutput() MultisiteSiteAssociationMapOutput {
	return o
}

func (o MultisiteSiteAssociationMapOutput) ToMultisiteSiteAssociationMapOutputWithContext(ctx context.Context) MultisiteSiteAssociationMapOutput {
	return o
}

func (o MultisiteSiteAssociationMapOutput) MapIndex(k pulumi.StringInput) MultisiteSiteAssociationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MultisiteSiteAssociation {
		return vs[0].(map[string]*MultisiteSiteAssociation)[vs[1].(string)]
	}).(MultisiteSiteAssociationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MultisiteSiteAssociationInput)(nil)).Elem(), &MultisiteSiteAssociation{})
	pulumi.RegisterInputType(reflect.TypeOf((*MultisiteSiteAssociationArrayInput)(nil)).Elem(), MultisiteSiteAssociationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MultisiteSiteAssociationMapInput)(nil)).Elem(), MultisiteSiteAssociationMap{})
	pulumi.RegisterOutputType(MultisiteSiteAssociationOutput{})
	pulumi.RegisterOutputType(MultisiteSiteAssociationArrayOutput{})
	pulumi.RegisterOutputType(MultisiteSiteAssociationMapOutput{})
}
