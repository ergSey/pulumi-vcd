// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

export class IpSpace extends pulumi.CustomResource {
    /**
     * Get an existing IpSpace resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IpSpaceState, opts?: pulumi.CustomResourceOptions): IpSpace {
        return new IpSpace(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/ipSpace:IpSpace';

    /**
     * Returns true if the given object is an instance of IpSpace.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IpSpace {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IpSpace.__pulumiType;
    }

    /**
     * Defines whether
     * default firewall rule creation should be enabled
     */
    public readonly defaultFirewallRuleCreationEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Defines whether NO SNAT
     * rule creation should be enabled
     */
    public readonly defaultNoSnatRuleCreationEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Defines whether SNAT rule
     * creation should be enabled
     *
     * <a id="ipspace-ip-range"></a>
     */
    public readonly defaultSnatRuleCreationEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * Description of IP Space
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The external scope defines the total span of IP addresses to which the IP
     * space has access, for example the internet or a WAN.
     */
    public readonly externalScope!: pulumi.Output<string | undefined>;
    /**
     * The internal scope of an IP space is a list of CIDR notations that
     * defines the exact span of IP addresses in which all ranges and blocks must be contained in.
     */
    public readonly internalScopes!: pulumi.Output<string[]>;
    /**
     * One or more IP prefixes (blocks) ip_prefix
     */
    public readonly ipPrefixes!: pulumi.Output<outputs.IpSpaceIpPrefix[] | undefined>;
    /**
     * If you entered at least one IP Range
     * (ip_range), enter a number of floating IP addresses to allocate individually.
     * `-1` is unlimited, while `0` means that no IPs can be allocated.
     */
    public readonly ipRangeQuota!: pulumi.Output<string>;
    /**
     * One or more ipRange for floating IP address
     * allocation. (Floating IP addresses are just IP addresses taken from the defined range)
     */
    public readonly ipRanges!: pulumi.Output<outputs.IpSpaceIpRange[] | undefined>;
    /**
     * A name for IP Space
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Required for `PRIVATE` type
     */
    public readonly orgId!: pulumi.Output<string | undefined>;
    /**
     * Toggle on the route advertisement option to
     * enable advertising networks with IP prefixes from this IP space (default `false`)
     */
    public readonly routeAdvertisementEnabled!: pulumi.Output<boolean | undefined>;
    /**
     * One of `PUBLIC`, `SHARED_SERVICES`, `PRIVATE`
     * * `PUBLIC` - A public IP space is *used by multiple organizations* and is *controlled by the service
     * provider* through a quota-based system.
     * * `SHARED_SERVICES` - An IP space for services and management networks that are required in the
     * tenant space, but as a service provider, you don't want to expose it to organizations in your
     * environment. The main difference from `PUBLIC` network is that IPs cannot be allocated by tenants.
     * * `PRIVATE` - Private IP spaces are dedicated to a single tenant - a private IP space is used by
     * only one organization that is specified during the space creation. For this organization, IP
     * consumption is unlimited.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a IpSpace resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IpSpaceArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IpSpaceArgs | IpSpaceState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IpSpaceState | undefined;
            resourceInputs["defaultFirewallRuleCreationEnabled"] = state ? state.defaultFirewallRuleCreationEnabled : undefined;
            resourceInputs["defaultNoSnatRuleCreationEnabled"] = state ? state.defaultNoSnatRuleCreationEnabled : undefined;
            resourceInputs["defaultSnatRuleCreationEnabled"] = state ? state.defaultSnatRuleCreationEnabled : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["externalScope"] = state ? state.externalScope : undefined;
            resourceInputs["internalScopes"] = state ? state.internalScopes : undefined;
            resourceInputs["ipPrefixes"] = state ? state.ipPrefixes : undefined;
            resourceInputs["ipRangeQuota"] = state ? state.ipRangeQuota : undefined;
            resourceInputs["ipRanges"] = state ? state.ipRanges : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
            resourceInputs["routeAdvertisementEnabled"] = state ? state.routeAdvertisementEnabled : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
        } else {
            const args = argsOrState as IpSpaceArgs | undefined;
            if ((!args || args.internalScopes === undefined) && !opts.urn) {
                throw new Error("Missing required property 'internalScopes'");
            }
            if ((!args || args.type === undefined) && !opts.urn) {
                throw new Error("Missing required property 'type'");
            }
            resourceInputs["defaultFirewallRuleCreationEnabled"] = args ? args.defaultFirewallRuleCreationEnabled : undefined;
            resourceInputs["defaultNoSnatRuleCreationEnabled"] = args ? args.defaultNoSnatRuleCreationEnabled : undefined;
            resourceInputs["defaultSnatRuleCreationEnabled"] = args ? args.defaultSnatRuleCreationEnabled : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["externalScope"] = args ? args.externalScope : undefined;
            resourceInputs["internalScopes"] = args ? args.internalScopes : undefined;
            resourceInputs["ipPrefixes"] = args ? args.ipPrefixes : undefined;
            resourceInputs["ipRangeQuota"] = args ? args.ipRangeQuota : undefined;
            resourceInputs["ipRanges"] = args ? args.ipRanges : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["orgId"] = args ? args.orgId : undefined;
            resourceInputs["routeAdvertisementEnabled"] = args ? args.routeAdvertisementEnabled : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IpSpace.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IpSpace resources.
 */
export interface IpSpaceState {
    /**
     * Defines whether
     * default firewall rule creation should be enabled
     */
    defaultFirewallRuleCreationEnabled?: pulumi.Input<boolean>;
    /**
     * Defines whether NO SNAT
     * rule creation should be enabled
     */
    defaultNoSnatRuleCreationEnabled?: pulumi.Input<boolean>;
    /**
     * Defines whether SNAT rule
     * creation should be enabled
     *
     * <a id="ipspace-ip-range"></a>
     */
    defaultSnatRuleCreationEnabled?: pulumi.Input<boolean>;
    /**
     * Description of IP Space
     */
    description?: pulumi.Input<string>;
    /**
     * The external scope defines the total span of IP addresses to which the IP
     * space has access, for example the internet or a WAN.
     */
    externalScope?: pulumi.Input<string>;
    /**
     * The internal scope of an IP space is a list of CIDR notations that
     * defines the exact span of IP addresses in which all ranges and blocks must be contained in.
     */
    internalScopes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * One or more IP prefixes (blocks) ip_prefix
     */
    ipPrefixes?: pulumi.Input<pulumi.Input<inputs.IpSpaceIpPrefix>[]>;
    /**
     * If you entered at least one IP Range
     * (ip_range), enter a number of floating IP addresses to allocate individually.
     * `-1` is unlimited, while `0` means that no IPs can be allocated.
     */
    ipRangeQuota?: pulumi.Input<string>;
    /**
     * One or more ipRange for floating IP address
     * allocation. (Floating IP addresses are just IP addresses taken from the defined range)
     */
    ipRanges?: pulumi.Input<pulumi.Input<inputs.IpSpaceIpRange>[]>;
    /**
     * A name for IP Space
     */
    name?: pulumi.Input<string>;
    /**
     * Required for `PRIVATE` type
     */
    orgId?: pulumi.Input<string>;
    /**
     * Toggle on the route advertisement option to
     * enable advertising networks with IP prefixes from this IP space (default `false`)
     */
    routeAdvertisementEnabled?: pulumi.Input<boolean>;
    /**
     * One of `PUBLIC`, `SHARED_SERVICES`, `PRIVATE`
     * * `PUBLIC` - A public IP space is *used by multiple organizations* and is *controlled by the service
     * provider* through a quota-based system.
     * * `SHARED_SERVICES` - An IP space for services and management networks that are required in the
     * tenant space, but as a service provider, you don't want to expose it to organizations in your
     * environment. The main difference from `PUBLIC` network is that IPs cannot be allocated by tenants.
     * * `PRIVATE` - Private IP spaces are dedicated to a single tenant - a private IP space is used by
     * only one organization that is specified during the space creation. For this organization, IP
     * consumption is unlimited.
     */
    type?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IpSpace resource.
 */
export interface IpSpaceArgs {
    /**
     * Defines whether
     * default firewall rule creation should be enabled
     */
    defaultFirewallRuleCreationEnabled?: pulumi.Input<boolean>;
    /**
     * Defines whether NO SNAT
     * rule creation should be enabled
     */
    defaultNoSnatRuleCreationEnabled?: pulumi.Input<boolean>;
    /**
     * Defines whether SNAT rule
     * creation should be enabled
     *
     * <a id="ipspace-ip-range"></a>
     */
    defaultSnatRuleCreationEnabled?: pulumi.Input<boolean>;
    /**
     * Description of IP Space
     */
    description?: pulumi.Input<string>;
    /**
     * The external scope defines the total span of IP addresses to which the IP
     * space has access, for example the internet or a WAN.
     */
    externalScope?: pulumi.Input<string>;
    /**
     * The internal scope of an IP space is a list of CIDR notations that
     * defines the exact span of IP addresses in which all ranges and blocks must be contained in.
     */
    internalScopes: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * One or more IP prefixes (blocks) ip_prefix
     */
    ipPrefixes?: pulumi.Input<pulumi.Input<inputs.IpSpaceIpPrefix>[]>;
    /**
     * If you entered at least one IP Range
     * (ip_range), enter a number of floating IP addresses to allocate individually.
     * `-1` is unlimited, while `0` means that no IPs can be allocated.
     */
    ipRangeQuota?: pulumi.Input<string>;
    /**
     * One or more ipRange for floating IP address
     * allocation. (Floating IP addresses are just IP addresses taken from the defined range)
     */
    ipRanges?: pulumi.Input<pulumi.Input<inputs.IpSpaceIpRange>[]>;
    /**
     * A name for IP Space
     */
    name?: pulumi.Input<string>;
    /**
     * Required for `PRIVATE` type
     */
    orgId?: pulumi.Input<string>;
    /**
     * Toggle on the route advertisement option to
     * enable advertising networks with IP prefixes from this IP space (default `false`)
     */
    routeAdvertisementEnabled?: pulumi.Input<boolean>;
    /**
     * One of `PUBLIC`, `SHARED_SERVICES`, `PRIVATE`
     * * `PUBLIC` - A public IP space is *used by multiple organizations* and is *controlled by the service
     * provider* through a quota-based system.
     * * `SHARED_SERVICES` - An IP space for services and management networks that are required in the
     * tenant space, but as a service provider, you don't want to expose it to organizations in your
     * environment. The main difference from `PUBLIC` network is that IPs cannot be allocated by tenants.
     * * `PRIVATE` - Private IP spaces are dedicated to a single tenant - a private IP space is used by
     * only one organization that is specified during the space creation. For this organization, IP
     * consumption is unlimited.
     */
    type: pulumi.Input<string>;
}
