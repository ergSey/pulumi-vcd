// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class NsxvDistributedFirewallRuleAppliedToArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the applied-to entity
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Type of the applied-to entity (one of Network, Edge, VirtualMachine, IPSet, VDC, Ipv4Address)
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Value of the applied-to entity
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public NsxvDistributedFirewallRuleAppliedToArgs()
        {
        }
        public static new NsxvDistributedFirewallRuleAppliedToArgs Empty => new NsxvDistributedFirewallRuleAppliedToArgs();
    }
}
