// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/nsxtDynamicSecurityGroup:NsxtDynamicSecurityGroup")]
    public partial class NsxtDynamicSecurityGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Up to 3 criteria for matching VMs. List of criteria is matched with boolean
        /// `OR` operation and matching any of defined criteria will include objects. Each `criteria` can
        /// contains up to 4 `rule` definitions.
        /// </summary>
        [Output("criterias")]
        public Output<ImmutableArray<Outputs.NsxtDynamicSecurityGroupCriteria>> Criterias { get; private set; } = null!;

        /// <summary>
        /// An optional description of the Dynamic Security Group
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A set of member VMs (if exist). see Member VMs below for details.
        /// </summary>
        [Output("memberVms")]
        public Output<ImmutableArray<Outputs.NsxtDynamicSecurityGroupMemberVm>> MemberVms { get; private set; } = null!;

        /// <summary>
        /// A unique name for Dynamic Security Group
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful
        /// when connected as sysadmin working across different organisations.
        /// </summary>
        [Output("org")]
        public Output<string?> Org { get; private set; } = null!;

        /// <summary>
        /// VDC Group ID for Dynamic Security Group creation.
        /// </summary>
        [Output("vdcGroupId")]
        public Output<string> VdcGroupId { get; private set; } = null!;


        /// <summary>
        /// Create a NsxtDynamicSecurityGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NsxtDynamicSecurityGroup(string name, NsxtDynamicSecurityGroupArgs args, CustomResourceOptions? options = null)
            : base("vcd:index/nsxtDynamicSecurityGroup:NsxtDynamicSecurityGroup", name, args ?? new NsxtDynamicSecurityGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NsxtDynamicSecurityGroup(string name, Input<string> id, NsxtDynamicSecurityGroupState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/nsxtDynamicSecurityGroup:NsxtDynamicSecurityGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NsxtDynamicSecurityGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NsxtDynamicSecurityGroup Get(string name, Input<string> id, NsxtDynamicSecurityGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new NsxtDynamicSecurityGroup(name, id, state, options);
        }
    }

    public sealed class NsxtDynamicSecurityGroupArgs : global::Pulumi.ResourceArgs
    {
        [Input("criterias")]
        private InputList<Inputs.NsxtDynamicSecurityGroupCriteriaArgs>? _criterias;

        /// <summary>
        /// Up to 3 criteria for matching VMs. List of criteria is matched with boolean
        /// `OR` operation and matching any of defined criteria will include objects. Each `criteria` can
        /// contains up to 4 `rule` definitions.
        /// </summary>
        public InputList<Inputs.NsxtDynamicSecurityGroupCriteriaArgs> Criterias
        {
            get => _criterias ?? (_criterias = new InputList<Inputs.NsxtDynamicSecurityGroupCriteriaArgs>());
            set => _criterias = value;
        }

        /// <summary>
        /// An optional description of the Dynamic Security Group
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A unique name for Dynamic Security Group
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful
        /// when connected as sysadmin working across different organisations.
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// VDC Group ID for Dynamic Security Group creation.
        /// </summary>
        [Input("vdcGroupId", required: true)]
        public Input<string> VdcGroupId { get; set; } = null!;

        public NsxtDynamicSecurityGroupArgs()
        {
        }
        public static new NsxtDynamicSecurityGroupArgs Empty => new NsxtDynamicSecurityGroupArgs();
    }

    public sealed class NsxtDynamicSecurityGroupState : global::Pulumi.ResourceArgs
    {
        [Input("criterias")]
        private InputList<Inputs.NsxtDynamicSecurityGroupCriteriaGetArgs>? _criterias;

        /// <summary>
        /// Up to 3 criteria for matching VMs. List of criteria is matched with boolean
        /// `OR` operation and matching any of defined criteria will include objects. Each `criteria` can
        /// contains up to 4 `rule` definitions.
        /// </summary>
        public InputList<Inputs.NsxtDynamicSecurityGroupCriteriaGetArgs> Criterias
        {
            get => _criterias ?? (_criterias = new InputList<Inputs.NsxtDynamicSecurityGroupCriteriaGetArgs>());
            set => _criterias = value;
        }

        /// <summary>
        /// An optional description of the Dynamic Security Group
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("memberVms")]
        private InputList<Inputs.NsxtDynamicSecurityGroupMemberVmGetArgs>? _memberVms;

        /// <summary>
        /// A set of member VMs (if exist). see Member VMs below for details.
        /// </summary>
        public InputList<Inputs.NsxtDynamicSecurityGroupMemberVmGetArgs> MemberVms
        {
            get => _memberVms ?? (_memberVms = new InputList<Inputs.NsxtDynamicSecurityGroupMemberVmGetArgs>());
            set => _memberVms = value;
        }

        /// <summary>
        /// A unique name for Dynamic Security Group
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful
        /// when connected as sysadmin working across different organisations.
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// VDC Group ID for Dynamic Security Group creation.
        /// </summary>
        [Input("vdcGroupId")]
        public Input<string>? VdcGroupId { get; set; }

        public NsxtDynamicSecurityGroupState()
        {
        }
        public static new NsxtDynamicSecurityGroupState Empty => new NsxtDynamicSecurityGroupState();
    }
}
