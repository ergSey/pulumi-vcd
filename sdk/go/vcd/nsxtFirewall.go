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

type NsxtFirewall struct {
	pulumi.CustomResourceState

	// The ID of the Edge Gateway (NSX-T only). Can be looked up using
	// `NsxtEdgegateway` datasource
	EdgeGatewayId pulumi.StringOutput `pulumi:"edgeGatewayId"`
	// The name of organization to use, optional if defined at provider level. Useful
	// when connected as sysadmin working across different organisations.
	Org pulumi.StringPtrOutput `pulumi:"org"`
	// One or more blocks with Firewall Rule definitions
	//
	// <a id="firewall-rule"></a>
	Rules NsxtFirewallRuleArrayOutput `pulumi:"rules"`
	// The name of VDC to use, optional if defined at provider level
	//
	// Deprecated: Edge Gateway will be looked up based on 'edge_gateway_id' field
	Vdc pulumi.StringOutput `pulumi:"vdc"`
}

// NewNsxtFirewall registers a new resource with the given unique name, arguments, and options.
func NewNsxtFirewall(ctx *pulumi.Context,
	name string, args *NsxtFirewallArgs, opts ...pulumi.ResourceOption) (*NsxtFirewall, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EdgeGatewayId == nil {
		return nil, errors.New("invalid value for required argument 'EdgeGatewayId'")
	}
	if args.Rules == nil {
		return nil, errors.New("invalid value for required argument 'Rules'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NsxtFirewall
	err := ctx.RegisterResource("vcd:index/nsxtFirewall:NsxtFirewall", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNsxtFirewall gets an existing NsxtFirewall resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNsxtFirewall(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NsxtFirewallState, opts ...pulumi.ResourceOption) (*NsxtFirewall, error) {
	var resource NsxtFirewall
	err := ctx.ReadResource("vcd:index/nsxtFirewall:NsxtFirewall", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NsxtFirewall resources.
type nsxtFirewallState struct {
	// The ID of the Edge Gateway (NSX-T only). Can be looked up using
	// `NsxtEdgegateway` datasource
	EdgeGatewayId *string `pulumi:"edgeGatewayId"`
	// The name of organization to use, optional if defined at provider level. Useful
	// when connected as sysadmin working across different organisations.
	Org *string `pulumi:"org"`
	// One or more blocks with Firewall Rule definitions
	//
	// <a id="firewall-rule"></a>
	Rules []NsxtFirewallRule `pulumi:"rules"`
	// The name of VDC to use, optional if defined at provider level
	//
	// Deprecated: Edge Gateway will be looked up based on 'edge_gateway_id' field
	Vdc *string `pulumi:"vdc"`
}

type NsxtFirewallState struct {
	// The ID of the Edge Gateway (NSX-T only). Can be looked up using
	// `NsxtEdgegateway` datasource
	EdgeGatewayId pulumi.StringPtrInput
	// The name of organization to use, optional if defined at provider level. Useful
	// when connected as sysadmin working across different organisations.
	Org pulumi.StringPtrInput
	// One or more blocks with Firewall Rule definitions
	//
	// <a id="firewall-rule"></a>
	Rules NsxtFirewallRuleArrayInput
	// The name of VDC to use, optional if defined at provider level
	//
	// Deprecated: Edge Gateway will be looked up based on 'edge_gateway_id' field
	Vdc pulumi.StringPtrInput
}

func (NsxtFirewallState) ElementType() reflect.Type {
	return reflect.TypeOf((*nsxtFirewallState)(nil)).Elem()
}

type nsxtFirewallArgs struct {
	// The ID of the Edge Gateway (NSX-T only). Can be looked up using
	// `NsxtEdgegateway` datasource
	EdgeGatewayId string `pulumi:"edgeGatewayId"`
	// The name of organization to use, optional if defined at provider level. Useful
	// when connected as sysadmin working across different organisations.
	Org *string `pulumi:"org"`
	// One or more blocks with Firewall Rule definitions
	//
	// <a id="firewall-rule"></a>
	Rules []NsxtFirewallRule `pulumi:"rules"`
	// The name of VDC to use, optional if defined at provider level
	//
	// Deprecated: Edge Gateway will be looked up based on 'edge_gateway_id' field
	Vdc *string `pulumi:"vdc"`
}

// The set of arguments for constructing a NsxtFirewall resource.
type NsxtFirewallArgs struct {
	// The ID of the Edge Gateway (NSX-T only). Can be looked up using
	// `NsxtEdgegateway` datasource
	EdgeGatewayId pulumi.StringInput
	// The name of organization to use, optional if defined at provider level. Useful
	// when connected as sysadmin working across different organisations.
	Org pulumi.StringPtrInput
	// One or more blocks with Firewall Rule definitions
	//
	// <a id="firewall-rule"></a>
	Rules NsxtFirewallRuleArrayInput
	// The name of VDC to use, optional if defined at provider level
	//
	// Deprecated: Edge Gateway will be looked up based on 'edge_gateway_id' field
	Vdc pulumi.StringPtrInput
}

func (NsxtFirewallArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nsxtFirewallArgs)(nil)).Elem()
}

type NsxtFirewallInput interface {
	pulumi.Input

	ToNsxtFirewallOutput() NsxtFirewallOutput
	ToNsxtFirewallOutputWithContext(ctx context.Context) NsxtFirewallOutput
}

func (*NsxtFirewall) ElementType() reflect.Type {
	return reflect.TypeOf((**NsxtFirewall)(nil)).Elem()
}

func (i *NsxtFirewall) ToNsxtFirewallOutput() NsxtFirewallOutput {
	return i.ToNsxtFirewallOutputWithContext(context.Background())
}

func (i *NsxtFirewall) ToNsxtFirewallOutputWithContext(ctx context.Context) NsxtFirewallOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxtFirewallOutput)
}

// NsxtFirewallArrayInput is an input type that accepts NsxtFirewallArray and NsxtFirewallArrayOutput values.
// You can construct a concrete instance of `NsxtFirewallArrayInput` via:
//
//	NsxtFirewallArray{ NsxtFirewallArgs{...} }
type NsxtFirewallArrayInput interface {
	pulumi.Input

	ToNsxtFirewallArrayOutput() NsxtFirewallArrayOutput
	ToNsxtFirewallArrayOutputWithContext(context.Context) NsxtFirewallArrayOutput
}

type NsxtFirewallArray []NsxtFirewallInput

func (NsxtFirewallArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NsxtFirewall)(nil)).Elem()
}

