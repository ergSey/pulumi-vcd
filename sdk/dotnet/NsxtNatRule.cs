// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/nsxtNatRule:NsxtNatRule")]
    public partial class NsxtNatRule : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Application Port Profile to which to apply the rule. The
        /// Application Port Profile includes a port, and a protocol that the incoming traffic uses on the edge
        /// gateway to connect to the internal network.  Can be looked up using `vcd.NsxtAppPortProfile`
        /// data source or created using `vcd.NsxtAppPortProfile` resource
        /// </summary>
        [Output("appPortProfileId")]
        public Output<string?> AppPortProfileId { get; private set; } = null!;

        /// <summary>
        /// An optional description of the NAT rule
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// For `DNAT` only. This represents the external port number or port range when doing 
        /// `DNAT` port forwarding from external to internal. The default dnatExternalPort is “ANY” meaning traffic on any port
        /// for the given IPs selected will be translated.
        /// </summary>
        [Output("dnatExternalPort")]
        public Output<string?> DnatExternalPort { get; private set; } = null!;

        /// <summary>
        /// The ID of the Edge Gateway (NSX-T only). Can be looked up using
        /// `vcd.NsxtEdgegateway` data source
        /// </summary>
        [Output("edgeGatewayId")]
        public Output<string> EdgeGatewayId { get; private set; } = null!;

        /// <summary>
        /// Enables or disables NAT rule (default `true`)
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// The external address for the NAT Rule. This must be supplied as a single IP or Network
        /// CIDR. For a `DNAT` rule, this is the external facing IP Address for incoming traffic. For an `SNAT` rule, this is the
        /// external facing IP Address for outgoing traffic. These IPs are typically allocated/suballocated IP Addresses on the
        /// Edge Gateway. For a `REFLEXIVE` rule, these are the external facing IPs.
        /// </summary>
        [Output("externalAddress")]
        public Output<string?> ExternalAddress { get; private set; } = null!;

        /// <summary>
        /// You can set a firewall match rule to determine how
        /// firewall is applied during NAT. One of `MATCH_INTERNAL_ADDRESS`, `MATCH_EXTERNAL_ADDRESS`,
        /// `BYPASS`
        /// * `MATCH_INTERNAL_ADDRESS` - applies firewall rules to the internal address of a NAT rule
        /// * `MATCH_EXTERNAL_ADDRESS` - applies firewall rules to the external address of a NAT rule
        /// * `BYPASS` - skip applying firewall rules to NAT rule
        /// </summary>
        [Output("firewallMatch")]
        public Output<string> FirewallMatch { get; private set; } = null!;

        /// <summary>
        /// The internal address for the NAT Rule. This must be supplied as a single IP or
        /// Network CIDR. For a `DNAT` rule, this is the internal IP address for incoming traffic. For an `SNAT` rule, this is the
        /// internal IP Address for outgoing traffic. For a `REFLEXIVE` rule, these are the internal IPs.
        /// These IPs are typically the Private IPs that are allocated to workloads.
        /// </summary>
        [Output("internalAddress")]
        public Output<string?> InternalAddress { get; private set; } = null!;

        /// <summary>
        /// Enable to have the address translation performed by this rule logged
        /// (default `false`). **Note** User might lack rights (**Organization Administrator** role by default
        /// is missing **Gateway &gt; Configure System Logging** right) to enable logging, but API does not
        /// return error and it is not possible to validate it. `pulumi preview` might show difference on
        /// every update.
        /// </summary>
        [Output("logging")]
        public Output<bool?> Logging { get; private set; } = null!;

        /// <summary>
        /// A name for NAT rule
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful
        /// when connected as sysadmin working across different organisations.
        /// </summary>
        [Output("org")]
        public Output<string?> Org { get; private set; } = null!;

        /// <summary>
        /// if an address has multiple NAT rules, you can assign these
        /// rules different priorities to determine the order in which they are applied. A lower value means a
        /// higher priority for this rule.
        /// </summary>
        [Output("priority")]
        public Output<int> Priority { get; private set; } = null!;

        /// <summary>
        /// One of `DNAT`, `NO_DNAT`, `SNAT`, `NO_SNAT`, `REFLEXIVE`
        /// * `DNAT` rule translates the external IP to an internal IP and is used for inbound traffic
        /// * `NO_DNAT` prevents external IP translation
        /// * `SNAT` translates an internal IP to an external IP and is used for outbound traffic
        /// * `NO_SNAT` prevents internal IP translation
        /// * `REFLEXIVE` (VCD 10.3+)  is also known as Stateless NAT. This translates an internal IP to an external IP and vice
        /// versa. The number of internal addresses should be exactly the same as that of external addresses.
        /// </summary>
        [Output("ruleType")]
        public Output<string> RuleType { get; private set; } = null!;

        /// <summary>
        /// For `SNAT` only. The destination addresses to match in the `SNAT` Rule. This 
        /// must be supplied as a single IP or Network CIDR. Providing no value for this field results in match with ANY
        /// destination network.
        /// </summary>
        [Output("snatDestinationAddress")]
        public Output<string?> SnatDestinationAddress { get; private set; } = null!;

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Output("vdc")]
        public Output<string> Vdc { get; private set; } = null!;


        /// <summary>
        /// Create a NsxtNatRule resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NsxtNatRule(string name, NsxtNatRuleArgs args, CustomResourceOptions? options = null)
            : base("vcd:index/nsxtNatRule:NsxtNatRule", name, args ?? new NsxtNatRuleArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NsxtNatRule(string name, Input<string> id, NsxtNatRuleState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/nsxtNatRule:NsxtNatRule", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NsxtNatRule resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NsxtNatRule Get(string name, Input<string> id, NsxtNatRuleState? state = null, CustomResourceOptions? options = null)
        {
            return new NsxtNatRule(name, id, state, options);
        }
    }

    public sealed class NsxtNatRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Application Port Profile to which to apply the rule. The
        /// Application Port Profile includes a port, and a protocol that the incoming traffic uses on the edge
        /// gateway to connect to the internal network.  Can be looked up using `vcd.NsxtAppPortProfile`
        /// data source or created using `vcd.NsxtAppPortProfile` resource
        /// </summary>
        [Input("appPortProfileId")]
        public Input<string>? AppPortProfileId { get; set; }

        /// <summary>
        /// An optional description of the NAT rule
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// For `DNAT` only. This represents the external port number or port range when doing 
        /// `DNAT` port forwarding from external to internal. The default dnatExternalPort is “ANY” meaning traffic on any port
        /// for the given IPs selected will be translated.
        /// </summary>
        [Input("dnatExternalPort")]
        public Input<string>? DnatExternalPort { get; set; }

        /// <summary>
        /// The ID of the Edge Gateway (NSX-T only). Can be looked up using
        /// `vcd.NsxtEdgegateway` data source
        /// </summary>
        [Input("edgeGatewayId", required: true)]
        public Input<string> EdgeGatewayId { get; set; } = null!;

        /// <summary>
        /// Enables or disables NAT rule (default `true`)
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The external address for the NAT Rule. This must be supplied as a single IP or Network
        /// CIDR. For a `DNAT` rule, this is the external facing IP Address for incoming traffic. For an `SNAT` rule, this is the
        /// external facing IP Address for outgoing traffic. These IPs are typically allocated/suballocated IP Addresses on the
        /// Edge Gateway. For a `REFLEXIVE` rule, these are the external facing IPs.
        /// </summary>
        [Input("externalAddress")]
        public Input<string>? ExternalAddress { get; set; }

        /// <summary>
        /// You can set a firewall match rule to determine how
        /// firewall is applied during NAT. One of `MATCH_INTERNAL_ADDRESS`, `MATCH_EXTERNAL_ADDRESS`,
        /// `BYPASS`
        /// * `MATCH_INTERNAL_ADDRESS` - applies firewall rules to the internal address of a NAT rule
        /// * `MATCH_EXTERNAL_ADDRESS` - applies firewall rules to the external address of a NAT rule
        /// * `BYPASS` - skip applying firewall rules to NAT rule
        /// </summary>
        [Input("firewallMatch")]
        public Input<string>? FirewallMatch { get; set; }

        /// <summary>
        /// The internal address for the NAT Rule. This must be supplied as a single IP or
        /// Network CIDR. For a `DNAT` rule, this is the internal IP address for incoming traffic. For an `SNAT` rule, this is the
        /// internal IP Address for outgoing traffic. For a `REFLEXIVE` rule, these are the internal IPs.
        /// These IPs are typically the Private IPs that are allocated to workloads.
        /// </summary>
        [Input("internalAddress")]
        public Input<string>? InternalAddress { get; set; }

        /// <summary>
        /// Enable to have the address translation performed by this rule logged
        /// (default `false`). **Note** User might lack rights (**Organization Administrator** role by default
        /// is missing **Gateway &gt; Configure System Logging** right) to enable logging, but API does not
        /// return error and it is not possible to validate it. `pulumi preview` might show difference on
        /// every update.
        /// </summary>
        [Input("logging")]
        public Input<bool>? Logging { get; set; }

        /// <summary>
        /// A name for NAT rule
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful
        /// when connected as sysadmin working across different organisations.
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// if an address has multiple NAT rules, you can assign these
        /// rules different priorities to determine the order in which they are applied. A lower value means a
        /// higher priority for this rule.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// One of `DNAT`, `NO_DNAT`, `SNAT`, `NO_SNAT`, `REFLEXIVE`
        /// * `DNAT` rule translates the external IP to an internal IP and is used for inbound traffic
        /// * `NO_DNAT` prevents external IP translation
        /// * `SNAT` translates an internal IP to an external IP and is used for outbound traffic
        /// * `NO_SNAT` prevents internal IP translation
        /// * `REFLEXIVE` (VCD 10.3+)  is also known as Stateless NAT. This translates an internal IP to an external IP and vice
        /// versa. The number of internal addresses should be exactly the same as that of external addresses.
        /// </summary>
        [Input("ruleType", required: true)]
        public Input<string> RuleType { get; set; } = null!;

        /// <summary>
        /// For `SNAT` only. The destination addresses to match in the `SNAT` Rule. This 
        /// must be supplied as a single IP or Network CIDR. Providing no value for this field results in match with ANY
        /// destination network.
        /// </summary>
        [Input("snatDestinationAddress")]
        public Input<string>? SnatDestinationAddress { get; set; }

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public NsxtNatRuleArgs()
        {
        }
        public static new NsxtNatRuleArgs Empty => new NsxtNatRuleArgs();
    }

    public sealed class NsxtNatRuleState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Application Port Profile to which to apply the rule. The
        /// Application Port Profile includes a port, and a protocol that the incoming traffic uses on the edge
        /// gateway to connect to the internal network.  Can be looked up using `vcd.NsxtAppPortProfile`
        /// data source or created using `vcd.NsxtAppPortProfile` resource
        /// </summary>
        [Input("appPortProfileId")]
        public Input<string>? AppPortProfileId { get; set; }

        /// <summary>
        /// An optional description of the NAT rule
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// For `DNAT` only. This represents the external port number or port range when doing 
        /// `DNAT` port forwarding from external to internal. The default dnatExternalPort is “ANY” meaning traffic on any port
        /// for the given IPs selected will be translated.
        /// </summary>
        [Input("dnatExternalPort")]
        public Input<string>? DnatExternalPort { get; set; }

        /// <summary>
        /// The ID of the Edge Gateway (NSX-T only). Can be looked up using
        /// `vcd.NsxtEdgegateway` data source
        /// </summary>
        [Input("edgeGatewayId")]
        public Input<string>? EdgeGatewayId { get; set; }

        /// <summary>
        /// Enables or disables NAT rule (default `true`)
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The external address for the NAT Rule. This must be supplied as a single IP or Network
        /// CIDR. For a `DNAT` rule, this is the external facing IP Address for incoming traffic. For an `SNAT` rule, this is the
        /// external facing IP Address for outgoing traffic. These IPs are typically allocated/suballocated IP Addresses on the
        /// Edge Gateway. For a `REFLEXIVE` rule, these are the external facing IPs.
        /// </summary>
        [Input("externalAddress")]
        public Input<string>? ExternalAddress { get; set; }

        /// <summary>
        /// You can set a firewall match rule to determine how
        /// firewall is applied during NAT. One of `MATCH_INTERNAL_ADDRESS`, `MATCH_EXTERNAL_ADDRESS`,
        /// `BYPASS`
        /// * `MATCH_INTERNAL_ADDRESS` - applies firewall rules to the internal address of a NAT rule
        /// * `MATCH_EXTERNAL_ADDRESS` - applies firewall rules to the external address of a NAT rule
        /// * `BYPASS` - skip applying firewall rules to NAT rule
        /// </summary>
        [Input("firewallMatch")]
        public Input<string>? FirewallMatch { get; set; }

        /// <summary>
        /// The internal address for the NAT Rule. This must be supplied as a single IP or
        /// Network CIDR. For a `DNAT` rule, this is the internal IP address for incoming traffic. For an `SNAT` rule, this is the
        /// internal IP Address for outgoing traffic. For a `REFLEXIVE` rule, these are the internal IPs.
        /// These IPs are typically the Private IPs that are allocated to workloads.
        /// </summary>
        [Input("internalAddress")]
        public Input<string>? InternalAddress { get; set; }

        /// <summary>
        /// Enable to have the address translation performed by this rule logged
        /// (default `false`). **Note** User might lack rights (**Organization Administrator** role by default
        /// is missing **Gateway &gt; Configure System Logging** right) to enable logging, but API does not
        /// return error and it is not possible to validate it. `pulumi preview` might show difference on
        /// every update.
        /// </summary>
        [Input("logging")]
        public Input<bool>? Logging { get; set; }

        /// <summary>
        /// A name for NAT rule
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful
        /// when connected as sysadmin working across different organisations.
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// if an address has multiple NAT rules, you can assign these
        /// rules different priorities to determine the order in which they are applied. A lower value means a
        /// higher priority for this rule.
        /// </summary>
        [Input("priority")]
        public Input<int>? Priority { get; set; }

        /// <summary>
        /// One of `DNAT`, `NO_DNAT`, `SNAT`, `NO_SNAT`, `REFLEXIVE`
        /// * `DNAT` rule translates the external IP to an internal IP and is used for inbound traffic
        /// * `NO_DNAT` prevents external IP translation
        /// * `SNAT` translates an internal IP to an external IP and is used for outbound traffic
        /// * `NO_SNAT` prevents internal IP translation
        /// * `REFLEXIVE` (VCD 10.3+)  is also known as Stateless NAT. This translates an internal IP to an external IP and vice
        /// versa. The number of internal addresses should be exactly the same as that of external addresses.
        /// </summary>
        [Input("ruleType")]
        public Input<string>? RuleType { get; set; }

        /// <summary>
        /// For `SNAT` only. The destination addresses to match in the `SNAT` Rule. This 
        /// must be supplied as a single IP or Network CIDR. Providing no value for this field results in match with ANY
        /// destination network.
        /// </summary>
        [Input("snatDestinationAddress")]
        public Input<string>? SnatDestinationAddress { get; set; }

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public NsxtNatRuleState()
        {
        }
        public static new NsxtNatRuleState Empty => new NsxtNatRuleState();
    }
}
