// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/ergSey/pulumi-vcd/sdk/go/vcd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VMware Cloud Director Org VDC isolated Network data source to read data or reference existing network.
//
// Supported in provider *v3.2+* for both NSX-T and NSX-V VDCs.
//
// ## Example Usage
//
// ### Looking Up Isolated Network In VDC)
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
//			main, err := vcd.LookupOrgVdc(ctx, &vcd.LookupOrgVdcArgs{
//				Org:  pulumi.StringRef("my-org"),
//				Name: "main-edge",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vcd.LookupNetworkIsolatedV2(ctx, &vcd.LookupNetworkIsolatedV2Args{
//				Org:     pulumi.StringRef("my-org"),
//				OwnerId: pulumi.StringRef(main.Id),
//				Name:    pulumi.StringRef("my-net"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ### Looking Up Isolated Network In VDC Group)
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
//			main, err := vcd.LookupVdcGroup(ctx, &vcd.LookupVdcGroupArgs{
//				Org:  pulumi.StringRef("my-org"),
//				Name: pulumi.StringRef("main-group"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vcd.LookupNetworkIsolatedV2(ctx, &vcd.LookupNetworkIsolatedV2Args{
//				Org:     pulumi.StringRef("my-org"),
//				OwnerId: pulumi.StringRef(main.Id),
//				Name:    pulumi.StringRef("my-net"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Filter arguments
//
// * `nameRegex` - (Optional) matches the name using a regular expression.
// * `ip` - (Optional) matches the IP of the resource using a regular expression.
//
// See [Filters reference](https://www.terraform.io/providers/vmware/vcd/latest/docs/guides/data_source_filters) for details and examples.
func LookupNetworkIsolatedV2(ctx *pulumi.Context, args *LookupNetworkIsolatedV2Args, opts ...pulumi.InvokeOption) (*LookupNetworkIsolatedV2Result, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNetworkIsolatedV2Result
	err := ctx.Invoke("vcd:index/getNetworkIsolatedV2:getNetworkIsolatedV2", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNetworkIsolatedV2.
type LookupNetworkIsolatedV2Args struct {
	// Retrieves the data source using one or more filter parameters. **Note**
	// filters do not support searching for networks in VDC Groups.
	Filter *GetNetworkIsolatedV2Filter `pulumi:"filter"`
	// A unique name for the network (optional when `filter` is used)
	Name *string `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level
	Org *string `pulumi:"org"`
	// VDC or VDC Group ID. Always takes precedence over `vdc` fields (in resource
	// and inherited from provider configuration)
	OwnerId *string `pulumi:"ownerId"`
	// The name of VDC to use. **Deprecated**  in favor of new field
	// `ownerId` which supports VDC and VDC Group IDs.
	//
	// Deprecated: This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
	Vdc *string `pulumi:"vdc"`
}

// A collection of values returned by getNetworkIsolatedV2.
type LookupNetworkIsolatedV2Result struct {
	Description      string                      `pulumi:"description"`
	Dns1             string                      `pulumi:"dns1"`
	Dns2             string                      `pulumi:"dns2"`
	DnsSuffix        string                      `pulumi:"dnsSuffix"`
	DualStackEnabled bool                        `pulumi:"dualStackEnabled"`
	Filter           *GetNetworkIsolatedV2Filter `pulumi:"filter"`
	Gateway          string                      `pulumi:"gateway"`
	GuestVlanAllowed bool                        `pulumi:"guestVlanAllowed"`
	// The provider-assigned unique ID for this managed resource.
	Id       string `pulumi:"id"`
	IsShared bool   `pulumi:"isShared"`
	// Deprecated: Use metadataEntry instead
	Metadata               map[string]string                           `pulumi:"metadata"`
	MetadataEntries        []GetNetworkIsolatedV2MetadataEntry         `pulumi:"metadataEntries"`
	Name                   *string                                     `pulumi:"name"`
	Org                    *string                                     `pulumi:"org"`
	OwnerId                string                                      `pulumi:"ownerId"`
	PrefixLength           int                                         `pulumi:"prefixLength"`
	SecondaryGateway       string                                      `pulumi:"secondaryGateway"`
	SecondaryPrefixLength  string                                      `pulumi:"secondaryPrefixLength"`
	SecondaryStaticIpPools []GetNetworkIsolatedV2SecondaryStaticIpPool `pulumi:"secondaryStaticIpPools"`
	StaticIpPools          []GetNetworkIsolatedV2StaticIpPool          `pulumi:"staticIpPools"`
	// Deprecated: This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
	Vdc string `pulumi:"vdc"`
}

func LookupNetworkIsolatedV2Output(ctx *pulumi.Context, args LookupNetworkIsolatedV2OutputArgs, opts ...pulumi.InvokeOption) LookupNetworkIsolatedV2ResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupNetworkIsolatedV2ResultOutput, error) {
			args := v.(LookupNetworkIsolatedV2Args)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vcd:index/getNetworkIsolatedV2:getNetworkIsolatedV2", args, LookupNetworkIsolatedV2ResultOutput{}, options).(LookupNetworkIsolatedV2ResultOutput), nil
		}).(LookupNetworkIsolatedV2ResultOutput)
}

// A collection of arguments for invoking getNetworkIsolatedV2.
type LookupNetworkIsolatedV2OutputArgs struct {
	// Retrieves the data source using one or more filter parameters. **Note**
	// filters do not support searching for networks in VDC Groups.
	Filter GetNetworkIsolatedV2FilterPtrInput `pulumi:"filter"`
	// A unique name for the network (optional when `filter` is used)
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level
	Org pulumi.StringPtrInput `pulumi:"org"`
	// VDC or VDC Group ID. Always takes precedence over `vdc` fields (in resource
	// and inherited from provider configuration)
	OwnerId pulumi.StringPtrInput `pulumi:"ownerId"`
	// The name of VDC to use. **Deprecated**  in favor of new field
	// `ownerId` which supports VDC and VDC Group IDs.
	//
	// Deprecated: This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
	Vdc pulumi.StringPtrInput `pulumi:"vdc"`
}

func (LookupNetworkIsolatedV2OutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkIsolatedV2Args)(nil)).Elem()
}

// A collection of values returned by getNetworkIsolatedV2.
type LookupNetworkIsolatedV2ResultOutput struct{ *pulumi.OutputState }

func (LookupNetworkIsolatedV2ResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNetworkIsolatedV2Result)(nil)).Elem()
}

func (o LookupNetworkIsolatedV2ResultOutput) ToLookupNetworkIsolatedV2ResultOutput() LookupNetworkIsolatedV2ResultOutput {
	return o
}

func (o LookupNetworkIsolatedV2ResultOutput) ToLookupNetworkIsolatedV2ResultOutputWithContext(ctx context.Context) LookupNetworkIsolatedV2ResultOutput {
	return o
}

func (o LookupNetworkIsolatedV2ResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkIsolatedV2Result) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupNetworkIsolatedV2ResultOutput) Dns1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkIsolatedV2Result) string { return v.Dns1 }).(pulumi.StringOutput)
}

func (o LookupNetworkIsolatedV2ResultOutput) Dns2() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkIsolatedV2Result) string { return v.Dns2 }).(pulumi.StringOutput)
}

func (o LookupNetworkIsolatedV2ResultOutput) DnsSuffix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkIsolatedV2Result) string { return v.DnsSuffix }).(pulumi.StringOutput)
}

func (o LookupNetworkIsolatedV2ResultOutput) DualStackEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNetworkIsolatedV2Result) bool { return v.DualStackEnabled }).(pulumi.BoolOutput)
}

func (o LookupNetworkIsolatedV2ResultOutput) Filter() GetNetworkIsolatedV2FilterPtrOutput {
	return o.ApplyT(func(v LookupNetworkIsolatedV2Result) *GetNetworkIsolatedV2Filter { return v.Filter }).(GetNetworkIsolatedV2FilterPtrOutput)
}

func (o LookupNetworkIsolatedV2ResultOutput) Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkIsolatedV2Result) string { return v.Gateway }).(pulumi.StringOutput)
}

func (o LookupNetworkIsolatedV2ResultOutput) GuestVlanAllowed() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNetworkIsolatedV2Result) bool { return v.GuestVlanAllowed }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNetworkIsolatedV2ResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkIsolatedV2Result) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNetworkIsolatedV2ResultOutput) IsShared() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNetworkIsolatedV2Result) bool { return v.IsShared }).(pulumi.BoolOutput)
}

// Deprecated: Use metadataEntry instead
func (o LookupNetworkIsolatedV2ResultOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupNetworkIsolatedV2Result) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

func (o LookupNetworkIsolatedV2ResultOutput) MetadataEntries() GetNetworkIsolatedV2MetadataEntryArrayOutput {
	return o.ApplyT(func(v LookupNetworkIsolatedV2Result) []GetNetworkIsolatedV2MetadataEntry { return v.MetadataEntries }).(GetNetworkIsolatedV2MetadataEntryArrayOutput)
}

func (o LookupNetworkIsolatedV2ResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkIsolatedV2Result) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkIsolatedV2ResultOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNetworkIsolatedV2Result) *string { return v.Org }).(pulumi.StringPtrOutput)
}

func (o LookupNetworkIsolatedV2ResultOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkIsolatedV2Result) string { return v.OwnerId }).(pulumi.StringOutput)
}

func (o LookupNetworkIsolatedV2ResultOutput) PrefixLength() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNetworkIsolatedV2Result) int { return v.PrefixLength }).(pulumi.IntOutput)
}

func (o LookupNetworkIsolatedV2ResultOutput) SecondaryGateway() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkIsolatedV2Result) string { return v.SecondaryGateway }).(pulumi.StringOutput)
}

func (o LookupNetworkIsolatedV2ResultOutput) SecondaryPrefixLength() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkIsolatedV2Result) string { return v.SecondaryPrefixLength }).(pulumi.StringOutput)
}

func (o LookupNetworkIsolatedV2ResultOutput) SecondaryStaticIpPools() GetNetworkIsolatedV2SecondaryStaticIpPoolArrayOutput {
	return o.ApplyT(func(v LookupNetworkIsolatedV2Result) []GetNetworkIsolatedV2SecondaryStaticIpPool {
		return v.SecondaryStaticIpPools
	}).(GetNetworkIsolatedV2SecondaryStaticIpPoolArrayOutput)
}

func (o LookupNetworkIsolatedV2ResultOutput) StaticIpPools() GetNetworkIsolatedV2StaticIpPoolArrayOutput {
	return o.ApplyT(func(v LookupNetworkIsolatedV2Result) []GetNetworkIsolatedV2StaticIpPool { return v.StaticIpPools }).(GetNetworkIsolatedV2StaticIpPoolArrayOutput)
}

// Deprecated: This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
func (o LookupNetworkIsolatedV2ResultOutput) Vdc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNetworkIsolatedV2Result) string { return v.Vdc }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNetworkIsolatedV2ResultOutput{})
}
