// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class RdeInterfaceBehavior extends pulumi.CustomResource {
    /**
     * Get an existing RdeInterfaceBehavior resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: RdeInterfaceBehaviorState, opts?: pulumi.CustomResourceOptions): RdeInterfaceBehavior {
        return new RdeInterfaceBehavior(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/rdeInterfaceBehavior:RdeInterfaceBehavior';

    /**
     * Returns true if the given object is an instance of RdeInterfaceBehavior.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RdeInterfaceBehavior {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RdeInterfaceBehavior.__pulumiType;
    }

    /**
     * Useful to update execution properties marked with `_secure_` and `_internal_`
     * as these are not retrievable from VCD, so they are not saved in state. Setting this to `true` will make the provider
     * to ask for updates whenever there is a secure property in the execution of the Behavior
     */
    public readonly alwaysUpdateSecureExecutionProperties!: pulumi.Output<boolean | undefined>;
    /**
     * A description specifying the contract of the Behavior
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A map that specifies the Behavior execution mechanism, this is just a simplification of `executionJson` that
     * can make the configuration more readable for simpler Behaviors. One of `executionJson` or `execution` must be set.
     */
    public readonly execution!: pulumi.Output<{[key: string]: string}>;
    /**
     * A string representing a valid JSON that specifies the Behavior execution mechanism.
     * You can find more information about the different execution types, like `WebHook`, `noop`, `Activity`, `MQTT`, `VRO`, `AWSLambdaFaaS`
     * and others [in the Extensibility SDK documentation](https://vmware.github.io/vcd-ext-sdk/docs/defined_entities_api/behaviors).
     * One of `executionJson` or `execution` must be set.
     */
    public readonly executionJson!: pulumi.Output<string>;
    /**
     * Name of the Behavior
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the RDE Interface that owns the Behavior
     */
    public readonly rdeInterfaceId!: pulumi.Output<string>;
    /**
     * The Behavior invocation reference to be used for polymorphic behavior invocations
     */
    public /*out*/ readonly ref!: pulumi.Output<string>;

    /**
     * Create a RdeInterfaceBehavior resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RdeInterfaceBehaviorArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: RdeInterfaceBehaviorArgs | RdeInterfaceBehaviorState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as RdeInterfaceBehaviorState | undefined;
            resourceInputs["alwaysUpdateSecureExecutionProperties"] = state ? state.alwaysUpdateSecureExecutionProperties : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["execution"] = state ? state.execution : undefined;
            resourceInputs["executionJson"] = state ? state.executionJson : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["rdeInterfaceId"] = state ? state.rdeInterfaceId : undefined;
            resourceInputs["ref"] = state ? state.ref : undefined;
        } else {
            const args = argsOrState as RdeInterfaceBehaviorArgs | undefined;
            if ((!args || args.rdeInterfaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'rdeInterfaceId'");
            }
            resourceInputs["alwaysUpdateSecureExecutionProperties"] = args ? args.alwaysUpdateSecureExecutionProperties : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["execution"] = args ? args.execution : undefined;
            resourceInputs["executionJson"] = args ? args.executionJson : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["rdeInterfaceId"] = args ? args.rdeInterfaceId : undefined;
            resourceInputs["ref"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(RdeInterfaceBehavior.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering RdeInterfaceBehavior resources.
 */
export interface RdeInterfaceBehaviorState {
    /**
     * Useful to update execution properties marked with `_secure_` and `_internal_`
     * as these are not retrievable from VCD, so they are not saved in state. Setting this to `true` will make the provider
     * to ask for updates whenever there is a secure property in the execution of the Behavior
     */
    alwaysUpdateSecureExecutionProperties?: pulumi.Input<boolean>;
    /**
     * A description specifying the contract of the Behavior
     */
    description?: pulumi.Input<string>;
    /**
     * A map that specifies the Behavior execution mechanism, this is just a simplification of `executionJson` that
     * can make the configuration more readable for simpler Behaviors. One of `executionJson` or `execution` must be set.
     */
    execution?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A string representing a valid JSON that specifies the Behavior execution mechanism.
     * You can find more information about the different execution types, like `WebHook`, `noop`, `Activity`, `MQTT`, `VRO`, `AWSLambdaFaaS`
     * and others [in the Extensibility SDK documentation](https://vmware.github.io/vcd-ext-sdk/docs/defined_entities_api/behaviors).
     * One of `executionJson` or `execution` must be set.
     */
    executionJson?: pulumi.Input<string>;
    /**
     * Name of the Behavior
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the RDE Interface that owns the Behavior
     */
    rdeInterfaceId?: pulumi.Input<string>;
    /**
     * The Behavior invocation reference to be used for polymorphic behavior invocations
     */
    ref?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a RdeInterfaceBehavior resource.
 */
export interface RdeInterfaceBehaviorArgs {
    /**
     * Useful to update execution properties marked with `_secure_` and `_internal_`
     * as these are not retrievable from VCD, so they are not saved in state. Setting this to `true` will make the provider
     * to ask for updates whenever there is a secure property in the execution of the Behavior
     */
    alwaysUpdateSecureExecutionProperties?: pulumi.Input<boolean>;
    /**
     * A description specifying the contract of the Behavior
     */
    description?: pulumi.Input<string>;
    /**
     * A map that specifies the Behavior execution mechanism, this is just a simplification of `executionJson` that
     * can make the configuration more readable for simpler Behaviors. One of `executionJson` or `execution` must be set.
     */
    execution?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A string representing a valid JSON that specifies the Behavior execution mechanism.
     * You can find more information about the different execution types, like `WebHook`, `noop`, `Activity`, `MQTT`, `VRO`, `AWSLambdaFaaS`
     * and others [in the Extensibility SDK documentation](https://vmware.github.io/vcd-ext-sdk/docs/defined_entities_api/behaviors).
     * One of `executionJson` or `execution` must be set.
     */
    executionJson?: pulumi.Input<string>;
    /**
     * Name of the Behavior
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the RDE Interface that owns the Behavior
     */
    rdeInterfaceId: pulumi.Input<string>;
}
