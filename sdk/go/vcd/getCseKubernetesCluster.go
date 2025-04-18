// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/ergSey/pulumi-vcd/sdk/go/vcd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a data source to read Kubernetes clusters in VMware Cloud Director with Container Service Extension (CSE) installed and running.
//
// Supported in provider *v3.12+*
//
// Supports the following **Container Service Extension** versions:
//
// * 4.1.0
// * 4.1.1 / 4.1.1a
// * 4.2.0
// * 4.2.1
// * 4.2.2
// * 4.2.3
//
// > To install CSE in VMware Cloud Director, please follow [this guide](https://www.terraform.io/providers/vmware/vcd/latest/docs/guides/container_service_extension_4_x_install)
//
// ## Example Usage
//
// ### With ID
//
// The cluster ID identifies unequivocally the cluster within VCD, and can be obtained with the CSE Kubernetes Clusters UI Plugin, by selecting
// the desired cluster and obtaining the ID from the displayed information.
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
//			_, err := vcd.LookupCseKubernetesCluster(ctx, &vcd.LookupCseKubernetesClusterArgs{
//				ClusterId: pulumi.StringRef("urn:vcloud:entity:vmware:capvcdCluster:e8e82bcc-50a1-484f-9dd0-20965ab3e865"),
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func LookupCseKubernetesCluster(ctx *pulumi.Context, args *LookupCseKubernetesClusterArgs, opts ...pulumi.InvokeOption) (*LookupCseKubernetesClusterResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupCseKubernetesClusterResult
	err := ctx.Invoke("vcd:index/getCseKubernetesCluster:getCseKubernetesCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCseKubernetesCluster.
type LookupCseKubernetesClusterArgs struct {
	// Unequivocally identifies a cluster in VCD. Either `clusterId` or `name` must be set.
	ClusterId *string `pulumi:"clusterId"`
	// Specifies the CSE Version of the cluster to find when `name` is used instead of `clusterId`.
	CseVersion *string `pulumi:"cseVersion"`
	// Allows to find a Kubernetes cluster by name inside the given Organization with ID `orgId`. Either `clusterId` or `name` must be set. This argument requires `cseVersion` and `orgId` to be set.
	Name *string `pulumi:"name"`
	// The ID of the Organization to which the Kubernetes cluster belongs. Only used if `clusterId` is not set. Must be present if `name` is used.
	OrgId *string `pulumi:"orgId"`
}

// A collection of values returned by getCseKubernetesCluster.
type LookupCseKubernetesClusterResult struct {
	AutoRepairOnErrors         bool                                         `pulumi:"autoRepairOnErrors"`
	CapvcdVersion              string                                       `pulumi:"capvcdVersion"`
	ClusterId                  *string                                      `pulumi:"clusterId"`
	ClusterResourceSetBindings []string                                     `pulumi:"clusterResourceSetBindings"`
	ControlPlanes              []GetCseKubernetesClusterControlPlane        `pulumi:"controlPlanes"`
	CpiVersion                 string                                       `pulumi:"cpiVersion"`
	CseVersion                 *string                                      `pulumi:"cseVersion"`
	CsiVersion                 string                                       `pulumi:"csiVersion"`
	DefaultStorageClasses      []GetCseKubernetesClusterDefaultStorageClass `pulumi:"defaultStorageClasses"`
	Events                     []GetCseKubernetesClusterEvent               `pulumi:"events"`
	// The provider-assigned unique ID for this managed resource.
	Id                   string                              `pulumi:"id"`
	Kubeconfig           string                              `pulumi:"kubeconfig"`
	KubernetesTemplateId string                              `pulumi:"kubernetesTemplateId"`
	KubernetesVersion    string                              `pulumi:"kubernetesVersion"`
	Name                 *string                             `pulumi:"name"`
	NetworkId            string                              `pulumi:"networkId"`
	NodeHealthCheck      bool                                `pulumi:"nodeHealthCheck"`
	OrgId                *string                             `pulumi:"orgId"`
	Owner                string                              `pulumi:"owner"`
	PodsCidr             string                              `pulumi:"podsCidr"`
	Runtime              string                              `pulumi:"runtime"`
	ServicesCidr         string                              `pulumi:"servicesCidr"`
	SshPublicKey         string                              `pulumi:"sshPublicKey"`
	State                string                              `pulumi:"state"`
	SupportedUpgrades    []string                            `pulumi:"supportedUpgrades"`
	TkgProductVersion    string                              `pulumi:"tkgProductVersion"`
	VdcId                string                              `pulumi:"vdcId"`
	VirtualIpSubnet      string                              `pulumi:"virtualIpSubnet"`
	WorkerPools          []GetCseKubernetesClusterWorkerPool `pulumi:"workerPools"`
}

func LookupCseKubernetesClusterOutput(ctx *pulumi.Context, args LookupCseKubernetesClusterOutputArgs, opts ...pulumi.InvokeOption) LookupCseKubernetesClusterResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupCseKubernetesClusterResultOutput, error) {
			args := v.(LookupCseKubernetesClusterArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vcd:index/getCseKubernetesCluster:getCseKubernetesCluster", args, LookupCseKubernetesClusterResultOutput{}, options).(LookupCseKubernetesClusterResultOutput), nil
		}).(LookupCseKubernetesClusterResultOutput)
}

