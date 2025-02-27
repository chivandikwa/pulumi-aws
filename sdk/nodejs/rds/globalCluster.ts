// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Manages an RDS Global Cluster, which is an Aurora global database spread across multiple regions. The global database contains a single primary cluster with read-write capability, and a read-only secondary cluster that receives data from the primary cluster through high-speed replication performed by the Aurora storage subsystem.
 *
 * More information about Aurora global databases can be found in the [Aurora User Guide](https://docs.aws.amazon.com/AmazonRDS/latest/AuroraUserGuide/aurora-global-database.html#aurora-global-database-creating).
 *
 * ## Example Usage
 * ### New MySQL Global Cluster
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.rds.GlobalCluster("example", {
 *     globalClusterIdentifier: "global-test",
 *     engine: "aurora",
 *     engineVersion: "5.6.mysql_aurora.1.22.2",
 *     databaseName: "example_db",
 * });
 * const primaryCluster = new aws.rds.Cluster("primaryCluster", {
 *     engine: example.engine,
 *     engineVersion: example.engineVersion,
 *     clusterIdentifier: "test-primary-cluster",
 *     masterUsername: "username",
 *     masterPassword: "somepass123",
 *     databaseName: "example_db",
 *     globalClusterIdentifier: example.id,
 *     dbSubnetGroupName: "default",
 * }, {
 *     provider: aws.primary,
 * });
 * const primaryClusterInstance = new aws.rds.ClusterInstance("primaryClusterInstance", {
 *     engine: example.engine,
 *     engineVersion: example.engineVersion,
 *     identifier: "test-primary-cluster-instance",
 *     clusterIdentifier: primaryCluster.id,
 *     instanceClass: "db.r4.large",
 *     dbSubnetGroupName: "default",
 * }, {
 *     provider: aws.primary,
 * });
 * const secondaryCluster = new aws.rds.Cluster("secondaryCluster", {
 *     engine: example.engine,
 *     engineVersion: example.engineVersion,
 *     clusterIdentifier: "test-secondary-cluster",
 *     globalClusterIdentifier: example.id,
 *     dbSubnetGroupName: "default",
 * }, {
 *     provider: aws.secondary,
 * });
 * const secondaryClusterInstance = new aws.rds.ClusterInstance("secondaryClusterInstance", {
 *     engine: example.engine,
 *     engineVersion: example.engineVersion,
 *     identifier: "test-secondary-cluster-instance",
 *     clusterIdentifier: secondaryCluster.id,
 *     instanceClass: "db.r4.large",
 *     dbSubnetGroupName: "default",
 * }, {
 *     provider: aws.secondary,
 *     dependsOn: [primaryClusterInstance],
 * });
 * ```
 * ### New PostgreSQL Global Cluster
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const primary = new aws.Provider("primary", {region: "us-east-2"});
 * const secondary = new aws.Provider("secondary", {region: "us-east-1"});
 * const example = new aws.rds.GlobalCluster("example", {
 *     globalClusterIdentifier: "global-test",
 *     engine: "aurora-postgresql",
 *     engineVersion: "11.9",
 *     databaseName: "example_db",
 * });
 * const primaryCluster = new aws.rds.Cluster("primaryCluster", {
 *     engine: example.engine,
 *     engineVersion: example.engineVersion,
 *     clusterIdentifier: "test-primary-cluster",
 *     masterUsername: "username",
 *     masterPassword: "somepass123",
 *     databaseName: "example_db",
 *     globalClusterIdentifier: example.id,
 *     dbSubnetGroupName: "default",
 * }, {
 *     provider: aws.primary,
 * });
 * const primaryClusterInstance = new aws.rds.ClusterInstance("primaryClusterInstance", {
 *     engine: example.engine,
 *     engineVersion: example.engineVersion,
 *     identifier: "test-primary-cluster-instance",
 *     clusterIdentifier: primaryCluster.id,
 *     instanceClass: "db.r4.large",
 *     dbSubnetGroupName: "default",
 * }, {
 *     provider: aws.primary,
 * });
 * const secondaryCluster = new aws.rds.Cluster("secondaryCluster", {
 *     engine: example.engine,
 *     engineVersion: example.engineVersion,
 *     clusterIdentifier: "test-secondary-cluster",
 *     globalClusterIdentifier: example.id,
 *     skipFinalSnapshot: true,
 *     dbSubnetGroupName: "default",
 * }, {
 *     provider: aws.secondary,
 *     dependsOn: [primaryClusterInstance],
 * });
 * const secondaryClusterInstance = new aws.rds.ClusterInstance("secondaryClusterInstance", {
 *     engine: example.engine,
 *     engineVersion: example.engineVersion,
 *     identifier: "test-secondary-cluster-instance",
 *     clusterIdentifier: secondaryCluster.id,
 *     instanceClass: "db.r4.large",
 *     dbSubnetGroupName: "default",
 * }, {
 *     provider: aws.secondary,
 * });
 * ```
 * ### New Global Cluster From Existing DB Cluster
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * // ... other configuration ...
 * const exampleCluster = new aws.rds.Cluster("exampleCluster", {});
 * const exampleGlobalCluster = new aws.rds.GlobalCluster("exampleGlobalCluster", {
 *     forceDestroy: true,
 *     globalClusterIdentifier: "example",
 *     sourceDbClusterIdentifier: exampleCluster.arn,
 * });
 * ```
 *
 * ## Import
 *
 * `aws_rds_global_cluster` can be imported by using the RDS Global Cluster identifier, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:rds/globalCluster:GlobalCluster example example
 * ```
 *
 *  Certain resource arguments, like `force_destroy`, only exist within this provider. If the argument is set in the the provider configuration on an imported resource, This provider will show a difference on the first plan after import to update the state value. This change is safe to apply immediately so the state matches the desired configuration. Certain resource arguments, like `source_db_cluster_identifier`, do not have an API method for reading the information after creation. If the argument is set in the provider configuration on an imported resource, the provider will always show a difference. To workaround this behavior, either omit the argument from the the provider configuration or use `ignore_changes` to hide the difference, e.g. terraform resource "aws_rds_global_cluster" "example" {
 *
 * # ... other configuration ...
 *
 * # There is no API for reading source_db_cluster_identifier
 *
 *  lifecycle {
 *
 *  ignore_changes = [source_db_cluster_identifier]
 *
 *  } }
 */
export class GlobalCluster extends pulumi.CustomResource {
    /**
     * Get an existing GlobalCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: GlobalClusterState, opts?: pulumi.CustomResourceOptions): GlobalCluster {
        return new GlobalCluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:rds/globalCluster:GlobalCluster';

    /**
     * Returns true if the given object is an instance of GlobalCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is GlobalCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === GlobalCluster.__pulumiType;
    }

    /**
     * RDS Global Cluster Amazon Resource Name (ARN)
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    /**
     * Name for an automatically created database on cluster creation.
     */
    public readonly databaseName!: pulumi.Output<string | undefined>;
    /**
     * If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
     */
    public readonly deletionProtection!: pulumi.Output<boolean | undefined>;
    /**
     * Name of the database engine to be used for this DB cluster. The provider will only perform drift detection if a configuration value is provided. Valid values: `aurora`, `aurora-mysql`, `aurora-postgresql`. Defaults to `aurora`. Conflicts with `sourceDbClusterIdentifier`.
     */
    public readonly engine!: pulumi.Output<string>;
    /**
     * Engine version of the Aurora global database.
     * * **NOTE:** When the engine is set to `aurora-mysql`, an engine version compatible with global database is required. The earliest available version is `5.7.mysql_aurora.2.06.0`.
     */
    public readonly engineVersion!: pulumi.Output<string>;
    /**
     * Enable to remove DB Cluster members from Global Cluster on destroy. Required with `sourceDbClusterIdentifier`.
     */
    public readonly forceDestroy!: pulumi.Output<boolean | undefined>;
    /**
     * The global cluster identifier.
     */
    public readonly globalClusterIdentifier!: pulumi.Output<string>;
    /**
     * Set of objects containing Global Cluster members.
     */
    public /*out*/ readonly globalClusterMembers!: pulumi.Output<outputs.rds.GlobalClusterGlobalClusterMember[]>;
    /**
     * AWS Region-unique, immutable identifier for the global database cluster. This identifier is found in AWS CloudTrail log entries whenever the AWS KMS key for the DB cluster is accessed
     */
    public /*out*/ readonly globalClusterResourceId!: pulumi.Output<string>;
    /**
     * Amazon Resource Name (ARN) to use as the primary DB Cluster of the Global Cluster on creation. The provider cannot perform drift detection of this value.
     */
    public readonly sourceDbClusterIdentifier!: pulumi.Output<string>;
    /**
     * Specifies whether the DB cluster is encrypted. The default is `false` unless `sourceDbClusterIdentifier` is specified and encrypted. The provider will only perform drift detection if a configuration value is provided.
     */
    public readonly storageEncrypted!: pulumi.Output<boolean>;

    /**
     * Create a GlobalCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: GlobalClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: GlobalClusterArgs | GlobalClusterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as GlobalClusterState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["databaseName"] = state ? state.databaseName : undefined;
            resourceInputs["deletionProtection"] = state ? state.deletionProtection : undefined;
            resourceInputs["engine"] = state ? state.engine : undefined;
            resourceInputs["engineVersion"] = state ? state.engineVersion : undefined;
            resourceInputs["forceDestroy"] = state ? state.forceDestroy : undefined;
            resourceInputs["globalClusterIdentifier"] = state ? state.globalClusterIdentifier : undefined;
            resourceInputs["globalClusterMembers"] = state ? state.globalClusterMembers : undefined;
            resourceInputs["globalClusterResourceId"] = state ? state.globalClusterResourceId : undefined;
            resourceInputs["sourceDbClusterIdentifier"] = state ? state.sourceDbClusterIdentifier : undefined;
            resourceInputs["storageEncrypted"] = state ? state.storageEncrypted : undefined;
        } else {
            const args = argsOrState as GlobalClusterArgs | undefined;
            if ((!args || args.globalClusterIdentifier === undefined) && !opts.urn) {
                throw new Error("Missing required property 'globalClusterIdentifier'");
            }
            resourceInputs["databaseName"] = args ? args.databaseName : undefined;
            resourceInputs["deletionProtection"] = args ? args.deletionProtection : undefined;
            resourceInputs["engine"] = args ? args.engine : undefined;
            resourceInputs["engineVersion"] = args ? args.engineVersion : undefined;
            resourceInputs["forceDestroy"] = args ? args.forceDestroy : undefined;
            resourceInputs["globalClusterIdentifier"] = args ? args.globalClusterIdentifier : undefined;
            resourceInputs["sourceDbClusterIdentifier"] = args ? args.sourceDbClusterIdentifier : undefined;
            resourceInputs["storageEncrypted"] = args ? args.storageEncrypted : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["globalClusterMembers"] = undefined /*out*/;
            resourceInputs["globalClusterResourceId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(GlobalCluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering GlobalCluster resources.
 */
export interface GlobalClusterState {
    /**
     * RDS Global Cluster Amazon Resource Name (ARN)
     */
    arn?: pulumi.Input<string>;
    /**
     * Name for an automatically created database on cluster creation.
     */
    databaseName?: pulumi.Input<string>;
    /**
     * If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
     */
    deletionProtection?: pulumi.Input<boolean>;
    /**
     * Name of the database engine to be used for this DB cluster. The provider will only perform drift detection if a configuration value is provided. Valid values: `aurora`, `aurora-mysql`, `aurora-postgresql`. Defaults to `aurora`. Conflicts with `sourceDbClusterIdentifier`.
     */
    engine?: pulumi.Input<string>;
    /**
     * Engine version of the Aurora global database.
     * * **NOTE:** When the engine is set to `aurora-mysql`, an engine version compatible with global database is required. The earliest available version is `5.7.mysql_aurora.2.06.0`.
     */
    engineVersion?: pulumi.Input<string>;
    /**
     * Enable to remove DB Cluster members from Global Cluster on destroy. Required with `sourceDbClusterIdentifier`.
     */
    forceDestroy?: pulumi.Input<boolean>;
    /**
     * The global cluster identifier.
     */
    globalClusterIdentifier?: pulumi.Input<string>;
    /**
     * Set of objects containing Global Cluster members.
     */
    globalClusterMembers?: pulumi.Input<pulumi.Input<inputs.rds.GlobalClusterGlobalClusterMember>[]>;
    /**
     * AWS Region-unique, immutable identifier for the global database cluster. This identifier is found in AWS CloudTrail log entries whenever the AWS KMS key for the DB cluster is accessed
     */
    globalClusterResourceId?: pulumi.Input<string>;
    /**
     * Amazon Resource Name (ARN) to use as the primary DB Cluster of the Global Cluster on creation. The provider cannot perform drift detection of this value.
     */
    sourceDbClusterIdentifier?: pulumi.Input<string>;
    /**
     * Specifies whether the DB cluster is encrypted. The default is `false` unless `sourceDbClusterIdentifier` is specified and encrypted. The provider will only perform drift detection if a configuration value is provided.
     */
    storageEncrypted?: pulumi.Input<boolean>;
}

/**
 * The set of arguments for constructing a GlobalCluster resource.
 */
export interface GlobalClusterArgs {
    /**
     * Name for an automatically created database on cluster creation.
     */
    databaseName?: pulumi.Input<string>;
    /**
     * If the Global Cluster should have deletion protection enabled. The database can't be deleted when this value is set to `true`. The default is `false`.
     */
    deletionProtection?: pulumi.Input<boolean>;
    /**
     * Name of the database engine to be used for this DB cluster. The provider will only perform drift detection if a configuration value is provided. Valid values: `aurora`, `aurora-mysql`, `aurora-postgresql`. Defaults to `aurora`. Conflicts with `sourceDbClusterIdentifier`.
     */
    engine?: pulumi.Input<string>;
    /**
     * Engine version of the Aurora global database.
     * * **NOTE:** When the engine is set to `aurora-mysql`, an engine version compatible with global database is required. The earliest available version is `5.7.mysql_aurora.2.06.0`.
     */
    engineVersion?: pulumi.Input<string>;
    /**
     * Enable to remove DB Cluster members from Global Cluster on destroy. Required with `sourceDbClusterIdentifier`.
     */
    forceDestroy?: pulumi.Input<boolean>;
    /**
     * The global cluster identifier.
     */
    globalClusterIdentifier: pulumi.Input<string>;
    /**
     * Amazon Resource Name (ARN) to use as the primary DB Cluster of the Global Cluster on creation. The provider cannot perform drift detection of this value.
     */
    sourceDbClusterIdentifier?: pulumi.Input<string>;
    /**
     * Specifies whether the DB cluster is encrypted. The default is `false` unless `sourceDbClusterIdentifier` is specified and encrypted. The provider will only perform drift detection if a configuration value is provided.
     */
    storageEncrypted?: pulumi.Input<boolean>;
}
