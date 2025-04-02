// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/nsxtNetworkImported:NsxtNetworkImported")]
    public partial class NsxtNetworkImported : global::Pulumi.CustomResource
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
        /// ID of Distributed Virtual Port Group used by this network
        /// </summary>
        [Output("dvpgId")]
        public Output<string> DvpgId { get; private set; } = null!;

        /// <summary>
        /// Unique name of an existing Distributed Virtual Port Group (DVPG). 
        /// **Note** it will never be refreshed because API does not allow reading this name after it is
        /// consumed. Instead ID will be stored in `dvpg_id` attribute.
        /// 
        /// &gt; One of `nsxt_logical_switch_name` or `dvpg_name` must be provided.
        /// </summary>
        [Output("dvpgName")]
        public Output<string?> DvpgName { get; private set; } = null!;

        /// <summary>
        /// The gateway for this network (e.g. 192.168.1.1, 2002:0:0:1234:abcd:ffff:c0a7:121)
        /// </summary>
        [Output("gateway")]
        public Output<string> Gateway { get; private set; } = null!;

        /// <summary>
        /// A unique name for the network
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// ID of NSX-T logical switch used by this network
        /// </summary>
        [Output("nsxtLogicalSwitchId")]
        public Output<string> NsxtLogicalSwitchId { get; private set; } = null!;

        /// <summary>
        /// Unique name of an existing NSX-T segment. 
        /// **Note** it will never be refreshed because API does not allow reading this name after it is
        /// consumed. Instead ID will be stored in `nsxt_logical_switch_id` attribute.
        /// 
        /// This resource **will fail** if multiple segments with the same name are available. One can rename
        /// them in NSX-T manager to make them unique.
        /// </summary>
        [Output("nsxtLogicalSwitchName")]
        public Output<string?> NsxtLogicalSwitchName { get; private set; } = null!;

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when
        /// connected as sysadmin working across different organisations
        /// </summary>
        [Output("org")]
        public Output<string?> Org { get; private set; } = null!;

        /// <summary>
        /// VDC or VDC Group ID. Always takes precedence over `vdc` fields (in resource
        /// and inherited from provider configuration)
        /// </summary>
        [Output("ownerId")]
        public Output<string> OwnerId { get; private set; } = null!;

        /// <summary>
        /// The prefix length for the new network (e.g. 24 for netmask 255.255.255.0).
        /// </summary>
        [Output("prefixLength")]
        public Output<int> PrefixLength { get; private set; } = null!;

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
        public Output<ImmutableArray<Outputs.NsxtNetworkImportedSecondaryStaticIpPool>> SecondaryStaticIpPools { get; private set; } = null!;

        /// <summary>
        /// A range of IPs permitted to be used as static IPs for
        /// virtual machines; see IP Pools below for details.
        /// </summary>
        [Output("staticIpPools")]
        public Output<ImmutableArray<Outputs.NsxtNetworkImportedStaticIpPool>> StaticIpPools { get; private set; } = null!;

        /// <summary>
        /// The name of VDC to use. **Deprecated**  in favor of new field
        /// `owner_id` which supports VDC and VDC Group IDs.
        /// </summary>
        [Output("vdc")]
        public Output<string> Vdc { get; private set; } = null!;


        /// <summary>
        /// Create a NsxtNetworkImported resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NsxtNetworkImported(string name, NsxtNetworkImportedArgs args, CustomResourceOptions? options = null)
            : base("vcd:index/nsxtNetworkImported:NsxtNetworkImported", name, args ?? new NsxtNetworkImportedArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NsxtNetworkImported(string name, Input<string> id, NsxtNetworkImportedState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/nsxtNetworkImported:NsxtNetworkImported", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NsxtNetworkImported resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NsxtNetworkImported Get(string name, Input<string> id, NsxtNetworkImportedState? state = null, CustomResourceOptions? options = null)
        {
            return new NsxtNetworkImported(name, id, state, options);
        }
    }

    public sealed class NsxtNetworkImportedArgs : global::Pulumi.ResourceArgs
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
        /// Unique name of an existing Distributed Virtual Port Group (DVPG). 
        /// **Note** it will never be refreshed because API does not allow reading this name after it is
        /// consumed. Instead ID will be stored in `dvpg_id` attribute.
        /// 
        /// &gt; One of `nsxt_logical_switch_name` or `dvpg_name` must be provided.
        /// </summary>
        [Input("dvpgName")]
        public Input<string>? DvpgName { get; set; }

        /// <summary>
        /// The gateway for this network (e.g. 192.168.1.1, 2002:0:0:1234:abcd:ffff:c0a7:121)
        /// </summary>
        [Input("gateway", required: true)]
        public Input<string> Gateway { get; set; } = null!;

        /// <summary>
        /// A unique name for the network
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Unique name of an existing NSX-T segment. 
        /// **Note** it will never be refreshed because API does not allow reading this name after it is
        /// consumed. Instead ID will be stored in `nsxt_logical_switch_id` attribute.
        /// 
        /// This resource **will fail** if multiple segments with the same name are available. One can rename
        /// them in NSX-T manager to make them unique.
        /// </summary>
        [Input("nsxtLogicalSwitchName")]
        public Input<string>? NsxtLogicalSwitchName { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when
        /// connected as sysadmin working across different organisations
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// VDC or VDC Group ID. Always takes precedence over `vdc` fields (in resource
        /// and inherited from provider configuration)
        /// </summary>
        [Input("ownerId")]
        public Input<string>? OwnerId { get; set; }

        /// <summary>
        /// The prefix length for the new network (e.g. 24 for netmask 255.255.255.0).
        /// </summary>
        [Input("prefixLength", required: true)]
        public Input<int> PrefixLength { get; set; } = null!;

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
        private InputList<Inputs.NsxtNetworkImportedSecondaryStaticIpPoolArgs>? _secondaryStaticIpPools;

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
        public InputList<Inputs.NsxtNetworkImportedSecondaryStaticIpPoolArgs> SecondaryStaticIpPools
        {
            get => _secondaryStaticIpPools ?? (_secondaryStaticIpPools = new InputList<Inputs.NsxtNetworkImportedSecondaryStaticIpPoolArgs>());
            set => _secondaryStaticIpPools = value;
        }

        [Input("staticIpPools")]
        private InputList<Inputs.NsxtNetworkImportedStaticIpPoolArgs>? _staticIpPools;

        /// <summary>
        /// A range of IPs permitted to be used as static IPs for
        /// virtual machines; see IP Pools below for details.
        /// </summary>
        public InputList<Inputs.NsxtNetworkImportedStaticIpPoolArgs> StaticIpPools
        {
            get => _staticIpPools ?? (_staticIpPools = new InputList<Inputs.NsxtNetworkImportedStaticIpPoolArgs>());
            set => _staticIpPools = value;
        }

        /// <summary>
        /// The name of VDC to use. **Deprecated**  in favor of new field
        /// `owner_id` which supports VDC and VDC Group IDs.
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public NsxtNetworkImportedArgs()
        {
        }
        public static new NsxtNetworkImportedArgs Empty => new NsxtNetworkImportedArgs();
    }

    public sealed class NsxtNetworkImportedState : global::Pulumi.ResourceArgs
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
        /// ID of Distributed Virtual Port Group used by this network
        /// </summary>
        [Input("dvpgId")]
        public Input<string>? DvpgId { get; set; }

        /// <summary>
        /// Unique name of an existing Distributed Virtual Port Group (DVPG). 
        /// **Note** it will never be refreshed because API does not allow reading this name after it is
        /// consumed. Instead ID will be stored in `dvpg_id` attribute.
        /// 
        /// &gt; One of `nsxt_logical_switch_name` or `dvpg_name` must be provided.
        /// </summary>
        [Input("dvpgName")]
        public Input<string>? DvpgName { get; set; }

        /// <summary>
        /// The gateway for this network (e.g. 192.168.1.1, 2002:0:0:1234:abcd:ffff:c0a7:121)
        /// </summary>
        [Input("gateway")]
        public Input<string>? Gateway { get; set; }

        /// <summary>
        /// A unique name for the network
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// ID of NSX-T logical switch used by this network
        /// </summary>
        [Input("nsxtLogicalSwitchId")]
        public Input<string>? NsxtLogicalSwitchId { get; set; }

        /// <summary>
        /// Unique name of an existing NSX-T segment. 
        /// **Note** it will never be refreshed because API does not allow reading this name after it is
        /// consumed. Instead ID will be stored in `nsxt_logical_switch_id` attribute.
        /// 
        /// This resource **will fail** if multiple segments with the same name are available. One can rename
        /// them in NSX-T manager to make them unique.
        /// </summary>
        [Input("nsxtLogicalSwitchName")]
        public Input<string>? NsxtLogicalSwitchName { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when
        /// connected as sysadmin working across different organisations
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// VDC or VDC Group ID. Always takes precedence over `vdc` fields (in resource
        /// and inherited from provider configuration)
        /// </summary>
        [Input("ownerId")]
        public Input<string>? OwnerId { get; set; }

        /// <summary>
        /// The prefix length for the new network (e.g. 24 for netmask 255.255.255.0).
        /// </summary>
        [Input("prefixLength")]
        public Input<int>? PrefixLength { get; set; }

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
        private InputList<Inputs.NsxtNetworkImportedSecondaryStaticIpPoolGetArgs>? _secondaryStaticIpPools;

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
        public InputList<Inputs.NsxtNetworkImportedSecondaryStaticIpPoolGetArgs> SecondaryStaticIpPools
        {
            get => _secondaryStaticIpPools ?? (_secondaryStaticIpPools = new InputList<Inputs.NsxtNetworkImportedSecondaryStaticIpPoolGetArgs>());
            set => _secondaryStaticIpPools = value;
        }

        [Input("staticIpPools")]
        private InputList<Inputs.NsxtNetworkImportedStaticIpPoolGetArgs>? _staticIpPools;

        /// <summary>
        /// A range of IPs permitted to be used as static IPs for
        /// virtual machines; see IP Pools below for details.
        /// </summary>
        public InputList<Inputs.NsxtNetworkImportedStaticIpPoolGetArgs> StaticIpPools
        {
            get => _staticIpPools ?? (_staticIpPools = new InputList<Inputs.NsxtNetworkImportedStaticIpPoolGetArgs>());
            set => _staticIpPools = value;
        }

        /// <summary>
        /// The name of VDC to use. **Deprecated**  in favor of new field
        /// `owner_id` which supports VDC and VDC Group IDs.
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public NsxtNetworkImportedState()
        {
        }
        public static new NsxtNetworkImportedState Empty => new NsxtNetworkImportedState();
    }
}
