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

__all__ = ['NsxvIpSetArgs', 'NsxvIpSet']

@pulumi.input_type
class NsxvIpSetArgs:
    def __init__(__self__, *,
                 ip_addresses: pulumi.Input[Sequence[pulumi.Input[str]]],
                 description: Optional[pulumi.Input[str]] = None,
                 is_inheritance_allowed: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 vdc: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a NsxvIpSet resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_addresses: A set of IP addresses, CIDRs and ranges as strings.
        :param pulumi.Input[str] description: An optional description for IP set.
        :param pulumi.Input[bool] is_inheritance_allowed: Toggle to enable inheritance to allow visibility at underlying scopes. Default `true`
        :param pulumi.Input[str] name: Unique IP set name.
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
        :param pulumi.Input[str] vdc: The name of VDC to use, optional if defined at provider level
        """
        pulumi.set(__self__, "ip_addresses", ip_addresses)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if is_inheritance_allowed is not None:
            pulumi.set(__self__, "is_inheritance_allowed", is_inheritance_allowed)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org is not None:
            pulumi.set(__self__, "org", org)
        if vdc is not None:
            pulumi.set(__self__, "vdc", vdc)

    @property
    @pulumi.getter(name="ipAddresses")
    def ip_addresses(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        """
        A set of IP addresses, CIDRs and ranges as strings.
        """
        return pulumi.get(self, "ip_addresses")

    @ip_addresses.setter
    def ip_addresses(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "ip_addresses", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description for IP set.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="isInheritanceAllowed")
    def is_inheritance_allowed(self) -> Optional[pulumi.Input[bool]]:
        """
        Toggle to enable inheritance to allow visibility at underlying scopes. Default `true`
        """
        return pulumi.get(self, "is_inheritance_allowed")

    @is_inheritance_allowed.setter
    def is_inheritance_allowed(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_inheritance_allowed", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Unique IP set name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def org(self) -> Optional[pulumi.Input[str]]:
        """
        The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
        """
        return pulumi.get(self, "org")

    @org.setter
    def org(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org", value)

    @property
    @pulumi.getter
    def vdc(self) -> Optional[pulumi.Input[str]]:
        """
        The name of VDC to use, optional if defined at provider level
        """
        return pulumi.get(self, "vdc")

    @vdc.setter
    def vdc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdc", value)


@pulumi.input_type
class _NsxvIpSetState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 is_inheritance_allowed: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 vdc: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering NsxvIpSet resources.
        :param pulumi.Input[str] description: An optional description for IP set.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_addresses: A set of IP addresses, CIDRs and ranges as strings.
        :param pulumi.Input[bool] is_inheritance_allowed: Toggle to enable inheritance to allow visibility at underlying scopes. Default `true`
        :param pulumi.Input[str] name: Unique IP set name.
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
        :param pulumi.Input[str] vdc: The name of VDC to use, optional if defined at provider level
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ip_addresses is not None:
            pulumi.set(__self__, "ip_addresses", ip_addresses)
        if is_inheritance_allowed is not None:
            pulumi.set(__self__, "is_inheritance_allowed", is_inheritance_allowed)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org is not None:
            pulumi.set(__self__, "org", org)
        if vdc is not None:
            pulumi.set(__self__, "vdc", vdc)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description for IP set.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="ipAddresses")
    def ip_addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A set of IP addresses, CIDRs and ranges as strings.
        """
        return pulumi.get(self, "ip_addresses")

    @ip_addresses.setter
    def ip_addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ip_addresses", value)

    @property
    @pulumi.getter(name="isInheritanceAllowed")
    def is_inheritance_allowed(self) -> Optional[pulumi.Input[bool]]:
        """
        Toggle to enable inheritance to allow visibility at underlying scopes. Default `true`
        """
        return pulumi.get(self, "is_inheritance_allowed")

    @is_inheritance_allowed.setter
    def is_inheritance_allowed(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_inheritance_allowed", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Unique IP set name.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def org(self) -> Optional[pulumi.Input[str]]:
        """
        The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
        """
        return pulumi.get(self, "org")

    @org.setter
    def org(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org", value)

    @property
    @pulumi.getter
    def vdc(self) -> Optional[pulumi.Input[str]]:
        """
        The name of VDC to use, optional if defined at provider level
        """
        return pulumi.get(self, "vdc")

    @vdc.setter
    def vdc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdc", value)


class NsxvIpSet(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 is_inheritance_allowed: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 vdc: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a NsxvIpSet resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: An optional description for IP set.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_addresses: A set of IP addresses, CIDRs and ranges as strings.
        :param pulumi.Input[bool] is_inheritance_allowed: Toggle to enable inheritance to allow visibility at underlying scopes. Default `true`
        :param pulumi.Input[str] name: Unique IP set name.
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
        :param pulumi.Input[str] vdc: The name of VDC to use, optional if defined at provider level
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NsxvIpSetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a NsxvIpSet resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param NsxvIpSetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NsxvIpSetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 is_inheritance_allowed: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 vdc: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NsxvIpSetArgs.__new__(NsxvIpSetArgs)

            __props__.__dict__["description"] = description
            if ip_addresses is None and not opts.urn:
                raise TypeError("Missing required property 'ip_addresses'")
            __props__.__dict__["ip_addresses"] = ip_addresses
            __props__.__dict__["is_inheritance_allowed"] = is_inheritance_allowed
            __props__.__dict__["name"] = name
            __props__.__dict__["org"] = org
            __props__.__dict__["vdc"] = vdc
        super(NsxvIpSet, __self__).__init__(
            'vcd:index/nsxvIpSet:NsxvIpSet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            is_inheritance_allowed: Optional[pulumi.Input[bool]] = None,
            name: Optional[pulumi.Input[str]] = None,
            org: Optional[pulumi.Input[str]] = None,
            vdc: Optional[pulumi.Input[str]] = None) -> 'NsxvIpSet':
        """
        Get an existing NsxvIpSet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: An optional description for IP set.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_addresses: A set of IP addresses, CIDRs and ranges as strings.
        :param pulumi.Input[bool] is_inheritance_allowed: Toggle to enable inheritance to allow visibility at underlying scopes. Default `true`
        :param pulumi.Input[str] name: Unique IP set name.
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
        :param pulumi.Input[str] vdc: The name of VDC to use, optional if defined at provider level
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NsxvIpSetState.__new__(_NsxvIpSetState)

        __props__.__dict__["description"] = description
        __props__.__dict__["ip_addresses"] = ip_addresses
        __props__.__dict__["is_inheritance_allowed"] = is_inheritance_allowed
        __props__.__dict__["name"] = name
        __props__.__dict__["org"] = org
        __props__.__dict__["vdc"] = vdc
        return NsxvIpSet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        An optional description for IP set.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="ipAddresses")
    def ip_addresses(self) -> pulumi.Output[Sequence[str]]:
        """
        A set of IP addresses, CIDRs and ranges as strings.
        """
        return pulumi.get(self, "ip_addresses")

    @property
    @pulumi.getter(name="isInheritanceAllowed")
    def is_inheritance_allowed(self) -> pulumi.Output[Optional[bool]]:
        """
        Toggle to enable inheritance to allow visibility at underlying scopes. Default `true`
        """
        return pulumi.get(self, "is_inheritance_allowed")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Unique IP set name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def org(self) -> pulumi.Output[Optional[str]]:
        """
        The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
        """
        return pulumi.get(self, "org")

    @property
    @pulumi.getter
    def vdc(self) -> pulumi.Output[Optional[str]]:
        """
        The name of VDC to use, optional if defined at provider level
        """
        return pulumi.get(self, "vdc")

