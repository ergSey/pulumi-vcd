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
    /// Provides a VMware Cloud Director VM Placement Policy resource. This can be used to create, modify, and delete VM Placement Policy.
    /// 
    /// Supported in provider *v3.8+* and requires VCD 10.2+
    /// 
    /// &gt; **Note:** This resource requires system administrator privileges.
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
    ///     var pvdc = Vcd.GetProviderVdc.Invoke(new()
    ///     {
    ///         Name = "my-pvdc",
    ///     });
    /// 
    ///     var vm_group = Vcd.GetVmGroup.Invoke(new()
    ///     {
    ///         Name = "vmware-vm-group",
    ///         ProviderVdcId = pvdc.Apply(getProviderVdcResult =&gt; getProviderVdcResult.Id),
    ///     });
    /// 
    ///     var test_placement_pol = new Vcd.VmPlacementPolicy("test-placement-pol", new()
    ///     {
    ///         Name = "my-placement-pol",
    ///         Description = "My awesome VM Placement Policy",
    ///         ProviderVdcId = pvdc.Apply(getProviderVdcResult =&gt; getProviderVdcResult.Id),
    ///         VmGroupIds = new[]
    ///         {
    ///             vm_group.Apply(vm_group =&gt; vm_group.Apply(getVmGroupResult =&gt; getVmGroupResult.Id)),
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// </summary>
    [VcdResourceType("vcd:index/vmPlacementPolicy:VmPlacementPolicy")]
    public partial class VmPlacementPolicy : global::Pulumi.CustomResource
    {
        /// <summary>
        /// description of VM Placement Policy.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// IDs of one or more Logical VM Groups to define this VM Placement Policy. There is an AND relationship among all the
        /// entries set in this attribute
        /// </summary>
        [Output("logicalVmGroupIds")]
        public Output<ImmutableArray<string>> LogicalVmGroupIds { get; private set; } = null!;

        /// <summary>
        /// The name of VM Placement Policy.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the Provider VDC to which this VM Placement Policy belongs.
        /// </summary>
        [Output("providerVdcId")]
        public Output<string> ProviderVdcId { get; private set; } = null!;

        /// <summary>
        /// IDs of the collection of VMs with similar host requirements. **Note:** Either `vm_group_ids` or `logical_vm_group_ids` must be set.
        /// </summary>
        [Output("vmGroupIds")]
        public Output<ImmutableArray<string>> VmGroupIds { get; private set; } = null!;


        /// <summary>
        /// Create a VmPlacementPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VmPlacementPolicy(string name, VmPlacementPolicyArgs args, CustomResourceOptions? options = null)
            : base("vcd:index/vmPlacementPolicy:VmPlacementPolicy", name, args ?? new VmPlacementPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VmPlacementPolicy(string name, Input<string> id, VmPlacementPolicyState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/vmPlacementPolicy:VmPlacementPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VmPlacementPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VmPlacementPolicy Get(string name, Input<string> id, VmPlacementPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new VmPlacementPolicy(name, id, state, options);
        }
    }

    public sealed class VmPlacementPolicyArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// description of VM Placement Policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("logicalVmGroupIds")]
        private InputList<string>? _logicalVmGroupIds;

        /// <summary>
        /// IDs of one or more Logical VM Groups to define this VM Placement Policy. There is an AND relationship among all the
        /// entries set in this attribute
        /// </summary>
        public InputList<string> LogicalVmGroupIds
        {
            get => _logicalVmGroupIds ?? (_logicalVmGroupIds = new InputList<string>());
            set => _logicalVmGroupIds = value;
        }

        /// <summary>
        /// The name of VM Placement Policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the Provider VDC to which this VM Placement Policy belongs.
        /// </summary>
        [Input("providerVdcId", required: true)]
        public Input<string> ProviderVdcId { get; set; } = null!;

        [Input("vmGroupIds")]
        private InputList<string>? _vmGroupIds;

        /// <summary>
        /// IDs of the collection of VMs with similar host requirements. **Note:** Either `vm_group_ids` or `logical_vm_group_ids` must be set.
        /// </summary>
        public InputList<string> VmGroupIds
        {
            get => _vmGroupIds ?? (_vmGroupIds = new InputList<string>());
            set => _vmGroupIds = value;
        }

        public VmPlacementPolicyArgs()
        {
        }
        public static new VmPlacementPolicyArgs Empty => new VmPlacementPolicyArgs();
    }

    public sealed class VmPlacementPolicyState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// description of VM Placement Policy.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("logicalVmGroupIds")]
        private InputList<string>? _logicalVmGroupIds;

        /// <summary>
        /// IDs of one or more Logical VM Groups to define this VM Placement Policy. There is an AND relationship among all the
        /// entries set in this attribute
        /// </summary>
        public InputList<string> LogicalVmGroupIds
        {
            get => _logicalVmGroupIds ?? (_logicalVmGroupIds = new InputList<string>());
            set => _logicalVmGroupIds = value;
        }

        /// <summary>
        /// The name of VM Placement Policy.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the Provider VDC to which this VM Placement Policy belongs.
        /// </summary>
        [Input("providerVdcId")]
        public Input<string>? ProviderVdcId { get; set; }

        [Input("vmGroupIds")]
        private InputList<string>? _vmGroupIds;

        /// <summary>
        /// IDs of the collection of VMs with similar host requirements. **Note:** Either `vm_group_ids` or `logical_vm_group_ids` must be set.
        /// </summary>
        public InputList<string> VmGroupIds
        {
            get => _vmGroupIds ?? (_vmGroupIds = new InputList<string>());
            set => _vmGroupIds = value;
        }

        public VmPlacementPolicyState()
        {
        }
        public static new VmPlacementPolicyState Empty => new VmPlacementPolicyState();
    }
}
