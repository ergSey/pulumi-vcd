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
    public sealed class OrgVdcStorageProfile
    {
        /// <summary>
        /// True if this is default storage profile for this VDC. The default storage profile is used when an object that can specify a storage profile is created with no storage profile specified.
        /// </summary>
        public readonly bool Default;
        /// <summary>
        /// True if this VDC is enabled for use by the organization VDCs. Default is true.
        /// </summary>
        public readonly bool? Enabled;
        /// <summary>
        /// Maximum number of MB allocated for this storage profile. A value of 0 specifies unlimited MB.
        /// </summary>
        public readonly int Limit;
        /// <summary>
        /// VDC name
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Storage used in MB
        /// </summary>
        public readonly int? StorageUsedInMb;

        [OutputConstructor]
        private OrgVdcStorageProfile(
            bool @default,

            bool? enabled,

            int limit,

            string name,

            int? storageUsedInMb)
        {
            Default = @default;
            Enabled = enabled;
            Limit = limit;
            Name = name;
            StorageUsedInMb = storageUsedInMb;
        }
    }
}
