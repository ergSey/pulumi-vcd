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
    public sealed class GetNsxtAlbVirtualServiceHttpSecRulesRuleResult
    {
        /// <summary>
        /// Actions to perform with the rule that matches
        /// </summary>
        public readonly ImmutableArray<Outputs.GetNsxtAlbVirtualServiceHttpSecRulesRuleActionResult> Actions;
        /// <summary>
        /// Defines is the rule is active or not
        /// </summary>
        public readonly bool Active;
        /// <summary>
        /// Defines whether to enable logging with headers on rule match or not
        /// </summary>
        public readonly bool Logging;
        /// <summary>
        /// Rule matching Criteria
        /// </summary>
        public readonly ImmutableArray<Outputs.GetNsxtAlbVirtualServiceHttpSecRulesRuleMatchCriteriaResult> MatchCriterias;
        /// <summary>
        /// Name of the rule
        /// </summary>
        public readonly string Name;

        [OutputConstructor]
        private GetNsxtAlbVirtualServiceHttpSecRulesRuleResult(
            ImmutableArray<Outputs.GetNsxtAlbVirtualServiceHttpSecRulesRuleActionResult> actions,

            bool active,

            bool logging,

            ImmutableArray<Outputs.GetNsxtAlbVirtualServiceHttpSecRulesRuleMatchCriteriaResult> matchCriterias,

            string name)
        {
            Actions = actions;
            Active = active;
            Logging = logging;
            MatchCriterias = matchCriterias;
            Name = name;
        }
    }
}
