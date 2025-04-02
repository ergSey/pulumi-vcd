// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/vappFirewallRules:VappFirewallRules")]
    public partial class VappFirewallRules : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Either 'allow' or 'drop'. Specifies what to do should none of the rules match.
        /// </summary>
        [Output("defaultAction")]
        public Output<string> DefaultAction { get; private set; } = null!;

        /// <summary>
        /// Enable or disable firewall. Default is `true`.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// Flag to enable logging for default action. Default value is `false`.
        /// </summary>
        [Output("logDefaultAction")]
        public Output<bool?> LogDefaultAction { get; private set; } = null!;

        /// <summary>
        /// The identifier of [vApp network](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp_network).
        /// </summary>
        [Output("networkId")]
        public Output<string> NetworkId { get; private set; } = null!;

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.
        /// </summary>
        [Output("org")]
        public Output<string?> Org { get; private set; } = null!;

        /// <summary>
        /// Configures a firewall rule; see Rules below for details.
        /// 
        /// &lt;a id="rules"&gt;&lt;/a&gt;
        /// </summary>
        [Output("rules")]
        public Output<ImmutableArray<Outputs.VappFirewallRulesRule>> Rules { get; private set; } = null!;

        /// <summary>
        /// The identifier of [vApp](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp).
        /// </summary>
        [Output("vappId")]
        public Output<string> VappId { get; private set; } = null!;

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level.
        /// </summary>
        [Output("vdc")]
        public Output<string?> Vdc { get; private set; } = null!;


        /// <summary>
        /// Create a VappFirewallRules resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public VappFirewallRules(string name, VappFirewallRulesArgs args, CustomResourceOptions? options = null)
            : base("vcd:index/vappFirewallRules:VappFirewallRules", name, args ?? new VappFirewallRulesArgs(), MakeResourceOptions(options, ""))
        {
        }

        private VappFirewallRules(string name, Input<string> id, VappFirewallRulesState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/vappFirewallRules:VappFirewallRules", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing VappFirewallRules resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static VappFirewallRules Get(string name, Input<string> id, VappFirewallRulesState? state = null, CustomResourceOptions? options = null)
        {
            return new VappFirewallRules(name, id, state, options);
        }
    }

    public sealed class VappFirewallRulesArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Either 'allow' or 'drop'. Specifies what to do should none of the rules match.
        /// </summary>
        [Input("defaultAction", required: true)]
        public Input<string> DefaultAction { get; set; } = null!;

        /// <summary>
        /// Enable or disable firewall. Default is `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Flag to enable logging for default action. Default value is `false`.
        /// </summary>
        [Input("logDefaultAction")]
        public Input<bool>? LogDefaultAction { get; set; }

        /// <summary>
        /// The identifier of [vApp network](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp_network).
        /// </summary>
        [Input("networkId", required: true)]
        public Input<string> NetworkId { get; set; } = null!;

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        [Input("rules")]
        private InputList<Inputs.VappFirewallRulesRuleArgs>? _rules;

        /// <summary>
        /// Configures a firewall rule; see Rules below for details.
        /// 
        /// &lt;a id="rules"&gt;&lt;/a&gt;
        /// </summary>
        public InputList<Inputs.VappFirewallRulesRuleArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.VappFirewallRulesRuleArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// The identifier of [vApp](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp).
        /// </summary>
        [Input("vappId", required: true)]
        public Input<string> VappId { get; set; } = null!;

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level.
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public VappFirewallRulesArgs()
        {
        }
        public static new VappFirewallRulesArgs Empty => new VappFirewallRulesArgs();
    }

    public sealed class VappFirewallRulesState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Either 'allow' or 'drop'. Specifies what to do should none of the rules match.
        /// </summary>
        [Input("defaultAction")]
        public Input<string>? DefaultAction { get; set; }

        /// <summary>
        /// Enable or disable firewall. Default is `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// Flag to enable logging for default action. Default value is `false`.
        /// </summary>
        [Input("logDefaultAction")]
        public Input<bool>? LogDefaultAction { get; set; }

        /// <summary>
        /// The identifier of [vApp network](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp_network).
        /// </summary>
        [Input("networkId")]
        public Input<string>? NetworkId { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        [Input("rules")]
        private InputList<Inputs.VappFirewallRulesRuleGetArgs>? _rules;

        /// <summary>
        /// Configures a firewall rule; see Rules below for details.
        /// 
        /// &lt;a id="rules"&gt;&lt;/a&gt;
        /// </summary>
        public InputList<Inputs.VappFirewallRulesRuleGetArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.VappFirewallRulesRuleGetArgs>());
            set => _rules = value;
        }

        /// <summary>
        /// The identifier of [vApp](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp).
        /// </summary>
        [Input("vappId")]
        public Input<string>? VappId { get; set; }

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level.
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public VappFirewallRulesState()
        {
        }
        public static new VappFirewallRulesState Empty => new VappFirewallRulesState();
    }
}
