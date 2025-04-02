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
    'GetNsxtFirewallResult',
    'AwaitableGetNsxtFirewallResult',
    'get_nsxt_firewall',
    'get_nsxt_firewall_output',
]

@pulumi.output_type
class GetNsxtFirewallResult:
    """
    A collection of values returned by getNsxtFirewall.
    """
    def __init__(__self__, edge_gateway_id=None, id=None, org=None, rules=None, vdc=None):
        if edge_gateway_id and not isinstance(edge_gateway_id, str):
            raise TypeError("Expected argument 'edge_gateway_id' to be a str")
        pulumi.set(__self__, "edge_gateway_id", edge_gateway_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if org and not isinstance(org, str):
            raise TypeError("Expected argument 'org' to be a str")
        pulumi.set(__self__, "org", org)
        if rules and not isinstance(rules, list):
            raise TypeError("Expected argument 'rules' to be a list")
        pulumi.set(__self__, "rules", rules)
        if vdc and not isinstance(vdc, str):
            raise TypeError("Expected argument 'vdc' to be a str")
        pulumi.set(__self__, "vdc", vdc)

    @property
    @pulumi.getter(name="edgeGatewayId")
    def edge_gateway_id(self) -> str:
        return pulumi.get(self, "edge_gateway_id")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def org(self) -> Optional[str]:
        return pulumi.get(self, "org")

    @property
    @pulumi.getter
    def rules(self) -> Sequence['outputs.GetNsxtFirewallRuleResult']:
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter
    @_utilities.deprecated("""Edge Gateway will be looked up based on 'edge_gateway_id' field""")
    def vdc(self) -> Optional[str]:
        return pulumi.get(self, "vdc")


class AwaitableGetNsxtFirewallResult(GetNsxtFirewallResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNsxtFirewallResult(
            edge_gateway_id=self.edge_gateway_id,
            id=self.id,
            org=self.org,
            rules=self.rules,
            vdc=self.vdc)


def get_nsxt_firewall(edge_gateway_id: Optional[str] = None,
                      org: Optional[str] = None,
                      vdc: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNsxtFirewallResult:
    """
    Supported in provider *v3.3+* and VCD 10.1+ with NSX-T backed Edge Gateways.

    Provides a data source to read NSX-T Firewall configuration of an Edge Gateway. Firewalls allow
    user to control the incoming and outgoing network traffic to and from an NSX-T Data Center
    Edge Gateway.

    ## Example Usage

    ### Read A List Of Firewall Rules On Edge Gateway)
    ```python
    import pulumi
    import pulumi_vcd as vcd

    testing = vcd.get_nsxt_firewall(org="my-org",
        edge_gateway_id=testing_vcd_nsxt_edgegateway["id"])
    ```


    :param str edge_gateway_id: The ID of the Edge Gateway (NSX-T only). Can be looked up using
           `NsxtEdgegateway` data source
    :param str org: The name of organization to use, optional if defined at provider level. Useful
           when connected as sysadmin working across different organizations.
    """
    __args__ = dict()
    __args__['edgeGatewayId'] = edge_gateway_id
    __args__['org'] = org
    __args__['vdc'] = vdc
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getNsxtFirewall:getNsxtFirewall', __args__, opts=opts, typ=GetNsxtFirewallResult).value

    return AwaitableGetNsxtFirewallResult(
        edge_gateway_id=pulumi.get(__ret__, 'edge_gateway_id'),
        id=pulumi.get(__ret__, 'id'),
        org=pulumi.get(__ret__, 'org'),
        rules=pulumi.get(__ret__, 'rules'),
        vdc=pulumi.get(__ret__, 'vdc'))
def get_nsxt_firewall_output(edge_gateway_id: Optional[pulumi.Input[str]] = None,
                             org: Optional[pulumi.Input[Optional[str]]] = None,
                             vdc: Optional[pulumi.Input[Optional[str]]] = None,
                             opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetNsxtFirewallResult]:
    """
    Supported in provider *v3.3+* and VCD 10.1+ with NSX-T backed Edge Gateways.

    Provides a data source to read NSX-T Firewall configuration of an Edge Gateway. Firewalls allow
    user to control the incoming and outgoing network traffic to and from an NSX-T Data Center
    Edge Gateway.

    ## Example Usage

    ### Read A List Of Firewall Rules On Edge Gateway)
    ```python
    import pulumi
    import pulumi_vcd as vcd

    testing = vcd.get_nsxt_firewall(org="my-org",
        edge_gateway_id=testing_vcd_nsxt_edgegateway["id"])
    ```


    :param str edge_gateway_id: The ID of the Edge Gateway (NSX-T only). Can be looked up using
           `NsxtEdgegateway` data source
    :param str org: The name of organization to use, optional if defined at provider level. Useful
           when connected as sysadmin working across different organizations.
    """
    __args__ = dict()
    __args__['edgeGatewayId'] = edge_gateway_id
    __args__['org'] = org
    __args__['vdc'] = vdc
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vcd:index/getNsxtFirewall:getNsxtFirewall', __args__, opts=opts, typ=GetNsxtFirewallResult)
    return __ret__.apply(lambda __response__: GetNsxtFirewallResult(
        edge_gateway_id=pulumi.get(__response__, 'edge_gateway_id'),
        id=pulumi.get(__response__, 'id'),
        org=pulumi.get(__response__, 'org'),
        rules=pulumi.get(__response__, 'rules'),
        vdc=pulumi.get(__response__, 'vdc')))
