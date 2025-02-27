// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.AutoScaling.Inputs
{

    public sealed class PolicyStepAdjustmentGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The lower bound for the
        /// difference between the alarm threshold and the CloudWatch metric.
        /// Without a value, AWS will treat this bound as negative infinity.
        /// </summary>
        [Input("metricIntervalLowerBound")]
        public Input<string>? MetricIntervalLowerBound { get; set; }

        /// <summary>
        /// The upper bound for the
        /// difference between the alarm threshold and the CloudWatch metric.
        /// Without a value, AWS will treat this bound as positive infinity. The upper bound
        /// must be greater than the lower bound.
        /// </summary>
        [Input("metricIntervalUpperBound")]
        public Input<string>? MetricIntervalUpperBound { get; set; }

        /// <summary>
        /// The number of members by which to
        /// scale, when the adjustment bounds are breached. A positive value scales
        /// up. A negative value scales down.
        /// </summary>
        [Input("scalingAdjustment", required: true)]
        public Input<int> ScalingAdjustment { get; set; } = null!;

        public PolicyStepAdjustmentGetArgs()
        {
        }
    }
}
