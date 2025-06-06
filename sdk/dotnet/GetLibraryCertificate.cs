// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Vcd
{
    public static class GetLibraryCertificate
    {
        public static Task<GetLibraryCertificateResult> InvokeAsync(GetLibraryCertificateArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLibraryCertificateResult>("vcd:index/getLibraryCertificate:getLibraryCertificate", args ?? new GetLibraryCertificateArgs(), options.WithDefaults());

        public static Output<GetLibraryCertificateResult> Invoke(GetLibraryCertificateInvokeArgs? args = null, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLibraryCertificateResult>("vcd:index/getLibraryCertificate:getLibraryCertificate", args ?? new GetLibraryCertificateInvokeArgs(), options.WithDefaults());

        public static Output<GetLibraryCertificateResult> Invoke(GetLibraryCertificateInvokeArgs args, InvokeOutputOptions options)
            => global::Pulumi.Deployment.Instance.Invoke<GetLibraryCertificateResult>("vcd:index/getLibraryCertificate:getLibraryCertificate", args ?? new GetLibraryCertificateInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLibraryCertificateArgs : global::Pulumi.InvokeArgs
    {
        [Input("alias")]
        public string? Alias { get; set; }

        [Input("certificate")]
        public string? Certificate { get; set; }

        [Input("description")]
        public string? Description { get; set; }

        [Input("id")]
        public string? Id { get; set; }

        [Input("org")]
        public string? Org { get; set; }

        public GetLibraryCertificateArgs()
        {
        }
        public static new GetLibraryCertificateArgs Empty => new GetLibraryCertificateArgs();
    }

    public sealed class GetLibraryCertificateInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("alias")]
        public Input<string>? Alias { get; set; }

        [Input("certificate")]
        public Input<string>? Certificate { get; set; }

        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("org")]
        public Input<string>? Org { get; set; }

        public GetLibraryCertificateInvokeArgs()
        {
        }
        public static new GetLibraryCertificateInvokeArgs Empty => new GetLibraryCertificateInvokeArgs();
    }


    [OutputType]
    public sealed class GetLibraryCertificateResult
    {
        public readonly string Alias;
        public readonly string Certificate;
        public readonly string Description;
        public readonly string Id;
        public readonly string? Org;

        [OutputConstructor]
        private GetLibraryCertificateResult(
            string alias,

            string certificate,

            string description,

            string id,

            string? org)
        {
            Alias = alias;
            Certificate = certificate;
            Description = description;
            Id = id;
            Org = org;
        }
    }
}
