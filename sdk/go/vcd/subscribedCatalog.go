// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"errors"
	"github.com/ergSey/pulumi-vcd/sdk/go/vcd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SubscribedCatalog struct {
	pulumi.CustomResourceState

	// When `true`, the subscribed catalog will attempt canceling failed tasks.
	CancelFailedTasks pulumi.BoolPtrOutput `pulumi:"cancelFailedTasks"`
	// Version number from this catalog. This is inherited from the publishing catalog and updated on sync.
	CatalogVersion pulumi.IntOutput `pulumi:"catalogVersion"`
	// Date and time of catalog creation. This is the creation date of the subscription, not the original published catalog.
	Created pulumi.StringOutput `pulumi:"created"`
	// When destroying use `delete_force=true` with `delete_recursive=true` to remove a catalog and any objects it contains, regardless of their state.
	DeleteForce pulumi.BoolPtrOutput `pulumi:"deleteForce"`
	// When destroying use `delete_recursive=true` to remove the catalog and any objects it contains that are in a state that normally allows removal.
	DeleteRecursive pulumi.BoolPtrOutput `pulumi:"deleteRecursive"`
	// Description of catalog. This is inherited from the publishing catalog and updated on sync.
	Description pulumi.StringOutput `pulumi:"description"`
	// List of synchronization tasks that are have failed. They can refer to the catalog or any of its catalog items.
	FailedTasks pulumi.StringArrayOutput `pulumi:"failedTasks"`
	// the catalog's Hyper reference.
	Href pulumi.StringOutput `pulumi:"href"`
	// (*v3.8.1+*) Indicates if this catalog was created in the current organization.
	IsLocal pulumi.BoolOutput `pulumi:"isLocal"`
	// Indicates if this catalog is available for subscription. (Always false)
	IsPublished pulumi.BoolOutput `pulumi:"isPublished"`
	// Indicates if the catalog is shared.
	IsShared pulumi.BoolOutput `pulumi:"isShared"`
	// If `true`, subscription to a catalog creates a local copy of all items. Defaults to `false`, which does not create a local copy of catalog items unless a sync operation is performed.
	// It can only be `false` if the user configured in the provider is the System administrator.
	MakeLocalCopy pulumi.BoolPtrOutput `pulumi:"makeLocalCopy"`
	// List of media item names in this catalog, in alphabetical order.
	MediaItemLists pulumi.StringArrayOutput `pulumi:"mediaItemLists"`
	// (*Available until VCD 10.5*) Optional metadata of the catalog. This is inherited from the publishing catalog and updated on sync.
	Metadata pulumi.StringMapOutput `pulumi:"metadata"`
	// Catalog name
	Name pulumi.StringOutput `pulumi:"name"`
	// Number of media items available in this catalog.
	NumberOfMedia pulumi.IntOutput `pulumi:"numberOfMedia"`
	// Number of vApp templates available in this catalog.
	NumberOfVappTemplates pulumi.IntOutput `pulumi:"numberOfVappTemplates"`
	// The name of organization to use, optional if defined at provider level.
	Org pulumi.StringPtrOutput `pulumi:"org"`
	// Owner of the catalog.
	OwnerName pulumi.StringOutput `pulumi:"ownerName"`
	// Shows if the catalog is published, if it is a subscription from another one or none of those. (Always `SUBSCRIBED`)
	PublishSubscriptionType pulumi.StringOutput `pulumi:"publishSubscriptionType"`
	// List of running synchronization tasks that are still running. They can refer to the catalog or any of its catalog items.
	RunningTasks pulumi.StringArrayOutput `pulumi:"runningTasks"`
	// Allows to set specific storage profile to be used for catalog.
	StorageProfileId pulumi.StringPtrOutput `pulumi:"storageProfileId"`
	// if `true`, saves the list of tasks to a file for later update.
	StoreTasks pulumi.BoolPtrOutput `pulumi:"storeTasks"`
	// An optional password to access the catalog. Only ASCII characters are allowed in a valid password.
	// The password is only required when set by the publishing catalog. Passing in six asterisks '******' indicates to keep current password.
	// Passing in an empty string indicates to remove password.
	SubscriptionPassword pulumi.StringOutput `pulumi:"subscriptionPassword"`
	// The URL to subscribe to the external catalog.
	SubscriptionUrl pulumi.StringOutput `pulumi:"subscriptionUrl"`
	// If `true`, synchronise this catalog and all items.
	SyncAll pulumi.BoolPtrOutput `pulumi:"syncAll"`
	// If `true`, synchronise all media items. Not to be used when `syncAll` is set.
	SyncAllMediaItems pulumi.BoolPtrOutput `pulumi:"syncAllMediaItems"`
	// If `true`, synchronise all vApp templates. Not to be used when `syncAll` is set.
	SyncAllVappTemplates pulumi.BoolPtrOutput `pulumi:"syncAllVappTemplates"`
	// If `true`, synchronise this catalog. Not to be used when `syncAll` is set. This operation fetches the list of items. If `makeLocalCopy` is set, it also synchronises all the items.
	SyncCatalog pulumi.BoolPtrOutput `pulumi:"syncCatalog"`
	// Synchronise a list of media items. Not to be used when `syncAll` or `syncAllMediaItems` are set.
	SyncMediaItems pulumi.StringArrayOutput `pulumi:"syncMediaItems"`
	// Boolean value that shows if sync should be performed on every refresh.
	SyncOnRefresh pulumi.BoolPtrOutput `pulumi:"syncOnRefresh"`
	// Synchronise a list of vApp templates. Not to be used when `syncAll` or `syncAllVappTemplates` are set.
	SyncVappTemplates pulumi.StringArrayOutput `pulumi:"syncVappTemplates"`
	// Where the running tasks IDs have been stored. Only if `storeTasks` is set.
	TasksFileName pulumi.StringOutput `pulumi:"tasksFileName"`
	// List of vApp template names in this catalog, in alphabetical order.
	VappTemplateLists pulumi.StringArrayOutput `pulumi:"vappTemplateLists"`
}

