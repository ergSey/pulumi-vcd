// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Outputs
{

    [OutputType]
    public sealed class GetNsxtEdgegatewayBgpIpPrefixListIpPrefixResult
    {
        /// <summary>
        /// Action 'PERMIT' or 'DENY'
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// Greater than or equal to (ge) subnet mask
        /// </summary>
        public readonly int GreaterThanOrEqualTo;
        /// <summary>
        /// Less than or equal to (le) subnet mask
        /// </summary>
        public readonly int LessThanOrEqualTo;
        /// <summary>
        /// Network in CIDR notation (e.g. '192.168.100.0/24', '2001:db8::/48')
        /// </summary>
        public readonly string Network;

        [OutputConstructor]
        private GetNsxtEdgegatewayBgpIpPrefixListIpPrefixResult(
            string action,

            int greaterThanOrEqualTo,

            int lessThanOrEqualTo,

            string network)
        {
            Action = action;
            GreaterThanOrEqualTo = greaterThanOrEqualTo;
            LessThanOrEqualTo = lessThanOrEqualTo;
            Network = network;
        }
    }
}
