// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class VmVgpuPolicyProviderVdcScopeGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("clusterNames")]
        private InputList<string>? _clusterNames;

        /// <summary>
        /// A set of vCenter cluster names on which the provider VDC is hosted. 
        /// If none are provided, the provider attempts to find one automatically. Can be fetched using `data.vcd_resource_pool.cluster_moref` attribute.
        /// </summary>
        public InputList<string> ClusterNames
        {
            get => _clusterNames ?? (_clusterNames = new InputList<string>());
            set => _clusterNames = value;
        }

        /// <summary>
        /// The ID of the provider VDC that should be in the scope.
        /// </summary>
        [Input("providerVdcId", required: true)]
        public Input<string> ProviderVdcId { get; set; } = null!;

        /// <summary>
        /// Optional identifier for a VM group within the provider VDC scope.
        /// </summary>
        [Input("vmGroupId")]
        public Input<string>? VmGroupId { get; set; }

        public VmVgpuPolicyProviderVdcScopeGetArgs()
        {
        }
        public static new VmVgpuPolicyProviderVdcScopeGetArgs Empty => new VmVgpuPolicyProviderVdcScopeGetArgs();
    }
}
