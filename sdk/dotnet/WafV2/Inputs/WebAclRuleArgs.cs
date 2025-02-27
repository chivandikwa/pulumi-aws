// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Aws.WafV2.Inputs
{

    public sealed class WebAclRuleArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Action that AWS WAF should take on a web request when it matches the rule's statement. This is used only for rules whose **statements do not reference a rule group**. See Action below for details.
        /// </summary>
        [Input("action")]
        public Input<Inputs.WebAclRuleActionArgs>? Action { get; set; }

        /// <summary>
        /// Friendly name of the rule.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Override action to apply to the rules in a rule group. Used only for rule **statements that reference a rule group**, like `rule_group_reference_statement` and `managed_rule_group_statement`. See Override Action below for details.
        /// </summary>
        [Input("overrideAction")]
        public Input<Inputs.WebAclRuleOverrideActionArgs>? OverrideAction { get; set; }

        /// <summary>
        /// If you define more than one Rule in a WebACL, AWS WAF evaluates each request against the `rules` in order based on the value of `priority`. AWS WAF processes rules with lower priority first.
        /// </summary>
        [Input("priority", required: true)]
        public Input<int> Priority { get; set; } = null!;

        [Input("ruleLabels")]
        private InputList<Inputs.WebAclRuleRuleLabelArgs>? _ruleLabels;

        /// <summary>
        /// Labels to apply to web requests that match the rule match statement. See Rule Label below for details.
        /// </summary>
        public InputList<Inputs.WebAclRuleRuleLabelArgs> RuleLabels
        {
            get => _ruleLabels ?? (_ruleLabels = new InputList<Inputs.WebAclRuleRuleLabelArgs>());
            set => _ruleLabels = value;
        }

        /// <summary>
        /// The AWS WAF processing statement for the rule, for example `byte_match_statement` or `geo_match_statement`. See Statement below for details.
        /// </summary>
        [Input("statement", required: true)]
        public Input<Inputs.WebAclRuleStatementArgs> Statement { get; set; } = null!;

        /// <summary>
        /// Defines and enables Amazon CloudWatch metrics and web request sample collection. See Visibility Configuration below for details.
        /// </summary>
        [Input("visibilityConfig", required: true)]
        public Input<Inputs.WebAclRuleVisibilityConfigArgs> VisibilityConfig { get; set; } = null!;

        public WebAclRuleArgs()
        {
        }
    }
}
