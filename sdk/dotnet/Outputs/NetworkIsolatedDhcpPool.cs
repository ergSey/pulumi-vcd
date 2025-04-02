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
    public sealed class NetworkIsolatedDhcpPool
    {
        /// <summary>
        /// The default DHCP lease time to use
        /// </summary>
        public readonly int? DefaultLeaseTime;
        /// <summary>
        /// The final address in the IP Range
        /// </summary>
        public readonly string EndAddress;
        /// <summary>
        /// The maximum DHCP lease time to use
        /// </summary>
        public readonly int? MaxLeaseTime;
        /// <summary>
        /// The first address in the IP Range
        /// </summary>
        public readonly string StartAddress;

        [OutputConstructor]
        private NetworkIsolatedDhcpPool(
            int? defaultLeaseTime,

            string endAddress,

            int? maxLeaseTime,

            string startAddress)
        {
            DefaultLeaseTime = defaultLeaseTime;
            EndAddress = endAddress;
            MaxLeaseTime = maxLeaseTime;
            StartAddress = startAddress;
        }
    }
}
