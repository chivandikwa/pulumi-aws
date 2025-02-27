// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package imagebuilder

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an Image Builder Image Recipe.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/imagebuilder"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := imagebuilder.NewImageRecipe(ctx, "example", &imagebuilder.ImageRecipeArgs{
// 			BlockDeviceMappings: imagebuilder.ImageRecipeBlockDeviceMappingArray{
// 				&imagebuilder.ImageRecipeBlockDeviceMappingArgs{
// 					DeviceName: pulumi.String("/dev/xvdb"),
// 					Ebs: &imagebuilder.ImageRecipeBlockDeviceMappingEbsArgs{
// 						DeleteOnTermination: pulumi.String("true"),
// 						VolumeSize:          pulumi.Int(100),
// 						VolumeType:          pulumi.String("gp2"),
// 					},
// 				},
// 			},
// 			Components: imagebuilder.ImageRecipeComponentArray{
// 				&imagebuilder.ImageRecipeComponentArgs{
// 					ComponentArn: pulumi.Any(aws_imagebuilder_component.Example.Arn),
// 					Parameters: imagebuilder.ImageRecipeComponentParameterArray{
// 						&imagebuilder.ImageRecipeComponentParameterArgs{
// 							Name:  pulumi.String("Parameter1"),
// 							Value: pulumi.String("Value1"),
// 						},
// 						&imagebuilder.ImageRecipeComponentParameterArgs{
// 							Name:  pulumi.String("Parameter2"),
// 							Value: pulumi.String("Value2"),
// 						},
// 					},
// 				},
// 			},
// 			ParentImage: pulumi.String(fmt.Sprintf("%v%v%v%v%v", "arn:", data.Aws_partition.Current.Partition, ":imagebuilder:", data.Aws_region.Current.Name, ":aws:image/amazon-linux-2-x86/x.x.x")),
// 			Version:     pulumi.String("1.0.0"),
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
// `aws_imagebuilder_image_recipe` resources can be imported by using the Amazon Resource Name (ARN), e.g.,
//
// ```sh
//  $ pulumi import aws:imagebuilder/imageRecipe:ImageRecipe example arn:aws:imagebuilder:us-east-1:123456789012:image-recipe/example/1.0.0
// ```
type ImageRecipe struct {
	pulumi.CustomResourceState

	// (Required) Amazon Resource Name (ARN) of the image recipe.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Configuration block(s) with block device mappings for the the image recipe. Detailed below.
	BlockDeviceMappings ImageRecipeBlockDeviceMappingArrayOutput `pulumi:"blockDeviceMappings"`
	// Ordered configuration block(s) with components for the image recipe. Detailed below.
	Components ImageRecipeComponentArrayOutput `pulumi:"components"`
	// Date the image recipe was created.
	DateCreated pulumi.StringOutput `pulumi:"dateCreated"`
	// Description of the image recipe.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The name of the component parameter.
	Name pulumi.StringOutput `pulumi:"name"`
	// Owner of the image recipe.
	Owner pulumi.StringOutput `pulumi:"owner"`
	// Platform of the image recipe.
	ParentImage pulumi.StringOutput `pulumi:"parentImage"`
	// Platform of the image recipe.
	Platform pulumi.StringOutput `pulumi:"platform"`
	// Key-value map of resource tags for the image recipe. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Base64 encoded user data. Use this to provide commands or a command script to run when you launch your build instance.
	UserDataBase64 pulumi.StringOutput `pulumi:"userDataBase64"`
	// Version of the image recipe.
	Version pulumi.StringOutput `pulumi:"version"`
	// The working directory to be used during build and test workflows.
	WorkingDirectory pulumi.StringPtrOutput `pulumi:"workingDirectory"`
}

