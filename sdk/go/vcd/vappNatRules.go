// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"errors"
	"github.com/ergSey/pulumi-vcd/sdk/go/vcd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VappNatRules struct {
	pulumi.CustomResourceState

	// When enabled translates a virtual machine's private, internal IP address to a public IP address for outbound traffic. Default value is `false`.
	EnableIpMasquerade pulumi.BoolPtrOutput `pulumi:"enableIpMasquerade"`
	// Enable or disable NAT. Default is `true`. To enable the NAT service, [VappFirewallRules](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp_firewall_rules) needs to be enabled as well.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// "One of: `ipTranslation` (use IP translation), `portForwarding` (use port forwarding). For `ipTranslation` fields `vmId`, `vmNicId`, `mappingMode` are required and `externalIp` is optional. For `portForwarding` fields `vmId`, `vmNicId`, `protocol`, `externalPort` and `forwardToPort` are required.
	NatType pulumi.StringOutput `pulumi:"natType"`
	// The identifier of [vApp network](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp_network).
	NetworkId pulumi.StringOutput `pulumi:"networkId"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.
	Org pulumi.StringPtrOutput `pulumi:"org"`
	// Configures a NAT rule; see Rules below for details.
	//
	// <a id="rules"></a>
	Rules VappNatRulesRuleArrayOutput `pulumi:"rules"`
	// The identifier of [vApp](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp).
	VappId pulumi.StringOutput `pulumi:"vappId"`
	// The name of VDC to use, optional if defined at provider level.
	Vdc pulumi.StringPtrOutput `pulumi:"vdc"`
}

// NewVappNatRules registers a new resource with the given unique name, arguments, and options.
func NewVappNatRules(ctx *pulumi.Context,
	name string, args *VappNatRulesArgs, opts ...pulumi.ResourceOption) (*VappNatRules, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NatType == nil {
		return nil, errors.New("invalid value for required argument 'NatType'")
	}
	if args.NetworkId == nil {
		return nil, errors.New("invalid value for required argument 'NetworkId'")
	}
	if args.VappId == nil {
		return nil, errors.New("invalid value for required argument 'VappId'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource VappNatRules
	err := ctx.RegisterResource("vcd:index/vappNatRules:VappNatRules", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVappNatRules gets an existing VappNatRules resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVappNatRules(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VappNatRulesState, opts ...pulumi.ResourceOption) (*VappNatRules, error) {
	var resource VappNatRules
	err := ctx.ReadResource("vcd:index/vappNatRules:VappNatRules", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering VappNatRules resources.
type vappNatRulesState struct {
	// When enabled translates a virtual machine's private, internal IP address to a public IP address for outbound traffic. Default value is `false`.
	EnableIpMasquerade *bool `pulumi:"enableIpMasquerade"`
	// Enable or disable NAT. Default is `true`. To enable the NAT service, [VappFirewallRules](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp_firewall_rules) needs to be enabled as well.
	Enabled *bool `pulumi:"enabled"`
	// "One of: `ipTranslation` (use IP translation), `portForwarding` (use port forwarding). For `ipTranslation` fields `vmId`, `vmNicId`, `mappingMode` are required and `externalIp` is optional. For `portForwarding` fields `vmId`, `vmNicId`, `protocol`, `externalPort` and `forwardToPort` are required.
	NatType *string `pulumi:"natType"`
	// The identifier of [vApp network](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp_network).
	NetworkId *string `pulumi:"networkId"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.
	Org *string `pulumi:"org"`
	// Configures a NAT rule; see Rules below for details.
	//
	// <a id="rules"></a>
	Rules []VappNatRulesRule `pulumi:"rules"`
	// The identifier of [vApp](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp).
	VappId *string `pulumi:"vappId"`
	// The name of VDC to use, optional if defined at provider level.
	Vdc *string `pulumi:"vdc"`
}

type VappNatRulesState struct {
	// When enabled translates a virtual machine's private, internal IP address to a public IP address for outbound traffic. Default value is `false`.
	EnableIpMasquerade pulumi.BoolPtrInput
	// Enable or disable NAT. Default is `true`. To enable the NAT service, [VappFirewallRules](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp_firewall_rules) needs to be enabled as well.
	Enabled pulumi.BoolPtrInput
	// "One of: `ipTranslation` (use IP translation), `portForwarding` (use port forwarding). For `ipTranslation` fields `vmId`, `vmNicId`, `mappingMode` are required and `externalIp` is optional. For `portForwarding` fields `vmId`, `vmNicId`, `protocol`, `externalPort` and `forwardToPort` are required.
	NatType pulumi.StringPtrInput
	// The identifier of [vApp network](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp_network).
	NetworkId pulumi.StringPtrInput
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.
	Org pulumi.StringPtrInput
	// Configures a NAT rule; see Rules below for details.
	//
	// <a id="rules"></a>
	Rules VappNatRulesRuleArrayInput
	// The identifier of [vApp](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp).
	VappId pulumi.StringPtrInput
	// The name of VDC to use, optional if defined at provider level.
	Vdc pulumi.StringPtrInput
}

