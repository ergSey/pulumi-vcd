// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/ergSey/pulumi-vcd/sdk/go/vcd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Supported in provider *v3.8+*.
//
// Provides a data source to read LDAP configuration for an organization.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/ergSey/pulumi-vcd/sdk/go/vcd"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			my_org, err := vcd.LookupOrg(ctx, &vcd.LookupOrgArgs{
//				Name: "my-org",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vcd.LookupOrgLdap(ctx, &vcd.LookupOrgLdapArgs{
//				OrgId: my_org.Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupOrgLdap(ctx *pulumi.Context, args *LookupOrgLdapArgs, opts ...pulumi.InvokeOption) (*LookupOrgLdapResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupOrgLdapResult
	err := ctx.Invoke("vcd:index/getOrgLdap:getOrgLdap", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrgLdap.
type LookupOrgLdapArgs struct {
	// ID of the organization containing the LDAP settings
	OrgId string `pulumi:"orgId"`
}

// A collection of values returned by getOrgLdap.
type LookupOrgLdapResult struct {
	CustomSettings []GetOrgLdapCustomSetting `pulumi:"customSettings"`
	CustomUserOu   string                    `pulumi:"customUserOu"`
	// The provider-assigned unique ID for this managed resource.
	Id       string `pulumi:"id"`
	LdapMode string `pulumi:"ldapMode"`
	OrgId    string `pulumi:"orgId"`
}

func LookupOrgLdapOutput(ctx *pulumi.Context, args LookupOrgLdapOutputArgs, opts ...pulumi.InvokeOption) LookupOrgLdapResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupOrgLdapResultOutput, error) {
			args := v.(LookupOrgLdapArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vcd:index/getOrgLdap:getOrgLdap", args, LookupOrgLdapResultOutput{}, options).(LookupOrgLdapResultOutput), nil
		}).(LookupOrgLdapResultOutput)
}

// A collection of arguments for invoking getOrgLdap.
type LookupOrgLdapOutputArgs struct {
	// ID of the organization containing the LDAP settings
	OrgId pulumi.StringInput `pulumi:"orgId"`
}

func (LookupOrgLdapOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrgLdapArgs)(nil)).Elem()
}

// A collection of values returned by getOrgLdap.
type LookupOrgLdapResultOutput struct{ *pulumi.OutputState }

func (LookupOrgLdapResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrgLdapResult)(nil)).Elem()
}

func (o LookupOrgLdapResultOutput) ToLookupOrgLdapResultOutput() LookupOrgLdapResultOutput {
	return o
}

func (o LookupOrgLdapResultOutput) ToLookupOrgLdapResultOutputWithContext(ctx context.Context) LookupOrgLdapResultOutput {
	return o
}

func (o LookupOrgLdapResultOutput) CustomSettings() GetOrgLdapCustomSettingArrayOutput {
	return o.ApplyT(func(v LookupOrgLdapResult) []GetOrgLdapCustomSetting { return v.CustomSettings }).(GetOrgLdapCustomSettingArrayOutput)
}

func (o LookupOrgLdapResultOutput) CustomUserOu() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgLdapResult) string { return v.CustomUserOu }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupOrgLdapResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgLdapResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupOrgLdapResultOutput) LdapMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgLdapResult) string { return v.LdapMode }).(pulumi.StringOutput)
}

func (o LookupOrgLdapResultOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgLdapResult) string { return v.OrgId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrgLdapResultOutput{})
}
