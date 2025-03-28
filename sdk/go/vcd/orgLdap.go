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

type OrgLdap struct {
	pulumi.CustomResourceState

	// LDAP server configuration. Becomes mandatory if `ldapMode` is set to `CUSTOM`. See Custom Settings below for details
	//
	// <a id="custom-settings"></a>
	CustomSettings OrgLdapCustomSettingsPtrOutput `pulumi:"customSettings"`
	// If `ldapMode` is `SYSTEM`, specifies an LDAP `attribute=value` pair to use for OU (organizational unit)
	CustomUserOu pulumi.StringPtrOutput `pulumi:"customUserOu"`
	// One of `NONE`, `CUSTOM`, `SYSTEM`. Note that using `NONE` has the effect of removing the LDAP settings
	LdapMode pulumi.StringOutput `pulumi:"ldapMode"`
	// Org ID: there is only one LDAP configuration available for an organization. Thus, the resource can be identified by the Org.
	OrgId pulumi.StringOutput `pulumi:"orgId"`
}

// NewOrgLdap registers a new resource with the given unique name, arguments, and options.
func NewOrgLdap(ctx *pulumi.Context,
	name string, args *OrgLdapArgs, opts ...pulumi.ResourceOption) (*OrgLdap, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LdapMode == nil {
		return nil, errors.New("invalid value for required argument 'LdapMode'")
	}
	if args.OrgId == nil {
		return nil, errors.New("invalid value for required argument 'OrgId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource OrgLdap
	err := ctx.RegisterResource("vcd:index/orgLdap:OrgLdap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrgLdap gets an existing OrgLdap resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrgLdap(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrgLdapState, opts ...pulumi.ResourceOption) (*OrgLdap, error) {
	var resource OrgLdap
	err := ctx.ReadResource("vcd:index/orgLdap:OrgLdap", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrgLdap resources.
type orgLdapState struct {
	// LDAP server configuration. Becomes mandatory if `ldapMode` is set to `CUSTOM`. See Custom Settings below for details
	//
	// <a id="custom-settings"></a>
	CustomSettings *OrgLdapCustomSettings `pulumi:"customSettings"`
	// If `ldapMode` is `SYSTEM`, specifies an LDAP `attribute=value` pair to use for OU (organizational unit)
	CustomUserOu *string `pulumi:"customUserOu"`
	// One of `NONE`, `CUSTOM`, `SYSTEM`. Note that using `NONE` has the effect of removing the LDAP settings
	LdapMode *string `pulumi:"ldapMode"`
	// Org ID: there is only one LDAP configuration available for an organization. Thus, the resource can be identified by the Org.
	OrgId *string `pulumi:"orgId"`
}

type OrgLdapState struct {
	// LDAP server configuration. Becomes mandatory if `ldapMode` is set to `CUSTOM`. See Custom Settings below for details
	//
	// <a id="custom-settings"></a>
	CustomSettings OrgLdapCustomSettingsPtrInput
	// If `ldapMode` is `SYSTEM`, specifies an LDAP `attribute=value` pair to use for OU (organizational unit)
	CustomUserOu pulumi.StringPtrInput
	// One of `NONE`, `CUSTOM`, `SYSTEM`. Note that using `NONE` has the effect of removing the LDAP settings
	LdapMode pulumi.StringPtrInput
	// Org ID: there is only one LDAP configuration available for an organization. Thus, the resource can be identified by the Org.
	OrgId pulumi.StringPtrInput
}

func (OrgLdapState) ElementType() reflect.Type {
	return reflect.TypeOf((*orgLdapState)(nil)).Elem()
}

type orgLdapArgs struct {
	// LDAP server configuration. Becomes mandatory if `ldapMode` is set to `CUSTOM`. See Custom Settings below for details
	//
	// <a id="custom-settings"></a>
	CustomSettings *OrgLdapCustomSettings `pulumi:"customSettings"`
	// If `ldapMode` is `SYSTEM`, specifies an LDAP `attribute=value` pair to use for OU (organizational unit)
	CustomUserOu *string `pulumi:"customUserOu"`
	// One of `NONE`, `CUSTOM`, `SYSTEM`. Note that using `NONE` has the effect of removing the LDAP settings
	LdapMode string `pulumi:"ldapMode"`
	// Org ID: there is only one LDAP configuration available for an organization. Thus, the resource can be identified by the Org.
	OrgId string `pulumi:"orgId"`
}

// The set of arguments for constructing a OrgLdap resource.
type OrgLdapArgs struct {
	// LDAP server configuration. Becomes mandatory if `ldapMode` is set to `CUSTOM`. See Custom Settings below for details
	//
	// <a id="custom-settings"></a>
	CustomSettings OrgLdapCustomSettingsPtrInput
	// If `ldapMode` is `SYSTEM`, specifies an LDAP `attribute=value` pair to use for OU (organizational unit)
	CustomUserOu pulumi.StringPtrInput
	// One of `NONE`, `CUSTOM`, `SYSTEM`. Note that using `NONE` has the effect of removing the LDAP settings
	LdapMode pulumi.StringInput
	// Org ID: there is only one LDAP configuration available for an organization. Thus, the resource can be identified by the Org.
	OrgId pulumi.StringInput
}

func (OrgLdapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*orgLdapArgs)(nil)).Elem()
}

type OrgLdapInput interface {
	pulumi.Input

	ToOrgLdapOutput() OrgLdapOutput
	ToOrgLdapOutputWithContext(ctx context.Context) OrgLdapOutput
}

func (*OrgLdap) ElementType() reflect.Type {
	return reflect.TypeOf((**OrgLdap)(nil)).Elem()
}

func (i *OrgLdap) ToOrgLdapOutput() OrgLdapOutput {
	return i.ToOrgLdapOutputWithContext(context.Background())
}

func (i *OrgLdap) ToOrgLdapOutputWithContext(ctx context.Context) OrgLdapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgLdapOutput)
}

// OrgLdapArrayInput is an input type that accepts OrgLdapArray and OrgLdapArrayOutput values.
// You can construct a concrete instance of `OrgLdapArrayInput` via:
//
//	OrgLdapArray{ OrgLdapArgs{...} }
type OrgLdapArrayInput interface {
	pulumi.Input

	ToOrgLdapArrayOutput() OrgLdapArrayOutput
	ToOrgLdapArrayOutputWithContext(context.Context) OrgLdapArrayOutput
}

type OrgLdapArray []OrgLdapInput

func (OrgLdapArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrgLdap)(nil)).Elem()
}

