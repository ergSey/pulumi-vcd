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

__all__ = [
    'GetNsxtSegmentIpDiscoveryProfileResult',
    'AwaitableGetNsxtSegmentIpDiscoveryProfileResult',
    'get_nsxt_segment_ip_discovery_profile',
    'get_nsxt_segment_ip_discovery_profile_output',
]

@pulumi.output_type
class GetNsxtSegmentIpDiscoveryProfileResult:
    """
    A collection of values returned by getNsxtSegmentIpDiscoveryProfile.
    """
    def __init__(__self__, arp_binding_limit=None, arp_binding_timeout=None, description=None, id=None, is_arp_snooping_enabled=None, is_dhcp_snooping_v4_enabled=None, is_dhcp_snooping_v6_enabled=None, is_duplicate_ip_detection_enabled=None, is_nd_snooping_enabled=None, is_tofu_enabled=None, is_vmtools_v4_enabled=None, is_vmtools_v6_enabled=None, name=None, nd_snooping_limit=None, nsxt_manager_id=None, vdc_group_id=None, vdc_id=None):
        if arp_binding_limit and not isinstance(arp_binding_limit, int):
            raise TypeError("Expected argument 'arp_binding_limit' to be a int")
        pulumi.set(__self__, "arp_binding_limit", arp_binding_limit)
        if arp_binding_timeout and not isinstance(arp_binding_timeout, int):
            raise TypeError("Expected argument 'arp_binding_timeout' to be a int")
        pulumi.set(__self__, "arp_binding_timeout", arp_binding_timeout)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if is_arp_snooping_enabled and not isinstance(is_arp_snooping_enabled, bool):
            raise TypeError("Expected argument 'is_arp_snooping_enabled' to be a bool")
        pulumi.set(__self__, "is_arp_snooping_enabled", is_arp_snooping_enabled)
        if is_dhcp_snooping_v4_enabled and not isinstance(is_dhcp_snooping_v4_enabled, bool):
            raise TypeError("Expected argument 'is_dhcp_snooping_v4_enabled' to be a bool")
        pulumi.set(__self__, "is_dhcp_snooping_v4_enabled", is_dhcp_snooping_v4_enabled)
        if is_dhcp_snooping_v6_enabled and not isinstance(is_dhcp_snooping_v6_enabled, bool):
            raise TypeError("Expected argument 'is_dhcp_snooping_v6_enabled' to be a bool")
        pulumi.set(__self__, "is_dhcp_snooping_v6_enabled", is_dhcp_snooping_v6_enabled)
        if is_duplicate_ip_detection_enabled and not isinstance(is_duplicate_ip_detection_enabled, bool):
            raise TypeError("Expected argument 'is_duplicate_ip_detection_enabled' to be a bool")
        pulumi.set(__self__, "is_duplicate_ip_detection_enabled", is_duplicate_ip_detection_enabled)
        if is_nd_snooping_enabled and not isinstance(is_nd_snooping_enabled, bool):
            raise TypeError("Expected argument 'is_nd_snooping_enabled' to be a bool")
        pulumi.set(__self__, "is_nd_snooping_enabled", is_nd_snooping_enabled)
        if is_tofu_enabled and not isinstance(is_tofu_enabled, bool):
            raise TypeError("Expected argument 'is_tofu_enabled' to be a bool")
        pulumi.set(__self__, "is_tofu_enabled", is_tofu_enabled)
        if is_vmtools_v4_enabled and not isinstance(is_vmtools_v4_enabled, bool):
            raise TypeError("Expected argument 'is_vmtools_v4_enabled' to be a bool")
        pulumi.set(__self__, "is_vmtools_v4_enabled", is_vmtools_v4_enabled)
        if is_vmtools_v6_enabled and not isinstance(is_vmtools_v6_enabled, bool):
            raise TypeError("Expected argument 'is_vmtools_v6_enabled' to be a bool")
        pulumi.set(__self__, "is_vmtools_v6_enabled", is_vmtools_v6_enabled)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if nd_snooping_limit and not isinstance(nd_snooping_limit, int):
            raise TypeError("Expected argument 'nd_snooping_limit' to be a int")
        pulumi.set(__self__, "nd_snooping_limit", nd_snooping_limit)
        if nsxt_manager_id and not isinstance(nsxt_manager_id, str):
            raise TypeError("Expected argument 'nsxt_manager_id' to be a str")
        pulumi.set(__self__, "nsxt_manager_id", nsxt_manager_id)
        if vdc_group_id and not isinstance(vdc_group_id, str):
            raise TypeError("Expected argument 'vdc_group_id' to be a str")
        pulumi.set(__self__, "vdc_group_id", vdc_group_id)
        if vdc_id and not isinstance(vdc_id, str):
            raise TypeError("Expected argument 'vdc_id' to be a str")
        pulumi.set(__self__, "vdc_id", vdc_id)

    @property
    @pulumi.getter(name="arpBindingLimit")
    def arp_binding_limit(self) -> int:
        """
        Indicates the number of ARP snooped IP addresses to be remembered per
        logical port
        """
        return pulumi.get(self, "arp_binding_limit")

    @property
    @pulumi.getter(name="arpBindingTimeout")
    def arp_binding_timeout(self) -> int:
        """
        ARP and ND (Neighbor Discovery) cache timeout (in minutes)
        """
        return pulumi.get(self, "arp_binding_timeout")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Description of IP Discovery Profile
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="isArpSnoopingEnabled")
    def is_arp_snooping_enabled(self) -> bool:
        """
        Defines whether ARP snooping is enabled
        """
        return pulumi.get(self, "is_arp_snooping_enabled")

    @property
    @pulumi.getter(name="isDhcpSnoopingV4Enabled")
    def is_dhcp_snooping_v4_enabled(self) -> bool:
        """
        Defines whether DHCP snooping for IPv4 is enabled
        """
        return pulumi.get(self, "is_dhcp_snooping_v4_enabled")

    @property
    @pulumi.getter(name="isDhcpSnoopingV6Enabled")
    def is_dhcp_snooping_v6_enabled(self) -> bool:
        """
        Defines whether DHCP snooping for IPv6 is enabled
        """
        return pulumi.get(self, "is_dhcp_snooping_v6_enabled")

    @property
    @pulumi.getter(name="isDuplicateIpDetectionEnabled")
    def is_duplicate_ip_detection_enabled(self) -> bool:
        """
        Defines whether duplicate IP detection is enabled. Duplicate
        IP detection is used to determine if there is any IP conflict with any other port on the same
        logical switch. If a conflict is detected, then the IP is marked as a duplicate on the port where
        the IP was discovered last
        """
        return pulumi.get(self, "is_duplicate_ip_detection_enabled")

    @property
    @pulumi.getter(name="isNdSnoopingEnabled")
    def is_nd_snooping_enabled(self) -> bool:
        """
        Defines whether ND (Neighbor Discovery) snooping is enabled. If true,
        this method will snoop the NS (Neighbor Solicitation) and NA (Neighbor Advertisement) messages in
        the ND (Neighbor Discovery Protocol) family of messages which are transmitted by a VM. From the NS
        messages, we will learn about the source which sent this NS message. From the NA message, we will
        learn the resolved address in the message which the VM is a recipient of. Addresses snooped by
        this method are subject to TOFU
        """
        return pulumi.get(self, "is_nd_snooping_enabled")

    @property
    @pulumi.getter(name="isTofuEnabled")
    def is_tofu_enabled(self) -> bool:
        """
        Defines whether `Trust on First Use(TOFU)` paradigm is enabled
        """
        return pulumi.get(self, "is_tofu_enabled")

    @property
    @pulumi.getter(name="isVmtoolsV4Enabled")
    def is_vmtools_v4_enabled(self) -> bool:
        """
        Defines whether fetching IPv4 address using vm-tools is enabled. This
        option is only supported on ESX where vm-tools is installed
        """
        return pulumi.get(self, "is_vmtools_v4_enabled")

    @property
    @pulumi.getter(name="isVmtoolsV6Enabled")
    def is_vmtools_v6_enabled(self) -> bool:
        """
        Defines whether fetching IPv6 address using vm-tools is enabled. This
        will learn the IPv6 addresses which are configured on interfaces of a VM with the help of the
        VMTools software
        """
        return pulumi.get(self, "is_vmtools_v6_enabled")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="ndSnoopingLimit")
    def nd_snooping_limit(self) -> int:
        """
        Maximum number of ND (Neighbor Discovery Protocol) snooped IPv6 addresses
        """
        return pulumi.get(self, "nd_snooping_limit")

    @property
    @pulumi.getter(name="nsxtManagerId")
    def nsxt_manager_id(self) -> Optional[str]:
        return pulumi.get(self, "nsxt_manager_id")

    @property
    @pulumi.getter(name="vdcGroupId")
    def vdc_group_id(self) -> Optional[str]:
        return pulumi.get(self, "vdc_group_id")

    @property
    @pulumi.getter(name="vdcId")
    def vdc_id(self) -> Optional[str]:
        return pulumi.get(self, "vdc_id")


