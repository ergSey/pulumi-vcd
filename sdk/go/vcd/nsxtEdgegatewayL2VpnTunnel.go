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

type NsxtEdgegatewayL2VpnTunnel struct {
	pulumi.CustomResourceState

	// Mode in which the connection is formed.
	// Required for `SERVER` mode sessions. One of:
	// * `INITIATOR` - Local endpoint initiates tunnel setup and will also respond to
	//   incoming tunnel setup requests from the peer gateway.
	// * `RESPOND_ONLY` - Local endpoint shall only respond to incoming tunnel setup
	//   requests, it shall not initiate the tunnel setup.
	// * `ON_DEMAND` - In this mode local endpoint will initiate tunnel creation once
	//   first packet matching the policy rule is received, and will also respond to
	//   incoming initiation requests.
	ConnectorInitiationMode pulumi.StringPtrOutput `pulumi:"connectorInitiationMode"`
	// The description of the tunnel.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The ID of the Edge Gateway (NSX-T only).
	// Can be looked up using [`NsxtEdgegateway`](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_edgegateway) data source
	EdgeGatewayId pulumi.StringOutput `pulumi:"edgeGatewayId"`
	// State of the `SERVER` mode session, always set to `true` for `CLIENT`
	// mode sessions. Default is `true`.
	Enabled pulumi.BoolPtrOutput `pulumi:"enabled"`
	// The IP address corresponding to the Edge
	// Gateway the tunnel is being configured on. The IP must be sub-allocated
	// on the Edge Gateway.
	LocalEndpointIp pulumi.StringOutput `pulumi:"localEndpointIp"`
	// The name of the tunnel.
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of organization to use, optional if defined at
	// provider level. Useful when connected as sysadmin working across different organisations
	Org pulumi.StringPtrOutput `pulumi:"org"`
	// Base64 encoded string of the full configuration of the tunnel provided by the SERVER session. It is a computed field for
	// SERVER sessions and is a required field for CLIENT sessions.
	PeerCode pulumi.StringOutput `pulumi:"peerCode"`
	// The key that is used for authenticating the
	// connection. Required for `SERVER` mode sessions.
	PreSharedKey pulumi.StringPtrOutput `pulumi:"preSharedKey"`
	// The IP address of the remote endpoint, which
	// corresponds to the device on the remote site terminating the VPN tunnel.
	RemoteEndpointIp pulumi.StringOutput `pulumi:"remoteEndpointIp"`
	// Mode of the tunnel session (SERVER or CLIENT).
	SessionMode pulumi.StringOutput `pulumi:"sessionMode"`
	// One or more stretched networks for the tunnel.
	// See `stretchedNetwork` for more detail.
	StretchedNetworks NsxtEdgegatewayL2VpnTunnelStretchedNetworkArrayOutput `pulumi:"stretchedNetworks"`
	// The network CIDR block over which the session
	// interfaces. Relevant only for `SERVER` mode sessions. If not provided, Cloud
	// Director will attempt to automatically allocate a tunnel interface.
	TunnelInterface pulumi.StringOutput `pulumi:"tunnelInterface"`
}

