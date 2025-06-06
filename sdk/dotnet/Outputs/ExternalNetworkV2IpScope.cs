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
    public sealed class ExternalNetworkV2IpScope
    {
        /// <summary>
        /// Primary DNS server
        /// </summary>
        public readonly string? Dns1;
        /// <summary>
        /// Secondary DNS server
        /// </summary>
        public readonly string? Dns2;
        /// <summary>
        /// DNS suffix
        /// </summary>
        public readonly string? DnsSuffix;
        /// <summary>
        /// If subnet is enabled
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// Gateway of the network
        /// </summary>
        public readonly string Gateway;
        /// <summary>
        /// Network mask
        /// </summary>
        public readonly int PrefixLength;
        /// <summary>
        /// IP ranges used for static pool allocation in the network
        /// </summary>
        public readonly ImmutableArray<Outputs.ExternalNetworkV2IpScopeStaticIpPool> StaticIpPools;

        [OutputConstructor]
        private ExternalNetworkV2IpScope(
            string? dns1,

            string? dns2,

            string? dnsSuffix,

            bool? enabled,

            string gateway,

            int prefixLength,

            ImmutableArray<Outputs.ExternalNetworkV2IpScopeStaticIpPool> staticIpPools)
        {
            Dns1 = dns1;
            Dns2 = dns2;
            DnsSuffix = dnsSuffix;
            Enabled = enabled;
            Gateway = gateway;
            PrefixLength = prefixLength;
            StaticIpPools = staticIpPools;
        }
    }
}
