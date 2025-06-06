// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetNsxtAlbSettings
    {
        /// <summary>
        /// Supported in provider *v3.5+* and VCD 10.2+ with NSX-T and ALB.
        /// 
        /// Provides a data source to read ALB General Settings for particular NSX-T Edge Gateway.
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
        ///     var existing = Vcd.GetNsxtEdgegateway.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         Vdc = "nsxt-vdc",
        ///         Name = "nsxt-gw",
        ///     });
        /// 
        ///     var test = Vcd.GetNsxtAlbSettings.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         EdgeGatewayId = existing.Apply(getNsxtEdgegatewayResult =&gt; getNsxtEdgegatewayResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetNsxtAlbSettingsResult> InvokeAsync(GetNsxtAlbSettingsArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNsxtAlbSettingsResult>("vcd:index/getNsxtAlbSettings:getNsxtAlbSettings", args ?? new GetNsxtAlbSettingsArgs(), options.WithDefaults());

        /// <summary>
        /// Supported in provider *v3.5+* and VCD 10.2+ with NSX-T and ALB.
        /// 
        /// Provides a data source to read ALB General Settings for particular NSX-T Edge Gateway.
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
        ///     var existing = Vcd.GetNsxtEdgegateway.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         Vdc = "nsxt-vdc",
        ///         Name = "nsxt-gw",
        ///     });
        /// 
        ///     var test = Vcd.GetNsxtAlbSettings.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         EdgeGatewayId = existing.Apply(getNsxtEdgegatewayResult =&gt; getNsxtEdgegatewayResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetNsxtAlbSettingsResult> Invoke(GetNsxtAlbSettingsInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNsxtAlbSettingsResult>("vcd:index/getNsxtAlbSettings:getNsxtAlbSettings", args ?? new GetNsxtAlbSettingsInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Supported in provider *v3.5+* and VCD 10.2+ with NSX-T and ALB.
        /// 
        /// Provides a data source to read ALB General Settings for particular NSX-T Edge Gateway.
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
        ///     var existing = Vcd.GetNsxtEdgegateway.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         Vdc = "nsxt-vdc",
        ///         Name = "nsxt-gw",
        ///     });
        /// 
        ///     var test = Vcd.GetNsxtAlbSettings.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         EdgeGatewayId = existing.Apply(getNsxtEdgegatewayResult =&gt; getNsxtEdgegatewayResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetNsxtAlbSettingsResult> Invoke(GetNsxtAlbSettingsInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetNsxtAlbSettingsResult>("vcd:index/getNsxtAlbSettings:getNsxtAlbSettings", args ?? new GetNsxtAlbSettingsInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNsxtAlbSettingsArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// An ID of NSX-T Edge Gateway. Can be looked up using
        /// [vcd.NsxtEdgegateway](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_edgegateway) data source
        /// </summary>
        [Input("edgeGatewayId", required: true)]
        public string EdgeGatewayId { get; set; } = null!;

        /// <summary>
        /// The name of organization to which the edge gateway belongs. Optional if defined at provider level.
        /// </summary>
        [Input("org")]
        public string? Org { get; set; }

        [Input("serviceNetworkSpecification")]
        public string? ServiceNetworkSpecification { get; set; }

        [Input("vdc")]
        public string? Vdc { get; set; }

        public GetNsxtAlbSettingsArgs()
        {
        }
        public static new GetNsxtAlbSettingsArgs Empty => new GetNsxtAlbSettingsArgs();
    }

    public sealed class GetNsxtAlbSettingsInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// An ID of NSX-T Edge Gateway. Can be looked up using
        /// [vcd.NsxtEdgegateway](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_edgegateway) data source
        /// </summary>
        [Input("edgeGatewayId", required: true)]
        public Input<string> EdgeGatewayId { get; set; } = null!;

        /// <summary>
        /// The name of organization to which the edge gateway belongs. Optional if defined at provider level.
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        [Input("serviceNetworkSpecification")]
        public Input<string>? ServiceNetworkSpecification { get; set; }

        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public GetNsxtAlbSettingsInvokeArgs()
        {
        }
        public static new GetNsxtAlbSettingsInvokeArgs Empty => new GetNsxtAlbSettingsInvokeArgs();
    }


    [OutputType]
    public sealed class GetNsxtAlbSettingsResult
    {
        public readonly string EdgeGatewayId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Ipv6ServiceNetworkSpecification;
        public readonly bool IsActive;
        public readonly bool IsTransparentModeEnabled;
        public readonly string? Org;
        public readonly string ServiceNetworkSpecification;
        public readonly string SupportedFeatureSet;
        public readonly string? Vdc;

        [OutputConstructor]
        private GetNsxtAlbSettingsResult(
            string edgeGatewayId,

            string id,

            string ipv6ServiceNetworkSpecification,

            bool isActive,

            bool isTransparentModeEnabled,

            string? org,

            string serviceNetworkSpecification,

            string supportedFeatureSet,

            string? vdc)
        {
            EdgeGatewayId = edgeGatewayId;
            Id = id;
            Ipv6ServiceNetworkSpecification = ipv6ServiceNetworkSpecification;
            IsActive = isActive;
            IsTransparentModeEnabled = isTransparentModeEnabled;
            Org = org;
            ServiceNetworkSpecification = serviceNetworkSpecification;
            SupportedFeatureSet = supportedFeatureSet;
            Vdc = vdc;
        }
    }
}
