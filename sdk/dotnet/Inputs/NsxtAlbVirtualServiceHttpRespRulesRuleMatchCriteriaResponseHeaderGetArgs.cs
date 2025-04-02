// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class NsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaResponseHeaderGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Criteria to use for matching headers and cookies in the HTTP request amd response. Options - EXISTS, DOES_NOT_EXIST, BEGINS_WITH, DOES_NOT_BEGIN_WITH, CONTAINS, DOES_NOT_CONTAIN, ENDS_WITH, DOES_NOT_END_WITH, EQUALS, DOES_NOT_EQUAL
        /// </summary>
        [Input("criteria", required: true)]
        public Input<string> Criteria { get; set; } = null!;

        /// <summary>
        /// Name of the HTTP header whose value is to be matched
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("values")]
        private InputList<string>? _values;

        /// <summary>
        /// A set of values to match for an HTTP header
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public NsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaResponseHeaderGetArgs()
        {
        }
        public static new NsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaResponseHeaderGetArgs Empty => new NsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteriaResponseHeaderGetArgs();
    }
}
