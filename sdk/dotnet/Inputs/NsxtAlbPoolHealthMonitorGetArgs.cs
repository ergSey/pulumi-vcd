// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class NsxtAlbPoolHealthMonitorGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A name for ALB Pool
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("systemDefined")]
        public Input<bool>? SystemDefined { get; set; }

        /// <summary>
        /// Type of health monitor. One of `HTTP`, `HTTPS`, `TCP`, `UDP`, `PING`
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public NsxtAlbPoolHealthMonitorGetArgs()
        {
        }
        public static new NsxtAlbPoolHealthMonitorGetArgs Empty => new NsxtAlbPoolHealthMonitorGetArgs();
    }
}
