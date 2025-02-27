// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package mwaa

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a MWAA Environment resource.
//
// ## Example Usage
//
// A MWAA Environment requires an IAM role (`iam.Role`), two subnets in the private zone (`ec2.Subnet`) and a versioned S3 bucket (`s3.BucketV2`).
//
// ## Import
//
// MWAA Environment can be imported using `Name` e.g.,
//
// ```sh
//  $ pulumi import aws:mwaa/environment:Environment example MyAirflowEnvironment
// ```
type Environment struct {
	pulumi.CustomResourceState

	// The `airflowConfigurationOptions` parameter specifies airflow override options. Check the [Official documentation](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-env-variables.html#configuring-env-variables-reference) for all possible configuration options.
	AirflowConfigurationOptions pulumi.StringMapOutput `pulumi:"airflowConfigurationOptions"`
	// Airflow version of your environment, will be set by default to the latest version that MWAA supports.
	AirflowVersion pulumi.StringOutput `pulumi:"airflowVersion"`
	// The ARN of the MWAA Environment
	Arn pulumi.StringOutput `pulumi:"arn"`
	// The Created At date of the MWAA Environment
	// * `logging_configuration[0].<LOG_CONFIGURATION_TYPE>[0].cloud_watch_log_group_arn` - Provides the ARN for the CloudWatch group where the logs will be published
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// The relative path to the DAG folder on your Amazon S3 storage bucket. For example, dags. For more information, see [Importing DAGs on Amazon MWAA](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import.html).
	DagS3Path pulumi.StringOutput `pulumi:"dagS3Path"`
	// Environment class for the cluster. Possible options are `mw1.small`, `mw1.medium`, `mw1.large`. Will be set by default to `mw1.small`. Please check the [AWS Pricing](https://aws.amazon.com/de/managed-workflows-for-apache-airflow/pricing/) for more information about the environment classes.
	EnvironmentClass pulumi.StringOutput `pulumi:"environmentClass"`
	// The Amazon Resource Name (ARN) of the task execution role that the Amazon MWAA and its environment can assume. Check the [official AWS documentation](https://docs.aws.amazon.com/mwaa/latest/userguide/mwaa-create-role.html) for the detailed role specification.
	ExecutionRoleArn pulumi.StringOutput `pulumi:"executionRoleArn"`
	// The Amazon Resource Name (ARN) of your KMS key that you want to use for encryption. Will be set to the ARN of the managed KMS key `aws/airflow` by default. Please check the [Official Documentation](https://docs.aws.amazon.com/mwaa/latest/userguide/custom-keys-certs.html) for more information.
	KmsKey       pulumi.StringPtrOutput            `pulumi:"kmsKey"`
	LastUpdateds EnvironmentLastUpdatedArrayOutput `pulumi:"lastUpdateds"`
	// The Apache Airflow logs you want to send to Amazon CloudWatch Logs.
	LoggingConfiguration EnvironmentLoggingConfigurationOutput `pulumi:"loggingConfiguration"`
	// The maximum number of workers that can be automatically scaled up. Value need to be between `1` and `25`. Will be `10` by default.
	MaxWorkers pulumi.IntOutput `pulumi:"maxWorkers"`
	// The minimum number of workers that you want to run in your environment. Will be `1` by default.
	MinWorkers pulumi.IntOutput `pulumi:"minWorkers"`
	// The name of the Apache Airflow Environment
	Name pulumi.StringOutput `pulumi:"name"`
	// Specifies the network configuration for your Apache Airflow Environment. This includes two private subnets as well as security groups for the Airflow environment. Each subnet requires internet connection, otherwise the deployment will fail. See Network configuration below for details.
	NetworkConfiguration EnvironmentNetworkConfigurationOutput `pulumi:"networkConfiguration"`
	// The plugins.zip file version you want to use.
	PluginsS3ObjectVersion pulumi.StringOutput `pulumi:"pluginsS3ObjectVersion"`
	// The relative path to the plugins.zip file on your Amazon S3 storage bucket. For example, plugins.zip. If a relative path is provided in the request, then pluginsS3ObjectVersion is required. For more information, see [Importing DAGs on Amazon MWAA](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import.html).
	PluginsS3Path pulumi.StringPtrOutput `pulumi:"pluginsS3Path"`
	// The requirements.txt file version you want to use.
	RequirementsS3ObjectVersion pulumi.StringOutput `pulumi:"requirementsS3ObjectVersion"`
	// The relative path to the requirements.txt file on your Amazon S3 storage bucket. For example, requirements.txt. If a relative path is provided in the request, then requirementsS3ObjectVersion is required. For more information, see [Importing DAGs on Amazon MWAA](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import.html).
	RequirementsS3Path pulumi.StringPtrOutput `pulumi:"requirementsS3Path"`
	// The Service Role ARN of the Amazon MWAA Environment
	ServiceRoleArn pulumi.StringOutput `pulumi:"serviceRoleArn"`
	// The Amazon Resource Name (ARN) of your Amazon S3 storage bucket. For example, arn:aws:s3:::airflow-mybucketname.
	SourceBucketArn pulumi.StringOutput `pulumi:"sourceBucketArn"`
	// The status of the Amazon MWAA Environment
	Status pulumi.StringOutput `pulumi:"status"`
	// A map of resource tags to associate with the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapOutput `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapOutput `pulumi:"tagsAll"`
	// Specifies whether the webserver should be accessible over the internet or via your specified VPC. Possible options: `PRIVATE_ONLY` (default) and `PUBLIC_ONLY`.
	WebserverAccessMode pulumi.StringOutput `pulumi:"webserverAccessMode"`
	// The webserver URL of the MWAA Environment
	WebserverUrl pulumi.StringOutput `pulumi:"webserverUrl"`
	// Specifies the start date for the weekly maintenance window.
	WeeklyMaintenanceWindowStart pulumi.StringOutput `pulumi:"weeklyMaintenanceWindowStart"`
}

