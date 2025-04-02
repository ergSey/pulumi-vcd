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

type NsxtEdgegatewayStaticRoute struct {
	pulumi.CustomResourceState

	// Description for NSX-T Edge Gateway Static Route
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// NSX-T Edge Gateway ID
	EdgeGatewayId pulumi.StringOutput `pulumi:"edgeGatewayId"`
	// Name for NSX-T Edge Gateway Static Route
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies network prefix in CIDR format. Both IPv4 and IPv6 formats
	// are supported.
	NetworkCidr pulumi.StringOutput `pulumi:"networkCidr"`
	// A set of next hops to use within the static route. At least one is
	// required. See Next Hop for definition structure.
	//
	// <a id="next-hop"></a>
	NextHops NsxtEdgegatewayStaticRouteNextHopArrayOutput `pulumi:"nextHops"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrOutput `pulumi:"org"`
}

// NewNsxtEdgegatewayStaticRoute registers a new resource with the given unique name, arguments, and options.
func NewNsxtEdgegatewayStaticRoute(ctx *pulumi.Context,
	name string, args *NsxtEdgegatewayStaticRouteArgs, opts ...pulumi.ResourceOption) (*NsxtEdgegatewayStaticRoute, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EdgeGatewayId == nil {
		return nil, errors.New("invalid value for required argument 'EdgeGatewayId'")
	}
	if args.NetworkCidr == nil {
		return nil, errors.New("invalid value for required argument 'NetworkCidr'")
	}
	if args.NextHops == nil {
		return nil, errors.New("invalid value for required argument 'NextHops'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NsxtEdgegatewayStaticRoute
	err := ctx.RegisterResource("vcd:index/nsxtEdgegatewayStaticRoute:NsxtEdgegatewayStaticRoute", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNsxtEdgegatewayStaticRoute gets an existing NsxtEdgegatewayStaticRoute resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNsxtEdgegatewayStaticRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NsxtEdgegatewayStaticRouteState, opts ...pulumi.ResourceOption) (*NsxtEdgegatewayStaticRoute, error) {
	var resource NsxtEdgegatewayStaticRoute
	err := ctx.ReadResource("vcd:index/nsxtEdgegatewayStaticRoute:NsxtEdgegatewayStaticRoute", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NsxtEdgegatewayStaticRoute resources.
type nsxtEdgegatewayStaticRouteState struct {
	// Description for NSX-T Edge Gateway Static Route
	Description *string `pulumi:"description"`
	// NSX-T Edge Gateway ID
	EdgeGatewayId *string `pulumi:"edgeGatewayId"`
	// Name for NSX-T Edge Gateway Static Route
	Name *string `pulumi:"name"`
	// Specifies network prefix in CIDR format. Both IPv4 and IPv6 formats
	// are supported.
	NetworkCidr *string `pulumi:"networkCidr"`
	// A set of next hops to use within the static route. At least one is
	// required. See Next Hop for definition structure.
	//
	// <a id="next-hop"></a>
	NextHops []NsxtEdgegatewayStaticRouteNextHop `pulumi:"nextHops"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org *string `pulumi:"org"`
}

type NsxtEdgegatewayStaticRouteState struct {
	// Description for NSX-T Edge Gateway Static Route
	Description pulumi.StringPtrInput
	// NSX-T Edge Gateway ID
	EdgeGatewayId pulumi.StringPtrInput
	// Name for NSX-T Edge Gateway Static Route
	Name pulumi.StringPtrInput
	// Specifies network prefix in CIDR format. Both IPv4 and IPv6 formats
	// are supported.
	NetworkCidr pulumi.StringPtrInput
	// A set of next hops to use within the static route. At least one is
	// required. See Next Hop for definition structure.
	//
	// <a id="next-hop"></a>
	NextHops NsxtEdgegatewayStaticRouteNextHopArrayInput
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrInput
}

func (NsxtEdgegatewayStaticRouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*nsxtEdgegatewayStaticRouteState)(nil)).Elem()
}

