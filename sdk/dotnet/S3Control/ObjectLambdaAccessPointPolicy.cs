// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.S3Control
{
    /// <summary>
    /// Provides a resource to manage an S3 Object Lambda Access Point resource policy.
    /// 
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var exampleBucketV2 = new Aws.S3.BucketV2("exampleBucketV2", new Aws.S3.BucketV2Args
    ///         {
    ///         });
    ///         var exampleAccessPoint = new Aws.S3.AccessPoint("exampleAccessPoint", new Aws.S3.AccessPointArgs
    ///         {
    ///             Bucket = exampleBucketV2.Id,
    ///         });
    ///         var exampleObjectLambdaAccessPoint = new Aws.S3Control.ObjectLambdaAccessPoint("exampleObjectLambdaAccessPoint", new Aws.S3Control.ObjectLambdaAccessPointArgs
    ///         {
    ///             Configuration = new Aws.S3Control.Inputs.ObjectLambdaAccessPointConfigurationArgs
    ///             {
    ///                 SupportingAccessPoint = exampleAccessPoint.Arn,
    ///                 TransformationConfigurations = 
    ///                 {
    ///                     new Aws.S3Control.Inputs.ObjectLambdaAccessPointConfigurationTransformationConfigurationArgs
    ///                     {
    ///                         Actions = 
    ///                         {
    ///                             "GetObject",
    ///                         },
    ///                         ContentTransformation = new Aws.S3Control.Inputs.ObjectLambdaAccessPointConfigurationTransformationConfigurationContentTransformationArgs
    ///                         {
    ///                             AwsLambda = new Aws.S3Control.Inputs.ObjectLambdaAccessPointConfigurationTransformationConfigurationContentTransformationAwsLambdaArgs
    ///                             {
    ///                                 FunctionArn = aws_lambda_function.Example.Arn,
    ///                             },
    ///                         },
    ///                     },
    ///                 },
    ///             },
    ///         });
    ///         var exampleObjectLambdaAccessPointPolicy = new Aws.S3Control.ObjectLambdaAccessPointPolicy("exampleObjectLambdaAccessPointPolicy", new Aws.S3Control.ObjectLambdaAccessPointPolicyArgs
    ///         {
    ///             Policy = exampleObjectLambdaAccessPoint.Arn.Apply(arn =&gt; JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///             {
    ///                 { "Version", "2008-10-17" },
    ///                 { "Statement", new[]
    ///                     {
    ///                         new Dictionary&lt;string, object?&gt;
    ///                         {
    ///                             { "Effect", "Allow" },
    ///                             { "Action", "s3-object-lambda:GetObject" },
    ///                             { "Principal", new Dictionary&lt;string, object?&gt;
    ///                             {
    ///                                 { "AWS", data.Aws_caller_identity.Current.Account_id },
    ///                             } },
    ///                             { "Resource", arn },
    ///                         },
    ///                     }
    ///                  },
    ///             })),
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// 
    /// ## Import
    /// 
    /// Object Lambda Access Point policies can be imported using the `account_id` and `name`, separated by a colon (`:`), e.g.
    /// 
    /// ```sh
    ///  $ pulumi import aws:s3control/objectLambdaAccessPointPolicy:ObjectLambdaAccessPointPolicy example 123456789012:example
    /// ```
    /// </summary>
    [AwsResourceType("aws:s3control/objectLambdaAccessPointPolicy:ObjectLambdaAccessPointPolicy")]
    public partial class ObjectLambdaAccessPointPolicy : Pulumi.CustomResource
    {
        [Output("accountId")]
        public Output<string> AccountId { get; private set; } = null!;

        /// <summary>
        /// Indicates whether this access point currently has a policy that allows public access.
        /// </summary>
        [Output("hasPublicAccessPolicy")]
        public Output<bool> HasPublicAccessPolicy { get; private set; } = null!;

        /// <summary>
        /// The name of the Object Lambda Access Point.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The Object Lambda Access Point resource policy document.
        /// </summary>
        [Output("policy")]
        public Output<string> Policy { get; private set; } = null!;


        /// <summary>
        /// Create a ObjectLambdaAccessPointPolicy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ObjectLambdaAccessPointPolicy(string name, ObjectLambdaAccessPointPolicyArgs args, CustomResourceOptions? options = null)
            : base("aws:s3control/objectLambdaAccessPointPolicy:ObjectLambdaAccessPointPolicy", name, args ?? new ObjectLambdaAccessPointPolicyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ObjectLambdaAccessPointPolicy(string name, Input<string> id, ObjectLambdaAccessPointPolicyState? state = null, CustomResourceOptions? options = null)
            : base("aws:s3control/objectLambdaAccessPointPolicy:ObjectLambdaAccessPointPolicy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ObjectLambdaAccessPointPolicy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ObjectLambdaAccessPointPolicy Get(string name, Input<string> id, ObjectLambdaAccessPointPolicyState? state = null, CustomResourceOptions? options = null)
        {
            return new ObjectLambdaAccessPointPolicy(name, id, state, options);
        }
    }

    public sealed class ObjectLambdaAccessPointPolicyArgs : Pulumi.ResourceArgs
    {
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        /// <summary>
        /// The name of the Object Lambda Access Point.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Object Lambda Access Point resource policy document.
        /// </summary>
        [Input("policy", required: true)]
        public Input<string> Policy { get; set; } = null!;

        public ObjectLambdaAccessPointPolicyArgs()
        {
        }
    }

    public sealed class ObjectLambdaAccessPointPolicyState : Pulumi.ResourceArgs
    {
        [Input("accountId")]
        public Input<string>? AccountId { get; set; }

        /// <summary>
        /// Indicates whether this access point currently has a policy that allows public access.
        /// </summary>
        [Input("hasPublicAccessPolicy")]
        public Input<bool>? HasPublicAccessPolicy { get; set; }

        /// <summary>
        /// The name of the Object Lambda Access Point.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The Object Lambda Access Point resource policy document.
        /// </summary>
        [Input("policy")]
        public Input<string>? Policy { get; set; }

        public ObjectLambdaAccessPointPolicyState()
        {
        }
    }
}