// A collection of arguments for invoking getCseKubernetesCluster.
type LookupCseKubernetesClusterOutputArgs struct {
	// Unequivocally identifies a cluster in VCD. Either `clusterId` or `name` must be set.
	ClusterId pulumi.StringPtrInput `pulumi:"clusterId"`
	// Specifies the CSE Version of the cluster to find when `name` is used instead of `clusterId`.
	CseVersion pulumi.StringPtrInput `pulumi:"cseVersion"`
	// Allows to find a Kubernetes cluster by name inside the given Organization with ID `orgId`. Either `clusterId` or `name` must be set. This argument requires `cseVersion` and `orgId` to be set.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The ID of the Organization to which the Kubernetes cluster belongs. Only used if `clusterId` is not set. Must be present if `name` is used.
	OrgId pulumi.StringPtrInput `pulumi:"orgId"`
}

func (LookupCseKubernetesClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCseKubernetesClusterArgs)(nil)).Elem()
}

// A collection of values returned by getCseKubernetesCluster.
type LookupCseKubernetesClusterResultOutput struct{ *pulumi.OutputState }

func (LookupCseKubernetesClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCseKubernetesClusterResult)(nil)).Elem()
}

func (o LookupCseKubernetesClusterResultOutput) ToLookupCseKubernetesClusterResultOutput() LookupCseKubernetesClusterResultOutput {
	return o
}

func (o LookupCseKubernetesClusterResultOutput) ToLookupCseKubernetesClusterResultOutputWithContext(ctx context.Context) LookupCseKubernetesClusterResultOutput {
	return o
}

func (o LookupCseKubernetesClusterResultOutput) AutoRepairOnErrors() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupCseKubernetesClusterResult) bool { return v.AutoRepairOnErrors }).(pulumi.BoolOutput)
}

func (o LookupCseKubernetesClusterResultOutput) CapvcdVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCseKubernetesClusterResult) string { return v.CapvcdVersion }).(pulumi.StringOutput)
}

func (o LookupCseKubernetesClusterResultOutput) ClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCseKubernetesClusterResult) *string { return v.ClusterId }).(pulumi.StringPtrOutput)
}

func (o LookupCseKubernetesClusterResultOutput) ClusterResourceSetBindings() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCseKubernetesClusterResult) []string { return v.ClusterResourceSetBindings }).(pulumi.StringArrayOutput)
}

func (o LookupCseKubernetesClusterResultOutput) ControlPlanes() GetCseKubernetesClusterControlPlaneArrayOutput {
	return o.ApplyT(func(v LookupCseKubernetesClusterResult) []GetCseKubernetesClusterControlPlane { return v.ControlPlanes }).(GetCseKubernetesClusterControlPlaneArrayOutput)
}

func (o LookupCseKubernetesClusterResultOutput) CpiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCseKubernetesClusterResult) string { return v.CpiVersion }).(pulumi.StringOutput)
}

func (o LookupCseKubernetesClusterResultOutput) CseVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCseKubernetesClusterResult) *string { return v.CseVersion }).(pulumi.StringPtrOutput)
}

