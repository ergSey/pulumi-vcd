// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

/**
 * Provides a VMware Cloud Director Org data source. An organization can be used to manage catalogs, virtual
 * data centers, and users.
 *
 * Supported in provider *v2.5+*
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
 * const my_org_clone = new vcd.Org("my-org-clone", {
 *     name: "my-org-clone",
 *     fullName: my_org.then(my_org => my_org.fullName),
 *     canPublishCatalogs: my_org.then(my_org => my_org.canPublishCatalogs),
 *     deployedVmQuota: my_org.then(my_org => my_org.deployedVmQuota),
 *     storedVmQuota: my_org.then(my_org => my_org.storedVmQuota),
 *     isEnabled: my_org.then(my_org => my_org.isEnabled),
 *     deleteForce: true,
 *     deleteRecursive: true,
 *     vappLease: {
 *         maximumRuntimeLeaseInSec: my_org.then(my_org => my_org.vappLeases?.[0]?.maximumRuntimeLeaseInSec),
 *         powerOffOnRuntimeLeaseExpiration: my_org.then(my_org => my_org.vappLeases?.[0]?.powerOffOnRuntimeLeaseExpiration),
 *         maximumStorageLeaseInSec: my_org.then(my_org => my_org.vappLeases?.[0]?.maximumStorageLeaseInSec),
 *         deleteOnStorageLeaseExpiration: my_org.then(my_org => my_org.vappLeases?.[0]?.deleteOnStorageLeaseExpiration),
 *     },
 *     vappTemplateLease: {
 *         maximumStorageLeaseInSec: my_org.then(my_org => my_org.vappTemplateLeases?.[0]?.maximumStorageLeaseInSec),
 *         deleteOnStorageLeaseExpiration: my_org.then(my_org => my_org.vappTemplateLeases?.[0]?.deleteOnStorageLeaseExpiration),
 *     },
 * });
 * ```
 *
 * ## vApp Lease
 *
 * The `vappLease` section contains lease parameters for vApps created in the current organization, as defined below:
 *
 * * `maximumRuntimeLeaseInSec` - How long vApps can run before they are automatically stopped (in seconds)
 * * `powerOffOnRuntimeLeaseExpiration` - When true, vApps are powered off when the runtime lease expires. When false, vApps are suspended when the runtime lease expires.
 * * `maximumStorageLeaseInSec` - How long stopped vApps are available before being automatically cleaned up (in seconds)
 * * `deleteOnStorageLeaseExpiration` - If true, storage for a vApp is deleted when the vApp's lease expires. If false, the storage is flagged for deletion, but not deleted.
 *
 * <a id="vapp-template-lease"></a>
 * ## vApp Template Lease
 *
 * The `vappTemplateLease` section contains lease parameters for vApp templates created in the current organization, as defined below:
 *
 * * `maximumStorageLeaseInSec` - How long vApp templates are available before being automatically cleaned up (in seconds)
 * * `deleteOnStorageLeaseExpiration` - If true, storage for a vAppTemplate is deleted when the vAppTemplate lease expires. If false, the storage is flagged for deletion, but not deleted
 *
 * <a id="metadata"></a>
 * ## Metadata
 *
 * The `metadataEntry` (*v3.8+*) is a set of metadata entries that have the following structure:
 *
 * * `key` - Key of this metadata entry.
 * * `value` - Value of this metadata entry.
 * * `type` - Type of this metadata entry. One of: `MetadataStringValue`, `MetadataNumberValue`, `MetadataDateTimeValue`, `MetadataBooleanValue`.
 * * `userAccess` - User access level for this metadata entry. One of: `PRIVATE` (hidden), `READONLY` (read only), `READWRITE` (read/write).
 * * `isSystem` - Domain for this metadata entry. true if it belongs to `SYSTEM`, false if it belongs to `GENERAL`.
 */
export function getOrg(args: GetOrgArgs, opts?: pulumi.InvokeOptions): Promise<GetOrgResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vcd:index/getOrg:getOrg", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getOrg.
 */
export interface GetOrgArgs {
    /**
     * Org name
     */
    name: string;
}

/**
 * A collection of values returned by getOrg.
 */