// NewImageRecipe registers a new resource with the given unique name, arguments, and options.
func NewImageRecipe(ctx *pulumi.Context,
	name string, args *ImageRecipeArgs, opts ...pulumi.ResourceOption) (*ImageRecipe, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Components == nil {
		return nil, errors.New("invalid value for required argument 'Components'")
	}
	if args.ParentImage == nil {
		return nil, errors.New("invalid value for required argument 'ParentImage'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	var resource ImageRecipe
	err := ctx.RegisterResource("aws:imagebuilder/imageRecipe:ImageRecipe", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetImageRecipe gets an existing ImageRecipe resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetImageRecipe(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ImageRecipeState, opts ...pulumi.ResourceOption) (*ImageRecipe, error) {
	var resource ImageRecipe
	err := ctx.ReadResource("aws:imagebuilder/imageRecipe:ImageRecipe", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ImageRecipe resources.
type imageRecipeState struct {
	// (Required) Amazon Resource Name (ARN) of the image recipe.
	Arn *string `pulumi:"arn"`
	// Configuration block(s) with block device mappings for the the image recipe. Detailed below.
	BlockDeviceMappings []ImageRecipeBlockDeviceMapping `pulumi:"blockDeviceMappings"`
	// Ordered configuration block(s) with components for the image recipe. Detailed below.
	Components []ImageRecipeComponent `pulumi:"components"`
	// Date the image recipe was created.
	DateCreated *string `pulumi:"dateCreated"`
	// Description of the image recipe.
	Description *string `pulumi:"description"`
	// The name of the component parameter.
	Name *string `pulumi:"name"`
	// Owner of the image recipe.
	Owner *string `pulumi:"owner"`
	// Platform of the image recipe.
	ParentImage *string `pulumi:"parentImage"`
	// Platform of the image recipe.
	Platform *string `pulumi:"platform"`
	// Key-value map of resource tags for the image recipe. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Base64 encoded user data. Use this to provide commands or a command script to run when you launch your build instance.
	UserDataBase64 *string `pulumi:"userDataBase64"`
	// Version of the image recipe.
	Version *string `pulumi:"version"`
	// The working directory to be used during build and test workflows.
	WorkingDirectory *string `pulumi:"workingDirectory"`
}

type ImageRecipeState struct {
	// (Required) Amazon Resource Name (ARN) of the image recipe.
	Arn pulumi.StringPtrInput
	// Configuration block(s) with block device mappings for the the image recipe. Detailed below.
	BlockDeviceMappings ImageRecipeBlockDeviceMappingArrayInput
	// Ordered configuration block(s) with components for the image recipe. Detailed below.
	Components ImageRecipeComponentArrayInput
	// Date the image recipe was created.
	DateCreated pulumi.StringPtrInput
	// Description of the image recipe.
	Description pulumi.StringPtrInput
	// The name of the component parameter.
	Name pulumi.StringPtrInput
	// Owner of the image recipe.
	Owner pulumi.StringPtrInput
	// Platform of the image recipe.
	ParentImage pulumi.StringPtrInput
	// Platform of the image recipe.
	Platform pulumi.StringPtrInput
	// Key-value map of resource tags for the image recipe. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// Base64 encoded user data. Use this to provide commands or a command script to run when you launch your build instance.
	UserDataBase64 pulumi.StringPtrInput
	// Version of the image recipe.
	Version pulumi.StringPtrInput
	// The working directory to be used during build and test workflows.
	WorkingDirectory pulumi.StringPtrInput
}

func (ImageRecipeState) ElementType() reflect.Type {
	return reflect.TypeOf((*imageRecipeState)(nil)).Elem()
}

type imageRecipeArgs struct {
	// Configuration block(s) with block device mappings for the the image recipe. Detailed below.
	BlockDeviceMappings []ImageRecipeBlockDeviceMapping `pulumi:"blockDeviceMappings"`
	// Ordered configuration block(s) with components for the image recipe. Detailed below.
	Components []ImageRecipeComponent `pulumi:"components"`
	// Description of the image recipe.
	Description *string `pulumi:"description"`
	// The name of the component parameter.
	Name *string `pulumi:"name"`
	// Platform of the image recipe.
	ParentImage string `pulumi:"parentImage"`
	// Key-value map of resource tags for the image recipe. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Base64 encoded user data. Use this to provide commands or a command script to run when you launch your build instance.
	UserDataBase64 *string `pulumi:"userDataBase64"`
	// Version of the image recipe.
	Version string `pulumi:"version"`
	// The working directory to be used during build and test workflows.
	WorkingDirectory *string `pulumi:"workingDirectory"`
}

// The set of arguments for constructing a ImageRecipe resource.
type ImageRecipeArgs struct {
	// Configuration block(s) with block device mappings for the the image recipe. Detailed below.
	BlockDeviceMappings ImageRecipeBlockDeviceMappingArrayInput
	// Ordered configuration block(s) with components for the image recipe. Detailed below.
	Components ImageRecipeComponentArrayInput
	// Description of the image recipe.
	Description pulumi.StringPtrInput
	// The name of the component parameter.
	Name pulumi.StringPtrInput
	// Platform of the image recipe.
	ParentImage pulumi.StringInput
	// Key-value map of resource tags for the image recipe. If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Base64 encoded user data. Use this to provide commands or a command script to run when you launch your build instance.
	UserDataBase64 pulumi.StringPtrInput
	// Version of the image recipe.
	Version pulumi.StringInput
	// The working directory to be used during build and test workflows.
	WorkingDirectory pulumi.StringPtrInput
}

func (ImageRecipeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*imageRecipeArgs)(nil)).Elem()
}

type ImageRecipeInput interface {
	pulumi.Input

	ToImageRecipeOutput() ImageRecipeOutput
	ToImageRecipeOutputWithContext(ctx context.Context) ImageRecipeOutput
}

func (*ImageRecipe) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageRecipe)(nil)).Elem()
}

func (i *ImageRecipe) ToImageRecipeOutput() ImageRecipeOutput {
	return i.ToImageRecipeOutputWithContext(context.Background())
}

func (i *ImageRecipe) ToImageRecipeOutputWithContext(ctx context.Context) ImageRecipeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageRecipeOutput)
}

