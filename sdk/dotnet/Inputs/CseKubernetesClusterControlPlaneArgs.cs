// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class CseKubernetesClusterControlPlaneArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Disk size, in **Gibibytes (Gi)**, for the control plane VMs. Must be at least `20`. Defaults to `20`
        /// </summary>
        [Input("diskSizeGi")]
        public Input<int>? DiskSizeGi { get; set; }

        /// <summary>
        /// IP for the control plane. It will be automatically assigned during cluster creation if left empty
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// The number of nodes that the control plane has. Must be an odd number and higher than `0`. Defaults to `3`
        /// </summary>
        [Input("machineCount")]
        public Input<int>? MachineCount { get; set; }

        /// <summary>
        /// VM Placement policy for the control plane VMs
        /// </summary>
        [Input("placementPolicyId")]
        public Input<string>? PlacementPolicyId { get; set; }

        /// <summary>
        /// VM Sizing policy for the control plane VMs. Must be one of the ones made available during CSE installation
        /// </summary>
        [Input("sizingPolicyId")]
        public Input<string>? SizingPolicyId { get; set; }

        /// <summary>
        /// Storage profile for the control plane VMs
        /// </summary>
        [Input("storageProfileId")]
        public Input<string>? StorageProfileId { get; set; }

        public CseKubernetesClusterControlPlaneArgs()
        {
        }
        public static new CseKubernetesClusterControlPlaneArgs Empty => new CseKubernetesClusterControlPlaneArgs();
    }
}
