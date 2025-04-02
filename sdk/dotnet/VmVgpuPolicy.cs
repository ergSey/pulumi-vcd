// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    /// <summary>
    /// Experimental in provider *3.11*.
    /// 
    /// &gt; **Note:** This resource requires system administrator privileges.
    /// 
    /// Provides a resource to manage vGPU policies for virtual machines in VMware Cloud Director.
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
    ///     var exampleOrg = Vcd.GetOrg.Invoke(new()
    ///     {
    ///         Name = "test_org",
    ///     });
    /// 
    ///     var exampleVgpuProfile = Vcd.GetVgpuProfile.Invoke(new()
    ///     {
    ///         Name = "grid_a100-10c",
    ///     });
    /// 
    ///     var exampleProviderVdc = Vcd.GetProviderVdc.Invoke(new()
    ///     {
    ///         Name = "example_provider_vdc",
    ///     });
    /// 
    ///     var vmGroupExample = Vcd.GetVmGroup.Invoke(new()
    ///     {
    ///         Name = "vm-group-1",
    ///     });
    /// 
    ///     var exampleVgpuPolicy = new Vcd.VmVgpuPolicy("example_vgpu_policy", new()
    ///     {
    ///         Name = "example-vgpu-policy",
    ///         Description = "An example vGPU policy configuration",
    ///         VgpuProfile = new Vcd.Inputs.VmVgpuPolicyVgpuProfileArgs
    ///         {
    ///             Id = exampleVgpuProfile.Apply(getVgpuProfileResult =&gt; getVgpuProfileResult.Id),
    ///             Count = 1,
    ///         },
    ///         Cpu = new Vcd.Inputs.VmVgpuPolicyCpuArgs
    ///         {
    ///             Shares = "886",
    ///             LimitInMhz = "2400",
    ///             Count = "9",
    ///             SpeedInMhz = "2500",
    ///             CoresPerSocket = "3",
    ///             ReservationGuarantee = "0.55",
    ///         },
    ///         Memory = new Vcd.Inputs.VmVgpuPolicyMemoryArgs
    ///         {
    ///             Shares = "1580",
    ///             SizeInMb = "3200",
    ///             LimitInMb = "2800",
    ///         },
    ///         ProviderVdcScopes = new[]
    ///         {
    ///             new Vcd.Inputs.VmVgpuPolicyProviderVdcScopeArgs
    ///             {
    ///                 ProviderVdcId = exampleProviderVdc.Apply(getProviderVdcResult =&gt; getProviderVdcResult.Id),
    ///                 ClusterNames = new[]
    ///                 {
    ///                     "cluster1",
    ///                 },
    ///                 VmGroupId = vmGroupExample.Apply(getVmGroupResult =&gt; getVmGroupResult.Id),
    ///             },
    ///         },
    ///     });
    /// 
    ///     var exampleOrgVdc = new Vcd.OrgVdc("example_org_vdc", new()
    ///     {
    ///         Org = exampleOrg.Apply(getOrgResult =&gt; getOrgResult.Name),
    ///         Name = "test-org-vdc",
    ///         ProviderVdcName = exampleProviderVdc.Apply(getProviderVdcResult =&gt; getProviderVdcResult.Name),
    ///         AllocationModel = "Flex",
    ///         DeleteForce = true,
    ///         ComputeCapacity = new Vcd.Inputs.OrgVdcComputeCapacityArgs
    ///         {
    ///             Cpu = new Vcd.Inputs.OrgVdcComputeCapacityCpuArgs
    ///             {
    ///                 Allocated = 2048,
    ///             },
    ///             Memory = new Vcd.Inputs.OrgVdcComputeCapacityMemoryArgs
    ///             {
    ///                 Allocated = 2048,
    ///             },
    ///         },
    ///         StorageProfiles = new[]
    ///         {
    ///             new Vcd.Inputs.OrgVdcStorageProfileArgs
    ///             {
    ///                 Name = "*",
    ///                 Limit = 10240,
    ///                 Default = true,
    ///             },
    ///         },
    ///         Elasticity = true,
    ///         IncludeVmMemoryOverhead = true,
    ///         MemoryGuaranteed = 1,
    ///         DefaultComputePolicyId = exampleVgpuPolicy.Id,
    ///         VmVgpuPolicyIds = new[]
    ///         {
    ///             exampleVgpuPolicy.Id,
    ///         },
    ///     });
    /// 
    ///     var testVm = new Vcd.Vm("test_vm", new()
    ///     {
    ///         Org = exampleOrg.Apply(getOrgResult =&gt; getOrgResult.Name),
    ///         Vdc = exampleOrgVdc.Name,
    ///         Name = "terraform-provider-vm",
    ///         ComputerName = "emptyVM",
    ///         Memory = 2048,
    ///         Cpus = 2,
    ///         CpuCores = 1,
    ///         PowerOn = false,
    ///         OsType = "sles11_64Guest",
    ///         HardwareVersion = "vmx-19",
    ///         PlacementPolicyId = exampleVgpuPolicy.Id,
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ### Without A Sizing Policy)
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using Pulumi;
    /// using Vcd = Pulumi.Vcd;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var exampleVgpuPolicyWithoutSizing = new Vcd.VmVgpuPolicy("example_vgpu_policy_without_sizing", new()
    ///     {
    ///         Name = "example-vgpu-policy-without-sizing",
    ///         Description = "An example vGPU policy configuration",
    ///         VgpuProfile = new Vcd.Inputs.VmVgpuPolicyVgpuProfileArgs
    ///         {
    ///             Id = exampleVgpuProfile.Id,
    ///             Count = 1,
    ///         },
    ///         ProviderVdcScopes = new[]
    ///         {
    ///             new Vcd.Inputs.VmVgpuPolicyProviderVdcScopeArgs
    ///             {
    ///                 ProviderVdcId = exampleProviderVdc.Id,
    ///                 ClusterNames = new[]
    ///                 {
    ///                     "cluster1",
    ///                 },
    ///                 VmGroupId = vmGroupExample.Id,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [VcdResourceType("vcd:index/vmVgpuPolicy:VmVgpuPolicy")]
    public partial class VmVgpuPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Configuration options for CPU resources. If this is set, 
        /// a VM created with this policy can't specify a custom sizing policy. See [cpu] for more details.
        /// </summary>
        [Output("cpu")]
        public Output<Outputs.VmVgpuPolicyCpu?> Cpu { get; private set; } = null!;

        /// <summary>
        /// A brief description of the vGPU policy.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Memory resource configuration settings. If this is set, 
        /// a VM created with this policy can't specify a custom sizing policy. See [memory] for more details.
        /// </summary>
        [Output("memory")]
        public Output<Outputs.VmVgpuPolicyMemory?> Memory { get; private set; } = null!;

        /// <summary>
        /// The unique name assigned to the vGPU policy for a virtual machine.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Defines the scope of the policy within 
        /// provider virtual data centers. If not provided, applies to all the current ant future PVDCs.
        /// See `provider_vdc_scope` for more details.
        /// </summary>
        [Output("providerVdcScopes")]
        public Output<ImmutableArray<Outputs.VmVgpuPolicyProviderVdcScope>> ProviderVdcScopes { get; private set; } = null!;

        /// <summary>
        /// Defines the vGPU profile ID and count.
        /// </summary>
        [Output("vgpuProfile")]
        public Output<Outputs.VmVgpuPolicyVgpuProfile> VgpuProfile { get; private set; } = null!;


        /// <summary>
        /// Create a VmVgpuPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VmVgpuPolicy(string name, VmVgpuPolicyArgs args, CustomResourceOptions? options = null)
            : base("vcd:index/vmVgpuPolicy:VmVgpuPolicy", name, args ?? new VmVgpuPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VmVgpuPolicy(string name, Input<string> id, VmVgpuPolicyState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/vmVgpuPolicy:VmVgpuPolicy", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/vmware/terraform-provider-vcd/releases/download/v3.14.1/",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing VmVgpuPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VmVgpuPolicy Get(string name, Input<string> id, VmVgpuPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new VmVgpuPolicy(name, id, state, options);
        }
    }

    public sealed class VmVgpuPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration options for CPU resources. If this is set, 
        /// a VM created with this policy can't specify a custom sizing policy. See [cpu] for more details.
        /// </summary>
        [Input("cpu")]
        public Input<Inputs.VmVgpuPolicyCpuArgs>? Cpu { get; set; }

        /// <summary>
        /// A brief description of the vGPU policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Memory resource configuration settings. If this is set, 
        /// a VM created with this policy can't specify a custom sizing policy. See [memory] for more details.
        /// </summary>
        [Input("memory")]
        public Input<Inputs.VmVgpuPolicyMemoryArgs>? Memory { get; set; }

        /// <summary>
        /// The unique name assigned to the vGPU policy for a virtual machine.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("providerVdcScopes")]
        private InputList<Inputs.VmVgpuPolicyProviderVdcScopeArgs>? _providerVdcScopes;

        /// <summary>
        /// Defines the scope of the policy within 
        /// provider virtual data centers. If not provided, applies to all the current ant future PVDCs.
        /// See `provider_vdc_scope` for more details.
        /// </summary>
        public InputList<Inputs.VmVgpuPolicyProviderVdcScopeArgs> ProviderVdcScopes
        {
            get => _providerVdcScopes ?? (_providerVdcScopes = new InputList<Inputs.VmVgpuPolicyProviderVdcScopeArgs>());
            set => _providerVdcScopes = value;
        }

        /// <summary>
        /// Defines the vGPU profile ID and count.
        /// </summary>
        [Input("vgpuProfile", required: true)]
        public Input<Inputs.VmVgpuPolicyVgpuProfileArgs> VgpuProfile { get; set; } = null!;

        public VmVgpuPolicyArgs()
        {
        }
        public static new VmVgpuPolicyArgs Empty => new VmVgpuPolicyArgs();
    }

    public sealed class VmVgpuPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Configuration options for CPU resources. If this is set, 
        /// a VM created with this policy can't specify a custom sizing policy. See [cpu] for more details.
        /// </summary>
        [Input("cpu")]
        public Input<Inputs.VmVgpuPolicyCpuGetArgs>? Cpu { get; set; }

        /// <summary>
        /// A brief description of the vGPU policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Memory resource configuration settings. If this is set, 
        /// a VM created with this policy can't specify a custom sizing policy. See [memory] for more details.
        /// </summary>
        [Input("memory")]
        public Input<Inputs.VmVgpuPolicyMemoryGetArgs>? Memory { get; set; }

        /// <summary>
        /// The unique name assigned to the vGPU policy for a virtual machine.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("providerVdcScopes")]
        private InputList<Inputs.VmVgpuPolicyProviderVdcScopeGetArgs>? _providerVdcScopes;

        /// <summary>
        /// Defines the scope of the policy within 
        /// provider virtual data centers. If not provided, applies to all the current ant future PVDCs.
        /// See `provider_vdc_scope` for more details.
        /// </summary>
        public InputList<Inputs.VmVgpuPolicyProviderVdcScopeGetArgs> ProviderVdcScopes
        {
            get => _providerVdcScopes ?? (_providerVdcScopes = new InputList<Inputs.VmVgpuPolicyProviderVdcScopeGetArgs>());
            set => _providerVdcScopes = value;
        }

        /// <summary>
        /// Defines the vGPU profile ID and count.
        /// </summary>
        [Input("vgpuProfile")]
        public Input<Inputs.VmVgpuPolicyVgpuProfileGetArgs>? VgpuProfile { get; set; }

        public VmVgpuPolicyState()
        {
        }
        public static new VmVgpuPolicyState Empty => new VmVgpuPolicyState();
    }
}
