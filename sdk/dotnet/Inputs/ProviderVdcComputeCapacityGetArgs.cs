// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class ProviderVdcComputeCapacityGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("cpus")]
        private InputList<Inputs.ProviderVdcComputeCapacityCpusGetArgs>? _cpus;

        /// <summary>
        /// Single-element list with an indicator of CPU capacity available in the Provider VDC
        /// </summary>
        public InputList<Inputs.ProviderVdcComputeCapacityCpusGetArgs> Cpus
        {
            get => _cpus ?? (_cpus = new InputList<Inputs.ProviderVdcComputeCapacityCpusGetArgs>());
            set => _cpus = value;
        }

        /// <summary>
        /// True if compute capacity can grow or shrink based on demand
        /// </summary>
        [Input("isElastic")]
        public Input<bool>? IsElastic { get; set; }

        /// <summary>
        /// True if compute capacity is highly available
        /// </summary>
        [Input("isHa")]
        public Input<bool>? IsHa { get; set; }

        [Input("memories")]
        private InputList<Inputs.ProviderVdcComputeCapacityMemoryGetArgs>? _memories;

        /// <summary>
        /// Single-element list with an indicator of Memory capacity available in the Provider VDC
        /// </summary>
        public InputList<Inputs.ProviderVdcComputeCapacityMemoryGetArgs> Memories
        {
            get => _memories ?? (_memories = new InputList<Inputs.ProviderVdcComputeCapacityMemoryGetArgs>());
            set => _memories = value;
        }

        public ProviderVdcComputeCapacityGetArgs()
        {
        }
        public static new ProviderVdcComputeCapacityGetArgs Empty => new ProviderVdcComputeCapacityGetArgs();
    }
}