// NewSubscribedCatalog registers a new resource with the given unique name, arguments, and options.
func NewSubscribedCatalog(ctx *pulumi.Context,
	name string, args *SubscribedCatalogArgs, opts ...pulumi.ResourceOption) (*SubscribedCatalog, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.SubscriptionUrl == nil {
		return nil, errors.New("invalid value for required argument 'SubscriptionUrl'")
	}
	if args.SubscriptionPassword != nil {
		args.SubscriptionPassword = pulumi.ToSecret(args.SubscriptionPassword).(pulumi.StringPtrInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"subscriptionPassword",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SubscribedCatalog
	err := ctx.RegisterResource("vcd:index/subscribedCatalog:SubscribedCatalog", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSubscribedCatalog gets an existing SubscribedCatalog resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSubscribedCatalog(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubscribedCatalogState, opts ...pulumi.ResourceOption) (*SubscribedCatalog, error) {
	var resource SubscribedCatalog
	err := ctx.ReadResource("vcd:index/subscribedCatalog:SubscribedCatalog", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SubscribedCatalog resources.
type subscribedCatalogState struct {
	// When `true`, the subscribed catalog will attempt canceling failed tasks.
	CancelFailedTasks *bool `pulumi:"cancelFailedTasks"`
	// Version number from this catalog. This is inherited from the publishing catalog and updated on sync.
	CatalogVersion *int `pulumi:"catalogVersion"`
	// Date and time of catalog creation. This is the creation date of the subscription, not the original published catalog.
	Created *string `pulumi:"created"`
	// When destroying use `delete_force=true` with `delete_recursive=true` to remove a catalog and any objects it contains, regardless of their state.
	DeleteForce *bool `pulumi:"deleteForce"`
	// When destroying use `delete_recursive=true` to remove the catalog and any objects it contains that are in a state that normally allows removal.
	DeleteRecursive *bool `pulumi:"deleteRecursive"`
	// Description of catalog. This is inherited from the publishing catalog and updated on sync.
	Description *string `pulumi:"description"`
	// List of synchronization tasks that are have failed. They can refer to the catalog or any of its catalog items.
	FailedTasks []string `pulumi:"failedTasks"`
	// the catalog's Hyper reference.
	Href *string `pulumi:"href"`
	// (*v3.8.1+*) Indicates if this catalog was created in the current organization.
	IsLocal *bool `pulumi:"isLocal"`
	// Indicates if this catalog is available for subscription. (Always false)
	IsPublished *bool `pulumi:"isPublished"`
	// Indicates if the catalog is shared.
	IsShared *bool `pulumi:"isShared"`
	// If `true`, subscription to a catalog creates a local copy of all items. Defaults to `false`, which does not create a local copy of catalog items unless a sync operation is performed.
	// It can only be `false` if the user configured in the provider is the System administrator.
	MakeLocalCopy *bool `pulumi:"makeLocalCopy"`
	// List of media item names in this catalog, in alphabetical order.
	MediaItemLists []string `pulumi:"mediaItemLists"`
	// (*Available until VCD 10.5*) Optional metadata of the catalog. This is inherited from the publishing catalog and updated on sync.
	Metadata map[string]string `pulumi:"metadata"`
	// Catalog name
	Name *string `pulumi:"name"`
	// Number of media items available in this catalog.
	NumberOfMedia *int `pulumi:"numberOfMedia"`
	// Number of vApp templates available in this catalog.
	NumberOfVappTemplates *int `pulumi:"numberOfVappTemplates"`
	// The name of organization to use, optional if defined at provider level.
	Org *string `pulumi:"org"`
	// Owner of the catalog.
	OwnerName *string `pulumi:"ownerName"`
	// Shows if the catalog is published, if it is a subscription from another one or none of those. (Always `SUBSCRIBED`)
	PublishSubscriptionType *string `pulumi:"publishSubscriptionType"`
	// List of running synchronization tasks that are still running. They can refer to the catalog or any of its catalog items.
	RunningTasks []string `pulumi:"runningTasks"`
	// Allows to set specific storage profile to be used for catalog.
	StorageProfileId *string `pulumi:"storageProfileId"`
	// if `true`, saves the list of tasks to a file for later update.
	StoreTasks *bool `pulumi:"storeTasks"`
	// An optional password to access the catalog. Only ASCII characters are allowed in a valid password.
	// The password is only required when set by the publishing catalog. Passing in six asterisks '******' indicates to keep current password.
	// Passing in an empty string indicates to remove password.
	SubscriptionPassword *string `pulumi:"subscriptionPassword"`
	// The URL to subscribe to the external catalog.
	SubscriptionUrl *string `pulumi:"subscriptionUrl"`
	// If `true`, synchronise this catalog and all items.
	SyncAll *bool `pulumi:"syncAll"`
	// If `true`, synchronise all media items. Not to be used when `syncAll` is set.
	SyncAllMediaItems *bool `pulumi:"syncAllMediaItems"`
	// If `true`, synchronise all vApp templates. Not to be used when `syncAll` is set.
	SyncAllVappTemplates *bool `pulumi:"syncAllVappTemplates"`
	// If `true`, synchronise this catalog. Not to be used when `syncAll` is set. This operation fetches the list of items. If `makeLocalCopy` is set, it also synchronises all the items.
	SyncCatalog *bool `pulumi:"syncCatalog"`
	// Synchronise a list of media items. Not to be used when `syncAll` or `syncAllMediaItems` are set.
	SyncMediaItems []string `pulumi:"syncMediaItems"`
	// Boolean value that shows if sync should be performed on every refresh.
	SyncOnRefresh *bool `pulumi:"syncOnRefresh"`
	// Synchronise a list of vApp templates. Not to be used when `syncAll` or `syncAllVappTemplates` are set.
	SyncVappTemplates []string `pulumi:"syncVappTemplates"`
	// Where the running tasks IDs have been stored. Only if `storeTasks` is set.
	TasksFileName *string `pulumi:"tasksFileName"`
	// List of vApp template names in this catalog, in alphabetical order.
	VappTemplateLists []string `pulumi:"vappTemplateLists"`
}

type SubscribedCatalogState struct {
	// When `true`, the subscribed catalog will attempt canceling failed tasks.
	CancelFailedTasks pulumi.BoolPtrInput
	// Version number from this catalog. This is inherited from the publishing catalog and updated on sync.
	CatalogVersion pulumi.IntPtrInput
	// Date and time of catalog creation. This is the creation date of the subscription, not the original published catalog.
	Created pulumi.StringPtrInput
	// When destroying use `delete_force=true` with `delete_recursive=true` to remove a catalog and any objects it contains, regardless of their state.
	DeleteForce pulumi.BoolPtrInput
	// When destroying use `delete_recursive=true` to remove the catalog and any objects it contains that are in a state that normally allows removal.
	DeleteRecursive pulumi.BoolPtrInput
	// Description of catalog. This is inherited from the publishing catalog and updated on sync.
	Description pulumi.StringPtrInput
	// List of synchronization tasks that are have failed. They can refer to the catalog or any of its catalog items.
	FailedTasks pulumi.StringArrayInput
	// the catalog's Hyper reference.
	Href pulumi.StringPtrInput
	// (*v3.8.1+*) Indicates if this catalog was created in the current organization.
	IsLocal pulumi.BoolPtrInput
	// Indicates if this catalog is available for subscription. (Always false)
	IsPublished pulumi.BoolPtrInput
	// Indicates if the catalog is shared.
	IsShared pulumi.BoolPtrInput
	// If `true`, subscription to a catalog creates a local copy of all items. Defaults to `false`, which does not create a local copy of catalog items unless a sync operation is performed.
	// It can only be `false` if the user configured in the provider is the System administrator.
	MakeLocalCopy pulumi.BoolPtrInput
	// List of media item names in this catalog, in alphabetical order.
	MediaItemLists pulumi.StringArrayInput
	// (*Available until VCD 10.5*) Optional metadata of the catalog. This is inherited from the publishing catalog and updated on sync.
	Metadata pulumi.StringMapInput
	// Catalog name
	Name pulumi.StringPtrInput
	// Number of media items available in this catalog.
	NumberOfMedia pulumi.IntPtrInput
	// Number of vApp templates available in this catalog.
	NumberOfVappTemplates pulumi.IntPtrInput
	// The name of organization to use, optional if defined at provider level.
	Org pulumi.StringPtrInput
	// Owner of the catalog.
	OwnerName pulumi.StringPtrInput
	// Shows if the catalog is published, if it is a subscription from another one or none of those. (Always `SUBSCRIBED`)
	PublishSubscriptionType pulumi.StringPtrInput
	// List of running synchronization tasks that are still running. They can refer to the catalog or any of its catalog items.
	RunningTasks pulumi.StringArrayInput
	// Allows to set specific storage profile to be used for catalog.
	StorageProfileId pulumi.StringPtrInput
	// if `true`, saves the list of tasks to a file for later update.
	StoreTasks pulumi.BoolPtrInput
	// An optional password to access the catalog. Only ASCII characters are allowed in a valid password.
	// The password is only required when set by the publishing catalog. Passing in six asterisks '******' indicates to keep current password.
	// Passing in an empty string indicates to remove password.
	SubscriptionPassword pulumi.StringPtrInput
	// The URL to subscribe to the external catalog.
	SubscriptionUrl pulumi.StringPtrInput
	// If `true`, synchronise this catalog and all items.
	SyncAll pulumi.BoolPtrInput
	// If `true`, synchronise all media items. Not to be used when `syncAll` is set.
	SyncAllMediaItems pulumi.BoolPtrInput
	// If `true`, synchronise all vApp templates. Not to be used when `syncAll` is set.
	SyncAllVappTemplates pulumi.BoolPtrInput
	// If `true`, synchronise this catalog. Not to be used when `syncAll` is set. This operation fetches the list of items. If `makeLocalCopy` is set, it also synchronises all the items.
	SyncCatalog pulumi.BoolPtrInput
	// Synchronise a list of media items. Not to be used when `syncAll` or `syncAllMediaItems` are set.
	SyncMediaItems pulumi.StringArrayInput
	// Boolean value that shows if sync should be performed on every refresh.
	SyncOnRefresh pulumi.BoolPtrInput
	// Synchronise a list of vApp templates. Not to be used when `syncAll` or `syncAllVappTemplates` are set.
	SyncVappTemplates pulumi.StringArrayInput
	// Where the running tasks IDs have been stored. Only if `storeTasks` is set.
	TasksFileName pulumi.StringPtrInput
	// List of vApp template names in this catalog, in alphabetical order.
	VappTemplateLists pulumi.StringArrayInput
}

func (SubscribedCatalogState) ElementType() reflect.Type {
	return reflect.TypeOf((*subscribedCatalogState)(nil)).Elem()
}

type subscribedCatalogArgs struct {
	// When `true`, the subscribed catalog will attempt canceling failed tasks.
	CancelFailedTasks *bool `pulumi:"cancelFailedTasks"`
	// When destroying use `delete_force=true` with `delete_recursive=true` to remove a catalog and any objects it contains, regardless of their state.
	DeleteForce *bool `pulumi:"deleteForce"`
	// When destroying use `delete_recursive=true` to remove the catalog and any objects it contains that are in a state that normally allows removal.
	DeleteRecursive *bool `pulumi:"deleteRecursive"`
	// If `true`, subscription to a catalog creates a local copy of all items. Defaults to `false`, which does not create a local copy of catalog items unless a sync operation is performed.
	// It can only be `false` if the user configured in the provider is the System administrator.
	MakeLocalCopy *bool `pulumi:"makeLocalCopy"`
	// Catalog name
	Name *string `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level.
	Org *string `pulumi:"org"`
	// Allows to set specific storage profile to be used for catalog.
	StorageProfileId *string `pulumi:"storageProfileId"`
	// if `true`, saves the list of tasks to a file for later update.
	StoreTasks *bool `pulumi:"storeTasks"`
	// An optional password to access the catalog. Only ASCII characters are allowed in a valid password.
	// The password is only required when set by the publishing catalog. Passing in six asterisks '******' indicates to keep current password.
	// Passing in an empty string indicates to remove password.
	SubscriptionPassword *string `pulumi:"subscriptionPassword"`
	// The URL to subscribe to the external catalog.
	SubscriptionUrl string `pulumi:"subscriptionUrl"`
	// If `true`, synchronise this catalog and all items.
	SyncAll *bool `pulumi:"syncAll"`
	// If `true`, synchronise all media items. Not to be used when `syncAll` is set.
	SyncAllMediaItems *bool `pulumi:"syncAllMediaItems"`
	// If `true`, synchronise all vApp templates. Not to be used when `syncAll` is set.
	SyncAllVappTemplates *bool `pulumi:"syncAllVappTemplates"`
	// If `true`, synchronise this catalog. Not to be used when `syncAll` is set. This operation fetches the list of items. If `makeLocalCopy` is set, it also synchronises all the items.
	SyncCatalog *bool `pulumi:"syncCatalog"`
	// Synchronise a list of media items. Not to be used when `syncAll` or `syncAllMediaItems` are set.
	SyncMediaItems []string `pulumi:"syncMediaItems"`
	// Boolean value that shows if sync should be performed on every refresh.
	SyncOnRefresh *bool `pulumi:"syncOnRefresh"`
	// Synchronise a list of vApp templates. Not to be used when `syncAll` or `syncAllVappTemplates` are set.
	SyncVappTemplates []string `pulumi:"syncVappTemplates"`
}

// The set of arguments for constructing a SubscribedCatalog resource.
type SubscribedCatalogArgs struct {
	// When `true`, the subscribed catalog will attempt canceling failed tasks.
	CancelFailedTasks pulumi.BoolPtrInput
	// When destroying use `delete_force=true` with `delete_recursive=true` to remove a catalog and any objects it contains, regardless of their state.
	DeleteForce pulumi.BoolPtrInput
	// When destroying use `delete_recursive=true` to remove the catalog and any objects it contains that are in a state that normally allows removal.
	DeleteRecursive pulumi.BoolPtrInput
	// If `true`, subscription to a catalog creates a local copy of all items. Defaults to `false`, which does not create a local copy of catalog items unless a sync operation is performed.
	// It can only be `false` if the user configured in the provider is the System administrator.
	MakeLocalCopy pulumi.BoolPtrInput
	// Catalog name
	Name pulumi.StringPtrInput
	// The name of organization to use, optional if defined at provider level.
	Org pulumi.StringPtrInput
	// Allows to set specific storage profile to be used for catalog.
	StorageProfileId pulumi.StringPtrInput
	// if `true`, saves the list of tasks to a file for later update.
	StoreTasks pulumi.BoolPtrInput
	// An optional password to access the catalog. Only ASCII characters are allowed in a valid password.
	// The password is only required when set by the publishing catalog. Passing in six asterisks '******' indicates to keep current password.
	// Passing in an empty string indicates to remove password.
	SubscriptionPassword pulumi.StringPtrInput
	// The URL to subscribe to the external catalog.
	SubscriptionUrl pulumi.StringInput
	// If `true`, synchronise this catalog and all items.
	SyncAll pulumi.BoolPtrInput
	// If `true`, synchronise all media items. Not to be used when `syncAll` is set.
	SyncAllMediaItems pulumi.BoolPtrInput
	// If `true`, synchronise all vApp templates. Not to be used when `syncAll` is set.
	SyncAllVappTemplates pulumi.BoolPtrInput
	// If `true`, synchronise this catalog. Not to be used when `syncAll` is set. This operation fetches the list of items. If `makeLocalCopy` is set, it also synchronises all the items.
	SyncCatalog pulumi.BoolPtrInput
	// Synchronise a list of media items. Not to be used when `syncAll` or `syncAllMediaItems` are set.
	SyncMediaItems pulumi.StringArrayInput
	// Boolean value that shows if sync should be performed on every refresh.
	SyncOnRefresh pulumi.BoolPtrInput
	// Synchronise a list of vApp templates. Not to be used when `syncAll` or `syncAllVappTemplates` are set.
	SyncVappTemplates pulumi.StringArrayInput
}

func (SubscribedCatalogArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subscribedCatalogArgs)(nil)).Elem()
}

type SubscribedCatalogInput interface {
	pulumi.Input

	ToSubscribedCatalogOutput() SubscribedCatalogOutput
	ToSubscribedCatalogOutputWithContext(ctx context.Context) SubscribedCatalogOutput
}

func (*SubscribedCatalog) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscribedCatalog)(nil)).Elem()
}

func (i *SubscribedCatalog) ToSubscribedCatalogOutput() SubscribedCatalogOutput {
	return i.ToSubscribedCatalogOutputWithContext(context.Background())
}

func (i *SubscribedCatalog) ToSubscribedCatalogOutputWithContext(ctx context.Context) SubscribedCatalogOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscribedCatalogOutput)
}

// SubscribedCatalogArrayInput is an input type that accepts SubscribedCatalogArray and SubscribedCatalogArrayOutput values.
// You can construct a concrete instance of `SubscribedCatalogArrayInput` via:
//
//	SubscribedCatalogArray{ SubscribedCatalogArgs{...} }
type SubscribedCatalogArrayInput interface {
	pulumi.Input

	ToSubscribedCatalogArrayOutput() SubscribedCatalogArrayOutput
	ToSubscribedCatalogArrayOutputWithContext(context.Context) SubscribedCatalogArrayOutput
}

type SubscribedCatalogArray []SubscribedCatalogInput

func (SubscribedCatalogArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SubscribedCatalog)(nil)).Elem()
}

func (i SubscribedCatalogArray) ToSubscribedCatalogArrayOutput() SubscribedCatalogArrayOutput {
	return i.ToSubscribedCatalogArrayOutputWithContext(context.Background())
}

func (i SubscribedCatalogArray) ToSubscribedCatalogArrayOutputWithContext(ctx context.Context) SubscribedCatalogArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscribedCatalogArrayOutput)
}

// SubscribedCatalogMapInput is an input type that accepts SubscribedCatalogMap and SubscribedCatalogMapOutput values.
// You can construct a concrete instance of `SubscribedCatalogMapInput` via:
//
//	SubscribedCatalogMap{ "key": SubscribedCatalogArgs{...} }
type SubscribedCatalogMapInput interface {
	pulumi.Input

	ToSubscribedCatalogMapOutput() SubscribedCatalogMapOutput
	ToSubscribedCatalogMapOutputWithContext(context.Context) SubscribedCatalogMapOutput
}

type SubscribedCatalogMap map[string]SubscribedCatalogInput

func (SubscribedCatalogMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SubscribedCatalog)(nil)).Elem()
}

func (i SubscribedCatalogMap) ToSubscribedCatalogMapOutput() SubscribedCatalogMapOutput {
	return i.ToSubscribedCatalogMapOutputWithContext(context.Background())
}

func (i SubscribedCatalogMap) ToSubscribedCatalogMapOutputWithContext(ctx context.Context) SubscribedCatalogMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscribedCatalogMapOutput)
}

type SubscribedCatalogOutput struct{ *pulumi.OutputState }

func (SubscribedCatalogOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscribedCatalog)(nil)).Elem()
}

func (o SubscribedCatalogOutput) ToSubscribedCatalogOutput() SubscribedCatalogOutput {
	return o
}

func (o SubscribedCatalogOutput) ToSubscribedCatalogOutputWithContext(ctx context.Context) SubscribedCatalogOutput {
	return o
}

// When `true`, the subscribed catalog will attempt canceling failed tasks.
func (o SubscribedCatalogOutput) CancelFailedTasks() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.BoolPtrOutput { return v.CancelFailedTasks }).(pulumi.BoolPtrOutput)
}

// Version number from this catalog. This is inherited from the publishing catalog and updated on sync.
func (o SubscribedCatalogOutput) CatalogVersion() pulumi.IntOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.IntOutput { return v.CatalogVersion }).(pulumi.IntOutput)
}

// Date and time of catalog creation. This is the creation date of the subscription, not the original published catalog.
func (o SubscribedCatalogOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.StringOutput { return v.Created }).(pulumi.StringOutput)
}

// When destroying use `delete_force=true` with `delete_recursive=true` to remove a catalog and any objects it contains, regardless of their state.
func (o SubscribedCatalogOutput) DeleteForce() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.BoolPtrOutput { return v.DeleteForce }).(pulumi.BoolPtrOutput)
}

// When destroying use `delete_recursive=true` to remove the catalog and any objects it contains that are in a state that normally allows removal.
func (o SubscribedCatalogOutput) DeleteRecursive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.BoolPtrOutput { return v.DeleteRecursive }).(pulumi.BoolPtrOutput)
}

// Description of catalog. This is inherited from the publishing catalog and updated on sync.
func (o SubscribedCatalogOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// List of synchronization tasks that are have failed. They can refer to the catalog or any of its catalog items.
func (o SubscribedCatalogOutput) FailedTasks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.StringArrayOutput { return v.FailedTasks }).(pulumi.StringArrayOutput)
}

// the catalog's Hyper reference.
func (o SubscribedCatalogOutput) Href() pulumi.StringOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.StringOutput { return v.Href }).(pulumi.StringOutput)
}

// (*v3.8.1+*) Indicates if this catalog was created in the current organization.
func (o SubscribedCatalogOutput) IsLocal() pulumi.BoolOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.BoolOutput { return v.IsLocal }).(pulumi.BoolOutput)
}

// Indicates if this catalog is available for subscription. (Always false)
func (o SubscribedCatalogOutput) IsPublished() pulumi.BoolOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.BoolOutput { return v.IsPublished }).(pulumi.BoolOutput)
}

// Indicates if the catalog is shared.
func (o SubscribedCatalogOutput) IsShared() pulumi.BoolOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.BoolOutput { return v.IsShared }).(pulumi.BoolOutput)
}

// If `true`, subscription to a catalog creates a local copy of all items. Defaults to `false`, which does not create a local copy of catalog items unless a sync operation is performed.
// It can only be `false` if the user configured in the provider is the System administrator.
func (o SubscribedCatalogOutput) MakeLocalCopy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.BoolPtrOutput { return v.MakeLocalCopy }).(pulumi.BoolPtrOutput)
}

// List of media item names in this catalog, in alphabetical order.
func (o SubscribedCatalogOutput) MediaItemLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.StringArrayOutput { return v.MediaItemLists }).(pulumi.StringArrayOutput)
}

// (*Available until VCD 10.5*) Optional metadata of the catalog. This is inherited from the publishing catalog and updated on sync.
func (o SubscribedCatalogOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.StringMapOutput { return v.Metadata }).(pulumi.StringMapOutput)
}

// Catalog name
func (o SubscribedCatalogOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Number of media items available in this catalog.
func (o SubscribedCatalogOutput) NumberOfMedia() pulumi.IntOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.IntOutput { return v.NumberOfMedia }).(pulumi.IntOutput)
}

// Number of vApp templates available in this catalog.
func (o SubscribedCatalogOutput) NumberOfVappTemplates() pulumi.IntOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.IntOutput { return v.NumberOfVappTemplates }).(pulumi.IntOutput)
}

// The name of organization to use, optional if defined at provider level.
func (o SubscribedCatalogOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.StringPtrOutput { return v.Org }).(pulumi.StringPtrOutput)
}

// Owner of the catalog.
func (o SubscribedCatalogOutput) OwnerName() pulumi.StringOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.StringOutput { return v.OwnerName }).(pulumi.StringOutput)
}

// Shows if the catalog is published, if it is a subscription from another one or none of those. (Always `SUBSCRIBED`)
func (o SubscribedCatalogOutput) PublishSubscriptionType() pulumi.StringOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.StringOutput { return v.PublishSubscriptionType }).(pulumi.StringOutput)
}

// List of running synchronization tasks that are still running. They can refer to the catalog or any of its catalog items.
func (o SubscribedCatalogOutput) RunningTasks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.StringArrayOutput { return v.RunningTasks }).(pulumi.StringArrayOutput)
}

// Allows to set specific storage profile to be used for catalog.
func (o SubscribedCatalogOutput) StorageProfileId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.StringPtrOutput { return v.StorageProfileId }).(pulumi.StringPtrOutput)
}

// if `true`, saves the list of tasks to a file for later update.
func (o SubscribedCatalogOutput) StoreTasks() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.BoolPtrOutput { return v.StoreTasks }).(pulumi.BoolPtrOutput)
}

// An optional password to access the catalog. Only ASCII characters are allowed in a valid password.
// The password is only required when set by the publishing catalog. Passing in six asterisks '******' indicates to keep current password.
// Passing in an empty string indicates to remove password.
func (o SubscribedCatalogOutput) SubscriptionPassword() pulumi.StringOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.StringOutput { return v.SubscriptionPassword }).(pulumi.StringOutput)
}

// The URL to subscribe to the external catalog.
func (o SubscribedCatalogOutput) SubscriptionUrl() pulumi.StringOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.StringOutput { return v.SubscriptionUrl }).(pulumi.StringOutput)
}

// If `true`, synchronise this catalog and all items.
func (o SubscribedCatalogOutput) SyncAll() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.BoolPtrOutput { return v.SyncAll }).(pulumi.BoolPtrOutput)
}

// If `true`, synchronise all media items. Not to be used when `syncAll` is set.
func (o SubscribedCatalogOutput) SyncAllMediaItems() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.BoolPtrOutput { return v.SyncAllMediaItems }).(pulumi.BoolPtrOutput)
}

// If `true`, synchronise all vApp templates. Not to be used when `syncAll` is set.
func (o SubscribedCatalogOutput) SyncAllVappTemplates() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.BoolPtrOutput { return v.SyncAllVappTemplates }).(pulumi.BoolPtrOutput)
}

// If `true`, synchronise this catalog. Not to be used when `syncAll` is set. This operation fetches the list of items. If `makeLocalCopy` is set, it also synchronises all the items.
func (o SubscribedCatalogOutput) SyncCatalog() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.BoolPtrOutput { return v.SyncCatalog }).(pulumi.BoolPtrOutput)
}

// Synchronise a list of media items. Not to be used when `syncAll` or `syncAllMediaItems` are set.
func (o SubscribedCatalogOutput) SyncMediaItems() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.StringArrayOutput { return v.SyncMediaItems }).(pulumi.StringArrayOutput)
}

// Boolean value that shows if sync should be performed on every refresh.
func (o SubscribedCatalogOutput) SyncOnRefresh() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.BoolPtrOutput { return v.SyncOnRefresh }).(pulumi.BoolPtrOutput)
}

// Synchronise a list of vApp templates. Not to be used when `syncAll` or `syncAllVappTemplates` are set.
func (o SubscribedCatalogOutput) SyncVappTemplates() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.StringArrayOutput { return v.SyncVappTemplates }).(pulumi.StringArrayOutput)
}

// Where the running tasks IDs have been stored. Only if `storeTasks` is set.
func (o SubscribedCatalogOutput) TasksFileName() pulumi.StringOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.StringOutput { return v.TasksFileName }).(pulumi.StringOutput)
}

// List of vApp template names in this catalog, in alphabetical order.
func (o SubscribedCatalogOutput) VappTemplateLists() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SubscribedCatalog) pulumi.StringArrayOutput { return v.VappTemplateLists }).(pulumi.StringArrayOutput)
}

type SubscribedCatalogArrayOutput struct{ *pulumi.OutputState }

func (SubscribedCatalogArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SubscribedCatalog)(nil)).Elem()
}

func (o SubscribedCatalogArrayOutput) ToSubscribedCatalogArrayOutput() SubscribedCatalogArrayOutput {
	return o
}

func (o SubscribedCatalogArrayOutput) ToSubscribedCatalogArrayOutputWithContext(ctx context.Context) SubscribedCatalogArrayOutput {
	return o
}

func (o SubscribedCatalogArrayOutput) Index(i pulumi.IntInput) SubscribedCatalogOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SubscribedCatalog {
		return vs[0].([]*SubscribedCatalog)[vs[1].(int)]
	}).(SubscribedCatalogOutput)
}

type SubscribedCatalogMapOutput struct{ *pulumi.OutputState }

func (SubscribedCatalogMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SubscribedCatalog)(nil)).Elem()
}

func (o SubscribedCatalogMapOutput) ToSubscribedCatalogMapOutput() SubscribedCatalogMapOutput {
	return o
}

func (o SubscribedCatalogMapOutput) ToSubscribedCatalogMapOutputWithContext(ctx context.Context) SubscribedCatalogMapOutput {
	return o
}

func (o SubscribedCatalogMapOutput) MapIndex(k pulumi.StringInput) SubscribedCatalogOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SubscribedCatalog {
		return vs[0].(map[string]*SubscribedCatalog)[vs[1].(string)]
	}).(SubscribedCatalogOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SubscribedCatalogInput)(nil)).Elem(), &SubscribedCatalog{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubscribedCatalogArrayInput)(nil)).Elem(), SubscribedCatalogArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SubscribedCatalogMapInput)(nil)).Elem(), SubscribedCatalogMap{})
	pulumi.RegisterOutputType(SubscribedCatalogOutput{})
	pulumi.RegisterOutputType(SubscribedCatalogArrayOutput{})
	pulumi.RegisterOutputType(SubscribedCatalogMapOutput{})
}
