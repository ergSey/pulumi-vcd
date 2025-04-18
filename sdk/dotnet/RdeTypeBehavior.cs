// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    [VcdResourceType("vcd:index/rdeTypeBehavior:RdeTypeBehavior")]
    public partial class RdeTypeBehavior : global::Pulumi.CustomResource
    {
        /// <summary>
        /// Useful to update execution properties marked with `_secure_` and `_internal_`
        /// as these are not retrievable from VCD, so they are not saved in state. Setting this to `true` will make the provider
        /// to ask for updates whenever there is a secure property in the execution of the Behavior
        /// </summary>
        [Output("alwaysUpdateSecureExecutionProperties")]
        public Output<bool?> AlwaysUpdateSecureExecutionProperties { get; private set; } = null!;

        /// <summary>
        /// The description of the RDE Type Behavior.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// A map that specifies the Behavior execution mechanism, this is just a simplification of `execution_json` that
        /// can make the configuration more readable for simpler Behaviors. One of `execution_json` or `execution` must be set.
        /// </summary>
        [Output("execution")]
        public Output<ImmutableDictionary<string, string>> Execution { get; private set; } = null!;

        /// <summary>
        /// A string representing a valid JSON that specifies the Behavior execution mechanism.
        /// You can find more information about the different execution types, like `WebHook`, `noop`, `Activity`, `MQTT`, `VRO`, `AWSLambdaFaaS`
        /// and others [in the Extensibility SDK documentation](https://vmware.github.io/vcd-ext-sdk/docs/defined_entities_api/behaviors).
        /// One of `execution_json` or `execution` must be set.
        /// </summary>
        [Output("executionJson")]
        public Output<string> ExecutionJson { get; private set; } = null!;

        /// <summary>
        /// Name of the overridden Behavior
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The ID of the [RDE Interface Behavior](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/rde_interface_behavior) to override
        /// </summary>
        [Output("rdeInterfaceBehaviorId")]
        public Output<string> RdeInterfaceBehaviorId { get; private set; } = null!;

        /// <summary>
        /// The ID of the RDE Type that owns the Behavior
        /// </summary>
        [Output("rdeTypeId")]
        public Output<string> RdeTypeId { get; private set; } = null!;

        /// <summary>
        /// The Behavior invocation reference to be used for polymorphic behavior invocations
        /// </summary>
        [Output("ref")]
        public Output<string> Ref { get; private set; } = null!;


        /// <summary>
        /// Create a RdeTypeBehavior resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RdeTypeBehavior(string name, RdeTypeBehaviorArgs args, CustomResourceOptions? options = null)
            : base("vcd:index/rdeTypeBehavior:RdeTypeBehavior", name, args ?? new RdeTypeBehaviorArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RdeTypeBehavior(string name, Input<string> id, RdeTypeBehaviorState? state = null, CustomResourceOptions? options = null)
            : base("vcd:index/rdeTypeBehavior:RdeTypeBehavior", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RdeTypeBehavior resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RdeTypeBehavior Get(string name, Input<string> id, RdeTypeBehaviorState? state = null, CustomResourceOptions? options = null)
        {
            return new RdeTypeBehavior(name, id, state, options);
        }
    }

    public sealed class RdeTypeBehaviorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Useful to update execution properties marked with `_secure_` and `_internal_`
        /// as these are not retrievable from VCD, so they are not saved in state. Setting this to `true` will make the provider
        /// to ask for updates whenever there is a secure property in the execution of the Behavior
        /// </summary>
        [Input("alwaysUpdateSecureExecutionProperties")]
        public Input<bool>? AlwaysUpdateSecureExecutionProperties { get; set; }

        /// <summary>
        /// The description of the RDE Type Behavior.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("execution")]
        private InputMap<string>? _execution;

        /// <summary>
        /// A map that specifies the Behavior execution mechanism, this is just a simplification of `execution_json` that
        /// can make the configuration more readable for simpler Behaviors. One of `execution_json` or `execution` must be set.
        /// </summary>
        public InputMap<string> Execution
        {
            get => _execution ?? (_execution = new InputMap<string>());
            set => _execution = value;
        }

        /// <summary>
        /// A string representing a valid JSON that specifies the Behavior execution mechanism.
        /// You can find more information about the different execution types, like `WebHook`, `noop`, `Activity`, `MQTT`, `VRO`, `AWSLambdaFaaS`
        /// and others [in the Extensibility SDK documentation](https://vmware.github.io/vcd-ext-sdk/docs/defined_entities_api/behaviors).
        /// One of `execution_json` or `execution` must be set.
        /// </summary>
        [Input("executionJson")]
        public Input<string>? ExecutionJson { get; set; }

        /// <summary>
        /// The ID of the [RDE Interface Behavior](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/rde_interface_behavior) to override
        /// </summary>
        [Input("rdeInterfaceBehaviorId", required: true)]
        public Input<string> RdeInterfaceBehaviorId { get; set; } = null!;

        /// <summary>
        /// The ID of the RDE Type that owns the Behavior
        /// </summary>
        [Input("rdeTypeId", required: true)]
        public Input<string> RdeTypeId { get; set; } = null!;

        public RdeTypeBehaviorArgs()
        {
        }
        public static new RdeTypeBehaviorArgs Empty => new RdeTypeBehaviorArgs();
    }

    public sealed class RdeTypeBehaviorState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Useful to update execution properties marked with `_secure_` and `_internal_`
        /// as these are not retrievable from VCD, so they are not saved in state. Setting this to `true` will make the provider
        /// to ask for updates whenever there is a secure property in the execution of the Behavior
        /// </summary>
        [Input("alwaysUpdateSecureExecutionProperties")]
        public Input<bool>? AlwaysUpdateSecureExecutionProperties { get; set; }

        /// <summary>
        /// The description of the RDE Type Behavior.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("execution")]
        private InputMap<string>? _execution;

        /// <summary>
        /// A map that specifies the Behavior execution mechanism, this is just a simplification of `execution_json` that
        /// can make the configuration more readable for simpler Behaviors. One of `execution_json` or `execution` must be set.
        /// </summary>
        public InputMap<string> Execution
        {
            get => _execution ?? (_execution = new InputMap<string>());
            set => _execution = value;
        }

        /// <summary>
        /// A string representing a valid JSON that specifies the Behavior execution mechanism.
        /// You can find more information about the different execution types, like `WebHook`, `noop`, `Activity`, `MQTT`, `VRO`, `AWSLambdaFaaS`
        /// and others [in the Extensibility SDK documentation](https://vmware.github.io/vcd-ext-sdk/docs/defined_entities_api/behaviors).
        /// One of `execution_json` or `execution` must be set.
        /// </summary>
        [Input("executionJson")]
        public Input<string>? ExecutionJson { get; set; }

        /// <summary>
        /// Name of the overridden Behavior
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The ID of the [RDE Interface Behavior](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/rde_interface_behavior) to override
        /// </summary>
        [Input("rdeInterfaceBehaviorId")]
        public Input<string>? RdeInterfaceBehaviorId { get; set; }

        /// <summary>
        /// The ID of the RDE Type that owns the Behavior
        /// </summary>
        [Input("rdeTypeId")]
        public Input<string>? RdeTypeId { get; set; }

        /// <summary>
        /// The Behavior invocation reference to be used for polymorphic behavior invocations
        /// </summary>
        [Input("ref")]
        public Input<string>? Ref { get; set; }

        public RdeTypeBehaviorState()
        {
        }
        public static new RdeTypeBehaviorState Empty => new RdeTypeBehaviorState();
    }
}
