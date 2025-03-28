// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a data source to read a VMware Cloud Director Site in the context of multi-site operatioos
 *
 * Supported in provider *v3.13+*
 *
 * > Note: this data source requires System Administrator privileges
 *
 * ## Example Usage
 *
 * Note: there is only one site available for each VCD. No ID or name is necessary to identify it.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const currentSite = vcd.getMultisiteSite({});
 * ```
 *
 * ## More information
 *
 * See [Site and Org association](https://www.terraform.io/providers/vmware/vcd/latest/docs/guides/site_org_association) for a broader description
 * of association workflows.
 */
export function getMultisiteSite(opts?: pulumi.InvokeOptions): Promise<GetMultisiteSiteResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vcd:index/getMultisiteSite:getMultisiteSite", {
    }, opts);
}

/**
 * A collection of values returned by getMultisiteSite.
 */
export interface GetMultisiteSiteResult {
    /**
     * An alphabetically sorted list of current associations.
     */
    readonly associations: string[];
    /**
     * An optional description of the site.
     */
    readonly description: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The name of the site, which usually corresponds to its host name.
     */
    readonly name: string;
    /**
     * The number of current associations with other sites.
     */
    readonly numberOfAssociations: number;
}
/**
 * Provides a data source to read a VMware Cloud Director Site in the context of multi-site operatioos
 *
 * Supported in provider *v3.13+*
 *
 * > Note: this data source requires System Administrator privileges
 *
 * ## Example Usage
 *
 * Note: there is only one site available for each VCD. No ID or name is necessary to identify it.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const currentSite = vcd.getMultisiteSite({});
 * ```
 *
 * ## More information
 *
 * See [Site and Org association](https://www.terraform.io/providers/vmware/vcd/latest/docs/guides/site_org_association) for a broader description
 * of association workflows.
 */
export function getMultisiteSiteOutput(opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetMultisiteSiteResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vcd:index/getMultisiteSite:getMultisiteSite", {
    }, opts);
}
