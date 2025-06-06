// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

export class OrgVdcTemplate extends pulumi.CustomResource {
    /**
     * Get an existing OrgVdcTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrgVdcTemplateState, opts?: pulumi.CustomResourceOptions): OrgVdcTemplate {
        return new OrgVdcTemplate(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/orgVdcTemplate:OrgVdcTemplate';

    /**
     * Returns true if the given object is an instance of OrgVdcTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrgVdcTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrgVdcTemplate.__pulumiType;
    }

    /**
     * Allocation model that the VDCs instantiated from this template will use.
     * Must be one of: `AllocationVApp`, `AllocationPool`, `ReservationPool` or  `Flex`
     */
    public readonly allocationModel!: pulumi.Output<string>;
    /**
     * The compute configuration for the VDCs instantiated from this template:
     */
    public readonly computeConfiguration!: pulumi.Output<outputs.OrgVdcTemplateComputeConfiguration>;
    /**
     * Description of the Organization VDC Template, as seen by System administrators
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * VDCs instantiated from this template will create a new Edge Gateway with the provided setup. Required if any `providerVdc` block
     * has defined a `gatewayEdgeClusterId`. This **unique** block has the following properties:
     */
    public readonly edgeGateway!: pulumi.Output<outputs.OrgVdcTemplateEdgeGateway | undefined>;
    /**
     * If `true`, the VDCs instantiated from this template will have Fast provisioning enabled. Defaults to `false`
     */
    public readonly enableFastProvisioning!: pulumi.Output<boolean | undefined>;
    /**
     * If `true`, the VDCs instantiated from this template will have Thin provisioning enabled. Defaults to `false`
     */
    public readonly enableThinProvisioning!: pulumi.Output<boolean | undefined>;
    /**
     * Name to give to the Organization VDC Template, as seen by System administrators
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * If set, specifies the Network pool for the instantiated VDCs. Otherwise, it is automatically chosen
     */
    public readonly networkPoolId!: pulumi.Output<string | undefined>;
    /**
     * Quota for the NICs of the instantiated VDCs. Defaults to 100
     */
    public readonly nicQuota!: pulumi.Output<number | undefined>;
    /**
     * A block that defines a candidate location for the instantiated VDCs. There must be **at least one**, which has the following properties:
     */
    public readonly providerVdcs!: pulumi.Output<outputs.OrgVdcTemplateProviderVdc[]>;
    /**
     * Quota for the provisioned networks of the instantiated VDCs. Defaults to 1000
     */
    public readonly provisionedNetworkQuota!: pulumi.Output<number | undefined>;
    /**
     * A set of Organization IDs that will be able to view and read this VDC template, they can be obtained with
     * [`vcd.Org` data source](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/org)
     */
    public readonly readableByOrgIds!: pulumi.Output<string[] | undefined>;
    /**
     * A block that defines a storage profile that the VDCs instantiated from this template will use. Must be **at least one**, which has the following properties:
     */
    public readonly storageProfiles!: pulumi.Output<outputs.OrgVdcTemplateStorageProfile[]>;
    /**
     * Description of the Organization VDC Template, as seen by the allowed tenants
     */
    public readonly tenantDescription!: pulumi.Output<string | undefined>;
    /**
     * Name to give to the Organization VDC Template, as seen by the allowed tenants
     */
    public readonly tenantName!: pulumi.Output<string>;
    /**
     * Quota for the VMs of the instantiated VDCs. 0 means unlimited. Defaults to 0
     */
    public readonly vmQuota!: pulumi.Output<number | undefined>;

