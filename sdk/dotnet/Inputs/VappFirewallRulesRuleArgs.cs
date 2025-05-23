// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class VappFirewallRulesRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Destination IP address to which the rule applies. A value of `Any` matches any IP address.
        /// </summary>
        [Input("destinationIp")]
        public Input<string>? DestinationIp { get; set; }

        /// <summary>
        /// Destination port to which this rule applies.
        /// </summary>
        [Input("destinationPort")]
        public Input<string>? DestinationPort { get; set; }

        /// <summary>
        /// Destination VM identifier
        /// </summary>
        [Input("destinationVmId")]
        public Input<string>? DestinationVmId { get; set; }

        /// <summary>
        /// The value can be one of: `assigned` - assigned internal IP will be automatically chosen. `NAT`: NATed external IP will be automatically chosen.
        /// </summary>
        [Input("destinationVmIpType")]
        public Input<string>? DestinationVmIpType { get; set; }

        /// <summary>
        /// Destination VM NIC ID to which this rule applies.
        /// </summary>
        [Input("destinationVmNicId")]
        public Input<int>? DestinationVmNicId { get; set; }

        /// <summary>
        /// 'true' value will enable rule logging. Default is false
        /// </summary>
        [Input("enableLogging")]
        public Input<bool>? EnableLogging { get; set; }

        /// <summary>
        /// Enable or disable firewall. Default is `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Rule name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// One of: `drop` (drop packets that match the rule), `allow` (allow packets that match the rule to pass through the firewall)
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        /// <summary>
        /// Specify the protocols to which the rule should be applied. One of: `any`, `icmp`, `tcp`, `udp`, `tcp&amp;udp`
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Source IP address to which the rule applies. A value of `Any` matches any IP address.
        /// </summary>
        [Input("sourceIp")]
        public Input<string>? SourceIp { get; set; }

        /// <summary>
        /// Source port to which this rule applies.
        /// </summary>
        [Input("sourcePort")]
        public Input<string>? SourcePort { get; set; }

        /// <summary>
        /// Source VM identifier
        /// </summary>
        [Input("sourceVmId")]
        public Input<string>? SourceVmId { get; set; }

        /// <summary>
        /// The value can be one of: `assigned` - assigned internal IP will be automatically chosen. `NAT`: NATed external IP will be automatically chosen.
        /// </summary>
        [Input("sourceVmIpType")]
        public Input<string>? SourceVmIpType { get; set; }

        /// <summary>
        /// Source VM NIC ID to which this rule applies.
        /// </summary>
        [Input("sourceVmNicId")]
        public Input<int>? SourceVmNicId { get; set; }

        public VappFirewallRulesRuleArgs()
        {
        }
        public static new VappFirewallRulesRuleArgs Empty => new VappFirewallRulesRuleArgs();
    }
}
