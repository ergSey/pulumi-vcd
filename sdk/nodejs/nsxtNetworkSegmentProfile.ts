// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class NsxtNetworkSegmentProfile extends pulumi.CustomResource {
    /**
     * Get an existing NsxtNetworkSegmentProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NsxtNetworkSegmentProfileState, opts?: pulumi.CustomResourceOptions): NsxtNetworkSegmentProfile {
        return new NsxtNetworkSegmentProfile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/nsxtNetworkSegmentProfile:NsxtNetworkSegmentProfile';

    /**
     * Returns true if the given object is an instance of NsxtNetworkSegmentProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NsxtNetworkSegmentProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NsxtNetworkSegmentProfile.__pulumiType;
    }

    /**
     * IP Discovery Profile ID. Can be referenced using
     * [`vcd.getNsxtSegmentIpDiscoveryProfile`](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_segment_ip_discovery_profile)
     * data source.
     */
    public readonly ipDiscoveryProfileId!: pulumi.Output<string>;
    /**
     * MAC Discovery Profile ID. Can be referenced using
     * [`vcd.getNsxtSegmentMacDiscoveryProfile`](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_segment_mac_discovery_profile)
     * data source.
     */
    public readonly macDiscoveryProfileId!: pulumi.Output<string>;
    /**
     * The name of organization to use, optional if defined at provider level
     */
    public readonly org!: pulumi.Output<string | undefined>;
    /**
     * Org VDC Network ID
     */
    public readonly orgNetworkId!: pulumi.Output<string>;
    /**
     * QoS Profile ID. Can be referenced using
     * [`vcd.getNsxtSegmentQosProfile`](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_segment_qos_profile)
     * data source.
     */
    public readonly qosProfileId!: pulumi.Output<string>;
    /**
     * Segment Profile Template ID to be applied for this Org
     * VDC Network
     */
    public readonly segmentProfileTemplateId!: pulumi.Output<string | undefined>;
    /**
     * Segment Profile Template Name
     */
    public /*out*/ readonly segmentProfileTemplateName!: pulumi.Output<string>;
    /**
     * Segment Security Profile ID. Can be referenced using
     * [`vcd.getNsxtSegmentSecurityProfile`](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_segment_security_profile)
     * data source.
     */
    public readonly segmentSecurityProfileId!: pulumi.Output<string>;
    /**
     * Spoof Guard Profile ID. Can be referenced using
     * [`vcd.getNsxtSegmentSpoofGuardProfile`](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_segment_spoof_guard_profile)
     * data source.
     */
    public readonly spoofGuardProfileId!: pulumi.Output<string>;

    /**
     * Create a NsxtNetworkSegmentProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NsxtNetworkSegmentProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NsxtNetworkSegmentProfileArgs | NsxtNetworkSegmentProfileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NsxtNetworkSegmentProfileState | undefined;
            resourceInputs["ipDiscoveryProfileId"] = state ? state.ipDiscoveryProfileId : undefined;
            resourceInputs["macDiscoveryProfileId"] = state ? state.macDiscoveryProfileId : undefined;
            resourceInputs["org"] = state ? state.org : undefined;
            resourceInputs["orgNetworkId"] = state ? state.orgNetworkId : undefined;
            resourceInputs["qosProfileId"] = state ? state.qosProfileId : undefined;
            resourceInputs["segmentProfileTemplateId"] = state ? state.segmentProfileTemplateId : undefined;
            resourceInputs["segmentProfileTemplateName"] = state ? state.segmentProfileTemplateName : undefined;
            resourceInputs["segmentSecurityProfileId"] = state ? state.segmentSecurityProfileId : undefined;
            resourceInputs["spoofGuardProfileId"] = state ? state.spoofGuardProfileId : undefined;
        } else {
            const args = argsOrState as NsxtNetworkSegmentProfileArgs | undefined;
            if ((!args || args.orgNetworkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'orgNetworkId'");
            }
            resourceInputs["ipDiscoveryProfileId"] = args ? args.ipDiscoveryProfileId : undefined;
            resourceInputs["macDiscoveryProfileId"] = args ? args.macDiscoveryProfileId : undefined;
            resourceInputs["org"] = args ? args.org : undefined;
            resourceInputs["orgNetworkId"] = args ? args.orgNetworkId : undefined;
            resourceInputs["qosProfileId"] = args ? args.qosProfileId : undefined;
            resourceInputs["segmentProfileTemplateId"] = args ? args.segmentProfileTemplateId : undefined;
            resourceInputs["segmentSecurityProfileId"] = args ? args.segmentSecurityProfileId : undefined;
            resourceInputs["spoofGuardProfileId"] = args ? args.spoofGuardProfileId : undefined;
            resourceInputs["segmentProfileTemplateName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NsxtNetworkSegmentProfile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NsxtNetworkSegmentProfile resources.
 */
