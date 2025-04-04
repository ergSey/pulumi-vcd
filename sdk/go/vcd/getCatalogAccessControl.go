// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/ergSey/pulumi-vcd/sdk/go/vcd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a data source to read Access Control details from a Catalog in VMware Cloud Director.
//
// > **Note:** Access control reads run in tenant context, meaning that, even if the user is a system administrator,
// in every request it uses headers items that define the tenant context as restricted to the organization to which the Catalog belongs.
//
// Supported in provider *v3.14+*
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
//			catalog, err := vcd.LookupCatalog(ctx, &vcd.LookupCatalogArgs{
//				Name: pulumi.StringRef("my-catalog"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ac, err := vcd.LookupCatalogAccessControl(ctx, &vcd.LookupCatalogAccessControlArgs{
//				CatalogId: catalog.Id,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("sharedWith", ac.SharedWiths)
//			return nil
//		})
//	}
//
// ```
func LookupCatalogAccessControl(ctx *pulumi.Context, args *LookupCatalogAccessControlArgs, opts ...pulumi.InvokeOption) (*LookupCatalogAccessControlResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupCatalogAccessControlResult
	err := ctx.Invoke("vcd:index/getCatalogAccessControl:getCatalogAccessControl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCatalogAccessControl.
type LookupCatalogAccessControlArgs struct {
	// A unique identifier for the Catalog.
	CatalogId string `pulumi:"catalogId"`
	// The name of organization to which the Catalog belongs. Optional if defined at provider level.
	Org *string `pulumi:"org"`
}

// A collection of values returned by getCatalogAccessControl.
type LookupCatalogAccessControlResult struct {
	CatalogId           string `pulumi:"catalogId"`
	EveryoneAccessLevel string `pulumi:"everyoneAccessLevel"`
	// The provider-assigned unique ID for this managed resource.
	Id                        string                              `pulumi:"id"`
	Org                       *string                             `pulumi:"org"`
	ReadOnlySharedWithAllOrgs bool                                `pulumi:"readOnlySharedWithAllOrgs"`
	SharedWithEveryone        bool                                `pulumi:"sharedWithEveryone"`
	SharedWiths               []GetCatalogAccessControlSharedWith `pulumi:"sharedWiths"`
}

func LookupCatalogAccessControlOutput(ctx *pulumi.Context, args LookupCatalogAccessControlOutputArgs, opts ...pulumi.InvokeOption) LookupCatalogAccessControlResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupCatalogAccessControlResultOutput, error) {
			args := v.(LookupCatalogAccessControlArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vcd:index/getCatalogAccessControl:getCatalogAccessControl", args, LookupCatalogAccessControlResultOutput{}, options).(LookupCatalogAccessControlResultOutput), nil
		}).(LookupCatalogAccessControlResultOutput)
}

// A collection of arguments for invoking getCatalogAccessControl.
type LookupCatalogAccessControlOutputArgs struct {
	// A unique identifier for the Catalog.
	CatalogId pulumi.StringInput `pulumi:"catalogId"`
	// The name of organization to which the Catalog belongs. Optional if defined at provider level.
	Org pulumi.StringPtrInput `pulumi:"org"`
}

func (LookupCatalogAccessControlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCatalogAccessControlArgs)(nil)).Elem()
}

// A collection of values returned by getCatalogAccessControl.
type LookupCatalogAccessControlResultOutput struct{ *pulumi.OutputState }

func (LookupCatalogAccessControlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCatalogAccessControlResult)(nil)).Elem()
}

func (o LookupCatalogAccessControlResultOutput) ToLookupCatalogAccessControlResultOutput() LookupCatalogAccessControlResultOutput {
	return o
}

func (o LookupCatalogAccessControlResultOutput) ToLookupCatalogAccessControlResultOutputWithContext(ctx context.Context) LookupCatalogAccessControlResultOutput {
	return o
}

func (o LookupCatalogAccessControlResultOutput) CatalogId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCatalogAccessControlResult) string { return v.CatalogId }).(pulumi.StringOutput)
}

func (o LookupCatalogAccessControlResultOutput) EveryoneAccessLevel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCatalogAccessControlResult) string { return v.EveryoneAccessLevel }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupCatalogAccessControlResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCatalogAccessControlResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCatalogAccessControlResultOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCatalogAccessControlResult) *string { return v.Org }).(pulumi.StringPtrOutput)
}

func (o LookupCatalogAccessControlResultOutput) ReadOnlySharedWithAllOrgs() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupCatalogAccessControlResult) bool { return v.ReadOnlySharedWithAllOrgs }).(pulumi.BoolOutput)
}

func (o LookupCatalogAccessControlResultOutput) SharedWithEveryone() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupCatalogAccessControlResult) bool { return v.SharedWithEveryone }).(pulumi.BoolOutput)
}

func (o LookupCatalogAccessControlResultOutput) SharedWiths() GetCatalogAccessControlSharedWithArrayOutput {
	return o.ApplyT(func(v LookupCatalogAccessControlResult) []GetCatalogAccessControlSharedWith { return v.SharedWiths }).(GetCatalogAccessControlSharedWithArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCatalogAccessControlResultOutput{})
}
