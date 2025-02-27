// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package opsworks

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an OpsWorks haproxy layer resource.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/opsworks"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := opsworks.NewHaproxyLayer(ctx, "lb", &opsworks.HaproxyLayerArgs{
// 			StackId:       pulumi.Any(aws_opsworks_stack.Main.Id),
// 			StatsPassword: pulumi.String("foobarbaz"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type HaproxyLayer struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name(ARN) of the layer.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps pulumi.BoolPtrOutput `pulumi:"autoAssignElasticIps"`
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps pulumi.BoolPtrOutput `pulumi:"autoAssignPublicIps"`
	// Whether to enable auto-healing for the layer.
	AutoHealing             pulumi.BoolPtrOutput                         `pulumi:"autoHealing"`
	CloudwatchConfiguration HaproxyLayerCloudwatchConfigurationPtrOutput `pulumi:"cloudwatchConfiguration"`
	CustomConfigureRecipes  pulumi.StringArrayOutput                     `pulumi:"customConfigureRecipes"`
	CustomDeployRecipes     pulumi.StringArrayOutput                     `pulumi:"customDeployRecipes"`
	// The ARN of an IAM profile that will be used for the layer's instances.
	CustomInstanceProfileArn pulumi.StringPtrOutput `pulumi:"customInstanceProfileArn"`
	// Custom JSON attributes to apply to the layer.
	CustomJson pulumi.StringPtrOutput `pulumi:"customJson"`
	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds pulumi.StringArrayOutput `pulumi:"customSecurityGroupIds"`
	CustomSetupRecipes     pulumi.StringArrayOutput `pulumi:"customSetupRecipes"`
	CustomShutdownRecipes  pulumi.StringArrayOutput `pulumi:"customShutdownRecipes"`
	CustomUndeployRecipes  pulumi.StringArrayOutput `pulumi:"customUndeployRecipes"`
	// Whether to enable Elastic Load Balancing connection draining.
	DrainElbOnShutdown pulumi.BoolPtrOutput `pulumi:"drainElbOnShutdown"`
	// `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	EbsVolumes HaproxyLayerEbsVolumeArrayOutput `pulumi:"ebsVolumes"`
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer pulumi.StringPtrOutput `pulumi:"elasticLoadBalancer"`
	// HTTP method to use for instance healthchecks. Defaults to "OPTIONS".
	HealthcheckMethod pulumi.StringPtrOutput `pulumi:"healthcheckMethod"`
	// URL path to use for instance healthchecks. Defaults to "/".
	HealthcheckUrl pulumi.StringPtrOutput `pulumi:"healthcheckUrl"`
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot pulumi.BoolPtrOutput `pulumi:"installUpdatesOnBoot"`
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout pulumi.IntPtrOutput `pulumi:"instanceShutdownTimeout"`
	// A human-readable name for the layer.
	Name pulumi.StringOutput `pulumi:"name"`
	// The id of the stack the layer will belong to.
	StackId pulumi.StringOutput `pulumi:"stackId"`
	// Whether to enable HAProxy stats.
	StatsEnabled pulumi.BoolPtrOutput `pulumi:"statsEnabled"`
	// The password to use for HAProxy stats.
	StatsPassword pulumi.StringOutput `pulumi:"statsPassword"`
	// The HAProxy stats URL. Defaults to "/haproxy?stats".
	StatsUrl pulumi.StringPtrOutput `pulumi:"statsUrl"`
	// The username for HAProxy stats. Defaults to "opsworks".
	StatsUser pulumi.StringPtrOutput `pulumi:"statsUser"`
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages pulumi.StringArrayOutput `pulumi:"systemPackages"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances pulumi.BoolPtrOutput `pulumi:"useEbsOptimizedInstances"`
}

// NewHaproxyLayer registers a new resource with the given unique name, arguments, and options.
func NewHaproxyLayer(ctx *pulumi.Context,
	name string, args *HaproxyLayerArgs, opts ...pulumi.ResourceOption) (*HaproxyLayer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.StackId == nil {
		return nil, errors.New("invalid value for required argument 'StackId'")
	}
	if args.StatsPassword == nil {
		return nil, errors.New("invalid value for required argument 'StatsPassword'")
	}
	var resource HaproxyLayer
	err := ctx.RegisterResource("aws:opsworks/haproxyLayer:HaproxyLayer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetHaproxyLayer gets an existing HaproxyLayer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetHaproxyLayer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HaproxyLayerState, opts ...pulumi.ResourceOption) (*HaproxyLayer, error) {
	var resource HaproxyLayer
	err := ctx.ReadResource("aws:opsworks/haproxyLayer:HaproxyLayer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering HaproxyLayer resources.
type haproxyLayerState struct {
	// The Amazon Resource Name(ARN) of the layer.
	Arn *string `pulumi:"arn"`
	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps *bool `pulumi:"autoAssignElasticIps"`
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps *bool `pulumi:"autoAssignPublicIps"`
	// Whether to enable auto-healing for the layer.
	AutoHealing             *bool                                `pulumi:"autoHealing"`
	CloudwatchConfiguration *HaproxyLayerCloudwatchConfiguration `pulumi:"cloudwatchConfiguration"`
	CustomConfigureRecipes  []string                             `pulumi:"customConfigureRecipes"`
	CustomDeployRecipes     []string                             `pulumi:"customDeployRecipes"`
	// The ARN of an IAM profile that will be used for the layer's instances.
	CustomInstanceProfileArn *string `pulumi:"customInstanceProfileArn"`
	// Custom JSON attributes to apply to the layer.
	CustomJson *string `pulumi:"customJson"`
	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds []string `pulumi:"customSecurityGroupIds"`
	CustomSetupRecipes     []string `pulumi:"customSetupRecipes"`
	CustomShutdownRecipes  []string `pulumi:"customShutdownRecipes"`
	CustomUndeployRecipes  []string `pulumi:"customUndeployRecipes"`
	// Whether to enable Elastic Load Balancing connection draining.
	DrainElbOnShutdown *bool `pulumi:"drainElbOnShutdown"`
	// `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	EbsVolumes []HaproxyLayerEbsVolume `pulumi:"ebsVolumes"`
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer *string `pulumi:"elasticLoadBalancer"`
	// HTTP method to use for instance healthchecks. Defaults to "OPTIONS".
	HealthcheckMethod *string `pulumi:"healthcheckMethod"`
	// URL path to use for instance healthchecks. Defaults to "/".
	HealthcheckUrl *string `pulumi:"healthcheckUrl"`
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot *bool `pulumi:"installUpdatesOnBoot"`
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout *int `pulumi:"instanceShutdownTimeout"`
	// A human-readable name for the layer.
	Name *string `pulumi:"name"`
	// The id of the stack the layer will belong to.
	StackId *string `pulumi:"stackId"`
	// Whether to enable HAProxy stats.
	StatsEnabled *bool `pulumi:"statsEnabled"`
	// The password to use for HAProxy stats.
	StatsPassword *string `pulumi:"statsPassword"`
	// The HAProxy stats URL. Defaults to "/haproxy?stats".
	StatsUrl *string `pulumi:"statsUrl"`
	// The username for HAProxy stats. Defaults to "opsworks".
	StatsUser *string `pulumi:"statsUser"`
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages []string `pulumi:"systemPackages"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances *bool `pulumi:"useEbsOptimizedInstances"`
}

type HaproxyLayerState struct {
	// The Amazon Resource Name(ARN) of the layer.
	Arn pulumi.StringPtrInput
	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps pulumi.BoolPtrInput
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps pulumi.BoolPtrInput
	// Whether to enable auto-healing for the layer.
	AutoHealing             pulumi.BoolPtrInput
	CloudwatchConfiguration HaproxyLayerCloudwatchConfigurationPtrInput
	CustomConfigureRecipes  pulumi.StringArrayInput
	CustomDeployRecipes     pulumi.StringArrayInput
	// The ARN of an IAM profile that will be used for the layer's instances.
	CustomInstanceProfileArn pulumi.StringPtrInput
	// Custom JSON attributes to apply to the layer.
	CustomJson pulumi.StringPtrInput
	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds pulumi.StringArrayInput
	CustomSetupRecipes     pulumi.StringArrayInput
	CustomShutdownRecipes  pulumi.StringArrayInput
	CustomUndeployRecipes  pulumi.StringArrayInput
	// Whether to enable Elastic Load Balancing connection draining.
	DrainElbOnShutdown pulumi.BoolPtrInput
	// `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	EbsVolumes HaproxyLayerEbsVolumeArrayInput
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer pulumi.StringPtrInput
	// HTTP method to use for instance healthchecks. Defaults to "OPTIONS".
	HealthcheckMethod pulumi.StringPtrInput
	// URL path to use for instance healthchecks. Defaults to "/".
	HealthcheckUrl pulumi.StringPtrInput
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot pulumi.BoolPtrInput
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout pulumi.IntPtrInput
	// A human-readable name for the layer.
	Name pulumi.StringPtrInput
	// The id of the stack the layer will belong to.
	StackId pulumi.StringPtrInput
	// Whether to enable HAProxy stats.
	StatsEnabled pulumi.BoolPtrInput
	// The password to use for HAProxy stats.
	StatsPassword pulumi.StringPtrInput
	// The HAProxy stats URL. Defaults to "/haproxy?stats".
	StatsUrl pulumi.StringPtrInput
	// The username for HAProxy stats. Defaults to "opsworks".
	StatsUser pulumi.StringPtrInput
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages pulumi.StringArrayInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances pulumi.BoolPtrInput
}

