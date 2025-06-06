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
    public sealed class GetIndependentDiskMetadataEntryResult
    {
        /// <summary>
        /// Domain for this metadata entry. true, if it belongs to SYSTEM. false, if it belongs to GENERAL
        /// </summary>
        public readonly bool IsSystem;
        /// <summary>
        /// Key of this metadata entry
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// Type of this metadata entry. One of: 'MetadataStringValue', 'MetadataNumberValue', 'MetadataBooleanValue', 'MetadataDateTimeValue'
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// User access level for this metadata entry. One of: 'READWRITE', 'READONLY', 'PRIVATE'
        /// </summary>
        public readonly string UserAccess;
        /// <summary>
        /// Value of this metadata entry
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetIndependentDiskMetadataEntryResult(
            bool isSystem,

            string key,

            string type,

            string userAccess,

            string value)
        {
            IsSystem = isSystem;
            Key = key;
            Type = type;
            UserAccess = userAccess;
            Value = value;
        }
    }
}
