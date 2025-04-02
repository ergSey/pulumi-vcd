// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class NsxvDistributedFirewallRuleApplicationGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Destination port for this application. Leaving it empty means 'any' port
        /// </summary>
        [Input("destinationPort")]
        public Input<string>? DestinationPort { get; set; }

        /// <summary>
        /// Name of application (Application, ApplicationGroup)
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Protocol of the application (one of TCP, UDP, ICMP) (When not using name/value)
        /// </summary>
        [Input("protocol")]
        public Input<string>? Protocol { get; set; }

        /// <summary>
        /// Source port for this application. Leaving it empty means 'any' port
        /// </summary>
        [Input("sourcePort")]
        public Input<string>? SourcePort { get; set; }

        /// <summary>
        /// Type of application
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Value of the application
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public NsxvDistributedFirewallRuleApplicationGetArgs()
        {
        }
        public static new NsxvDistributedFirewallRuleApplicationGetArgs Empty => new NsxvDistributedFirewallRuleApplicationGetArgs();
    }
}
