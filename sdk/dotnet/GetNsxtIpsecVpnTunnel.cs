// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetNsxtIpsecVpnTunnel
    {
        /// <summary>
        /// Supported in provider *v3.3+* and VCD 10.1+ with NSX-T backed VDCs.
        /// 
        /// Provides a data source to read NSX-T IPsec VPN Tunnel. You can configure site-to-site connectivity between an NSX-T Data
        /// Center Edge Gateway and remote sites. The remote sites must use NSX-T Data Center, have third-party hardware routers,
        /// or VPN gateways that support IPSec.
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
        ///     var tunnel1 = Vcd.GetNsxtIpsecVpnTunnel.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         EdgeGatewayId = existing.Id,
        ///         Name = "tunnel-1",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetNsxtIpsecVpnTunnelResult> InvokeAsync(GetNsxtIpsecVpnTunnelArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNsxtIpsecVpnTunnelResult>("vcd:index/getNsxtIpsecVpnTunnel:getNsxtIpsecVpnTunnel", args ?? new GetNsxtIpsecVpnTunnelArgs(), options.WithDefaults());

        /// <summary>
        /// Supported in provider *v3.3+* and VCD 10.1+ with NSX-T backed VDCs.
        /// 
        /// Provides a data source to read NSX-T IPsec VPN Tunnel. You can configure site-to-site connectivity between an NSX-T Data
        /// Center Edge Gateway and remote sites. The remote sites must use NSX-T Data Center, have third-party hardware routers,
        /// or VPN gateways that support IPSec.
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
        ///     var tunnel1 = Vcd.GetNsxtIpsecVpnTunnel.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         EdgeGatewayId = existing.Id,
        ///         Name = "tunnel-1",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetNsxtIpsecVpnTunnelResult> Invoke(GetNsxtIpsecVpnTunnelInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNsxtIpsecVpnTunnelResult>("vcd:index/getNsxtIpsecVpnTunnel:getNsxtIpsecVpnTunnel", args ?? new GetNsxtIpsecVpnTunnelInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Supported in provider *v3.3+* and VCD 10.1+ with NSX-T backed VDCs.
        /// 
        /// Provides a data source to read NSX-T IPsec VPN Tunnel. You can configure site-to-site connectivity between an NSX-T Data
        /// Center Edge Gateway and remote sites. The remote sites must use NSX-T Data Center, have third-party hardware routers,
        /// or VPN gateways that support IPSec.
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
        ///     var tunnel1 = Vcd.GetNsxtIpsecVpnTunnel.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         EdgeGatewayId = existing.Id,
        ///         Name = "tunnel-1",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetNsxtIpsecVpnTunnelResult> Invoke(GetNsxtIpsecVpnTunnelInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetNsxtIpsecVpnTunnelResult>("vcd:index/getNsxtIpsecVpnTunnel:getNsxtIpsecVpnTunnel", args ?? new GetNsxtIpsecVpnTunnelInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNsxtIpsecVpnTunnelArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Edge Gateway (NSX-T only). Can be looked up using `vcd.NsxtEdgegateway`
        /// data source
        /// </summary>
        [Input("edgeGatewayId", required: true)]
        public string EdgeGatewayId { get; set; } = null!;

        /// <summary>
        /// Name of existing IPsec VPN Tunnel
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful
        /// when connected as sysadmin working across different organisations.
        /// </summary>
        [Input("org")]
        public string? Org { get; set; }

        [Input("vdc")]
        public string? Vdc { get; set; }

        public GetNsxtIpsecVpnTunnelArgs()
        {
        }
        public static new GetNsxtIpsecVpnTunnelArgs Empty => new GetNsxtIpsecVpnTunnelArgs();
    }

    public sealed class GetNsxtIpsecVpnTunnelInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The ID of the Edge Gateway (NSX-T only). Can be looked up using `vcd.NsxtEdgegateway`
        /// data source
        /// </summary>
        [Input("edgeGatewayId", required: true)]
        public Input<string> EdgeGatewayId { get; set; } = null!;

        /// <summary>
        /// Name of existing IPsec VPN Tunnel
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful
        /// when connected as sysadmin working across different organisations.
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public GetNsxtIpsecVpnTunnelInvokeArgs()
        {
        }
        public static new GetNsxtIpsecVpnTunnelInvokeArgs Empty => new GetNsxtIpsecVpnTunnelInvokeArgs();
    }


    [OutputType]
    public sealed class GetNsxtIpsecVpnTunnelResult
    {
        public readonly string AuthenticationMode;
        public readonly string CaCertificateId;
        public readonly string CertificateId;
        public readonly string Description;
        public readonly string EdgeGatewayId;
        public readonly bool Enabled;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string IkeFailReason;
        public readonly string IkeServiceStatus;
        public readonly string LocalIpAddress;
        public readonly ImmutableArray<string> LocalNetworks;
        public readonly bool Logging;
        public readonly string Name;
        public readonly string? Org;
        public readonly string PreSharedKey;
        public readonly string RemoteId;
        public readonly string RemoteIpAddress;
        public readonly ImmutableArray<string> RemoteNetworks;
        public readonly string SecurityProfile;
        public readonly ImmutableArray<Outputs.GetNsxtIpsecVpnTunnelSecurityProfileCustomizationResult> SecurityProfileCustomizations;
        public readonly string Status;
        public readonly string? Vdc;

        [OutputConstructor]
        private GetNsxtIpsecVpnTunnelResult(
            string authenticationMode,

            string caCertificateId,

            string certificateId,

            string description,

            string edgeGatewayId,

            bool enabled,

            string id,

            string ikeFailReason,

            string ikeServiceStatus,

            string localIpAddress,

            ImmutableArray<string> localNetworks,

            bool logging,

            string name,

            string? org,

            string preSharedKey,

            string remoteId,

            string remoteIpAddress,

            ImmutableArray<string> remoteNetworks,

            string securityProfile,

            ImmutableArray<Outputs.GetNsxtIpsecVpnTunnelSecurityProfileCustomizationResult> securityProfileCustomizations,

            string status,

            string? vdc)
        {
            AuthenticationMode = authenticationMode;
            CaCertificateId = caCertificateId;
            CertificateId = certificateId;
            Description = description;
            EdgeGatewayId = edgeGatewayId;
            Enabled = enabled;
            Id = id;
            IkeFailReason = ikeFailReason;
            IkeServiceStatus = ikeServiceStatus;
            LocalIpAddress = localIpAddress;
            LocalNetworks = localNetworks;
            Logging = logging;
            Name = name;
            Org = org;
            PreSharedKey = preSharedKey;
            RemoteId = remoteId;
            RemoteIpAddress = remoteIpAddress;
            RemoteNetworks = remoteNetworks;
            SecurityProfile = securityProfile;
            SecurityProfileCustomizations = securityProfileCustomizations;
            Status = status;
            Vdc = vdc;
        }
    }
}
