// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Supported in provider *v3.4+* and VCD 10.2+ with NSX-T and ALB.
 *
 * Provides a data source to reference existing ALB Importable Clouds. An NSX-T Importable Cloud is a reference to a
 * Cloud configured in ALB Controller.
 *
 * > Only `System Administrator` can use this data source.
 *
 * > VCD 10.3.0 has a caching bug which prevents listing importable clouds immediately after [ALB
 * Controller](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_controller) is created. This data should be
 * available 15 minutes after the Controller is created.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const first = vcd.getNsxtAlbController({
 *     name: "alb-controller",
 * });
 * const cld = first.then(first => vcd.getNsxtAlbImportableCloud({
 *     name: "NSXT Importable Cloud",
 *     controllerId: first.id,
 * }));
 * ```
 */
export function getNsxtAlbImportableCloud(args: GetNsxtAlbImportableCloudArgs, opts?: pulumi.InvokeOptions): Promise<GetNsxtAlbImportableCloudResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vcd:index/getNsxtAlbImportableCloud:getNsxtAlbImportableCloud", {
        "controllerId": args.controllerId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getNsxtAlbImportableCloud.
 */
export interface GetNsxtAlbImportableCloudArgs {
    /**
     * ALB Controller ID
     */
    controllerId: string;
    /**
     * Name of ALB Importable Cloud
     */
    name: string;
}

/**
 * A collection of values returned by getNsxtAlbImportableCloud.
 */
export interface GetNsxtAlbImportableCloudResult {
    /**
     * boolean value which displays if the ALB Importable Cloud is already consumed
     */
    readonly alreadyImported: boolean;
    readonly controllerId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    /**
     * backing network pool ID
     */
    readonly networkPoolId: string;
    /**
     * backing network pool ID
     */
    readonly networkPoolName: string;
    /**
     * backing transport zone name
     */
    readonly transportZoneName: string;
}
/**
 * Supported in provider *v3.4+* and VCD 10.2+ with NSX-T and ALB.
 *
 * Provides a data source to reference existing ALB Importable Clouds. An NSX-T Importable Cloud is a reference to a
 * Cloud configured in ALB Controller.
 *
 * > Only `System Administrator` can use this data source.
 *
 * > VCD 10.3.0 has a caching bug which prevents listing importable clouds immediately after [ALB
 * Controller](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_controller) is created. This data should be
 * available 15 minutes after the Controller is created.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const first = vcd.getNsxtAlbController({
 *     name: "alb-controller",
 * });
 * const cld = first.then(first => vcd.getNsxtAlbImportableCloud({
 *     name: "NSXT Importable Cloud",
 *     controllerId: first.id,
 * }));
 * ```
 */
export function getNsxtAlbImportableCloudOutput(args: GetNsxtAlbImportableCloudOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetNsxtAlbImportableCloudResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vcd:index/getNsxtAlbImportableCloud:getNsxtAlbImportableCloud", {
        "controllerId": args.controllerId,
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getNsxtAlbImportableCloud.
 */
export interface GetNsxtAlbImportableCloudOutputArgs {
    /**
     * ALB Controller ID
     */
    controllerId: pulumi.Input<string>;
    /**
     * Name of ALB Importable Cloud
     */
    name: pulumi.Input<string>;
}
