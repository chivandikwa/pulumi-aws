// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../utilities";

// Export members:
export * from "./account";
export * from "./apiKey";
export * from "./authorizer";
export * from "./basePathMapping";
export * from "./clientCertificate";
export * from "./deployment";
export * from "./documentationPart";
export * from "./documentationVersion";
export * from "./domainName";
export * from "./getDomainName";
export * from "./getExport";
export * from "./getKey";
export * from "./getResource";
export * from "./getRestApi";
export * from "./getSdk";
export * from "./getVpcLink";
export * from "./integration";
export * from "./integrationResponse";
export * from "./method";
export * from "./methodResponse";
export * from "./methodSettings";
export * from "./model";
export * from "./requestValidator";
export * from "./resource";
export * from "./response";
export * from "./restApi";
export * from "./restApiPolicy";
export * from "./stage";
export * from "./usagePlan";
export * from "./usagePlanKey";
export * from "./vpcLink";

// Import resources to register:
import { Account } from "./account";
import { ApiKey } from "./apiKey";
import { Authorizer } from "./authorizer";
import { BasePathMapping } from "./basePathMapping";
import { ClientCertificate } from "./clientCertificate";
import { Deployment } from "./deployment";
import { DocumentationPart } from "./documentationPart";
import { DocumentationVersion } from "./documentationVersion";
import { DomainName } from "./domainName";
import { Integration } from "./integration";
import { IntegrationResponse } from "./integrationResponse";
import { Method } from "./method";
import { MethodResponse } from "./methodResponse";
import { MethodSettings } from "./methodSettings";
import { Model } from "./model";
import { RequestValidator } from "./requestValidator";
import { Resource } from "./resource";
import { Response } from "./response";
import { RestApi } from "./restApi";
import { RestApiPolicy } from "./restApiPolicy";
import { Stage } from "./stage";
import { UsagePlan } from "./usagePlan";
import { UsagePlanKey } from "./usagePlanKey";
import { VpcLink } from "./vpcLink";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "aws:apigateway/account:Account":
                return new Account(name, <any>undefined, { urn })
            case "aws:apigateway/apiKey:ApiKey":
                return new ApiKey(name, <any>undefined, { urn })
            case "aws:apigateway/authorizer:Authorizer":
                return new Authorizer(name, <any>undefined, { urn })
            case "aws:apigateway/basePathMapping:BasePathMapping":
                return new BasePathMapping(name, <any>undefined, { urn })
            case "aws:apigateway/clientCertificate:ClientCertificate":
                return new ClientCertificate(name, <any>undefined, { urn })
            case "aws:apigateway/deployment:Deployment":
                return new Deployment(name, <any>undefined, { urn })
            case "aws:apigateway/documentationPart:DocumentationPart":
                return new DocumentationPart(name, <any>undefined, { urn })
            case "aws:apigateway/documentationVersion:DocumentationVersion":
                return new DocumentationVersion(name, <any>undefined, { urn })
            case "aws:apigateway/domainName:DomainName":
                return new DomainName(name, <any>undefined, { urn })
            case "aws:apigateway/integration:Integration":
                return new Integration(name, <any>undefined, { urn })
            case "aws:apigateway/integrationResponse:IntegrationResponse":
                return new IntegrationResponse(name, <any>undefined, { urn })
            case "aws:apigateway/method:Method":
                return new Method(name, <any>undefined, { urn })
            case "aws:apigateway/methodResponse:MethodResponse":
                return new MethodResponse(name, <any>undefined, { urn })
            case "aws:apigateway/methodSettings:MethodSettings":
                return new MethodSettings(name, <any>undefined, { urn })
            case "aws:apigateway/model:Model":
                return new Model(name, <any>undefined, { urn })
            case "aws:apigateway/requestValidator:RequestValidator":
                return new RequestValidator(name, <any>undefined, { urn })
            case "aws:apigateway/resource:Resource":
                return new Resource(name, <any>undefined, { urn })
            case "aws:apigateway/response:Response":
                return new Response(name, <any>undefined, { urn })
            case "aws:apigateway/restApi:RestApi":
                return new RestApi(name, <any>undefined, { urn })
            case "aws:apigateway/restApiPolicy:RestApiPolicy":
                return new RestApiPolicy(name, <any>undefined, { urn })
            case "aws:apigateway/stage:Stage":
                return new Stage(name, <any>undefined, { urn })
            case "aws:apigateway/usagePlan:UsagePlan":
                return new UsagePlan(name, <any>undefined, { urn })
            case "aws:apigateway/usagePlanKey:UsagePlanKey":
                return new UsagePlanKey(name, <any>undefined, { urn })
            case "aws:apigateway/vpcLink:VpcLink":
                return new VpcLink(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("aws", "apigateway/account", _module)
pulumi.runtime.registerResourceModule("aws", "apigateway/apiKey", _module)
pulumi.runtime.registerResourceModule("aws", "apigateway/authorizer", _module)
pulumi.runtime.registerResourceModule("aws", "apigateway/basePathMapping", _module)
pulumi.runtime.registerResourceModule("aws", "apigateway/clientCertificate", _module)
pulumi.runtime.registerResourceModule("aws", "apigateway/deployment", _module)
pulumi.runtime.registerResourceModule("aws", "apigateway/documentationPart", _module)
pulumi.runtime.registerResourceModule("aws", "apigateway/documentationVersion", _module)
pulumi.runtime.registerResourceModule("aws", "apigateway/domainName", _module)
pulumi.runtime.registerResourceModule("aws", "apigateway/integration", _module)
pulumi.runtime.registerResourceModule("aws", "apigateway/integrationResponse", _module)
pulumi.runtime.registerResourceModule("aws", "apigateway/method", _module)
pulumi.runtime.registerResourceModule("aws", "apigateway/methodResponse", _module)
pulumi.runtime.registerResourceModule("aws", "apigateway/methodSettings", _module)
pulumi.runtime.registerResourceModule("aws", "apigateway/model", _module)
pulumi.runtime.registerResourceModule("aws", "apigateway/requestValidator", _module)
pulumi.runtime.registerResourceModule("aws", "apigateway/resource", _module)
pulumi.runtime.registerResourceModule("aws", "apigateway/response", _module)
pulumi.runtime.registerResourceModule("aws", "apigateway/restApi", _module)
pulumi.runtime.registerResourceModule("aws", "apigateway/restApiPolicy", _module)
pulumi.runtime.registerResourceModule("aws", "apigateway/stage", _module)
pulumi.runtime.registerResourceModule("aws", "apigateway/usagePlan", _module)
pulumi.runtime.registerResourceModule("aws", "apigateway/usagePlanKey", _module)
pulumi.runtime.registerResourceModule("aws", "apigateway/vpcLink", _module)
