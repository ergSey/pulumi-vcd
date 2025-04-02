// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/nsxtNetworkDhcp:NsxtNetworkDhcp")]
    public partial class NsxtNetworkDhcp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The DNS server IPs to be assigned by this
        /// DHCP service. Maximum two values.
        /// </summary>
        [Output("dnsServers")]
        public Output<ImmutableArray<string>> DnsServers { get; private set; } = null!;

        /// <summary>
        /// Lease time in seconds. Minimum value is 60s
        /// and maximum is 4294967295s (~ 49 days).
        /// </summary>
        [Output("leaseTime")]
        public Output<int> LeaseTime { get; private set; } = null!;

        /// <summary>
        /// IP address of DHCP server in network. Must match
        /// subnet. **Only** used when `mode=NETWORK`.
        /// </summary>
        [Output("listenerIpAddress")]
        public Output<string?> ListenerIpAddress { get; private set; } = null!;

        /// <summary>
        /// One of `EDGE`, `NETWORK` or `RELAY`. Default is `EDGE`
        /// * `EDGE` can be used with Routed Org VDC networks.
        /// * `NETWORK` can be used for Isolated and Routed Org VDC networks. It requires
        /// `listener_ip_address` to be set and Edge Cluster must be assigned to VDC.
        /// * `RELAY` can be used with Routed Org VDC networks, but requires DHCP forwarding configuration in
        /// NSX-T Edge Gateway.
        /// </summary>
        [Output("mode")]
        public Output<string?> Mode { get; private set; } = null!;

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful
        /// when connected as sysadmin working across different organisations.
        /// </summary>
        [Output("org")]
        public Output<string?> Org { get; private set; } = null!;

        /// <summary>
        /// ID of parent Org VDC Routed network.
        /// </summary>
        [Output("orgNetworkId")]
        public Output<string> OrgNetworkId { get; private set; } = null!;

        /// <summary>
        /// One or more blocks to define DHCP pool ranges. Must not be set when
        /// `mode=RELAY`. See Pools and example for usage details.
        /// </summary>
        [Output("pools")]
        public Output<ImmutableArray<Outputs.NsxtNetworkDhcpPool>> Pools { get; private set; } = null!;

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Output("vdc")]
        public Output<string> Vdc { get; private set; } = null!;


        /// <summary>
        /// Create a NsxtNetworkDhcp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NsxtNetworkDhcp(string name, NsxtNetworkDhcpArgs args, CustomResourceOptions? options = null)
            : base("vcd:index/nsxtNetworkDhcp:NsxtNetworkDhcp", name, args ?? new NsxtNetworkDhcpArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NsxtNetworkDhcp(string name, Input<string> id, NsxtNetworkDhcpState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/nsxtNetworkDhcp:NsxtNetworkDhcp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NsxtNetworkDhcp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NsxtNetworkDhcp Get(string name, Input<string> id, NsxtNetworkDhcpState? state = null, CustomResourceOptions? options = null)
        {
            return new NsxtNetworkDhcp(name, id, state, options);
        }
    }

    public sealed class NsxtNetworkDhcpArgs : global::Pulumi.ResourceArgs
    {
        [Input("dnsServers")]
        private InputList<string>? _dnsServers;

        /// <summary>
        /// The DNS server IPs to be assigned by this
        /// DHCP service. Maximum two values.
        /// </summary>
        public InputList<string> DnsServers
        {
            get => _dnsServers ?? (_dnsServers = new InputList<string>());
            set => _dnsServers = value;
        }

        /// <summary>
        /// Lease time in seconds. Minimum value is 60s
        /// and maximum is 4294967295s (~ 49 days).
        /// </summary>
        [Input("leaseTime")]
        public Input<int>? LeaseTime { get; set; }

        /// <summary>
        /// IP address of DHCP server in network. Must match
        /// subnet. **Only** used when `mode=NETWORK`.
        /// </summary>
        [Input("listenerIpAddress")]
        public Input<string>? ListenerIpAddress { get; set; }

        /// <summary>
        /// One of `EDGE`, `NETWORK` or `RELAY`. Default is `EDGE`
        /// * `EDGE` can be used with Routed Org VDC networks.
        /// * `NETWORK` can be used for Isolated and Routed Org VDC networks. It requires
        /// `listener_ip_address` to be set and Edge Cluster must be assigned to VDC.
        /// * `RELAY` can be used with Routed Org VDC networks, but requires DHCP forwarding configuration in
        /// NSX-T Edge Gateway.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful
        /// when connected as sysadmin working across different organisations.
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// ID of parent Org VDC Routed network.
        /// </summary>
        [Input("orgNetworkId", required: true)]
        public Input<string> OrgNetworkId { get; set; } = null!;

        [Input("pools")]
        private InputList<Inputs.NsxtNetworkDhcpPoolArgs>? _pools;

        /// <summary>
        /// One or more blocks to define DHCP pool ranges. Must not be set when
        /// `mode=RELAY`. See Pools and example for usage details.
        /// </summary>
        public InputList<Inputs.NsxtNetworkDhcpPoolArgs> Pools
        {
            get => _pools ?? (_pools = new InputList<Inputs.NsxtNetworkDhcpPoolArgs>());
            set => _pools = value;
        }

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public NsxtNetworkDhcpArgs()
        {
        }
        public static new NsxtNetworkDhcpArgs Empty => new NsxtNetworkDhcpArgs();
    }

    public sealed class NsxtNetworkDhcpState : global::Pulumi.ResourceArgs
    {
        [Input("dnsServers")]
        private InputList<string>? _dnsServers;

        /// <summary>
        /// The DNS server IPs to be assigned by this
        /// DHCP service. Maximum two values.
        /// </summary>
        public InputList<string> DnsServers
        {
            get => _dnsServers ?? (_dnsServers = new InputList<string>());
            set => _dnsServers = value;
        }

        /// <summary>
        /// Lease time in seconds. Minimum value is 60s
        /// and maximum is 4294967295s (~ 49 days).
        /// </summary>
        [Input("leaseTime")]
        public Input<int>? LeaseTime { get; set; }

        /// <summary>
        /// IP address of DHCP server in network. Must match
        /// subnet. **Only** used when `mode=NETWORK`.
        /// </summary>
        [Input("listenerIpAddress")]
        public Input<string>? ListenerIpAddress { get; set; }

        /// <summary>
        /// One of `EDGE`, `NETWORK` or `RELAY`. Default is `EDGE`
        /// * `EDGE` can be used with Routed Org VDC networks.
        /// * `NETWORK` can be used for Isolated and Routed Org VDC networks. It requires
        /// `listener_ip_address` to be set and Edge Cluster must be assigned to VDC.
        /// * `RELAY` can be used with Routed Org VDC networks, but requires DHCP forwarding configuration in
        /// NSX-T Edge Gateway.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful
        /// when connected as sysadmin working across different organisations.
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// ID of parent Org VDC Routed network.
        /// </summary>
        [Input("orgNetworkId")]
        public Input<string>? OrgNetworkId { get; set; }

        [Input("pools")]
        private InputList<Inputs.NsxtNetworkDhcpPoolGetArgs>? _pools;

        /// <summary>
        /// One or more blocks to define DHCP pool ranges. Must not be set when
        /// `mode=RELAY`. See Pools and example for usage details.
        /// </summary>
        public InputList<Inputs.NsxtNetworkDhcpPoolGetArgs> Pools
        {
            get => _pools ?? (_pools = new InputList<Inputs.NsxtNetworkDhcpPoolGetArgs>());
            set => _pools = value;
        }

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public NsxtNetworkDhcpState()
        {
        }
        public static new NsxtNetworkDhcpState Empty => new NsxtNetworkDhcpState();
    }
}
