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
    public sealed class DseRegistryConfigurationContainerRegistry
    {
        /// <summary>
        /// Registry description
        /// </summary>
        public readonly string Description;
        /// <summary>
        /// Registry host
        /// </summary>
        public readonly string Host;
        /// <summary>
        /// Password for registry user
        /// </summary>
        public readonly string? Password;
        /// <summary>
        /// Username for registry access
        /// </summary>
        public readonly string? Username;

        [OutputConstructor]
        private DseRegistryConfigurationContainerRegistry(
            string description,

            string host,

            string? password,

            string? username)
        {
            Description = description;
            Host = host;
            Password = password;
            Username = username;
        }
    }
}