// NewNsxtEdgegatewayL2VpnTunnel registers a new resource with the given unique name, arguments, and options.
func NewNsxtEdgegatewayL2VpnTunnel(ctx *pulumi.Context,
	name string, args *NsxtEdgegatewayL2VpnTunnelArgs, opts ...pulumi.ResourceOption) (*NsxtEdgegatewayL2VpnTunnel, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EdgeGatewayId == nil {
		return nil, errors.New("invalid value for required argument 'EdgeGatewayId'")
	}
	if args.LocalEndpointIp == nil {
		return nil, errors.New("invalid value for required argument 'LocalEndpointIp'")
	}
	if args.RemoteEndpointIp == nil {
		return nil, errors.New("invalid value for required argument 'RemoteEndpointIp'")
	}
	if args.SessionMode == nil {
		return nil, errors.New("invalid value for required argument 'SessionMode'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NsxtEdgegatewayL2VpnTunnel
	err := ctx.RegisterResource("vcd:index/nsxtEdgegatewayL2VpnTunnel:NsxtEdgegatewayL2VpnTunnel", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNsxtEdgegatewayL2VpnTunnel gets an existing NsxtEdgegatewayL2VpnTunnel resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNsxtEdgegatewayL2VpnTunnel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NsxtEdgegatewayL2VpnTunnelState, opts ...pulumi.ResourceOption) (*NsxtEdgegatewayL2VpnTunnel, error) {
	var resource NsxtEdgegatewayL2VpnTunnel
	err := ctx.ReadResource("vcd:index/nsxtEdgegatewayL2VpnTunnel:NsxtEdgegatewayL2VpnTunnel", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NsxtEdgegatewayL2VpnTunnel resources.
type nsxtEdgegatewayL2VpnTunnelState struct {
	// Mode in which the connection is formed.
	// Required for `SERVER` mode sessions. One of:
	// * `INITIATOR` - Local endpoint initiates tunnel setup and will also respond to
	//   incoming tunnel setup requests from the peer gateway.
	// * `RESPOND_ONLY` - Local endpoint shall only respond to incoming tunnel setup
	//   requests, it shall not initiate the tunnel setup.
	// * `ON_DEMAND` - In this mode local endpoint will initiate tunnel creation once
	//   first packet matching the policy rule is received, and will also respond to
	//   incoming initiation requests.
	ConnectorInitiationMode *string `pulumi:"connectorInitiationMode"`
	// The description of the tunnel.
	Description *string `pulumi:"description"`
	// The ID of the Edge Gateway (NSX-T only).
	// Can be looked up using [`NsxtEdgegateway`](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_edgegateway) data source
	EdgeGatewayId *string `pulumi:"edgeGatewayId"`
	// State of the `SERVER` mode session, always set to `true` for `CLIENT`
	// mode sessions. Default is `true`.
	Enabled *bool `pulumi:"enabled"`
	// The IP address corresponding to the Edge
	// Gateway the tunnel is being configured on. The IP must be sub-allocated
	// on the Edge Gateway.
	LocalEndpointIp *string `pulumi:"localEndpointIp"`
	// The name of the tunnel.
	Name *string `pulumi:"name"`
	// The name of organization to use, optional if defined at
	// provider level. Useful when connected as sysadmin working across different organisations
	Org *string `pulumi:"org"`
	// Base64 encoded string of the full configuration of the tunnel provided by the SERVER session. It is a computed field for
	// SERVER sessions and is a required field for CLIENT sessions.
	PeerCode *string `pulumi:"peerCode"`
	// The key that is used for authenticating the
	// connection. Required for `SERVER` mode sessions.
	PreSharedKey *string `pulumi:"preSharedKey"`
	// The IP address of the remote endpoint, which
	// corresponds to the device on the remote site terminating the VPN tunnel.
	RemoteEndpointIp *string `pulumi:"remoteEndpointIp"`
	// Mode of the tunnel session (SERVER or CLIENT).
	SessionMode *string `pulumi:"sessionMode"`
	// One or more stretched networks for the tunnel.
	// See `stretchedNetwork` for more detail.
	StretchedNetworks []NsxtEdgegatewayL2VpnTunnelStretchedNetwork `pulumi:"stretchedNetworks"`
	// The network CIDR block over which the session
	// interfaces. Relevant only for `SERVER` mode sessions. If not provided, Cloud
	// Director will attempt to automatically allocate a tunnel interface.
	TunnelInterface *string `pulumi:"tunnelInterface"`
}

type NsxtEdgegatewayL2VpnTunnelState struct {
	// Mode in which the connection is formed.
	// Required for `SERVER` mode sessions. One of:
	// * `INITIATOR` - Local endpoint initiates tunnel setup and will also respond to
	//   incoming tunnel setup requests from the peer gateway.
	// * `RESPOND_ONLY` - Local endpoint shall only respond to incoming tunnel setup
	//   requests, it shall not initiate the tunnel setup.
	// * `ON_DEMAND` - In this mode local endpoint will initiate tunnel creation once
	//   first packet matching the policy rule is received, and will also respond to
	//   incoming initiation requests.
	ConnectorInitiationMode pulumi.StringPtrInput
	// The description of the tunnel.
	Description pulumi.StringPtrInput
	// The ID of the Edge Gateway (NSX-T only).
	// Can be looked up using [`NsxtEdgegateway`](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_edgegateway) data source
	EdgeGatewayId pulumi.StringPtrInput
	// State of the `SERVER` mode session, always set to `true` for `CLIENT`
	// mode sessions. Default is `true`.
	Enabled pulumi.BoolPtrInput
	// The IP address corresponding to the Edge
	// Gateway the tunnel is being configured on. The IP must be sub-allocated
	// on the Edge Gateway.
	LocalEndpointIp pulumi.StringPtrInput
	// The name of the tunnel.
	Name pulumi.StringPtrInput
	// The name of organization to use, optional if defined at
	// provider level. Useful when connected as sysadmin working across different organisations
	Org pulumi.StringPtrInput
	// Base64 encoded string of the full configuration of the tunnel provided by the SERVER session. It is a computed field for
	// SERVER sessions and is a required field for CLIENT sessions.
	PeerCode pulumi.StringPtrInput
	// The key that is used for authenticating the
	// connection. Required for `SERVER` mode sessions.
	PreSharedKey pulumi.StringPtrInput
	// The IP address of the remote endpoint, which
	// corresponds to the device on the remote site terminating the VPN tunnel.
	RemoteEndpointIp pulumi.StringPtrInput
	// Mode of the tunnel session (SERVER or CLIENT).
	SessionMode pulumi.StringPtrInput
	// One or more stretched networks for the tunnel.
	// See `stretchedNetwork` for more detail.
	StretchedNetworks NsxtEdgegatewayL2VpnTunnelStretchedNetworkArrayInput
	// The network CIDR block over which the session
	// interfaces. Relevant only for `SERVER` mode sessions. If not provided, Cloud
	// Director will attempt to automatically allocate a tunnel interface.
	TunnelInterface pulumi.StringPtrInput
}

func (NsxtEdgegatewayL2VpnTunnelState) ElementType() reflect.Type {
	return reflect.TypeOf((*nsxtEdgegatewayL2VpnTunnelState)(nil)).Elem()
}

type nsxtEdgegatewayL2VpnTunnelArgs struct {
	// Mode in which the connection is formed.
	// Required for `SERVER` mode sessions. One of:
	// * `INITIATOR` - Local endpoint initiates tunnel setup and will also respond to
	//   incoming tunnel setup requests from the peer gateway.
	// * `RESPOND_ONLY` - Local endpoint shall only respond to incoming tunnel setup
	//   requests, it shall not initiate the tunnel setup.
	// * `ON_DEMAND` - In this mode local endpoint will initiate tunnel creation once
	//   first packet matching the policy rule is received, and will also respond to
	//   incoming initiation requests.
	ConnectorInitiationMode *string `pulumi:"connectorInitiationMode"`
	// The description of the tunnel.
	Description *string `pulumi:"description"`
	// The ID of the Edge Gateway (NSX-T only).
	// Can be looked up using [`NsxtEdgegateway`](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_edgegateway) data source
	EdgeGatewayId string `pulumi:"edgeGatewayId"`
	// State of the `SERVER` mode session, always set to `true` for `CLIENT`
	// mode sessions. Default is `true`.
	Enabled *bool `pulumi:"enabled"`
	// The IP address corresponding to the Edge
	// Gateway the tunnel is being configured on. The IP must be sub-allocated
	// on the Edge Gateway.
	LocalEndpointIp string `pulumi:"localEndpointIp"`
	// The name of the tunnel.
	Name *string `pulumi:"name"`
	// The name of organization to use, optional if defined at
	// provider level. Useful when connected as sysadmin working across different organisations
	Org *string `pulumi:"org"`
	// Base64 encoded string of the full configuration of the tunnel provided by the SERVER session. It is a computed field for
	// SERVER sessions and is a required field for CLIENT sessions.
	PeerCode *string `pulumi:"peerCode"`
	// The key that is used for authenticating the
	// connection. Required for `SERVER` mode sessions.
	PreSharedKey *string `pulumi:"preSharedKey"`
	// The IP address of the remote endpoint, which
	// corresponds to the device on the remote site terminating the VPN tunnel.
	RemoteEndpointIp string `pulumi:"remoteEndpointIp"`
	// Mode of the tunnel session (SERVER or CLIENT).
	SessionMode string `pulumi:"sessionMode"`
	// One or more stretched networks for the tunnel.
	// See `stretchedNetwork` for more detail.
	StretchedNetworks []NsxtEdgegatewayL2VpnTunnelStretchedNetwork `pulumi:"stretchedNetworks"`
	// The network CIDR block over which the session
	// interfaces. Relevant only for `SERVER` mode sessions. If not provided, Cloud
	// Director will attempt to automatically allocate a tunnel interface.
	TunnelInterface *string `pulumi:"tunnelInterface"`
}

// The set of arguments for constructing a NsxtEdgegatewayL2VpnTunnel resource.
type NsxtEdgegatewayL2VpnTunnelArgs struct {
	// Mode in which the connection is formed.
	// Required for `SERVER` mode sessions. One of:
	// * `INITIATOR` - Local endpoint initiates tunnel setup and will also respond to
	//   incoming tunnel setup requests from the peer gateway.
	// * `RESPOND_ONLY` - Local endpoint shall only respond to incoming tunnel setup
	//   requests, it shall not initiate the tunnel setup.
	// * `ON_DEMAND` - In this mode local endpoint will initiate tunnel creation once
	//   first packet matching the policy rule is received, and will also respond to
	//   incoming initiation requests.
	ConnectorInitiationMode pulumi.StringPtrInput
	// The description of the tunnel.
	Description pulumi.StringPtrInput
	// The ID of the Edge Gateway (NSX-T only).
	// Can be looked up using [`NsxtEdgegateway`](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_edgegateway) data source
	EdgeGatewayId pulumi.StringInput
	// State of the `SERVER` mode session, always set to `true` for `CLIENT`
	// mode sessions. Default is `true`.
	Enabled pulumi.BoolPtrInput
	// The IP address corresponding to the Edge
	// Gateway the tunnel is being configured on. The IP must be sub-allocated
	// on the Edge Gateway.
	LocalEndpointIp pulumi.StringInput
	// The name of the tunnel.
	Name pulumi.StringPtrInput
	// The name of organization to use, optional if defined at
	// provider level. Useful when connected as sysadmin working across different organisations
	Org pulumi.StringPtrInput
	// Base64 encoded string of the full configuration of the tunnel provided by the SERVER session. It is a computed field for
	// SERVER sessions and is a required field for CLIENT sessions.
	PeerCode pulumi.StringPtrInput
	// The key that is used for authenticating the
	// connection. Required for `SERVER` mode sessions.
	PreSharedKey pulumi.StringPtrInput
	// The IP address of the remote endpoint, which
	// corresponds to the device on the remote site terminating the VPN tunnel.
	RemoteEndpointIp pulumi.StringInput
	// Mode of the tunnel session (SERVER or CLIENT).
	SessionMode pulumi.StringInput
	// One or more stretched networks for the tunnel.
	// See `stretchedNetwork` for more detail.
	StretchedNetworks NsxtEdgegatewayL2VpnTunnelStretchedNetworkArrayInput
	// The network CIDR block over which the session
	// interfaces. Relevant only for `SERVER` mode sessions. If not provided, Cloud
	// Director will attempt to automatically allocate a tunnel interface.
	TunnelInterface pulumi.StringPtrInput
}

func (NsxtEdgegatewayL2VpnTunnelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nsxtEdgegatewayL2VpnTunnelArgs)(nil)).Elem()
}

type NsxtEdgegatewayL2VpnTunnelInput interface {
	pulumi.Input

	ToNsxtEdgegatewayL2VpnTunnelOutput() NsxtEdgegatewayL2VpnTunnelOutput
	ToNsxtEdgegatewayL2VpnTunnelOutputWithContext(ctx context.Context) NsxtEdgegatewayL2VpnTunnelOutput
}

func (*NsxtEdgegatewayL2VpnTunnel) ElementType() reflect.Type {
	return reflect.TypeOf((**NsxtEdgegatewayL2VpnTunnel)(nil)).Elem()
}

func (i *NsxtEdgegatewayL2VpnTunnel) ToNsxtEdgegatewayL2VpnTunnelOutput() NsxtEdgegatewayL2VpnTunnelOutput {
	return i.ToNsxtEdgegatewayL2VpnTunnelOutputWithContext(context.Background())
}

func (i *NsxtEdgegatewayL2VpnTunnel) ToNsxtEdgegatewayL2VpnTunnelOutputWithContext(ctx context.Context) NsxtEdgegatewayL2VpnTunnelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxtEdgegatewayL2VpnTunnelOutput)
}

// NsxtEdgegatewayL2VpnTunnelArrayInput is an input type that accepts NsxtEdgegatewayL2VpnTunnelArray and NsxtEdgegatewayL2VpnTunnelArrayOutput values.
// You can construct a concrete instance of `NsxtEdgegatewayL2VpnTunnelArrayInput` via:
//
//	NsxtEdgegatewayL2VpnTunnelArray{ NsxtEdgegatewayL2VpnTunnelArgs{...} }
type NsxtEdgegatewayL2VpnTunnelArrayInput interface {
	pulumi.Input

	ToNsxtEdgegatewayL2VpnTunnelArrayOutput() NsxtEdgegatewayL2VpnTunnelArrayOutput
	ToNsxtEdgegatewayL2VpnTunnelArrayOutputWithContext(context.Context) NsxtEdgegatewayL2VpnTunnelArrayOutput
}

type NsxtEdgegatewayL2VpnTunnelArray []NsxtEdgegatewayL2VpnTunnelInput

func (NsxtEdgegatewayL2VpnTunnelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NsxtEdgegatewayL2VpnTunnel)(nil)).Elem()
}

