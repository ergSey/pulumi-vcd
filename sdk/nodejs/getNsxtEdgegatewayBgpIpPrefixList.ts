// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

/**
 * Supported in provider *v3.7+* and VCD 10.2+ with NSX-T
 *
 * Provides a resource to manage NSX-T Edge Gateway BGP IP Prefix Lists. IP prefix lists can contain
 * single or multiple IP addresses and can be used to assign BGP neighbors with access permissions
 * for route advertisement.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const g1 = vcd.getVdcGroup({
 *     org: "my-org",
 *     name: "my-vdc-group",
 * });
 * const testing = g1.then(g1 => vcd.getNsxtEdgegateway({
 *     org: "my-org",
 *     ownerId: g1.id,
 *     name: "my-edge-gateway",
 * }));
 * const testingGetNsxtEdgegatewayBgpIpPrefixList = testing.then(testing => vcd.getNsxtEdgegatewayBgpIpPrefixList({
 *     org: "my-org",
 *     edgeGatewayId: testing.id,
 *     name: "my-bgp-prefix-list",
 * }));
 * ```
 */
export function getNsxtEdgegatewayBgpIpPrefixList(args: GetNsxtEdgegatewayBgpIpPrefixListArgs, opts?: pulumi.InvokeOptions): Promise<GetNsxtEdgegatewayBgpIpPrefixListResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vcd:index/getNsxtEdgegatewayBgpIpPrefixList:getNsxtEdgegatewayBgpIpPrefixList", {
        "edgeGatewayId": args.edgeGatewayId,
        "name": args.name,
        "org": args.org,
    }, opts);
}

/**
 * A collection of arguments for invoking getNsxtEdgegatewayBgpIpPrefixList.
 */
export interface GetNsxtEdgegatewayBgpIpPrefixListArgs {
    /**
     * An ID of NSX-T Edge Gateway. Can be looked up using
     * [vcd.NsxtEdgegateway](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_edgegateway) data source
     */
    edgeGatewayId: string;
    /**
     * A name of existing BGP IP Prefix List in specified Edge Gateway
     */
    name: string;
    /**
     * The name of organization to which the edge gateway belongs. Optional if defined at provider level.
     */
    org?: string;
}

/**
 * A collection of values returned by getNsxtEdgegatewayBgpIpPrefixList.
 */
export interface GetNsxtEdgegatewayBgpIpPrefixListResult {
    readonly description: string;
    readonly edgeGatewayId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ipPrefixes: outputs.GetNsxtEdgegatewayBgpIpPrefixListIpPrefix[];
    readonly name: string;
    readonly org?: string;
}
/**
 * Supported in provider *v3.7+* and VCD 10.2+ with NSX-T
 *
 * Provides a resource to manage NSX-T Edge Gateway BGP IP Prefix Lists. IP prefix lists can contain
 * single or multiple IP addresses and can be used to assign BGP neighbors with access permissions
 * for route advertisement.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const g1 = vcd.getVdcGroup({
 *     org: "my-org",
 *     name: "my-vdc-group",
 * });
 * const testing = g1.then(g1 => vcd.getNsxtEdgegateway({
 *     org: "my-org",
 *     ownerId: g1.id,
 *     name: "my-edge-gateway",
 * }));
 * const testingGetNsxtEdgegatewayBgpIpPrefixList = testing.then(testing => vcd.getNsxtEdgegatewayBgpIpPrefixList({
 *     org: "my-org",
 *     edgeGatewayId: testing.id,
 *     name: "my-bgp-prefix-list",
 * }));
 * ```
 */
export function getNsxtEdgegatewayBgpIpPrefixListOutput(args: GetNsxtEdgegatewayBgpIpPrefixListOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetNsxtEdgegatewayBgpIpPrefixListResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vcd:index/getNsxtEdgegatewayBgpIpPrefixList:getNsxtEdgegatewayBgpIpPrefixList", {
        "edgeGatewayId": args.edgeGatewayId,
        "name": args.name,
        "org": args.org,
    }, opts);
}

/**
 * A collection of arguments for invoking getNsxtEdgegatewayBgpIpPrefixList.
 */
export interface GetNsxtEdgegatewayBgpIpPrefixListOutputArgs {
    /**
     * An ID of NSX-T Edge Gateway. Can be looked up using
     * [vcd.NsxtEdgegateway](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_edgegateway) data source
     */
    edgeGatewayId: pulumi.Input<string>;
    /**
     * A name of existing BGP IP Prefix List in specified Edge Gateway
     */
    name: pulumi.Input<string>;
    /**
     * The name of organization to which the edge gateway belongs. Optional if defined at provider level.
     */
    org?: pulumi.Input<string>;
}