func (i NsxtFirewallArray) ToNsxtFirewallArrayOutput() NsxtFirewallArrayOutput {
	return i.ToNsxtFirewallArrayOutputWithContext(context.Background())
}

func (i NsxtFirewallArray) ToNsxtFirewallArrayOutputWithContext(ctx context.Context) NsxtFirewallArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxtFirewallArrayOutput)
}

// NsxtFirewallMapInput is an input type that accepts NsxtFirewallMap and NsxtFirewallMapOutput values.
// You can construct a concrete instance of `NsxtFirewallMapInput` via:
//
//	NsxtFirewallMap{ "key": NsxtFirewallArgs{...} }
type NsxtFirewallMapInput interface {
	pulumi.Input

	ToNsxtFirewallMapOutput() NsxtFirewallMapOutput
	ToNsxtFirewallMapOutputWithContext(context.Context) NsxtFirewallMapOutput
}

type NsxtFirewallMap map[string]NsxtFirewallInput

func (NsxtFirewallMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NsxtFirewall)(nil)).Elem()
}

func (i NsxtFirewallMap) ToNsxtFirewallMapOutput() NsxtFirewallMapOutput {
	return i.ToNsxtFirewallMapOutputWithContext(context.Background())
}

func (i NsxtFirewallMap) ToNsxtFirewallMapOutputWithContext(ctx context.Context) NsxtFirewallMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxtFirewallMapOutput)
}

