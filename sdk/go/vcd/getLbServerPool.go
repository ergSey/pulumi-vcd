// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/ergSey/pulumi-vcd/sdk/go/vcd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VMware Cloud Director Edge Gateway Load Balancer Server Pool data source. A Server Pool defines
// a group of backend servers (defined as pool members), manages load balancer distribution methods, and has a service
// monitor attached to it for health check parameters.
//
// > **Note:** See additional support notes in [server pool resource page](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/lb_server_pool).
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
//			_, err := vcd.LookupLbServerPool(ctx, &vcd.LookupLbServerPoolArgs{
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
func LookupLbServerPool(ctx *pulumi.Context, args *LookupLbServerPoolArgs, opts ...pulumi.InvokeOption) (*LookupLbServerPoolResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupLbServerPoolResult
	err := ctx.Invoke("vcd:index/getLbServerPool:getLbServerPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getLbServerPool.
type LookupLbServerPoolArgs struct {
	// The name of the edge gateway on which the server pool is defined
	EdgeGateway string `pulumi:"edgeGateway"`
	// Server Pool name for identifying the exact server pool
	Name string `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
	Org *string `pulumi:"org"`
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `pulumi:"vdc"`
}

// A collection of values returned by getLbServerPool.
type LookupLbServerPoolResult struct {
	Algorithm           string `pulumi:"algorithm"`
	AlgorithmParameters string `pulumi:"algorithmParameters"`
	Description         string `pulumi:"description"`
	EdgeGateway         string `pulumi:"edgeGateway"`
	EnableTransparency  bool   `pulumi:"enableTransparency"`
	// The provider-assigned unique ID for this managed resource.
	Id        string                  `pulumi:"id"`
	Members   []GetLbServerPoolMember `pulumi:"members"`
	MonitorId string                  `pulumi:"monitorId"`
	Name      string                  `pulumi:"name"`
	Org       *string                 `pulumi:"org"`
	Vdc       *string                 `pulumi:"vdc"`
}

func LookupLbServerPoolOutput(ctx *pulumi.Context, args LookupLbServerPoolOutputArgs, opts ...pulumi.InvokeOption) LookupLbServerPoolResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupLbServerPoolResultOutput, error) {
			args := v.(LookupLbServerPoolArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vcd:index/getLbServerPool:getLbServerPool", args, LookupLbServerPoolResultOutput{}, options).(LookupLbServerPoolResultOutput), nil
		}).(LookupLbServerPoolResultOutput)
}

// A collection of arguments for invoking getLbServerPool.
type LookupLbServerPoolOutputArgs struct {
	// The name of the edge gateway on which the server pool is defined
	EdgeGateway pulumi.StringInput `pulumi:"edgeGateway"`
	// Server Pool name for identifying the exact server pool
	Name pulumi.StringInput `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
	Org pulumi.StringPtrInput `pulumi:"org"`
	// The name of VDC to use, optional if defined at provider level
	Vdc pulumi.StringPtrInput `pulumi:"vdc"`
}

func (LookupLbServerPoolOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLbServerPoolArgs)(nil)).Elem()
}

// A collection of values returned by getLbServerPool.
type LookupLbServerPoolResultOutput struct{ *pulumi.OutputState }

func (LookupLbServerPoolResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLbServerPoolResult)(nil)).Elem()
}

func (o LookupLbServerPoolResultOutput) ToLookupLbServerPoolResultOutput() LookupLbServerPoolResultOutput {
	return o
}

func (o LookupLbServerPoolResultOutput) ToLookupLbServerPoolResultOutputWithContext(ctx context.Context) LookupLbServerPoolResultOutput {
	return o
}

func (o LookupLbServerPoolResultOutput) Algorithm() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbServerPoolResult) string { return v.Algorithm }).(pulumi.StringOutput)
}

func (o LookupLbServerPoolResultOutput) AlgorithmParameters() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbServerPoolResult) string { return v.AlgorithmParameters }).(pulumi.StringOutput)
}

func (o LookupLbServerPoolResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbServerPoolResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupLbServerPoolResultOutput) EdgeGateway() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbServerPoolResult) string { return v.EdgeGateway }).(pulumi.StringOutput)
}

func (o LookupLbServerPoolResultOutput) EnableTransparency() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupLbServerPoolResult) bool { return v.EnableTransparency }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupLbServerPoolResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbServerPoolResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLbServerPoolResultOutput) Members() GetLbServerPoolMemberArrayOutput {
	return o.ApplyT(func(v LookupLbServerPoolResult) []GetLbServerPoolMember { return v.Members }).(GetLbServerPoolMemberArrayOutput)
}

func (o LookupLbServerPoolResultOutput) MonitorId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbServerPoolResult) string { return v.MonitorId }).(pulumi.StringOutput)
}

func (o LookupLbServerPoolResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLbServerPoolResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLbServerPoolResultOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLbServerPoolResult) *string { return v.Org }).(pulumi.StringPtrOutput)
}

func (o LookupLbServerPoolResultOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLbServerPoolResult) *string { return v.Vdc }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLbServerPoolResultOutput{})
}
