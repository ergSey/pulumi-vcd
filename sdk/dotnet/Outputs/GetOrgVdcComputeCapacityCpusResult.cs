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
    public sealed class GetOrgVdcComputeCapacityCpusResult
    {
        /// <summary>
        /// Capacity that is committed to be available. Value in MB or MHz. Used with AllocationPool (Allocation pool) and ReservationPool (Reservation pool).
        /// </summary>
        public readonly int Allocated;
        /// <summary>
        /// Capacity limit relative to the value specified for Allocation. It must not be less than that value. If it is greater than that value, it implies over provisioning. A value of 0 specifies unlimited units. Value in MB or MHz. Used with AllocationVApp (Pay as you go).
        /// </summary>
        public readonly int Limit;
        public readonly int Reserved;
        public readonly int Used;

        [OutputConstructor]
        private GetOrgVdcComputeCapacityCpusResult(
            int allocated,

            int limit,

            int reserved,

            int used)
        {
            Allocated = allocated;
            Limit = limit;
            Reserved = reserved;
            Used = used;
        }
    }
}
