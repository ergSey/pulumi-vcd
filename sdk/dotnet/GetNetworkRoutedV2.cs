// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetNetworkRoutedV2
    {
        /// <summary>
        /// Provides a VMware Cloud Director Org VDC routed Network data source to read data or reference  existing network
        /// (backed by NSX-T or NSX-V).
        /// 
        /// Supported in provider *v3.2+* for both NSX-T and NSX-V VDCs.
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
        ///     var main = Vcd.GetNsxtEdgegateway.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         Name = "main-edge",
        ///     });
        /// 
        ///     var net = Vcd.GetNetworkRoutedV2.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         EdgeGatewayId = main.Apply(getNsxtEdgegatewayResult =&gt; getNsxtEdgegatewayResult.Id),
        ///         Name = "my-net",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Filter arguments
        /// 
        /// * `name_regex` - (Optional) matches the name using a regular expression.
        /// * `ip` - (Optional) matches the IP of the resource using a regular expression.
        /// 
        /// See [Filters reference](https://www.terraform.io/providers/vmware/vcd/latest/docs/guides/data_source_filters) for details and examples.
        /// </summary>
        public static Task<GetNetworkRoutedV2Result> InvokeAsync(GetNetworkRoutedV2Args? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNetworkRoutedV2Result>("vcd:index/getNetworkRoutedV2:getNetworkRoutedV2", args ?? new GetNetworkRoutedV2Args(), options.WithDefaults());

        /// <summary>
        /// Provides a VMware Cloud Director Org VDC routed Network data source to read data or reference  existing network
        /// (backed by NSX-T or NSX-V).
        /// 
        /// Supported in provider *v3.2+* for both NSX-T and NSX-V VDCs.
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
        ///     var main = Vcd.GetNsxtEdgegateway.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         Name = "main-edge",
        ///     });
        /// 
        ///     var net = Vcd.GetNetworkRoutedV2.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         EdgeGatewayId = main.Apply(getNsxtEdgegatewayResult =&gt; getNsxtEdgegatewayResult.Id),
        ///         Name = "my-net",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Filter arguments
        /// 
        /// * `name_regex` - (Optional) matches the name using a regular expression.
        /// * `ip` - (Optional) matches the IP of the resource using a regular expression.
        /// 
        /// See [Filters reference](https://www.terraform.io/providers/vmware/vcd/latest/docs/guides/data_source_filters) for details and examples.
        /// </summary>
        public static Output<GetNetworkRoutedV2Result> Invoke(GetNetworkRoutedV2InvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNetworkRoutedV2Result>("vcd:index/getNetworkRoutedV2:getNetworkRoutedV2", args ?? new GetNetworkRoutedV2InvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a VMware Cloud Director Org VDC routed Network data source to read data or reference  existing network
        /// (backed by NSX-T or NSX-V).
        /// 
        /// Supported in provider *v3.2+* for both NSX-T and NSX-V VDCs.
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
        ///     var main = Vcd.GetNsxtEdgegateway.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         Name = "main-edge",
        ///     });
        /// 
        ///     var net = Vcd.GetNetworkRoutedV2.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         EdgeGatewayId = main.Apply(getNsxtEdgegatewayResult =&gt; getNsxtEdgegatewayResult.Id),
        ///         Name = "my-net",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Filter arguments
        /// 
        /// * `name_regex` - (Optional) matches the name using a regular expression.
        /// * `ip` - (Optional) matches the IP of the resource using a regular expression.
        /// 
        /// See [Filters reference](https://www.terraform.io/providers/vmware/vcd/latest/docs/guides/data_source_filters) for details and examples.
        /// </summary>
        public static Output<GetNetworkRoutedV2Result> Invoke(GetNetworkRoutedV2InvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetNetworkRoutedV2Result>("vcd:index/getNetworkRoutedV2:getNetworkRoutedV2", args ?? new GetNetworkRoutedV2InvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNetworkRoutedV2Args : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Replaces `vdc` field and helps to identify exact Org
        /// Network
        /// </summary>
        [Input("edgeGatewayId")]
        public string? EdgeGatewayId { get; set; }

        /// <summary>
        /// Retrieves the data source using one or more filter parameters. **Note**
        /// filters do not support searching for networks in VDC Groups.
        /// </summary>
        [Input("filter")]
        public Inputs.GetNetworkRoutedV2FilterArgs? Filter { get; set; }

        /// <summary>
        /// A unique name for the network (optional when `filter` is used)
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level
        /// </summary>
        [Input("org")]
        public string? Org { get; set; }

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level. **Deprecated**
        /// in favor of `edge_gateway_id` field.
        /// </summary>
        [Input("vdc")]
        public string? Vdc { get; set; }

        public GetNetworkRoutedV2Args()
        {
        }
        public static new GetNetworkRoutedV2Args Empty => new GetNetworkRoutedV2Args();
    }

    public sealed class GetNetworkRoutedV2InvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Replaces `vdc` field and helps to identify exact Org
        /// Network
        /// </summary>
        [Input("edgeGatewayId")]
        public Input<string>? EdgeGatewayId { get; set; }

        /// <summary>
        /// Retrieves the data source using one or more filter parameters. **Note**
        /// filters do not support searching for networks in VDC Groups.
        /// </summary>
        [Input("filter")]
        public Input<Inputs.GetNetworkRoutedV2FilterInputArgs>? Filter { get; set; }

        /// <summary>
        /// A unique name for the network (optional when `filter` is used)
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level. **Deprecated**
        /// in favor of `edge_gateway_id` field.
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public GetNetworkRoutedV2InvokeArgs()
        {
        }
        public static new GetNetworkRoutedV2InvokeArgs Empty => new GetNetworkRoutedV2InvokeArgs();
    }


    [OutputType]
    public sealed class GetNetworkRoutedV2Result
    {
        public readonly string Description;
        public readonly string Dns1;
        public readonly string Dns2;
        public readonly string DnsSuffix;
        public readonly bool DualStackEnabled;
        public readonly string EdgeGatewayId;
        public readonly Outputs.GetNetworkRoutedV2FilterResult? Filter;
        public readonly string Gateway;
        public readonly bool GuestVlanAllowed;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string InterfaceType;
        public readonly ImmutableDictionary<string, string> Metadata;
        public readonly ImmutableArray<Outputs.GetNetworkRoutedV2MetadataEntryResult> MetadataEntries;
        public readonly string? Name;
        public readonly string? Org;
        /// <summary>
        /// Parent VDC or VDC Group ID.
        /// </summary>
        public readonly string OwnerId;
        public readonly int PrefixLength;
        public readonly bool RouteAdvertisementEnabled;
        public readonly string SecondaryGateway;
        public readonly string SecondaryPrefixLength;
        public readonly ImmutableArray<Outputs.GetNetworkRoutedV2SecondaryStaticIpPoolResult> SecondaryStaticIpPools;
        public readonly ImmutableArray<Outputs.GetNetworkRoutedV2StaticIpPoolResult> StaticIpPools;
        public readonly string? Vdc;

        [OutputConstructor]
        private GetNetworkRoutedV2Result(
            string description,

            string dns1,

            string dns2,

            string dnsSuffix,

            bool dualStackEnabled,

            string edgeGatewayId,

            Outputs.GetNetworkRoutedV2FilterResult? filter,

            string gateway,

            bool guestVlanAllowed,

            string id,

            string interfaceType,

            ImmutableDictionary<string, string> metadata,

            ImmutableArray<Outputs.GetNetworkRoutedV2MetadataEntryResult> metadataEntries,

            string? name,

            string? org,

            string ownerId,

            int prefixLength,

            bool routeAdvertisementEnabled,

            string secondaryGateway,

            string secondaryPrefixLength,

            ImmutableArray<Outputs.GetNetworkRoutedV2SecondaryStaticIpPoolResult> secondaryStaticIpPools,

            ImmutableArray<Outputs.GetNetworkRoutedV2StaticIpPoolResult> staticIpPools,

            string? vdc)
        {
            Description = description;
            Dns1 = dns1;
            Dns2 = dns2;
            DnsSuffix = dnsSuffix;
            DualStackEnabled = dualStackEnabled;
            EdgeGatewayId = edgeGatewayId;
            Filter = filter;
            Gateway = gateway;
            GuestVlanAllowed = guestVlanAllowed;
            Id = id;
            InterfaceType = interfaceType;
            Metadata = metadata;
            MetadataEntries = metadataEntries;
            Name = name;
            Org = org;
            OwnerId = ownerId;
            PrefixLength = prefixLength;
            RouteAdvertisementEnabled = routeAdvertisementEnabled;
            SecondaryGateway = secondaryGateway;
            SecondaryPrefixLength = secondaryPrefixLength;
            SecondaryStaticIpPools = secondaryStaticIpPools;
            StaticIpPools = staticIpPools;
            Vdc = vdc;
        }
    }
}
