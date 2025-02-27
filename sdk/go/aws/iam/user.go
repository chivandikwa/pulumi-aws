// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package iam

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides an IAM user.
//
// > *NOTE:* If policies are attached to the user via the `iam.PolicyAttachment` resource and you are modifying the user `name` or `path`, the `forceDestroy` argument must be set to `true` and applied before attempting the operation otherwise you will encounter a `DeleteConflict` error. The `iam.UserPolicyAttachment` resource (recommended) does not have this requirement.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		lbUser, err := iam.NewUser(ctx, "lbUser", &iam.UserArgs{
// 			Path: pulumi.String("/system/"),
// 			Tags: pulumi.StringMap{
// 				"tag-key": pulumi.String("tag-value"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewAccessKey(ctx, "lbAccessKey", &iam.AccessKeyArgs{
// 			User: lbUser.Name,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewUserPolicy(ctx, "lbRo", &iam.UserPolicyArgs{
// 			User:   lbUser.Name,
// 			Policy: pulumi.Any(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Action\": [\n", "        \"ec2:Describe*\"\n", "      ],\n", "      \"Effect\": \"Allow\",\n", "      \"Resource\": \"*\"\n", "    }\n", "  ]\n", "}\n")),
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
// IAM Users can be imported using the `name`, e.g.,
//
// ```sh
//  $ pulumi import aws:iam/user:User lb loadbalancer
// ```
type User struct {
	pulumi.CustomResourceState

	// The ARN assigned by AWS for this user.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// When destroying this user, destroy even if it
	// has non-provider-managed IAM access keys, login profile or MFA devices. Without `forceDestroy`
	// a user with non-provider-managed access keys and login profile will fail to be destroyed.
	ForceDestroy pulumi.BoolPtrOutput `pulumi:"forceDestroy"`
	// The user's name. The name must consist of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: `=,.@-_.`. User names are not distinguished by case. For example, you cannot create users named both "TESTUSER" and "testuser".
	Name pulumi.StringOutput `pulumi:"name"`
	// Path in which to create the user.
	Path pulumi.StringPtrOutput `pulumi:"path"`
	// The ARN of the policy that is used to set the permissions boundary for the user.
	PermissionsBoundary pulumi.StringPtrOutput `pulumi:"permissionsBoundary"`
	// Key-value mapping of tags for the IAM user
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// The [unique ID][1] assigned by AWS.
	UniqueId pulumi.StringOutput `pulumi:"uniqueId"`
}

// NewUser registers a new resource with the given unique name, arguments, and options.
func NewUser(ctx *pulumi.Context,
	name string, args *UserArgs, opts ...pulumi.ResourceOption) (*User, error) {
	if args == nil {
		args = &UserArgs{}
	}

	var resource User
	err := ctx.RegisterResource("aws:iam/user:User", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetUser gets an existing User resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetUser(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *UserState, opts ...pulumi.ResourceOption) (*User, error) {
	var resource User
	err := ctx.ReadResource("aws:iam/user:User", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering User resources.
type userState struct {
	// The ARN assigned by AWS for this user.
	Arn *string `pulumi:"arn"`
	// When destroying this user, destroy even if it
	// has non-provider-managed IAM access keys, login profile or MFA devices. Without `forceDestroy`
	// a user with non-provider-managed access keys and login profile will fail to be destroyed.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// The user's name. The name must consist of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: `=,.@-_.`. User names are not distinguished by case. For example, you cannot create users named both "TESTUSER" and "testuser".
	Name *string `pulumi:"name"`
	// Path in which to create the user.
	Path *string `pulumi:"path"`
	// The ARN of the policy that is used to set the permissions boundary for the user.
	PermissionsBoundary *string `pulumi:"permissionsBoundary"`
	// Key-value mapping of tags for the IAM user
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider.
	TagsAll map[string]string `pulumi:"tagsAll"`
	// The [unique ID][1] assigned by AWS.
	UniqueId *string `pulumi:"uniqueId"`
}

type UserState struct {
	// The ARN assigned by AWS for this user.
	Arn pulumi.StringPtrInput
	// When destroying this user, destroy even if it
	// has non-provider-managed IAM access keys, login profile or MFA devices. Without `forceDestroy`
	// a user with non-provider-managed access keys and login profile will fail to be destroyed.
	ForceDestroy pulumi.BoolPtrInput
	// The user's name. The name must consist of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: `=,.@-_.`. User names are not distinguished by case. For example, you cannot create users named both "TESTUSER" and "testuser".
	Name pulumi.StringPtrInput
	// Path in which to create the user.
	Path pulumi.StringPtrInput
	// The ARN of the policy that is used to set the permissions boundary for the user.
	PermissionsBoundary pulumi.StringPtrInput
	// Key-value mapping of tags for the IAM user
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider.
	TagsAll pulumi.StringMapInput
	// The [unique ID][1] assigned by AWS.
	UniqueId pulumi.StringPtrInput
}

func (UserState) ElementType() reflect.Type {
	return reflect.TypeOf((*userState)(nil)).Elem()
}

type userArgs struct {
	// When destroying this user, destroy even if it
	// has non-provider-managed IAM access keys, login profile or MFA devices. Without `forceDestroy`
	// a user with non-provider-managed access keys and login profile will fail to be destroyed.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// The user's name. The name must consist of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: `=,.@-_.`. User names are not distinguished by case. For example, you cannot create users named both "TESTUSER" and "testuser".
	Name *string `pulumi:"name"`
	// Path in which to create the user.
	Path *string `pulumi:"path"`
	// The ARN of the policy that is used to set the permissions boundary for the user.
	PermissionsBoundary *string `pulumi:"permissionsBoundary"`
	// Key-value mapping of tags for the IAM user
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a User resource.
type UserArgs struct {
	// When destroying this user, destroy even if it
	// has non-provider-managed IAM access keys, login profile or MFA devices. Without `forceDestroy`
	// a user with non-provider-managed access keys and login profile will fail to be destroyed.
	ForceDestroy pulumi.BoolPtrInput
	// The user's name. The name must consist of upper and lowercase alphanumeric characters with no spaces. You can also include any of the following characters: `=,.@-_.`. User names are not distinguished by case. For example, you cannot create users named both "TESTUSER" and "testuser".
	Name pulumi.StringPtrInput
	// Path in which to create the user.
	Path pulumi.StringPtrInput
	// The ARN of the policy that is used to set the permissions boundary for the user.
	PermissionsBoundary pulumi.StringPtrInput
	// Key-value mapping of tags for the IAM user
	Tags pulumi.StringMapInput
}

func (UserArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*userArgs)(nil)).Elem()
}

type UserInput interface {
	pulumi.Input

	ToUserOutput() UserOutput
	ToUserOutputWithContext(ctx context.Context) UserOutput
}

func (*User) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (i *User) ToUserOutput() UserOutput {
	return i.ToUserOutputWithContext(context.Background())
}

func (i *User) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserOutput)
}

// UserArrayInput is an input type that accepts UserArray and UserArrayOutput values.
// You can construct a concrete instance of `UserArrayInput` via:
//
//          UserArray{ UserArgs{...} }
type UserArrayInput interface {
	pulumi.Input

	ToUserArrayOutput() UserArrayOutput
	ToUserArrayOutputWithContext(context.Context) UserArrayOutput
}

type UserArray []UserInput

func (UserArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*User)(nil)).Elem()
}

