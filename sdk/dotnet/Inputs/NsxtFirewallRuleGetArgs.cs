// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class NsxtFirewallRuleGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defines if the rule should 'ALLOW', 'DROP' or 'REJECT' matching traffic
        /// </summary>
        [Input("action", required: true)]
        public Input<string> Action { get; set; } = null!;

        [Input("appPortProfileIds")]
        private InputList<string>? _appPortProfileIds;

        /// <summary>
        /// A set of Application Port Profile IDs. Leaving it empty means 'Any'
        /// </summary>
        public InputList<string> AppPortProfileIds
        {
            get => _appPortProfileIds ?? (_appPortProfileIds = new InputList<string>());
            set => _appPortProfileIds = value;
        }

        [Input("destinationIds")]
        private InputList<string>? _destinationIds;

        /// <summary>
        /// A set of Destination Firewall Group IDs (IP Sets or Security Groups). Leaving it empty means 'Any'
        /// </summary>
        public InputList<string> DestinationIds
        {
            get => _destinationIds ?? (_destinationIds = new InputList<string>());
            set => _destinationIds = value;
        }

        /// <summary>
        /// Direction on which Firewall Rule applies (One of 'IN', 'OUT', 'IN_OUT')
        /// </summary>
        [Input("direction", required: true)]
        public Input<string> Direction { get; set; } = null!;

        /// <summary>
        /// Defined if Firewall Rule is active
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Firewall Rule ID
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Firewall Rule Protocol (One of 'IPV4', 'IPV6', 'IPV4_IPV6')
        /// </summary>
        [Input("ipProtocol", required: true)]
        public Input<string> IpProtocol { get; set; } = null!;

        /// <summary>
        /// Defines if matching traffic should be logged
        /// </summary>
        [Input("logging")]
        public Input<bool>? Logging { get; set; }

        /// <summary>
        /// Firewall Rule name
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("sourceIds")]
        private InputList<string>? _sourceIds;

        /// <summary>
        /// A set of Source Firewall Group IDs (IP Sets or Security Groups). Leaving it empty means 'Any'
        /// </summary>
        public InputList<string> SourceIds
        {
            get => _sourceIds ?? (_sourceIds = new InputList<string>());
            set => _sourceIds = value;
        }

        public NsxtFirewallRuleGetArgs()
        {
        }
        public static new NsxtFirewallRuleGetArgs Empty => new NsxtFirewallRuleGetArgs();
    }
}
