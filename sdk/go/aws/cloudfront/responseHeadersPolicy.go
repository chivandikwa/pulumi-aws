// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudfront

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a CloudFront response headers policy resource.
// A response headers policy contains information about a set of HTTP response headers and their values.
// After you create a response headers policy, you can use its ID to attach it to one or more cache behaviors in a CloudFront distribution.
// When it’s attached to a cache behavior, CloudFront adds the headers in the policy to every response that it sends for requests that match the cache behavior.
//
// ## Example Usage
//
// The example below creates a CloudFront response headers policy.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/cloudfront"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudfront.NewResponseHeadersPolicy(ctx, "example", &cloudfront.ResponseHeadersPolicyArgs{
// 			Comment: pulumi.String("test comment"),
// 			CorsConfig: &cloudfront.ResponseHeadersPolicyCorsConfigArgs{
// 				AccessControlAllowCredentials: pulumi.Bool(true),
// 				AccessControlAllowHeaders: &cloudfront.ResponseHeadersPolicyCorsConfigAccessControlAllowHeadersArgs{
// 					Items: pulumi.StringArray{
// 						pulumi.String("test"),
// 					},
// 				},
// 				AccessControlAllowMethods: &cloudfront.ResponseHeadersPolicyCorsConfigAccessControlAllowMethodsArgs{
// 					Items: pulumi.StringArray{
// 						pulumi.String("GET"),
// 					},
// 				},
// 				AccessControlAllowOrigins: &cloudfront.ResponseHeadersPolicyCorsConfigAccessControlAllowOriginsArgs{
// 					Items: pulumi.StringArray{
// 						pulumi.String("test.example.comtest"),
// 					},
// 				},
// 				OriginOverride: pulumi.Bool(true),
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
// The example below creates a CloudFront response headers policy with a custom headers config.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/cloudfront"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := cloudfront.NewResponseHeadersPolicy(ctx, "example", &cloudfront.ResponseHeadersPolicyArgs{
// 			CustomHeadersConfig: &cloudfront.ResponseHeadersPolicyCustomHeadersConfigArgs{
// 				Items: cloudfront.ResponseHeadersPolicyCustomHeadersConfigItemArray{
// 					&cloudfront.ResponseHeadersPolicyCustomHeadersConfigItemArgs{
// 						Header:   pulumi.String("X-Permitted-Cross-Domain-Policies"),
// 						Override: pulumi.Bool(true),
// 						Value:    pulumi.String("none"),
// 					},
// 					&cloudfront.ResponseHeadersPolicyCustomHeadersConfigItemArgs{
// 						Header:   pulumi.String("X-Test"),
// 						Override: pulumi.Bool(true),
// 						Value:    pulumi.String("none"),
// 					},
// 				},
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
// Cloudfront Response Headers Policies can be imported using the `id`, e.g.
//
// ```sh
//  $ pulumi import aws:cloudfront/responseHeadersPolicy:ResponseHeadersPolicy policy 658327ea-f89d-4fab-a63d-7e88639e58f9
// ```
type ResponseHeadersPolicy struct {
	pulumi.CustomResourceState

	// A comment to describe the response headers policy. The comment cannot be longer than 128 characters.
	Comment pulumi.StringPtrOutput `pulumi:"comment"`
	// A configuration for a set of HTTP response headers that are used for Cross-Origin Resource Sharing (CORS). See Cors Config for more information.
	CorsConfig ResponseHeadersPolicyCorsConfigPtrOutput `pulumi:"corsConfig"`
	// Object that contains an attribute `items` that contains a list of custom headers. See Custom Header for more information.
	CustomHeadersConfig ResponseHeadersPolicyCustomHeadersConfigPtrOutput `pulumi:"customHeadersConfig"`
	// The current version of the response headers policy.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// A unique name to identify the response headers policy.
	Name pulumi.StringOutput `pulumi:"name"`
	// A configuration for a set of security-related HTTP response headers. See Security Headers Config for more information.
	SecurityHeadersConfig ResponseHeadersPolicySecurityHeadersConfigPtrOutput `pulumi:"securityHeadersConfig"`
}

