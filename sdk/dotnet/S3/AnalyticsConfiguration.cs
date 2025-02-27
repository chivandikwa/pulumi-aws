// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.S3
{
    /// <summary>
    /// Provides a S3 bucket [analytics configuration](https://docs.aws.amazon.com/AmazonS3/latest/dev/analytics-storage-class.html) resource.
    /// 
    /// ## Example Usage
    /// ### Add analytics configuration for entire S3 bucket and export results to a second S3 bucket
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.S3.BucketV2("example", new Aws.S3.BucketV2Args
    ///         {
    ///         });
    ///         var analytics = new Aws.S3.BucketV2("analytics", new Aws.S3.BucketV2Args
    ///         {
    ///         });
    ///         var example_entire_bucket = new Aws.S3.AnalyticsConfiguration("example-entire-bucket", new Aws.S3.AnalyticsConfigurationArgs
    ///         {
    ///             Bucket = example.Bucket,
    ///             StorageClassAnalysis = new Aws.S3.Inputs.AnalyticsConfigurationStorageClassAnalysisArgs
    ///             {
    ///                 DataExport = new Aws.S3.Inputs.AnalyticsConfigurationStorageClassAnalysisDataExportArgs
    ///                 {
    ///                     Destination = new Aws.S3.Inputs.AnalyticsConfigurationStorageClassAnalysisDataExportDestinationArgs
    ///                     {
    ///                         S3BucketDestination = new Aws.S3.Inputs.AnalyticsConfigurationStorageClassAnalysisDataExportDestinationS3BucketDestinationArgs
    ///                         {
    ///                             BucketArn = analytics.Arn,
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// ### Add analytics configuration with S3 object filter
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.S3.BucketV2("example", new Aws.S3.BucketV2Args
    ///         {
    ///         });
    ///         var example_filtered = new Aws.S3.AnalyticsConfiguration("example-filtered", new Aws.S3.AnalyticsConfigurationArgs
    ///         {
    ///             Bucket = example.Bucket,
    ///             Filter = new Aws.S3.Inputs.AnalyticsConfigurationFilterArgs
    ///             {
    ///                 Prefix = "documents/",
    ///                 Tags = 
    ///                 {
    ///                     { "priority", "high" },
    ///                     { "class", "blue" },
    ///                 },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// S3 bucket analytics configurations can be imported using `bucket:analytics`, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:s3/analyticsConfiguration:AnalyticsConfiguration my-bucket-entire-bucket my-bucket:EntireBucket
    /// ```
    /// </summary>
    [AwsResourceType("aws:s3/analyticsConfiguration:AnalyticsConfiguration")]
    public partial class AnalyticsConfiguration : Pulumi.CustomResource
    {
        /// <summary>
        /// The name of the bucket this analytics configuration is associated with.
        /// </summary>
        [Output("bucket")]
        public Output<string> Bucket { get; private set; } = null!;

        /// <summary>
        /// Object filtering that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
        /// </summary>
        [Output("filter")]
        public Output<Outputs.AnalyticsConfigurationFilter?> Filter { get; private set; } = null!;

        /// <summary>
        /// Unique identifier of the analytics configuration for the bucket.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// Configuration for the analytics data export (documented below).
        /// </summary>
        [Output("storageClassAnalysis")]
        public Output<Outputs.AnalyticsConfigurationStorageClassAnalysis?> StorageClassAnalysis { get; private set; } = null!;


        /// <summary>
        /// Create a AnalyticsConfiguration resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AnalyticsConfiguration(string name, AnalyticsConfigurationArgs args, CustomResourceOptions? options = null)
            : base("aws:s3/analyticsConfiguration:AnalyticsConfiguration", name, args ?? new AnalyticsConfigurationArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AnalyticsConfiguration(string name, Input<string> id, AnalyticsConfigurationState? state = null, CustomResourceOptions? options = null)
            : base("aws:s3/analyticsConfiguration:AnalyticsConfiguration", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing AnalyticsConfiguration resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AnalyticsConfiguration Get(string name, Input<string> id, AnalyticsConfigurationState? state = null, CustomResourceOptions? options = null)
        {
            return new AnalyticsConfiguration(name, id, state, options);
        }
    }

    public sealed class AnalyticsConfigurationArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the bucket this analytics configuration is associated with.
        /// </summary>
        [Input("bucket", required: true)]
        public Input<string> Bucket { get; set; } = null!;

        /// <summary>
        /// Object filtering that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
        /// </summary>
        [Input("filter")]
        public Input<Inputs.AnalyticsConfigurationFilterArgs>? Filter { get; set; }

        /// <summary>
        /// Unique identifier of the analytics configuration for the bucket.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Configuration for the analytics data export (documented below).
        /// </summary>
        [Input("storageClassAnalysis")]
        public Input<Inputs.AnalyticsConfigurationStorageClassAnalysisArgs>? StorageClassAnalysis { get; set; }

        public AnalyticsConfigurationArgs()
        {
        }
    }

    public sealed class AnalyticsConfigurationState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The name of the bucket this analytics configuration is associated with.
        /// </summary>
        [Input("bucket")]
        public Input<string>? Bucket { get; set; }

        /// <summary>
        /// Object filtering that accepts a prefix, tags, or a logical AND of prefix and tags (documented below).
        /// </summary>
        [Input("filter")]
        public Input<Inputs.AnalyticsConfigurationFilterGetArgs>? Filter { get; set; }

        /// <summary>
        /// Unique identifier of the analytics configuration for the bucket.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Configuration for the analytics data export (documented below).
        /// </summary>
        [Input("storageClassAnalysis")]
        public Input<Inputs.AnalyticsConfigurationStorageClassAnalysisGetArgs>? StorageClassAnalysis { get; set; }

        public AnalyticsConfigurationState()
        {
        }
    }
}
