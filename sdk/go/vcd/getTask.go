// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package vcd

import (
	"context"
	"reflect"

	"github.com/ergSey/pulumi-vcd/sdk/go/vcd/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a data source for available tasks.
//
// Supported in provider *v3.8+*
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
//			some_task, err := vcd.GetTask(ctx, &vcd.GetTaskArgs{
//				Id: "d4fdcaa9-8db4-45a9-80b8-69de49901bc7",
//			}, nil)
//			if err != nil {
//				return err
//			}
//			ctx.Export("some-task", some_task)
//			return nil
//		})
//	}
//
// ```
func GetTask(ctx *pulumi.Context, args *GetTaskArgs, opts ...pulumi.InvokeOption) (*GetTaskResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetTaskResult
	err := ctx.Invoke("vcd:index/getTask:getTask", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getTask.
type GetTaskArgs struct {
	// The ID of the task
	Id string `pulumi:"id"`
}

// A collection of values returned by getTask.
type GetTaskResult struct {
	// Whether user has requested this processing to be canceled (`true` or `false`).
	CancelRequested bool `pulumi:"cancelRequested"`
	// An optional description of the task.
	Description string `pulumi:"description"`
	// The date and time that processing of the task was completed. May not be present if the task is still being executed.
	EndTime string `pulumi:"endTime"`
	// error information from a failed task.
	Error string `pulumi:"error"`
	// The date and time at which the task resource will be destroyed and no longer available for retrieval. May not be present if the task has not been executed or is still being executed.
	ExpiryTime string `pulumi:"expiryTime"`
	// The URI of the task.
	Href string `pulumi:"href"`
	Id   string `pulumi:"id"`
	// Name of the task. May not be unique. Defines the general operation being performed.
	Name string `pulumi:"name"`
	// A message describing the operation that is tracked by this task.
	Operation string `pulumi:"operation"`
	// The short name of the operation that is tracked by this task.
	OperationName string `pulumi:"operationName"`
	// The unique identifier of the user org.
	OrgId string `pulumi:"orgId"`
	// The name of the org to which the user belongs.
	OrgName string `pulumi:"orgName"`
	// The unique identifier of the task owner.
	OwnerId string `pulumi:"ownerId"`
	// The name of the task owner. This is typically the object that the task is creating or updating.
	OwnerName string `pulumi:"ownerName"`
	// The type of the task owner.
	OwnerType string `pulumi:"ownerType"`
	// Indicator of task progress as an approximate percentage between 0 and 100. Not available for all tasks.
	Progress int `pulumi:"progress"`
	// The date and time the system started executing the task. May not be present if the task has not been executed yet.
	StartTime string `pulumi:"startTime"`
	// The execution status of the task. One of queued, preRunning, running, success, error, aborted.
	Status string `pulumi:"status"`
	// Type of the task.
	Type string `pulumi:"type"`
	// The unique identifier of the task user.
	UserId string `pulumi:"userId"`
	// The name of the user who started the task.
	UserName string `pulumi:"userName"`
}

func GetTaskOutput(ctx *pulumi.Context, args GetTaskOutputArgs, opts ...pulumi.InvokeOption) GetTaskResultOutput {
	return pulumi.ToOutputWithContext(ctx.Context(), args).
		ApplyT(func(v interface{}) (GetTaskResultOutput, error) {
			args := v.(GetTaskArgs)
			options := pulumi.InvokeOutputOptions{InvokeOptions: internal.PkgInvokeDefaultOpts(opts)}
			return ctx.InvokeOutput("vcd:index/getTask:getTask", args, GetTaskResultOutput{}, options).(GetTaskResultOutput), nil
		}).(GetTaskResultOutput)
}

// A collection of arguments for invoking getTask.
type GetTaskOutputArgs struct {
	// The ID of the task
	Id pulumi.StringInput `pulumi:"id"`
}

func (GetTaskOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTaskArgs)(nil)).Elem()
}

// A collection of values returned by getTask.
type GetTaskResultOutput struct{ *pulumi.OutputState }

func (GetTaskResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetTaskResult)(nil)).Elem()
}

func (o GetTaskResultOutput) ToGetTaskResultOutput() GetTaskResultOutput {
	return o
}

func (o GetTaskResultOutput) ToGetTaskResultOutputWithContext(ctx context.Context) GetTaskResultOutput {
	return o
}

