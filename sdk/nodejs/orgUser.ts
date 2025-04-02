// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class OrgUser extends pulumi.CustomResource {
    /**
     * Get an existing OrgUser resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrgUserState, opts?: pulumi.CustomResourceOptions): OrgUser {
        return new OrgUser(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/orgUser:OrgUser';

    /**
     * Returns true if the given object is an instance of OrgUser.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrgUser {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrgUser.__pulumiType;
    }

    /**
     * Quota of vApps that this user can deploy. A value of 0 specifies an unlimited quota.
     * The default is 0.
     */
    public readonly deployedVmQuota!: pulumi.Output<number>;
    /**
     * An optional description of the user.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The Org User email address. Needs to be a properly formatted email address.
     */
    public readonly emailAddress!: pulumi.Output<string>;
    /**
     * True if the user is enabled and can log in. The default is `true`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * The full name of the user.
     */
    public readonly fullName!: pulumi.Output<string>;
    /**
     * The set of group names to which this user belongs. It's only populated if the users
     * are created after the group (with this user having a `dependsOn` of the given group).
     */
    public /*out*/ readonly groupNames!: pulumi.Output<string[]>;
    /**
     * The Org User instant messaging.
     */
    public readonly instantMessaging!: pulumi.Output<string>;
    /**
     * If the user account is going to be imported from an external resource, like an LDAP.
     * In this case, `password` nor `passwordFile` are not required. Defaults to `false`.
     */
    public readonly isExternal!: pulumi.Output<boolean | undefined>;
    /**
     * True if this user has a group role. The default is `false`.
     */
    public readonly isGroupRole!: pulumi.Output<boolean | undefined>;
    /**
     * If the user account has been locked due to too many invalid login attempts, the value will 
     * change to true (only the system can lock the user). To unlock the user re-set this flag to false.
     */
    public readonly isLocked!: pulumi.Output<boolean | undefined>;
    /**
     * A unique name for the user.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of organization to which the user belongs. Optional if defined at provider level. If we 
     * want to create a user at provider level, use "System" as org name.
     */
    public readonly org!: pulumi.Output<string | undefined>;
    /**
     * The user's password. This value is never returned on read. Either "password" or "passwordFile" must be included on
     * creation unless isExternal is true.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * Name of a file containing the user's password. Either "passwordFile" or "password" must be included on creation unless
     * isExternal is true.
     */
    public readonly passwordFile!: pulumi.Output<string | undefined>;
    /**
     * Identity provider type for this user. One of: `INTEGRATED`, `SAML`, `OAUTH`. The default
     * is `INTEGRATED`.
     */
    public readonly providerType!: pulumi.Output<string | undefined>;
    /**
     * The role of the user. Role names can be retrieved from the organization. Both built-in roles and
     * custom built can be used. The roles normally available are:
     * * `Organization Administrator`
     * * `Catalog Author`
     * * `vApp Author`
     * * `vApp User`
     * * `Console Access Only`
     * * `Defer to Identity Provider`
     */
    public readonly role!: pulumi.Output<string>;
    /**
     * Quota of vApps that this user can store. A value of 0 specifies an unlimited quota.
     * The default is 0.
     */
    public readonly storedVmQuota!: pulumi.Output<number>;
    /**
     * Take ownership of user's objects on deletion.
     */
    public readonly takeOwnership!: pulumi.Output<boolean | undefined>;
    /**
     * The Org User telephone number.
     */
    public readonly telephone!: pulumi.Output<string>;

    /**
     * Create a OrgUser resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrgUserArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrgUserArgs | OrgUserState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OrgUserState | undefined;
            resourceInputs["deployedVmQuota"] = state ? state.deployedVmQuota : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["emailAddress"] = state ? state.emailAddress : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["fullName"] = state ? state.fullName : undefined;
            resourceInputs["groupNames"] = state ? state.groupNames : undefined;
            resourceInputs["instantMessaging"] = state ? state.instantMessaging : undefined;
            resourceInputs["isExternal"] = state ? state.isExternal : undefined;
            resourceInputs["isGroupRole"] = state ? state.isGroupRole : undefined;
            resourceInputs["isLocked"] = state ? state.isLocked : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["org"] = state ? state.org : undefined;
            resourceInputs["password"] = state ? state.password : undefined;
            resourceInputs["passwordFile"] = state ? state.passwordFile : undefined;
            resourceInputs["providerType"] = state ? state.providerType : undefined;
            resourceInputs["role"] = state ? state.role : undefined;
            resourceInputs["storedVmQuota"] = state ? state.storedVmQuota : undefined;
            resourceInputs["takeOwnership"] = state ? state.takeOwnership : undefined;
            resourceInputs["telephone"] = state ? state.telephone : undefined;
        } else {
            const args = argsOrState as OrgUserArgs | undefined;
            if ((!args || args.role === undefined) && !opts.urn) {
                throw new Error("Missing required property 'role'");
            }
            resourceInputs["deployedVmQuota"] = args ? args.deployedVmQuota : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["emailAddress"] = args ? args.emailAddress : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["fullName"] = args ? args.fullName : undefined;
            resourceInputs["instantMessaging"] = args ? args.instantMessaging : undefined;
            resourceInputs["isExternal"] = args ? args.isExternal : undefined;
            resourceInputs["isGroupRole"] = args ? args.isGroupRole : undefined;
            resourceInputs["isLocked"] = args ? args.isLocked : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["org"] = args ? args.org : undefined;
            resourceInputs["password"] = args?.password ? pulumi.secret(args.password) : undefined;
            resourceInputs["passwordFile"] = args ? args.passwordFile : undefined;
            resourceInputs["providerType"] = args ? args.providerType : undefined;
            resourceInputs["role"] = args ? args.role : undefined;
            resourceInputs["storedVmQuota"] = args ? args.storedVmQuota : undefined;
            resourceInputs["takeOwnership"] = args ? args.takeOwnership : undefined;
            resourceInputs["telephone"] = args ? args.telephone : undefined;
            resourceInputs["groupNames"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["password"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(OrgUser.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OrgUser resources.
 */
