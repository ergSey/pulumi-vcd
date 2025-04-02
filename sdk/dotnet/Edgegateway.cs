// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/edgegateway:Edgegateway")]
    public partial class Edgegateway : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Configuration of the vShield edge VM for this gateway. One of: `compact`, `full` ("Large"), `x-large`, `full4` ("Quad Large").
        /// </summary>
        [Output("configuration")]
        public Output<string> Configuration { get; private set; } = null!;

        /// <summary>
        /// (*v2.6+*) - IP address of edge gateway used for default network
        /// </summary>
        [Output("defaultExternalNetworkIp")]
        public Output<string> DefaultExternalNetworkIp { get; private set; } = null!;

        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// If advanced networking enabled, also enable distributed
        /// routing. Default is `false`.
        /// </summary>
        [Output("distributedRouting")]
        public Output<bool?> DistributedRouting { get; private set; } = null!;

        /// <summary>
        /// (*v2.6+*) - A list of IP addresses assigned to edge gateway interfaces
        /// connected to external networks.
        /// </summary>
        [Output("externalNetworkIps")]
        public Output<ImmutableArray<string>> ExternalNetworkIps { get; private set; } = null!;

        /// <summary>
        /// One or more blocks defining external networks, their
        /// subnets, IP addresses and  IP pool suballocation attached to edge gateway interfaces. Details are
        /// in external network block below.
        /// </summary>
        [Output("externalNetworks")]
        public Output<ImmutableArray<Outputs.EdgegatewayExternalNetwork>> ExternalNetworks { get; private set; } = null!;

        /// <summary>
        /// When FIPS mode is enabled, any secure communication to or from
        /// the NSX Edge uses cryptographic algorithms or protocols that are allowed by United States Federal
        /// Information Processing Standards (FIPS). FIPS mode turns on the cipher suites that comply with
        /// FIPS. Default is `false`. **Note:** to use FIPS mode it must be enabled in vCD system settings.
        /// </summary>
        [Output("fipsModeEnabled")]
        public Output<bool?> FipsModeEnabled { get; private set; } = null!;

        /// <summary>
        /// Default firewall rule (last in the processing order) action.
        /// One of `accept` or `deny`. Default `deny`.
        /// 
        /// &lt;a id="external-network"&gt;&lt;/a&gt;
        /// </summary>
        [Output("fwDefaultRuleAction")]
        public Output<string?> FwDefaultRuleAction { get; private set; } = null!;

        /// <summary>
        /// Enable default firewall rule (last in the processing 
        /// order) logging. Default `false`.
        /// </summary>
        [Output("fwDefaultRuleLoggingEnabled")]
        public Output<bool?> FwDefaultRuleLoggingEnabled { get; private set; } = null!;

        /// <summary>
        /// Enable firewall. Default `true`. **Note:** Disabling Firewall will also
        /// disable NAT and other NAT dependent features like Load Balancer.
        /// </summary>
        [Output("fwEnabled")]
        public Output<bool?> FwEnabled { get; private set; } = null!;

        /// <summary>
        /// Enable high availability on this edge gateway. Default is `false`.
        /// </summary>
        [Output("haEnabled")]
        public Output<bool?> HaEnabled { get; private set; } = null!;

        /// <summary>
        /// Enable to configure the load balancer to use the faster L4
        /// engine rather than L7 engine. The L4 TCP VIP is processed before the edge gateway firewall so no
        /// `allow` firewall rule is required. Default is `false`. **Note:** L7 VIPs for HTTP and HTTPS are
        /// processed after the firewall, so when Acceleration Enabled is not selected, an edge gateway firewall
        /// rule must exist to allow access to the L7 VIP for those protocols. When Acceleration Enabled is
        /// selected and the server pool is in non-transparent mode, an SNAT rule is added, so you must ensure
        /// that the firewall is enabled on the edge gateway.
        /// </summary>
        [Output("lbAccelerationEnabled")]
        public Output<bool?> LbAccelerationEnabled { get; private set; } = null!;

        /// <summary>
        /// Enable load balancing. Default is `false`.
        /// </summary>
        [Output("lbEnabled")]
        public Output<bool?> LbEnabled { get; private set; } = null!;

        /// <summary>
        /// Enables the edge gateway load balancer to collect traffic logs.
        /// Default is `false`.
        /// </summary>
        [Output("lbLoggingEnabled")]
        public Output<bool?> LbLoggingEnabled { get; private set; } = null!;

        /// <summary>
        /// Choose the severity of events to be logged. One of `emergency`,
        /// `alert`, `critical`, `error`, `warning`, `notice`, `info`, `debug`
        /// </summary>
        [Output("lbLoglevel")]
        public Output<string?> LbLoglevel { get; private set; } = null!;

        /// <summary>
        /// A unique name for the edge gateway.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of organization to which the VDC belongs. Optional if defined at provider level.
        /// </summary>
        [Output("org")]
        public Output<string?> Org { get; private set; } = null!;

        /// <summary>
        /// When default route is set, it will be used for
        /// gateways' default routing and DNS forwarding. Default is `false`.
        /// </summary>
        [Output("useDefaultRouteForDnsRelay")]
        public Output<bool> UseDefaultRouteForDnsRelay { get; private set; } = null!;

        /// <summary>
        /// The name of VDC that owns the edge gateway. Optional if defined at provider level.
        /// </summary>
        [Output("vdc")]
        public Output<string?> Vdc { get; private set; } = null!;


        /// <summary>
        /// Create a Edgegateway resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Edgegateway(string name, EdgegatewayArgs args, CustomResourceOptions? options = null)
            : base("vcd:index/edgegateway:Edgegateway", name, args ?? new EdgegatewayArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Edgegateway(string name, Input<string> id, EdgegatewayState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/edgegateway:Edgegateway", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Edgegateway resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Edgegateway Get(string name, Input<string> id, EdgegatewayState? state = null, CustomResourceOptions? options = null)
        {
            return new Edgegateway(name, id, state, options);
        }
    }

    public sealed class EdgegatewayArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration of the vShield edge VM for this gateway. One of: `compact`, `full` ("Large"), `x-large`, `full4` ("Quad Large").
        /// </summary>
        [Input("configuration", required: true)]
        public Input<string> Configuration { get; set; } = null!;

        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// If advanced networking enabled, also enable distributed
        /// routing. Default is `false`.
        /// </summary>
        [Input("distributedRouting")]
        public Input<bool>? DistributedRouting { get; set; }

        [Input("externalNetworks", required: true)]
        private InputList<Inputs.EdgegatewayExternalNetworkArgs>? _externalNetworks;

        /// <summary>
        /// One or more blocks defining external networks, their
        /// subnets, IP addresses and  IP pool suballocation attached to edge gateway interfaces. Details are
        /// in external network block below.
        /// </summary>
        public InputList<Inputs.EdgegatewayExternalNetworkArgs> ExternalNetworks
        {
            get => _externalNetworks ?? (_externalNetworks = new InputList<Inputs.EdgegatewayExternalNetworkArgs>());
            set => _externalNetworks = value;
        }

        /// <summary>
        /// When FIPS mode is enabled, any secure communication to or from
        /// the NSX Edge uses cryptographic algorithms or protocols that are allowed by United States Federal
        /// Information Processing Standards (FIPS). FIPS mode turns on the cipher suites that comply with
        /// FIPS. Default is `false`. **Note:** to use FIPS mode it must be enabled in vCD system settings.
        /// </summary>
        [Input("fipsModeEnabled")]
        public Input<bool>? FipsModeEnabled { get; set; }

        /// <summary>
        /// Default firewall rule (last in the processing order) action.
        /// One of `accept` or `deny`. Default `deny`.
        /// 
        /// &lt;a id="external-network"&gt;&lt;/a&gt;
        /// </summary>
        [Input("fwDefaultRuleAction")]
        public Input<string>? FwDefaultRuleAction { get; set; }

        /// <summary>
        /// Enable default firewall rule (last in the processing 
        /// order) logging. Default `false`.
        /// </summary>
        [Input("fwDefaultRuleLoggingEnabled")]
        public Input<bool>? FwDefaultRuleLoggingEnabled { get; set; }

        /// <summary>
        /// Enable firewall. Default `true`. **Note:** Disabling Firewall will also
        /// disable NAT and other NAT dependent features like Load Balancer.
        /// </summary>
        [Input("fwEnabled")]
        public Input<bool>? FwEnabled { get; set; }

        /// <summary>
        /// Enable high availability on this edge gateway. Default is `false`.
        /// </summary>
        [Input("haEnabled")]
        public Input<bool>? HaEnabled { get; set; }

        /// <summary>
        /// Enable to configure the load balancer to use the faster L4
        /// engine rather than L7 engine. The L4 TCP VIP is processed before the edge gateway firewall so no
        /// `allow` firewall rule is required. Default is `false`. **Note:** L7 VIPs for HTTP and HTTPS are
        /// processed after the firewall, so when Acceleration Enabled is not selected, an edge gateway firewall
        /// rule must exist to allow access to the L7 VIP for those protocols. When Acceleration Enabled is
        /// selected and the server pool is in non-transparent mode, an SNAT rule is added, so you must ensure
        /// that the firewall is enabled on the edge gateway.
        /// </summary>
        [Input("lbAccelerationEnabled")]
        public Input<bool>? LbAccelerationEnabled { get; set; }

        /// <summary>
        /// Enable load balancing. Default is `false`.
        /// </summary>
        [Input("lbEnabled")]
        public Input<bool>? LbEnabled { get; set; }

        /// <summary>
        /// Enables the edge gateway load balancer to collect traffic logs.
        /// Default is `false`.
        /// </summary>
        [Input("lbLoggingEnabled")]
        public Input<bool>? LbLoggingEnabled { get; set; }

        /// <summary>
        /// Choose the severity of events to be logged. One of `emergency`,
        /// `alert`, `critical`, `error`, `warning`, `notice`, `info`, `debug`
        /// </summary>
        [Input("lbLoglevel")]
        public Input<string>? LbLoglevel { get; set; }

        /// <summary>
        /// A unique name for the edge gateway.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of organization to which the VDC belongs. Optional if defined at provider level.
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// When default route is set, it will be used for
        /// gateways' default routing and DNS forwarding. Default is `false`.
        /// </summary>
        [Input("useDefaultRouteForDnsRelay")]
        public Input<bool>? UseDefaultRouteForDnsRelay { get; set; }

        /// <summary>
        /// The name of VDC that owns the edge gateway. Optional if defined at provider level.
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public EdgegatewayArgs()
        {
        }
        public static new EdgegatewayArgs Empty => new EdgegatewayArgs();
    }

    public sealed class EdgegatewayState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration of the vShield edge VM for this gateway. One of: `compact`, `full` ("Large"), `x-large`, `full4` ("Quad Large").
        /// </summary>
        [Input("configuration")]
        public Input<string>? Configuration { get; set; }

        /// <summary>
        /// (*v2.6+*) - IP address of edge gateway used for default network
        /// </summary>
        [Input("defaultExternalNetworkIp")]
        public Input<string>? DefaultExternalNetworkIp { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// If advanced networking enabled, also enable distributed
        /// routing. Default is `false`.
        /// </summary>
        [Input("distributedRouting")]
        public Input<bool>? DistributedRouting { get; set; }

        [Input("externalNetworkIps")]
        private InputList<string>? _externalNetworkIps;

        /// <summary>
        /// (*v2.6+*) - A list of IP addresses assigned to edge gateway interfaces
        /// connected to external networks.
        /// </summary>
        public InputList<string> ExternalNetworkIps
        {
            get => _externalNetworkIps ?? (_externalNetworkIps = new InputList<string>());
            set => _externalNetworkIps = value;
        }

        [Input("externalNetworks")]
        private InputList<Inputs.EdgegatewayExternalNetworkGetArgs>? _externalNetworks;

        /// <summary>
        /// One or more blocks defining external networks, their
        /// subnets, IP addresses and  IP pool suballocation attached to edge gateway interfaces. Details are
        /// in external network block below.
        /// </summary>
        public InputList<Inputs.EdgegatewayExternalNetworkGetArgs> ExternalNetworks
        {
            get => _externalNetworks ?? (_externalNetworks = new InputList<Inputs.EdgegatewayExternalNetworkGetArgs>());
            set => _externalNetworks = value;
        }

        /// <summary>
        /// When FIPS mode is enabled, any secure communication to or from
        /// the NSX Edge uses cryptographic algorithms or protocols that are allowed by United States Federal
        /// Information Processing Standards (FIPS). FIPS mode turns on the cipher suites that comply with
        /// FIPS. Default is `false`. **Note:** to use FIPS mode it must be enabled in vCD system settings.
        /// </summary>
        [Input("fipsModeEnabled")]
        public Input<bool>? FipsModeEnabled { get; set; }

        /// <summary>
        /// Default firewall rule (last in the processing order) action.
        /// One of `accept` or `deny`. Default `deny`.
        /// 
        /// &lt;a id="external-network"&gt;&lt;/a&gt;
        /// </summary>
        [Input("fwDefaultRuleAction")]
        public Input<string>? FwDefaultRuleAction { get; set; }

        /// <summary>
        /// Enable default firewall rule (last in the processing 
        /// order) logging. Default `false`.
        /// </summary>
        [Input("fwDefaultRuleLoggingEnabled")]
        public Input<bool>? FwDefaultRuleLoggingEnabled { get; set; }

        /// <summary>
        /// Enable firewall. Default `true`. **Note:** Disabling Firewall will also
        /// disable NAT and other NAT dependent features like Load Balancer.
        /// </summary>
        [Input("fwEnabled")]
        public Input<bool>? FwEnabled { get; set; }

        /// <summary>
        /// Enable high availability on this edge gateway. Default is `false`.
        /// </summary>
        [Input("haEnabled")]
        public Input<bool>? HaEnabled { get; set; }

        /// <summary>
        /// Enable to configure the load balancer to use the faster L4
        /// engine rather than L7 engine. The L4 TCP VIP is processed before the edge gateway firewall so no
        /// `allow` firewall rule is required. Default is `false`. **Note:** L7 VIPs for HTTP and HTTPS are
        /// processed after the firewall, so when Acceleration Enabled is not selected, an edge gateway firewall
        /// rule must exist to allow access to the L7 VIP for those protocols. When Acceleration Enabled is
        /// selected and the server pool is in non-transparent mode, an SNAT rule is added, so you must ensure
        /// that the firewall is enabled on the edge gateway.
        /// </summary>
        [Input("lbAccelerationEnabled")]
        public Input<bool>? LbAccelerationEnabled { get; set; }

        /// <summary>
        /// Enable load balancing. Default is `false`.
        /// </summary>
        [Input("lbEnabled")]
        public Input<bool>? LbEnabled { get; set; }

        /// <summary>
        /// Enables the edge gateway load balancer to collect traffic logs.
        /// Default is `false`.
        /// </summary>
        [Input("lbLoggingEnabled")]
        public Input<bool>? LbLoggingEnabled { get; set; }

        /// <summary>
        /// Choose the severity of events to be logged. One of `emergency`,
        /// `alert`, `critical`, `error`, `warning`, `notice`, `info`, `debug`
        /// </summary>
        [Input("lbLoglevel")]
        public Input<string>? LbLoglevel { get; set; }

        /// <summary>
        /// A unique name for the edge gateway.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of organization to which the VDC belongs. Optional if defined at provider level.
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// When default route is set, it will be used for
        /// gateways' default routing and DNS forwarding. Default is `false`.
        /// </summary>
        [Input("useDefaultRouteForDnsRelay")]
        public Input<bool>? UseDefaultRouteForDnsRelay { get; set; }

        /// <summary>
        /// The name of VDC that owns the edge gateway. Optional if defined at provider level.
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public EdgegatewayState()
        {
        }
        public static new EdgegatewayState Empty => new EdgegatewayState();
    }
}
