// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetCatalogVappTemplate
    {
        /// <summary>
        /// Provides a VMware Cloud Director vApp Template data source. A vApp Template can be used to reference an already existing
        /// vApp Template in VCD and use its data within other resources or data sources.
        /// 
        /// Supported in provider *v3.8+*
        /// 
        /// ## Example: Fetching a vApp Template from a Catalog
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
        ///     var my_first_vapp_template = Vcd.GetCatalogVappTemplate.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         CatalogId = my_catalog.Apply(getCatalogResult =&gt; getCatalogResult.Id),
        ///         Name = "my-first-vapp-template",
        ///     });
        /// 
        ///     var my_second_vappTemplate = new Vcd.CatalogVappTemplate("my-second-vapp_template", new()
        ///     {
        ///         MetadataEntries = .Select(entry =&gt; 
        ///         {
        ///             return new Vcd.Inputs.CatalogVappTemplateMetadataEntryArgs
        ///             {
        ///                 Key = entry.Value.Key,
        ///                 Value = entry.Value.Value,
        ///                 Type = entry.Value.Type,
        ///                 IsSystem = entry.Value.Is_system,
        ///                 UserAccess = entry.Value.User_access,
        ///             };
        ///         }).ToList(),
        ///         Org = my_first_vapp_template.Apply(my_first_vapp_template =&gt; my_first_vapp_template.Apply(getCatalogVappTemplateResult =&gt; getCatalogVappTemplateResult.Org)),
        ///         CatalogId = my_first_vapp_template.Apply(my_first_vapp_template =&gt; my_first_vapp_template.Apply(getCatalogVappTemplateResult =&gt; getCatalogVappTemplateResult.CatalogId)),
        ///         Name = "my-second-item",
        ///         Description = my_catalog.Apply(my_catalog =&gt; $"Belongs to {my_catalog.Apply(getCatalogResult =&gt; getCatalogResult.Name)}"),
        ///         OvaPath = "/path/to/test_vapp_template.ova",
        ///         UploadPieceSize = 5,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Example: Fetching a vApp Template from a VDC
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vcd = Pulumi.Vcd;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var my_vdc = Vcd.GetOrgVdc.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         Name = "my-vdc",
        ///     });
        /// 
        ///     var my_first_vapp_template = Vcd.GetCatalogVappTemplate.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         VdcId = my_vdc.Apply(getOrgVdcResult =&gt; getOrgVdcResult.Id),
        ///         Name = "my-first-vapp-template",
        ///     });
        /// 
        ///     var my_second_vappTemplate = new Vcd.CatalogVappTemplate("my-second-vapp_template", new()
        ///     {
        ///         MetadataEntries = .Select(entry =&gt; 
        ///         {
        ///             return new Vcd.Inputs.CatalogVappTemplateMetadataEntryArgs
        ///             {
        ///                 Key = entry.Value.Key,
        ///                 Value = entry.Value.Value,
        ///                 Type = entry.Value.Type,
        ///                 IsSystem = entry.Value.Is_system,
        ///                 UserAccess = entry.Value.User_access,
        ///             };
        ///         }).ToList(),
        ///         Org = my_first_vapp_template.Apply(my_first_vapp_template =&gt; my_first_vapp_template.Apply(getCatalogVappTemplateResult =&gt; getCatalogVappTemplateResult.Org)),
        ///         CatalogId = my_first_vapp_template.Apply(my_first_vapp_template =&gt; my_first_vapp_template.Apply(getCatalogVappTemplateResult =&gt; getCatalogVappTemplateResult.CatalogId)),
        ///         Name = "my-second-item",
        ///         Description = my_vdc.Apply(my_vdc =&gt; $"Belongs to {my_vdc.Apply(getOrgVdcResult =&gt; getOrgVdcResult.Name)}"),
        ///         OvaPath = "/path/to/test_vapp_template.ova",
        ///         UploadPieceSize = 5,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Filter arguments
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
        public static Task<GetCatalogVappTemplateResult> InvokeAsync(GetCatalogVappTemplateArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetCatalogVappTemplateResult>("vcd:index/getCatalogVappTemplate:getCatalogVappTemplate", args ?? new GetCatalogVappTemplateArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a VMware Cloud Director vApp Template data source. A vApp Template can be used to reference an already existing
        /// vApp Template in VCD and use its data within other resources or data sources.
        /// 
        /// Supported in provider *v3.8+*
        /// 
        /// ## Example: Fetching a vApp Template from a Catalog
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
        ///     var my_first_vapp_template = Vcd.GetCatalogVappTemplate.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         CatalogId = my_catalog.Apply(getCatalogResult =&gt; getCatalogResult.Id),
        ///         Name = "my-first-vapp-template",
        ///     });
        /// 
        ///     var my_second_vappTemplate = new Vcd.CatalogVappTemplate("my-second-vapp_template", new()
        ///     {
        ///         MetadataEntries = .Select(entry =&gt; 
        ///         {
        ///             return new Vcd.Inputs.CatalogVappTemplateMetadataEntryArgs
        ///             {
        ///                 Key = entry.Value.Key,
        ///                 Value = entry.Value.Value,
        ///                 Type = entry.Value.Type,
        ///                 IsSystem = entry.Value.Is_system,
        ///                 UserAccess = entry.Value.User_access,
        ///             };
        ///         }).ToList(),
        ///         Org = my_first_vapp_template.Apply(my_first_vapp_template =&gt; my_first_vapp_template.Apply(getCatalogVappTemplateResult =&gt; getCatalogVappTemplateResult.Org)),
        ///         CatalogId = my_first_vapp_template.Apply(my_first_vapp_template =&gt; my_first_vapp_template.Apply(getCatalogVappTemplateResult =&gt; getCatalogVappTemplateResult.CatalogId)),
        ///         Name = "my-second-item",
        ///         Description = my_catalog.Apply(my_catalog =&gt; $"Belongs to {my_catalog.Apply(getCatalogResult =&gt; getCatalogResult.Name)}"),
        ///         OvaPath = "/path/to/test_vapp_template.ova",
        ///         UploadPieceSize = 5,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Example: Fetching a vApp Template from a VDC
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vcd = Pulumi.Vcd;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var my_vdc = Vcd.GetOrgVdc.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         Name = "my-vdc",
        ///     });
        /// 
        ///     var my_first_vapp_template = Vcd.GetCatalogVappTemplate.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         VdcId = my_vdc.Apply(getOrgVdcResult =&gt; getOrgVdcResult.Id),
        ///         Name = "my-first-vapp-template",
        ///     });
        /// 
        ///     var my_second_vappTemplate = new Vcd.CatalogVappTemplate("my-second-vapp_template", new()
        ///     {
        ///         MetadataEntries = .Select(entry =&gt; 
        ///         {
        ///             return new Vcd.Inputs.CatalogVappTemplateMetadataEntryArgs
        ///             {
        ///                 Key = entry.Value.Key,
        ///                 Value = entry.Value.Value,
        ///                 Type = entry.Value.Type,
        ///                 IsSystem = entry.Value.Is_system,
        ///                 UserAccess = entry.Value.User_access,
        ///             };
        ///         }).ToList(),
        ///         Org = my_first_vapp_template.Apply(my_first_vapp_template =&gt; my_first_vapp_template.Apply(getCatalogVappTemplateResult =&gt; getCatalogVappTemplateResult.Org)),
        ///         CatalogId = my_first_vapp_template.Apply(my_first_vapp_template =&gt; my_first_vapp_template.Apply(getCatalogVappTemplateResult =&gt; getCatalogVappTemplateResult.CatalogId)),
        ///         Name = "my-second-item",
        ///         Description = my_vdc.Apply(my_vdc =&gt; $"Belongs to {my_vdc.Apply(getOrgVdcResult =&gt; getOrgVdcResult.Name)}"),
        ///         OvaPath = "/path/to/test_vapp_template.ova",
        ///         UploadPieceSize = 5,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Filter arguments
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
        public static Output<GetCatalogVappTemplateResult> Invoke(GetCatalogVappTemplateInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetCatalogVappTemplateResult>("vcd:index/getCatalogVappTemplate:getCatalogVappTemplate", args ?? new GetCatalogVappTemplateInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a VMware Cloud Director vApp Template data source. A vApp Template can be used to reference an already existing
        /// vApp Template in VCD and use its data within other resources or data sources.
        /// 
        /// Supported in provider *v3.8+*
        /// 
        /// ## Example: Fetching a vApp Template from a Catalog
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
        ///     var my_first_vapp_template = Vcd.GetCatalogVappTemplate.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         CatalogId = my_catalog.Apply(getCatalogResult =&gt; getCatalogResult.Id),
        ///         Name = "my-first-vapp-template",
        ///     });
        /// 
        ///     var my_second_vappTemplate = new Vcd.CatalogVappTemplate("my-second-vapp_template", new()
        ///     {
        ///         MetadataEntries = .Select(entry =&gt; 
        ///         {
        ///             return new Vcd.Inputs.CatalogVappTemplateMetadataEntryArgs
        ///             {
        ///                 Key = entry.Value.Key,
        ///                 Value = entry.Value.Value,
        ///                 Type = entry.Value.Type,
        ///                 IsSystem = entry.Value.Is_system,
        ///                 UserAccess = entry.Value.User_access,
        ///             };
        ///         }).ToList(),
        ///         Org = my_first_vapp_template.Apply(my_first_vapp_template =&gt; my_first_vapp_template.Apply(getCatalogVappTemplateResult =&gt; getCatalogVappTemplateResult.Org)),
        ///         CatalogId = my_first_vapp_template.Apply(my_first_vapp_template =&gt; my_first_vapp_template.Apply(getCatalogVappTemplateResult =&gt; getCatalogVappTemplateResult.CatalogId)),
        ///         Name = "my-second-item",
        ///         Description = my_catalog.Apply(my_catalog =&gt; $"Belongs to {my_catalog.Apply(getCatalogResult =&gt; getCatalogResult.Name)}"),
        ///         OvaPath = "/path/to/test_vapp_template.ova",
        ///         UploadPieceSize = 5,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Example: Fetching a vApp Template from a VDC
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vcd = Pulumi.Vcd;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var my_vdc = Vcd.GetOrgVdc.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         Name = "my-vdc",
        ///     });
        /// 
        ///     var my_first_vapp_template = Vcd.GetCatalogVappTemplate.Invoke(new()
        ///     {
        ///         Org = "my-org",
        ///         VdcId = my_vdc.Apply(getOrgVdcResult =&gt; getOrgVdcResult.Id),
        ///         Name = "my-first-vapp-template",
        ///     });
        /// 
        ///     var my_second_vappTemplate = new Vcd.CatalogVappTemplate("my-second-vapp_template", new()
        ///     {
        ///         MetadataEntries = .Select(entry =&gt; 
        ///         {
        ///             return new Vcd.Inputs.CatalogVappTemplateMetadataEntryArgs
        ///             {
        ///                 Key = entry.Value.Key,
        ///                 Value = entry.Value.Value,
        ///                 Type = entry.Value.Type,
        ///                 IsSystem = entry.Value.Is_system,
        ///                 UserAccess = entry.Value.User_access,
        ///             };
        ///         }).ToList(),
        ///         Org = my_first_vapp_template.Apply(my_first_vapp_template =&gt; my_first_vapp_template.Apply(getCatalogVappTemplateResult =&gt; getCatalogVappTemplateResult.Org)),
        ///         CatalogId = my_first_vapp_template.Apply(my_first_vapp_template =&gt; my_first_vapp_template.Apply(getCatalogVappTemplateResult =&gt; getCatalogVappTemplateResult.CatalogId)),
        ///         Name = "my-second-item",
        ///         Description = my_vdc.Apply(my_vdc =&gt; $"Belongs to {my_vdc.Apply(getOrgVdcResult =&gt; getOrgVdcResult.Name)}"),
        ///         OvaPath = "/path/to/test_vapp_template.ova",
        ///         UploadPieceSize = 5,
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// ## Filter arguments
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
        public static Output<GetCatalogVappTemplateResult> Invoke(GetCatalogVappTemplateInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetCatalogVappTemplateResult>("vcd:index/getCatalogVappTemplate:getCatalogVappTemplate", args ?? new GetCatalogVappTemplateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetCatalogVappTemplateArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the catalog containing the vApp Template. Can't be used if a specific VDC identifier is set (`vdc_id`).
        /// </summary>
        [Input("catalogId")]
        public string? CatalogId { get; set; }

        /// <summary>
        /// Retrieves the data source using one or more filter parameters
        /// </summary>
        [Input("filter")]
        public Inputs.GetCatalogVappTemplateFilterArgs? Filter { get; set; }

        /// <summary>
        /// vApp Template name (optional when `filter` is used)
        /// </summary>
        [Input("name")]
        public string? Name { get; set; }

        /// <summary>
        /// Org name
        /// </summary>
        [Input("org")]
        public string? Org { get; set; }

        /// <summary>
        /// ID of the VDC to which the vApp Template belongs. Can't be used if a specific Catalog is set (`catalog_id`).
        /// </summary>
        [Input("vdcId")]
        public string? VdcId { get; set; }

        public GetCatalogVappTemplateArgs()
        {
        }
        public static new GetCatalogVappTemplateArgs Empty => new GetCatalogVappTemplateArgs();
    }

    public sealed class GetCatalogVappTemplateInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// ID of the catalog containing the vApp Template. Can't be used if a specific VDC identifier is set (`vdc_id`).
        /// </summary>
        [Input("catalogId")]
        public Input<string>? CatalogId { get; set; }

        /// <summary>
        /// Retrieves the data source using one or more filter parameters
        /// </summary>
        [Input("filter")]
        public Input<Inputs.GetCatalogVappTemplateFilterInputArgs>? Filter { get; set; }

        /// <summary>
        /// vApp Template name (optional when `filter` is used)
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Org name
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// ID of the VDC to which the vApp Template belongs. Can't be used if a specific Catalog is set (`catalog_id`).
        /// </summary>
        [Input("vdcId")]
        public Input<string>? VdcId { get; set; }

        public GetCatalogVappTemplateInvokeArgs()
        {
        }
        public static new GetCatalogVappTemplateInvokeArgs Empty => new GetCatalogVappTemplateInvokeArgs();
    }


    [OutputType]
    public sealed class GetCatalogVappTemplateResult
    {
        public readonly string? CatalogId;
        public readonly string CatalogItemId;
        public readonly string Created;
        /// <summary>
        /// vApp Template description
        /// </summary>
        public readonly string Description;
        public readonly Outputs.GetCatalogVappTemplateFilterResult? Filter;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableDictionary<string, string> InheritedMetadata;
        /// <summary>
        /// (*v3.11+*) - The information about the vApp Template lease. It includes the following field:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCatalogVappTemplateLeaseResult> Leases;
        /// <summary>
        /// (Deprecated) Use `metadata_entry` instead. Key/value map of metadata for the associated vApp template.
        /// </summary>
        public readonly ImmutableDictionary<string, string> Metadata;
        /// <summary>
        /// A set of metadata entries assigned to this vApp Template. See [Metadata](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/catalog_vapp_template#metadata) section for details.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetCatalogVappTemplateMetadataEntryResult> MetadataEntries;
        public readonly string? Name;
        public readonly string? Org;
        public readonly string? VdcId;
        /// <summary>
        /// Set of VM names within the vApp template
        /// </summary>
        public readonly ImmutableArray<string> VmNames;

        [OutputConstructor]
        private GetCatalogVappTemplateResult(
            string? catalogId,

            string catalogItemId,

            string created,

            string description,

            Outputs.GetCatalogVappTemplateFilterResult? filter,

            string id,

            ImmutableDictionary<string, string> inheritedMetadata,

            ImmutableArray<Outputs.GetCatalogVappTemplateLeaseResult> leases,

            ImmutableDictionary<string, string> metadata,

            ImmutableArray<Outputs.GetCatalogVappTemplateMetadataEntryResult> metadataEntries,

            string? name,

            string? org,

            string? vdcId,

            ImmutableArray<string> vmNames)
        {
            CatalogId = catalogId;
            CatalogItemId = catalogItemId;
            Created = created;
            Description = description;
            Filter = filter;
            Id = id;
            InheritedMetadata = inheritedMetadata;
            Leases = leases;
            Metadata = metadata;
            MetadataEntries = metadataEntries;
            Name = name;
            Org = org;
            VdcId = vdcId;
            VmNames = vmNames;
        }
    }
}
