// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/serviceAccount:ServiceAccount")]
    public partial class ServiceAccount : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Status of the Service Account. Can be set to `false` and back to `true` if
        /// the access token was lost to get a new one.
        /// </summary>
        [Output("active")]
        public Output<bool?> Active { get; private set; } = null!;

        /// <summary>
        /// If set to false, will output a warning about the service account file
        /// containing sensitive information.
        /// </summary>
        [Output("allowTokenFile")]
        public Output<bool?> AllowTokenFile { get; private set; } = null!;

        /// <summary>
        /// Required only when `active` is set to `true`. Contains the access token
        /// that can be used for authenticating to VCD.
        /// </summary>
        [Output("fileName")]
        public Output<string?> FileName { get; private set; } = null!;

        /// <summary>
        /// A unique name for the Service Account in an organisation.
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
        /// Role ID of service account
        /// </summary>
        [Output("roleId")]
        public Output<string> RoleId { get; private set; } = null!;

        /// <summary>
        /// UUID of the Service Account.
        /// </summary>
        [Output("softwareId")]
        public Output<string> SoftwareId { get; private set; } = null!;

        /// <summary>
        /// Version of the service using the Service Account
        /// </summary>
        [Output("softwareVersion")]
        public Output<string?> SoftwareVersion { get; private set; } = null!;

        /// <summary>
        /// URI of the service using the Service Account
        /// </summary>
        [Output("uri")]
        public Output<string?> Uri { get; private set; } = null!;


        /// <summary>
        /// Create a ServiceAccount resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ServiceAccount(string name, ServiceAccountArgs args, CustomResourceOptions? options = null)
            : base("vcd:index/serviceAccount:ServiceAccount", name, args ?? new ServiceAccountArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ServiceAccount(string name, Input<string> id, ServiceAccountState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/serviceAccount:ServiceAccount", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ServiceAccount resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ServiceAccount Get(string name, Input<string> id, ServiceAccountState? state = null, CustomResourceOptions? options = null)
        {
            return new ServiceAccount(name, id, state, options);
        }
    }

    public sealed class ServiceAccountArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Status of the Service Account. Can be set to `false` and back to `true` if
        /// the access token was lost to get a new one.
        /// </summary>
        [Input("active")]
        public Input<bool>? Active { get; set; }

        /// <summary>
        /// If set to false, will output a warning about the service account file
        /// containing sensitive information.
        /// </summary>
        [Input("allowTokenFile")]
        public Input<bool>? AllowTokenFile { get; set; }

        /// <summary>
        /// Required only when `active` is set to `true`. Contains the access token
        /// that can be used for authenticating to VCD.
        /// </summary>
        [Input("fileName")]
        public Input<string>? FileName { get; set; }

        /// <summary>
        /// A unique name for the Service Account in an organisation.
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
        /// Role ID of service account
        /// </summary>
        [Input("roleId", required: true)]
        public Input<string> RoleId { get; set; } = null!;

        /// <summary>
        /// UUID of the Service Account.
        /// </summary>
        [Input("softwareId", required: true)]
        public Input<string> SoftwareId { get; set; } = null!;

        /// <summary>
        /// Version of the service using the Service Account
        /// </summary>
        [Input("softwareVersion")]
        public Input<string>? SoftwareVersion { get; set; }

        /// <summary>
        /// URI of the service using the Service Account
        /// </summary>
        [Input("uri")]
        public Input<string>? Uri { get; set; }

        public ServiceAccountArgs()
        {
        }
        public static new ServiceAccountArgs Empty => new ServiceAccountArgs();
    }

    public sealed class ServiceAccountState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Status of the Service Account. Can be set to `false` and back to `true` if
        /// the access token was lost to get a new one.
        /// </summary>
        [Input("active")]
        public Input<bool>? Active { get; set; }

        /// <summary>
        /// If set to false, will output a warning about the service account file
        /// containing sensitive information.
        /// </summary>
        [Input("allowTokenFile")]
        public Input<bool>? AllowTokenFile { get; set; }

        /// <summary>
        /// Required only when `active` is set to `true`. Contains the access token
        /// that can be used for authenticating to VCD.
        /// </summary>
        [Input("fileName")]
        public Input<string>? FileName { get; set; }

        /// <summary>
        /// A unique name for the Service Account in an organisation.
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
        /// Role ID of service account
        /// </summary>
        [Input("roleId")]
        public Input<string>? RoleId { get; set; }

        /// <summary>
        /// UUID of the Service Account.
        /// </summary>
        [Input("softwareId")]
        public Input<string>? SoftwareId { get; set; }

        /// <summary>
        /// Version of the service using the Service Account
        /// </summary>
        [Input("softwareVersion")]
        public Input<string>? SoftwareVersion { get; set; }

        /// <summary>
        /// URI of the service using the Service Account
        /// </summary>
        [Input("uri")]
        public Input<string>? Uri { get; set; }

        public ServiceAccountState()
        {
        }
        public static new ServiceAccountState Empty => new ServiceAccountState();
    }
}
