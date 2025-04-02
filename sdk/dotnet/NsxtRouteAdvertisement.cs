// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/nsxtRouteAdvertisement:NsxtRouteAdvertisement")]
    public partial class NsxtRouteAdvertisement : global::Pulumi.CustomResource
    {
        /// <summary>
        /// NSX-T Edge Gateway ID in which route advertisement is located.
        /// </summary>
        [Output("edgeGatewayId")]
        public Output<string> EdgeGatewayId { get; private set; } = null!;

        /// <summary>
        /// Define if route advertisement is active. Default `true`.
        /// </summary>
        [Output("enabled")]
        public Output<bool?> Enabled { get; private set; } = null!;

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful
        /// when connected as sysadmin working across different organizations.
        /// </summary>
        [Output("org")]
        public Output<string?> Org { get; private set; } = null!;

        /// <summary>
        /// Set of subnets that will be advertised to Tier-0 gateway. Leaving it empty means none.
        /// </summary>
        [Output("subnets")]
        public Output<ImmutableArray<string>> Subnets { get; private set; } = null!;


        /// <summary>
        /// Create a NsxtRouteAdvertisement resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NsxtRouteAdvertisement(string name, NsxtRouteAdvertisementArgs args, CustomResourceOptions? options = null)
            : base("vcd:index/nsxtRouteAdvertisement:NsxtRouteAdvertisement", name, args ?? new NsxtRouteAdvertisementArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NsxtRouteAdvertisement(string name, Input<string> id, NsxtRouteAdvertisementState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/nsxtRouteAdvertisement:NsxtRouteAdvertisement", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing NsxtRouteAdvertisement resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NsxtRouteAdvertisement Get(string name, Input<string> id, NsxtRouteAdvertisementState? state = null, CustomResourceOptions? options = null)
        {
            return new NsxtRouteAdvertisement(name, id, state, options);
        }
    }

    public sealed class NsxtRouteAdvertisementArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// NSX-T Edge Gateway ID in which route advertisement is located.
        /// </summary>
        [Input("edgeGatewayId", required: true)]
        public Input<string> EdgeGatewayId { get; set; } = null!;

        /// <summary>
        /// Define if route advertisement is active. Default `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful
        /// when connected as sysadmin working across different organizations.
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        [Input("subnets")]
        private InputList<string>? _subnets;

        /// <summary>
        /// Set of subnets that will be advertised to Tier-0 gateway. Leaving it empty means none.
        /// </summary>
        public InputList<string> Subnets
        {
            get => _subnets ?? (_subnets = new InputList<string>());
            set => _subnets = value;
        }

        public NsxtRouteAdvertisementArgs()
        {
        }
        public static new NsxtRouteAdvertisementArgs Empty => new NsxtRouteAdvertisementArgs();
    }

    public sealed class NsxtRouteAdvertisementState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// NSX-T Edge Gateway ID in which route advertisement is located.
        /// </summary>
        [Input("edgeGatewayId")]
        public Input<string>? EdgeGatewayId { get; set; }

        /// <summary>
        /// Define if route advertisement is active. Default `true`.
        /// </summary>
        [Input("enabled")]
        public Input<bool>? Enabled { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful
        /// when connected as sysadmin working across different organizations.
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        [Input("subnets")]
        private InputList<string>? _subnets;

        /// <summary>
        /// Set of subnets that will be advertised to Tier-0 gateway. Leaving it empty means none.
        /// </summary>
        public InputList<string> Subnets
        {
            get => _subnets ?? (_subnets = new InputList<string>());
            set => _subnets = value;
        }

        public NsxtRouteAdvertisementState()
        {
        }
        public static new NsxtRouteAdvertisementState Empty => new NsxtRouteAdvertisementState();
    }
}
