// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/ergSey/pulumi-vcd/sdk/go/vcd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Supported in provider *v3.7+* and VCD 10.3+ with NSX-T backed VDC Groups.
//
// Provides a data source to read NSX-T Dynamic Security Groups. Dynamic Security Groups group Virtual
// Machines based on specific criteria (VM Names or Security tags) to which Distributed Firewall Rules
// apply.
//
// ## Example Usage
//
// ### 1 (Existing Dynamic Security Group Lookup)
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
//			group1, err := vcd.LookupVdcGroup(ctx, &vcd.LookupVdcGroupArgs{
//				Org:  pulumi.StringRef("cloud"),
//				Name: pulumi.StringRef("vdc-group-cloud"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vcd.LookupNsxtDynamicSecurityGroup(ctx, &vcd.LookupNsxtDynamicSecurityGroupArgs{
//				Org:        pulumi.StringRef("cloud"),
//				VdcGroupId: group1.Id,
//				Name:       "cloud-dynamic-security-group",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupNsxtDynamicSecurityGroup(ctx *pulumi.Context, args *LookupNsxtDynamicSecurityGroupArgs, opts ...pulumi.InvokeOption) (*LookupNsxtDynamicSecurityGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNsxtDynamicSecurityGroupResult
	err := ctx.Invoke("vcd:index/getNsxtDynamicSecurityGroup:getNsxtDynamicSecurityGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNsxtDynamicSecurityGroup.
type LookupNsxtDynamicSecurityGroupArgs struct {
	Description *string `pulumi:"description"`
	// A unique name for existing Dynamic Security Group
	//
	// All the arguments and attributes defined in
	// [`NsxtDynamicSecurityGroup`](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_dynamic_security_group) resource are available.
	Name string `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful
	// when connected as sysadmin working across different organisations.
	Org *string `pulumi:"org"`
	// VDC Group ID hosting existing Dynamic Security Group.
	VdcGroupId string `pulumi:"vdcGroupId"`
}

// A collection of values returned by getNsxtDynamicSecurityGroup.
type LookupNsxtDynamicSecurityGroupResult struct {
	Criterias   []GetNsxtDynamicSecurityGroupCriteria `pulumi:"criterias"`
	Description *string                               `pulumi:"description"`
	// The provider-assigned unique ID for this managed resource.
	Id         string                                `pulumi:"id"`
	MemberVms  []GetNsxtDynamicSecurityGroupMemberVm `pulumi:"memberVms"`
	Name       string                                `pulumi:"name"`
	Org        *string                               `pulumi:"org"`
	VdcGroupId string                                `pulumi:"vdcGroupId"`
}

func LookupNsxtDynamicSecurityGroupOutput(ctx *pulumi.Context, args LookupNsxtDynamicSecurityGroupOutputArgs, opts ...pulumi.InvokeOption) LookupNsxtDynamicSecurityGroupResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupNsxtDynamicSecurityGroupResultOutput, error) {
			args := v.(LookupNsxtDynamicSecurityGroupArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vcd:index/getNsxtDynamicSecurityGroup:getNsxtDynamicSecurityGroup", args, LookupNsxtDynamicSecurityGroupResultOutput{}, options).(LookupNsxtDynamicSecurityGroupResultOutput), nil
		}).(LookupNsxtDynamicSecurityGroupResultOutput)
}

// A collection of arguments for invoking getNsxtDynamicSecurityGroup.
type LookupNsxtDynamicSecurityGroupOutputArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	// A unique name for existing Dynamic Security Group
	//
	// All the arguments and attributes defined in
	// [`NsxtDynamicSecurityGroup`](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_dynamic_security_group) resource are available.
	Name pulumi.StringInput `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful
	// when connected as sysadmin working across different organisations.
	Org pulumi.StringPtrInput `pulumi:"org"`
	// VDC Group ID hosting existing Dynamic Security Group.
	VdcGroupId pulumi.StringInput `pulumi:"vdcGroupId"`
}

func (LookupNsxtDynamicSecurityGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxtDynamicSecurityGroupArgs)(nil)).Elem()
}

// A collection of values returned by getNsxtDynamicSecurityGroup.
type LookupNsxtDynamicSecurityGroupResultOutput struct{ *pulumi.OutputState }

func (LookupNsxtDynamicSecurityGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxtDynamicSecurityGroupResult)(nil)).Elem()
}

func (o LookupNsxtDynamicSecurityGroupResultOutput) ToLookupNsxtDynamicSecurityGroupResultOutput() LookupNsxtDynamicSecurityGroupResultOutput {
	return o
}

func (o LookupNsxtDynamicSecurityGroupResultOutput) ToLookupNsxtDynamicSecurityGroupResultOutputWithContext(ctx context.Context) LookupNsxtDynamicSecurityGroupResultOutput {
	return o
}

func (o LookupNsxtDynamicSecurityGroupResultOutput) Criterias() GetNsxtDynamicSecurityGroupCriteriaArrayOutput {
	return o.ApplyT(func(v LookupNsxtDynamicSecurityGroupResult) []GetNsxtDynamicSecurityGroupCriteria { return v.Criterias }).(GetNsxtDynamicSecurityGroupCriteriaArrayOutput)
}

func (o LookupNsxtDynamicSecurityGroupResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNsxtDynamicSecurityGroupResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNsxtDynamicSecurityGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtDynamicSecurityGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNsxtDynamicSecurityGroupResultOutput) MemberVms() GetNsxtDynamicSecurityGroupMemberVmArrayOutput {
	return o.ApplyT(func(v LookupNsxtDynamicSecurityGroupResult) []GetNsxtDynamicSecurityGroupMemberVm { return v.MemberVms }).(GetNsxtDynamicSecurityGroupMemberVmArrayOutput)
}

func (o LookupNsxtDynamicSecurityGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtDynamicSecurityGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNsxtDynamicSecurityGroupResultOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNsxtDynamicSecurityGroupResult) *string { return v.Org }).(pulumi.StringPtrOutput)
}

func (o LookupNsxtDynamicSecurityGroupResultOutput) VdcGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxtDynamicSecurityGroupResult) string { return v.VdcGroupId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNsxtDynamicSecurityGroupResultOutput{})
}
