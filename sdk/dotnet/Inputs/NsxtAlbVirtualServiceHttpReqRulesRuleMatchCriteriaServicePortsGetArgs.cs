// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class NsxtAlbVirtualServiceHttpReqRulesRuleMatchCriteriaServicePortsGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Criteria to use for IP address matching the HTTP request. Options - IS_IN, IS_NOT_IN
        /// </summary>
        [Input("criteria", required: true)]
        public Input<string> Criteria { get; set; } = null!;

        [Input("ports", required: true)]
        private InputList<int>? _ports;

        /// <summary>
        /// A set of TCP ports. Allowed values are 1-65535
        /// </summary>
        public InputList<int> Ports
        {
            get => _ports ?? (_ports = new InputList<int>());
            set => _ports = value;
        }

        public NsxtAlbVirtualServiceHttpReqRulesRuleMatchCriteriaServicePortsGetArgs()
        {
        }
        public static new NsxtAlbVirtualServiceHttpReqRulesRuleMatchCriteriaServicePortsGetArgs Empty => new NsxtAlbVirtualServiceHttpReqRulesRuleMatchCriteriaServicePortsGetArgs();
    }
}
