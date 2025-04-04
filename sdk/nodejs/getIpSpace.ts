// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

/**
 * Provides a data source to read IP Spaces. IP Spaces provide structured approach to allocating public
 * and private IP addresses by preventing the use of overlapping IP addresses across organizations and
 * organization VDCs.
 *
 * IP Spaces require VCD 10.4.1+ with NSX-T.
 *
 * ## Example Usage
 *
 * ### Private IP Space Within An Org)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const space1 = vcd.getIpSpace({
 *     orgId: org1.id,
 *     name: "private-ip-space",
 * });
 * ```
 *
 * ### Public Or Shared IP Space)
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const space1 = vcd.getIpSpace({
 *     name: "public-or-shared-ip-space",
 * });
 * ```
 */
export function getIpSpace(args: GetIpSpaceArgs, opts?: pulumi.InvokeOptions): Promise<GetIpSpaceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vcd:index/getIpSpace:getIpSpace", {
        "name": args.name,
        "orgId": args.orgId,
    }, opts);
}

/**
 * A collection of arguments for invoking getIpSpace.
 */
export interface GetIpSpaceArgs {
    /**
     * The name of IP Space.
     */
    name: string;
    /**
     * Org ID for Private IP Space.
     */
    orgId?: string;
}

/**
 * A collection of values returned by getIpSpace.
 */
export interface GetIpSpaceResult {
    readonly defaultFirewallRuleCreationEnabled: boolean;
    readonly defaultNoSnatRuleCreationEnabled: boolean;
    readonly defaultSnatRuleCreationEnabled: boolean;
    readonly description: string;
    readonly externalScope: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly internalScopes: string[];
    readonly ipPrefixes: outputs.GetIpSpaceIpPrefix[];
    readonly ipRangeQuota: string;
    readonly ipRanges: outputs.GetIpSpaceIpRange[];
    readonly name: string;
    readonly orgId?: string;
    readonly routeAdvertisementEnabled: boolean;
    readonly type: string;
}
/**
 * Provides a data source to read IP Spaces. IP Spaces provide structured approach to allocating public
 * and private IP addresses by preventing the use of overlapping IP addresses across organizations and
 * organization VDCs.
 *
 * IP Spaces require VCD 10.4.1+ with NSX-T.
 *
 * ## Example Usage
 *
 * ### Private IP Space Within An Org)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const space1 = vcd.getIpSpace({
 *     orgId: org1.id,
 *     name: "private-ip-space",
 * });
 * ```
 *
 * ### Public Or Shared IP Space)
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const space1 = vcd.getIpSpace({
 *     name: "public-or-shared-ip-space",
 * });
 * ```
 */
export function getIpSpaceOutput(args: GetIpSpaceOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetIpSpaceResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vcd:index/getIpSpace:getIpSpace", {
        "name": args.name,
        "orgId": args.orgId,
    }, opts);
}

/**
 * A collection of arguments for invoking getIpSpace.
 */
export interface GetIpSpaceOutputArgs {
    /**
     * The name of IP Space.
     */
    name: pulumi.Input<string>;
    /**
     * Org ID for Private IP Space.
     */
    orgId?: pulumi.Input<string>;
}
