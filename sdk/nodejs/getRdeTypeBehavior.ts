// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides the capability of reading RDE Type Behaviors in VMware Cloud Director, which override an existing [RDE Interface
 * Behavior](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/rde_interface_behavior).
 *
 * Supported in provider *v3.10+*. Requires System Administrator privileges.
 */
export function getRdeTypeBehavior(args: GetRdeTypeBehaviorArgs, opts?: pulumi.InvokeOptions): Promise<GetRdeTypeBehaviorResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vcd:index/getRdeTypeBehavior:getRdeTypeBehavior", {
        "behaviorId": args.behaviorId,
        "rdeTypeId": args.rdeTypeId,
    }, opts);
}

/**
 * A collection of arguments for invoking getRdeTypeBehavior.
 */
export interface GetRdeTypeBehaviorArgs {
    behaviorId: string;
    /**
     * The ID of the [RDE Type](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/rde_type) that owns the Behavior override
     */
    rdeTypeId: string;
}

/**
 * A collection of values returned by getRdeTypeBehavior.
 */
export interface GetRdeTypeBehaviorResult {
    readonly behaviorId: string;
    readonly description: string;
    /**
     * @deprecated This argument cannot consider complex execution structures. For that purpose, use 'execution_json' instead
     */
    readonly execution: {[key: string]: string};
    readonly executionJson: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly name: string;
    readonly rdeTypeId: string;
    readonly ref: string;
}
/**
 * Provides the capability of reading RDE Type Behaviors in VMware Cloud Director, which override an existing [RDE Interface
 * Behavior](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/rde_interface_behavior).
 *
 * Supported in provider *v3.10+*. Requires System Administrator privileges.
 */
export function getRdeTypeBehaviorOutput(args: GetRdeTypeBehaviorOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetRdeTypeBehaviorResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vcd:index/getRdeTypeBehavior:getRdeTypeBehavior", {
        "behaviorId": args.behaviorId,
        "rdeTypeId": args.rdeTypeId,
    }, opts);
}

/**
 * A collection of arguments for invoking getRdeTypeBehavior.
 */
export interface GetRdeTypeBehaviorOutputArgs {
    behaviorId: pulumi.Input<string>;
    /**
     * The ID of the [RDE Type](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/rde_type) that owns the Behavior override
     */
    rdeTypeId: pulumi.Input<string>;
}
