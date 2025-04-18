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
from . import outputs

__all__ = [
    'GetNsxvDhcpRelayResult',
    'AwaitableGetNsxvDhcpRelayResult',
    'get_nsxv_dhcp_relay',
    'get_nsxv_dhcp_relay_output',
]

@pulumi.output_type
class GetNsxvDhcpRelayResult:
    """
    A collection of values returned by getNsxvDhcpRelay.
    """
    def __init__(__self__, domain_names=None, edge_gateway=None, id=None, ip_addresses=None, ip_sets=None, org=None, relay_agents=None, vdc=None):
        if domain_names and not isinstance(domain_names, list):
            raise TypeError("Expected argument 'domain_names' to be a list")
        pulumi.set(__self__, "domain_names", domain_names)
        if edge_gateway and not isinstance(edge_gateway, str):
            raise TypeError("Expected argument 'edge_gateway' to be a str")
        pulumi.set(__self__, "edge_gateway", edge_gateway)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_addresses and not isinstance(ip_addresses, list):
            raise TypeError("Expected argument 'ip_addresses' to be a list")
        pulumi.set(__self__, "ip_addresses", ip_addresses)
        if ip_sets and not isinstance(ip_sets, list):
            raise TypeError("Expected argument 'ip_sets' to be a list")
        pulumi.set(__self__, "ip_sets", ip_sets)
        if org and not isinstance(org, str):
            raise TypeError("Expected argument 'org' to be a str")
        pulumi.set(__self__, "org", org)
        if relay_agents and not isinstance(relay_agents, list):
            raise TypeError("Expected argument 'relay_agents' to be a list")
        pulumi.set(__self__, "relay_agents", relay_agents)
        if vdc and not isinstance(vdc, str):
            raise TypeError("Expected argument 'vdc' to be a str")
        pulumi.set(__self__, "vdc", vdc)

    @property
    @pulumi.getter(name="domainNames")
    def domain_names(self) -> Sequence[str]:
        return pulumi.get(self, "domain_names")

    @property
    @pulumi.getter(name="edgeGateway")
    def edge_gateway(self) -> str:
        return pulumi.get(self, "edge_gateway")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ipAddresses")
    def ip_addresses(self) -> Sequence[str]:
        return pulumi.get(self, "ip_addresses")

    @property
    @pulumi.getter(name="ipSets")
    def ip_sets(self) -> Sequence[str]:
        return pulumi.get(self, "ip_sets")

    @property
    @pulumi.getter
    def org(self) -> Optional[str]:
        return pulumi.get(self, "org")

    @property
    @pulumi.getter(name="relayAgents")
    def relay_agents(self) -> Sequence['outputs.GetNsxvDhcpRelayRelayAgentResult']:
        return pulumi.get(self, "relay_agents")

    @property
    @pulumi.getter
    def vdc(self) -> Optional[str]:
        return pulumi.get(self, "vdc")


