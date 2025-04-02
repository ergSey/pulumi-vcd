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
    public sealed class EdgegatewayVpnPeerSubnet
    {
        public readonly string PeerSubnetGateway;
        public readonly string PeerSubnetMask;
        public readonly string PeerSubnetName;

        [OutputConstructor]
        private EdgegatewayVpnPeerSubnet(
            string peerSubnetGateway,

            string peerSubnetMask,

            string peerSubnetName)
        {
            PeerSubnetGateway = peerSubnetGateway;
            PeerSubnetMask = peerSubnetMask;
            PeerSubnetName = peerSubnetName;
        }
    }
}
