// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class EdgegatewayVpnPeerSubnetArgs : global::Pulumi.ResourceArgs
    {
        [Input("peerSubnetGateway", required: true)]
        public Input<string> PeerSubnetGateway { get; set; } = null!;

        [Input("peerSubnetMask", required: true)]
        public Input<string> PeerSubnetMask { get; set; } = null!;

        [Input("peerSubnetName", required: true)]
        public Input<string> PeerSubnetName { get; set; } = null!;

        public EdgegatewayVpnPeerSubnetArgs()
        {
        }
        public static new EdgegatewayVpnPeerSubnetArgs Empty => new EdgegatewayVpnPeerSubnetArgs();
    }
}
