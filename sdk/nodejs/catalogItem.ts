// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

export class CatalogItem extends pulumi.CustomResource {
    /**
     * Get an existing CatalogItem resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CatalogItemState, opts?: pulumi.CustomResourceOptions): CatalogItem {
        return new CatalogItem(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/catalogItem:CatalogItem';

    /**
     * Returns true if the given object is an instance of CatalogItem.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CatalogItem {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CatalogItem.__pulumiType;
    }

    /**
     * The name of the catalog where to upload OVA file
     */
    public readonly catalog!: pulumi.Output<string>;
    /**
     * Use `metadataEntry` instead.  Key value map of metadata to assign to the Catalog Item
     *
     * > This resource handles metadata in the following way: `metadata` attribute assigns metadata to the associated **vApp Template**.
     * `metadataEntry` attribute assigns metadata to the **Catalog Item**. `catalogItemMetadata` is deprecated and should not be used.
     *
     * <a id="metadata"></a>
     *
     * @deprecated Use metadataEntry instead
     */
    public readonly catalogItemMetadata!: pulumi.Output<{[key: string]: string}>;
    /**
     * Time stamp of when the item was created
     */
    public /*out*/ readonly created!: pulumi.Output<string>;
    /**
     * Description of item
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Key value map of metadata to assign to the associated vApp Template
     */
    public readonly metadata!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A set of metadata entries to assign to the Catalog Item. See Metadata section for details.
     */
    public readonly metadataEntries!: pulumi.Output<outputs.CatalogItemMetadataEntry[]>;
    /**
     * Item name in catalog
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
     */
    public readonly org!: pulumi.Output<string | undefined>;
    /**
     * Absolute or relative path to file to upload
     */
    public readonly ovaPath!: pulumi.Output<string | undefined>;
    /**
     * URL to OVF file. Only OVF (not OVA) files are supported by VCD uploading by URL
     */
    public readonly ovfUrl!: pulumi.Output<string | undefined>;
    /**
     * Default false. Allows seeing upload progress. (See note below)
     */
    public readonly showUploadProgress!: pulumi.Output<boolean | undefined>;
    /**
     * Size in MB for splitting upload size. It can possibly impact upload performance. Default 1MB.
     */
    public readonly uploadPieceSize!: pulumi.Output<number | undefined>;

    /**
     * Create a CatalogItem resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CatalogItemArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CatalogItemArgs | CatalogItemState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CatalogItemState | undefined;
            resourceInputs["catalog"] = state ? state.catalog : undefined;
            resourceInputs["catalogItemMetadata"] = state ? state.catalogItemMetadata : undefined;
            resourceInputs["created"] = state ? state.created : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["metadata"] = state ? state.metadata : undefined;
            resourceInputs["metadataEntries"] = state ? state.metadataEntries : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["org"] = state ? state.org : undefined;
            resourceInputs["ovaPath"] = state ? state.ovaPath : undefined;
            resourceInputs["ovfUrl"] = state ? state.ovfUrl : undefined;
            resourceInputs["showUploadProgress"] = state ? state.showUploadProgress : undefined;
            resourceInputs["uploadPieceSize"] = state ? state.uploadPieceSize : undefined;
        } else {
            const args = argsOrState as CatalogItemArgs | undefined;
            if ((!args || args.catalog === undefined) && !opts.urn) {
                throw new Error("Missing required property 'catalog'");
            }
            resourceInputs["catalog"] = args ? args.catalog : undefined;
            resourceInputs["catalogItemMetadata"] = args ? args.catalogItemMetadata : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["metadataEntries"] = args ? args.metadataEntries : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["org"] = args ? args.org : undefined;
            resourceInputs["ovaPath"] = args ? args.ovaPath : undefined;
            resourceInputs["ovfUrl"] = args ? args.ovfUrl : undefined;
            resourceInputs["showUploadProgress"] = args ? args.showUploadProgress : undefined;
            resourceInputs["uploadPieceSize"] = args ? args.uploadPieceSize : undefined;
            resourceInputs["created"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CatalogItem.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CatalogItem resources.
 */
export interface CatalogItemState {
    /**
     * The name of the catalog where to upload OVA file
     */
    catalog?: pulumi.Input<string>;
    /**
     * Use `metadataEntry` instead.  Key value map of metadata to assign to the Catalog Item
     *
     * > This resource handles metadata in the following way: `metadata` attribute assigns metadata to the associated **vApp Template**.
     * `metadataEntry` attribute assigns metadata to the **Catalog Item**. `catalogItemMetadata` is deprecated and should not be used.
     *
     * <a id="metadata"></a>
     *
     * @deprecated Use metadataEntry instead
     */
    catalogItemMetadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Time stamp of when the item was created
     */
    created?: pulumi.Input<string>;
    /**
     * Description of item
     */
    description?: pulumi.Input<string>;
    /**
     * Key value map of metadata to assign to the associated vApp Template
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A set of metadata entries to assign to the Catalog Item. See Metadata section for details.
     */
    metadataEntries?: pulumi.Input<pulumi.Input<inputs.CatalogItemMetadataEntry>[]>;
    /**
     * Item name in catalog
     */
    name?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
     */
    org?: pulumi.Input<string>;
    /**
     * Absolute or relative path to file to upload
     */
    ovaPath?: pulumi.Input<string>;
    /**
     * URL to OVF file. Only OVF (not OVA) files are supported by VCD uploading by URL
     */
    ovfUrl?: pulumi.Input<string>;
    /**
     * Default false. Allows seeing upload progress. (See note below)
     */
    showUploadProgress?: pulumi.Input<boolean>;
    /**
     * Size in MB for splitting upload size. It can possibly impact upload performance. Default 1MB.
     */
    uploadPieceSize?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a CatalogItem resource.
 */
export interface CatalogItemArgs {
    /**
     * The name of the catalog where to upload OVA file
     */
    catalog: pulumi.Input<string>;
    /**
     * Use `metadataEntry` instead.  Key value map of metadata to assign to the Catalog Item
     *
     * > This resource handles metadata in the following way: `metadata` attribute assigns metadata to the associated **vApp Template**.
     * `metadataEntry` attribute assigns metadata to the **Catalog Item**. `catalogItemMetadata` is deprecated and should not be used.
     *
     * <a id="metadata"></a>
     *
     * @deprecated Use metadataEntry instead
     */
    catalogItemMetadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Description of item
     */
    description?: pulumi.Input<string>;
    /**
     * Key value map of metadata to assign to the associated vApp Template
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A set of metadata entries to assign to the Catalog Item. See Metadata section for details.
     */
    metadataEntries?: pulumi.Input<pulumi.Input<inputs.CatalogItemMetadataEntry>[]>;
    /**
     * Item name in catalog
     */
    name?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
     */
    org?: pulumi.Input<string>;
    /**
     * Absolute or relative path to file to upload
     */
    ovaPath?: pulumi.Input<string>;
    /**
     * URL to OVF file. Only OVF (not OVA) files are supported by VCD uploading by URL
     */
    ovfUrl?: pulumi.Input<string>;
    /**
     * Default false. Allows seeing upload progress. (See note below)
     */
    showUploadProgress?: pulumi.Input<boolean>;
    /**
     * Size in MB for splitting upload size. It can possibly impact upload performance. Default 1MB.
     */
    uploadPieceSize?: pulumi.Input<number>;
}
