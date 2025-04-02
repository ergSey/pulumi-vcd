// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class NsxtAlbVirtualServiceHttpSecRulesRuleActionsRateLimitGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Set to true if the connection should be closed
        /// </summary>
        [Input("actionCloseConnection")]
        public Input<bool>? ActionCloseConnection { get; set; }

        [Input("actionLocalResponses")]
        private InputList<Inputs.NsxtAlbVirtualServiceHttpSecRulesRuleActionsRateLimitActionLocalResponseGetArgs>? _actionLocalResponses;

        /// <summary>
        /// Send custom response
        /// </summary>
        public InputList<Inputs.NsxtAlbVirtualServiceHttpSecRulesRuleActionsRateLimitActionLocalResponseGetArgs> ActionLocalResponses
        {
            get => _actionLocalResponses ?? (_actionLocalResponses = new InputList<Inputs.NsxtAlbVirtualServiceHttpSecRulesRuleActionsRateLimitActionLocalResponseGetArgs>());
            set => _actionLocalResponses = value;
        }

        [Input("actionRedirects")]
        private InputList<Inputs.NsxtAlbVirtualServiceHttpSecRulesRuleActionsRateLimitActionRedirectGetArgs>? _actionRedirects;

        /// <summary>
        /// Redirect based on rate limits
        /// </summary>
        public InputList<Inputs.NsxtAlbVirtualServiceHttpSecRulesRuleActionsRateLimitActionRedirectGetArgs> ActionRedirects
        {
            get => _actionRedirects ?? (_actionRedirects = new InputList<Inputs.NsxtAlbVirtualServiceHttpSecRulesRuleActionsRateLimitActionRedirectGetArgs>());
            set => _actionRedirects = value;
        }

        /// <summary>
        /// Maximum number of connections, requests or packets permitted each period. The count must be between 1 and 1000000000
        /// </summary>
        [Input("count", required: true)]
        public Input<string> Count { get; set; } = null!;

        /// <summary>
        /// Time value in seconds to enforce rate count. The period must be between 1 and 1000000000
        /// </summary>
        [Input("period", required: true)]
        public Input<string> Period { get; set; } = null!;

        public NsxtAlbVirtualServiceHttpSecRulesRuleActionsRateLimitGetArgs()
        {
        }
        public static new NsxtAlbVirtualServiceHttpSecRulesRuleActionsRateLimitGetArgs Empty => new NsxtAlbVirtualServiceHttpSecRulesRuleActionsRateLimitGetArgs();
    }
}
