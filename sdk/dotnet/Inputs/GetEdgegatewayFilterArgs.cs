// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class GetEdgegatewayFilterInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Search by name with a regular expression
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        public GetEdgegatewayFilterInputArgs()
        {
        }
        public static new GetEdgegatewayFilterInputArgs Empty => new GetEdgegatewayFilterInputArgs();
    }
}
