// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/ergSey/pulumi-vcd/sdk/go/vcd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EdgegatewaySettings struct {
	pulumi.CustomResourceState

	// The edge gateway ID. (Required if `edgeGatewayName` is not set)
	EdgeGatewayId pulumi.StringOutput `pulumi:"edgeGatewayId"`
	// A unique name for the edge gateway. (Required if `edgeGatewayId` is not set)
	EdgeGatewayName pulumi.StringOutput `pulumi:"edgeGatewayName"`
	// Default firewall rule (last in the processing order) action.
	// One of `accept` or `deny`. Default `deny`.
	FwDefaultRuleAction pulumi.StringPtrOutput `pulumi:"fwDefaultRuleAction"`
	// Enable default firewall rule (last in the processing
	// order) logging. Default `false`.
	FwDefaultRuleLoggingEnabled pulumi.BoolPtrOutput `pulumi:"fwDefaultRuleLoggingEnabled"`
	// Enable firewall. Default `true`.
	FwEnabled pulumi.BoolPtrOutput `pulumi:"fwEnabled"`
	// Enable to configure the load balancer.
	LbAccelerationEnabled pulumi.BoolPtrOutput `pulumi:"lbAccelerationEnabled"`
	// Enable load balancing. Default is `false`.
	LbEnabled pulumi.BoolPtrOutput `pulumi:"lbEnabled"`
	// Enables the edge gateway load balancer to collect traffic logs.
	// Default is `false`. Note: **only System administrators can change this property**. It is ignored by API for Org users.
	LbLoggingEnabled pulumi.BoolPtrOutput `pulumi:"lbLoggingEnabled"`
	// Choose the severity of events to be logged. One of `emergency`,
	// `alert`, `critical`, `error`, `warning`, `notice`, `info`, `debug`. Note: **only System administrators can change this property**. It is ignored by API for Org users.
	LbLoglevel pulumi.StringPtrOutput `pulumi:"lbLoglevel"`
	// The name of organization to which the VDC belongs. Optional if defined at provider level.
	Org pulumi.StringPtrOutput `pulumi:"org"`
	// The name of VDC that owns the edge gateway. Optional if defined at provider level.
	Vdc pulumi.StringPtrOutput `pulumi:"vdc"`
}

