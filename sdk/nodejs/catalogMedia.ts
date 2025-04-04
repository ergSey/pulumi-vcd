// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

export class CatalogMedia extends pulumi.CustomResource {
    /**
     * Get an existing CatalogMedia resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CatalogMediaState, opts?: pulumi.CustomResourceOptions): CatalogMedia {
        return new CatalogMedia(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/catalogMedia:CatalogMedia';

    /**
     * Returns true if the given object is an instance of CatalogMedia.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CatalogMedia {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CatalogMedia.__pulumiType;
    }

    /**
     * The name of the catalog where to upload media file. It's mandatory if `catalogId` is not used.
     *
     * @deprecated Use catalogId instead
     */
    public readonly catalog!: pulumi.Output<string>;
    /**
     * The ID of the catalog where to upload media file. It's mandatory if `catalog` field is not used.
     */
    public readonly catalogId!: pulumi.Output<string>;
    /**
     * Catalog Item ID of this media item
     */
    public /*out*/ readonly catalogItemId!: pulumi.Output<string>;
    /**
     * (Computed) returns creation date
     */
    public /*out*/ readonly creationDate!: pulumi.Output<string>;
    /**
     * Description of media file
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * (Computed) returns True if this media file is ISO
     */
    public /*out*/ readonly isIso!: pulumi.Output<boolean>;
    /**
     * (Computed) returns True if this media file is in a published catalog
     */
    public /*out*/ readonly isPublished!: pulumi.Output<boolean>;
    /**
     * Absolute or relative path to file to upload
     */
    public readonly mediaPath!: pulumi.Output<string | undefined>;
    /**
     * Use `metadataEntry` instead. Key value map of metadata to assign
     *
     * @deprecated Use metadataEntry instead
     */
    public readonly metadata!: pulumi.Output<{[key: string]: string}>;
    /**
     * A set of metadata entries to assign. See Metadata section for details.
     */
    public readonly metadataEntries!: pulumi.Output<outputs.CatalogMediaMetadataEntry[]>;
    /**
     * Media file name in catalog
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
     */
    public readonly org!: pulumi.Output<string>;
    /**
     * (Computed) returns owner name
     */
    public /*out*/ readonly ownerName!: pulumi.Output<string>;
    /**
     * Default false. Allows to see upload progress. (See note below)
     */
    public readonly showUploadProgress!: pulumi.Output<boolean | undefined>;
    /**
     * (Computed) returns media storage in Bytes
     */
    public /*out*/ readonly size!: pulumi.Output<number>;
    /**
     * (Computed) returns media status
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * (Computed) returns storage profile name
     */
    public /*out*/ readonly storageProfileName!: pulumi.Output<string>;
    /**
     * If `true`, allows uploading any file type. With the default `false`, we can only upload `.ISO` files.
     */
    public readonly uploadAnyFile!: pulumi.Output<boolean | undefined>;
    /**
     * size in MB for splitting upload size. It can possibly impact upload performance. Default 1MB.
     */
    public readonly uploadPieceSize!: pulumi.Output<number | undefined>;

