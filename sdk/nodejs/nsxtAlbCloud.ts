// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class NsxtAlbCloud extends pulumi.CustomResource {
    /**
     * Get an existing NsxtAlbCloud resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NsxtAlbCloudState, opts?: pulumi.CustomResourceOptions): NsxtAlbCloud {
        return new NsxtAlbCloud(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/nsxtAlbCloud:NsxtAlbCloud';

    /**
     * Returns true if the given object is an instance of NsxtAlbCloud.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NsxtAlbCloud {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NsxtAlbCloud.__pulumiType;
    }

    /**
     * ALB Controller ID
     */
    public readonly controllerId!: pulumi.Output<string>;
    /**
     * An optional description ALB Cloud
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * DetailedHealthMessage contains detailed message on the health of the Cloud
     */
    public /*out*/ readonly healthMessage!: pulumi.Output<string>;
    /**
     * HealthStatus contains status of the Load Balancer Cloud. Possible values are:
     * * UP - The cloud is healthy and ready to enable Load Balancer for an Edge Gateway
     * * DOWN - The cloud is in a failure state. Enabling Load balancer on an Edge Gateway may not be possible
     * * RUNNING - The cloud is currently processing. An example is if it's enabling a Load Balancer for an Edge Gateway
     * * UNAVAILABLE - The cloud is unavailable
     * * UNKNOWN - The cloud state is unknown
     */
    public /*out*/ readonly healthStatus!: pulumi.Output<string>;
    /**
     * Importable Cloud ID. Can be looked up using `vcd.getNsxtAlbImportableCloud` data
     * source
     */
    public readonly importableCloudId!: pulumi.Output<string>;
    /**
     * A name for ALB Cloud
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Network pool ID for ALB Cloud. Can be looked up using `vcd.getNsxtAlbImportableCloud` data
     * source
     */
    public readonly networkPoolId!: pulumi.Output<string>;
    /**
     * Network Pool Name used by the Cloud
     */
    public /*out*/ readonly networkPoolName!: pulumi.Output<string>;

    /**
     * Create a NsxtAlbCloud resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NsxtAlbCloudArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NsxtAlbCloudArgs | NsxtAlbCloudState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NsxtAlbCloudState | undefined;
            resourceInputs["controllerId"] = state ? state.controllerId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["healthMessage"] = state ? state.healthMessage : undefined;
            resourceInputs["healthStatus"] = state ? state.healthStatus : undefined;
            resourceInputs["importableCloudId"] = state ? state.importableCloudId : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkPoolId"] = state ? state.networkPoolId : undefined;
            resourceInputs["networkPoolName"] = state ? state.networkPoolName : undefined;
        } else {
            const args = argsOrState as NsxtAlbCloudArgs | undefined;
            if ((!args || args.controllerId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'controllerId'");
            }
            if ((!args || args.importableCloudId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'importableCloudId'");
            }
            if ((!args || args.networkPoolId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkPoolId'");
            }
            resourceInputs["controllerId"] = args ? args.controllerId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["importableCloudId"] = args ? args.importableCloudId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkPoolId"] = args ? args.networkPoolId : undefined;
            resourceInputs["healthMessage"] = undefined /*out*/;
            resourceInputs["healthStatus"] = undefined /*out*/;
            resourceInputs["networkPoolName"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NsxtAlbCloud.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NsxtAlbCloud resources.
 */
export interface NsxtAlbCloudState {
    /**
     * ALB Controller ID
     */
    controllerId?: pulumi.Input<string>;
    /**
     * An optional description ALB Cloud
     */
    description?: pulumi.Input<string>;
    /**
     * DetailedHealthMessage contains detailed message on the health of the Cloud
     */
    healthMessage?: pulumi.Input<string>;
    /**
     * HealthStatus contains status of the Load Balancer Cloud. Possible values are:
     * * UP - The cloud is healthy and ready to enable Load Balancer for an Edge Gateway
     * * DOWN - The cloud is in a failure state. Enabling Load balancer on an Edge Gateway may not be possible
     * * RUNNING - The cloud is currently processing. An example is if it's enabling a Load Balancer for an Edge Gateway
     * * UNAVAILABLE - The cloud is unavailable
     * * UNKNOWN - The cloud state is unknown
     */
    healthStatus?: pulumi.Input<string>;
    /**
     * Importable Cloud ID. Can be looked up using `vcd.getNsxtAlbImportableCloud` data
     * source
     */
    importableCloudId?: pulumi.Input<string>;
    /**
     * A name for ALB Cloud
     */
    name?: pulumi.Input<string>;
    /**
     * Network pool ID for ALB Cloud. Can be looked up using `vcd.getNsxtAlbImportableCloud` data
     * source
     */
    networkPoolId?: pulumi.Input<string>;
    /**
     * Network Pool Name used by the Cloud
     */
    networkPoolName?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NsxtAlbCloud resource.
 */
export interface NsxtAlbCloudArgs {
    /**
     * ALB Controller ID
     */
    controllerId: pulumi.Input<string>;
    /**
     * An optional description ALB Cloud
     */
    description?: pulumi.Input<string>;
    /**
     * Importable Cloud ID. Can be looked up using `vcd.getNsxtAlbImportableCloud` data
     * source
     */
    importableCloudId: pulumi.Input<string>;
    /**
     * A name for ALB Cloud
     */
    name?: pulumi.Input<string>;
    /**
     * Network pool ID for ALB Cloud. Can be looked up using `vcd.getNsxtAlbImportableCloud` data
     * source
     */
    networkPoolId: pulumi.Input<string>;
}