// NewEdgegatewaySettings registers a new resource with the given unique name, arguments, and options.
func NewEdgegatewaySettings(ctx *pulumi.Context,
	name string, args *EdgegatewaySettingsArgs, opts ...pulumi.ResourceOption) (*EdgegatewaySettings, error) {
	if args == nil {
		args = &EdgegatewaySettingsArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource EdgegatewaySettings
	err := ctx.RegisterResource("vcd:index/edgegatewaySettings:EdgegatewaySettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEdgegatewaySettings gets an existing EdgegatewaySettings resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEdgegatewaySettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EdgegatewaySettingsState, opts ...pulumi.ResourceOption) (*EdgegatewaySettings, error) {
	var resource EdgegatewaySettings
	err := ctx.ReadResource("vcd:index/edgegatewaySettings:EdgegatewaySettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering EdgegatewaySettings resources.
type edgegatewaySettingsState struct {
	// The edge gateway ID. (Required if `edgeGatewayName` is not set)
	EdgeGatewayId *string `pulumi:"edgeGatewayId"`
	// A unique name for the edge gateway. (Required if `edgeGatewayId` is not set)
	EdgeGatewayName *string `pulumi:"edgeGatewayName"`
	// Default firewall rule (last in the processing order) action.
	// One of `accept` or `deny`. Default `deny`.
	FwDefaultRuleAction *string `pulumi:"fwDefaultRuleAction"`
	// Enable default firewall rule (last in the processing
	// order) logging. Default `false`.
	FwDefaultRuleLoggingEnabled *bool `pulumi:"fwDefaultRuleLoggingEnabled"`
	// Enable firewall. Default `true`.
	FwEnabled *bool `pulumi:"fwEnabled"`
	// Enable to configure the load balancer.
	LbAccelerationEnabled *bool `pulumi:"lbAccelerationEnabled"`
	// Enable load balancing. Default is `false`.
	LbEnabled *bool `pulumi:"lbEnabled"`
	// Enables the edge gateway load balancer to collect traffic logs.
	// Default is `false`. Note: **only System administrators can change this property**. It is ignored by API for Org users.
	LbLoggingEnabled *bool `pulumi:"lbLoggingEnabled"`
	// Choose the severity of events to be logged. One of `emergency`,
	// `alert`, `critical`, `error`, `warning`, `notice`, `info`, `debug`. Note: **only System administrators can change this property**. It is ignored by API for Org users.
	LbLoglevel *string `pulumi:"lbLoglevel"`
	// The name of organization to which the VDC belongs. Optional if defined at provider level.
	Org *string `pulumi:"org"`
	// The name of VDC that owns the edge gateway. Optional if defined at provider level.
	Vdc *string `pulumi:"vdc"`
}

type EdgegatewaySettingsState struct {
	// The edge gateway ID. (Required if `edgeGatewayName` is not set)
	EdgeGatewayId pulumi.StringPtrInput
	// A unique name for the edge gateway. (Required if `edgeGatewayId` is not set)
	EdgeGatewayName pulumi.StringPtrInput
	// Default firewall rule (last in the processing order) action.
	// One of `accept` or `deny`. Default `deny`.
	FwDefaultRuleAction pulumi.StringPtrInput
	// Enable default firewall rule (last in the processing
	// order) logging. Default `false`.
	FwDefaultRuleLoggingEnabled pulumi.BoolPtrInput
	// Enable firewall. Default `true`.
	FwEnabled pulumi.BoolPtrInput
	// Enable to configure the load balancer.
	LbAccelerationEnabled pulumi.BoolPtrInput
	// Enable load balancing. Default is `false`.
	LbEnabled pulumi.BoolPtrInput
	// Enables the edge gateway load balancer to collect traffic logs.
	// Default is `false`. Note: **only System administrators can change this property**. It is ignored by API for Org users.
	LbLoggingEnabled pulumi.BoolPtrInput
	// Choose the severity of events to be logged. One of `emergency`,
	// `alert`, `critical`, `error`, `warning`, `notice`, `info`, `debug`. Note: **only System administrators can change this property**. It is ignored by API for Org users.
	LbLoglevel pulumi.StringPtrInput
	// The name of organization to which the VDC belongs. Optional if defined at provider level.
	Org pulumi.StringPtrInput
	// The name of VDC that owns the edge gateway. Optional if defined at provider level.
	Vdc pulumi.StringPtrInput
}

func (EdgegatewaySettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*edgegatewaySettingsState)(nil)).Elem()
}

type edgegatewaySettingsArgs struct {
	// The edge gateway ID. (Required if `edgeGatewayName` is not set)
	EdgeGatewayId *string `pulumi:"edgeGatewayId"`
	// A unique name for the edge gateway. (Required if `edgeGatewayId` is not set)
	EdgeGatewayName *string `pulumi:"edgeGatewayName"`
	// Default firewall rule (last in the processing order) action.
	// One of `accept` or `deny`. Default `deny`.
	FwDefaultRuleAction *string `pulumi:"fwDefaultRuleAction"`
	// Enable default firewall rule (last in the processing
	// order) logging. Default `false`.
	FwDefaultRuleLoggingEnabled *bool `pulumi:"fwDefaultRuleLoggingEnabled"`
	// Enable firewall. Default `true`.
	FwEnabled *bool `pulumi:"fwEnabled"`
	// Enable to configure the load balancer.
	LbAccelerationEnabled *bool `pulumi:"lbAccelerationEnabled"`
	// Enable load balancing. Default is `false`.
	LbEnabled *bool `pulumi:"lbEnabled"`
	// Enables the edge gateway load balancer to collect traffic logs.
	// Default is `false`. Note: **only System administrators can change this property**. It is ignored by API for Org users.
	LbLoggingEnabled *bool `pulumi:"lbLoggingEnabled"`
	// Choose the severity of events to be logged. One of `emergency`,
	// `alert`, `critical`, `error`, `warning`, `notice`, `info`, `debug`. Note: **only System administrators can change this property**. It is ignored by API for Org users.
	LbLoglevel *string `pulumi:"lbLoglevel"`
	// The name of organization to which the VDC belongs. Optional if defined at provider level.
	Org *string `pulumi:"org"`
	// The name of VDC that owns the edge gateway. Optional if defined at provider level.
	Vdc *string `pulumi:"vdc"`
}

// The set of arguments for constructing a EdgegatewaySettings resource.
type EdgegatewaySettingsArgs struct {
	// The edge gateway ID. (Required if `edgeGatewayName` is not set)
	EdgeGatewayId pulumi.StringPtrInput
	// A unique name for the edge gateway. (Required if `edgeGatewayId` is not set)
	EdgeGatewayName pulumi.StringPtrInput
	// Default firewall rule (last in the processing order) action.
	// One of `accept` or `deny`. Default `deny`.
	FwDefaultRuleAction pulumi.StringPtrInput
	// Enable default firewall rule (last in the processing
	// order) logging. Default `false`.
	FwDefaultRuleLoggingEnabled pulumi.BoolPtrInput
	// Enable firewall. Default `true`.
	FwEnabled pulumi.BoolPtrInput
	// Enable to configure the load balancer.
	LbAccelerationEnabled pulumi.BoolPtrInput
	// Enable load balancing. Default is `false`.
	LbEnabled pulumi.BoolPtrInput
	// Enables the edge gateway load balancer to collect traffic logs.
	// Default is `false`. Note: **only System administrators can change this property**. It is ignored by API for Org users.
	LbLoggingEnabled pulumi.BoolPtrInput
	// Choose the severity of events to be logged. One of `emergency`,
	// `alert`, `critical`, `error`, `warning`, `notice`, `info`, `debug`. Note: **only System administrators can change this property**. It is ignored by API for Org users.
	LbLoglevel pulumi.StringPtrInput
	// The name of organization to which the VDC belongs. Optional if defined at provider level.
	Org pulumi.StringPtrInput
	// The name of VDC that owns the edge gateway. Optional if defined at provider level.
	Vdc pulumi.StringPtrInput
}

func (EdgegatewaySettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*edgegatewaySettingsArgs)(nil)).Elem()
}

