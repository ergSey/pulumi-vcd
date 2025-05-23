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

type NsxtAlbController struct {
	pulumi.CustomResourceState

	// An optional description ALB Controller
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// License type of ALB Controller (`ENTERPRISE` or `BASIC`)
	//
	// > The attribute `licenseType` must not be used in VCD 10.4+, it is replaced by [nsxtAlbServiceEngineGroup](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_service_engine_group) and [nsxtAlbSettings](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_settings) attribute `supportedFeatureSet`.
	LicenseType pulumi.StringPtrOutput `pulumi:"licenseType"`
	// A name for ALB Controller
	Name pulumi.StringOutput `pulumi:"name"`
	// The password for ALB Controller. Password will not be refreshed.
	Password pulumi.StringOutput `pulumi:"password"`
	// The URL of ALB Controller
	Url pulumi.StringOutput `pulumi:"url"`
	// The username for ALB Controller
	Username pulumi.StringOutput `pulumi:"username"`
	// ALB Controller version (e.g. 20.1.3)
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewNsxtAlbController registers a new resource with the given unique name, arguments, and options.
func NewNsxtAlbController(ctx *pulumi.Context,
	name string, args *NsxtAlbControllerArgs, opts ...pulumi.ResourceOption) (*NsxtAlbController, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Password == nil {
		return nil, errors.New("invalid value for required argument 'Password'")
	}
	if args.Url == nil {
		return nil, errors.New("invalid value for required argument 'Url'")
	}
	if args.Username == nil {
		return nil, errors.New("invalid value for required argument 'Username'")
	}
	if args.Password != nil {
		args.Password = pulumi.ToSecret(args.Password).(pulumi.StringInput)
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"password",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NsxtAlbController
	err := ctx.RegisterResource("vcd:index/nsxtAlbController:NsxtAlbController", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNsxtAlbController gets an existing NsxtAlbController resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNsxtAlbController(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NsxtAlbControllerState, opts ...pulumi.ResourceOption) (*NsxtAlbController, error) {
	var resource NsxtAlbController
	err := ctx.ReadResource("vcd:index/nsxtAlbController:NsxtAlbController", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NsxtAlbController resources.
type nsxtAlbControllerState struct {
	// An optional description ALB Controller
	Description *string `pulumi:"description"`
	// License type of ALB Controller (`ENTERPRISE` or `BASIC`)
	//
	// > The attribute `licenseType` must not be used in VCD 10.4+, it is replaced by [nsxtAlbServiceEngineGroup](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_service_engine_group) and [nsxtAlbSettings](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_settings) attribute `supportedFeatureSet`.
	LicenseType *string `pulumi:"licenseType"`
	// A name for ALB Controller
	Name *string `pulumi:"name"`
	// The password for ALB Controller. Password will not be refreshed.
	Password *string `pulumi:"password"`
	// The URL of ALB Controller
	Url *string `pulumi:"url"`
	// The username for ALB Controller
	Username *string `pulumi:"username"`
	// ALB Controller version (e.g. 20.1.3)
	Version *string `pulumi:"version"`
}

type NsxtAlbControllerState struct {
	// An optional description ALB Controller
	Description pulumi.StringPtrInput
	// License type of ALB Controller (`ENTERPRISE` or `BASIC`)
	//
	// > The attribute `licenseType` must not be used in VCD 10.4+, it is replaced by [nsxtAlbServiceEngineGroup](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_service_engine_group) and [nsxtAlbSettings](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_settings) attribute `supportedFeatureSet`.
	LicenseType pulumi.StringPtrInput
	// A name for ALB Controller
	Name pulumi.StringPtrInput
	// The password for ALB Controller. Password will not be refreshed.
	Password pulumi.StringPtrInput
	// The URL of ALB Controller
	Url pulumi.StringPtrInput
	// The username for ALB Controller
	Username pulumi.StringPtrInput
	// ALB Controller version (e.g. 20.1.3)
	Version pulumi.StringPtrInput
}

func (NsxtAlbControllerState) ElementType() reflect.Type {
	return reflect.TypeOf((*nsxtAlbControllerState)(nil)).Elem()
}

type nsxtAlbControllerArgs struct {
	// An optional description ALB Controller
	Description *string `pulumi:"description"`
	// License type of ALB Controller (`ENTERPRISE` or `BASIC`)
	//
	// > The attribute `licenseType` must not be used in VCD 10.4+, it is replaced by [nsxtAlbServiceEngineGroup](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_service_engine_group) and [nsxtAlbSettings](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_settings) attribute `supportedFeatureSet`.
	LicenseType *string `pulumi:"licenseType"`
	// A name for ALB Controller
	Name *string `pulumi:"name"`
	// The password for ALB Controller. Password will not be refreshed.
	Password string `pulumi:"password"`
	// The URL of ALB Controller
	Url string `pulumi:"url"`
	// The username for ALB Controller
	Username string `pulumi:"username"`
}

// The set of arguments for constructing a NsxtAlbController resource.
type NsxtAlbControllerArgs struct {
	// An optional description ALB Controller
	Description pulumi.StringPtrInput
	// License type of ALB Controller (`ENTERPRISE` or `BASIC`)
	//
	// > The attribute `licenseType` must not be used in VCD 10.4+, it is replaced by [nsxtAlbServiceEngineGroup](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_service_engine_group) and [nsxtAlbSettings](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_settings) attribute `supportedFeatureSet`.
	LicenseType pulumi.StringPtrInput
	// A name for ALB Controller
	Name pulumi.StringPtrInput
	// The password for ALB Controller. Password will not be refreshed.
	Password pulumi.StringInput
	// The URL of ALB Controller
	Url pulumi.StringInput
	// The username for ALB Controller
	Username pulumi.StringInput
}

func (NsxtAlbControllerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nsxtAlbControllerArgs)(nil)).Elem()
}

type NsxtAlbControllerInput interface {
	pulumi.Input

	ToNsxtAlbControllerOutput() NsxtAlbControllerOutput
	ToNsxtAlbControllerOutputWithContext(ctx context.Context) NsxtAlbControllerOutput
}

func (*NsxtAlbController) ElementType() reflect.Type {
	return reflect.TypeOf((**NsxtAlbController)(nil)).Elem()
}

func (i *NsxtAlbController) ToNsxtAlbControllerOutput() NsxtAlbControllerOutput {
	return i.ToNsxtAlbControllerOutputWithContext(context.Background())
}

func (i *NsxtAlbController) ToNsxtAlbControllerOutputWithContext(ctx context.Context) NsxtAlbControllerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxtAlbControllerOutput)
}

// NsxtAlbControllerArrayInput is an input type that accepts NsxtAlbControllerArray and NsxtAlbControllerArrayOutput values.
// You can construct a concrete instance of `NsxtAlbControllerArrayInput` via:
//
//	NsxtAlbControllerArray{ NsxtAlbControllerArgs{...} }
type NsxtAlbControllerArrayInput interface {
	pulumi.Input

	ToNsxtAlbControllerArrayOutput() NsxtAlbControllerArrayOutput
	ToNsxtAlbControllerArrayOutputWithContext(context.Context) NsxtAlbControllerArrayOutput
}

type NsxtAlbControllerArray []NsxtAlbControllerInput

func (NsxtAlbControllerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NsxtAlbController)(nil)).Elem()
}

