// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

/**
 * Provides a VMware Cloud Director NSX-T edge gateway data source. This can be used to read NSX-T edge gateway configurations.
 *
 * Supported in provider *v3.1+*.
 *
 * ## Example Usage
 *
 * ### NSX-T Edge Gateway Belonging To VDC Group)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const group1 = vcd.getVdcGroup({
 *     name: "existing-group",
 * });
 * const t1 = group1.then(group1 => vcd.getNsxtEdgegateway({
 *     org: "myorg",
 *     ownerId: group1.id,
 *     name: "nsxt-edge-gateway",
 * }));
 * ```
 *
 * ### NSX-T Edge Gateway Belonging To VDC)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const vdc1 = vcd.getOrgVdc({
 *     name: "existing-vdc",
 * });
 * const t1 = vdc1.then(vdc1 => vcd.getNsxtEdgegateway({
 *     org: "myorg",
 *     ownerId: vdc1.id,
 *     name: "nsxt-edge-gateway",
 * }));
 * ```
 */
export function getNsxtEdgegateway(args: GetNsxtEdgegatewayArgs, opts?: pulumi.InvokeOptions): Promise<GetNsxtEdgegatewayResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vcd:index/getNsxtEdgegateway:getNsxtEdgegateway", {
        "edgeClusterId": args.edgeClusterId,
        "ipCountReadLimit": args.ipCountReadLimit,
        "name": args.name,
        "org": args.org,
        "ownerId": args.ownerId,
        "vdc": args.vdc,
    }, opts);
}

/**
 * A collection of arguments for invoking getNsxtEdgegateway.
 */
export interface GetNsxtEdgegatewayArgs {
    edgeClusterId?: string;
    /**
     * Sets a limit of IPs to count for
     * `usedIpCount` and `unusedIpCount` attributes to avoid exhausting compute resource while
     * counting IPs in large IPv6 subnets. It does not affect operation of Edge Gateway configuration,
     * only IP count reporting. Defaults to `1000000`. While it is unlikely that a single Edge Gateway
     * can effectively manage more IPs, one can specify `0` for *unlimited* value.
     */
    ipCountReadLimit?: number;
    /**
     * NSX-T Edge Gateway name.
     */
    name: string;
    /**
     * The name of organization to which the NSX-T Edge Gateway belongs. Optional if
     * defined at provider level.
     */
    org?: string;
    /**
     * The ID of VDC or VDC Group. **Note:** Data sources
     * [vcd.VdcGroup](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/vdc_group) or
     * [vcd.OrgVdc](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/org_vdc) can be used to lookup IDs by
     * name.
     *
     * > Only one of `vdc` or `ownerId` can be specified. `ownerId` takes precedence over `vdc`
     * definition at provider level.
     */
    ownerId?: string;
    /**
     * **Deprecated** - please use `ownerId` field. The name of VDC that owns the
     * NSX-T Edge Gateway. Optional if defined at provider level.
     *
     * @deprecated This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
     */
    vdc?: string;
}

/**
 * A collection of values returned by getNsxtEdgegateway.
 */
export interface GetNsxtEdgegatewayResult {
    readonly dedicateExternalNetwork: boolean;
    readonly deploymentMode: string;
    readonly description: string;
    readonly edgeClusterId?: string;
    readonly externalNetworkAllocatedIpCount: number;
    readonly externalNetworkId: string;
    readonly externalNetworks: outputs.GetNsxtEdgegatewayExternalNetwork[];
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ipCountReadLimit?: number;
    readonly name: string;
    readonly nonDistributedRoutingEnabled: boolean;
    readonly org?: string;
    readonly ownerId: string;
    readonly primaryIp: string;
    readonly subnetWithIpCounts: outputs.GetNsxtEdgegatewaySubnetWithIpCount[];
    readonly subnetWithTotalIpCounts: outputs.GetNsxtEdgegatewaySubnetWithTotalIpCount[];
    readonly subnets: outputs.GetNsxtEdgegatewaySubnet[];
    readonly totalAllocatedIpCount: number;
    readonly unusedIpCount: number;
    readonly useIpSpaces: boolean;
    readonly usedIpCount: number;
    /**
     * @deprecated This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
     */
    readonly vdc: string;
}
/**
 * Provides a VMware Cloud Director NSX-T edge gateway data source. This can be used to read NSX-T edge gateway configurations.
 *
 * Supported in provider *v3.1+*.
 *
 * ## Example Usage
 *
 * ### NSX-T Edge Gateway Belonging To VDC Group)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const group1 = vcd.getVdcGroup({
 *     name: "existing-group",
 * });
 * const t1 = group1.then(group1 => vcd.getNsxtEdgegateway({
 *     org: "myorg",
 *     ownerId: group1.id,
 *     name: "nsxt-edge-gateway",
 * }));
 * ```
 *
 * ### NSX-T Edge Gateway Belonging To VDC)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const vdc1 = vcd.getOrgVdc({
 *     name: "existing-vdc",
 * });
 * const t1 = vdc1.then(vdc1 => vcd.getNsxtEdgegateway({
 *     org: "myorg",
 *     ownerId: vdc1.id,
 *     name: "nsxt-edge-gateway",
 * }));
 * ```
 */
export function getNsxtEdgegatewayOutput(args: GetNsxtEdgegatewayOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetNsxtEdgegatewayResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vcd:index/getNsxtEdgegateway:getNsxtEdgegateway", {
        "edgeClusterId": args.edgeClusterId,
        "ipCountReadLimit": args.ipCountReadLimit,
        "name": args.name,
        "org": args.org,
        "ownerId": args.ownerId,
        "vdc": args.vdc,
    }, opts);
}

/**
 * A collection of arguments for invoking getNsxtEdgegateway.
 */
export interface GetNsxtEdgegatewayOutputArgs {
    edgeClusterId?: pulumi.Input<string>;
    /**
     * Sets a limit of IPs to count for
     * `usedIpCount` and `unusedIpCount` attributes to avoid exhausting compute resource while
     * counting IPs in large IPv6 subnets. It does not affect operation of Edge Gateway configuration,
     * only IP count reporting. Defaults to `1000000`. While it is unlikely that a single Edge Gateway
     * can effectively manage more IPs, one can specify `0` for *unlimited* value.
     */
    ipCountReadLimit?: pulumi.Input<number>;
    /**
     * NSX-T Edge Gateway name.
     */
    name: pulumi.Input<string>;
    /**
     * The name of organization to which the NSX-T Edge Gateway belongs. Optional if
     * defined at provider level.
     */
    org?: pulumi.Input<string>;
    /**
     * The ID of VDC or VDC Group. **Note:** Data sources
     * [vcd.VdcGroup](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/vdc_group) or
     * [vcd.OrgVdc](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/org_vdc) can be used to lookup IDs by
     * name.
     *
     * > Only one of `vdc` or `ownerId` can be specified. `ownerId` takes precedence over `vdc`
     * definition at provider level.
     */
    ownerId?: pulumi.Input<string>;
    /**
     * **Deprecated** - please use `ownerId` field. The name of VDC that owns the
     * NSX-T Edge Gateway. Optional if defined at provider level.
     *
     * @deprecated This field is deprecated in favor of 'owner_id' which supports both - VDC and VDC Group IDs
     */
    vdc?: pulumi.Input<string>;
}
