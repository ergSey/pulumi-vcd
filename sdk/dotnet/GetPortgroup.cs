// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetPortgroup
    {
        /// <summary>
        /// Provides a data source for available vCenter Port Groups.
        /// 
        /// Supported in provider *v3.0+*
        /// 
        /// 
        /// ## Example Usage
        /// 
        /// ### VSwitch Port Group
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vcd = Pulumi.Vcd;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var first_pg_vswitch = Vcd.GetPortgroup.Invoke(new()
        ///     {
        ///         Name = "pg-name",
        ///         Type = "NETWORK",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// 
        /// ### Distributed Port Group
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vcd = Pulumi.Vcd;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var first_pg_dvswitch = Vcd.GetPortgroup.Invoke(new()
        ///     {
        ///         Name = "pg-name",
        ///         Type = "DV_PORTGROUP",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetPortgroupResult> InvokeAsync(GetPortgroupArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetPortgroupResult>("vcd:index/getPortgroup:getPortgroup", args ?? new GetPortgroupArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a data source for available vCenter Port Groups.
        /// 
        /// Supported in provider *v3.0+*
        /// 
        /// 
        /// ## Example Usage
        /// 
        /// ### VSwitch Port Group
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vcd = Pulumi.Vcd;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var first_pg_vswitch = Vcd.GetPortgroup.Invoke(new()
        ///     {
        ///         Name = "pg-name",
        ///         Type = "NETWORK",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// 
        /// ### Distributed Port Group
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vcd = Pulumi.Vcd;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var first_pg_dvswitch = Vcd.GetPortgroup.Invoke(new()
        ///     {
        ///         Name = "pg-name",
        ///         Type = "DV_PORTGROUP",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetPortgroupResult> Invoke(GetPortgroupInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetPortgroupResult>("vcd:index/getPortgroup:getPortgroup", args ?? new GetPortgroupInvokeArgs(), options.WithDefaults());

        /// <summary>
        /// Provides a data source for available vCenter Port Groups.
        /// 
        /// Supported in provider *v3.0+*
        /// 
        /// 
        /// ## Example Usage
        /// 
        /// ### VSwitch Port Group
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vcd = Pulumi.Vcd;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var first_pg_vswitch = Vcd.GetPortgroup.Invoke(new()
        ///     {
        ///         Name = "pg-name",
        ///         Type = "NETWORK",
        ///     });
        /// 
        /// });
        /// ```
        /// 
        /// 
        /// ### Distributed Port Group
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vcd = Pulumi.Vcd;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var first_pg_dvswitch = Vcd.GetPortgroup.Invoke(new()
        ///     {
        ///         Name = "pg-name",
        ///         Type = "DV_PORTGROUP",
        ///     });
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetPortgroupResult> Invoke(GetPortgroupInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetPortgroupResult>("vcd:index/getPortgroup:getPortgroup", args ?? new GetPortgroupInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetPortgroupArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Organization VDC name
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        /// <summary>
        /// `NETWORK` for vSwitch port group or `DV_PORTGROUP` for distributed port group.
        /// </summary>
        [Input("type", required: true)]
        public string Type { get; set; } = null!;

        public GetPortgroupArgs()
        {
        }
        public static new GetPortgroupArgs Empty => new GetPortgroupArgs();
    }

    public sealed class GetPortgroupInvokeArgs : global::Pulumi.InvokeArgs
    {
        /// <summary>
        /// Organization VDC name
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// `NETWORK` for vSwitch port group or `DV_PORTGROUP` for distributed port group.
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public GetPortgroupInvokeArgs()
        {
        }
        public static new GetPortgroupInvokeArgs Empty => new GetPortgroupInvokeArgs();
    }


    [OutputType]
    public sealed class GetPortgroupResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly string Name;
        public readonly string Type;

        [OutputConstructor]
        private GetPortgroupResult(
            string id,

            string name,

            string type)
        {
            Id = id;
            Name = name;
            Type = type;
        }
    }
}
