// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/vdcGroup:VdcGroup")]
    public partial class VdcGroup : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Whether this security policy is enabled. `dfw_enabled` must be `true`.
        /// </summary>
        [Output("defaultPolicyStatus")]
        public Output<bool> DefaultPolicyStatus { get; private set; } = null!;

        /// <summary>
        /// VDC group description
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// Whether Distributed Firewall is enabled for this VDC group.
        /// </summary>
        [Output("dfwEnabled")]
        public Output<bool> DfwEnabled { get; private set; } = null!;

        /// <summary>
        /// More detailed error message when VDC group has error status
        /// </summary>
        [Output("errorMessage")]
        public Output<string> ErrorMessage { get; private set; } = null!;

        /// <summary>
        /// When `true`, will request VCD to force VDC Group deletion. It
        /// should clean up child components. Default `false` (VCD may fail removing VDC Group if there are
        /// child components remaining). **Note:** when setting it to `true` for existing resource, it will
        /// cause a plan change (update), but this will not alter the resource in any way.
        /// </summary>
        [Output("forceDelete")]
        public Output<bool?> ForceDelete { get; private set; } = null!;

        /// <summary>
        /// Status whether local egress is enabled for a universal router belonging to a universal VDC group.
        /// </summary>
        [Output("localEgress")]
        public Output<bool> LocalEgress { get; private set; } = null!;

        /// <summary>
        /// The name for VDC group
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// ID of used network pool.
        /// </summary>
        [Output("networkPoolId")]
        public Output<string> NetworkPoolId { get; private set; } = null!;

        /// <summary>
        /// The network provider’s universal id that is backing the universal network pool.
        /// </summary>
        [Output("networkPoolUniversalId")]
        public Output<string> NetworkPoolUniversalId { get; private set; } = null!;

        /// <summary>
        /// Defines the networking provider backing the VDC group.
        /// </summary>
        [Output("networkProviderType")]
        public Output<string> NetworkProviderType { get; private set; } = null!;

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
        /// </summary>
        [Output("org")]
        public Output<string?> Org { get; private set; } = null!;

        /// <summary>
        /// A list of blocks providing organization VDCs that are participating in this group details. See Participating Org VDCs below for details.
        /// </summary>
        [Output("participatingOrgVdcs")]
        public Output<ImmutableArray<Outputs.VdcGroupParticipatingOrgVdc>> ParticipatingOrgVdcs { get; private set; } = null!;

        /// <summary>
        /// The list of organization VDCs that are participating in this group. **Note**: `starting_vdc_id` isn't automatically included in this list.
        /// </summary>
        [Output("participatingVdcIds")]
        public Output<ImmutableArray<string>> ParticipatingVdcIds { get; private set; } = null!;

        /// <summary>
        /// Marks whether default firewall rule should be
        /// removed after activating. Both `dfw_enabled` and `default_policy_status` must be true. **Note.**
        /// This is mainly useful when using
        /// [`vcd.NsxtDistributedFirewallRule`](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_distributed_firewall_rule)
        /// resource as it cannot remove the default rule.
        /// </summary>
        [Output("removeDefaultFirewallRule")]
        public Output<bool?> RemoveDefaultFirewallRule { get; private set; } = null!;

        /// <summary>
        /// With selecting a starting VDC you will be able to create a group in which this VDC can participate. **Note**: `starting_vdc_id` must be included in `participating_vdc_ids` to participate in this group.
        /// </summary>
        [Output("startingVdcId")]
        public Output<string> StartingVdcId { get; private set; } = null!;

        /// <summary>
        /// The status that the group can be in (e.g. 'SAVING', 'SAVED', 'CONFIGURING', 'REALIZED', 'REALIZATION_FAILED', 'DELETING', 'DELETE_FAILED', 'OBJECT_NOT_FOUND', 'UNCONFIGURED').
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// Defines the group as LOCAL or UNIVERSAL.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// True means that a VDC group router has been created.
        /// </summary>
        [Output("universalNetworkingEnabled")]
        public Output<bool> UniversalNetworkingEnabled { get; private set; } = null!;


        /// <summary>
        /// Create a VdcGroup resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VdcGroup(string name, VdcGroupArgs args, CustomResourceOptions? options = null)
            : base("vcd:index/vdcGroup:VdcGroup", name, args ?? new VdcGroupArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VdcGroup(string name, Input<string> id, VdcGroupState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/vdcGroup:VdcGroup", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VdcGroup resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VdcGroup Get(string name, Input<string> id, VdcGroupState? state = null, CustomResourceOptions? options = null)
        {
            return new VdcGroup(name, id, state, options);
        }
    }

    public sealed class VdcGroupArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether this security policy is enabled. `dfw_enabled` must be `true`.
        /// </summary>
        [Input("defaultPolicyStatus")]
        public Input<bool>? DefaultPolicyStatus { get; set; }

        /// <summary>
        /// VDC group description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether Distributed Firewall is enabled for this VDC group.
        /// </summary>
        [Input("dfwEnabled")]
        public Input<bool>? DfwEnabled { get; set; }

        /// <summary>
        /// When `true`, will request VCD to force VDC Group deletion. It
        /// should clean up child components. Default `false` (VCD may fail removing VDC Group if there are
        /// child components remaining). **Note:** when setting it to `true` for existing resource, it will
        /// cause a plan change (update), but this will not alter the resource in any way.
        /// </summary>
        [Input("forceDelete")]
        public Input<bool>? ForceDelete { get; set; }

        /// <summary>
        /// The name for VDC group
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        [Input("participatingVdcIds", required: true)]
        private InputList<string>? _participatingVdcIds;

        /// <summary>
        /// The list of organization VDCs that are participating in this group. **Note**: `starting_vdc_id` isn't automatically included in this list.
        /// </summary>
        public InputList<string> ParticipatingVdcIds
        {
            get => _participatingVdcIds ?? (_participatingVdcIds = new InputList<string>());
            set => _participatingVdcIds = value;
        }

        /// <summary>
        /// Marks whether default firewall rule should be
        /// removed after activating. Both `dfw_enabled` and `default_policy_status` must be true. **Note.**
        /// This is mainly useful when using
        /// [`vcd.NsxtDistributedFirewallRule`](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_distributed_firewall_rule)
        /// resource as it cannot remove the default rule.
        /// </summary>
        [Input("removeDefaultFirewallRule")]
        public Input<bool>? RemoveDefaultFirewallRule { get; set; }

        /// <summary>
        /// With selecting a starting VDC you will be able to create a group in which this VDC can participate. **Note**: `starting_vdc_id` must be included in `participating_vdc_ids` to participate in this group.
        /// </summary>
        [Input("startingVdcId", required: true)]
        public Input<string> StartingVdcId { get; set; } = null!;

        public VdcGroupArgs()
        {
        }
        public static new VdcGroupArgs Empty => new VdcGroupArgs();
    }

    public sealed class VdcGroupState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether this security policy is enabled. `dfw_enabled` must be `true`.
        /// </summary>
        [Input("defaultPolicyStatus")]
        public Input<bool>? DefaultPolicyStatus { get; set; }

        /// <summary>
        /// VDC group description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether Distributed Firewall is enabled for this VDC group.
        /// </summary>
        [Input("dfwEnabled")]
        public Input<bool>? DfwEnabled { get; set; }

        /// <summary>
        /// More detailed error message when VDC group has error status
        /// </summary>
        [Input("errorMessage")]
        public Input<string>? ErrorMessage { get; set; }

        /// <summary>
        /// When `true`, will request VCD to force VDC Group deletion. It
        /// should clean up child components. Default `false` (VCD may fail removing VDC Group if there are
        /// child components remaining). **Note:** when setting it to `true` for existing resource, it will
        /// cause a plan change (update), but this will not alter the resource in any way.
        /// </summary>
        [Input("forceDelete")]
        public Input<bool>? ForceDelete { get; set; }

        /// <summary>
        /// Status whether local egress is enabled for a universal router belonging to a universal VDC group.
        /// </summary>
        [Input("localEgress")]
        public Input<bool>? LocalEgress { get; set; }

        /// <summary>
        /// The name for VDC group
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of used network pool.
        /// </summary>
        [Input("networkPoolId")]
        public Input<string>? NetworkPoolId { get; set; }

        /// <summary>
        /// The network provider’s universal id that is backing the universal network pool.
        /// </summary>
        [Input("networkPoolUniversalId")]
        public Input<string>? NetworkPoolUniversalId { get; set; }

        /// <summary>
        /// Defines the networking provider backing the VDC group.
        /// </summary>
        [Input("networkProviderType")]
        public Input<string>? NetworkProviderType { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organizations
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        [Input("participatingOrgVdcs")]
        private InputList<Inputs.VdcGroupParticipatingOrgVdcGetArgs>? _participatingOrgVdcs;

        /// <summary>
        /// A list of blocks providing organization VDCs that are participating in this group details. See Participating Org VDCs below for details.
        /// </summary>
        public InputList<Inputs.VdcGroupParticipatingOrgVdcGetArgs> ParticipatingOrgVdcs
        {
            get => _participatingOrgVdcs ?? (_participatingOrgVdcs = new InputList<Inputs.VdcGroupParticipatingOrgVdcGetArgs>());
            set => _participatingOrgVdcs = value;
        }

        [Input("participatingVdcIds")]
        private InputList<string>? _participatingVdcIds;

        /// <summary>
        /// The list of organization VDCs that are participating in this group. **Note**: `starting_vdc_id` isn't automatically included in this list.
        /// </summary>
        public InputList<string> ParticipatingVdcIds
        {
            get => _participatingVdcIds ?? (_participatingVdcIds = new InputList<string>());
            set => _participatingVdcIds = value;
        }

        /// <summary>
        /// Marks whether default firewall rule should be
        /// removed after activating. Both `dfw_enabled` and `default_policy_status` must be true. **Note.**
        /// This is mainly useful when using
        /// [`vcd.NsxtDistributedFirewallRule`](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_distributed_firewall_rule)
        /// resource as it cannot remove the default rule.
        /// </summary>
        [Input("removeDefaultFirewallRule")]
        public Input<bool>? RemoveDefaultFirewallRule { get; set; }

        /// <summary>
        /// With selecting a starting VDC you will be able to create a group in which this VDC can participate. **Note**: `starting_vdc_id` must be included in `participating_vdc_ids` to participate in this group.
        /// </summary>
        [Input("startingVdcId")]
        public Input<string>? StartingVdcId { get; set; }

        /// <summary>
        /// The status that the group can be in (e.g. 'SAVING', 'SAVED', 'CONFIGURING', 'REALIZED', 'REALIZATION_FAILED', 'DELETING', 'DELETE_FAILED', 'OBJECT_NOT_FOUND', 'UNCONFIGURED').
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// Defines the group as LOCAL or UNIVERSAL.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// True means that a VDC group router has been created.
        /// </summary>
        [Input("universalNetworkingEnabled")]
        public Input<bool>? UniversalNetworkingEnabled { get; set; }

        public VdcGroupState()
        {
        }
        public static new VdcGroupState Empty => new VdcGroupState();
    }
}
