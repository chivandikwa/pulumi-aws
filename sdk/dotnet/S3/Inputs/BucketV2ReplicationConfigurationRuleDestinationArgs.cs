// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.S3.Inputs
{

    public sealed class BucketV2ReplicationConfigurationRuleDestinationArgs : Pulumi.ResourceArgs
    {
        [Input("accessControlTranslations")]
        private InputList<Inputs.BucketV2ReplicationConfigurationRuleDestinationAccessControlTranslationArgs>? _accessControlTranslations;

        /// <summary>
        /// The overrides to use for object owners on replication.
        /// </summary>
        [Obsolete(@"Use the aws_s3_bucket_replication_configuration resource instead")]
        public InputList<Inputs.BucketV2ReplicationConfigurationRuleDestinationAccessControlTranslationArgs> AccessControlTranslations
        {
            get => _accessControlTranslations ?? (_accessControlTranslations = new InputList<Inputs.BucketV2ReplicationConfigurationRuleDestinationAccessControlTranslationArgs>());
            set => _accessControlTranslations = value;
        }

        /// <summary>
        /// The Account ID to use for overriding the object owner on replication.
        /// </summary>
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        /// <summary>
        /// The name of the bucket. If omitted, this provider will assign a random, unique name. Must be lowercase and less than or equal to 63 characters in length. A full list of bucket naming rules [may be found here](https://docs.aws.amazon.com/AmazonS3/latest/userguide/bucketnamingrules.html).
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        [Input("metrics")]
        private InputList<Inputs.BucketV2ReplicationConfigurationRuleDestinationMetricArgs>? _metrics;

        /// <summary>
        /// Replication metrics.
        /// </summary>
        [Obsolete(@"Use the aws_s3_bucket_replication_configuration resource instead")]
        public InputList<Inputs.BucketV2ReplicationConfigurationRuleDestinationMetricArgs> Metrics
        {
            get => _metrics ?? (_metrics = new InputList<Inputs.BucketV2ReplicationConfigurationRuleDestinationMetricArgs>());
            set => _metrics = value;
        }

        /// <summary>
        /// Destination KMS encryption key ARN for SSE-KMS replication.
        /// </summary>
        [Input("replicaKmsKeyId")]
        public Input<string>? ReplicaKmsKeyId { get; set; }

        [Input("replicationTimes")]
        private InputList<Inputs.BucketV2ReplicationConfigurationRuleDestinationReplicationTimeArgs>? _replicationTimes;

        /// <summary>
        /// S3 Replication Time Control (S3 RTC).
        /// </summary>
        [Obsolete(@"Use the aws_s3_bucket_replication_configuration resource instead")]
        public InputList<Inputs.BucketV2ReplicationConfigurationRuleDestinationReplicationTimeArgs> ReplicationTimes
        {
            get => _replicationTimes ?? (_replicationTimes = new InputList<Inputs.BucketV2ReplicationConfigurationRuleDestinationReplicationTimeArgs>());
            set => _replicationTimes = value;
        }

        /// <summary>
        /// The [storage class](https://docs.aws.amazon.com/AmazonS3/latest/API/API_Destination.html#AmazonS3-Type-Destination-StorageClass) used to store the object.
        /// </summary>
        [Input("storageClass")]
        public Input<string>? StorageClass { get; set; }

        public BucketV2ReplicationConfigurationRuleDestinationArgs()
        {
        }
    }
}
