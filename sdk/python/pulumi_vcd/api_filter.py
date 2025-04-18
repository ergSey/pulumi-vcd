# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities

__all__ = ['ApiFilterArgs', 'ApiFilter']

@pulumi.input_type
class ApiFilterArgs:
    def __init__(__self__, *,
                 external_endpoint_id: pulumi.Input[str],
                 url_matcher_pattern: pulumi.Input[str],
                 url_matcher_scope: pulumi.Input[str]):
        """
        The set of arguments for constructing a ApiFilter resource.
        :param pulumi.Input[str] external_endpoint_id: ID of the [External Endpoint](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/external_endpoint) where this API Filter will process the requests to
        :param pulumi.Input[str] url_matcher_pattern: Request URL pattern, written as a regular expression. This argument cannot exceed 1024 characters.
               In most cases, it should end with `.*` (it is like a suffix) which specifies that all the parts of the URL coming after (like parameters) will be redirected to an external endpoint.
               It is important to note that in the case of `url_matcher_scope=EXT_UI_TENANT`, the tenant name is not part of the pattern, it will match the request after the tenant name - if request
               is *"/ext-ui/tenant/testOrg/custom/test"*, the pattern will match against */custom/test*
        :param pulumi.Input[str] url_matcher_scope: Allowed values are `EXT_API`, `EXT_UI_PROVIDER`, `EXT_UI_TENANT` corresponding to
               */ext-api*, */ext-ui/provider*, */ext-ui/tenant/<tenant-name>*
        """
        pulumi.set(__self__, "external_endpoint_id", external_endpoint_id)
        pulumi.set(__self__, "url_matcher_pattern", url_matcher_pattern)
        pulumi.set(__self__, "url_matcher_scope", url_matcher_scope)

    @property
    @pulumi.getter(name="externalEndpointId")
    def external_endpoint_id(self) -> pulumi.Input[str]:
        """
        ID of the [External Endpoint](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/external_endpoint) where this API Filter will process the requests to
        """
        return pulumi.get(self, "external_endpoint_id")

    @external_endpoint_id.setter
    def external_endpoint_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "external_endpoint_id", value)

    @property
    @pulumi.getter(name="urlMatcherPattern")
    def url_matcher_pattern(self) -> pulumi.Input[str]:
        """
        Request URL pattern, written as a regular expression. This argument cannot exceed 1024 characters.
        In most cases, it should end with `.*` (it is like a suffix) which specifies that all the parts of the URL coming after (like parameters) will be redirected to an external endpoint.
        It is important to note that in the case of `url_matcher_scope=EXT_UI_TENANT`, the tenant name is not part of the pattern, it will match the request after the tenant name - if request
        is *"/ext-ui/tenant/testOrg/custom/test"*, the pattern will match against */custom/test*
        """
        return pulumi.get(self, "url_matcher_pattern")

    @url_matcher_pattern.setter
    def url_matcher_pattern(self, value: pulumi.Input[str]):
        pulumi.set(self, "url_matcher_pattern", value)

    @property
    @pulumi.getter(name="urlMatcherScope")
    def url_matcher_scope(self) -> pulumi.Input[str]:
        """
        Allowed values are `EXT_API`, `EXT_UI_PROVIDER`, `EXT_UI_TENANT` corresponding to
        */ext-api*, */ext-ui/provider*, */ext-ui/tenant/<tenant-name>*
        """
        return pulumi.get(self, "url_matcher_scope")

    @url_matcher_scope.setter
    def url_matcher_scope(self, value: pulumi.Input[str]):
        pulumi.set(self, "url_matcher_scope", value)


