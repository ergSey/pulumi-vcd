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
    public sealed class GetSubscribedCatalogFilterResult
    {
        /// <summary>
        /// Search by date comparison ({&gt;|&gt;=|&lt;|&lt;=|==} yyyy-mm-dd[ hh[:mm[:ss]]])
        /// </summary>
        public readonly string? Date;
        /// <summary>
        /// Retrieves the oldest item
        /// </summary>
        public readonly bool? Earliest;
        /// <summary>
        /// Retrieves the newest item
        /// </summary>
        public readonly bool? Latest;
        /// <summary>
        /// Optional metadata of the catalog. This is inherited from the publishing catalog
        /// </summary>
        public readonly ImmutableArray<Outputs.GetSubscribedCatalogFilterMetadataResult> Metadatas;
        /// <summary>
        /// Search by name with a regular expression
        /// </summary>
        public readonly string? NameRegex;

        [OutputConstructor]
        private GetSubscribedCatalogFilterResult(
            string? date,

            bool? earliest,

            bool? latest,

            ImmutableArray<Outputs.GetSubscribedCatalogFilterMetadataResult> metadatas,

            string? nameRegex)
        {
            Date = date;
            Earliest = earliest;
            Latest = latest;
            Metadatas = metadatas;
            NameRegex = nameRegex;
        }
    }
}