export interface NsxtNetworkSegmentProfileState {
    /**
     * IP Discovery Profile ID. Can be referenced using
     * [`vcd.getNsxtSegmentIpDiscoveryProfile`](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_segment_ip_discovery_profile)
     * data source.
     */
    ipDiscoveryProfileId?: pulumi.Input<string>;
    /**
     * MAC Discovery Profile ID. Can be referenced using
     * [`vcd.getNsxtSegmentMacDiscoveryProfile`](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_segment_mac_discovery_profile)
     * data source.
     */
    macDiscoveryProfileId?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level
     */
    org?: pulumi.Input<string>;
    /**
     * Org VDC Network ID
     */
    orgNetworkId?: pulumi.Input<string>;
    /**
     * QoS Profile ID. Can be referenced using
     * [`vcd.getNsxtSegmentQosProfile`](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_segment_qos_profile)
     * data source.
     */
    qosProfileId?: pulumi.Input<string>;
    /**
     * Segment Profile Template ID to be applied for this Org
     * VDC Network
     */
    segmentProfileTemplateId?: pulumi.Input<string>;
    /**
     * Segment Profile Template Name
     */
    segmentProfileTemplateName?: pulumi.Input<string>;
    /**
     * Segment Security Profile ID. Can be referenced using
     * [`vcd.getNsxtSegmentSecurityProfile`](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_segment_security_profile)
     * data source.
     */
    segmentSecurityProfileId?: pulumi.Input<string>;
    /**
     * Spoof Guard Profile ID. Can be referenced using
     * [`vcd.getNsxtSegmentSpoofGuardProfile`](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_segment_spoof_guard_profile)
     * data source.
     */
    spoofGuardProfileId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NsxtNetworkSegmentProfile resource.
 */
export interface NsxtNetworkSegmentProfileArgs {
    /**
     * IP Discovery Profile ID. Can be referenced using
     * [`vcd.getNsxtSegmentIpDiscoveryProfile`](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_segment_ip_discovery_profile)
     * data source.
     */
    ipDiscoveryProfileId?: pulumi.Input<string>;
    /**
     * MAC Discovery Profile ID. Can be referenced using
     * [`vcd.getNsxtSegmentMacDiscoveryProfile`](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_segment_mac_discovery_profile)
     * data source.
     */
    macDiscoveryProfileId?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level
     */
    org?: pulumi.Input<string>;
    /**
     * Org VDC Network ID
     */
    orgNetworkId: pulumi.Input<string>;
    /**
     * QoS Profile ID. Can be referenced using
     * [`vcd.getNsxtSegmentQosProfile`](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_segment_qos_profile)
     * data source.
     */
    qosProfileId?: pulumi.Input<string>;
    /**
     * Segment Profile Template ID to be applied for this Org
     * VDC Network
     */
    segmentProfileTemplateId?: pulumi.Input<string>;
    /**
     * Segment Security Profile ID. Can be referenced using
     * [`vcd.getNsxtSegmentSecurityProfile`](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_segment_security_profile)
     * data source.
     */
    segmentSecurityProfileId?: pulumi.Input<string>;
    /**
     * Spoof Guard Profile ID. Can be referenced using
     * [`vcd.getNsxtSegmentSpoofGuardProfile`](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_segment_spoof_guard_profile)
     * data source.
     */
    spoofGuardProfileId?: pulumi.Input<string>;
}
