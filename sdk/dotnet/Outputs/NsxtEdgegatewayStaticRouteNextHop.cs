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
    public sealed class NsxtEdgegatewayStaticRouteNextHop
    {
        /// <summary>
        /// Admin distance of next hop
        /// </summary>
        public readonly int AdminDistance;
        /// <summary>
        /// IP Address of next hop
        /// </summary>
        public readonly string IpAddress;
        public readonly Outputs.NsxtEdgegatewayStaticRouteNextHopScope? Scope;

        [OutputConstructor]
        private NsxtEdgegatewayStaticRouteNextHop(
            int adminDistance,

            string ipAddress,

            Outputs.NsxtEdgegatewayStaticRouteNextHopScope? scope)
        {
            AdminDistance = adminDistance;
            IpAddress = ipAddress;
            Scope = scope;
        }
    }
}
