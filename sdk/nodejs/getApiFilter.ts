// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * Supported in provider *v3.14+* and VCD 10.4.3+.
 *
 * Provides a data source to read API Filters in VMware Cloud Director. An API Filter allows to extend VCD API with customised URLs
 * that can be redirected to an [`vcd.ExternalEndpoint`](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/external_endpoint).
 *
 * > Only `System Administrator` can use this data source.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const apiFilter1 = vcd.getApiFilter({
 *     apiFilterId: "urn:vcloud:apiFilter:4252ab09-eed8-4bc6-86d7-6019090273f5",
 * });
 * ```
 */
export function getApiFilter(args: GetApiFilterArgs, opts?: pulumi.InvokeOptions): Promise<GetApiFilterResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vcd:index/getApiFilter:getApiFilter", {
        "apiFilterId": args.apiFilterId,
    }, opts);
}

/**
 * A collection of arguments for invoking getApiFilter.
 */
export interface GetApiFilterArgs {
    /**
     * ID of the API Filter. This is the only way of unequivocally identify an API Filter. A list of
     * available API Filters can be obtained by using the `list@` option of the import mechanism of the [resource](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/api_filter#importing)
     */
    apiFilterId: string;
}

/**
 * A collection of values returned by getApiFilter.
 */
export interface GetApiFilterResult {
    readonly apiFilterId: string;
    readonly externalEndpointId: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly urlMatcherPattern: string;
    readonly urlMatcherScope: string;
}
/**
 * Supported in provider *v3.14+* and VCD 10.4.3+.
 *
 * Provides a data source to read API Filters in VMware Cloud Director. An API Filter allows to extend VCD API with customised URLs
 * that can be redirected to an [`vcd.ExternalEndpoint`](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/external_endpoint).
 *
 * > Only `System Administrator` can use this data source.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const apiFilter1 = vcd.getApiFilter({
 *     apiFilterId: "urn:vcloud:apiFilter:4252ab09-eed8-4bc6-86d7-6019090273f5",
 * });
 * ```
 */
export function getApiFilterOutput(args: GetApiFilterOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetApiFilterResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vcd:index/getApiFilter:getApiFilter", {
        "apiFilterId": args.apiFilterId,
    }, opts);
}

/**
 * A collection of arguments for invoking getApiFilter.
 */
export interface GetApiFilterOutputArgs {
    /**
     * ID of the API Filter. This is the only way of unequivocally identify an API Filter. A list of
     * available API Filters can be obtained by using the `list@` option of the import mechanism of the [resource](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/api_filter#importing)
     */
    apiFilterId: pulumi.Input<string>;
}
