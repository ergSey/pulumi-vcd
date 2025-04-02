// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * The Distributed Firewall data source reads a single rule for a particular VDC Group.
 *
 * > There is a different data source
 * [`vcd.NsxtDistributedFirewall`](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_distributed_firewall)
 * resource available that can fetch all firewall rules.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const g1 = vcd.getVdcGroup({
 *     org: "my-org",
 *     name: "my-vdc-group",
 * });
 * const r1 = g1.then(g1 => vcd.getNsxtDistributedFirewallRule({
 *     org: "my-org",
 *     vdcGroupId: g1.id,
 *     name: "rule1",
 * }));
 * ```
 */
export function getNsxtDistributedFirewallRule(args: GetNsxtDistributedFirewallRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetNsxtDistributedFirewallRuleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vcd:index/getNsxtDistributedFirewallRule:getNsxtDistributedFirewallRule", {
        "name": args.name,
        "org": args.org,
        "vdcGroupId": args.vdcGroupId,
    }, opts);
}

/**
 * A collection of arguments for invoking getNsxtDistributedFirewallRule.
 */
export interface GetNsxtDistributedFirewallRuleArgs {
    /**
     * The name of firewall rule
     */
    name: string;
    /**
     * The name of organization in which Distributed Firewall is located. Optional if
     * defined at provider level.
     */
    org?: string;
    /**
     * The ID of a VDC Group
     */
    vdcGroupId: string;
}

/**
 * A collection of values returned by getNsxtDistributedFirewallRule.
 */
export interface GetNsxtDistributedFirewallRuleResult {
    readonly action: string;
    readonly appPortProfileIds: string[];
    readonly comment: string;
    readonly description: string;
    readonly destinationGroupsExcluded: boolean;
    readonly destinationIds: string[];
    readonly direction: string;
    readonly enabled: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ipProtocol: string;
    readonly logging: boolean;
    readonly name: string;
    readonly networkContextProfileIds: string[];
    readonly org?: string;
    readonly sourceGroupsExcluded: boolean;
    readonly sourceIds: string[];
    readonly vdcGroupId: string;
}
/**
 * The Distributed Firewall data source reads a single rule for a particular VDC Group.
 *
 * > There is a different data source
 * [`vcd.NsxtDistributedFirewall`](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_distributed_firewall)
 * resource available that can fetch all firewall rules.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const g1 = vcd.getVdcGroup({
 *     org: "my-org",
 *     name: "my-vdc-group",
 * });
 * const r1 = g1.then(g1 => vcd.getNsxtDistributedFirewallRule({
 *     org: "my-org",
 *     vdcGroupId: g1.id,
 *     name: "rule1",
 * }));
 * ```
 */
export function getNsxtDistributedFirewallRuleOutput(args: GetNsxtDistributedFirewallRuleOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetNsxtDistributedFirewallRuleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vcd:index/getNsxtDistributedFirewallRule:getNsxtDistributedFirewallRule", {
        "name": args.name,
        "org": args.org,
        "vdcGroupId": args.vdcGroupId,
    }, opts);
}

/**
 * A collection of arguments for invoking getNsxtDistributedFirewallRule.
 */
export interface GetNsxtDistributedFirewallRuleOutputArgs {
    /**
     * The name of firewall rule
     */
    name: pulumi.Input<string>;
    /**
     * The name of organization in which Distributed Firewall is located. Optional if
     * defined at provider level.
     */
    org?: pulumi.Input<string>;
    /**
     * The ID of a VDC Group
     */
    vdcGroupId: pulumi.Input<string>;
}
