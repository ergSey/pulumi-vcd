// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a data source to read IP Space Uplinks in External Networks (Provider Gateways).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const u1 = vcd.getIpSpaceUplink({
 *     name: "ip-space-uplink-1",
 *     externalNetworkId: provider_gateway.id,
 * });
 * ```
 */
export function getIpSpaceUplink(args: GetIpSpaceUplinkArgs, opts?: pulumi.InvokeOptions): Promise<GetIpSpaceUplinkResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vcd:index/getIpSpaceUplink:getIpSpaceUplink", {
        "description": args.description,
        "externalNetworkId": args.externalNetworkId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getIpSpaceUplink.
 */
export interface GetIpSpaceUplinkArgs {
    description?: string;
    /**
     * Parent External Network ID
     */
    externalNetworkId: string;
    /**
     * Name of IP Space Uplink
     */
    name: string;
}

/**
 * A collection of values returned by getIpSpaceUplink.
 */
export interface GetIpSpaceUplinkResult {
    readonly associatedInterfaceIds: string[];
    readonly description?: string;
    readonly externalNetworkId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ipSpaceId: string;
    readonly ipSpaceType: string;
    readonly name: string;
    readonly status: string;
}
/**
 * Provides a data source to read IP Space Uplinks in External Networks (Provider Gateways).
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const u1 = vcd.getIpSpaceUplink({
 *     name: "ip-space-uplink-1",
 *     externalNetworkId: provider_gateway.id,
 * });
 * ```
 */
export function getIpSpaceUplinkOutput(args: GetIpSpaceUplinkOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetIpSpaceUplinkResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vcd:index/getIpSpaceUplink:getIpSpaceUplink", {
        "description": args.description,
        "externalNetworkId": args.externalNetworkId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getIpSpaceUplink.
 */
export interface GetIpSpaceUplinkOutputArgs {
    description?: pulumi.Input<string>;
    /**
     * Parent External Network ID
     */
    externalNetworkId: pulumi.Input<string>;
    /**
     * Name of IP Space Uplink
     */
    name: pulumi.Input<string>;
}
