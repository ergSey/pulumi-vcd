// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Supported in provider *v3.10+*.
 *
 * Provides a data source to read service provider SAML metadata for an organization.
 * This service provider metadata is used to configure the identity provider.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const my_org = vcd.getOrg({
 *     name: "my-org",
 * });
 * const first = my_org.then(my_org => vcd.getOrgSamlMetadata({
 *     orgId: my_org.id,
 *     fileName: "vcd-metadata.txt",
 * }));
 * ```
 */
export function getOrgSamlMetadata(args: GetOrgSamlMetadataArgs, opts?: pulumi.InvokeOptions): Promise<GetOrgSamlMetadataResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vcd:index/getOrgSamlMetadata:getOrgSamlMetadata", {
        "fileName": args.fileName,
        "orgId": args.orgId,
    }, opts);
}

/**
 * A collection of arguments for invoking getOrgSamlMetadata.
 */
export interface GetOrgSamlMetadataArgs {
    /**
     * name of the file where to store the metadata.
     */
    fileName?: string;
    /**
     * ID of the organization containing the SAML metadata
     */
    orgId: string;
}

/**
 * A collection of values returned by getOrgSamlMetadata.
 */
export interface GetOrgSamlMetadataResult {
    readonly fileName?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * the text of the metadata for this organization.
     */
    readonly metadataText: string;
    readonly orgId: string;
}
/**
 * Supported in provider *v3.10+*.
 *
 * Provides a data source to read service provider SAML metadata for an organization.
 * This service provider metadata is used to configure the identity provider.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const my_org = vcd.getOrg({
 *     name: "my-org",
 * });
 * const first = my_org.then(my_org => vcd.getOrgSamlMetadata({
 *     orgId: my_org.id,
 *     fileName: "vcd-metadata.txt",
 * }));
 * ```
 */
export function getOrgSamlMetadataOutput(args: GetOrgSamlMetadataOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetOrgSamlMetadataResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vcd:index/getOrgSamlMetadata:getOrgSamlMetadata", {
        "fileName": args.fileName,
        "orgId": args.orgId,
    }, opts);
}

/**
 * A collection of arguments for invoking getOrgSamlMetadata.
 */
export interface GetOrgSamlMetadataOutputArgs {
    /**
     * name of the file where to store the metadata.
     */
    fileName?: pulumi.Input<string>;
    /**
     * ID of the organization containing the SAML metadata
     */
    orgId: pulumi.Input<string>;
}
