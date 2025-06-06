// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

export class NsxtDynamicSecurityGroup extends pulumi.CustomResource {
    /**
     * Get an existing NsxtDynamicSecurityGroup resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NsxtDynamicSecurityGroupState, opts?: pulumi.CustomResourceOptions): NsxtDynamicSecurityGroup {
        return new NsxtDynamicSecurityGroup(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/nsxtDynamicSecurityGroup:NsxtDynamicSecurityGroup';

    /**
     * Returns true if the given object is an instance of NsxtDynamicSecurityGroup.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NsxtDynamicSecurityGroup {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NsxtDynamicSecurityGroup.__pulumiType;
    }

    /**
     * Up to 3 criteria for matching VMs. List of criteria is matched with boolean
     * `OR` operation and matching any of defined criteria will include objects. Each `criteria` can
     * contains up to 4 `rule` definitions.
     */
    public readonly criterias!: pulumi.Output<outputs.NsxtDynamicSecurityGroupCriteria[] | undefined>;
    /**
     * An optional description of the Dynamic Security Group
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * A set of member VMs (if exist). see Member VMs below for details.
     */
    public /*out*/ readonly memberVms!: pulumi.Output<outputs.NsxtDynamicSecurityGroupMemberVm[]>;
    /**
     * A unique name for Dynamic Security Group
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful
     * when connected as sysadmin working across different organisations.
     */
    public readonly org!: pulumi.Output<string | undefined>;
    /**
     * VDC Group ID for Dynamic Security Group creation.
     */
    public readonly vdcGroupId!: pulumi.Output<string>;

    /**
     * Create a NsxtDynamicSecurityGroup resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NsxtDynamicSecurityGroupArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NsxtDynamicSecurityGroupArgs | NsxtDynamicSecurityGroupState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NsxtDynamicSecurityGroupState | undefined;
            resourceInputs["criterias"] = state ? state.criterias : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["memberVms"] = state ? state.memberVms : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["org"] = state ? state.org : undefined;
            resourceInputs["vdcGroupId"] = state ? state.vdcGroupId : undefined;
        } else {
            const args = argsOrState as NsxtDynamicSecurityGroupArgs | undefined;
            if ((!args || args.vdcGroupId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vdcGroupId'");
            }
            resourceInputs["criterias"] = args ? args.criterias : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["org"] = args ? args.org : undefined;
            resourceInputs["vdcGroupId"] = args ? args.vdcGroupId : undefined;
            resourceInputs["memberVms"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(NsxtDynamicSecurityGroup.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NsxtDynamicSecurityGroup resources.
 */
export interface NsxtDynamicSecurityGroupState {
    /**
     * Up to 3 criteria for matching VMs. List of criteria is matched with boolean
     * `OR` operation and matching any of defined criteria will include objects. Each `criteria` can
     * contains up to 4 `rule` definitions.
     */
    criterias?: pulumi.Input<pulumi.Input<inputs.NsxtDynamicSecurityGroupCriteria>[]>;
    /**
     * An optional description of the Dynamic Security Group
     */
    description?: pulumi.Input<string>;
    /**
     * A set of member VMs (if exist). see Member VMs below for details.
     */
    memberVms?: pulumi.Input<pulumi.Input<inputs.NsxtDynamicSecurityGroupMemberVm>[]>;
    /**
     * A unique name for Dynamic Security Group
     */
    name?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful
     * when connected as sysadmin working across different organisations.
     */
    org?: pulumi.Input<string>;
    /**
     * VDC Group ID for Dynamic Security Group creation.
     */
    vdcGroupId?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NsxtDynamicSecurityGroup resource.
 */
export interface NsxtDynamicSecurityGroupArgs {
    /**
     * Up to 3 criteria for matching VMs. List of criteria is matched with boolean
     * `OR` operation and matching any of defined criteria will include objects. Each `criteria` can
     * contains up to 4 `rule` definitions.
     */
    criterias?: pulumi.Input<pulumi.Input<inputs.NsxtDynamicSecurityGroupCriteria>[]>;
    /**
     * An optional description of the Dynamic Security Group
     */
    description?: pulumi.Input<string>;
    /**
     * A unique name for Dynamic Security Group
     */
    name?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful
     * when connected as sysadmin working across different organisations.
     */
    org?: pulumi.Input<string>;
    /**
     * VDC Group ID for Dynamic Security Group creation.
     */
    vdcGroupId: pulumi.Input<string>;
}
