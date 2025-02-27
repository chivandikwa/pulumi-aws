// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package opsworks

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an OpsWorks MySQL layer resource.
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
// 		_, err := opsworks.NewMysqlLayer(ctx, "db", &opsworks.MysqlLayerArgs{
// 			StackId: pulumi.Any(aws_opsworks_stack.Main.Id),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type MysqlLayer struct {
	pulumi.CustomResourceState

	// The Amazon Resource Name(ARN) of the layer.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps pulumi.BoolPtrOutput `pulumi:"autoAssignElasticIps"`
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps pulumi.BoolPtrOutput `pulumi:"autoAssignPublicIps"`
	// Whether to enable auto-healing for the layer.
	AutoHealing             pulumi.BoolPtrOutput                       `pulumi:"autoHealing"`
	CloudwatchConfiguration MysqlLayerCloudwatchConfigurationPtrOutput `pulumi:"cloudwatchConfiguration"`
	CustomConfigureRecipes  pulumi.StringArrayOutput                   `pulumi:"customConfigureRecipes"`
	CustomDeployRecipes     pulumi.StringArrayOutput                   `pulumi:"customDeployRecipes"`
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
	EbsVolumes MysqlLayerEbsVolumeArrayOutput `pulumi:"ebsVolumes"`
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer pulumi.StringPtrOutput `pulumi:"elasticLoadBalancer"`
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot pulumi.BoolPtrOutput `pulumi:"installUpdatesOnBoot"`
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout pulumi.IntPtrOutput `pulumi:"instanceShutdownTimeout"`
	// A human-readable name for the layer.
	Name pulumi.StringOutput `pulumi:"name"`
	// Root password to use for MySQL.
	RootPassword pulumi.StringPtrOutput `pulumi:"rootPassword"`
	// Whether to set the root user password to all instances in the stack so they can access the instances in this layer.
	RootPasswordOnAllInstances pulumi.BoolPtrOutput `pulumi:"rootPasswordOnAllInstances"`
	// The id of the stack the layer will belong to.
	StackId pulumi.StringOutput `pulumi:"stackId"`
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages pulumi.StringArrayOutput `pulumi:"systemPackages"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances pulumi.BoolPtrOutput `pulumi:"useEbsOptimizedInstances"`
}

// NewMysqlLayer registers a new resource with the given unique name, arguments, and options.
func NewMysqlLayer(ctx *pulumi.Context,
	name string, args *MysqlLayerArgs, opts ...pulumi.ResourceOption) (*MysqlLayer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.StackId == nil {
		return nil, errors.New("invalid value for required argument 'StackId'")
	}
	var resource MysqlLayer
	err := ctx.RegisterResource("aws:opsworks/mysqlLayer:MysqlLayer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMysqlLayer gets an existing MysqlLayer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMysqlLayer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MysqlLayerState, opts ...pulumi.ResourceOption) (*MysqlLayer, error) {
	var resource MysqlLayer
	err := ctx.ReadResource("aws:opsworks/mysqlLayer:MysqlLayer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MysqlLayer resources.
type mysqlLayerState struct {
	// The Amazon Resource Name(ARN) of the layer.
	Arn *string `pulumi:"arn"`
	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps *bool `pulumi:"autoAssignElasticIps"`
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps *bool `pulumi:"autoAssignPublicIps"`
	// Whether to enable auto-healing for the layer.
	AutoHealing             *bool                              `pulumi:"autoHealing"`
	CloudwatchConfiguration *MysqlLayerCloudwatchConfiguration `pulumi:"cloudwatchConfiguration"`
	CustomConfigureRecipes  []string                           `pulumi:"customConfigureRecipes"`
	CustomDeployRecipes     []string                           `pulumi:"customDeployRecipes"`
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
	EbsVolumes []MysqlLayerEbsVolume `pulumi:"ebsVolumes"`
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer *string `pulumi:"elasticLoadBalancer"`
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot *bool `pulumi:"installUpdatesOnBoot"`
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout *int `pulumi:"instanceShutdownTimeout"`
	// A human-readable name for the layer.
	Name *string `pulumi:"name"`
	// Root password to use for MySQL.
	RootPassword *string `pulumi:"rootPassword"`
	// Whether to set the root user password to all instances in the stack so they can access the instances in this layer.
	RootPasswordOnAllInstances *bool `pulumi:"rootPasswordOnAllInstances"`
	// The id of the stack the layer will belong to.
	StackId *string `pulumi:"stackId"`
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages []string `pulumi:"systemPackages"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances *bool `pulumi:"useEbsOptimizedInstances"`
}

