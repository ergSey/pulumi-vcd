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
    public sealed class VmBootOptions
    {
        /// <summary>
        /// Number of milliseconds to wait between powering-on and booting the VM
        /// </summary>
        public readonly int? BootDelay;
        /// <summary>
        /// Delay in milliseconds before a boot retry. Only works if 'boot_retry_enabled' is set to true.
        /// </summary>
        public readonly int? BootRetryDelay;
        /// <summary>
        /// If set to true, a VM that fails to boot will try again after the 'boot_retry_delay' time period has expired
        /// </summary>
        public readonly bool? BootRetryEnabled;
        /// <summary>
        /// If set to true, enables EFI Secure Boot for the VM. Can only be changed when the VM is powered off.
        /// </summary>
        public readonly bool? EfiSecureBoot;
        public readonly bool? EnterBiosSetupOnNextBoot;

        [OutputConstructor]
        private VmBootOptions(
            int? bootDelay,

            int? bootRetryDelay,

            bool? bootRetryEnabled,

            bool? efiSecureBoot,

            bool? enterBiosSetupOnNextBoot)
        {
            BootDelay = bootDelay;
            BootRetryDelay = bootRetryDelay;
            BootRetryEnabled = bootRetryEnabled;
            EfiSecureBoot = efiSecureBoot;
            EnterBiosSetupOnNextBoot = enterBiosSetupOnNextBoot;
        }
    }
}
