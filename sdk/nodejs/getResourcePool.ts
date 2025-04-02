// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a data source for a resource pool attached to a vCenter. A resource pool is an essential component of a Provider VDC.
 *
 * > Note 1: this data source requires System Administrator privileges
 *
 * > Note 2: you can create or modify a resource pool using vSphere provider
 *
 * Supported in provider *v3.10+*
 *
 * ## Example Usage
 *
 * ### 1
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const vcenter1 = vcd.getVcenter({
 *     name: "vc1",
 * });
 * const rp1 = vcenter1.then(vcenter1 => vcd.getResourcePool({
 *     name: "resource-pool-for-vcd-01",
 *     vcenterId: vcenter1.id,
 * }));
 * ```
 *
 * ### 2
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const rp1 = vcd.getResourcePool({
 *     name: "common-name",
 *     vcenterId: vcenter1.id,
 * });
 * ```
 *
 * When you receive such error, you can run the script again, but using the resource pool ID instead of the name.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const rp1 = vcd.getResourcePool({
 *     name: "resgroup-241",
 *     vcenterId: vcenter1.id,
 * });
 * ```
 */
export function getResourcePool(args: GetResourcePoolArgs, opts?: pulumi.InvokeOptions): Promise<GetResourcePoolResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vcd:index/getResourcePool:getResourcePool", {
        "name": args.name,
        "vcenterId": args.vcenterId,
    }, opts);
}

/**
 * A collection of arguments for invoking getResourcePool.
 */
export interface GetResourcePoolArgs {
    /**
     * resource pool name. The name may not be unique within the vCenter. If that happens, you will get an
     * error message with the list of IDs for the pools with the same name, and can subsequently enter the resource pool ID instead of the name.
     * (See Example Usage 2)
     */
    name: string;
    /**
     * ID of the vCenter to which this resource pool belongs.
     */
    vcenterId: string;
}

/**
 * A collection of values returned by getResourcePool.
 */
export interface GetResourcePoolResult {
    /**
     * managed object reference of the vCenter cluster that this resource pool is hosted on.
     */
    readonly clusterMoref: string;
    /**
     * default hardware version available to this resource pool.
     */
    readonly hardwareVersion: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly vcenterId: string;
}
/**
 * Provides a data source for a resource pool attached to a vCenter. A resource pool is an essential component of a Provider VDC.
 *
 * > Note 1: this data source requires System Administrator privileges
 *
 * > Note 2: you can create or modify a resource pool using vSphere provider
 *
 * Supported in provider *v3.10+*
 *
 * ## Example Usage
 *
 * ### 1
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const vcenter1 = vcd.getVcenter({
 *     name: "vc1",
 * });
 * const rp1 = vcenter1.then(vcenter1 => vcd.getResourcePool({
 *     name: "resource-pool-for-vcd-01",
 *     vcenterId: vcenter1.id,
 * }));
 * ```
 *
 * ### 2
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const rp1 = vcd.getResourcePool({
 *     name: "common-name",
 *     vcenterId: vcenter1.id,
 * });
 * ```
 *
 * When you receive such error, you can run the script again, but using the resource pool ID instead of the name.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const rp1 = vcd.getResourcePool({
 *     name: "resgroup-241",
 *     vcenterId: vcenter1.id,
 * });
 * ```
 */
export function getResourcePoolOutput(args: GetResourcePoolOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetResourcePoolResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vcd:index/getResourcePool:getResourcePool", {
        "name": args.name,
        "vcenterId": args.vcenterId,
    }, opts);
}

/**
 * A collection of arguments for invoking getResourcePool.
 */
export interface GetResourcePoolOutputArgs {
    /**
     * resource pool name. The name may not be unique within the vCenter. If that happens, you will get an
     * error message with the list of IDs for the pools with the same name, and can subsequently enter the resource pool ID instead of the name.
     * (See Example Usage 2)
     */
    name: pulumi.Input<string>;
    /**
     * ID of the vCenter to which this resource pool belongs.
     */
    vcenterId: pulumi.Input<string>;
}
