﻿// Copyright 2016-2019, Pulumi Corporation.  All rights reserved.

using System.Collections.Generic;
using System.Threading.Tasks;

using Pulumi;
using Pulumi.Aws;
using Pulumi.Aws.Ec2;
using Pulumi.Aws.Ec2.Inputs;
using Pulumi.Aws.Inputs;

class Program
{
    static Task<int> Main()
    {
        return Deployment.RunAsync(async () => {

            var ami = await Pulumi.Aws.Ec2.GetAmi.InvokeAsync(new Pulumi.Aws.Ec2.GetAmiArgs
            {
                MostRecent = true,
                Owners = { "137112412989" },
                Filters = { new Pulumi.Aws.Ec2.Inputs.GetAmiFilterArgs { Name = "name", Values = { "amzn-ami-hvm-*" } } },
            });


            var group = new SecurityGroup("web-secgrp", new SecurityGroupArgs
            {
                Description = "Enable HTTP access",
                Ingress =
            {
                new SecurityGroupIngressArgs
                {
                    Protocol = "tcp",
                    FromPort = 80,
                    ToPort = 80,
                    CidrBlocks = { "0.0.0.0/0" }
                }
            }
            });

            var userData = @"
#!/bin/bash
echo ""Hello, World!"" > index.html
nohup python -m SimpleHTTPServer 80 &
";

            var server = new Instance("web-server-www", new InstanceArgs
            {
                InstanceType = Pulumi.Aws.Ec2.InstanceType.T2_Micro,
                VpcSecurityGroupIds = { group.Id },
                UserData = userData,
                Ami = ami.Id,
            });

            return new Dictionary<string, object>
            {
                { "publicIp",  server.PublicIp },
                { "publicDns",  server.PublicDns },
                { "url", Output.Format($"http://{server.PublicDns}") },
            };
        });
    }
}
