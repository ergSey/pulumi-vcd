// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a VMware Cloud Director SNAT data source for advanced edge gateways (NSX-V). This can be used to
 * read existing rule by ID and use its attributes in other resources.
 *
 * > **Note:** This data source requires advanced edge gateway.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const my_rule = vcd.getNsxvSnat({
 *     org: "my-org",
 *     vdc: "my-org-vdc",
 *     edgeGateway: "my-edge-gw",
 *     ruleId: "197867",
 * });
 * ```
 */
export function getNsxvSnat(args: GetNsxvSnatArgs, opts?: pulumi.InvokeOptions): Promise<GetNsxvSnatResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vcd:index/getNsxvSnat:getNsxvSnat", {
        "edgeGateway": args.edgeGateway,
        "org": args.org,
        "ruleId": args.ruleId,
        "vdc": args.vdc,
    }, opts);
}

/**
 * A collection of arguments for invoking getNsxvSnat.
 */
export interface GetNsxvSnatArgs {
    /**
     * The name of the edge gateway on which to apply the SNAT rule.
     */
    edgeGateway: string;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.
     */
    org?: string;
    /**
     * ID of SNAT rule as shown in the UI.
     */
    ruleId: string;
    /**
     * The name of VDC to use, optional if defined at provider level.
     */
    vdc?: string;
}

/**
 * A collection of values returned by getNsxvSnat.
 */
export interface GetNsxvSnatResult {
    readonly description: string;
    readonly edgeGateway: string;
    readonly enabled: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly loggingEnabled: boolean;
    readonly networkName: string;
    readonly networkType: string;
    readonly org?: string;
    readonly originalAddress: string;
    readonly ruleId: string;
    readonly ruleTag: number;
    readonly ruleType: string;
    readonly translatedAddress: string;
    readonly vdc?: string;
}
/**
 * Provides a VMware Cloud Director SNAT data source for advanced edge gateways (NSX-V). This can be used to
 * read existing rule by ID and use its attributes in other resources.
 *
 * > **Note:** This data source requires advanced edge gateway.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const my_rule = vcd.getNsxvSnat({
 *     org: "my-org",
 *     vdc: "my-org-vdc",
 *     edgeGateway: "my-edge-gw",
 *     ruleId: "197867",
 * });
 * ```
 */
export function getNsxvSnatOutput(args: GetNsxvSnatOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetNsxvSnatResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vcd:index/getNsxvSnat:getNsxvSnat", {
        "edgeGateway": args.edgeGateway,
        "org": args.org,
        "ruleId": args.ruleId,
        "vdc": args.vdc,
    }, opts);
}

/**
 * A collection of arguments for invoking getNsxvSnat.
 */
export interface GetNsxvSnatOutputArgs {
    /**
     * The name of the edge gateway on which to apply the SNAT rule.
     */
    edgeGateway: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.
     */
    org?: pulumi.Input<string>;
    /**
     * ID of SNAT rule as shown in the UI.
     */
    ruleId: pulumi.Input<string>;
    /**
     * The name of VDC to use, optional if defined at provider level.
     */
    vdc?: pulumi.Input<string>;
}