type NsxtFirewallOutput struct{ *pulumi.OutputState }

func (NsxtFirewallOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NsxtFirewall)(nil)).Elem()
}

func (o NsxtFirewallOutput) ToNsxtFirewallOutput() NsxtFirewallOutput {
	return o
}

func (o NsxtFirewallOutput) ToNsxtFirewallOutputWithContext(ctx context.Context) NsxtFirewallOutput {
	return o
}

// The ID of the Edge Gateway (NSX-T only). Can be looked up using
// `NsxtEdgegateway` datasource
func (o NsxtFirewallOutput) EdgeGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxtFirewall) pulumi.StringOutput { return v.EdgeGatewayId }).(pulumi.StringOutput)
}

// The name of organization to use, optional if defined at provider level. Useful
// when connected as sysadmin working across different organisations.
func (o NsxtFirewallOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsxtFirewall) pulumi.StringPtrOutput { return v.Org }).(pulumi.StringPtrOutput)
}

// One or more blocks with Firewall Rule definitions
//
// <a id="firewall-rule"></a>
func (o NsxtFirewallOutput) Rules() NsxtFirewallRuleArrayOutput {
	return o.ApplyT(func(v *NsxtFirewall) NsxtFirewallRuleArrayOutput { return v.Rules }).(NsxtFirewallRuleArrayOutput)
}

// The name of VDC to use, optional if defined at provider level
//
// Deprecated: Edge Gateway will be looked up based on 'edge_gateway_id' field
func (o NsxtFirewallOutput) Vdc() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxtFirewall) pulumi.StringOutput { return v.Vdc }).(pulumi.StringOutput)
}

type NsxtFirewallArrayOutput struct{ *pulumi.OutputState }

func (NsxtFirewallArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NsxtFirewall)(nil)).Elem()
}

func (o NsxtFirewallArrayOutput) ToNsxtFirewallArrayOutput() NsxtFirewallArrayOutput {
	return o
}

func (o NsxtFirewallArrayOutput) ToNsxtFirewallArrayOutputWithContext(ctx context.Context) NsxtFirewallArrayOutput {
	return o
}

func (o NsxtFirewallArrayOutput) Index(i pulumi.IntInput) NsxtFirewallOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NsxtFirewall {
		return vs[0].([]*NsxtFirewall)[vs[1].(int)]
	}).(NsxtFirewallOutput)
}

type NsxtFirewallMapOutput struct{ *pulumi.OutputState }

func (NsxtFirewallMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NsxtFirewall)(nil)).Elem()
}

func (o NsxtFirewallMapOutput) ToNsxtFirewallMapOutput() NsxtFirewallMapOutput {
	return o
}

func (o NsxtFirewallMapOutput) ToNsxtFirewallMapOutputWithContext(ctx context.Context) NsxtFirewallMapOutput {
	return o
}

func (o NsxtFirewallMapOutput) MapIndex(k pulumi.StringInput) NsxtFirewallOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NsxtFirewall {
		return vs[0].(map[string]*NsxtFirewall)[vs[1].(string)]
	}).(NsxtFirewallOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NsxtFirewallInput)(nil)).Elem(), &NsxtFirewall{})
	pulumi.RegisterInputType(reflect.TypeOf((*NsxtFirewallArrayInput)(nil)).Elem(), NsxtFirewallArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NsxtFirewallMapInput)(nil)).Elem(), NsxtFirewallMap{})
	pulumi.RegisterOutputType(NsxtFirewallOutput{})
	pulumi.RegisterOutputType(NsxtFirewallArrayOutput{})
	pulumi.RegisterOutputType(NsxtFirewallMapOutput{})
}
