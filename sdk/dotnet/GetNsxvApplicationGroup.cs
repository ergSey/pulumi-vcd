// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetNsxvApplicationGroup
    {
        /// <summary>
        /// Provides a VMware Cloud Director NSX-V Distributed Firewall data source used to read an existing application group.
        /// 
        /// Supported in provider *v3.9+*
        /// </summary>
        public static Task<GetNsxvApplicationGroupResult> InvokeAsync(GetNsxvApplicationGroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNsxvApplicationGroupResult>("vcd:index/getNsxvApplicationGroup:getNsxvApplicationGroup", args ?? new GetNsxvApplicationGroupArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a VMware Cloud Director NSX-V Distributed Firewall data source used to read an existing application group.
        /// 
        /// Supported in provider *v3.9+*
        /// </summary>
        public static Output<GetNsxvApplicationGroupResult> Invoke(GetNsxvApplicationGroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNsxvApplicationGroupResult>("vcd:index/getNsxvApplicationGroup:getNsxvApplicationGroup", args ?? new GetNsxvApplicationGroupInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a VMware Cloud Director NSX-V Distributed Firewall data source used to read an existing application group.
        /// 
        /// Supported in provider *v3.9+*
        /// </summary>
        public static Output<GetNsxvApplicationGroupResult> Invoke(GetNsxvApplicationGroupInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetNsxvApplicationGroupResult>("vcd:index/getNsxvApplicationGroup:getNsxvApplicationGroup", args ?? new GetNsxvApplicationGroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetNsxvApplicationGroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the application group
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// The ID of VDC to use
        /// </summary>
        [Input("vdcId", required: true)]
        public string VdcId { get; set; } = null!;

        public GetNsxvApplicationGroupArgs()
        {
        }
        public static new GetNsxvApplicationGroupArgs Empty => new GetNsxvApplicationGroupArgs();
    }

    public sealed class GetNsxvApplicationGroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// The name of the application group
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// The ID of VDC to use
        /// </summary>
        [Input("vdcId", required: true)]
        public Input<string> VdcId { get; set; } = null!;

        public GetNsxvApplicationGroupInvokeArgs()
        {
        }
        public static new GetNsxvApplicationGroupInvokeArgs Empty => new GetNsxvApplicationGroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetNsxvApplicationGroupResult
    {
        /// <summary>
        /// The list of the applications belonging to this group. For each one we get the following:
        /// </summary>
        public readonly ImmutableArray<Outputs.GetNsxvApplicationGroupApplicationResult> Applications;
        /// <summary>
        /// The identifier of the application groups
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// The name of the application
        /// </summary>
        public readonly string Name;
        public readonly string VdcId;

        [OutputConstructor]
        private GetNsxvApplicationGroupResult(
            ImmutableArray<Outputs.GetNsxvApplicationGroupApplicationResult> applications,

            string id,

            string name,

            string vdcId)
        {
            Applications = applications;
            Id = id;
            Name = name;
            VdcId = vdcId;
        }
    }
}
