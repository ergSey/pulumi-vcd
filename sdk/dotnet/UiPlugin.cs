// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/uiPlugin:UiPlugin")]
    public partial class UiPlugin : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The description of the UI Plugin
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// Whether the UI Plugin will be enabled (`true`) or not (`false`)
        /// </summary>
        [Output("enabled")]
        public Output<bool> Enabled { get; private set; } = null!;

        /// <summary>
        /// The license of the UI Plugin
        /// </summary>
        [Output("license")]
        public Output<string> License { get; private set; } = null!;

        /// <summary>
        /// The website or custom URL of the UI Plugin
        /// </summary>
        [Output("link")]
        public Output<string> Link { get; private set; } = null!;

        /// <summary>
        /// The name of the UI Plugin
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Path to a .zip file that contains the bundled UI Plugin
        /// </summary>
        [Output("pluginPath")]
        public Output<string?> PluginPath { get; private set; } = null!;

        /// <summary>
        /// **Can only be set on updates**, the initial value is taken from the JSON manifest.
        /// Changes the scope of the UI Plugin for System providers. It should be set to `true` when the UI Plugin is published to the System organization, to prevent
        /// unwanted updates-in-place.
        /// </summary>
        [Output("providerScoped")]
        public Output<bool> ProviderScoped { get; private set; } = null!;

        /// <summary>
        /// The status of the UI Plugin (for example, `ready`, `unavailable`, etc)
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// The identifiers of the [Organizations](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/org)
        /// that will be able to use the UI Plugin if enabled. If not set, it doesn't publish to any Organization.
        /// </summary>
        [Output("tenantIds")]
        public Output<ImmutableArray<string>> TenantIds { get; private set; } = null!;

        /// <summary>
        /// **Can only be set on updates**, the initial value is taken from the JSON manifest.
        /// Changes the scope of the UI Plugin for Organization users. It should be set to `true` when the UI Plugin is published to any organization, to prevent
        /// unwanted updates-in-place.
        /// </summary>
        [Output("tenantScoped")]
        public Output<bool> TenantScoped { get; private set; } = null!;

        /// <summary>
        /// The vendor of the UI Plugin
        /// </summary>
        [Output("vendor")]
        public Output<string> Vendor { get; private set; } = null!;

        /// <summary>
        /// The version of the UI Plugin
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a UiPlugin resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public UiPlugin(string name, UiPluginArgs args, CustomResourceOptions? options = null)
            : base("vcd:index/uiPlugin:UiPlugin", name, args ?? new UiPluginArgs(), MakeResourceOptions(options, ""))
        {
        }

        private UiPlugin(string name, Input<string> id, UiPluginState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/uiPlugin:UiPlugin", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing UiPlugin resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static UiPlugin Get(string name, Input<string> id, UiPluginState? state = null, CustomResourceOptions? options = null)
        {
            return new UiPlugin(name, id, state, options);
        }
    }

    public sealed class UiPluginArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the UI Plugin will be enabled (`true`) or not (`false`)
        /// </summary>
        [Input("enabled", required: true)]
        public Input<bool> Enabled { get; set; } = null!;

        /// <summary>
        /// Path to a .zip file that contains the bundled UI Plugin
        /// </summary>
        [Input("pluginPath")]
        public Input<string>? PluginPath { get; set; }

        /// <summary>
        /// **Can only be set on updates**, the initial value is taken from the JSON manifest.
        /// Changes the scope of the UI Plugin for System providers. It should be set to `true` when the UI Plugin is published to the System organization, to prevent
        /// unwanted updates-in-place.
        /// </summary>
        [Input("providerScoped")]
        public Input<bool>? ProviderScoped { get; set; }

        [Input("tenantIds")]
        private InputList<string>? _tenantIds;

        /// <summary>
        /// The identifiers of the [Organizations](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/org)
        /// that will be able to use the UI Plugin if enabled. If not set, it doesn't publish to any Organization.
        /// </summary>
        public InputList<string> TenantIds
        {
            get => _tenantIds ?? (_tenantIds = new InputList<string>());
            set => _tenantIds = value;
        }

        /// <summary>
        /// **Can only be set on updates**, the initial value is taken from the JSON manifest.
        /// Changes the scope of the UI Plugin for Organization users. It should be set to `true` when the UI Plugin is published to any organization, to prevent
        /// unwanted updates-in-place.
        /// </summary>
        [Input("tenantScoped")]
        public Input<bool>? TenantScoped { get; set; }

        public UiPluginArgs()
        {
        }
        public static new UiPluginArgs Empty => new UiPluginArgs();
    }

    public sealed class UiPluginState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The description of the UI Plugin
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Whether the UI Plugin will be enabled (`true`) or not (`false`)
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The license of the UI Plugin
        /// </summary>
        [Input("license")]
        public Input<string>? License { get; set; }

        /// <summary>
        /// The website or custom URL of the UI Plugin
        /// </summary>
        [Input("link")]
        public Input<string>? Link { get; set; }

        /// <summary>
        /// The name of the UI Plugin
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Path to a .zip file that contains the bundled UI Plugin
        /// </summary>
        [Input("pluginPath")]
        public Input<string>? PluginPath { get; set; }

        /// <summary>
        /// **Can only be set on updates**, the initial value is taken from the JSON manifest.
        /// Changes the scope of the UI Plugin for System providers. It should be set to `true` when the UI Plugin is published to the System organization, to prevent
        /// unwanted updates-in-place.
        /// </summary>
        [Input("providerScoped")]
        public Input<bool>? ProviderScoped { get; set; }

        /// <summary>
        /// The status of the UI Plugin (for example, `ready`, `unavailable`, etc)
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        [Input("tenantIds")]
        private InputList<string>? _tenantIds;

        /// <summary>
        /// The identifiers of the [Organizations](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/org)
        /// that will be able to use the UI Plugin if enabled. If not set, it doesn't publish to any Organization.
        /// </summary>
        public InputList<string> TenantIds
        {
            get => _tenantIds ?? (_tenantIds = new InputList<string>());
            set => _tenantIds = value;
        }

        /// <summary>
        /// **Can only be set on updates**, the initial value is taken from the JSON manifest.
        /// Changes the scope of the UI Plugin for Organization users. It should be set to `true` when the UI Plugin is published to any organization, to prevent
        /// unwanted updates-in-place.
        /// </summary>
        [Input("tenantScoped")]
        public Input<bool>? TenantScoped { get; set; }

        /// <summary>
        /// The vendor of the UI Plugin
        /// </summary>
        [Input("vendor")]
        public Input<string>? Vendor { get; set; }

        /// <summary>
        /// The version of the UI Plugin
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public UiPluginState()
        {
        }
        public static new UiPluginState Empty => new UiPluginState();
    }
}
