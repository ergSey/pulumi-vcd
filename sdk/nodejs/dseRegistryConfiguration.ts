// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

export class DseRegistryConfiguration extends pulumi.CustomResource {
    /**
     * Get an existing DseRegistryConfiguration resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DseRegistryConfigurationState, opts?: pulumi.CustomResourceOptions): DseRegistryConfiguration {
        return new DseRegistryConfiguration(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/dseRegistryConfiguration:DseRegistryConfiguration';

    /**
     * Returns true if the given object is an instance of DseRegistryConfiguration.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DseRegistryConfiguration {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DseRegistryConfiguration.__pulumiType;
    }

    /**
     * Chart repository for Helm based images
     */
    public readonly chartRepository!: pulumi.Output<string>;
    /**
     * A set of version constrains that this Data Solution defines
     */
    public /*out*/ readonly compatibleVersionConstraints!: pulumi.Output<string[]>;
    /**
     * Only applies to `VCD Data Solutions` configuration. Specifies
     * credentials that can be used to authenticate to repositories. See Container Registry
     * Configuration 
     *
     *
     * <a id="container-registry"></a>
     */
    public readonly containerRegistries!: pulumi.Output<outputs.DseRegistryConfigurationContainerRegistry[] | undefined>;
    /**
     * Default chart repository as provided by Data Solution
     */
    public /*out*/ readonly defaultChartRepository!: pulumi.Output<string>;
    /**
     * Default package name as provided by Data Solution
     */
    public /*out*/ readonly defaultPackageName!: pulumi.Output<string>;
    /**
     * Default container repository as provided by Data Solution
     */
    public /*out*/ readonly defaultRepository!: pulumi.Output<string>;
    /**
     * Default package version as provided by Data Solution
     */
    public /*out*/ readonly defaultVersion!: pulumi.Output<string>;
    /**
     * The name of Data Solution as it appears in repository configuration
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Helm package name. Only for Helm based images
     */
    public readonly packageName!: pulumi.Output<string>;
    /**
     * Package repository for container based images
     */
    public readonly packageRepository!: pulumi.Output<string>;
    /**
     * State of parent Runtime Defined Entity (RDE)
     */
    public /*out*/ readonly rdeState!: pulumi.Output<string>;
    /**
     * Boolean flag as defined in Data Solution
     */
    public /*out*/ readonly requiresVersionCompatibility!: pulumi.Output<boolean>;
    /**
     * Type of repository settings. It can be one of `PackageRepository`, `ChartRepository`
     */
    public /*out*/ readonly type!: pulumi.Output<string>;
    /**
     * Defines if repository settings should be inherited from Data
     * Solution itself. Default `false`
     */
    public readonly useDefaultValues!: pulumi.Output<boolean | undefined>;
    /**
     * Version of package to use. Required when `useDefaultValues` is not used.
     */
    public readonly version!: pulumi.Output<string>;

