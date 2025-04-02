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

__all__ = ['NsxtIpSetArgs', 'NsxtIpSet']

@pulumi.input_type
class NsxtIpSetArgs:
    def __init__(__self__, *,
                 edge_gateway_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 vdc: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a NsxtIpSet resource.
        :param pulumi.Input[str] edge_gateway_id: The ID of the Edge Gateway (NSX-T only). Can be looked up using
               `NsxtEdgegateway` data source.
        :param pulumi.Input[str] description: An optional description of the IP Set
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_addresses: A set of IP addresses, subnets or ranges (IPv4 or IPv6)
        :param pulumi.Input[str] name: A unique name for IP Set
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful
               when connected as sysadmin working across different organisations.
        :param pulumi.Input[str] vdc: The name of VDC to use, optional if defined at provider level. **Deprecated**
               in favor of `edge_gateway_id` field.
        """
        pulumi.set(__self__, "edge_gateway_id", edge_gateway_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ip_addresses is not None:
            pulumi.set(__self__, "ip_addresses", ip_addresses)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org is not None:
            pulumi.set(__self__, "org", org)
        if vdc is not None:
            warnings.warn("""Deprecated in favor of `edge_gateway_id`. IP Sets will inherit VDC from parent Edge Gateway.""", DeprecationWarning)
            pulumi.log.warn("""vdc is deprecated: Deprecated in favor of `edge_gateway_id`. IP Sets will inherit VDC from parent Edge Gateway.""")
        if vdc is not None:
            pulumi.set(__self__, "vdc", vdc)

    @property
    @pulumi.getter(name="edgeGatewayId")
    def edge_gateway_id(self) -> pulumi.Input[str]:
        """
        The ID of the Edge Gateway (NSX-T only). Can be looked up using
        `NsxtEdgegateway` data source.
        """
        return pulumi.get(self, "edge_gateway_id")

    @edge_gateway_id.setter
    def edge_gateway_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "edge_gateway_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of the IP Set
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="ipAddresses")
    def ip_addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A set of IP addresses, subnets or ranges (IPv4 or IPv6)
        """
        return pulumi.get(self, "ip_addresses")

    @ip_addresses.setter
    def ip_addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ip_addresses", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A unique name for IP Set
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def org(self) -> Optional[pulumi.Input[str]]:
        """
        The name of organization to use, optional if defined at provider level. Useful
        when connected as sysadmin working across different organisations.
        """
        return pulumi.get(self, "org")

    @org.setter
    def org(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org", value)

    @property
    @pulumi.getter
    @_utilities.deprecated("""Deprecated in favor of `edge_gateway_id`. IP Sets will inherit VDC from parent Edge Gateway.""")
    def vdc(self) -> Optional[pulumi.Input[str]]:
        """
        The name of VDC to use, optional if defined at provider level. **Deprecated**
        in favor of `edge_gateway_id` field.
        """
        return pulumi.get(self, "vdc")

    @vdc.setter
    def vdc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdc", value)


@pulumi.input_type
class _NsxtIpSetState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 edge_gateway_id: Optional[pulumi.Input[str]] = None,
                 ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 owner_id: Optional[pulumi.Input[str]] = None,
                 vdc: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering NsxtIpSet resources.
        :param pulumi.Input[str] description: An optional description of the IP Set
        :param pulumi.Input[str] edge_gateway_id: The ID of the Edge Gateway (NSX-T only). Can be looked up using
               `NsxtEdgegateway` data source.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_addresses: A set of IP addresses, subnets or ranges (IPv4 or IPv6)
        :param pulumi.Input[str] name: A unique name for IP Set
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful
               when connected as sysadmin working across different organisations.
        :param pulumi.Input[str] owner_id: ID of VDC or VDC Group
        :param pulumi.Input[str] vdc: The name of VDC to use, optional if defined at provider level. **Deprecated**
               in favor of `edge_gateway_id` field.
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if edge_gateway_id is not None:
            pulumi.set(__self__, "edge_gateway_id", edge_gateway_id)
        if ip_addresses is not None:
            pulumi.set(__self__, "ip_addresses", ip_addresses)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org is not None:
            pulumi.set(__self__, "org", org)
        if owner_id is not None:
            pulumi.set(__self__, "owner_id", owner_id)
        if vdc is not None:
            warnings.warn("""Deprecated in favor of `edge_gateway_id`. IP Sets will inherit VDC from parent Edge Gateway.""", DeprecationWarning)
            pulumi.log.warn("""vdc is deprecated: Deprecated in favor of `edge_gateway_id`. IP Sets will inherit VDC from parent Edge Gateway.""")
        if vdc is not None:
            pulumi.set(__self__, "vdc", vdc)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description of the IP Set
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="edgeGatewayId")
    def edge_gateway_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the Edge Gateway (NSX-T only). Can be looked up using
        `NsxtEdgegateway` data source.
        """
        return pulumi.get(self, "edge_gateway_id")

    @edge_gateway_id.setter
    def edge_gateway_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "edge_gateway_id", value)

    @property
    @pulumi.getter(name="ipAddresses")
    def ip_addresses(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A set of IP addresses, subnets or ranges (IPv4 or IPv6)
        """
        return pulumi.get(self, "ip_addresses")

    @ip_addresses.setter
    def ip_addresses(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "ip_addresses", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A unique name for IP Set
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def org(self) -> Optional[pulumi.Input[str]]:
        """
        The name of organization to use, optional if defined at provider level. Useful
        when connected as sysadmin working across different organisations.
        """
        return pulumi.get(self, "org")

    @org.setter
    def org(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org", value)

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of VDC or VDC Group
        """
        return pulumi.get(self, "owner_id")

    @owner_id.setter
    def owner_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "owner_id", value)

    @property
    @pulumi.getter
    @_utilities.deprecated("""Deprecated in favor of `edge_gateway_id`. IP Sets will inherit VDC from parent Edge Gateway.""")
    def vdc(self) -> Optional[pulumi.Input[str]]:
        """
        The name of VDC to use, optional if defined at provider level. **Deprecated**
        in favor of `edge_gateway_id` field.
        """
        return pulumi.get(self, "vdc")

    @vdc.setter
    def vdc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdc", value)


class NsxtIpSet(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 edge_gateway_id: Optional[pulumi.Input[str]] = None,
                 ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 vdc: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a NsxtIpSet resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: An optional description of the IP Set
        :param pulumi.Input[str] edge_gateway_id: The ID of the Edge Gateway (NSX-T only). Can be looked up using
               `NsxtEdgegateway` data source.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_addresses: A set of IP addresses, subnets or ranges (IPv4 or IPv6)
        :param pulumi.Input[str] name: A unique name for IP Set
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful
               when connected as sysadmin working across different organisations.
        :param pulumi.Input[str] vdc: The name of VDC to use, optional if defined at provider level. **Deprecated**
               in favor of `edge_gateway_id` field.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NsxtIpSetArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a NsxtIpSet resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param NsxtIpSetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NsxtIpSetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 edge_gateway_id: Optional[pulumi.Input[str]] = None,
                 ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
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
            __props__ = NsxtIpSetArgs.__new__(NsxtIpSetArgs)

            __props__.__dict__["description"] = description
            if edge_gateway_id is None and not opts.urn:
                raise TypeError("Missing required property 'edge_gateway_id'")
            __props__.__dict__["edge_gateway_id"] = edge_gateway_id
            __props__.__dict__["ip_addresses"] = ip_addresses
            __props__.__dict__["name"] = name
            __props__.__dict__["org"] = org
            __props__.__dict__["vdc"] = vdc
            __props__.__dict__["owner_id"] = None
        super(NsxtIpSet, __self__).__init__(
            'vcd:index/nsxtIpSet:NsxtIpSet',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            edge_gateway_id: Optional[pulumi.Input[str]] = None,
            ip_addresses: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            org: Optional[pulumi.Input[str]] = None,
            owner_id: Optional[pulumi.Input[str]] = None,
            vdc: Optional[pulumi.Input[str]] = None) -> 'NsxtIpSet':
        """
        Get an existing NsxtIpSet resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: An optional description of the IP Set
        :param pulumi.Input[str] edge_gateway_id: The ID of the Edge Gateway (NSX-T only). Can be looked up using
               `NsxtEdgegateway` data source.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] ip_addresses: A set of IP addresses, subnets or ranges (IPv4 or IPv6)
        :param pulumi.Input[str] name: A unique name for IP Set
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful
               when connected as sysadmin working across different organisations.
        :param pulumi.Input[str] owner_id: ID of VDC or VDC Group
        :param pulumi.Input[str] vdc: The name of VDC to use, optional if defined at provider level. **Deprecated**
               in favor of `edge_gateway_id` field.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NsxtIpSetState.__new__(_NsxtIpSetState)

        __props__.__dict__["description"] = description
        __props__.__dict__["edge_gateway_id"] = edge_gateway_id
        __props__.__dict__["ip_addresses"] = ip_addresses
        __props__.__dict__["name"] = name
        __props__.__dict__["org"] = org
        __props__.__dict__["owner_id"] = owner_id
        __props__.__dict__["vdc"] = vdc
        return NsxtIpSet(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        An optional description of the IP Set
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="edgeGatewayId")
    def edge_gateway_id(self) -> pulumi.Output[str]:
        """
        The ID of the Edge Gateway (NSX-T only). Can be looked up using
        `NsxtEdgegateway` data source.
        """
        return pulumi.get(self, "edge_gateway_id")

    @property
    @pulumi.getter(name="ipAddresses")
    def ip_addresses(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A set of IP addresses, subnets or ranges (IPv4 or IPv6)
        """
        return pulumi.get(self, "ip_addresses")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A unique name for IP Set
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def org(self) -> pulumi.Output[Optional[str]]:
        """
        The name of organization to use, optional if defined at provider level. Useful
        when connected as sysadmin working across different organisations.
        """
        return pulumi.get(self, "org")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> pulumi.Output[str]:
        """
        ID of VDC or VDC Group
        """
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter
    @_utilities.deprecated("""Deprecated in favor of `edge_gateway_id`. IP Sets will inherit VDC from parent Edge Gateway.""")
    def vdc(self) -> pulumi.Output[str]:
        """
        The name of VDC to use, optional if defined at provider level. **Deprecated**
        in favor of `edge_gateway_id` field.
        """
        return pulumi.get(self, "vdc")

