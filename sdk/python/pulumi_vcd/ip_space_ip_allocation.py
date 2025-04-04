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

__all__ = ['IpSpaceIpAllocationArgs', 'IpSpaceIpAllocation']

@pulumi.input_type
class IpSpaceIpAllocationArgs:
    def __init__(__self__, *,
                 org_id: pulumi.Input[str],
                 type: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 ip_space_id: Optional[pulumi.Input[str]] = None,
                 prefix_length: Optional[pulumi.Input[str]] = None,
                 usage_state: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a IpSpaceIpAllocation resource.
        :param pulumi.Input[str] org_id: Org ID in which the IP is allocated
        :param pulumi.Input[str] type: One of `FLOATING_IP`, `IP_PREFIX`
               * `FLOATING_IP` - allocates single IP from defined ranges in IP Space
               * `IP_PREFIX` - allocates subnets. **Note** field `prefix_length` is required to allocate IP
               Prefix
        :param pulumi.Input[str] description: Can only be set when `usage_state=USED_MANUAL`
               
               > IP Allocation resources can be created only if there is a NSX-T Edge Gateway
               (`NsxtEdgegateway`) that is backed by the Provider Gateway (`ExternalNetworkV2`) with IP
               Space Uplinks (`IpSpaceUplink`). Attempting to allocate IP Addresses before having an
               Edge Gateway withing VDC will return errors of type `This operation is denied`.
        :param pulumi.Input[str] ip_space_id: IP Space ID to use for IP Allocations
        :param pulumi.Input[str] prefix_length: Required when `type=IP_PREFIX`
        :param pulumi.Input[str] usage_state: (Optional) Only used with manual reservations. Value `USED_MANUAL`
               enables manual IP reservation. Value `UNUSED` is set to release manual allocation of IP.
        :param pulumi.Input[str] value: An option to request a specific IP or subnet from IP Space.
               **Note:** This field does not support IP ranges because it would cause multiple allocations
               created in one resource. Please use multiple resource instances to allocate IP ranges.
        """
        pulumi.set(__self__, "org_id", org_id)
        pulumi.set(__self__, "type", type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ip_space_id is not None:
            pulumi.set(__self__, "ip_space_id", ip_space_id)
        if prefix_length is not None:
            pulumi.set(__self__, "prefix_length", prefix_length)
        if usage_state is not None:
            pulumi.set(__self__, "usage_state", usage_state)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Input[str]:
        """
        Org ID in which the IP is allocated
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "org_id", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        One of `FLOATING_IP`, `IP_PREFIX`
        * `FLOATING_IP` - allocates single IP from defined ranges in IP Space
        * `IP_PREFIX` - allocates subnets. **Note** field `prefix_length` is required to allocate IP
        Prefix
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Can only be set when `usage_state=USED_MANUAL`

        > IP Allocation resources can be created only if there is a NSX-T Edge Gateway
        (`NsxtEdgegateway`) that is backed by the Provider Gateway (`ExternalNetworkV2`) with IP
        Space Uplinks (`IpSpaceUplink`). Attempting to allocate IP Addresses before having an
        Edge Gateway withing VDC will return errors of type `This operation is denied`.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="ipSpaceId")
    def ip_space_id(self) -> Optional[pulumi.Input[str]]:
        """
        IP Space ID to use for IP Allocations
        """
        return pulumi.get(self, "ip_space_id")

    @ip_space_id.setter
    def ip_space_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_space_id", value)

    @property
    @pulumi.getter(name="prefixLength")
    def prefix_length(self) -> Optional[pulumi.Input[str]]:
        """
        Required when `type=IP_PREFIX`
        """
        return pulumi.get(self, "prefix_length")

    @prefix_length.setter
    def prefix_length(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prefix_length", value)

    @property
    @pulumi.getter(name="usageState")
    def usage_state(self) -> Optional[pulumi.Input[str]]:
        """
        (Optional) Only used with manual reservations. Value `USED_MANUAL`
        enables manual IP reservation. Value `UNUSED` is set to release manual allocation of IP.
        """
        return pulumi.get(self, "usage_state")

    @usage_state.setter
    def usage_state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "usage_state", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        An option to request a specific IP or subnet from IP Space.
        **Note:** This field does not support IP ranges because it would cause multiple allocations
        created in one resource. Please use multiple resource instances to allocate IP ranges.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class _IpSpaceIpAllocationState:
    def __init__(__self__, *,
                 allocation_date: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ip: Optional[pulumi.Input[str]] = None,
                 ip_address: Optional[pulumi.Input[str]] = None,
                 ip_space_id: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 prefix_length: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 usage_state: Optional[pulumi.Input[str]] = None,
                 used_by_id: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering IpSpaceIpAllocation resources.
        :param pulumi.Input[str] allocation_date: allocation date in formated as `2023-06-07T09:57:58.721Z` (ISO 8601)
        :param pulumi.Input[str] description: Can only be set when `usage_state=USED_MANUAL`
               
               > IP Allocation resources can be created only if there is a NSX-T Edge Gateway
               (`NsxtEdgegateway`) that is backed by the Provider Gateway (`ExternalNetworkV2`) with IP
               Space Uplinks (`IpSpaceUplink`). Attempting to allocate IP Addresses before having an
               Edge Gateway withing VDC will return errors of type `This operation is denied`.
        :param pulumi.Input[str] ip: convenience field. For `type=IP_PREFIX` it will contain only the IP from CIDR returned
        :param pulumi.Input[str] ip_address: IP address or CIDR
        :param pulumi.Input[str] ip_space_id: IP Space ID to use for IP Allocations
        :param pulumi.Input[str] org_id: Org ID in which the IP is allocated
        :param pulumi.Input[str] prefix_length: Required when `type=IP_PREFIX`
        :param pulumi.Input[str] type: One of `FLOATING_IP`, `IP_PREFIX`
               * `FLOATING_IP` - allocates single IP from defined ranges in IP Space
               * `IP_PREFIX` - allocates subnets. **Note** field `prefix_length` is required to allocate IP
               Prefix
        :param pulumi.Input[str] usage_state: (Optional) Only used with manual reservations. Value `USED_MANUAL`
               enables manual IP reservation. Value `UNUSED` is set to release manual allocation of IP.
        :param pulumi.Input[str] used_by_id: contains entity ID that is using the IP if `usage_state=USED`
        :param pulumi.Input[str] value: An option to request a specific IP or subnet from IP Space.
               **Note:** This field does not support IP ranges because it would cause multiple allocations
               created in one resource. Please use multiple resource instances to allocate IP ranges.
        """
        if allocation_date is not None:
            pulumi.set(__self__, "allocation_date", allocation_date)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ip is not None:
            pulumi.set(__self__, "ip", ip)
        if ip_address is not None:
            pulumi.set(__self__, "ip_address", ip_address)
        if ip_space_id is not None:
            pulumi.set(__self__, "ip_space_id", ip_space_id)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)
        if prefix_length is not None:
            pulumi.set(__self__, "prefix_length", prefix_length)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if usage_state is not None:
            pulumi.set(__self__, "usage_state", usage_state)
        if used_by_id is not None:
            pulumi.set(__self__, "used_by_id", used_by_id)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter(name="allocationDate")
    def allocation_date(self) -> Optional[pulumi.Input[str]]:
        """
        allocation date in formated as `2023-06-07T09:57:58.721Z` (ISO 8601)
        """
        return pulumi.get(self, "allocation_date")

    @allocation_date.setter
    def allocation_date(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "allocation_date", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Can only be set when `usage_state=USED_MANUAL`

        > IP Allocation resources can be created only if there is a NSX-T Edge Gateway
        (`NsxtEdgegateway`) that is backed by the Provider Gateway (`ExternalNetworkV2`) with IP
        Space Uplinks (`IpSpaceUplink`). Attempting to allocate IP Addresses before having an
        Edge Gateway withing VDC will return errors of type `This operation is denied`.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def ip(self) -> Optional[pulumi.Input[str]]:
        """
        convenience field. For `type=IP_PREFIX` it will contain only the IP from CIDR returned
        """
        return pulumi.get(self, "ip")

    @ip.setter
    def ip(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip", value)

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> Optional[pulumi.Input[str]]:
        """
        IP address or CIDR
        """
        return pulumi.get(self, "ip_address")

    @ip_address.setter
    def ip_address(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_address", value)

    @property
    @pulumi.getter(name="ipSpaceId")
    def ip_space_id(self) -> Optional[pulumi.Input[str]]:
        """
        IP Space ID to use for IP Allocations
        """
        return pulumi.get(self, "ip_space_id")

    @ip_space_id.setter
    def ip_space_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ip_space_id", value)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[pulumi.Input[str]]:
        """
        Org ID in which the IP is allocated
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org_id", value)

    @property
    @pulumi.getter(name="prefixLength")
    def prefix_length(self) -> Optional[pulumi.Input[str]]:
        """
        Required when `type=IP_PREFIX`
        """
        return pulumi.get(self, "prefix_length")

    @prefix_length.setter
    def prefix_length(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "prefix_length", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        One of `FLOATING_IP`, `IP_PREFIX`
        * `FLOATING_IP` - allocates single IP from defined ranges in IP Space
        * `IP_PREFIX` - allocates subnets. **Note** field `prefix_length` is required to allocate IP
        Prefix
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="usageState")
    def usage_state(self) -> Optional[pulumi.Input[str]]:
        """
        (Optional) Only used with manual reservations. Value `USED_MANUAL`
        enables manual IP reservation. Value `UNUSED` is set to release manual allocation of IP.
        """
        return pulumi.get(self, "usage_state")

    @usage_state.setter
    def usage_state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "usage_state", value)

    @property
    @pulumi.getter(name="usedById")
    def used_by_id(self) -> Optional[pulumi.Input[str]]:
        """
        contains entity ID that is using the IP if `usage_state=USED`
        """
        return pulumi.get(self, "used_by_id")

    @used_by_id.setter
    def used_by_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "used_by_id", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        """
        An option to request a specific IP or subnet from IP Space.
        **Note:** This field does not support IP ranges because it would cause multiple allocations
        created in one resource. Please use multiple resource instances to allocate IP ranges.
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


class IpSpaceIpAllocation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ip_space_id: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 prefix_length: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 usage_state: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a IpSpaceIpAllocation resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Can only be set when `usage_state=USED_MANUAL`
               
               > IP Allocation resources can be created only if there is a NSX-T Edge Gateway
               (`NsxtEdgegateway`) that is backed by the Provider Gateway (`ExternalNetworkV2`) with IP
               Space Uplinks (`IpSpaceUplink`). Attempting to allocate IP Addresses before having an
               Edge Gateway withing VDC will return errors of type `This operation is denied`.
        :param pulumi.Input[str] ip_space_id: IP Space ID to use for IP Allocations
        :param pulumi.Input[str] org_id: Org ID in which the IP is allocated
        :param pulumi.Input[str] prefix_length: Required when `type=IP_PREFIX`
        :param pulumi.Input[str] type: One of `FLOATING_IP`, `IP_PREFIX`
               * `FLOATING_IP` - allocates single IP from defined ranges in IP Space
               * `IP_PREFIX` - allocates subnets. **Note** field `prefix_length` is required to allocate IP
               Prefix
        :param pulumi.Input[str] usage_state: (Optional) Only used with manual reservations. Value `USED_MANUAL`
               enables manual IP reservation. Value `UNUSED` is set to release manual allocation of IP.
        :param pulumi.Input[str] value: An option to request a specific IP or subnet from IP Space.
               **Note:** This field does not support IP ranges because it would cause multiple allocations
               created in one resource. Please use multiple resource instances to allocate IP ranges.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: IpSpaceIpAllocationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a IpSpaceIpAllocation resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param IpSpaceIpAllocationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(IpSpaceIpAllocationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ip_space_id: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 prefix_length: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 usage_state: Optional[pulumi.Input[str]] = None,
                 value: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = IpSpaceIpAllocationArgs.__new__(IpSpaceIpAllocationArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["ip_space_id"] = ip_space_id
            if org_id is None and not opts.urn:
                raise TypeError("Missing required property 'org_id'")
            __props__.__dict__["org_id"] = org_id
            __props__.__dict__["prefix_length"] = prefix_length
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["usage_state"] = usage_state
            __props__.__dict__["value"] = value
            __props__.__dict__["allocation_date"] = None
            __props__.__dict__["ip"] = None
            __props__.__dict__["ip_address"] = None
            __props__.__dict__["used_by_id"] = None
        super(IpSpaceIpAllocation, __self__).__init__(
            'vcd:index/ipSpaceIpAllocation:IpSpaceIpAllocation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allocation_date: Optional[pulumi.Input[str]] = None,
            description: Optional[pulumi.Input[str]] = None,
            ip: Optional[pulumi.Input[str]] = None,
            ip_address: Optional[pulumi.Input[str]] = None,
            ip_space_id: Optional[pulumi.Input[str]] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            prefix_length: Optional[pulumi.Input[str]] = None,
            type: Optional[pulumi.Input[str]] = None,
            usage_state: Optional[pulumi.Input[str]] = None,
            used_by_id: Optional[pulumi.Input[str]] = None,
            value: Optional[pulumi.Input[str]] = None) -> 'IpSpaceIpAllocation':
        """
        Get an existing IpSpaceIpAllocation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] allocation_date: allocation date in formated as `2023-06-07T09:57:58.721Z` (ISO 8601)
        :param pulumi.Input[str] description: Can only be set when `usage_state=USED_MANUAL`
               
               > IP Allocation resources can be created only if there is a NSX-T Edge Gateway
               (`NsxtEdgegateway`) that is backed by the Provider Gateway (`ExternalNetworkV2`) with IP
               Space Uplinks (`IpSpaceUplink`). Attempting to allocate IP Addresses before having an
               Edge Gateway withing VDC will return errors of type `This operation is denied`.
        :param pulumi.Input[str] ip: convenience field. For `type=IP_PREFIX` it will contain only the IP from CIDR returned
        :param pulumi.Input[str] ip_address: IP address or CIDR
        :param pulumi.Input[str] ip_space_id: IP Space ID to use for IP Allocations
        :param pulumi.Input[str] org_id: Org ID in which the IP is allocated
        :param pulumi.Input[str] prefix_length: Required when `type=IP_PREFIX`
        :param pulumi.Input[str] type: One of `FLOATING_IP`, `IP_PREFIX`
               * `FLOATING_IP` - allocates single IP from defined ranges in IP Space
               * `IP_PREFIX` - allocates subnets. **Note** field `prefix_length` is required to allocate IP
               Prefix
        :param pulumi.Input[str] usage_state: (Optional) Only used with manual reservations. Value `USED_MANUAL`
               enables manual IP reservation. Value `UNUSED` is set to release manual allocation of IP.
        :param pulumi.Input[str] used_by_id: contains entity ID that is using the IP if `usage_state=USED`
        :param pulumi.Input[str] value: An option to request a specific IP or subnet from IP Space.
               **Note:** This field does not support IP ranges because it would cause multiple allocations
               created in one resource. Please use multiple resource instances to allocate IP ranges.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _IpSpaceIpAllocationState.__new__(_IpSpaceIpAllocationState)

        __props__.__dict__["allocation_date"] = allocation_date
        __props__.__dict__["description"] = description
        __props__.__dict__["ip"] = ip
        __props__.__dict__["ip_address"] = ip_address
        __props__.__dict__["ip_space_id"] = ip_space_id
        __props__.__dict__["org_id"] = org_id
        __props__.__dict__["prefix_length"] = prefix_length
        __props__.__dict__["type"] = type
        __props__.__dict__["usage_state"] = usage_state
        __props__.__dict__["used_by_id"] = used_by_id
        __props__.__dict__["value"] = value
        return IpSpaceIpAllocation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allocationDate")
    def allocation_date(self) -> pulumi.Output[str]:
        """
        allocation date in formated as `2023-06-07T09:57:58.721Z` (ISO 8601)
        """
        return pulumi.get(self, "allocation_date")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Can only be set when `usage_state=USED_MANUAL`

        > IP Allocation resources can be created only if there is a NSX-T Edge Gateway
        (`NsxtEdgegateway`) that is backed by the Provider Gateway (`ExternalNetworkV2`) with IP
        Space Uplinks (`IpSpaceUplink`). Attempting to allocate IP Addresses before having an
        Edge Gateway withing VDC will return errors of type `This operation is denied`.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def ip(self) -> pulumi.Output[str]:
        """
        convenience field. For `type=IP_PREFIX` it will contain only the IP from CIDR returned
        """
        return pulumi.get(self, "ip")

    @property
    @pulumi.getter(name="ipAddress")
    def ip_address(self) -> pulumi.Output[str]:
        """
        IP address or CIDR
        """
        return pulumi.get(self, "ip_address")

    @property
    @pulumi.getter(name="ipSpaceId")
    def ip_space_id(self) -> pulumi.Output[Optional[str]]:
        """
        IP Space ID to use for IP Allocations
        """
        return pulumi.get(self, "ip_space_id")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[str]:
        """
        Org ID in which the IP is allocated
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter(name="prefixLength")
    def prefix_length(self) -> pulumi.Output[str]:
        """
        Required when `type=IP_PREFIX`
        """
        return pulumi.get(self, "prefix_length")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        One of `FLOATING_IP`, `IP_PREFIX`
        * `FLOATING_IP` - allocates single IP from defined ranges in IP Space
        * `IP_PREFIX` - allocates subnets. **Note** field `prefix_length` is required to allocate IP
        Prefix
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="usageState")
    def usage_state(self) -> pulumi.Output[str]:
        """
        (Optional) Only used with manual reservations. Value `USED_MANUAL`
        enables manual IP reservation. Value `UNUSED` is set to release manual allocation of IP.
        """
        return pulumi.get(self, "usage_state")

    @property
    @pulumi.getter(name="usedById")
    def used_by_id(self) -> pulumi.Output[str]:
        """
        contains entity ID that is using the IP if `usage_state=USED`
        """
        return pulumi.get(self, "used_by_id")

    @property
    @pulumi.getter
    def value(self) -> pulumi.Output[Optional[str]]:
        """
        An option to request a specific IP or subnet from IP Space.
        **Note:** This field does not support IP ranges because it would cause multiple allocations
        created in one resource. Please use multiple resource instances to allocate IP ranges.
        """
        return pulumi.get(self, "value")

