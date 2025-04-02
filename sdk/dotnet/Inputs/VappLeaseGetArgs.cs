// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class VappLeaseGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// How long any of the VMs in the vApp can run before the vApp is automatically powered off or suspended. 0 means never expires (or maximum allowed by Org). Regular values accepted from 3600+.
        /// </summary>
        [Input("runtimeLeaseInSec", required: true)]
        public Input<int> RuntimeLeaseInSec { get; set; } = null!;

        /// <summary>
        /// How long the vApp is available before being automatically deleted or marked as expired. 0 means never expires (or maximum allowed by Org). Regular values accepted from 3600+.
        /// </summary>
        [Input("storageLeaseInSec", required: true)]
        public Input<int> StorageLeaseInSec { get; set; } = null!;

        public VappLeaseGetArgs()
        {
        }
        public static new VappLeaseGetArgs Empty => new VappLeaseGetArgs();
    }
}