type EdgegatewaySettingsInput interface {
	pulumi.Input

	ToEdgegatewaySettingsOutput() EdgegatewaySettingsOutput
	ToEdgegatewaySettingsOutputWithContext(ctx context.Context) EdgegatewaySettingsOutput
}

func (*EdgegatewaySettings) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgegatewaySettings)(nil)).Elem()
}

func (i *EdgegatewaySettings) ToEdgegatewaySettingsOutput() EdgegatewaySettingsOutput {
	return i.ToEdgegatewaySettingsOutputWithContext(context.Background())
}

func (i *EdgegatewaySettings) ToEdgegatewaySettingsOutputWithContext(ctx context.Context) EdgegatewaySettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgegatewaySettingsOutput)
}

// EdgegatewaySettingsArrayInput is an input type that accepts EdgegatewaySettingsArray and EdgegatewaySettingsArrayOutput values.
// You can construct a concrete instance of `EdgegatewaySettingsArrayInput` via:
//
//	EdgegatewaySettingsArray{ EdgegatewaySettingsArgs{...} }
type EdgegatewaySettingsArrayInput interface {
	pulumi.Input

	ToEdgegatewaySettingsArrayOutput() EdgegatewaySettingsArrayOutput
	ToEdgegatewaySettingsArrayOutputWithContext(context.Context) EdgegatewaySettingsArrayOutput
}

