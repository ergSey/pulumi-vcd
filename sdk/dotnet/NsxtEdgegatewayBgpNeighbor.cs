// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/nsxtEdgegatewayBgpNeighbor:NsxtEdgegatewayBgpNeighbor")]
    public partial class NsxtEdgegatewayBgpNeighbor : global::Pulumi.CustomResource
    {
        /// <summary>
        /// BGP Allow-as-in feature is used to allow the BGP speaker to accept the BGP updates even if its own BGP AS number is in the AS-Path attribute.
        /// </summary>
        [Output("allowAsIn")]
        public Output<bool> AllowAsIn { get; private set; } = null!;

        /// <summary>
        /// Number of times a heartbeat packet is missed before BFD declares that the neighbor is down
        /// </summary>
        [Output("bfdDeadMultiple")]
        public Output<int> BfdDeadMultiple { get; private set; } = null!;

        /// <summary>
        /// Should Bidirectional Forwarding Detection (BFD) be enabled
        /// </summary>
        [Output("bfdEnabled")]
        public Output<bool> BfdEnabled { get; private set; } = null!;

        /// <summary>
        /// Time interval (in milliseconds) between heartbeat packets
        /// </summary>
        [Output("bfdInterval")]
        public Output<int> BfdInterval { get; private set; } = null!;

        /// <summary>
        /// The ID of the edge gateway (NSX-T only). Can be looked up using
        /// `vcd.NsxtEdgegateway` datasource
        /// </summary>
        [Output("edgeGatewayId")]
        public Output<string> EdgeGatewayId { get; private set; } = null!;

        /// <summary>
        /// BGP Neighbor Graceful Restart Mode. One of:
        /// * `DISABLE` - Overrides the global edge gateway settings and disables graceful restart mode for this neighbor.
        /// * `HELPER_ONLY` - Overrides the global edge gateway settings and configures graceful restart mode as Helper only for this neighbor.
        /// * `GRACEFUL_AND_HELPER` - Overrides the global edge gateway settings and configures graceful restart mode as Graceful restart and Helper for this neighbor.
        /// </summary>
        [Output("gracefulRestartMode")]
        public Output<string> GracefulRestartMode { get; private set; } = null!;

        /// <summary>
        /// Time interval (in seconds) before declaring a BGP peer dead
        /// </summary>
        [Output("holdDownTimer")]
        public Output<int> HoldDownTimer { get; private set; } = null!;

        /// <summary>
        /// The ID of the IP Prefix List to be used for filtering incoming BGP routes
        /// </summary>
        [Output("inFilterIpPrefixListId")]
        public Output<string> InFilterIpPrefixListId { get; private set; } = null!;

        /// <summary>
        /// BGP Neighbor IP Address (IPv4 or IPv6)
        /// </summary>
        [Output("ipAddress")]
        public Output<string> IpAddress { get; private set; } = null!;

        /// <summary>
        /// Time interval (in seconds) between sending keep-alive messages to a BGP peer
        /// </summary>
        [Output("keepAliveTimer")]
        public Output<int> KeepAliveTimer { get; private set; } = null!;

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful
        /// when connected as sysadmin working across different organisations
        /// </summary>
        [Output("org")]
        public Output<string?> Org { get; private set; } = null!;

        /// <summary>
        /// The ID of the IP Prefix List to be used for filtering outgoing BGP routes
        /// </summary>
        [Output("outFilterIpPrefixListId")]
        public Output<string> OutFilterIpPrefixListId { get; private set; } = null!;

        /// <summary>
        /// BGP Neighbor Password
        /// </summary>
        [Output("password")]
        public Output<string?> Password { get; private set; } = null!;

        /// <summary>
        /// BGP Neighbor Remote Autonomous System (AS) Number
        /// </summary>
        [Output("remoteAsNumber")]
        public Output<string> RemoteAsNumber { get; private set; } = null!;

        /// <summary>
        /// Route filtering by IP Address family. One of `DISABLED`, `IPV4`, `IPV6`
        /// </summary>
        [Output("routeFiltering")]
        public Output<string> RouteFiltering { get; private set; } = null!;


        /// <summary>
        /// Create a NsxtEdgegatewayBgpNeighbor resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NsxtEdgegatewayBgpNeighbor(string name, NsxtEdgegatewayBgpNeighborArgs args, CustomResourceOptions? options = null)
            : base("vcd:index/nsxtEdgegatewayBgpNeighbor:NsxtEdgegatewayBgpNeighbor", name, args ?? new NsxtEdgegatewayBgpNeighborArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NsxtEdgegatewayBgpNeighbor(string name, Input<string> id, NsxtEdgegatewayBgpNeighborState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/nsxtEdgegatewayBgpNeighbor:NsxtEdgegatewayBgpNeighbor", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/vmware/terraform-provider-vcd/releases/download/v3.14.1/",
                AdditionalSecretOutputs =
                {
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing NsxtEdgegatewayBgpNeighbor resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NsxtEdgegatewayBgpNeighbor Get(string name, Input<string> id, NsxtEdgegatewayBgpNeighborState? state = null, CustomResourceOptions? options = null)
        {
            return new NsxtEdgegatewayBgpNeighbor(name, id, state, options);
        }
    }

    public sealed class NsxtEdgegatewayBgpNeighborArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// BGP Allow-as-in feature is used to allow the BGP speaker to accept the BGP updates even if its own BGP AS number is in the AS-Path attribute.
        /// </summary>
        [Input("allowAsIn")]
        public Input<bool>? AllowAsIn { get; set; }

        /// <summary>
        /// Number of times a heartbeat packet is missed before BFD declares that the neighbor is down
        /// </summary>
        [Input("bfdDeadMultiple")]
        public Input<int>? BfdDeadMultiple { get; set; }

        /// <summary>
        /// Should Bidirectional Forwarding Detection (BFD) be enabled
        /// </summary>
        [Input("bfdEnabled")]
        public Input<bool>? BfdEnabled { get; set; }

        /// <summary>
        /// Time interval (in milliseconds) between heartbeat packets
        /// </summary>
        [Input("bfdInterval")]
        public Input<int>? BfdInterval { get; set; }

        /// <summary>
        /// The ID of the edge gateway (NSX-T only). Can be looked up using
        /// `vcd.NsxtEdgegateway` datasource
        /// </summary>
        [Input("edgeGatewayId", required: true)]
        public Input<string> EdgeGatewayId { get; set; } = null!;

        /// <summary>
        /// BGP Neighbor Graceful Restart Mode. One of:
        /// * `DISABLE` - Overrides the global edge gateway settings and disables graceful restart mode for this neighbor.
        /// * `HELPER_ONLY` - Overrides the global edge gateway settings and configures graceful restart mode as Helper only for this neighbor.
        /// * `GRACEFUL_AND_HELPER` - Overrides the global edge gateway settings and configures graceful restart mode as Graceful restart and Helper for this neighbor.
        /// </summary>
        [Input("gracefulRestartMode")]
        public Input<string>? GracefulRestartMode { get; set; }

        /// <summary>
        /// Time interval (in seconds) before declaring a BGP peer dead
        /// </summary>
        [Input("holdDownTimer")]
        public Input<int>? HoldDownTimer { get; set; }

        /// <summary>
        /// The ID of the IP Prefix List to be used for filtering incoming BGP routes
        /// </summary>
        [Input("inFilterIpPrefixListId")]
        public Input<string>? InFilterIpPrefixListId { get; set; }

        /// <summary>
        /// BGP Neighbor IP Address (IPv4 or IPv6)
        /// </summary>
        [Input("ipAddress", required: true)]
        public Input<string> IpAddress { get; set; } = null!;

        /// <summary>
        /// Time interval (in seconds) between sending keep-alive messages to a BGP peer
        /// </summary>
        [Input("keepAliveTimer")]
        public Input<int>? KeepAliveTimer { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful
        /// when connected as sysadmin working across different organisations
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// The ID of the IP Prefix List to be used for filtering outgoing BGP routes
        /// </summary>
        [Input("outFilterIpPrefixListId")]
        public Input<string>? OutFilterIpPrefixListId { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// BGP Neighbor Password
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// BGP Neighbor Remote Autonomous System (AS) Number
        /// </summary>
        [Input("remoteAsNumber", required: true)]
        public Input<string> RemoteAsNumber { get; set; } = null!;

        /// <summary>
        /// Route filtering by IP Address family. One of `DISABLED`, `IPV4`, `IPV6`
        /// </summary>
        [Input("routeFiltering")]
        public Input<string>? RouteFiltering { get; set; }

        public NsxtEdgegatewayBgpNeighborArgs()
        {
        }
        public static new NsxtEdgegatewayBgpNeighborArgs Empty => new NsxtEdgegatewayBgpNeighborArgs();
    }

    public sealed class NsxtEdgegatewayBgpNeighborState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// BGP Allow-as-in feature is used to allow the BGP speaker to accept the BGP updates even if its own BGP AS number is in the AS-Path attribute.
        /// </summary>
        [Input("allowAsIn")]
        public Input<bool>? AllowAsIn { get; set; }

        /// <summary>
        /// Number of times a heartbeat packet is missed before BFD declares that the neighbor is down
        /// </summary>
        [Input("bfdDeadMultiple")]
        public Input<int>? BfdDeadMultiple { get; set; }

        /// <summary>
        /// Should Bidirectional Forwarding Detection (BFD) be enabled
        /// </summary>
        [Input("bfdEnabled")]
        public Input<bool>? BfdEnabled { get; set; }

        /// <summary>
        /// Time interval (in milliseconds) between heartbeat packets
        /// </summary>
        [Input("bfdInterval")]
        public Input<int>? BfdInterval { get; set; }

        /// <summary>
        /// The ID of the edge gateway (NSX-T only). Can be looked up using
        /// `vcd.NsxtEdgegateway` datasource
        /// </summary>
        [Input("edgeGatewayId")]
        public Input<string>? EdgeGatewayId { get; set; }

        /// <summary>
        /// BGP Neighbor Graceful Restart Mode. One of:
        /// * `DISABLE` - Overrides the global edge gateway settings and disables graceful restart mode for this neighbor.
        /// * `HELPER_ONLY` - Overrides the global edge gateway settings and configures graceful restart mode as Helper only for this neighbor.
        /// * `GRACEFUL_AND_HELPER` - Overrides the global edge gateway settings and configures graceful restart mode as Graceful restart and Helper for this neighbor.
        /// </summary>
        [Input("gracefulRestartMode")]
        public Input<string>? GracefulRestartMode { get; set; }

        /// <summary>
        /// Time interval (in seconds) before declaring a BGP peer dead
        /// </summary>
        [Input("holdDownTimer")]
        public Input<int>? HoldDownTimer { get; set; }

        /// <summary>
        /// The ID of the IP Prefix List to be used for filtering incoming BGP routes
        /// </summary>
        [Input("inFilterIpPrefixListId")]
        public Input<string>? InFilterIpPrefixListId { get; set; }

        /// <summary>
        /// BGP Neighbor IP Address (IPv4 or IPv6)
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// Time interval (in seconds) between sending keep-alive messages to a BGP peer
        /// </summary>
        [Input("keepAliveTimer")]
        public Input<int>? KeepAliveTimer { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful
        /// when connected as sysadmin working across different organisations
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// The ID of the IP Prefix List to be used for filtering outgoing BGP routes
        /// </summary>
        [Input("outFilterIpPrefixListId")]
        public Input<string>? OutFilterIpPrefixListId { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// BGP Neighbor Password
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// BGP Neighbor Remote Autonomous System (AS) Number
        /// </summary>
        [Input("remoteAsNumber")]
        public Input<string>? RemoteAsNumber { get; set; }

        /// <summary>
        /// Route filtering by IP Address family. One of `DISABLED`, `IPV4`, `IPV6`
        /// </summary>
        [Input("routeFiltering")]
        public Input<string>? RouteFiltering { get; set; }

        public NsxtEdgegatewayBgpNeighborState()
        {
        }
        public static new NsxtEdgegatewayBgpNeighborState Empty => new NsxtEdgegatewayBgpNeighborState();
    }
}
