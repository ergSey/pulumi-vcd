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

type LbServiceMonitor struct {
	pulumi.CustomResourceState

	// The name of the edge gateway on which the service monitor is to be created
	EdgeGateway pulumi.StringOutput `pulumi:"edgeGateway"`
	// For types `http` and `https`. String that the monitor expects to match in the status line of
	// the HTTP or HTTPS response (for example, `HTTP/1.1`)
	Expected pulumi.StringPtrOutput `pulumi:"expected"`
	// A map of advanced monitor parameters as key=value pairs (i.e. `max-age=SECONDS`, `invert-regex`)
	// **Note**: When you need a value of `key` only format just set value to empty string (i.e. `linespan = ""`)
	Extension pulumi.StringMapOutput `pulumi:"extension"`
	// Interval in seconds at which a server is to be monitored using the specified Method.
	// Defaults to 10
	Interval pulumi.IntPtrOutput `pulumi:"interval"`
	// Number of times the specified monitoring Method must fail sequentially before the server is
	// declared down. Defaults to 3
	MaxRetries pulumi.IntPtrOutput `pulumi:"maxRetries"`
	// For types `http` and `https`. Select http method to be used to detect server status. One of
	// OPTIONS, GET, HEAD, POST, PUT, DELETE, TRACE, or CONNECT
	Method pulumi.StringPtrOutput `pulumi:"method"`
	// Service Monitor name
	Name pulumi.StringOutput `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
	Org pulumi.StringPtrOutput `pulumi:"org"`
	// For types `http`,  `https`, and `udp`. The string to be matched in the response content.
	// **Note**: When `expected` is not matched, the monitor does not try to match the Receive content
	Receive pulumi.StringPtrOutput `pulumi:"receive"`
	// For types `http`,  `https`, and `udp`. The data to be sent.
	Send pulumi.StringPtrOutput `pulumi:"send"`
	// Maximum time in seconds within which a response from the server must be received. Defaults to 15
	Timeout pulumi.IntPtrOutput `pulumi:"timeout"`
	// Select the way in which you want to send the health check request to the server — `http`, `https`,
	// `tcp`, `icmp`, or `udp`. Depending on the type selected, the remaining attributes are allowed or not
	Type pulumi.StringOutput `pulumi:"type"`
	// For types `http` and `https`. URL to be used in the server status request
	Url pulumi.StringPtrOutput `pulumi:"url"`
	// The name of VDC to use, optional if defined at provider level
	Vdc pulumi.StringPtrOutput `pulumi:"vdc"`
}

// NewLbServiceMonitor registers a new resource with the given unique name, arguments, and options.
func NewLbServiceMonitor(ctx *pulumi.Context,
	name string, args *LbServiceMonitorArgs, opts ...pulumi.ResourceOption) (*LbServiceMonitor, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.EdgeGateway == nil {
		return nil, errors.New("invalid value for required argument 'EdgeGateway'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource LbServiceMonitor
	err := ctx.RegisterResource("vcd:index/lbServiceMonitor:LbServiceMonitor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetLbServiceMonitor gets an existing LbServiceMonitor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetLbServiceMonitor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LbServiceMonitorState, opts ...pulumi.ResourceOption) (*LbServiceMonitor, error) {
	var resource LbServiceMonitor
	err := ctx.ReadResource("vcd:index/lbServiceMonitor:LbServiceMonitor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering LbServiceMonitor resources.
type lbServiceMonitorState struct {
	// The name of the edge gateway on which the service monitor is to be created
	EdgeGateway *string `pulumi:"edgeGateway"`
	// For types `http` and `https`. String that the monitor expects to match in the status line of
	// the HTTP or HTTPS response (for example, `HTTP/1.1`)
	Expected *string `pulumi:"expected"`
	// A map of advanced monitor parameters as key=value pairs (i.e. `max-age=SECONDS`, `invert-regex`)
	// **Note**: When you need a value of `key` only format just set value to empty string (i.e. `linespan = ""`)
	Extension map[string]string `pulumi:"extension"`
	// Interval in seconds at which a server is to be monitored using the specified Method.
	// Defaults to 10
	Interval *int `pulumi:"interval"`
	// Number of times the specified monitoring Method must fail sequentially before the server is
	// declared down. Defaults to 3
	MaxRetries *int `pulumi:"maxRetries"`
	// For types `http` and `https`. Select http method to be used to detect server status. One of
	// OPTIONS, GET, HEAD, POST, PUT, DELETE, TRACE, or CONNECT
	Method *string `pulumi:"method"`
	// Service Monitor name
	Name *string `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
	Org *string `pulumi:"org"`
	// For types `http`,  `https`, and `udp`. The string to be matched in the response content.
	// **Note**: When `expected` is not matched, the monitor does not try to match the Receive content
	Receive *string `pulumi:"receive"`
	// For types `http`,  `https`, and `udp`. The data to be sent.
	Send *string `pulumi:"send"`
	// Maximum time in seconds within which a response from the server must be received. Defaults to 15
	Timeout *int `pulumi:"timeout"`
	// Select the way in which you want to send the health check request to the server — `http`, `https`,
	// `tcp`, `icmp`, or `udp`. Depending on the type selected, the remaining attributes are allowed or not
	Type *string `pulumi:"type"`
	// For types `http` and `https`. URL to be used in the server status request
	Url *string `pulumi:"url"`
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `pulumi:"vdc"`
}