class AwaitableGetNsxvDhcpRelayResult(GetNsxvDhcpRelayResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNsxvDhcpRelayResult(
            domain_names=self.domain_names,
            edge_gateway=self.edge_gateway,
            id=self.id,
            ip_addresses=self.ip_addresses,
            ip_sets=self.ip_sets,
            org=self.org,
            relay_agents=self.relay_agents,
            vdc=self.vdc)


def get_nsxv_dhcp_relay(edge_gateway: Optional[str] = None,
                        org: Optional[str] = None,
                        vdc: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNsxvDhcpRelayResult:
    """
    Provides a VMware Cloud Director Edge Gateway DHCP relay configuration data source. The DHCP relay
    capability provided by NSX in VMware Cloud Director environment allows to leverage existing DHCP
    infrastructure from within VMware Cloud Director environment without any interruption to the IP address
    management in existing DHCP infrastructure. DHCP messages are relayed from virtual machines to the
    designated DHCP servers in your physical DHCP infrastructure, which allows IP addresses controlled
    by the NSX software to continue to be in sync with IP addresses in the rest of your DHCP-controlled
    environments.

    Supported in provider *v2.6+*

    ## Example Usage

    ### 1

    ```python
    import pulumi
    import pulumi_vcd as vcd

    relay_config = vcd.get_nsxv_dhcp_relay(org="my-org",
        vdc="my-org-vdc",
        edge_gateway="my-edge-gw")
    ```


    :param str edge_gateway: The name of the edge gateway on which DHCP relay is to be configured.
    :param str org: The name of organization to use, optional if defined at provider level. Useful
           when connected as sysadmin working across different organisations.
    :param str vdc: The name of VDC to use, optional if defined at provider level.
    """
    __args__ = dict()
    __args__['edgeGateway'] = edge_gateway
    __args__['org'] = org
    __args__['vdc'] = vdc
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getNsxvDhcpRelay:getNsxvDhcpRelay', __args__, opts=opts, typ=GetNsxvDhcpRelayResult).value

    return AwaitableGetNsxvDhcpRelayResult(
        domain_names=pulumi.get(__ret__, 'domain_names'),
        edge_gateway=pulumi.get(__ret__, 'edge_gateway'),
        id=pulumi.get(__ret__, 'id'),
        ip_addresses=pulumi.get(__ret__, 'ip_addresses'),
        ip_sets=pulumi.get(__ret__, 'ip_sets'),
        org=pulumi.get(__ret__, 'org'),
        relay_agents=pulumi.get(__ret__, 'relay_agents'),
        vdc=pulumi.get(__ret__, 'vdc'))
def get_nsxv_dhcp_relay_output(edge_gateway: Optional[pulumi.Input[str]] = None,
                               org: Optional[pulumi.Input[Optional[str]]] = None,
                               vdc: Optional[pulumi.Input[Optional[str]]] = None,
                               opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetNsxvDhcpRelayResult]:
    """
    Provides a VMware Cloud Director Edge Gateway DHCP relay configuration data source. The DHCP relay
    capability provided by NSX in VMware Cloud Director environment allows to leverage existing DHCP
    infrastructure from within VMware Cloud Director environment without any interruption to the IP address
    management in existing DHCP infrastructure. DHCP messages are relayed from virtual machines to the
    designated DHCP servers in your physical DHCP infrastructure, which allows IP addresses controlled
    by the NSX software to continue to be in sync with IP addresses in the rest of your DHCP-controlled
    environments.

    Supported in provider *v2.6+*

    ## Example Usage

    ### 1

    ```python
    import pulumi
    import pulumi_vcd as vcd

    relay_config = vcd.get_nsxv_dhcp_relay(org="my-org",
        vdc="my-org-vdc",
        edge_gateway="my-edge-gw")
    ```


    :param str edge_gateway: The name of the edge gateway on which DHCP relay is to be configured.
    :param str org: The name of organization to use, optional if defined at provider level. Useful
           when connected as sysadmin working across different organisations.
    :param str vdc: The name of VDC to use, optional if defined at provider level.
    """
    __args__ = dict()
    __args__['edgeGateway'] = edge_gateway
    __args__['org'] = org
    __args__['vdc'] = vdc
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vcd:index/getNsxvDhcpRelay:getNsxvDhcpRelay', __args__, opts=opts, typ=GetNsxvDhcpRelayResult)
    return __ret__.apply(lambda __response__: GetNsxvDhcpRelayResult(
        domain_names=pulumi.get(__response__, 'domain_names'),
        edge_gateway=pulumi.get(__response__, 'edge_gateway'),
        id=pulumi.get(__response__, 'id'),
        ip_addresses=pulumi.get(__response__, 'ip_addresses'),
        ip_sets=pulumi.get(__response__, 'ip_sets'),
        org=pulumi.get(__response__, 'org'),
        relay_agents=pulumi.get(__response__, 'relay_agents'),
        vdc=pulumi.get(__response__, 'vdc')))
