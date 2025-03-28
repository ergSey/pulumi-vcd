// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetCatalogMedia
    {
        /// <summary>
        /// Provides a VMware Cloud Director Catalog media data source. A Catalog media can be used to reference a catalog media and use its 
        /// data within other resources or data sources.
        /// 
        /// Supported in provider *v2.5+*
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vcd = Pulumi.Vcd;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var my_catalog = Vcd.GetCatalog.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         Name = "my-catalog",
        ///     });
        /// 
        ///     var existing_media = Vcd.GetCatalogMedia.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         CatalogId = my_catalog.Apply(getCatalogResult =&gt; getCatalogResult.Id),
        ///         Name = "my-media",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["mediaSize"] = existing_media.Apply(existing_media =&gt; existing_media.Apply(getCatalogMediaResult =&gt; getCatalogMediaResult.Size)),
        ///         ["typeIsIso"] = existing_media.Apply(existing_media =&gt; existing_media.Apply(getCatalogMediaResult =&gt; getCatalogMediaResult.IsIso)),
        ///     };
        /// });
        /// ```
        /// 
        /// ## Filter arguments
        /// 
        /// (Supported in provider *v2.9+*)
        /// 
        /// * `name_regex` - (Optional) matches the name using a regular expression.
        /// * `date` - (Optional) is an expression starting with an operator (`&gt;`, `&lt;`, `&gt;=`, `&lt;=`, `==`), followed by a date, with
        ///   optional spaces in between. For example: `&gt; 2020-02-01 12:35:00.523Z`
        ///   The filter recognizes several formats, but one of `yyyy-mm-dd [hh:mm[:ss[.nnnZ]]]` or `dd-MMM-yyyy [hh:mm[:ss[.nnnZ]]]`
        ///   is recommended.
        ///   Comparison with equality operator (`==`) need to define the date to the microseconds.
        /// * `latest` - (Optional) If `true`, retrieve the latest item among the ones matching other parameters. If no other parameters
        ///   are set, it retrieves the newest item.
        /// * `earliest` - (Optional) If `true`, retrieve the earliest item among the ones matching other parameters. If no other parameters
        ///   are set, it retrieves the oldest item.
        /// * `metadata` - (Optional) One or more parameters that will match metadata contents.
        /// 
        /// See [Filters reference](https://www.terraform.io/providers/vmware/vcd/latest/docs/guides/data_source_filters) for details and examples.
        /// </summary>
        public static Task<GetCatalogMediaResult> InvokeAsync(GetCatalogMediaArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCatalogMediaResult>("vcd:index/getCatalogMedia:getCatalogMedia", args ?? new GetCatalogMediaArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a VMware Cloud Director Catalog media data source. A Catalog media can be used to reference a catalog media and use its 
        /// data within other resources or data sources.
        /// 
        /// Supported in provider *v2.5+*
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vcd = Pulumi.Vcd;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var my_catalog = Vcd.GetCatalog.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         Name = "my-catalog",
        ///     });
        /// 
        ///     var existing_media = Vcd.GetCatalogMedia.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         CatalogId = my_catalog.Apply(getCatalogResult =&gt; getCatalogResult.Id),
        ///         Name = "my-media",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["mediaSize"] = existing_media.Apply(existing_media =&gt; existing_media.Apply(getCatalogMediaResult =&gt; getCatalogMediaResult.Size)),
        ///         ["typeIsIso"] = existing_media.Apply(existing_media =&gt; existing_media.Apply(getCatalogMediaResult =&gt; getCatalogMediaResult.IsIso)),
        ///     };
        /// });
        /// ```
        /// 
        /// ## Filter arguments
        /// 
        /// (Supported in provider *v2.9+*)
        /// 
        /// * `name_regex` - (Optional) matches the name using a regular expression.
        /// * `date` - (Optional) is an expression starting with an operator (`&gt;`, `&lt;`, `&gt;=`, `&lt;=`, `==`), followed by a date, with
        ///   optional spaces in between. For example: `&gt; 2020-02-01 12:35:00.523Z`
        ///   The filter recognizes several formats, but one of `yyyy-mm-dd [hh:mm[:ss[.nnnZ]]]` or `dd-MMM-yyyy [hh:mm[:ss[.nnnZ]]]`
        ///   is recommended.
        ///   Comparison with equality operator (`==`) need to define the date to the microseconds.
        /// * `latest` - (Optional) If `true`, retrieve the latest item among the ones matching other parameters. If no other parameters
        ///   are set, it retrieves the newest item.
        /// * `earliest` - (Optional) If `true`, retrieve the earliest item among the ones matching other parameters. If no other parameters
        ///   are set, it retrieves the oldest item.
        /// * `metadata` - (Optional) One or more parameters that will match metadata contents.
        /// 
        /// See [Filters reference](https://www.terraform.io/providers/vmware/vcd/latest/docs/guides/data_source_filters) for details and examples.
        /// </summary>
        public static Output<GetCatalogMediaResult> Invoke(GetCatalogMediaInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCatalogMediaResult>("vcd:index/getCatalogMedia:getCatalogMedia", args ?? new GetCatalogMediaInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a VMware Cloud Director Catalog media data source. A Catalog media can be used to reference a catalog media and use its 
        /// data within other resources or data sources.
        /// 
        /// Supported in provider *v2.5+*
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vcd = Pulumi.Vcd;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var my_catalog = Vcd.GetCatalog.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         Name = "my-catalog",
        ///     });
        /// 
        ///     var existing_media = Vcd.GetCatalogMedia.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         CatalogId = my_catalog.Apply(getCatalogResult =&gt; getCatalogResult.Id),
        ///         Name = "my-media",
        ///     });
        /// 
        ///     return new Dictionary&lt;string, object?&gt;
        ///     {
        ///         ["mediaSize"] = existing_media.Apply(existing_media =&gt; existing_media.Apply(getCatalogMediaResult =&gt; getCatalogMediaResult.Size)),
        ///         ["typeIsIso"] = existing_media.Apply(existing_media =&gt; existing_media.Apply(getCatalogMediaResult =&gt; getCatalogMediaResult.IsIso)),
        ///     };
        /// });
        /// ```
        /// 
        /// ## Filter arguments
        /// 
        /// (Supported in provider *v2.9+*)
        /// 
        /// * `name_regex` - (Optional) matches the name using a regular expression.
        /// * `date` - (Optional) is an expression starting with an operator (`&gt;`, `&lt;`, `&gt;=`, `&lt;=`, `==`), followed by a date, with
        ///   optional spaces in between. For example: `&gt; 2020-02-01 12:35:00.523Z`
        ///   The filter recognizes several formats, but one of `yyyy-mm-dd [hh:mm[:ss[.nnnZ]]]` or `dd-MMM-yyyy [hh:mm[:ss[.nnnZ]]]`
        ///   is recommended.
        ///   Comparison with equality operator (`==`) need to define the date to the microseconds.
        /// * `latest` - (Optional) If `true`, retrieve the latest item among the ones matching other parameters. If no other parameters
        ///   are set, it retrieves the newest item.
        /// * `earliest` - (Optional) If `true`, retrieve the earliest item among the ones matching other parameters. If no other parameters
        ///   are set, it retrieves the oldest item.
        /// * `metadata` - (Optional) One or more parameters that will match metadata contents.
        /// 
        /// See [Filters reference](https://www.terraform.io/providers/vmware/vcd/latest/docs/guides/data_source_filters) for details and examples.
        /// </summary>
        public static Output<GetCatalogMediaResult> Invoke(GetCatalogMediaInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetCatalogMediaResult>("vcd:index/getCatalogMedia:getCatalogMedia", args ?? new GetCatalogMediaInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCatalogMediaArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the catalog to which media file belongs. It's mandatory if `catalog_id` is not used.
        /// </summary>
        [Input("catalog")]
        public string? Catalog { get; set; }

        /// <summary>
        /// The ID of the catalog to which the media file belongs. It's mandatory if `catalog` field is not used.
        /// </summary>
        [Input("catalogId")]
        public string? CatalogId { get; set; }

        [Input("downloadToFile")]
        public string? DownloadToFile { get; set; }

        /// <summary>
        /// Retrieves the data source using one or more filter parameters
        /// </summary>
        [Input("filter")]
        public Inputs.GetCatalogMediaFilterArgs? Filter { get; set; }

        /// <summary>
        /// Media name in catalog (optional when `filter` is used)
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level
        /// </summary>
        [Input("org")]
        public string? Org { get; set; }

        public GetCatalogMediaArgs()
        {
        }
        public static new GetCatalogMediaArgs Empty => new GetCatalogMediaArgs();
    }

    public sealed class GetCatalogMediaInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the catalog to which media file belongs. It's mandatory if `catalog_id` is not used.
        /// </summary>
        [Input("catalog")]
        public Input<string>? Catalog { get; set; }

        /// <summary>
        /// The ID of the catalog to which the media file belongs. It's mandatory if `catalog` field is not used.
        /// </summary>
        [Input("catalogId")]
        public Input<string>? CatalogId { get; set; }

        [Input("downloadToFile")]
        public Input<string>? DownloadToFile { get; set; }

        /// <summary>
        /// Retrieves the data source using one or more filter parameters
        /// </summary>
        [Input("filter")]
        public Input<Inputs.GetCatalogMediaFilterInputArgs>? Filter { get; set; }

        /// <summary>
        /// Media name in catalog (optional when `filter` is used)
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        public GetCatalogMediaInvokeArgs()
        {
        }
        public static new GetCatalogMediaInvokeArgs Empty => new GetCatalogMediaInvokeArgs();
    }


    [OutputType]
    public sealed class GetCatalogMediaResult
    {
        public readonly string Catalog;
        public readonly string? CatalogId;
        public readonly string CatalogItemId;
        public readonly string CreationDate;
        public readonly string Description;
        public readonly string? DownloadToFile;
        public readonly Outputs.GetCatalogMediaFilterResult? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly bool IsIso;
        public readonly bool IsPublished;
        public readonly ImmutableDictionary<string, string> Metadata;
        public readonly ImmutableArray<Outputs.GetCatalogMediaMetadataEntryResult> MetadataEntries;
        public readonly string? Name;
        public readonly string? Org;
        public readonly string OwnerName;
        public readonly int Size;
        public readonly string Status;
        public readonly string StorageProfileName;

        [OutputConstructor]
        private GetCatalogMediaResult(
            string catalog,

            string? catalogId,

            string catalogItemId,

            string creationDate,

            string description,

            string? downloadToFile,

            Outputs.GetCatalogMediaFilterResult? filter,

            string id,

            bool isIso,

            bool isPublished,

            ImmutableDictionary<string, string> metadata,

            ImmutableArray<Outputs.GetCatalogMediaMetadataEntryResult> metadataEntries,

            string? name,

            string? org,

            string ownerName,

            int size,

            string status,

            string storageProfileName)
        {
            Catalog = catalog;
            CatalogId = catalogId;
            CatalogItemId = catalogItemId;
            CreationDate = creationDate;
            Description = description;
            DownloadToFile = downloadToFile;
            Filter = filter;
            Id = id;
            IsIso = isIso;
            IsPublished = isPublished;
            Metadata = metadata;
            MetadataEntries = metadataEntries;
            Name = name;
            Org = org;
            OwnerName = ownerName;
            Size = size;
            Status = status;
            StorageProfileName = storageProfileName;
        }
    }
}
