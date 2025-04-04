// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a data source to read Network Profile for NSX-T VDCs.
 *
 * Supported in provider *v3.11+* and VCD 10.4.0+ with NSX-T.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const nsxt = vcd.getOrgVdcNsxtNetworkProfile({
 *     org: "my-org",
 *     vdc: "my-vdc",
 * });
 * ```
 */
export function getOrgVdcNsxtNetworkProfile(args?: GetOrgVdcNsxtNetworkProfileArgs, opts?: pulumi.InvokeOptions): Promise<GetOrgVdcNsxtNetworkProfileResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vcd:index/getOrgVdcNsxtNetworkProfile:getOrgVdcNsxtNetworkProfile", {
        "org": args.org,
        "vdc": args.vdc,
    }, opts);
}

/**
 * A collection of arguments for invoking getOrgVdcNsxtNetworkProfile.
 */
export interface GetOrgVdcNsxtNetworkProfileArgs {
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
 * A collection of values returned by getOrgVdcNsxtNetworkProfile.
 */
export interface GetOrgVdcNsxtNetworkProfileResult {
    /**
     * An ID of NSX-T Edge Cluster which should provide vApp
     * Networking Services or DHCP for Isolated Networks. This value **might be unavailable** in data
     * source if user has insufficient rights.
     */
    readonly edgeClusterId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly org?: string;
    /**
     * Default Segment Profile ID for all vApp
     * networks withing this VDC
     */
    readonly vappNetworksDefaultSegmentProfileTemplateId: string;
    readonly vdc?: string;
    /**
     * Default Segment Profile ID for all Org VDC
     * networks withing this VDC
     */
    readonly vdcNetworksDefaultSegmentProfileTemplateId: string;
}
/**
 * Provides a data source to read Network Profile for NSX-T VDCs.
 *
 * Supported in provider *v3.11+* and VCD 10.4.0+ with NSX-T.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const nsxt = vcd.getOrgVdcNsxtNetworkProfile({
 *     org: "my-org",
 *     vdc: "my-vdc",
 * });
 * ```
 */
export function getOrgVdcNsxtNetworkProfileOutput(args?: GetOrgVdcNsxtNetworkProfileOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetOrgVdcNsxtNetworkProfileResult> {
    args = args || {};
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vcd:index/getOrgVdcNsxtNetworkProfile:getOrgVdcNsxtNetworkProfile", {
        "org": args.org,
        "vdc": args.vdc,
    }, opts);
}

/**
 * A collection of arguments for invoking getOrgVdcNsxtNetworkProfile.
 */
export interface GetOrgVdcNsxtNetworkProfileOutputArgs {
    /**
     * The name of organization to use, optional if defined at provider level
     */
    org?: pulumi.Input<string>;
    /**
     * The name of VDC to use, optional if defined at provider level
     */
    vdc?: pulumi.Input<string>;
}
