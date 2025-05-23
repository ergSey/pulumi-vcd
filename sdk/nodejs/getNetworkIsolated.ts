// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

/**
 * Provides a VMware Cloud Director Org VDC isolated Network data source. This can be used to reference
 * internal networks for vApps to connect. This network is not attached to external networks or routers.
 *
 * Supported in provider *v2.5+*
 *
 * > **Note:** This data source supports only NSX-V backed Org VDC networks.
 * Please use newer [`vcd.NetworkIsolatedV2`](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/network_isolated_v2)
 * data source which is compatible with NSX-T.
 */
export function getNetworkIsolated(args?: GetNetworkIsolatedArgs, opts?: pulumi.InvokeOptions): Promise<GetNetworkIsolatedResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vcd:index/getNetworkIsolated:getNetworkIsolated", {
        "filter": args.filter,
        "name": args.name,
        "org": args.org,
        "vdc": args.vdc,
    }, opts);
}

/**
 * A collection of arguments for invoking getNetworkIsolated.
 */
export interface GetNetworkIsolatedArgs {
    /**
     * Retrieves the data source using one or more filter parameters
     */
    filter?: inputs.GetNetworkIsolatedFilter;
    /**
     * A unique name for the network (optional when `filter` is used)
     */
    name?: string;
    /**
     * The name of organization to use, optional if defined at provider level
     */
    org?: string;
    /**
     * The name of VDC to use, optional if defined at provider level
     */
    vdc?: string;
}

/**
 * A collection of values returned by getNetworkIsolated.
 */
export interface GetNetworkIsolatedResult {
    readonly description: string;
    readonly dhcpPools: outputs.GetNetworkIsolatedDhcpPool[];
    readonly dns1: string;
    readonly dns2: string;
    readonly dnsSuffix: string;
    readonly filter?: outputs.GetNetworkIsolatedFilter;
    readonly gateway: string;
    readonly href: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * @deprecated Use metadataEntry instead
     */
    readonly metadata: {[key: string]: string};
    readonly metadataEntries: outputs.GetNetworkIsolatedMetadataEntry[];
    readonly name?: string;
    readonly netmask: string;
    readonly org?: string;
    readonly shared: boolean;
    readonly staticIpPools: outputs.GetNetworkIsolatedStaticIpPool[];
    readonly vdc?: string;
}
/**
 * Provides a VMware Cloud Director Org VDC isolated Network data source. This can be used to reference
 * internal networks for vApps to connect. This network is not attached to external networks or routers.
 *
 * Supported in provider *v2.5+*
 *
 * > **Note:** This data source supports only NSX-V backed Org VDC networks.
 * Please use newer [`vcd.NetworkIsolatedV2`](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/network_isolated_v2)
 * data source which is compatible with NSX-T.
 */
export function getNetworkIsolatedOutput(args?: GetNetworkIsolatedOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetNetworkIsolatedResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vcd:index/getNetworkIsolated:getNetworkIsolated", {
        "filter": args.filter,
        "name": args.name,
        "org": args.org,
        "vdc": args.vdc,
    }, opts);
}

/**
 * A collection of arguments for invoking getNetworkIsolated.
 */
export interface GetNetworkIsolatedOutputArgs {
    /**
     * Retrieves the data source using one or more filter parameters
     */
    filter?: pulumi.Input<inputs.GetNetworkIsolatedFilterArgs>;
    /**
     * A unique name for the network (optional when `filter` is used)
     */
    name?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level
     */
    org?: pulumi.Input<string>;
    /**
     * The name of VDC to use, optional if defined at provider level
     */
    vdc?: pulumi.Input<string>;
}
