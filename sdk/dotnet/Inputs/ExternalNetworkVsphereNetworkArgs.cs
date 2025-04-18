// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class ExternalNetworkVsphereNetworkArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique name for the network
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The vSphere port group type. One of: DV_PORTGROUP (distributed virtual port group), NETWORK
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// The vCenter server name
        /// </summary>
        [Input("vcenter", required: true)]
        public Input<string> Vcenter { get; set; } = null!;

        public ExternalNetworkVsphereNetworkArgs()
        {
        }
        public static new ExternalNetworkVsphereNetworkArgs Empty => new ExternalNetworkVsphereNetworkArgs();
    }
}
