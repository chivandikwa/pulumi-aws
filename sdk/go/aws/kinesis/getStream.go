// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kinesis

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Use this data source to get information about a Kinesis Stream for use in other
// resources.
//
// For more details, see the [Amazon Kinesis Documentation](https://aws.amazon.com/documentation/kinesis/).
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/kinesis"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := kinesis.LookupStream(ctx, &kinesis.LookupStreamArgs{
// 			Name: "stream-name",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
func LookupStream(ctx *pulumi.Context, args *LookupStreamArgs, opts ...pulumi.InvokeOption) (*LookupStreamResult, error) {
	var rv LookupStreamResult
	err := ctx.Invoke("aws:kinesis/getStream:getStream", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getStream.
type LookupStreamArgs struct {
	// The name of the Kinesis Stream.
	Name string `pulumi:"name"`
	// A map of tags to assigned to the stream.
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getStream.
type LookupStreamResult struct {
	// The Amazon Resource Name (ARN) of the Kinesis Stream (same as id).
	Arn string `pulumi:"arn"`
	// The list of shard ids in the CLOSED state. See [Shard State](https://docs.aws.amazon.com/streams/latest/dev/kinesis-using-sdk-java-after-resharding.html#kinesis-using-sdk-java-resharding-data-routing) for more.
	ClosedShards []string `pulumi:"closedShards"`
	// The approximate UNIX timestamp that the stream was created.
	CreationTimestamp int `pulumi:"creationTimestamp"`
	// The provider-assigned unique ID for this managed resource.
	Id string `pulumi:"id"`
	// The name of the Kinesis Stream.
	Name string `pulumi:"name"`
	// The list of shard ids in the OPEN state. See [Shard State](https://docs.aws.amazon.com/streams/latest/dev/kinesis-using-sdk-java-after-resharding.html#kinesis-using-sdk-java-resharding-data-routing) for more.
	OpenShards []string `pulumi:"openShards"`
	// Length of time (in hours) data records are accessible after they are added to the stream.
	RetentionPeriod int `pulumi:"retentionPeriod"`
	// A list of shard-level CloudWatch metrics which are enabled for the stream. See [Monitoring with CloudWatch](https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html) for more.
	ShardLevelMetrics []string `pulumi:"shardLevelMetrics"`
	// The current status of the stream. The stream status is one of CREATING, DELETING, ACTIVE, or UPDATING.
	Status string `pulumi:"status"`
	// Indicates the [capacity mode](https://docs.aws.amazon.com/streams/latest/dev/how-do-i-size-a-stream.html) of the data stream. Detailed below.
	StreamModeDetails []GetStreamStreamModeDetail `pulumi:"streamModeDetails"`
	// A map of tags to assigned to the stream.
	Tags map[string]string `pulumi:"tags"`
}

func LookupStreamOutput(ctx *pulumi.Context, args LookupStreamOutputArgs, opts ...pulumi.InvokeOption) LookupStreamResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupStreamResult, error) {
			args := v.(LookupStreamArgs)
			r, err := LookupStream(ctx, &args, opts...)
			return *r, err
		}).(LookupStreamResultOutput)
}

// A collection of arguments for invoking getStream.
type LookupStreamOutputArgs struct {
	// The name of the Kinesis Stream.
	Name pulumi.StringInput `pulumi:"name"`
	// A map of tags to assigned to the stream.
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (LookupStreamOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStreamArgs)(nil)).Elem()
}

// A collection of values returned by getStream.
type LookupStreamResultOutput struct{ *pulumi.OutputState }

func (LookupStreamResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupStreamResult)(nil)).Elem()
}

func (o LookupStreamResultOutput) ToLookupStreamResultOutput() LookupStreamResultOutput {
	return o
}

func (o LookupStreamResultOutput) ToLookupStreamResultOutputWithContext(ctx context.Context) LookupStreamResultOutput {
	return o
}

// The Amazon Resource Name (ARN) of the Kinesis Stream (same as id).
func (o LookupStreamResultOutput) Arn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamResult) string { return v.Arn }).(pulumi.StringOutput)
}

// The list of shard ids in the CLOSED state. See [Shard State](https://docs.aws.amazon.com/streams/latest/dev/kinesis-using-sdk-java-after-resharding.html#kinesis-using-sdk-java-resharding-data-routing) for more.
func (o LookupStreamResultOutput) ClosedShards() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupStreamResult) []string { return v.ClosedShards }).(pulumi.StringArrayOutput)
}

// The approximate UNIX timestamp that the stream was created.
func (o LookupStreamResultOutput) CreationTimestamp() pulumi.IntOutput {
	return o.ApplyT(func(v LookupStreamResult) int { return v.CreationTimestamp }).(pulumi.IntOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o LookupStreamResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamResult) string { return v.Id }).(pulumi.StringOutput)
}

// The name of the Kinesis Stream.
func (o LookupStreamResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamResult) string { return v.Name }).(pulumi.StringOutput)
}

// The list of shard ids in the OPEN state. See [Shard State](https://docs.aws.amazon.com/streams/latest/dev/kinesis-using-sdk-java-after-resharding.html#kinesis-using-sdk-java-resharding-data-routing) for more.
func (o LookupStreamResultOutput) OpenShards() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupStreamResult) []string { return v.OpenShards }).(pulumi.StringArrayOutput)
}

// Length of time (in hours) data records are accessible after they are added to the stream.
func (o LookupStreamResultOutput) RetentionPeriod() pulumi.IntOutput {
	return o.ApplyT(func(v LookupStreamResult) int { return v.RetentionPeriod }).(pulumi.IntOutput)
}

// A list of shard-level CloudWatch metrics which are enabled for the stream. See [Monitoring with CloudWatch](https://docs.aws.amazon.com/streams/latest/dev/monitoring-with-cloudwatch.html) for more.
func (o LookupStreamResultOutput) ShardLevelMetrics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupStreamResult) []string { return v.ShardLevelMetrics }).(pulumi.StringArrayOutput)
}

// The current status of the stream. The stream status is one of CREATING, DELETING, ACTIVE, or UPDATING.
func (o LookupStreamResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupStreamResult) string { return v.Status }).(pulumi.StringOutput)
}

// Indicates the [capacity mode](https://docs.aws.amazon.com/streams/latest/dev/how-do-i-size-a-stream.html) of the data stream. Detailed below.
func (o LookupStreamResultOutput) StreamModeDetails() GetStreamStreamModeDetailArrayOutput {
	return o.ApplyT(func(v LookupStreamResult) []GetStreamStreamModeDetail { return v.StreamModeDetails }).(GetStreamStreamModeDetailArrayOutput)
}

// A map of tags to assigned to the stream.
func (o LookupStreamResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupStreamResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupStreamResultOutput{})
}
