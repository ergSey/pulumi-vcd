// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

/**
 * Provides a VMware Cloud Director firewall rule data source for advanced edge gateways (NSX-V). This can be
 * used to read existing rules by ID and use its attributes in other resources.
 *
 * > **Note:** This data source requires advanced edge gateway.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const my_rule = vcd.getNsxvFirewallRule({
 *     org: "my-org",
 *     vdc: "my-org-vdc",
 *     edgeGateway: "my-edge-gw",
 *     ruleId: "133048",
 * });
 * ```
 */
export function getNsxvFirewallRule(args: GetNsxvFirewallRuleArgs, opts?: pulumi.InvokeOptions): Promise<GetNsxvFirewallRuleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vcd:index/getNsxvFirewallRule:getNsxvFirewallRule", {
        "edgeGateway": args.edgeGateway,
        "org": args.org,
        "ruleId": args.ruleId,
        "vdc": args.vdc,
    }, opts);
}

/**
 * A collection of arguments for invoking getNsxvFirewallRule.
 */
export interface GetNsxvFirewallRuleArgs {
    /**
     * The name of the edge gateway on which to apply the DNAT rule.
     */
    edgeGateway: string;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.
     */
    org?: string;
    /**
     * ID of firewall rule (not UI number). See more information about firewall
     * rule ID in `vcd.NsxvFirewallRule` [import section](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxv_firewall_rule#listing-real-firewall-rule-ids).
     */
    ruleId: string;
    /**
     * The name of VDC to use, optional if defined at provider level.
     */
    vdc?: string;
}

/**
 * A collection of values returned by getNsxvFirewallRule.
 */
export interface GetNsxvFirewallRuleResult {
    readonly action: string;
    readonly destinations: outputs.GetNsxvFirewallRuleDestination[];
    readonly edgeGateway: string;
    readonly enabled: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly loggingEnabled: boolean;
    readonly name: string;
    readonly org?: string;
    readonly ruleId: string;
    readonly ruleTag: number;
    readonly ruleType: string;
    readonly services: outputs.GetNsxvFirewallRuleService[];
    readonly sources: outputs.GetNsxvFirewallRuleSource[];
    readonly vdc?: string;
}
/**
 * Provides a VMware Cloud Director firewall rule data source for advanced edge gateways (NSX-V). This can be
 * used to read existing rules by ID and use its attributes in other resources.
 *
 * > **Note:** This data source requires advanced edge gateway.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const my_rule = vcd.getNsxvFirewallRule({
 *     org: "my-org",
 *     vdc: "my-org-vdc",
 *     edgeGateway: "my-edge-gw",
 *     ruleId: "133048",
 * });
 * ```
 */
export function getNsxvFirewallRuleOutput(args: GetNsxvFirewallRuleOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetNsxvFirewallRuleResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vcd:index/getNsxvFirewallRule:getNsxvFirewallRule", {
        "edgeGateway": args.edgeGateway,
        "org": args.org,
        "ruleId": args.ruleId,
        "vdc": args.vdc,
    }, opts);
}

/**
 * A collection of arguments for invoking getNsxvFirewallRule.
 */
export interface GetNsxvFirewallRuleOutputArgs {
    /**
     * The name of the edge gateway on which to apply the DNAT rule.
     */
    edgeGateway: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.
     */
    org?: pulumi.Input<string>;
    /**
     * ID of firewall rule (not UI number). See more information about firewall
     * rule ID in `vcd.NsxvFirewallRule` [import section](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxv_firewall_rule#listing-real-firewall-rule-ids).
     */
    ruleId: pulumi.Input<string>;
    /**
     * The name of VDC to use, optional if defined at provider level.
     */
    vdc?: pulumi.Input<string>;
}
