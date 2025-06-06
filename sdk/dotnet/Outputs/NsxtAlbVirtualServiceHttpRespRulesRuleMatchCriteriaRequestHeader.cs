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
    public sealed class NsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaRequestHeader
    {
        /// <summary>
        /// Criteria to use for matching headers and cookies in the HTTP request amd response. Options - EXISTS, DOES_NOT_EXIST, BEGINS_WITH, DOES_NOT_BEGIN_WITH, CONTAINS, DOES_NOT_CONTAIN, ENDS_WITH, DOES_NOT_END_WITH, EQUALS, DOES_NOT_EQUAL
        /// </summary>
        public readonly string Criteria;
        /// <summary>
        /// Name of the HTTP header whose value is to be matched
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// String values to match for an HTTP header
        /// </summary>
        public readonly ImmutableArray<string> Values;

        [OutputConstructor]
        private NsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaRequestHeader(
            string criteria,

            string name,

            ImmutableArray<string> values)
        {
            Criteria = criteria;
            Name = name;
            Values = values;
        }
    }
}
