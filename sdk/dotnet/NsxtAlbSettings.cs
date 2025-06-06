// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/nsxtAlbSettings:NsxtAlbSettings")]
    public partial class NsxtAlbSettings : global::Pulumi.CustomResource
    {
        /// <summary>
        /// An ID of NSX-T Edge Gateway. Can be looked up using
        /// [vcd.NsxtEdgegateway](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_edgegateway) data source
        /// </summary>
        [Output("edgeGatewayId")]
        public Output<string> EdgeGatewayId { get; private set; } = null!;

        /// <summary>
        /// The IPv6 network
        /// definition in CIDR format which will be used by Load Balancer service in the Edge Gateway. All the
        /// load balancer service engines associated with the Service Engine Group will be attached to this
        /// network. This field cannot be updated
        /// 
        /// &gt; IPv4 service network will be used if both the `service_network_specification` and
        /// `ipv6_service_network_specification` properties are unset. If both are set, it will still be one
        /// service network with a dual IPv4 and IPv6 stack. If only `ipv6_service_network_specification` is
        /// used, then only IPv6 will be used.
        /// 
        /// &gt; The attribute `supported_feature_set` must not be used in VCD versions lower than 10.4. Starting with 10.4, it replaces `license_type` field in [nsxt_alb_controller](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_controller).
        /// </summary>
        [Output("ipv6ServiceNetworkSpecification")]
        public Output<string> Ipv6ServiceNetworkSpecification { get; private set; } = null!;

        /// <summary>
        /// Boolean value `true` or `false` if ALB is enabled. **Note** Delete operation of this resource
        /// will set it to `false`
        /// </summary>
        [Output("isActive")]
        public Output<bool> IsActive { get; private set; } = null!;

        /// <summary>
        /// When enabled, it allows to
        /// configure Preserve Client IP on a Virtual Service
        /// </summary>
        [Output("isTransparentModeEnabled")]
        public Output<bool> IsTransparentModeEnabled { get; private set; } = null!;

        /// <summary>
        /// The name of organization to which the edge gateway belongs. Optional if defined at provider level.
        /// </summary>
        [Output("org")]
        public Output<string?> Org { get; private set; } = null!;

        /// <summary>
        /// Gateway CIDR format which will be used by Load
        /// Balancer service. All the load balancer service engines associated with the Service Engine Group
        /// will be attached to this network. The subnet prefix length must be 25. If nothing is set and
        /// `ipv6_service_network_specification` is not used, the **default is 192.168.255.125/25**. This
        /// field cannot be updated
        /// </summary>
        [Output("serviceNetworkSpecification")]
        public Output<string> ServiceNetworkSpecification { get; private set; } = null!;

        /// <summary>
        /// Feature set of this Edge Gateway if ALB is enabled (`STANDARD` or `PREMIUM`)
        /// </summary>
        [Output("supportedFeatureSet")]
        public Output<string> SupportedFeatureSet { get; private set; } = null!;

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Output("vdc")]
        public Output<string> Vdc { get; private set; } = null!;


        /// <summary>
        /// Create a NsxtAlbSettings resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NsxtAlbSettings(string name, NsxtAlbSettingsArgs args, CustomResourceOptions? options = null)
            : base("vcd:index/nsxtAlbSettings:NsxtAlbSettings", name, args ?? new NsxtAlbSettingsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NsxtAlbSettings(string name, Input<string> id, NsxtAlbSettingsState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/nsxtAlbSettings:NsxtAlbSettings", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NsxtAlbSettings resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NsxtAlbSettings Get(string name, Input<string> id, NsxtAlbSettingsState? state = null, CustomResourceOptions? options = null)
        {
            return new NsxtAlbSettings(name, id, state, options);
        }
    }

    public sealed class NsxtAlbSettingsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An ID of NSX-T Edge Gateway. Can be looked up using
        /// [vcd.NsxtEdgegateway](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_edgegateway) data source
        /// </summary>
        [Input("edgeGatewayId", required: true)]
        public Input<string> EdgeGatewayId { get; set; } = null!;

        /// <summary>
        /// The IPv6 network
        /// definition in CIDR format which will be used by Load Balancer service in the Edge Gateway. All the
        /// load balancer service engines associated with the Service Engine Group will be attached to this
        /// network. This field cannot be updated
        /// 
        /// &gt; IPv4 service network will be used if both the `service_network_specification` and
        /// `ipv6_service_network_specification` properties are unset. If both are set, it will still be one
        /// service network with a dual IPv4 and IPv6 stack. If only `ipv6_service_network_specification` is
        /// used, then only IPv6 will be used.
        /// 
        /// &gt; The attribute `supported_feature_set` must not be used in VCD versions lower than 10.4. Starting with 10.4, it replaces `license_type` field in [nsxt_alb_controller](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_controller).
        /// </summary>
        [Input("ipv6ServiceNetworkSpecification")]
        public Input<string>? Ipv6ServiceNetworkSpecification { get; set; }

        /// <summary>
        /// Boolean value `true` or `false` if ALB is enabled. **Note** Delete operation of this resource
        /// will set it to `false`
        /// </summary>
        [Input("isActive", required: true)]
        public Input<bool> IsActive { get; set; } = null!;

        /// <summary>
        /// When enabled, it allows to
        /// configure Preserve Client IP on a Virtual Service
        /// </summary>
        [Input("isTransparentModeEnabled")]
        public Input<bool>? IsTransparentModeEnabled { get; set; }

        /// <summary>
        /// The name of organization to which the edge gateway belongs. Optional if defined at provider level.
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// Gateway CIDR format which will be used by Load
        /// Balancer service. All the load balancer service engines associated with the Service Engine Group
        /// will be attached to this network. The subnet prefix length must be 25. If nothing is set and
        /// `ipv6_service_network_specification` is not used, the **default is 192.168.255.125/25**. This
        /// field cannot be updated
        /// </summary>
        [Input("serviceNetworkSpecification")]
        public Input<string>? ServiceNetworkSpecification { get; set; }

        /// <summary>
        /// Feature set of this Edge Gateway if ALB is enabled (`STANDARD` or `PREMIUM`)
        /// </summary>
        [Input("supportedFeatureSet")]
        public Input<string>? SupportedFeatureSet { get; set; }

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public NsxtAlbSettingsArgs()
        {
        }
        public static new NsxtAlbSettingsArgs Empty => new NsxtAlbSettingsArgs();
    }

    public sealed class NsxtAlbSettingsState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An ID of NSX-T Edge Gateway. Can be looked up using
        /// [vcd.NsxtEdgegateway](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_edgegateway) data source
        /// </summary>
        [Input("edgeGatewayId")]
        public Input<string>? EdgeGatewayId { get; set; }

        /// <summary>
        /// The IPv6 network
        /// definition in CIDR format which will be used by Load Balancer service in the Edge Gateway. All the
        /// load balancer service engines associated with the Service Engine Group will be attached to this
        /// network. This field cannot be updated
        /// 
        /// &gt; IPv4 service network will be used if both the `service_network_specification` and
        /// `ipv6_service_network_specification` properties are unset. If both are set, it will still be one
        /// service network with a dual IPv4 and IPv6 stack. If only `ipv6_service_network_specification` is
        /// used, then only IPv6 will be used.
        /// 
        /// &gt; The attribute `supported_feature_set` must not be used in VCD versions lower than 10.4. Starting with 10.4, it replaces `license_type` field in [nsxt_alb_controller](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_controller).
        /// </summary>
        [Input("ipv6ServiceNetworkSpecification")]
        public Input<string>? Ipv6ServiceNetworkSpecification { get; set; }

        /// <summary>
        /// Boolean value `true` or `false` if ALB is enabled. **Note** Delete operation of this resource
        /// will set it to `false`
        /// </summary>
        [Input("isActive")]
        public Input<bool>? IsActive { get; set; }

        /// <summary>
        /// When enabled, it allows to
        /// configure Preserve Client IP on a Virtual Service
        /// </summary>
        [Input("isTransparentModeEnabled")]
        public Input<bool>? IsTransparentModeEnabled { get; set; }

        /// <summary>
        /// The name of organization to which the edge gateway belongs. Optional if defined at provider level.
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// Gateway CIDR format which will be used by Load
        /// Balancer service. All the load balancer service engines associated with the Service Engine Group
        /// will be attached to this network. The subnet prefix length must be 25. If nothing is set and
        /// `ipv6_service_network_specification` is not used, the **default is 192.168.255.125/25**. This
        /// field cannot be updated
        /// </summary>
        [Input("serviceNetworkSpecification")]
        public Input<string>? ServiceNetworkSpecification { get; set; }

        /// <summary>
        /// Feature set of this Edge Gateway if ALB is enabled (`STANDARD` or `PREMIUM`)
        /// </summary>
        [Input("supportedFeatureSet")]
        public Input<string>? SupportedFeatureSet { get; set; }

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public NsxtAlbSettingsState()
        {
        }
        public static new NsxtAlbSettingsState Empty => new NsxtAlbSettingsState();
    }
}
