// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/ipSpace:IpSpace")]
    public partial class IpSpace : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Defines whether
        /// default firewall rule creation should be enabled
        /// </summary>
        [Output("defaultFirewallRuleCreationEnabled")]
        public Output<bool?> DefaultFirewallRuleCreationEnabled { get; private set; } = null!;

        /// <summary>
        /// Defines whether NO SNAT
        /// rule creation should be enabled
        /// </summary>
        [Output("defaultNoSnatRuleCreationEnabled")]
        public Output<bool?> DefaultNoSnatRuleCreationEnabled { get; private set; } = null!;

        /// <summary>
        /// Defines whether SNAT rule
        /// creation should be enabled
        /// 
        /// &lt;a id="ipspace-ip-range"&gt;&lt;/a&gt;
        /// </summary>
        [Output("defaultSnatRuleCreationEnabled")]
        public Output<bool?> DefaultSnatRuleCreationEnabled { get; private set; } = null!;

        /// <summary>
        /// Description of IP Space
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The external scope defines the total span of IP addresses to which the IP
        /// space has access, for example the internet or a WAN.
        /// </summary>
        [Output("externalScope")]
        public Output<string?> ExternalScope { get; private set; } = null!;

        /// <summary>
        /// The internal scope of an IP space is a list of CIDR notations that
        /// defines the exact span of IP addresses in which all ranges and blocks must be contained in.
        /// </summary>
        [Output("internalScopes")]
        public Output<ImmutableArray<string>> InternalScopes { get; private set; } = null!;

        /// <summary>
        /// One or more IP prefixes (blocks) ip_prefix
        /// </summary>
        [Output("ipPrefixes")]
        public Output<ImmutableArray<Outputs.IpSpaceIpPrefix>> IpPrefixes { get; private set; } = null!;

        /// <summary>
        /// If you entered at least one IP Range
        /// (ip_range), enter a number of floating IP addresses to allocate individually.
        /// `-1` is unlimited, while `0` means that no IPs can be allocated.
        /// </summary>
        [Output("ipRangeQuota")]
        public Output<string> IpRangeQuota { get; private set; } = null!;

        /// <summary>
        /// One or more ip_range for floating IP address
        /// allocation. (Floating IP addresses are just IP addresses taken from the defined range)
        /// </summary>
        [Output("ipRanges")]
        public Output<ImmutableArray<Outputs.IpSpaceIpRange>> IpRanges { get; private set; } = null!;

        /// <summary>
        /// A name for IP Space
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Required for `PRIVATE` type
        /// </summary>
        [Output("orgId")]
        public Output<string?> OrgId { get; private set; } = null!;

        /// <summary>
        /// Toggle on the route advertisement option to
        /// enable advertising networks with IP prefixes from this IP space (default `false`)
        /// </summary>
        [Output("routeAdvertisementEnabled")]
        public Output<bool?> RouteAdvertisementEnabled { get; private set; } = null!;

        /// <summary>
        /// One of `PUBLIC`, `SHARED_SERVICES`, `PRIVATE`
        /// * `PUBLIC` - A public IP space is *used by multiple organizations* and is *controlled by the service
        /// provider* through a quota-based system.
        /// * `SHARED_SERVICES` - An IP space for services and management networks that are required in the
        /// tenant space, but as a service provider, you don't want to expose it to organizations in your
        /// environment. The main difference from `PUBLIC` network is that IPs cannot be allocated by tenants.
        /// * `PRIVATE` - Private IP spaces are dedicated to a single tenant - a private IP space is used by
        /// only one organization that is specified during the space creation. For this organization, IP
        /// consumption is unlimited.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a IpSpace resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IpSpace(string name, IpSpaceArgs args, CustomResourceOptions? options = null)
            : base("vcd:index/ipSpace:IpSpace", name, args ?? new IpSpaceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IpSpace(string name, Input<string> id, IpSpaceState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/ipSpace:IpSpace", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IpSpace resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IpSpace Get(string name, Input<string> id, IpSpaceState? state = null, CustomResourceOptions? options = null)
        {
            return new IpSpace(name, id, state, options);
        }
    }

    public sealed class IpSpaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defines whether
        /// default firewall rule creation should be enabled
        /// </summary>
        [Input("defaultFirewallRuleCreationEnabled")]
        public Input<bool>? DefaultFirewallRuleCreationEnabled { get; set; }

        /// <summary>
        /// Defines whether NO SNAT
        /// rule creation should be enabled
        /// </summary>
        [Input("defaultNoSnatRuleCreationEnabled")]
        public Input<bool>? DefaultNoSnatRuleCreationEnabled { get; set; }

        /// <summary>
        /// Defines whether SNAT rule
        /// creation should be enabled
        /// 
        /// &lt;a id="ipspace-ip-range"&gt;&lt;/a&gt;
        /// </summary>
        [Input("defaultSnatRuleCreationEnabled")]
        public Input<bool>? DefaultSnatRuleCreationEnabled { get; set; }

        /// <summary>
        /// Description of IP Space
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The external scope defines the total span of IP addresses to which the IP
        /// space has access, for example the internet or a WAN.
        /// </summary>
        [Input("externalScope")]
        public Input<string>? ExternalScope { get; set; }

        [Input("internalScopes", required: true)]
        private InputList<string>? _internalScopes;

        /// <summary>
        /// The internal scope of an IP space is a list of CIDR notations that
        /// defines the exact span of IP addresses in which all ranges and blocks must be contained in.
        /// </summary>
        public InputList<string> InternalScopes
        {
            get => _internalScopes ?? (_internalScopes = new InputList<string>());
            set => _internalScopes = value;
        }

        [Input("ipPrefixes")]
        private InputList<Inputs.IpSpaceIpPrefixArgs>? _ipPrefixes;

        /// <summary>
        /// One or more IP prefixes (blocks) ip_prefix
        /// </summary>
        public InputList<Inputs.IpSpaceIpPrefixArgs> IpPrefixes
        {
            get => _ipPrefixes ?? (_ipPrefixes = new InputList<Inputs.IpSpaceIpPrefixArgs>());
            set => _ipPrefixes = value;
        }

        /// <summary>
        /// If you entered at least one IP Range
        /// (ip_range), enter a number of floating IP addresses to allocate individually.
        /// `-1` is unlimited, while `0` means that no IPs can be allocated.
        /// </summary>
        [Input("ipRangeQuota")]
        public Input<string>? IpRangeQuota { get; set; }

        [Input("ipRanges")]
        private InputList<Inputs.IpSpaceIpRangeArgs>? _ipRanges;

        /// <summary>
        /// One or more ip_range for floating IP address
        /// allocation. (Floating IP addresses are just IP addresses taken from the defined range)
        /// </summary>
        public InputList<Inputs.IpSpaceIpRangeArgs> IpRanges
        {
            get => _ipRanges ?? (_ipRanges = new InputList<Inputs.IpSpaceIpRangeArgs>());
            set => _ipRanges = value;
        }

        /// <summary>
        /// A name for IP Space
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Required for `PRIVATE` type
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// Toggle on the route advertisement option to
        /// enable advertising networks with IP prefixes from this IP space (default `false`)
        /// </summary>
        [Input("routeAdvertisementEnabled")]
        public Input<bool>? RouteAdvertisementEnabled { get; set; }

        /// <summary>
        /// One of `PUBLIC`, `SHARED_SERVICES`, `PRIVATE`
        /// * `PUBLIC` - A public IP space is *used by multiple organizations* and is *controlled by the service
        /// provider* through a quota-based system.
        /// * `SHARED_SERVICES` - An IP space for services and management networks that are required in the
        /// tenant space, but as a service provider, you don't want to expose it to organizations in your
        /// environment. The main difference from `PUBLIC` network is that IPs cannot be allocated by tenants.
        /// * `PRIVATE` - Private IP spaces are dedicated to a single tenant - a private IP space is used by
        /// only one organization that is specified during the space creation. For this organization, IP
        /// consumption is unlimited.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public IpSpaceArgs()
        {
        }
        public static new IpSpaceArgs Empty => new IpSpaceArgs();
    }

    public sealed class IpSpaceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Defines whether
        /// default firewall rule creation should be enabled
        /// </summary>
        [Input("defaultFirewallRuleCreationEnabled")]
        public Input<bool>? DefaultFirewallRuleCreationEnabled { get; set; }

        /// <summary>
        /// Defines whether NO SNAT
        /// rule creation should be enabled
        /// </summary>
        [Input("defaultNoSnatRuleCreationEnabled")]
        public Input<bool>? DefaultNoSnatRuleCreationEnabled { get; set; }

        /// <summary>
        /// Defines whether SNAT rule
        /// creation should be enabled
        /// 
        /// &lt;a id="ipspace-ip-range"&gt;&lt;/a&gt;
        /// </summary>
        [Input("defaultSnatRuleCreationEnabled")]
        public Input<bool>? DefaultSnatRuleCreationEnabled { get; set; }

        /// <summary>
        /// Description of IP Space
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The external scope defines the total span of IP addresses to which the IP
        /// space has access, for example the internet or a WAN.
        /// </summary>
        [Input("externalScope")]
        public Input<string>? ExternalScope { get; set; }

        [Input("internalScopes")]
        private InputList<string>? _internalScopes;

        /// <summary>
        /// The internal scope of an IP space is a list of CIDR notations that
        /// defines the exact span of IP addresses in which all ranges and blocks must be contained in.
        /// </summary>
        public InputList<string> InternalScopes
        {
            get => _internalScopes ?? (_internalScopes = new InputList<string>());
            set => _internalScopes = value;
        }

        [Input("ipPrefixes")]
        private InputList<Inputs.IpSpaceIpPrefixGetArgs>? _ipPrefixes;

        /// <summary>
        /// One or more IP prefixes (blocks) ip_prefix
        /// </summary>
        public InputList<Inputs.IpSpaceIpPrefixGetArgs> IpPrefixes
        {
            get => _ipPrefixes ?? (_ipPrefixes = new InputList<Inputs.IpSpaceIpPrefixGetArgs>());
            set => _ipPrefixes = value;
        }

        /// <summary>
        /// If you entered at least one IP Range
        /// (ip_range), enter a number of floating IP addresses to allocate individually.
        /// `-1` is unlimited, while `0` means that no IPs can be allocated.
        /// </summary>
        [Input("ipRangeQuota")]
        public Input<string>? IpRangeQuota { get; set; }

        [Input("ipRanges")]
        private InputList<Inputs.IpSpaceIpRangeGetArgs>? _ipRanges;

        /// <summary>
        /// One or more ip_range for floating IP address
        /// allocation. (Floating IP addresses are just IP addresses taken from the defined range)
        /// </summary>
        public InputList<Inputs.IpSpaceIpRangeGetArgs> IpRanges
        {
            get => _ipRanges ?? (_ipRanges = new InputList<Inputs.IpSpaceIpRangeGetArgs>());
            set => _ipRanges = value;
        }

        /// <summary>
        /// A name for IP Space
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Required for `PRIVATE` type
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// Toggle on the route advertisement option to
        /// enable advertising networks with IP prefixes from this IP space (default `false`)
        /// </summary>
        [Input("routeAdvertisementEnabled")]
        public Input<bool>? RouteAdvertisementEnabled { get; set; }

        /// <summary>
        /// One of `PUBLIC`, `SHARED_SERVICES`, `PRIVATE`
        /// * `PUBLIC` - A public IP space is *used by multiple organizations* and is *controlled by the service
        /// provider* through a quota-based system.
        /// * `SHARED_SERVICES` - An IP space for services and management networks that are required in the
        /// tenant space, but as a service provider, you don't want to expose it to organizations in your
        /// environment. The main difference from `PUBLIC` network is that IPs cannot be allocated by tenants.
        /// * `PRIVATE` - Private IP spaces are dedicated to a single tenant - a private IP space is used by
        /// only one organization that is specified during the space creation. For this organization, IP
        /// consumption is unlimited.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public IpSpaceState()
        {
        }
        public static new IpSpaceState Empty => new IpSpaceState();
    }
}
