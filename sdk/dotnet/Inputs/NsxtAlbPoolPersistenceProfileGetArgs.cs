// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd.Inputs
{

    public sealed class NsxtAlbPoolPersistenceProfileGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// A name for ALB Pool
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Type of persistence strategy. One of `CLIENT_IP`, `HTTP_COOKIE`, `CUSTOM_HTTP_HEADER`, `APP_COOKIE`, `TLS`
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// Value of attribute based on persistence type
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public NsxtAlbPoolPersistenceProfileGetArgs()
        {
        }
        public static new NsxtAlbPoolPersistenceProfileGetArgs Empty => new NsxtAlbPoolPersistenceProfileGetArgs();
    }
}
