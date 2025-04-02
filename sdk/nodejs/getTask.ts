// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a data source for available tasks.
 *
 * Supported in provider *v3.8+*
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * export = async () => {
 *     const some_task = await vcd.getTask({
 *         id: "d4fdcaa9-8db4-45a9-80b8-69de49901bc7",
 *     });
 *     return {
 *         "some-task": some_task,
 *     };
 * }
 * ```
 */
export function getTask(args: GetTaskArgs, opts?: pulumi.InvokeOptions): Promise<GetTaskResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vcd:index/getTask:getTask", {
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getTask.
 */
export interface GetTaskArgs {
    /**
     * The ID of the task
     */
    id: string;
}

/**
 * A collection of values returned by getTask.
 */
export interface GetTaskResult {
    /**
     * Whether user has requested this processing to be canceled (`true` or `false`).
     */
    readonly cancelRequested: boolean;
    /**
     * An optional description of the task.
     */
    readonly description: string;
    /**
     * The date and time that processing of the task was completed. May not be present if the task is still being executed.
     */
    readonly endTime: string;
    /**
     * error information from a failed task.
     */
    readonly error: string;
    /**
     * The date and time at which the task resource will be destroyed and no longer available for retrieval. May not be present if the task has not been executed or is still being executed.
     */
    readonly expiryTime: string;
    /**
     * The URI of the task.
     */
    readonly href: string;
    readonly id: string;
    /**
     * Name of the task. May not be unique. Defines the general operation being performed.
     */
    readonly name: string;
    /**
     * A message describing the operation that is tracked by this task.
     */
    readonly operation: string;
    /**
     * The short name of the operation that is tracked by this task.
     */
    readonly operationName: string;
    /**
     * The unique identifier of the user org.
     */
    readonly orgId: string;
    /**
     * The name of the org to which the user belongs.
     */
    readonly orgName: string;
    /**
     * The unique identifier of the task owner.
     */
    readonly ownerId: string;
    /**
     * The name of the task owner. This is typically the object that the task is creating or updating.
     */
    readonly ownerName: string;
    /**
     * The type of the task owner.
     */
    readonly ownerType: string;
    /**
     * Indicator of task progress as an approximate percentage between 0 and 100. Not available for all tasks.
     */
    readonly progress: number;
    /**
     * The date and time the system started executing the task. May not be present if the task has not been executed yet.
     */
    readonly startTime: string;
    /**
     * The execution status of the task. One of queued, preRunning, running, success, error, aborted.
     */
    readonly status: string;
    /**
     * Type of the task.
     */
    readonly type: string;
    /**
     * The unique identifier of the task user.
     */
    readonly userId: string;
    /**
     * The name of the user who started the task.
     */
    readonly userName: string;
}
/**
 * Provides a data source for available tasks.
 *
 * Supported in provider *v3.8+*
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * export = async () => {
 *     const some_task = await vcd.getTask({
 *         id: "d4fdcaa9-8db4-45a9-80b8-69de49901bc7",
 *     });
 *     return {
 *         "some-task": some_task,
 *     };
 * }
 * ```
 */
export function getTaskOutput(args: GetTaskOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetTaskResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vcd:index/getTask:getTask", {
        "id": args.id,
    }, opts);
}

/**
 * A collection of arguments for invoking getTask.
 */
export interface GetTaskOutputArgs {
    /**
     * The ID of the task
     */
    id: pulumi.Input<string>;
}