type EdgegatewaySettingsArray []EdgegatewaySettingsInput

func (EdgegatewaySettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EdgegatewaySettings)(nil)).Elem()
}

func (i EdgegatewaySettingsArray) ToEdgegatewaySettingsArrayOutput() EdgegatewaySettingsArrayOutput {
	return i.ToEdgegatewaySettingsArrayOutputWithContext(context.Background())
}

func (i EdgegatewaySettingsArray) ToEdgegatewaySettingsArrayOutputWithContext(ctx context.Context) EdgegatewaySettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgegatewaySettingsArrayOutput)
}

// EdgegatewaySettingsMapInput is an input type that accepts EdgegatewaySettingsMap and EdgegatewaySettingsMapOutput values.
// You can construct a concrete instance of `EdgegatewaySettingsMapInput` via:
//
//	EdgegatewaySettingsMap{ "key": EdgegatewaySettingsArgs{...} }
type EdgegatewaySettingsMapInput interface {
	pulumi.Input

	ToEdgegatewaySettingsMapOutput() EdgegatewaySettingsMapOutput
	ToEdgegatewaySettingsMapOutputWithContext(context.Context) EdgegatewaySettingsMapOutput
}

type EdgegatewaySettingsMap map[string]EdgegatewaySettingsInput

func (EdgegatewaySettingsMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EdgegatewaySettings)(nil)).Elem()
}

func (i EdgegatewaySettingsMap) ToEdgegatewaySettingsMapOutput() EdgegatewaySettingsMapOutput {
	return i.ToEdgegatewaySettingsMapOutputWithContext(context.Background())
}

func (i EdgegatewaySettingsMap) ToEdgegatewaySettingsMapOutputWithContext(ctx context.Context) EdgegatewaySettingsMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EdgegatewaySettingsMapOutput)
}

type EdgegatewaySettingsOutput struct{ *pulumi.OutputState }

func (EdgegatewaySettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EdgegatewaySettings)(nil)).Elem()
}

func (o EdgegatewaySettingsOutput) ToEdgegatewaySettingsOutput() EdgegatewaySettingsOutput {
	return o
}

func (o EdgegatewaySettingsOutput) ToEdgegatewaySettingsOutputWithContext(ctx context.Context) EdgegatewaySettingsOutput {
	return o
}

// The edge gateway ID. (Required if `edgeGatewayName` is not set)
func (o EdgegatewaySettingsOutput) EdgeGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgegatewaySettings) pulumi.StringOutput { return v.EdgeGatewayId }).(pulumi.StringOutput)
}

// A unique name for the edge gateway. (Required if `edgeGatewayId` is not set)
func (o EdgegatewaySettingsOutput) EdgeGatewayName() pulumi.StringOutput {
	return o.ApplyT(func(v *EdgegatewaySettings) pulumi.StringOutput { return v.EdgeGatewayName }).(pulumi.StringOutput)
}

// Default firewall rule (last in the processing order) action.
// One of `accept` or `deny`. Default `deny`.
func (o EdgegatewaySettingsOutput) FwDefaultRuleAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgegatewaySettings) pulumi.StringPtrOutput { return v.FwDefaultRuleAction }).(pulumi.StringPtrOutput)
}

// Enable default firewall rule (last in the processing
// order) logging. Default `false`.
func (o EdgegatewaySettingsOutput) FwDefaultRuleLoggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdgegatewaySettings) pulumi.BoolPtrOutput { return v.FwDefaultRuleLoggingEnabled }).(pulumi.BoolPtrOutput)
}

// Enable firewall. Default `true`.
func (o EdgegatewaySettingsOutput) FwEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdgegatewaySettings) pulumi.BoolPtrOutput { return v.FwEnabled }).(pulumi.BoolPtrOutput)
}

