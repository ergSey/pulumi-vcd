// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

export class OrgVdc extends pulumi.CustomResource {
    /**
     * Get an existing OrgVdc resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrgVdcState, opts?: pulumi.CustomResourceOptions): OrgVdc {
        return new OrgVdc(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/orgVdc:OrgVdc';

    /**
     * Returns true if the given object is an instance of OrgVdc.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrgVdc {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrgVdc.__pulumiType;
    }

    /**
     * The allocation model used by this VDC; must be one of 
     * * AllocationVApp ("Pay as you go")
     * * AllocationPool ("Allocation pool")
     * * ReservationPool ("Reservation pool")
     * * Flex ("Flex") (*v2.7+*, *VCD 9.7+*)
     */
    public readonly allocationModel!: pulumi.Output<string>;
    /**
     * Set to false to disallow creation of the VDC if the `allocationModel` is AllocationPool or ReservationPool and the ComputeCapacity you specified is greater than what the backing Provider VDC can supply. Default is true.
     */
    public readonly allowOverCommit!: pulumi.Output<boolean>;
    /**
     * The compute capacity allocated to this VDC.  See Compute Capacity below for details.
     */
    public readonly computeCapacity!: pulumi.Output<outputs.OrgVdcComputeCapacity>;
    /**
     * Percentage of allocated CPU resources guaranteed to vApps deployed in this VDC. For example, if this value is 0.75, then 75% of allocated resources are guaranteed. Required when `allocationModel` is AllocationVApp, AllocationPool or Flex. If left empty, VCD sets a value.
     */
    public readonly cpuGuaranteed!: pulumi.Output<number>;
    /**
     * Specifies the clock frequency, in Megahertz, for any virtual CPU that is allocated to a VM. A VM with 2 vCPUs will consume twice as much of this value. Ignored for ReservationPool. Required when `allocationModel` is AllocationVApp, AllocationPool or Flex, and may not be less than 256 MHz. Defaults to 1000 MHz if value isn't provided.
     */
    public readonly cpuSpeed!: pulumi.Output<number>;
    /**
     * ID of the default Compute Policy for this VDC. It can be a VM Sizing Policy, a VM Placement Policy or a vGPU Policy.
     */
    public readonly defaultComputePolicyId!: pulumi.Output<string>;
    /**
     * ID of the default Compute Policy for this VDC. It can be a VM Sizing Policy, a VM Placement Policy or a vGPU Policy. Deprecated in favor of `defaultComputePolicyId`.
     *
     * @deprecated Use `defaultComputePolicyId` attribute instead, which can support VM Sizing Policies, VM Placement Policies and vGPU Policies
     */
    public readonly defaultVmSizingPolicyId!: pulumi.Output<string>;
    /**
     * When destroying use `delete_force=true` to remove a VDC and any objects it contains, regardless of their state. Default is `false`
     */
    public readonly deleteForce!: pulumi.Output<boolean | undefined>;
    /**
     * When destroying use `delete_recursive=true` to remove the VDC and any objects it contains that are in a state that normally allows removal. Default is `false`
     */
    public readonly deleteRecursive!: pulumi.Output<boolean | undefined>;
    /**
     * VDC friendly description
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * An ID of NSX-T Edge Cluster which
     * should provide vApp Networking Services or DHCP for isolated networks. Can be looked up using
     * `vcd.getNsxtEdgeCluster` data source. This field is **deprecated** in favor of
     * [`vcd.OrgVdcNsxtNetworkProfile`](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/org_vdc_nsxt_network_profile).
     *
     * @deprecated Please use 'vcd_org_vdc_nsxt_network_profile' resource to manage Edge Cluster and Segment Profile Templates
     */
    public readonly edgeClusterId!: pulumi.Output<string>;
    /**
     * Indicates if the Flex VDC should be elastic. Required with the Flex allocation model.
     */
    public readonly elasticity!: pulumi.Output<boolean>;
    /**
     * Request fast provisioning. Request will be honored only if the underlying datastore supports it. Fast provisioning can reduce the time it takes to create virtual machines by using vSphere linked clones. If you disable fast provisioning, all provisioning operations will result in full clones.
     */
    public readonly enableFastProvisioning!: pulumi.Output<boolean | undefined>;
    /**
     * Enables or disables the NSX-V distributed firewall.
     *
     * <a id="storageprofile"></a>
     */
    public readonly enableNsxvDistributedFirewall!: pulumi.Output<boolean>;
    /**
     * Boolean to request thin provisioning. Request will be honored only if the underlying data store supports it. Thin provisioning saves storage space by committing it on demand. This allows over-allocation of storage.
     */
    public readonly enableThinProvisioning!: pulumi.Output<boolean | undefined>;
    /**
     * If true, discovery of vCenter VMs is enabled for resource pools backing this VDC. If false, discovery is disabled. If left unspecified, the actual behaviour depends on enablement at the organization level and at the system level.
     */
    public readonly enableVmDiscovery!: pulumi.Output<boolean | undefined>;
    /**
     * True if this VDC is enabled for use by the organization VDCs. Default is true.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Indicates if the Flex VDC should include memory overhead into its accounting for admission control. Required with the Flex allocation model. `memoryGuaranteed` must also be specified together with this parameter.
     */
    public readonly includeVmMemoryOverhead!: pulumi.Output<boolean>;
    /**
     * Percentage of allocated memory resources guaranteed to vApps deployed in this VDC. For example, if this value is 0.75, then 75% of allocated resources are guaranteed. Required when `allocationModel` is AllocationVApp, AllocationPool or Flex. When Allocation model is AllocationPool minimum value is 0.2. If left empty, VCD sets a value.
     */
    public readonly memoryGuaranteed!: pulumi.Output<number>;
    /**
     * Use `metadataEntry` instead. Key value map of metadata to assign to this VDC
     *
     * @deprecated Use metadataEntry instead
     */
    public readonly metadata!: pulumi.Output<{[key: string]: string}>;
    /**
     * A set of metadata entries to assign. See Metadata section for details.
     */
    public readonly metadataEntries!: pulumi.Output<outputs.OrgVdcMetadataEntry[]>;
    /**
     * VDC name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Reference to a network pool in the Provider VDC. Required if this VDC will contain routed or isolated networks.
     */
    public readonly networkPoolName!: pulumi.Output<string | undefined>;
    /**
     * Maximum number of network objects that can be deployed in this VDC. Defaults to 0, which means no networks can be deployed.
     */
    public readonly networkQuota!: pulumi.Output<number | undefined>;
    /**
     * Maximum number of virtual NICs allowed in this VDC. Defaults to 0, which specifies an unlimited number.
     */
    public readonly nicQuota!: pulumi.Output<number | undefined>;
    /**
     * Organization to create the VDC in, optional if defined at provider level
     */
    public readonly org!: pulumi.Output<string | undefined>;
    /**
     * Name of the Provider VDC from which this organization VDC is provisioned.
     */
    public readonly providerVdcName!: pulumi.Output<string>;
    /**
     * Storage profiles supported by this VDC.  See Storage Profile below for details.
     */
    public readonly storageProfiles!: pulumi.Output<outputs.OrgVdcStorageProfile[]>;
    /**
     * Set of IDs of VM Placement policies that are assigned to this VDC. This field requires `defaultComputePolicyId` to be configured together.
     */
    public readonly vmPlacementPolicyIds!: pulumi.Output<string[]>;
    /**
     * The maximum number of VMs that can be created in this VDC. Includes deployed and undeployed VMs in vApps and vApp templates. Defaults to 0, which specifies an unlimited number.
     */
    public readonly vmQuota!: pulumi.Output<number | undefined>;
    /**
     * Set of IDs of VM Sizing policies that are assigned to this VDC. This field requires `defaultComputePolicyId` to be configured together.
     */
    public readonly vmSizingPolicyIds!: pulumi.Output<string[]>;
    /**
     * Set of IDs of VM vGPU policies that are assigned to this VDC. This field requires `defaultComputePolicyId` to be configured together.
     */
    public readonly vmVgpuPolicyIds!: pulumi.Output<string[]>;

