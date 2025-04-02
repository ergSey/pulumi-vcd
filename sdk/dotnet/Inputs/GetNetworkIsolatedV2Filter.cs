// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class GetNetworkIsolatedV2FilterArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Search by IP. The value can be a regular expression
        /// </summary>
        [Input("ip")]
        public string? Ip { get; set; }

        /// <summary>
        /// Search by name with a regular expression
        /// </summary>
        [Input("nameRegex")]
        public string? NameRegex { get; set; }

        public GetNetworkIsolatedV2FilterArgs()
        {
        }
        public static new GetNetworkIsolatedV2FilterArgs Empty => new GetNetworkIsolatedV2FilterArgs();
    }
}