// Enable to configure the load balancer.
func (o EdgegatewaySettingsOutput) LbAccelerationEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdgegatewaySettings) pulumi.BoolPtrOutput { return v.LbAccelerationEnabled }).(pulumi.BoolPtrOutput)
}

// Enable load balancing. Default is `false`.
func (o EdgegatewaySettingsOutput) LbEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdgegatewaySettings) pulumi.BoolPtrOutput { return v.LbEnabled }).(pulumi.BoolPtrOutput)
}

// Enables the edge gateway load balancer to collect traffic logs.
// Default is `false`. Note: **only System administrators can change this property**. It is ignored by API for Org users.
func (o EdgegatewaySettingsOutput) LbLoggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EdgegatewaySettings) pulumi.BoolPtrOutput { return v.LbLoggingEnabled }).(pulumi.BoolPtrOutput)
}

// Choose the severity of events to be logged. One of `emergency`,
// `alert`, `critical`, `error`, `warning`, `notice`, `info`, `debug`. Note: **only System administrators can change this property**. It is ignored by API for Org users.
func (o EdgegatewaySettingsOutput) LbLoglevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgegatewaySettings) pulumi.StringPtrOutput { return v.LbLoglevel }).(pulumi.StringPtrOutput)
}

// The name of organization to which the VDC belongs. Optional if defined at provider level.
func (o EdgegatewaySettingsOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgegatewaySettings) pulumi.StringPtrOutput { return v.Org }).(pulumi.StringPtrOutput)
}

// The name of VDC that owns the edge gateway. Optional if defined at provider level.
func (o EdgegatewaySettingsOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EdgegatewaySettings) pulumi.StringPtrOutput { return v.Vdc }).(pulumi.StringPtrOutput)
}

type EdgegatewaySettingsArrayOutput struct{ *pulumi.OutputState }

func (EdgegatewaySettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*EdgegatewaySettings)(nil)).Elem()
}

func (o EdgegatewaySettingsArrayOutput) ToEdgegatewaySettingsArrayOutput() EdgegatewaySettingsArrayOutput {
	return o
}

func (o EdgegatewaySettingsArrayOutput) ToEdgegatewaySettingsArrayOutputWithContext(ctx context.Context) EdgegatewaySettingsArrayOutput {
	return o
}

func (o EdgegatewaySettingsArrayOutput) Index(i pulumi.IntInput) EdgegatewaySettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *EdgegatewaySettings {
		return vs[0].([]*EdgegatewaySettings)[vs[1].(int)]
	}).(EdgegatewaySettingsOutput)
}

type EdgegatewaySettingsMapOutput struct{ *pulumi.OutputState }

func (EdgegatewaySettingsMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*EdgegatewaySettings)(nil)).Elem()
}

func (o EdgegatewaySettingsMapOutput) ToEdgegatewaySettingsMapOutput() EdgegatewaySettingsMapOutput {
	return o
}

func (o EdgegatewaySettingsMapOutput) ToEdgegatewaySettingsMapOutputWithContext(ctx context.Context) EdgegatewaySettingsMapOutput {
	return o
}

func (o EdgegatewaySettingsMapOutput) MapIndex(k pulumi.StringInput) EdgegatewaySettingsOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *EdgegatewaySettings {
		return vs[0].(map[string]*EdgegatewaySettings)[vs[1].(string)]
	}).(EdgegatewaySettingsOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EdgegatewaySettingsInput)(nil)).Elem(), &EdgegatewaySettings{})
	pulumi.RegisterInputType(reflect.TypeOf((*EdgegatewaySettingsArrayInput)(nil)).Elem(), EdgegatewaySettingsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EdgegatewaySettingsMapInput)(nil)).Elem(), EdgegatewaySettingsMap{})
	pulumi.RegisterOutputType(EdgegatewaySettingsOutput{})
	pulumi.RegisterOutputType(EdgegatewaySettingsArrayOutput{})
	pulumi.RegisterOutputType(EdgegatewaySettingsMapOutput{})
}
