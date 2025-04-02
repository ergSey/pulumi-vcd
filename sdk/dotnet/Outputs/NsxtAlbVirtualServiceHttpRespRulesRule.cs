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
    public sealed class NsxtAlbVirtualServiceHttpRespRulesRule
    {
        /// <summary>
        /// Actions to perform with the rule that matches
        /// </summary>
        public readonly Outputs.NsxtAlbVirtualServiceHttpRespRulesRuleActions Actions;
        /// <summary>
        /// Defines if the rule is active or not
        /// </summary>
        public readonly bool? Active;
        /// <summary>
        /// Defines whether to enable logging with headers on rule match or not
        /// </summary>
        public readonly bool? Logging;
        /// <summary>
        /// Rule matching Criteria
        /// </summary>
        public readonly Outputs.NsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteria MatchCriteria;
        /// <summary>
        /// Name of the rule
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private NsxtAlbVirtualServiceHttpRespRulesRule(
            Outputs.NsxtAlbVirtualServiceHttpRespRulesRuleActions actions,

            bool? active,

            bool? logging,

            Outputs.NsxtAlbVirtualServiceHttpRespRulesRuleMatchCriteria matchCriteria,

            string name)
        {
            Actions = actions;
            Active = active;
            Logging = logging;
            MatchCriteria = matchCriteria;
            Name = name;
        }
    }
}
