// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class SolutionLandingZoneVdcStoragePolicyGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("capabilities")]
        private InputList<string>? _capabilities;

        /// <summary>
        /// Set of capabilities for Storage Policy
        /// </summary>
        public InputList<string> Capabilities
        {
            get => _capabilities ?? (_capabilities = new InputList<string>());
            set => _capabilities = value;
        }

        /// <summary>
        /// ID of Storage Policy
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// Boolean value that marks if this Storage Policy should be default
        /// </summary>
        [Input("isDefault")]
        public Input<bool>? IsDefault { get; set; }

        /// <summary>
        /// Name of Storage Policy
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public SolutionLandingZoneVdcStoragePolicyGetArgs()
        {
        }
        public static new SolutionLandingZoneVdcStoragePolicyGetArgs Empty => new SolutionLandingZoneVdcStoragePolicyGetArgs();
    }
}