func (i UserArray) ToUserArrayOutput() UserArrayOutput {
	return i.ToUserArrayOutputWithContext(context.Background())
}

func (i UserArray) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserArrayOutput)
}

// UserMapInput is an input type that accepts UserMap and UserMapOutput values.
// You can construct a concrete instance of `UserMapInput` via:
//
//          UserMap{ "key": UserArgs{...} }
type UserMapInput interface {
	pulumi.Input

	ToUserMapOutput() UserMapOutput
	ToUserMapOutputWithContext(context.Context) UserMapOutput
}

type UserMap map[string]UserInput

func (UserMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*User)(nil)).Elem()
}

func (i UserMap) ToUserMapOutput() UserMapOutput {
	return i.ToUserMapOutputWithContext(context.Background())
}

func (i UserMap) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserMapOutput)
}

type UserOutput struct{ *pulumi.OutputState }

func (UserOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**User)(nil)).Elem()
}

func (o UserOutput) ToUserOutput() UserOutput {
	return o
}

func (o UserOutput) ToUserOutputWithContext(ctx context.Context) UserOutput {
	return o
}

type UserArrayOutput struct{ *pulumi.OutputState }

func (UserArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*User)(nil)).Elem()
}

func (o UserArrayOutput) ToUserArrayOutput() UserArrayOutput {
	return o
}

func (o UserArrayOutput) ToUserArrayOutputWithContext(ctx context.Context) UserArrayOutput {
	return o
}

func (o UserArrayOutput) Index(i pulumi.IntInput) UserOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *User {
		return vs[0].([]*User)[vs[1].(int)]
	}).(UserOutput)
}

type UserMapOutput struct{ *pulumi.OutputState }

func (UserMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*User)(nil)).Elem()
}

func (o UserMapOutput) ToUserMapOutput() UserMapOutput {
	return o
}

func (o UserMapOutput) ToUserMapOutputWithContext(ctx context.Context) UserMapOutput {
	return o
}

func (o UserMapOutput) MapIndex(k pulumi.StringInput) UserOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *User {
		return vs[0].(map[string]*User)[vs[1].(string)]
	}).(UserOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*UserInput)(nil)).Elem(), &User{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserArrayInput)(nil)).Elem(), UserArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserMapInput)(nil)).Elem(), UserMap{})
	pulumi.RegisterOutputType(UserOutput{})
	pulumi.RegisterOutputType(UserArrayOutput{})
	pulumi.RegisterOutputType(UserMapOutput{})
}
