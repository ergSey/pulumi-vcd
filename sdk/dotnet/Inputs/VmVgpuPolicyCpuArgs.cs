// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class VmVgpuPolicyCpuArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The number of cores per socket for a VM. This is a VM hardware configuration. The number of vCPUs that is defined in the VM sizing policy must be divisible by the number of cores per socket. If the number of vCPUs is not divisible by the number of cores per socket, the number of cores per socket becomes invalid.
        /// </summary>
        [Input("coresPerSocket")]
        public Input<string>? CoresPerSocket { get; set; }

        /// <summary>
        /// Defines the number of vCPUs configured for a VM. This is a VM hardware configuration. When a tenant assigns the VM sizing policy to a VM, this count becomes the configured number of vCPUs for the VM.
        /// </summary>
        [Input("count")]
        public Input<string>? Count { get; set; }

        /// <summary>
        /// Defines the CPU limit in MHz for a VM. If not defined in the VDC compute policy, CPU limit is equal to the vCPU speed multiplied by the number of vCPUs. -1 means unlimited
        /// </summary>
        [Input("limitInMhz")]
        public Input<string>? LimitInMhz { get; set; }

        /// <summary>
        /// Defines how much of the CPU resources of a VM are reserved. The allocated CPU for a VM equals the number of vCPUs times the vCPU speed in MHz. The value of the attribute ranges between 0 and one. Value of 0 CPU reservation guarantee defines no CPU reservation. Value of 1 defines 100% of CPU reserved.
        /// </summary>
        [Input("reservationGuarantee")]
        public Input<string>? ReservationGuarantee { get; set; }

        /// <summary>
        /// Defines the number of CPU shares for a VM. Shares specify the relative importance of a VM within a virtual data center. If a VM has twice as many shares of CPU as another VM, it is entitled to consume twice as much CPU when these two virtual machines are competing for resources. If not defined in the VDC compute policy, normal shares are applied to the VM.
        /// </summary>
        [Input("shares")]
        public Input<string>? Shares { get; set; }

        /// <summary>
        /// Defines the vCPU speed of a core in MHz.
        /// </summary>
        [Input("speedInMhz")]
        public Input<string>? SpeedInMhz { get; set; }

        public VmVgpuPolicyCpuArgs()
        {
        }
        public static new VmVgpuPolicyCpuArgs Empty => new VmVgpuPolicyCpuArgs();
    }
}
