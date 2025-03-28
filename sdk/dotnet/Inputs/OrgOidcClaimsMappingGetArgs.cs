// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class OrgOidcClaimsMappingGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Required if `wellknown_endpoint` doesn't give info about it
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        /// <summary>
        /// Required if `wellknown_endpoint` doesn't give info about it
        /// </summary>
        [Input("firstName")]
        public Input<string>? FirstName { get; set; }

        /// <summary>
        /// Required if `wellknown_endpoint` doesn't give info about it
        /// </summary>
        [Input("fullName")]
        public Input<string>? FullName { get; set; }

        /// <summary>
        /// Optional
        /// </summary>
        [Input("groups")]
        public Input<string>? Groups { get; set; }

        /// <summary>
        /// Required if `wellknown_endpoint` doesn't give info about it
        /// </summary>
        [Input("lastName")]
        public Input<string>? LastName { get; set; }

        /// <summary>
        /// Optional
        /// </summary>
        [Input("roles")]
        public Input<string>? Roles { get; set; }

        /// <summary>
        /// Required if `wellknown_endpoint` doesn't give info about it
        /// </summary>
        [Input("subject")]
        public Input<string>? Subject { get; set; }

        public OrgOidcClaimsMappingGetArgs()
        {
        }
        public static new OrgOidcClaimsMappingGetArgs Empty => new OrgOidcClaimsMappingGetArgs();
    }
}