type MysqlLayerState struct {
	// The Amazon Resource Name(ARN) of the layer.
	Arn pulumi.StringPtrInput
	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps pulumi.BoolPtrInput
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps pulumi.BoolPtrInput
	// Whether to enable auto-healing for the layer.
	AutoHealing             pulumi.BoolPtrInput
	CloudwatchConfiguration MysqlLayerCloudwatchConfigurationPtrInput
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
	EbsVolumes MysqlLayerEbsVolumeArrayInput
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer pulumi.StringPtrInput
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot pulumi.BoolPtrInput
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout pulumi.IntPtrInput
	// A human-readable name for the layer.
	Name pulumi.StringPtrInput
	// Root password to use for MySQL.
	RootPassword pulumi.StringPtrInput
	// Whether to set the root user password to all instances in the stack so they can access the instances in this layer.
	RootPasswordOnAllInstances pulumi.BoolPtrInput
	// The id of the stack the layer will belong to.
	StackId pulumi.StringPtrInput
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages pulumi.StringArrayInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances pulumi.BoolPtrInput
}

func (MysqlLayerState) ElementType() reflect.Type {
	return reflect.TypeOf((*mysqlLayerState)(nil)).Elem()
}

type mysqlLayerArgs struct {
	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps *bool `pulumi:"autoAssignElasticIps"`
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps *bool `pulumi:"autoAssignPublicIps"`
	// Whether to enable auto-healing for the layer.
	AutoHealing             *bool                              `pulumi:"autoHealing"`
	CloudwatchConfiguration *MysqlLayerCloudwatchConfiguration `pulumi:"cloudwatchConfiguration"`
	CustomConfigureRecipes  []string                           `pulumi:"customConfigureRecipes"`
	CustomDeployRecipes     []string                           `pulumi:"customDeployRecipes"`
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
	EbsVolumes []MysqlLayerEbsVolume `pulumi:"ebsVolumes"`
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer *string `pulumi:"elasticLoadBalancer"`
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot *bool `pulumi:"installUpdatesOnBoot"`
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout *int `pulumi:"instanceShutdownTimeout"`
	// A human-readable name for the layer.
	Name *string `pulumi:"name"`
	// Root password to use for MySQL.
	RootPassword *string `pulumi:"rootPassword"`
	// Whether to set the root user password to all instances in the stack so they can access the instances in this layer.
	RootPasswordOnAllInstances *bool `pulumi:"rootPasswordOnAllInstances"`
	// The id of the stack the layer will belong to.
	StackId string `pulumi:"stackId"`
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages []string `pulumi:"systemPackages"`
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances *bool `pulumi:"useEbsOptimizedInstances"`
}

// The set of arguments for constructing a MysqlLayer resource.
type MysqlLayerArgs struct {
	// Whether to automatically assign an elastic IP address to the layer's instances.
	AutoAssignElasticIps pulumi.BoolPtrInput
	// For stacks belonging to a VPC, whether to automatically assign a public IP address to each of the layer's instances.
	AutoAssignPublicIps pulumi.BoolPtrInput
	// Whether to enable auto-healing for the layer.
	AutoHealing             pulumi.BoolPtrInput
	CloudwatchConfiguration MysqlLayerCloudwatchConfigurationPtrInput
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
	EbsVolumes MysqlLayerEbsVolumeArrayInput
	// Name of an Elastic Load Balancer to attach to this layer
	ElasticLoadBalancer pulumi.StringPtrInput
	// Whether to install OS and package updates on each instance when it boots.
	InstallUpdatesOnBoot pulumi.BoolPtrInput
	// The time, in seconds, that OpsWorks will wait for Chef to complete after triggering the Shutdown event.
	InstanceShutdownTimeout pulumi.IntPtrInput
	// A human-readable name for the layer.
	Name pulumi.StringPtrInput
	// Root password to use for MySQL.
	RootPassword pulumi.StringPtrInput
	// Whether to set the root user password to all instances in the stack so they can access the instances in this layer.
	RootPasswordOnAllInstances pulumi.BoolPtrInput
	// The id of the stack the layer will belong to.
	StackId pulumi.StringInput
	// Names of a set of system packages to install on the layer's instances.
	SystemPackages pulumi.StringArrayInput
	// A map of tags to assign to the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Whether to use EBS-optimized instances.
	UseEbsOptimizedInstances pulumi.BoolPtrInput
}

func (MysqlLayerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mysqlLayerArgs)(nil)).Elem()
}

type MysqlLayerInput interface {
	pulumi.Input

	ToMysqlLayerOutput() MysqlLayerOutput
	ToMysqlLayerOutputWithContext(ctx context.Context) MysqlLayerOutput
}

