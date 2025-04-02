// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

export class NsxtNetworkDhcp extends pulumi.CustomResource {
    /**
     * Get an existing NsxtNetworkDhcp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NsxtNetworkDhcpState, opts?: pulumi.CustomResourceOptions): NsxtNetworkDhcp {
        return new NsxtNetworkDhcp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/nsxtNetworkDhcp:NsxtNetworkDhcp';

    /**
     * Returns true if the given object is an instance of NsxtNetworkDhcp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NsxtNetworkDhcp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NsxtNetworkDhcp.__pulumiType;
    }

    /**
     * The DNS server IPs to be assigned by this
     * DHCP service. Maximum two values.
     */
    public readonly dnsServers!: pulumi.Output<string[] | undefined>;
    /**
     * Lease time in seconds. Minimum value is 60s
     * and maximum is 4294967295s (~ 49 days).
     */
    public readonly leaseTime!: pulumi.Output<number>;
    /**
     * IP address of DHCP server in network. Must match
     * subnet. **Only** used when `mode=NETWORK`.
     */
    public readonly listenerIpAddress!: pulumi.Output<string | undefined>;
    /**
     * One of `EDGE`, `NETWORK` or `RELAY`. Default is `EDGE`
     * * `EDGE` can be used with Routed Org VDC networks.
     * * `NETWORK` can be used for Isolated and Routed Org VDC networks. It requires
     * `listenerIpAddress` to be set and Edge Cluster must be assigned to VDC.
     * * `RELAY` can be used with Routed Org VDC networks, but requires DHCP forwarding configuration in
     * NSX-T Edge Gateway.
     */
    public readonly mode!: pulumi.Output<string | undefined>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful
     * when connected as sysadmin working across different organisations.
     */
    public readonly org!: pulumi.Output<string | undefined>;
    /**
     * ID of parent Org VDC Routed network.
     */
    public readonly orgNetworkId!: pulumi.Output<string>;
    /**
     * One or more blocks to define DHCP pool ranges. Must not be set when
     * `mode=RELAY`. See Pools and example for usage details.
     */
    public readonly pools!: pulumi.Output<outputs.NsxtNetworkDhcpPool[] | undefined>;
    /**
     * The name of VDC to use, optional if defined at provider level
     *
     * @deprecated Org network will be looked up based on 'org_network_id' field
     */
    public readonly vdc!: pulumi.Output<string>;

    /**
     * Create a NsxtNetworkDhcp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NsxtNetworkDhcpArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NsxtNetworkDhcpArgs | NsxtNetworkDhcpState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NsxtNetworkDhcpState | undefined;
            resourceInputs["dnsServers"] = state ? state.dnsServers : undefined;
            resourceInputs["leaseTime"] = state ? state.leaseTime : undefined;
            resourceInputs["listenerIpAddress"] = state ? state.listenerIpAddress : undefined;
            resourceInputs["mode"] = state ? state.mode : undefined;
            resourceInputs["org"] = state ? state.org : undefined;
            resourceInputs["orgNetworkId"] = state ? state.orgNetworkId : undefined;
            resourceInputs["pools"] = state ? state.pools : undefined;
            resourceInputs["vdc"] = state ? state.vdc : undefined;
        } else {
            const args = argsOrState as NsxtNetworkDhcpArgs | undefined;
            if ((!args || args.orgNetworkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'orgNetworkId'");
            }
            resourceInputs["dnsServers"] = args ? args.dnsServers : undefined;
            resourceInputs["leaseTime"] = args ? args.leaseTime : undefined;
            resourceInputs["listenerIpAddress"] = args ? args.listenerIpAddress : undefined;
            resourceInputs["mode"] = args ? args.mode : undefined;
            resourceInputs["org"] = args ? args.org : undefined;
            resourceInputs["orgNetworkId"] = args ? args.orgNetworkId : undefined;
            resourceInputs["pools"] = args ? args.pools : undefined;
            resourceInputs["vdc"] = args ? args.vdc : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NsxtNetworkDhcp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NsxtNetworkDhcp resources.
 */
export interface NsxtNetworkDhcpState {
    /**
     * The DNS server IPs to be assigned by this
     * DHCP service. Maximum two values.
     */
    dnsServers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Lease time in seconds. Minimum value is 60s
     * and maximum is 4294967295s (~ 49 days).
     */
    leaseTime?: pulumi.Input<number>;
    /**
     * IP address of DHCP server in network. Must match
     * subnet. **Only** used when `mode=NETWORK`.
     */
    listenerIpAddress?: pulumi.Input<string>;
    /**
     * One of `EDGE`, `NETWORK` or `RELAY`. Default is `EDGE`
     * * `EDGE` can be used with Routed Org VDC networks.
     * * `NETWORK` can be used for Isolated and Routed Org VDC networks. It requires
     * `listenerIpAddress` to be set and Edge Cluster must be assigned to VDC.
     * * `RELAY` can be used with Routed Org VDC networks, but requires DHCP forwarding configuration in
     * NSX-T Edge Gateway.
     */
    mode?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful
     * when connected as sysadmin working across different organisations.
     */
    org?: pulumi.Input<string>;
    /**
     * ID of parent Org VDC Routed network.
     */
    orgNetworkId?: pulumi.Input<string>;
    /**
     * One or more blocks to define DHCP pool ranges. Must not be set when
     * `mode=RELAY`. See Pools and example for usage details.
     */
    pools?: pulumi.Input<pulumi.Input<inputs.NsxtNetworkDhcpPool>[]>;
    /**
     * The name of VDC to use, optional if defined at provider level
     *
     * @deprecated Org network will be looked up based on 'org_network_id' field
     */
    vdc?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NsxtNetworkDhcp resource.
 */
export interface NsxtNetworkDhcpArgs {
    /**
     * The DNS server IPs to be assigned by this
     * DHCP service. Maximum two values.
     */
    dnsServers?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Lease time in seconds. Minimum value is 60s
     * and maximum is 4294967295s (~ 49 days).
     */
    leaseTime?: pulumi.Input<number>;
    /**
     * IP address of DHCP server in network. Must match
     * subnet. **Only** used when `mode=NETWORK`.
     */
    listenerIpAddress?: pulumi.Input<string>;
    /**
     * One of `EDGE`, `NETWORK` or `RELAY`. Default is `EDGE`
     * * `EDGE` can be used with Routed Org VDC networks.
     * * `NETWORK` can be used for Isolated and Routed Org VDC networks. It requires
     * `listenerIpAddress` to be set and Edge Cluster must be assigned to VDC.
     * * `RELAY` can be used with Routed Org VDC networks, but requires DHCP forwarding configuration in
     * NSX-T Edge Gateway.
     */
    mode?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful
     * when connected as sysadmin working across different organisations.
     */
    org?: pulumi.Input<string>;
    /**
     * ID of parent Org VDC Routed network.
     */
    orgNetworkId: pulumi.Input<string>;
    /**
     * One or more blocks to define DHCP pool ranges. Must not be set when
     * `mode=RELAY`. See Pools and example for usage details.
     */
    pools?: pulumi.Input<pulumi.Input<inputs.NsxtNetworkDhcpPool>[]>;
    /**
     * The name of VDC to use, optional if defined at provider level
     *
     * @deprecated Org network will be looked up based on 'org_network_id' field
     */
    vdc?: pulumi.Input<string>;
}