func (o LookupCseKubernetesClusterResultOutput) CsiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCseKubernetesClusterResult) string { return v.CsiVersion }).(pulumi.StringOutput)
}

func (o LookupCseKubernetesClusterResultOutput) DefaultStorageClasses() GetCseKubernetesClusterDefaultStorageClassArrayOutput {
	return o.ApplyT(func(v LookupCseKubernetesClusterResult) []GetCseKubernetesClusterDefaultStorageClass {
		return v.DefaultStorageClasses
	}).(GetCseKubernetesClusterDefaultStorageClassArrayOutput)
}

func (o LookupCseKubernetesClusterResultOutput) Events() GetCseKubernetesClusterEventArrayOutput {
	return o.ApplyT(func(v LookupCseKubernetesClusterResult) []GetCseKubernetesClusterEvent { return v.Events }).(GetCseKubernetesClusterEventArrayOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupCseKubernetesClusterResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCseKubernetesClusterResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCseKubernetesClusterResultOutput) Kubeconfig() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCseKubernetesClusterResult) string { return v.Kubeconfig }).(pulumi.StringOutput)
}

func (o LookupCseKubernetesClusterResultOutput) KubernetesTemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCseKubernetesClusterResult) string { return v.KubernetesTemplateId }).(pulumi.StringOutput)
}

func (o LookupCseKubernetesClusterResultOutput) KubernetesVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCseKubernetesClusterResult) string { return v.KubernetesVersion }).(pulumi.StringOutput)
}

func (o LookupCseKubernetesClusterResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCseKubernetesClusterResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupCseKubernetesClusterResultOutput) NetworkId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCseKubernetesClusterResult) string { return v.NetworkId }).(pulumi.StringOutput)
}

func (o LookupCseKubernetesClusterResultOutput) NodeHealthCheck() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupCseKubernetesClusterResult) bool { return v.NodeHealthCheck }).(pulumi.BoolOutput)
}

func (o LookupCseKubernetesClusterResultOutput) OrgId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCseKubernetesClusterResult) *string { return v.OrgId }).(pulumi.StringPtrOutput)
}

func (o LookupCseKubernetesClusterResultOutput) Owner() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCseKubernetesClusterResult) string { return v.Owner }).(pulumi.StringOutput)
}

func (o LookupCseKubernetesClusterResultOutput) PodsCidr() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCseKubernetesClusterResult) string { return v.PodsCidr }).(pulumi.StringOutput)
}

func (o LookupCseKubernetesClusterResultOutput) Runtime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCseKubernetesClusterResult) string { return v.Runtime }).(pulumi.StringOutput)
}

func (o LookupCseKubernetesClusterResultOutput) ServicesCidr() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCseKubernetesClusterResult) string { return v.ServicesCidr }).(pulumi.StringOutput)
}

func (o LookupCseKubernetesClusterResultOutput) SshPublicKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCseKubernetesClusterResult) string { return v.SshPublicKey }).(pulumi.StringOutput)
}

func (o LookupCseKubernetesClusterResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCseKubernetesClusterResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupCseKubernetesClusterResultOutput) SupportedUpgrades() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCseKubernetesClusterResult) []string { return v.SupportedUpgrades }).(pulumi.StringArrayOutput)
}

func (o LookupCseKubernetesClusterResultOutput) TkgProductVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCseKubernetesClusterResult) string { return v.TkgProductVersion }).(pulumi.StringOutput)
}

func (o LookupCseKubernetesClusterResultOutput) VdcId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCseKubernetesClusterResult) string { return v.VdcId }).(pulumi.StringOutput)
}

func (o LookupCseKubernetesClusterResultOutput) VirtualIpSubnet() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCseKubernetesClusterResult) string { return v.VirtualIpSubnet }).(pulumi.StringOutput)
}

func (o LookupCseKubernetesClusterResultOutput) WorkerPools() GetCseKubernetesClusterWorkerPoolArrayOutput {
	return o.ApplyT(func(v LookupCseKubernetesClusterResult) []GetCseKubernetesClusterWorkerPool { return v.WorkerPools }).(GetCseKubernetesClusterWorkerPoolArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCseKubernetesClusterResultOutput{})
}
