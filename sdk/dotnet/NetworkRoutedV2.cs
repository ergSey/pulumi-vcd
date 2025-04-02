// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/networkRoutedV2:NetworkRoutedV2")]
    public partial class NetworkRoutedV2 : global::Pulumi.CustomResource
    {
        /// <summary>
        /// An optional description of the network
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// First DNS server to use.
        /// </summary>
        [Output("dns1")]
        public Output<string?> Dns1 { get; private set; } = null!;

        /// <summary>
        /// Second DNS server to use.
        /// </summary>
        [Output("dns2")]
        public Output<string?> Dns2 { get; private set; } = null!;

        /// <summary>
        /// A FQDN for the virtual machines on this network
        /// </summary>
        [Output("dnsSuffix")]
        public Output<string?> DnsSuffix { get; private set; } = null!;

        /// <summary>
        /// Enables Dual-Stack mode so that one can configure one
        /// IPv4 and one IPv6 networks. **Note** In such case *IPv4* addresses must be used in `gateway`,
        /// `prefix_length` and `static_ip_pool` while *IPv6* addresses in `secondary_gateway`,
        /// `secondary_prefix_length` and `secondary_static_ip_pool` fields.
        /// </summary>
        [Output("dualStackEnabled")]
        public Output<bool?> DualStackEnabled { get; private set; } = null!;

        /// <summary>
        /// The ID of the Edge Gateway (NSX-V or NSX-T)
        /// </summary>
        [Output("edgeGatewayId")]
        public Output<string> EdgeGatewayId { get; private set; } = null!;

        /// <summary>
        /// The gateway for this network (e.g. 192.168.1.1, 2002:0:0:1234:abcd:ffff:c0a7:121)
        /// </summary>
        [Output("gateway")]
        public Output<string> Gateway { get; private set; } = null!;

        /// <summary>
        /// Set to `true` if network should allow guest VLAN tagging.
        /// Default `false`.
        /// </summary>
        [Output("guestVlanAllowed")]
        public Output<bool?> GuestVlanAllowed { get; private set; } = null!;

        /// <summary>
        /// An interface for the network. One of `internal` (default),
        /// `subinterface`, `distributed`, `non_distributed` (requires the Edge Gateway to support distributed
        /// networks). NSX-T supports only `internal` and `non_distributed` (*v3.14+*, requires Edge Gateway
        /// to have [non-distributed routing
        /// enabled](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_edgegateway#non_distributed_routing_enabled)).
        /// </summary>
        [Output("interfaceType")]
        public Output<string?> InterfaceType { get; private set; } = null!;

        /// <summary>
        /// Use `metadata_entry` instead. Key value map of metadata to assign to this network. **Not supported** if the owner edge gateway belongs to a VDC Group.
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, string>> Metadata { get; private set; } = null!;

        /// <summary>
        /// A set of metadata entries to assign. See Metadata section for details.
        /// </summary>
        [Output("metadataEntries")]
        public Output<ImmutableArray<Outputs.NetworkRoutedV2MetadataEntry>> MetadataEntries { get; private set; } = null!;

        /// <summary>
        /// A unique name for the network
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when
        /// connected as sysadmin working across different organisations
        /// </summary>
        [Output("org")]
        public Output<string?> Org { get; private set; } = null!;

        /// <summary>
        /// ID of VDC or VDC Group
        /// </summary>
        [Output("ownerId")]
        public Output<string> OwnerId { get; private set; } = null!;

        /// <summary>
        /// The prefix length for the new network (e.g. 24 for netmask 255.255.255.0).
        /// </summary>
        [Output("prefixLength")]
        public Output<int> PrefixLength { get; private set; } = null!;

        /// <summary>
        /// Enables route advertising for
        /// this network. **Note** This requires Edge Gateway to use IP Spaces and IP Space *must have* route
        /// advertisement
        /// enabled.
        /// </summary>
        [Output("routeAdvertisementEnabled")]
        public Output<bool?> RouteAdvertisementEnabled { get; private set; } = null!;

        /// <summary>
        /// IPv6 gateway *when Dual-Stack mode is enabled*
        /// </summary>
        [Output("secondaryGateway")]
        public Output<string?> SecondaryGateway { get; private set; } = null!;

        /// <summary>
        /// IPv6 prefix length *when Dual-Stack mode is
        /// enabled*
        /// </summary>
        [Output("secondaryPrefixLength")]
        public Output<string?> SecondaryPrefixLength { get; private set; } = null!;

        /// <summary>
        /// One or more IPv6 static
        /// pools *when Dual-Stack mode is enabled*
        /// 
        /// &gt; When using IPv6, VCD API will expand IP Addresses if they are specified using *double colon*
        /// notation and it will cause inconsistent plan. (e.g. `2002::1234:abcd:ffff:c0a6:121` will be
        /// converted to `2002:0:0:1234:abcd:ffff:c0a6:121`)
        /// 
        /// &lt;a id="ip-pools"&gt;&lt;/a&gt;
        /// </summary>
        [Output("secondaryStaticIpPools")]
        public Output<ImmutableArray<Outputs.NetworkRoutedV2SecondaryStaticIpPool>> SecondaryStaticIpPools { get; private set; } = null!;

        /// <summary>
        /// A range of IPs permitted to be used as static IPs for
        /// virtual machines; see IP Pools below for details.
        /// </summary>
        [Output("staticIpPools")]
        public Output<ImmutableArray<Outputs.NetworkRoutedV2StaticIpPool>> StaticIpPools { get; private set; } = null!;

        /// <summary>
        /// The name of VDC to use. *v3.6+* inherits parent VDC or VDC Group
        /// from `edge_gateway_id`)
        /// </summary>
        [Output("vdc")]
        public Output<string> Vdc { get; private set; } = null!;


        /// <summary>
        /// Create a NetworkRoutedV2 resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NetworkRoutedV2(string name, NetworkRoutedV2Args args, CustomResourceOptions? options = null)
            : base("vcd:index/networkRoutedV2:NetworkRoutedV2", name, args ?? new NetworkRoutedV2Args(), MakeResourceOptions(options, ""))
        {
        }

        private NetworkRoutedV2(string name, Input<string> id, NetworkRoutedV2State? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/networkRoutedV2:NetworkRoutedV2", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NetworkRoutedV2 resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NetworkRoutedV2 Get(string name, Input<string> id, NetworkRoutedV2State? state = null, CustomResourceOptions? options = null)
        {
            return new NetworkRoutedV2(name, id, state, options);
        }
    }

    public sealed class NetworkRoutedV2Args : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional description of the network
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// First DNS server to use.
        /// </summary>
        [Input("dns1")]
        public Input<string>? Dns1 { get; set; }

        /// <summary>
        /// Second DNS server to use.
        /// </summary>
        [Input("dns2")]
        public Input<string>? Dns2 { get; set; }

        /// <summary>
        /// A FQDN for the virtual machines on this network
        /// </summary>
        [Input("dnsSuffix")]
        public Input<string>? DnsSuffix { get; set; }

        /// <summary>
        /// Enables Dual-Stack mode so that one can configure one
        /// IPv4 and one IPv6 networks. **Note** In such case *IPv4* addresses must be used in `gateway`,
        /// `prefix_length` and `static_ip_pool` while *IPv6* addresses in `secondary_gateway`,
        /// `secondary_prefix_length` and `secondary_static_ip_pool` fields.
        /// </summary>
        [Input("dualStackEnabled")]
        public Input<bool>? DualStackEnabled { get; set; }

        /// <summary>
        /// The ID of the Edge Gateway (NSX-V or NSX-T)
        /// </summary>
        [Input("edgeGatewayId", required: true)]
        public Input<string> EdgeGatewayId { get; set; } = null!;

        /// <summary>
        /// The gateway for this network (e.g. 192.168.1.1, 2002:0:0:1234:abcd:ffff:c0a7:121)
        /// </summary>
        [Input("gateway", required: true)]
        public Input<string> Gateway { get; set; } = null!;

        /// <summary>
        /// Set to `true` if network should allow guest VLAN tagging.
        /// Default `false`.
        /// </summary>
        [Input("guestVlanAllowed")]
        public Input<bool>? GuestVlanAllowed { get; set; }

        /// <summary>
        /// An interface for the network. One of `internal` (default),
        /// `subinterface`, `distributed`, `non_distributed` (requires the Edge Gateway to support distributed
        /// networks). NSX-T supports only `internal` and `non_distributed` (*v3.14+*, requires Edge Gateway
        /// to have [non-distributed routing
        /// enabled](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_edgegateway#non_distributed_routing_enabled)).
        /// </summary>
        [Input("interfaceType")]
        public Input<string>? InterfaceType { get; set; }

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// Use `metadata_entry` instead. Key value map of metadata to assign to this network. **Not supported** if the owner edge gateway belongs to a VDC Group.
        /// </summary>
        [Obsolete(@"Use metadata_entry instead")]
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        [Input("metadataEntries")]
        private InputList<Inputs.NetworkRoutedV2MetadataEntryArgs>? _metadataEntries;

        /// <summary>
        /// A set of metadata entries to assign. See Metadata section for details.
        /// </summary>
        public InputList<Inputs.NetworkRoutedV2MetadataEntryArgs> MetadataEntries
        {
            get => _metadataEntries ?? (_metadataEntries = new InputList<Inputs.NetworkRoutedV2MetadataEntryArgs>());
            set => _metadataEntries = value;
        }

        /// <summary>
        /// A unique name for the network
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when
        /// connected as sysadmin working across different organisations
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// The prefix length for the new network (e.g. 24 for netmask 255.255.255.0).
        /// </summary>
        [Input("prefixLength", required: true)]
        public Input<int> PrefixLength { get; set; } = null!;

        /// <summary>
        /// Enables route advertising for
        /// this network. **Note** This requires Edge Gateway to use IP Spaces and IP Space *must have* route
        /// advertisement
        /// enabled.
        /// </summary>
        [Input("routeAdvertisementEnabled")]
        public Input<bool>? RouteAdvertisementEnabled { get; set; }

        /// <summary>
        /// IPv6 gateway *when Dual-Stack mode is enabled*
        /// </summary>
        [Input("secondaryGateway")]
        public Input<string>? SecondaryGateway { get; set; }

        /// <summary>
        /// IPv6 prefix length *when Dual-Stack mode is
        /// enabled*
        /// </summary>
        [Input("secondaryPrefixLength")]
        public Input<string>? SecondaryPrefixLength { get; set; }

        [Input("secondaryStaticIpPools")]
        private InputList<Inputs.NetworkRoutedV2SecondaryStaticIpPoolArgs>? _secondaryStaticIpPools;

        /// <summary>
        /// One or more IPv6 static
        /// pools *when Dual-Stack mode is enabled*
        /// 
        /// &gt; When using IPv6, VCD API will expand IP Addresses if they are specified using *double colon*
        /// notation and it will cause inconsistent plan. (e.g. `2002::1234:abcd:ffff:c0a6:121` will be
        /// converted to `2002:0:0:1234:abcd:ffff:c0a6:121`)
        /// 
        /// &lt;a id="ip-pools"&gt;&lt;/a&gt;
        /// </summary>
        public InputList<Inputs.NetworkRoutedV2SecondaryStaticIpPoolArgs> SecondaryStaticIpPools
        {
            get => _secondaryStaticIpPools ?? (_secondaryStaticIpPools = new InputList<Inputs.NetworkRoutedV2SecondaryStaticIpPoolArgs>());
            set => _secondaryStaticIpPools = value;
        }

        [Input("staticIpPools")]
        private InputList<Inputs.NetworkRoutedV2StaticIpPoolArgs>? _staticIpPools;

        /// <summary>
        /// A range of IPs permitted to be used as static IPs for
        /// virtual machines; see IP Pools below for details.
        /// </summary>
        public InputList<Inputs.NetworkRoutedV2StaticIpPoolArgs> StaticIpPools
        {
            get => _staticIpPools ?? (_staticIpPools = new InputList<Inputs.NetworkRoutedV2StaticIpPoolArgs>());
            set => _staticIpPools = value;
        }

        /// <summary>
        /// The name of VDC to use. *v3.6+* inherits parent VDC or VDC Group
        /// from `edge_gateway_id`)
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public NetworkRoutedV2Args()
        {
        }
        public static new NetworkRoutedV2Args Empty => new NetworkRoutedV2Args();
    }

    public sealed class NetworkRoutedV2State : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional description of the network
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// First DNS server to use.
        /// </summary>
        [Input("dns1")]
        public Input<string>? Dns1 { get; set; }

        /// <summary>
        /// Second DNS server to use.
        /// </summary>
        [Input("dns2")]
        public Input<string>? Dns2 { get; set; }

        /// <summary>
        /// A FQDN for the virtual machines on this network
        /// </summary>
        [Input("dnsSuffix")]
        public Input<string>? DnsSuffix { get; set; }

        /// <summary>
        /// Enables Dual-Stack mode so that one can configure one
        /// IPv4 and one IPv6 networks. **Note** In such case *IPv4* addresses must be used in `gateway`,
        /// `prefix_length` and `static_ip_pool` while *IPv6* addresses in `secondary_gateway`,
        /// `secondary_prefix_length` and `secondary_static_ip_pool` fields.
        /// </summary>
        [Input("dualStackEnabled")]
        public Input<bool>? DualStackEnabled { get; set; }

        /// <summary>
        /// The ID of the Edge Gateway (NSX-V or NSX-T)
        /// </summary>
        [Input("edgeGatewayId")]
        public Input<string>? EdgeGatewayId { get; set; }

        /// <summary>
        /// The gateway for this network (e.g. 192.168.1.1, 2002:0:0:1234:abcd:ffff:c0a7:121)
        /// </summary>
        [Input("gateway")]
        public Input<string>? Gateway { get; set; }

        /// <summary>
        /// Set to `true` if network should allow guest VLAN tagging.
        /// Default `false`.
        /// </summary>
        [Input("guestVlanAllowed")]
        public Input<bool>? GuestVlanAllowed { get; set; }

        /// <summary>
        /// An interface for the network. One of `internal` (default),
        /// `subinterface`, `distributed`, `non_distributed` (requires the Edge Gateway to support distributed
        /// networks). NSX-T supports only `internal` and `non_distributed` (*v3.14+*, requires Edge Gateway
        /// to have [non-distributed routing
        /// enabled](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_edgegateway#non_distributed_routing_enabled)).
        /// </summary>
        [Input("interfaceType")]
        public Input<string>? InterfaceType { get; set; }

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// Use `metadata_entry` instead. Key value map of metadata to assign to this network. **Not supported** if the owner edge gateway belongs to a VDC Group.
        /// </summary>
        [Obsolete(@"Use metadata_entry instead")]
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        [Input("metadataEntries")]
        private InputList<Inputs.NetworkRoutedV2MetadataEntryGetArgs>? _metadataEntries;

        /// <summary>
        /// A set of metadata entries to assign. See Metadata section for details.
        /// </summary>
        public InputList<Inputs.NetworkRoutedV2MetadataEntryGetArgs> MetadataEntries
        {
            get => _metadataEntries ?? (_metadataEntries = new InputList<Inputs.NetworkRoutedV2MetadataEntryGetArgs>());
            set => _metadataEntries = value;
        }

        /// <summary>
        /// A unique name for the network
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when
        /// connected as sysadmin working across different organisations
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// ID of VDC or VDC Group
        /// </summary>
        [Input("ownerId")]
        public Input<string>? OwnerId { get; set; }

        /// <summary>
        /// The prefix length for the new network (e.g. 24 for netmask 255.255.255.0).
        /// </summary>
        [Input("prefixLength")]
        public Input<int>? PrefixLength { get; set; }

        /// <summary>
        /// Enables route advertising for
        /// this network. **Note** This requires Edge Gateway to use IP Spaces and IP Space *must have* route
        /// advertisement
        /// enabled.
        /// </summary>
        [Input("routeAdvertisementEnabled")]
        public Input<bool>? RouteAdvertisementEnabled { get; set; }

        /// <summary>
        /// IPv6 gateway *when Dual-Stack mode is enabled*
        /// </summary>
        [Input("secondaryGateway")]
        public Input<string>? SecondaryGateway { get; set; }

        /// <summary>
        /// IPv6 prefix length *when Dual-Stack mode is
        /// enabled*
        /// </summary>
        [Input("secondaryPrefixLength")]
        public Input<string>? SecondaryPrefixLength { get; set; }

        [Input("secondaryStaticIpPools")]
        private InputList<Inputs.NetworkRoutedV2SecondaryStaticIpPoolGetArgs>? _secondaryStaticIpPools;

        /// <summary>
        /// One or more IPv6 static
        /// pools *when Dual-Stack mode is enabled*
        /// 
        /// &gt; When using IPv6, VCD API will expand IP Addresses if they are specified using *double colon*
        /// notation and it will cause inconsistent plan. (e.g. `2002::1234:abcd:ffff:c0a6:121` will be
        /// converted to `2002:0:0:1234:abcd:ffff:c0a6:121`)
        /// 
        /// &lt;a id="ip-pools"&gt;&lt;/a&gt;
        /// </summary>
        public InputList<Inputs.NetworkRoutedV2SecondaryStaticIpPoolGetArgs> SecondaryStaticIpPools
        {
            get => _secondaryStaticIpPools ?? (_secondaryStaticIpPools = new InputList<Inputs.NetworkRoutedV2SecondaryStaticIpPoolGetArgs>());
            set => _secondaryStaticIpPools = value;
        }

        [Input("staticIpPools")]
        private InputList<Inputs.NetworkRoutedV2StaticIpPoolGetArgs>? _staticIpPools;

        /// <summary>
        /// A range of IPs permitted to be used as static IPs for
        /// virtual machines; see IP Pools below for details.
        /// </summary>
        public InputList<Inputs.NetworkRoutedV2StaticIpPoolGetArgs> StaticIpPools
        {
            get => _staticIpPools ?? (_staticIpPools = new InputList<Inputs.NetworkRoutedV2StaticIpPoolGetArgs>());
            set => _staticIpPools = value;
        }

        /// <summary>
        /// The name of VDC to use. *v3.6+* inherits parent VDC or VDC Group
        /// from `edge_gateway_id`)
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public NetworkRoutedV2State()
        {
        }
        public static new NetworkRoutedV2State Empty => new NetworkRoutedV2State();
    }
}
