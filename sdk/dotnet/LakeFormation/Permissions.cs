// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.LakeFormation
{
    /// <summary>
    /// ## Example Usage
    /// ### Grant Permissions For A Lake Formation S3 Resource
    /// 
    /// ```csharp
    /// using Pulumi;
    /// using Aws = Pulumi.Aws;
    /// 
    /// class MyStack : Stack
    /// {
    ///     public MyStack()
    ///     {
    ///         var example = new Aws.LakeFormation.Permissions("example", new Aws.LakeFormation.PermissionsArgs
    ///         {
    ///             Principal = aws_iam_role.Workflow_role.Arn,
    ///             Permissions = 
    ///             {
    ///                 "ALL",
    ///             },
    ///             DataLocation = new Aws.LakeFormation.Inputs.PermissionsDataLocationArgs
    ///             {
    ///                 Arn = aws_lakeformation_resource.Example.Arn,
    ///             },
    ///         });
    ///     }
    /// 
    /// }
    /// ```
    /// </summary>
    [AwsResourceType("aws:lakeformation/permissions:Permissions")]
    public partial class Permissions : Pulumi.CustomResource
    {
        /// <summary>
        /// Identifier for the Data Catalog. By default, it is the account ID of the caller.
        /// </summary>
        [Output("catalogId")]
        public Output<string?> CatalogId { get; private set; } = null!;

        /// <summary>
        /// Whether the permissions are to be granted for the Data Catalog. Defaults to `false`.
        /// </summary>
        [Output("catalogResource")]
        public Output<bool?> CatalogResource { get; private set; } = null!;

        /// <summary>
        /// Configuration block for a data location resource. Detailed below.
        /// </summary>
        [Output("dataLocation")]
        public Output<Outputs.PermissionsDataLocation> DataLocation { get; private set; } = null!;

        /// <summary>
        /// Configuration block for a database resource. Detailed below.
        /// </summary>
        [Output("database")]
        public Output<Outputs.PermissionsDatabase> Database { get; private set; } = null!;

        /// <summary>
        /// List of permissions granted to the principal. Valid values may include `ALL`, `ALTER`, `CREATE_DATABASE`, `CREATE_TABLE`, `DATA_LOCATION_ACCESS`, `DELETE`, `DESCRIBE`, `DROP`, `INSERT`, and `SELECT`. For details on each permission, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
        /// </summary>
        [Output("permissions")]
        public Output<ImmutableArray<string>> PermissionDetails { get; private set; } = null!;

        /// <summary>
        /// Subset of `permissions` which the principal can pass.
        /// </summary>
        [Output("permissionsWithGrantOptions")]
        public Output<ImmutableArray<string>> PermissionsWithGrantOptions { get; private set; } = null!;

        /// <summary>
        /// Principal to be granted the permissions on the resource. Supported principals include `IAM_ALLOWED_PRINCIPALS` (see Default Behavior and `IAMAllowedPrincipals` above), IAM roles, users, groups, SAML groups and users, QuickSight groups, OUs, and organizations as well as AWS account IDs for cross-account permissions. For more information, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
        /// </summary>
        [Output("principal")]
        public Output<string> Principal { get; private set; } = null!;

        /// <summary>
        /// Configuration block for a table resource. Detailed below.
        /// </summary>
        [Output("table")]
        public Output<Outputs.PermissionsTable> Table { get; private set; } = null!;

        /// <summary>
        /// Configuration block for a table with columns resource. Detailed below.
        /// </summary>
        [Output("tableWithColumns")]
        public Output<Outputs.PermissionsTableWithColumns> TableWithColumns { get; private set; } = null!;


        /// <summary>
        /// Create a Permissions resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Permissions(string name, PermissionsArgs args, CustomResourceOptions? options = null)
            : base("aws:lakeformation/permissions:Permissions", name, args ?? new PermissionsArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Permissions(string name, Input<string> id, PermissionsState? state = null, CustomResourceOptions? options = null)
            : base("aws:lakeformation/permissions:Permissions", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Permissions resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Permissions Get(string name, Input<string> id, PermissionsState? state = null, CustomResourceOptions? options = null)
        {
            return new Permissions(name, id, state, options);
        }
    }

    public sealed class PermissionsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Identifier for the Data Catalog. By default, it is the account ID of the caller.
        /// </summary>
        [Input("catalogId")]
        public Input<string>? CatalogId { get; set; }

        /// <summary>
        /// Whether the permissions are to be granted for the Data Catalog. Defaults to `false`.
        /// </summary>
        [Input("catalogResource")]
        public Input<bool>? CatalogResource { get; set; }

        /// <summary>
        /// Configuration block for a data location resource. Detailed below.
        /// </summary>
        [Input("dataLocation")]
        public Input<Inputs.PermissionsDataLocationArgs>? DataLocation { get; set; }

        /// <summary>
        /// Configuration block for a database resource. Detailed below.
        /// </summary>
        [Input("database")]
        public Input<Inputs.PermissionsDatabaseArgs>? Database { get; set; }

        [Input("permissions", required: true)]
        private InputList<string>? _permissions;

        /// <summary>
        /// List of permissions granted to the principal. Valid values may include `ALL`, `ALTER`, `CREATE_DATABASE`, `CREATE_TABLE`, `DATA_LOCATION_ACCESS`, `DELETE`, `DESCRIBE`, `DROP`, `INSERT`, and `SELECT`. For details on each permission, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
        /// </summary>
        public InputList<string> PermissionDetails
        {
            get => _permissions ?? (_permissions = new InputList<string>());
            set => _permissions = value;
        }

        [Input("permissionsWithGrantOptions")]
        private InputList<string>? _permissionsWithGrantOptions;

        /// <summary>
        /// Subset of `permissions` which the principal can pass.
        /// </summary>
        public InputList<string> PermissionsWithGrantOptions
        {
            get => _permissionsWithGrantOptions ?? (_permissionsWithGrantOptions = new InputList<string>());
            set => _permissionsWithGrantOptions = value;
        }

        /// <summary>
        /// Principal to be granted the permissions on the resource. Supported principals include `IAM_ALLOWED_PRINCIPALS` (see Default Behavior and `IAMAllowedPrincipals` above), IAM roles, users, groups, SAML groups and users, QuickSight groups, OUs, and organizations as well as AWS account IDs for cross-account permissions. For more information, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
        /// </summary>
        [Input("principal", required: true)]
        public Input<string> Principal { get; set; } = null!;

        /// <summary>
        /// Configuration block for a table resource. Detailed below.
        /// </summary>
        [Input("table")]
        public Input<Inputs.PermissionsTableArgs>? Table { get; set; }

        /// <summary>
        /// Configuration block for a table with columns resource. Detailed below.
        /// </summary>
        [Input("tableWithColumns")]
        public Input<Inputs.PermissionsTableWithColumnsArgs>? TableWithColumns { get; set; }

        public PermissionsArgs()
        {
        }
    }

    public sealed class PermissionsState : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Identifier for the Data Catalog. By default, it is the account ID of the caller.
        /// </summary>
        [Input("catalogId")]
        public Input<string>? CatalogId { get; set; }

        /// <summary>
        /// Whether the permissions are to be granted for the Data Catalog. Defaults to `false`.
        /// </summary>
        [Input("catalogResource")]
        public Input<bool>? CatalogResource { get; set; }

        /// <summary>
        /// Configuration block for a data location resource. Detailed below.
        /// </summary>
        [Input("dataLocation")]
        public Input<Inputs.PermissionsDataLocationGetArgs>? DataLocation { get; set; }

        /// <summary>
        /// Configuration block for a database resource. Detailed below.
        /// </summary>
        [Input("database")]
        public Input<Inputs.PermissionsDatabaseGetArgs>? Database { get; set; }

        [Input("permissions")]
        private InputList<string>? _permissions;

        /// <summary>
        /// List of permissions granted to the principal. Valid values may include `ALL`, `ALTER`, `CREATE_DATABASE`, `CREATE_TABLE`, `DATA_LOCATION_ACCESS`, `DELETE`, `DESCRIBE`, `DROP`, `INSERT`, and `SELECT`. For details on each permission, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
        /// </summary>
        public InputList<string> PermissionDetails
        {
            get => _permissions ?? (_permissions = new InputList<string>());
            set => _permissions = value;
        }

        [Input("permissionsWithGrantOptions")]
        private InputList<string>? _permissionsWithGrantOptions;

        /// <summary>
        /// Subset of `permissions` which the principal can pass.
        /// </summary>
        public InputList<string> PermissionsWithGrantOptions
        {
            get => _permissionsWithGrantOptions ?? (_permissionsWithGrantOptions = new InputList<string>());
            set => _permissionsWithGrantOptions = value;
        }

        /// <summary>
        /// Principal to be granted the permissions on the resource. Supported principals include `IAM_ALLOWED_PRINCIPALS` (see Default Behavior and `IAMAllowedPrincipals` above), IAM roles, users, groups, SAML groups and users, QuickSight groups, OUs, and organizations as well as AWS account IDs for cross-account permissions. For more information, see [Lake Formation Permissions Reference](https://docs.aws.amazon.com/lake-formation/latest/dg/lf-permissions-reference.html).
        /// </summary>
        [Input("principal")]
        public Input<string>? Principal { get; set; }

        /// <summary>
        /// Configuration block for a table resource. Detailed below.
        /// </summary>
        [Input("table")]
        public Input<Inputs.PermissionsTableGetArgs>? Table { get; set; }

        /// <summary>
        /// Configuration block for a table with columns resource. Detailed below.
        /// </summary>
        [Input("tableWithColumns")]
        public Input<Inputs.PermissionsTableWithColumnsGetArgs>? TableWithColumns { get; set; }

        public PermissionsState()
        {
        }
    }
}