func (i NsxtEdgegatewayL2VpnTunnelArray) ToNsxtEdgegatewayL2VpnTunnelArrayOutput() NsxtEdgegatewayL2VpnTunnelArrayOutput {
	return i.ToNsxtEdgegatewayL2VpnTunnelArrayOutputWithContext(context.Background())
}

func (i NsxtEdgegatewayL2VpnTunnelArray) ToNsxtEdgegatewayL2VpnTunnelArrayOutputWithContext(ctx context.Context) NsxtEdgegatewayL2VpnTunnelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxtEdgegatewayL2VpnTunnelArrayOutput)
}

// NsxtEdgegatewayL2VpnTunnelMapInput is an input type that accepts NsxtEdgegatewayL2VpnTunnelMap and NsxtEdgegatewayL2VpnTunnelMapOutput values.
// You can construct a concrete instance of `NsxtEdgegatewayL2VpnTunnelMapInput` via:
//
//	NsxtEdgegatewayL2VpnTunnelMap{ "key": NsxtEdgegatewayL2VpnTunnelArgs{...} }
type NsxtEdgegatewayL2VpnTunnelMapInput interface {
	pulumi.Input

	ToNsxtEdgegatewayL2VpnTunnelMapOutput() NsxtEdgegatewayL2VpnTunnelMapOutput
	ToNsxtEdgegatewayL2VpnTunnelMapOutputWithContext(context.Context) NsxtEdgegatewayL2VpnTunnelMapOutput
}