@pulumi.input_type
class _ApiFilterState:
    def __init__(__self__, *,
                 external_endpoint_id: Optional[pulumi.Input[str]] = None,
                 url_matcher_pattern: Optional[pulumi.Input[str]] = None,
                 url_matcher_scope: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ApiFilter resources.
        :param pulumi.Input[str] external_endpoint_id: ID of the [External Endpoint](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/external_endpoint) where this API Filter will process the requests to
        :param pulumi.Input[str] url_matcher_pattern: Request URL pattern, written as a regular expression. This argument cannot exceed 1024 characters.
               In most cases, it should end with `.*` (it is like a suffix) which specifies that all the parts of the URL coming after (like parameters) will be redirected to an external endpoint.
               It is important to note that in the case of `url_matcher_scope=EXT_UI_TENANT`, the tenant name is not part of the pattern, it will match the request after the tenant name - if request
               is *"/ext-ui/tenant/testOrg/custom/test"*, the pattern will match against */custom/test*
        :param pulumi.Input[str] url_matcher_scope: Allowed values are `EXT_API`, `EXT_UI_PROVIDER`, `EXT_UI_TENANT` corresponding to
               */ext-api*, */ext-ui/provider*, */ext-ui/tenant/<tenant-name>*
        """
        if external_endpoint_id is not None:
            pulumi.set(__self__, "external_endpoint_id", external_endpoint_id)
        if url_matcher_pattern is not None:
            pulumi.set(__self__, "url_matcher_pattern", url_matcher_pattern)
        if url_matcher_scope is not None:
            pulumi.set(__self__, "url_matcher_scope", url_matcher_scope)

    @property
    @pulumi.getter(name="externalEndpointId")
    def external_endpoint_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the [External Endpoint](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/external_endpoint) where this API Filter will process the requests to
        """
        return pulumi.get(self, "external_endpoint_id")

    @external_endpoint_id.setter
    def external_endpoint_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "external_endpoint_id", value)

    @property
    @pulumi.getter(name="urlMatcherPattern")
    def url_matcher_pattern(self) -> Optional[pulumi.Input[str]]:
        """
        Request URL pattern, written as a regular expression. This argument cannot exceed 1024 characters.
        In most cases, it should end with `.*` (it is like a suffix) which specifies that all the parts of the URL coming after (like parameters) will be redirected to an external endpoint.
        It is important to note that in the case of `url_matcher_scope=EXT_UI_TENANT`, the tenant name is not part of the pattern, it will match the request after the tenant name - if request
        is *"/ext-ui/tenant/testOrg/custom/test"*, the pattern will match against */custom/test*
        """
        return pulumi.get(self, "url_matcher_pattern")

    @url_matcher_pattern.setter
    def url_matcher_pattern(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url_matcher_pattern", value)

    @property
    @pulumi.getter(name="urlMatcherScope")
    def url_matcher_scope(self) -> Optional[pulumi.Input[str]]:
        """
        Allowed values are `EXT_API`, `EXT_UI_PROVIDER`, `EXT_UI_TENANT` corresponding to
        */ext-api*, */ext-ui/provider*, */ext-ui/tenant/<tenant-name>*
        """
        return pulumi.get(self, "url_matcher_scope")

    @url_matcher_scope.setter
    def url_matcher_scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "url_matcher_scope", value)