// NewEnvironment registers a new resource with the given unique name, arguments, and options.
func NewEnvironment(ctx *pulumi.Context,
	name string, args *EnvironmentArgs, opts ...pulumi.ResourceOption) (*Environment, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DagS3Path == nil {
		return nil, errors.New("invalid value for required argument 'DagS3Path'")
	}
	if args.ExecutionRoleArn == nil {
		return nil, errors.New("invalid value for required argument 'ExecutionRoleArn'")
	}
	if args.NetworkConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'NetworkConfiguration'")
	}
	if args.SourceBucketArn == nil {
		return nil, errors.New("invalid value for required argument 'SourceBucketArn'")
	}
	var resource Environment
	err := ctx.RegisterResource("aws:mwaa/environment:Environment", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetEnvironment gets an existing Environment resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetEnvironment(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnvironmentState, opts ...pulumi.ResourceOption) (*Environment, error) {
	var resource Environment
	err := ctx.ReadResource("aws:mwaa/environment:Environment", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Environment resources.
type environmentState struct {
	// The `airflowConfigurationOptions` parameter specifies airflow override options. Check the [Official documentation](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-env-variables.html#configuring-env-variables-reference) for all possible configuration options.
	AirflowConfigurationOptions map[string]string `pulumi:"airflowConfigurationOptions"`
	// Airflow version of your environment, will be set by default to the latest version that MWAA supports.
	AirflowVersion *string `pulumi:"airflowVersion"`
	// The ARN of the MWAA Environment
	Arn *string `pulumi:"arn"`
	// The Created At date of the MWAA Environment
	// * `logging_configuration[0].<LOG_CONFIGURATION_TYPE>[0].cloud_watch_log_group_arn` - Provides the ARN for the CloudWatch group where the logs will be published
	CreatedAt *string `pulumi:"createdAt"`
	// The relative path to the DAG folder on your Amazon S3 storage bucket. For example, dags. For more information, see [Importing DAGs on Amazon MWAA](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import.html).
	DagS3Path *string `pulumi:"dagS3Path"`
	// Environment class for the cluster. Possible options are `mw1.small`, `mw1.medium`, `mw1.large`. Will be set by default to `mw1.small`. Please check the [AWS Pricing](https://aws.amazon.com/de/managed-workflows-for-apache-airflow/pricing/) for more information about the environment classes.
	EnvironmentClass *string `pulumi:"environmentClass"`
	// The Amazon Resource Name (ARN) of the task execution role that the Amazon MWAA and its environment can assume. Check the [official AWS documentation](https://docs.aws.amazon.com/mwaa/latest/userguide/mwaa-create-role.html) for the detailed role specification.
	ExecutionRoleArn *string `pulumi:"executionRoleArn"`
	// The Amazon Resource Name (ARN) of your KMS key that you want to use for encryption. Will be set to the ARN of the managed KMS key `aws/airflow` by default. Please check the [Official Documentation](https://docs.aws.amazon.com/mwaa/latest/userguide/custom-keys-certs.html) for more information.
	KmsKey       *string                  `pulumi:"kmsKey"`
	LastUpdateds []EnvironmentLastUpdated `pulumi:"lastUpdateds"`
	// The Apache Airflow logs you want to send to Amazon CloudWatch Logs.
	LoggingConfiguration *EnvironmentLoggingConfiguration `pulumi:"loggingConfiguration"`
	// The maximum number of workers that can be automatically scaled up. Value need to be between `1` and `25`. Will be `10` by default.
	MaxWorkers *int `pulumi:"maxWorkers"`
	// The minimum number of workers that you want to run in your environment. Will be `1` by default.
	MinWorkers *int `pulumi:"minWorkers"`
	// The name of the Apache Airflow Environment
	Name *string `pulumi:"name"`
	// Specifies the network configuration for your Apache Airflow Environment. This includes two private subnets as well as security groups for the Airflow environment. Each subnet requires internet connection, otherwise the deployment will fail. See Network configuration below for details.
	NetworkConfiguration *EnvironmentNetworkConfiguration `pulumi:"networkConfiguration"`
	// The plugins.zip file version you want to use.
	PluginsS3ObjectVersion *string `pulumi:"pluginsS3ObjectVersion"`
	// The relative path to the plugins.zip file on your Amazon S3 storage bucket. For example, plugins.zip. If a relative path is provided in the request, then pluginsS3ObjectVersion is required. For more information, see [Importing DAGs on Amazon MWAA](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import.html).
	PluginsS3Path *string `pulumi:"pluginsS3Path"`
	// The requirements.txt file version you want to use.
	RequirementsS3ObjectVersion *string `pulumi:"requirementsS3ObjectVersion"`
	// The relative path to the requirements.txt file on your Amazon S3 storage bucket. For example, requirements.txt. If a relative path is provided in the request, then requirementsS3ObjectVersion is required. For more information, see [Importing DAGs on Amazon MWAA](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import.html).
	RequirementsS3Path *string `pulumi:"requirementsS3Path"`
	// The Service Role ARN of the Amazon MWAA Environment
	ServiceRoleArn *string `pulumi:"serviceRoleArn"`
	// The Amazon Resource Name (ARN) of your Amazon S3 storage bucket. For example, arn:aws:s3:::airflow-mybucketname.
	SourceBucketArn *string `pulumi:"sourceBucketArn"`
	// The status of the Amazon MWAA Environment
	Status *string `pulumi:"status"`
	// A map of resource tags to associate with the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll map[string]string `pulumi:"tagsAll"`
	// Specifies whether the webserver should be accessible over the internet or via your specified VPC. Possible options: `PRIVATE_ONLY` (default) and `PUBLIC_ONLY`.
	WebserverAccessMode *string `pulumi:"webserverAccessMode"`
	// The webserver URL of the MWAA Environment
	WebserverUrl *string `pulumi:"webserverUrl"`
	// Specifies the start date for the weekly maintenance window.
	WeeklyMaintenanceWindowStart *string `pulumi:"weeklyMaintenanceWindowStart"`
}

type EnvironmentState struct {
	// The `airflowConfigurationOptions` parameter specifies airflow override options. Check the [Official documentation](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-env-variables.html#configuring-env-variables-reference) for all possible configuration options.
	AirflowConfigurationOptions pulumi.StringMapInput
	// Airflow version of your environment, will be set by default to the latest version that MWAA supports.
	AirflowVersion pulumi.StringPtrInput
	// The ARN of the MWAA Environment
	Arn pulumi.StringPtrInput
	// The Created At date of the MWAA Environment
	// * `logging_configuration[0].<LOG_CONFIGURATION_TYPE>[0].cloud_watch_log_group_arn` - Provides the ARN for the CloudWatch group where the logs will be published
	CreatedAt pulumi.StringPtrInput
	// The relative path to the DAG folder on your Amazon S3 storage bucket. For example, dags. For more information, see [Importing DAGs on Amazon MWAA](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import.html).
	DagS3Path pulumi.StringPtrInput
	// Environment class for the cluster. Possible options are `mw1.small`, `mw1.medium`, `mw1.large`. Will be set by default to `mw1.small`. Please check the [AWS Pricing](https://aws.amazon.com/de/managed-workflows-for-apache-airflow/pricing/) for more information about the environment classes.
	EnvironmentClass pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the task execution role that the Amazon MWAA and its environment can assume. Check the [official AWS documentation](https://docs.aws.amazon.com/mwaa/latest/userguide/mwaa-create-role.html) for the detailed role specification.
	ExecutionRoleArn pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of your KMS key that you want to use for encryption. Will be set to the ARN of the managed KMS key `aws/airflow` by default. Please check the [Official Documentation](https://docs.aws.amazon.com/mwaa/latest/userguide/custom-keys-certs.html) for more information.
	KmsKey       pulumi.StringPtrInput
	LastUpdateds EnvironmentLastUpdatedArrayInput
	// The Apache Airflow logs you want to send to Amazon CloudWatch Logs.
	LoggingConfiguration EnvironmentLoggingConfigurationPtrInput
	// The maximum number of workers that can be automatically scaled up. Value need to be between `1` and `25`. Will be `10` by default.
	MaxWorkers pulumi.IntPtrInput
	// The minimum number of workers that you want to run in your environment. Will be `1` by default.
	MinWorkers pulumi.IntPtrInput
	// The name of the Apache Airflow Environment
	Name pulumi.StringPtrInput
	// Specifies the network configuration for your Apache Airflow Environment. This includes two private subnets as well as security groups for the Airflow environment. Each subnet requires internet connection, otherwise the deployment will fail. See Network configuration below for details.
	NetworkConfiguration EnvironmentNetworkConfigurationPtrInput
	// The plugins.zip file version you want to use.
	PluginsS3ObjectVersion pulumi.StringPtrInput
	// The relative path to the plugins.zip file on your Amazon S3 storage bucket. For example, plugins.zip. If a relative path is provided in the request, then pluginsS3ObjectVersion is required. For more information, see [Importing DAGs on Amazon MWAA](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import.html).
	PluginsS3Path pulumi.StringPtrInput
	// The requirements.txt file version you want to use.
	RequirementsS3ObjectVersion pulumi.StringPtrInput
	// The relative path to the requirements.txt file on your Amazon S3 storage bucket. For example, requirements.txt. If a relative path is provided in the request, then requirementsS3ObjectVersion is required. For more information, see [Importing DAGs on Amazon MWAA](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import.html).
	RequirementsS3Path pulumi.StringPtrInput
	// The Service Role ARN of the Amazon MWAA Environment
	ServiceRoleArn pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of your Amazon S3 storage bucket. For example, arn:aws:s3:::airflow-mybucketname.
	SourceBucketArn pulumi.StringPtrInput
	// The status of the Amazon MWAA Environment
	Status pulumi.StringPtrInput
	// A map of resource tags to associate with the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// A map of tags assigned to the resource, including those inherited from the provider .
	TagsAll pulumi.StringMapInput
	// Specifies whether the webserver should be accessible over the internet or via your specified VPC. Possible options: `PRIVATE_ONLY` (default) and `PUBLIC_ONLY`.
	WebserverAccessMode pulumi.StringPtrInput
	// The webserver URL of the MWAA Environment
	WebserverUrl pulumi.StringPtrInput
	// Specifies the start date for the weekly maintenance window.
	WeeklyMaintenanceWindowStart pulumi.StringPtrInput
}

func (EnvironmentState) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentState)(nil)).Elem()
}

type environmentArgs struct {
	// The `airflowConfigurationOptions` parameter specifies airflow override options. Check the [Official documentation](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-env-variables.html#configuring-env-variables-reference) for all possible configuration options.
	AirflowConfigurationOptions map[string]string `pulumi:"airflowConfigurationOptions"`
	// Airflow version of your environment, will be set by default to the latest version that MWAA supports.
	AirflowVersion *string `pulumi:"airflowVersion"`
	// The relative path to the DAG folder on your Amazon S3 storage bucket. For example, dags. For more information, see [Importing DAGs on Amazon MWAA](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import.html).
	DagS3Path string `pulumi:"dagS3Path"`
	// Environment class for the cluster. Possible options are `mw1.small`, `mw1.medium`, `mw1.large`. Will be set by default to `mw1.small`. Please check the [AWS Pricing](https://aws.amazon.com/de/managed-workflows-for-apache-airflow/pricing/) for more information about the environment classes.
	EnvironmentClass *string `pulumi:"environmentClass"`
	// The Amazon Resource Name (ARN) of the task execution role that the Amazon MWAA and its environment can assume. Check the [official AWS documentation](https://docs.aws.amazon.com/mwaa/latest/userguide/mwaa-create-role.html) for the detailed role specification.
	ExecutionRoleArn string `pulumi:"executionRoleArn"`
	// The Amazon Resource Name (ARN) of your KMS key that you want to use for encryption. Will be set to the ARN of the managed KMS key `aws/airflow` by default. Please check the [Official Documentation](https://docs.aws.amazon.com/mwaa/latest/userguide/custom-keys-certs.html) for more information.
	KmsKey *string `pulumi:"kmsKey"`
	// The Apache Airflow logs you want to send to Amazon CloudWatch Logs.
	LoggingConfiguration *EnvironmentLoggingConfiguration `pulumi:"loggingConfiguration"`
	// The maximum number of workers that can be automatically scaled up. Value need to be between `1` and `25`. Will be `10` by default.
	MaxWorkers *int `pulumi:"maxWorkers"`
	// The minimum number of workers that you want to run in your environment. Will be `1` by default.
	MinWorkers *int `pulumi:"minWorkers"`
	// The name of the Apache Airflow Environment
	Name *string `pulumi:"name"`
	// Specifies the network configuration for your Apache Airflow Environment. This includes two private subnets as well as security groups for the Airflow environment. Each subnet requires internet connection, otherwise the deployment will fail. See Network configuration below for details.
	NetworkConfiguration EnvironmentNetworkConfiguration `pulumi:"networkConfiguration"`
	// The plugins.zip file version you want to use.
	PluginsS3ObjectVersion *string `pulumi:"pluginsS3ObjectVersion"`
	// The relative path to the plugins.zip file on your Amazon S3 storage bucket. For example, plugins.zip. If a relative path is provided in the request, then pluginsS3ObjectVersion is required. For more information, see [Importing DAGs on Amazon MWAA](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import.html).
	PluginsS3Path *string `pulumi:"pluginsS3Path"`
	// The requirements.txt file version you want to use.
	RequirementsS3ObjectVersion *string `pulumi:"requirementsS3ObjectVersion"`
	// The relative path to the requirements.txt file on your Amazon S3 storage bucket. For example, requirements.txt. If a relative path is provided in the request, then requirementsS3ObjectVersion is required. For more information, see [Importing DAGs on Amazon MWAA](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import.html).
	RequirementsS3Path *string `pulumi:"requirementsS3Path"`
	// The Amazon Resource Name (ARN) of your Amazon S3 storage bucket. For example, arn:aws:s3:::airflow-mybucketname.
	SourceBucketArn string `pulumi:"sourceBucketArn"`
	// A map of resource tags to associate with the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `pulumi:"tags"`
	// Specifies whether the webserver should be accessible over the internet or via your specified VPC. Possible options: `PRIVATE_ONLY` (default) and `PUBLIC_ONLY`.
	WebserverAccessMode *string `pulumi:"webserverAccessMode"`
	// Specifies the start date for the weekly maintenance window.
	WeeklyMaintenanceWindowStart *string `pulumi:"weeklyMaintenanceWindowStart"`
}

// The set of arguments for constructing a Environment resource.
type EnvironmentArgs struct {
	// The `airflowConfigurationOptions` parameter specifies airflow override options. Check the [Official documentation](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-env-variables.html#configuring-env-variables-reference) for all possible configuration options.
	AirflowConfigurationOptions pulumi.StringMapInput
	// Airflow version of your environment, will be set by default to the latest version that MWAA supports.
	AirflowVersion pulumi.StringPtrInput
	// The relative path to the DAG folder on your Amazon S3 storage bucket. For example, dags. For more information, see [Importing DAGs on Amazon MWAA](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import.html).
	DagS3Path pulumi.StringInput
	// Environment class for the cluster. Possible options are `mw1.small`, `mw1.medium`, `mw1.large`. Will be set by default to `mw1.small`. Please check the [AWS Pricing](https://aws.amazon.com/de/managed-workflows-for-apache-airflow/pricing/) for more information about the environment classes.
	EnvironmentClass pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of the task execution role that the Amazon MWAA and its environment can assume. Check the [official AWS documentation](https://docs.aws.amazon.com/mwaa/latest/userguide/mwaa-create-role.html) for the detailed role specification.
	ExecutionRoleArn pulumi.StringInput
	// The Amazon Resource Name (ARN) of your KMS key that you want to use for encryption. Will be set to the ARN of the managed KMS key `aws/airflow` by default. Please check the [Official Documentation](https://docs.aws.amazon.com/mwaa/latest/userguide/custom-keys-certs.html) for more information.
	KmsKey pulumi.StringPtrInput
	// The Apache Airflow logs you want to send to Amazon CloudWatch Logs.
	LoggingConfiguration EnvironmentLoggingConfigurationPtrInput
	// The maximum number of workers that can be automatically scaled up. Value need to be between `1` and `25`. Will be `10` by default.
	MaxWorkers pulumi.IntPtrInput
	// The minimum number of workers that you want to run in your environment. Will be `1` by default.
	MinWorkers pulumi.IntPtrInput
	// The name of the Apache Airflow Environment
	Name pulumi.StringPtrInput
	// Specifies the network configuration for your Apache Airflow Environment. This includes two private subnets as well as security groups for the Airflow environment. Each subnet requires internet connection, otherwise the deployment will fail. See Network configuration below for details.
	NetworkConfiguration EnvironmentNetworkConfigurationInput
	// The plugins.zip file version you want to use.
	PluginsS3ObjectVersion pulumi.StringPtrInput
	// The relative path to the plugins.zip file on your Amazon S3 storage bucket. For example, plugins.zip. If a relative path is provided in the request, then pluginsS3ObjectVersion is required. For more information, see [Importing DAGs on Amazon MWAA](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import.html).
	PluginsS3Path pulumi.StringPtrInput
	// The requirements.txt file version you want to use.
	RequirementsS3ObjectVersion pulumi.StringPtrInput
	// The relative path to the requirements.txt file on your Amazon S3 storage bucket. For example, requirements.txt. If a relative path is provided in the request, then requirementsS3ObjectVersion is required. For more information, see [Importing DAGs on Amazon MWAA](https://docs.aws.amazon.com/mwaa/latest/userguide/configuring-dag-import.html).
	RequirementsS3Path pulumi.StringPtrInput
	// The Amazon Resource Name (ARN) of your Amazon S3 storage bucket. For example, arn:aws:s3:::airflow-mybucketname.
	SourceBucketArn pulumi.StringInput
	// A map of resource tags to associate with the resource. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags pulumi.StringMapInput
	// Specifies whether the webserver should be accessible over the internet or via your specified VPC. Possible options: `PRIVATE_ONLY` (default) and `PUBLIC_ONLY`.
	WebserverAccessMode pulumi.StringPtrInput
	// Specifies the start date for the weekly maintenance window.
	WeeklyMaintenanceWindowStart pulumi.StringPtrInput
}

func (EnvironmentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*environmentArgs)(nil)).Elem()
}

type EnvironmentInput interface {
	pulumi.Input

	ToEnvironmentOutput() EnvironmentOutput
	ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput
}

func (*Environment) ElementType() reflect.Type {
	return reflect.TypeOf((**Environment)(nil)).Elem()
}

func (i *Environment) ToEnvironmentOutput() EnvironmentOutput {
	return i.ToEnvironmentOutputWithContext(context.Background())
}

func (i *Environment) ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentOutput)
}

