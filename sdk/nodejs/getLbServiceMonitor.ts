// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Provides a VMware Cloud Director Edge Gateway Load Balancer Service Monitor data source. A service monitor
 * defines health check parameters for a particular type of network traffic. It can be associated with
 * a pool. Pool members are monitored according to the service monitor parameters. See example usage of
 * this data source in [server pool resource page](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/lb_server_pool).
 *
 * > **Note:** See additional support notes in [service monitor resource page](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/lb_service_monitor).
 *
 * Supported in provider *v2.4+*
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const my_monitor = vcd.getLbServiceMonitor({
 *     org: "my-org",
 *     vdc: "my-org-vdc",
 *     edgeGateway: "my-edge-gw",
 *     name: "not-managed",
 * });
 * ```
 */
export function getLbServiceMonitor(args: GetLbServiceMonitorArgs, opts?: pulumi.InvokeOptions): Promise<GetLbServiceMonitorResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vcd:index/getLbServiceMonitor:getLbServiceMonitor", {
        "edgeGateway": args.edgeGateway,
        "name": args.name,
        "org": args.org,
        "vdc": args.vdc,
    }, opts);
}

/**
 * A collection of arguments for invoking getLbServiceMonitor.
 */
export interface GetLbServiceMonitorArgs {
    /**
     * The name of the edge gateway on which the service monitor is defined
     */
    edgeGateway: string;
    /**
     * Service Monitor name for identifying the exact service monitor
     */
    name: string;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
     */
    org?: string;
    /**
     * The name of VDC to use, optional if defined at provider level
     */
    vdc?: string;
}

/**
 * A collection of values returned by getLbServiceMonitor.
 */
export interface GetLbServiceMonitorResult {
    readonly edgeGateway: string;
    readonly expected: string;
    readonly extension: {[key: string]: string};
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly interval: number;
    readonly maxRetries: number;
    readonly method: string;
    readonly name: string;
    readonly org?: string;
    readonly receive: string;
    readonly send: string;
    readonly timeout: number;
    readonly type: string;
    readonly url: string;
    readonly vdc?: string;
}
/**
 * Provides a VMware Cloud Director Edge Gateway Load Balancer Service Monitor data source. A service monitor
 * defines health check parameters for a particular type of network traffic. It can be associated with
 * a pool. Pool members are monitored according to the service monitor parameters. See example usage of
 * this data source in [server pool resource page](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/lb_server_pool).
 *
 * > **Note:** See additional support notes in [service monitor resource page](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/lb_service_monitor).
 *
 * Supported in provider *v2.4+*
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const my_monitor = vcd.getLbServiceMonitor({
 *     org: "my-org",
 *     vdc: "my-org-vdc",
 *     edgeGateway: "my-edge-gw",
 *     name: "not-managed",
 * });
 * ```
 */
export function getLbServiceMonitorOutput(args: GetLbServiceMonitorOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetLbServiceMonitorResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vcd:index/getLbServiceMonitor:getLbServiceMonitor", {
        "edgeGateway": args.edgeGateway,
        "name": args.name,
        "org": args.org,
        "vdc": args.vdc,
    }, opts);
}

/**
 * A collection of arguments for invoking getLbServiceMonitor.
 */
export interface GetLbServiceMonitorOutputArgs {
    /**
     * The name of the edge gateway on which the service monitor is defined
     */
    edgeGateway: pulumi.Input<string>;
    /**
     * Service Monitor name for identifying the exact service monitor
     */
    name: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
     */
    org?: pulumi.Input<string>;
    /**
     * The name of VDC to use, optional if defined at provider level
     */
    vdc?: pulumi.Input<string>;
}
