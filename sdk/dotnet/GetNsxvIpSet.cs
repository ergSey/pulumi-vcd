// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetNsxvIpSet
    {
        /// <summary>
        /// Provides a VMware Cloud Director IP set data source. An IP set is a group of IP addresses that you can add
        ///   as the source or destination in a firewall rule or in DHCP relay configuration.
        /// 
        /// Supported in provider *v2.6+*
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
        ///     var ip_set_DS = Vcd.GetNsxvIpSet.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         Vdc = "my-org-vdc",
        ///         Name = "not-managed",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetNsxvIpSetResult> InvokeAsync(GetNsxvIpSetArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNsxvIpSetResult>("vcd:index/getNsxvIpSet:getNsxvIpSet", args ?? new GetNsxvIpSetArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a VMware Cloud Director IP set data source. An IP set is a group of IP addresses that you can add
        ///   as the source or destination in a firewall rule or in DHCP relay configuration.
        /// 
        /// Supported in provider *v2.6+*
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
        ///     var ip_set_DS = Vcd.GetNsxvIpSet.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         Vdc = "my-org-vdc",
        ///         Name = "not-managed",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetNsxvIpSetResult> Invoke(GetNsxvIpSetInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNsxvIpSetResult>("vcd:index/getNsxvIpSet:getNsxvIpSet", args ?? new GetNsxvIpSetInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a VMware Cloud Director IP set data source. An IP set is a group of IP addresses that you can add
        ///   as the source or destination in a firewall rule or in DHCP relay configuration.
        /// 
        /// Supported in provider *v2.6+*
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
        ///     var ip_set_DS = Vcd.GetNsxvIpSet.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         Vdc = "my-org-vdc",
        ///         Name = "not-managed",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetNsxvIpSetResult> Invoke(GetNsxvIpSetInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetNsxvIpSetResult>("vcd:index/getNsxvIpSet:getNsxvIpSet", args ?? new GetNsxvIpSetInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNsxvIpSetArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// IP set name for identifying the exact IP set
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of organization to use, optional if defined at provider level
        /// </summary>
        [Input("org")]
        public string? Org { get; set; }

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Input("vdc")]
        public string? Vdc { get; set; }

        public GetNsxvIpSetArgs()
        {
        }
        public static new GetNsxvIpSetArgs Empty => new GetNsxvIpSetArgs();
    }

    public sealed class GetNsxvIpSetInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// IP set name for identifying the exact IP set
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of organization to use, optional if defined at provider level
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public GetNsxvIpSetInvokeArgs()
        {
        }
        public static new GetNsxvIpSetInvokeArgs Empty => new GetNsxvIpSetInvokeArgs();
    }


    [OutputType]
    public sealed class GetNsxvIpSetResult
    {
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<string> IpAddresses;
        public readonly bool IsInheritanceAllowed;
        public readonly string Name;
        public readonly string? Org;
        public readonly string? Vdc;

        [OutputConstructor]
        private GetNsxvIpSetResult(
            string description,

            string id,

            ImmutableArray<string> ipAddresses,

            bool isInheritanceAllowed,

            string name,

            string? org,

            string? vdc)
        {
            Description = description;
            Id = id;
            IpAddresses = ipAddresses;
            IsInheritanceAllowed = isInheritanceAllowed;
            Name = name;
            Org = org;
            Vdc = vdc;
        }
    }
}
