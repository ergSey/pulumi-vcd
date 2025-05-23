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
    public sealed class NsxtEdgegatewayStaticRouteNextHopScope
    {
        /// <summary>
        /// ID of Scope element
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Name for NSX-T Edge Gateway Static Route
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Scope type - One of 'NETWORK', 'SYSTEM_OWNED'
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private NsxtEdgegatewayStaticRouteNextHopScope(
            string id,

            string? name,

            string type)
        {
            Id = id;
            Name = name;
            Type = type;
        }
    }
}