class ApiFilter(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 external_endpoint_id: Optional[pulumi.Input[str]] = None,
                 url_matcher_pattern: Optional[pulumi.Input[str]] = None,
                 url_matcher_scope: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a ApiFilter resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] external_endpoint_id: ID of the [External Endpoint](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/external_endpoint) where this API Filter will process the requests to
        :param pulumi.Input[str] url_matcher_pattern: Request URL pattern, written as a regular expression. This argument cannot exceed 1024 characters.
               In most cases, it should end with `.*` (it is like a suffix) which specifies that all the parts of the URL coming after (like parameters) will be redirected to an external endpoint.
               It is important to note that in the case of `url_matcher_scope=EXT_UI_TENANT`, the tenant name is not part of the pattern, it will match the request after the tenant name - if request
               is *"/ext-ui/tenant/testOrg/custom/test"*, the pattern will match against */custom/test*
        :param pulumi.Input[str] url_matcher_scope: Allowed values are `EXT_API`, `EXT_UI_PROVIDER`, `EXT_UI_TENANT` corresponding to
               */ext-api*, */ext-ui/provider*, */ext-ui/tenant/<tenant-name>*
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApiFilterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ApiFilter resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ApiFilterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApiFilterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 external_endpoint_id: Optional[pulumi.Input[str]] = None,
                 url_matcher_pattern: Optional[pulumi.Input[str]] = None,
                 url_matcher_scope: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApiFilterArgs.__new__(ApiFilterArgs)

            if external_endpoint_id is None and not opts.urn:
                raise TypeError("Missing required property 'external_endpoint_id'")
            __props__.__dict__["external_endpoint_id"] = external_endpoint_id
            if url_matcher_pattern is None and not opts.urn:
                raise TypeError("Missing required property 'url_matcher_pattern'")
            __props__.__dict__["url_matcher_pattern"] = url_matcher_pattern
            if url_matcher_scope is None and not opts.urn:
                raise TypeError("Missing required property 'url_matcher_scope'")
            __props__.__dict__["url_matcher_scope"] = url_matcher_scope
        super(ApiFilter, __self__).__init__(
            'vcd:index/apiFilter:ApiFilter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            external_endpoint_id: Optional[pulumi.Input[str]] = None,
            url_matcher_pattern: Optional[pulumi.Input[str]] = None,
            url_matcher_scope: Optional[pulumi.Input[str]] = None) -> 'ApiFilter':
        """
        Get an existing ApiFilter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] external_endpoint_id: ID of the [External Endpoint](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/external_endpoint) where this API Filter will process the requests to
        :param pulumi.Input[str] url_matcher_pattern: Request URL pattern, written as a regular expression. This argument cannot exceed 1024 characters.
               In most cases, it should end with `.*` (it is like a suffix) which specifies that all the parts of the URL coming after (like parameters) will be redirected to an external endpoint.
               It is important to note that in the case of `url_matcher_scope=EXT_UI_TENANT`, the tenant name is not part of the pattern, it will match the request after the tenant name - if request
               is *"/ext-ui/tenant/testOrg/custom/test"*, the pattern will match against */custom/test*
        :param pulumi.Input[str] url_matcher_scope: Allowed values are `EXT_API`, `EXT_UI_PROVIDER`, `EXT_UI_TENANT` corresponding to
               */ext-api*, */ext-ui/provider*, */ext-ui/tenant/<tenant-name>*
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApiFilterState.__new__(_ApiFilterState)

        __props__.__dict__["external_endpoint_id"] = external_endpoint_id
        __props__.__dict__["url_matcher_pattern"] = url_matcher_pattern
        __props__.__dict__["url_matcher_scope"] = url_matcher_scope
        return ApiFilter(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="externalEndpointId")
    def external_endpoint_id(self) -> pulumi.Output[str]:
        """
        ID of the [External Endpoint](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/external_endpoint) where this API Filter will process the requests to
        """
        return pulumi.get(self, "external_endpoint_id")

    @property
    @pulumi.getter(name="urlMatcherPattern")
    def url_matcher_pattern(self) -> pulumi.Output[str]:
        """
        Request URL pattern, written as a regular expression. This argument cannot exceed 1024 characters.
        In most cases, it should end with `.*` (it is like a suffix) which specifies that all the parts of the URL coming after (like parameters) will be redirected to an external endpoint.
        It is important to note that in the case of `url_matcher_scope=EXT_UI_TENANT`, the tenant name is not part of the pattern, it will match the request after the tenant name - if request
        is *"/ext-ui/tenant/testOrg/custom/test"*, the pattern will match against */custom/test*
        """
        return pulumi.get(self, "url_matcher_pattern")

    @property
    @pulumi.getter(name="urlMatcherScope")
    def url_matcher_scope(self) -> pulumi.Output[str]:
        """
        Allowed values are `EXT_API`, `EXT_UI_PROVIDER`, `EXT_UI_TENANT` corresponding to
        */ext-api*, */ext-ui/provider*, */ext-ui/tenant/<tenant-name>*
        """
        return pulumi.get(self, "url_matcher_scope")

