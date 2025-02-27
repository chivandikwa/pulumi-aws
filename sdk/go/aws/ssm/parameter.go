// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package ssm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an SSM Parameter resource.
//
// ## Example Usage
//
// To store a basic string parameter:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ssm"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := ssm.NewParameter(ctx, "foo", &ssm.ParameterArgs{
// 			Type:  pulumi.String("String"),
// 			Value: pulumi.String("bar"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// To store an encrypted string using the default SSM KMS key:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/rds"
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ssm"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := rds.NewInstance(ctx, "default", &rds.InstanceArgs{
// 			AllocatedStorage:   pulumi.Int(10),
// 			StorageType:        pulumi.String("gp2"),
// 			Engine:             pulumi.String("mysql"),
// 			EngineVersion:      pulumi.String("5.7.16"),
// 			InstanceClass:      pulumi.String("db.t2.micro"),
// 			Name:               pulumi.String("mydb"),
// 			Username:           pulumi.String("foo"),
// 			Password:           pulumi.Any(_var.Database_master_password),
// 			DbSubnetGroupName:  pulumi.String("my_database_subnet_group"),
// 			ParameterGroupName: pulumi.String("default.mysql5.7"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = ssm.NewParameter(ctx, "secret", &ssm.ParameterArgs{
// 			Description: pulumi.String("The parameter description"),
// 			Type:        pulumi.String("SecureString"),
// 			Value:       pulumi.Any(_var.Database_master_password),
// 			Tags: pulumi.StringMap{
// 				"environment": pulumi.String("production"),
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
// SSM Parameters can be imported using the `parameter store name`, e.g.,
//
// ```sh
//  $ pulumi import aws:ssm/parameter:Parameter my_param /my_path/my_paramname
// ```
type Parameter struct {
	pulumi.CustomResourceState

	// A regular expression used to validate the parameter value.
	AllowedPattern pulumi.StringPtrOutput `pulumi:"allowedPattern"`
	// The ARN of the parameter.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The dataType of the parameter. Valid values: text and aws:ec2:image for AMI format, see the [Native parameter support for Amazon Machine Image IDs
	// ](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-ec2-aliases.html)
	DataType pulumi.StringOutput `pulumi:"dataType"`
	// The description of the parameter.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The KMS key id or arn for encrypting a SecureString.
	KeyId pulumi.StringOutput `pulumi:"keyId"`
	// The name of the parameter. If the name contains a path (e.g., any forward slashes (`/`)), it must be fully qualified with a leading forward slash (`/`). For additional requirements and constraints, see the [AWS SSM User Guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-parameter-name-constraints.html).
	Name pulumi.StringOutput `pulumi:"name"`
	// Overwrite an existing parameter. If not specified, will default to `false` if the resource has not been created by this provider to avoid overwrite of existing resource and will default to `true` otherwise (lifecycle rules should then be used to manage the update behavior).
	Overwrite pulumi.BoolPtrOutput `pulumi:"overwrite"`
	// A map of tags to assign to the object. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The tier of the parameter. If not specified, will default to `Standard`. Valid tiers are `Standard`, `Advanced`, and `Intelligent-Tiering`. For more information on parameter tiers, see the [AWS SSM Parameter tier comparison and guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html).
	Tier pulumi.StringPtrOutput `pulumi:"tier"`
	// The type of the parameter. Valid types are `String`, `StringList` and `SecureString`.
	Type pulumi.StringOutput `pulumi:"type"`
	// The value of the parameter.
	Value pulumi.StringOutput `pulumi:"value"`
	// The version of the parameter.
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewParameter registers a new resource with the given unique name, arguments, and options.
func NewParameter(ctx *pulumi.Context,
	name string, args *ParameterArgs, opts ...pulumi.ResourceOption) (*Parameter, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	if args.Value == nil {
		return nil, errors.New("invalid value for required argument 'Value'")
	}
	var resource Parameter
	err := ctx.RegisterResource("aws:ssm/parameter:Parameter", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetParameter gets an existing Parameter resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetParameter(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ParameterState, opts ...pulumi.ResourceOption) (*Parameter, error) {
	var resource Parameter
	err := ctx.ReadResource("aws:ssm/parameter:Parameter", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Parameter resources.
type parameterState struct {
	// A regular expression used to validate the parameter value.
	AllowedPattern *string `pulumi:"allowedPattern"`
	// The ARN of the parameter.
	Arn *string `pulumi:"arn"`
	// The dataType of the parameter. Valid values: text and aws:ec2:image for AMI format, see the [Native parameter support for Amazon Machine Image IDs
	// ](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-ec2-aliases.html)
	DataType *string `pulumi:"dataType"`
	// The description of the parameter.
	Description *string `pulumi:"description"`
	// The KMS key id or arn for encrypting a SecureString.
	KeyId *string `pulumi:"keyId"`
	// The name of the parameter. If the name contains a path (e.g., any forward slashes (`/`)), it must be fully qualified with a leading forward slash (`/`). For additional requirements and constraints, see the [AWS SSM User Guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-parameter-name-constraints.html).
	Name *string `pulumi:"name"`
	// Overwrite an existing parameter. If not specified, will default to `false` if the resource has not been created by this provider to avoid overwrite of existing resource and will default to `true` otherwise (lifecycle rules should then be used to manage the update behavior).
	Overwrite *bool `pulumi:"overwrite"`
	// A map of tags to assign to the object. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The tier of the parameter. If not specified, will default to `Standard`. Valid tiers are `Standard`, `Advanced`, and `Intelligent-Tiering`. For more information on parameter tiers, see the [AWS SSM Parameter tier comparison and guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html).
	Tier *string `pulumi:"tier"`
	// The type of the parameter. Valid types are `String`, `StringList` and `SecureString`.
	Type *string `pulumi:"type"`
	// The value of the parameter.
	Value *string `pulumi:"value"`
	// The version of the parameter.
	Version *int `pulumi:"version"`
}

type ParameterState struct {
	// A regular expression used to validate the parameter value.
	AllowedPattern pulumi.StringPtrInput
	// The ARN of the parameter.
	Arn pulumi.StringPtrInput
	// The dataType of the parameter. Valid values: text and aws:ec2:image for AMI format, see the [Native parameter support for Amazon Machine Image IDs
	// ](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-ec2-aliases.html)
	DataType pulumi.StringPtrInput
	// The description of the parameter.
	Description pulumi.StringPtrInput
	// The KMS key id or arn for encrypting a SecureString.
	KeyId pulumi.StringPtrInput
	// The name of the parameter. If the name contains a path (e.g., any forward slashes (`/`)), it must be fully qualified with a leading forward slash (`/`). For additional requirements and constraints, see the [AWS SSM User Guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-parameter-name-constraints.html).
	Name pulumi.StringPtrInput
	// Overwrite an existing parameter. If not specified, will default to `false` if the resource has not been created by this provider to avoid overwrite of existing resource and will default to `true` otherwise (lifecycle rules should then be used to manage the update behavior).
	Overwrite pulumi.BoolPtrInput
	// A map of tags to assign to the object. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// The tier of the parameter. If not specified, will default to `Standard`. Valid tiers are `Standard`, `Advanced`, and `Intelligent-Tiering`. For more information on parameter tiers, see the [AWS SSM Parameter tier comparison and guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html).
	Tier pulumi.StringPtrInput
	// The type of the parameter. Valid types are `String`, `StringList` and `SecureString`.
	Type pulumi.StringPtrInput
	// The value of the parameter.
	Value pulumi.StringPtrInput
	// The version of the parameter.
	Version pulumi.IntPtrInput
}

func (ParameterState) ElementType() reflect.Type {
	return reflect.TypeOf((*parameterState)(nil)).Elem()
}

type parameterArgs struct {
	// A regular expression used to validate the parameter value.
	AllowedPattern *string `pulumi:"allowedPattern"`
	// The ARN of the parameter.
	Arn *string `pulumi:"arn"`
	// The dataType of the parameter. Valid values: text and aws:ec2:image for AMI format, see the [Native parameter support for Amazon Machine Image IDs
	// ](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-ec2-aliases.html)
	DataType *string `pulumi:"dataType"`
	// The description of the parameter.
	Description *string `pulumi:"description"`
	// The KMS key id or arn for encrypting a SecureString.
	KeyId *string `pulumi:"keyId"`
	// The name of the parameter. If the name contains a path (e.g., any forward slashes (`/`)), it must be fully qualified with a leading forward slash (`/`). For additional requirements and constraints, see the [AWS SSM User Guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-parameter-name-constraints.html).
	Name *string `pulumi:"name"`
	// Overwrite an existing parameter. If not specified, will default to `false` if the resource has not been created by this provider to avoid overwrite of existing resource and will default to `true` otherwise (lifecycle rules should then be used to manage the update behavior).
	Overwrite *bool `pulumi:"overwrite"`
	// A map of tags to assign to the object. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// The tier of the parameter. If not specified, will default to `Standard`. Valid tiers are `Standard`, `Advanced`, and `Intelligent-Tiering`. For more information on parameter tiers, see the [AWS SSM Parameter tier comparison and guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html).
	Tier *string `pulumi:"tier"`
	// The type of the parameter. Valid types are `String`, `StringList` and `SecureString`.
	Type string `pulumi:"type"`
	// The value of the parameter.
	Value string `pulumi:"value"`
}

// The set of arguments for constructing a Parameter resource.
type ParameterArgs struct {
	// A regular expression used to validate the parameter value.
	AllowedPattern pulumi.StringPtrInput
	// The ARN of the parameter.
	Arn pulumi.StringPtrInput
	// The dataType of the parameter. Valid values: text and aws:ec2:image for AMI format, see the [Native parameter support for Amazon Machine Image IDs
	// ](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-ec2-aliases.html)
	DataType pulumi.StringPtrInput
	// The description of the parameter.
	Description pulumi.StringPtrInput
	// The KMS key id or arn for encrypting a SecureString.
	KeyId pulumi.StringPtrInput
	// The name of the parameter. If the name contains a path (e.g., any forward slashes (`/`)), it must be fully qualified with a leading forward slash (`/`). For additional requirements and constraints, see the [AWS SSM User Guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/sysman-parameter-name-constraints.html).
	Name pulumi.StringPtrInput
	// Overwrite an existing parameter. If not specified, will default to `false` if the resource has not been created by this provider to avoid overwrite of existing resource and will default to `true` otherwise (lifecycle rules should then be used to manage the update behavior).
	Overwrite pulumi.BoolPtrInput
	// A map of tags to assign to the object. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// The tier of the parameter. If not specified, will default to `Standard`. Valid tiers are `Standard`, `Advanced`, and `Intelligent-Tiering`. For more information on parameter tiers, see the [AWS SSM Parameter tier comparison and guide](https://docs.aws.amazon.com/systems-manager/latest/userguide/parameter-store-advanced-parameters.html).
	Tier pulumi.StringPtrInput
	// The type of the parameter. Valid types are `String`, `StringList` and `SecureString`.
	Type pulumi.StringInput
	// The value of the parameter.
	Value pulumi.StringInput
}

func (ParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*parameterArgs)(nil)).Elem()
}

type ParameterInput interface {
	pulumi.Input

	ToParameterOutput() ParameterOutput
	ToParameterOutputWithContext(ctx context.Context) ParameterOutput
}

func (*Parameter) ElementType() reflect.Type {
	return reflect.TypeOf((**Parameter)(nil)).Elem()
}

func (i *Parameter) ToParameterOutput() ParameterOutput {
	return i.ToParameterOutputWithContext(context.Background())
}

func (i *Parameter) ToParameterOutputWithContext(ctx context.Context) ParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterOutput)
}

// ParameterArrayInput is an input type that accepts ParameterArray and ParameterArrayOutput values.
// You can construct a concrete instance of `ParameterArrayInput` via:
//
//          ParameterArray{ ParameterArgs{...} }
type ParameterArrayInput interface {
	pulumi.Input

	ToParameterArrayOutput() ParameterArrayOutput
	ToParameterArrayOutputWithContext(context.Context) ParameterArrayOutput
}

type ParameterArray []ParameterInput

func (ParameterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Parameter)(nil)).Elem()
}

