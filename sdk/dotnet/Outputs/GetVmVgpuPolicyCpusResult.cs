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
    public sealed class GetVmVgpuPolicyCpusResult
    {
        /// <summary>
        /// The number of cores per socket for a VM. This is a VM hardware configuration. The number of vCPUs that is defined in the VM sizing policy must be divisible by the number of cores per socket. If the number of vCPUs is not divisible by the number of cores per socket, the number of cores per socket becomes invalid.
        /// </summary>
        public readonly string CoresPerSocket;
        /// <summary>
        /// Defines the number of vCPUs configured for a VM. This is a VM hardware configuration. When a tenant assigns the VM sizing policy to a VM, this count becomes the configured number of vCPUs for the VM.
        /// </summary>
        public readonly string Count;
        /// <summary>
        /// Defines the CPU limit in MHz for a VM. If not defined in the VDC compute policy, CPU limit is equal to the vCPU speed multiplied by the number of vCPUs.
        /// </summary>
        public readonly string LimitInMhz;
        /// <summary>
        /// Defines how much of the CPU resources of a VM are reserved. The allocated CPU for a VM equals the number of vCPUs times the vCPU speed in MHz. The value of the attribute ranges between 0 and one. Value of 0 CPU reservation guarantee defines no CPU reservation. Value of 1 defines 100% of CPU reserved.
        /// </summary>
        public readonly string ReservationGuarantee;
        /// <summary>
        /// Defines the number of CPU shares for a VM. Shares specify the relative importance of a VM within a virtual data center. If a VM has twice as many shares of CPU as another VM, it is entitled to consume twice as much CPU when these two virtual machines are competing for resources. If not defined in the VDC compute policy, normal shares are applied to the VM.
        /// </summary>
        public readonly string Shares;
        /// <summary>
        /// Defines the vCPU speed of a core in MHz.
        /// </summary>
        public readonly string SpeedInMhz;

        [OutputConstructor]
        private GetVmVgpuPolicyCpusResult(
            string coresPerSocket,

            string count,

            string limitInMhz,

            string reservationGuarantee,

            string shares,

            string speedInMhz)
        {
            CoresPerSocket = coresPerSocket;
            Count = count;
            LimitInMhz = limitInMhz;
            ReservationGuarantee = reservationGuarantee;
            Shares = shares;
            SpeedInMhz = speedInMhz;
        }
    }
}
