// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class IpSpaceUplink extends pulumi.CustomResource {
    /**
     * Get an existing IpSpaceUplink resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IpSpaceUplinkState, opts?: pulumi.CustomResourceOptions): IpSpaceUplink {
        return new IpSpaceUplink(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/ipSpaceUplink:IpSpaceUplink';

    /**
     * Returns true if the given object is an instance of IpSpaceUplink.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IpSpaceUplink {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IpSpaceUplink.__pulumiType;
    }

    /**
     * A set of Tier-0 Router Interfaces to associate with this uplink
     */
    public readonly associatedInterfaceIds!: pulumi.Output<string[] | undefined>;
    /**
     * An optional description for IP Space Uplink
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * External Network ID For IP Space Uplink configuration
     */
    public readonly externalNetworkId!: pulumi.Output<string>;
    /**
     * IP Space ID configuration
     */
    public readonly ipSpaceId!: pulumi.Output<string>;
    /**
     * Backing IP Space type
     */
    public /*out*/ readonly ipSpaceType!: pulumi.Output<string>;
    /**
     * A tenant facing name for IP Space Uplink
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Status of IP Space Uplink
     */
    public /*out*/ readonly status!: pulumi.Output<string>;

    /**
     * Create a IpSpaceUplink resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IpSpaceUplinkArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IpSpaceUplinkArgs | IpSpaceUplinkState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IpSpaceUplinkState | undefined;
            resourceInputs["associatedInterfaceIds"] = state ? state.associatedInterfaceIds : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["externalNetworkId"] = state ? state.externalNetworkId : undefined;
            resourceInputs["ipSpaceId"] = state ? state.ipSpaceId : undefined;
            resourceInputs["ipSpaceType"] = state ? state.ipSpaceType : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
        } else {
            const args = argsOrState as IpSpaceUplinkArgs | undefined;
            if ((!args || args.externalNetworkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'externalNetworkId'");
            }
            if ((!args || args.ipSpaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ipSpaceId'");
            }
            resourceInputs["associatedInterfaceIds"] = args ? args.associatedInterfaceIds : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["externalNetworkId"] = args ? args.externalNetworkId : undefined;
            resourceInputs["ipSpaceId"] = args ? args.ipSpaceId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["ipSpaceType"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IpSpaceUplink.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IpSpaceUplink resources.
 */
export interface IpSpaceUplinkState {
    /**
     * A set of Tier-0 Router Interfaces to associate with this uplink
     */
    associatedInterfaceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An optional description for IP Space Uplink
     */
    description?: pulumi.Input<string>;
    /**
     * External Network ID For IP Space Uplink configuration
     */
    externalNetworkId?: pulumi.Input<string>;
    /**
     * IP Space ID configuration
     */
    ipSpaceId?: pulumi.Input<string>;
    /**
     * Backing IP Space type
     */
    ipSpaceType?: pulumi.Input<string>;
    /**
     * A tenant facing name for IP Space Uplink
     */
    name?: pulumi.Input<string>;
    /**
     * Status of IP Space Uplink
     */
    status?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IpSpaceUplink resource.
 */
export interface IpSpaceUplinkArgs {
    /**
     * A set of Tier-0 Router Interfaces to associate with this uplink
     */
    associatedInterfaceIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * An optional description for IP Space Uplink
     */
    description?: pulumi.Input<string>;
    /**
     * External Network ID For IP Space Uplink configuration
     */
    externalNetworkId: pulumi.Input<string>;
    /**
     * IP Space ID configuration
     */
    ipSpaceId: pulumi.Input<string>;
    /**
     * A tenant facing name for IP Space Uplink
     */
    name?: pulumi.Input<string>;
}
