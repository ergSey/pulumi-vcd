// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

/**
 * Provides a data source to read Access Control details from a Catalog in VMware Cloud Director.
 *
 * > **Note:** Access control reads run in tenant context, meaning that, even if the user is a system administrator,
 * in every request it uses headers items that define the tenant context as restricted to the organization to which the Catalog belongs.
 *
 * Supported in provider *v3.14+*
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const catalog = vcd.getCatalog({
 *     name: "my-catalog",
 * });
 * const ac = catalog.then(catalog => vcd.getCatalogAccessControl({
 *     catalogId: catalog.id,
 * }));
 * export const sharedWith = ac.then(ac => ac.sharedWiths);
 * ```
 */
export function getCatalogAccessControl(args: GetCatalogAccessControlArgs, opts?: pulumi.InvokeOptions): Promise<GetCatalogAccessControlResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vcd:index/getCatalogAccessControl:getCatalogAccessControl", {
        "catalogId": args.catalogId,
        "org": args.org,
    }, opts);
}

/**
 * A collection of arguments for invoking getCatalogAccessControl.
 */
export interface GetCatalogAccessControlArgs {
    /**
     * A unique identifier for the Catalog.
     */
    catalogId: string;
    /**
     * The name of organization to which the Catalog belongs. Optional if defined at provider level.
     */
    org?: string;
}

/**
 * A collection of values returned by getCatalogAccessControl.
 */
export interface GetCatalogAccessControlResult {
    readonly catalogId: string;
    readonly everyoneAccessLevel: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly org?: string;
    readonly readOnlySharedWithAllOrgs: boolean;
    readonly sharedWithEveryone: boolean;
    readonly sharedWiths: outputs.GetCatalogAccessControlSharedWith[];
}
/**
 * Provides a data source to read Access Control details from a Catalog in VMware Cloud Director.
 *
 * > **Note:** Access control reads run in tenant context, meaning that, even if the user is a system administrator,
 * in every request it uses headers items that define the tenant context as restricted to the organization to which the Catalog belongs.
 *
 * Supported in provider *v3.14+*
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const catalog = vcd.getCatalog({
 *     name: "my-catalog",
 * });
 * const ac = catalog.then(catalog => vcd.getCatalogAccessControl({
 *     catalogId: catalog.id,
 * }));
 * export const sharedWith = ac.then(ac => ac.sharedWiths);
 * ```
 */
export function getCatalogAccessControlOutput(args: GetCatalogAccessControlOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetCatalogAccessControlResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vcd:index/getCatalogAccessControl:getCatalogAccessControl", {
        "catalogId": args.catalogId,
        "org": args.org,
    }, opts);
}

/**
 * A collection of arguments for invoking getCatalogAccessControl.
 */
export interface GetCatalogAccessControlOutputArgs {
    /**
     * A unique identifier for the Catalog.
     */
    catalogId: pulumi.Input<string>;
    /**
     * The name of organization to which the Catalog belongs. Optional if defined at provider level.
     */
    org?: pulumi.Input<string>;
}
