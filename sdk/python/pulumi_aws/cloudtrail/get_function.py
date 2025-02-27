# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities

__all__ = [
    'GetFunctionResult',
    'AwaitableGetFunctionResult',
    'get_function',
    'get_function_output',
]

warnings.warn("""aws.cloudtrail.getFunction has been deprecated in favor of aws.cloudfront.getFunction""", DeprecationWarning)

@pulumi.output_type
class GetFunctionResult:
    """
    A collection of values returned by getFunction.
    """
    def __init__(__self__, arn=None, code=None, comment=None, etag=None, id=None, last_modified_time=None, name=None, runtime=None, stage=None, status=None):
        if arn and not isinstance(arn, str):
            raise TypeError("Expected argument 'arn' to be a str")
        pulumi.set(__self__, "arn", arn)
        if code and not isinstance(code, str):
            raise TypeError("Expected argument 'code' to be a str")
        pulumi.set(__self__, "code", code)
        if comment and not isinstance(comment, str):
            raise TypeError("Expected argument 'comment' to be a str")
        pulumi.set(__self__, "comment", comment)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if last_modified_time and not isinstance(last_modified_time, str):
            raise TypeError("Expected argument 'last_modified_time' to be a str")
        pulumi.set(__self__, "last_modified_time", last_modified_time)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if runtime and not isinstance(runtime, str):
            raise TypeError("Expected argument 'runtime' to be a str")
        pulumi.set(__self__, "runtime", runtime)
        if stage and not isinstance(stage, str):
            raise TypeError("Expected argument 'stage' to be a str")
        pulumi.set(__self__, "stage", stage)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter
    def arn(self) -> str:
        """
        Amazon Resource Name (ARN) identifying your CloudFront Function.
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def code(self) -> str:
        """
        Source code of the function
        """
        return pulumi.get(self, "code")

    @property
    @pulumi.getter
    def comment(self) -> str:
        """
        Comment.
        """
        return pulumi.get(self, "comment")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        ETag hash of the function
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="lastModifiedTime")
    def last_modified_time(self) -> str:
        """
        When this resource was last modified.
        """
        return pulumi.get(self, "last_modified_time")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def runtime(self) -> str:
        """
        Identifier of the function's runtime.
        """
        return pulumi.get(self, "runtime")

    @property
    @pulumi.getter
    def stage(self) -> str:
        return pulumi.get(self, "stage")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        Status of the function. Can be `UNPUBLISHED`, `UNASSOCIATED` or `ASSOCIATED`.
        """
        return pulumi.get(self, "status")


class AwaitableGetFunctionResult(GetFunctionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFunctionResult(
            arn=self.arn,
            code=self.code,
            comment=self.comment,
            etag=self.etag,
            id=self.id,
            last_modified_time=self.last_modified_time,
            name=self.name,
            runtime=self.runtime,
            stage=self.stage,
            status=self.status)


def get_function(name: Optional[str] = None,
                 stage: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFunctionResult:
    """
    Provides information about a CloudFront Function.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    config = pulumi.Config()
    function_name = config.require("functionName")
    existing = aws.cloudfront.get_function(name=function_name)
    ```


    :param str name: Name of the CloudFront function.
    :param str stage: The function’s stage, either `DEVELOPMENT` or `LIVE`.
    """
    pulumi.log.warn("""get_function is deprecated: aws.cloudtrail.getFunction has been deprecated in favor of aws.cloudfront.getFunction""")
    __args__ = dict()
    __args__['name'] = name
    __args__['stage'] = stage
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('aws:cloudtrail/getFunction:getFunction', __args__, opts=opts, typ=GetFunctionResult).value

    return AwaitableGetFunctionResult(
        arn=__ret__.arn,
        code=__ret__.code,
        comment=__ret__.comment,
        etag=__ret__.etag,
        id=__ret__.id,
        last_modified_time=__ret__.last_modified_time,
        name=__ret__.name,
        runtime=__ret__.runtime,
        stage=__ret__.stage,
        status=__ret__.status)


@_utilities.lift_output_func(get_function)
def get_function_output(name: Optional[pulumi.Input[str]] = None,
                        stage: Optional[pulumi.Input[str]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFunctionResult]:
    """
    Provides information about a CloudFront Function.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_aws as aws

    config = pulumi.Config()
    function_name = config.require("functionName")
    existing = aws.cloudfront.get_function(name=function_name)
    ```


    :param str name: Name of the CloudFront function.
    :param str stage: The function’s stage, either `DEVELOPMENT` or `LIVE`.
    """
    pulumi.log.warn("""get_function is deprecated: aws.cloudtrail.getFunction has been deprecated in favor of aws.cloudfront.getFunction""")
    ...