// EnvironmentArrayInput is an input type that accepts EnvironmentArray and EnvironmentArrayOutput values.
// You can construct a concrete instance of `EnvironmentArrayInput` via:
//
//          EnvironmentArray{ EnvironmentArgs{...} }
type EnvironmentArrayInput interface {
	pulumi.Input

	ToEnvironmentArrayOutput() EnvironmentArrayOutput
	ToEnvironmentArrayOutputWithContext(context.Context) EnvironmentArrayOutput
}

type EnvironmentArray []EnvironmentInput

func (EnvironmentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Environment)(nil)).Elem()
}

func (i EnvironmentArray) ToEnvironmentArrayOutput() EnvironmentArrayOutput {
	return i.ToEnvironmentArrayOutputWithContext(context.Background())
}

func (i EnvironmentArray) ToEnvironmentArrayOutputWithContext(ctx context.Context) EnvironmentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentArrayOutput)
}

// EnvironmentMapInput is an input type that accepts EnvironmentMap and EnvironmentMapOutput values.
// You can construct a concrete instance of `EnvironmentMapInput` via:
//
//          EnvironmentMap{ "key": EnvironmentArgs{...} }
type EnvironmentMapInput interface {
	pulumi.Input

	ToEnvironmentMapOutput() EnvironmentMapOutput
	ToEnvironmentMapOutputWithContext(context.Context) EnvironmentMapOutput
}

