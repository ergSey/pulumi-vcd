// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/nsxtAlbVirtualService:NsxtAlbVirtualService")]
    public partial class NsxtAlbVirtualService : global::Pulumi.CustomResource
    {
        /// <summary>
        /// One of `HTTP`, `HTTPS`, `L4`, `L4_TLS`.
        /// </summary>
        [Output("applicationProfileType")]
        public Output<string> ApplicationProfileType { get; private set; } = null!;

        /// <summary>
        /// ID reference of CA certificate. Required when `application_profile_type` is `HTTPS`
        /// or `L4_TLS`
        /// </summary>
        [Output("caCertificateId")]
        public Output<string?> CaCertificateId { get; private set; } = null!;

        /// <summary>
        /// An optional description ALB Virtual Service
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// An ID of NSX-T Edge Gateway. Can be looked up using
        /// [vcd.NsxtEdgegateway](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_edgegateway) data source
        /// </summary>
        [Output("edgeGatewayId")]
        public Output<string> EdgeGatewayId { get; private set; } = null!;

        /// <summary>
        /// Virtual Service is enabled or disabled (default true)
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// IPv6 Address for the service to listen on.
        /// </summary>
        [Output("ipv6VirtualIpAddress")]
        public Output<string?> Ipv6VirtualIpAddress { get; private set; } = null!;

        /// <summary>
        /// Preserves Client IP on a
        /// Virtual Service. **Note** - the following criteria must be matched to make transparent mode work:
        /// * ALB Pool membership must be configured in Group mode
        /// * Backing Avi Service Engine Group must be in Legacy Active Standby mode
        /// 
        /// &lt;a id="service-port-block"&gt;&lt;/a&gt;
        /// </summary>
        [Output("isTransparentModeEnabled")]
        public Output<bool> IsTransparentModeEnabled { get; private set; } = null!;

        /// <summary>
        /// A name for ALB Virtual Service
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful
        /// when connected as sysadmin working across different organisations.
        /// </summary>
        [Output("org")]
        public Output<string?> Org { get; private set; } = null!;

        /// <summary>
        /// A reference to ALB Pool. Can be looked up using `vcd.NsxtAlbPool` resource or data
        /// source
        /// </summary>
        [Output("poolId")]
        public Output<string> PoolId { get; private set; } = null!;

        /// <summary>
        /// A reference to ALB Service Engine Group. Can be looked up using
        /// `vcd.NsxtAlbEdgegatewayServiceEngineGroup` resource or data source
        /// </summary>
        [Output("serviceEngineGroupId")]
        public Output<string> ServiceEngineGroupId { get; private set; } = null!;

        /// <summary>
        /// A block to define port, port range and traffic type. Multiple can be used. See
        /// service_port and example for usage details.
        /// </summary>
        [Output("servicePorts")]
        public Output<ImmutableArray<Outputs.NsxtAlbVirtualServiceServicePort>> ServicePorts { get; private set; } = null!;

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Output("vdc")]
        public Output<string> Vdc { get; private set; } = null!;

        /// <summary>
        /// IP Address for the service to listen on.
        /// </summary>
        [Output("virtualIpAddress")]
        public Output<string> VirtualIpAddress { get; private set; } = null!;


        /// <summary>
        /// Create a NsxtAlbVirtualService resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NsxtAlbVirtualService(string name, NsxtAlbVirtualServiceArgs args, CustomResourceOptions? options = null)
            : base("vcd:index/nsxtAlbVirtualService:NsxtAlbVirtualService", name, args ?? new NsxtAlbVirtualServiceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NsxtAlbVirtualService(string name, Input<string> id, NsxtAlbVirtualServiceState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/nsxtAlbVirtualService:NsxtAlbVirtualService", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NsxtAlbVirtualService resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NsxtAlbVirtualService Get(string name, Input<string> id, NsxtAlbVirtualServiceState? state = null, CustomResourceOptions? options = null)
        {
            return new NsxtAlbVirtualService(name, id, state, options);
        }
    }

    public sealed class NsxtAlbVirtualServiceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// One of `HTTP`, `HTTPS`, `L4`, `L4_TLS`.
        /// </summary>
        [Input("applicationProfileType", required: true)]
        public Input<string> ApplicationProfileType { get; set; } = null!;

        /// <summary>
        /// ID reference of CA certificate. Required when `application_profile_type` is `HTTPS`
        /// or `L4_TLS`
        /// </summary>
        [Input("caCertificateId")]
        public Input<string>? CaCertificateId { get; set; }

        /// <summary>
        /// An optional description ALB Virtual Service
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// An ID of NSX-T Edge Gateway. Can be looked up using
        /// [vcd.NsxtEdgegateway](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_edgegateway) data source
        /// </summary>
        [Input("edgeGatewayId", required: true)]
        public Input<string> EdgeGatewayId { get; set; } = null!;

        /// <summary>
        /// Virtual Service is enabled or disabled (default true)
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// IPv6 Address for the service to listen on.
        /// </summary>
        [Input("ipv6VirtualIpAddress")]
        public Input<string>? Ipv6VirtualIpAddress { get; set; }

        /// <summary>
        /// Preserves Client IP on a
        /// Virtual Service. **Note** - the following criteria must be matched to make transparent mode work:
        /// * ALB Pool membership must be configured in Group mode
        /// * Backing Avi Service Engine Group must be in Legacy Active Standby mode
        /// 
        /// &lt;a id="service-port-block"&gt;&lt;/a&gt;
        /// </summary>
        [Input("isTransparentModeEnabled")]
        public Input<bool>? IsTransparentModeEnabled { get; set; }

        /// <summary>
        /// A name for ALB Virtual Service
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful
        /// when connected as sysadmin working across different organisations.
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// A reference to ALB Pool. Can be looked up using `vcd.NsxtAlbPool` resource or data
        /// source
        /// </summary>
        [Input("poolId", required: true)]
        public Input<string> PoolId { get; set; } = null!;

        /// <summary>
        /// A reference to ALB Service Engine Group. Can be looked up using
        /// `vcd.NsxtAlbEdgegatewayServiceEngineGroup` resource or data source
        /// </summary>
        [Input("serviceEngineGroupId", required: true)]
        public Input<string> ServiceEngineGroupId { get; set; } = null!;

        [Input("servicePorts")]
        private InputList<Inputs.NsxtAlbVirtualServiceServicePortArgs>? _servicePorts;

        /// <summary>
        /// A block to define port, port range and traffic type. Multiple can be used. See
        /// service_port and example for usage details.
        /// </summary>
        public InputList<Inputs.NsxtAlbVirtualServiceServicePortArgs> ServicePorts
        {
            get => _servicePorts ?? (_servicePorts = new InputList<Inputs.NsxtAlbVirtualServiceServicePortArgs>());
            set => _servicePorts = value;
        }

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        /// <summary>
        /// IP Address for the service to listen on.
        /// </summary>
        [Input("virtualIpAddress", required: true)]
        public Input<string> VirtualIpAddress { get; set; } = null!;

        public NsxtAlbVirtualServiceArgs()
        {
        }
        public static new NsxtAlbVirtualServiceArgs Empty => new NsxtAlbVirtualServiceArgs();
    }

    public sealed class NsxtAlbVirtualServiceState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// One of `HTTP`, `HTTPS`, `L4`, `L4_TLS`.
        /// </summary>
        [Input("applicationProfileType")]
        public Input<string>? ApplicationProfileType { get; set; }

        /// <summary>
        /// ID reference of CA certificate. Required when `application_profile_type` is `HTTPS`
        /// or `L4_TLS`
        /// </summary>
        [Input("caCertificateId")]
        public Input<string>? CaCertificateId { get; set; }

        /// <summary>
        /// An optional description ALB Virtual Service
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// An ID of NSX-T Edge Gateway. Can be looked up using
        /// [vcd.NsxtEdgegateway](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_edgegateway) data source
        /// </summary>
        [Input("edgeGatewayId")]
        public Input<string>? EdgeGatewayId { get; set; }

        /// <summary>
        /// Virtual Service is enabled or disabled (default true)
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// IPv6 Address for the service to listen on.
        /// </summary>
        [Input("ipv6VirtualIpAddress")]
        public Input<string>? Ipv6VirtualIpAddress { get; set; }

        /// <summary>
        /// Preserves Client IP on a
        /// Virtual Service. **Note** - the following criteria must be matched to make transparent mode work:
        /// * ALB Pool membership must be configured in Group mode
        /// * Backing Avi Service Engine Group must be in Legacy Active Standby mode
        /// 
        /// &lt;a id="service-port-block"&gt;&lt;/a&gt;
        /// </summary>
        [Input("isTransparentModeEnabled")]
        public Input<bool>? IsTransparentModeEnabled { get; set; }

        /// <summary>
        /// A name for ALB Virtual Service
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful
        /// when connected as sysadmin working across different organisations.
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// A reference to ALB Pool. Can be looked up using `vcd.NsxtAlbPool` resource or data
        /// source
        /// </summary>
        [Input("poolId")]
        public Input<string>? PoolId { get; set; }

        /// <summary>
        /// A reference to ALB Service Engine Group. Can be looked up using
        /// `vcd.NsxtAlbEdgegatewayServiceEngineGroup` resource or data source
        /// </summary>
        [Input("serviceEngineGroupId")]
        public Input<string>? ServiceEngineGroupId { get; set; }

        [Input("servicePorts")]
        private InputList<Inputs.NsxtAlbVirtualServiceServicePortGetArgs>? _servicePorts;

        /// <summary>
        /// A block to define port, port range and traffic type. Multiple can be used. See
        /// service_port and example for usage details.
        /// </summary>
        public InputList<Inputs.NsxtAlbVirtualServiceServicePortGetArgs> ServicePorts
        {
            get => _servicePorts ?? (_servicePorts = new InputList<Inputs.NsxtAlbVirtualServiceServicePortGetArgs>());
            set => _servicePorts = value;
        }

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        /// <summary>
        /// IP Address for the service to listen on.
        /// </summary>
        [Input("virtualIpAddress")]
        public Input<string>? VirtualIpAddress { get; set; }

        public NsxtAlbVirtualServiceState()
        {
        }
        public static new NsxtAlbVirtualServiceState Empty => new NsxtAlbVirtualServiceState();
    }
}