type NsxtEdgegatewayL2VpnTunnelMap map[string]NsxtEdgegatewayL2VpnTunnelInput

func (NsxtEdgegatewayL2VpnTunnelMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NsxtEdgegatewayL2VpnTunnel)(nil)).Elem()
}

func (i NsxtEdgegatewayL2VpnTunnelMap) ToNsxtEdgegatewayL2VpnTunnelMapOutput() NsxtEdgegatewayL2VpnTunnelMapOutput {
	return i.ToNsxtEdgegatewayL2VpnTunnelMapOutputWithContext(context.Background())
}

func (i NsxtEdgegatewayL2VpnTunnelMap) ToNsxtEdgegatewayL2VpnTunnelMapOutputWithContext(ctx context.Context) NsxtEdgegatewayL2VpnTunnelMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxtEdgegatewayL2VpnTunnelMapOutput)
}

type NsxtEdgegatewayL2VpnTunnelOutput struct{ *pulumi.OutputState }

func (NsxtEdgegatewayL2VpnTunnelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NsxtEdgegatewayL2VpnTunnel)(nil)).Elem()
}

func (o NsxtEdgegatewayL2VpnTunnelOutput) ToNsxtEdgegatewayL2VpnTunnelOutput() NsxtEdgegatewayL2VpnTunnelOutput {
	return o
}

