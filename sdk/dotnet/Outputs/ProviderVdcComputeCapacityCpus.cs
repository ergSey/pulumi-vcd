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
    public sealed class ProviderVdcComputeCapacityCpus
    {
        /// <summary>
        /// Allocated CPU for this Provider VDC
        /// </summary>
        public readonly int? Allocation;
        /// <summary>
        /// CPU overhead for this Provider VDC
        /// </summary>
        public readonly int? Overhead;
        /// <summary>
        /// Reserved CPU for this Provider VDC
        /// </summary>
        public readonly int? Reserved;
        /// <summary>
        /// Total CPU for this Provider VDC
        /// </summary>
        public readonly int? Total;
        /// <summary>
        /// Units for the CPU of this Provider VDC
        /// </summary>
        public readonly string? Units;
        /// <summary>
        /// Used CPU in this Provider VDC
        /// </summary>
        public readonly int? Used;

        [OutputConstructor]
        private ProviderVdcComputeCapacityCpus(
            int? allocation,

            int? overhead,

            int? reserved,

            int? total,

            string? units,

            int? used)
        {
            Allocation = allocation;
            Overhead = overhead;
            Reserved = reserved;
            Total = total;
            Units = units;
            Used = used;
        }
    }
}