func (*MysqlLayer) ElementType() reflect.Type {
	return reflect.TypeOf((**MysqlLayer)(nil)).Elem()
}

func (i *MysqlLayer) ToMysqlLayerOutput() MysqlLayerOutput {
	return i.ToMysqlLayerOutputWithContext(context.Background())
}

func (i *MysqlLayer) ToMysqlLayerOutputWithContext(ctx context.Context) MysqlLayerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MysqlLayerOutput)
}

// MysqlLayerArrayInput is an input type that accepts MysqlLayerArray and MysqlLayerArrayOutput values.
// You can construct a concrete instance of `MysqlLayerArrayInput` via:
//
//          MysqlLayerArray{ MysqlLayerArgs{...} }
type MysqlLayerArrayInput interface {
	pulumi.Input

	ToMysqlLayerArrayOutput() MysqlLayerArrayOutput
	ToMysqlLayerArrayOutputWithContext(context.Context) MysqlLayerArrayOutput
}

type MysqlLayerArray []MysqlLayerInput

func (MysqlLayerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MysqlLayer)(nil)).Elem()
}

func (i MysqlLayerArray) ToMysqlLayerArrayOutput() MysqlLayerArrayOutput {
	return i.ToMysqlLayerArrayOutputWithContext(context.Background())
}

func (i MysqlLayerArray) ToMysqlLayerArrayOutputWithContext(ctx context.Context) MysqlLayerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MysqlLayerArrayOutput)
}

// MysqlLayerMapInput is an input type that accepts MysqlLayerMap and MysqlLayerMapOutput values.
// You can construct a concrete instance of `MysqlLayerMapInput` via:
//
//          MysqlLayerMap{ "key": MysqlLayerArgs{...} }
type MysqlLayerMapInput interface {
	pulumi.Input

	ToMysqlLayerMapOutput() MysqlLayerMapOutput
	ToMysqlLayerMapOutputWithContext(context.Context) MysqlLayerMapOutput
}

type MysqlLayerMap map[string]MysqlLayerInput

func (MysqlLayerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MysqlLayer)(nil)).Elem()
}

func (i MysqlLayerMap) ToMysqlLayerMapOutput() MysqlLayerMapOutput {
	return i.ToMysqlLayerMapOutputWithContext(context.Background())
}

func (i MysqlLayerMap) ToMysqlLayerMapOutputWithContext(ctx context.Context) MysqlLayerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MysqlLayerMapOutput)
}

type MysqlLayerOutput struct{ *pulumi.OutputState }

func (MysqlLayerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MysqlLayer)(nil)).Elem()
}

func (o MysqlLayerOutput) ToMysqlLayerOutput() MysqlLayerOutput {
	return o
}

func (o MysqlLayerOutput) ToMysqlLayerOutputWithContext(ctx context.Context) MysqlLayerOutput {
	return o
}

type MysqlLayerArrayOutput struct{ *pulumi.OutputState }

func (MysqlLayerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*MysqlLayer)(nil)).Elem()
}

func (o MysqlLayerArrayOutput) ToMysqlLayerArrayOutput() MysqlLayerArrayOutput {
	return o
}

func (o MysqlLayerArrayOutput) ToMysqlLayerArrayOutputWithContext(ctx context.Context) MysqlLayerArrayOutput {
	return o
}

func (o MysqlLayerArrayOutput) Index(i pulumi.IntInput) MysqlLayerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *MysqlLayer {
		return vs[0].([]*MysqlLayer)[vs[1].(int)]
	}).(MysqlLayerOutput)
}

type MysqlLayerMapOutput struct{ *pulumi.OutputState }

func (MysqlLayerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*MysqlLayer)(nil)).Elem()
}

func (o MysqlLayerMapOutput) ToMysqlLayerMapOutput() MysqlLayerMapOutput {
	return o
}

func (o MysqlLayerMapOutput) ToMysqlLayerMapOutputWithContext(ctx context.Context) MysqlLayerMapOutput {
	return o
}

func (o MysqlLayerMapOutput) MapIndex(k pulumi.StringInput) MysqlLayerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *MysqlLayer {
		return vs[0].(map[string]*MysqlLayer)[vs[1].(string)]
	}).(MysqlLayerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MysqlLayerInput)(nil)).Elem(), &MysqlLayer{})
	pulumi.RegisterInputType(reflect.TypeOf((*MysqlLayerArrayInput)(nil)).Elem(), MysqlLayerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MysqlLayerMapInput)(nil)).Elem(), MysqlLayerMap{})
	pulumi.RegisterOutputType(MysqlLayerOutput{})
	pulumi.RegisterOutputType(MysqlLayerArrayOutput{})
	pulumi.RegisterOutputType(MysqlLayerMapOutput{})
}