export interface GetOrgResult {
    /**
     * (*v3.14+*) Contains the account lockout properties of the read organization:
     */
    readonly accountLockouts: outputs.GetOrgAccountLockout[];
    /**
     * True if this organization is allowed to share catalogs.
     */
    readonly canPublishCatalogs: boolean;
    /**
     * (*v3.6+*) True if this organization is allowed to publish external catalogs.
     */
    readonly canPublishExternalCatalogs: boolean;
    /**
     * (*v3.6+*) True if this organization is allowed to subscribe to external catalogs.
     */
    readonly canSubscribeExternalCatalogs: boolean;
    /**
     * Specifies this organization's default for virtual machine boot delay after power on.
     */
    readonly delayAfterPowerOnSeconds: number;
    /**
     * Maximum number of virtual machines that can be deployed simultaneously by a member of this organization.
     */
    readonly deployedVmQuota: number;
    /**
     * Org description.
     */
    readonly description: string;
    /**
     * Org full name
     */
    readonly fullName: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * True if this organization is enabled (allows login and all other operations).
     */
    readonly isEnabled: boolean;
    /**
     * (*v3.11+*) List of catalogs (sorted alphabetically), owned or shared, available to this organization.
     */
    readonly listOfCatalogs: string[];
    /**
     * (*v3.11+*) List of VDCs (sorted alphabetically), owned or shared, available to this organization.
     */
    readonly listOfVdcs: string[];
    /**
     * (Deprecated; *v3.6+*) Use `metadataEntry` instead. Key value map of metadata assigned to this organization.
     *
     * @deprecated Use metadataEntry instead
     */
    readonly metadata: {[key: string]: string};
    /**
     * A set of metadata entries assigned to the organization. See Metadata section for details.
     */
    readonly metadataEntries: outputs.GetOrgMetadataEntry[];
    readonly name: string;
    /**
     * (*v3.11+*) Number of catalogs owned or shared, available to this organization.
     */
    readonly numberOfCatalogs: number;
    /**
     * (*v3.11+*) Number of VDCs owned or shared, available to this organization.
     */
    readonly numberOfVdcs: number;
    /**
     * Maximum number of virtual machines in vApps or vApp templates that can be stored in an undeployed state by a member of this organization.
     */
    readonly storedVmQuota: number;
    /**
     * (*v2.7+*) Defines lease parameters for vApps created in this organization. See vApp Lease below for details.
     */
    readonly vappLeases: outputs.GetOrgVappLease[];
    /**
     * (*v2.7+*) Defines lease parameters for vApp templates created in this organization. See vApp Template Lease below for details.
     */
    readonly vappTemplateLeases: outputs.GetOrgVappTemplateLease[];
}
/**
 * Provides a VMware Cloud Director Org data source. An organization can be used to manage catalogs, virtual
 * data centers, and users.
 *
 * Supported in provider *v2.5+*
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
 * const my_org_clone = new vcd.Org("my-org-clone", {
 *     name: "my-org-clone",
 *     fullName: my_org.then(my_org => my_org.fullName),
 *     canPublishCatalogs: my_org.then(my_org => my_org.canPublishCatalogs),
 *     deployedVmQuota: my_org.then(my_org => my_org.deployedVmQuota),
 *     storedVmQuota: my_org.then(my_org => my_org.storedVmQuota),
 *     isEnabled: my_org.then(my_org => my_org.isEnabled),
 *     deleteForce: true,
 *     deleteRecursive: true,
 *     vappLease: {
 *         maximumRuntimeLeaseInSec: my_org.then(my_org => my_org.vappLeases?.[0]?.maximumRuntimeLeaseInSec),
 *         powerOffOnRuntimeLeaseExpiration: my_org.then(my_org => my_org.vappLeases?.[0]?.powerOffOnRuntimeLeaseExpiration),
 *         maximumStorageLeaseInSec: my_org.then(my_org => my_org.vappLeases?.[0]?.maximumStorageLeaseInSec),
 *         deleteOnStorageLeaseExpiration: my_org.then(my_org => my_org.vappLeases?.[0]?.deleteOnStorageLeaseExpiration),
 *     },
 *     vappTemplateLease: {
 *         maximumStorageLeaseInSec: my_org.then(my_org => my_org.vappTemplateLeases?.[0]?.maximumStorageLeaseInSec),
 *         deleteOnStorageLeaseExpiration: my_org.then(my_org => my_org.vappTemplateLeases?.[0]?.deleteOnStorageLeaseExpiration),
 *     },
 * });
 * ```
 *
 * ## vApp Lease
 *
 * The `vappLease` section contains lease parameters for vApps created in the current organization, as defined below:
 *
 * * `maximumRuntimeLeaseInSec` - How long vApps can run before they are automatically stopped (in seconds)
 * * `powerOffOnRuntimeLeaseExpiration` - When true, vApps are powered off when the runtime lease expires. When false, vApps are suspended when the runtime lease expires.
 * * `maximumStorageLeaseInSec` - How long stopped vApps are available before being automatically cleaned up (in seconds)
 * * `deleteOnStorageLeaseExpiration` - If true, storage for a vApp is deleted when the vApp's lease expires. If false, the storage is flagged for deletion, but not deleted.
 *
 * <a id="vapp-template-lease"></a>
 * ## vApp Template Lease
 *
 * The `vappTemplateLease` section contains lease parameters for vApp templates created in the current organization, as defined below:
 *
 * * `maximumStorageLeaseInSec` - How long vApp templates are available before being automatically cleaned up (in seconds)
 * * `deleteOnStorageLeaseExpiration` - If true, storage for a vAppTemplate is deleted when the vAppTemplate lease expires. If false, the storage is flagged for deletion, but not deleted
 *
 * <a id="metadata"></a>
 * ## Metadata
 *
 * The `metadataEntry` (*v3.8+*) is a set of metadata entries that have the following structure:
 *
 * * `key` - Key of this metadata entry.
 * * `value` - Value of this metadata entry.
 * * `type` - Type of this metadata entry. One of: `MetadataStringValue`, `MetadataNumberValue`, `MetadataDateTimeValue`, `MetadataBooleanValue`.
 * * `userAccess` - User access level for this metadata entry. One of: `PRIVATE` (hidden), `READONLY` (read only), `READWRITE` (read/write).
 * * `isSystem` - Domain for this metadata entry. true if it belongs to `SYSTEM`, false if it belongs to `GENERAL`.
 */
export function getOrgOutput(args: GetOrgOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetOrgResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vcd:index/getOrg:getOrg", {
        "name": args.name,
    }, opts);
}

/**
 * A collection of arguments for invoking getOrg.
 */
export interface GetOrgOutputArgs {
    /**
     * Org name
     */
    name: pulumi.Input<string>;
}
