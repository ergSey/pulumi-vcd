// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class OrgSaml extends pulumi.CustomResource {
    /**
     * Get an existing OrgSaml resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrgSamlState, opts?: pulumi.CustomResourceOptions): OrgSaml {
        return new OrgSaml(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/orgSaml:OrgSaml';

    /**
     * Returns true if the given object is an instance of OrgSaml.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrgSaml {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrgSaml.__pulumiType;
    }

    /**
     * The name of the SAML attribute that returns the email address of the user
     */
    public readonly email!: pulumi.Output<string | undefined>;
    /**
     * If true, the organization will use SAML for authentication
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * Your service provider entity ID. Once you set this field, it cannot be changed back to empty
     */
    public readonly entityId!: pulumi.Output<string | undefined>;
    /**
     * The name of the SAML attribute that returns the first name of the user
     */
    public readonly firstName!: pulumi.Output<string | undefined>;
    /**
     * The name of the SAML attribute that returns the full name of the user
     */
    public readonly fullName!: pulumi.Output<string | undefined>;
    /**
     * The name of the SAML attribute that returns the identifiers of all the groups of which the user is a member
     */
    public readonly group!: pulumi.Output<string | undefined>;
    /**
     * Name of a file containing the metadata text from a SAML Identity Provider. Required if `identityProviderMetadataText` is not defined
     */
    public readonly identityProviderMetadataFile!: pulumi.Output<string | undefined>;
    /**
     * Text of the metadata text from a SAML Identity Provider. Required if `identityProviderMetadataFile` is not defined
     */
    public readonly identityProviderMetadataText!: pulumi.Output<string | undefined>;
    /**
     * Since there is only one SAML configuration available for an organization, the resource can be identified by the Org itself
     */
    public readonly orgId!: pulumi.Output<string>;
    /**
     * The name of the SAML attribute that returns the identifiers of all roles of the user
     */
    public readonly role!: pulumi.Output<string | undefined>;
    /**
     * The name of the SAML attribute that returns the surname of the user
     */
    public readonly surname!: pulumi.Output<string | undefined>;
    /**
     * The name of the SAML attribute that returns the username of the user
     */
    public readonly userName!: pulumi.Output<string | undefined>;

    /**
     * Create a OrgSaml resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrgSamlArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrgSamlArgs | OrgSamlState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OrgSamlState | undefined;
            resourceInputs["email"] = state ? state.email : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["entityId"] = state ? state.entityId : undefined;
            resourceInputs["firstName"] = state ? state.firstName : undefined;
            resourceInputs["fullName"] = state ? state.fullName : undefined;
            resourceInputs["group"] = state ? state.group : undefined;
            resourceInputs["identityProviderMetadataFile"] = state ? state.identityProviderMetadataFile : undefined;
            resourceInputs["identityProviderMetadataText"] = state ? state.identityProviderMetadataText : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
            resourceInputs["surname"] = state ? state.surname : undefined;
            resourceInputs["userName"] = state ? state.userName : undefined;
        } else {
            const args = argsOrState as OrgSamlArgs | undefined;
            if ((!args || args.enabled === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enabled'");
            }
            if ((!args || args.orgId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'orgId'");
            }
            resourceInputs["email"] = args ? args.email : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["entityId"] = args ? args.entityId : undefined;
            resourceInputs["firstName"] = args ? args.firstName : undefined;
            resourceInputs["fullName"] = args ? args.fullName : undefined;
            resourceInputs["group"] = args ? args.group : undefined;
            resourceInputs["identityProviderMetadataFile"] = args ? args.identityProviderMetadataFile : undefined;
            resourceInputs["identityProviderMetadataText"] = args ? args.identityProviderMetadataText : undefined;
            resourceInputs["orgId"] = args ? args.orgId : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["surname"] = args ? args.surname : undefined;
            resourceInputs["userName"] = args ? args.userName : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OrgSaml.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OrgSaml resources.
 */
export interface OrgSamlState {
    /**
     * The name of the SAML attribute that returns the email address of the user
     */
    email?: pulumi.Input<string>;
    /**
     * If true, the organization will use SAML for authentication
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Your service provider entity ID. Once you set this field, it cannot be changed back to empty
     */
    entityId?: pulumi.Input<string>;
    /**
     * The name of the SAML attribute that returns the first name of the user
     */
    firstName?: pulumi.Input<string>;
    /**
     * The name of the SAML attribute that returns the full name of the user
     */
    fullName?: pulumi.Input<string>;
    /**
     * The name of the SAML attribute that returns the identifiers of all the groups of which the user is a member
     */
    group?: pulumi.Input<string>;
    /**
     * Name of a file containing the metadata text from a SAML Identity Provider. Required if `identityProviderMetadataText` is not defined
     */
    identityProviderMetadataFile?: pulumi.Input<string>;
    /**
     * Text of the metadata text from a SAML Identity Provider. Required if `identityProviderMetadataFile` is not defined
     */
    identityProviderMetadataText?: pulumi.Input<string>;
    /**
     * Since there is only one SAML configuration available for an organization, the resource can be identified by the Org itself
     */
    orgId?: pulumi.Input<string>;
    /**
     * The name of the SAML attribute that returns the identifiers of all roles of the user
     */
    role?: pulumi.Input<string>;
    /**
     * The name of the SAML attribute that returns the surname of the user
     */
    surname?: pulumi.Input<string>;
    /**
     * The name of the SAML attribute that returns the username of the user
     */
    userName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OrgSaml resource.
 */
export interface OrgSamlArgs {
    /**
     * The name of the SAML attribute that returns the email address of the user
     */
    email?: pulumi.Input<string>;
    /**
     * If true, the organization will use SAML for authentication
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Your service provider entity ID. Once you set this field, it cannot be changed back to empty
     */
    entityId?: pulumi.Input<string>;
    /**
     * The name of the SAML attribute that returns the first name of the user
     */
    firstName?: pulumi.Input<string>;
    /**
     * The name of the SAML attribute that returns the full name of the user
     */
    fullName?: pulumi.Input<string>;
    /**
     * The name of the SAML attribute that returns the identifiers of all the groups of which the user is a member
     */
    group?: pulumi.Input<string>;
    /**
     * Name of a file containing the metadata text from a SAML Identity Provider. Required if `identityProviderMetadataText` is not defined
     */
    identityProviderMetadataFile?: pulumi.Input<string>;
    /**
     * Text of the metadata text from a SAML Identity Provider. Required if `identityProviderMetadataFile` is not defined
     */
    identityProviderMetadataText?: pulumi.Input<string>;
    /**
     * Since there is only one SAML configuration available for an organization, the resource can be identified by the Org itself
     */
    orgId: pulumi.Input<string>;
    /**
     * The name of the SAML attribute that returns the identifiers of all roles of the user
     */
    role?: pulumi.Input<string>;
    /**
     * The name of the SAML attribute that returns the surname of the user
     */
    surname?: pulumi.Input<string>;
    /**
     * The name of the SAML attribute that returns the username of the user
     */
    userName?: pulumi.Input<string>;
}
