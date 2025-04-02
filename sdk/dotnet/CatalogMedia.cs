// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/catalogMedia:CatalogMedia")]
    public partial class CatalogMedia : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the catalog where to upload media file. It's mandatory if `catalog_id` is not used.
        /// </summary>
        [Output("catalog")]
        public Output<string> Catalog { get; private set; } = null!;

        /// <summary>
        /// The ID of the catalog where to upload media file. It's mandatory if `catalog` field is not used.
        /// </summary>
        [Output("catalogId")]
        public Output<string> CatalogId { get; private set; } = null!;

        /// <summary>
        /// Catalog Item ID of this media item
        /// </summary>
        [Output("catalogItemId")]
        public Output<string> CatalogItemId { get; private set; } = null!;

        /// <summary>
        /// (Computed) returns creation date
        /// </summary>
        [Output("creationDate")]
        public Output<string> CreationDate { get; private set; } = null!;

        /// <summary>
        /// Description of media file
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// (Computed) returns True if this media file is ISO
        /// </summary>
        [Output("isIso")]
        public Output<bool> IsIso { get; private set; } = null!;

        /// <summary>
        /// (Computed) returns True if this media file is in a published catalog
        /// </summary>
        [Output("isPublished")]
        public Output<bool> IsPublished { get; private set; } = null!;

        /// <summary>
        /// Absolute or relative path to file to upload
        /// </summary>
        [Output("mediaPath")]
        public Output<string?> MediaPath { get; private set; } = null!;

        /// <summary>
        /// Use `metadata_entry` instead. Key value map of metadata to assign
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, string>> Metadata { get; private set; } = null!;

        /// <summary>
        /// A set of metadata entries to assign. See Metadata section for details.
        /// </summary>
        [Output("metadataEntries")]
        public Output<ImmutableArray<Outputs.CatalogMediaMetadataEntry>> MetadataEntries { get; private set; } = null!;

        /// <summary>
        /// Media file name in catalog
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
        /// </summary>
        [Output("org")]
        public Output<string> Org { get; private set; } = null!;

        /// <summary>
        /// (Computed) returns owner name
        /// </summary>
        [Output("ownerName")]
        public Output<string> OwnerName { get; private set; } = null!;

        /// <summary>
        /// Default false. Allows to see upload progress. (See note below)
        /// </summary>
        [Output("showUploadProgress")]
        public Output<bool?> ShowUploadProgress { get; private set; } = null!;

        /// <summary>
        /// (Computed) returns media storage in Bytes
        /// </summary>
        [Output("size")]
        public Output<int> Size { get; private set; } = null!;

        /// <summary>
        /// (Computed) returns media status
        /// </summary>
        [Output("status")]
        public Output<string> Status { get; private set; } = null!;

        /// <summary>
        /// (Computed) returns storage profile name
        /// </summary>
        [Output("storageProfileName")]
        public Output<string> StorageProfileName { get; private set; } = null!;

        /// <summary>
        /// If `true`, allows uploading any file type. With the default `false`, we can only upload `.ISO` files.
        /// </summary>
        [Output("uploadAnyFile")]
        public Output<bool?> UploadAnyFile { get; private set; } = null!;

        /// <summary>
        /// size in MB for splitting upload size. It can possibly impact upload performance. Default 1MB.
        /// </summary>
        [Output("uploadPieceSize")]
        public Output<int?> UploadPieceSize { get; private set; } = null!;


        /// <summary>
        /// Create a CatalogMedia resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public CatalogMedia(string name, CatalogMediaArgs? args = null, CustomResourceOptions? options = null)
            : base("vcd:index/catalogMedia:CatalogMedia", name, args ?? new CatalogMediaArgs(), MakeResourceOptions(options, ""))
        {
        }

        private CatalogMedia(string name, Input<string> id, CatalogMediaState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/catalogMedia:CatalogMedia", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing CatalogMedia resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static CatalogMedia Get(string name, Input<string> id, CatalogMediaState? state = null, CustomResourceOptions? options = null)
        {
            return new CatalogMedia(name, id, state, options);
        }
    }

    public sealed class CatalogMediaArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the catalog where to upload media file. It's mandatory if `catalog_id` is not used.
        /// </summary>
        [Input("catalog")]
        public Input<string>? Catalog { get; set; }

        /// <summary>
        /// The ID of the catalog where to upload media file. It's mandatory if `catalog` field is not used.
        /// </summary>
        [Input("catalogId")]
        public Input<string>? CatalogId { get; set; }

        /// <summary>
        /// Description of media file
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Absolute or relative path to file to upload
        /// </summary>
        [Input("mediaPath")]
        public Input<string>? MediaPath { get; set; }

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// Use `metadata_entry` instead. Key value map of metadata to assign
        /// </summary>
        [Obsolete(@"Use metadata_entry instead")]
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        [Input("metadataEntries")]
        private InputList<Inputs.CatalogMediaMetadataEntryArgs>? _metadataEntries;

        /// <summary>
        /// A set of metadata entries to assign. See Metadata section for details.
        /// </summary>
        public InputList<Inputs.CatalogMediaMetadataEntryArgs> MetadataEntries
        {
            get => _metadataEntries ?? (_metadataEntries = new InputList<Inputs.CatalogMediaMetadataEntryArgs>());
            set => _metadataEntries = value;
        }

        /// <summary>
        /// Media file name in catalog
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// Default false. Allows to see upload progress. (See note below)
        /// </summary>
        [Input("showUploadProgress")]
        public Input<bool>? ShowUploadProgress { get; set; }

        /// <summary>
        /// If `true`, allows uploading any file type. With the default `false`, we can only upload `.ISO` files.
        /// </summary>
        [Input("uploadAnyFile")]
        public Input<bool>? UploadAnyFile { get; set; }

        /// <summary>
        /// size in MB for splitting upload size. It can possibly impact upload performance. Default 1MB.
        /// </summary>
        [Input("uploadPieceSize")]
        public Input<int>? UploadPieceSize { get; set; }

        public CatalogMediaArgs()
        {
        }
        public static new CatalogMediaArgs Empty => new CatalogMediaArgs();
    }

    public sealed class CatalogMediaState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the catalog where to upload media file. It's mandatory if `catalog_id` is not used.
        /// </summary>
        [Input("catalog")]
        public Input<string>? Catalog { get; set; }

        /// <summary>
        /// The ID of the catalog where to upload media file. It's mandatory if `catalog` field is not used.
        /// </summary>
        [Input("catalogId")]
        public Input<string>? CatalogId { get; set; }

        /// <summary>
        /// Catalog Item ID of this media item
        /// </summary>
        [Input("catalogItemId")]
        public Input<string>? CatalogItemId { get; set; }

        /// <summary>
        /// (Computed) returns creation date
        /// </summary>
        [Input("creationDate")]
        public Input<string>? CreationDate { get; set; }

        /// <summary>
        /// Description of media file
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// (Computed) returns True if this media file is ISO
        /// </summary>
        [Input("isIso")]
        public Input<bool>? IsIso { get; set; }

        /// <summary>
        /// (Computed) returns True if this media file is in a published catalog
        /// </summary>
        [Input("isPublished")]
        public Input<bool>? IsPublished { get; set; }

        /// <summary>
        /// Absolute or relative path to file to upload
        /// </summary>
        [Input("mediaPath")]
        public Input<string>? MediaPath { get; set; }

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// Use `metadata_entry` instead. Key value map of metadata to assign
        /// </summary>
        [Obsolete(@"Use metadata_entry instead")]
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        [Input("metadataEntries")]
        private InputList<Inputs.CatalogMediaMetadataEntryGetArgs>? _metadataEntries;

        /// <summary>
        /// A set of metadata entries to assign. See Metadata section for details.
        /// </summary>
        public InputList<Inputs.CatalogMediaMetadataEntryGetArgs> MetadataEntries
        {
            get => _metadataEntries ?? (_metadataEntries = new InputList<Inputs.CatalogMediaMetadataEntryGetArgs>());
            set => _metadataEntries = value;
        }

        /// <summary>
        /// Media file name in catalog
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// (Computed) returns owner name
        /// </summary>
        [Input("ownerName")]
        public Input<string>? OwnerName { get; set; }

        /// <summary>
        /// Default false. Allows to see upload progress. (See note below)
        /// </summary>
        [Input("showUploadProgress")]
        public Input<bool>? ShowUploadProgress { get; set; }

        /// <summary>
        /// (Computed) returns media storage in Bytes
        /// </summary>
        [Input("size")]
        public Input<int>? Size { get; set; }

        /// <summary>
        /// (Computed) returns media status
        /// </summary>
        [Input("status")]
        public Input<string>? Status { get; set; }

        /// <summary>
        /// (Computed) returns storage profile name
        /// </summary>
        [Input("storageProfileName")]
        public Input<string>? StorageProfileName { get; set; }

        /// <summary>
        /// If `true`, allows uploading any file type. With the default `false`, we can only upload `.ISO` files.
        /// </summary>
        [Input("uploadAnyFile")]
        public Input<bool>? UploadAnyFile { get; set; }

        /// <summary>
        /// size in MB for splitting upload size. It can possibly impact upload performance. Default 1MB.
        /// </summary>
        [Input("uploadPieceSize")]
        public Input<int>? UploadPieceSize { get; set; }

        public CatalogMediaState()
        {
        }
        public static new CatalogMediaState Empty => new CatalogMediaState();
    }
}