type LbServiceMonitorState struct {
	// The name of the edge gateway on which the service monitor is to be created
	EdgeGateway pulumi.StringPtrInput
	// For types `http` and `https`. String that the monitor expects to match in the status line of
	// the HTTP or HTTPS response (for example, `HTTP/1.1`)
	Expected pulumi.StringPtrInput
	// A map of advanced monitor parameters as key=value pairs (i.e. `max-age=SECONDS`, `invert-regex`)
	// **Note**: When you need a value of `key` only format just set value to empty string (i.e. `linespan = ""`)
	Extension pulumi.StringMapInput
	// Interval in seconds at which a server is to be monitored using the specified Method.
	// Defaults to 10
	Interval pulumi.IntPtrInput
	// Number of times the specified monitoring Method must fail sequentially before the server is
	// declared down. Defaults to 3
	MaxRetries pulumi.IntPtrInput
	// For types `http` and `https`. Select http method to be used to detect server status. One of
	// OPTIONS, GET, HEAD, POST, PUT, DELETE, TRACE, or CONNECT
	Method pulumi.StringPtrInput
	// Service Monitor name
	Name pulumi.StringPtrInput
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
	Org pulumi.StringPtrInput
	// For types `http`,  `https`, and `udp`. The string to be matched in the response content.
	// **Note**: When `expected` is not matched, the monitor does not try to match the Receive content
	Receive pulumi.StringPtrInput
	// For types `http`,  `https`, and `udp`. The data to be sent.
	Send pulumi.StringPtrInput
	// Maximum time in seconds within which a response from the server must be received. Defaults to 15
	Timeout pulumi.IntPtrInput
	// Select the way in which you want to send the health check request to the server — `http`, `https`,
	// `tcp`, `icmp`, or `udp`. Depending on the type selected, the remaining attributes are allowed or not
	Type pulumi.StringPtrInput
	// For types `http` and `https`. URL to be used in the server status request
	Url pulumi.StringPtrInput
	// The name of VDC to use, optional if defined at provider level
	Vdc pulumi.StringPtrInput
}

func (LbServiceMonitorState) ElementType() reflect.Type {
	return reflect.TypeOf((*lbServiceMonitorState)(nil)).Elem()
}

