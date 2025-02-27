// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.S3.Inputs
{

    public sealed class BucketV2ReplicationConfigurationGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The ARN of the IAM role for Amazon S3 assumed when replicating the objects.
        /// </summary>
        [Input("role")]
        public Input<string>? Role { get; set; }

        [Input("rules")]
        private InputList<Inputs.BucketV2ReplicationConfigurationRuleGetArgs>? _rules;

        /// <summary>
        /// The rules managing the replication.
        /// </summary>
        [Obsolete(@"Use the aws_s3_bucket_replication_configuration resource instead")]
        public InputList<Inputs.BucketV2ReplicationConfigurationRuleGetArgs> Rules
        {
            get => _rules ?? (_rules = new InputList<Inputs.BucketV2ReplicationConfigurationRuleGetArgs>());
            set => _rules = value;
        }

        public BucketV2ReplicationConfigurationGetArgs()
        {
        }
    }
}
