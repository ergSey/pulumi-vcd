// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

export class NetworkDirect extends pulumi.CustomResource {
    /**
     * Get an existing NetworkDirect resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkDirectState, opts?: pulumi.CustomResourceOptions): NetworkDirect {
        return new NetworkDirect(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/networkDirect:NetworkDirect';

    /**
     * Returns true if the given object is an instance of NetworkDirect.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkDirect {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkDirect.__pulumiType;
    }

    /**
     * An optional description of the network
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the external network.
     */
    public readonly externalNetwork!: pulumi.Output<string>;
    /**
     * (Computed) returns the first DNS from the external network
     */
    public /*out*/ readonly externalNetworkDns1!: pulumi.Output<string>;
    /**
     * (Computed) returns the second DNS from the external network
     */
    public /*out*/ readonly externalNetworkDns2!: pulumi.Output<string>;
    /**
     * (Computed) returns the DNS suffix from the external network
     */
    public /*out*/ readonly externalNetworkDnsSuffix!: pulumi.Output<string>;
    /**
     * (Computed) returns the gateway from the external network
     */
    public /*out*/ readonly externalNetworkGateway!: pulumi.Output<string>;
    /**
     * (Computed) returns the netmask from the external network
     */
    public /*out*/ readonly externalNetworkNetmask!: pulumi.Output<string>;
    /**
     * Network Hypertext Reference
     */
    public /*out*/ readonly href!: pulumi.Output<string>;
    /**
     * Use `metadataEntry` instead. Key value map of metadata to assign to this network.
     *
     * @deprecated Use metadataEntry instead
     */
    public readonly metadata!: pulumi.Output<{[key: string]: string}>;
    /**
     * A set of metadata entries to assign. See Metadata section for details.
     */
    public readonly metadataEntries!: pulumi.Output<outputs.NetworkDirectMetadataEntry[]>;
    /**
     * A unique name for the network
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when
     * connected as sysadmin working across different organisations
     */
    public readonly org!: pulumi.Output<string | undefined>;
    /**
     * Defines if this network is shared between multiple VDCs
     * in the Org.  Defaults to `false`.
     */
    public readonly shared!: pulumi.Output<boolean | undefined>;
    /**
     * The name of VDC to use, optional if defined at provider level
     */
    public readonly vdc!: pulumi.Output<string | undefined>;

    /**
     * Create a NetworkDirect resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NetworkDirectArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkDirectArgs | NetworkDirectState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetworkDirectState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["externalNetwork"] = state ? state.externalNetwork : undefined;
            resourceInputs["externalNetworkDns1"] = state ? state.externalNetworkDns1 : undefined;
            resourceInputs["externalNetworkDns2"] = state ? state.externalNetworkDns2 : undefined;
            resourceInputs["externalNetworkDnsSuffix"] = state ? state.externalNetworkDnsSuffix : undefined;
            resourceInputs["externalNetworkGateway"] = state ? state.externalNetworkGateway : undefined;
            resourceInputs["externalNetworkNetmask"] = state ? state.externalNetworkNetmask : undefined;
            resourceInputs["href"] = state ? state.href : undefined;
            resourceInputs["metadata"] = state ? state.metadata : undefined;
            resourceInputs["metadataEntries"] = state ? state.metadataEntries : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["org"] = state ? state.org : undefined;
            resourceInputs["shared"] = state ? state.shared : undefined;
            resourceInputs["vdc"] = state ? state.vdc : undefined;
        } else {
            const args = argsOrState as NetworkDirectArgs | undefined;
            if ((!args || args.externalNetwork === undefined) && !opts.urn) {
                throw new Error("Missing required property 'externalNetwork'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["externalNetwork"] = args ? args.externalNetwork : undefined;
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["metadataEntries"] = args ? args.metadataEntries : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["org"] = args ? args.org : undefined;
            resourceInputs["shared"] = args ? args.shared : undefined;
            resourceInputs["vdc"] = args ? args.vdc : undefined;
            resourceInputs["externalNetworkDns1"] = undefined /*out*/;
            resourceInputs["externalNetworkDns2"] = undefined /*out*/;
            resourceInputs["externalNetworkDnsSuffix"] = undefined /*out*/;
            resourceInputs["externalNetworkGateway"] = undefined /*out*/;
            resourceInputs["externalNetworkNetmask"] = undefined /*out*/;
            resourceInputs["href"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NetworkDirect.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkDirect resources.
 */
export interface NetworkDirectState {
    /**
     * An optional description of the network
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the external network.
     */
    externalNetwork?: pulumi.Input<string>;
    /**
     * (Computed) returns the first DNS from the external network
     */
    externalNetworkDns1?: pulumi.Input<string>;
    /**
     * (Computed) returns the second DNS from the external network
     */
    externalNetworkDns2?: pulumi.Input<string>;
    /**
     * (Computed) returns the DNS suffix from the external network
     */
    externalNetworkDnsSuffix?: pulumi.Input<string>;
    /**
     * (Computed) returns the gateway from the external network
     */
    externalNetworkGateway?: pulumi.Input<string>;
    /**
     * (Computed) returns the netmask from the external network
     */
    externalNetworkNetmask?: pulumi.Input<string>;
    /**
     * Network Hypertext Reference
     */
    href?: pulumi.Input<string>;
    /**
     * Use `metadataEntry` instead. Key value map of metadata to assign to this network.
     *
     * @deprecated Use metadataEntry instead
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A set of metadata entries to assign. See Metadata section for details.
     */
    metadataEntries?: pulumi.Input<pulumi.Input<inputs.NetworkDirectMetadataEntry>[]>;
    /**
     * A unique name for the network
     */
    name?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when
     * connected as sysadmin working across different organisations
     */
    org?: pulumi.Input<string>;
    /**
     * Defines if this network is shared between multiple VDCs
     * in the Org.  Defaults to `false`.
     */
    shared?: pulumi.Input<boolean>;
    /**
     * The name of VDC to use, optional if defined at provider level
     */
    vdc?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NetworkDirect resource.
 */
export interface NetworkDirectArgs {
    /**
     * An optional description of the network
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the external network.
     */
    externalNetwork: pulumi.Input<string>;
    /**
     * Use `metadataEntry` instead. Key value map of metadata to assign to this network.
     *
     * @deprecated Use metadataEntry instead
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A set of metadata entries to assign. See Metadata section for details.
     */
    metadataEntries?: pulumi.Input<pulumi.Input<inputs.NetworkDirectMetadataEntry>[]>;
    /**
     * A unique name for the network
     */
    name?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when
     * connected as sysadmin working across different organisations
     */
    org?: pulumi.Input<string>;
    /**
     * Defines if this network is shared between multiple VDCs
     * in the Org.  Defaults to `false`.
     */
    shared?: pulumi.Input<boolean>;
    /**
     * The name of VDC to use, optional if defined at provider level
     */
    vdc?: pulumi.Input<string>;
}
