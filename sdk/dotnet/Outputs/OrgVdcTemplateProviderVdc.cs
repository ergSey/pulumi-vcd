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
    public sealed class OrgVdcTemplateProviderVdc
    {
        /// <summary>
        /// ID of the Provider Gateway to use, can be obtained with
        /// [`vcd.ExternalNetworkV2` data source](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/external_network_v2)
        /// </summary>
        public readonly string ExternalNetworkId;
        /// <summary>
        /// ID of the Edge Cluster that the VDCs instantiated from this template will use with the Edge Gateway.
        /// Can be obtained with [`vcd.getNsxtEdgeCluster` data source](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_edge_cluster).
        /// If set, a `edge_gateway` block **must** be present in the VDC Template configuration (see below).
        /// </summary>
        public readonly string? GatewayEdgeClusterId;
        /// <summary>
        /// ID of the Provider VDC, can be obtained with
        /// [`vcd.ProviderVdc` data source](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/provider_vdc)
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// ID of the Edge Cluster that the VDCs instantiated from this template will use for services.
        /// Can be obtained with [`vcd.getNsxtEdgeCluster` data source](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_edge_cluster)
        /// </summary>
        public readonly string? ServicesEdgeClusterId;

        [OutputConstructor]
        private OrgVdcTemplateProviderVdc(
            string externalNetworkId,

            string? gatewayEdgeClusterId,

            string id,

            string? servicesEdgeClusterId)
        {
            ExternalNetworkId = externalNetworkId;
            GatewayEdgeClusterId = gatewayEdgeClusterId;
            Id = id;
            ServicesEdgeClusterId = servicesEdgeClusterId;
        }
    }
}