func (VappNatRulesState) ElementType() reflect.Type {
	return reflect.TypeOf((*vappNatRulesState)(nil)).Elem()
}

type vappNatRulesArgs struct {
	// When enabled translates a virtual machine's private, internal IP address to a public IP address for outbound traffic. Default value is `false`.
	EnableIpMasquerade *bool `pulumi:"enableIpMasquerade"`
	// Enable or disable NAT. Default is `true`. To enable the NAT service, [VappFirewallRules](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp_firewall_rules) needs to be enabled as well.
	Enabled *bool `pulumi:"enabled"`
	// "One of: `ipTranslation` (use IP translation), `portForwarding` (use port forwarding). For `ipTranslation` fields `vmId`, `vmNicId`, `mappingMode` are required and `externalIp` is optional. For `portForwarding` fields `vmId`, `vmNicId`, `protocol`, `externalPort` and `forwardToPort` are required.
	NatType string `pulumi:"natType"`
	// The identifier of [vApp network](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp_network).
	NetworkId string `pulumi:"networkId"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.
	Org *string `pulumi:"org"`
	// Configures a NAT rule; see Rules below for details.
	//
	// <a id="rules"></a>
	Rules []VappNatRulesRule `pulumi:"rules"`
	// The identifier of [vApp](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp).
	VappId string `pulumi:"vappId"`
	// The name of VDC to use, optional if defined at provider level.
	Vdc *string `pulumi:"vdc"`
}

// The set of arguments for constructing a VappNatRules resource.
type VappNatRulesArgs struct {
	// When enabled translates a virtual machine's private, internal IP address to a public IP address for outbound traffic. Default value is `false`.
	EnableIpMasquerade pulumi.BoolPtrInput
	// Enable or disable NAT. Default is `true`. To enable the NAT service, [VappFirewallRules](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp_firewall_rules) needs to be enabled as well.
	Enabled pulumi.BoolPtrInput
	// "One of: `ipTranslation` (use IP translation), `portForwarding` (use port forwarding). For `ipTranslation` fields `vmId`, `vmNicId`, `mappingMode` are required and `externalIp` is optional. For `portForwarding` fields `vmId`, `vmNicId`, `protocol`, `externalPort` and `forwardToPort` are required.
	NatType pulumi.StringInput
	// The identifier of [vApp network](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp_network).
	NetworkId pulumi.StringInput
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.
	Org pulumi.StringPtrInput
	// Configures a NAT rule; see Rules below for details.
	//
	// <a id="rules"></a>
	Rules VappNatRulesRuleArrayInput
	// The identifier of [vApp](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp).
	VappId pulumi.StringInput
	// The name of VDC to use, optional if defined at provider level.
	Vdc pulumi.StringPtrInput
}

func (VappNatRulesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*vappNatRulesArgs)(nil)).Elem()
}

type VappNatRulesInput interface {
	pulumi.Input

	ToVappNatRulesOutput() VappNatRulesOutput
	ToVappNatRulesOutputWithContext(ctx context.Context) VappNatRulesOutput
}

func (*VappNatRules) ElementType() reflect.Type {
	return reflect.TypeOf((**VappNatRules)(nil)).Elem()
}

func (i *VappNatRules) ToVappNatRulesOutput() VappNatRulesOutput {
	return i.ToVappNatRulesOutputWithContext(context.Background())
}

func (i *VappNatRules) ToVappNatRulesOutputWithContext(ctx context.Context) VappNatRulesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VappNatRulesOutput)
}

// VappNatRulesArrayInput is an input type that accepts VappNatRulesArray and VappNatRulesArrayOutput values.
// You can construct a concrete instance of `VappNatRulesArrayInput` via:
//
//	VappNatRulesArray{ VappNatRulesArgs{...} }
type VappNatRulesArrayInput interface {
	pulumi.Input

	ToVappNatRulesArrayOutput() VappNatRulesArrayOutput
	ToVappNatRulesArrayOutputWithContext(context.Context) VappNatRulesArrayOutput
}

type VappNatRulesArray []VappNatRulesInput

func (VappNatRulesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VappNatRules)(nil)).Elem()
}

func (i VappNatRulesArray) ToVappNatRulesArrayOutput() VappNatRulesArrayOutput {
	return i.ToVappNatRulesArrayOutputWithContext(context.Background())
}

func (i VappNatRulesArray) ToVappNatRulesArrayOutputWithContext(ctx context.Context) VappNatRulesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VappNatRulesArrayOutput)
}

// VappNatRulesMapInput is an input type that accepts VappNatRulesMap and VappNatRulesMapOutput values.
// You can construct a concrete instance of `VappNatRulesMapInput` via:
//
//	VappNatRulesMap{ "key": VappNatRulesArgs{...} }
type VappNatRulesMapInput interface {
	pulumi.Input

	ToVappNatRulesMapOutput() VappNatRulesMapOutput
	ToVappNatRulesMapOutputWithContext(context.Context) VappNatRulesMapOutput
}

type VappNatRulesMap map[string]VappNatRulesInput

func (VappNatRulesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VappNatRules)(nil)).Elem()
}

func (i VappNatRulesMap) ToVappNatRulesMapOutput() VappNatRulesMapOutput {
	return i.ToVappNatRulesMapOutputWithContext(context.Background())
}

func (i VappNatRulesMap) ToVappNatRulesMapOutputWithContext(ctx context.Context) VappNatRulesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VappNatRulesMapOutput)
}

type VappNatRulesOutput struct{ *pulumi.OutputState }

func (VappNatRulesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VappNatRules)(nil)).Elem()
}

func (o VappNatRulesOutput) ToVappNatRulesOutput() VappNatRulesOutput {
	return o
}

func (o VappNatRulesOutput) ToVappNatRulesOutputWithContext(ctx context.Context) VappNatRulesOutput {
	return o
}

// When enabled translates a virtual machine's private, internal IP address to a public IP address for outbound traffic. Default value is `false`.
func (o VappNatRulesOutput) EnableIpMasquerade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VappNatRules) pulumi.BoolPtrOutput { return v.EnableIpMasquerade }).(pulumi.BoolPtrOutput)
}

// Enable or disable NAT. Default is `true`. To enable the NAT service, [VappFirewallRules](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp_firewall_rules) needs to be enabled as well.
func (o VappNatRulesOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VappNatRules) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// "One of: `ipTranslation` (use IP translation), `portForwarding` (use port forwarding). For `ipTranslation` fields `vmId`, `vmNicId`, `mappingMode` are required and `externalIp` is optional. For `portForwarding` fields `vmId`, `vmNicId`, `protocol`, `externalPort` and `forwardToPort` are required.
func (o VappNatRulesOutput) NatType() pulumi.StringOutput {
	return o.ApplyT(func(v *VappNatRules) pulumi.StringOutput { return v.NatType }).(pulumi.StringOutput)
}

// The identifier of [vApp network](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp_network).
func (o VappNatRulesOutput) NetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v *VappNatRules) pulumi.StringOutput { return v.NetworkId }).(pulumi.StringOutput)
}

// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.
func (o VappNatRulesOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VappNatRules) pulumi.StringPtrOutput { return v.Org }).(pulumi.StringPtrOutput)
}

// Configures a NAT rule; see Rules below for details.
//
// <a id="rules"></a>
func (o VappNatRulesOutput) Rules() VappNatRulesRuleArrayOutput {
	return o.ApplyT(func(v *VappNatRules) VappNatRulesRuleArrayOutput { return v.Rules }).(VappNatRulesRuleArrayOutput)
}

// The identifier of [vApp](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp).
func (o VappNatRulesOutput) VappId() pulumi.StringOutput {
	return o.ApplyT(func(v *VappNatRules) pulumi.StringOutput { return v.VappId }).(pulumi.StringOutput)
}

// The name of VDC to use, optional if defined at provider level.
func (o VappNatRulesOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VappNatRules) pulumi.StringPtrOutput { return v.Vdc }).(pulumi.StringPtrOutput)
}

type VappNatRulesArrayOutput struct{ *pulumi.OutputState }

func (VappNatRulesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*VappNatRules)(nil)).Elem()
}

func (o VappNatRulesArrayOutput) ToVappNatRulesArrayOutput() VappNatRulesArrayOutput {
	return o
}

func (o VappNatRulesArrayOutput) ToVappNatRulesArrayOutputWithContext(ctx context.Context) VappNatRulesArrayOutput {
	return o
}

func (o VappNatRulesArrayOutput) Index(i pulumi.IntInput) VappNatRulesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *VappNatRules {
		return vs[0].([]*VappNatRules)[vs[1].(int)]
	}).(VappNatRulesOutput)
}

type VappNatRulesMapOutput struct{ *pulumi.OutputState }

func (VappNatRulesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*VappNatRules)(nil)).Elem()
}

func (o VappNatRulesMapOutput) ToVappNatRulesMapOutput() VappNatRulesMapOutput {
	return o
}

func (o VappNatRulesMapOutput) ToVappNatRulesMapOutputWithContext(ctx context.Context) VappNatRulesMapOutput {
	return o
}

func (o VappNatRulesMapOutput) MapIndex(k pulumi.StringInput) VappNatRulesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *VappNatRules {
		return vs[0].(map[string]*VappNatRules)[vs[1].(string)]
	}).(VappNatRulesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*VappNatRulesInput)(nil)).Elem(), &VappNatRules{})
	pulumi.RegisterInputType(reflect.TypeOf((*VappNatRulesArrayInput)(nil)).Elem(), VappNatRulesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VappNatRulesMapInput)(nil)).Elem(), VappNatRulesMap{})
	pulumi.RegisterOutputType(VappNatRulesOutput{})
	pulumi.RegisterOutputType(VappNatRulesArrayOutput{})
	pulumi.RegisterOutputType(VappNatRulesMapOutput{})
}
