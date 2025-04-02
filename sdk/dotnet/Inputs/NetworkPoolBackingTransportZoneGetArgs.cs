// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class NetworkPoolBackingTransportZoneGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// (Computed) The ID of the backing element
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Unique name of network pool
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Type of the network pool (one of `GENEVE`, `VLAN`, `PORTGROUP_BACKED`)
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public NetworkPoolBackingTransportZoneGetArgs()
        {
        }
        public static new NetworkPoolBackingTransportZoneGetArgs Empty => new NetworkPoolBackingTransportZoneGetArgs();
    }
}
