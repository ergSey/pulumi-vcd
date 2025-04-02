// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetExternalNetworkV2
    {
        /// <summary>
        /// Provides a VMware Cloud Director External Network data source (version 2). New version of this data source uses new VCD
        /// API and is capable of handling NSX-T backed external networks as well as port group backed ones.
        /// 
        /// &gt; **Note:** This resource uses new VMware Cloud Director
        /// [OpenAPI](https://code.vmware.com/docs/11982/getting-started-with-vmware-cloud-director-openapi) and
        /// requires at least VCD *10.0+*. It supports both NSX-T and NSX-V backed networks (NSX-T *3.0+* requires VCD *10.1.1+*)
        /// 
        /// Supported in provider *v3.0+*.
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
        ///     var extNet = Vcd.GetExternalNetworkV2.Invoke(new()
        ///     {
        ///         Name = "my-nsxt-net",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetExternalNetworkV2Result> InvokeAsync(GetExternalNetworkV2Args args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetExternalNetworkV2Result>("vcd:index/getExternalNetworkV2:getExternalNetworkV2", args ?? new GetExternalNetworkV2Args(), options.WithDefaults());

        /// <summary>
        /// Provides a VMware Cloud Director External Network data source (version 2). New version of this data source uses new VCD
        /// API and is capable of handling NSX-T backed external networks as well as port group backed ones.
        /// 
        /// &gt; **Note:** This resource uses new VMware Cloud Director
        /// [OpenAPI](https://code.vmware.com/docs/11982/getting-started-with-vmware-cloud-director-openapi) and
        /// requires at least VCD *10.0+*. It supports both NSX-T and NSX-V backed networks (NSX-T *3.0+* requires VCD *10.1.1+*)
        /// 
        /// Supported in provider *v3.0+*.
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
        ///     var extNet = Vcd.GetExternalNetworkV2.Invoke(new()
        ///     {
        ///         Name = "my-nsxt-net",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetExternalNetworkV2Result> Invoke(GetExternalNetworkV2InvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetExternalNetworkV2Result>("vcd:index/getExternalNetworkV2:getExternalNetworkV2", args ?? new GetExternalNetworkV2InvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a VMware Cloud Director External Network data source (version 2). New version of this data source uses new VCD
        /// API and is capable of handling NSX-T backed external networks as well as port group backed ones.
        /// 
        /// &gt; **Note:** This resource uses new VMware Cloud Director
        /// [OpenAPI](https://code.vmware.com/docs/11982/getting-started-with-vmware-cloud-director-openapi) and
        /// requires at least VCD *10.0+*. It supports both NSX-T and NSX-V backed networks (NSX-T *3.0+* requires VCD *10.1.1+*)
        /// 
        /// Supported in provider *v3.0+*.
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
        ///     var extNet = Vcd.GetExternalNetworkV2.Invoke(new()
        ///     {
        ///         Name = "my-nsxt-net",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetExternalNetworkV2Result> Invoke(GetExternalNetworkV2InvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetExternalNetworkV2Result>("vcd:index/getExternalNetworkV2:getExternalNetworkV2", args ?? new GetExternalNetworkV2InvokeArgs(), options.WithDefaults());
    }


    public sealed class GetExternalNetworkV2Args : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// external network name
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        public GetExternalNetworkV2Args()
        {
        }
        public static new GetExternalNetworkV2Args Empty => new GetExternalNetworkV2Args();
    }

    public sealed class GetExternalNetworkV2InvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// external network name
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public GetExternalNetworkV2InvokeArgs()
        {
        }
        public static new GetExternalNetworkV2InvokeArgs Empty => new GetExternalNetworkV2InvokeArgs();
    }


    [OutputType]
    public sealed class GetExternalNetworkV2Result
    {
        public readonly string DedicatedOrgId;
        public readonly string Description;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetExternalNetworkV2IpScopeResult> IpScopes;
        public readonly string Name;
        public readonly string NatAndFirewallServiceIntention;
        public readonly ImmutableArray<Outputs.GetExternalNetworkV2NsxtNetworkResult> NsxtNetworks;
        public readonly string RouteAdvertisementIntention;
        public readonly bool UseIpSpaces;
        public readonly ImmutableArray<Outputs.GetExternalNetworkV2VsphereNetworkResult> VsphereNetworks;

        [OutputConstructor]
        private GetExternalNetworkV2Result(
            string dedicatedOrgId,

            string description,

            string id,

            ImmutableArray<Outputs.GetExternalNetworkV2IpScopeResult> ipScopes,

            string name,

            string natAndFirewallServiceIntention,

            ImmutableArray<Outputs.GetExternalNetworkV2NsxtNetworkResult> nsxtNetworks,

            string routeAdvertisementIntention,

            bool useIpSpaces,

            ImmutableArray<Outputs.GetExternalNetworkV2VsphereNetworkResult> vsphereNetworks)
        {
            DedicatedOrgId = dedicatedOrgId;
            Description = description;
            Id = id;
            IpScopes = ipScopes;
            Name = name;
            NatAndFirewallServiceIntention = natAndFirewallServiceIntention;
            NsxtNetworks = nsxtNetworks;
            RouteAdvertisementIntention = routeAdvertisementIntention;
            UseIpSpaces = useIpSpaces;
            VsphereNetworks = vsphereNetworks;
        }
    }
}
