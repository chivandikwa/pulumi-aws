// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Cfg
{
    /// <summary>
    /// Manages status (recording / stopped) of an AWS Config Configuration Recorder.
    /// 
    /// &gt; **Note:** Starting Configuration Recorder requires a `Delivery Channel` to be present. Use of `depends_on` (as shown below) is recommended to avoid race conditions.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var bucketV2 = new Aws.S3.BucketV2("bucketV2", new Aws.S3.BucketV2Args
    ///         {
    ///         });
    ///         var fooDeliveryChannel = new Aws.Cfg.DeliveryChannel("fooDeliveryChannel", new Aws.Cfg.DeliveryChannelArgs
    ///         {
    ///             S3BucketName = bucketV2.Bucket,
    ///         });
    ///         var fooRecorderStatus = new Aws.Cfg.RecorderStatus("fooRecorderStatus", new Aws.Cfg.RecorderStatusArgs
    ///         {
    ///             IsEnabled = true,
    ///         }, new CustomResourceOptions
    ///         {
    ///             DependsOn = 
    ///             {
    ///                 fooDeliveryChannel,
    ///             },
    ///         });
    ///         var role = new Aws.Iam.Role("role", new Aws.Iam.RoleArgs
    ///         {
    ///             AssumeRolePolicy = @"{
    ///   ""Version"": ""2012-10-17"",
    ///   ""Statement"": [
    ///     {
    ///       ""Action"": ""sts:AssumeRole"",
    ///       ""Principal"": {
    ///         ""Service"": ""config.amazonaws.com""
    ///       },
    ///       ""Effect"": ""Allow"",
    ///       ""Sid"": """"
    ///     }
    ///   ]
    /// }
    /// ",
    ///         });
    ///         var rolePolicyAttachment = new Aws.Iam.RolePolicyAttachment("rolePolicyAttachment", new Aws.Iam.RolePolicyAttachmentArgs
    ///         {
    ///             Role = role.Name,
    ///             PolicyArn = "arn:aws:iam::aws:policy/service-role/AWSConfigRole",
    ///         });
    ///         var fooRecorder = new Aws.Cfg.Recorder("fooRecorder", new Aws.Cfg.RecorderArgs
    ///         {
    ///             RoleArn = role.Arn,
    ///         });
    ///         var rolePolicy = new Aws.Iam.RolePolicy("rolePolicy", new Aws.Iam.RolePolicyArgs
    ///         {
    ///             Role = role.Id,
    ///             Policy = Output.Tuple(bucketV2.Arn, bucketV2.Arn).Apply(values =&gt;
    ///             {
    ///                 var bucketV2Arn = values.Item1;
    ///                 var bucketV2Arn1 = values.Item2;
    ///                 return @$"{{
    ///   ""Version"": ""2012-10-17"",
    ///   ""Statement"": [
    ///     {{
    ///       ""Action"": [
    ///         ""s3:*""
    ///       ],
    ///       ""Effect"": ""Allow"",
    ///       ""Resource"": [
    ///         ""{bucketV2Arn}"",
    ///         ""{bucketV2Arn1}/*""
    ///       ]
    ///     }}
    ///   ]
    /// }}
    /// ";
    ///             }),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Configuration Recorder Status can be imported using the name of the Configuration Recorder, e.g.,
    /// 
    /// ```sh
    ///  $ pulumi import aws:cfg/recorderStatus:RecorderStatus foo example
    /// ```
    /// </summary>
    [AwsResourceType("aws:cfg/recorderStatus:RecorderStatus")]
    public partial class RecorderStatus : Pulumi.CustomResource
    {
        /// <summary>
        /// Whether the configuration recorder should be enabled or disabled.
        /// </summary>
        [Output("isEnabled")]
        public Output<bool> IsEnabled { get; private set; } = null!;

        /// <summary>
        /// The name of the recorder
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;


        /// <summary>
        /// Create a RecorderStatus resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public RecorderStatus(string name, RecorderStatusArgs args, CustomResourceOptions? options = null)
            : base("aws:cfg/recorderStatus:RecorderStatus", name, args ?? new RecorderStatusArgs(), MakeResourceOptions(options, ""))
        {
        }

        private RecorderStatus(string name, Input<string> id, RecorderStatusState? state = null, CustomResourceOptions? options = null)
            : base("aws:cfg/recorderStatus:RecorderStatus", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing RecorderStatus resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static RecorderStatus Get(string name, Input<string> id, RecorderStatusState? state = null, CustomResourceOptions? options = null)
        {
            return new RecorderStatus(name, id, state, options);
        }
    }

    public sealed class RecorderStatusArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the configuration recorder should be enabled or disabled.
        /// </summary>
        [Input("isEnabled", required: true)]
        public Input<bool> IsEnabled { get; set; } = null!;

        /// <summary>
        /// The name of the recorder
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public RecorderStatusArgs()
        {
        }
    }

    public sealed class RecorderStatusState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Whether the configuration recorder should be enabled or disabled.
        /// </summary>
        [Input("isEnabled")]
        public Input<bool>? IsEnabled { get; set; }

        /// <summary>
        /// The name of the recorder
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public RecorderStatusState()
        {
        }
    }
}
