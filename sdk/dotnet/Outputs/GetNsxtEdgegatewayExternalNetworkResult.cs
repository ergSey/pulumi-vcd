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
    public sealed class GetNsxtEdgegatewayExternalNetworkResult
    {
        /// <summary>
        /// Number of allocated IPs
        /// </summary>
        public readonly int AllocatedIpCount;
        /// <summary>
        /// NSX-T Segment backed External Network ID
        /// </summary>
        public readonly string ExternalNetworkId;
        /// <summary>
        /// Gateway IP Address
        /// </summary>
        public readonly string Gateway;
        /// <summary>
        /// Prefix length for a subnet (e.g. 24)
        /// </summary>
        public readonly int PrefixLength;
        /// <summary>
        /// Primary IP address for the Edge Gateway
        /// </summary>
        public readonly string PrimaryIp;

        [OutputConstructor]
        private GetNsxtEdgegatewayExternalNetworkResult(
            int allocatedIpCount,

            string externalNetworkId,

            string gateway,

            int prefixLength,

            string primaryIp)
        {
            AllocatedIpCount = allocatedIpCount;
            ExternalNetworkId = externalNetworkId;
            Gateway = gateway;
            PrefixLength = prefixLength;
            PrimaryIp = primaryIp;
        }
    }
}
