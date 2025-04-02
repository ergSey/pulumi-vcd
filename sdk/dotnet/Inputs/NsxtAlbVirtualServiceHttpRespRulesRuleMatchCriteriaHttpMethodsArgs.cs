// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class NsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaHttpMethodsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Criteria to use for IP address matching the HTTP request. Options - IS_IN, IS_NOT_IN
        /// </summary>
        [Input("criteria", required: true)]
        public Input<string> Criteria { get; set; } = null!;

        [Input("methods", required: true)]
        private InputList<string>? _methods;

        /// <summary>
        /// HTTP methods to match. Options - GET, PUT, POST, DELETE, HEAD, OPTIONS, TRACE, CONNECT, PATCH, PROPFIND, PROPPATCH, MKCOL, COPY, MOVE, LOCK, UNLOCK
        /// </summary>
        public InputList<string> Methods
        {
            get => _methods ?? (_methods = new InputList<string>());
            set => _methods = value;
        }

        public NsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaHttpMethodsArgs()
        {
        }
        public static new NsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaHttpMethodsArgs Empty => new NsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaHttpMethodsArgs();
    }
}
