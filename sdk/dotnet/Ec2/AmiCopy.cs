// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Ec2
{
    /// <summary>
    /// The "AMI copy" resource allows duplication of an Amazon Machine Image (AMI),
    /// including cross-region copies.
    /// 
    /// If the source AMI has associated EBS snapshots, those will also be duplicated
    /// along with the AMI.
    /// 
    /// This is useful for taking a single AMI provisioned in one region and making
    /// it available in another for a multi-region deployment.
    /// 
    /// Copying an AMI can take several minutes. The creation of this resource will
    /// block until the new AMI is available for use on new instances.
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
    ///         var example = new Aws.Ec2.AmiCopy("example", new Aws.Ec2.AmiCopyArgs
    ///         {
    ///             Description = "A copy of ami-xxxxxxxx",
    ///             SourceAmiId = "ami-xxxxxxxx",
    ///             SourceAmiRegion = "us-west-1",
    ///             Tags = 
    ///             {
    ///                 { "Name", "HelloWorld" },
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [AwsResourceType("aws:ec2/amiCopy:AmiCopy")]
    public partial class AmiCopy : Pulumi.CustomResource
    {
        /// <summary>
        /// Machine architecture for created instances. Defaults to "x86_64".
        /// </summary>
        [Output("architecture")]
        public Output<string> Architecture { get; private set; } = null!;

        /// <summary>
        /// The ARN of the AMI.
        /// </summary>
        [Output("arn")]
        public Output<string> Arn { get; private set; } = null!;

        /// <summary>
        /// The boot mode of the AMI. For more information, see [Boot modes](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ami-boot.html) in the Amazon Elastic Compute Cloud User Guide.
        /// </summary>
        [Output("bootMode")]
        public Output<string> BootMode { get; private set; } = null!;

        /// <summary>
        /// A longer, human-readable description for the AMI.
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// The ARN of the Outpost to which to copy the AMI.
        /// Only specify this parameter when copying an AMI from an AWS Region to an Outpost. The AMI must be in the Region of the destination Outpost.
        /// </summary>
        [Output("destinationOutpostArn")]
        public Output<string?> DestinationOutpostArn { get; private set; } = null!;

        /// <summary>
        /// Nested block describing an EBS block device that should be
        /// attached to created instances. The structure of this block is described below.
        /// </summary>
        [Output("ebsBlockDevices")]
        public Output<ImmutableArray<Outputs.AmiCopyEbsBlockDevice>> EbsBlockDevices { get; private set; } = null!;

        /// <summary>
        /// Specifies whether enhanced networking with ENA is enabled. Defaults to `false`.
        /// </summary>
        [Output("enaSupport")]
        public Output<bool> EnaSupport { get; private set; } = null!;

        /// <summary>
        /// Boolean controlling whether the created EBS volumes will be encrypted. Can't be used with `snapshot_id`.
        /// </summary>
        [Output("encrypted")]
        public Output<bool?> Encrypted { get; private set; } = null!;

        /// <summary>
        /// Nested block describing an ephemeral block device that
        /// should be attached to created instances. The structure of this block is described below.
        /// </summary>
        [Output("ephemeralBlockDevices")]
        public Output<ImmutableArray<Outputs.AmiCopyEphemeralBlockDevice>> EphemeralBlockDevices { get; private set; } = null!;

        [Output("hypervisor")]
        public Output<string> Hypervisor { get; private set; } = null!;

        /// <summary>
        /// Path to an S3 object containing an image manifest, e.g., created
        /// by the `ec2-upload-bundle` command in the EC2 command line tools.
        /// </summary>
        [Output("imageLocation")]
        public Output<string> ImageLocation { get; private set; } = null!;

        [Output("imageOwnerAlias")]
        public Output<string> ImageOwnerAlias { get; private set; } = null!;

        [Output("imageType")]
        public Output<string> ImageType { get; private set; } = null!;

        /// <summary>
        /// The id of the kernel image (AKI) that will be used as the paravirtual
        /// kernel in created instances.
        /// </summary>
        [Output("kernelId")]
        public Output<string> KernelId { get; private set; } = null!;

        /// <summary>
        /// The full ARN of the AWS Key Management Service (AWS KMS) CMK to use when encrypting the snapshots of
        /// an image during a copy operation. This parameter is only required if you want to use a non-default CMK;
        /// if this parameter is not specified, the default CMK for EBS is used
        /// </summary>
        [Output("kmsKeyId")]
        public Output<string> KmsKeyId { get; private set; } = null!;

        [Output("manageEbsSnapshots")]
        public Output<bool> ManageEbsSnapshots { get; private set; } = null!;

        /// <summary>
        /// A region-unique name for the AMI.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("ownerId")]
        public Output<string> OwnerId { get; private set; } = null!;

        [Output("platform")]
        public Output<string> Platform { get; private set; } = null!;

        [Output("platformDetails")]
        public Output<string> PlatformDetails { get; private set; } = null!;

        [Output("public")]
        public Output<bool> Public { get; private set; } = null!;

        /// <summary>
        /// The id of an initrd image (ARI) that will be used when booting the
        /// created instances.
        /// </summary>
        [Output("ramdiskId")]
        public Output<string> RamdiskId { get; private set; } = null!;

        /// <summary>
        /// The name of the root device (for example, `/dev/sda1`, or `/dev/xvda`).
        /// </summary>
        [Output("rootDeviceName")]
        public Output<string> RootDeviceName { get; private set; } = null!;

        [Output("rootSnapshotId")]
        public Output<string> RootSnapshotId { get; private set; } = null!;

        /// <summary>
        /// The id of the AMI to copy. This id must be valid in the region
        /// given by `source_ami_region`.
        /// </summary>
        [Output("sourceAmiId")]
        public Output<string> SourceAmiId { get; private set; } = null!;

        /// <summary>
        /// The region from which the AMI will be copied. This may be the
        /// same as the AWS provider region in order to create a copy within the same region.
        /// </summary>
        [Output("sourceAmiRegion")]
        public Output<string> SourceAmiRegion { get; private set; } = null!;

        /// <summary>
        /// When set to "simple" (the default), enables enhanced networking
        /// for created instances. No other value is supported at this time.
        /// </summary>
        [Output("sriovNetSupport")]
        public Output<string> SriovNetSupport { get; private set; } = null!;

        /// <summary>
        /// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        [Output("tags")]
        public Output<ImmutableDictionary<string, string>?> Tags { get; private set; } = null!;

        [Output("tagsAll")]
        public Output<ImmutableDictionary<string, string>> TagsAll { get; private set; } = null!;

        [Output("usageOperation")]
        public Output<string> UsageOperation { get; private set; } = null!;

        /// <summary>
        /// Keyword to choose what virtualization mode created instances
        /// will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
        /// changes the set of further arguments that are required, as described below.
        /// </summary>
        [Output("virtualizationType")]
        public Output<string> VirtualizationType { get; private set; } = null!;


        /// <summary>
        /// Create a AmiCopy resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AmiCopy(string name, AmiCopyArgs args, CustomResourceOptions? options = null)
            : base("aws:ec2/amiCopy:AmiCopy", name, args ?? new AmiCopyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AmiCopy(string name, Input<string> id, AmiCopyState? state = null, CustomResourceOptions? options = null)
            : base("aws:ec2/amiCopy:AmiCopy", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AmiCopy resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AmiCopy Get(string name, Input<string> id, AmiCopyState? state = null, CustomResourceOptions? options = null)
        {
            return new AmiCopy(name, id, state, options);
        }
    }

    public sealed class AmiCopyArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A longer, human-readable description for the AMI.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ARN of the Outpost to which to copy the AMI.
        /// Only specify this parameter when copying an AMI from an AWS Region to an Outpost. The AMI must be in the Region of the destination Outpost.
        /// </summary>
        [Input("destinationOutpostArn")]
        public Input<string>? DestinationOutpostArn { get; set; }

        [Input("ebsBlockDevices")]
        private InputList<Inputs.AmiCopyEbsBlockDeviceArgs>? _ebsBlockDevices;

        /// <summary>
        /// Nested block describing an EBS block device that should be
        /// attached to created instances. The structure of this block is described below.
        /// </summary>
        public InputList<Inputs.AmiCopyEbsBlockDeviceArgs> EbsBlockDevices
        {
            get => _ebsBlockDevices ?? (_ebsBlockDevices = new InputList<Inputs.AmiCopyEbsBlockDeviceArgs>());
            set => _ebsBlockDevices = value;
        }

        /// <summary>
        /// Boolean controlling whether the created EBS volumes will be encrypted. Can't be used with `snapshot_id`.
        /// </summary>
        [Input("encrypted")]
        public Input<bool>? Encrypted { get; set; }

        [Input("ephemeralBlockDevices")]
        private InputList<Inputs.AmiCopyEphemeralBlockDeviceArgs>? _ephemeralBlockDevices;

        /// <summary>
        /// Nested block describing an ephemeral block device that
        /// should be attached to created instances. The structure of this block is described below.
        /// </summary>
        public InputList<Inputs.AmiCopyEphemeralBlockDeviceArgs> EphemeralBlockDevices
        {
            get => _ephemeralBlockDevices ?? (_ephemeralBlockDevices = new InputList<Inputs.AmiCopyEphemeralBlockDeviceArgs>());
            set => _ephemeralBlockDevices = value;
        }

        /// <summary>
        /// The full ARN of the AWS Key Management Service (AWS KMS) CMK to use when encrypting the snapshots of
        /// an image during a copy operation. This parameter is only required if you want to use a non-default CMK;
        /// if this parameter is not specified, the default CMK for EBS is used
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        /// <summary>
        /// A region-unique name for the AMI.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The id of the AMI to copy. This id must be valid in the region
        /// given by `source_ami_region`.
        /// </summary>
        [Input("sourceAmiId", required: true)]
        public Input<string> SourceAmiId { get; set; } = null!;

        /// <summary>
        /// The region from which the AMI will be copied. This may be the
        /// same as the AWS provider region in order to create a copy within the same region.
        /// </summary>
        [Input("sourceAmiRegion", required: true)]
        public Input<string> SourceAmiRegion { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public AmiCopyArgs()
        {
        }
    }

    public sealed class AmiCopyState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Machine architecture for created instances. Defaults to "x86_64".
        /// </summary>
        [Input("architecture")]
        public Input<string>? Architecture { get; set; }

        /// <summary>
        /// The ARN of the AMI.
        /// </summary>
        [Input("arn")]
        public Input<string>? Arn { get; set; }

        /// <summary>
        /// The boot mode of the AMI. For more information, see [Boot modes](https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/ami-boot.html) in the Amazon Elastic Compute Cloud User Guide.
        /// </summary>
        [Input("bootMode")]
        public Input<string>? BootMode { get; set; }

        /// <summary>
        /// A longer, human-readable description for the AMI.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// The ARN of the Outpost to which to copy the AMI.
        /// Only specify this parameter when copying an AMI from an AWS Region to an Outpost. The AMI must be in the Region of the destination Outpost.
        /// </summary>
        [Input("destinationOutpostArn")]
        public Input<string>? DestinationOutpostArn { get; set; }

        [Input("ebsBlockDevices")]
        private InputList<Inputs.AmiCopyEbsBlockDeviceGetArgs>? _ebsBlockDevices;

        /// <summary>
        /// Nested block describing an EBS block device that should be
        /// attached to created instances. The structure of this block is described below.
        /// </summary>
        public InputList<Inputs.AmiCopyEbsBlockDeviceGetArgs> EbsBlockDevices
        {
            get => _ebsBlockDevices ?? (_ebsBlockDevices = new InputList<Inputs.AmiCopyEbsBlockDeviceGetArgs>());
            set => _ebsBlockDevices = value;
        }

        /// <summary>
        /// Specifies whether enhanced networking with ENA is enabled. Defaults to `false`.
        /// </summary>
        [Input("enaSupport")]
        public Input<bool>? EnaSupport { get; set; }

        /// <summary>
        /// Boolean controlling whether the created EBS volumes will be encrypted. Can't be used with `snapshot_id`.
        /// </summary>
        [Input("encrypted")]
        public Input<bool>? Encrypted { get; set; }

        [Input("ephemeralBlockDevices")]
        private InputList<Inputs.AmiCopyEphemeralBlockDeviceGetArgs>? _ephemeralBlockDevices;

        /// <summary>
        /// Nested block describing an ephemeral block device that
        /// should be attached to created instances. The structure of this block is described below.
        /// </summary>
        public InputList<Inputs.AmiCopyEphemeralBlockDeviceGetArgs> EphemeralBlockDevices
        {
            get => _ephemeralBlockDevices ?? (_ephemeralBlockDevices = new InputList<Inputs.AmiCopyEphemeralBlockDeviceGetArgs>());
            set => _ephemeralBlockDevices = value;
        }

        [Input("hypervisor")]
        public Input<string>? Hypervisor { get; set; }

        /// <summary>
        /// Path to an S3 object containing an image manifest, e.g., created
        /// by the `ec2-upload-bundle` command in the EC2 command line tools.
        /// </summary>
        [Input("imageLocation")]
        public Input<string>? ImageLocation { get; set; }

        [Input("imageOwnerAlias")]
        public Input<string>? ImageOwnerAlias { get; set; }

        [Input("imageType")]
        public Input<string>? ImageType { get; set; }

        /// <summary>
        /// The id of the kernel image (AKI) that will be used as the paravirtual
        /// kernel in created instances.
        /// </summary>
        [Input("kernelId")]
        public Input<string>? KernelId { get; set; }

        /// <summary>
        /// The full ARN of the AWS Key Management Service (AWS KMS) CMK to use when encrypting the snapshots of
        /// an image during a copy operation. This parameter is only required if you want to use a non-default CMK;
        /// if this parameter is not specified, the default CMK for EBS is used
        /// </summary>
        [Input("kmsKeyId")]
        public Input<string>? KmsKeyId { get; set; }

        [Input("manageEbsSnapshots")]
        public Input<bool>? ManageEbsSnapshots { get; set; }

        /// <summary>
        /// A region-unique name for the AMI.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("ownerId")]
        public Input<string>? OwnerId { get; set; }

        [Input("platform")]
        public Input<string>? Platform { get; set; }

        [Input("platformDetails")]
        public Input<string>? PlatformDetails { get; set; }

        [Input("public")]
        public Input<bool>? Public { get; set; }

        /// <summary>
        /// The id of an initrd image (ARI) that will be used when booting the
        /// created instances.
        /// </summary>
        [Input("ramdiskId")]
        public Input<string>? RamdiskId { get; set; }

        /// <summary>
        /// The name of the root device (for example, `/dev/sda1`, or `/dev/xvda`).
        /// </summary>
        [Input("rootDeviceName")]
        public Input<string>? RootDeviceName { get; set; }

        [Input("rootSnapshotId")]
        public Input<string>? RootSnapshotId { get; set; }

        /// <summary>
        /// The id of the AMI to copy. This id must be valid in the region
        /// given by `source_ami_region`.
        /// </summary>
        [Input("sourceAmiId")]
        public Input<string>? SourceAmiId { get; set; }

        /// <summary>
        /// The region from which the AMI will be copied. This may be the
        /// same as the AWS provider region in order to create a copy within the same region.
        /// </summary>
        [Input("sourceAmiRegion")]
        public Input<string>? SourceAmiRegion { get; set; }

        /// <summary>
        /// When set to "simple" (the default), enables enhanced networking
        /// for created instances. No other value is supported at this time.
        /// </summary>
        [Input("sriovNetSupport")]
        public Input<string>? SriovNetSupport { get; set; }

        [Input("tags")]
        private InputMap<string>? _tags;

        /// <summary>
        /// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        /// </summary>
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        [Input("tagsAll")]
        private InputMap<string>? _tagsAll;
        public InputMap<string> TagsAll
        {
            get => _tagsAll ?? (_tagsAll = new InputMap<string>());
            set => _tagsAll = value;
        }

        [Input("usageOperation")]
        public Input<string>? UsageOperation { get; set; }

        /// <summary>
        /// Keyword to choose what virtualization mode created instances
        /// will use. Can be either "paravirtual" (the default) or "hvm". The choice of virtualization type
        /// changes the set of further arguments that are required, as described below.
        /// </summary>
        [Input("virtualizationType")]
        public Input<string>? VirtualizationType { get; set; }

        public AmiCopyState()
        {
        }
    }
}
