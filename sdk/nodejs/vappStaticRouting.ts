// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

export class VappStaticRouting extends pulumi.CustomResource {
    /**
     * Get an existing VappStaticRouting resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VappStaticRoutingState, opts?: pulumi.CustomResourceOptions): VappStaticRouting {
        return new VappStaticRouting(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/vappStaticRouting:VappStaticRouting';

    /**
     * Returns true if the given object is an instance of VappStaticRouting.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VappStaticRouting {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VappStaticRouting.__pulumiType;
    }

    /**
     * Enable or disable static Routing. Default is `true`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * The identifier of [vApp network](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp_network).
     */
    public readonly networkId!: pulumi.Output<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.
     */
    public readonly org!: pulumi.Output<string | undefined>;
    /**
     * Configures a static routing rule; see Rules below for details.
     *
     * <a id="rules"></a>
     */
    public readonly rules!: pulumi.Output<outputs.VappStaticRoutingRule[] | undefined>;
    /**
     * The identifier of [vApp](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp).
     */
    public readonly vappId!: pulumi.Output<string>;
    /**
     * The name of VDC to use, optional if defined at provider level.
     */
    public readonly vdc!: pulumi.Output<string | undefined>;

    /**
     * Create a VappStaticRouting resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VappStaticRoutingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VappStaticRoutingArgs | VappStaticRoutingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VappStaticRoutingState | undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["networkId"] = state ? state.networkId : undefined;
            resourceInputs["org"] = state ? state.org : undefined;
            resourceInputs["rules"] = state ? state.rules : undefined;
            resourceInputs["vappId"] = state ? state.vappId : undefined;
            resourceInputs["vdc"] = state ? state.vdc : undefined;
        } else {
            const args = argsOrState as VappStaticRoutingArgs | undefined;
            if ((!args || args.networkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkId'");
            }
            if ((!args || args.vappId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vappId'");
            }
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["networkId"] = args ? args.networkId : undefined;
            resourceInputs["org"] = args ? args.org : undefined;
            resourceInputs["rules"] = args ? args.rules : undefined;
            resourceInputs["vappId"] = args ? args.vappId : undefined;
            resourceInputs["vdc"] = args ? args.vdc : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VappStaticRouting.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering VappStaticRouting resources.
 */
export interface VappStaticRoutingState {
    /**
     * Enable or disable static Routing. Default is `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The identifier of [vApp network](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp_network).
     */
    networkId?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.
     */
    org?: pulumi.Input<string>;
    /**
     * Configures a static routing rule; see Rules below for details.
     *
     * <a id="rules"></a>
     */
    rules?: pulumi.Input<pulumi.Input<inputs.VappStaticRoutingRule>[]>;
    /**
     * The identifier of [vApp](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp).
     */
    vappId?: pulumi.Input<string>;
    /**
     * The name of VDC to use, optional if defined at provider level.
     */
    vdc?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a VappStaticRouting resource.
 */
export interface VappStaticRoutingArgs {
    /**
     * Enable or disable static Routing. Default is `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The identifier of [vApp network](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp_network).
     */
    networkId: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.
     */
    org?: pulumi.Input<string>;
    /**
     * Configures a static routing rule; see Rules below for details.
     *
     * <a id="rules"></a>
     */
    rules?: pulumi.Input<pulumi.Input<inputs.VappStaticRoutingRule>[]>;
    /**
     * The identifier of [vApp](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp).
     */
    vappId: pulumi.Input<string>;
    /**
     * The name of VDC to use, optional if defined at provider level.
     */
    vdc?: pulumi.Input<string>;
}
