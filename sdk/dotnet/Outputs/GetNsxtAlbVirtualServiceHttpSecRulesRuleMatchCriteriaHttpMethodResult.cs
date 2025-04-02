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
    public sealed class GetNsxtAlbVirtualServiceHttpSecRulesRuleMatchCriteriaHttpMethodResult
    {
        /// <summary>
        /// Criteria to use for HTTP method matching in the HTTP request
        /// </summary>
        public readonly string Criteria;
        /// <summary>
        /// HTTP methods that will be matched
        /// </summary>
        public readonly ImmutableArray<string> Methods;

        [OutputConstructor]
        private GetNsxtAlbVirtualServiceHttpSecRulesRuleMatchCriteriaHttpMethodResult(
            string criteria,

            ImmutableArray<string> methods)
        {
            Criteria = criteria;
            Methods = methods;
        }
    }
}
