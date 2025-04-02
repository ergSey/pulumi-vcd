// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/catalogAccessControl:CatalogAccessControl")]
    public partial class CatalogAccessControl : global::Pulumi.CustomResource
    {
        /// <summary>
        /// A unique identifier for the Catalog.
        /// </summary>
        [Output("catalogId")]
        public Output<string> CatalogId { get; private set; } = null!;

        /// <summary>
        /// Access level when the Catalog is shared with everyone (it can only be set to
        /// `ReadOnly`). Required if `shared_with_everyone` is set.
        /// </summary>
        [Output("everyoneAccessLevel")]
        public Output<string?> EveryoneAccessLevel { get; private set; } = null!;

        /// <summary>
        /// The name of organization to which the Catalog belongs. Optional if defined at provider level.
        /// </summary>
        [Output("org")]
        public Output<string?> Org { get; private set; } = null!;

        /// <summary>
        /// If true, the catalog is shared as read-only with all organizations.
        /// </summary>
        [Output("readOnlySharedWithAllOrgs")]
        public Output<bool> ReadOnlySharedWithAllOrgs { get; private set; } = null!;

        /// <summary>
        /// Whether the Catalog is shared with everyone. If any `shared_with` blocks are included,
        /// this property must be set to `false`.
        /// </summary>
        [Output("sharedWithEveryone")]
        public Output<bool> SharedWithEveryone { get; private set; } = null!;

        /// <summary>
        /// one or more blocks defining a subject (one of Organization, User, or Group) to which we are sharing. 
        /// See shared_with below for detail. It cannot be used if `shared_with_everyone` is true.
        /// </summary>
        [Output("sharedWiths")]
        public Output<ImmutableArray<Outputs.CatalogAccessControlSharedWith>> SharedWiths { get; private set; } = null!;


        /// <summary>
        /// Create a CatalogAccessControl resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CatalogAccessControl(string name, CatalogAccessControlArgs args, CustomResourceOptions? options = null)
            : base("vcd:index/catalogAccessControl:CatalogAccessControl", name, args ?? new CatalogAccessControlArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CatalogAccessControl(string name, Input<string> id, CatalogAccessControlState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/catalogAccessControl:CatalogAccessControl", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CatalogAccessControl resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CatalogAccessControl Get(string name, Input<string> id, CatalogAccessControlState? state = null, CustomResourceOptions? options = null)
        {
            return new CatalogAccessControl(name, id, state, options);
        }
    }

    public sealed class CatalogAccessControlArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique identifier for the Catalog.
        /// </summary>
        [Input("catalogId", required: true)]
        public Input<string> CatalogId { get; set; } = null!;

        /// <summary>
        /// Access level when the Catalog is shared with everyone (it can only be set to
        /// `ReadOnly`). Required if `shared_with_everyone` is set.
        /// </summary>
        [Input("everyoneAccessLevel")]
        public Input<string>? EveryoneAccessLevel { get; set; }

        /// <summary>
        /// The name of organization to which the Catalog belongs. Optional if defined at provider level.
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// If true, the catalog is shared as read-only with all organizations.
        /// </summary>
        [Input("readOnlySharedWithAllOrgs")]
        public Input<bool>? ReadOnlySharedWithAllOrgs { get; set; }

        /// <summary>
        /// Whether the Catalog is shared with everyone. If any `shared_with` blocks are included,
        /// this property must be set to `false`.
        /// </summary>
        [Input("sharedWithEveryone", required: true)]
        public Input<bool> SharedWithEveryone { get; set; } = null!;

        [Input("sharedWiths")]
        private InputList<Inputs.CatalogAccessControlSharedWithArgs>? _sharedWiths;

        /// <summary>
        /// one or more blocks defining a subject (one of Organization, User, or Group) to which we are sharing. 
        /// See shared_with below for detail. It cannot be used if `shared_with_everyone` is true.
        /// </summary>
        public InputList<Inputs.CatalogAccessControlSharedWithArgs> SharedWiths
        {
            get => _sharedWiths ?? (_sharedWiths = new InputList<Inputs.CatalogAccessControlSharedWithArgs>());
            set => _sharedWiths = value;
        }

        public CatalogAccessControlArgs()
        {
        }
        public static new CatalogAccessControlArgs Empty => new CatalogAccessControlArgs();
    }

    public sealed class CatalogAccessControlState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A unique identifier for the Catalog.
        /// </summary>
        [Input("catalogId")]
        public Input<string>? CatalogId { get; set; }

        /// <summary>
        /// Access level when the Catalog is shared with everyone (it can only be set to
        /// `ReadOnly`). Required if `shared_with_everyone` is set.
        /// </summary>
        [Input("everyoneAccessLevel")]
        public Input<string>? EveryoneAccessLevel { get; set; }

        /// <summary>
        /// The name of organization to which the Catalog belongs. Optional if defined at provider level.
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// If true, the catalog is shared as read-only with all organizations.
        /// </summary>
        [Input("readOnlySharedWithAllOrgs")]
        public Input<bool>? ReadOnlySharedWithAllOrgs { get; set; }

        /// <summary>
        /// Whether the Catalog is shared with everyone. If any `shared_with` blocks are included,
        /// this property must be set to `false`.
        /// </summary>
        [Input("sharedWithEveryone")]
        public Input<bool>? SharedWithEveryone { get; set; }

        [Input("sharedWiths")]
        private InputList<Inputs.CatalogAccessControlSharedWithGetArgs>? _sharedWiths;

        /// <summary>
        /// one or more blocks defining a subject (one of Organization, User, or Group) to which we are sharing. 
        /// See shared_with below for detail. It cannot be used if `shared_with_everyone` is true.
        /// </summary>
        public InputList<Inputs.CatalogAccessControlSharedWithGetArgs> SharedWiths
        {
            get => _sharedWiths ?? (_sharedWiths = new InputList<Inputs.CatalogAccessControlSharedWithGetArgs>());
            set => _sharedWiths = value;
        }

        public CatalogAccessControlState()
        {
        }
        public static new CatalogAccessControlState Empty => new CatalogAccessControlState();
    }
}