export interface OrgUserState {
    /**
     * Quota of vApps that this user can deploy. A value of 0 specifies an unlimited quota.
     * The default is 0.
     */
    deployedVmQuota?: pulumi.Input<number>;
    /**
     * An optional description of the user.
     */
    description?: pulumi.Input<string>;
    /**
     * The Org User email address. Needs to be a properly formatted email address.
     */
    emailAddress?: pulumi.Input<string>;
    /**
     * True if the user is enabled and can log in. The default is `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The full name of the user.
     */
    fullName?: pulumi.Input<string>;
    /**
     * The set of group names to which this user belongs. It's only populated if the users
     * are created after the group (with this user having a `dependsOn` of the given group).
     */
    groupNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The Org User instant messaging.
     */
    instantMessaging?: pulumi.Input<string>;
    /**
     * If the user account is going to be imported from an external resource, like an LDAP.
     * In this case, `password` nor `passwordFile` are not required. Defaults to `false`.
     */
    isExternal?: pulumi.Input<boolean>;
    /**
     * True if this user has a group role. The default is `false`.
     */
    isGroupRole?: pulumi.Input<boolean>;
    /**
     * If the user account has been locked due to too many invalid login attempts, the value will 
     * change to true (only the system can lock the user). To unlock the user re-set this flag to false.
     */
    isLocked?: pulumi.Input<boolean>;
    /**
     * A unique name for the user.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of organization to which the user belongs. Optional if defined at provider level. If we 
     * want to create a user at provider level, use "System" as org name.
     */
    org?: pulumi.Input<string>;
    /**
     * The user's password. This value is never returned on read. Either "password" or "passwordFile" must be included on
     * creation unless isExternal is true.
     */
    password?: pulumi.Input<string>;
    /**
     * Name of a file containing the user's password. Either "passwordFile" or "password" must be included on creation unless
     * isExternal is true.
     */
    passwordFile?: pulumi.Input<string>;
    /**
     * Identity provider type for this user. One of: `INTEGRATED`, `SAML`, `OAUTH`. The default
     * is `INTEGRATED`.
     */
    providerType?: pulumi.Input<string>;
    /**
     * The role of the user. Role names can be retrieved from the organization. Both built-in roles and
     * custom built can be used. The roles normally available are:
     * * `Organization Administrator`
     * * `Catalog Author`
     * * `vApp Author`
     * * `vApp User`
     * * `Console Access Only`
     * * `Defer to Identity Provider`
     */
    role?: pulumi.Input<string>;
    /**
     * Quota of vApps that this user can store. A value of 0 specifies an unlimited quota.
     * The default is 0.
     */
    storedVmQuota?: pulumi.Input<number>;
    /**
     * Take ownership of user's objects on deletion.
     */
    takeOwnership?: pulumi.Input<boolean>;
    /**
     * The Org User telephone number.
     */
    telephone?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OrgUser resource.
 */
export interface OrgUserArgs {
    /**
     * Quota of vApps that this user can deploy. A value of 0 specifies an unlimited quota.
     * The default is 0.
     */
    deployedVmQuota?: pulumi.Input<number>;
    /**
     * An optional description of the user.
     */
    description?: pulumi.Input<string>;
    /**
     * The Org User email address. Needs to be a properly formatted email address.
     */
    emailAddress?: pulumi.Input<string>;
    /**
     * True if the user is enabled and can log in. The default is `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The full name of the user.
     */
    fullName?: pulumi.Input<string>;
    /**
     * The Org User instant messaging.
     */
    instantMessaging?: pulumi.Input<string>;
    /**
     * If the user account is going to be imported from an external resource, like an LDAP.
     * In this case, `password` nor `passwordFile` are not required. Defaults to `false`.
     */
    isExternal?: pulumi.Input<boolean>;
    /**
     * True if this user has a group role. The default is `false`.
     */
    isGroupRole?: pulumi.Input<boolean>;
    /**
     * If the user account has been locked due to too many invalid login attempts, the value will 
     * change to true (only the system can lock the user). To unlock the user re-set this flag to false.
     */
    isLocked?: pulumi.Input<boolean>;
    /**
     * A unique name for the user.
     */
    name?: pulumi.Input<string>;
    /**
     * The name of organization to which the user belongs. Optional if defined at provider level. If we 
     * want to create a user at provider level, use "System" as org name.
     */
    org?: pulumi.Input<string>;
    /**
     * The user's password. This value is never returned on read. Either "password" or "passwordFile" must be included on
     * creation unless isExternal is true.
     */
    password?: pulumi.Input<string>;
    /**
     * Name of a file containing the user's password. Either "passwordFile" or "password" must be included on creation unless
     * isExternal is true.
     */
    passwordFile?: pulumi.Input<string>;
    /**
     * Identity provider type for this user. One of: `INTEGRATED`, `SAML`, `OAUTH`. The default
     * is `INTEGRATED`.
     */
    providerType?: pulumi.Input<string>;
    /**
     * The role of the user. Role names can be retrieved from the organization. Both built-in roles and
     * custom built can be used. The roles normally available are:
     * * `Organization Administrator`
     * * `Catalog Author`
     * * `vApp Author`
     * * `vApp User`
     * * `Console Access Only`
     * * `Defer to Identity Provider`
     */
    role: pulumi.Input<string>;
    /**
     * Quota of vApps that this user can store. A value of 0 specifies an unlimited quota.
     * The default is 0.
     */
    storedVmQuota?: pulumi.Input<number>;
    /**
     * Take ownership of user's objects on deletion.
     */
    takeOwnership?: pulumi.Input<boolean>;
    /**
     * The Org User telephone number.
     */
    telephone?: pulumi.Input<string>;
}