// ImageRecipeArrayInput is an input type that accepts ImageRecipeArray and ImageRecipeArrayOutput values.
// You can construct a concrete instance of `ImageRecipeArrayInput` via:
//
//          ImageRecipeArray{ ImageRecipeArgs{...} }
type ImageRecipeArrayInput interface {
	pulumi.Input

	ToImageRecipeArrayOutput() ImageRecipeArrayOutput
	ToImageRecipeArrayOutputWithContext(context.Context) ImageRecipeArrayOutput
}

type ImageRecipeArray []ImageRecipeInput

func (ImageRecipeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ImageRecipe)(nil)).Elem()
}

func (i ImageRecipeArray) ToImageRecipeArrayOutput() ImageRecipeArrayOutput {
	return i.ToImageRecipeArrayOutputWithContext(context.Background())
}

func (i ImageRecipeArray) ToImageRecipeArrayOutputWithContext(ctx context.Context) ImageRecipeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageRecipeArrayOutput)
}

// ImageRecipeMapInput is an input type that accepts ImageRecipeMap and ImageRecipeMapOutput values.
// You can construct a concrete instance of `ImageRecipeMapInput` via:
//
//          ImageRecipeMap{ "key": ImageRecipeArgs{...} }
type ImageRecipeMapInput interface {
	pulumi.Input

	ToImageRecipeMapOutput() ImageRecipeMapOutput
	ToImageRecipeMapOutputWithContext(context.Context) ImageRecipeMapOutput
}

type ImageRecipeMap map[string]ImageRecipeInput

func (ImageRecipeMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ImageRecipe)(nil)).Elem()
}

func (i ImageRecipeMap) ToImageRecipeMapOutput() ImageRecipeMapOutput {
	return i.ToImageRecipeMapOutputWithContext(context.Background())
}

func (i ImageRecipeMap) ToImageRecipeMapOutputWithContext(ctx context.Context) ImageRecipeMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageRecipeMapOutput)
}

type ImageRecipeOutput struct{ *pulumi.OutputState }

func (ImageRecipeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageRecipe)(nil)).Elem()
}

func (o ImageRecipeOutput) ToImageRecipeOutput() ImageRecipeOutput {
	return o
}

func (o ImageRecipeOutput) ToImageRecipeOutputWithContext(ctx context.Context) ImageRecipeOutput {
	return o
}

type ImageRecipeArrayOutput struct{ *pulumi.OutputState }

func (ImageRecipeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ImageRecipe)(nil)).Elem()
}

func (o ImageRecipeArrayOutput) ToImageRecipeArrayOutput() ImageRecipeArrayOutput {
	return o
}

func (o ImageRecipeArrayOutput) ToImageRecipeArrayOutputWithContext(ctx context.Context) ImageRecipeArrayOutput {
	return o
}

func (o ImageRecipeArrayOutput) Index(i pulumi.IntInput) ImageRecipeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ImageRecipe {
		return vs[0].([]*ImageRecipe)[vs[1].(int)]
	}).(ImageRecipeOutput)
}

type ImageRecipeMapOutput struct{ *pulumi.OutputState }

func (ImageRecipeMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ImageRecipe)(nil)).Elem()
}

func (o ImageRecipeMapOutput) ToImageRecipeMapOutput() ImageRecipeMapOutput {
	return o
}

func (o ImageRecipeMapOutput) ToImageRecipeMapOutputWithContext(ctx context.Context) ImageRecipeMapOutput {
	return o
}

func (o ImageRecipeMapOutput) MapIndex(k pulumi.StringInput) ImageRecipeOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ImageRecipe {
		return vs[0].(map[string]*ImageRecipe)[vs[1].(string)]
	}).(ImageRecipeOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ImageRecipeInput)(nil)).Elem(), &ImageRecipe{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageRecipeArrayInput)(nil)).Elem(), ImageRecipeArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ImageRecipeMapInput)(nil)).Elem(), ImageRecipeMap{})
	pulumi.RegisterOutputType(ImageRecipeOutput{})
	pulumi.RegisterOutputType(ImageRecipeArrayOutput{})
	pulumi.RegisterOutputType(ImageRecipeMapOutput{})
}
