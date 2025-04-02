// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class NsxtAlbVirtualServiceHttpReqRulesRuleMatchCriteriaClientIpAddressArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Criteria to use for IP address matching the HTTP request. Options - IS_IN, IS_NOT_IN.
        /// </summary>
        [Input("criteria", required: true)]
        public Input<string> Criteria { get; set; } = null!;

        [Input("ipAddresses", required: true)]
        private InputList<string>? _ipAddresses;

        /// <summary>
        /// A set of IP addresses
        /// </summary>
        public InputList<string> IpAddresses
        {
            get => _ipAddresses ?? (_ipAddresses = new InputList<string>());
            set => _ipAddresses = value;
        }

        public NsxtAlbVirtualServiceHttpReqRulesRuleMatchCriteriaClientIpAddressArgs()
        {
        }
        public static new NsxtAlbVirtualServiceHttpReqRulesRuleMatchCriteriaClientIpAddressArgs Empty => new NsxtAlbVirtualServiceHttpReqRulesRuleMatchCriteriaClientIpAddressArgs();
    }
}