type nsxtEdgegatewayStaticRouteArgs struct {
	// Description for NSX-T Edge Gateway Static Route
	Description *string `pulumi:"description"`
	// NSX-T Edge Gateway ID
	EdgeGatewayId string `pulumi:"edgeGatewayId"`
	// Name for NSX-T Edge Gateway Static Route
	Name *string `pulumi:"name"`
	// Specifies network prefix in CIDR format. Both IPv4 and IPv6 formats
	// are supported.
	NetworkCidr string `pulumi:"networkCidr"`
	// A set of next hops to use within the static route. At least one is
	// required. See Next Hop for definition structure.
	//
	// <a id="next-hop"></a>
	NextHops []NsxtEdgegatewayStaticRouteNextHop `pulumi:"nextHops"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org *string `pulumi:"org"`
}

// The set of arguments for constructing a NsxtEdgegatewayStaticRoute resource.
type NsxtEdgegatewayStaticRouteArgs struct {
	// Description for NSX-T Edge Gateway Static Route
	Description pulumi.StringPtrInput
	// NSX-T Edge Gateway ID
	EdgeGatewayId pulumi.StringInput
	// Name for NSX-T Edge Gateway Static Route
	Name pulumi.StringPtrInput
	// Specifies network prefix in CIDR format. Both IPv4 and IPv6 formats
	// are supported.
	NetworkCidr pulumi.StringInput
	// A set of next hops to use within the static route. At least one is
	// required. See Next Hop for definition structure.
	//
	// <a id="next-hop"></a>
	NextHops NsxtEdgegatewayStaticRouteNextHopArrayInput
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
	// different organizations
	Org pulumi.StringPtrInput
}

func (NsxtEdgegatewayStaticRouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nsxtEdgegatewayStaticRouteArgs)(nil)).Elem()
}

type NsxtEdgegatewayStaticRouteInput interface {
	pulumi.Input

	ToNsxtEdgegatewayStaticRouteOutput() NsxtEdgegatewayStaticRouteOutput
	ToNsxtEdgegatewayStaticRouteOutputWithContext(ctx context.Context) NsxtEdgegatewayStaticRouteOutput
}

func (*NsxtEdgegatewayStaticRoute) ElementType() reflect.Type {
	return reflect.TypeOf((**NsxtEdgegatewayStaticRoute)(nil)).Elem()
}

func (i *NsxtEdgegatewayStaticRoute) ToNsxtEdgegatewayStaticRouteOutput() NsxtEdgegatewayStaticRouteOutput {
	return i.ToNsxtEdgegatewayStaticRouteOutputWithContext(context.Background())
}

func (i *NsxtEdgegatewayStaticRoute) ToNsxtEdgegatewayStaticRouteOutputWithContext(ctx context.Context) NsxtEdgegatewayStaticRouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxtEdgegatewayStaticRouteOutput)
}

// NsxtEdgegatewayStaticRouteArrayInput is an input type that accepts NsxtEdgegatewayStaticRouteArray and NsxtEdgegatewayStaticRouteArrayOutput values.
// You can construct a concrete instance of `NsxtEdgegatewayStaticRouteArrayInput` via:
//
//	NsxtEdgegatewayStaticRouteArray{ NsxtEdgegatewayStaticRouteArgs{...} }
type NsxtEdgegatewayStaticRouteArrayInput interface {
	pulumi.Input

	ToNsxtEdgegatewayStaticRouteArrayOutput() NsxtEdgegatewayStaticRouteArrayOutput
	ToNsxtEdgegatewayStaticRouteArrayOutputWithContext(context.Context) NsxtEdgegatewayStaticRouteArrayOutput
}

type NsxtEdgegatewayStaticRouteArray []NsxtEdgegatewayStaticRouteInput

func (NsxtEdgegatewayStaticRouteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NsxtEdgegatewayStaticRoute)(nil)).Elem()
}

func (i NsxtEdgegatewayStaticRouteArray) ToNsxtEdgegatewayStaticRouteArrayOutput() NsxtEdgegatewayStaticRouteArrayOutput {
	return i.ToNsxtEdgegatewayStaticRouteArrayOutputWithContext(context.Background())
}

func (i NsxtEdgegatewayStaticRouteArray) ToNsxtEdgegatewayStaticRouteArrayOutputWithContext(ctx context.Context) NsxtEdgegatewayStaticRouteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxtEdgegatewayStaticRouteArrayOutput)
}

// NsxtEdgegatewayStaticRouteMapInput is an input type that accepts NsxtEdgegatewayStaticRouteMap and NsxtEdgegatewayStaticRouteMapOutput values.
// You can construct a concrete instance of `NsxtEdgegatewayStaticRouteMapInput` via:
//
//	NsxtEdgegatewayStaticRouteMap{ "key": NsxtEdgegatewayStaticRouteArgs{...} }
type NsxtEdgegatewayStaticRouteMapInput interface {
	pulumi.Input

	ToNsxtEdgegatewayStaticRouteMapOutput() NsxtEdgegatewayStaticRouteMapOutput
	ToNsxtEdgegatewayStaticRouteMapOutputWithContext(context.Context) NsxtEdgegatewayStaticRouteMapOutput
}

type NsxtEdgegatewayStaticRouteMap map[string]NsxtEdgegatewayStaticRouteInput

func (NsxtEdgegatewayStaticRouteMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NsxtEdgegatewayStaticRoute)(nil)).Elem()
}

func (i NsxtEdgegatewayStaticRouteMap) ToNsxtEdgegatewayStaticRouteMapOutput() NsxtEdgegatewayStaticRouteMapOutput {
	return i.ToNsxtEdgegatewayStaticRouteMapOutputWithContext(context.Background())
}

func (i NsxtEdgegatewayStaticRouteMap) ToNsxtEdgegatewayStaticRouteMapOutputWithContext(ctx context.Context) NsxtEdgegatewayStaticRouteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxtEdgegatewayStaticRouteMapOutput)
}

type NsxtEdgegatewayStaticRouteOutput struct{ *pulumi.OutputState }

func (NsxtEdgegatewayStaticRouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NsxtEdgegatewayStaticRoute)(nil)).Elem()
}

func (o NsxtEdgegatewayStaticRouteOutput) ToNsxtEdgegatewayStaticRouteOutput() NsxtEdgegatewayStaticRouteOutput {
	return o
}

func (o NsxtEdgegatewayStaticRouteOutput) ToNsxtEdgegatewayStaticRouteOutputWithContext(ctx context.Context) NsxtEdgegatewayStaticRouteOutput {
	return o
}

// Description for NSX-T Edge Gateway Static Route
func (o NsxtEdgegatewayStaticRouteOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsxtEdgegatewayStaticRoute) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// NSX-T Edge Gateway ID
func (o NsxtEdgegatewayStaticRouteOutput) EdgeGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxtEdgegatewayStaticRoute) pulumi.StringOutput { return v.EdgeGatewayId }).(pulumi.StringOutput)
}

// Name for NSX-T Edge Gateway Static Route
func (o NsxtEdgegatewayStaticRouteOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxtEdgegatewayStaticRoute) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Specifies network prefix in CIDR format. Both IPv4 and IPv6 formats
// are supported.
func (o NsxtEdgegatewayStaticRouteOutput) NetworkCidr() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxtEdgegatewayStaticRoute) pulumi.StringOutput { return v.NetworkCidr }).(pulumi.StringOutput)
}

// A set of next hops to use within the static route. At least one is
// required. See Next Hop for definition structure.
//
// <a id="next-hop"></a>
func (o NsxtEdgegatewayStaticRouteOutput) NextHops() NsxtEdgegatewayStaticRouteNextHopArrayOutput {
	return o.ApplyT(func(v *NsxtEdgegatewayStaticRoute) NsxtEdgegatewayStaticRouteNextHopArrayOutput { return v.NextHops }).(NsxtEdgegatewayStaticRouteNextHopArrayOutput)
}

// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
// different organizations
func (o NsxtEdgegatewayStaticRouteOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsxtEdgegatewayStaticRoute) pulumi.StringPtrOutput { return v.Org }).(pulumi.StringPtrOutput)
}

type NsxtEdgegatewayStaticRouteArrayOutput struct{ *pulumi.OutputState }

func (NsxtEdgegatewayStaticRouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NsxtEdgegatewayStaticRoute)(nil)).Elem()
}

func (o NsxtEdgegatewayStaticRouteArrayOutput) ToNsxtEdgegatewayStaticRouteArrayOutput() NsxtEdgegatewayStaticRouteArrayOutput {
	return o
}

func (o NsxtEdgegatewayStaticRouteArrayOutput) ToNsxtEdgegatewayStaticRouteArrayOutputWithContext(ctx context.Context) NsxtEdgegatewayStaticRouteArrayOutput {
	return o
}

func (o NsxtEdgegatewayStaticRouteArrayOutput) Index(i pulumi.IntInput) NsxtEdgegatewayStaticRouteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NsxtEdgegatewayStaticRoute {
		return vs[0].([]*NsxtEdgegatewayStaticRoute)[vs[1].(int)]
	}).(NsxtEdgegatewayStaticRouteOutput)
}

type NsxtEdgegatewayStaticRouteMapOutput struct{ *pulumi.OutputState }

func (NsxtEdgegatewayStaticRouteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NsxtEdgegatewayStaticRoute)(nil)).Elem()
}

func (o NsxtEdgegatewayStaticRouteMapOutput) ToNsxtEdgegatewayStaticRouteMapOutput() NsxtEdgegatewayStaticRouteMapOutput {
	return o
}

func (o NsxtEdgegatewayStaticRouteMapOutput) ToNsxtEdgegatewayStaticRouteMapOutputWithContext(ctx context.Context) NsxtEdgegatewayStaticRouteMapOutput {
	return o
}

func (o NsxtEdgegatewayStaticRouteMapOutput) MapIndex(k pulumi.StringInput) NsxtEdgegatewayStaticRouteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NsxtEdgegatewayStaticRoute {
		return vs[0].(map[string]*NsxtEdgegatewayStaticRoute)[vs[1].(string)]
	}).(NsxtEdgegatewayStaticRouteOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NsxtEdgegatewayStaticRouteInput)(nil)).Elem(), &NsxtEdgegatewayStaticRoute{})
	pulumi.RegisterInputType(reflect.TypeOf((*NsxtEdgegatewayStaticRouteArrayInput)(nil)).Elem(), NsxtEdgegatewayStaticRouteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NsxtEdgegatewayStaticRouteMapInput)(nil)).Elem(), NsxtEdgegatewayStaticRouteMap{})
	pulumi.RegisterOutputType(NsxtEdgegatewayStaticRouteOutput{})
	pulumi.RegisterOutputType(NsxtEdgegatewayStaticRouteArrayOutput{})
	pulumi.RegisterOutputType(NsxtEdgegatewayStaticRouteMapOutput{})
}
