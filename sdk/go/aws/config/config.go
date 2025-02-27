// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
)

// The access key for API operations. You can retrieve this from the 'Security & Credentials' section of the AWS console.
func GetAccessKey(ctx *pulumi.Context) string {
	return config.Get(ctx, "aws:accessKey")
}
func GetAllowedAccountIds(ctx *pulumi.Context) string {
	return config.Get(ctx, "aws:allowedAccountIds")
}
func GetAssumeRole(ctx *pulumi.Context) string {
	return config.Get(ctx, "aws:assumeRole")
}

// Configuration block with settings to default resource tags across all resources.
func GetDefaultTags(ctx *pulumi.Context) string {
	return config.Get(ctx, "aws:defaultTags")
}

// Address of the EC2 metadata service endpoint to use. Can also be configured using the
// `AWS_EC2_METADATA_SERVICE_ENDPOINT` environment variable.
func GetEc2MetadataServiceEndpoint(ctx *pulumi.Context) string {
	return config.Get(ctx, "aws:ec2MetadataServiceEndpoint")
}

// Protocol to use with EC2 metadata service endpoint.Valid values are `IPv4` and `IPv6`. Can also be configured using the
// `AWS_EC2_METADATA_SERVICE_ENDPOINT_MODE` environment variable.
func GetEc2MetadataServiceEndpointMode(ctx *pulumi.Context) string {
	return config.Get(ctx, "aws:ec2MetadataServiceEndpointMode")
}
func GetEndpoints(ctx *pulumi.Context) string {
	return config.Get(ctx, "aws:endpoints")
}
func GetForbiddenAccountIds(ctx *pulumi.Context) string {
	return config.Get(ctx, "aws:forbiddenAccountIds")
}

// The address of an HTTP proxy to use when accessing the AWS API. Can also be configured using the `HTTP_PROXY` or
// `HTTPS_PROXY` environment variables.
func GetHttpProxy(ctx *pulumi.Context) string {
	return config.Get(ctx, "aws:httpProxy")
}

// Configuration block with settings to ignore resource tags across all resources.
func GetIgnoreTags(ctx *pulumi.Context) string {
	return config.Get(ctx, "aws:ignoreTags")
}

// Explicitly allow the provider to perform "insecure" SSL requests. If omitted, default value is `false`
func GetInsecure(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "aws:insecure")
}

// The maximum number of times an AWS API request is being executed. If the API request still fails, an error is thrown.
func GetMaxRetries(ctx *pulumi.Context) int {
	return config.GetInt(ctx, "aws:maxRetries")
}

// The profile for API operations. If not set, the default profile created with `aws configure` will be used.
func GetProfile(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "aws:profile")
	if err == nil {
		return v
	}
	return getEnvOrDefault("", nil, "AWS_PROFILE").(string)
}

// The region where AWS operations will take place. Examples are us-east-1, us-west-2, etc.
func GetRegion(ctx *pulumi.Context) string {
	v, err := config.Try(ctx, "aws:region")
	if err == nil {
		return v
	}
	return getEnvOrDefault("", nil, "AWS_REGION", "AWS_DEFAULT_REGION").(string)
}

// Set this to true to enable the request to use path-style addressing, i.e., https://s3.amazonaws.com/BUCKET/KEY. By
// default, the S3 client will use virtual hosted bucket addressing when possible (https://BUCKET.s3.amazonaws.com/KEY).
// Specific to the Amazon S3 service.
//
// Deprecated: Use s3_use_path_style instead.
func GetS3ForcePathStyle(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "aws:s3ForcePathStyle")
}

// Set this to true to enable the request to use path-style addressing, i.e., https://s3.amazonaws.com/BUCKET/KEY. By
// default, the S3 client will use virtual hosted bucket addressing when possible (https://BUCKET.s3.amazonaws.com/KEY).
// Specific to the Amazon S3 service.
func GetS3UsePathStyle(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "aws:s3UsePathStyle")
}

// The secret key for API operations. You can retrieve this from the 'Security & Credentials' section of the AWS console.
func GetSecretKey(ctx *pulumi.Context) string {
	return config.Get(ctx, "aws:secretKey")
}

// List of paths to shared config files. If not set, defaults to [~/.aws/config].
func GetSharedConfigFiles(ctx *pulumi.Context) string {
	return config.Get(ctx, "aws:sharedConfigFiles")
}

// The path to the shared credentials file. If not set, defaults to ~/.aws/credentials.
//
// Deprecated: Use shared_credentials_files instead.
func GetSharedCredentialsFile(ctx *pulumi.Context) string {
	return config.Get(ctx, "aws:sharedCredentialsFile")
}

// List of paths to shared credentials files. If not set, defaults to [~/.aws/credentials].
func GetSharedCredentialsFiles(ctx *pulumi.Context) string {
	return config.Get(ctx, "aws:sharedCredentialsFiles")
}

// Skip the credentials validation via STS API. Used for AWS API implementations that do not have STS
// available/implemented.
func GetSkipCredentialsValidation(ctx *pulumi.Context) bool {
	v, err := config.TryBool(ctx, "aws:skipCredentialsValidation")
	if err == nil {
		return v
	}
	return true
}

// Skip getting the supported EC2 platforms. Used by users that don't have ec2:DescribeAccountAttributes permissions.
func GetSkipGetEc2Platforms(ctx *pulumi.Context) bool {
	v, err := config.TryBool(ctx, "aws:skipGetEc2Platforms")
	if err == nil {
		return v
	}
	return true
}

// Skip the AWS Metadata API check. Used for AWS API implementations that do not have a metadata api endpoint.
func GetSkipMetadataApiCheck(ctx *pulumi.Context) bool {
	v, err := config.TryBool(ctx, "aws:skipMetadataApiCheck")
	if err == nil {
		return v
	}
	return true
}

// Skip static validation of region name. Used by users of alternative AWS-like APIs or users w/ access to regions that are
// not public (yet).
func GetSkipRegionValidation(ctx *pulumi.Context) bool {
	v, err := config.TryBool(ctx, "aws:skipRegionValidation")
	if err == nil {
		return v
	}
	return true
}

// Skip requesting the account ID. Used for AWS API implementations that do not have IAM/STS API and/or metadata API.
func GetSkipRequestingAccountId(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "aws:skipRequestingAccountId")
}

// session token. A session token is only required if you are using temporary security credentials.
func GetToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "aws:token")
}

// Resolve an endpoint with DualStack capability
func GetUseDualstackEndpoint(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "aws:useDualstackEndpoint")
}

// Resolve an endpoint with FIPS capability
func GetUseFipsEndpoint(ctx *pulumi.Context) bool {
	return config.GetBool(ctx, "aws:useFipsEndpoint")
}
