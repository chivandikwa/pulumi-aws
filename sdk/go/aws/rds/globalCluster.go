// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package rds

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Manages an RDS Global Cluster, which is an Aurora global database spread across multiple regions. The global database contains a single primary cluster with read-write capability, and a read-only secondary cluster that receives data from the primary cluster through high-speed replication performed by the Aurora storage subsystem.
//
// More information about Aurora global databases can be found in the [Aurora User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database.html#aurora-global-database-creating).
//
// ## Example Usage
// ### New MySQL Global Cluster
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/rds"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		example, err := rds.NewGlobalCluster(ctx, "example", &rds.GlobalClusterArgs{
// 			GlobalClusterIdentifier: pulumi.String("global-test"),
// 			Engine:                  pulumi.String("aurora"),
// 			EngineVersion:           pulumi.String("5.6.mysql_aurora.1.22.2"),
// 			DatabaseName:            pulumi.String("example_db"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		primaryCluster, err := rds.NewCluster(ctx, "primaryCluster", &rds.ClusterArgs{
// 			Engine:                  example.Engine,
// 			EngineVersion:           example.EngineVersion,
// 			ClusterIdentifier:       pulumi.String("test-primary-cluster"),
// 			MasterUsername:          pulumi.String("username"),
// 			MasterPassword:          pulumi.String("somepass123"),
// 			DatabaseName:            pulumi.String("example_db"),
// 			GlobalClusterIdentifier: example.ID(),
// 			DbSubnetGroupName:       pulumi.String("default"),
// 		}, pulumi.Provider(aws.Primary))
// 		if err != nil {
// 			return err
// 		}
// 		primaryClusterInstance, err := rds.NewClusterInstance(ctx, "primaryClusterInstance", &rds.ClusterInstanceArgs{
// 			Engine:            example.Engine,
// 			EngineVersion:     example.EngineVersion,
// 			Identifier:        pulumi.String("test-primary-cluster-instance"),
// 			ClusterIdentifier: primaryCluster.ID(),
// 			InstanceClass:     pulumi.String("db.r4.large"),
// 			DbSubnetGroupName: pulumi.String("default"),
// 		}, pulumi.Provider(aws.Primary))
// 		if err != nil {
// 			return err
// 		}
// 		secondaryCluster, err := rds.NewCluster(ctx, "secondaryCluster", &rds.ClusterArgs{
// 			Engine:                  example.Engine,
// 			EngineVersion:           example.EngineVersion,
// 			ClusterIdentifier:       pulumi.String("test-secondary-cluster"),
// 			GlobalClusterIdentifier: example.ID(),
// 			DbSubnetGroupName:       pulumi.String("default"),
// 		}, pulumi.Provider(aws.Secondary))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = rds.NewClusterInstance(ctx, "secondaryClusterInstance", &rds.ClusterInstanceArgs{
// 			Engine:            example.Engine,
// 			EngineVersion:     example.EngineVersion,
// 			Identifier:        pulumi.String("test-secondary-cluster-instance"),
// 			ClusterIdentifier: secondaryCluster.ID(),
// 			InstanceClass:     pulumi.String("db.r4.large"),
// 			DbSubnetGroupName: pulumi.String("default"),
// 		}, pulumi.Provider(aws.Secondary), pulumi.DependsOn([]pulumi.Resource{
// 			primaryClusterInstance,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### New PostgreSQL Global Cluster
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/providers"
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/rds"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := providers.Newaws(ctx, "primary", &providers.awsArgs{
// 			Region: "us-east-2",
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = providers.Newaws(ctx, "secondary", &providers.awsArgs{
// 			Region: "us-east-1",
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		example, err := rds.NewGlobalCluster(ctx, "example", &rds.GlobalClusterArgs{
// 			GlobalClusterIdentifier: pulumi.String("global-test"),
// 			Engine:                  pulumi.String("aurora-postgresql"),
// 			EngineVersion:           pulumi.String("11.9"),
// 			DatabaseName:            pulumi.String("example_db"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		primaryCluster, err := rds.NewCluster(ctx, "primaryCluster", &rds.ClusterArgs{
// 			Engine:                  example.Engine,
// 			EngineVersion:           example.EngineVersion,
// 			ClusterIdentifier:       pulumi.String("test-primary-cluster"),
// 			MasterUsername:          pulumi.String("username"),
// 			MasterPassword:          pulumi.String("somepass123"),
// 			DatabaseName:            pulumi.String("example_db"),
// 			GlobalClusterIdentifier: example.ID(),
// 			DbSubnetGroupName:       pulumi.String("default"),
// 		}, pulumi.Provider(aws.Primary))
// 		if err != nil {
// 			return err
// 		}
// 		primaryClusterInstance, err := rds.NewClusterInstance(ctx, "primaryClusterInstance", &rds.ClusterInstanceArgs{
// 			Engine:            example.Engine,
// 			EngineVersion:     example.EngineVersion,
// 			Identifier:        pulumi.String("test-primary-cluster-instance"),
// 			ClusterIdentifier: primaryCluster.ID(),
// 			InstanceClass:     pulumi.String("db.r4.large"),
// 			DbSubnetGroupName: pulumi.String("default"),
// 		}, pulumi.Provider(aws.Primary))
// 		if err != nil {
// 			return err
// 		}
// 		secondaryCluster, err := rds.NewCluster(ctx, "secondaryCluster", &rds.ClusterArgs{
// 			Engine:                  example.Engine,
// 			EngineVersion:           example.EngineVersion,
// 			ClusterIdentifier:       pulumi.String("test-secondary-cluster"),
// 			GlobalClusterIdentifier: example.ID(),
// 			SkipFinalSnapshot:       pulumi.Bool(true),
// 			DbSubnetGroupName:       pulumi.String("default"),
// 		}, pulumi.Provider(aws.Secondary), pulumi.DependsOn([]pulumi.Resource{
// 			primaryClusterInstance,
// 		}))
// 		if err != nil {
// 			return err
// 		}
// 		_, err = rds.NewClusterInstance(ctx, "secondaryClusterInstance", &rds.ClusterInstanceArgs{
// 			Engine:            example.Engine,
// 			EngineVersion:     example.EngineVersion,
// 			Identifier:        pulumi.String("test-secondary-cluster-instance"),
// 			ClusterIdentifier: secondaryCluster.ID(),
// 			InstanceClass:     pulumi.String("db.r4.large"),
// 			DbSubnetGroupName: pulumi.String("default"),
// 		}, pulumi.Provider(aws.Secondary))
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
// ### New Global Cluster From Existing DB Cluster
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/rds"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		exampleCluster, err := rds.NewCluster(ctx, "exampleCluster", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = rds.NewGlobalCluster(ctx, "exampleGlobalCluster", &rds.GlobalClusterArgs{
// 			ForceDestroy:              pulumi.Bool(true),
// 			GlobalClusterIdentifier:   pulumi.String("example"),
// 			SourceDbClusterIdentifier: exampleCluster.Arn,
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
// `aws_rds_global_cluster` can be imported by using the RDS Global Cluster identifier, e.g.,
//
// ```sh
//  $ pulumi import aws:rds/globalCluster:GlobalCluster example example
// ```
//
//  Certain resource arguments, like `force_destroy`, only exist within this provider. If the argument is set in the the provider configuration on an imported resource, This provider will show a difference on the first plan after import to update the state value. This change is safe to apply immediately so the state matches the desired configuration. Certain resource arguments, like `source_db_cluster_identifier`, do not have an API method for reading the information after creation. If the argument is set in the provider configuration on an imported resource, the provider will always show a difference. To workaround this behavior, either omit the argument from the the provider configuration or use `ignore_changes` to hide the difference, e.g. terraform resource "aws_rds_global_cluster" "example" {
//
// # ... other configuration ...
//
// # There is no API for reading source_db_cluster_identifier
//
//  lifecycle {
//
//  ignore_changes = [source_db_cluster_identifier]
//
//  } }
type GlobalCluster struct {
	pulumi.CustomResourceState

	// RDS Global Cluster Amazon Resource Name (ARN)
	Arn pulumi.StringOutput `pulumi:"arn"`
	// Name for an automatically created database on cluster creation.
	DatabaseName pulumi.StringPtrOutput `pulumi:"databaseName"`
	// If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
	DeletionProtection pulumi.BoolPtrOutput `pulumi:"deletionProtection"`
	// Name of the database engine to be used for this DB cluster. The provider will only perform drift detection if a configuration value is provided. Valid values: `aurora`, `aurora-mysql`, `aurora-postgresql`. Defaults to `aurora`. Conflicts with `sourceDbClusterIdentifier`.
	Engine pulumi.StringOutput `pulumi:"engine"`
	// Engine version of the Aurora global database.
	// * **NOTE:** When the engine is set to `aurora-mysql`, an engine version compatible with global database is required. The earliest available version is `5.7.mysql_aurora.2.06.0`.
	EngineVersion pulumi.StringOutput `pulumi:"engineVersion"`
	// Enable to remove DB Cluster members from Global Cluster on destroy. Required with `sourceDbClusterIdentifier`.
	ForceDestroy pulumi.BoolPtrOutput `pulumi:"forceDestroy"`
	// The global cluster identifier.
	GlobalClusterIdentifier pulumi.StringOutput `pulumi:"globalClusterIdentifier"`
	// Set of objects containing Global Cluster members.
	GlobalClusterMembers GlobalClusterGlobalClusterMemberArrayOutput `pulumi:"globalClusterMembers"`
	// AWS Region-unique, immutable identifier for the global database cluster. This identifier is found in AWS CloudTrail log entries whenever the AWS KMS key for the DB cluster is accessed
	GlobalClusterResourceId pulumi.StringOutput `pulumi:"globalClusterResourceId"`
	// Amazon Resource Name (ARN) to use as the primary DB Cluster of the Global Cluster on creation. The provider cannot perform drift detection of this value.
	SourceDbClusterIdentifier pulumi.StringOutput `pulumi:"sourceDbClusterIdentifier"`
	// Specifies whether the DB cluster is encrypted. The default is `false` unless `sourceDbClusterIdentifier` is specified and encrypted. The provider will only perform drift detection if a configuration value is provided.
	StorageEncrypted pulumi.BoolOutput `pulumi:"storageEncrypted"`
}

// NewGlobalCluster registers a new resource with the given unique name, arguments, and options.
func NewGlobalCluster(ctx *pulumi.Context,
	name string, args *GlobalClusterArgs, opts ...pulumi.ResourceOption) (*GlobalCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GlobalClusterIdentifier == nil {
		return nil, errors.New("invalid value for required argument 'GlobalClusterIdentifier'")
	}
	var resource GlobalCluster
	err := ctx.RegisterResource("aws:rds/globalCluster:GlobalCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGlobalCluster gets an existing GlobalCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGlobalCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GlobalClusterState, opts ...pulumi.ResourceOption) (*GlobalCluster, error) {
	var resource GlobalCluster
	err := ctx.ReadResource("aws:rds/globalCluster:GlobalCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering GlobalCluster resources.
type globalClusterState struct {
	// RDS Global Cluster Amazon Resource Name (ARN)
	Arn *string `pulumi:"arn"`
	// Name for an automatically created database on cluster creation.
	DatabaseName *string `pulumi:"databaseName"`
	// If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// Name of the database engine to be used for this DB cluster. The provider will only perform drift detection if a configuration value is provided. Valid values: `aurora`, `aurora-mysql`, `aurora-postgresql`. Defaults to `aurora`. Conflicts with `sourceDbClusterIdentifier`.
	Engine *string `pulumi:"engine"`
	// Engine version of the Aurora global database.
	// * **NOTE:** When the engine is set to `aurora-mysql`, an engine version compatible with global database is required. The earliest available version is `5.7.mysql_aurora.2.06.0`.
	EngineVersion *string `pulumi:"engineVersion"`
	// Enable to remove DB Cluster members from Global Cluster on destroy. Required with `sourceDbClusterIdentifier`.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// The global cluster identifier.
	GlobalClusterIdentifier *string `pulumi:"globalClusterIdentifier"`
	// Set of objects containing Global Cluster members.
	GlobalClusterMembers []GlobalClusterGlobalClusterMember `pulumi:"globalClusterMembers"`
	// AWS Region-unique, immutable identifier for the global database cluster. This identifier is found in AWS CloudTrail log entries whenever the AWS KMS key for the DB cluster is accessed
	GlobalClusterResourceId *string `pulumi:"globalClusterResourceId"`
	// Amazon Resource Name (ARN) to use as the primary DB Cluster of the Global Cluster on creation. The provider cannot perform drift detection of this value.
	SourceDbClusterIdentifier *string `pulumi:"sourceDbClusterIdentifier"`
	// Specifies whether the DB cluster is encrypted. The default is `false` unless `sourceDbClusterIdentifier` is specified and encrypted. The provider will only perform drift detection if a configuration value is provided.
	StorageEncrypted *bool `pulumi:"storageEncrypted"`
}

type GlobalClusterState struct {
	// RDS Global Cluster Amazon Resource Name (ARN)
	Arn pulumi.StringPtrInput
	// Name for an automatically created database on cluster creation.
	DatabaseName pulumi.StringPtrInput
	// If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
	DeletionProtection pulumi.BoolPtrInput
	// Name of the database engine to be used for this DB cluster. The provider will only perform drift detection if a configuration value is provided. Valid values: `aurora`, `aurora-mysql`, `aurora-postgresql`. Defaults to `aurora`. Conflicts with `sourceDbClusterIdentifier`.
	Engine pulumi.StringPtrInput
	// Engine version of the Aurora global database.
	// * **NOTE:** When the engine is set to `aurora-mysql`, an engine version compatible with global database is required. The earliest available version is `5.7.mysql_aurora.2.06.0`.
	EngineVersion pulumi.StringPtrInput
	// Enable to remove DB Cluster members from Global Cluster on destroy. Required with `sourceDbClusterIdentifier`.
	ForceDestroy pulumi.BoolPtrInput
	// The global cluster identifier.
	GlobalClusterIdentifier pulumi.StringPtrInput
	// Set of objects containing Global Cluster members.
	GlobalClusterMembers GlobalClusterGlobalClusterMemberArrayInput
	// AWS Region-unique, immutable identifier for the global database cluster. This identifier is found in AWS CloudTrail log entries whenever the AWS KMS key for the DB cluster is accessed
	GlobalClusterResourceId pulumi.StringPtrInput
	// Amazon Resource Name (ARN) to use as the primary DB Cluster of the Global Cluster on creation. The provider cannot perform drift detection of this value.
	SourceDbClusterIdentifier pulumi.StringPtrInput
	// Specifies whether the DB cluster is encrypted. The default is `false` unless `sourceDbClusterIdentifier` is specified and encrypted. The provider will only perform drift detection if a configuration value is provided.
	StorageEncrypted pulumi.BoolPtrInput
}

func (GlobalClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*globalClusterState)(nil)).Elem()
}

type globalClusterArgs struct {
	// Name for an automatically created database on cluster creation.
	DatabaseName *string `pulumi:"databaseName"`
	// If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// Name of the database engine to be used for this DB cluster. The provider will only perform drift detection if a configuration value is provided. Valid values: `aurora`, `aurora-mysql`, `aurora-postgresql`. Defaults to `aurora`. Conflicts with `sourceDbClusterIdentifier`.
	Engine *string `pulumi:"engine"`
	// Engine version of the Aurora global database.
	// * **NOTE:** When the engine is set to `aurora-mysql`, an engine version compatible with global database is required. The earliest available version is `5.7.mysql_aurora.2.06.0`.
	EngineVersion *string `pulumi:"engineVersion"`
	// Enable to remove DB Cluster members from Global Cluster on destroy. Required with `sourceDbClusterIdentifier`.
	ForceDestroy *bool `pulumi:"forceDestroy"`
	// The global cluster identifier.
	GlobalClusterIdentifier string `pulumi:"globalClusterIdentifier"`
	// Amazon Resource Name (ARN) to use as the primary DB Cluster of the Global Cluster on creation. The provider cannot perform drift detection of this value.
	SourceDbClusterIdentifier *string `pulumi:"sourceDbClusterIdentifier"`
	// Specifies whether the DB cluster is encrypted. The default is `false` unless `sourceDbClusterIdentifier` is specified and encrypted. The provider will only perform drift detection if a configuration value is provided.
	StorageEncrypted *bool `pulumi:"storageEncrypted"`
}

// The set of arguments for constructing a GlobalCluster resource.
type GlobalClusterArgs struct {
	// Name for an automatically created database on cluster creation.
	DatabaseName pulumi.StringPtrInput
	// If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
	DeletionProtection pulumi.BoolPtrInput
	// Name of the database engine to be used for this DB cluster. The provider will only perform drift detection if a configuration value is provided. Valid values: `aurora`, `aurora-mysql`, `aurora-postgresql`. Defaults to `aurora`. Conflicts with `sourceDbClusterIdentifier`.
	Engine pulumi.StringPtrInput
	// Engine version of the Aurora global database.
	// * **NOTE:** When the engine is set to `aurora-mysql`, an engine version compatible with global database is required. The earliest available version is `5.7.mysql_aurora.2.06.0`.
	EngineVersion pulumi.StringPtrInput
	// Enable to remove DB Cluster members from Global Cluster on destroy. Required with `sourceDbClusterIdentifier`.
	ForceDestroy pulumi.BoolPtrInput
	// The global cluster identifier.
	GlobalClusterIdentifier pulumi.StringInput
	// Amazon Resource Name (ARN) to use as the primary DB Cluster of the Global Cluster on creation. The provider cannot perform drift detection of this value.
	SourceDbClusterIdentifier pulumi.StringPtrInput
	// Specifies whether the DB cluster is encrypted. The default is `false` unless `sourceDbClusterIdentifier` is specified and encrypted. The provider will only perform drift detection if a configuration value is provided.
	StorageEncrypted pulumi.BoolPtrInput
}

func (GlobalClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*globalClusterArgs)(nil)).Elem()
}

type GlobalClusterInput interface {
	pulumi.Input

	ToGlobalClusterOutput() GlobalClusterOutput
	ToGlobalClusterOutputWithContext(ctx context.Context) GlobalClusterOutput
}

func (*GlobalCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalCluster)(nil)).Elem()
}

func (i *GlobalCluster) ToGlobalClusterOutput() GlobalClusterOutput {
	return i.ToGlobalClusterOutputWithContext(context.Background())
}

func (i *GlobalCluster) ToGlobalClusterOutputWithContext(ctx context.Context) GlobalClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalClusterOutput)
}

