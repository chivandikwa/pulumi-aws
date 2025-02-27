// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.Dlm.Inputs
{

    public sealed class LifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleGetArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The amount of time to retain each snapshot. The maximum is 100 years. This is equivalent to 1200 months, 5200 weeks, or 36500 days.
        /// </summary>
        [Input("interval", required: true)]
        public Input<int> Interval { get; set; } = null!;

        /// <summary>
        /// The unit of time for time-based retention. Valid values: `DAYS`, `WEEKS`, `MONTHS`, or `YEARS`.
        /// </summary>
        [Input("intervalUnit", required: true)]
        public Input<string> IntervalUnit { get; set; } = null!;

        public LifecyclePolicyPolicyDetailsScheduleCrossRegionCopyRuleRetainRuleGetArgs()
        {
        }
    }
}