func (o NsxtEdgegatewayL2VpnTunnelOutput) ToNsxtEdgegatewayL2VpnTunnelOutputWithContext(ctx context.Context) NsxtEdgegatewayL2VpnTunnelOutput {
	return o
}

// Mode in which the connection is formed.
// Required for `SERVER` mode sessions. One of:
//   - `INITIATOR` - Local endpoint initiates tunnel setup and will also respond to
//     incoming tunnel setup requests from the peer gateway.
//   - `RESPOND_ONLY` - Local endpoint shall only respond to incoming tunnel setup
//     requests, it shall not initiate the tunnel setup.
//   - `ON_DEMAND` - In this mode local endpoint will initiate tunnel creation once
//     first packet matching the policy rule is received, and will also respond to
//     incoming initiation requests.
func (o NsxtEdgegatewayL2VpnTunnelOutput) ConnectorInitiationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsxtEdgegatewayL2VpnTunnel) pulumi.StringPtrOutput { return v.ConnectorInitiationMode }).(pulumi.StringPtrOutput)
}

// The description of the tunnel.
func (o NsxtEdgegatewayL2VpnTunnelOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsxtEdgegatewayL2VpnTunnel) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The ID of the Edge Gateway (NSX-T only).
// Can be looked up using [`NsxtEdgegateway`](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_edgegateway) data source
func (o NsxtEdgegatewayL2VpnTunnelOutput) EdgeGatewayId() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxtEdgegatewayL2VpnTunnel) pulumi.StringOutput { return v.EdgeGatewayId }).(pulumi.StringOutput)
}

