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
    public sealed class GetNsxtAlbVirtualServiceHttpReqRulesRuleMatchCriteriaServicePortResult
    {
        /// <summary>
        /// Criteria to use for service port matching the HTTP request
        /// </summary>
        public readonly string Criteria;
        /// <summary>
        /// A set of TCP ports
        /// </summary>
        public readonly ImmutableArray<int> Ports;

        [OutputConstructor]
        private GetNsxtAlbVirtualServiceHttpReqRulesRuleMatchCriteriaServicePortResult(
            string criteria,

            ImmutableArray<int> ports)
        {
            Criteria = criteria;
            Ports = ports;
        }
    }
}
