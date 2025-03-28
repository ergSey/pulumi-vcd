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
    public sealed class VmVgpuPolicyVgpuProfile
    {
        /// <summary>
        /// Specifies the number of vGPU profiles. Must be at least 1.
        /// </summary>
        public readonly int Count;
        /// <summary>
        /// The identifier of the vGPU profile.
        /// </summary>
        public readonly string Id;

        [OutputConstructor]
        private VmVgpuPolicyVgpuProfile(
            int count,

            string id)
        {
            Count = count;
            Id = id;
        }
    }
}
