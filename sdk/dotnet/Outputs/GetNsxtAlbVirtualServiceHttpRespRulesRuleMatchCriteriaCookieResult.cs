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
    public sealed class GetNsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaCookieResult
    {
        /// <summary>
        /// Criteria to use for matching cookies in the HTTP request
        /// </summary>
        public readonly string Criteria;
        /// <summary>
        /// Name of the HTTP cookie whose value is to be matched
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// String values to match for an HTTP cookie
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetNsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaCookieResult(
            string criteria,

            string name,

            string value)
        {
            Criteria = criteria;
            Name = name;
            Value = value;
        }
    }
}
