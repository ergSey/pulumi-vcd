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
    public sealed class NsxtSecurityGroupMemberVm
    {
        /// <summary>
        /// Parent vApp name (if exists) for member VM
        /// </summary>
        public readonly string? VappId;
        /// <summary>
        /// Parent vApp ID (if exists) for member VM
        /// </summary>
        public readonly string? VappName;
        /// <summary>
        /// Member VM ID
        /// </summary>
        public readonly string? VmId;
        /// <summary>
        /// Member VM Name
        /// </summary>
        public readonly string? VmName;

        [OutputConstructor]
        private NsxtSecurityGroupMemberVm(
            string? vappId,

            string? vappName,

            string? vmId,

            string? vmName)
        {
            VappId = vappId;
            VappName = vappName;
            VmId = vmId;
            VmName = vmName;
        }
    }
}
