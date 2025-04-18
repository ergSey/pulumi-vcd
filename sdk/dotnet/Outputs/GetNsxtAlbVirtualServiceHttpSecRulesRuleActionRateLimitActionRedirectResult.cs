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
    public sealed class GetNsxtAlbVirtualServiceHttpSecRulesRuleActionRateLimitActionRedirectResult
    {
        /// <summary>
        /// Host to which redirect the request. Default is the original host
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// Path to which redirect the request. Default is the original path
        /// </summary>
        public readonly bool KeepQuery;
        /// <summary>
        /// Port to which redirect the request. Default is 80 for HTTP and 443 for HTTPS protocol
        /// </summary>
        public readonly string Path;
        /// <summary>
        /// Port to which redirect the request. Default is 80 for HTTP and 443 for HTTPS protocol
        /// </summary>
        public readonly string Port;
        /// <summary>
        /// HTTP or HTTPS protocol
        /// </summary>
        public readonly string Protocol;
        /// <summary>
        /// One of the redirect status codes - 301, 302, 307
        /// </summary>
        public readonly int StatusCode;

        [OutputConstructor]
        private GetNsxtAlbVirtualServiceHttpSecRulesRuleActionRateLimitActionRedirectResult(
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
