// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.S3.Outputs
{

    [OutputType]
    public sealed class BucketV2LifecycleRuleTransition
    {
        /// <summary>
        /// The date after which you want the corresponding action to take effect.
        /// </summary>
        public readonly string? Date;
        /// <summary>
        /// The number of days specified for the default retention period.
        /// </summary>
        public readonly int? Days;
        /// <summary>
        /// The [storage class](https://docs.aws.amazon.com/AmazonS3/latest/API/API_Destination.html#AmazonS3-Type-Destination-StorageClass) used to store the object.
        /// </summary>
        public readonly string? StorageClass;

        [OutputConstructor]
        private BucketV2LifecycleRuleTransition(
            string? date,

            int? days,

            string? storageClass)
        {
            Date = date;
            Days = days;
            StorageClass = storageClass;
        }
    }
}
