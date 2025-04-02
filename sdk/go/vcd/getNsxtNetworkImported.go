// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/ergSey/pulumi-vcd/sdk/go/vcd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## Example Usage
//
// ### Looking Up Imported Network In VDC)
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
//			_, err = vcd.LookupNsxtNetworkImported(ctx, &vcd.LookupNsxtNetworkImportedArgs{
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
// ### Looking Up Imported Network In VDC Group)
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
//			_, err = vcd.LookupNsxtNetworkImported(ctx, &vcd.LookupNsxtNetworkImportedArgs{
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
func LookupNsxtNetworkImported(ctx *pulumi.Context, args *LookupNsxtNetworkImportedArgs, opts ...pulumi.InvokeOption) (*LookupNsxtNetworkImportedResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNsxtNetworkImportedResult
	err := ctx.Invoke("vcd:index/getNsxtNetworkImported:getNsxtNetworkImported", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNsxtNetworkImported.
type LookupNsxtNetworkImportedArgs struct {
	// Retrieves the data source using one or more filter parameters
	Filter *GetNsxtNetworkImportedFilter `pulumi:"filter"`
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

// A collection of values returned by getNsxtNetworkImported.
type LookupNsxtNetworkImportedResult struct {
	Description      string                        `pulumi:"description"`
	Dns1             string                        `pulumi:"dns1"`
	Dns2             string                        `pulumi:"dns2"`
	DnsSuffix        string                        `pulumi:"dnsSuffix"`
	DualStackEnabled bool                          `pulumi:"dualStackEnabled"`
	DvpgId           string                        `pulumi:"dvpgId"`
	Filter           *GetNsxtNetworkImportedFilter `pulumi:"filter"`
	Gateway          string                        `pulumi:"gateway"`
	// The provider-assigned unique ID for this managed resource.
	Id                     string                                        `pulumi:"id"`
	Name                   *string                                       `pulumi:"name"`
	NsxtLogicalSwitchId    string                                        `pulumi:"nsxtLogicalSwitchId"`
	Org                    *string                                       `pulumi:"org"`
	OwnerId                string                                        `pulumi:"ownerId"`
	PrefixLength           int                                           `pulumi:"prefixLength"`
	SecondaryGateway       string                                        `pulumi:"secondaryGateway"`
	SecondaryPrefixLength  string                                        `pulumi:"secondaryPrefixLength"`
	SecondaryStaticIpPools []GetNsxtNetworkImportedSecondaryStaticIpPool `pulumi:"secondaryStaticIpPools"`
	StaticIpPools          []GetNsxtNetworkImportedStaticIpPool          `pulumi:"staticIpPools"`
	// Deprecated: This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
	Vdc *string `pulumi:"vdc"`
}

func LookupNsxtNetworkImportedOutput(ctx *pulumi.Context, args LookupNsxtNetworkImportedOutputArgs, opts ...pulumi.InvokeOption) LookupNsxtNetworkImportedResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupNsxtNetworkImportedResultOutput, error) {
			args := v.(LookupNsxtNetworkImportedArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vcd:index/getNsxtNetworkImported:getNsxtNetworkImported", args, LookupNsxtNetworkImportedResultOutput{}, options).(LookupNsxtNetworkImportedResultOutput), nil
		}).(LookupNsxtNetworkImportedResultOutput)
}

// A collection of arguments for invoking getNsxtNetworkImported.
type LookupNsxtNetworkImportedOutputArgs struct {
	// Retrieves the data source using one or more filter parameters
	Filter GetNsxtNetworkImportedFilterPtrInput `pulumi:"filter"`
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

func (LookupNsxtNetworkImportedOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxtNetworkImportedArgs)(nil)).Elem()
}

// A collection of values returned by getNsxtNetworkImported.
type LookupNsxtNetworkImportedResultOutput struct{ *pulumi.OutputState }

func (LookupNsxtNetworkImportedResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxtNetworkImportedResult)(nil)).Elem()
}

