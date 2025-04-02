// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class ServiceAccount extends pulumi.CustomResource {
    /**
     * Get an existing ServiceAccount resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: ServiceAccountState, opts?: pulumi.CustomResourceOptions): ServiceAccount {
        return new ServiceAccount(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/serviceAccount:ServiceAccount';

    /**
     * Returns true if the given object is an instance of ServiceAccount.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceAccount {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceAccount.__pulumiType;
    }

    /**
     * Status of the Service Account. Can be set to `false` and back to `true` if
     * the access token was lost to get a new one.
     */
    public readonly active!: pulumi.Output<boolean | undefined>;
    /**
     * If set to false, will output a warning about the service account file
     * containing sensitive information.
     */
    public readonly allowTokenFile!: pulumi.Output<boolean | undefined>;
    /**
     * Required only when `active` is set to `true`. Contains the access token
     * that can be used for authenticating to VCD.
     */
    public readonly fileName!: pulumi.Output<string | undefined>;
    /**
     * A unique name for the Service Account in an organisation.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful
     * when connected as sysadmin working across different organisations.
     */
    public readonly org!: pulumi.Output<string | undefined>;
    /**
     * Role ID of service account
     */
    public readonly roleId!: pulumi.Output<string>;
    /**
     * UUID of the Service Account.
     */
    public readonly softwareId!: pulumi.Output<string>;
    /**
     * Version of the service using the Service Account
     */
    public readonly softwareVersion!: pulumi.Output<string | undefined>;
    /**
     * URI of the service using the Service Account
     */
    public readonly uri!: pulumi.Output<string | undefined>;

    /**
     * Create a ServiceAccount resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceAccountArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: ServiceAccountArgs | ServiceAccountState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as ServiceAccountState | undefined;
            resourceInputs["active"] = state ? state.active : undefined;
            resourceInputs["allowTokenFile"] = state ? state.allowTokenFile : undefined;
            resourceInputs["fileName"] = state ? state.fileName : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["org"] = state ? state.org : undefined;
            resourceInputs["roleId"] = state ? state.roleId : undefined;
            resourceInputs["softwareId"] = state ? state.softwareId : undefined;
            resourceInputs["softwareVersion"] = state ? state.softwareVersion : undefined;
            resourceInputs["uri"] = state ? state.uri : undefined;
        } else {
            const args = argsOrState as ServiceAccountArgs | undefined;
            if ((!args || args.roleId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'roleId'");
            }
            if ((!args || args.softwareId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'softwareId'");
            }
            resourceInputs["active"] = args ? args.active : undefined;
            resourceInputs["allowTokenFile"] = args ? args.allowTokenFile : undefined;
            resourceInputs["fileName"] = args ? args.fileName : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["org"] = args ? args.org : undefined;
            resourceInputs["roleId"] = args ? args.roleId : undefined;
            resourceInputs["softwareId"] = args ? args.softwareId : undefined;
            resourceInputs["softwareVersion"] = args ? args.softwareVersion : undefined;
            resourceInputs["uri"] = args ? args.uri : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ServiceAccount.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering ServiceAccount resources.
 */
export interface ServiceAccountState {
    /**
     * Status of the Service Account. Can be set to `false` and back to `true` if
     * the access token was lost to get a new one.
     */
    active?: pulumi.Input<boolean>;
    /**
     * If set to false, will output a warning about the service account file
     * containing sensitive information.
     */
    allowTokenFile?: pulumi.Input<boolean>;
    /**
     * Required only when `active` is set to `true`. Contains the access token
     * that can be used for authenticating to VCD.
     */
    fileName?: pulumi.Input<string>;
    /**
     * A unique name for the Service Account in an organisation.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful
     * when connected as sysadmin working across different organisations.
     */
    org?: pulumi.Input<string>;
    /**
     * Role ID of service account
     */
    roleId?: pulumi.Input<string>;
    /**
     * UUID of the Service Account.
     */
    softwareId?: pulumi.Input<string>;
    /**
     * Version of the service using the Service Account
     */
    softwareVersion?: pulumi.Input<string>;
    /**
     * URI of the service using the Service Account
     */
    uri?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a ServiceAccount resource.
 */
export interface ServiceAccountArgs {
    /**
     * Status of the Service Account. Can be set to `false` and back to `true` if
     * the access token was lost to get a new one.
     */
    active?: pulumi.Input<boolean>;
    /**
     * If set to false, will output a warning about the service account file
     * containing sensitive information.
     */
    allowTokenFile?: pulumi.Input<boolean>;
    /**
     * Required only when `active` is set to `true`. Contains the access token
     * that can be used for authenticating to VCD.
     */
    fileName?: pulumi.Input<string>;
    /**
     * A unique name for the Service Account in an organisation.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful
     * when connected as sysadmin working across different organisations.
     */
    org?: pulumi.Input<string>;
    /**
     * Role ID of service account
     */
    roleId: pulumi.Input<string>;
    /**
     * UUID of the Service Account.
     */
    softwareId: pulumi.Input<string>;
    /**
     * Version of the service using the Service Account
     */
    softwareVersion?: pulumi.Input<string>;
    /**
     * URI of the service using the Service Account
     */
    uri?: pulumi.Input<string>;
}
