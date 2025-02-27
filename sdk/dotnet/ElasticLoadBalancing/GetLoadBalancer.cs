// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.ElasticLoadBalancing
{
    [Obsolete(@"aws.elasticloadbalancing.getLoadBalancer has been deprecated in favor of aws.elb.getLoadBalancer")]
    public static class GetLoadBalancer
    {
        /// <summary>
        /// Provides information about a "classic" Elastic Load Balancer (ELB).
        /// See `LB` Data Source if you are looking for "v2"
        /// Application Load Balancer (ALB) or Network Load Balancer (NLB).
        /// 
        /// This data source can prove useful when a module accepts an LB as an input
        /// variable and needs to, for example, determine the security groups associated
        /// with it, etc.
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
        ///         var config = new Config();
        ///         var lbName = config.Get("lbName") ?? "";
        ///         var test = Output.Create(Aws.Elb.GetLoadBalancer.InvokeAsync(new Aws.Elb.GetLoadBalancerArgs
        ///         {
        ///             Name = lbName,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Task<GetLoadBalancerResult> InvokeAsync(GetLoadBalancerArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetLoadBalancerResult>("aws:elasticloadbalancing/getLoadBalancer:getLoadBalancer", args ?? new GetLoadBalancerArgs(), options.WithDefaults());

        /// <summary>
        /// Provides information about a "classic" Elastic Load Balancer (ELB).
        /// See `LB` Data Source if you are looking for "v2"
        /// Application Load Balancer (ALB) or Network Load Balancer (NLB).
        /// 
        /// This data source can prove useful when a module accepts an LB as an input
        /// variable and needs to, for example, determine the security groups associated
        /// with it, etc.
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
        ///         var config = new Config();
        ///         var lbName = config.Get("lbName") ?? "";
        ///         var test = Output.Create(Aws.Elb.GetLoadBalancer.InvokeAsync(new Aws.Elb.GetLoadBalancerArgs
        ///         {
        ///             Name = lbName,
        ///         }));
        ///     }
        /// 
        /// }
        /// ```
        /// {{% /example %}}
        /// {{% /examples %}}
        /// </summary>
        public static Output<GetLoadBalancerResult> Invoke(GetLoadBalancerInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetLoadBalancerResult>("aws:elasticloadbalancing/getLoadBalancer:getLoadBalancer", args ?? new GetLoadBalancerInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetLoadBalancerArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique name of the load balancer.
        /// </summary>
        [Input("name", required: true)]
        public string Name { get; set; } = null!;

        [Input("tags")]
        private Dictionary<string, string>? _tags;
        public Dictionary<string, string> Tags
        {
            get => _tags ?? (_tags = new Dictionary<string, string>());
            set => _tags = value;
        }

        public GetLoadBalancerArgs()
        {
        }
    }

    public sealed class GetLoadBalancerInvokeArgs : Pulumi.InvokeArgs
    {
        /// <summary>
        /// The unique name of the load balancer.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("tags")]
        private InputMap<string>? _tags;
        public InputMap<string> Tags
        {
            get => _tags ?? (_tags = new InputMap<string>());
            set => _tags = value;
        }

        public GetLoadBalancerInvokeArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetLoadBalancerResult
    {
        public readonly Outputs.GetLoadBalancerAccessLogsResult AccessLogs;
        public readonly string Arn;
        public readonly ImmutableArray<string> AvailabilityZones;
        public readonly bool ConnectionDraining;
        public readonly int ConnectionDrainingTimeout;
        public readonly bool CrossZoneLoadBalancing;
        public readonly string DesyncMitigationMode;
        public readonly string DnsName;
        public readonly Outputs.GetLoadBalancerHealthCheckResult HealthCheck;
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly int IdleTimeout;
        public readonly ImmutableArray<string> Instances;
        public readonly bool Internal;
        public readonly ImmutableArray<Outputs.GetLoadBalancerListenerResult> Listeners;
        public readonly string Name;
        public readonly ImmutableArray<string> SecurityGroups;
        public readonly string SourceSecurityGroup;
        public readonly string SourceSecurityGroupId;
        public readonly ImmutableArray<string> Subnets;
        public readonly ImmutableDictionary<string, string> Tags;
        public readonly string ZoneId;

        [OutputConstructor]
        private GetLoadBalancerResult(
            Outputs.GetLoadBalancerAccessLogsResult accessLogs,

            string arn,

            ImmutableArray<string> availabilityZones,

            bool connectionDraining,

            int connectionDrainingTimeout,

            bool crossZoneLoadBalancing,

            string desyncMitigationMode,

            string dnsName,

            Outputs.GetLoadBalancerHealthCheckResult healthCheck,

            string id,

            int idleTimeout,

            ImmutableArray<string> instances,

            bool @internal,

            ImmutableArray<Outputs.GetLoadBalancerListenerResult> listeners,

            string name,

            ImmutableArray<string> securityGroups,

            string sourceSecurityGroup,

            string sourceSecurityGroupId,

            ImmutableArray<string> subnets,

            ImmutableDictionary<string, string> tags,

            string zoneId)
        {
            AccessLogs = accessLogs;
            Arn = arn;
            AvailabilityZones = availabilityZones;
            ConnectionDraining = connectionDraining;
            ConnectionDrainingTimeout = connectionDrainingTimeout;
            CrossZoneLoadBalancing = crossZoneLoadBalancing;
            DesyncMitigationMode = desyncMitigationMode;
            DnsName = dnsName;
            HealthCheck = healthCheck;
            Id = id;
            IdleTimeout = idleTimeout;
            Instances = instances;
            Internal = @internal;
            Listeners = listeners;
            Name = name;
            SecurityGroups = securityGroups;
            SourceSecurityGroup = sourceSecurityGroup;
            SourceSecurityGroupId = sourceSecurityGroupId;
            Subnets = subnets;
            Tags = tags;
            ZoneId = zoneId;
        }
    }
}
