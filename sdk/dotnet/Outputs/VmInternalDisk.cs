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
    public sealed class VmInternalDisk
    {
        /// <summary>
        /// The number of the SCSI or IDE controller itself.
        /// </summary>
        public readonly int? BusNumber;
        /// <summary>
        /// The type of disk controller. Possible values: ide, parallel( LSI Logic Parallel SCSI), sas(LSI Logic SAS (SCSI)), paravirtual(Paravirtual (SCSI)), sata, nvme
        /// </summary>
        public readonly string? BusType;
        /// <summary>
        /// The disk ID.
        /// </summary>
        public readonly string? DiskId;
        /// <summary>
        /// Specifies the IOPS for the disk. Default is 0.
        /// </summary>
        public readonly int? Iops;
        /// <summary>
        /// The size of the disk in MB.
        /// </summary>
        public readonly int? SizeInMb;
        /// <summary>
        /// Storage profile to override the VM default one
        /// </summary>
        public readonly string? StorageProfile;
        /// <summary>
        /// Specifies whether the disk storage is pre-allocated or allocated on demand.
        /// </summary>
        public readonly bool? ThinProvisioned;
        /// <summary>
        /// The device number on the SCSI or IDE controller of the disk.
        /// </summary>
        public readonly int? UnitNumber;

        [OutputConstructor]
        private VmInternalDisk(
            int? busNumber,

            string? busType,

            string? diskId,

            int? iops,

            int? sizeInMb,

            string? storageProfile,

            bool? thinProvisioned,

            int? unitNumber)
        {
            BusNumber = busNumber;
            BusType = busType;
            DiskId = diskId;
            Iops = iops;
            SizeInMb = sizeInMb;
            StorageProfile = storageProfile;
            ThinProvisioned = thinProvisioned;
            UnitNumber = unitNumber;
        }
    }
}
