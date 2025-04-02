// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/clonedVapp:ClonedVapp")]
    public partial class ClonedVapp : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A boolean value of `true` or `false` stating if the source entity should be deleted after creation.
        /// A source vApp can only be deleted if it is fully powered off.
        /// </summary>
        [Output("deleteSource")]
        public Output<bool?> DeleteSource { get; private set; } = null!;

        /// <summary>
        /// An optional description for the vApp, up to 256 characters.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// (Computed) The vApp Hyper Reference.
        /// </summary>
        [Output("href")]
        public Output<string> Href { get; private set; } = null!;

        /// <summary>
        /// A unique name for the vApp
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
        /// </summary>
        [Output("org")]
        public Output<string?> Org { get; private set; } = null!;

        /// <summary>
        /// A boolean value stating if this vApp should be powered on. Default is `false`.
        /// </summary>
        [Output("powerOn")]
        public Output<bool?> PowerOn { get; private set; } = null!;

        /// <summary>
        /// The ID of the source to use.
        /// </summary>
        [Output("sourceId")]
        public Output<string> SourceId { get; private set; } = null!;

        /// <summary>
        /// The type of the source to use: one of `template` or `vapp`.
        /// </summary>
        [Output("sourceType")]
        public Output<string> SourceType { get; private set; } = null!;

        /// <summary>
        /// (Computed) The vApp status as a numeric code.
        /// </summary>
        [Output("status")]
        public Output<int> Status { get; private set; } = null!;

        /// <summary>
        /// (Computed) The vApp status as text.
        /// </summary>
        [Output("statusText")]
        public Output<string> StatusText { get; private set; } = null!;

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Output("vdc")]
        public Output<string?> Vdc { get; private set; } = null!;

        /// <summary>
        /// (Computed) The list of VM names included in this vApp, in alphabetic order.
        /// </summary>
        [Output("vmLists")]
        public Output<ImmutableArray<string>> VmLists { get; private set; } = null!;


        /// <summary>
        /// Create a ClonedVapp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClonedVapp(string name, ClonedVappArgs args, CustomResourceOptions? options = null)
            : base("vcd:index/clonedVapp:ClonedVapp", name, args ?? new ClonedVappArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ClonedVapp(string name, Input<string> id, ClonedVappState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/clonedVapp:ClonedVapp", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ClonedVapp resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ClonedVapp Get(string name, Input<string> id, ClonedVappState? state = null, CustomResourceOptions? options = null)
        {
            return new ClonedVapp(name, id, state, options);
        }
    }

    public sealed class ClonedVappArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A boolean value of `true` or `false` stating if the source entity should be deleted after creation.
        /// A source vApp can only be deleted if it is fully powered off.
        /// </summary>
        [Input("deleteSource")]
        public Input<bool>? DeleteSource { get; set; }

        /// <summary>
        /// An optional description for the vApp, up to 256 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// A unique name for the vApp
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// A boolean value stating if this vApp should be powered on. Default is `false`.
        /// </summary>
        [Input("powerOn")]
        public Input<bool>? PowerOn { get; set; }

        /// <summary>
        /// The ID of the source to use.
        /// </summary>
        [Input("sourceId", required: true)]
        public Input<string> SourceId { get; set; } = null!;

        /// <summary>
        /// The type of the source to use: one of `template` or `vapp`.
        /// </summary>
        [Input("sourceType", required: true)]
        public Input<string> SourceType { get; set; } = null!;

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public ClonedVappArgs()
        {
        }
        public static new ClonedVappArgs Empty => new ClonedVappArgs();
    }

    public sealed class ClonedVappState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A boolean value of `true` or `false` stating if the source entity should be deleted after creation.
        /// A source vApp can only be deleted if it is fully powered off.
        /// </summary>
        [Input("deleteSource")]
        public Input<bool>? DeleteSource { get; set; }

        /// <summary>
        /// An optional description for the vApp, up to 256 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// (Computed) The vApp Hyper Reference.
        /// </summary>
        [Input("href")]
        public Input<string>? Href { get; set; }

        /// <summary>
        /// A unique name for the vApp
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// A boolean value stating if this vApp should be powered on. Default is `false`.
        /// </summary>
        [Input("powerOn")]
        public Input<bool>? PowerOn { get; set; }

        /// <summary>
        /// The ID of the source to use.
        /// </summary>
        [Input("sourceId")]
        public Input<string>? SourceId { get; set; }

        /// <summary>
        /// The type of the source to use: one of `template` or `vapp`.
        /// </summary>
        [Input("sourceType")]
        public Input<string>? SourceType { get; set; }

        /// <summary>
        /// (Computed) The vApp status as a numeric code.
        /// </summary>
        [Input("status")]
        public Input<int>? Status { get; set; }

        /// <summary>
        /// (Computed) The vApp status as text.
        /// </summary>
        [Input("statusText")]
        public Input<string>? StatusText { get; set; }

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        [Input("vmLists")]
        private InputList<string>? _vmLists;

        /// <summary>
        /// (Computed) The list of VM names included in this vApp, in alphabetic order.
        /// </summary>
        public InputList<string> VmLists
        {
            get => _vmLists ?? (_vmLists = new InputList<string>());
            set => _vmLists = value;
        }

        public ClonedVappState()
        {
        }
        public static new ClonedVappState Empty => new ClonedVappState();
    }
}
