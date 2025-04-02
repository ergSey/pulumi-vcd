// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class NetworkRoutedMetadataEntryGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Domain for this metadata entry. true, if it belongs to SYSTEM. false, if it belongs to GENERAL
        /// </summary>
        [Input("isSystem")]
        public Input<bool>? IsSystem { get; set; }

        /// <summary>
        /// Key of this metadata entry. Required if the metadata entry is not empty
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        /// <summary>
        /// Type of this metadata entry. One of: 'MetadataStringValue', 'MetadataNumberValue', 'MetadataBooleanValue', 'MetadataDateTimeValue'
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// User access level for this metadata entry. One of: 'READWRITE', 'READONLY', 'PRIVATE'
        /// </summary>
        [Input("userAccess")]
        public Input<string>? UserAccess { get; set; }

        /// <summary>
        /// Value of this metadata entry. Required if the metadata entry is not empty
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public NetworkRoutedMetadataEntryGetArgs()
        {
        }
        public static new NetworkRoutedMetadataEntryGetArgs Empty => new NetworkRoutedMetadataEntryGetArgs();
    }
}