// GlobalClusterArrayInput is an input type that accepts GlobalClusterArray and GlobalClusterArrayOutput values.
// You can construct a concrete instance of `GlobalClusterArrayInput` via:
//
//          GlobalClusterArray{ GlobalClusterArgs{...} }
type GlobalClusterArrayInput interface {
	pulumi.Input

	ToGlobalClusterArrayOutput() GlobalClusterArrayOutput
	ToGlobalClusterArrayOutputWithContext(context.Context) GlobalClusterArrayOutput
}

type GlobalClusterArray []GlobalClusterInput

func (GlobalClusterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GlobalCluster)(nil)).Elem()
}

func (i GlobalClusterArray) ToGlobalClusterArrayOutput() GlobalClusterArrayOutput {
	return i.ToGlobalClusterArrayOutputWithContext(context.Background())
}

func (i GlobalClusterArray) ToGlobalClusterArrayOutputWithContext(ctx context.Context) GlobalClusterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalClusterArrayOutput)
}

// GlobalClusterMapInput is an input type that accepts GlobalClusterMap and GlobalClusterMapOutput values.
// You can construct a concrete instance of `GlobalClusterMapInput` via:
//
//          GlobalClusterMap{ "key": GlobalClusterArgs{...} }
type GlobalClusterMapInput interface {
	pulumi.Input

	ToGlobalClusterMapOutput() GlobalClusterMapOutput
	ToGlobalClusterMapOutputWithContext(context.Context) GlobalClusterMapOutput
}

