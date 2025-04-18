// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/independentDisk:IndependentDisk")]
    public partial class IndependentDisk : global::Pulumi.CustomResource
    {
        /// <summary>
        /// (Computed, *v3.6+*) IDs of VM which are using the disk
        /// </summary>
        [Output("attachedVmIds")]
        public Output<ImmutableArray<string>> AttachedVmIds { get; private set; } = null!;

        /// <summary>
        /// Disk bus subtype. Values can be: `buslogic`, `lsilogic`, `lsilogicsas`, `VirtualSCSI` for `SCSI`, `ahci` for `SATA` and (*v3.6+*) `nvmecontroller` for `NVME`
        /// </summary>
        [Output("busSubType")]
        public Output<string> BusSubType { get; private set; } = null!;

        /// <summary>
        /// Disk bus type. Values can be: `IDE`, `SCSI`, `SATA`, (*v3.6+*) `NVME`. **Note** When the disk type is IDE then VM is required to be powered off
        /// </summary>
        [Output("busType")]
        public Output<string> BusType { get; private set; } = null!;

        /// <summary>
        /// (Computed) Data store name. Readable only for system user.
        /// </summary>
        [Output("datastoreName")]
        public Output<string> DatastoreName { get; private set; } = null!;

        /// <summary>
        /// independent disk description
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// (Computed, *v3.6+* and VCD 10.2+) True if disk is encrypted
        /// </summary>
        [Output("encrypted")]
        public Output<bool> Encrypted { get; private set; } = null!;

        /// <summary>
        /// (Computed) IOPS request for the created disk
        /// </summary>
        [Output("iops")]
        public Output<int> Iops { get; private set; } = null!;

        /// <summary>
        /// (Computed) True if the disk is already attached
        /// </summary>
        [Output("isAttached")]
        public Output<bool> IsAttached { get; private set; } = null!;

        /// <summary>
        /// Use `metadata_entry` instead. Key value map of metadata to assign to this independent disk.
        /// </summary>
        [Output("metadata")]
        public Output<ImmutableDictionary<string, string>> Metadata { get; private set; } = null!;

        /// <summary>
        /// A set of metadata entries to assign. See Metadata section for details.
        /// </summary>
        [Output("metadataEntries")]
        public Output<ImmutableArray<Outputs.IndependentDiskMetadataEntry>> MetadataEntries { get; private set; } = null!;

        /// <summary>
        /// Disk name
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
        /// </summary>
        [Output("org")]
        public Output<string?> Org { get; private set; } = null!;

        /// <summary>
        /// (Computed) The owner name of the disk
        /// </summary>
        [Output("ownerName")]
        public Output<string> OwnerName { get; private set; } = null!;

        /// <summary>
        /// This is the sharing type. Values can be: `DiskSharing`,`ControllerSharing`, or `None`
        /// </summary>
        [Output("sharingType")]
        public Output<string> SharingType { get; private set; } = null!;

        /// <summary>
        /// Size of disk in MB.
        /// </summary>
        [Output("sizeInMb")]
        public Output<int> SizeInMb { get; private set; } = null!;

        /// <summary>
        /// The name of storage profile where disk will be created
        /// </summary>
        [Output("storageProfile")]
        public Output<string> StorageProfile { get; private set; } = null!;

        /// <summary>
        /// (Computed, *v3.6+* and VCD 10.2+) The UUID of this named disk's device backing
        /// </summary>
        [Output("uuid")]
        public Output<string> Uuid { get; private set; } = null!;

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Output("vdc")]
        public Output<string?> Vdc { get; private set; } = null!;


        /// <summary>
        /// Create a IndependentDisk resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public IndependentDisk(string name, IndependentDiskArgs args, CustomResourceOptions? options = null)
            : base("vcd:index/independentDisk:IndependentDisk", name, args ?? new IndependentDiskArgs(), MakeResourceOptions(options, ""))
        {
        }

        private IndependentDisk(string name, Input<string> id, IndependentDiskState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/independentDisk:IndependentDisk", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing IndependentDisk resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static IndependentDisk Get(string name, Input<string> id, IndependentDiskState? state = null, CustomResourceOptions? options = null)
        {
            return new IndependentDisk(name, id, state, options);
        }
    }

    public sealed class IndependentDiskArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Disk bus subtype. Values can be: `buslogic`, `lsilogic`, `lsilogicsas`, `VirtualSCSI` for `SCSI`, `ahci` for `SATA` and (*v3.6+*) `nvmecontroller` for `NVME`
        /// </summary>
        [Input("busSubType")]
        public Input<string>? BusSubType { get; set; }

        /// <summary>
        /// Disk bus type. Values can be: `IDE`, `SCSI`, `SATA`, (*v3.6+*) `NVME`. **Note** When the disk type is IDE then VM is required to be powered off
        /// </summary>
        [Input("busType")]
        public Input<string>? BusType { get; set; }

        /// <summary>
        /// independent disk description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// Use `metadata_entry` instead. Key value map of metadata to assign to this independent disk.
        /// </summary>
        [Obsolete(@"Use metadata_entry instead")]
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        [Input("metadataEntries")]
        private InputList<Inputs.IndependentDiskMetadataEntryArgs>? _metadataEntries;

        /// <summary>
        /// A set of metadata entries to assign. See Metadata section for details.
        /// </summary>
        public InputList<Inputs.IndependentDiskMetadataEntryArgs> MetadataEntries
        {
            get => _metadataEntries ?? (_metadataEntries = new InputList<Inputs.IndependentDiskMetadataEntryArgs>());
            set => _metadataEntries = value;
        }

        /// <summary>
        /// Disk name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// This is the sharing type. Values can be: `DiskSharing`,`ControllerSharing`, or `None`
        /// </summary>
        [Input("sharingType")]
        public Input<string>? SharingType { get; set; }

        /// <summary>
        /// Size of disk in MB.
        /// </summary>
        [Input("sizeInMb", required: true)]
        public Input<int> SizeInMb { get; set; } = null!;

        /// <summary>
        /// The name of storage profile where disk will be created
        /// </summary>
        [Input("storageProfile")]
        public Input<string>? StorageProfile { get; set; }

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public IndependentDiskArgs()
        {
        }
        public static new IndependentDiskArgs Empty => new IndependentDiskArgs();
    }

    public sealed class IndependentDiskState : global::Pulumi.ResourceArgs
    {
        [Input("attachedVmIds")]
        private InputList<string>? _attachedVmIds;

        /// <summary>
        /// (Computed, *v3.6+*) IDs of VM which are using the disk
        /// </summary>
        public InputList<string> AttachedVmIds
        {
            get => _attachedVmIds ?? (_attachedVmIds = new InputList<string>());
            set => _attachedVmIds = value;
        }

        /// <summary>
        /// Disk bus subtype. Values can be: `buslogic`, `lsilogic`, `lsilogicsas`, `VirtualSCSI` for `SCSI`, `ahci` for `SATA` and (*v3.6+*) `nvmecontroller` for `NVME`
        /// </summary>
        [Input("busSubType")]
        public Input<string>? BusSubType { get; set; }

        /// <summary>
        /// Disk bus type. Values can be: `IDE`, `SCSI`, `SATA`, (*v3.6+*) `NVME`. **Note** When the disk type is IDE then VM is required to be powered off
        /// </summary>
        [Input("busType")]
        public Input<string>? BusType { get; set; }

        /// <summary>
        /// (Computed) Data store name. Readable only for system user.
        /// </summary>
        [Input("datastoreName")]
        public Input<string>? DatastoreName { get; set; }

        /// <summary>
        /// independent disk description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// (Computed, *v3.6+* and VCD 10.2+) True if disk is encrypted
        /// </summary>
        [Input("encrypted")]
        public Input<bool>? Encrypted { get; set; }

        /// <summary>
        /// (Computed) IOPS request for the created disk
        /// </summary>
        [Input("iops")]
        public Input<int>? Iops { get; set; }

        /// <summary>
        /// (Computed) True if the disk is already attached
        /// </summary>
        [Input("isAttached")]
        public Input<bool>? IsAttached { get; set; }

        [Input("metadata")]
        private InputMap<string>? _metadata;

        /// <summary>
        /// Use `metadata_entry` instead. Key value map of metadata to assign to this independent disk.
        /// </summary>
        [Obsolete(@"Use metadata_entry instead")]
        public InputMap<string> Metadata
        {
            get => _metadata ?? (_metadata = new InputMap<string>());
            set => _metadata = value;
        }

        [Input("metadataEntries")]
        private InputList<Inputs.IndependentDiskMetadataEntryGetArgs>? _metadataEntries;

        /// <summary>
        /// A set of metadata entries to assign. See Metadata section for details.
        /// </summary>
        public InputList<Inputs.IndependentDiskMetadataEntryGetArgs> MetadataEntries
        {
            get => _metadataEntries ?? (_metadataEntries = new InputList<Inputs.IndependentDiskMetadataEntryGetArgs>());
            set => _metadataEntries = value;
        }

        /// <summary>
        /// Disk name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
        /// </summary>
        [Input("org")]
        public Input<string>? Org { get; set; }

        /// <summary>
        /// (Computed) The owner name of the disk
        /// </summary>
        [Input("ownerName")]
        public Input<string>? OwnerName { get; set; }

        /// <summary>
        /// This is the sharing type. Values can be: `DiskSharing`,`ControllerSharing`, or `None`
        /// </summary>
        [Input("sharingType")]
        public Input<string>? SharingType { get; set; }

        /// <summary>
        /// Size of disk in MB.
        /// </summary>
        [Input("sizeInMb")]
        public Input<int>? SizeInMb { get; set; }

        /// <summary>
        /// The name of storage profile where disk will be created
        /// </summary>
        [Input("storageProfile")]
        public Input<string>? StorageProfile { get; set; }

        /// <summary>
        /// (Computed, *v3.6+* and VCD 10.2+) The UUID of this named disk's device backing
        /// </summary>
        [Input("uuid")]
        public Input<string>? Uuid { get; set; }

        /// <summary>
        /// The name of VDC to use, optional if defined at provider level
        /// </summary>
        [Input("vdc")]
        public Input<string>? Vdc { get; set; }

        public IndependentDiskState()
        {
        }
        public static new IndependentDiskState Empty => new IndependentDiskState();
    }
}
