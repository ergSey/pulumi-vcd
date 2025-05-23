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
    public sealed class ExternalNetworkV2NsxtNetwork
    {
        /// <summary>
        /// ID of NSX-T manager
        /// </summary>
        public readonly string NsxtManagerId;
        /// <summary>
        /// Name of NSX-T segment (for NSX-T segment backed external network)
        /// </summary>
        public readonly string? NsxtSegmentName;
        /// <summary>
        /// ID of NSX-T Tier-0 router (for T0 gateway backed external network)
        /// </summary>
        public readonly string? NsxtTier0RouterId;

        [OutputConstructor]
        private ExternalNetworkV2NsxtNetwork(
            string nsxtManagerId,

            string? nsxtSegmentName,

            string? nsxtTier0RouterId)
        {
            NsxtManagerId = nsxtManagerId;
            NsxtSegmentName = nsxtSegmentName;
            NsxtTier0RouterId = nsxtTier0RouterId;
        }
    }
}