    /**
     * Create a OrgVdcTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrgVdcTemplateArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrgVdcTemplateArgs | OrgVdcTemplateState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OrgVdcTemplateState | undefined;
            resourceInputs["allocationModel"] = state ? state.allocationModel : undefined;
            resourceInputs["computeConfiguration"] = state ? state.computeConfiguration : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["edgeGateway"] = state ? state.edgeGateway : undefined;
            resourceInputs["enableFastProvisioning"] = state ? state.enableFastProvisioning : undefined;
            resourceInputs["enableThinProvisioning"] = state ? state.enableThinProvisioning : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkPoolId"] = state ? state.networkPoolId : undefined;
            resourceInputs["nicQuota"] = state ? state.nicQuota : undefined;
            resourceInputs["providerVdcs"] = state ? state.providerVdcs : undefined;
            resourceInputs["provisionedNetworkQuota"] = state ? state.provisionedNetworkQuota : undefined;
            resourceInputs["readableByOrgIds"] = state ? state.readableByOrgIds : undefined;
            resourceInputs["storageProfiles"] = state ? state.storageProfiles : undefined;
            resourceInputs["tenantDescription"] = state ? state.tenantDescription : undefined;
            resourceInputs["tenantName"] = state ? state.tenantName : undefined;
            resourceInputs["vmQuota"] = state ? state.vmQuota : undefined;
        } else {
            const args = argsOrState as OrgVdcTemplateArgs | undefined;
            if ((!args || args.allocationModel === undefined) && !opts.urn) {
                throw new Error("Missing required property 'allocationModel'");
            }
            if ((!args || args.computeConfiguration === undefined) && !opts.urn) {
                throw new Error("Missing required property 'computeConfiguration'");
            }
            if ((!args || args.providerVdcs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'providerVdcs'");
            }
            if ((!args || args.storageProfiles === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storageProfiles'");
            }
            if ((!args || args.tenantName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tenantName'");
            }
            resourceInputs["allocationModel"] = args ? args.allocationModel : undefined;
            resourceInputs["computeConfiguration"] = args ? args.computeConfiguration : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["edgeGateway"] = args ? args.edgeGateway : undefined;
            resourceInputs["enableFastProvisioning"] = args ? args.enableFastProvisioning : undefined;
            resourceInputs["enableThinProvisioning"] = args ? args.enableThinProvisioning : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkPoolId"] = args ? args.networkPoolId : undefined;
            resourceInputs["nicQuota"] = args ? args.nicQuota : undefined;
            resourceInputs["providerVdcs"] = args ? args.providerVdcs : undefined;
            resourceInputs["provisionedNetworkQuota"] = args ? args.provisionedNetworkQuota : undefined;
            resourceInputs["readableByOrgIds"] = args ? args.readableByOrgIds : undefined;
            resourceInputs["storageProfiles"] = args ? args.storageProfiles : undefined;
            resourceInputs["tenantDescription"] = args ? args.tenantDescription : undefined;
            resourceInputs["tenantName"] = args ? args.tenantName : undefined;
            resourceInputs["vmQuota"] = args ? args.vmQuota : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OrgVdcTemplate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OrgVdcTemplate resources.
 */
export interface OrgVdcTemplateState {
    /**
     * Allocation model that the VDCs instantiated from this template will use.
     * Must be one of: `AllocationVApp`, `AllocationPool`, `ReservationPool` or  `Flex`
     */
    allocationModel?: pulumi.Input<string>;
    /**
     * The compute configuration for the VDCs instantiated from this template:
     */
    computeConfiguration?: pulumi.Input<inputs.OrgVdcTemplateComputeConfiguration>;
    /**
     * Description of the Organization VDC Template, as seen by System administrators
     */
    description?: pulumi.Input<string>;
    /**
     * VDCs instantiated from this template will create a new Edge Gateway with the provided setup. Required if any `providerVdc` block
     * has defined a `gatewayEdgeClusterId`. This **unique** block has the following properties:
     */
    edgeGateway?: pulumi.Input<inputs.OrgVdcTemplateEdgeGateway>;
    /**
     * If `true`, the VDCs instantiated from this template will have Fast provisioning enabled. Defaults to `false`
     */
    enableFastProvisioning?: pulumi.Input<boolean>;
    /**
     * If `true`, the VDCs instantiated from this template will have Thin provisioning enabled. Defaults to `false`
     */
    enableThinProvisioning?: pulumi.Input<boolean>;
    /**
     * Name to give to the Organization VDC Template, as seen by System administrators
     */
    name?: pulumi.Input<string>;
    /**
     * If set, specifies the Network pool for the instantiated VDCs. Otherwise, it is automatically chosen
     */
    networkPoolId?: pulumi.Input<string>;
    /**
     * Quota for the NICs of the instantiated VDCs. Defaults to 100
     */
    nicQuota?: pulumi.Input<number>;
    /**
     * A block that defines a candidate location for the instantiated VDCs. There must be **at least one**, which has the following properties:
     */
    providerVdcs?: pulumi.Input<pulumi.Input<inputs.OrgVdcTemplateProviderVdc>[]>;
    /**
     * Quota for the provisioned networks of the instantiated VDCs. Defaults to 1000
     */
    provisionedNetworkQuota?: pulumi.Input<number>;
    /**
     * A set of Organization IDs that will be able to view and read this VDC template, they can be obtained with
     * [`vcd.Org` data source](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/org)
     */
    readableByOrgIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A block that defines a storage profile that the VDCs instantiated from this template will use. Must be **at least one**, which has the following properties:
     */
    storageProfiles?: pulumi.Input<pulumi.Input<inputs.OrgVdcTemplateStorageProfile>[]>;
    /**
     * Description of the Organization VDC Template, as seen by the allowed tenants
     */
    tenantDescription?: pulumi.Input<string>;
    /**
     * Name to give to the Organization VDC Template, as seen by the allowed tenants
     */
    tenantName?: pulumi.Input<string>;
    /**
     * Quota for the VMs of the instantiated VDCs. 0 means unlimited. Defaults to 0
     */
    vmQuota?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a OrgVdcTemplate resource.
 */
export interface OrgVdcTemplateArgs {
    /**
     * Allocation model that the VDCs instantiated from this template will use.
     * Must be one of: `AllocationVApp`, `AllocationPool`, `ReservationPool` or  `Flex`
     */
    allocationModel: pulumi.Input<string>;
    /**
     * The compute configuration for the VDCs instantiated from this template:
     */
    computeConfiguration: pulumi.Input<inputs.OrgVdcTemplateComputeConfiguration>;
    /**
     * Description of the Organization VDC Template, as seen by System administrators
     */
    description?: pulumi.Input<string>;
    /**
     * VDCs instantiated from this template will create a new Edge Gateway with the provided setup. Required if any `providerVdc` block
     * has defined a `gatewayEdgeClusterId`. This **unique** block has the following properties:
     */
    edgeGateway?: pulumi.Input<inputs.OrgVdcTemplateEdgeGateway>;
    /**
     * If `true`, the VDCs instantiated from this template will have Fast provisioning enabled. Defaults to `false`
     */
    enableFastProvisioning?: pulumi.Input<boolean>;
    /**
     * If `true`, the VDCs instantiated from this template will have Thin provisioning enabled. Defaults to `false`
     */
    enableThinProvisioning?: pulumi.Input<boolean>;
    /**
     * Name to give to the Organization VDC Template, as seen by System administrators
     */
    name?: pulumi.Input<string>;
    /**
     * If set, specifies the Network pool for the instantiated VDCs. Otherwise, it is automatically chosen
     */
    networkPoolId?: pulumi.Input<string>;
    /**
     * Quota for the NICs of the instantiated VDCs. Defaults to 100
     */
    nicQuota?: pulumi.Input<number>;
    /**
     * A block that defines a candidate location for the instantiated VDCs. There must be **at least one**, which has the following properties:
     */
    providerVdcs: pulumi.Input<pulumi.Input<inputs.OrgVdcTemplateProviderVdc>[]>;
    /**
     * Quota for the provisioned networks of the instantiated VDCs. Defaults to 1000
     */
    provisionedNetworkQuota?: pulumi.Input<number>;
    /**
     * A set of Organization IDs that will be able to view and read this VDC template, they can be obtained with
     * [`vcd.Org` data source](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/org)
     */
    readableByOrgIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A block that defines a storage profile that the VDCs instantiated from this template will use. Must be **at least one**, which has the following properties:
     */
    storageProfiles: pulumi.Input<pulumi.Input<inputs.OrgVdcTemplateStorageProfile>[]>;
    /**
     * Description of the Organization VDC Template, as seen by the allowed tenants
     */
    tenantDescription?: pulumi.Input<string>;
    /**
     * Name to give to the Organization VDC Template, as seen by the allowed tenants
     */
    tenantName: pulumi.Input<string>;
    /**
     * Quota for the VMs of the instantiated VDCs. 0 means unlimited. Defaults to 0
     */
    vmQuota?: pulumi.Input<number>;
}
