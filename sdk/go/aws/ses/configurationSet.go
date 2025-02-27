// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ses

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an SES configuration set resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ses"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ses.NewConfigurationSet(ctx, "test", nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Require TLS Connections
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ses"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ses.NewConfigurationSet(ctx, "test", &ses.ConfigurationSetArgs{
// 			DeliveryOptions: &ses.ConfigurationSetDeliveryOptionsArgs{
// 				TlsPolicy: pulumi.String("Require"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// SES Configuration Sets can be imported using their `name`, e.g.,
//
// ```sh
//  $ pulumi import aws:ses/configurationSet:ConfigurationSet test some-configuration-set-test
// ```
type ConfigurationSet struct {
	pulumi.CustomResourceState

	// SES configuration set ARN.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Configuration block. Detailed below.
	DeliveryOptions ConfigurationSetDeliveryOptionsPtrOutput `pulumi:"deliveryOptions"`
	// The date and time at which the reputation metrics for the configuration set were last reset. Resetting these metrics is known as a fresh start.
	LastFreshStart pulumi.StringOutput `pulumi:"lastFreshStart"`
	// Name of the configuration set.
	Name pulumi.StringOutput `pulumi:"name"`
	// Whether or not Amazon SES publishes reputation metrics for the configuration set, such as bounce and complaint rates, to Amazon CloudWatch. The default value is `false`.
	ReputationMetricsEnabled pulumi.BoolPtrOutput `pulumi:"reputationMetricsEnabled"`
	// Whether email sending is enabled or disabled for the configuration set. The default value is `true`.
	SendingEnabled pulumi.BoolPtrOutput `pulumi:"sendingEnabled"`
}

// NewConfigurationSet registers a new resource with the given unique name, arguments, and options.
func NewConfigurationSet(ctx *pulumi.Context,
	name string, args *ConfigurationSetArgs, opts ...pulumi.ResourceOption) (*ConfigurationSet, error) {
	if args == nil {
		args = &ConfigurationSetArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("aws:ses/confgurationSet:ConfgurationSet"),
		},
	})
	opts = append(opts, aliases)
	var resource ConfigurationSet
	err := ctx.RegisterResource("aws:ses/configurationSet:ConfigurationSet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfigurationSet gets an existing ConfigurationSet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfigurationSet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigurationSetState, opts ...pulumi.ResourceOption) (*ConfigurationSet, error) {
	var resource ConfigurationSet
	err := ctx.ReadResource("aws:ses/configurationSet:ConfigurationSet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConfigurationSet resources.
type configurationSetState struct {
	// SES configuration set ARN.
	Arn *string `pulumi:"arn"`
	// Configuration block. Detailed below.
	DeliveryOptions *ConfigurationSetDeliveryOptions `pulumi:"deliveryOptions"`
	// The date and time at which the reputation metrics for the configuration set were last reset. Resetting these metrics is known as a fresh start.
	LastFreshStart *string `pulumi:"lastFreshStart"`
	// Name of the configuration set.
	Name *string `pulumi:"name"`
	// Whether or not Amazon SES publishes reputation metrics for the configuration set, such as bounce and complaint rates, to Amazon CloudWatch. The default value is `false`.
	ReputationMetricsEnabled *bool `pulumi:"reputationMetricsEnabled"`
	// Whether email sending is enabled or disabled for the configuration set. The default value is `true`.
	SendingEnabled *bool `pulumi:"sendingEnabled"`
}

type ConfigurationSetState struct {
	// SES configuration set ARN.
	Arn pulumi.StringPtrInput
	// Configuration block. Detailed below.
	DeliveryOptions ConfigurationSetDeliveryOptionsPtrInput
	// The date and time at which the reputation metrics for the configuration set were last reset. Resetting these metrics is known as a fresh start.
	LastFreshStart pulumi.StringPtrInput
	// Name of the configuration set.
	Name pulumi.StringPtrInput
	// Whether or not Amazon SES publishes reputation metrics for the configuration set, such as bounce and complaint rates, to Amazon CloudWatch. The default value is `false`.
	ReputationMetricsEnabled pulumi.BoolPtrInput
	// Whether email sending is enabled or disabled for the configuration set. The default value is `true`.
	SendingEnabled pulumi.BoolPtrInput
}

func (ConfigurationSetState) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationSetState)(nil)).Elem()
}

type configurationSetArgs struct {
	// Configuration block. Detailed below.
	DeliveryOptions *ConfigurationSetDeliveryOptions `pulumi:"deliveryOptions"`
	// Name of the configuration set.
	Name *string `pulumi:"name"`
	// Whether or not Amazon SES publishes reputation metrics for the configuration set, such as bounce and complaint rates, to Amazon CloudWatch. The default value is `false`.
	ReputationMetricsEnabled *bool `pulumi:"reputationMetricsEnabled"`
	// Whether email sending is enabled or disabled for the configuration set. The default value is `true`.
	SendingEnabled *bool `pulumi:"sendingEnabled"`
}

// The set of arguments for constructing a ConfigurationSet resource.
type ConfigurationSetArgs struct {
	// Configuration block. Detailed below.
	DeliveryOptions ConfigurationSetDeliveryOptionsPtrInput
	// Name of the configuration set.
	Name pulumi.StringPtrInput
	// Whether or not Amazon SES publishes reputation metrics for the configuration set, such as bounce and complaint rates, to Amazon CloudWatch. The default value is `false`.
	ReputationMetricsEnabled pulumi.BoolPtrInput
	// Whether email sending is enabled or disabled for the configuration set. The default value is `true`.
	SendingEnabled pulumi.BoolPtrInput
}

