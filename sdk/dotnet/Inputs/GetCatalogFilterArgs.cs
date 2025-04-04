// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class GetCatalogFilterInputArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Search by date comparison ({&gt;|&gt;=|&lt;|&lt;=|==} yyyy-mm-dd[ hh[:mm[:ss]]])
        /// </summary>
        [Input("date")]
        public Input<string>? Date { get; set; }

        /// <summary>
        /// Retrieves the oldest item
        /// </summary>
        [Input("earliest")]
        public Input<bool>? Earliest { get; set; }

        /// <summary>
        /// Retrieves the newest item
        /// </summary>
        [Input("latest")]
        public Input<bool>? Latest { get; set; }

        [Input("metadatas")]
        private InputList<Inputs.GetCatalogFilterMetadataInputArgs>? _metadatas;

        /// <summary>
        /// (Deprecated; *v3.6+*) Use `metadata_entry` instead. Key value map of metadata.
        /// </summary>
        public InputList<Inputs.GetCatalogFilterMetadataInputArgs> Metadatas
        {
            get => _metadatas ?? (_metadatas = new InputList<Inputs.GetCatalogFilterMetadataInputArgs>());
            set => _metadatas = value;
        }

        /// <summary>
        /// Search by name with a regular expression
        /// </summary>
        [Input("nameRegex")]
        public Input<string>? NameRegex { get; set; }

        public GetCatalogFilterInputArgs()
        {
        }
        public static new GetCatalogFilterInputArgs Empty => new GetCatalogFilterInputArgs();
    }
}