type GlobalClusterMap map[string]GlobalClusterInput

func (GlobalClusterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GlobalCluster)(nil)).Elem()
}

func (i GlobalClusterMap) ToGlobalClusterMapOutput() GlobalClusterMapOutput {
	return i.ToGlobalClusterMapOutputWithContext(context.Background())
}

func (i GlobalClusterMap) ToGlobalClusterMapOutputWithContext(ctx context.Context) GlobalClusterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalClusterMapOutput)
}

type GlobalClusterOutput struct{ *pulumi.OutputState }

func (GlobalClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GlobalCluster)(nil)).Elem()
}

func (o GlobalClusterOutput) ToGlobalClusterOutput() GlobalClusterOutput {
	return o
}

func (o GlobalClusterOutput) ToGlobalClusterOutputWithContext(ctx context.Context) GlobalClusterOutput {
	return o
}

type GlobalClusterArrayOutput struct{ *pulumi.OutputState }

func (GlobalClusterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*GlobalCluster)(nil)).Elem()
}

func (o GlobalClusterArrayOutput) ToGlobalClusterArrayOutput() GlobalClusterArrayOutput {
	return o
}

func (o GlobalClusterArrayOutput) ToGlobalClusterArrayOutputWithContext(ctx context.Context) GlobalClusterArrayOutput {
	return o
}