func (ConfigurationSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configurationSetArgs)(nil)).Elem()
}

type ConfigurationSetInput interface {
	pulumi.Input

	ToConfigurationSetOutput() ConfigurationSetOutput
	ToConfigurationSetOutputWithContext(ctx context.Context) ConfigurationSetOutput
}

func (*ConfigurationSet) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationSet)(nil)).Elem()
}

func (i *ConfigurationSet) ToConfigurationSetOutput() ConfigurationSetOutput {
	return i.ToConfigurationSetOutputWithContext(context.Background())
}

func (i *ConfigurationSet) ToConfigurationSetOutputWithContext(ctx context.Context) ConfigurationSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationSetOutput)
}

// ConfigurationSetArrayInput is an input type that accepts ConfigurationSetArray and ConfigurationSetArrayOutput values.
// You can construct a concrete instance of `ConfigurationSetArrayInput` via:
//
//          ConfigurationSetArray{ ConfigurationSetArgs{...} }
type ConfigurationSetArrayInput interface {
	pulumi.Input

	ToConfigurationSetArrayOutput() ConfigurationSetArrayOutput
	ToConfigurationSetArrayOutputWithContext(context.Context) ConfigurationSetArrayOutput
}

type ConfigurationSetArray []ConfigurationSetInput

func (ConfigurationSetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConfigurationSet)(nil)).Elem()
}

func (i ConfigurationSetArray) ToConfigurationSetArrayOutput() ConfigurationSetArrayOutput {
	return i.ToConfigurationSetArrayOutputWithContext(context.Background())
}

func (i ConfigurationSetArray) ToConfigurationSetArrayOutputWithContext(ctx context.Context) ConfigurationSetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationSetArrayOutput)
}

// ConfigurationSetMapInput is an input type that accepts ConfigurationSetMap and ConfigurationSetMapOutput values.
// You can construct a concrete instance of `ConfigurationSetMapInput` via:
//
//          ConfigurationSetMap{ "key": ConfigurationSetArgs{...} }
type ConfigurationSetMapInput interface {
	pulumi.Input

	ToConfigurationSetMapOutput() ConfigurationSetMapOutput
	ToConfigurationSetMapOutputWithContext(context.Context) ConfigurationSetMapOutput
}

type ConfigurationSetMap map[string]ConfigurationSetInput

func (ConfigurationSetMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConfigurationSet)(nil)).Elem()
}

func (i ConfigurationSetMap) ToConfigurationSetMapOutput() ConfigurationSetMapOutput {
	return i.ToConfigurationSetMapOutputWithContext(context.Background())
}

func (i ConfigurationSetMap) ToConfigurationSetMapOutputWithContext(ctx context.Context) ConfigurationSetMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationSetMapOutput)
}

type ConfigurationSetOutput struct{ *pulumi.OutputState }

func (ConfigurationSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationSet)(nil)).Elem()
}

func (o ConfigurationSetOutput) ToConfigurationSetOutput() ConfigurationSetOutput {
	return o
}

func (o ConfigurationSetOutput) ToConfigurationSetOutputWithContext(ctx context.Context) ConfigurationSetOutput {
	return o
}

type ConfigurationSetArrayOutput struct{ *pulumi.OutputState }

func (ConfigurationSetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConfigurationSet)(nil)).Elem()
}

func (o ConfigurationSetArrayOutput) ToConfigurationSetArrayOutput() ConfigurationSetArrayOutput {
	return o
}

func (o ConfigurationSetArrayOutput) ToConfigurationSetArrayOutputWithContext(ctx context.Context) ConfigurationSetArrayOutput {
	return o
}

func (o ConfigurationSetArrayOutput) Index(i pulumi.IntInput) ConfigurationSetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ConfigurationSet {
		return vs[0].([]*ConfigurationSet)[vs[1].(int)]
	}).(ConfigurationSetOutput)
}

type ConfigurationSetMapOutput struct{ *pulumi.OutputState }

func (ConfigurationSetMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConfigurationSet)(nil)).Elem()
}

func (o ConfigurationSetMapOutput) ToConfigurationSetMapOutput() ConfigurationSetMapOutput {
	return o
}

func (o ConfigurationSetMapOutput) ToConfigurationSetMapOutputWithContext(ctx context.Context) ConfigurationSetMapOutput {
	return o
}

func (o ConfigurationSetMapOutput) MapIndex(k pulumi.StringInput) ConfigurationSetOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ConfigurationSet {
		return vs[0].(map[string]*ConfigurationSet)[vs[1].(string)]
	}).(ConfigurationSetOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationSetInput)(nil)).Elem(), &ConfigurationSet{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationSetArrayInput)(nil)).Elem(), ConfigurationSetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationSetMapInput)(nil)).Elem(), ConfigurationSetMap{})
	pulumi.RegisterOutputType(ConfigurationSetOutput{})
	pulumi.RegisterOutputType(ConfigurationSetArrayOutput{})
	pulumi.RegisterOutputType(ConfigurationSetMapOutput{})
}
