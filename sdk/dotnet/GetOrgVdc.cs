// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetOrgVdc
    {
        /// <summary>
        /// Provides a VMware Cloud Director Organization VDC data source. An Organization VDC can be used to
        /// reference a VDC and use its data within other resources or data sources.
        /// 
        /// &gt; **Note:** This resource supports NSX-T and NSX-V based Org VDCs
        /// 
        /// Supported in provider *v2.5+*
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vcd = Pulumi.Vcd;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var my_org_vdc = Vcd.GetOrgVdc.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         Name = "my-vdc",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["providerVdc"] = my_org_vdc.Apply(my_org_vdc =&gt; my_org_vdc.Apply(getOrgVdcResult =&gt; getOrgVdcResult.ProviderVdcName)),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Task<GetOrgVdcResult> InvokeAsync(GetOrgVdcArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOrgVdcResult>("vcd:index/getOrgVdc:getOrgVdc", args ?? new GetOrgVdcArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a VMware Cloud Director Organization VDC data source. An Organization VDC can be used to
        /// reference a VDC and use its data within other resources or data sources.
        /// 
        /// &gt; **Note:** This resource supports NSX-T and NSX-V based Org VDCs
        /// 
        /// Supported in provider *v2.5+*
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vcd = Pulumi.Vcd;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var my_org_vdc = Vcd.GetOrgVdc.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         Name = "my-vdc",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["providerVdc"] = my_org_vdc.Apply(my_org_vdc =&gt; my_org_vdc.Apply(getOrgVdcResult =&gt; getOrgVdcResult.ProviderVdcName)),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetOrgVdcResult> Invoke(GetOrgVdcInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOrgVdcResult>("vcd:index/getOrgVdc:getOrgVdc", args ?? new GetOrgVdcInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a VMware Cloud Director Organization VDC data source. An Organization VDC can be used to
        /// reference a VDC and use its data within other resources or data sources.
        /// 
        /// &gt; **Note:** This resource supports NSX-T and NSX-V based Org VDCs
        /// 
        /// Supported in provider *v2.5+*
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vcd = Pulumi.Vcd;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var my_org_vdc = Vcd.GetOrgVdc.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         Name = "my-vdc",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["providerVdc"] = my_org_vdc.Apply(my_org_vdc =&gt; my_org_vdc.Apply(getOrgVdcResult =&gt; getOrgVdcResult.ProviderVdcName)),
        ///     };
        /// });
        /// ```
        /// </summary>
        public static Output<GetOrgVdcResult> Invoke(GetOrgVdcInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetOrgVdcResult>("vcd:index/getOrgVdc:getOrgVdc", args ?? new GetOrgVdcInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOrgVdcArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Organization VDC name
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Org name
        /// </summary>
        [Input("org")]
        public string? Org { get; set; }

        public GetOrgVdcArgs()
        {
        }
        public static new GetOrgVdcArgs Empty => new GetOrgVdcArgs();
    }

    public sealed class GetOrgVdcInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Organization VDC name
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Org name
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        public GetOrgVdcInvokeArgs()
        {
        }
        public static new GetOrgVdcInvokeArgs Empty => new GetOrgVdcInvokeArgs();
    }


    [OutputType]
    public sealed class GetOrgVdcResult
    {
        public readonly string AllocationModel;
        public readonly bool AllowOverCommit;
        public readonly ImmutableArray<Outputs.GetOrgVdcComputeCapacityResult> ComputeCapacities;
        public readonly double CpuGuaranteed;
        public readonly int CpuSpeed;
        public readonly string DefaultComputePolicyId;
        public readonly string DefaultVmSizingPolicyId;
        public readonly string Description;
        /// <summary>
        /// (*v3.8+*, *VCD 10.3+*) An ID of NSX-T Edge Cluster which should provide vApp
        /// Networking Services or DHCP for Isolated Networks. This value **might be unavailable** in data
        /// source if user has insufficient rights.
        /// </summary>
        public readonly string EdgeClusterId;
        public readonly bool Elasticity;
        public readonly bool EnableFastProvisioning;
        public readonly bool EnableNsxvDistributedFirewall;
        public readonly bool EnableThinProvisioning;
        public readonly bool EnableVmDiscovery;
        public readonly bool Enabled;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool IncludeVmMemoryOverhead;
        public readonly double MemoryGuaranteed;
        public readonly ImmutableDictionary<string, string> Metadata;
        public readonly ImmutableArray<Outputs.GetOrgVdcMetadataEntryResult> MetadataEntries;
        public readonly string Name;
        public readonly string NetworkPoolName;
        public readonly int NetworkQuota;
        public readonly int NicQuota;
        public readonly string? Org;
        public readonly string ProviderVdcName;
        public readonly ImmutableArray<Outputs.GetOrgVdcStorageProfileResult> StorageProfiles;
        public readonly ImmutableArray<string> VmPlacementPolicyIds;
        public readonly int VmQuota;
        public readonly ImmutableArray<string> VmSizingPolicyIds;
        public readonly ImmutableArray<string> VmVgpuPolicyIds;

        [OutputConstructor]
        private GetOrgVdcResult(
            string allocationModel,

            bool allowOverCommit,

            ImmutableArray<Outputs.GetOrgVdcComputeCapacityResult> computeCapacities,

            double cpuGuaranteed,

            int cpuSpeed,

            string defaultComputePolicyId,

            string defaultVmSizingPolicyId,

            string description,

            string edgeClusterId,

            bool elasticity,

            bool enableFastProvisioning,

            bool enableNsxvDistributedFirewall,

            bool enableThinProvisioning,

            bool enableVmDiscovery,

            bool enabled,

            string id,

            bool includeVmMemoryOverhead,

            double memoryGuaranteed,

            ImmutableDictionary<string, string> metadata,

            ImmutableArray<Outputs.GetOrgVdcMetadataEntryResult> metadataEntries,

            string name,

            string networkPoolName,

            int networkQuota,

            int nicQuota,

            string? org,

            string providerVdcName,

            ImmutableArray<Outputs.GetOrgVdcStorageProfileResult> storageProfiles,

            ImmutableArray<string> vmPlacementPolicyIds,

            int vmQuota,

            ImmutableArray<string> vmSizingPolicyIds,

            ImmutableArray<string> vmVgpuPolicyIds)
        {
            AllocationModel = allocationModel;
            AllowOverCommit = allowOverCommit;
            ComputeCapacities = computeCapacities;
            CpuGuaranteed = cpuGuaranteed;
            CpuSpeed = cpuSpeed;
            DefaultComputePolicyId = defaultComputePolicyId;
            DefaultVmSizingPolicyId = defaultVmSizingPolicyId;
            Description = description;
            EdgeClusterId = edgeClusterId;
            Elasticity = elasticity;
            EnableFastProvisioning = enableFastProvisioning;
            EnableNsxvDistributedFirewall = enableNsxvDistributedFirewall;
            EnableThinProvisioning = enableThinProvisioning;
            EnableVmDiscovery = enableVmDiscovery;
            Enabled = enabled;
            Id = id;
            IncludeVmMemoryOverhead = includeVmMemoryOverhead;
            MemoryGuaranteed = memoryGuaranteed;
            Metadata = metadata;
            MetadataEntries = metadataEntries;
            Name = name;
            NetworkPoolName = networkPoolName;
            NetworkQuota = networkQuota;
            NicQuota = nicQuota;
            Org = org;
            ProviderVdcName = providerVdcName;
            StorageProfiles = storageProfiles;
            VmPlacementPolicyIds = vmPlacementPolicyIds;
            VmQuota = vmQuota;
            VmSizingPolicyIds = vmSizingPolicyIds;
            VmVgpuPolicyIds = vmVgpuPolicyIds;
        }
    }
}
