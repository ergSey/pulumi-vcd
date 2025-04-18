// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/ergSey/pulumi-vcd/sdk/go/vcd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VMware Cloud Director generic structure data source. It shows the structure of any VCD resource.
//
// Supported in provider *v3.1+*
//
// ## Example Usage
//
// ### 1
//
// # Showing a structure with simple attributes only
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
//			orgStruct, err := vcd.GetResourceSchema(ctx, &vcd.GetResourceSchemaArgs{
//				Name:         "org_struct",
//				ResourceType: "vcd_org",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("orgStruct", orgStruct.Attributes)
//			return nil
//		})
//	}
//
// ```
//
// ### 2
//
// # Showing a structure with both simple and compound attributes
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
//			_, err := vcd.GetResourceSchema(ctx, &vcd.GetResourceSchemaArgs{
//				Name:         "net_struct",
//				ResourceType: "vcd_network_isolated",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("netStruct", netStructVcdResourceSchema)
//			return nil
//		})
//	}
//
// ```
func GetResourceSchema(ctx *pulumi.Context, args *GetResourceSchemaArgs, opts ...pulumi.InvokeOption) (*GetResourceSchemaResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetResourceSchemaResult
	err := ctx.Invoke("vcd:index/getResourceSchema:getResourceSchema", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getResourceSchema.
type GetResourceSchemaArgs struct {
	// An unique name to identify the data source
	Name string `pulumi:"name"`
	// Which resource we want to list. It needs to use the full name of the resource (i.e. "Org",
	// not simply "org")
	ResourceType string `pulumi:"resourceType"`
}

// A collection of values returned by getResourceSchema.
type GetResourceSchemaResult struct {
	// (Computed) Same composition of the simple `attributes` above.
	Attributes []GetResourceSchemaAttribute `pulumi:"attributes"`
	// (Computed) The list of compound attributes
	// Each bock attribute is made of:
	BlockAttributes []GetResourceSchemaBlockAttribute `pulumi:"blockAttributes"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// the attribute name
	Name         string `pulumi:"name"`
	ResourceType string `pulumi:"resourceType"`
}

func GetResourceSchemaOutput(ctx *pulumi.Context, args GetResourceSchemaOutputArgs, opts ...pulumi.InvokeOption) GetResourceSchemaResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetResourceSchemaResultOutput, error) {
			args := v.(GetResourceSchemaArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vcd:index/getResourceSchema:getResourceSchema", args, GetResourceSchemaResultOutput{}, options).(GetResourceSchemaResultOutput), nil
		}).(GetResourceSchemaResultOutput)
}

// A collection of arguments for invoking getResourceSchema.
type GetResourceSchemaOutputArgs struct {
	// An unique name to identify the data source
	Name pulumi.StringInput `pulumi:"name"`
	// Which resource we want to list. It needs to use the full name of the resource (i.e. "Org",
	// not simply "org")
	ResourceType pulumi.StringInput `pulumi:"resourceType"`
}

func (GetResourceSchemaOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetResourceSchemaArgs)(nil)).Elem()
}

// A collection of values returned by getResourceSchema.
type GetResourceSchemaResultOutput struct{ *pulumi.OutputState }

func (GetResourceSchemaResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetResourceSchemaResult)(nil)).Elem()
}

func (o GetResourceSchemaResultOutput) ToGetResourceSchemaResultOutput() GetResourceSchemaResultOutput {
	return o
}

func (o GetResourceSchemaResultOutput) ToGetResourceSchemaResultOutputWithContext(ctx context.Context) GetResourceSchemaResultOutput {
	return o
}

// (Computed) Same composition of the simple `attributes` above.
func (o GetResourceSchemaResultOutput) Attributes() GetResourceSchemaAttributeArrayOutput {
	return o.ApplyT(func(v GetResourceSchemaResult) []GetResourceSchemaAttribute { return v.Attributes }).(GetResourceSchemaAttributeArrayOutput)
}

// (Computed) The list of compound attributes
// Each bock attribute is made of:
func (o GetResourceSchemaResultOutput) BlockAttributes() GetResourceSchemaBlockAttributeArrayOutput {
	return o.ApplyT(func(v GetResourceSchemaResult) []GetResourceSchemaBlockAttribute { return v.BlockAttributes }).(GetResourceSchemaBlockAttributeArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetResourceSchemaResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetResourceSchemaResult) string { return v.Id }).(pulumi.StringOutput)
}

// the attribute name
func (o GetResourceSchemaResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetResourceSchemaResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetResourceSchemaResultOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v GetResourceSchemaResult) string { return v.ResourceType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetResourceSchemaResultOutput{})
}