// State of the `SERVER` mode session, always set to `true` for `CLIENT`
// mode sessions. Default is `true`.
func (o NsxtEdgegatewayL2VpnTunnelOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *NsxtEdgegatewayL2VpnTunnel) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

// The IP address corresponding to the Edge
// Gateway the tunnel is being configured on. The IP must be sub-allocated
// on the Edge Gateway.
func (o NsxtEdgegatewayL2VpnTunnelOutput) LocalEndpointIp() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxtEdgegatewayL2VpnTunnel) pulumi.StringOutput { return v.LocalEndpointIp }).(pulumi.StringOutput)
}

// The name of the tunnel.
func (o NsxtEdgegatewayL2VpnTunnelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxtEdgegatewayL2VpnTunnel) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of organization to use, optional if defined at
// provider level. Useful when connected as sysadmin working across different organisations
func (o NsxtEdgegatewayL2VpnTunnelOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsxtEdgegatewayL2VpnTunnel) pulumi.StringPtrOutput { return v.Org }).(pulumi.StringPtrOutput)
}

// Base64 encoded string of the full configuration of the tunnel provided by the SERVER session. It is a computed field for
// SERVER sessions and is a required field for CLIENT sessions.
func (o NsxtEdgegatewayL2VpnTunnelOutput) PeerCode() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxtEdgegatewayL2VpnTunnel) pulumi.StringOutput { return v.PeerCode }).(pulumi.StringOutput)
}

// The key that is used for authenticating the
// connection. Required for `SERVER` mode sessions.
func (o NsxtEdgegatewayL2VpnTunnelOutput) PreSharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsxtEdgegatewayL2VpnTunnel) pulumi.StringPtrOutput { return v.PreSharedKey }).(pulumi.StringPtrOutput)
}

// The IP address of the remote endpoint, which
// corresponds to the device on the remote site terminating the VPN tunnel.
func (o NsxtEdgegatewayL2VpnTunnelOutput) RemoteEndpointIp() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxtEdgegatewayL2VpnTunnel) pulumi.StringOutput { return v.RemoteEndpointIp }).(pulumi.StringOutput)
}

// Mode of the tunnel session (SERVER or CLIENT).
func (o NsxtEdgegatewayL2VpnTunnelOutput) SessionMode() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxtEdgegatewayL2VpnTunnel) pulumi.StringOutput { return v.SessionMode }).(pulumi.StringOutput)
}