// NewResponseHeadersPolicy registers a new resource with the given unique name, arguments, and options.
func NewResponseHeadersPolicy(ctx *pulumi.Context,
	name string, args *ResponseHeadersPolicyArgs, opts ...pulumi.ResourceOption) (*ResponseHeadersPolicy, error) {
	if args == nil {
		args = &ResponseHeadersPolicyArgs{}
	}

	var resource ResponseHeadersPolicy
	err := ctx.RegisterResource("aws:cloudfront/responseHeadersPolicy:ResponseHeadersPolicy", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetResponseHeadersPolicy gets an existing ResponseHeadersPolicy resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetResponseHeadersPolicy(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ResponseHeadersPolicyState, opts ...pulumi.ResourceOption) (*ResponseHeadersPolicy, error) {
	var resource ResponseHeadersPolicy
	err := ctx.ReadResource("aws:cloudfront/responseHeadersPolicy:ResponseHeadersPolicy", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ResponseHeadersPolicy resources.
type responseHeadersPolicyState struct {
	// A comment to describe the response headers policy. The comment cannot be longer than 128 characters.
	Comment *string `pulumi:"comment"`
	// A configuration for a set of HTTP response headers that are used for Cross-Origin Resource Sharing (CORS). See Cors Config for more information.
	CorsConfig *ResponseHeadersPolicyCorsConfig `pulumi:"corsConfig"`
	// Object that contains an attribute `items` that contains a list of custom headers. See Custom Header for more information.
	CustomHeadersConfig *ResponseHeadersPolicyCustomHeadersConfig `pulumi:"customHeadersConfig"`
	// The current version of the response headers policy.
	Etag *string `pulumi:"etag"`
	// A unique name to identify the response headers policy.
	Name *string `pulumi:"name"`
	// A configuration for a set of security-related HTTP response headers. See Security Headers Config for more information.
	SecurityHeadersConfig *ResponseHeadersPolicySecurityHeadersConfig `pulumi:"securityHeadersConfig"`
}

type ResponseHeadersPolicyState struct {
	// A comment to describe the response headers policy. The comment cannot be longer than 128 characters.
	Comment pulumi.StringPtrInput
	// A configuration for a set of HTTP response headers that are used for Cross-Origin Resource Sharing (CORS). See Cors Config for more information.
	CorsConfig ResponseHeadersPolicyCorsConfigPtrInput
	// Object that contains an attribute `items` that contains a list of custom headers. See Custom Header for more information.
	CustomHeadersConfig ResponseHeadersPolicyCustomHeadersConfigPtrInput
	// The current version of the response headers policy.
	Etag pulumi.StringPtrInput
	// A unique name to identify the response headers policy.
	Name pulumi.StringPtrInput
	// A configuration for a set of security-related HTTP response headers. See Security Headers Config for more information.
	SecurityHeadersConfig ResponseHeadersPolicySecurityHeadersConfigPtrInput
}

func (ResponseHeadersPolicyState) ElementType() reflect.Type {
	return reflect.TypeOf((*responseHeadersPolicyState)(nil)).Elem()
}

type responseHeadersPolicyArgs struct {
	// A comment to describe the response headers policy. The comment cannot be longer than 128 characters.
	Comment *string `pulumi:"comment"`
	// A configuration for a set of HTTP response headers that are used for Cross-Origin Resource Sharing (CORS). See Cors Config for more information.
	CorsConfig *ResponseHeadersPolicyCorsConfig `pulumi:"corsConfig"`
	// Object that contains an attribute `items` that contains a list of custom headers. See Custom Header for more information.
	CustomHeadersConfig *ResponseHeadersPolicyCustomHeadersConfig `pulumi:"customHeadersConfig"`
	// The current version of the response headers policy.
	Etag *string `pulumi:"etag"`
	// A unique name to identify the response headers policy.
	Name *string `pulumi:"name"`
	// A configuration for a set of security-related HTTP response headers. See Security Headers Config for more information.
	SecurityHeadersConfig *ResponseHeadersPolicySecurityHeadersConfig `pulumi:"securityHeadersConfig"`
}

// The set of arguments for constructing a ResponseHeadersPolicy resource.
type ResponseHeadersPolicyArgs struct {
	// A comment to describe the response headers policy. The comment cannot be longer than 128 characters.
	Comment pulumi.StringPtrInput
	// A configuration for a set of HTTP response headers that are used for Cross-Origin Resource Sharing (CORS). See Cors Config for more information.
	CorsConfig ResponseHeadersPolicyCorsConfigPtrInput
	// Object that contains an attribute `items` that contains a list of custom headers. See Custom Header for more information.
	CustomHeadersConfig ResponseHeadersPolicyCustomHeadersConfigPtrInput
	// The current version of the response headers policy.
	Etag pulumi.StringPtrInput
	// A unique name to identify the response headers policy.
	Name pulumi.StringPtrInput
	// A configuration for a set of security-related HTTP response headers. See Security Headers Config for more information.
	SecurityHeadersConfig ResponseHeadersPolicySecurityHeadersConfigPtrInput
}

func (ResponseHeadersPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*responseHeadersPolicyArgs)(nil)).Elem()
}

type ResponseHeadersPolicyInput interface {
	pulumi.Input

	ToResponseHeadersPolicyOutput() ResponseHeadersPolicyOutput
	ToResponseHeadersPolicyOutputWithContext(ctx context.Context) ResponseHeadersPolicyOutput
}

func (*ResponseHeadersPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((**ResponseHeadersPolicy)(nil)).Elem()
}

func (i *ResponseHeadersPolicy) ToResponseHeadersPolicyOutput() ResponseHeadersPolicyOutput {
	return i.ToResponseHeadersPolicyOutputWithContext(context.Background())
}

func (i *ResponseHeadersPolicy) ToResponseHeadersPolicyOutputWithContext(ctx context.Context) ResponseHeadersPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResponseHeadersPolicyOutput)
}

