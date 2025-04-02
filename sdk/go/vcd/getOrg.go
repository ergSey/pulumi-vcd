// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/ergSey/pulumi-vcd/sdk/go/vcd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a VMware Cloud Director Org data source. An organization can be used to manage catalogs, virtual
// data centers, and users.
//
// Supported in provider *v2.5+*
//
// ## Example Usage
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
//			my_org, err := vcd.LookupOrg(ctx, &vcd.LookupOrgArgs{
//				Name: "my-org",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			_, err = vcd.NewOrg(ctx, "my-org-clone", &vcd.OrgArgs{
//				Name:               pulumi.String("my-org-clone"),
//				FullName:           pulumi.String(my_org.FullName),
//				CanPublishCatalogs: pulumi.Bool(my_org.CanPublishCatalogs),
//				DeployedVmQuota:    pulumi.Int(my_org.DeployedVmQuota),
//				StoredVmQuota:      pulumi.Int(my_org.StoredVmQuota),
//				IsEnabled:          pulumi.Bool(my_org.IsEnabled),
//				DeleteForce:        pulumi.Bool(true),
//				DeleteRecursive:    pulumi.Bool(true),
//				VappLease: &vcd.OrgVappLeaseArgs{
//					MaximumRuntimeLeaseInSec:         pulumi.Int(my_org.VappLeases[0].MaximumRuntimeLeaseInSec),
//					PowerOffOnRuntimeLeaseExpiration: pulumi.Bool(my_org.VappLeases[0].PowerOffOnRuntimeLeaseExpiration),
//					MaximumStorageLeaseInSec:         pulumi.Int(my_org.VappLeases[0].MaximumStorageLeaseInSec),
//					DeleteOnStorageLeaseExpiration:   pulumi.Bool(my_org.VappLeases[0].DeleteOnStorageLeaseExpiration),
//				},
//				VappTemplateLease: &vcd.OrgVappTemplateLeaseArgs{
//					MaximumStorageLeaseInSec:       pulumi.Int(my_org.VappTemplateLeases[0].MaximumStorageLeaseInSec),
//					DeleteOnStorageLeaseExpiration: pulumi.Bool(my_org.VappTemplateLeases[0].DeleteOnStorageLeaseExpiration),
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## vApp Lease
//
// The `vappLease` section contains lease parameters for vApps created in the current organization, as defined below:
//
// * `maximumRuntimeLeaseInSec` - How long vApps can run before they are automatically stopped (in seconds)
// * `powerOffOnRuntimeLeaseExpiration` - When true, vApps are powered off when the runtime lease expires. When false, vApps are suspended when the runtime lease expires.
// * `maximumStorageLeaseInSec` - How long stopped vApps are available before being automatically cleaned up (in seconds)
// * `deleteOnStorageLeaseExpiration` - If true, storage for a vApp is deleted when the vApp's lease expires. If false, the storage is flagged for deletion, but not deleted.
//
// <a id="vapp-template-lease"></a>
// ## vApp Template Lease
//
// The `vappTemplateLease` section contains lease parameters for vApp templates created in the current organization, as defined below:
//
// * `maximumStorageLeaseInSec` - How long vApp templates are available before being automatically cleaned up (in seconds)
// * `deleteOnStorageLeaseExpiration` - If true, storage for a vAppTemplate is deleted when the vAppTemplate lease expires. If false, the storage is flagged for deletion, but not deleted
//
// <a id="metadata"></a>
// ## Metadata
//
// The `metadataEntry` (*v3.8+*) is a set of metadata entries that have the following structure:
//
// * `key` - Key of this metadata entry.
// * `value` - Value of this metadata entry.
// * `type` - Type of this metadata entry. One of: `MetadataStringValue`, `MetadataNumberValue`, `MetadataDateTimeValue`, `MetadataBooleanValue`.
// * `userAccess` - User access level for this metadata entry. One of: `PRIVATE` (hidden), `READONLY` (read only), `READWRITE` (read/write).
// * `isSystem` - Domain for this metadata entry. true if it belongs to `SYSTEM`, false if it belongs to `GENERAL`.
func LookupOrg(ctx *pulumi.Context, args *LookupOrgArgs, opts ...pulumi.InvokeOption) (*LookupOrgResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupOrgResult
	err := ctx.Invoke("vcd:index/getOrg:getOrg", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getOrg.
type LookupOrgArgs struct {
	// Org name
	Name string `pulumi:"name"`
}

// A collection of values returned by getOrg.
type LookupOrgResult struct {
	// (*v3.14+*) Contains the account lockout properties of the read organization:
	AccountLockouts []GetOrgAccountLockout `pulumi:"accountLockouts"`
	// True if this organization is allowed to share catalogs.
	CanPublishCatalogs bool `pulumi:"canPublishCatalogs"`
	// (*v3.6+*) True if this organization is allowed to publish external catalogs.
	CanPublishExternalCatalogs bool `pulumi:"canPublishExternalCatalogs"`
	// (*v3.6+*) True if this organization is allowed to subscribe to external catalogs.
	CanSubscribeExternalCatalogs bool `pulumi:"canSubscribeExternalCatalogs"`
	// Specifies this organization's default for virtual machine boot delay after power on.
	DelayAfterPowerOnSeconds int `pulumi:"delayAfterPowerOnSeconds"`
	// Maximum number of virtual machines that can be deployed simultaneously by a member of this organization.
	DeployedVmQuota int `pulumi:"deployedVmQuota"`
	// Org description.
	Description string `pulumi:"description"`
	// Org full name
	FullName string `pulumi:"fullName"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// True if this organization is enabled (allows login and all other operations).
	IsEnabled bool `pulumi:"isEnabled"`
	// (*v3.11+*) List of catalogs (sorted alphabetically), owned or shared, available to this organization.
	ListOfCatalogs []string `pulumi:"listOfCatalogs"`
	// (*v3.11+*) List of VDCs (sorted alphabetically), owned or shared, available to this organization.
	ListOfVdcs []string `pulumi:"listOfVdcs"`
	// (Deprecated; *v3.6+*) Use `metadataEntry` instead. Key value map of metadata assigned to this organization.
	//
	// Deprecated: Use metadataEntry instead
	Metadata map[string]string `pulumi:"metadata"`
	// A set of metadata entries assigned to the organization. See Metadata section for details.
	MetadataEntries []GetOrgMetadataEntry `pulumi:"metadataEntries"`
	Name            string                `pulumi:"name"`
	// (*v3.11+*) Number of catalogs owned or shared, available to this organization.
	NumberOfCatalogs int `pulumi:"numberOfCatalogs"`
	// (*v3.11+*) Number of VDCs owned or shared, available to this organization.
	NumberOfVdcs int `pulumi:"numberOfVdcs"`
	// Maximum number of virtual machines in vApps or vApp templates that can be stored in an undeployed state by a member of this organization.
	StoredVmQuota int `pulumi:"storedVmQuota"`
	// (*v2.7+*) Defines lease parameters for vApps created in this organization. See vApp Lease below for details.
	VappLeases []GetOrgVappLease `pulumi:"vappLeases"`
	// (*v2.7+*) Defines lease parameters for vApp templates created in this organization. See vApp Template Lease below for details.
	VappTemplateLeases []GetOrgVappTemplateLease `pulumi:"vappTemplateLeases"`
}

func LookupOrgOutput(ctx *pulumi.Context, args LookupOrgOutputArgs, opts ...pulumi.InvokeOption) LookupOrgResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (LookupOrgResultOutput, error) {
			args := v.(LookupOrgArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vcd:index/getOrg:getOrg", args, LookupOrgResultOutput{}, options).(LookupOrgResultOutput), nil
		}).(LookupOrgResultOutput)
}

// A collection of arguments for invoking getOrg.
type LookupOrgOutputArgs struct {
	// Org name
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupOrgOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrgArgs)(nil)).Elem()
}

// A collection of values returned by getOrg.
type LookupOrgResultOutput struct{ *pulumi.OutputState }

func (LookupOrgResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOrgResult)(nil)).Elem()
}

func (o LookupOrgResultOutput) ToLookupOrgResultOutput() LookupOrgResultOutput {
	return o
}

func (o LookupOrgResultOutput) ToLookupOrgResultOutputWithContext(ctx context.Context) LookupOrgResultOutput {
	return o
}

// (*v3.14+*) Contains the account lockout properties of the read organization:
func (o LookupOrgResultOutput) AccountLockouts() GetOrgAccountLockoutArrayOutput {
	return o.ApplyT(func(v LookupOrgResult) []GetOrgAccountLockout { return v.AccountLockouts }).(GetOrgAccountLockoutArrayOutput)
}

// True if this organization is allowed to share catalogs.
func (o LookupOrgResultOutput) CanPublishCatalogs() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOrgResult) bool { return v.CanPublishCatalogs }).(pulumi.BoolOutput)
}

// (*v3.6+*) True if this organization is allowed to publish external catalogs.
func (o LookupOrgResultOutput) CanPublishExternalCatalogs() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOrgResult) bool { return v.CanPublishExternalCatalogs }).(pulumi.BoolOutput)
}

// (*v3.6+*) True if this organization is allowed to subscribe to external catalogs.
func (o LookupOrgResultOutput) CanSubscribeExternalCatalogs() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOrgResult) bool { return v.CanSubscribeExternalCatalogs }).(pulumi.BoolOutput)
}

// Specifies this organization's default for virtual machine boot delay after power on.
func (o LookupOrgResultOutput) DelayAfterPowerOnSeconds() pulumi.IntOutput {
	return o.ApplyT(func(v LookupOrgResult) int { return v.DelayAfterPowerOnSeconds }).(pulumi.IntOutput)
}

// Maximum number of virtual machines that can be deployed simultaneously by a member of this organization.
func (o LookupOrgResultOutput) DeployedVmQuota() pulumi.IntOutput {
	return o.ApplyT(func(v LookupOrgResult) int { return v.DeployedVmQuota }).(pulumi.IntOutput)
}

// Org description.
func (o LookupOrgResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgResult) string { return v.Description }).(pulumi.StringOutput)
}

// Org full name
func (o LookupOrgResultOutput) FullName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgResult) string { return v.FullName }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupOrgResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgResult) string { return v.Id }).(pulumi.StringOutput)
}

// True if this organization is enabled (allows login and all other operations).
func (o LookupOrgResultOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupOrgResult) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

// (*v3.11+*) List of catalogs (sorted alphabetically), owned or shared, available to this organization.
func (o LookupOrgResultOutput) ListOfCatalogs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupOrgResult) []string { return v.ListOfCatalogs }).(pulumi.StringArrayOutput)
}

// (*v3.11+*) List of VDCs (sorted alphabetically), owned or shared, available to this organization.
func (o LookupOrgResultOutput) ListOfVdcs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupOrgResult) []string { return v.ListOfVdcs }).(pulumi.StringArrayOutput)
}

// (Deprecated; *v3.6+*) Use `metadataEntry` instead. Key value map of metadata assigned to this organization.
//
// Deprecated: Use metadataEntry instead
func (o LookupOrgResultOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupOrgResult) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

// A set of metadata entries assigned to the organization. See Metadata section for details.
func (o LookupOrgResultOutput) MetadataEntries() GetOrgMetadataEntryArrayOutput {
	return o.ApplyT(func(v LookupOrgResult) []GetOrgMetadataEntry { return v.MetadataEntries }).(GetOrgMetadataEntryArrayOutput)
}

func (o LookupOrgResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOrgResult) string { return v.Name }).(pulumi.StringOutput)
}

// (*v3.11+*) Number of catalogs owned or shared, available to this organization.
func (o LookupOrgResultOutput) NumberOfCatalogs() pulumi.IntOutput {
	return o.ApplyT(func(v LookupOrgResult) int { return v.NumberOfCatalogs }).(pulumi.IntOutput)
}

// (*v3.11+*) Number of VDCs owned or shared, available to this organization.
func (o LookupOrgResultOutput) NumberOfVdcs() pulumi.IntOutput {
	return o.ApplyT(func(v LookupOrgResult) int { return v.NumberOfVdcs }).(pulumi.IntOutput)
}

// Maximum number of virtual machines in vApps or vApp templates that can be stored in an undeployed state by a member of this organization.
func (o LookupOrgResultOutput) StoredVmQuota() pulumi.IntOutput {
	return o.ApplyT(func(v LookupOrgResult) int { return v.StoredVmQuota }).(pulumi.IntOutput)
}

// (*v2.7+*) Defines lease parameters for vApps created in this organization. See vApp Lease below for details.
func (o LookupOrgResultOutput) VappLeases() GetOrgVappLeaseArrayOutput {
	return o.ApplyT(func(v LookupOrgResult) []GetOrgVappLease { return v.VappLeases }).(GetOrgVappLeaseArrayOutput)
}

// (*v2.7+*) Defines lease parameters for vApp templates created in this organization. See vApp Template Lease below for details.
func (o LookupOrgResultOutput) VappTemplateLeases() GetOrgVappTemplateLeaseArrayOutput {
	return o.ApplyT(func(v LookupOrgResult) []GetOrgVappTemplateLease { return v.VappTemplateLeases }).(GetOrgVappTemplateLeaseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOrgResultOutput{})
}
