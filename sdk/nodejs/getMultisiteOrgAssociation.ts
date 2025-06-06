// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a data source to read a VMware Cloud Director Org association information.
 *
 * Supported in provider *v3.13+*
 *
 * ## Example Usage
 *
 * ### 1
 *
 * Retrieving an Org association using the associated Org ID.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const my_org = vcd.getOrg({
 *     name: "my-org",
 * });
 * const org1_org2 = my_org.then(my_org => vcd.getMultisiteOrgAssociation({
 *     orgId: my_org.id,
 *     associatedOrgId: "urn:vcloud:org:3901d87d-1596-4a5a-a74b-57a7313737cf",
 * }));
 * ```
 *
 * ### 2
 *
 * Retrieving an Org association using the association data file.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const my_org = vcd.getOrg({
 *     name: "my-org",
 * });
 * const org1_org2 = my_org.then(my_org => vcd.getMultisiteOrgAssociation({
 *     orgId: my_org.id,
 *     associationDataFile: "remote-org.xml",
 * }));
 * ```
 *
 * ## More information
 *
 * See [Site and Org association](https://www.terraform.io/providers/vmware/vcd/latest/docs/guides/site_org_association) for a broader description
 * of association workflows.
 */
export function getMultisiteOrgAssociation(args: GetMultisiteOrgAssociationArgs, opts?: pulumi.InvokeOptions): Promise<GetMultisiteOrgAssociationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vcd:index/getMultisiteOrgAssociation:getMultisiteOrgAssociation", {
        "associatedOrgId": args.associatedOrgId,
        "associationDataFile": args.associationDataFile,
        "orgId": args.orgId,
    }, opts);
}

/**
 * A collection of arguments for invoking getMultisiteOrgAssociation.
 */
export interface GetMultisiteOrgAssociationArgs {
    /**
     * ID of the remote organization associated with the current one. (Used in alternative to
     * `associatedDataFile`)
     */
    associatedOrgId?: string;
    /**
     * Name of the file containing the data used to associate this Org to another one.
     * (Used when `associatedOrgId` is not known)
     */
    associationDataFile?: string;
    /**
     * The ID of the organization for which we need to collect the data.
     */
    orgId: string;
}

/**
 * A collection of values returned by getMultisiteOrgAssociation.
 */
export interface GetMultisiteOrgAssociationResult {
    readonly associatedOrgId?: string;
    /**
     * The name of the associated Org.
     */
    readonly associatedOrgName: string;
    /**
     * The ID of the associated site.
     */
    readonly associatedSiteId: string;
    readonly associationDataFile?: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly orgId: string;
    /**
     * The status of the association (one of `ASYMMETRIC`, `ACTIVE`, `UNREACHABLE`, `ERROR`)
     */
    readonly status: string;
}
/**
 * Provides a data source to read a VMware Cloud Director Org association information.
 *
 * Supported in provider *v3.13+*
 *
 * ## Example Usage
 *
 * ### 1
 *
 * Retrieving an Org association using the associated Org ID.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const my_org = vcd.getOrg({
 *     name: "my-org",
 * });
 * const org1_org2 = my_org.then(my_org => vcd.getMultisiteOrgAssociation({
 *     orgId: my_org.id,
 *     associatedOrgId: "urn:vcloud:org:3901d87d-1596-4a5a-a74b-57a7313737cf",
 * }));
 * ```
 *
 * ### 2
 *
 * Retrieving an Org association using the association data file.
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const my_org = vcd.getOrg({
 *     name: "my-org",
 * });
 * const org1_org2 = my_org.then(my_org => vcd.getMultisiteOrgAssociation({
 *     orgId: my_org.id,
 *     associationDataFile: "remote-org.xml",
 * }));
 * ```
 *
 * ## More information
 *
 * See [Site and Org association](https://www.terraform.io/providers/vmware/vcd/latest/docs/guides/site_org_association) for a broader description
 * of association workflows.
 */
export function getMultisiteOrgAssociationOutput(args: GetMultisiteOrgAssociationOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetMultisiteOrgAssociationResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vcd:index/getMultisiteOrgAssociation:getMultisiteOrgAssociation", {
        "associatedOrgId": args.associatedOrgId,
        "associationDataFile": args.associationDataFile,
        "orgId": args.orgId,
    }, opts);
}

/**
 * A collection of arguments for invoking getMultisiteOrgAssociation.
 */
export interface GetMultisiteOrgAssociationOutputArgs {
    /**
     * ID of the remote organization associated with the current one. (Used in alternative to
     * `associatedDataFile`)
     */
    associatedOrgId?: pulumi.Input<string>;
    /**
     * Name of the file containing the data used to associate this Org to another one.
     * (Used when `associatedOrgId` is not known)
     */
    associationDataFile?: pulumi.Input<string>;
    /**
     * The ID of the organization for which we need to collect the data.
     */
    orgId: pulumi.Input<string>;
}
