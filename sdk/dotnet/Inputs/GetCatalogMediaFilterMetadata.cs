// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class GetCatalogMediaFilterMetadataArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// True if is a metadata@SYSTEM key
        /// </summary>
        [Input("isSystem")]
        public bool? IsSystem { get; set; }

        /// <summary>
        /// Metadata key (field name)
        /// </summary>
        [Input("key", required: true)]
        public string Key { get; set; } = null!;

        /// <summary>
        /// Type of metadata value (needed only if "use_api_search" is true)
        /// </summary>
        [Input("type")]
        public string? Type { get; set; }

        /// <summary>
        /// If true, will search the vCD using native metadata query (without regular expressions)
        /// </summary>
        [Input("useApiSearch")]
        public bool? UseApiSearch { get; set; }

        /// <summary>
        /// Metadata value (can be a regular expression if "use_api_search" is false)
        /// </summary>
        [Input("value", required: true)]
        public string Value { get; set; } = null!;

        public GetCatalogMediaFilterMetadataArgs()
        {
        }
        public static new GetCatalogMediaFilterMetadataArgs Empty => new GetCatalogMediaFilterMetadataArgs();
    }
}
