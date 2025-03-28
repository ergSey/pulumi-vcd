// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/ergSey/pulumi-vcd/sdk/go/vcd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Supported in provider *v3.3+* and VCD 10.1+ with NSX-T backed VDCs.
//
// Provides a data source to access NSX-T Security Group configuration. Security Groups are groups of
// data center group networks to which distributed firewall rules apply. Grouping networks helps you to
// reduce the total number of distributed firewall rules to be created.
//
// ## Example Usage
//
// ### 1
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
//			main, err := vcd.LookupNsxtEdgegateway(ctx, &vcd.LookupNsxtEdgegatewayArgs{
//				Org:  pulumi.StringRef("my-org"),
//				Name: "main-edge",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vcd.LookupNsxtSecurityGroup(ctx, &vcd.LookupNsxtSecurityGroupArgs{
//				Org:           pulumi.StringRef("my-org"),
//				EdgeGatewayId: main.Id,
//				Name:          "test-security-group-changed",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupNsxtSecurityGroup(ctx *pulumi.Context, args *LookupNsxtSecurityGroupArgs, opts ...pulumi.InvokeOption) (*LookupNsxtSecurityGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNsxtSecurityGroupResult
	err := ctx.Invoke("vcd:index/getNsxtSecurityGroup:getNsxtSecurityGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNsxtSecurityGroup.
type LookupNsxtSecurityGroupArgs struct {
	// The ID of the Edge Gateway (NSX-T only). Can be looked up using
	EdgeGatewayId string `pulumi:"edgeGatewayId"`
	// Unique name of existing Security Group.
	Name string `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful
	// when connected as sysadmin working across different organisations.
	Org *string `pulumi:"org"`
	// The name of VDC to use, optional if defined at provider level. **Deprecated**
	// in favor of `edgeGatewayId` field.
	//
	// Deprecated: Deprecated in favor of `edgeGatewayId`. Security Group will inherit VDC from parent Edge Gateway.
	Vdc *string `pulumi:"vdc"`
}

// A collection of values returned by getNsxtSecurityGroup.
type LookupNsxtSecurityGroupResult struct {
	Description   string `pulumi:"description"`
	EdgeGatewayId string `pulumi:"edgeGatewayId"`
	// The provider-assigned unique ID for this managed resource.
	Id                  string                         `pulumi:"id"`
	MemberOrgNetworkIds []string                       `pulumi:"memberOrgNetworkIds"`
	MemberVms           []GetNsxtSecurityGroupMemberVm `pulumi:"memberVms"`
	Name                string                         `pulumi:"name"`
	Org                 *string                        `pulumi:"org"`
	// Parent VDC or VDC Group ID.
	OwnerId string `pulumi:"ownerId"`
	// Deprecated: Deprecated in favor of `edgeGatewayId`. Security Group will inherit VDC from parent Edge Gateway.
	Vdc *string `pulumi:"vdc"`
}

func LookupNsxtSecurityGroupOutput(ctx *pulumi.Context, args LookupNsxtSecurityGroupOutputArgs, opts ...pulumi.InvokeOption) LookupNsxtSecurityGroupResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupNsxtSecurityGroupResultOutput, error) {
			args := v.(LookupNsxtSecurityGroupArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vcd:index/getNsxtSecurityGroup:getNsxtSecurityGroup", args, LookupNsxtSecurityGroupResultOutput{}, options).(LookupNsxtSecurityGroupResultOutput), nil
		}).(LookupNsxtSecurityGroupResultOutput)
}

// A collection of arguments for invoking getNsxtSecurityGroup.
type LookupNsxtSecurityGroupOutputArgs struct {
	// The ID of the Edge Gateway (NSX-T only). Can be looked up using
	EdgeGatewayId pulumi.StringInput `pulumi:"edgeGatewayId"`
	// Unique name of existing Security Group.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful
	// when connected as sysadmin working across different organisations.
	Org pulumi.StringPtrInput `pulumi:"org"`
	// The name of VDC to use, optional if defined at provider level. **Deprecated**
	// in favor of `edgeGatewayId` field.
	//
	// Deprecated: Deprecated in favor of `edgeGatewayId`. Security Group will inherit VDC from parent Edge Gateway.
	Vdc pulumi.StringPtrInput `pulumi:"vdc"`
}

func (LookupNsxtSecurityGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxtSecurityGroupArgs)(nil)).Elem()
}

// A collection of values returned by getNsxtSecurityGroup.
type LookupNsxtSecurityGroupResultOutput struct{ *pulumi.OutputState }

func (LookupNsxtSecurityGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxtSecurityGroupResult)(nil)).Elem()
}

func (o LookupNsxtSecurityGroupResultOutput) ToLookupNsxtSecurityGroupResultOutput() LookupNsxtSecurityGroupResultOutput {
	return o
}

func (o LookupNsxtSecurityGroupResultOutput) ToLookupNsxtSecurityGroupResultOutputWithContext(ctx context.Context) LookupNsxtSecurityGroupResultOutput {
	return o
}

func (o LookupNsxtSecurityGroupResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtSecurityGroupResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupNsxtSecurityGroupResultOutput) EdgeGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtSecurityGroupResult) string { return v.EdgeGatewayId }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNsxtSecurityGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtSecurityGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNsxtSecurityGroupResultOutput) MemberOrgNetworkIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNsxtSecurityGroupResult) []string { return v.MemberOrgNetworkIds }).(pulumi.StringArrayOutput)
}

func (o LookupNsxtSecurityGroupResultOutput) MemberVms() GetNsxtSecurityGroupMemberVmArrayOutput {
	return o.ApplyT(func(v LookupNsxtSecurityGroupResult) []GetNsxtSecurityGroupMemberVm { return v.MemberVms }).(GetNsxtSecurityGroupMemberVmArrayOutput)
}

func (o LookupNsxtSecurityGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtSecurityGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNsxtSecurityGroupResultOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNsxtSecurityGroupResult) *string { return v.Org }).(pulumi.StringPtrOutput)
}

// Parent VDC or VDC Group ID.
func (o LookupNsxtSecurityGroupResultOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtSecurityGroupResult) string { return v.OwnerId }).(pulumi.StringOutput)
}

// Deprecated: Deprecated in favor of `edgeGatewayId`. Security Group will inherit VDC from parent Edge Gateway.
func (o LookupNsxtSecurityGroupResultOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNsxtSecurityGroupResult) *string { return v.Vdc }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNsxtSecurityGroupResultOutput{})
}
