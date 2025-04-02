// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetOrgOidc
    {
        /// <summary>
        /// Provides a data source to read the OpenID Connect (OIDC) configuration of an Organization in VMware Cloud Director.
        /// 
        /// Supported in provider *v3.13+*.
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
        ///     var myOrg = Vcd.GetOrg.Invoke(new()
        ///     {
        ///         Name = "my-org",
        ///     });
        /// 
        ///     var oidcSettings = Vcd.GetOrgOidc.Invoke(new()
        ///     {
        ///         OrgId = myOrg.Apply(getOrgResult =&gt; getOrgResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetOrgOidcResult> InvokeAsync(GetOrgOidcArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetOrgOidcResult>("vcd:index/getOrgOidc:getOrgOidc", args ?? new GetOrgOidcArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a data source to read the OpenID Connect (OIDC) configuration of an Organization in VMware Cloud Director.
        /// 
        /// Supported in provider *v3.13+*.
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
        ///     var myOrg = Vcd.GetOrg.Invoke(new()
        ///     {
        ///         Name = "my-org",
        ///     });
        /// 
        ///     var oidcSettings = Vcd.GetOrgOidc.Invoke(new()
        ///     {
        ///         OrgId = myOrg.Apply(getOrgResult =&gt; getOrgResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetOrgOidcResult> Invoke(GetOrgOidcInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetOrgOidcResult>("vcd:index/getOrgOidc:getOrgOidc", args ?? new GetOrgOidcInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a data source to read the OpenID Connect (OIDC) configuration of an Organization in VMware Cloud Director.
        /// 
        /// Supported in provider *v3.13+*.
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
        ///     var myOrg = Vcd.GetOrg.Invoke(new()
        ///     {
        ///         Name = "my-org",
        ///     });
        /// 
        ///     var oidcSettings = Vcd.GetOrgOidc.Invoke(new()
        ///     {
        ///         OrgId = myOrg.Apply(getOrgResult =&gt; getOrgResult.Id),
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetOrgOidcResult> Invoke(GetOrgOidcInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetOrgOidcResult>("vcd:index/getOrgOidc:getOrgOidc", args ?? new GetOrgOidcInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetOrgOidcArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the organization containing the OIDC settings
        /// </summary>
        [Input("orgId", required: true)]
        public string OrgId { get; set; } = null!;

        public GetOrgOidcArgs()
        {
        }
        public static new GetOrgOidcArgs Empty => new GetOrgOidcArgs();
    }

    public sealed class GetOrgOidcInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the organization containing the OIDC settings
        /// </summary>
        [Input("orgId", required: true)]
        public Input<string> OrgId { get; set; } = null!;

        public GetOrgOidcInvokeArgs()
        {
        }
        public static new GetOrgOidcInvokeArgs Empty => new GetOrgOidcInvokeArgs();
    }


    [OutputType]
    public sealed class GetOrgOidcResult
    {
        public readonly string AccessTokenEndpoint;
        public readonly ImmutableArray<Outputs.GetOrgOidcClaimsMappingResult> ClaimsMappings;
        public readonly string ClientId;
        public readonly string ClientSecret;
        public readonly bool Enabled;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string IssuerId;
        public readonly int KeyExpireDurationHours;
        public readonly string KeyRefreshEndpoint;
        public readonly int KeyRefreshPeriodHours;
        public readonly string KeyRefreshStrategy;
        public readonly ImmutableArray<Outputs.GetOrgOidcKeyResult> Keys;
        public readonly int MaxClockSkewSeconds;
        public readonly string OrgId;
        public readonly bool PreferIdToken;
        public readonly string RedirectUri;
        public readonly ImmutableArray<string> Scopes;
        public readonly string UiButtonLabel;
        public readonly string UserAuthorizationEndpoint;
        public readonly string UserinfoEndpoint;
        public readonly string WellknownEndpoint;

        [OutputConstructor]
        private GetOrgOidcResult(
            string accessTokenEndpoint,

            ImmutableArray<Outputs.GetOrgOidcClaimsMappingResult> claimsMappings,

            string clientId,

            string clientSecret,

            bool enabled,

            string id,

            string issuerId,

            int keyExpireDurationHours,

            string keyRefreshEndpoint,

            int keyRefreshPeriodHours,

            string keyRefreshStrategy,

            ImmutableArray<Outputs.GetOrgOidcKeyResult> keys,

            int maxClockSkewSeconds,

            string orgId,

            bool preferIdToken,

            string redirectUri,

            ImmutableArray<string> scopes,

            string uiButtonLabel,

            string userAuthorizationEndpoint,

            string userinfoEndpoint,

            string wellknownEndpoint)
        {
            AccessTokenEndpoint = accessTokenEndpoint;
            ClaimsMappings = claimsMappings;
            ClientId = clientId;
            ClientSecret = clientSecret;
            Enabled = enabled;
            Id = id;
            IssuerId = issuerId;
            KeyExpireDurationHours = keyExpireDurationHours;
            KeyRefreshEndpoint = keyRefreshEndpoint;
            KeyRefreshPeriodHours = keyRefreshPeriodHours;
            KeyRefreshStrategy = keyRefreshStrategy;
            Keys = keys;
            MaxClockSkewSeconds = maxClockSkewSeconds;
            OrgId = orgId;
            PreferIdToken = preferIdToken;
            RedirectUri = redirectUri;
            Scopes = scopes;
            UiButtonLabel = uiButtonLabel;
            UserAuthorizationEndpoint = userAuthorizationEndpoint;
            UserinfoEndpoint = userinfoEndpoint;
            WellknownEndpoint = wellknownEndpoint;
        }
    }
}
