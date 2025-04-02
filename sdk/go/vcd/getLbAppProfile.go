// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/ergSey/pulumi-vcd/sdk/go/vcd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VMware Cloud Director Edge Gateway Load Balancer Application Profile data source. An
// application profile defines the behavior of the load balancer for a particular type of network
// traffic. After configuring a profile, you associate it with a virtual server. The virtual server
// then processes traffic according to the values specified in the profile.
//
// > **Note:** See additional support notes in [application profile resource page]
// (/providers/vmware/vcd/latest/docs/resources/lb_app_profile).
//
// Supported in provider *v2.4+*
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
//			_, err := vcd.LookupLbAppProfile(ctx, &vcd.LookupLbAppProfileArgs{
//				Org:         pulumi.StringRef("my-org"),
//				Vdc:         pulumi.StringRef("my-org-vdc"),
//				EdgeGateway: "my-edge-gw",
//				Name:        "not-managed",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupLbAppProfile(ctx *pulumi.Context, args *LookupLbAppProfileArgs, opts ...pulumi.InvokeOption) (*LookupLbAppProfileResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupLbAppProfileResult
	err := ctx.Invoke("vcd:index/getLbAppProfile:getLbAppProfile", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLbAppProfile.
type LookupLbAppProfileArgs struct {
	// The name of the edge gateway on which the service monitor is defined
	EdgeGateway string `pulumi:"edgeGateway"`
	// Application profile name for identifying the exact application profile
	Name string `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level
	Org *string `pulumi:"org"`
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `pulumi:"vdc"`
}

// A collection of values returned by getLbAppProfile.
type LookupLbAppProfileResult struct {
	CookieMode           string `pulumi:"cookieMode"`
	CookieName           string `pulumi:"cookieName"`
	EdgeGateway          string `pulumi:"edgeGateway"`
	EnablePoolSideSsl    bool   `pulumi:"enablePoolSideSsl"`
	EnableSslPassthrough bool   `pulumi:"enableSslPassthrough"`
	Expiration           int    `pulumi:"expiration"`
	HttpRedirectUrl      string `pulumi:"httpRedirectUrl"`
	// The provider-assigned unique ID for this managed resource.
	Id                         string  `pulumi:"id"`
	InsertXForwardedHttpHeader bool    `pulumi:"insertXForwardedHttpHeader"`
	Name                       string  `pulumi:"name"`
	Org                        *string `pulumi:"org"`
	PersistenceMechanism       string  `pulumi:"persistenceMechanism"`
	Type                       string  `pulumi:"type"`
	Vdc                        *string `pulumi:"vdc"`
}

func LookupLbAppProfileOutput(ctx *pulumi.Context, args LookupLbAppProfileOutputArgs, opts ...pulumi.InvokeOption) LookupLbAppProfileResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupLbAppProfileResultOutput, error) {
			args := v.(LookupLbAppProfileArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vcd:index/getLbAppProfile:getLbAppProfile", args, LookupLbAppProfileResultOutput{}, options).(LookupLbAppProfileResultOutput), nil
		}).(LookupLbAppProfileResultOutput)
}

// A collection of arguments for invoking getLbAppProfile.
type LookupLbAppProfileOutputArgs struct {
	// The name of the edge gateway on which the service monitor is defined
	EdgeGateway pulumi.StringInput `pulumi:"edgeGateway"`
	// Application profile name for identifying the exact application profile
	Name pulumi.StringInput `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level
	Org pulumi.StringPtrInput `pulumi:"org"`
	// The name of VDC to use, optional if defined at provider level
	Vdc pulumi.StringPtrInput `pulumi:"vdc"`
}

func (LookupLbAppProfileOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLbAppProfileArgs)(nil)).Elem()
}

// A collection of values returned by getLbAppProfile.
type LookupLbAppProfileResultOutput struct{ *pulumi.OutputState }

func (LookupLbAppProfileResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLbAppProfileResult)(nil)).Elem()
}

func (o LookupLbAppProfileResultOutput) ToLookupLbAppProfileResultOutput() LookupLbAppProfileResultOutput {
	return o
}

func (o LookupLbAppProfileResultOutput) ToLookupLbAppProfileResultOutputWithContext(ctx context.Context) LookupLbAppProfileResultOutput {
	return o
}

func (o LookupLbAppProfileResultOutput) CookieMode() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbAppProfileResult) string { return v.CookieMode }).(pulumi.StringOutput)
}

func (o LookupLbAppProfileResultOutput) CookieName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbAppProfileResult) string { return v.CookieName }).(pulumi.StringOutput)
}

func (o LookupLbAppProfileResultOutput) EdgeGateway() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbAppProfileResult) string { return v.EdgeGateway }).(pulumi.StringOutput)
}

func (o LookupLbAppProfileResultOutput) EnablePoolSideSsl() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupLbAppProfileResult) bool { return v.EnablePoolSideSsl }).(pulumi.BoolOutput)
}

func (o LookupLbAppProfileResultOutput) EnableSslPassthrough() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupLbAppProfileResult) bool { return v.EnableSslPassthrough }).(pulumi.BoolOutput)
}

func (o LookupLbAppProfileResultOutput) Expiration() pulumi.IntOutput {
	return o.ApplyT(func(v LookupLbAppProfileResult) int { return v.Expiration }).(pulumi.IntOutput)
}

func (o LookupLbAppProfileResultOutput) HttpRedirectUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbAppProfileResult) string { return v.HttpRedirectUrl }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupLbAppProfileResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbAppProfileResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLbAppProfileResultOutput) InsertXForwardedHttpHeader() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupLbAppProfileResult) bool { return v.InsertXForwardedHttpHeader }).(pulumi.BoolOutput)
}

func (o LookupLbAppProfileResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbAppProfileResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLbAppProfileResultOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLbAppProfileResult) *string { return v.Org }).(pulumi.StringPtrOutput)
}

func (o LookupLbAppProfileResultOutput) PersistenceMechanism() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbAppProfileResult) string { return v.PersistenceMechanism }).(pulumi.StringOutput)
}

func (o LookupLbAppProfileResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbAppProfileResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupLbAppProfileResultOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLbAppProfileResult) *string { return v.Vdc }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLbAppProfileResultOutput{})
}
