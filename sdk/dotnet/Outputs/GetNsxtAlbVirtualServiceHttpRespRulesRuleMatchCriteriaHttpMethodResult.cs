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
    public sealed class GetNsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaHttpMethodResult
    {
        /// <summary>
        /// Criteria to use for HTTP methods matching the request
        /// </summary>
        public readonly string Criteria;
        /// <summary>
        /// HTTP methods to match
        /// </summary>
        public readonly ImmutableArray<string> Methods;

        [OutputConstructor]
        private GetNsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaHttpMethodResult(
            string criteria,

            ImmutableArray<string> methods)
        {
            Criteria = criteria;
            Methods = methods;
        }
    }
}
