// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class GetNetworkDirectFilterInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Search by IP. The value can be a regular expression
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        [Input("metadatas")]
        private InputList<Inputs.GetNetworkDirectFilterMetadataInputArgs>? _metadatas;

        /// <summary>
        /// metadata filter
        /// </summary>
        public InputList<Inputs.GetNetworkDirectFilterMetadataInputArgs> Metadatas
        {
            get => _metadatas ?? (_metadatas = new InputList<Inputs.GetNetworkDirectFilterMetadataInputArgs>());
            set => _metadatas = value;
        }

        /// <summary>
        /// Search by name with a regular expression
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        public GetNetworkDirectFilterInputArgs()
        {
        }
        public static new GetNetworkDirectFilterInputArgs Empty => new GetNetworkDirectFilterInputArgs();
    }
}
