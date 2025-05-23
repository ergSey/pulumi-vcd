// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/nsxtAlbController:NsxtAlbController")]
    public partial class NsxtAlbController : global::Pulumi.CustomResource
    {
        /// <summary>
        /// An optional description ALB Controller
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// License type of ALB Controller (`ENTERPRISE` or `BASIC`)
        /// 
        /// &gt; The attribute `license_type` must not be used in VCD 10.4+, it is replaced by [nsxt_alb_service_engine_group](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_service_engine_group) and [nsxt_alb_settings](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_settings) attribute `supported_feature_set`.
        /// </summary>
        [Output("licenseType")]
        public Output<string?> LicenseType { get; private set; } = null!;

        /// <summary>
        /// A name for ALB Controller
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The password for ALB Controller. Password will not be refreshed.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// The URL of ALB Controller
        /// </summary>
        [Output("url")]
        public Output<string> Url { get; private set; } = null!;

        /// <summary>
        /// The username for ALB Controller
        /// </summary>
        [Output("username")]
        public Output<string> Username { get; private set; } = null!;

        /// <summary>
        /// ALB Controller version (e.g. 20.1.3)
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a NsxtAlbController resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public NsxtAlbController(string name, NsxtAlbControllerArgs args, CustomResourceOptions? options = null)
            : base("vcd:index/nsxtAlbController:NsxtAlbController", name, args ?? new NsxtAlbControllerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private NsxtAlbController(string name, Input<string> id, NsxtAlbControllerState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/nsxtAlbController:NsxtAlbController", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/vmware/terraform-provider-vcd/releases/download/v3.14.1/",
                AdditionalSecretOutputs =
                {
                    "password",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing NsxtAlbController resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static NsxtAlbController Get(string name, Input<string> id, NsxtAlbControllerState? state = null, CustomResourceOptions? options = null)
        {
            return new NsxtAlbController(name, id, state, options);
        }
    }

    public sealed class NsxtAlbControllerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional description ALB Controller
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// License type of ALB Controller (`ENTERPRISE` or `BASIC`)
        /// 
        /// &gt; The attribute `license_type` must not be used in VCD 10.4+, it is replaced by [nsxt_alb_service_engine_group](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_service_engine_group) and [nsxt_alb_settings](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_settings) attribute `supported_feature_set`.
        /// </summary>
        [Input("licenseType")]
        public Input<string>? LicenseType { get; set; }

        /// <summary>
        /// A name for ALB Controller
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password", required: true)]
        private Input<string>? _password;

        /// <summary>
        /// The password for ALB Controller. Password will not be refreshed.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The URL of ALB Controller
        /// </summary>
        [Input("url", required: true)]
        public Input<string> Url { get; set; } = null!;

        /// <summary>
        /// The username for ALB Controller
        /// </summary>
        [Input("username", required: true)]
        public Input<string> Username { get; set; } = null!;

        public NsxtAlbControllerArgs()
        {
        }
        public static new NsxtAlbControllerArgs Empty => new NsxtAlbControllerArgs();
    }

    public sealed class NsxtAlbControllerState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional description ALB Controller
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// License type of ALB Controller (`ENTERPRISE` or `BASIC`)
        /// 
        /// &gt; The attribute `license_type` must not be used in VCD 10.4+, it is replaced by [nsxt_alb_service_engine_group](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_service_engine_group) and [nsxt_alb_settings](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_settings) attribute `supported_feature_set`.
        /// </summary>
        [Input("licenseType")]
        public Input<string>? LicenseType { get; set; }

        /// <summary>
        /// A name for ALB Controller
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("password")]
        private Input<string>? _password;

        /// <summary>
        /// The password for ALB Controller. Password will not be refreshed.
        /// </summary>
        public Input<string>? Password
        {
            get => _password;
            set
            {
                var emptySecret = Output.CreateSecret(0);
                _password = Output.Tuple<Input<string>?, int>(value, emptySecret).Apply(t => t.Item1);
            }
        }

        /// <summary>
        /// The URL of ALB Controller
        /// </summary>
        [Input("url")]
        public Input<string>? Url { get; set; }

        /// <summary>
        /// The username for ALB Controller
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        /// <summary>
        /// ALB Controller version (e.g. 20.1.3)
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public NsxtAlbControllerState()
        {
        }
        public static new NsxtAlbControllerState Empty => new NsxtAlbControllerState();
    }
}
