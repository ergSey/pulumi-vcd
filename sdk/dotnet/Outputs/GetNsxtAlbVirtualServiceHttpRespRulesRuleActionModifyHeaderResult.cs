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
    public sealed class GetNsxtAlbVirtualServiceHttpRespRulesRuleActionModifyHeaderResult
    {
        /// <summary>
        /// One of the following HTTP header actions
        /// </summary>
        public readonly string Action;
        /// <summary>
        /// HTTP header name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// HTTP header value
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetNsxtAlbVirtualServiceHttpRespRulesRuleActionModifyHeaderResult(
            string action,

            string name,

            string value)
        {
            Action = action;
            Name = name;
            Value = value;
        }
    }
}
