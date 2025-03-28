// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

export class Vapp extends pulumi.CustomResource {
    /**
     * Get an existing Vapp resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: VappState, opts?: pulumi.CustomResourceOptions): Vapp {
        return new Vapp(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/vapp:Vapp';

    /**
     * Returns true if the given object is an instance of Vapp.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Vapp {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Vapp.__pulumiType;
    }

    /**
     * An optional description for the vApp, up to 256 characters.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * Key value map of vApp guest properties
     */
    public readonly guestProperties!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * (Computed) The vApp Hyper Reference.
     */
    public /*out*/ readonly href!: pulumi.Output<string>;
    /**
     * (Computed; *v3.11+*; *VCD 10.5.1+*) A map that contains read-only metadata that is automatically added by VCD (10.5.1+) and provides
     * details on the origin of the vApp (e.g. `vapp.origin.id`, `vapp.origin.name`, `vapp.origin.type`).
     */
    public /*out*/ readonly inheritedMetadata!: pulumi.Output<{[key: string]: string}>;
    /**
     * the information about the vApp lease. It includes the fields below. When this section is 
     * included, both fields are mandatory. If lease values are higher than the ones allowed for the whole Org, the values
     * are **silently** reduced to the highest value allowed.
     */
    public readonly lease!: pulumi.Output<outputs.VappLease>;
    /**
     * Use `metadataEntry` instead. Key value map of metadata to assign to this vApp. Key and value can be any string. (Since *v2.2+* metadata is added directly to vApp instead of first VM in vApp)
     *
     * @deprecated Use metadataEntry instead
     */
    public readonly metadata!: pulumi.Output<{[key: string]: string}>;
    /**
     * A set of metadata entries to assign. See Metadata section for details.
     */
    public readonly metadataEntries!: pulumi.Output<outputs.VappMetadataEntry[]>;
    /**
     * A unique name for the vApp
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
     */
    public readonly org!: pulumi.Output<string | undefined>;
    /**
     * A boolean value stating if this vApp should be powered on. Default is `false`. Works only on update when vApp already has VMs.
     */
    public readonly powerOn!: pulumi.Output<boolean | undefined>;
    /**
     * (Computed; *v2.5+*) The vApp status as a numeric code.
     */
    public /*out*/ readonly status!: pulumi.Output<number>;
    /**
     * (Computed; *v2.5+*) The vApp status as text.
     */
    public /*out*/ readonly statusText!: pulumi.Output<string>;
    /**
     * (*v3.13.0+*) A list of vApp network names included in this vApp
     */
    public /*out*/ readonly vappNetworkNames!: pulumi.Output<string[]>;
    /**
     * (*v3.13.0+*) A list of vApp Org network names included in this vApp
     */
    public /*out*/ readonly vappOrgNetworkNames!: pulumi.Output<string[]>;
    /**
     * The name of VDC to use, optional if defined at provider level
     */
    public readonly vdc!: pulumi.Output<string | undefined>;
    /**
     * (*v3.13.0+*) A list of VM names included in this vApp
     */
    public /*out*/ readonly vmNames!: pulumi.Output<string[]>;