func (i NsxtAlbControllerArray) ToNsxtAlbControllerArrayOutput() NsxtAlbControllerArrayOutput {
	return i.ToNsxtAlbControllerArrayOutputWithContext(context.Background())
}

func (i NsxtAlbControllerArray) ToNsxtAlbControllerArrayOutputWithContext(ctx context.Context) NsxtAlbControllerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxtAlbControllerArrayOutput)
}

// NsxtAlbControllerMapInput is an input type that accepts NsxtAlbControllerMap and NsxtAlbControllerMapOutput values.
// You can construct a concrete instance of `NsxtAlbControllerMapInput` via:
//
//	NsxtAlbControllerMap{ "key": NsxtAlbControllerArgs{...} }
type NsxtAlbControllerMapInput interface {
	pulumi.Input

	ToNsxtAlbControllerMapOutput() NsxtAlbControllerMapOutput
	ToNsxtAlbControllerMapOutputWithContext(context.Context) NsxtAlbControllerMapOutput
}

type NsxtAlbControllerMap map[string]NsxtAlbControllerInput

func (NsxtAlbControllerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NsxtAlbController)(nil)).Elem()
}

func (i NsxtAlbControllerMap) ToNsxtAlbControllerMapOutput() NsxtAlbControllerMapOutput {
	return i.ToNsxtAlbControllerMapOutputWithContext(context.Background())
}

