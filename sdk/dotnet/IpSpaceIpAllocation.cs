// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/ipSpaceIpAllocation:IpSpaceIpAllocation")]
    public partial class IpSpaceIpAllocation : global::Pulumi.CustomResource
    {
        /// <summary>
        /// allocation date in formated as `2023-06-07T09:57:58.721Z` (ISO 8601)
        /// </summary>
        [Output("allocationDate")]
        public Output<string> AllocationDate { get; private set; } = null!;

        /// <summary>
        /// Can only be set when `usage_state=USED_MANUAL`
        /// 
        /// &gt; IP Allocation resources can be created only if there is a NSX-T Edge Gateway
        /// (`vcd.NsxtEdgegateway`) that is backed by the Provider Gateway (`vcd.ExternalNetworkV2`) with IP
        /// Space Uplinks (`vcd.IpSpaceUplink`). Attempting to allocate IP Addresses before having an
        /// Edge Gateway withing VDC will return errors of type `This operation is denied`.
        /// </summary>
        [Output("description")]
        public Output<string> Description { get; private set; } = null!;

        /// <summary>
        /// convenience field. For `type=IP_PREFIX` it will contain only the IP from CIDR returned
        /// </summary>
        [Output("ip")]
        public Output<string> Ip { get; private set; } = null!;

        /// <summary>
        /// IP address or CIDR
        /// </summary>
        [Output("ipAddress")]
        public Output<string> IpAddress { get; private set; } = null!;

        /// <summary>
        /// IP Space ID to use for IP Allocations
        /// </summary>
        [Output("ipSpaceId")]
        public Output<string?> IpSpaceId { get; private set; } = null!;

        /// <summary>
        /// Org ID in which the IP is allocated
        /// </summary>
        [Output("orgId")]
        public Output<string> OrgId { get; private set; } = null!;

        /// <summary>
        /// Required when `type=IP_PREFIX`
        /// </summary>
        [Output("prefixLength")]
        public Output<string> PrefixLength { get; private set; } = null!;

        /// <summary>
        /// One of `FLOATING_IP`, `IP_PREFIX`
        /// * `FLOATING_IP` - allocates single IP from defined ranges in IP Space
        /// * `IP_PREFIX` - allocates subnets. **Note** field `prefix_length` is required to allocate IP
        /// Prefix
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// (Optional) Only used with manual reservations. Value `USED_MANUAL`
        /// enables manual IP reservation. Value `UNUSED` is set to release manual allocation of IP.
        /// </summary>
        [Output("usageState")]
        public Output<string> UsageState { get; private set; } = null!;

        /// <summary>
        /// contains entity ID that is using the IP if `usage_state=USED`
        /// </summary>
        [Output("usedById")]
        public Output<string> UsedById { get; private set; } = null!;

        /// <summary>
        /// An option to request a specific IP or subnet from IP Space.
        /// **Note:** This field does not support IP ranges because it would cause multiple allocations
        /// created in one resource. Please use multiple resource instances to allocate IP ranges.
        /// </summary>
        [Output("value")]
        public Output<string?> Value { get; private set; } = null!;


        /// <summary>
        /// Create a IpSpaceIpAllocation resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IpSpaceIpAllocation(string name, IpSpaceIpAllocationArgs args, CustomResourceOptions? options = null)
            : base("vcd:index/ipSpaceIpAllocation:IpSpaceIpAllocation", name, args ?? new IpSpaceIpAllocationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IpSpaceIpAllocation(string name, Input<string> id, IpSpaceIpAllocationState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/ipSpaceIpAllocation:IpSpaceIpAllocation", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/vmware/terraform-provider-vcd/releases/download/v3.14.1/",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing IpSpaceIpAllocation resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IpSpaceIpAllocation Get(string name, Input<string> id, IpSpaceIpAllocationState? state = null, CustomResourceOptions? options = null)
        {
            return new IpSpaceIpAllocation(name, id, state, options);
        }
    }

    public sealed class IpSpaceIpAllocationArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Can only be set when `usage_state=USED_MANUAL`
        /// 
        /// &gt; IP Allocation resources can be created only if there is a NSX-T Edge Gateway
        /// (`vcd.NsxtEdgegateway`) that is backed by the Provider Gateway (`vcd.ExternalNetworkV2`) with IP
        /// Space Uplinks (`vcd.IpSpaceUplink`). Attempting to allocate IP Addresses before having an
        /// Edge Gateway withing VDC will return errors of type `This operation is denied`.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// IP Space ID to use for IP Allocations
        /// </summary>
        [Input("ipSpaceId")]
        public Input<string>? IpSpaceId { get; set; }

        /// <summary>
        /// Org ID in which the IP is allocated
        /// </summary>
        [Input("orgId", required: true)]
        public Input<string> OrgId { get; set; } = null!;

        /// <summary>
        /// Required when `type=IP_PREFIX`
        /// </summary>
        [Input("prefixLength")]
        public Input<string>? PrefixLength { get; set; }

        /// <summary>
        /// One of `FLOATING_IP`, `IP_PREFIX`
        /// * `FLOATING_IP` - allocates single IP from defined ranges in IP Space
        /// * `IP_PREFIX` - allocates subnets. **Note** field `prefix_length` is required to allocate IP
        /// Prefix
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// (Optional) Only used with manual reservations. Value `USED_MANUAL`
        /// enables manual IP reservation. Value `UNUSED` is set to release manual allocation of IP.
        /// </summary>
        [Input("usageState")]
        public Input<string>? UsageState { get; set; }

        /// <summary>
        /// An option to request a specific IP or subnet from IP Space.
        /// **Note:** This field does not support IP ranges because it would cause multiple allocations
        /// created in one resource. Please use multiple resource instances to allocate IP ranges.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public IpSpaceIpAllocationArgs()
        {
        }
        public static new IpSpaceIpAllocationArgs Empty => new IpSpaceIpAllocationArgs();
    }

    public sealed class IpSpaceIpAllocationState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// allocation date in formated as `2023-06-07T09:57:58.721Z` (ISO 8601)
        /// </summary>
        [Input("allocationDate")]
        public Input<string>? AllocationDate { get; set; }

        /// <summary>
        /// Can only be set when `usage_state=USED_MANUAL`
        /// 
        /// &gt; IP Allocation resources can be created only if there is a NSX-T Edge Gateway
        /// (`vcd.NsxtEdgegateway`) that is backed by the Provider Gateway (`vcd.ExternalNetworkV2`) with IP
        /// Space Uplinks (`vcd.IpSpaceUplink`). Attempting to allocate IP Addresses before having an
        /// Edge Gateway withing VDC will return errors of type `This operation is denied`.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// convenience field. For `type=IP_PREFIX` it will contain only the IP from CIDR returned
        /// </summary>
        [Input("ip")]
        public Input<string>? Ip { get; set; }

        /// <summary>
        /// IP address or CIDR
        /// </summary>
        [Input("ipAddress")]
        public Input<string>? IpAddress { get; set; }

        /// <summary>
        /// IP Space ID to use for IP Allocations
        /// </summary>
        [Input("ipSpaceId")]
        public Input<string>? IpSpaceId { get; set; }

        /// <summary>
        /// Org ID in which the IP is allocated
        /// </summary>
        [Input("orgId")]
        public Input<string>? OrgId { get; set; }

        /// <summary>
        /// Required when `type=IP_PREFIX`
        /// </summary>
        [Input("prefixLength")]
        public Input<string>? PrefixLength { get; set; }

        /// <summary>
        /// One of `FLOATING_IP`, `IP_PREFIX`
        /// * `FLOATING_IP` - allocates single IP from defined ranges in IP Space
        /// * `IP_PREFIX` - allocates subnets. **Note** field `prefix_length` is required to allocate IP
        /// Prefix
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// (Optional) Only used with manual reservations. Value `USED_MANUAL`
        /// enables manual IP reservation. Value `UNUSED` is set to release manual allocation of IP.
        /// </summary>
        [Input("usageState")]
        public Input<string>? UsageState { get; set; }

        /// <summary>
        /// contains entity ID that is using the IP if `usage_state=USED`
        /// </summary>
        [Input("usedById")]
        public Input<string>? UsedById { get; set; }

        /// <summary>
        /// An option to request a specific IP or subnet from IP Space.
        /// **Note:** This field does not support IP ranges because it would cause multiple allocations
        /// created in one resource. Please use multiple resource instances to allocate IP ranges.
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public IpSpaceIpAllocationState()
        {
        }
        public static new IpSpaceIpAllocationState Empty => new IpSpaceIpAllocationState();
    }
}