// ResponseHeadersPolicyArrayInput is an input type that accepts ResponseHeadersPolicyArray and ResponseHeadersPolicyArrayOutput values.
// You can construct a concrete instance of `ResponseHeadersPolicyArrayInput` via:
//
//          ResponseHeadersPolicyArray{ ResponseHeadersPolicyArgs{...} }
type ResponseHeadersPolicyArrayInput interface {
	pulumi.Input

	ToResponseHeadersPolicyArrayOutput() ResponseHeadersPolicyArrayOutput
	ToResponseHeadersPolicyArrayOutputWithContext(context.Context) ResponseHeadersPolicyArrayOutput
}

type ResponseHeadersPolicyArray []ResponseHeadersPolicyInput

func (ResponseHeadersPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResponseHeadersPolicy)(nil)).Elem()
}

func (i ResponseHeadersPolicyArray) ToResponseHeadersPolicyArrayOutput() ResponseHeadersPolicyArrayOutput {
	return i.ToResponseHeadersPolicyArrayOutputWithContext(context.Background())
}

func (i ResponseHeadersPolicyArray) ToResponseHeadersPolicyArrayOutputWithContext(ctx context.Context) ResponseHeadersPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResponseHeadersPolicyArrayOutput)
}

// ResponseHeadersPolicyMapInput is an input type that accepts ResponseHeadersPolicyMap and ResponseHeadersPolicyMapOutput values.
// You can construct a concrete instance of `ResponseHeadersPolicyMapInput` via:
//
//          ResponseHeadersPolicyMap{ "key": ResponseHeadersPolicyArgs{...} }
type ResponseHeadersPolicyMapInput interface {
	pulumi.Input

	ToResponseHeadersPolicyMapOutput() ResponseHeadersPolicyMapOutput
	ToResponseHeadersPolicyMapOutputWithContext(context.Context) ResponseHeadersPolicyMapOutput
}

type ResponseHeadersPolicyMap map[string]ResponseHeadersPolicyInput

func (ResponseHeadersPolicyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResponseHeadersPolicy)(nil)).Elem()
}

func (i ResponseHeadersPolicyMap) ToResponseHeadersPolicyMapOutput() ResponseHeadersPolicyMapOutput {
	return i.ToResponseHeadersPolicyMapOutputWithContext(context.Background())
}

func (i ResponseHeadersPolicyMap) ToResponseHeadersPolicyMapOutputWithContext(ctx context.Context) ResponseHeadersPolicyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResponseHeadersPolicyMapOutput)
}

type ResponseHeadersPolicyOutput struct{ *pulumi.OutputState }

func (ResponseHeadersPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResponseHeadersPolicy)(nil)).Elem()
}

func (o ResponseHeadersPolicyOutput) ToResponseHeadersPolicyOutput() ResponseHeadersPolicyOutput {
	return o
}

func (o ResponseHeadersPolicyOutput) ToResponseHeadersPolicyOutputWithContext(ctx context.Context) ResponseHeadersPolicyOutput {
	return o
}

type ResponseHeadersPolicyArrayOutput struct{ *pulumi.OutputState }

func (ResponseHeadersPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ResponseHeadersPolicy)(nil)).Elem()
}

func (o ResponseHeadersPolicyArrayOutput) ToResponseHeadersPolicyArrayOutput() ResponseHeadersPolicyArrayOutput {
	return o
}

func (o ResponseHeadersPolicyArrayOutput) ToResponseHeadersPolicyArrayOutputWithContext(ctx context.Context) ResponseHeadersPolicyArrayOutput {
	return o
}

func (o ResponseHeadersPolicyArrayOutput) Index(i pulumi.IntInput) ResponseHeadersPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ResponseHeadersPolicy {
		return vs[0].([]*ResponseHeadersPolicy)[vs[1].(int)]
	}).(ResponseHeadersPolicyOutput)
}

type ResponseHeadersPolicyMapOutput struct{ *pulumi.OutputState }

func (ResponseHeadersPolicyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ResponseHeadersPolicy)(nil)).Elem()
}

func (o ResponseHeadersPolicyMapOutput) ToResponseHeadersPolicyMapOutput() ResponseHeadersPolicyMapOutput {
	return o
}

func (o ResponseHeadersPolicyMapOutput) ToResponseHeadersPolicyMapOutputWithContext(ctx context.Context) ResponseHeadersPolicyMapOutput {
	return o
}

func (o ResponseHeadersPolicyMapOutput) MapIndex(k pulumi.StringInput) ResponseHeadersPolicyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ResponseHeadersPolicy {
		return vs[0].(map[string]*ResponseHeadersPolicy)[vs[1].(string)]
	}).(ResponseHeadersPolicyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ResponseHeadersPolicyInput)(nil)).Elem(), &ResponseHeadersPolicy{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResponseHeadersPolicyArrayInput)(nil)).Elem(), ResponseHeadersPolicyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ResponseHeadersPolicyMapInput)(nil)).Elem(), ResponseHeadersPolicyMap{})
	pulumi.RegisterOutputType(ResponseHeadersPolicyOutput{})
	pulumi.RegisterOutputType(ResponseHeadersPolicyArrayOutput{})
	pulumi.RegisterOutputType(ResponseHeadersPolicyMapOutput{})
}
