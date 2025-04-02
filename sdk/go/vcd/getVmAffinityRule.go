// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/ergSey/pulumi-vcd/sdk/go/vcd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VMware Cloud Director VM affinity rule data source. This can be
// used to read VM affinity and anti-affinity rules.
//
// Supported in provider *v2.9+*
//
// > **Note:** The vCD UI defines two different entities (*Affinity Rules* and *Anti-Affinity Rules*). This data source combines both
// entities: they are differentiated by the `polarity` property (See below).
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
//			_, err := vcd.LookupVmAffinityRule(ctx, &vcd.LookupVmAffinityRuleArgs{
//				Name: pulumi.StringRef("my-rule"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vcd.LookupVmAffinityRule(ctx, &vcd.LookupVmAffinityRuleArgs{
//				RuleId: pulumi.StringRef("eda9011c-6841-4060-9336-d2f609c110c3"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupVmAffinityRule(ctx *pulumi.Context, args *LookupVmAffinityRuleArgs, opts ...pulumi.InvokeOption) (*LookupVmAffinityRuleResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupVmAffinityRuleResult
	err := ctx.Invoke("vcd:index/getVmAffinityRule:getVmAffinityRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getVmAffinityRule.
type LookupVmAffinityRuleArgs struct {
	// The name of VM affinity rule. Needed if we don't provide `ruleId`
	Name *string `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
	Org *string `pulumi:"org"`
	// Is the ID of the affinity rule. It's the preferred way to retrieve the affinity
	// rule, especially if the rule name could have duplicates
	RuleId *string `pulumi:"ruleId"`
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `pulumi:"vdc"`
}

// A collection of values returned by getVmAffinityRule.
type LookupVmAffinityRuleResult struct {
	// True if this affinity rule is enabled.
	Enabled bool `pulumi:"enabled"`
	// The provider-assigned unique ID for this managed resource.
	Id   string  `pulumi:"id"`
	Name *string `pulumi:"name"`
	Org  *string `pulumi:"org"`
	// One of `Affinity` or `Anti-Affinity`. This property cannot be changed. Once created, if we
	// need to change polarity, we need to remove the rule and create a new one.
	Polarity string `pulumi:"polarity"`
	// True if this affinity rule is required. When a rule is mandatory, a host failover will not
	// power on the VM if doing so would violate the rule.
	Required bool    `pulumi:"required"`
	RuleId   *string `pulumi:"ruleId"`
	Vdc      *string `pulumi:"vdc"`
	// A set of virtual machine IDs that compose this rule.
	VmIds []string `pulumi:"vmIds"`
}

func LookupVmAffinityRuleOutput(ctx *pulumi.Context, args LookupVmAffinityRuleOutputArgs, opts ...pulumi.InvokeOption) LookupVmAffinityRuleResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupVmAffinityRuleResultOutput, error) {
			args := v.(LookupVmAffinityRuleArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vcd:index/getVmAffinityRule:getVmAffinityRule", args, LookupVmAffinityRuleResultOutput{}, options).(LookupVmAffinityRuleResultOutput), nil
		}).(LookupVmAffinityRuleResultOutput)
}

// A collection of arguments for invoking getVmAffinityRule.
type LookupVmAffinityRuleOutputArgs struct {
	// The name of VM affinity rule. Needed if we don't provide `ruleId`
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
	Org pulumi.StringPtrInput `pulumi:"org"`
	// Is the ID of the affinity rule. It's the preferred way to retrieve the affinity
	// rule, especially if the rule name could have duplicates
	RuleId pulumi.StringPtrInput `pulumi:"ruleId"`
	// The name of VDC to use, optional if defined at provider level
	Vdc pulumi.StringPtrInput `pulumi:"vdc"`
}

func (LookupVmAffinityRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVmAffinityRuleArgs)(nil)).Elem()
}

// A collection of values returned by getVmAffinityRule.
type LookupVmAffinityRuleResultOutput struct{ *pulumi.OutputState }

func (LookupVmAffinityRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVmAffinityRuleResult)(nil)).Elem()
}

func (o LookupVmAffinityRuleResultOutput) ToLookupVmAffinityRuleResultOutput() LookupVmAffinityRuleResultOutput {
	return o
}

func (o LookupVmAffinityRuleResultOutput) ToLookupVmAffinityRuleResultOutputWithContext(ctx context.Context) LookupVmAffinityRuleResultOutput {
	return o
}

// True if this affinity rule is enabled.
func (o LookupVmAffinityRuleResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVmAffinityRuleResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupVmAffinityRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVmAffinityRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVmAffinityRuleResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVmAffinityRuleResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupVmAffinityRuleResultOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVmAffinityRuleResult) *string { return v.Org }).(pulumi.StringPtrOutput)
}

// One of `Affinity` or `Anti-Affinity`. This property cannot be changed. Once created, if we
// need to change polarity, we need to remove the rule and create a new one.
func (o LookupVmAffinityRuleResultOutput) Polarity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVmAffinityRuleResult) string { return v.Polarity }).(pulumi.StringOutput)
}

// True if this affinity rule is required. When a rule is mandatory, a host failover will not
// power on the VM if doing so would violate the rule.
func (o LookupVmAffinityRuleResultOutput) Required() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupVmAffinityRuleResult) bool { return v.Required }).(pulumi.BoolOutput)
}

func (o LookupVmAffinityRuleResultOutput) RuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVmAffinityRuleResult) *string { return v.RuleId }).(pulumi.StringPtrOutput)
}

func (o LookupVmAffinityRuleResultOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVmAffinityRuleResult) *string { return v.Vdc }).(pulumi.StringPtrOutput)
}

// A set of virtual machine IDs that compose this rule.
func (o LookupVmAffinityRuleResultOutput) VmIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVmAffinityRuleResult) []string { return v.VmIds }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVmAffinityRuleResultOutput{})
}
