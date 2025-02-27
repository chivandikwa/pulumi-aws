// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package cloudtrail

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Provides a CloudTrail resource.
//
// > **Tip:** For a multi-region trail, this resource must be in the home region of the trail.
//
// > **Tip:** For an organization trail, this resource must be in the master account of the organization.
//
// ## Example Usage
// ### Basic
//
// Enable CloudTrail to capture all compatible management events in region.
// For capturing events from services like IAM, `includeGlobalServiceEvents` must be enabled.
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws"
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/cloudtrail"
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		current, err := aws.GetCallerIdentity(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		bucketV2, err := s3.NewBucketV2(ctx, "bucketV2", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = s3.NewBucketPolicy(ctx, "bucketPolicy", &s3.BucketPolicyArgs{
// 			Bucket: bucketV2.ID(),
// 			Policy: pulumi.All(bucketV2.ID(), bucketV2.ID()).ApplyT(func(_args []interface{}) (string, error) {
// 				bucketV2Id := _args[0].(string)
// 				bucketV2Id1 := _args[1].(string)
// 				return fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "  {\n", "      \"Version\": \"2012-10-17\",\n", "      \"Statement\": [\n", "          {\n", "              \"Sid\": \"AWSCloudTrailAclCheck\",\n", "              \"Effect\": \"Allow\",\n", "              \"Principal\": {\n", "                \"Service\": \"cloudtrail.amazonaws.com\"\n", "              },\n", "              \"Action\": \"s3:GetBucketAcl\",\n", "              \"Resource\": \"arn:aws:s3:::", bucketV2Id, "\"\n", "          },\n", "          {\n", "              \"Sid\": \"AWSCloudTrailWrite\",\n", "              \"Effect\": \"Allow\",\n", "              \"Principal\": {\n", "                \"Service\": \"cloudtrail.amazonaws.com\"\n", "              },\n", "              \"Action\": \"s3:PutObject\",\n", "              \"Resource\": \"arn:aws:s3:::", bucketV2Id1, "/prefix/AWSLogs/", current.AccountId, "/*\",\n", "              \"Condition\": {\n", "                  \"StringEquals\": {\n", "                      \"s3:x-amz-acl\": \"bucket-owner-full-control\"\n", "                  }\n", "              }\n", "          }\n", "      ]\n", "  }\n"), nil
// 			}).(pulumi.StringOutput),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudtrail.NewTrail(ctx, "foobar", &cloudtrail.TrailArgs{
// 			S3BucketName:               bucketV2.ID(),
// 			S3KeyPrefix:                pulumi.String("prefix"),
// 			IncludeGlobalServiceEvents: pulumi.Bool(false),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### Data Event Logging
//
// CloudTrail can log [Data Events](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/logging-data-events-with-cloudtrail.html) for certain services such as S3 objects and Lambda function invocations. Additional information about data event configuration can be found in the following links:
//
// * [CloudTrail API DataResource documentation](https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_DataResource.html) (for basic event selector).
// * [CloudTrail API AdvancedFieldSelector documentation](https://docs.aws.amazon.com/awscloudtrail/latest/APIReference/API_AdvancedFieldSelector.html) (for advanced event selector).
// ### Logging All Lambda Function Invocations By Using Basic Event Selectors
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/cloudtrail"
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		bucketV2, err := s3.NewBucketV2(ctx, "bucketV2", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudtrail.NewTrail(ctx, "example", &cloudtrail.TrailArgs{
// 			S3BucketName: bucketV2.ID(),
// 			S3KeyPrefix:  pulumi.String("prefix"),
// 			EventSelectors: cloudtrail.TrailEventSelectorArray{
// 				&cloudtrail.TrailEventSelectorArgs{
// 					ReadWriteType:           pulumi.String("All"),
// 					IncludeManagementEvents: pulumi.Bool(true),
// 					DataResources: cloudtrail.TrailEventSelectorDataResourceArray{
// 						&cloudtrail.TrailEventSelectorDataResourceArgs{
// 							Type: pulumi.String("AWS::Lambda::Function"),
// 							Values: pulumi.StringArray{
// 								pulumi.String("arn:aws:lambda"),
// 							},
// 						},
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
// ### Logging All S3 Object Events By Using Basic Event Selectors
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/cloudtrail"
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		bucketV2, err := s3.NewBucketV2(ctx, "bucketV2", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudtrail.NewTrail(ctx, "example", &cloudtrail.TrailArgs{
// 			S3BucketName: bucketV2.ID(),
// 			S3KeyPrefix:  pulumi.String("prefix"),
// 			EventSelectors: cloudtrail.TrailEventSelectorArray{
// 				&cloudtrail.TrailEventSelectorArgs{
// 					ReadWriteType:           pulumi.String("All"),
// 					IncludeManagementEvents: pulumi.Bool(true),
// 					DataResources: cloudtrail.TrailEventSelectorDataResourceArray{
// 						&cloudtrail.TrailEventSelectorDataResourceArgs{
// 							Type: pulumi.String("AWS::S3::Object"),
// 							Values: pulumi.StringArray{
// 								pulumi.String("arn:aws:s3:::"),
// 							},
// 						},
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
// ### Logging Individual S3 Bucket Events By Using Basic Event Selectors
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/cloudtrail"
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		important_bucket, err := s3.LookupBucket(ctx, &s3.LookupBucketArgs{
// 			Bucket: "important-bucket",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudtrail.NewTrail(ctx, "example", &cloudtrail.TrailArgs{
// 			S3BucketName: pulumi.String(important_bucket.Id),
// 			S3KeyPrefix:  pulumi.String("prefix"),
// 			EventSelectors: cloudtrail.TrailEventSelectorArray{
// 				&cloudtrail.TrailEventSelectorArgs{
// 					ReadWriteType:           pulumi.String("All"),
// 					IncludeManagementEvents: pulumi.Bool(true),
// 					DataResources: cloudtrail.TrailEventSelectorDataResourceArray{
// 						&cloudtrail.TrailEventSelectorDataResourceArgs{
// 							Type: pulumi.String("AWS::S3::Object"),
// 							Values: pulumi.StringArray{
// 								pulumi.String(fmt.Sprintf("%v%v", important_bucket.Arn, "/")),
// 							},
// 						},
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
// ### Logging All S3 Object Events Except For Two S3 Buckets By Using Advanced Event Selectors
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/cloudtrail"
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		not_important_bucket_1, err := s3.LookupBucket(ctx, &s3.LookupBucketArgs{
// 			Bucket: "not-important-bucket-1",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		not_important_bucket_2, err := s3.LookupBucket(ctx, &s3.LookupBucketArgs{
// 			Bucket: "not-important-bucket-2",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudtrail.NewTrail(ctx, "example", &cloudtrail.TrailArgs{
// 			AdvancedEventSelectors: cloudtrail.TrailAdvancedEventSelectorArray{
// 				&cloudtrail.TrailAdvancedEventSelectorArgs{
// 					FieldSelectors: cloudtrail.TrailAdvancedEventSelectorFieldSelectorArray{
// 						&cloudtrail.TrailAdvancedEventSelectorFieldSelectorArgs{
// 							Equals: pulumi.StringArray{
// 								pulumi.String("Data"),
// 							},
// 							Field: pulumi.String("eventCategory"),
// 						},
// 						&cloudtrail.TrailAdvancedEventSelectorFieldSelectorArgs{
// 							Field: pulumi.String("resources.ARN"),
// 							NotEquals: pulumi.StringArray{
// 								pulumi.String(fmt.Sprintf("%v%v", not_important_bucket_1.Arn, "/")),
// 								pulumi.String(fmt.Sprintf("%v%v", not_important_bucket_2.Arn, "/")),
// 							},
// 						},
// 						&cloudtrail.TrailAdvancedEventSelectorFieldSelectorArgs{
// 							Equals: pulumi.StringArray{
// 								pulumi.String("AWS::S3::Object"),
// 							},
// 							Field: pulumi.String("resources.type"),
// 						},
// 					},
// 					Name: pulumi.String("Log all S3 objects events except for two S3 buckets"),
// 				},
// 				&cloudtrail.TrailAdvancedEventSelectorArgs{
// 					FieldSelectors: cloudtrail.TrailAdvancedEventSelectorFieldSelectorArray{
// 						&cloudtrail.TrailAdvancedEventSelectorFieldSelectorArgs{
// 							Equals: pulumi.StringArray{
// 								pulumi.String("Management"),
// 							},
// 							Field: pulumi.String("eventCategory"),
// 						},
// 					},
// 					Name: pulumi.String("Log readOnly and writeOnly management events"),
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
// ### Logging Individual S3 Buckets And Specific Event Names By Using Advanced Event Selectors
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/cloudtrail"
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		important_bucket_1, err := s3.LookupBucket(ctx, &s3.LookupBucketArgs{
// 			Bucket: "important-bucket-1",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		important_bucket_2, err := s3.LookupBucket(ctx, &s3.LookupBucketArgs{
// 			Bucket: "important-bucket-2",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		important_bucket_3, err := s3.LookupBucket(ctx, &s3.LookupBucketArgs{
// 			Bucket: "important-bucket-3",
// 		}, nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudtrail.NewTrail(ctx, "example", &cloudtrail.TrailArgs{
// 			AdvancedEventSelectors: cloudtrail.TrailAdvancedEventSelectorArray{
// 				&cloudtrail.TrailAdvancedEventSelectorArgs{
// 					FieldSelectors: cloudtrail.TrailAdvancedEventSelectorFieldSelectorArray{
// 						&cloudtrail.TrailAdvancedEventSelectorFieldSelectorArgs{
// 							Equals: pulumi.StringArray{
// 								pulumi.String("Data"),
// 							},
// 							Field: pulumi.String("eventCategory"),
// 						},
// 						&cloudtrail.TrailAdvancedEventSelectorFieldSelectorArgs{
// 							Equals: pulumi.StringArray{
// 								pulumi.String("PutObject"),
// 								pulumi.String("DeleteObject"),
// 							},
// 							Field: pulumi.String("eventName"),
// 						},
// 						&cloudtrail.TrailAdvancedEventSelectorFieldSelectorArgs{
// 							Equals: pulumi.StringArray{
// 								pulumi.String(fmt.Sprintf("%v%v", important_bucket_1.Arn, "/")),
// 								pulumi.String(fmt.Sprintf("%v%v", important_bucket_2.Arn, "/")),
// 							},
// 							Field: pulumi.String("resources.ARN"),
// 						},
// 						&cloudtrail.TrailAdvancedEventSelectorFieldSelectorArgs{
// 							Equals: pulumi.StringArray{
// 								pulumi.String("false"),
// 							},
// 							Field: pulumi.String("readOnly"),
// 						},
// 						&cloudtrail.TrailAdvancedEventSelectorFieldSelectorArgs{
// 							Equals: pulumi.StringArray{
// 								pulumi.String("AWS::S3::Object"),
// 							},
// 							Field: pulumi.String("resources.type"),
// 						},
// 					},
// 					Name: pulumi.String("Log PutObject and DeleteObject events for two S3 buckets"),
// 				},
// 				&cloudtrail.TrailAdvancedEventSelectorArgs{
// 					FieldSelectors: cloudtrail.TrailAdvancedEventSelectorFieldSelectorArray{
// 						&cloudtrail.TrailAdvancedEventSelectorFieldSelectorArgs{
// 							Equals: pulumi.StringArray{
// 								pulumi.String("Data"),
// 							},
// 							Field: pulumi.String("eventCategory"),
// 						},
// 						&cloudtrail.TrailAdvancedEventSelectorFieldSelectorArgs{
// 							Field: pulumi.String("eventName"),
// 							StartsWith: []string{
// 								"Delete",
// 							},
// 						},
// 						&cloudtrail.TrailAdvancedEventSelectorFieldSelectorArgs{
// 							Equals: pulumi.StringArray{
// 								pulumi.String(fmt.Sprintf("%v%v", important_bucket_3.Arn, "/important-prefix")),
// 							},
// 							Field: pulumi.String("resources.ARN"),
// 						},
// 						&cloudtrail.TrailAdvancedEventSelectorFieldSelectorArgs{
// 							Equals: pulumi.StringArray{
// 								pulumi.String("false"),
// 							},
// 							Field: pulumi.String("readOnly"),
// 						},
// 						&cloudtrail.TrailAdvancedEventSelectorFieldSelectorArgs{
// 							Equals: pulumi.StringArray{
// 								pulumi.String("AWS::S3::Object"),
// 							},
// 							Field: pulumi.String("resources.type"),
// 						},
// 					},
// 					Name: pulumi.String("Log Delete* events for one S3 bucket"),
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
// ### Sending Events to CloudWatch Logs
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws"
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/cloudtrail"
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/cloudwatch"
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		current, err := aws.GetPartition(ctx, nil, nil)
// 		if err != nil {
// 			return err
// 		}
// 		exampleLogGroup, err := cloudwatch.NewLogGroup(ctx, "exampleLogGroup", nil)
// 		if err != nil {
// 			return err
// 		}
// 		testRole, err := iam.NewRole(ctx, "testRole", &iam.RoleArgs{
// 			AssumeRolePolicy: pulumi.Any(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Sid\": \"\",\n", "      \"Effect\": \"Allow\",\n", "      \"Principal\": {\n", "        \"Service\": \"cloudtrail.", current.DnsSuffix, "\"\n", "      },\n", "      \"Action\": \"sts:AssumeRole\"\n", "    }\n", "  ]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = iam.NewRolePolicy(ctx, "testRolePolicy", &iam.RolePolicyArgs{
// 			Role:   testRole.ID(),
// 			Policy: pulumi.Any(fmt.Sprintf("%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v%v", "{\n", "  \"Version\": \"2012-10-17\",\n", "  \"Statement\": [\n", "    {\n", "      \"Sid\": \"AWSCloudTrailCreateLogStream\",\n", "      \"Effect\": \"Allow\",\n", "      \"Action\": [\n", "        \"logs:CreateLogStream\",\n", "        \"logs:PutLogEvents\"\n", "      ],\n", "      \"Resource\": \"", aws_cloudwatch_log_group.Test.Arn, ":*\"\n", "    }\n", "  ]\n", "}\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = s3.NewBucketV2(ctx, "bucketV2", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = cloudtrail.NewTrail(ctx, "exampleTrail", &cloudtrail.TrailArgs{
// 			S3BucketName:          pulumi.Any(data.Aws_s3_bucket.Important - bucket.Id),
// 			S3KeyPrefix:           pulumi.String("prefix"),
// 			CloudWatchLogsRoleArn: testRole.Arn,
// 			CloudWatchLogsGroupArn: exampleLogGroup.Arn.ApplyT(func(arn string) (string, error) {
// 				return fmt.Sprintf("%v%v", arn, ":*"), nil
// 			}).(pulumi.StringOutput),
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
// Cloudtrails can be imported using the `name`, e.g.,
//
// ```sh
//  $ pulumi import aws:cloudtrail/trail:Trail sample my-sample-trail
// ```
type Trail struct {
	pulumi.CustomResourceState

	// Specifies an advanced event selector for enabling data event logging. Fields documented below. Conflicts with `eventSelector`.
	AdvancedEventSelectors TrailAdvancedEventSelectorArrayOutput `pulumi:"advancedEventSelectors"`
	// ARN of the trail.
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Log group name using an ARN that represents the log group to which CloudTrail logs will be delivered. Note that CloudTrail requires the Log Stream wildcard.
	CloudWatchLogsGroupArn pulumi.StringPtrOutput `pulumi:"cloudWatchLogsGroupArn"`
	// Role for the CloudWatch Logs endpoint to assume to write to a user’s log group.
	CloudWatchLogsRoleArn pulumi.StringPtrOutput `pulumi:"cloudWatchLogsRoleArn"`
	// Whether log file integrity validation is enabled. Defaults to `false`.
	EnableLogFileValidation pulumi.BoolPtrOutput `pulumi:"enableLogFileValidation"`
	// Enables logging for the trail. Defaults to `true`. Setting this to `false` will pause logging.
	EnableLogging pulumi.BoolPtrOutput `pulumi:"enableLogging"`
	// Specifies an event selector for enabling data event logging. Fields documented below. Please note the [CloudTrail limits](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/WhatIsCloudTrail-Limits.html) when configuring these. Conflicts with `advancedEventSelector`.
	EventSelectors TrailEventSelectorArrayOutput `pulumi:"eventSelectors"`
	// Region in which the trail was created.
	HomeRegion pulumi.StringOutput `pulumi:"homeRegion"`
	// Whether the trail is publishing events from global services such as IAM to the log files. Defaults to `true`.
	IncludeGlobalServiceEvents pulumi.BoolPtrOutput `pulumi:"includeGlobalServiceEvents"`
	// Configuration block for identifying unusual operational activity. See details below.
	InsightSelectors TrailInsightSelectorArrayOutput `pulumi:"insightSelectors"`
	// Whether the trail is created in the current region or in all regions. Defaults to `false`.
	IsMultiRegionTrail pulumi.BoolPtrOutput `pulumi:"isMultiRegionTrail"`
	// Whether the trail is an AWS Organizations trail. Organization trails log events for the master account and all member accounts. Can only be created in the organization master account. Defaults to `false`.
	IsOrganizationTrail pulumi.BoolPtrOutput `pulumi:"isOrganizationTrail"`
	// KMS key ARN to use to encrypt the logs delivered by CloudTrail.
	KmsKeyId pulumi.StringPtrOutput `pulumi:"kmsKeyId"`
	// Specifies the name of the advanced event selector.
	Name pulumi.StringOutput `pulumi:"name"`
	// Name of the S3 bucket designated for publishing log files.
	S3BucketName pulumi.StringOutput `pulumi:"s3BucketName"`
	// S3 key prefix that follows the name of the bucket you have designated for log file delivery.
	S3KeyPrefix pulumi.StringPtrOutput `pulumi:"s3KeyPrefix"`
	// Name of the Amazon SNS topic defined for notification of log file delivery.
	SnsTopicName pulumi.StringPtrOutput `pulumi:"snsTopicName"`
	// Map of tags to assign to the trail. If configured with provider defaultTags present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider.
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
}

// NewTrail registers a new resource with the given unique name, arguments, and options.
func NewTrail(ctx *pulumi.Context,
	name string, args *TrailArgs, opts ...pulumi.ResourceOption) (*Trail, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.S3BucketName == nil {
		return nil, errors.New("invalid value for required argument 'S3BucketName'")
	}
	var resource Trail
	err := ctx.RegisterResource("aws:cloudtrail/trail:Trail", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTrail gets an existing Trail resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTrail(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TrailState, opts ...pulumi.ResourceOption) (*Trail, error) {
	var resource Trail
	err := ctx.ReadResource("aws:cloudtrail/trail:Trail", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Trail resources.
type trailState struct {
	// Specifies an advanced event selector for enabling data event logging. Fields documented below. Conflicts with `eventSelector`.
	AdvancedEventSelectors []TrailAdvancedEventSelector `pulumi:"advancedEventSelectors"`
	// ARN of the trail.
	Arn *string `pulumi:"arn"`
	// Log group name using an ARN that represents the log group to which CloudTrail logs will be delivered. Note that CloudTrail requires the Log Stream wildcard.
	CloudWatchLogsGroupArn *string `pulumi:"cloudWatchLogsGroupArn"`
	// Role for the CloudWatch Logs endpoint to assume to write to a user’s log group.
	CloudWatchLogsRoleArn *string `pulumi:"cloudWatchLogsRoleArn"`
	// Whether log file integrity validation is enabled. Defaults to `false`.
	EnableLogFileValidation *bool `pulumi:"enableLogFileValidation"`
	// Enables logging for the trail. Defaults to `true`. Setting this to `false` will pause logging.
	EnableLogging *bool `pulumi:"enableLogging"`
	// Specifies an event selector for enabling data event logging. Fields documented below. Please note the [CloudTrail limits](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/WhatIsCloudTrail-Limits.html) when configuring these. Conflicts with `advancedEventSelector`.
	EventSelectors []TrailEventSelector `pulumi:"eventSelectors"`
	// Region in which the trail was created.
	HomeRegion *string `pulumi:"homeRegion"`
	// Whether the trail is publishing events from global services such as IAM to the log files. Defaults to `true`.
	IncludeGlobalServiceEvents *bool `pulumi:"includeGlobalServiceEvents"`
	// Configuration block for identifying unusual operational activity. See details below.
	InsightSelectors []TrailInsightSelector `pulumi:"insightSelectors"`
	// Whether the trail is created in the current region or in all regions. Defaults to `false`.
	IsMultiRegionTrail *bool `pulumi:"isMultiRegionTrail"`
	// Whether the trail is an AWS Organizations trail. Organization trails log events for the master account and all member accounts. Can only be created in the organization master account. Defaults to `false`.
	IsOrganizationTrail *bool `pulumi:"isOrganizationTrail"`
	// KMS key ARN to use to encrypt the logs delivered by CloudTrail.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// Specifies the name of the advanced event selector.
	Name *string `pulumi:"name"`
	// Name of the S3 bucket designated for publishing log files.
	S3BucketName *string `pulumi:"s3BucketName"`
	// S3 key prefix that follows the name of the bucket you have designated for log file delivery.
	S3KeyPrefix *string `pulumi:"s3KeyPrefix"`
	// Name of the Amazon SNS topic defined for notification of log file delivery.
	SnsTopicName *string `pulumi:"snsTopicName"`
	// Map of tags to assign to the trail. If configured with provider defaultTags present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Map of tags assigned to the resource, including those inherited from the provider.
	TagsAll map[string]string `pulumi:"tagsAll"`
}

type TrailState struct {
	// Specifies an advanced event selector for enabling data event logging. Fields documented below. Conflicts with `eventSelector`.
	AdvancedEventSelectors TrailAdvancedEventSelectorArrayInput
	// ARN of the trail.
	Arn pulumi.StringPtrInput
	// Log group name using an ARN that represents the log group to which CloudTrail logs will be delivered. Note that CloudTrail requires the Log Stream wildcard.
	CloudWatchLogsGroupArn pulumi.StringPtrInput
	// Role for the CloudWatch Logs endpoint to assume to write to a user’s log group.
	CloudWatchLogsRoleArn pulumi.StringPtrInput
	// Whether log file integrity validation is enabled. Defaults to `false`.
	EnableLogFileValidation pulumi.BoolPtrInput
	// Enables logging for the trail. Defaults to `true`. Setting this to `false` will pause logging.
	EnableLogging pulumi.BoolPtrInput
	// Specifies an event selector for enabling data event logging. Fields documented below. Please note the [CloudTrail limits](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/WhatIsCloudTrail-Limits.html) when configuring these. Conflicts with `advancedEventSelector`.
	EventSelectors TrailEventSelectorArrayInput
	// Region in which the trail was created.
	HomeRegion pulumi.StringPtrInput
	// Whether the trail is publishing events from global services such as IAM to the log files. Defaults to `true`.
	IncludeGlobalServiceEvents pulumi.BoolPtrInput
	// Configuration block for identifying unusual operational activity. See details below.
	InsightSelectors TrailInsightSelectorArrayInput
	// Whether the trail is created in the current region or in all regions. Defaults to `false`.
	IsMultiRegionTrail pulumi.BoolPtrInput
	// Whether the trail is an AWS Organizations trail. Organization trails log events for the master account and all member accounts. Can only be created in the organization master account. Defaults to `false`.
	IsOrganizationTrail pulumi.BoolPtrInput
	// KMS key ARN to use to encrypt the logs delivered by CloudTrail.
	KmsKeyId pulumi.StringPtrInput
	// Specifies the name of the advanced event selector.
	Name pulumi.StringPtrInput
	// Name of the S3 bucket designated for publishing log files.
	S3BucketName pulumi.StringPtrInput
	// S3 key prefix that follows the name of the bucket you have designated for log file delivery.
	S3KeyPrefix pulumi.StringPtrInput
	// Name of the Amazon SNS topic defined for notification of log file delivery.
	SnsTopicName pulumi.StringPtrInput
	// Map of tags to assign to the trail. If configured with provider defaultTags present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Map of tags assigned to the resource, including those inherited from the provider.
	TagsAll pulumi.StringMapInput
}

func (TrailState) ElementType() reflect.Type {
	return reflect.TypeOf((*trailState)(nil)).Elem()
}

type trailArgs struct {
	// Specifies an advanced event selector for enabling data event logging. Fields documented below. Conflicts with `eventSelector`.
	AdvancedEventSelectors []TrailAdvancedEventSelector `pulumi:"advancedEventSelectors"`
	// Log group name using an ARN that represents the log group to which CloudTrail logs will be delivered. Note that CloudTrail requires the Log Stream wildcard.
	CloudWatchLogsGroupArn *string `pulumi:"cloudWatchLogsGroupArn"`
	// Role for the CloudWatch Logs endpoint to assume to write to a user’s log group.
	CloudWatchLogsRoleArn *string `pulumi:"cloudWatchLogsRoleArn"`
	// Whether log file integrity validation is enabled. Defaults to `false`.
	EnableLogFileValidation *bool `pulumi:"enableLogFileValidation"`
	// Enables logging for the trail. Defaults to `true`. Setting this to `false` will pause logging.
	EnableLogging *bool `pulumi:"enableLogging"`
	// Specifies an event selector for enabling data event logging. Fields documented below. Please note the [CloudTrail limits](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/WhatIsCloudTrail-Limits.html) when configuring these. Conflicts with `advancedEventSelector`.
	EventSelectors []TrailEventSelector `pulumi:"eventSelectors"`
	// Whether the trail is publishing events from global services such as IAM to the log files. Defaults to `true`.
	IncludeGlobalServiceEvents *bool `pulumi:"includeGlobalServiceEvents"`
	// Configuration block for identifying unusual operational activity. See details below.
	InsightSelectors []TrailInsightSelector `pulumi:"insightSelectors"`
	// Whether the trail is created in the current region or in all regions. Defaults to `false`.
	IsMultiRegionTrail *bool `pulumi:"isMultiRegionTrail"`
	// Whether the trail is an AWS Organizations trail. Organization trails log events for the master account and all member accounts. Can only be created in the organization master account. Defaults to `false`.
	IsOrganizationTrail *bool `pulumi:"isOrganizationTrail"`
	// KMS key ARN to use to encrypt the logs delivered by CloudTrail.
	KmsKeyId *string `pulumi:"kmsKeyId"`
	// Specifies the name of the advanced event selector.
	Name *string `pulumi:"name"`
	// Name of the S3 bucket designated for publishing log files.
	S3BucketName string `pulumi:"s3BucketName"`
	// S3 key prefix that follows the name of the bucket you have designated for log file delivery.
	S3KeyPrefix *string `pulumi:"s3KeyPrefix"`
	// Name of the Amazon SNS topic defined for notification of log file delivery.
	SnsTopicName *string `pulumi:"snsTopicName"`
	// Map of tags to assign to the trail. If configured with provider defaultTags present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
}

// The set of arguments for constructing a Trail resource.
type TrailArgs struct {
	// Specifies an advanced event selector for enabling data event logging. Fields documented below. Conflicts with `eventSelector`.
	AdvancedEventSelectors TrailAdvancedEventSelectorArrayInput
	// Log group name using an ARN that represents the log group to which CloudTrail logs will be delivered. Note that CloudTrail requires the Log Stream wildcard.
	CloudWatchLogsGroupArn pulumi.StringPtrInput
	// Role for the CloudWatch Logs endpoint to assume to write to a user’s log group.
	CloudWatchLogsRoleArn pulumi.StringPtrInput
	// Whether log file integrity validation is enabled. Defaults to `false`.
	EnableLogFileValidation pulumi.BoolPtrInput
	// Enables logging for the trail. Defaults to `true`. Setting this to `false` will pause logging.
	EnableLogging pulumi.BoolPtrInput
	// Specifies an event selector for enabling data event logging. Fields documented below. Please note the [CloudTrail limits](https://docs.aws.amazon.com/awscloudtrail/latest/userguide/WhatIsCloudTrail-Limits.html) when configuring these. Conflicts with `advancedEventSelector`.
	EventSelectors TrailEventSelectorArrayInput
	// Whether the trail is publishing events from global services such as IAM to the log files. Defaults to `true`.
	IncludeGlobalServiceEvents pulumi.BoolPtrInput
	// Configuration block for identifying unusual operational activity. See details below.
	InsightSelectors TrailInsightSelectorArrayInput
	// Whether the trail is created in the current region or in all regions. Defaults to `false`.
	IsMultiRegionTrail pulumi.BoolPtrInput
	// Whether the trail is an AWS Organizations trail. Organization trails log events for the master account and all member accounts. Can only be created in the organization master account. Defaults to `false`.
	IsOrganizationTrail pulumi.BoolPtrInput
	// KMS key ARN to use to encrypt the logs delivered by CloudTrail.
	KmsKeyId pulumi.StringPtrInput
	// Specifies the name of the advanced event selector.
	Name pulumi.StringPtrInput
	// Name of the S3 bucket designated for publishing log files.
	S3BucketName pulumi.StringInput
	// S3 key prefix that follows the name of the bucket you have designated for log file delivery.
	S3KeyPrefix pulumi.StringPtrInput
	// Name of the Amazon SNS topic defined for notification of log file delivery.
	SnsTopicName pulumi.StringPtrInput
	// Map of tags to assign to the trail. If configured with provider defaultTags present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
}

func (TrailArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*trailArgs)(nil)).Elem()
}

type TrailInput interface {
	pulumi.Input

	ToTrailOutput() TrailOutput
	ToTrailOutputWithContext(ctx context.Context) TrailOutput
}

func (*Trail) ElementType() reflect.Type {
	return reflect.TypeOf((**Trail)(nil)).Elem()
}

func (i *Trail) ToTrailOutput() TrailOutput {
	return i.ToTrailOutputWithContext(context.Background())
}

func (i *Trail) ToTrailOutputWithContext(ctx context.Context) TrailOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrailOutput)
}

// TrailArrayInput is an input type that accepts TrailArray and TrailArrayOutput values.
// You can construct a concrete instance of `TrailArrayInput` via:
//
//          TrailArray{ TrailArgs{...} }
type TrailArrayInput interface {
	pulumi.Input

	ToTrailArrayOutput() TrailArrayOutput
	ToTrailArrayOutputWithContext(context.Context) TrailArrayOutput
}

type TrailArray []TrailInput

func (TrailArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Trail)(nil)).Elem()
}

func (i TrailArray) ToTrailArrayOutput() TrailArrayOutput {
	return i.ToTrailArrayOutputWithContext(context.Background())
}

func (i TrailArray) ToTrailArrayOutputWithContext(ctx context.Context) TrailArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrailArrayOutput)
}

// TrailMapInput is an input type that accepts TrailMap and TrailMapOutput values.
// You can construct a concrete instance of `TrailMapInput` via:
//
//          TrailMap{ "key": TrailArgs{...} }
type TrailMapInput interface {
	pulumi.Input

	ToTrailMapOutput() TrailMapOutput
	ToTrailMapOutputWithContext(context.Context) TrailMapOutput
}

type TrailMap map[string]TrailInput

func (TrailMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Trail)(nil)).Elem()
}

func (i TrailMap) ToTrailMapOutput() TrailMapOutput {
	return i.ToTrailMapOutputWithContext(context.Background())
}

func (i TrailMap) ToTrailMapOutputWithContext(ctx context.Context) TrailMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrailMapOutput)
}

type TrailOutput struct{ *pulumi.OutputState }

func (TrailOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Trail)(nil)).Elem()
}

func (o TrailOutput) ToTrailOutput() TrailOutput {
	return o
}

func (o TrailOutput) ToTrailOutputWithContext(ctx context.Context) TrailOutput {
	return o
}

type TrailArrayOutput struct{ *pulumi.OutputState }

func (TrailArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Trail)(nil)).Elem()
}

func (o TrailArrayOutput) ToTrailArrayOutput() TrailArrayOutput {
	return o
}

func (o TrailArrayOutput) ToTrailArrayOutputWithContext(ctx context.Context) TrailArrayOutput {
	return o
}

func (o TrailArrayOutput) Index(i pulumi.IntInput) TrailOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Trail {
		return vs[0].([]*Trail)[vs[1].(int)]
	}).(TrailOutput)
}

type TrailMapOutput struct{ *pulumi.OutputState }

func (TrailMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Trail)(nil)).Elem()
}

func (o TrailMapOutput) ToTrailMapOutput() TrailMapOutput {
	return o
}

func (o TrailMapOutput) ToTrailMapOutputWithContext(ctx context.Context) TrailMapOutput {
	return o
}

func (o TrailMapOutput) MapIndex(k pulumi.StringInput) TrailOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Trail {
		return vs[0].(map[string]*Trail)[vs[1].(string)]
	}).(TrailOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TrailInput)(nil)).Elem(), &Trail{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrailArrayInput)(nil)).Elem(), TrailArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TrailMapInput)(nil)).Elem(), TrailMap{})
	pulumi.RegisterOutputType(TrailOutput{})
	pulumi.RegisterOutputType(TrailArrayOutput{})
	pulumi.RegisterOutputType(TrailMapOutput{})
}
