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
    public sealed class GetNetworkDirectFilterMetadataResult
    {
        /// <summary>
        /// True if is a metadata@SYSTEM key
        /// </summary>
        public readonly bool? IsSystem;
        /// <summary>
        /// Metadata key (field name)
        /// </summary>
        public readonly string Key;
        /// <summary>
        /// Type of metadata value (needed only if "use_api_search" is true)
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// If true, will search the vCD using native metadata query (without regular expressions)
        /// </summary>
        public readonly bool? UseApiSearch;
        /// <summary>
        /// Metadata value (can be a regular expression if "use_api_search" is false)
        /// </summary>
        public readonly string Value;

        [OutputConstructor]
        private GetNetworkDirectFilterMetadataResult(
            bool? isSystem,

            string key,

            string? type,

            bool? useApiSearch,

            string value)
        {
            IsSystem = isSystem;
            Key = key;
            Type = type;
            UseApiSearch = useApiSearch;
            Value = value;
        }
    }
}
