// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a VMware Cloud Director role data source. This can be used to read roles.
 *
 * Supported in provider *v3.3+*
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const vapp_author = vcd.getRole({
 *     org: "my-org",
 *     name: "vApp Author",
 * });
 * ```
 *
 * ## More information
 *
 * See [Roles management](https://www.terraform.io/providers/vmware/vcd/latest/docs/guides/roles_management) for a broader description of how roles and
 * rights work together.
 */
export function getRole(args: GetRoleArgs, opts?: pulumi.InvokeOptions): Promise<GetRoleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vcd:index/getRole:getRole", {
        "name": args.name,
        "org": args.org,
    }, opts);
}

/**
 * A collection of arguments for invoking getRole.
 */
export interface GetRoleArgs {
    /**
     * The name of the role.
     */
    name: string;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
     */
    org?: string;
}

/**
 * A collection of values returned by getRole.
 */
export interface GetRoleResult {
    /**
     * Key used for internationalization.
     */
    readonly bundleKey: string;
    /**
     * A description of the role
     */
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly org?: string;
    /**
     * Whether this role is read-only
     */
    readonly readOnly: boolean;
    /**
     * Set of rights assigned to this role
     */
    readonly rights: string[];
}
/**
 * Provides a VMware Cloud Director role data source. This can be used to read roles.
 *
 * Supported in provider *v3.3+*
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const vapp_author = vcd.getRole({
 *     org: "my-org",
 *     name: "vApp Author",
 * });
 * ```
 *
 * ## More information
 *
 * See [Roles management](https://www.terraform.io/providers/vmware/vcd/latest/docs/guides/roles_management) for a broader description of how roles and
 * rights work together.
 */
export function getRoleOutput(args: GetRoleOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetRoleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vcd:index/getRole:getRole", {
        "name": args.name,
        "org": args.org,
    }, opts);
}

/**
 * A collection of arguments for invoking getRole.
 */
export interface GetRoleOutputArgs {
    /**
     * The name of the role.
     */
    name: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
     */
    org?: pulumi.Input<string>;
}