    /**
     * Create a OrgVdc resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrgVdcArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrgVdcArgs | OrgVdcState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OrgVdcState | undefined;
            resourceInputs["allocationModel"] = state ? state.allocationModel : undefined;
            resourceInputs["allowOverCommit"] = state ? state.allowOverCommit : undefined;
            resourceInputs["computeCapacity"] = state ? state.computeCapacity : undefined;
            resourceInputs["cpuGuaranteed"] = state ? state.cpuGuaranteed : undefined;
            resourceInputs["cpuSpeed"] = state ? state.cpuSpeed : undefined;
            resourceInputs["defaultComputePolicyId"] = state ? state.defaultComputePolicyId : undefined;
            resourceInputs["defaultVmSizingPolicyId"] = state ? state.defaultVmSizingPolicyId : undefined;
            resourceInputs["deleteForce"] = state ? state.deleteForce : undefined;
            resourceInputs["deleteRecursive"] = state ? state.deleteRecursive : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["edgeClusterId"] = state ? state.edgeClusterId : undefined;
            resourceInputs["elasticity"] = state ? state.elasticity : undefined;
            resourceInputs["enableFastProvisioning"] = state ? state.enableFastProvisioning : undefined;
            resourceInputs["enableNsxvDistributedFirewall"] = state ? state.enableNsxvDistributedFirewall : undefined;
            resourceInputs["enableThinProvisioning"] = state ? state.enableThinProvisioning : undefined;
            resourceInputs["enableVmDiscovery"] = state ? state.enableVmDiscovery : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["includeVmMemoryOverhead"] = state ? state.includeVmMemoryOverhead : undefined;
            resourceInputs["memoryGuaranteed"] = state ? state.memoryGuaranteed : undefined;
            resourceInputs["metadata"] = state ? state.metadata : undefined;
            resourceInputs["metadataEntries"] = state ? state.metadataEntries : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkPoolName"] = state ? state.networkPoolName : undefined;
            resourceInputs["networkQuota"] = state ? state.networkQuota : undefined;
            resourceInputs["nicQuota"] = state ? state.nicQuota : undefined;
            resourceInputs["org"] = state ? state.org : undefined;
            resourceInputs["providerVdcName"] = state ? state.providerVdcName : undefined;
            resourceInputs["storageProfiles"] = state ? state.storageProfiles : undefined;
            resourceInputs["vmPlacementPolicyIds"] = state ? state.vmPlacementPolicyIds : undefined;
            resourceInputs["vmQuota"] = state ? state.vmQuota : undefined;
            resourceInputs["vmSizingPolicyIds"] = state ? state.vmSizingPolicyIds : undefined;
            resourceInputs["vmVgpuPolicyIds"] = state ? state.vmVgpuPolicyIds : undefined;
        } else {
            const args = argsOrState as OrgVdcArgs | undefined;
            if ((!args || args.allocationModel === undefined) && !opts.urn) {
                throw new Error("Missing required property 'allocationModel'");
            }
            if ((!args || args.computeCapacity === undefined) && !opts.urn) {
                throw new Error("Missing required property 'computeCapacity'");
            }
            if ((!args || args.providerVdcName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'providerVdcName'");
            }
            if ((!args || args.storageProfiles === undefined) && !opts.urn) {
                throw new Error("Missing required property 'storageProfiles'");
            }
            resourceInputs["allocationModel"] = args ? args.allocationModel : undefined;
            resourceInputs["allowOverCommit"] = args ? args.allowOverCommit : undefined;
            resourceInputs["computeCapacity"] = args ? args.computeCapacity : undefined;
            resourceInputs["cpuGuaranteed"] = args ? args.cpuGuaranteed : undefined;
            resourceInputs["cpuSpeed"] = args ? args.cpuSpeed : undefined;
            resourceInputs["defaultComputePolicyId"] = args ? args.defaultComputePolicyId : undefined;
            resourceInputs["defaultVmSizingPolicyId"] = args ? args.defaultVmSizingPolicyId : undefined;
            resourceInputs["deleteForce"] = args ? args.deleteForce : undefined;
            resourceInputs["deleteRecursive"] = args ? args.deleteRecursive : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["edgeClusterId"] = args ? args.edgeClusterId : undefined;
            resourceInputs["elasticity"] = args ? args.elasticity : undefined;
            resourceInputs["enableFastProvisioning"] = args ? args.enableFastProvisioning : undefined;
            resourceInputs["enableNsxvDistributedFirewall"] = args ? args.enableNsxvDistributedFirewall : undefined;
            resourceInputs["enableThinProvisioning"] = args ? args.enableThinProvisioning : undefined;
            resourceInputs["enableVmDiscovery"] = args ? args.enableVmDiscovery : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["includeVmMemoryOverhead"] = args ? args.includeVmMemoryOverhead : undefined;
            resourceInputs["memoryGuaranteed"] = args ? args.memoryGuaranteed : undefined;
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["metadataEntries"] = args ? args.metadataEntries : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkPoolName"] = args ? args.networkPoolName : undefined;
            resourceInputs["networkQuota"] = args ? args.networkQuota : undefined;
            resourceInputs["nicQuota"] = args ? args.nicQuota : undefined;
            resourceInputs["org"] = args ? args.org : undefined;
            resourceInputs["providerVdcName"] = args ? args.providerVdcName : undefined;
            resourceInputs["storageProfiles"] = args ? args.storageProfiles : undefined;
            resourceInputs["vmPlacementPolicyIds"] = args ? args.vmPlacementPolicyIds : undefined;
            resourceInputs["vmQuota"] = args ? args.vmQuota : undefined;
            resourceInputs["vmSizingPolicyIds"] = args ? args.vmSizingPolicyIds : undefined;
            resourceInputs["vmVgpuPolicyIds"] = args ? args.vmVgpuPolicyIds : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OrgVdc.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OrgVdc resources.
 */
export interface OrgVdcState {
    /**
     * The allocation model used by this VDC; must be one of 
     * * AllocationVApp ("Pay as you go")
     * * AllocationPool ("Allocation pool")
     * * ReservationPool ("Reservation pool")
     * * Flex ("Flex") (*v2.7+*, *VCD 9.7+*)
     */
    allocationModel?: pulumi.Input<string>;
    /**
     * Set to false to disallow creation of the VDC if the `allocationModel` is AllocationPool or ReservationPool and the ComputeCapacity you specified is greater than what the backing Provider VDC can supply. Default is true.
     */
    allowOverCommit?: pulumi.Input<boolean>;
    /**
     * The compute capacity allocated to this VDC.  See Compute Capacity below for details.
     */
    computeCapacity?: pulumi.Input<inputs.OrgVdcComputeCapacity>;
    /**
     * Percentage of allocated CPU resources guaranteed to vApps deployed in this VDC. For example, if this value is 0.75, then 75% of allocated resources are guaranteed. Required when `allocationModel` is AllocationVApp, AllocationPool or Flex. If left empty, VCD sets a value.
     */
    cpuGuaranteed?: pulumi.Input<number>;
    /**
     * Specifies the clock frequency, in Megahertz, for any virtual CPU that is allocated to a VM. A VM with 2 vCPUs will consume twice as much of this value. Ignored for ReservationPool. Required when `allocationModel` is AllocationVApp, AllocationPool or Flex, and may not be less than 256 MHz. Defaults to 1000 MHz if value isn't provided.
     */
    cpuSpeed?: pulumi.Input<number>;
    /**
     * ID of the default Compute Policy for this VDC. It can be a VM Sizing Policy, a VM Placement Policy or a vGPU Policy.
     */
    defaultComputePolicyId?: pulumi.Input<string>;
    /**
     * ID of the default Compute Policy for this VDC. It can be a VM Sizing Policy, a VM Placement Policy or a vGPU Policy. Deprecated in favor of `defaultComputePolicyId`.
     *
     * @deprecated Use `defaultComputePolicyId` attribute instead, which can support VM Sizing Policies, VM Placement Policies and vGPU Policies
     */
    defaultVmSizingPolicyId?: pulumi.Input<string>;
    /**
     * When destroying use `delete_force=true` to remove a VDC and any objects it contains, regardless of their state. Default is `false`
     */
    deleteForce?: pulumi.Input<boolean>;
    /**
     * When destroying use `delete_recursive=true` to remove the VDC and any objects it contains that are in a state that normally allows removal. Default is `false`
     */
    deleteRecursive?: pulumi.Input<boolean>;
    /**
     * VDC friendly description
     */
    description?: pulumi.Input<string>;
    /**
     * An ID of NSX-T Edge Cluster which
     * should provide vApp Networking Services or DHCP for isolated networks. Can be looked up using
     * `vcd.getNsxtEdgeCluster` data source. This field is **deprecated** in favor of
     * [`vcd.OrgVdcNsxtNetworkProfile`](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/org_vdc_nsxt_network_profile).
     *
     * @deprecated Please use 'vcd_org_vdc_nsxt_network_profile' resource to manage Edge Cluster and Segment Profile Templates
     */
    edgeClusterId?: pulumi.Input<string>;
    /**
     * Indicates if the Flex VDC should be elastic. Required with the Flex allocation model.
     */
    elasticity?: pulumi.Input<boolean>;
    /**
     * Request fast provisioning. Request will be honored only if the underlying datastore supports it. Fast provisioning can reduce the time it takes to create virtual machines by using vSphere linked clones. If you disable fast provisioning, all provisioning operations will result in full clones.
     */
    enableFastProvisioning?: pulumi.Input<boolean>;
    /**
     * Enables or disables the NSX-V distributed firewall.
     *
     * <a id="storageprofile"></a>
     */
    enableNsxvDistributedFirewall?: pulumi.Input<boolean>;
    /**
     * Boolean to request thin provisioning. Request will be honored only if the underlying data store supports it. Thin provisioning saves storage space by committing it on demand. This allows over-allocation of storage.
     */
    enableThinProvisioning?: pulumi.Input<boolean>;
    /**
     * If true, discovery of vCenter VMs is enabled for resource pools backing this VDC. If false, discovery is disabled. If left unspecified, the actual behaviour depends on enablement at the organization level and at the system level.
     */
    enableVmDiscovery?: pulumi.Input<boolean>;
    /**
     * True if this VDC is enabled for use by the organization VDCs. Default is true.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Indicates if the Flex VDC should include memory overhead into its accounting for admission control. Required with the Flex allocation model. `memoryGuaranteed` must also be specified together with this parameter.
     */
    includeVmMemoryOverhead?: pulumi.Input<boolean>;
    /**
     * Percentage of allocated memory resources guaranteed to vApps deployed in this VDC. For example, if this value is 0.75, then 75% of allocated resources are guaranteed. Required when `allocationModel` is AllocationVApp, AllocationPool or Flex. When Allocation model is AllocationPool minimum value is 0.2. If left empty, VCD sets a value.
     */
    memoryGuaranteed?: pulumi.Input<number>;
    /**
     * Use `metadataEntry` instead. Key value map of metadata to assign to this VDC
     *
     * @deprecated Use metadataEntry instead
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A set of metadata entries to assign. See Metadata section for details.
     */
    metadataEntries?: pulumi.Input<pulumi.Input<inputs.OrgVdcMetadataEntry>[]>;
    /**
     * VDC name
     */
    name?: pulumi.Input<string>;
    /**
     * Reference to a network pool in the Provider VDC. Required if this VDC will contain routed or isolated networks.
     */
    networkPoolName?: pulumi.Input<string>;
    /**
     * Maximum number of network objects that can be deployed in this VDC. Defaults to 0, which means no networks can be deployed.
     */
    networkQuota?: pulumi.Input<number>;
    /**
     * Maximum number of virtual NICs allowed in this VDC. Defaults to 0, which specifies an unlimited number.
     */
    nicQuota?: pulumi.Input<number>;
    /**
     * Organization to create the VDC in, optional if defined at provider level
     */
    org?: pulumi.Input<string>;
    /**
     * Name of the Provider VDC from which this organization VDC is provisioned.
     */
    providerVdcName?: pulumi.Input<string>;
    /**
     * Storage profiles supported by this VDC.  See Storage Profile below for details.
     */
    storageProfiles?: pulumi.Input<pulumi.Input<inputs.OrgVdcStorageProfile>[]>;
    /**
     * Set of IDs of VM Placement policies that are assigned to this VDC. This field requires `defaultComputePolicyId` to be configured together.
     */
    vmPlacementPolicyIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The maximum number of VMs that can be created in this VDC. Includes deployed and undeployed VMs in vApps and vApp templates. Defaults to 0, which specifies an unlimited number.
     */
    vmQuota?: pulumi.Input<number>;
    /**
     * Set of IDs of VM Sizing policies that are assigned to this VDC. This field requires `defaultComputePolicyId` to be configured together.
     */
    vmSizingPolicyIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Set of IDs of VM vGPU policies that are assigned to this VDC. This field requires `defaultComputePolicyId` to be configured together.
     */
    vmVgpuPolicyIds?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a OrgVdc resource.
 */
export interface OrgVdcArgs {
    /**
     * The allocation model used by this VDC; must be one of 
     * * AllocationVApp ("Pay as you go")
     * * AllocationPool ("Allocation pool")
     * * ReservationPool ("Reservation pool")
     * * Flex ("Flex") (*v2.7+*, *VCD 9.7+*)
     */
    allocationModel: pulumi.Input<string>;
    /**
     * Set to false to disallow creation of the VDC if the `allocationModel` is AllocationPool or ReservationPool and the ComputeCapacity you specified is greater than what the backing Provider VDC can supply. Default is true.
     */
    allowOverCommit?: pulumi.Input<boolean>;
    /**
     * The compute capacity allocated to this VDC.  See Compute Capacity below for details.
     */
    computeCapacity: pulumi.Input<inputs.OrgVdcComputeCapacity>;
    /**
     * Percentage of allocated CPU resources guaranteed to vApps deployed in this VDC. For example, if this value is 0.75, then 75% of allocated resources are guaranteed. Required when `allocationModel` is AllocationVApp, AllocationPool or Flex. If left empty, VCD sets a value.
     */
    cpuGuaranteed?: pulumi.Input<number>;
    /**
     * Specifies the clock frequency, in Megahertz, for any virtual CPU that is allocated to a VM. A VM with 2 vCPUs will consume twice as much of this value. Ignored for ReservationPool. Required when `allocationModel` is AllocationVApp, AllocationPool or Flex, and may not be less than 256 MHz. Defaults to 1000 MHz if value isn't provided.
     */
    cpuSpeed?: pulumi.Input<number>;
    /**
     * ID of the default Compute Policy for this VDC. It can be a VM Sizing Policy, a VM Placement Policy or a vGPU Policy.
     */
    defaultComputePolicyId?: pulumi.Input<string>;
    /**
     * ID of the default Compute Policy for this VDC. It can be a VM Sizing Policy, a VM Placement Policy or a vGPU Policy. Deprecated in favor of `defaultComputePolicyId`.
     *
     * @deprecated Use `defaultComputePolicyId` attribute instead, which can support VM Sizing Policies, VM Placement Policies and vGPU Policies
     */
    defaultVmSizingPolicyId?: pulumi.Input<string>;
    /**
     * When destroying use `delete_force=true` to remove a VDC and any objects it contains, regardless of their state. Default is `false`
     */
    deleteForce?: pulumi.Input<boolean>;
    /**
     * When destroying use `delete_recursive=true` to remove the VDC and any objects it contains that are in a state that normally allows removal. Default is `false`
     */
    deleteRecursive?: pulumi.Input<boolean>;
    /**
     * VDC friendly description
     */
    description?: pulumi.Input<string>;
    /**
     * An ID of NSX-T Edge Cluster which
     * should provide vApp Networking Services or DHCP for isolated networks. Can be looked up using
     * `vcd.getNsxtEdgeCluster` data source. This field is **deprecated** in favor of
     * [`vcd.OrgVdcNsxtNetworkProfile`](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/org_vdc_nsxt_network_profile).
     *
     * @deprecated Please use 'vcd_org_vdc_nsxt_network_profile' resource to manage Edge Cluster and Segment Profile Templates
     */
    edgeClusterId?: pulumi.Input<string>;
    /**
     * Indicates if the Flex VDC should be elastic. Required with the Flex allocation model.
     */
    elasticity?: pulumi.Input<boolean>;
    /**
     * Request fast provisioning. Request will be honored only if the underlying datastore supports it. Fast provisioning can reduce the time it takes to create virtual machines by using vSphere linked clones. If you disable fast provisioning, all provisioning operations will result in full clones.
     */
    enableFastProvisioning?: pulumi.Input<boolean>;
    /**
     * Enables or disables the NSX-V distributed firewall.
     *
     * <a id="storageprofile"></a>
     */
    enableNsxvDistributedFirewall?: pulumi.Input<boolean>;
    /**
     * Boolean to request thin provisioning. Request will be honored only if the underlying data store supports it. Thin provisioning saves storage space by committing it on demand. This allows over-allocation of storage.
     */
    enableThinProvisioning?: pulumi.Input<boolean>;
    /**
     * If true, discovery of vCenter VMs is enabled for resource pools backing this VDC. If false, discovery is disabled. If left unspecified, the actual behaviour depends on enablement at the organization level and at the system level.
     */
    enableVmDiscovery?: pulumi.Input<boolean>;
    /**
     * True if this VDC is enabled for use by the organization VDCs. Default is true.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Indicates if the Flex VDC should include memory overhead into its accounting for admission control. Required with the Flex allocation model. `memoryGuaranteed` must also be specified together with this parameter.
     */
    includeVmMemoryOverhead?: pulumi.Input<boolean>;
    /**
     * Percentage of allocated memory resources guaranteed to vApps deployed in this VDC. For example, if this value is 0.75, then 75% of allocated resources are guaranteed. Required when `allocationModel` is AllocationVApp, AllocationPool or Flex. When Allocation model is AllocationPool minimum value is 0.2. If left empty, VCD sets a value.
     */
    memoryGuaranteed?: pulumi.Input<number>;
    /**
     * Use `metadataEntry` instead. Key value map of metadata to assign to this VDC
     *
     * @deprecated Use metadataEntry instead
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A set of metadata entries to assign. See Metadata section for details.
     */
    metadataEntries?: pulumi.Input<pulumi.Input<inputs.OrgVdcMetadataEntry>[]>;
    /**
     * VDC name
     */
    name?: pulumi.Input<string>;
    /**
     * Reference to a network pool in the Provider VDC. Required if this VDC will contain routed or isolated networks.
     */
    networkPoolName?: pulumi.Input<string>;
    /**
     * Maximum number of network objects that can be deployed in this VDC. Defaults to 0, which means no networks can be deployed.
     */
    networkQuota?: pulumi.Input<number>;
    /**
     * Maximum number of virtual NICs allowed in this VDC. Defaults to 0, which specifies an unlimited number.
     */
    nicQuota?: pulumi.Input<number>;
    /**
     * Organization to create the VDC in, optional if defined at provider level
     */
    org?: pulumi.Input<string>;
    /**
     * Name of the Provider VDC from which this organization VDC is provisioned.
     */
    providerVdcName: pulumi.Input<string>;
    /**
     * Storage profiles supported by this VDC.  See Storage Profile below for details.
     */
    storageProfiles: pulumi.Input<pulumi.Input<inputs.OrgVdcStorageProfile>[]>;
    /**
     * Set of IDs of VM Placement policies that are assigned to this VDC. This field requires `defaultComputePolicyId` to be configured together.
     */
    vmPlacementPolicyIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The maximum number of VMs that can be created in this VDC. Includes deployed and undeployed VMs in vApps and vApp templates. Defaults to 0, which specifies an unlimited number.
     */
    vmQuota?: pulumi.Input<number>;
    /**
     * Set of IDs of VM Sizing policies that are assigned to this VDC. This field requires `defaultComputePolicyId` to be configured together.
     */
    vmSizingPolicyIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Set of IDs of VM vGPU policies that are assigned to this VDC. This field requires `defaultComputePolicyId` to be configured together.
     */
    vmVgpuPolicyIds?: pulumi.Input<pulumi.Input<string>[]>;
}
