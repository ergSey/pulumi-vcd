// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

export class NetworkIsolated extends pulumi.CustomResource {
    /**
     * Get an existing NetworkIsolated resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NetworkIsolatedState, opts?: pulumi.CustomResourceOptions): NetworkIsolated {
        return new NetworkIsolated(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/networkIsolated:NetworkIsolated';

    /**
     * Returns true if the given object is an instance of NetworkIsolated.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NetworkIsolated {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NetworkIsolated.__pulumiType;
    }

    /**
     * An optional description of the network
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A range of IPs to issue to virtual machines that don't
     * have a static IP; see IP Pools below for details.
     */
    public readonly dhcpPools!: pulumi.Output<outputs.NetworkIsolatedDhcpPool[] | undefined>;
    /**
     * First DNS server to use.
     */
    public readonly dns1!: pulumi.Output<string | undefined>;
    /**
     * Second DNS server to use.
     */
    public readonly dns2!: pulumi.Output<string | undefined>;
    /**
     * A FQDN for the virtual machines on this network
     */
    public readonly dnsSuffix!: pulumi.Output<string | undefined>;
    /**
     * The gateway for this network
     */
    public readonly gateway!: pulumi.Output<string | undefined>;
    /**
     * Network Hyper Reference
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
     *
     * <a id="ip-pools"></a>
     */
    public readonly metadataEntries!: pulumi.Output<outputs.NetworkIsolatedMetadataEntry[]>;
    /**
     * A unique name for the network
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The netmask for the new network. Defaults to `255.255.255.0`
     */
    public readonly netmask!: pulumi.Output<string | undefined>;
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
     * A range of IPs permitted to be used as static IPs for
     * virtual machines; see IP Pools below for details.
     */
    public readonly staticIpPools!: pulumi.Output<outputs.NetworkIsolatedStaticIpPool[] | undefined>;
    /**
     * The name of VDC to use, optional if defined at provider level
     */
    public readonly vdc!: pulumi.Output<string | undefined>;

    /**
     * Create a NetworkIsolated resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: NetworkIsolatedArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NetworkIsolatedArgs | NetworkIsolatedState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NetworkIsolatedState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["dhcpPools"] = state ? state.dhcpPools : undefined;
            resourceInputs["dns1"] = state ? state.dns1 : undefined;
            resourceInputs["dns2"] = state ? state.dns2 : undefined;
            resourceInputs["dnsSuffix"] = state ? state.dnsSuffix : undefined;
            resourceInputs["gateway"] = state ? state.gateway : undefined;
            resourceInputs["href"] = state ? state.href : undefined;
            resourceInputs["metadata"] = state ? state.metadata : undefined;
            resourceInputs["metadataEntries"] = state ? state.metadataEntries : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["netmask"] = state ? state.netmask : undefined;
            resourceInputs["org"] = state ? state.org : undefined;
            resourceInputs["shared"] = state ? state.shared : undefined;
            resourceInputs["staticIpPools"] = state ? state.staticIpPools : undefined;
            resourceInputs["vdc"] = state ? state.vdc : undefined;
        } else {
            const args = argsOrState as NetworkIsolatedArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dhcpPools"] = args ? args.dhcpPools : undefined;
            resourceInputs["dns1"] = args ? args.dns1 : undefined;
            resourceInputs["dns2"] = args ? args.dns2 : undefined;
            resourceInputs["dnsSuffix"] = args ? args.dnsSuffix : undefined;
            resourceInputs["gateway"] = args ? args.gateway : undefined;
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["metadataEntries"] = args ? args.metadataEntries : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["netmask"] = args ? args.netmask : undefined;
            resourceInputs["org"] = args ? args.org : undefined;
            resourceInputs["shared"] = args ? args.shared : undefined;
            resourceInputs["staticIpPools"] = args ? args.staticIpPools : undefined;
            resourceInputs["vdc"] = args ? args.vdc : undefined;
            resourceInputs["href"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NetworkIsolated.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NetworkIsolated resources.
 */
export interface NetworkIsolatedState {
    /**
     * An optional description of the network
     */
    description?: pulumi.Input<string>;
    /**
     * A range of IPs to issue to virtual machines that don't
     * have a static IP; see IP Pools below for details.
     */
    dhcpPools?: pulumi.Input<pulumi.Input<inputs.NetworkIsolatedDhcpPool>[]>;
    /**
     * First DNS server to use.
     */
    dns1?: pulumi.Input<string>;
    /**
     * Second DNS server to use.
     */
    dns2?: pulumi.Input<string>;
    /**
     * A FQDN for the virtual machines on this network
     */
    dnsSuffix?: pulumi.Input<string>;
    /**
     * The gateway for this network
     */
    gateway?: pulumi.Input<string>;
    /**
     * Network Hyper Reference
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
     *
     * <a id="ip-pools"></a>
     */
    metadataEntries?: pulumi.Input<pulumi.Input<inputs.NetworkIsolatedMetadataEntry>[]>;
    /**
     * A unique name for the network
     */
    name?: pulumi.Input<string>;
    /**
     * The netmask for the new network. Defaults to `255.255.255.0`
     */
    netmask?: pulumi.Input<string>;
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
     * A range of IPs permitted to be used as static IPs for
     * virtual machines; see IP Pools below for details.
     */
    staticIpPools?: pulumi.Input<pulumi.Input<inputs.NetworkIsolatedStaticIpPool>[]>;
    /**
     * The name of VDC to use, optional if defined at provider level
     */
    vdc?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NetworkIsolated resource.
 */
export interface NetworkIsolatedArgs {
    /**
     * An optional description of the network
     */
    description?: pulumi.Input<string>;
    /**
     * A range of IPs to issue to virtual machines that don't
     * have a static IP; see IP Pools below for details.
     */
    dhcpPools?: pulumi.Input<pulumi.Input<inputs.NetworkIsolatedDhcpPool>[]>;
    /**
     * First DNS server to use.
     */
    dns1?: pulumi.Input<string>;
    /**
     * Second DNS server to use.
     */
    dns2?: pulumi.Input<string>;
    /**
     * A FQDN for the virtual machines on this network
     */
    dnsSuffix?: pulumi.Input<string>;
    /**
     * The gateway for this network
     */
    gateway?: pulumi.Input<string>;
    /**
     * Use `metadataEntry` instead. Key value map of metadata to assign to this network.
     *
     * @deprecated Use metadataEntry instead
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A set of metadata entries to assign. See Metadata section for details.
     *
     * <a id="ip-pools"></a>
     */
    metadataEntries?: pulumi.Input<pulumi.Input<inputs.NetworkIsolatedMetadataEntry>[]>;
    /**
     * A unique name for the network
     */
    name?: pulumi.Input<string>;
    /**
     * The netmask for the new network. Defaults to `255.255.255.0`
     */
    netmask?: pulumi.Input<string>;
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
     * A range of IPs permitted to be used as static IPs for
     * virtual machines; see IP Pools below for details.
     */
    staticIpPools?: pulumi.Input<pulumi.Input<inputs.NetworkIsolatedStaticIpPool>[]>;
    /**
     * The name of VDC to use, optional if defined at provider level
     */
    vdc?: pulumi.Input<string>;
}