// One or more stretched networks for the tunnel.
// See `stretchedNetwork` for more detail.
func (o NsxtEdgegatewayL2VpnTunnelOutput) StretchedNetworks() NsxtEdgegatewayL2VpnTunnelStretchedNetworkArrayOutput {
	return o.ApplyT(func(v *NsxtEdgegatewayL2VpnTunnel) NsxtEdgegatewayL2VpnTunnelStretchedNetworkArrayOutput {
		return v.StretchedNetworks
	}).(NsxtEdgegatewayL2VpnTunnelStretchedNetworkArrayOutput)
}

// The network CIDR block over which the session
// interfaces. Relevant only for `SERVER` mode sessions. If not provided, Cloud
// Director will attempt to automatically allocate a tunnel interface.
func (o NsxtEdgegatewayL2VpnTunnelOutput) TunnelInterface() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxtEdgegatewayL2VpnTunnel) pulumi.StringOutput { return v.TunnelInterface }).(pulumi.StringOutput)
}

type NsxtEdgegatewayL2VpnTunnelArrayOutput struct{ *pulumi.OutputState }

func (NsxtEdgegatewayL2VpnTunnelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NsxtEdgegatewayL2VpnTunnel)(nil)).Elem()
}

func (o NsxtEdgegatewayL2VpnTunnelArrayOutput) ToNsxtEdgegatewayL2VpnTunnelArrayOutput() NsxtEdgegatewayL2VpnTunnelArrayOutput {
	return o
}

func (o NsxtEdgegatewayL2VpnTunnelArrayOutput) ToNsxtEdgegatewayL2VpnTunnelArrayOutputWithContext(ctx context.Context) NsxtEdgegatewayL2VpnTunnelArrayOutput {
	return o
}

func (o NsxtEdgegatewayL2VpnTunnelArrayOutput) Index(i pulumi.IntInput) NsxtEdgegatewayL2VpnTunnelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NsxtEdgegatewayL2VpnTunnel {
		return vs[0].([]*NsxtEdgegatewayL2VpnTunnel)[vs[1].(int)]
	}).(NsxtEdgegatewayL2VpnTunnelOutput)
}

type NsxtEdgegatewayL2VpnTunnelMapOutput struct{ *pulumi.OutputState }

func (NsxtEdgegatewayL2VpnTunnelMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NsxtEdgegatewayL2VpnTunnel)(nil)).Elem()
}

func (o NsxtEdgegatewayL2VpnTunnelMapOutput) ToNsxtEdgegatewayL2VpnTunnelMapOutput() NsxtEdgegatewayL2VpnTunnelMapOutput {
	return o
}

func (o NsxtEdgegatewayL2VpnTunnelMapOutput) ToNsxtEdgegatewayL2VpnTunnelMapOutputWithContext(ctx context.Context) NsxtEdgegatewayL2VpnTunnelMapOutput {
	return o
}

func (o NsxtEdgegatewayL2VpnTunnelMapOutput) MapIndex(k pulumi.StringInput) NsxtEdgegatewayL2VpnTunnelOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NsxtEdgegatewayL2VpnTunnel {
		return vs[0].(map[string]*NsxtEdgegatewayL2VpnTunnel)[vs[1].(string)]
	}).(NsxtEdgegatewayL2VpnTunnelOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NsxtEdgegatewayL2VpnTunnelInput)(nil)).Elem(), &NsxtEdgegatewayL2VpnTunnel{})
	pulumi.RegisterInputType(reflect.TypeOf((*NsxtEdgegatewayL2VpnTunnelArrayInput)(nil)).Elem(), NsxtEdgegatewayL2VpnTunnelArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NsxtEdgegatewayL2VpnTunnelMapInput)(nil)).Elem(), NsxtEdgegatewayL2VpnTunnelMap{})
	pulumi.RegisterOutputType(NsxtEdgegatewayL2VpnTunnelOutput{})
	pulumi.RegisterOutputType(NsxtEdgegatewayL2VpnTunnelArrayOutput{})
	pulumi.RegisterOutputType(NsxtEdgegatewayL2VpnTunnelMapOutput{})
}
