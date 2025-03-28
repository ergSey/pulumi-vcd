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

__all__ = ['ApiTokenArgs', 'ApiToken']

@pulumi.input_type
class ApiTokenArgs:
    def __init__(__self__, *,
                 allow_token_file: pulumi.Input[bool],
                 file_name: pulumi.Input[str],
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ApiToken resource.
        :param pulumi.Input[bool] allow_token_file: An additional check that the user is aware that the file contains
               SENSITIVE information. Must be set to `true` or it will return a validation error.
        :param pulumi.Input[str] file_name: The name of the file which will be created containing the API token
        :param pulumi.Input[str] name: The unique name of the API token for a specific user.
        """
        pulumi.set(__self__, "allow_token_file", allow_token_file)
        pulumi.set(__self__, "file_name", file_name)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="allowTokenFile")
    def allow_token_file(self) -> pulumi.Input[bool]:
        """
        An additional check that the user is aware that the file contains
        SENSITIVE information. Must be set to `true` or it will return a validation error.
        """
        return pulumi.get(self, "allow_token_file")

    @allow_token_file.setter
    def allow_token_file(self, value: pulumi.Input[bool]):
        pulumi.set(self, "allow_token_file", value)

    @property
    @pulumi.getter(name="fileName")
    def file_name(self) -> pulumi.Input[str]:
        """
        The name of the file which will be created containing the API token
        """
        return pulumi.get(self, "file_name")

    @file_name.setter
    def file_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "file_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The unique name of the API token for a specific user.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _ApiTokenState:
    def __init__(__self__, *,
                 allow_token_file: Optional[pulumi.Input[bool]] = None,
                 file_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering ApiToken resources.
        :param pulumi.Input[bool] allow_token_file: An additional check that the user is aware that the file contains
               SENSITIVE information. Must be set to `true` or it will return a validation error.
        :param pulumi.Input[str] file_name: The name of the file which will be created containing the API token
        :param pulumi.Input[str] name: The unique name of the API token for a specific user.
        """
        if allow_token_file is not None:
            pulumi.set(__self__, "allow_token_file", allow_token_file)
        if file_name is not None:
            pulumi.set(__self__, "file_name", file_name)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="allowTokenFile")
    def allow_token_file(self) -> Optional[pulumi.Input[bool]]:
        """
        An additional check that the user is aware that the file contains
        SENSITIVE information. Must be set to `true` or it will return a validation error.
        """
        return pulumi.get(self, "allow_token_file")

    @allow_token_file.setter
    def allow_token_file(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_token_file", value)

    @property
    @pulumi.getter(name="fileName")
    def file_name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the file which will be created containing the API token
        """
        return pulumi.get(self, "file_name")

    @file_name.setter
    def file_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "file_name", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The unique name of the API token for a specific user.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class ApiToken(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_token_file: Optional[pulumi.Input[bool]] = None,
                 file_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a ApiToken resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_token_file: An additional check that the user is aware that the file contains
               SENSITIVE information. Must be set to `true` or it will return a validation error.
        :param pulumi.Input[str] file_name: The name of the file which will be created containing the API token
        :param pulumi.Input[str] name: The unique name of the API token for a specific user.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ApiTokenArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ApiToken resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ApiTokenArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ApiTokenArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_token_file: Optional[pulumi.Input[bool]] = None,
                 file_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ApiTokenArgs.__new__(ApiTokenArgs)

            if allow_token_file is None and not opts.urn:
                raise TypeError("Missing required property 'allow_token_file'")
            __props__.__dict__["allow_token_file"] = allow_token_file
            if file_name is None and not opts.urn:
                raise TypeError("Missing required property 'file_name'")
            __props__.__dict__["file_name"] = file_name
            __props__.__dict__["name"] = name
        super(ApiToken, __self__).__init__(
            'vcd:index/apiToken:ApiToken',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allow_token_file: Optional[pulumi.Input[bool]] = None,
            file_name: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'ApiToken':
        """
        Get an existing ApiToken resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_token_file: An additional check that the user is aware that the file contains
               SENSITIVE information. Must be set to `true` or it will return a validation error.
        :param pulumi.Input[str] file_name: The name of the file which will be created containing the API token
        :param pulumi.Input[str] name: The unique name of the API token for a specific user.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _ApiTokenState.__new__(_ApiTokenState)

        __props__.__dict__["allow_token_file"] = allow_token_file
        __props__.__dict__["file_name"] = file_name
        __props__.__dict__["name"] = name
        return ApiToken(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowTokenFile")
    def allow_token_file(self) -> pulumi.Output[bool]:
        """
        An additional check that the user is aware that the file contains
        SENSITIVE information. Must be set to `true` or it will return a validation error.
        """
        return pulumi.get(self, "allow_token_file")

    @property
    @pulumi.getter(name="fileName")
    def file_name(self) -> pulumi.Output[str]:
        """
        The name of the file which will be created containing the API token
        """
        return pulumi.get(self, "file_name")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The unique name of the API token for a specific user.
        """
        return pulumi.get(self, "name")

