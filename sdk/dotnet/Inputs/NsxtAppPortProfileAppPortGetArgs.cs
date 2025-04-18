// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class NsxtAppPortProfileAppPortGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("ports")]
        private InputList<string>? _ports;

        /// <summary>
        /// Set of ports or ranges
        /// </summary>
        public InputList<string> Ports
        {
            get => _ports ?? (_ports = new InputList<string>());
            set => _ports = value;
        }

        [Input("protocol", required: true)]
        public Input<string> Protocol { get; set; } = null!;

        public NsxtAppPortProfileAppPortGetArgs()
        {
        }
        public static new NsxtAppPortProfileAppPortGetArgs Empty => new NsxtAppPortProfileAppPortGetArgs();
    }
}
