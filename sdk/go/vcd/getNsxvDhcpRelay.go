// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/ergSey/pulumi-vcd/sdk/go/vcd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VMware Cloud Director Edge Gateway DHCP relay configuration data source. The DHCP relay
// capability provided by NSX in VMware Cloud Director environment allows to leverage existing DHCP
// infrastructure from within VMware Cloud Director environment without any interruption to the IP address
// management in existing DHCP infrastructure. DHCP messages are relayed from virtual machines to the
// designated DHCP servers in your physical DHCP infrastructure, which allows IP addresses controlled
// by the NSX software to continue to be in sync with IP addresses in the rest of your DHCP-controlled
// environments.
//
// Supported in provider *v2.6+*
//
// ## Example Usage
//
// ### 1
//
// ```go
// package main
//
// import (
//
//	"github.com/ergSey/pulumi-vcd/sdk/go/vcd"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := vcd.LookupNsxvDhcpRelay(ctx, &vcd.LookupNsxvDhcpRelayArgs{
//				Org:         pulumi.StringRef("my-org"),
//				Vdc:         pulumi.StringRef("my-org-vdc"),
//				EdgeGateway: "my-edge-gw",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupNsxvDhcpRelay(ctx *pulumi.Context, args *LookupNsxvDhcpRelayArgs, opts ...pulumi.InvokeOption) (*LookupNsxvDhcpRelayResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNsxvDhcpRelayResult
	err := ctx.Invoke("vcd:index/getNsxvDhcpRelay:getNsxvDhcpRelay", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNsxvDhcpRelay.
type LookupNsxvDhcpRelayArgs struct {
	// The name of the edge gateway on which DHCP relay is to be configured.
	EdgeGateway string `pulumi:"edgeGateway"`
	// The name of organization to use, optional if defined at provider level. Useful
	// when connected as sysadmin working across different organisations.
	Org *string `pulumi:"org"`
	// The name of VDC to use, optional if defined at provider level.
	Vdc *string `pulumi:"vdc"`
}

// A collection of values returned by getNsxvDhcpRelay.
type LookupNsxvDhcpRelayResult struct {
	DomainNames []string `pulumi:"domainNames"`
	EdgeGateway string   `pulumi:"edgeGateway"`
	// The provider-assigned unique ID for this managed resource.
	Id          string                       `pulumi:"id"`
	IpAddresses []string                     `pulumi:"ipAddresses"`
	IpSets      []string                     `pulumi:"ipSets"`
	Org         *string                      `pulumi:"org"`
	RelayAgents []GetNsxvDhcpRelayRelayAgent `pulumi:"relayAgents"`
	Vdc         *string                      `pulumi:"vdc"`
}

func LookupNsxvDhcpRelayOutput(ctx *pulumi.Context, args LookupNsxvDhcpRelayOutputArgs, opts ...pulumi.InvokeOption) LookupNsxvDhcpRelayResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupNsxvDhcpRelayResultOutput, error) {
			args := v.(LookupNsxvDhcpRelayArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vcd:index/getNsxvDhcpRelay:getNsxvDhcpRelay", args, LookupNsxvDhcpRelayResultOutput{}, options).(LookupNsxvDhcpRelayResultOutput), nil
		}).(LookupNsxvDhcpRelayResultOutput)
}

// A collection of arguments for invoking getNsxvDhcpRelay.
type LookupNsxvDhcpRelayOutputArgs struct {
	// The name of the edge gateway on which DHCP relay is to be configured.
	EdgeGateway pulumi.StringInput `pulumi:"edgeGateway"`
	// The name of organization to use, optional if defined at provider level. Useful
	// when connected as sysadmin working across different organisations.
	Org pulumi.StringPtrInput `pulumi:"org"`
	// The name of VDC to use, optional if defined at provider level.
	Vdc pulumi.StringPtrInput `pulumi:"vdc"`
}

func (LookupNsxvDhcpRelayOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxvDhcpRelayArgs)(nil)).Elem()
}

// A collection of values returned by getNsxvDhcpRelay.
type LookupNsxvDhcpRelayResultOutput struct{ *pulumi.OutputState }

func (LookupNsxvDhcpRelayResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNsxvDhcpRelayResult)(nil)).Elem()
}

func (o LookupNsxvDhcpRelayResultOutput) ToLookupNsxvDhcpRelayResultOutput() LookupNsxvDhcpRelayResultOutput {
	return o
}

func (o LookupNsxvDhcpRelayResultOutput) ToLookupNsxvDhcpRelayResultOutputWithContext(ctx context.Context) LookupNsxvDhcpRelayResultOutput {
	return o
}

func (o LookupNsxvDhcpRelayResultOutput) DomainNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNsxvDhcpRelayResult) []string { return v.DomainNames }).(pulumi.StringArrayOutput)
}

func (o LookupNsxvDhcpRelayResultOutput) EdgeGateway() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxvDhcpRelayResult) string { return v.EdgeGateway }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupNsxvDhcpRelayResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNsxvDhcpRelayResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNsxvDhcpRelayResultOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNsxvDhcpRelayResult) []string { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

func (o LookupNsxvDhcpRelayResultOutput) IpSets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNsxvDhcpRelayResult) []string { return v.IpSets }).(pulumi.StringArrayOutput)
}

func (o LookupNsxvDhcpRelayResultOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNsxvDhcpRelayResult) *string { return v.Org }).(pulumi.StringPtrOutput)
}

func (o LookupNsxvDhcpRelayResultOutput) RelayAgents() GetNsxvDhcpRelayRelayAgentArrayOutput {
	return o.ApplyT(func(v LookupNsxvDhcpRelayResult) []GetNsxvDhcpRelayRelayAgent { return v.RelayAgents }).(GetNsxvDhcpRelayRelayAgentArrayOutput)
}

func (o LookupNsxvDhcpRelayResultOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNsxvDhcpRelayResult) *string { return v.Vdc }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNsxvDhcpRelayResultOutput{})
}