// Whether user has requested this processing to be canceled (`true` or `false`).
func (o GetTaskResultOutput) CancelRequested() pulumi.BoolOutput {
	return o.ApplyT(func(v GetTaskResult) bool { return v.CancelRequested }).(pulumi.BoolOutput)
}

// An optional description of the task.
func (o GetTaskResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v GetTaskResult) string { return v.Description }).(pulumi.StringOutput)
}

// The date and time that processing of the task was completed. May not be present if the task is still being executed.
func (o GetTaskResultOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetTaskResult) string { return v.EndTime }).(pulumi.StringOutput)
}

// error information from a failed task.
func (o GetTaskResultOutput) Error() pulumi.StringOutput {
	return o.ApplyT(func(v GetTaskResult) string { return v.Error }).(pulumi.StringOutput)
}

// The date and time at which the task resource will be destroyed and no longer available for retrieval. May not be present if the task has not been executed or is still being executed.
func (o GetTaskResultOutput) ExpiryTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetTaskResult) string { return v.ExpiryTime }).(pulumi.StringOutput)
}

// The URI of the task.
func (o GetTaskResultOutput) Href() pulumi.StringOutput {
	return o.ApplyT(func(v GetTaskResult) string { return v.Href }).(pulumi.StringOutput)
}

func (o GetTaskResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetTaskResult) string { return v.Id }).(pulumi.StringOutput)
}

// Name of the task. May not be unique. Defines the general operation being performed.
func (o GetTaskResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetTaskResult) string { return v.Name }).(pulumi.StringOutput)
}

// A message describing the operation that is tracked by this task.
func (o GetTaskResultOutput) Operation() pulumi.StringOutput {
	return o.ApplyT(func(v GetTaskResult) string { return v.Operation }).(pulumi.StringOutput)
}

// The short name of the operation that is tracked by this task.
func (o GetTaskResultOutput) OperationName() pulumi.StringOutput {
	return o.ApplyT(func(v GetTaskResult) string { return v.OperationName }).(pulumi.StringOutput)
}

// The unique identifier of the user org.
func (o GetTaskResultOutput) OrgId() pulumi.StringOutput {
	return o.ApplyT(func(v GetTaskResult) string { return v.OrgId }).(pulumi.StringOutput)
}

// The name of the org to which the user belongs.
func (o GetTaskResultOutput) OrgName() pulumi.StringOutput {
	return o.ApplyT(func(v GetTaskResult) string { return v.OrgName }).(pulumi.StringOutput)
}

// The unique identifier of the task owner.
func (o GetTaskResultOutput) OwnerId() pulumi.StringOutput {
	return o.ApplyT(func(v GetTaskResult) string { return v.OwnerId }).(pulumi.StringOutput)
}

// The name of the task owner. This is typically the object that the task is creating or updating.
func (o GetTaskResultOutput) OwnerName() pulumi.StringOutput {
	return o.ApplyT(func(v GetTaskResult) string { return v.OwnerName }).(pulumi.StringOutput)
}

// The type of the task owner.
func (o GetTaskResultOutput) OwnerType() pulumi.StringOutput {
	return o.ApplyT(func(v GetTaskResult) string { return v.OwnerType }).(pulumi.StringOutput)
}

// Indicator of task progress as an approximate percentage between 0 and 100. Not available for all tasks.
func (o GetTaskResultOutput) Progress() pulumi.IntOutput {
	return o.ApplyT(func(v GetTaskResult) int { return v.Progress }).(pulumi.IntOutput)
}

// The date and time the system started executing the task. May not be present if the task has not been executed yet.
func (o GetTaskResultOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v GetTaskResult) string { return v.StartTime }).(pulumi.StringOutput)
}

// The execution status of the task. One of queued, preRunning, running, success, error, aborted.
func (o GetTaskResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetTaskResult) string { return v.Status }).(pulumi.StringOutput)
}

// Type of the task.
func (o GetTaskResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetTaskResult) string { return v.Type }).(pulumi.StringOutput)
}

// The unique identifier of the task user.
func (o GetTaskResultOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v GetTaskResult) string { return v.UserId }).(pulumi.StringOutput)
}

// The name of the user who started the task.
func (o GetTaskResultOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v GetTaskResult) string { return v.UserName }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetTaskResultOutput{})
}