class AwaitableGetNsxtSegmentIpDiscoveryProfileResult(GetNsxtSegmentIpDiscoveryProfileResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNsxtSegmentIpDiscoveryProfileResult(
            arp_binding_limit=self.arp_binding_limit,
            arp_binding_timeout=self.arp_binding_timeout,
            description=self.description,
            id=self.id,
            is_arp_snooping_enabled=self.is_arp_snooping_enabled,
            is_dhcp_snooping_v4_enabled=self.is_dhcp_snooping_v4_enabled,
            is_dhcp_snooping_v6_enabled=self.is_dhcp_snooping_v6_enabled,
            is_duplicate_ip_detection_enabled=self.is_duplicate_ip_detection_enabled,
            is_nd_snooping_enabled=self.is_nd_snooping_enabled,
            is_tofu_enabled=self.is_tofu_enabled,
            is_vmtools_v4_enabled=self.is_vmtools_v4_enabled,
            is_vmtools_v6_enabled=self.is_vmtools_v6_enabled,
            name=self.name,
            nd_snooping_limit=self.nd_snooping_limit,
            nsxt_manager_id=self.nsxt_manager_id,
            vdc_group_id=self.vdc_group_id,
            vdc_id=self.vdc_id)


def get_nsxt_segment_ip_discovery_profile(name: Optional[str] = None,
                                          nsxt_manager_id: Optional[str] = None,
                                          vdc_group_id: Optional[str] = None,
                                          vdc_id: Optional[str] = None,
                                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNsxtSegmentIpDiscoveryProfileResult:
    """
    Provides a VMware Cloud Director NSX-T IP Discovery Profile data source. This can be used to read NSX-T Segment Profile definitions.

    Supported in provider *v3.11+*.

    ## Example Usage

    ### IP Discovery Profile)

    ```python
    import pulumi
    import pulumi_vcd as vcd

    nsxt = vcd.get_nsxt_manager(name="nsxManager1")
    first = vcd.get_nsxt_segment_ip_discovery_profile(name="ip-discovery-profile-0",
        nsxt_manager_id=nsxt.id)
    ```


    :param str name: The name of Segment Profile
    :param str nsxt_manager_id: Segment Profile search context. Use when searching by NSX-T manager
    :param str vdc_group_id: Segment Profile search context. Use when searching by VDC group
           
           > Note: only one of `nsxt_manager_id`, `vdc_id`, `vdc_group_id` can be used
    :param str vdc_id: Segment Profile search context. Use when searching by VDC
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['nsxtManagerId'] = nsxt_manager_id
    __args__['vdcGroupId'] = vdc_group_id
    __args__['vdcId'] = vdc_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getNsxtSegmentIpDiscoveryProfile:getNsxtSegmentIpDiscoveryProfile', __args__, opts=opts, typ=GetNsxtSegmentIpDiscoveryProfileResult).value

    return AwaitableGetNsxtSegmentIpDiscoveryProfileResult(
        arp_binding_limit=pulumi.get(__ret__, 'arp_binding_limit'),
        arp_binding_timeout=pulumi.get(__ret__, 'arp_binding_timeout'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        is_arp_snooping_enabled=pulumi.get(__ret__, 'is_arp_snooping_enabled'),
        is_dhcp_snooping_v4_enabled=pulumi.get(__ret__, 'is_dhcp_snooping_v4_enabled'),
        is_dhcp_snooping_v6_enabled=pulumi.get(__ret__, 'is_dhcp_snooping_v6_enabled'),
        is_duplicate_ip_detection_enabled=pulumi.get(__ret__, 'is_duplicate_ip_detection_enabled'),
        is_nd_snooping_enabled=pulumi.get(__ret__, 'is_nd_snooping_enabled'),
        is_tofu_enabled=pulumi.get(__ret__, 'is_tofu_enabled'),
        is_vmtools_v4_enabled=pulumi.get(__ret__, 'is_vmtools_v4_enabled'),
        is_vmtools_v6_enabled=pulumi.get(__ret__, 'is_vmtools_v6_enabled'),
        name=pulumi.get(__ret__, 'name'),
        nd_snooping_limit=pulumi.get(__ret__, 'nd_snooping_limit'),
        nsxt_manager_id=pulumi.get(__ret__, 'nsxt_manager_id'),
        vdc_group_id=pulumi.get(__ret__, 'vdc_group_id'),
        vdc_id=pulumi.get(__ret__, 'vdc_id'))
def get_nsxt_segment_ip_discovery_profile_output(name: Optional[pulumi.Input[str]] = None,
                                                 nsxt_manager_id: Optional[pulumi.Input[Optional[str]]] = None,
                                                 vdc_group_id: Optional[pulumi.Input[Optional[str]]] = None,
                                                 vdc_id: Optional[pulumi.Input[Optional[str]]] = None,
                                                 opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetNsxtSegmentIpDiscoveryProfileResult]:
    """
    Provides a VMware Cloud Director NSX-T IP Discovery Profile data source. This can be used to read NSX-T Segment Profile definitions.

    Supported in provider *v3.11+*.

    ## Example Usage

    ### IP Discovery Profile)

    ```python
    import pulumi
    import pulumi_vcd as vcd

    nsxt = vcd.get_nsxt_manager(name="nsxManager1")
    first = vcd.get_nsxt_segment_ip_discovery_profile(name="ip-discovery-profile-0",
        nsxt_manager_id=nsxt.id)
    ```


    :param str name: The name of Segment Profile
    :param str nsxt_manager_id: Segment Profile search context. Use when searching by NSX-T manager
    :param str vdc_group_id: Segment Profile search context. Use when searching by VDC group
           
           > Note: only one of `nsxt_manager_id`, `vdc_id`, `vdc_group_id` can be used
    :param str vdc_id: Segment Profile search context. Use when searching by VDC
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['nsxtManagerId'] = nsxt_manager_id
    __args__['vdcGroupId'] = vdc_group_id
    __args__['vdcId'] = vdc_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vcd:index/getNsxtSegmentIpDiscoveryProfile:getNsxtSegmentIpDiscoveryProfile', __args__, opts=opts, typ=GetNsxtSegmentIpDiscoveryProfileResult)
    return __ret__.apply(lambda __response__: GetNsxtSegmentIpDiscoveryProfileResult(
        arp_binding_limit=pulumi.get(__response__, 'arp_binding_limit'),
        arp_binding_timeout=pulumi.get(__response__, 'arp_binding_timeout'),
        description=pulumi.get(__response__, 'description'),
        id=pulumi.get(__response__, 'id'),
        is_arp_snooping_enabled=pulumi.get(__response__, 'is_arp_snooping_enabled'),
        is_dhcp_snooping_v4_enabled=pulumi.get(__response__, 'is_dhcp_snooping_v4_enabled'),
        is_dhcp_snooping_v6_enabled=pulumi.get(__response__, 'is_dhcp_snooping_v6_enabled'),
        is_duplicate_ip_detection_enabled=pulumi.get(__response__, 'is_duplicate_ip_detection_enabled'),
        is_nd_snooping_enabled=pulumi.get(__response__, 'is_nd_snooping_enabled'),
        is_tofu_enabled=pulumi.get(__response__, 'is_tofu_enabled'),
        is_vmtools_v4_enabled=pulumi.get(__response__, 'is_vmtools_v4_enabled'),
        is_vmtools_v6_enabled=pulumi.get(__response__, 'is_vmtools_v6_enabled'),
        name=pulumi.get(__response__, 'name'),
        nd_snooping_limit=pulumi.get(__response__, 'nd_snooping_limit'),
        nsxt_manager_id=pulumi.get(__response__, 'nsxt_manager_id'),
        vdc_group_id=pulumi.get(__response__, 'vdc_group_id'),
        vdc_id=pulumi.get(__response__, 'vdc_id')))
