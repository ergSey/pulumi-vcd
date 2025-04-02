// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

/**
 * Provides a VMware Cloud Director vApp network data source. This can be used to access a vApp network.
 *
 * Supported in provider *v2.7+*
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const web = vcd.getVapp({
 *     name: "web",
 * });
 * const network1 = web.then(web => vcd.getVappNetwork({
 *     vappName: web.name,
 *     name: "isolated-network",
 * }));
 * export const gateway = network1.then(network1 => network1.gateway);
 * ```
 */
export function getVappNetwork(args: GetVappNetworkArgs, opts?: pulumi.InvokeOptions): Promise<GetVappNetworkResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vcd:index/getVappNetwork:getVappNetwork", {
        "name": args.name,
        "org": args.org,
        "vappName": args.vappName,
        "vdc": args.vdc,
    }, opts);
}

/**
 * A collection of arguments for invoking getVappNetwork.
 */
export interface GetVappNetworkArgs {
    /**
     * A name for the vApp network, unique within the vApp
     */
    name: string;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
     */
    org?: string;
    /**
     * The vApp name.
     */
    vappName: string;
    /**
     * The name of VDC to use, optional if defined at provider level
     */
    vdc?: string;
}

/**
 * A collection of values returned by getVappNetwork.
 */
export interface GetVappNetworkResult {
    readonly description: string;
    readonly dhcpPools: outputs.GetVappNetworkDhcpPool[];
    readonly dns1: string;
    readonly dns2: string;
    readonly dnsSuffix: string;
    readonly gateway: string;
    readonly guestVlanAllowed: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly netmask: string;
    readonly org?: string;
    readonly orgNetworkName: string;
    readonly prefixLength: string;
    readonly retainIpMacEnabled: boolean;
    readonly staticIpPools: outputs.GetVappNetworkStaticIpPool[];
    readonly vappName: string;
    readonly vdc?: string;
}
/**
 * Provides a VMware Cloud Director vApp network data source. This can be used to access a vApp network.
 *
 * Supported in provider *v2.7+*
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const web = vcd.getVapp({
 *     name: "web",
 * });
 * const network1 = web.then(web => vcd.getVappNetwork({
 *     vappName: web.name,
 *     name: "isolated-network",
 * }));
 * export const gateway = network1.then(network1 => network1.gateway);
 * ```
 */
export function getVappNetworkOutput(args: GetVappNetworkOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetVappNetworkResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vcd:index/getVappNetwork:getVappNetwork", {
        "name": args.name,
        "org": args.org,
        "vappName": args.vappName,
        "vdc": args.vdc,
    }, opts);
}

/**
 * A collection of arguments for invoking getVappNetwork.
 */
export interface GetVappNetworkOutputArgs {
    /**
     * A name for the vApp network, unique within the vApp
     */
    name: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
     */
    org?: pulumi.Input<string>;
    /**
     * The vApp name.
     */
    vappName: pulumi.Input<string>;
    /**
     * The name of VDC to use, optional if defined at provider level
     */
    vdc?: pulumi.Input<string>;
}
