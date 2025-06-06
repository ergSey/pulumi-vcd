// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a VMware Cloud Director rights bundle data source. This can be used to read rights bundles.
 *
 * Supported in provider *v3.3+*
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const default_set = vcd.getRightsBundle({
 *     name: "Default Rights Bundle",
 * });
 * ```
 *
 * ## More information
 *
 * See [Roles management](https://www.terraform.io/providers/vmware/vcd/latest/docs/guides/roles_management) for a broader description of how rights bundles and
 * rights work together.
 */
export function getRightsBundle(args: GetRightsBundleArgs, opts?: pulumi.InvokeOptions): Promise<GetRightsBundleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vcd:index/getRightsBundle:getRightsBundle", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getRightsBundle.
 */
export interface GetRightsBundleArgs {
    /**
     * The name of the rights bundle.
     */
    name: string;
}

/**
 * A collection of values returned by getRightsBundle.
 */
export interface GetRightsBundleResult {
    /**
     * Key used for internationalization.
     */
    readonly bundleKey: string;
    /**
     * A description of the rights bundle
     */
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    /**
     * When true, publishes the rights bundle to all tenants
     */
    readonly publishToAllTenants: boolean;
    /**
     * Whether this rights bundle is read-only
     */
    readonly readOnly: boolean;
    /**
     * Set of rights assigned to this role
     */
    readonly rights: string[];
    /**
     * Set of tenants to which this rights bundle gets published. Ignored if `publishToAllTenants` is true.
     */
    readonly tenants: string[];
}
/**
 * Provides a VMware Cloud Director rights bundle data source. This can be used to read rights bundles.
 *
 * Supported in provider *v3.3+*
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const default_set = vcd.getRightsBundle({
 *     name: "Default Rights Bundle",
 * });
 * ```
 *
 * ## More information
 *
 * See [Roles management](https://www.terraform.io/providers/vmware/vcd/latest/docs/guides/roles_management) for a broader description of how rights bundles and
 * rights work together.
 */
export function getRightsBundleOutput(args: GetRightsBundleOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetRightsBundleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vcd:index/getRightsBundle:getRightsBundle", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getRightsBundle.
 */
export interface GetRightsBundleOutputArgs {
    /**
     * The name of the rights bundle.
     */
    name: pulumi.Input<string>;
}
