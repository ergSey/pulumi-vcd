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
    public sealed class NetworkPoolBackingPortGroup
    {
        /// <summary>
        /// (Computed) The ID of the backing element
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Unique name of network pool
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Type of the network pool (one of `GENEVE`, `VLAN`, `PORTGROUP_BACKED`)
        /// </summary>
        public readonly string? Type;

        [OutputConstructor]
        private NetworkPoolBackingPortGroup(
            string? id,

            string? name,

            string? type)
        {
            Id = id;
            Name = name;
            Type = type;
        }
    }
}
