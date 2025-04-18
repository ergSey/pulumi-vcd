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
    public sealed class GetNsxtNetworkDhcpBindingDhcpV6ConfigResult
    {
        /// <summary>
        /// List of DNS servers to be used by the DHCP client
        /// </summary>
        public readonly ImmutableArray<string> DnsServers;
        /// <summary>
        /// List of SNTP servers to be used by the DHCP client
        /// </summary>
        public readonly ImmutableArray<string> SntpServers;

        [OutputConstructor]
        private GetNsxtNetworkDhcpBindingDhcpV6ConfigResult(
            ImmutableArray<string> dnsServers,

            ImmutableArray<string> sntpServers)
        {
            DnsServers = dnsServers;
            SntpServers = sntpServers;
        }
    }
}
