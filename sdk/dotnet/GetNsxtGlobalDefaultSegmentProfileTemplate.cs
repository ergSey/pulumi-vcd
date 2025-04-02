// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetNsxtGlobalDefaultSegmentProfileTemplate
    {
        /// <summary>
        /// Provides a resource to manage Global Default NSX-T Segment Profile Templates.
        /// 
        /// Supported in provider *v3.11+* and VCD 10.4.0+ with NSX-T. Requires System Administrator privileges.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vcd = Pulumi.Vcd;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var singleton = new Vcd.NsxtGlobalDefaultSegmentProfileTemplate("singleton");
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetNsxtGlobalDefaultSegmentProfileTemplateResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetNsxtGlobalDefaultSegmentProfileTemplateResult>("vcd:index/getNsxtGlobalDefaultSegmentProfileTemplate:getNsxtGlobalDefaultSegmentProfileTemplate", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// Provides a resource to manage Global Default NSX-T Segment Profile Templates.
        /// 
        /// Supported in provider *v3.11+* and VCD 10.4.0+ with NSX-T. Requires System Administrator privileges.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vcd = Pulumi.Vcd;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var singleton = new Vcd.NsxtGlobalDefaultSegmentProfileTemplate("singleton");
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetNsxtGlobalDefaultSegmentProfileTemplateResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetNsxtGlobalDefaultSegmentProfileTemplateResult>("vcd:index/getNsxtGlobalDefaultSegmentProfileTemplate:getNsxtGlobalDefaultSegmentProfileTemplate", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// Provides a resource to manage Global Default NSX-T Segment Profile Templates.
        /// 
        /// Supported in provider *v3.11+* and VCD 10.4.0+ with NSX-T. Requires System Administrator privileges.
        /// 
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Vcd = Pulumi.Vcd;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var singleton = new Vcd.NsxtGlobalDefaultSegmentProfileTemplate("singleton");
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetNsxtGlobalDefaultSegmentProfileTemplateResult> Invoke(InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetNsxtGlobalDefaultSegmentProfileTemplateResult>("vcd:index/getNsxtGlobalDefaultSegmentProfileTemplate:getNsxtGlobalDefaultSegmentProfileTemplate", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetNsxtGlobalDefaultSegmentProfileTemplateResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Global Default Segment Profile
        /// Template ID for all vApp Networks
        /// </summary>
        public readonly string VappNetworksDefaultSegmentProfileTemplateId;
        /// <summary>
        /// Global Default Segment Profile
        /// Template ID for all VDC Networks
        /// </summary>
        public readonly string VdcNetworksDefaultSegmentProfileTemplateId;

        [OutputConstructor]
        private GetNsxtGlobalDefaultSegmentProfileTemplateResult(
            string id,

            string vappNetworksDefaultSegmentProfileTemplateId,

            string vdcNetworksDefaultSegmentProfileTemplateId)
        {
            Id = id;
            VappNetworksDefaultSegmentProfileTemplateId = vappNetworksDefaultSegmentProfileTemplateId;
            VdcNetworksDefaultSegmentProfileTemplateId = vdcNetworksDefaultSegmentProfileTemplateId;
        }
    }
}
