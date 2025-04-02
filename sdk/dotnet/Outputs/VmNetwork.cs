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
    public sealed class VmNetwork
    {
        /// <summary>
        /// Network card adapter type. (e.g. 'E1000', 'E1000E', 'SRIOVETHERNETCARD', 'VMXNET3', 'PCNet32')
        /// </summary>
        public readonly string? AdapterType;
        /// <summary>
        /// It defines if NIC is connected or not.
        /// </summary>
        public readonly bool? Connected;
        /// <summary>
        /// IP of the VM. Settings depend on `ip_allocation_mode`. Omitted or empty for DHCP, POOL, NONE. Required for MANUAL
        /// </summary>
        public readonly string? Ip;
        /// <summary>
        /// IP address allocation mode. One of POOL, DHCP, MANUAL, NONE
        /// </summary>
        public readonly string IpAllocationMode;
        /// <summary>
        /// Set to true if network interface should be primary. First network card in the list will be primary by default
        /// </summary>
        public readonly bool? IsPrimary;
        /// <summary>
        /// Mac address of network interface
        /// </summary>
        public readonly string? Mac;
        /// <summary>
        /// Name of the network this VM should connect to. Always required except for `type` `NONE`
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// Secondary (IPv6) IP of the VM. Settings depend on `secondary_ip_allocation_mode`. Omitted or empty for DHCP, POOL, NONE. Required for MANUAL
        /// </summary>
        public readonly string? SecondaryIp;
        /// <summary>
        /// Secondary (IPv6) IP address allocation mode. One of POOL, DHCP, MANUAL, NONE
        /// </summary>
        public readonly string? SecondaryIpAllocationMode;
        /// <summary>
        /// Network type to use: 'vapp', 'org' or 'none'. Use 'vapp' for vApp network, 'org' to attach Org VDC network. 'none' for empty NIC.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private VmNetwork(
            string? adapterType,

            bool? connected,

            string? ip,

            string ipAllocationMode,

            bool? isPrimary,

            string? mac,

            string? name,

            string? secondaryIp,

            string? secondaryIpAllocationMode,

            string type)
        {
            AdapterType = adapterType;
            Connected = connected;
            Ip = ip;
            IpAllocationMode = ipAllocationMode;
            IsPrimary = isPrimary;
            Mac = mac;
            Name = name;
            SecondaryIp = secondaryIp;
            SecondaryIpAllocationMode = secondaryIpAllocationMode;
            Type = type;
        }
    }
}
