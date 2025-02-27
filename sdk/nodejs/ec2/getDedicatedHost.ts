// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../types";
import * as utilities from "../utilities";

/**
 * Use this data source to get information about an EC2 Dedicated Host.
 *
 * ## Example Usage
 * ### Filter Example
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as aws from "@pulumi/aws";
 *
 * const test = pulumi.output(aws.ec2.getDedicatedHost({
 *     filters: [{
 *         name: "instance-type",
 *         values: ["c5.18xlarge"],
 *     }],
 * }));
 * ```
 */
export function getDedicatedHost(args?: GetDedicatedHostArgs, opts?: pulumi.InvokeOptions): Promise<GetDedicatedHostResult> {
    args = args || {};
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("aws:ec2/getDedicatedHost:getDedicatedHost", {
        "filters": args.filters,
        "hostId": args.hostId,
        "tags": args.tags,
    }, opts);
}

/**
 * A collection of arguments for invoking getDedicatedHost.
 */
export interface GetDedicatedHostArgs {
    /**
     * Configuration block. Detailed below.
     */
    filters?: inputs.ec2.GetDedicatedHostFilter[];
    /**
     * The ID of the Dedicated Host.
     */
    hostId?: string;
    tags?: {[key: string]: string};
}

/**
 * A collection of values returned by getDedicatedHost.
 */
export interface GetDedicatedHostResult {
    /**
     * The ARN of the Dedicated Host.
     */
    readonly arn: string;
    /**
     * Whether auto-placement is on or off.
     */
    readonly autoPlacement: string;
    /**
     * The Availability Zone of the Dedicated Host.
     */
    readonly availabilityZone: string;
    /**
     * The number of cores on the Dedicated Host.
     */
    readonly cores: number;
    readonly filters?: outputs.ec2.GetDedicatedHostFilter[];
    readonly hostId: string;
    /**
     * Indicates whether host recovery is enabled or disabled for the Dedicated Host.
     */
    readonly hostRecovery: string;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    /**
     * The instance family supported by the Dedicated Host. For example, "m5".
     */
    readonly instanceFamily: string;
    /**
     * The instance type supported by the Dedicated Host. For example, "m5.large". If the host supports multiple instance types, no instanceType is returned.
     */
    readonly instanceType: string;
    /**
     * The ID of the AWS account that owns the Dedicated Host.
     */
    readonly ownerId: string;
    /**
     * The number of sockets on the Dedicated Host.
     */
    readonly sockets: number;
    readonly tags: {[key: string]: string};
    /**
     * The total number of vCPUs on the Dedicated Host.
     */
    readonly totalVcpus: number;
}

export function getDedicatedHostOutput(args?: GetDedicatedHostOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetDedicatedHostResult> {
    return pulumi.output(args).apply(a => getDedicatedHost(a, opts))
}

/**
 * A collection of arguments for invoking getDedicatedHost.
 */
export interface GetDedicatedHostOutputArgs {
    /**
     * Configuration block. Detailed below.
     */
    filters?: pulumi.Input<pulumi.Input<inputs.ec2.GetDedicatedHostFilterArgs>[]>;
    /**
     * The ID of the Dedicated Host.
     */
    hostId?: pulumi.Input<string>;
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
