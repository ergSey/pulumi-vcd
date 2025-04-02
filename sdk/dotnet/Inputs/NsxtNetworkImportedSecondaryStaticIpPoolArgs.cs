// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class NsxtNetworkImportedSecondaryStaticIpPoolArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// End address of the IP range
        /// </summary>
        [Input("endAddress", required: true)]
        public Input<string> EndAddress { get; set; } = null!;

        /// <summary>
        /// Start address of the IP range
        /// </summary>
        [Input("startAddress", required: true)]
        public Input<string> StartAddress { get; set; } = null!;

        public NsxtNetworkImportedSecondaryStaticIpPoolArgs()
        {
        }
        public static new NsxtNetworkImportedSecondaryStaticIpPoolArgs Empty => new NsxtNetworkImportedSecondaryStaticIpPoolArgs();
    }
}
