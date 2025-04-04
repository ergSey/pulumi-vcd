// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class NsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Client IP Address criteria
        /// </summary>
        [Input("clientIpAddress")]
        public Input<Inputs.NsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaClientIpAddressGetArgs>? ClientIpAddress { get; set; }

        /// <summary>
        /// Criteria for matching cookie
        /// </summary>
        [Input("cookie")]
        public Input<Inputs.NsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaCookieGetArgs>? Cookie { get; set; }

        /// <summary>
        /// HTTP methods that should be matched
        /// </summary>
        [Input("httpMethods")]
        public Input<Inputs.NsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaHttpMethodsGetArgs>? HttpMethods { get; set; }

        /// <summary>
        /// A matching criteria for Location header
        /// </summary>
        [Input("locationHeader")]
        public Input<Inputs.NsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaLocationHeaderGetArgs>? LocationHeader { get; set; }

        /// <summary>
        /// Request path criteria
        /// </summary>
        [Input("path")]
        public Input<Inputs.NsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaPathGetArgs>? Path { get; set; }

        /// <summary>
        /// Protocol to match - 'HTTP' or 'HTTPS'
        /// </summary>
        [Input("protocolType")]
        public Input<string>? ProtocolType { get; set; }

        [Input("queries")]
        private InputList<string>? _queries;

        /// <summary>
        /// HTTP request query strings to match
        /// </summary>
        public InputList<string> Queries
        {
            get => _queries ?? (_queries = new InputList<string>());
            set => _queries = value;
        }

        [Input("requestHeaders")]
        private InputList<Inputs.NsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaRequestHeaderGetArgs>? _requestHeaders;

        /// <summary>
        /// A set of rules for matching request headers
        /// </summary>
        public InputList<Inputs.NsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaRequestHeaderGetArgs> RequestHeaders
        {
            get => _requestHeaders ?? (_requestHeaders = new InputList<Inputs.NsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaRequestHeaderGetArgs>());
            set => _requestHeaders = value;
        }

        [Input("responseHeaders")]
        private InputList<Inputs.NsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaResponseHeaderGetArgs>? _responseHeaders;

        /// <summary>
        /// A set of criteria to match response headers
        /// </summary>
        public InputList<Inputs.NsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaResponseHeaderGetArgs> ResponseHeaders
        {
            get => _responseHeaders ?? (_responseHeaders = new InputList<Inputs.NsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaResponseHeaderGetArgs>());
            set => _responseHeaders = value;
        }

        /// <summary>
        /// Service Port criteria
        /// </summary>
        [Input("servicePorts")]
        public Input<Inputs.NsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaServicePortsGetArgs>? ServicePorts { get; set; }

        /// <summary>
        /// HTTP Status code to match
        /// </summary>
        [Input("statusCode")]
        public Input<Inputs.NsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaStatusCodeGetArgs>? StatusCode { get; set; }

        public NsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaGetArgs()
        {
        }
        public static new NsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaGetArgs Empty => new NsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaGetArgs();
    }
}
