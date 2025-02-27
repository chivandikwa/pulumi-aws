# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from .. import _utilities
from . import outputs
from ._inputs import *

__all__ = ['VpcIpamArgs', 'VpcIpam']

@pulumi.input_type
class VpcIpamArgs:
    def __init__(__self__, *,
                 operating_regions: pulumi.Input[Sequence[pulumi.Input['VpcIpamOperatingRegionArgs']]],
                 description: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        The set of arguments for constructing a VpcIpam resource.
        :param pulumi.Input[Sequence[pulumi.Input['VpcIpamOperatingRegionArgs']]] operating_regions: Determines which locales can be chosen when you create pools. Locale is the Region where you want to make an IPAM pool available for allocations. You can only create pools with locales that match the operating Regions of the IPAM. You can only create VPCs from a pool whose locale matches the VPC's Region. You specify a region using the region_name parameter. You **must** set your provider block region as an operating_region.
        :param pulumi.Input[str] description: A description for the IPAM.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        pulumi.set(__self__, "operating_regions", operating_regions)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="operatingRegions")
    def operating_regions(self) -> pulumi.Input[Sequence[pulumi.Input['VpcIpamOperatingRegionArgs']]]:
        """
        Determines which locales can be chosen when you create pools. Locale is the Region where you want to make an IPAM pool available for allocations. You can only create pools with locales that match the operating Regions of the IPAM. You can only create VPCs from a pool whose locale matches the VPC's Region. You specify a region using the region_name parameter. You **must** set your provider block region as an operating_region.
        """
        return pulumi.get(self, "operating_regions")

    @operating_regions.setter
    def operating_regions(self, value: pulumi.Input[Sequence[pulumi.Input['VpcIpamOperatingRegionArgs']]]):
        pulumi.set(self, "operating_regions", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description for the IPAM.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _VpcIpamState:
    def __init__(__self__, *,
                 arn: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 operating_regions: Optional[pulumi.Input[Sequence[pulumi.Input['VpcIpamOperatingRegionArgs']]]] = None,
                 private_default_scope_id: Optional[pulumi.Input[str]] = None,
                 public_default_scope_id: Optional[pulumi.Input[str]] = None,
                 scope_count: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None):
        """
        Input properties used for looking up and filtering VpcIpam resources.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of IPAM
        :param pulumi.Input[str] description: A description for the IPAM.
        :param pulumi.Input[Sequence[pulumi.Input['VpcIpamOperatingRegionArgs']]] operating_regions: Determines which locales can be chosen when you create pools. Locale is the Region where you want to make an IPAM pool available for allocations. You can only create pools with locales that match the operating Regions of the IPAM. You can only create VPCs from a pool whose locale matches the VPC's Region. You specify a region using the region_name parameter. You **must** set your provider block region as an operating_region.
        :param pulumi.Input[str] private_default_scope_id: The ID of the IPAM's private scope. A scope is a top-level container in IPAM. Each scope represents an IP-independent network. Scopes enable you to represent networks where you have overlapping IP space. When you create an IPAM, IPAM automatically creates two scopes: public and private. The private scope is intended for private IP space. The public scope is intended for all internet-routable IP space.
        :param pulumi.Input[str] public_default_scope_id: The ID of the IPAM's public scope. A scope is a top-level container in IPAM. Each scope represents an IP-independent network. Scopes enable you to represent networks where you have overlapping IP space. When you create an IPAM, IPAM automatically creates two scopes: public and private. The private scope is intended for private
               IP space. The public scope is intended for all internet-routable IP space.
        :param pulumi.Input[int] scope_count: The number of scopes in the IPAM.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        if arn is not None:
            pulumi.set(__self__, "arn", arn)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if operating_regions is not None:
            pulumi.set(__self__, "operating_regions", operating_regions)
        if private_default_scope_id is not None:
            pulumi.set(__self__, "private_default_scope_id", private_default_scope_id)
        if public_default_scope_id is not None:
            pulumi.set(__self__, "public_default_scope_id", public_default_scope_id)
        if scope_count is not None:
            pulumi.set(__self__, "scope_count", scope_count)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if tags_all is not None:
            pulumi.set(__self__, "tags_all", tags_all)

    @property
    @pulumi.getter
    def arn(self) -> Optional[pulumi.Input[str]]:
        """
        Amazon Resource Name (ARN) of IPAM
        """
        return pulumi.get(self, "arn")

    @arn.setter
    def arn(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "arn", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        A description for the IPAM.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="operatingRegions")
    def operating_regions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VpcIpamOperatingRegionArgs']]]]:
        """
        Determines which locales can be chosen when you create pools. Locale is the Region where you want to make an IPAM pool available for allocations. You can only create pools with locales that match the operating Regions of the IPAM. You can only create VPCs from a pool whose locale matches the VPC's Region. You specify a region using the region_name parameter. You **must** set your provider block region as an operating_region.
        """
        return pulumi.get(self, "operating_regions")

    @operating_regions.setter
    def operating_regions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VpcIpamOperatingRegionArgs']]]]):
        pulumi.set(self, "operating_regions", value)

    @property
    @pulumi.getter(name="privateDefaultScopeId")
    def private_default_scope_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the IPAM's private scope. A scope is a top-level container in IPAM. Each scope represents an IP-independent network. Scopes enable you to represent networks where you have overlapping IP space. When you create an IPAM, IPAM automatically creates two scopes: public and private. The private scope is intended for private IP space. The public scope is intended for all internet-routable IP space.
        """
        return pulumi.get(self, "private_default_scope_id")

    @private_default_scope_id.setter
    def private_default_scope_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_default_scope_id", value)

    @property
    @pulumi.getter(name="publicDefaultScopeId")
    def public_default_scope_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the IPAM's public scope. A scope is a top-level container in IPAM. Each scope represents an IP-independent network. Scopes enable you to represent networks where you have overlapping IP space. When you create an IPAM, IPAM automatically creates two scopes: public and private. The private scope is intended for private
        IP space. The public scope is intended for all internet-routable IP space.
        """
        return pulumi.get(self, "public_default_scope_id")

    @public_default_scope_id.setter
    def public_default_scope_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_default_scope_id", value)

    @property
    @pulumi.getter(name="scopeCount")
    def scope_count(self) -> Optional[pulumi.Input[int]]:
        """
        The number of scopes in the IPAM.
        """
        return pulumi.get(self, "scope_count")

    @scope_count.setter
    def scope_count(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "scope_count", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        return pulumi.get(self, "tags_all")

    @tags_all.setter
    def tags_all(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "tags_all", value)


class VpcIpam(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 operating_regions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpcIpamOperatingRegionArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        """
        Provides a IPAM resource.

        ## Import

        IPAMs can be imported using the `ipam id`, e.g.

        ```sh
         $ pulumi import aws:ec2/vpcIpam:VpcIpam example ipam-0178368ad2146a492
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: A description for the IPAM.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpcIpamOperatingRegionArgs']]]] operating_regions: Determines which locales can be chosen when you create pools. Locale is the Region where you want to make an IPAM pool available for allocations. You can only create pools with locales that match the operating Regions of the IPAM. You can only create VPCs from a pool whose locale matches the VPC's Region. You specify a region using the region_name parameter. You **must** set your provider block region as an operating_region.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VpcIpamArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Provides a IPAM resource.

        ## Import

        IPAMs can be imported using the `ipam id`, e.g.

        ```sh
         $ pulumi import aws:ec2/vpcIpam:VpcIpam example ipam-0178368ad2146a492
        ```

        :param str resource_name: The name of the resource.
        :param VpcIpamArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VpcIpamArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 operating_regions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpcIpamOperatingRegionArgs']]]]] = None,
                 tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VpcIpamArgs.__new__(VpcIpamArgs)

            __props__.__dict__["description"] = description
            if operating_regions is None and not opts.urn:
                raise TypeError("Missing required property 'operating_regions'")
            __props__.__dict__["operating_regions"] = operating_regions
            __props__.__dict__["tags"] = tags
            __props__.__dict__["arn"] = None
            __props__.__dict__["private_default_scope_id"] = None
            __props__.__dict__["public_default_scope_id"] = None
            __props__.__dict__["scope_count"] = None
            __props__.__dict__["tags_all"] = None
        super(VpcIpam, __self__).__init__(
            'aws:ec2/vpcIpam:VpcIpam',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            arn: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            operating_regions: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpcIpamOperatingRegionArgs']]]]] = None,
            private_default_scope_id: Optional[pulumi.Input[str]] = None,
            public_default_scope_id: Optional[pulumi.Input[str]] = None,
            scope_count: Optional[pulumi.Input[int]] = None,
            tags: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
            tags_all: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None) -> 'VpcIpam':
        """
        Get an existing VpcIpam resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] arn: Amazon Resource Name (ARN) of IPAM
        :param pulumi.Input[str] description: A description for the IPAM.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['VpcIpamOperatingRegionArgs']]]] operating_regions: Determines which locales can be chosen when you create pools. Locale is the Region where you want to make an IPAM pool available for allocations. You can only create pools with locales that match the operating Regions of the IPAM. You can only create VPCs from a pool whose locale matches the VPC's Region. You specify a region using the region_name parameter. You **must** set your provider block region as an operating_region.
        :param pulumi.Input[str] private_default_scope_id: The ID of the IPAM's private scope. A scope is a top-level container in IPAM. Each scope represents an IP-independent network. Scopes enable you to represent networks where you have overlapping IP space. When you create an IPAM, IPAM automatically creates two scopes: public and private. The private scope is intended for private IP space. The public scope is intended for all internet-routable IP space.
        :param pulumi.Input[str] public_default_scope_id: The ID of the IPAM's public scope. A scope is a top-level container in IPAM. Each scope represents an IP-independent network. Scopes enable you to represent networks where you have overlapping IP space. When you create an IPAM, IPAM automatically creates two scopes: public and private. The private scope is intended for private
               IP space. The public scope is intended for all internet-routable IP space.
        :param pulumi.Input[int] scope_count: The number of scopes in the IPAM.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags: A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] tags_all: A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VpcIpamState.__new__(_VpcIpamState)

        __props__.__dict__["arn"] = arn
        __props__.__dict__["description"] = description
        __props__.__dict__["operating_regions"] = operating_regions
        __props__.__dict__["private_default_scope_id"] = private_default_scope_id
        __props__.__dict__["public_default_scope_id"] = public_default_scope_id
        __props__.__dict__["scope_count"] = scope_count
        __props__.__dict__["tags"] = tags
        __props__.__dict__["tags_all"] = tags_all
        return VpcIpam(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def arn(self) -> pulumi.Output[str]:
        """
        Amazon Resource Name (ARN) of IPAM
        """
        return pulumi.get(self, "arn")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        A description for the IPAM.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="operatingRegions")
    def operating_regions(self) -> pulumi.Output[Sequence['outputs.VpcIpamOperatingRegion']]:
        """
        Determines which locales can be chosen when you create pools. Locale is the Region where you want to make an IPAM pool available for allocations. You can only create pools with locales that match the operating Regions of the IPAM. You can only create VPCs from a pool whose locale matches the VPC's Region. You specify a region using the region_name parameter. You **must** set your provider block region as an operating_region.
        """
        return pulumi.get(self, "operating_regions")

    @property
    @pulumi.getter(name="privateDefaultScopeId")
    def private_default_scope_id(self) -> pulumi.Output[str]:
        """
        The ID of the IPAM's private scope. A scope is a top-level container in IPAM. Each scope represents an IP-independent network. Scopes enable you to represent networks where you have overlapping IP space. When you create an IPAM, IPAM automatically creates two scopes: public and private. The private scope is intended for private IP space. The public scope is intended for all internet-routable IP space.
        """
        return pulumi.get(self, "private_default_scope_id")

    @property
    @pulumi.getter(name="publicDefaultScopeId")
    def public_default_scope_id(self) -> pulumi.Output[str]:
        """
        The ID of the IPAM's public scope. A scope is a top-level container in IPAM. Each scope represents an IP-independent network. Scopes enable you to represent networks where you have overlapping IP space. When you create an IPAM, IPAM automatically creates two scopes: public and private. The private scope is intended for private
        IP space. The public scope is intended for all internet-routable IP space.
        """
        return pulumi.get(self, "public_default_scope_id")

    @property
    @pulumi.getter(name="scopeCount")
    def scope_count(self) -> pulumi.Output[int]:
        """
        The number of scopes in the IPAM.
        """
        return pulumi.get(self, "scope_count")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Mapping[str, str]]]:
        """
        A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="tagsAll")
    def tags_all(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A map of tags assigned to the resource, including those inherited from the provider `default_tags` configuration block.
        """
        return pulumi.get(self, "tags_all")

