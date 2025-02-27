// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Provides a Lambda event source mapping. This allows Lambda functions to get events from Kinesis, DynamoDB, SQS, Amazon MQ and Managed Streaming for Apache Kafka (MSK).
 *
 * For information about Lambda and how to use it, see [What is AWS Lambda?](http://docs.aws.amazon.com/lambda/latest/dg/welcome.html).
 * For information about event source mappings, see [CreateEventSourceMapping](http://docs.aws.amazon.com/lambda/latest/dg/API_CreateEventSourceMapping.html) in the API docs.
 *
 * ## Example Usage
 * ### DynamoDB
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.lambda.EventSourceMapping("example", {
 *     eventSourceArn: aws_dynamodb_table.example.stream_arn,
 *     functionName: aws_lambda_function.example.arn,
 *     startingPosition: "LATEST",
 * });
 * ```
 * ### Kinesis
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.lambda.EventSourceMapping("example", {
 *     eventSourceArn: aws_kinesis_stream.example.arn,
 *     functionName: aws_lambda_function.example.arn,
 *     startingPosition: "LATEST",
 * });
 * ```
 * ### Managed Streaming for Apache Kafka (MSK)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.lambda.EventSourceMapping("example", {
 *     eventSourceArn: aws_msk_cluster.example.arn,
 *     functionName: aws_lambda_function.example.arn,
 *     topics: ["Example"],
 *     startingPosition: "TRIM_HORIZON",
 * });
 * ```
 * ### Self Managed Apache Kafka
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.lambda.EventSourceMapping("example", {
 *     functionName: aws_lambda_function.example.arn,
 *     topics: ["Example"],
 *     startingPosition: "TRIM_HORIZON",
 *     selfManagedEventSource: {
 *         endpoints: {
 *             KAFKA_BOOTSTRAP_SERVERS: "kafka1.example.com:9092,kafka2.example.com:9092",
 *         },
 *     },
 *     sourceAccessConfigurations: [
 *         {
 *             type: "VPC_SUBNET",
 *             uri: "subnet:subnet-example1",
 *         },
 *         {
 *             type: "VPC_SUBNET",
 *             uri: "subnet:subnet-example2",
 *         },
 *         {
 *             type: "VPC_SECURITY_GROUP",
 *             uri: "security_group:sg-example",
 *         },
 *     ],
 * });
 * ```
 * ### SQS
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.lambda.EventSourceMapping("example", {
 *     eventSourceArn: aws_sqs_queue.sqs_queue_test.arn,
 *     functionName: aws_lambda_function.example.arn,
 * });
 * ```
 * ### SQS with event filter
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.lambda.EventSourceMapping("example", {
 *     eventSourceArn: aws_sqs_queue.sqs_queue_test.arn,
 *     functionName: aws_lambda_function.example.arn,
 *     filterCriteria: {
 *         filters: [{
 *             pattern: JSON.stringify({
 *                 body: {
 *                     Temperature: [{
 *                         numeric: [
 *                             ">",
 *                             0,
 *                             "<=",
 *                             100,
 *                         ],
 *                     }],
 *                     Location: ["New York"],
 *                 },
 *             }),
 *         }],
 *     },
 * });
 * ```
 * ### Amazon MQ (ActiveMQ)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.lambda.EventSourceMapping("example", {
 *     batchSize: 10,
 *     eventSourceArn: aws_mq_broker.example.arn,
 *     enabled: true,
 *     functionName: aws_lambda_function.example.arn,
 *     queues: ["example"],
 *     sourceAccessConfigurations: [{
 *         type: "BASIC_AUTH",
 *         uri: aws_secretsmanager_secret_version.example.arn,
 *     }],
 * });
 * ```
 * ### Amazon MQ (RabbitMQ)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.lambda.EventSourceMapping("example", {
 *     batchSize: 1,
 *     eventSourceArn: aws_mq_broker.example.arn,
 *     enabled: true,
 *     functionName: aws_lambda_function.example.arn,
 *     queues: ["example"],
 *     sourceAccessConfigurations: [
 *         {
 *             type: "VIRTUAL_HOST",
 *             uri: "/example",
 *         },
 *         {
 *             type: "BASIC_AUTH",
 *             uri: aws_secretsmanager_secret_version.example.arn,
 *         },
 *     ],
 * });
 * ```
 * ### Managed Streaming for Kafka (MSK)
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const example = new aws.lambda.EventSourceMapping("example", {
 *     eventSourceArn: aws_msk_cluster.example.arn,
 *     functionName: aws_lambda_function.example.arn,
 *     topics: ["Example"],
 *     startingPosition: "TRIM_HORIZON",
 * });
 * ```
 *
 * ## Import
 *
 * Lambda event source mappings can be imported using the `UUID` (event source mapping identifier), e.g.,
 *
 * ```sh
 *  $ pulumi import aws:lambda/eventSourceMapping:EventSourceMapping event_source_mapping 12345kxodurf3443
 * ```
 */
export class EventSourceMapping extends pulumi.CustomResource {
    /**
     * Get an existing EventSourceMapping resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: EventSourceMappingState, opts?: pulumi.CustomResourceOptions): EventSourceMapping {
        return new EventSourceMapping(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:lambda/eventSourceMapping:EventSourceMapping';

    /**
     * Returns true if the given object is an instance of EventSourceMapping.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is EventSourceMapping {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === EventSourceMapping.__pulumiType;
    }

    /**
     * The largest number of records that Lambda will retrieve from your event source at the time of invocation. Defaults to `100` for DynamoDB, Kinesis, MQ and MSK, `10` for SQS.
     * * `bisectBatchOnFunctionError`: - (Optional) If the function returns an error, split the batch in two and retry. Only available for stream sources (DynamoDB and Kinesis). Defaults to `false`.
     * * `destinationConfig`: - (Optional) An Amazon SQS queue or Amazon SNS topic destination for failed records. Only available for stream sources (DynamoDB and Kinesis). Detailed below.
     */
    public readonly batchSize!: pulumi.Output<number | undefined>;
    public readonly bisectBatchOnFunctionError!: pulumi.Output<boolean | undefined>;
    public readonly destinationConfig!: pulumi.Output<outputs.lambda.EventSourceMappingDestinationConfig | undefined>;
    /**
     * Determines if the mapping will be enabled on creation. Defaults to `true`.
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * The event source ARN - this is required for Kinesis stream, DynamoDB stream, SQS queue, MQ broker or MSK cluster.  It is incompatible with a Self Managed Kafka source.
     */
    public readonly eventSourceArn!: pulumi.Output<string | undefined>;
    /**
     * The criteria to use for [event filtering](https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html) Kinesis stream, DynamoDB stream, SQS queue event sources. Detailed below.
     */
    public readonly filterCriteria!: pulumi.Output<outputs.lambda.EventSourceMappingFilterCriteria | undefined>;
    /**
     * The the ARN of the Lambda function the event source mapping is sending events to. (Note: this is a computed value that differs from `functionName` above.)
     */
    public /*out*/ readonly functionArn!: pulumi.Output<string>;
    /**
     * The name or the ARN of the Lambda function that will be subscribing to events.
     */
    public readonly functionName!: pulumi.Output<string>;
    /**
     * A list of current response type enums applied to the event source mapping for [AWS Lambda checkpointing](https://docs.aws.amazon.com/lambda/latest/dg/with-ddb.html#services-ddb-batchfailurereporting). Only available for SQS and stream sources (DynamoDB and Kinesis). Valid values: `ReportBatchItemFailures`.
     */
    public readonly functionResponseTypes!: pulumi.Output<string[] | undefined>;
    /**
     * The date this resource was last modified.
     */
    public /*out*/ readonly lastModified!: pulumi.Output<string>;
    /**
     * The result of the last AWS Lambda invocation of your Lambda function.
     */
    public /*out*/ readonly lastProcessingResult!: pulumi.Output<string>;
    /**
     * The maximum amount of time to gather records before invoking the function, in seconds (between 0 and 300). Records will continue to buffer (or accumulate in the case of an SQS queue event source) until either `maximumBatchingWindowInSeconds` expires or `batchSize` has been met. For streaming event sources, defaults to as soon as records are available in the stream. If the batch it reads from the stream/queue only has one record in it, Lambda only sends one record to the function. Only available for stream sources (DynamoDB and Kinesis) and SQS standard queues.
     * * `maximumRecordAgeInSeconds`: - (Optional) The maximum age of a record that Lambda sends to a function for processing. Only available for stream sources (DynamoDB and Kinesis). Must be either -1 (forever, and the default value) or between 60 and 604800 (inclusive).
     * * `maximumRetryAttempts`: - (Optional) The maximum number of times to retry when the function returns an error. Only available for stream sources (DynamoDB and Kinesis). Minimum and default of -1 (forever), maximum of 10000.
     * * `parallelizationFactor`: - (Optional) The number of batches to process from each shard concurrently. Only available for stream sources (DynamoDB and Kinesis). Minimum and default of 1, maximum of 10.
     */
    public readonly maximumBatchingWindowInSeconds!: pulumi.Output<number | undefined>;
    public readonly maximumRecordAgeInSeconds!: pulumi.Output<number>;
    public readonly maximumRetryAttempts!: pulumi.Output<number>;
    public readonly parallelizationFactor!: pulumi.Output<number>;
    /**
     * The name of the Amazon MQ broker destination queue to consume. Only available for MQ sources. A single queue name must be specified.
     * * `selfManagedEventSource`: - (Optional) For Self Managed Kafka sources, the location of the self managed cluster. If set, configuration must also include `sourceAccessConfiguration`. Detailed below.
     * * `sourceAccessConfiguration`: (Optional) For Self Managed Kafka sources, the access configuration for the source. If set, configuration must also include `selfManagedEventSource`. Detailed below.
     */
    public readonly queues!: pulumi.Output<string[] | undefined>;
    public readonly selfManagedEventSource!: pulumi.Output<outputs.lambda.EventSourceMappingSelfManagedEventSource | undefined>;
    public readonly sourceAccessConfigurations!: pulumi.Output<outputs.lambda.EventSourceMappingSourceAccessConfiguration[] | undefined>;
    /**
     * The position in the stream where AWS Lambda should start reading. Must be one of `AT_TIMESTAMP` (Kinesis only), `LATEST` or `TRIM_HORIZON` if getting events from Kinesis, DynamoDB or MSK. Must not be provided if getting events from SQS. More information about these positions can be found in the [AWS DynamoDB Streams API Reference](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_streams_GetShardIterator.html) and [AWS Kinesis API Reference](https://docs.aws.amazon.com/kinesis/latest/APIReference/API_GetShardIterator.html#Kinesis-GetShardIterator-request-ShardIteratorType).
     */
    public readonly startingPosition!: pulumi.Output<string | undefined>;
    /**
     * A timestamp in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) of the data record which to start reading when using `startingPosition` set to `AT_TIMESTAMP`. If a record with this exact timestamp does not exist, the next later record is chosen. If the timestamp is older than the current trim horizon, the oldest available record is chosen.
     */
    public readonly startingPositionTimestamp!: pulumi.Output<string | undefined>;
    /**
     * The state of the event source mapping.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * The reason the event source mapping is in its current state.
     */
    public /*out*/ readonly stateTransitionReason!: pulumi.Output<string>;
    /**
     * The name of the Kafka topics. Only available for MSK sources. A single topic name must be specified.
     */
    public readonly topics!: pulumi.Output<string[] | undefined>;
    /**
     * The duration in seconds of a processing window for [AWS Lambda streaming analytics](https://docs.aws.amazon.com/lambda/latest/dg/with-kinesis.html#services-kinesis-windows). The range is between 1 second up to 900 seconds. Only available for stream sources (DynamoDB and Kinesis).
     */
    public readonly tumblingWindowInSeconds!: pulumi.Output<number | undefined>;
    /**
     * The UUID of the created event source mapping.
     */
    public /*out*/ readonly uuid!: pulumi.Output<string>;

    /**
     * Create a EventSourceMapping resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EventSourceMappingArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: EventSourceMappingArgs | EventSourceMappingState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as EventSourceMappingState | undefined;
            resourceInputs["batchSize"] = state ? state.batchSize : undefined;
            resourceInputs["bisectBatchOnFunctionError"] = state ? state.bisectBatchOnFunctionError : undefined;
            resourceInputs["destinationConfig"] = state ? state.destinationConfig : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["eventSourceArn"] = state ? state.eventSourceArn : undefined;
            resourceInputs["filterCriteria"] = state ? state.filterCriteria : undefined;
            resourceInputs["functionArn"] = state ? state.functionArn : undefined;
            resourceInputs["functionName"] = state ? state.functionName : undefined;
            resourceInputs["functionResponseTypes"] = state ? state.functionResponseTypes : undefined;
            resourceInputs["lastModified"] = state ? state.lastModified : undefined;
            resourceInputs["lastProcessingResult"] = state ? state.lastProcessingResult : undefined;
            resourceInputs["maximumBatchingWindowInSeconds"] = state ? state.maximumBatchingWindowInSeconds : undefined;
            resourceInputs["maximumRecordAgeInSeconds"] = state ? state.maximumRecordAgeInSeconds : undefined;
            resourceInputs["maximumRetryAttempts"] = state ? state.maximumRetryAttempts : undefined;
            resourceInputs["parallelizationFactor"] = state ? state.parallelizationFactor : undefined;
            resourceInputs["queues"] = state ? state.queues : undefined;
            resourceInputs["selfManagedEventSource"] = state ? state.selfManagedEventSource : undefined;
            resourceInputs["sourceAccessConfigurations"] = state ? state.sourceAccessConfigurations : undefined;
            resourceInputs["startingPosition"] = state ? state.startingPosition : undefined;
            resourceInputs["startingPositionTimestamp"] = state ? state.startingPositionTimestamp : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["stateTransitionReason"] = state ? state.stateTransitionReason : undefined;
            resourceInputs["topics"] = state ? state.topics : undefined;
            resourceInputs["tumblingWindowInSeconds"] = state ? state.tumblingWindowInSeconds : undefined;
            resourceInputs["uuid"] = state ? state.uuid : undefined;
        } else {
            const args = argsOrState as EventSourceMappingArgs | undefined;
            if ((!args || args.functionName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'functionName'");
            }
            resourceInputs["batchSize"] = args ? args.batchSize : undefined;
            resourceInputs["bisectBatchOnFunctionError"] = args ? args.bisectBatchOnFunctionError : undefined;
            resourceInputs["destinationConfig"] = args ? args.destinationConfig : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["eventSourceArn"] = args ? args.eventSourceArn : undefined;
            resourceInputs["filterCriteria"] = args ? args.filterCriteria : undefined;
            resourceInputs["functionName"] = args ? args.functionName : undefined;
            resourceInputs["functionResponseTypes"] = args ? args.functionResponseTypes : undefined;
            resourceInputs["maximumBatchingWindowInSeconds"] = args ? args.maximumBatchingWindowInSeconds : undefined;
            resourceInputs["maximumRecordAgeInSeconds"] = args ? args.maximumRecordAgeInSeconds : undefined;
            resourceInputs["maximumRetryAttempts"] = args ? args.maximumRetryAttempts : undefined;
            resourceInputs["parallelizationFactor"] = args ? args.parallelizationFactor : undefined;
            resourceInputs["queues"] = args ? args.queues : undefined;
            resourceInputs["selfManagedEventSource"] = args ? args.selfManagedEventSource : undefined;
            resourceInputs["sourceAccessConfigurations"] = args ? args.sourceAccessConfigurations : undefined;
            resourceInputs["startingPosition"] = args ? args.startingPosition : undefined;
            resourceInputs["startingPositionTimestamp"] = args ? args.startingPositionTimestamp : undefined;
            resourceInputs["topics"] = args ? args.topics : undefined;
            resourceInputs["tumblingWindowInSeconds"] = args ? args.tumblingWindowInSeconds : undefined;
            resourceInputs["functionArn"] = undefined /*out*/;
            resourceInputs["lastModified"] = undefined /*out*/;
            resourceInputs["lastProcessingResult"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["stateTransitionReason"] = undefined /*out*/;
            resourceInputs["uuid"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(EventSourceMapping.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering EventSourceMapping resources.
 */
export interface EventSourceMappingState {
    /**
     * The largest number of records that Lambda will retrieve from your event source at the time of invocation. Defaults to `100` for DynamoDB, Kinesis, MQ and MSK, `10` for SQS.
     * * `bisectBatchOnFunctionError`: - (Optional) If the function returns an error, split the batch in two and retry. Only available for stream sources (DynamoDB and Kinesis). Defaults to `false`.
     * * `destinationConfig`: - (Optional) An Amazon SQS queue or Amazon SNS topic destination for failed records. Only available for stream sources (DynamoDB and Kinesis). Detailed below.
     */
    batchSize?: pulumi.Input<number>;
    bisectBatchOnFunctionError?: pulumi.Input<boolean>;
    destinationConfig?: pulumi.Input<inputs.lambda.EventSourceMappingDestinationConfig>;
    /**
     * Determines if the mapping will be enabled on creation. Defaults to `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The event source ARN - this is required for Kinesis stream, DynamoDB stream, SQS queue, MQ broker or MSK cluster.  It is incompatible with a Self Managed Kafka source.
     */
    eventSourceArn?: pulumi.Input<string>;
    /**
     * The criteria to use for [event filtering](https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html) Kinesis stream, DynamoDB stream, SQS queue event sources. Detailed below.
     */
    filterCriteria?: pulumi.Input<inputs.lambda.EventSourceMappingFilterCriteria>;
    /**
     * The the ARN of the Lambda function the event source mapping is sending events to. (Note: this is a computed value that differs from `functionName` above.)
     */
    functionArn?: pulumi.Input<string>;
    /**
     * The name or the ARN of the Lambda function that will be subscribing to events.
     */
    functionName?: pulumi.Input<string>;
    /**
     * A list of current response type enums applied to the event source mapping for [AWS Lambda checkpointing](https://docs.aws.amazon.com/lambda/latest/dg/with-ddb.html#services-ddb-batchfailurereporting). Only available for SQS and stream sources (DynamoDB and Kinesis). Valid values: `ReportBatchItemFailures`.
     */
    functionResponseTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The date this resource was last modified.
     */
    lastModified?: pulumi.Input<string>;
    /**
     * The result of the last AWS Lambda invocation of your Lambda function.
     */
    lastProcessingResult?: pulumi.Input<string>;
    /**
     * The maximum amount of time to gather records before invoking the function, in seconds (between 0 and 300). Records will continue to buffer (or accumulate in the case of an SQS queue event source) until either `maximumBatchingWindowInSeconds` expires or `batchSize` has been met. For streaming event sources, defaults to as soon as records are available in the stream. If the batch it reads from the stream/queue only has one record in it, Lambda only sends one record to the function. Only available for stream sources (DynamoDB and Kinesis) and SQS standard queues.
     * * `maximumRecordAgeInSeconds`: - (Optional) The maximum age of a record that Lambda sends to a function for processing. Only available for stream sources (DynamoDB and Kinesis). Must be either -1 (forever, and the default value) or between 60 and 604800 (inclusive).
     * * `maximumRetryAttempts`: - (Optional) The maximum number of times to retry when the function returns an error. Only available for stream sources (DynamoDB and Kinesis). Minimum and default of -1 (forever), maximum of 10000.
     * * `parallelizationFactor`: - (Optional) The number of batches to process from each shard concurrently. Only available for stream sources (DynamoDB and Kinesis). Minimum and default of 1, maximum of 10.
     */
    maximumBatchingWindowInSeconds?: pulumi.Input<number>;
    maximumRecordAgeInSeconds?: pulumi.Input<number>;
    maximumRetryAttempts?: pulumi.Input<number>;
    parallelizationFactor?: pulumi.Input<number>;
    /**
     * The name of the Amazon MQ broker destination queue to consume. Only available for MQ sources. A single queue name must be specified.
     * * `selfManagedEventSource`: - (Optional) For Self Managed Kafka sources, the location of the self managed cluster. If set, configuration must also include `sourceAccessConfiguration`. Detailed below.
     * * `sourceAccessConfiguration`: (Optional) For Self Managed Kafka sources, the access configuration for the source. If set, configuration must also include `selfManagedEventSource`. Detailed below.
     */
    queues?: pulumi.Input<pulumi.Input<string>[]>;
    selfManagedEventSource?: pulumi.Input<inputs.lambda.EventSourceMappingSelfManagedEventSource>;
    sourceAccessConfigurations?: pulumi.Input<pulumi.Input<inputs.lambda.EventSourceMappingSourceAccessConfiguration>[]>;
    /**
     * The position in the stream where AWS Lambda should start reading. Must be one of `AT_TIMESTAMP` (Kinesis only), `LATEST` or `TRIM_HORIZON` if getting events from Kinesis, DynamoDB or MSK. Must not be provided if getting events from SQS. More information about these positions can be found in the [AWS DynamoDB Streams API Reference](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_streams_GetShardIterator.html) and [AWS Kinesis API Reference](https://docs.aws.amazon.com/kinesis/latest/APIReference/API_GetShardIterator.html#Kinesis-GetShardIterator-request-ShardIteratorType).
     */
    startingPosition?: pulumi.Input<string>;
    /**
     * A timestamp in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) of the data record which to start reading when using `startingPosition` set to `AT_TIMESTAMP`. If a record with this exact timestamp does not exist, the next later record is chosen. If the timestamp is older than the current trim horizon, the oldest available record is chosen.
     */
    startingPositionTimestamp?: pulumi.Input<string>;
    /**
     * The state of the event source mapping.
     */
    state?: pulumi.Input<string>;
    /**
     * The reason the event source mapping is in its current state.
     */
    stateTransitionReason?: pulumi.Input<string>;
    /**
     * The name of the Kafka topics. Only available for MSK sources. A single topic name must be specified.
     */
    topics?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The duration in seconds of a processing window for [AWS Lambda streaming analytics](https://docs.aws.amazon.com/lambda/latest/dg/with-kinesis.html#services-kinesis-windows). The range is between 1 second up to 900 seconds. Only available for stream sources (DynamoDB and Kinesis).
     */
    tumblingWindowInSeconds?: pulumi.Input<number>;
    /**
     * The UUID of the created event source mapping.
     */
    uuid?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a EventSourceMapping resource.
 */
export interface EventSourceMappingArgs {
    /**
     * The largest number of records that Lambda will retrieve from your event source at the time of invocation. Defaults to `100` for DynamoDB, Kinesis, MQ and MSK, `10` for SQS.
     * * `bisectBatchOnFunctionError`: - (Optional) If the function returns an error, split the batch in two and retry. Only available for stream sources (DynamoDB and Kinesis). Defaults to `false`.
     * * `destinationConfig`: - (Optional) An Amazon SQS queue or Amazon SNS topic destination for failed records. Only available for stream sources (DynamoDB and Kinesis). Detailed below.
     */
    batchSize?: pulumi.Input<number>;
    bisectBatchOnFunctionError?: pulumi.Input<boolean>;
    destinationConfig?: pulumi.Input<inputs.lambda.EventSourceMappingDestinationConfig>;
    /**
     * Determines if the mapping will be enabled on creation. Defaults to `true`.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * The event source ARN - this is required for Kinesis stream, DynamoDB stream, SQS queue, MQ broker or MSK cluster.  It is incompatible with a Self Managed Kafka source.
     */
    eventSourceArn?: pulumi.Input<string>;
    /**
     * The criteria to use for [event filtering](https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html) Kinesis stream, DynamoDB stream, SQS queue event sources. Detailed below.
     */
    filterCriteria?: pulumi.Input<inputs.lambda.EventSourceMappingFilterCriteria>;
    /**
     * The name or the ARN of the Lambda function that will be subscribing to events.
     */
    functionName: pulumi.Input<string>;
    /**
     * A list of current response type enums applied to the event source mapping for [AWS Lambda checkpointing](https://docs.aws.amazon.com/lambda/latest/dg/with-ddb.html#services-ddb-batchfailurereporting). Only available for SQS and stream sources (DynamoDB and Kinesis). Valid values: `ReportBatchItemFailures`.
     */
    functionResponseTypes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The maximum amount of time to gather records before invoking the function, in seconds (between 0 and 300). Records will continue to buffer (or accumulate in the case of an SQS queue event source) until either `maximumBatchingWindowInSeconds` expires or `batchSize` has been met. For streaming event sources, defaults to as soon as records are available in the stream. If the batch it reads from the stream/queue only has one record in it, Lambda only sends one record to the function. Only available for stream sources (DynamoDB and Kinesis) and SQS standard queues.
     * * `maximumRecordAgeInSeconds`: - (Optional) The maximum age of a record that Lambda sends to a function for processing. Only available for stream sources (DynamoDB and Kinesis). Must be either -1 (forever, and the default value) or between 60 and 604800 (inclusive).
     * * `maximumRetryAttempts`: - (Optional) The maximum number of times to retry when the function returns an error. Only available for stream sources (DynamoDB and Kinesis). Minimum and default of -1 (forever), maximum of 10000.
     * * `parallelizationFactor`: - (Optional) The number of batches to process from each shard concurrently. Only available for stream sources (DynamoDB and Kinesis). Minimum and default of 1, maximum of 10.
     */
    maximumBatchingWindowInSeconds?: pulumi.Input<number>;
    maximumRecordAgeInSeconds?: pulumi.Input<number>;
    maximumRetryAttempts?: pulumi.Input<number>;
    parallelizationFactor?: pulumi.Input<number>;
    /**
     * The name of the Amazon MQ broker destination queue to consume. Only available for MQ sources. A single queue name must be specified.
     * * `selfManagedEventSource`: - (Optional) For Self Managed Kafka sources, the location of the self managed cluster. If set, configuration must also include `sourceAccessConfiguration`. Detailed below.
     * * `sourceAccessConfiguration`: (Optional) For Self Managed Kafka sources, the access configuration for the source. If set, configuration must also include `selfManagedEventSource`. Detailed below.
     */
    queues?: pulumi.Input<pulumi.Input<string>[]>;
    selfManagedEventSource?: pulumi.Input<inputs.lambda.EventSourceMappingSelfManagedEventSource>;
    sourceAccessConfigurations?: pulumi.Input<pulumi.Input<inputs.lambda.EventSourceMappingSourceAccessConfiguration>[]>;
    /**
     * The position in the stream where AWS Lambda should start reading. Must be one of `AT_TIMESTAMP` (Kinesis only), `LATEST` or `TRIM_HORIZON` if getting events from Kinesis, DynamoDB or MSK. Must not be provided if getting events from SQS. More information about these positions can be found in the [AWS DynamoDB Streams API Reference](https://docs.aws.amazon.com/amazondynamodb/latest/APIReference/API_streams_GetShardIterator.html) and [AWS Kinesis API Reference](https://docs.aws.amazon.com/kinesis/latest/APIReference/API_GetShardIterator.html#Kinesis-GetShardIterator-request-ShardIteratorType).
     */
    startingPosition?: pulumi.Input<string>;
    /**
     * A timestamp in [RFC3339 format](https://tools.ietf.org/html/rfc3339#section-5.8) of the data record which to start reading when using `startingPosition` set to `AT_TIMESTAMP`. If a record with this exact timestamp does not exist, the next later record is chosen. If the timestamp is older than the current trim horizon, the oldest available record is chosen.
     */
    startingPositionTimestamp?: pulumi.Input<string>;
    /**
     * The name of the Kafka topics. Only available for MSK sources. A single topic name must be specified.
     */
    topics?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The duration in seconds of a processing window for [AWS Lambda streaming analytics](https://docs.aws.amazon.com/lambda/latest/dg/with-kinesis.html#services-kinesis-windows). The range is between 1 second up to 900 seconds. Only available for stream sources (DynamoDB and Kinesis).
     */
    tumblingWindowInSeconds?: pulumi.Input<number>;
}
