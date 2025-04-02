// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetRde
    {
        /// <summary>
        /// Provides the capability of reading an existing Runtime Defined Entity in VMware Cloud Director.
        /// 
        /// &gt; VCD allows to have multiple RDEs of the same [RDE Type](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/rde_type) with
        /// the same name, meaning that the data source will not be able to fetch a RDE in this situation, as this data source
        /// can only retrieve **unique RDEs**.
        /// 
        /// Supported in provider *v3.9+*
        /// </summary>
        public static Task<GetRdeResult> InvokeAsync(GetRdeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetRdeResult>("vcd:index/getRde:getRde", args ?? new GetRdeArgs(), options.WithDefaults());

        /// <summary>
        /// Provides the capability of reading an existing Runtime Defined Entity in VMware Cloud Director.
        /// 
        /// &gt; VCD allows to have multiple RDEs of the same [RDE Type](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/rde_type) with
        /// the same name, meaning that the data source will not be able to fetch a RDE in this situation, as this data source
        /// can only retrieve **unique RDEs**.
        /// 
        /// Supported in provider *v3.9+*
        /// </summary>
        public static Output<GetRdeResult> Invoke(GetRdeInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetRdeResult>("vcd:index/getRde:getRde", args ?? new GetRdeInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Provides the capability of reading an existing Runtime Defined Entity in VMware Cloud Director.
        /// 
        /// &gt; VCD allows to have multiple RDEs of the same [RDE Type](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/rde_type) with
        /// the same name, meaning that the data source will not be able to fetch a RDE in this situation, as this data source
        /// can only retrieve **unique RDEs**.
        /// 
        /// Supported in provider *v3.9+*
        /// </summary>
        public static Output<GetRdeResult> Invoke(GetRdeInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetRdeResult>("vcd:index/getRde:getRde", args ?? new GetRdeInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetRdeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Runtime Defined Entity.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// Name of the [Organization](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/org) that owns the RDE, optional if defined at provider level.
        /// </summary>
        [Input("org")]
        public string? Org { get; set; }

        /// <summary>
        /// The ID of the [RDE Type](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/rde_type) of the RDE to fetch.
        /// </summary>
        [Input("rdeTypeId", required: true)]
        public string RdeTypeId { get; set; } = null!;

        public GetRdeArgs()
        {
        }
        public static new GetRdeArgs Empty => new GetRdeArgs();
    }

    public sealed class GetRdeInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the Runtime Defined Entity.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Name of the [Organization](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/org) that owns the RDE, optional if defined at provider level.
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// The ID of the [RDE Type](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/rde_type) of the RDE to fetch.
        /// </summary>
        [Input("rdeTypeId", required: true)]
        public Input<string> RdeTypeId { get; set; } = null!;

        public GetRdeInvokeArgs()
        {
        }
        public static new GetRdeInvokeArgs Empty => new GetRdeInvokeArgs();
    }


    [OutputType]
    public sealed class GetRdeResult
    {
        /// <summary>
        /// The entity JSON.
        /// </summary>
        public readonly string Entity;
        public readonly string ExternalId;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// A set of metadata entries that belong to the RDE.
        /// Read the [resource](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/rde#metadata) documentation for the details of the sub-attributes.
        /// </summary>
        public readonly ImmutableArray<Outputs.GetRdeMetadataEntryResult> MetadataEntries;
        public readonly string Name;
        public readonly string? Org;
        /// <summary>
        /// The ID of the [Organization](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/org) to which the Runtime Defined Entity belongs.
        /// </summary>
        public readonly string OrgId;
        /// <summary>
        /// The ID of the [Organization user](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/org_user) that owns this Runtime Defined Entity.
        /// </summary>
        public readonly string OwnerUserId;
        public readonly string RdeTypeId;
        /// <summary>
        /// It can be `RESOLVED`, `RESOLUTION_ERROR` or `PRE_CREATED`.
        /// </summary>
        public readonly string State;

        [OutputConstructor]
        private GetRdeResult(
            string entity,

            string externalId,

            string id,

            ImmutableArray<Outputs.GetRdeMetadataEntryResult> metadataEntries,

            string name,

            string? org,

            string orgId,

            string ownerUserId,

            string rdeTypeId,

            string state)
        {
            Entity = entity;
            ExternalId = externalId;
            Id = id;
            MetadataEntries = metadataEntries;
            Name = name;
            Org = org;
            OrgId = orgId;
            OwnerUserId = ownerUserId;
            RdeTypeId = rdeTypeId;
            State = state;
        }
    }
}
