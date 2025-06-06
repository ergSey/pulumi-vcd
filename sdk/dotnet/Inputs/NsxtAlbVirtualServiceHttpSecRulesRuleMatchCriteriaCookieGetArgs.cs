// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class NsxtAlbVirtualServiceHttpSecRulesRuleMatchCriteriaCookieGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Criteria to use for matching cookies in the HTTP request. Options - EXISTS, DOES_NOT_EXIST, BEGINS_WITH, DOES_NOT_BEGIN_WITH, CONTAINS, DOES_NOT_CONTAIN, ENDS_WITH, DOES_NOT_END_WITH, EQUALS, DOES_NOT_EQUAL
        /// </summary>
        [Input("criteria", required: true)]
        public Input<string> Criteria { get; set; } = null!;

        /// <summary>
        /// Name of the HTTP cookie whose value is to be matched
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// String values to match for an HTTP cookie
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public NsxtAlbVirtualServiceHttpSecRulesRuleMatchCriteriaCookieGetArgs()
        {
        }
        public static new NsxtAlbVirtualServiceHttpSecRulesRuleMatchCriteriaCookieGetArgs Empty => new NsxtAlbVirtualServiceHttpSecRulesRuleMatchCriteriaCookieGetArgs();
    }
}
