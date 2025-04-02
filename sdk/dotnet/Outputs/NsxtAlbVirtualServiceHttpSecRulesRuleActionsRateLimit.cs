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
    public sealed class NsxtAlbVirtualServiceHttpSecRulesRuleActionsRateLimit
    {
        /// <summary>
        /// Set to true if the connection should be closed
        /// </summary>
        public readonly bool? ActionCloseConnection;
        /// <summary>
        /// Send custom response
        /// </summary>
        public readonly ImmutableArray<Outputs.NsxtAlbVirtualServiceHttpSecRulesRuleActionsRateLimitActionLocalResponse> ActionLocalResponses;
        /// <summary>
        /// Redirect based on rate limits
        /// </summary>
        public readonly ImmutableArray<Outputs.NsxtAlbVirtualServiceHttpSecRulesRuleActionsRateLimitActionRedirect> ActionRedirects;
        /// <summary>
        /// Maximum number of connections, requests or packets permitted each period. The count must be between 1 and 1000000000
        /// </summary>
        public readonly string Count;
        /// <summary>
        /// Time value in seconds to enforce rate count. The period must be between 1 and 1000000000
        /// </summary>
        public readonly string Period;

        [OutputConstructor]
        private NsxtAlbVirtualServiceHttpSecRulesRuleActionsRateLimit(
            bool? actionCloseConnection,

            ImmutableArray<Outputs.NsxtAlbVirtualServiceHttpSecRulesRuleActionsRateLimitActionLocalResponse> actionLocalResponses,

            ImmutableArray<Outputs.NsxtAlbVirtualServiceHttpSecRulesRuleActionsRateLimitActionRedirect> actionRedirects,

            string count,

            string period)
        {
            ActionCloseConnection = actionCloseConnection;
            ActionLocalResponses = actionLocalResponses;
            ActionRedirects = actionRedirects;
            Count = count;
            Period = period;
        }
    }
}