    /**
     * Create a DseRegistryConfiguration resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: DseRegistryConfigurationArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DseRegistryConfigurationArgs | DseRegistryConfigurationState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DseRegistryConfigurationState | undefined;
            resourceInputs["chartRepository"] = state ? state.chartRepository : undefined;
            resourceInputs["compatibleVersionConstraints"] = state ? state.compatibleVersionConstraints : undefined;
            resourceInputs["containerRegistries"] = state ? state.containerRegistries : undefined;
            resourceInputs["defaultChartRepository"] = state ? state.defaultChartRepository : undefined;
            resourceInputs["defaultPackageName"] = state ? state.defaultPackageName : undefined;
            resourceInputs["defaultRepository"] = state ? state.defaultRepository : undefined;
            resourceInputs["defaultVersion"] = state ? state.defaultVersion : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["packageName"] = state ? state.packageName : undefined;
            resourceInputs["packageRepository"] = state ? state.packageRepository : undefined;
            resourceInputs["rdeState"] = state ? state.rdeState : undefined;
            resourceInputs["requiresVersionCompatibility"] = state ? state.requiresVersionCompatibility : undefined;
            resourceInputs["type"] = state ? state.type : undefined;
            resourceInputs["useDefaultValues"] = state ? state.useDefaultValues : undefined;
            resourceInputs["version"] = state ? state.version : undefined;
        } else {
            const args = argsOrState as DseRegistryConfigurationArgs | undefined;
            resourceInputs["chartRepository"] = args ? args.chartRepository : undefined;
            resourceInputs["containerRegistries"] = args ? args.containerRegistries : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["packageName"] = args ? args.packageName : undefined;
            resourceInputs["packageRepository"] = args ? args.packageRepository : undefined;
            resourceInputs["useDefaultValues"] = args ? args.useDefaultValues : undefined;
            resourceInputs["version"] = args ? args.version : undefined;
            resourceInputs["compatibleVersionConstraints"] = undefined /*out*/;
            resourceInputs["defaultChartRepository"] = undefined /*out*/;
            resourceInputs["defaultPackageName"] = undefined /*out*/;
            resourceInputs["defaultRepository"] = undefined /*out*/;
            resourceInputs["defaultVersion"] = undefined /*out*/;
            resourceInputs["rdeState"] = undefined /*out*/;
            resourceInputs["requiresVersionCompatibility"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DseRegistryConfiguration.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DseRegistryConfiguration resources.
 */
export interface DseRegistryConfigurationState {
    /**
     * Chart repository for Helm based images
     */
    chartRepository?: pulumi.Input<string>;
    /**
     * A set of version constrains that this Data Solution defines
     */
    compatibleVersionConstraints?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Only applies to `VCD Data Solutions` configuration. Specifies
     * credentials that can be used to authenticate to repositories. See Container Registry
     * Configuration 
     *
     *
     * <a id="container-registry"></a>
     */
    containerRegistries?: pulumi.Input<pulumi.Input<inputs.DseRegistryConfigurationContainerRegistry>[]>;
    /**
     * Default chart repository as provided by Data Solution
     */
    defaultChartRepository?: pulumi.Input<string>;
    /**
     * Default package name as provided by Data Solution
     */
    defaultPackageName?: pulumi.Input<string>;
    /**
     * Default container repository as provided by Data Solution
     */
    defaultRepository?: pulumi.Input<string>;
    /**
     * Default package version as provided by Data Solution
     */
    defaultVersion?: pulumi.Input<string>;
    /**
     * The name of Data Solution as it appears in repository configuration
     */
    name?: pulumi.Input<string>;
    /**
     * Helm package name. Only for Helm based images
     */
    packageName?: pulumi.Input<string>;
    /**
     * Package repository for container based images
     */
    packageRepository?: pulumi.Input<string>;
    /**
     * State of parent Runtime Defined Entity (RDE)
     */
    rdeState?: pulumi.Input<string>;
    /**
     * Boolean flag as defined in Data Solution
     */
    requiresVersionCompatibility?: pulumi.Input<boolean>;
    /**
     * Type of repository settings. It can be one of `PackageRepository`, `ChartRepository`
     */
    type?: pulumi.Input<string>;
    /**
     * Defines if repository settings should be inherited from Data
     * Solution itself. Default `false`
     */
    useDefaultValues?: pulumi.Input<boolean>;
    /**
     * Version of package to use. Required when `useDefaultValues` is not used.
     */
    version?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a DseRegistryConfiguration resource.
 */
export interface DseRegistryConfigurationArgs {
    /**
     * Chart repository for Helm based images
     */
    chartRepository?: pulumi.Input<string>;
    /**
     * Only applies to `VCD Data Solutions` configuration. Specifies
     * credentials that can be used to authenticate to repositories. See Container Registry
     * Configuration 
     *
     *
     * <a id="container-registry"></a>
     */
    containerRegistries?: pulumi.Input<pulumi.Input<inputs.DseRegistryConfigurationContainerRegistry>[]>;
    /**
     * The name of Data Solution as it appears in repository configuration
     */
    name?: pulumi.Input<string>;
    /**
     * Helm package name. Only for Helm based images
     */
    packageName?: pulumi.Input<string>;
    /**
     * Package repository for container based images
     */
    packageRepository?: pulumi.Input<string>;
    /**
     * Defines if repository settings should be inherited from Data
     * Solution itself. Default `false`
     */
    useDefaultValues?: pulumi.Input<boolean>;
    /**
     * Version of package to use. Required when `useDefaultValues` is not used.
     */
    version?: pulumi.Input<string>;
}
