// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/ergSey/pulumi-vcd/sdk/go/vcd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a data source to read the OpenID Connect (OIDC) configuration of an Organization in VMware Cloud Director.
//
// Supported in provider *v3.13+*.
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
//			myOrg, err := vcd.LookupOrg(ctx, &vcd.LookupOrgArgs{
//				Name: "my-org",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vcd.LookupOrgOidc(ctx, &vcd.LookupOrgOidcArgs{
//				OrgId: myOrg.Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupOrgOidc(ctx *pulumi.Context, args *LookupOrgOidcArgs, opts ...pulumi.InvokeOption) (*LookupOrgOidcResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupOrgOidcResult
	err := ctx.Invoke("vcd:index/getOrgOidc:getOrgOidc", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrgOidc.
type LookupOrgOidcArgs struct {
	// ID of the organization containing the OIDC settings
	OrgId string `pulumi:"orgId"`
}

// A collection of values returned by getOrgOidc.
type LookupOrgOidcResult struct {
	AccessTokenEndpoint string                    `pulumi:"accessTokenEndpoint"`
	ClaimsMappings      []GetOrgOidcClaimsMapping `pulumi:"claimsMappings"`
	ClientId            string                    `pulumi:"clientId"`
	ClientSecret        string                    `pulumi:"clientSecret"`
	Enabled             bool                      `pulumi:"enabled"`
	// The provider-assigned unique ID for this managed resource.
	Id                        string          `pulumi:"id"`
	IssuerId                  string          `pulumi:"issuerId"`
	KeyExpireDurationHours    int             `pulumi:"keyExpireDurationHours"`
	KeyRefreshEndpoint        string          `pulumi:"keyRefreshEndpoint"`
	KeyRefreshPeriodHours     int             `pulumi:"keyRefreshPeriodHours"`
	KeyRefreshStrategy        string          `pulumi:"keyRefreshStrategy"`
	Keys                      []GetOrgOidcKey `pulumi:"keys"`
	MaxClockSkewSeconds       int             `pulumi:"maxClockSkewSeconds"`
	OrgId                     string          `pulumi:"orgId"`
	PreferIdToken             bool            `pulumi:"preferIdToken"`
	RedirectUri               string          `pulumi:"redirectUri"`
	Scopes                    []string        `pulumi:"scopes"`
	UiButtonLabel             string          `pulumi:"uiButtonLabel"`
	UserAuthorizationEndpoint string          `pulumi:"userAuthorizationEndpoint"`
	UserinfoEndpoint          string          `pulumi:"userinfoEndpoint"`
	WellknownEndpoint         string          `pulumi:"wellknownEndpoint"`
}

func LookupOrgOidcOutput(ctx *pulumi.Context, args LookupOrgOidcOutputArgs, opts ...pulumi.InvokeOption) LookupOrgOidcResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupOrgOidcResultOutput, error) {
			args := v.(LookupOrgOidcArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vcd:index/getOrgOidc:getOrgOidc", args, LookupOrgOidcResultOutput{}, options).(LookupOrgOidcResultOutput), nil
		}).(LookupOrgOidcResultOutput)
}

// A collection of arguments for invoking getOrgOidc.
type LookupOrgOidcOutputArgs struct {
	// ID of the organization containing the OIDC settings
	OrgId pulumi.StringInput `pulumi:"orgId"`
}

func (LookupOrgOidcOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrgOidcArgs)(nil)).Elem()
}

// A collection of values returned by getOrgOidc.
type LookupOrgOidcResultOutput struct{ *pulumi.OutputState }

func (LookupOrgOidcResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrgOidcResult)(nil)).Elem()
}

func (o LookupOrgOidcResultOutput) ToLookupOrgOidcResultOutput() LookupOrgOidcResultOutput {
	return o
}

func (o LookupOrgOidcResultOutput) ToLookupOrgOidcResultOutputWithContext(ctx context.Context) LookupOrgOidcResultOutput {
	return o
}

func (o LookupOrgOidcResultOutput) AccessTokenEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgOidcResult) string { return v.AccessTokenEndpoint }).(pulumi.StringOutput)
}

func (o LookupOrgOidcResultOutput) ClaimsMappings() GetOrgOidcClaimsMappingArrayOutput {
	return o.ApplyT(func(v LookupOrgOidcResult) []GetOrgOidcClaimsMapping { return v.ClaimsMappings }).(GetOrgOidcClaimsMappingArrayOutput)
}

func (o LookupOrgOidcResultOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgOidcResult) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o LookupOrgOidcResultOutput) ClientSecret() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgOidcResult) string { return v.ClientSecret }).(pulumi.StringOutput)
}

func (o LookupOrgOidcResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOrgOidcResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupOrgOidcResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgOidcResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupOrgOidcResultOutput) IssuerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgOidcResult) string { return v.IssuerId }).(pulumi.StringOutput)
}

func (o LookupOrgOidcResultOutput) KeyExpireDurationHours() pulumi.IntOutput {
	return o.ApplyT(func(v LookupOrgOidcResult) int { return v.KeyExpireDurationHours }).(pulumi.IntOutput)
}

func (o LookupOrgOidcResultOutput) KeyRefreshEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgOidcResult) string { return v.KeyRefreshEndpoint }).(pulumi.StringOutput)
}

func (o LookupOrgOidcResultOutput) KeyRefreshPeriodHours() pulumi.IntOutput {
	return o.ApplyT(func(v LookupOrgOidcResult) int { return v.KeyRefreshPeriodHours }).(pulumi.IntOutput)
}

func (o LookupOrgOidcResultOutput) KeyRefreshStrategy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgOidcResult) string { return v.KeyRefreshStrategy }).(pulumi.StringOutput)
}

func (o LookupOrgOidcResultOutput) Keys() GetOrgOidcKeyArrayOutput {
	return o.ApplyT(func(v LookupOrgOidcResult) []GetOrgOidcKey { return v.Keys }).(GetOrgOidcKeyArrayOutput)
}

func (o LookupOrgOidcResultOutput) MaxClockSkewSeconds() pulumi.IntOutput {
	return o.ApplyT(func(v LookupOrgOidcResult) int { return v.MaxClockSkewSeconds }).(pulumi.IntOutput)
}

func (o LookupOrgOidcResultOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgOidcResult) string { return v.OrgId }).(pulumi.StringOutput)
}

func (o LookupOrgOidcResultOutput) PreferIdToken() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOrgOidcResult) bool { return v.PreferIdToken }).(pulumi.BoolOutput)
}

func (o LookupOrgOidcResultOutput) RedirectUri() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgOidcResult) string { return v.RedirectUri }).(pulumi.StringOutput)
}

func (o LookupOrgOidcResultOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupOrgOidcResult) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

func (o LookupOrgOidcResultOutput) UiButtonLabel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgOidcResult) string { return v.UiButtonLabel }).(pulumi.StringOutput)
}

func (o LookupOrgOidcResultOutput) UserAuthorizationEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgOidcResult) string { return v.UserAuthorizationEndpoint }).(pulumi.StringOutput)
}

func (o LookupOrgOidcResultOutput) UserinfoEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgOidcResult) string { return v.UserinfoEndpoint }).(pulumi.StringOutput)
}

func (o LookupOrgOidcResultOutput) WellknownEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgOidcResult) string { return v.WellknownEndpoint }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrgOidcResultOutput{})
}
