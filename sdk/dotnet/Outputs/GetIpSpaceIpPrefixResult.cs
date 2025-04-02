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
    public sealed class GetIpSpaceIpPrefixResult
    {
        /// <summary>
        /// Floating IP quota
        /// </summary>
        public readonly string DefaultQuota;
        /// <summary>
        /// IP Prefix
        /// </summary>
        public readonly ImmutableArray<Outputs.GetIpSpaceIpPrefixPrefixResult> Prefixes;

        [OutputConstructor]
        private GetIpSpaceIpPrefixResult(
            string defaultQuota,

            ImmutableArray<Outputs.GetIpSpaceIpPrefixPrefixResult> prefixes)
        {
            DefaultQuota = defaultQuota;
            Prefixes = prefixes;
        }
    }
}
