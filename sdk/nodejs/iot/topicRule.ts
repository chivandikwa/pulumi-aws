// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const mytopic = new aws.sns.Topic("mytopic", {});
 * const myerrortopic = new aws.sns.Topic("myerrortopic", {});
 * const role = new aws.iam.Role("role", {assumeRolePolicy: `{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *       "Effect": "Allow",
 *       "Principal": {
 *         "Service": "iot.amazonaws.com"
 *       },
 *       "Action": "sts:AssumeRole"
 *     }
 *   ]
 * }
 * `});
 * const rule = new aws.iot.TopicRule("rule", {
 *     description: "Example rule",
 *     enabled: true,
 *     sql: "SELECT * FROM 'topic/test'",
 *     sqlVersion: "2016-03-23",
 *     sns: {
 *         messageFormat: "RAW",
 *         roleArn: role.arn,
 *         targetArn: mytopic.arn,
 *     },
 *     errorAction: {
 *         sns: {
 *             messageFormat: "RAW",
 *             roleArn: role.arn,
 *             targetArn: myerrortopic.arn,
 *         },
 *     },
 * });
 * const iamPolicyForLambda = new aws.iam.RolePolicy("iamPolicyForLambda", {
 *     role: role.id,
 *     policy: pulumi.interpolate`{
 *   "Version": "2012-10-17",
 *   "Statement": [
 *     {
 *         "Effect": "Allow",
 *         "Action": [
 *             "sns:Publish"
 *         ],
 *         "Resource": "${mytopic.arn}"
 *     }
 *   ]
 * }
 * `,
 * });
 * ```
 *
 * ## Import
 *
 * IoT Topic Rules can be imported using the `name`, e.g.,
 *
 * ```sh
 *  $ pulumi import aws:iot/topicRule:TopicRule rule <name>
 * ```
 */
export class TopicRule extends pulumi.CustomResource {
    /**
     * Get an existing TopicRule resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: TopicRuleState, opts?: pulumi.CustomResourceOptions): TopicRule {
        return new TopicRule(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'aws:iot/topicRule:TopicRule';

    /**
     * Returns true if the given object is an instance of TopicRule.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is TopicRule {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === TopicRule.__pulumiType;
    }

    /**
     * The ARN of the topic rule
     */
    public /*out*/ readonly arn!: pulumi.Output<string>;
    public readonly cloudwatchAlarm!: pulumi.Output<outputs.iot.TopicRuleCloudwatchAlarm | undefined>;
    public readonly cloudwatchMetric!: pulumi.Output<outputs.iot.TopicRuleCloudwatchMetric | undefined>;
    /**
     * The description of the rule.
     */
    public readonly description!: pulumi.Output<string | undefined>;
    public readonly dynamodb!: pulumi.Output<outputs.iot.TopicRuleDynamodb | undefined>;
    public readonly dynamodbv2s!: pulumi.Output<outputs.iot.TopicRuleDynamodbv2[] | undefined>;
    public readonly elasticsearch!: pulumi.Output<outputs.iot.TopicRuleElasticsearch | undefined>;
    /**
     * Specifies whether the rule is enabled.
     */
    public readonly enabled!: pulumi.Output<boolean>;
    /**
     * Configuration block with error action to be associated with the rule. See the documentation for `cloudwatchAlarm`, `cloudwatchMetric`, `dynamodb`, `dynamodbv2`, `elasticsearch`, `firehose`, `iotAnalytics`, `iotEvents`, `kinesis`, `lambda`, `republish`, `s3`, `stepFunctions`, `sns`, `sqs` configuration blocks for further configuration details.
     */
    public readonly errorAction!: pulumi.Output<outputs.iot.TopicRuleErrorAction | undefined>;
    public readonly firehose!: pulumi.Output<outputs.iot.TopicRuleFirehose | undefined>;
    public readonly iotAnalytics!: pulumi.Output<outputs.iot.TopicRuleIotAnalytic[] | undefined>;
    public readonly iotEvents!: pulumi.Output<outputs.iot.TopicRuleIotEvent[] | undefined>;
    public readonly kinesis!: pulumi.Output<outputs.iot.TopicRuleKinesis | undefined>;
    public readonly lambda!: pulumi.Output<outputs.iot.TopicRuleLambda | undefined>;
    /**
     * The name of the rule.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly republish!: pulumi.Output<outputs.iot.TopicRuleRepublish | undefined>;
    public readonly s3!: pulumi.Output<outputs.iot.TopicRuleS3 | undefined>;
    public readonly sns!: pulumi.Output<outputs.iot.TopicRuleSns | undefined>;
    /**
     * The SQL statement used to query the topic. For more information, see AWS IoT SQL Reference (http://docs.aws.amazon.com/iot/latest/developerguide/iot-rules.html#aws-iot-sql-reference) in the AWS IoT Developer Guide.
     */
    public readonly sql!: pulumi.Output<string>;
    /**
     * The version of the SQL rules engine to use when evaluating the rule.
     */
    public readonly sqlVersion!: pulumi.Output<string>;
    public readonly sqs!: pulumi.Output<outputs.iot.TopicRuleSqs | undefined>;
    public readonly stepFunctions!: pulumi.Output<outputs.iot.TopicRuleStepFunction[] | undefined>;
    /**
     * Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    public readonly tags!: pulumi.Output<{[key: string]: string} | undefined>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    public /*out*/ readonly tagsAll!: pulumi.Output<{[key: string]: string}>;

    /**
     * Create a TopicRule resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: TopicRuleArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: TopicRuleArgs | TopicRuleState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as TopicRuleState | undefined;
            resourceInputs["arn"] = state ? state.arn : undefined;
            resourceInputs["cloudwatchAlarm"] = state ? state.cloudwatchAlarm : undefined;
            resourceInputs["cloudwatchMetric"] = state ? state.cloudwatchMetric : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["dynamodb"] = state ? state.dynamodb : undefined;
            resourceInputs["dynamodbv2s"] = state ? state.dynamodbv2s : undefined;
            resourceInputs["elasticsearch"] = state ? state.elasticsearch : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["errorAction"] = state ? state.errorAction : undefined;
            resourceInputs["firehose"] = state ? state.firehose : undefined;
            resourceInputs["iotAnalytics"] = state ? state.iotAnalytics : undefined;
            resourceInputs["iotEvents"] = state ? state.iotEvents : undefined;
            resourceInputs["kinesis"] = state ? state.kinesis : undefined;
            resourceInputs["lambda"] = state ? state.lambda : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["republish"] = state ? state.republish : undefined;
            resourceInputs["s3"] = state ? state.s3 : undefined;
            resourceInputs["sns"] = state ? state.sns : undefined;
            resourceInputs["sql"] = state ? state.sql : undefined;
            resourceInputs["sqlVersion"] = state ? state.sqlVersion : undefined;
            resourceInputs["sqs"] = state ? state.sqs : undefined;
            resourceInputs["stepFunctions"] = state ? state.stepFunctions : undefined;
            resourceInputs["tags"] = state ? state.tags : undefined;
            resourceInputs["tagsAll"] = state ? state.tagsAll : undefined;
        } else {
            const args = argsOrState as TopicRuleArgs | undefined;
            if ((!args || args.enabled === undefined) && !opts.urn) {
                throw new Error("Missing required property 'enabled'");
            }
            if ((!args || args.sql === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sql'");
            }
            if ((!args || args.sqlVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'sqlVersion'");
            }
            resourceInputs["cloudwatchAlarm"] = args ? args.cloudwatchAlarm : undefined;
            resourceInputs["cloudwatchMetric"] = args ? args.cloudwatchMetric : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dynamodb"] = args ? args.dynamodb : undefined;
            resourceInputs["dynamodbv2s"] = args ? args.dynamodbv2s : undefined;
            resourceInputs["elasticsearch"] = args ? args.elasticsearch : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["errorAction"] = args ? args.errorAction : undefined;
            resourceInputs["firehose"] = args ? args.firehose : undefined;
            resourceInputs["iotAnalytics"] = args ? args.iotAnalytics : undefined;
            resourceInputs["iotEvents"] = args ? args.iotEvents : undefined;
            resourceInputs["kinesis"] = args ? args.kinesis : undefined;
            resourceInputs["lambda"] = args ? args.lambda : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["republish"] = args ? args.republish : undefined;
            resourceInputs["s3"] = args ? args.s3 : undefined;
            resourceInputs["sns"] = args ? args.sns : undefined;
            resourceInputs["sql"] = args ? args.sql : undefined;
            resourceInputs["sqlVersion"] = args ? args.sqlVersion : undefined;
            resourceInputs["sqs"] = args ? args.sqs : undefined;
            resourceInputs["stepFunctions"] = args ? args.stepFunctions : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["arn"] = undefined /*out*/;
            resourceInputs["tagsAll"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(TopicRule.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering TopicRule resources.
 */
export interface TopicRuleState {
    /**
     * The ARN of the topic rule
     */
    arn?: pulumi.Input<string>;
    cloudwatchAlarm?: pulumi.Input<inputs.iot.TopicRuleCloudwatchAlarm>;
    cloudwatchMetric?: pulumi.Input<inputs.iot.TopicRuleCloudwatchMetric>;
    /**
     * The description of the rule.
     */
    description?: pulumi.Input<string>;
    dynamodb?: pulumi.Input<inputs.iot.TopicRuleDynamodb>;
    dynamodbv2s?: pulumi.Input<pulumi.Input<inputs.iot.TopicRuleDynamodbv2>[]>;
    elasticsearch?: pulumi.Input<inputs.iot.TopicRuleElasticsearch>;
    /**
     * Specifies whether the rule is enabled.
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Configuration block with error action to be associated with the rule. See the documentation for `cloudwatchAlarm`, `cloudwatchMetric`, `dynamodb`, `dynamodbv2`, `elasticsearch`, `firehose`, `iotAnalytics`, `iotEvents`, `kinesis`, `lambda`, `republish`, `s3`, `stepFunctions`, `sns`, `sqs` configuration blocks for further configuration details.
     */
    errorAction?: pulumi.Input<inputs.iot.TopicRuleErrorAction>;
    firehose?: pulumi.Input<inputs.iot.TopicRuleFirehose>;
    iotAnalytics?: pulumi.Input<pulumi.Input<inputs.iot.TopicRuleIotAnalytic>[]>;
    iotEvents?: pulumi.Input<pulumi.Input<inputs.iot.TopicRuleIotEvent>[]>;
    kinesis?: pulumi.Input<inputs.iot.TopicRuleKinesis>;
    lambda?: pulumi.Input<inputs.iot.TopicRuleLambda>;
    /**
     * The name of the rule.
     */
    name?: pulumi.Input<string>;
    republish?: pulumi.Input<inputs.iot.TopicRuleRepublish>;
    s3?: pulumi.Input<inputs.iot.TopicRuleS3>;
    sns?: pulumi.Input<inputs.iot.TopicRuleSns>;
    /**
     * The SQL statement used to query the topic. For more information, see AWS IoT SQL Reference (http://docs.aws.amazon.com/iot/latest/developerguide/iot-rules.html#aws-iot-sql-reference) in the AWS IoT Developer Guide.
     */
    sql?: pulumi.Input<string>;
    /**
     * The version of the SQL rules engine to use when evaluating the rule.
     */
    sqlVersion?: pulumi.Input<string>;
    sqs?: pulumi.Input<inputs.iot.TopicRuleSqs>;
    stepFunctions?: pulumi.Input<pulumi.Input<inputs.iot.TopicRuleStepFunction>[]>;
    /**
     * Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A map of tags assigned to the resource, including those inherited from the provider .
     */
    tagsAll?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}

/**
 * The set of arguments for constructing a TopicRule resource.
 */
export interface TopicRuleArgs {
    cloudwatchAlarm?: pulumi.Input<inputs.iot.TopicRuleCloudwatchAlarm>;
    cloudwatchMetric?: pulumi.Input<inputs.iot.TopicRuleCloudwatchMetric>;
    /**
     * The description of the rule.
     */
    description?: pulumi.Input<string>;
    dynamodb?: pulumi.Input<inputs.iot.TopicRuleDynamodb>;
    dynamodbv2s?: pulumi.Input<pulumi.Input<inputs.iot.TopicRuleDynamodbv2>[]>;
    elasticsearch?: pulumi.Input<inputs.iot.TopicRuleElasticsearch>;
    /**
     * Specifies whether the rule is enabled.
     */
    enabled: pulumi.Input<boolean>;
    /**
     * Configuration block with error action to be associated with the rule. See the documentation for `cloudwatchAlarm`, `cloudwatchMetric`, `dynamodb`, `dynamodbv2`, `elasticsearch`, `firehose`, `iotAnalytics`, `iotEvents`, `kinesis`, `lambda`, `republish`, `s3`, `stepFunctions`, `sns`, `sqs` configuration blocks for further configuration details.
     */
    errorAction?: pulumi.Input<inputs.iot.TopicRuleErrorAction>;
    firehose?: pulumi.Input<inputs.iot.TopicRuleFirehose>;
    iotAnalytics?: pulumi.Input<pulumi.Input<inputs.iot.TopicRuleIotAnalytic>[]>;
    iotEvents?: pulumi.Input<pulumi.Input<inputs.iot.TopicRuleIotEvent>[]>;
    kinesis?: pulumi.Input<inputs.iot.TopicRuleKinesis>;
    lambda?: pulumi.Input<inputs.iot.TopicRuleLambda>;
    /**
     * The name of the rule.
     */
    name?: pulumi.Input<string>;
    republish?: pulumi.Input<inputs.iot.TopicRuleRepublish>;
    s3?: pulumi.Input<inputs.iot.TopicRuleS3>;
    sns?: pulumi.Input<inputs.iot.TopicRuleSns>;
    /**
     * The SQL statement used to query the topic. For more information, see AWS IoT SQL Reference (http://docs.aws.amazon.com/iot/latest/developerguide/iot-rules.html#aws-iot-sql-reference) in the AWS IoT Developer Guide.
     */
    sql: pulumi.Input<string>;
    /**
     * The version of the SQL rules engine to use when evaluating the rule.
     */
    sqlVersion: pulumi.Input<string>;
    sqs?: pulumi.Input<inputs.iot.TopicRuleSqs>;
    stepFunctions?: pulumi.Input<pulumi.Input<inputs.iot.TopicRuleStepFunction>[]>;
    /**
     * Key-value map of resource tags. .If configured with a provider `defaultTags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
