// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class SubscribedCatalog extends pulumi.CustomResource {
    /**
     * Get an existing SubscribedCatalog resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: SubscribedCatalogState, opts?: pulumi.CustomResourceOptions): SubscribedCatalog {
        return new SubscribedCatalog(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/subscribedCatalog:SubscribedCatalog';

    /**
     * Returns true if the given object is an instance of SubscribedCatalog.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is SubscribedCatalog {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === SubscribedCatalog.__pulumiType;
    }

    /**
     * When `true`, the subscribed catalog will attempt canceling failed tasks.
     */
    public readonly cancelFailedTasks!: pulumi.Output<boolean | undefined>;
    /**
     * Version number from this catalog. This is inherited from the publishing catalog and updated on sync.
     */
    public /*out*/ readonly catalogVersion!: pulumi.Output<number>;
    /**
     * Date and time of catalog creation. This is the creation date of the subscription, not the original published catalog.
     */
    public /*out*/ readonly created!: pulumi.Output<string>;
    /**
     * When destroying use `delete_force=true` with `delete_recursive=true` to remove a catalog and any objects it contains, regardless of their state.
     */
    public readonly deleteForce!: pulumi.Output<boolean | undefined>;
    /**
     * When destroying use `delete_recursive=true` to remove the catalog and any objects it contains that are in a state that normally allows removal.
     */
    public readonly deleteRecursive!: pulumi.Output<boolean | undefined>;
    /**
     * Description of catalog. This is inherited from the publishing catalog and updated on sync.
     */
    public /*out*/ readonly description!: pulumi.Output<string>;
    /**
     * List of synchronization tasks that are have failed. They can refer to the catalog or any of its catalog items.
     */
    public /*out*/ readonly failedTasks!: pulumi.Output<string[]>;
    /**
     * the catalog's Hyper reference.
     */
    public /*out*/ readonly href!: pulumi.Output<string>;
    /**
     * (*v3.8.1+*) Indicates if this catalog was created in the current organization.
     */
    public /*out*/ readonly isLocal!: pulumi.Output<boolean>;
    /**
     * Indicates if this catalog is available for subscription. (Always false)
     */
    public /*out*/ readonly isPublished!: pulumi.Output<boolean>;
    /**
     * Indicates if the catalog is shared.
     */
    public /*out*/ readonly isShared!: pulumi.Output<boolean>;
    /**
     * If `true`, subscription to a catalog creates a local copy of all items. Defaults to `false`, which does not create a local copy of catalog items unless a sync operation is performed.
     * It can only be `false` if the user configured in the provider is the System administrator.
     */
    public readonly makeLocalCopy!: pulumi.Output<boolean | undefined>;
    /**
     * List of media item names in this catalog, in alphabetical order.
     */
    public /*out*/ readonly mediaItemLists!: pulumi.Output<string[]>;
    /**
     * (*Available until VCD 10.5*) Optional metadata of the catalog. This is inherited from the publishing catalog and updated on sync.
     */
    public /*out*/ readonly metadata!: pulumi.Output<{[key: string]: string}>;
    /**
     * Catalog name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Number of media items available in this catalog.
     */
    public /*out*/ readonly numberOfMedia!: pulumi.Output<number>;
    /**
     * Number of vApp templates available in this catalog.
     */
    public /*out*/ readonly numberOfVappTemplates!: pulumi.Output<number>;
    /**
     * The name of organization to use, optional if defined at provider level.
     */
    public readonly org!: pulumi.Output<string | undefined>;
    /**
     * Owner of the catalog.
     */
    public /*out*/ readonly ownerName!: pulumi.Output<string>;
    /**
     * Shows if the catalog is published, if it is a subscription from another one or none of those. (Always `SUBSCRIBED`)
     */
    public /*out*/ readonly publishSubscriptionType!: pulumi.Output<string>;
    /**
     * List of running synchronization tasks that are still running. They can refer to the catalog or any of its catalog items.
     */
    public /*out*/ readonly runningTasks!: pulumi.Output<string[]>;
    /**
     * Allows to set specific storage profile to be used for catalog.
     */
    public readonly storageProfileId!: pulumi.Output<string | undefined>;
    /**
     * if `true`, saves the list of tasks to a file for later update.
     */
    public readonly storeTasks!: pulumi.Output<boolean | undefined>;
    /**
     * An optional password to access the catalog. Only ASCII characters are allowed in a valid password. 
     * The password is only required when set by the publishing catalog. Passing in six asterisks '******' indicates to keep current password.
     * Passing in an empty string indicates to remove password.
     */
    public readonly subscriptionPassword!: pulumi.Output<string>;
    /**
     * The URL to subscribe to the external catalog.
     */
    public readonly subscriptionUrl!: pulumi.Output<string>;
    /**
     * If `true`, synchronise this catalog and all items.
     */
    public readonly syncAll!: pulumi.Output<boolean | undefined>;
    /**
     * If `true`, synchronise all media items. Not to be used when `syncAll` is set.
     */
    public readonly syncAllMediaItems!: pulumi.Output<boolean | undefined>;
    /**
     * If `true`, synchronise all vApp templates. Not to be used when `syncAll` is set.
     */
    public readonly syncAllVappTemplates!: pulumi.Output<boolean | undefined>;
    /**
     * If `true`, synchronise this catalog. Not to be used when `syncAll` is set. This operation fetches the list of items. If `makeLocalCopy` is set, it also synchronises all the items.
     */
    public readonly syncCatalog!: pulumi.Output<boolean | undefined>;
    /**
     * Synchronise a list of media items. Not to be used when `syncAll` or `syncAllMediaItems` are set.
     */
    public readonly syncMediaItems!: pulumi.Output<string[] | undefined>;
    /**
     * Boolean value that shows if sync should be performed on every refresh.
     */
    public readonly syncOnRefresh!: pulumi.Output<boolean | undefined>;
    /**
     * Synchronise a list of vApp templates. Not to be used when `syncAll` or `syncAllVappTemplates` are set.
     */
    public readonly syncVappTemplates!: pulumi.Output<string[] | undefined>;
    /**
     * Where the running tasks IDs have been stored. Only if `storeTasks` is set.
     */
    public /*out*/ readonly tasksFileName!: pulumi.Output<string>;
    /**
     * List of vApp template names in this catalog, in alphabetical order.
     */
    public /*out*/ readonly vappTemplateLists!: pulumi.Output<string[]>;

    /**
     * Create a SubscribedCatalog resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: SubscribedCatalogArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: SubscribedCatalogArgs | SubscribedCatalogState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as SubscribedCatalogState | undefined;
            resourceInputs["cancelFailedTasks"] = state ? state.cancelFailedTasks : undefined;
            resourceInputs["catalogVersion"] = state ? state.catalogVersion : undefined;
            resourceInputs["created"] = state ? state.created : undefined;
            resourceInputs["deleteForce"] = state ? state.deleteForce : undefined;
            resourceInputs["deleteRecursive"] = state ? state.deleteRecursive : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["failedTasks"] = state ? state.failedTasks : undefined;
            resourceInputs["href"] = state ? state.href : undefined;
            resourceInputs["isLocal"] = state ? state.isLocal : undefined;
            resourceInputs["isPublished"] = state ? state.isPublished : undefined;
            resourceInputs["isShared"] = state ? state.isShared : undefined;
            resourceInputs["makeLocalCopy"] = state ? state.makeLocalCopy : undefined;
            resourceInputs["mediaItemLists"] = state ? state.mediaItemLists : undefined;
            resourceInputs["metadata"] = state ? state.metadata : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["numberOfMedia"] = state ? state.numberOfMedia : undefined;
            resourceInputs["numberOfVappTemplates"] = state ? state.numberOfVappTemplates : undefined;
            resourceInputs["org"] = state ? state.org : undefined;
            resourceInputs["ownerName"] = state ? state.ownerName : undefined;
            resourceInputs["publishSubscriptionType"] = state ? state.publishSubscriptionType : undefined;
            resourceInputs["runningTasks"] = state ? state.runningTasks : undefined;
            resourceInputs["storageProfileId"] = state ? state.storageProfileId : undefined;
            resourceInputs["storeTasks"] = state ? state.storeTasks : undefined;
            resourceInputs["subscriptionPassword"] = state ? state.subscriptionPassword : undefined;
            resourceInputs["subscriptionUrl"] = state ? state.subscriptionUrl : undefined;
            resourceInputs["syncAll"] = state ? state.syncAll : undefined;
            resourceInputs["syncAllMediaItems"] = state ? state.syncAllMediaItems : undefined;
            resourceInputs["syncAllVappTemplates"] = state ? state.syncAllVappTemplates : undefined;
            resourceInputs["syncCatalog"] = state ? state.syncCatalog : undefined;
            resourceInputs["syncMediaItems"] = state ? state.syncMediaItems : undefined;
            resourceInputs["syncOnRefresh"] = state ? state.syncOnRefresh : undefined;
            resourceInputs["syncVappTemplates"] = state ? state.syncVappTemplates : undefined;
            resourceInputs["tasksFileName"] = state ? state.tasksFileName : undefined;
            resourceInputs["vappTemplateLists"] = state ? state.vappTemplateLists : undefined;
        } else {
            const args = argsOrState as SubscribedCatalogArgs | undefined;
            if ((!args || args.subscriptionUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subscriptionUrl'");
            }
            resourceInputs["cancelFailedTasks"] = args ? args.cancelFailedTasks : undefined;
            resourceInputs["deleteForce"] = args ? args.deleteForce : undefined;
            resourceInputs["deleteRecursive"] = args ? args.deleteRecursive : undefined;
            resourceInputs["makeLocalCopy"] = args ? args.makeLocalCopy : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["org"] = args ? args.org : undefined;
            resourceInputs["storageProfileId"] = args ? args.storageProfileId : undefined;
            resourceInputs["storeTasks"] = args ? args.storeTasks : undefined;
            resourceInputs["subscriptionPassword"] = args?.subscriptionPassword ? pulumi.secret(args.subscriptionPassword) : undefined;
            resourceInputs["subscriptionUrl"] = args ? args.subscriptionUrl : undefined;
            resourceInputs["syncAll"] = args ? args.syncAll : undefined;
            resourceInputs["syncAllMediaItems"] = args ? args.syncAllMediaItems : undefined;
            resourceInputs["syncAllVappTemplates"] = args ? args.syncAllVappTemplates : undefined;
            resourceInputs["syncCatalog"] = args ? args.syncCatalog : undefined;
            resourceInputs["syncMediaItems"] = args ? args.syncMediaItems : undefined;
            resourceInputs["syncOnRefresh"] = args ? args.syncOnRefresh : undefined;
            resourceInputs["syncVappTemplates"] = args ? args.syncVappTemplates : undefined;
            resourceInputs["catalogVersion"] = undefined /*out*/;
            resourceInputs["created"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["failedTasks"] = undefined /*out*/;
            resourceInputs["href"] = undefined /*out*/;
            resourceInputs["isLocal"] = undefined /*out*/;
            resourceInputs["isPublished"] = undefined /*out*/;
            resourceInputs["isShared"] = undefined /*out*/;
            resourceInputs["mediaItemLists"] = undefined /*out*/;
            resourceInputs["metadata"] = undefined /*out*/;
            resourceInputs["numberOfMedia"] = undefined /*out*/;
            resourceInputs["numberOfVappTemplates"] = undefined /*out*/;
            resourceInputs["ownerName"] = undefined /*out*/;
            resourceInputs["publishSubscriptionType"] = undefined /*out*/;
            resourceInputs["runningTasks"] = undefined /*out*/;
            resourceInputs["tasksFileName"] = undefined /*out*/;
            resourceInputs["vappTemplateLists"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["subscriptionPassword"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(SubscribedCatalog.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering SubscribedCatalog resources.
 */
export interface SubscribedCatalogState {
    /**
     * When `true`, the subscribed catalog will attempt canceling failed tasks.
     */
    cancelFailedTasks?: pulumi.Input<boolean>;
    /**
     * Version number from this catalog. This is inherited from the publishing catalog and updated on sync.
     */
    catalogVersion?: pulumi.Input<number>;
    /**
     * Date and time of catalog creation. This is the creation date of the subscription, not the original published catalog.
     */
    created?: pulumi.Input<string>;
    /**
     * When destroying use `delete_force=true` with `delete_recursive=true` to remove a catalog and any objects it contains, regardless of their state.
     */
    deleteForce?: pulumi.Input<boolean>;
    /**
     * When destroying use `delete_recursive=true` to remove the catalog and any objects it contains that are in a state that normally allows removal.
     */
    deleteRecursive?: pulumi.Input<boolean>;
    /**
     * Description of catalog. This is inherited from the publishing catalog and updated on sync.
     */
    description?: pulumi.Input<string>;
    /**
     * List of synchronization tasks that are have failed. They can refer to the catalog or any of its catalog items.
     */
    failedTasks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * the catalog's Hyper reference.
     */
    href?: pulumi.Input<string>;
    /**
     * (*v3.8.1+*) Indicates if this catalog was created in the current organization.
     */
    isLocal?: pulumi.Input<boolean>;
    /**
     * Indicates if this catalog is available for subscription. (Always false)
     */
    isPublished?: pulumi.Input<boolean>;
    /**
     * Indicates if the catalog is shared.
     */
    isShared?: pulumi.Input<boolean>;
    /**
     * If `true`, subscription to a catalog creates a local copy of all items. Defaults to `false`, which does not create a local copy of catalog items unless a sync operation is performed.
     * It can only be `false` if the user configured in the provider is the System administrator.
     */
    makeLocalCopy?: pulumi.Input<boolean>;
    /**
     * List of media item names in this catalog, in alphabetical order.
     */
    mediaItemLists?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (*Available until VCD 10.5*) Optional metadata of the catalog. This is inherited from the publishing catalog and updated on sync.
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Catalog name
     */
    name?: pulumi.Input<string>;
    /**
     * Number of media items available in this catalog.
     */
    numberOfMedia?: pulumi.Input<number>;
    /**
     * Number of vApp templates available in this catalog.
     */
    numberOfVappTemplates?: pulumi.Input<number>;
    /**
     * The name of organization to use, optional if defined at provider level.
     */
    org?: pulumi.Input<string>;
    /**
     * Owner of the catalog.
     */
    ownerName?: pulumi.Input<string>;
    /**
     * Shows if the catalog is published, if it is a subscription from another one or none of those. (Always `SUBSCRIBED`)
     */
    publishSubscriptionType?: pulumi.Input<string>;
    /**
     * List of running synchronization tasks that are still running. They can refer to the catalog or any of its catalog items.
     */
    runningTasks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Allows to set specific storage profile to be used for catalog.
     */
    storageProfileId?: pulumi.Input<string>;
    /**
     * if `true`, saves the list of tasks to a file for later update.
     */
    storeTasks?: pulumi.Input<boolean>;
    /**
     * An optional password to access the catalog. Only ASCII characters are allowed in a valid password. 
     * The password is only required when set by the publishing catalog. Passing in six asterisks '******' indicates to keep current password.
     * Passing in an empty string indicates to remove password.
     */
    subscriptionPassword?: pulumi.Input<string>;
    /**
     * The URL to subscribe to the external catalog.
     */
    subscriptionUrl?: pulumi.Input<string>;
    /**
     * If `true`, synchronise this catalog and all items.
     */
    syncAll?: pulumi.Input<boolean>;
    /**
     * If `true`, synchronise all media items. Not to be used when `syncAll` is set.
     */
    syncAllMediaItems?: pulumi.Input<boolean>;
    /**
     * If `true`, synchronise all vApp templates. Not to be used when `syncAll` is set.
     */
    syncAllVappTemplates?: pulumi.Input<boolean>;
    /**
     * If `true`, synchronise this catalog. Not to be used when `syncAll` is set. This operation fetches the list of items. If `makeLocalCopy` is set, it also synchronises all the items.
     */
    syncCatalog?: pulumi.Input<boolean>;
    /**
     * Synchronise a list of media items. Not to be used when `syncAll` or `syncAllMediaItems` are set.
     */
    syncMediaItems?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Boolean value that shows if sync should be performed on every refresh.
     */
    syncOnRefresh?: pulumi.Input<boolean>;
    /**
     * Synchronise a list of vApp templates. Not to be used when `syncAll` or `syncAllVappTemplates` are set.
     */
    syncVappTemplates?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Where the running tasks IDs have been stored. Only if `storeTasks` is set.
     */
    tasksFileName?: pulumi.Input<string>;
    /**
     * List of vApp template names in this catalog, in alphabetical order.
     */
    vappTemplateLists?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a SubscribedCatalog resource.
 */
export interface SubscribedCatalogArgs {
    /**
     * When `true`, the subscribed catalog will attempt canceling failed tasks.
     */
    cancelFailedTasks?: pulumi.Input<boolean>;
    /**
     * When destroying use `delete_force=true` with `delete_recursive=true` to remove a catalog and any objects it contains, regardless of their state.
     */
    deleteForce?: pulumi.Input<boolean>;
    /**
     * When destroying use `delete_recursive=true` to remove the catalog and any objects it contains that are in a state that normally allows removal.
     */
    deleteRecursive?: pulumi.Input<boolean>;
    /**
     * If `true`, subscription to a catalog creates a local copy of all items. Defaults to `false`, which does not create a local copy of catalog items unless a sync operation is performed.
     * It can only be `false` if the user configured in the provider is the System administrator.
     */
    makeLocalCopy?: pulumi.Input<boolean>;
    /**
     * Catalog name
     */
    name?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level.
     */
    org?: pulumi.Input<string>;
    /**
     * Allows to set specific storage profile to be used for catalog.
     */
    storageProfileId?: pulumi.Input<string>;
    /**
     * if `true`, saves the list of tasks to a file for later update.
     */
    storeTasks?: pulumi.Input<boolean>;
    /**
     * An optional password to access the catalog. Only ASCII characters are allowed in a valid password. 
     * The password is only required when set by the publishing catalog. Passing in six asterisks '******' indicates to keep current password.
     * Passing in an empty string indicates to remove password.
     */
    subscriptionPassword?: pulumi.Input<string>;
    /**
     * The URL to subscribe to the external catalog.
     */
    subscriptionUrl: pulumi.Input<string>;
    /**
     * If `true`, synchronise this catalog and all items.
     */
    syncAll?: pulumi.Input<boolean>;
    /**
     * If `true`, synchronise all media items. Not to be used when `syncAll` is set.
     */
    syncAllMediaItems?: pulumi.Input<boolean>;
    /**
     * If `true`, synchronise all vApp templates. Not to be used when `syncAll` is set.
     */
    syncAllVappTemplates?: pulumi.Input<boolean>;
    /**
     * If `true`, synchronise this catalog. Not to be used when `syncAll` is set. This operation fetches the list of items. If `makeLocalCopy` is set, it also synchronises all the items.
     */
    syncCatalog?: pulumi.Input<boolean>;
    /**
     * Synchronise a list of media items. Not to be used when `syncAll` or `syncAllMediaItems` are set.
     */
    syncMediaItems?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Boolean value that shows if sync should be performed on every refresh.
     */
    syncOnRefresh?: pulumi.Input<boolean>;
    /**
     * Synchronise a list of vApp templates. Not to be used when `syncAll` or `syncAllVappTemplates` are set.
     */
    syncVappTemplates?: pulumi.Input<pulumi.Input<string>[]>;
}
