// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

export class NetworkPool extends pulumi.CustomResource {
    /**
     * Get an existing NetworkPool resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkPoolState, opts?: pulumi.CustomResourceOptions): NetworkPool {
        return new NetworkPool(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/networkPool:NetworkPool';

    /**
     * Returns true if the given object is an instance of NetworkPool.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkPool {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkPool.__pulumiType;
    }

    /**
     * The components used by the network pool. See Backing below for details
     */
    public readonly backing!: pulumi.Output<outputs.NetworkPoolBacking>;
    /**
     * Define how the backing components are considered. It should be one of the following:
     * * `use-explicit-name` (Default) The backing components must be named explicitly;
     * * `use-when-only-one` The automatically selected backing component will be used if there is only one available;
     * * `use-first-available` Use the first available backing component.
     */
    public readonly backingSelectionConstraint!: pulumi.Output<string | undefined>;
    /**
     * Description of the network pool
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Unique name of network pool
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Id of the network provider (either vCenter or NSX-T manager)
     */
    public readonly networkProviderId!: pulumi.Output<string>;
    /**
     * Name of the network provider
     */
    public /*out*/ readonly networkProviderName!: pulumi.Output<string>;
    /**
     * Type of network provider
     */
    public /*out*/ readonly networkProviderType!: pulumi.Output<string>;
    /**
     * Whether the network pool is in promiscuous mode
     */
    public /*out*/ readonly promiscuousMode!: pulumi.Output<boolean>;
    /**
     * Status of the network pool
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Total number of backings
     */
    public /*out*/ readonly totalBackingsCount!: pulumi.Output<number>;
    /**
     * Type of the network pool (one of `GENEVE`, `VLAN`, `PORTGROUP_BACKED`)
     */
    public readonly type!: pulumi.Output<string>;
    /**
     * Number of used backings
     */
    public /*out*/ readonly usedBackingsCount!: pulumi.Output<number>;

    /**
     * Create a NetworkPool resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkPoolArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkPoolArgs | NetworkPoolState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetworkPoolState | undefined;
            resourceInputs["backing"] = state ? state.backing : undefined;
            resourceInputs["backingSelectionConstraint"] = state ? state.backingSelectionConstraint : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkProviderId"] = state ? state.networkProviderId : undefined;
            resourceInputs["networkProviderName"] = state ? state.networkProviderName : undefined;
            resourceInputs["networkProviderType"] = state ? state.networkProviderType : undefined;
            resourceInputs["promiscuousMode"] = state ? state.promiscuousMode : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["totalBackingsCount"] = state ? state.totalBackingsCount : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["usedBackingsCount"] = state ? state.usedBackingsCount : undefined;
        } else {
            const args = argsOrState as NetworkPoolArgs | undefined;
            if ((!args || args.networkProviderId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkProviderId'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["backing"] = args ? args.backing : undefined;
            resourceInputs["backingSelectionConstraint"] = args ? args.backingSelectionConstraint : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkProviderId"] = args ? args.networkProviderId : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["networkProviderName"] = undefined /*out*/;
            resourceInputs["networkProviderType"] = undefined /*out*/;
            resourceInputs["promiscuousMode"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["totalBackingsCount"] = undefined /*out*/;
            resourceInputs["usedBackingsCount"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NetworkPool.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkPool resources.
 */
export interface NetworkPoolState {
    /**
     * The components used by the network pool. See Backing below for details
     */
    backing?: pulumi.Input<inputs.NetworkPoolBacking>;
    /**
     * Define how the backing components are considered. It should be one of the following:
     * * `use-explicit-name` (Default) The backing components must be named explicitly;
     * * `use-when-only-one` The automatically selected backing component will be used if there is only one available;
     * * `use-first-available` Use the first available backing component.
     */
    backingSelectionConstraint?: pulumi.Input<string>;
    /**
     * Description of the network pool
     */
    description?: pulumi.Input<string>;
    /**
     * Unique name of network pool
     */
    name?: pulumi.Input<string>;
    /**
     * Id of the network provider (either vCenter or NSX-T manager)
     */
    networkProviderId?: pulumi.Input<string>;
    /**
     * Name of the network provider
     */
    networkProviderName?: pulumi.Input<string>;
    /**
     * Type of network provider
     */
    networkProviderType?: pulumi.Input<string>;
    /**
     * Whether the network pool is in promiscuous mode
     */
    promiscuousMode?: pulumi.Input<boolean>;
    /**
     * Status of the network pool
     */
    status?: pulumi.Input<string>;
    /**
     * Total number of backings
     */
    totalBackingsCount?: pulumi.Input<number>;
    /**
     * Type of the network pool (one of `GENEVE`, `VLAN`, `PORTGROUP_BACKED`)
     */
    type?: pulumi.Input<string>;
    /**
     * Number of used backings
     */
    usedBackingsCount?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a NetworkPool resource.
 */
export interface NetworkPoolArgs {
    /**
     * The components used by the network pool. See Backing below for details
     */
    backing?: pulumi.Input<inputs.NetworkPoolBacking>;
    /**
     * Define how the backing components are considered. It should be one of the following:
     * * `use-explicit-name` (Default) The backing components must be named explicitly;
     * * `use-when-only-one` The automatically selected backing component will be used if there is only one available;
     * * `use-first-available` Use the first available backing component.
     */
    backingSelectionConstraint?: pulumi.Input<string>;
    /**
     * Description of the network pool
     */
    description?: pulumi.Input<string>;
    /**
     * Unique name of network pool
     */
    name?: pulumi.Input<string>;
    /**
     * Id of the network provider (either vCenter or NSX-T manager)
     */
    networkProviderId: pulumi.Input<string>;
    /**
     * Type of the network pool (one of `GENEVE`, `VLAN`, `PORTGROUP_BACKED`)
     */
    type: pulumi.Input<string>;
}
