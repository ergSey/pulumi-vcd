// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Outputs
{

    [OutputType]
    public sealed class GetNsxtAlbVirtualServiceHttpReqRulesRuleActionRedirectResult
    {
        /// <summary>
        /// Host to which redirect the request
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// Path to which redirect the request
        /// </summary>
        public readonly bool KeepQuery;
        /// <summary>
        /// Port to which redirect the request
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// Port to which the request will be redirected
        /// </summary>
        public readonly string Port;
        /// <summary>
        /// HTTP or HTTPS protocol
        /// </summary>
        public readonly string Protocol;
        /// <summary>
        /// Redirect status code
        /// </summary>
        public readonly int StatusCode;

        [OutputConstructor]
        private GetNsxtAlbVirtualServiceHttpReqRulesRuleActionRedirectResult(
            string host,

            bool keepQuery,

            string path,

            string port,

            string protocol,

            int statusCode)
        {
            Host = host;
            KeepQuery = keepQuery;
            Path = path;
            Port = port;
            Protocol = protocol;
            StatusCode = statusCode;
        }
    }
}
