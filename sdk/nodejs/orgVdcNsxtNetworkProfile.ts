// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class OrgVdcNsxtNetworkProfile extends pulumi.CustomResource {
    /**
     * Get an existing OrgVdcNsxtNetworkProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: OrgVdcNsxtNetworkProfileState, opts?: pulumi.CustomResourceOptions): OrgVdcNsxtNetworkProfile {
        return new OrgVdcNsxtNetworkProfile(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/orgVdcNsxtNetworkProfile:OrgVdcNsxtNetworkProfile';

    /**
     * Returns true if the given object is an instance of OrgVdcNsxtNetworkProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrgVdcNsxtNetworkProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrgVdcNsxtNetworkProfile.__pulumiType;
    }

    /**
     * Edge Cluster ID to be used for this VDC
     */
    public readonly edgeClusterId!: pulumi.Output<string | undefined>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    public readonly org!: pulumi.Output<string | undefined>;
    /**
     * Default Segment Profile
     * Template ID for all vApp Networks in a VDC
     */
    public readonly vappNetworksDefaultSegmentProfileTemplateId!: pulumi.Output<string | undefined>;
    /**
     * The name of VDC to use, optional if defined at provider level
     */
    public readonly vdc!: pulumi.Output<string | undefined>;
    /**
     * Default Segment Profile
     * Template ID for all VDC Networks in a VDC
     */
    public readonly vdcNetworksDefaultSegmentProfileTemplateId!: pulumi.Output<string | undefined>;

    /**
     * Create a OrgVdcNsxtNetworkProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: OrgVdcNsxtNetworkProfileArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: OrgVdcNsxtNetworkProfileArgs | OrgVdcNsxtNetworkProfileState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as OrgVdcNsxtNetworkProfileState | undefined;
            resourceInputs["edgeClusterId"] = state ? state.edgeClusterId : undefined;
            resourceInputs["org"] = state ? state.org : undefined;
            resourceInputs["vappNetworksDefaultSegmentProfileTemplateId"] = state ? state.vappNetworksDefaultSegmentProfileTemplateId : undefined;
            resourceInputs["vdc"] = state ? state.vdc : undefined;
            resourceInputs["vdcNetworksDefaultSegmentProfileTemplateId"] = state ? state.vdcNetworksDefaultSegmentProfileTemplateId : undefined;
        } else {
            const args = argsOrState as OrgVdcNsxtNetworkProfileArgs | undefined;
            resourceInputs["edgeClusterId"] = args ? args.edgeClusterId : undefined;
            resourceInputs["org"] = args ? args.org : undefined;
            resourceInputs["vappNetworksDefaultSegmentProfileTemplateId"] = args ? args.vappNetworksDefaultSegmentProfileTemplateId : undefined;
            resourceInputs["vdc"] = args ? args.vdc : undefined;
            resourceInputs["vdcNetworksDefaultSegmentProfileTemplateId"] = args ? args.vdcNetworksDefaultSegmentProfileTemplateId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(OrgVdcNsxtNetworkProfile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering OrgVdcNsxtNetworkProfile resources.
 */
export interface OrgVdcNsxtNetworkProfileState {
    /**
     * Edge Cluster ID to be used for this VDC
     */
    edgeClusterId?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    org?: pulumi.Input<string>;
    /**
     * Default Segment Profile
     * Template ID for all vApp Networks in a VDC
     */
    vappNetworksDefaultSegmentProfileTemplateId?: pulumi.Input<string>;
    /**
     * The name of VDC to use, optional if defined at provider level
     */
    vdc?: pulumi.Input<string>;
    /**
     * Default Segment Profile
     * Template ID for all VDC Networks in a VDC
     */
    vdcNetworksDefaultSegmentProfileTemplateId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a OrgVdcNsxtNetworkProfile resource.
 */
export interface OrgVdcNsxtNetworkProfileArgs {
    /**
     * Edge Cluster ID to be used for this VDC
     */
    edgeClusterId?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
     * different organizations
     */
    org?: pulumi.Input<string>;
    /**
     * Default Segment Profile
     * Template ID for all vApp Networks in a VDC
     */
    vappNetworksDefaultSegmentProfileTemplateId?: pulumi.Input<string>;
    /**
     * The name of VDC to use, optional if defined at provider level
     */
    vdc?: pulumi.Input<string>;
    /**
     * Default Segment Profile
     * Template ID for all VDC Networks in a VDC
     */
    vdcNetworksDefaultSegmentProfileTemplateId?: pulumi.Input<string>;
}
