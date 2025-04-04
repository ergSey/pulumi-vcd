// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/dseRegistryConfiguration:DseRegistryConfiguration")]
    public partial class DseRegistryConfiguration : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Chart repository for Helm based images
        /// </summary>
        [Output("chartRepository")]
        public Output<string> ChartRepository { get; private set; } = null!;

        /// <summary>
        /// A set of version constrains that this Data Solution defines
        /// </summary>
        [Output("compatibleVersionConstraints")]
        public Output<ImmutableArray<string>> CompatibleVersionConstraints { get; private set; } = null!;

        /// <summary>
        /// Only applies to `VCD Data Solutions` configuration. Specifies
        /// credentials that can be used to authenticate to repositories. See Container Registry
        /// Configuration 
        /// 
        /// 
        /// &lt;a id="container-registry"&gt;&lt;/a&gt;
        /// </summary>
        [Output("containerRegistries")]
        public Output<ImmutableArray<Outputs.DseRegistryConfigurationContainerRegistry>> ContainerRegistries { get; private set; } = null!;

        /// <summary>
        /// Default chart repository as provided by Data Solution
        /// </summary>
        [Output("defaultChartRepository")]
        public Output<string> DefaultChartRepository { get; private set; } = null!;

        /// <summary>
        /// Default package name as provided by Data Solution
        /// </summary>
        [Output("defaultPackageName")]
        public Output<string> DefaultPackageName { get; private set; } = null!;

        /// <summary>
        /// Default container repository as provided by Data Solution
        /// </summary>
        [Output("defaultRepository")]
        public Output<string> DefaultRepository { get; private set; } = null!;

        /// <summary>
        /// Default package version as provided by Data Solution
        /// </summary>
        [Output("defaultVersion")]
        public Output<string> DefaultVersion { get; private set; } = null!;

        /// <summary>
        /// The name of Data Solution as it appears in repository configuration
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Helm package name. Only for Helm based images
        /// </summary>
        [Output("packageName")]
        public Output<string> PackageName { get; private set; } = null!;

        /// <summary>
        /// Package repository for container based images
        /// </summary>
        [Output("packageRepository")]
        public Output<string> PackageRepository { get; private set; } = null!;

        /// <summary>
        /// State of parent Runtime Defined Entity (RDE)
        /// </summary>
        [Output("rdeState")]
        public Output<string> RdeState { get; private set; } = null!;

        /// <summary>
        /// Boolean flag as defined in Data Solution
        /// </summary>
        [Output("requiresVersionCompatibility")]
        public Output<bool> RequiresVersionCompatibility { get; private set; } = null!;

        /// <summary>
        /// Type of repository settings. It can be one of `PackageRepository`, `ChartRepository`
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// Defines if repository settings should be inherited from Data
        /// Solution itself. Default `false`
        /// </summary>
        [Output("useDefaultValues")]
        public Output<bool?> UseDefaultValues { get; private set; } = null!;

        /// <summary>
        /// Version of package to use. Required when `use_default_values` is not used.
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a DseRegistryConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DseRegistryConfiguration(string name, DseRegistryConfigurationArgs? args = null, CustomResourceOptions? options = null)
            : base("vcd:index/dseRegistryConfiguration:DseRegistryConfiguration", name, args ?? new DseRegistryConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DseRegistryConfiguration(string name, Input<string> id, DseRegistryConfigurationState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/dseRegistryConfiguration:DseRegistryConfiguration", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/vmware/terraform-provider-vcd/releases/download/v3.14.1/",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DseRegistryConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DseRegistryConfiguration Get(string name, Input<string> id, DseRegistryConfigurationState? state = null, CustomResourceOptions? options = null)
        {
            return new DseRegistryConfiguration(name, id, state, options);
        }
    }

    public sealed class DseRegistryConfigurationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Chart repository for Helm based images
        /// </summary>
        [Input("chartRepository")]
        public Input<string>? ChartRepository { get; set; }

        [Input("containerRegistries")]
        private InputList<Inputs.DseRegistryConfigurationContainerRegistryArgs>? _containerRegistries;

        /// <summary>
        /// Only applies to `VCD Data Solutions` configuration. Specifies
        /// credentials that can be used to authenticate to repositories. See Container Registry
        /// Configuration 
        /// 
        /// 
        /// &lt;a id="container-registry"&gt;&lt;/a&gt;
        /// </summary>
        public InputList<Inputs.DseRegistryConfigurationContainerRegistryArgs> ContainerRegistries
        {
            get => _containerRegistries ?? (_containerRegistries = new InputList<Inputs.DseRegistryConfigurationContainerRegistryArgs>());
            set => _containerRegistries = value;
        }

        /// <summary>
        /// The name of Data Solution as it appears in repository configuration
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Helm package name. Only for Helm based images
        /// </summary>
        [Input("packageName")]
        public Input<string>? PackageName { get; set; }

        /// <summary>
        /// Package repository for container based images
        /// </summary>
        [Input("packageRepository")]
        public Input<string>? PackageRepository { get; set; }

        /// <summary>
        /// Defines if repository settings should be inherited from Data
        /// Solution itself. Default `false`
        /// </summary>
        [Input("useDefaultValues")]
        public Input<bool>? UseDefaultValues { get; set; }

        /// <summary>
        /// Version of package to use. Required when `use_default_values` is not used.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public DseRegistryConfigurationArgs()
        {
        }
        public static new DseRegistryConfigurationArgs Empty => new DseRegistryConfigurationArgs();
    }

    public sealed class DseRegistryConfigurationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Chart repository for Helm based images
        /// </summary>
        [Input("chartRepository")]
        public Input<string>? ChartRepository { get; set; }

        [Input("compatibleVersionConstraints")]
        private InputList<string>? _compatibleVersionConstraints;

        /// <summary>
        /// A set of version constrains that this Data Solution defines
        /// </summary>
        public InputList<string> CompatibleVersionConstraints
        {
            get => _compatibleVersionConstraints ?? (_compatibleVersionConstraints = new InputList<string>());
            set => _compatibleVersionConstraints = value;
        }

        [Input("containerRegistries")]
        private InputList<Inputs.DseRegistryConfigurationContainerRegistryGetArgs>? _containerRegistries;

        /// <summary>
        /// Only applies to `VCD Data Solutions` configuration. Specifies
        /// credentials that can be used to authenticate to repositories. See Container Registry
        /// Configuration 
        /// 
        /// 
        /// &lt;a id="container-registry"&gt;&lt;/a&gt;
        /// </summary>
        public InputList<Inputs.DseRegistryConfigurationContainerRegistryGetArgs> ContainerRegistries
        {
            get => _containerRegistries ?? (_containerRegistries = new InputList<Inputs.DseRegistryConfigurationContainerRegistryGetArgs>());
            set => _containerRegistries = value;
        }

        /// <summary>
        /// Default chart repository as provided by Data Solution
        /// </summary>
        [Input("defaultChartRepository")]
        public Input<string>? DefaultChartRepository { get; set; }

        /// <summary>
        /// Default package name as provided by Data Solution
        /// </summary>
        [Input("defaultPackageName")]
        public Input<string>? DefaultPackageName { get; set; }

        /// <summary>
        /// Default container repository as provided by Data Solution
        /// </summary>
        [Input("defaultRepository")]
        public Input<string>? DefaultRepository { get; set; }

        /// <summary>
        /// Default package version as provided by Data Solution
        /// </summary>
        [Input("defaultVersion")]
        public Input<string>? DefaultVersion { get; set; }

        /// <summary>
        /// The name of Data Solution as it appears in repository configuration
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Helm package name. Only for Helm based images
        /// </summary>
        [Input("packageName")]
        public Input<string>? PackageName { get; set; }

        /// <summary>
        /// Package repository for container based images
        /// </summary>
        [Input("packageRepository")]
        public Input<string>? PackageRepository { get; set; }

        /// <summary>
        /// State of parent Runtime Defined Entity (RDE)
        /// </summary>
        [Input("rdeState")]
        public Input<string>? RdeState { get; set; }

        /// <summary>
        /// Boolean flag as defined in Data Solution
        /// </summary>
        [Input("requiresVersionCompatibility")]
        public Input<bool>? RequiresVersionCompatibility { get; set; }

        /// <summary>
        /// Type of repository settings. It can be one of `PackageRepository`, `ChartRepository`
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// Defines if repository settings should be inherited from Data
        /// Solution itself. Default `false`
        /// </summary>
        [Input("useDefaultValues")]
        public Input<bool>? UseDefaultValues { get; set; }

        /// <summary>
        /// Version of package to use. Required when `use_default_values` is not used.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public DseRegistryConfigurationState()
        {
        }
        public static new DseRegistryConfigurationState Empty => new DseRegistryConfigurationState();
    }
}
