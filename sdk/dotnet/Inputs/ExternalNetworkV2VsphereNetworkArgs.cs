// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class ExternalNetworkV2VsphereNetworkArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the port group
        /// </summary>
        [Input("portgroupId", required: true)]
        public Input<string> PortgroupId { get; set; } = null!;

        /// <summary>
        /// The vCenter server name
        /// </summary>
        [Input("vcenterId", required: true)]
        public Input<string> VcenterId { get; set; } = null!;

        public ExternalNetworkV2VsphereNetworkArgs()
        {
        }
        public static new ExternalNetworkV2VsphereNetworkArgs Empty => new ExternalNetworkV2VsphereNetworkArgs();
    }
}
