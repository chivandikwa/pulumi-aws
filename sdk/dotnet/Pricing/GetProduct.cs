// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Pricing
{
    public static class GetProduct
    {
        /// <summary>
        /// Use this data source to get the pricing information of all products in AWS.
        /// This data source is only available in a us-east-1 or ap-south-1 provider.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Aws.Pricing.GetProduct.InvokeAsync(new Aws.Pricing.GetProductArgs
        ///         {
        ///             Filters = 
        ///             {
        ///                 new Aws.Pricing.Inputs.GetProductFilterArgs
        ///                 {
        ///                     Field = "instanceType",
        ///                     Value = "c5.xlarge",
        ///                 },
        ///                 new Aws.Pricing.Inputs.GetProductFilterArgs
        ///                 {
        ///                     Field = "operatingSystem",
        ///                     Value = "Linux",
        ///                 },
        ///                 new Aws.Pricing.Inputs.GetProductFilterArgs
        ///                 {
        ///                     Field = "location",
        ///                     Value = "US East (N. Virginia)",
        ///                 },
        ///                 new Aws.Pricing.Inputs.GetProductFilterArgs
        ///                 {
        ///                     Field = "preInstalledSw",
        ///                     Value = "NA",
        ///                 },
        ///                 new Aws.Pricing.Inputs.GetProductFilterArgs
        ///                 {
        ///                     Field = "licenseModel",
        ///                     Value = "No License required",
        ///                 },
        ///                 new Aws.Pricing.Inputs.GetProductFilterArgs
        ///                 {
        ///                     Field = "tenancy",
        ///                     Value = "Shared",
        ///                 },
        ///                 new Aws.Pricing.Inputs.GetProductFilterArgs
        ///                 {
        ///                     Field = "capacitystatus",
        ///                     Value = "Used",
        ///                 },
        ///             },
        ///             ServiceCode = "AmazonEC2",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Aws.Pricing.GetProduct.InvokeAsync(new Aws.Pricing.GetProductArgs
        ///         {
        ///             Filters = 
        ///             {
        ///                 new Aws.Pricing.Inputs.GetProductFilterArgs
        ///                 {
        ///                     Field = "instanceType",
        ///                     Value = "ds1.xlarge",
        ///                 },
        ///                 new Aws.Pricing.Inputs.GetProductFilterArgs
        ///                 {
        ///                     Field = "location",
        ///                     Value = "US East (N. Virginia)",
        ///                 },
        ///             },
        ///             ServiceCode = "AmazonRedshift",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetProductResult> InvokeAsync(GetProductArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetProductResult>("aws:pricing/getProduct:getProduct", args ?? new GetProductArgs(), options.WithDefaults());

        /// <summary>
        /// Use this data source to get the pricing information of all products in AWS.
        /// This data source is only available in a us-east-1 or ap-south-1 provider.
        /// 
        /// {{% examples %}}
        /// ## Example Usage
        /// {{% example %}}
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Aws.Pricing.GetProduct.InvokeAsync(new Aws.Pricing.GetProductArgs
        ///         {
        ///             Filters = 
        ///             {
        ///                 new Aws.Pricing.Inputs.GetProductFilterArgs
        ///                 {
        ///                     Field = "instanceType",
        ///                     Value = "c5.xlarge",
        ///                 },
        ///                 new Aws.Pricing.Inputs.GetProductFilterArgs
        ///                 {
        ///                     Field = "operatingSystem",
        ///                     Value = "Linux",
        ///                 },
        ///                 new Aws.Pricing.Inputs.GetProductFilterArgs
        ///                 {
        ///                     Field = "location",
        ///                     Value = "US East (N. Virginia)",
        ///                 },
        ///                 new Aws.Pricing.Inputs.GetProductFilterArgs
        ///                 {
        ///                     Field = "preInstalledSw",
        ///                     Value = "NA",
        ///                 },
        ///                 new Aws.Pricing.Inputs.GetProductFilterArgs
        ///                 {
        ///                     Field = "licenseModel",
        ///                     Value = "No License required",
        ///                 },
        ///                 new Aws.Pricing.Inputs.GetProductFilterArgs
        ///                 {
        ///                     Field = "tenancy",
        ///                     Value = "Shared",
        ///                 },
        ///                 new Aws.Pricing.Inputs.GetProductFilterArgs
        ///                 {
        ///                     Field = "capacitystatus",
        ///                     Value = "Used",
        ///                 },
        ///             },
        ///             ServiceCode = "AmazonEC2",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// 
        /// ```csharp
        /// using Pulumi;
        /// using Aws = Pulumi.Aws;
        /// 
        /// class MyStack : Stack
        /// {
        ///     public MyStack()
        ///     {
        ///         var example = Output.Create(Aws.Pricing.GetProduct.InvokeAsync(new Aws.Pricing.GetProductArgs
        ///         {
        ///             Filters = 
        ///             {
        ///                 new Aws.Pricing.Inputs.GetProductFilterArgs
        ///                 {
        ///                     Field = "instanceType",
        ///                     Value = "ds1.xlarge",
        ///                 },
        ///                 new Aws.Pricing.Inputs.GetProductFilterArgs
        ///                 {
        ///                     Field = "location",
        ///                     Value = "US East (N. Virginia)",
        ///                 },
        ///             },
        ///             ServiceCode = "AmazonRedshift",
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetProductResult> Invoke(GetProductInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetProductResult>("aws:pricing/getProduct:getProduct", args ?? new GetProductInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetProductArgs : Pulumi.InvokeArgs
    {
        [Input("filters", required: true)]
        private List<Inputs.GetProductFilterArgs>? _filters;

        /// <summary>
        /// A list of filters. Passed directly to the API (see GetProducts API reference). These filters must describe a single product, this resource will fail if more than one product is returned by the API.
        /// </summary>
        public List<Inputs.GetProductFilterArgs> Filters
        {
            get => _filters ?? (_filters = new List<Inputs.GetProductFilterArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// The code of the service. Available service codes can be fetched using the DescribeServices pricing API call.
        /// </summary>
        [Input("serviceCode", required: true)]
        public string ServiceCode { get; set; } = null!;

        public GetProductArgs()
        {
        }
    }

    public sealed class GetProductInvokeArgs : Pulumi.InvokeArgs
    {
        [Input("filters", required: true)]
        private InputList<Inputs.GetProductFilterInputArgs>? _filters;

        /// <summary>
        /// A list of filters. Passed directly to the API (see GetProducts API reference). These filters must describe a single product, this resource will fail if more than one product is returned by the API.
        /// </summary>
        public InputList<Inputs.GetProductFilterInputArgs> Filters
        {
            get => _filters ?? (_filters = new InputList<Inputs.GetProductFilterInputArgs>());
            set => _filters = value;
        }

        /// <summary>
        /// The code of the service. Available service codes can be fetched using the DescribeServices pricing API call.
        /// </summary>
        [Input("serviceCode", required: true)]
        public Input<string> ServiceCode { get; set; } = null!;

        public GetProductInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetProductResult
    {
        public readonly ImmutableArray<Outputs.GetProductFilterResult> Filters;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        /// <summary>
        /// Set to the product returned from the API.
        /// </summary>
        public readonly string Result;
        public readonly string ServiceCode;

        [OutputConstructor]
        private GetProductResult(
            ImmutableArray<Outputs.GetProductFilterResult> filters,

            string id,

            string result,

            string serviceCode)
        {
            Filters = filters;
            Id = id;
            Result = result;
            ServiceCode = serviceCode;
        }
    }
}