func (i NsxtAlbControllerMap) ToNsxtAlbControllerMapOutputWithContext(ctx context.Context) NsxtAlbControllerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NsxtAlbControllerMapOutput)
}

type NsxtAlbControllerOutput struct{ *pulumi.OutputState }

func (NsxtAlbControllerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NsxtAlbController)(nil)).Elem()
}

func (o NsxtAlbControllerOutput) ToNsxtAlbControllerOutput() NsxtAlbControllerOutput {
	return o
}

func (o NsxtAlbControllerOutput) ToNsxtAlbControllerOutputWithContext(ctx context.Context) NsxtAlbControllerOutput {
	return o
}

// An optional description ALB Controller
func (o NsxtAlbControllerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsxtAlbController) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// License type of ALB Controller (`ENTERPRISE` or `BASIC`)
//
// > The attribute `licenseType` must not be used in VCD 10.4+, it is replaced by [nsxtAlbServiceEngineGroup](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_service_engine_group) and [nsxtAlbSettings](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_settings) attribute `supportedFeatureSet`.
func (o NsxtAlbControllerOutput) LicenseType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NsxtAlbController) pulumi.StringPtrOutput { return v.LicenseType }).(pulumi.StringPtrOutput)
}

// A name for ALB Controller
func (o NsxtAlbControllerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxtAlbController) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The password for ALB Controller. Password will not be refreshed.
func (o NsxtAlbControllerOutput) Password() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxtAlbController) pulumi.StringOutput { return v.Password }).(pulumi.StringOutput)
}

// The URL of ALB Controller
func (o NsxtAlbControllerOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxtAlbController) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

// The username for ALB Controller
func (o NsxtAlbControllerOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxtAlbController) pulumi.StringOutput { return v.Username }).(pulumi.StringOutput)
}

// ALB Controller version (e.g. 20.1.3)
func (o NsxtAlbControllerOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *NsxtAlbController) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

type NsxtAlbControllerArrayOutput struct{ *pulumi.OutputState }

func (NsxtAlbControllerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NsxtAlbController)(nil)).Elem()
}

func (o NsxtAlbControllerArrayOutput) ToNsxtAlbControllerArrayOutput() NsxtAlbControllerArrayOutput {
	return o
}

func (o NsxtAlbControllerArrayOutput) ToNsxtAlbControllerArrayOutputWithContext(ctx context.Context) NsxtAlbControllerArrayOutput {
	return o
}

func (o NsxtAlbControllerArrayOutput) Index(i pulumi.IntInput) NsxtAlbControllerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NsxtAlbController {
		return vs[0].([]*NsxtAlbController)[vs[1].(int)]
	}).(NsxtAlbControllerOutput)
}

type NsxtAlbControllerMapOutput struct{ *pulumi.OutputState }

func (NsxtAlbControllerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NsxtAlbController)(nil)).Elem()
}

func (o NsxtAlbControllerMapOutput) ToNsxtAlbControllerMapOutput() NsxtAlbControllerMapOutput {
	return o
}

func (o NsxtAlbControllerMapOutput) ToNsxtAlbControllerMapOutputWithContext(ctx context.Context) NsxtAlbControllerMapOutput {
	return o
}

func (o NsxtAlbControllerMapOutput) MapIndex(k pulumi.StringInput) NsxtAlbControllerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NsxtAlbController {
		return vs[0].(map[string]*NsxtAlbController)[vs[1].(string)]
	}).(NsxtAlbControllerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NsxtAlbControllerInput)(nil)).Elem(), &NsxtAlbController{})
	pulumi.RegisterInputType(reflect.TypeOf((*NsxtAlbControllerArrayInput)(nil)).Elem(), NsxtAlbControllerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NsxtAlbControllerMapInput)(nil)).Elem(), NsxtAlbControllerMap{})
	pulumi.RegisterOutputType(NsxtAlbControllerOutput{})
	pulumi.RegisterOutputType(NsxtAlbControllerArrayOutput{})
	pulumi.RegisterOutputType(NsxtAlbControllerMapOutput{})
}