type lbServiceMonitorArgs struct {
	// The name of the edge gateway on which the service monitor is to be created
	EdgeGateway string `pulumi:"edgeGateway"`
	// For types `http` and `https`. String that the monitor expects to match in the status line of
	// the HTTP or HTTPS response (for example, `HTTP/1.1`)
	Expected *string `pulumi:"expected"`
	// A map of advanced monitor parameters as key=value pairs (i.e. `max-age=SECONDS`, `invert-regex`)
	// **Note**: When you need a value of `key` only format just set value to empty string (i.e. `linespan = ""`)
	Extension map[string]string `pulumi:"extension"`
	// Interval in seconds at which a server is to be monitored using the specified Method.
	// Defaults to 10
	Interval *int `pulumi:"interval"`
	// Number of times the specified monitoring Method must fail sequentially before the server is
	// declared down. Defaults to 3
	MaxRetries *int `pulumi:"maxRetries"`
	// For types `http` and `https`. Select http method to be used to detect server status. One of
	// OPTIONS, GET, HEAD, POST, PUT, DELETE, TRACE, or CONNECT
	Method *string `pulumi:"method"`
	// Service Monitor name
	Name *string `pulumi:"name"`
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
	Org *string `pulumi:"org"`
	// For types `http`,  `https`, and `udp`. The string to be matched in the response content.
	// **Note**: When `expected` is not matched, the monitor does not try to match the Receive content
	Receive *string `pulumi:"receive"`
	// For types `http`,  `https`, and `udp`. The data to be sent.
	Send *string `pulumi:"send"`
	// Maximum time in seconds within which a response from the server must be received. Defaults to 15
	Timeout *int `pulumi:"timeout"`
	// Select the way in which you want to send the health check request to the server — `http`, `https`,
	// `tcp`, `icmp`, or `udp`. Depending on the type selected, the remaining attributes are allowed or not
	Type string `pulumi:"type"`
	// For types `http` and `https`. URL to be used in the server status request
	Url *string `pulumi:"url"`
	// The name of VDC to use, optional if defined at provider level
	Vdc *string `pulumi:"vdc"`
}

// The set of arguments for constructing a LbServiceMonitor resource.
type LbServiceMonitorArgs struct {
	// The name of the edge gateway on which the service monitor is to be created
	EdgeGateway pulumi.StringInput
	// For types `http` and `https`. String that the monitor expects to match in the status line of
	// the HTTP or HTTPS response (for example, `HTTP/1.1`)
	Expected pulumi.StringPtrInput
	// A map of advanced monitor parameters as key=value pairs (i.e. `max-age=SECONDS`, `invert-regex`)
	// **Note**: When you need a value of `key` only format just set value to empty string (i.e. `linespan = ""`)
	Extension pulumi.StringMapInput
	// Interval in seconds at which a server is to be monitored using the specified Method.
	// Defaults to 10
	Interval pulumi.IntPtrInput
	// Number of times the specified monitoring Method must fail sequentially before the server is
	// declared down. Defaults to 3
	MaxRetries pulumi.IntPtrInput
	// For types `http` and `https`. Select http method to be used to detect server status. One of
	// OPTIONS, GET, HEAD, POST, PUT, DELETE, TRACE, or CONNECT
	Method pulumi.StringPtrInput
	// Service Monitor name
	Name pulumi.StringPtrInput
	// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
	Org pulumi.StringPtrInput
	// For types `http`,  `https`, and `udp`. The string to be matched in the response content.
	// **Note**: When `expected` is not matched, the monitor does not try to match the Receive content
	Receive pulumi.StringPtrInput
	// For types `http`,  `https`, and `udp`. The data to be sent.
	Send pulumi.StringPtrInput
	// Maximum time in seconds within which a response from the server must be received. Defaults to 15
	Timeout pulumi.IntPtrInput
	// Select the way in which you want to send the health check request to the server — `http`, `https`,
	// `tcp`, `icmp`, or `udp`. Depending on the type selected, the remaining attributes are allowed or not
	Type pulumi.StringInput
	// For types `http` and `https`. URL to be used in the server status request
	Url pulumi.StringPtrInput
	// The name of VDC to use, optional if defined at provider level
	Vdc pulumi.StringPtrInput
}

func (LbServiceMonitorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*lbServiceMonitorArgs)(nil)).Elem()
}

type LbServiceMonitorInput interface {
	pulumi.Input

	ToLbServiceMonitorOutput() LbServiceMonitorOutput
	ToLbServiceMonitorOutputWithContext(ctx context.Context) LbServiceMonitorOutput
}

func (*LbServiceMonitor) ElementType() reflect.Type {
	return reflect.TypeOf((**LbServiceMonitor)(nil)).Elem()
}

func (i *LbServiceMonitor) ToLbServiceMonitorOutput() LbServiceMonitorOutput {
	return i.ToLbServiceMonitorOutputWithContext(context.Background())
}

func (i *LbServiceMonitor) ToLbServiceMonitorOutputWithContext(ctx context.Context) LbServiceMonitorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LbServiceMonitorOutput)
}