func (i ParameterArray) ToParameterArrayOutput() ParameterArrayOutput {
	return i.ToParameterArrayOutputWithContext(context.Background())
}

func (i ParameterArray) ToParameterArrayOutputWithContext(ctx context.Context) ParameterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterArrayOutput)
}

// ParameterMapInput is an input type that accepts ParameterMap and ParameterMapOutput values.
// You can construct a concrete instance of `ParameterMapInput` via:
//
//          ParameterMap{ "key": ParameterArgs{...} }
type ParameterMapInput interface {
	pulumi.Input

	ToParameterMapOutput() ParameterMapOutput
	ToParameterMapOutputWithContext(context.Context) ParameterMapOutput
}

type ParameterMap map[string]ParameterInput

func (ParameterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Parameter)(nil)).Elem()
}

func (i ParameterMap) ToParameterMapOutput() ParameterMapOutput {
	return i.ToParameterMapOutputWithContext(context.Background())
}

func (i ParameterMap) ToParameterMapOutputWithContext(ctx context.Context) ParameterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParameterMapOutput)
}

type ParameterOutput struct{ *pulumi.OutputState }

func (ParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Parameter)(nil)).Elem()
}

func (o ParameterOutput) ToParameterOutput() ParameterOutput {
	return o
}

func (o ParameterOutput) ToParameterOutputWithContext(ctx context.Context) ParameterOutput {
	return o
}

