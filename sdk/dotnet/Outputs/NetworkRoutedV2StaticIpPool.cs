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
    public sealed class NetworkRoutedV2StaticIpPool
    {
        /// <summary>
        /// End address of the IP range
        /// </summary>
        public readonly string EndAddress;
        /// <summary>
        /// Start address of the IP range
        /// </summary>
        public readonly string StartAddress;

        [OutputConstructor]
        private NetworkRoutedV2StaticIpPool(
            string endAddress,

            string startAddress)
        {
            EndAddress = endAddress;
            StartAddress = startAddress;
        }
    }
}
