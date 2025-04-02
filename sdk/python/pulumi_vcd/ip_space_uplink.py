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

__all__ = ['IpSpaceUplinkArgs', 'IpSpaceUplink']

@pulumi.input_type
class IpSpaceUplinkArgs:
    def __init__(__self__, *,
                 external_network_id: pulumi.Input[str],
                 ip_space_id: pulumi.Input[str],
                 associated_interface_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a IpSpaceUplink resource.
        :param pulumi.Input[str] external_network_id: External Network ID For IP Space Uplink configuration
        :param pulumi.Input[str] ip_space_id: IP Space ID configuration
        :param pulumi.Input[Sequence[pulumi.Input[str]]] associated_interface_ids: A set of Tier-0 Router Interfaces to associate with this uplink
        :param pulumi.Input[str] description: An optional description for IP Space Uplink
        :param pulumi.Input[str] name: A tenant facing name for IP Space Uplink
        """
        pulumi.set(__self__, "external_network_id", external_network_id)
        pulumi.set(__self__, "ip_space_id", ip_space_id)
        if associated_interface_ids is not None:
            pulumi.set(__self__, "associated_interface_ids", associated_interface_ids)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="externalNetworkId")
    def external_network_id(self) -> pulumi.Input[str]:
        """
        External Network ID For IP Space Uplink configuration
        """
        return pulumi.get(self, "external_network_id")

    @external_network_id.setter
    def external_network_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "external_network_id", value)

    @property
    @pulumi.getter(name="ipSpaceId")
    def ip_space_id(self) -> pulumi.Input[str]:
        """
        IP Space ID configuration
        """
        return pulumi.get(self, "ip_space_id")

    @ip_space_id.setter
    def ip_space_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "ip_space_id", value)

    @property
    @pulumi.getter(name="associatedInterfaceIds")
    def associated_interface_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A set of Tier-0 Router Interfaces to associate with this uplink
        """
        return pulumi.get(self, "associated_interface_ids")

    @associated_interface_ids.setter
    def associated_interface_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "associated_interface_ids", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description for IP Space Uplink
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A tenant facing name for IP Space Uplink
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _IpSpaceUplinkState:
    def __init__(__self__, *,
                 associated_interface_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 external_network_id: Optional[pulumi.Input[str]] = None,
                 ip_space_id: Optional[pulumi.Input[str]] = None,
                 ip_space_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 status: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IpSpaceUplink resources.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] associated_interface_ids: A set of Tier-0 Router Interfaces to associate with this uplink
        :param pulumi.Input[str] description: An optional description for IP Space Uplink
        :param pulumi.Input[str] external_network_id: External Network ID For IP Space Uplink configuration
        :param pulumi.Input[str] ip_space_id: IP Space ID configuration
        :param pulumi.Input[str] ip_space_type: Backing IP Space type
        :param pulumi.Input[str] name: A tenant facing name for IP Space Uplink
        :param pulumi.Input[str] status: Status of IP Space Uplink
        """
        if associated_interface_ids is not None:
            pulumi.set(__self__, "associated_interface_ids", associated_interface_ids)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if external_network_id is not None:
            pulumi.set(__self__, "external_network_id", external_network_id)
        if ip_space_id is not None:
            pulumi.set(__self__, "ip_space_id", ip_space_id)
        if ip_space_type is not None:
            pulumi.set(__self__, "ip_space_type", ip_space_type)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if status is not None:
            pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="associatedInterfaceIds")
    def associated_interface_ids(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]:
        """
        A set of Tier-0 Router Interfaces to associate with this uplink
        """
        return pulumi.get(self, "associated_interface_ids")

    @associated_interface_ids.setter
    def associated_interface_ids(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]]):
        pulumi.set(self, "associated_interface_ids", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description for IP Space Uplink
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="externalNetworkId")
    def external_network_id(self) -> Optional[pulumi.Input[str]]:
        """
        External Network ID For IP Space Uplink configuration
        """
        return pulumi.get(self, "external_network_id")

    @external_network_id.setter
    def external_network_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "external_network_id", value)

    @property
    @pulumi.getter(name="ipSpaceId")
    def ip_space_id(self) -> Optional[pulumi.Input[str]]:
        """
        IP Space ID configuration
        """
        return pulumi.get(self, "ip_space_id")

    @ip_space_id.setter
    def ip_space_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_space_id", value)

    @property
    @pulumi.getter(name="ipSpaceType")
    def ip_space_type(self) -> Optional[pulumi.Input[str]]:
        """
        Backing IP Space type
        """
        return pulumi.get(self, "ip_space_type")

    @ip_space_type.setter
    def ip_space_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_space_type", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A tenant facing name for IP Space Uplink
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def status(self) -> Optional[pulumi.Input[str]]:
        """
        Status of IP Space Uplink
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status", value)


class IpSpaceUplink(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 associated_interface_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 external_network_id: Optional[pulumi.Input[str]] = None,
                 ip_space_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a IpSpaceUplink resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] associated_interface_ids: A set of Tier-0 Router Interfaces to associate with this uplink
        :param pulumi.Input[str] description: An optional description for IP Space Uplink
        :param pulumi.Input[str] external_network_id: External Network ID For IP Space Uplink configuration
        :param pulumi.Input[str] ip_space_id: IP Space ID configuration
        :param pulumi.Input[str] name: A tenant facing name for IP Space Uplink
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IpSpaceUplinkArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a IpSpaceUplink resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param IpSpaceUplinkArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IpSpaceUplinkArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 associated_interface_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 external_network_id: Optional[pulumi.Input[str]] = None,
                 ip_space_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IpSpaceUplinkArgs.__new__(IpSpaceUplinkArgs)

            __props__.__dict__["associated_interface_ids"] = associated_interface_ids
            __props__.__dict__["description"] = description
            if external_network_id is None and not opts.urn:
                raise TypeError("Missing required property 'external_network_id'")
            __props__.__dict__["external_network_id"] = external_network_id
            if ip_space_id is None and not opts.urn:
                raise TypeError("Missing required property 'ip_space_id'")
            __props__.__dict__["ip_space_id"] = ip_space_id
            __props__.__dict__["name"] = name
            __props__.__dict__["ip_space_type"] = None
            __props__.__dict__["status"] = None
        super(IpSpaceUplink, __self__).__init__(
            'vcd:index/ipSpaceUplink:IpSpaceUplink',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            associated_interface_ids: Optional[pulumi.Input[Sequence[pulumi.Input[str]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            external_network_id: Optional[pulumi.Input[str]] = None,
            ip_space_id: Optional[pulumi.Input[str]] = None,
            ip_space_type: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            status: Optional[pulumi.Input[str]] = None) -> 'IpSpaceUplink':
        """
        Get an existing IpSpaceUplink resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[str]]] associated_interface_ids: A set of Tier-0 Router Interfaces to associate with this uplink
        :param pulumi.Input[str] description: An optional description for IP Space Uplink
        :param pulumi.Input[str] external_network_id: External Network ID For IP Space Uplink configuration
        :param pulumi.Input[str] ip_space_id: IP Space ID configuration
        :param pulumi.Input[str] ip_space_type: Backing IP Space type
        :param pulumi.Input[str] name: A tenant facing name for IP Space Uplink
        :param pulumi.Input[str] status: Status of IP Space Uplink
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IpSpaceUplinkState.__new__(_IpSpaceUplinkState)

        __props__.__dict__["associated_interface_ids"] = associated_interface_ids
        __props__.__dict__["description"] = description
        __props__.__dict__["external_network_id"] = external_network_id
        __props__.__dict__["ip_space_id"] = ip_space_id
        __props__.__dict__["ip_space_type"] = ip_space_type
        __props__.__dict__["name"] = name
        __props__.__dict__["status"] = status
        return IpSpaceUplink(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="associatedInterfaceIds")
    def associated_interface_ids(self) -> pulumi.Output[Optional[Sequence[str]]]:
        """
        A set of Tier-0 Router Interfaces to associate with this uplink
        """
        return pulumi.get(self, "associated_interface_ids")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        An optional description for IP Space Uplink
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="externalNetworkId")
    def external_network_id(self) -> pulumi.Output[str]:
        """
        External Network ID For IP Space Uplink configuration
        """
        return pulumi.get(self, "external_network_id")

    @property
    @pulumi.getter(name="ipSpaceId")
    def ip_space_id(self) -> pulumi.Output[str]:
        """
        IP Space ID configuration
        """
        return pulumi.get(self, "ip_space_id")

    @property
    @pulumi.getter(name="ipSpaceType")
    def ip_space_type(self) -> pulumi.Output[str]:
        """
        Backing IP Space type
        """
        return pulumi.get(self, "ip_space_type")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A tenant facing name for IP Space Uplink
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output[str]:
        """
        Status of IP Space Uplink
        """
        return pulumi.get(self, "status")