type ParameterArrayOutput struct{ *pulumi.OutputState }

func (ParameterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Parameter)(nil)).Elem()
}

func (o ParameterArrayOutput) ToParameterArrayOutput() ParameterArrayOutput {
	return o
}

func (o ParameterArrayOutput) ToParameterArrayOutputWithContext(ctx context.Context) ParameterArrayOutput {
	return o
}

func (o ParameterArrayOutput) Index(i pulumi.IntInput) ParameterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Parameter {
		return vs[0].([]*Parameter)[vs[1].(int)]
	}).(ParameterOutput)
}

type ParameterMapOutput struct{ *pulumi.OutputState }

func (ParameterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Parameter)(nil)).Elem()
}

func (o ParameterMapOutput) ToParameterMapOutput() ParameterMapOutput {
	return o
}

func (o ParameterMapOutput) ToParameterMapOutputWithContext(ctx context.Context) ParameterMapOutput {
	return o
}

func (o ParameterMapOutput) MapIndex(k pulumi.StringInput) ParameterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Parameter {
		return vs[0].(map[string]*Parameter)[vs[1].(string)]
	}).(ParameterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterInput)(nil)).Elem(), &Parameter{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterArrayInput)(nil)).Elem(), ParameterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterMapInput)(nil)).Elem(), ParameterMap{})
	pulumi.RegisterOutputType(ParameterOutput{})
	pulumi.RegisterOutputType(ParameterArrayOutput{})
	pulumi.RegisterOutputType(ParameterMapOutput{})
}