// LbServiceMonitorArrayInput is an input type that accepts LbServiceMonitorArray and LbServiceMonitorArrayOutput values.
// You can construct a concrete instance of `LbServiceMonitorArrayInput` via:
//
//	LbServiceMonitorArray{ LbServiceMonitorArgs{...} }
type LbServiceMonitorArrayInput interface {
	pulumi.Input

	ToLbServiceMonitorArrayOutput() LbServiceMonitorArrayOutput
	ToLbServiceMonitorArrayOutputWithContext(context.Context) LbServiceMonitorArrayOutput
}

type LbServiceMonitorArray []LbServiceMonitorInput

func (LbServiceMonitorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LbServiceMonitor)(nil)).Elem()
}

func (i LbServiceMonitorArray) ToLbServiceMonitorArrayOutput() LbServiceMonitorArrayOutput {
	return i.ToLbServiceMonitorArrayOutputWithContext(context.Background())
}

func (i LbServiceMonitorArray) ToLbServiceMonitorArrayOutputWithContext(ctx context.Context) LbServiceMonitorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LbServiceMonitorArrayOutput)
}

// LbServiceMonitorMapInput is an input type that accepts LbServiceMonitorMap and LbServiceMonitorMapOutput values.
// You can construct a concrete instance of `LbServiceMonitorMapInput` via:
//
//	LbServiceMonitorMap{ "key": LbServiceMonitorArgs{...} }
type LbServiceMonitorMapInput interface {
	pulumi.Input

	ToLbServiceMonitorMapOutput() LbServiceMonitorMapOutput
	ToLbServiceMonitorMapOutputWithContext(context.Context) LbServiceMonitorMapOutput
}

type LbServiceMonitorMap map[string]LbServiceMonitorInput

func (LbServiceMonitorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LbServiceMonitor)(nil)).Elem()
}

func (i LbServiceMonitorMap) ToLbServiceMonitorMapOutput() LbServiceMonitorMapOutput {
	return i.ToLbServiceMonitorMapOutputWithContext(context.Background())
}

func (i LbServiceMonitorMap) ToLbServiceMonitorMapOutputWithContext(ctx context.Context) LbServiceMonitorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LbServiceMonitorMapOutput)
}

type LbServiceMonitorOutput struct{ *pulumi.OutputState }

func (LbServiceMonitorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LbServiceMonitor)(nil)).Elem()
}

func (o LbServiceMonitorOutput) ToLbServiceMonitorOutput() LbServiceMonitorOutput {
	return o
}

func (o LbServiceMonitorOutput) ToLbServiceMonitorOutputWithContext(ctx context.Context) LbServiceMonitorOutput {
	return o
}

// The name of the edge gateway on which the service monitor is to be created
func (o LbServiceMonitorOutput) EdgeGateway() pulumi.StringOutput {
	return o.ApplyT(func(v *LbServiceMonitor) pulumi.StringOutput { return v.EdgeGateway }).(pulumi.StringOutput)
}

// For types `http` and `https`. String that the monitor expects to match in the status line of
// the HTTP or HTTPS response (for example, `HTTP/1.1`)
func (o LbServiceMonitorOutput) Expected() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LbServiceMonitor) pulumi.StringPtrOutput { return v.Expected }).(pulumi.StringPtrOutput)
}

// A map of advanced monitor parameters as key=value pairs (i.e. `max-age=SECONDS`, `invert-regex`)
// **Note**: When you need a value of `key` only format just set value to empty string (i.e. `linespan = ""`)
func (o LbServiceMonitorOutput) Extension() pulumi.StringMapOutput {
	return o.ApplyT(func(v *LbServiceMonitor) pulumi.StringMapOutput { return v.Extension }).(pulumi.StringMapOutput)
}

// Interval in seconds at which a server is to be monitored using the specified Method.
// Defaults to 10
func (o LbServiceMonitorOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LbServiceMonitor) pulumi.IntPtrOutput { return v.Interval }).(pulumi.IntPtrOutput)
}

// Number of times the specified monitoring Method must fail sequentially before the server is
// declared down. Defaults to 3
func (o LbServiceMonitorOutput) MaxRetries() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LbServiceMonitor) pulumi.IntPtrOutput { return v.MaxRetries }).(pulumi.IntPtrOutput)
}