type EnvironmentMap map[string]EnvironmentInput

func (EnvironmentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Environment)(nil)).Elem()
}

func (i EnvironmentMap) ToEnvironmentMapOutput() EnvironmentMapOutput {
	return i.ToEnvironmentMapOutputWithContext(context.Background())
}

func (i EnvironmentMap) ToEnvironmentMapOutputWithContext(ctx context.Context) EnvironmentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentMapOutput)
}

type EnvironmentOutput struct{ *pulumi.OutputState }

func (EnvironmentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Environment)(nil)).Elem()
}

func (o EnvironmentOutput) ToEnvironmentOutput() EnvironmentOutput {
	return o
}

func (o EnvironmentOutput) ToEnvironmentOutputWithContext(ctx context.Context) EnvironmentOutput {
	return o
}

type EnvironmentArrayOutput struct{ *pulumi.OutputState }

func (EnvironmentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Environment)(nil)).Elem()
}

func (o EnvironmentArrayOutput) ToEnvironmentArrayOutput() EnvironmentArrayOutput {
	return o
}

func (o EnvironmentArrayOutput) ToEnvironmentArrayOutputWithContext(ctx context.Context) EnvironmentArrayOutput {
	return o
}

func (o EnvironmentArrayOutput) Index(i pulumi.IntInput) EnvironmentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Environment {
		return vs[0].([]*Environment)[vs[1].(int)]
	}).(EnvironmentOutput)
}

type EnvironmentMapOutput struct{ *pulumi.OutputState }

func (EnvironmentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Environment)(nil)).Elem()
}

func (o EnvironmentMapOutput) ToEnvironmentMapOutput() EnvironmentMapOutput {
	return o
}

func (o EnvironmentMapOutput) ToEnvironmentMapOutputWithContext(ctx context.Context) EnvironmentMapOutput {
	return o
}

func (o EnvironmentMapOutput) MapIndex(k pulumi.StringInput) EnvironmentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Environment {
		return vs[0].(map[string]*Environment)[vs[1].(string)]
	}).(EnvironmentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentInput)(nil)).Elem(), &Environment{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentArrayInput)(nil)).Elem(), EnvironmentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EnvironmentMapInput)(nil)).Elem(), EnvironmentMap{})
	pulumi.RegisterOutputType(EnvironmentOutput{})
	pulumi.RegisterOutputType(EnvironmentArrayOutput{})
	pulumi.RegisterOutputType(EnvironmentMapOutput{})
}
