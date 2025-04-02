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
    public sealed class GetNsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaResult
    {
        /// <summary>
        /// Criteria for matching client IP Address
        /// </summary>
        public readonly ImmutableArray<Outputs.GetNsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaClientIpAddressResult> ClientIpAddresses;
        /// <summary>
        /// Rule for matching cookie
        /// </summary>
        public readonly ImmutableArray<Outputs.GetNsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaCookieResult> Cookies;
        /// <summary>
        /// Criteria to match HTTP methods
        /// </summary>
        public readonly ImmutableArray<Outputs.GetNsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaHttpMethodResult> HttpMethods;
        /// <summary>
        /// A matching criteria for Location header
        /// </summary>
        public readonly ImmutableArray<Outputs.GetNsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaLocationHeaderResult> LocationHeaders;
        /// <summary>
        /// Criteria for matching request paths
        /// </summary>
        public readonly ImmutableArray<Outputs.GetNsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaPathResult> Paths;
        /// <summary>
        /// Protocol to match - 'HTTP' or 'HTTPS'
        /// </summary>
        public readonly string ProtocolType;
        /// <summary>
        /// HTTP request query strings to match
        /// </summary>
        public readonly ImmutableArray<string> Queries;
        /// <summary>
        /// A set of rules for matching request headers
        /// </summary>
        public readonly ImmutableArray<Outputs.GetNsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaRequestHeaderResult> RequestHeaders;
        /// <summary>
        /// A set of criteria to match response headers
        /// </summary>
        public readonly ImmutableArray<Outputs.GetNsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaResponseHeaderResult> ResponseHeaders;
        /// <summary>
        /// Criteria for matching service ports
        /// </summary>
        public readonly ImmutableArray<Outputs.GetNsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaServicePortResult> ServicePorts;
        /// <summary>
        /// HTTP Status code to match
        /// </summary>
        public readonly ImmutableArray<Outputs.GetNsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaStatusCodeResult> StatusCodes;

        [OutputConstructor]
        private GetNsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaResult(
            ImmutableArray<Outputs.GetNsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaClientIpAddressResult> clientIpAddresses,

            ImmutableArray<Outputs.GetNsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaCookieResult> cookies,

            ImmutableArray<Outputs.GetNsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaHttpMethodResult> httpMethods,

            ImmutableArray<Outputs.GetNsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaLocationHeaderResult> locationHeaders,

            ImmutableArray<Outputs.GetNsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaPathResult> paths,

            string protocolType,

            ImmutableArray<string> queries,

            ImmutableArray<Outputs.GetNsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaRequestHeaderResult> requestHeaders,

            ImmutableArray<Outputs.GetNsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaResponseHeaderResult> responseHeaders,

            ImmutableArray<Outputs.GetNsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaServicePortResult> servicePorts,

            ImmutableArray<Outputs.GetNsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaStatusCodeResult> statusCodes)
        {
            ClientIpAddresses = clientIpAddresses;
            Cookies = cookies;
            HttpMethods = httpMethods;
            LocationHeaders = locationHeaders;
            Paths = paths;
            ProtocolType = protocolType;
            Queries = queries;
            RequestHeaders = requestHeaders;
            ResponseHeaders = responseHeaders;
            ServicePorts = servicePorts;
            StatusCodes = statusCodes;
        }
    }
}
