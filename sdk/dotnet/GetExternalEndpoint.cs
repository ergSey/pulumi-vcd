// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetExternalEndpoint
    {
        /// <summary>
        /// Supported in provider *v3.14+* and VCD 10.4.3+.
        /// 
        /// Provides a data source to read External Endpoints in VMware Cloud Director. An External Endpoint holds information for the
        /// HTTPS endpoint which requests will be proxied to when using a [`vcd.ApiFilter`](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/api_filter).
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
        ///     var externalEndpoint1 = Vcd.GetExternalEndpoint.Invoke(new()
        ///     {
        ///         Vendor = "vmware",
        ///         Name = "my-endpoint",
        ///         Version = "1.0.0",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetExternalEndpointResult> InvokeAsync(GetExternalEndpointArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetExternalEndpointResult>("vcd:index/getExternalEndpoint:getExternalEndpoint", args ?? new GetExternalEndpointArgs(), options.WithDefaults());

        /// <summary>
        /// Supported in provider *v3.14+* and VCD 10.4.3+.
        /// 
        /// Provides a data source to read External Endpoints in VMware Cloud Director. An External Endpoint holds information for the
        /// HTTPS endpoint which requests will be proxied to when using a [`vcd.ApiFilter`](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/api_filter).
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
        ///     var externalEndpoint1 = Vcd.GetExternalEndpoint.Invoke(new()
        ///     {
        ///         Vendor = "vmware",
        ///         Name = "my-endpoint",
        ///         Version = "1.0.0",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetExternalEndpointResult> Invoke(GetExternalEndpointInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetExternalEndpointResult>("vcd:index/getExternalEndpoint:getExternalEndpoint", args ?? new GetExternalEndpointInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Supported in provider *v3.14+* and VCD 10.4.3+.
        /// 
        /// Provides a data source to read External Endpoints in VMware Cloud Director. An External Endpoint holds information for the
        /// HTTPS endpoint which requests will be proxied to when using a [`vcd.ApiFilter`](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/api_filter).
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
        ///     var externalEndpoint1 = Vcd.GetExternalEndpoint.Invoke(new()
        ///     {
        ///         Vendor = "vmware",
        ///         Name = "my-endpoint",
        ///         Version = "1.0.0",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetExternalEndpointResult> Invoke(GetExternalEndpointInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetExternalEndpointResult>("vcd:index/getExternalEndpoint:getExternalEndpoint", args ?? new GetExternalEndpointInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetExternalEndpointArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the External Endpoint
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The vendor name of the External Endpoint
        /// </summary>
        [Input("vendor", required: true)]
        public string Vendor { get; set; } = null!;

        /// <summary>
        /// The version of the External Endpoint
        /// </summary>
        [Input("version", required: true)]
        public string Version { get; set; } = null!;

        public GetExternalEndpointArgs()
        {
        }
        public static new GetExternalEndpointArgs Empty => new GetExternalEndpointArgs();
    }

    public sealed class GetExternalEndpointInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the External Endpoint
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The vendor name of the External Endpoint
        /// </summary>
        [Input("vendor", required: true)]
        public Input<string> Vendor { get; set; } = null!;

        /// <summary>
        /// The version of the External Endpoint
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public GetExternalEndpointInvokeArgs()
        {
        }
        public static new GetExternalEndpointInvokeArgs Empty => new GetExternalEndpointInvokeArgs();
    }


    [OutputType]
    public sealed class GetExternalEndpointResult
    {
        public readonly string Description;
        public readonly bool Enabled;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string RootUrl;
        public readonly string Vendor;
        public readonly string Version;

        [OutputConstructor]
        private GetExternalEndpointResult(
            string description,

            bool enabled,

            string id,

            string name,

            string rootUrl,

            string vendor,

            string version)
        {
            Description = description;
            Enabled = enabled;
            Id = id;
            Name = name;
            RootUrl = rootUrl;
            Vendor = vendor;
            Version = version;
        }
    }
}