// For types `http` and `https`. Select http method to be used to detect server status. One of
// OPTIONS, GET, HEAD, POST, PUT, DELETE, TRACE, or CONNECT
func (o LbServiceMonitorOutput) Method() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LbServiceMonitor) pulumi.StringPtrOutput { return v.Method }).(pulumi.StringPtrOutput)
}

// Service Monitor name
func (o LbServiceMonitorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LbServiceMonitor) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
func (o LbServiceMonitorOutput) Org() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LbServiceMonitor) pulumi.StringPtrOutput { return v.Org }).(pulumi.StringPtrOutput)
}

// For types `http`,  `https`, and `udp`. The string to be matched in the response content.
// **Note**: When `expected` is not matched, the monitor does not try to match the Receive content
func (o LbServiceMonitorOutput) Receive() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LbServiceMonitor) pulumi.StringPtrOutput { return v.Receive }).(pulumi.StringPtrOutput)
}

// For types `http`,  `https`, and `udp`. The data to be sent.
func (o LbServiceMonitorOutput) Send() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LbServiceMonitor) pulumi.StringPtrOutput { return v.Send }).(pulumi.StringPtrOutput)
}

// Maximum time in seconds within which a response from the server must be received. Defaults to 15
func (o LbServiceMonitorOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *LbServiceMonitor) pulumi.IntPtrOutput { return v.Timeout }).(pulumi.IntPtrOutput)
}

// Select the way in which you want to send the health check request to the server — `http`, `https`,
// `tcp`, `icmp`, or `udp`. Depending on the type selected, the remaining attributes are allowed or not
func (o LbServiceMonitorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LbServiceMonitor) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// For types `http` and `https`. URL to be used in the server status request
func (o LbServiceMonitorOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LbServiceMonitor) pulumi.StringPtrOutput { return v.Url }).(pulumi.StringPtrOutput)
}

// The name of VDC to use, optional if defined at provider level
func (o LbServiceMonitorOutput) Vdc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LbServiceMonitor) pulumi.StringPtrOutput { return v.Vdc }).(pulumi.StringPtrOutput)
}

type LbServiceMonitorArrayOutput struct{ *pulumi.OutputState }

func (LbServiceMonitorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*LbServiceMonitor)(nil)).Elem()
}

func (o LbServiceMonitorArrayOutput) ToLbServiceMonitorArrayOutput() LbServiceMonitorArrayOutput {
	return o
}

func (o LbServiceMonitorArrayOutput) ToLbServiceMonitorArrayOutputWithContext(ctx context.Context) LbServiceMonitorArrayOutput {
	return o
}

func (o LbServiceMonitorArrayOutput) Index(i pulumi.IntInput) LbServiceMonitorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *LbServiceMonitor {
		return vs[0].([]*LbServiceMonitor)[vs[1].(int)]
	}).(LbServiceMonitorOutput)
}

type LbServiceMonitorMapOutput struct{ *pulumi.OutputState }

func (LbServiceMonitorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*LbServiceMonitor)(nil)).Elem()
}

func (o LbServiceMonitorMapOutput) ToLbServiceMonitorMapOutput() LbServiceMonitorMapOutput {
	return o
}

func (o LbServiceMonitorMapOutput) ToLbServiceMonitorMapOutputWithContext(ctx context.Context) LbServiceMonitorMapOutput {
	return o
}

func (o LbServiceMonitorMapOutput) MapIndex(k pulumi.StringInput) LbServiceMonitorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *LbServiceMonitor {
		return vs[0].(map[string]*LbServiceMonitor)[vs[1].(string)]
	}).(LbServiceMonitorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LbServiceMonitorInput)(nil)).Elem(), &LbServiceMonitor{})
	pulumi.RegisterInputType(reflect.TypeOf((*LbServiceMonitorArrayInput)(nil)).Elem(), LbServiceMonitorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LbServiceMonitorMapInput)(nil)).Elem(), LbServiceMonitorMap{})
	pulumi.RegisterOutputType(LbServiceMonitorOutput{})
	pulumi.RegisterOutputType(LbServiceMonitorArrayOutput{})
	pulumi.RegisterOutputType(LbServiceMonitorMapOutput{})
}
