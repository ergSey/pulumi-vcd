// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/nsxtEdgegatewayL2VpnTunnel:NsxtEdgegatewayL2VpnTunnel")]
    public partial class NsxtEdgegatewayL2VpnTunnel : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Mode in which the connection is formed. 
        /// Required for `SERVER` mode sessions. One of:
        /// * `INITIATOR` - Local endpoint initiates tunnel setup and will also respond to
        /// incoming tunnel setup requests from the peer gateway.
        /// * `RESPOND_ONLY` - Local endpoint shall only respond to incoming tunnel setup
        /// requests, it shall not initiate the tunnel setup.
        /// * `ON_DEMAND` - In this mode local endpoint will initiate tunnel creation once
        /// first packet matching the policy rule is received, and will also respond to
        /// incoming initiation requests.
        /// </summary>
        [Output("connectorInitiationMode")]
        public Output<string?> ConnectorInitiationMode { get; private set; } = null!;

        /// <summary>
        /// The description of the tunnel.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The ID of the Edge Gateway (NSX-T only). 
        /// Can be looked up using [`vcd.NsxtEdgegateway`](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_edgegateway) data source
        /// </summary>
        [Output("edgeGatewayId")]
        public Output<string> EdgeGatewayId { get; private set; } = null!;

        /// <summary>
        /// State of the `SERVER` mode session, always set to `true` for `CLIENT` 
        /// mode sessions. Default is `true`.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// The IP address corresponding to the Edge 
        /// Gateway the tunnel is being configured on. The IP must be sub-allocated
        /// on the Edge Gateway.
        /// </summary>
        [Output("localEndpointIp")]
        public Output<string> LocalEndpointIp { get; private set; } = null!;

        /// <summary>
        /// The name of the tunnel.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of organization to use, optional if defined at 
        /// provider level. Useful when connected as sysadmin working across different organisations
        /// </summary>
        [Output("org")]
        public Output<string?> Org { get; private set; } = null!;

        /// <summary>
        /// Base64 encoded string of the full configuration of the tunnel provided by the SERVER session. It is a computed field for
        /// SERVER sessions and is a required field for CLIENT sessions.
        /// </summary>
        [Output("peerCode")]
        public Output<string> PeerCode { get; private set; } = null!;

        /// <summary>
        /// The key that is used for authenticating the 
        /// connection. Required for `SERVER` mode sessions.
        /// </summary>
        [Output("preSharedKey")]
        public Output<string?> PreSharedKey { get; private set; } = null!;

        /// <summary>
        /// The IP address of the remote endpoint, which 
        /// corresponds to the device on the remote site terminating the VPN tunnel.
        /// </summary>
        [Output("remoteEndpointIp")]
        public Output<string> RemoteEndpointIp { get; private set; } = null!;

        /// <summary>
        /// Mode of the tunnel session (SERVER or CLIENT).
        /// </summary>
        [Output("sessionMode")]
        public Output<string> SessionMode { get; private set; } = null!;

        /// <summary>
        /// One or more stretched networks for the tunnel. 
        /// See `stretched_network` for more detail.
        /// </summary>
        [Output("stretchedNetworks")]
        public Output<ImmutableArray<Outputs.NsxtEdgegatewayL2VpnTunnelStretchedNetwork>> StretchedNetworks { get; private set; } = null!;

        /// <summary>
        /// The network CIDR block over which the session 
        /// interfaces. Relevant only for `SERVER` mode sessions. If not provided, Cloud
        /// Director will attempt to automatically allocate a tunnel interface.
        /// </summary>
        [Output("tunnelInterface")]
        public Output<string> TunnelInterface { get; private set; } = null!;


        /// <summary>
        /// Create a NsxtEdgegatewayL2VpnTunnel resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NsxtEdgegatewayL2VpnTunnel(string name, NsxtEdgegatewayL2VpnTunnelArgs args, CustomResourceOptions? options = null)
            : base("vcd:index/nsxtEdgegatewayL2VpnTunnel:NsxtEdgegatewayL2VpnTunnel", name, args ?? new NsxtEdgegatewayL2VpnTunnelArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NsxtEdgegatewayL2VpnTunnel(string name, Input<string> id, NsxtEdgegatewayL2VpnTunnelState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/nsxtEdgegatewayL2VpnTunnel:NsxtEdgegatewayL2VpnTunnel", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/vmware/terraform-provider-vcd/releases/download/v3.14.1/",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing NsxtEdgegatewayL2VpnTunnel resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NsxtEdgegatewayL2VpnTunnel Get(string name, Input<string> id, NsxtEdgegatewayL2VpnTunnelState? state = null, CustomResourceOptions? options = null)
        {
            return new NsxtEdgegatewayL2VpnTunnel(name, id, state, options);
        }
    }

    public sealed class NsxtEdgegatewayL2VpnTunnelArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Mode in which the connection is formed. 
        /// Required for `SERVER` mode sessions. One of:
        /// * `INITIATOR` - Local endpoint initiates tunnel setup and will also respond to
        /// incoming tunnel setup requests from the peer gateway.
        /// * `RESPOND_ONLY` - Local endpoint shall only respond to incoming tunnel setup
        /// requests, it shall not initiate the tunnel setup.
        /// * `ON_DEMAND` - In this mode local endpoint will initiate tunnel creation once
        /// first packet matching the policy rule is received, and will also respond to
        /// incoming initiation requests.
        /// </summary>
        [Input("connectorInitiationMode")]
        public Input<string>? ConnectorInitiationMode { get; set; }

        /// <summary>
        /// The description of the tunnel.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the Edge Gateway (NSX-T only). 
        /// Can be looked up using [`vcd.NsxtEdgegateway`](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_edgegateway) data source
        /// </summary>
        [Input("edgeGatewayId", required: true)]
        public Input<string> EdgeGatewayId { get; set; } = null!;

        /// <summary>
        /// State of the `SERVER` mode session, always set to `true` for `CLIENT` 
        /// mode sessions. Default is `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The IP address corresponding to the Edge 
        /// Gateway the tunnel is being configured on. The IP must be sub-allocated
        /// on the Edge Gateway.
        /// </summary>
        [Input("localEndpointIp", required: true)]
        public Input<string> LocalEndpointIp { get; set; } = null!;

        /// <summary>
        /// The name of the tunnel.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at 
        /// provider level. Useful when connected as sysadmin working across different organisations
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// Base64 encoded string of the full configuration of the tunnel provided by the SERVER session. It is a computed field for
        /// SERVER sessions and is a required field for CLIENT sessions.
        /// </summary>
        [Input("peerCode")]
        public Input<string>? PeerCode { get; set; }

        /// <summary>
        /// The key that is used for authenticating the 
        /// connection. Required for `SERVER` mode sessions.
        /// </summary>
        [Input("preSharedKey")]
        public Input<string>? PreSharedKey { get; set; }

        /// <summary>
        /// The IP address of the remote endpoint, which 
        /// corresponds to the device on the remote site terminating the VPN tunnel.
        /// </summary>
        [Input("remoteEndpointIp", required: true)]
        public Input<string> RemoteEndpointIp { get; set; } = null!;

        /// <summary>
        /// Mode of the tunnel session (SERVER or CLIENT).
        /// </summary>
        [Input("sessionMode", required: true)]
        public Input<string> SessionMode { get; set; } = null!;

        [Input("stretchedNetworks")]
        private InputList<Inputs.NsxtEdgegatewayL2VpnTunnelStretchedNetworkArgs>? _stretchedNetworks;

        /// <summary>
        /// One or more stretched networks for the tunnel. 
        /// See `stretched_network` for more detail.
        /// </summary>
        public InputList<Inputs.NsxtEdgegatewayL2VpnTunnelStretchedNetworkArgs> StretchedNetworks
        {
            get => _stretchedNetworks ?? (_stretchedNetworks = new InputList<Inputs.NsxtEdgegatewayL2VpnTunnelStretchedNetworkArgs>());
            set => _stretchedNetworks = value;
        }

        /// <summary>
        /// The network CIDR block over which the session 
        /// interfaces. Relevant only for `SERVER` mode sessions. If not provided, Cloud
        /// Director will attempt to automatically allocate a tunnel interface.
        /// </summary>
        [Input("tunnelInterface")]
        public Input<string>? TunnelInterface { get; set; }

        public NsxtEdgegatewayL2VpnTunnelArgs()
        {
        }
        public static new NsxtEdgegatewayL2VpnTunnelArgs Empty => new NsxtEdgegatewayL2VpnTunnelArgs();
    }

    public sealed class NsxtEdgegatewayL2VpnTunnelState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Mode in which the connection is formed. 
        /// Required for `SERVER` mode sessions. One of:
        /// * `INITIATOR` - Local endpoint initiates tunnel setup and will also respond to
        /// incoming tunnel setup requests from the peer gateway.
        /// * `RESPOND_ONLY` - Local endpoint shall only respond to incoming tunnel setup
        /// requests, it shall not initiate the tunnel setup.
        /// * `ON_DEMAND` - In this mode local endpoint will initiate tunnel creation once
        /// first packet matching the policy rule is received, and will also respond to
        /// incoming initiation requests.
        /// </summary>
        [Input("connectorInitiationMode")]
        public Input<string>? ConnectorInitiationMode { get; set; }

        /// <summary>
        /// The description of the tunnel.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ID of the Edge Gateway (NSX-T only). 
        /// Can be looked up using [`vcd.NsxtEdgegateway`](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_edgegateway) data source
        /// </summary>
        [Input("edgeGatewayId")]
        public Input<string>? EdgeGatewayId { get; set; }

        /// <summary>
        /// State of the `SERVER` mode session, always set to `true` for `CLIENT` 
        /// mode sessions. Default is `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The IP address corresponding to the Edge 
        /// Gateway the tunnel is being configured on. The IP must be sub-allocated
        /// on the Edge Gateway.
        /// </summary>
        [Input("localEndpointIp")]
        public Input<string>? LocalEndpointIp { get; set; }

        /// <summary>
        /// The name of the tunnel.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at 
        /// provider level. Useful when connected as sysadmin working across different organisations
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// Base64 encoded string of the full configuration of the tunnel provided by the SERVER session. It is a computed field for
        /// SERVER sessions and is a required field for CLIENT sessions.
        /// </summary>
        [Input("peerCode")]
        public Input<string>? PeerCode { get; set; }

        /// <summary>
        /// The key that is used for authenticating the 
        /// connection. Required for `SERVER` mode sessions.
        /// </summary>
        [Input("preSharedKey")]
        public Input<string>? PreSharedKey { get; set; }

        /// <summary>
        /// The IP address of the remote endpoint, which 
        /// corresponds to the device on the remote site terminating the VPN tunnel.
        /// </summary>
        [Input("remoteEndpointIp")]
        public Input<string>? RemoteEndpointIp { get; set; }

        /// <summary>
        /// Mode of the tunnel session (SERVER or CLIENT).
        /// </summary>
        [Input("sessionMode")]
        public Input<string>? SessionMode { get; set; }

        [Input("stretchedNetworks")]
        private InputList<Inputs.NsxtEdgegatewayL2VpnTunnelStretchedNetworkGetArgs>? _stretchedNetworks;

        /// <summary>
        /// One or more stretched networks for the tunnel. 
        /// See `stretched_network` for more detail.
        /// </summary>
        public InputList<Inputs.NsxtEdgegatewayL2VpnTunnelStretchedNetworkGetArgs> StretchedNetworks
        {
            get => _stretchedNetworks ?? (_stretchedNetworks = new InputList<Inputs.NsxtEdgegatewayL2VpnTunnelStretchedNetworkGetArgs>());
            set => _stretchedNetworks = value;
        }

        /// <summary>
        /// The network CIDR block over which the session 
        /// interfaces. Relevant only for `SERVER` mode sessions. If not provided, Cloud
        /// Director will attempt to automatically allocate a tunnel interface.
        /// </summary>
        [Input("tunnelInterface")]
        public Input<string>? TunnelInterface { get; set; }

        public NsxtEdgegatewayL2VpnTunnelState()
        {
        }
        public static new NsxtEdgegatewayL2VpnTunnelState Empty => new NsxtEdgegatewayL2VpnTunnelState();
    }
}