    /**
     * Create a Vapp resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: VappArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: VappArgs | VappState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as VappState | undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["guestProperties"] = state ? state.guestProperties : undefined;
            resourceInputs["href"] = state ? state.href : undefined;
            resourceInputs["inheritedMetadata"] = state ? state.inheritedMetadata : undefined;
            resourceInputs["lease"] = state ? state.lease : undefined;
            resourceInputs["metadata"] = state ? state.metadata : undefined;
            resourceInputs["metadataEntries"] = state ? state.metadataEntries : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["org"] = state ? state.org : undefined;
            resourceInputs["powerOn"] = state ? state.powerOn : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["statusText"] = state ? state.statusText : undefined;
            resourceInputs["vappNetworkNames"] = state ? state.vappNetworkNames : undefined;
            resourceInputs["vappOrgNetworkNames"] = state ? state.vappOrgNetworkNames : undefined;
            resourceInputs["vdc"] = state ? state.vdc : undefined;
            resourceInputs["vmNames"] = state ? state.vmNames : undefined;
        } else {
            const args = argsOrState as VappArgs | undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["guestProperties"] = args ? args.guestProperties : undefined;
            resourceInputs["lease"] = args ? args.lease : undefined;
            resourceInputs["metadata"] = args ? args.metadata : undefined;
            resourceInputs["metadataEntries"] = args ? args.metadataEntries : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["org"] = args ? args.org : undefined;
            resourceInputs["powerOn"] = args ? args.powerOn : undefined;
            resourceInputs["vdc"] = args ? args.vdc : undefined;
            resourceInputs["href"] = undefined /*out*/;
            resourceInputs["inheritedMetadata"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["statusText"] = undefined /*out*/;
            resourceInputs["vappNetworkNames"] = undefined /*out*/;
            resourceInputs["vappOrgNetworkNames"] = undefined /*out*/;
            resourceInputs["vmNames"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Vapp.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering Vapp resources.
 */
export interface VappState {
    /**
     * An optional description for the vApp, up to 256 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Key value map of vApp guest properties
     */
    guestProperties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * (Computed) The vApp Hyper Reference.
     */
    href?: pulumi.Input<string>;
    /**
     * (Computed; *v3.11+*; *VCD 10.5.1+*) A map that contains read-only metadata that is automatically added by VCD (10.5.1+) and provides
     * details on the origin of the vApp (e.g. `vapp.origin.id`, `vapp.origin.name`, `vapp.origin.type`).
     */
    inheritedMetadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * the information about the vApp lease. It includes the fields below. When this section is 
     * included, both fields are mandatory. If lease values are higher than the ones allowed for the whole Org, the values
     * are **silently** reduced to the highest value allowed.
     */
    lease?: pulumi.Input<inputs.VappLease>;
    /**
     * Use `metadataEntry` instead. Key value map of metadata to assign to this vApp. Key and value can be any string. (Since *v2.2+* metadata is added directly to vApp instead of first VM in vApp)
     *
     * @deprecated Use metadataEntry instead
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A set of metadata entries to assign. See Metadata section for details.
     */
    metadataEntries?: pulumi.Input<pulumi.Input<inputs.VappMetadataEntry>[]>;
    /**
     * A unique name for the vApp
     */
    name?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
     */
    org?: pulumi.Input<string>;
    /**
     * A boolean value stating if this vApp should be powered on. Default is `false`. Works only on update when vApp already has VMs.
     */
    powerOn?: pulumi.Input<boolean>;
    /**
     * (Computed; *v2.5+*) The vApp status as a numeric code.
     */
    status?: pulumi.Input<number>;
    /**
     * (Computed; *v2.5+*) The vApp status as text.
     */
    statusText?: pulumi.Input<string>;
    /**
     * (*v3.13.0+*) A list of vApp network names included in this vApp
     */
    vappNetworkNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * (*v3.13.0+*) A list of vApp Org network names included in this vApp
     */
    vappOrgNetworkNames?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The name of VDC to use, optional if defined at provider level
     */
    vdc?: pulumi.Input<string>;
    /**
     * (*v3.13.0+*) A list of VM names included in this vApp
     */
    vmNames?: pulumi.Input<pulumi.Input<string>[]>;
}

/**
 * The set of arguments for constructing a Vapp resource.
 */
export interface VappArgs {
    /**
     * An optional description for the vApp, up to 256 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Key value map of vApp guest properties
     */
    guestProperties?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * the information about the vApp lease. It includes the fields below. When this section is 
     * included, both fields are mandatory. If lease values are higher than the ones allowed for the whole Org, the values
     * are **silently** reduced to the highest value allowed.
     */
    lease?: pulumi.Input<inputs.VappLease>;
    /**
     * Use `metadataEntry` instead. Key value map of metadata to assign to this vApp. Key and value can be any string. (Since *v2.2+* metadata is added directly to vApp instead of first VM in vApp)
     *
     * @deprecated Use metadataEntry instead
     */
    metadata?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A set of metadata entries to assign. See Metadata section for details.
     */
    metadataEntries?: pulumi.Input<pulumi.Input<inputs.VappMetadataEntry>[]>;
    /**
     * A unique name for the vApp
     */
    name?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
     */
    org?: pulumi.Input<string>;
    /**
     * A boolean value stating if this vApp should be powered on. Default is `false`. Works only on update when vApp already has VMs.
     */
    powerOn?: pulumi.Input<boolean>;
    /**
     * The name of VDC to use, optional if defined at provider level
     */
    vdc?: pulumi.Input<string>;
}
