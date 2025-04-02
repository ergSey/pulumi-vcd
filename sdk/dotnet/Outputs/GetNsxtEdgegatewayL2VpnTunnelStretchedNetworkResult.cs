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
    public sealed class GetNsxtEdgegatewayL2VpnTunnelStretchedNetworkResult
    {
        /// <summary>
        /// ID of the Org VDC network
        /// </summary>
        public readonly string NetworkId;
        /// <summary>
        /// Tunnel ID of the network for the tunnel. Read-only for `SERVER` sessions.
        /// </summary>
        public readonly int TunnelId;

        [OutputConstructor]
        private GetNsxtEdgegatewayL2VpnTunnelStretchedNetworkResult(
            string networkId,

            int tunnelId)
        {
            NetworkId = networkId;
            TunnelId = tunnelId;
        }
    }
}