func (i OrgLdapArray) ToOrgLdapArrayOutput() OrgLdapArrayOutput {
	return i.ToOrgLdapArrayOutputWithContext(context.Background())
}

func (i OrgLdapArray) ToOrgLdapArrayOutputWithContext(ctx context.Context) OrgLdapArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgLdapArrayOutput)
}

// OrgLdapMapInput is an input type that accepts OrgLdapMap and OrgLdapMapOutput values.
// You can construct a concrete instance of `OrgLdapMapInput` via:
//
//	OrgLdapMap{ "key": OrgLdapArgs{...} }
type OrgLdapMapInput interface {
	pulumi.Input

	ToOrgLdapMapOutput() OrgLdapMapOutput
	ToOrgLdapMapOutputWithContext(context.Context) OrgLdapMapOutput
}

type OrgLdapMap map[string]OrgLdapInput

func (OrgLdapMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrgLdap)(nil)).Elem()
}

func (i OrgLdapMap) ToOrgLdapMapOutput() OrgLdapMapOutput {
	return i.ToOrgLdapMapOutputWithContext(context.Background())
}

func (i OrgLdapMap) ToOrgLdapMapOutputWithContext(ctx context.Context) OrgLdapMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrgLdapMapOutput)
}

type OrgLdapOutput struct{ *pulumi.OutputState }

func (OrgLdapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OrgLdap)(nil)).Elem()
}

func (o OrgLdapOutput) ToOrgLdapOutput() OrgLdapOutput {
	return o
}

func (o OrgLdapOutput) ToOrgLdapOutputWithContext(ctx context.Context) OrgLdapOutput {
	return o
}

// LDAP server configuration. Becomes mandatory if `ldapMode` is set to `CUSTOM`. See Custom Settings below for details
//
// <a id="custom-settings"></a>
func (o OrgLdapOutput) CustomSettings() OrgLdapCustomSettingsPtrOutput {
	return o.ApplyT(func(v *OrgLdap) OrgLdapCustomSettingsPtrOutput { return v.CustomSettings }).(OrgLdapCustomSettingsPtrOutput)
}

// If `ldapMode` is `SYSTEM`, specifies an LDAP `attribute=value` pair to use for OU (organizational unit)
func (o OrgLdapOutput) CustomUserOu() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OrgLdap) pulumi.StringPtrOutput { return v.CustomUserOu }).(pulumi.StringPtrOutput)
}

// One of `NONE`, `CUSTOM`, `SYSTEM`. Note that using `NONE` has the effect of removing the LDAP settings
func (o OrgLdapOutput) LdapMode() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgLdap) pulumi.StringOutput { return v.LdapMode }).(pulumi.StringOutput)
}

// Org ID: there is only one LDAP configuration available for an organization. Thus, the resource can be identified by the Org.
func (o OrgLdapOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v *OrgLdap) pulumi.StringOutput { return v.OrgId }).(pulumi.StringOutput)
}

type OrgLdapArrayOutput struct{ *pulumi.OutputState }

func (OrgLdapArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*OrgLdap)(nil)).Elem()
}

func (o OrgLdapArrayOutput) ToOrgLdapArrayOutput() OrgLdapArrayOutput {
	return o
}

func (o OrgLdapArrayOutput) ToOrgLdapArrayOutputWithContext(ctx context.Context) OrgLdapArrayOutput {
	return o
}

func (o OrgLdapArrayOutput) Index(i pulumi.IntInput) OrgLdapOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *OrgLdap {
		return vs[0].([]*OrgLdap)[vs[1].(int)]
	}).(OrgLdapOutput)
}

type OrgLdapMapOutput struct{ *pulumi.OutputState }

func (OrgLdapMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*OrgLdap)(nil)).Elem()
}

func (o OrgLdapMapOutput) ToOrgLdapMapOutput() OrgLdapMapOutput {
	return o
}

func (o OrgLdapMapOutput) ToOrgLdapMapOutputWithContext(ctx context.Context) OrgLdapMapOutput {
	return o
}

func (o OrgLdapMapOutput) MapIndex(k pulumi.StringInput) OrgLdapOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *OrgLdap {
		return vs[0].(map[string]*OrgLdap)[vs[1].(string)]
	}).(OrgLdapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OrgLdapInput)(nil)).Elem(), &OrgLdap{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrgLdapArrayInput)(nil)).Elem(), OrgLdapArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*OrgLdapMapInput)(nil)).Elem(), OrgLdapMap{})
	pulumi.RegisterOutputType(OrgLdapOutput{})
	pulumi.RegisterOutputType(OrgLdapArrayOutput{})
	pulumi.RegisterOutputType(OrgLdapMapOutput{})
}