func (o LookupNsxtNetworkImportedResultOutput) ToLookupNsxtNetworkImportedResultOutput() LookupNsxtNetworkImportedResultOutput {
	return o
}

func (o LookupNsxtNetworkImportedResultOutput) ToLookupNsxtNetworkImportedResultOutputWithContext(ctx context.Context) LookupNsxtNetworkImportedResultOutput {
	return o
}

func (o LookupNsxtNetworkImportedResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtNetworkImportedResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupNsxtNetworkImportedResultOutput) Dns1() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtNetworkImportedResult) string { return v.Dns1 }).(pulumi.StringOutput)
}

func (o LookupNsxtNetworkImportedResultOutput) Dns2() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtNetworkImportedResult) string { return v.Dns2 }).(pulumi.StringOutput)
}

func (o LookupNsxtNetworkImportedResultOutput) DnsSuffix() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtNetworkImportedResult) string { return v.DnsSuffix }).(pulumi.StringOutput)
}

func (o LookupNsxtNetworkImportedResultOutput) DualStackEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupNsxtNetworkImportedResult) bool { return v.DualStackEnabled }).(pulumi.BoolOutput)
}

func (o LookupNsxtNetworkImportedResultOutput) DvpgId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtNetworkImportedResult) string { return v.DvpgId }).(pulumi.StringOutput)
}

func (o LookupNsxtNetworkImportedResultOutput) Filter() GetNsxtNetworkImportedFilterPtrOutput {
	return o.ApplyT(func(v LookupNsxtNetworkImportedResult) *GetNsxtNetworkImportedFilter { return v.Filter }).(GetNsxtNetworkImportedFilterPtrOutput)
}

func (o LookupNsxtNetworkImportedResultOutput) Gateway() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtNetworkImportedResult) string { return v.Gateway }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNsxtNetworkImportedResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtNetworkImportedResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNsxtNetworkImportedResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNsxtNetworkImportedResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupNsxtNetworkImportedResultOutput) NsxtLogicalSwitchId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtNetworkImportedResult) string { return v.NsxtLogicalSwitchId }).(pulumi.StringOutput)
}

func (o LookupNsxtNetworkImportedResultOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNsxtNetworkImportedResult) *string { return v.Org }).(pulumi.StringPtrOutput)
}

func (o LookupNsxtNetworkImportedResultOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtNetworkImportedResult) string { return v.OwnerId }).(pulumi.StringOutput)
}

func (o LookupNsxtNetworkImportedResultOutput) PrefixLength() pulumi.IntOutput {
	return o.ApplyT(func(v LookupNsxtNetworkImportedResult) int { return v.PrefixLength }).(pulumi.IntOutput)
}

func (o LookupNsxtNetworkImportedResultOutput) SecondaryGateway() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtNetworkImportedResult) string { return v.SecondaryGateway }).(pulumi.StringOutput)
}

func (o LookupNsxtNetworkImportedResultOutput) SecondaryPrefixLength() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtNetworkImportedResult) string { return v.SecondaryPrefixLength }).(pulumi.StringOutput)
}

func (o LookupNsxtNetworkImportedResultOutput) SecondaryStaticIpPools() GetNsxtNetworkImportedSecondaryStaticIpPoolArrayOutput {
	return o.ApplyT(func(v LookupNsxtNetworkImportedResult) []GetNsxtNetworkImportedSecondaryStaticIpPool {
		return v.SecondaryStaticIpPools
	}).(GetNsxtNetworkImportedSecondaryStaticIpPoolArrayOutput)
}

func (o LookupNsxtNetworkImportedResultOutput) StaticIpPools() GetNsxtNetworkImportedStaticIpPoolArrayOutput {
	return o.ApplyT(func(v LookupNsxtNetworkImportedResult) []GetNsxtNetworkImportedStaticIpPool { return v.StaticIpPools }).(GetNsxtNetworkImportedStaticIpPoolArrayOutput)
}

// Deprecated: This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
func (o LookupNsxtNetworkImportedResultOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNsxtNetworkImportedResult) *string { return v.Vdc }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNsxtNetworkImportedResultOutput{})
}
