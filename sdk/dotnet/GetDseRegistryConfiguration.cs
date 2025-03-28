// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetDseRegistryConfiguration
    {
        /// <summary>
        /// Supported in provider *v3.13+* with Data Solution Extension.
        /// 
        /// Provides a data source to read Data Solution Extension (DSE) registry configuration.
        /// 
        /// &gt; Only `System Administrator` can use this data source.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vcd = Pulumi.Vcd;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var mongodb = Vcd.GetDseRegistryConfiguration.Invoke(new()
        ///     {
        ///         Name = "MongoDB",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetDseRegistryConfigurationResult> InvokeAsync(GetDseRegistryConfigurationArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetDseRegistryConfigurationResult>("vcd:index/getDseRegistryConfiguration:getDseRegistryConfiguration", args ?? new GetDseRegistryConfigurationArgs(), options.WithDefaults());

        /// <summary>
        /// Supported in provider *v3.13+* with Data Solution Extension.
        /// 
        /// Provides a data source to read Data Solution Extension (DSE) registry configuration.
        /// 
        /// &gt; Only `System Administrator` can use this data source.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vcd = Pulumi.Vcd;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var mongodb = Vcd.GetDseRegistryConfiguration.Invoke(new()
        ///     {
        ///         Name = "MongoDB",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetDseRegistryConfigurationResult> Invoke(GetDseRegistryConfigurationInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetDseRegistryConfigurationResult>("vcd:index/getDseRegistryConfiguration:getDseRegistryConfiguration", args ?? new GetDseRegistryConfigurationInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Supported in provider *v3.13+* with Data Solution Extension.
        /// 
        /// Provides a data source to read Data Solution Extension (DSE) registry configuration.
        /// 
        /// &gt; Only `System Administrator` can use this data source.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vcd = Pulumi.Vcd;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var mongodb = Vcd.GetDseRegistryConfiguration.Invoke(new()
        ///     {
        ///         Name = "MongoDB",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetDseRegistryConfigurationResult> Invoke(GetDseRegistryConfigurationInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetDseRegistryConfigurationResult>("vcd:index/getDseRegistryConfiguration:getDseRegistryConfiguration", args ?? new GetDseRegistryConfigurationInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDseRegistryConfigurationArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of Data Solution as it appears in repository configuration
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetDseRegistryConfigurationArgs()
        {
        }
        public static new GetDseRegistryConfigurationArgs Empty => new GetDseRegistryConfigurationArgs();
    }

    public sealed class GetDseRegistryConfigurationInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of Data Solution as it appears in repository configuration
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetDseRegistryConfigurationInvokeArgs()
        {
        }
        public static new GetDseRegistryConfigurationInvokeArgs Empty => new GetDseRegistryConfigurationInvokeArgs();
    }


    [OutputType]
    public sealed class GetDseRegistryConfigurationResult
    {
        public readonly string ChartRepository;
        public readonly ImmutableArray<string> CompatibleVersionConstraints;
        public readonly ImmutableArray<Outputs.GetDseRegistryConfigurationContainerRegistryResult> ContainerRegistries;
        public readonly string DefaultChartRepository;
        public readonly string DefaultPackageName;
        public readonly string DefaultRepository;
        public readonly string DefaultVersion;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string PackageName;
        public readonly string PackageRepository;
        public readonly string RdeState;
        public readonly bool RequiresVersionCompatibility;
        public readonly string Type;
        public readonly string Version;

        [OutputConstructor]
        private GetDseRegistryConfigurationResult(
            string chartRepository,

            ImmutableArray<string> compatibleVersionConstraints,

            ImmutableArray<Outputs.GetDseRegistryConfigurationContainerRegistryResult> containerRegistries,

            string defaultChartRepository,

            string defaultPackageName,

            string defaultRepository,

            string defaultVersion,

            string id,

            string name,

            string packageName,

            string packageRepository,

            string rdeState,

            bool requiresVersionCompatibility,

            string type,

            string version)
        {
            ChartRepository = chartRepository;
            CompatibleVersionConstraints = compatibleVersionConstraints;
            ContainerRegistries = containerRegistries;
            DefaultChartRepository = defaultChartRepository;
            DefaultPackageName = defaultPackageName;
            DefaultRepository = defaultRepository;
            DefaultVersion = defaultVersion;
            Id = id;
            Name = name;
            PackageName = packageName;
            PackageRepository = packageRepository;
            RdeState = rdeState;
            RequiresVersionCompatibility = requiresVersionCompatibility;
            Type = type;
            Version = version;
        }
    }
}