    /**
     * Create a CatalogMedia resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: CatalogMediaArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CatalogMediaArgs | CatalogMediaState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CatalogMediaState | undefined;
            resourceInputs["catalog"] = state ? state.catalog : undefined;
            resourceInputs["catalogId"] = state ? state.catalogId : undefined;
            resourceInputs["catalogItemId"] = state ? state.catalogItemId : undefined;
            resourceInputs["creationDate"] = state ? state.creationDate : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["isIso"] = state ? state.isIso : undefined;
            resourceInputs["isPublished"] = state ? state.isPublished : undefined;
            resourceInputs["mediaPath"] = state ? state.mediaPath : undefined;
            resourceInputs["metadata"] = state ? state.metadata : undefined;
            resourceInputs["metadataEntries"] = state ? state.metadataEntries : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["org"] = state ? state.org : undefined;
            resourceInputs["ownerName"] = state ? state.ownerName : undefined;
            resourceInputs["showUploadProgress"] = state ? state.showUploadProgress : undefined;
            resourceInputs["size"] = state ? state.size : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["storageProfileName"] = state ? state.storageProfileName : undefined;
            resourceInputs["uploadAnyFile"] = state ? state.uploadAnyFile : undefined;
            resourceInputs["uploadPieceSize"] = state ? state.uploadPieceSize : undefined;
        } else {
            const args = argsOrState as CatalogMediaArgs | undefined;
            resourceInputs["catalog"] = args ? args.catalog : undefined;
            resourceInputs["catalogId"] = args ? args.catalogId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["mediaPath"] = args ? args.mediaPath : undefined;
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["metadataEntries"] = args ? args.metadataEntries : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["org"] = args ? args.org : undefined;
            resourceInputs["showUploadProgress"] = args ? args.showUploadProgress : undefined;
            resourceInputs["uploadAnyFile"] = args ? args.uploadAnyFile : undefined;
            resourceInputs["uploadPieceSize"] = args ? args.uploadPieceSize : undefined;
            resourceInputs["catalogItemId"] = undefined /*out*/;
            resourceInputs["creationDate"] = undefined /*out*/;
            resourceInputs["isIso"] = undefined /*out*/;
            resourceInputs["isPublished"] = undefined /*out*/;
            resourceInputs["ownerName"] = undefined /*out*/;
            resourceInputs["size"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["storageProfileName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(CatalogMedia.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CatalogMedia resources.
 */
export interface CatalogMediaState {
    /**
     * The name of the catalog where to upload media file. It's mandatory if `catalogId` is not used.
     *
     * @deprecated Use catalogId instead
     */
    catalog?: pulumi.Input<string>;
    /**
     * The ID of the catalog where to upload media file. It's mandatory if `catalog` field is not used.
     */
    catalogId?: pulumi.Input<string>;
    /**
     * Catalog Item ID of this media item
     */
    catalogItemId?: pulumi.Input<string>;
    /**
     * (Computed) returns creation date
     */
    creationDate?: pulumi.Input<string>;
    /**
     * Description of media file
     */
    description?: pulumi.Input<string>;
    /**
     * (Computed) returns True if this media file is ISO
     */
    isIso?: pulumi.Input<boolean>;
    /**
     * (Computed) returns True if this media file is in a published catalog
     */
    isPublished?: pulumi.Input<boolean>;
    /**
     * Absolute or relative path to file to upload
     */
    mediaPath?: pulumi.Input<string>;
    /**
     * Use `metadataEntry` instead. Key value map of metadata to assign
     *
     * @deprecated Use metadataEntry instead
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A set of metadata entries to assign. See Metadata section for details.
     */
    metadataEntries?: pulumi.Input<pulumi.Input<inputs.CatalogMediaMetadataEntry>[]>;
    /**
     * Media file name in catalog
     */
    name?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
     */
    org?: pulumi.Input<string>;
    /**
     * (Computed) returns owner name
     */
    ownerName?: pulumi.Input<string>;
    /**
     * Default false. Allows to see upload progress. (See note below)
     */
    showUploadProgress?: pulumi.Input<boolean>;
    /**
     * (Computed) returns media storage in Bytes
     */
    size?: pulumi.Input<number>;
    /**
     * (Computed) returns media status
     */
    status?: pulumi.Input<string>;
    /**
     * (Computed) returns storage profile name
     */
    storageProfileName?: pulumi.Input<string>;
    /**
     * If `true`, allows uploading any file type. With the default `false`, we can only upload `.ISO` files.
     */
    uploadAnyFile?: pulumi.Input<boolean>;
    /**
     * size in MB for splitting upload size. It can possibly impact upload performance. Default 1MB.
     */
    uploadPieceSize?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a CatalogMedia resource.
 */
export interface CatalogMediaArgs {
    /**
     * The name of the catalog where to upload media file. It's mandatory if `catalogId` is not used.
     *
     * @deprecated Use catalogId instead
     */
    catalog?: pulumi.Input<string>;
    /**
     * The ID of the catalog where to upload media file. It's mandatory if `catalog` field is not used.
     */
    catalogId?: pulumi.Input<string>;
    /**
     * Description of media file
     */
    description?: pulumi.Input<string>;
    /**
     * Absolute or relative path to file to upload
     */
    mediaPath?: pulumi.Input<string>;
    /**
     * Use `metadataEntry` instead. Key value map of metadata to assign
     *
     * @deprecated Use metadataEntry instead
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A set of metadata entries to assign. See Metadata section for details.
     */
    metadataEntries?: pulumi.Input<pulumi.Input<inputs.CatalogMediaMetadataEntry>[]>;
    /**
     * Media file name in catalog
     */
    name?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
     */
    org?: pulumi.Input<string>;
    /**
     * Default false. Allows to see upload progress. (See note below)
     */
    showUploadProgress?: pulumi.Input<boolean>;
    /**
     * If `true`, allows uploading any file type. With the default `false`, we can only upload `.ISO` files.
     */
    uploadAnyFile?: pulumi.Input<boolean>;
    /**
     * size in MB for splitting upload size. It can possibly impact upload performance. Default 1MB.
     */
    uploadPieceSize?: pulumi.Input<number>;
}
