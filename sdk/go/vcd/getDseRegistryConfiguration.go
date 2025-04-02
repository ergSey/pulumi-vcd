// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/ergSey/pulumi-vcd/sdk/go/vcd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Supported in provider *v3.13+* with Data Solution Extension.
//
// Provides a data source to read Data Solution Extension (DSE) registry configuration.
//
// > Only `System Administrator` can use this data source.
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
//			_, err := vcd.LookupDseRegistryConfiguration(ctx, &vcd.LookupDseRegistryConfigurationArgs{
//				Name: "MongoDB",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupDseRegistryConfiguration(ctx *pulumi.Context, args *LookupDseRegistryConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupDseRegistryConfigurationResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDseRegistryConfigurationResult
	err := ctx.Invoke("vcd:index/getDseRegistryConfiguration:getDseRegistryConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDseRegistryConfiguration.
type LookupDseRegistryConfigurationArgs struct {
	// The name of Data Solution as it appears in repository configuration
	Name string `pulumi:"name"`
}

// A collection of values returned by getDseRegistryConfiguration.
type LookupDseRegistryConfigurationResult struct {
	ChartRepository              string                                         `pulumi:"chartRepository"`
	CompatibleVersionConstraints []string                                       `pulumi:"compatibleVersionConstraints"`
	ContainerRegistries          []GetDseRegistryConfigurationContainerRegistry `pulumi:"containerRegistries"`
	DefaultChartRepository       string                                         `pulumi:"defaultChartRepository"`
	DefaultPackageName           string                                         `pulumi:"defaultPackageName"`
	DefaultRepository            string                                         `pulumi:"defaultRepository"`
	DefaultVersion               string                                         `pulumi:"defaultVersion"`
	// The provider-assigned unique ID for this managed resource.
	Id                           string `pulumi:"id"`
	Name                         string `pulumi:"name"`
	PackageName                  string `pulumi:"packageName"`
	PackageRepository            string `pulumi:"packageRepository"`
	RdeState                     string `pulumi:"rdeState"`
	RequiresVersionCompatibility bool   `pulumi:"requiresVersionCompatibility"`
	Type                         string `pulumi:"type"`
	Version                      string `pulumi:"version"`
}

func LookupDseRegistryConfigurationOutput(ctx *pulumi.Context, args LookupDseRegistryConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupDseRegistryConfigurationResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupDseRegistryConfigurationResultOutput, error) {
			args := v.(LookupDseRegistryConfigurationArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vcd:index/getDseRegistryConfiguration:getDseRegistryConfiguration", args, LookupDseRegistryConfigurationResultOutput{}, options).(LookupDseRegistryConfigurationResultOutput), nil
		}).(LookupDseRegistryConfigurationResultOutput)
}

// A collection of arguments for invoking getDseRegistryConfiguration.
type LookupDseRegistryConfigurationOutputArgs struct {
	// The name of Data Solution as it appears in repository configuration
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupDseRegistryConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDseRegistryConfigurationArgs)(nil)).Elem()
}

// A collection of values returned by getDseRegistryConfiguration.
type LookupDseRegistryConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupDseRegistryConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDseRegistryConfigurationResult)(nil)).Elem()
}

func (o LookupDseRegistryConfigurationResultOutput) ToLookupDseRegistryConfigurationResultOutput() LookupDseRegistryConfigurationResultOutput {
	return o
}

func (o LookupDseRegistryConfigurationResultOutput) ToLookupDseRegistryConfigurationResultOutputWithContext(ctx context.Context) LookupDseRegistryConfigurationResultOutput {
	return o
}

func (o LookupDseRegistryConfigurationResultOutput) ChartRepository() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDseRegistryConfigurationResult) string { return v.ChartRepository }).(pulumi.StringOutput)
}

func (o LookupDseRegistryConfigurationResultOutput) CompatibleVersionConstraints() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDseRegistryConfigurationResult) []string { return v.CompatibleVersionConstraints }).(pulumi.StringArrayOutput)
}

func (o LookupDseRegistryConfigurationResultOutput) ContainerRegistries() GetDseRegistryConfigurationContainerRegistryArrayOutput {
	return o.ApplyT(func(v LookupDseRegistryConfigurationResult) []GetDseRegistryConfigurationContainerRegistry {
		return v.ContainerRegistries
	}).(GetDseRegistryConfigurationContainerRegistryArrayOutput)
}

func (o LookupDseRegistryConfigurationResultOutput) DefaultChartRepository() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDseRegistryConfigurationResult) string { return v.DefaultChartRepository }).(pulumi.StringOutput)
}

func (o LookupDseRegistryConfigurationResultOutput) DefaultPackageName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDseRegistryConfigurationResult) string { return v.DefaultPackageName }).(pulumi.StringOutput)
}

func (o LookupDseRegistryConfigurationResultOutput) DefaultRepository() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDseRegistryConfigurationResult) string { return v.DefaultRepository }).(pulumi.StringOutput)
}

func (o LookupDseRegistryConfigurationResultOutput) DefaultVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDseRegistryConfigurationResult) string { return v.DefaultVersion }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupDseRegistryConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDseRegistryConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDseRegistryConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDseRegistryConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDseRegistryConfigurationResultOutput) PackageName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDseRegistryConfigurationResult) string { return v.PackageName }).(pulumi.StringOutput)
}

func (o LookupDseRegistryConfigurationResultOutput) PackageRepository() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDseRegistryConfigurationResult) string { return v.PackageRepository }).(pulumi.StringOutput)
}

func (o LookupDseRegistryConfigurationResultOutput) RdeState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDseRegistryConfigurationResult) string { return v.RdeState }).(pulumi.StringOutput)
}

func (o LookupDseRegistryConfigurationResultOutput) RequiresVersionCompatibility() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDseRegistryConfigurationResult) bool { return v.RequiresVersionCompatibility }).(pulumi.BoolOutput)
}

func (o LookupDseRegistryConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDseRegistryConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupDseRegistryConfigurationResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDseRegistryConfigurationResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDseRegistryConfigurationResultOutput{})
}
