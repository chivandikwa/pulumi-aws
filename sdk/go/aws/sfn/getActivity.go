// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package sfn

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a Step Functions Activity data source
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/sfn"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := sfn.LookupActivity(ctx, &sfn.LookupActivityArgs{
// 			Name: pulumi.StringRef("my-activity"),
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupActivity(ctx *pulumi.Context, args *LookupActivityArgs, opts ...pulumi.InvokeOption) (*LookupActivityResult, error) {
	var rv LookupActivityResult
	err := ctx.Invoke("aws:sfn/getActivity:getActivity", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getActivity.
type LookupActivityArgs struct {
	// The Amazon Resource Name (ARN) that identifies the activity.
	Arn *string `pulumi:"arn"`
	// The name that identifies the activity.
	Name *string `pulumi:"name"`
}

// A collection of values returned by getActivity.
type LookupActivityResult struct {
	Arn string `pulumi:"arn"`
	// The date the activity was created.
	CreationDate string `pulumi:"creationDate"`
	// The provider-assigned unique ID for this managed resource.
	Id   string `pulumi:"id"`
	Name string `pulumi:"name"`
}

func LookupActivityOutput(ctx *pulumi.Context, args LookupActivityOutputArgs, opts ...pulumi.InvokeOption) LookupActivityResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupActivityResult, error) {
			args := v.(LookupActivityArgs)
			r, err := LookupActivity(ctx, &args, opts...)
			return *r, err
		}).(LookupActivityResultOutput)
}

// A collection of arguments for invoking getActivity.
type LookupActivityOutputArgs struct {
	// The Amazon Resource Name (ARN) that identifies the activity.
	Arn pulumi.StringPtrInput `pulumi:"arn"`
	// The name that identifies the activity.
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupActivityOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupActivityArgs)(nil)).Elem()
}

// A collection of values returned by getActivity.
type LookupActivityResultOutput struct{ *pulumi.OutputState }

func (LookupActivityResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupActivityResult)(nil)).Elem()
}

func (o LookupActivityResultOutput) ToLookupActivityResultOutput() LookupActivityResultOutput {
	return o
}

func (o LookupActivityResultOutput) ToLookupActivityResultOutputWithContext(ctx context.Context) LookupActivityResultOutput {
	return o
}

func (o LookupActivityResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActivityResult) string { return v.Arn }).(pulumi.StringOutput)
}

// The date the activity was created.
func (o LookupActivityResultOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActivityResult) string { return v.CreationDate }).(pulumi.StringOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupActivityResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActivityResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupActivityResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupActivityResult) string { return v.Name }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupActivityResultOutput{})
}
