// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class NsxtEdgegatewayDnsDefaultForwarderZoneGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Unique ID of the forwarder zone.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Name of the forwarder zone.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("upstreamServers", required: true)]
        private InputList<string>? _upstreamServers;

        /// <summary>
        /// Servers to which DNS requests should be forwarded to.
        /// </summary>
        public InputList<string> UpstreamServers
        {
            get => _upstreamServers ?? (_upstreamServers = new InputList<string>());
            set => _upstreamServers = value;
        }

        public NsxtEdgegatewayDnsDefaultForwarderZoneGetArgs()
        {
        }
        public static new NsxtEdgegatewayDnsDefaultForwarderZoneGetArgs Empty => new NsxtEdgegatewayDnsDefaultForwarderZoneGetArgs();
    }
}