func (o GlobalClusterArrayOutput) Index(i pulumi.IntInput) GlobalClusterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *GlobalCluster {
		return vs[0].([]*GlobalCluster)[vs[1].(int)]
	}).(GlobalClusterOutput)
}

type GlobalClusterMapOutput struct{ *pulumi.OutputState }

func (GlobalClusterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*GlobalCluster)(nil)).Elem()
}

func (o GlobalClusterMapOutput) ToGlobalClusterMapOutput() GlobalClusterMapOutput {
	return o
}

func (o GlobalClusterMapOutput) ToGlobalClusterMapOutputWithContext(ctx context.Context) GlobalClusterMapOutput {
	return o
}

func (o GlobalClusterMapOutput) MapIndex(k pulumi.StringInput) GlobalClusterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *GlobalCluster {
		return vs[0].(map[string]*GlobalCluster)[vs[1].(string)]
	}).(GlobalClusterOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalClusterInput)(nil)).Elem(), &GlobalCluster{})
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalClusterArrayInput)(nil)).Elem(), GlobalClusterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GlobalClusterMapInput)(nil)).Elem(), GlobalClusterMap{})
	pulumi.RegisterOutputType(GlobalClusterOutput{})
	pulumi.RegisterOutputType(GlobalClusterArrayOutput{})
	pulumi.RegisterOutputType(GlobalClusterMapOutput{})
}