func (HaproxyLayerState) ElementType() reflect.Type {
	return reflect.TypeOf((*haproxyLayerState)(nil)).Elem()
}

type haproxyLayerArgs struct {
	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps *bool `pulumi:"autoAssignElasticIps"`
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps *bool `pulumi:"autoAssignPublicIps"`
	// Whether to enable auto-healing for the layer.
	AutoHealing             *bool                                `pulumi:"autoHealing"`
	CloudwatchConfiguration *HaproxyLayerCloudwatchConfiguration `pulumi:"cloudwatchConfiguration"`
	CustomConfigureRecipes  []string                             `pulumi:"customConfigureRecipes"`
	CustomDeployRecipes     []string                             `pulumi:"customDeployRecipes"`
	// The ARN of an IAM profile that will be used for the layer's instances.
	CustomInstanceProfileArn *string `pulumi:"customInstanceProfileArn"`
	// Custom JSON attributes to apply to the layer.
	CustomJson *string `pulumi:"customJson"`
	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds []string `pulumi:"customSecurityGroupIds"`
	CustomSetupRecipes     []string `pulumi:"customSetupRecipes"`
	CustomShutdownRecipes  []string `pulumi:"customShutdownRecipes"`
	CustomUndeployRecipes  []string `pulumi:"customUndeployRecipes"`
	// Whether to enable Elastic Load Balancing connection draining.
	DrainElbOnShutdown *bool `pulumi:"drainElbOnShutdown"`
	// `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	EbsVolumes []HaproxyLayerEbsVolume `pulumi:"ebsVolumes"`
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer *string `pulumi:"elasticLoadBalancer"`
	// HTTP method to use for instance healthchecks. Defaults to "OPTIONS".
	HealthcheckMethod *string `pulumi:"healthcheckMethod"`
	// URL path to use for instance healthchecks. Defaults to "/".
	HealthcheckUrl *string `pulumi:"healthcheckUrl"`
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot *bool `pulumi:"installUpdatesOnBoot"`
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout *int `pulumi:"instanceShutdownTimeout"`
	// A human-readable name for the layer.
	Name *string `pulumi:"name"`
	// The id of the stack the layer will belong to.
	StackId string `pulumi:"stackId"`
	// Whether to enable HAProxy stats.
	StatsEnabled *bool `pulumi:"statsEnabled"`
	// The password to use for HAProxy stats.
	StatsPassword string `pulumi:"statsPassword"`
	// The HAProxy stats URL. Defaults to "/haproxy?stats".
	StatsUrl *string `pulumi:"statsUrl"`
	// The username for HAProxy stats. Defaults to "opsworks".
	StatsUser *string `pulumi:"statsUser"`
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages []string `pulumi:"systemPackages"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances *bool `pulumi:"useEbsOptimizedInstances"`
}

// The set of arguments for constructing a HaproxyLayer resource.
type HaproxyLayerArgs struct {
	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps pulumi.BoolPtrInput
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps pulumi.BoolPtrInput
	// Whether to enable auto-healing for the layer.
	AutoHealing             pulumi.BoolPtrInput
	CloudwatchConfiguration HaproxyLayerCloudwatchConfigurationPtrInput
	CustomConfigureRecipes  pulumi.StringArrayInput
	CustomDeployRecipes     pulumi.StringArrayInput
	// The ARN of an IAM profile that will be used for the layer's instances.
	CustomInstanceProfileArn pulumi.StringPtrInput
	// Custom JSON attributes to apply to the layer.
	CustomJson pulumi.StringPtrInput
	// Ids for a set of security groups to apply to the layer's instances.
	CustomSecurityGroupIds pulumi.StringArrayInput
	CustomSetupRecipes     pulumi.StringArrayInput
	CustomShutdownRecipes  pulumi.StringArrayInput
	CustomUndeployRecipes  pulumi.StringArrayInput
	// Whether to enable Elastic Load Balancing connection draining.
	DrainElbOnShutdown pulumi.BoolPtrInput
	// `ebsVolume` blocks, as described below, will each create an EBS volume and connect it to the layer's instances.
	EbsVolumes HaproxyLayerEbsVolumeArrayInput
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer pulumi.StringPtrInput
	// HTTP method to use for instance healthchecks. Defaults to "OPTIONS".
	HealthcheckMethod pulumi.StringPtrInput
	// URL path to use for instance healthchecks. Defaults to "/".
	HealthcheckUrl pulumi.StringPtrInput
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot pulumi.BoolPtrInput
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout pulumi.IntPtrInput
	// A human-readable name for the layer.
	Name pulumi.StringPtrInput
	// The id of the stack the layer will belong to.
	StackId pulumi.StringInput
	// Whether to enable HAProxy stats.
	StatsEnabled pulumi.BoolPtrInput
	// The password to use for HAProxy stats.
	StatsPassword pulumi.StringInput
	// The HAProxy stats URL. Defaults to "/haproxy?stats".
	StatsUrl pulumi.StringPtrInput
	// The username for HAProxy stats. Defaults to "opsworks".
	StatsUser pulumi.StringPtrInput
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages pulumi.StringArrayInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances pulumi.BoolPtrInput
}

func (HaproxyLayerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*haproxyLayerArgs)(nil)).Elem()
}

type HaproxyLayerInput interface {
	pulumi.Input

	ToHaproxyLayerOutput() HaproxyLayerOutput
	ToHaproxyLayerOutputWithContext(ctx context.Context) HaproxyLayerOutput
}

func (*HaproxyLayer) ElementType() reflect.Type {
	return reflect.TypeOf((**HaproxyLayer)(nil)).Elem()
}

func (i *HaproxyLayer) ToHaproxyLayerOutput() HaproxyLayerOutput {
	return i.ToHaproxyLayerOutputWithContext(context.Background())
}

func (i *HaproxyLayer) ToHaproxyLayerOutputWithContext(ctx context.Context) HaproxyLayerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HaproxyLayerOutput)
}

// HaproxyLayerArrayInput is an input type that accepts HaproxyLayerArray and HaproxyLayerArrayOutput values.
// You can construct a concrete instance of `HaproxyLayerArrayInput` via:
//
//          HaproxyLayerArray{ HaproxyLayerArgs{...} }
type HaproxyLayerArrayInput interface {
	pulumi.Input

	ToHaproxyLayerArrayOutput() HaproxyLayerArrayOutput
	ToHaproxyLayerArrayOutputWithContext(context.Context) HaproxyLayerArrayOutput
}

type HaproxyLayerArray []HaproxyLayerInput

func (HaproxyLayerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HaproxyLayer)(nil)).Elem()
}

func (i HaproxyLayerArray) ToHaproxyLayerArrayOutput() HaproxyLayerArrayOutput {
	return i.ToHaproxyLayerArrayOutputWithContext(context.Background())
}

func (i HaproxyLayerArray) ToHaproxyLayerArrayOutputWithContext(ctx context.Context) HaproxyLayerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HaproxyLayerArrayOutput)
}

// HaproxyLayerMapInput is an input type that accepts HaproxyLayerMap and HaproxyLayerMapOutput values.
// You can construct a concrete instance of `HaproxyLayerMapInput` via:
//
//          HaproxyLayerMap{ "key": HaproxyLayerArgs{...} }
type HaproxyLayerMapInput interface {
	pulumi.Input

	ToHaproxyLayerMapOutput() HaproxyLayerMapOutput
	ToHaproxyLayerMapOutputWithContext(context.Context) HaproxyLayerMapOutput
}

type HaproxyLayerMap map[string]HaproxyLayerInput

func (HaproxyLayerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HaproxyLayer)(nil)).Elem()
}

func (i HaproxyLayerMap) ToHaproxyLayerMapOutput() HaproxyLayerMapOutput {
	return i.ToHaproxyLayerMapOutputWithContext(context.Background())
}

func (i HaproxyLayerMap) ToHaproxyLayerMapOutputWithContext(ctx context.Context) HaproxyLayerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HaproxyLayerMapOutput)
}

type HaproxyLayerOutput struct{ *pulumi.OutputState }

func (HaproxyLayerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HaproxyLayer)(nil)).Elem()
}

func (o HaproxyLayerOutput) ToHaproxyLayerOutput() HaproxyLayerOutput {
	return o
}

func (o HaproxyLayerOutput) ToHaproxyLayerOutputWithContext(ctx context.Context) HaproxyLayerOutput {
	return o
}

type HaproxyLayerArrayOutput struct{ *pulumi.OutputState }

func (HaproxyLayerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*HaproxyLayer)(nil)).Elem()
}

func (o HaproxyLayerArrayOutput) ToHaproxyLayerArrayOutput() HaproxyLayerArrayOutput {
	return o
}

func (o HaproxyLayerArrayOutput) ToHaproxyLayerArrayOutputWithContext(ctx context.Context) HaproxyLayerArrayOutput {
	return o
}

func (o HaproxyLayerArrayOutput) Index(i pulumi.IntInput) HaproxyLayerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *HaproxyLayer {
		return vs[0].([]*HaproxyLayer)[vs[1].(int)]
	}).(HaproxyLayerOutput)
}

type HaproxyLayerMapOutput struct{ *pulumi.OutputState }

func (HaproxyLayerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*HaproxyLayer)(nil)).Elem()
}

func (o HaproxyLayerMapOutput) ToHaproxyLayerMapOutput() HaproxyLayerMapOutput {
	return o
}

func (o HaproxyLayerMapOutput) ToHaproxyLayerMapOutputWithContext(ctx context.Context) HaproxyLayerMapOutput {
	return o
}

func (o HaproxyLayerMapOutput) MapIndex(k pulumi.StringInput) HaproxyLayerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *HaproxyLayer {
		return vs[0].(map[string]*HaproxyLayer)[vs[1].(string)]
	}).(HaproxyLayerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HaproxyLayerInput)(nil)).Elem(), &HaproxyLayer{})
	pulumi.RegisterInputType(reflect.TypeOf((*HaproxyLayerArrayInput)(nil)).Elem(), HaproxyLayerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HaproxyLayerMapInput)(nil)).Elem(), HaproxyLayerMap{})
	pulumi.RegisterOutputType(HaproxyLayerOutput{})
	pulumi.RegisterOutputType(HaproxyLayerArrayOutput{})
	pulumi.RegisterOutputType(HaproxyLayerMapOutput{})
}
