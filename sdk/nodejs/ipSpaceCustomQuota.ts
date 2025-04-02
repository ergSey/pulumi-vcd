// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

export class IpSpaceCustomQuota extends pulumi.CustomResource {
    /**
     * Get an existing IpSpaceCustomQuota resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: IpSpaceCustomQuotaState, opts?: pulumi.CustomResourceOptions): IpSpaceCustomQuota {
        return new IpSpaceCustomQuota(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/ipSpaceCustomQuota:IpSpaceCustomQuota';

    /**
     * Returns true if the given object is an instance of IpSpaceCustomQuota.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is IpSpaceCustomQuota {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === IpSpaceCustomQuota.__pulumiType;
    }

    /**
     * IP Prefix Quota set in ipPrefixQuota blocks.
     * Will inherit the default Quota set in `vcd.IpSpace` if not set
     *
     * > The resource `vcd.IpSpaceCustomQuota` can only be created for an Org after an NSX-T Edge
     * Gateway backed by Provider Gateway is created within the Org. An explicit `dependsOn` constraint
     * for an Edge Gateway to exist might be required. (See the example.)
     *
     * <a id="ip-prefix-quota"></a>
     */
    public readonly ipPrefixQuotas!: pulumi.Output<outputs.IpSpaceCustomQuotaIpPrefixQuota[] | undefined>;
    /**
     * Floating IP Quota. Will inherit the default Quota set in
     * `vcd.IpSpace` if not set
     */
    public readonly ipRangeQuota!: pulumi.Output<string | undefined>;
    /**
     * IP Space ID to set Custom Quotas
     */
    public readonly ipSpaceId!: pulumi.Output<string>;
    /**
     * Organization ID, for which the Quota should be customized
     */
    public readonly orgId!: pulumi.Output<string>;

    /**
     * Create a IpSpaceCustomQuota resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: IpSpaceCustomQuotaArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: IpSpaceCustomQuotaArgs | IpSpaceCustomQuotaState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as IpSpaceCustomQuotaState | undefined;
            resourceInputs["ipPrefixQuotas"] = state ? state.ipPrefixQuotas : undefined;
            resourceInputs["ipRangeQuota"] = state ? state.ipRangeQuota : undefined;
            resourceInputs["ipSpaceId"] = state ? state.ipSpaceId : undefined;
            resourceInputs["orgId"] = state ? state.orgId : undefined;
        } else {
            const args = argsOrState as IpSpaceCustomQuotaArgs | undefined;
            if ((!args || args.ipSpaceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ipSpaceId'");
            }
            if ((!args || args.orgId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'orgId'");
            }
            resourceInputs["ipPrefixQuotas"] = args ? args.ipPrefixQuotas : undefined;
            resourceInputs["ipRangeQuota"] = args ? args.ipRangeQuota : undefined;
            resourceInputs["ipSpaceId"] = args ? args.ipSpaceId : undefined;
            resourceInputs["orgId"] = args ? args.orgId : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(IpSpaceCustomQuota.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering IpSpaceCustomQuota resources.
 */
export interface IpSpaceCustomQuotaState {
    /**
     * IP Prefix Quota set in ipPrefixQuota blocks.
     * Will inherit the default Quota set in `vcd.IpSpace` if not set
     *
     * > The resource `vcd.IpSpaceCustomQuota` can only be created for an Org after an NSX-T Edge
     * Gateway backed by Provider Gateway is created within the Org. An explicit `dependsOn` constraint
     * for an Edge Gateway to exist might be required. (See the example.)
     *
     * <a id="ip-prefix-quota"></a>
     */
    ipPrefixQuotas?: pulumi.Input<pulumi.Input<inputs.IpSpaceCustomQuotaIpPrefixQuota>[]>;
    /**
     * Floating IP Quota. Will inherit the default Quota set in
     * `vcd.IpSpace` if not set
     */
    ipRangeQuota?: pulumi.Input<string>;
    /**
     * IP Space ID to set Custom Quotas
     */
    ipSpaceId?: pulumi.Input<string>;
    /**
     * Organization ID, for which the Quota should be customized
     */
    orgId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a IpSpaceCustomQuota resource.
 */
export interface IpSpaceCustomQuotaArgs {
    /**
     * IP Prefix Quota set in ipPrefixQuota blocks.
     * Will inherit the default Quota set in `vcd.IpSpace` if not set
     *
     * > The resource `vcd.IpSpaceCustomQuota` can only be created for an Org after an NSX-T Edge
     * Gateway backed by Provider Gateway is created within the Org. An explicit `dependsOn` constraint
     * for an Edge Gateway to exist might be required. (See the example.)
     *
     * <a id="ip-prefix-quota"></a>
     */
    ipPrefixQuotas?: pulumi.Input<pulumi.Input<inputs.IpSpaceCustomQuotaIpPrefixQuota>[]>;
    /**
     * Floating IP Quota. Will inherit the default Quota set in
     * `vcd.IpSpace` if not set
     */
    ipRangeQuota?: pulumi.Input<string>;
    /**
     * IP Space ID to set Custom Quotas
     */
    ipSpaceId: pulumi.Input<string>;
    /**
     * Organization ID, for which the Quota should be customized
     */
    orgId: pulumi.Input<string>;
}
