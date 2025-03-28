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
    public sealed class GetCseKubernetesClusterControlPlaneResult
    {
        /// <summary>
        /// Disk size, in Gibibytes (Gi), of the control plane nodes
        /// </summary>
        public readonly int DiskSizeGi;
        /// <summary>
        /// IP of the control plane
        /// </summary>
        public readonly string Ip;
        /// <summary>
        /// The number of nodes that the control plane has
        /// </summary>
        public readonly int MachineCount;
        /// <summary>
        /// VM Placement policy of the control plane nodes
        /// </summary>
        public readonly string PlacementPolicyId;
        /// <summary>
        /// VM Sizing policy of the control plane nodes
        /// </summary>
        public readonly string SizingPolicyId;
        /// <summary>
        /// Storage profile of the control plane nodes
        /// </summary>
        public readonly string StorageProfileId;

        [OutputConstructor]
        private GetCseKubernetesClusterControlPlaneResult(
            int diskSizeGi,

            string ip,

            int machineCount,

            string placementPolicyId,

            string sizingPolicyId,

            string storageProfileId)
        {
            DiskSizeGi = diskSizeGi;
            Ip = ip;
            MachineCount = machineCount;
            PlacementPolicyId = placementPolicyId;
            SizingPolicyId = sizingPolicyId;
            StorageProfileId = storageProfileId;
        }
    }
}
