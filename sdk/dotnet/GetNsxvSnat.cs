// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetNsxvSnat
    {
        /// <summary>
        /// Provides a VMware Cloud Director SNAT data source for advanced edge gateways (NSX-V). This can be used to
        /// read existing rule by ID and use its attributes in other resources.
        /// 
        /// &gt; **Note:** This data source requires advanced edge gateway.
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
        ///     var my_rule = Vcd.GetNsxvSnat.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         Vdc = "my-org-vdc",
        ///         EdgeGateway = "my-edge-gw",
        ///         RuleId = "197867",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetNsxvSnatResult> InvokeAsync(GetNsxvSnatArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNsxvSnatResult>("vcd:index/getNsxvSnat:getNsxvSnat", args ?? new GetNsxvSnatArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a VMware Cloud Director SNAT data source for advanced edge gateways (NSX-V). This can be used to
        /// read existing rule by ID and use its attributes in other resources.
        /// 
        /// &gt; **Note:** This data source requires advanced edge gateway.
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
        ///     var my_rule = Vcd.GetNsxvSnat.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         Vdc = "my-org-vdc",
        ///         EdgeGateway = "my-edge-gw",
        ///         RuleId = "197867",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetNsxvSnatResult> Invoke(GetNsxvSnatInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNsxvSnatResult>("vcd:index/getNsxvSnat:getNsxvSnat", args ?? new GetNsxvSnatInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a VMware Cloud Director SNAT data source for advanced edge gateways (NSX-V). This can be used to
        /// read existing rule by ID and use its attributes in other resources.
        /// 
        /// &gt; **Note:** This data source requires advanced edge gateway.
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
        ///     var my_rule = Vcd.GetNsxvSnat.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         Vdc = "my-org-vdc",
        ///         EdgeGateway = "my-edge-gw",
        ///         RuleId = "197867",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetNsxvSnatResult> Invoke(GetNsxvSnatInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetNsxvSnatResult>("vcd:index/getNsxvSnat:getNsxvSnat", args ?? new GetNsxvSnatInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNsxvSnatArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the edge gateway on which to apply the SNAT rule.
        /// </summary>
        [Input("edgeGateway", required: true)]
        public string EdgeGateway { get; set; } = null!;

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.
        /// </summary>
        [Input("org")]
        public string? Org { get; set; }

        /// <summary>
        /// ID of SNAT rule as shown in the UI.
        /// </summary>
        [Input("ruleId", required: true)]
        public string RuleId { get; set; } = null!;

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level.
        /// </summary>
        [Input("vdc")]
        public string? Vdc { get; set; }

        public GetNsxvSnatArgs()
        {
        }
        public static new GetNsxvSnatArgs Empty => new GetNsxvSnatArgs();
    }

    public sealed class GetNsxvSnatInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the edge gateway on which to apply the SNAT rule.
        /// </summary>
        [Input("edgeGateway", required: true)]
        public Input<string> EdgeGateway { get; set; } = null!;

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// ID of SNAT rule as shown in the UI.
        /// </summary>
        [Input("ruleId", required: true)]
        public Input<string> RuleId { get; set; } = null!;

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level.
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public GetNsxvSnatInvokeArgs()
        {
        }
        public static new GetNsxvSnatInvokeArgs Empty => new GetNsxvSnatInvokeArgs();
    }


    [OutputType]
    public sealed class GetNsxvSnatResult
    {
        public readonly string Description;
        public readonly string EdgeGateway;
        public readonly bool Enabled;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool LoggingEnabled;
        public readonly string NetworkName;
        public readonly string NetworkType;
        public readonly string? Org;
        public readonly string OriginalAddress;
        public readonly string RuleId;
        public readonly int RuleTag;
        public readonly string RuleType;
        public readonly string TranslatedAddress;
        public readonly string? Vdc;

        [OutputConstructor]
        private GetNsxvSnatResult(
            string description,

            string edgeGateway,

            bool enabled,

            string id,

            bool loggingEnabled,

            string networkName,

            string networkType,

            string? org,

            string originalAddress,

            string ruleId,

            int ruleTag,

            string ruleType,

            string translatedAddress,

            string? vdc)
        {
            Description = description;
            EdgeGateway = edgeGateway;
            Enabled = enabled;
            Id = id;
            LoggingEnabled = loggingEnabled;
            NetworkName = networkName;
            NetworkType = networkType;
            Org = org;
            OriginalAddress = originalAddress;
            RuleId = ruleId;
            RuleTag = ruleTag;
            RuleType = ruleType;
            TranslatedAddress = translatedAddress;
            Vdc = vdc;
        }
    }
}
