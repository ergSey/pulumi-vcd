// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/nsxtEdgegatewayStaticRoute:NsxtEdgegatewayStaticRoute")]
    public partial class NsxtEdgegatewayStaticRoute : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Description for NSX-T Edge Gateway Static Route
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// NSX-T Edge Gateway ID
        /// </summary>
        [Output("edgeGatewayId")]
        public Output<string> EdgeGatewayId { get; private set; } = null!;

        /// <summary>
        /// Name for NSX-T Edge Gateway Static Route
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Specifies network prefix in CIDR format. Both IPv4 and IPv6 formats
        /// are supported.
        /// </summary>
        [Output("networkCidr")]
        public Output<string> NetworkCidr { get; private set; } = null!;

        /// <summary>
        /// A set of next hops to use within the static route. At least one is
        /// required. See Next Hop for definition structure.
        /// 
        /// &lt;a id="next-hop"&gt;&lt;/a&gt;
        /// </summary>
        [Output("nextHops")]
        public Output<ImmutableArray<Outputs.NsxtEdgegatewayStaticRouteNextHop>> NextHops { get; private set; } = null!;

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        /// different organizations
        /// </summary>
        [Output("org")]
        public Output<string?> Org { get; private set; } = null!;


        /// <summary>
        /// Create a NsxtEdgegatewayStaticRoute resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NsxtEdgegatewayStaticRoute(string name, NsxtEdgegatewayStaticRouteArgs args, CustomResourceOptions? options = null)
            : base("vcd:index/nsxtEdgegatewayStaticRoute:NsxtEdgegatewayStaticRoute", name, args ?? new NsxtEdgegatewayStaticRouteArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NsxtEdgegatewayStaticRoute(string name, Input<string> id, NsxtEdgegatewayStaticRouteState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/nsxtEdgegatewayStaticRoute:NsxtEdgegatewayStaticRoute", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NsxtEdgegatewayStaticRoute resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NsxtEdgegatewayStaticRoute Get(string name, Input<string> id, NsxtEdgegatewayStaticRouteState? state = null, CustomResourceOptions? options = null)
        {
            return new NsxtEdgegatewayStaticRoute(name, id, state, options);
        }
    }

    public sealed class NsxtEdgegatewayStaticRouteArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description for NSX-T Edge Gateway Static Route
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// NSX-T Edge Gateway ID
        /// </summary>
        [Input("edgeGatewayId", required: true)]
        public Input<string> EdgeGatewayId { get; set; } = null!;

        /// <summary>
        /// Name for NSX-T Edge Gateway Static Route
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies network prefix in CIDR format. Both IPv4 and IPv6 formats
        /// are supported.
        /// </summary>
        [Input("networkCidr", required: true)]
        public Input<string> NetworkCidr { get; set; } = null!;

        [Input("nextHops", required: true)]
        private InputList<Inputs.NsxtEdgegatewayStaticRouteNextHopArgs>? _nextHops;

        /// <summary>
        /// A set of next hops to use within the static route. At least one is
        /// required. See Next Hop for definition structure.
        /// 
        /// &lt;a id="next-hop"&gt;&lt;/a&gt;
        /// </summary>
        public InputList<Inputs.NsxtEdgegatewayStaticRouteNextHopArgs> NextHops
        {
            get => _nextHops ?? (_nextHops = new InputList<Inputs.NsxtEdgegatewayStaticRouteNextHopArgs>());
            set => _nextHops = value;
        }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        /// different organizations
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        public NsxtEdgegatewayStaticRouteArgs()
        {
        }
        public static new NsxtEdgegatewayStaticRouteArgs Empty => new NsxtEdgegatewayStaticRouteArgs();
    }

    public sealed class NsxtEdgegatewayStaticRouteState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Description for NSX-T Edge Gateway Static Route
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// NSX-T Edge Gateway ID
        /// </summary>
        [Input("edgeGatewayId")]
        public Input<string>? EdgeGatewayId { get; set; }

        /// <summary>
        /// Name for NSX-T Edge Gateway Static Route
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Specifies network prefix in CIDR format. Both IPv4 and IPv6 formats
        /// are supported.
        /// </summary>
        [Input("networkCidr")]
        public Input<string>? NetworkCidr { get; set; }

        [Input("nextHops")]
        private InputList<Inputs.NsxtEdgegatewayStaticRouteNextHopGetArgs>? _nextHops;

        /// <summary>
        /// A set of next hops to use within the static route. At least one is
        /// required. See Next Hop for definition structure.
        /// 
        /// &lt;a id="next-hop"&gt;&lt;/a&gt;
        /// </summary>
        public InputList<Inputs.NsxtEdgegatewayStaticRouteNextHopGetArgs> NextHops
        {
            get => _nextHops ?? (_nextHops = new InputList<Inputs.NsxtEdgegatewayStaticRouteNextHopGetArgs>());
            set => _nextHops = value;
        }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        /// different organizations
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        public NsxtEdgegatewayStaticRouteState()
        {
        }
        public static new NsxtEdgegatewayStaticRouteState Empty => new NsxtEdgegatewayStaticRouteState();
    }
}
