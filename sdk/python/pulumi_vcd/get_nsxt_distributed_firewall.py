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
    'GetNsxtDistributedFirewallResult',
    'AwaitableGetNsxtDistributedFirewallResult',
    'get_nsxt_distributed_firewall',
    'get_nsxt_distributed_firewall_output',
]

@pulumi.output_type
class GetNsxtDistributedFirewallResult:
    """
    A collection of values returned by getNsxtDistributedFirewall.
    """
    def __init__(__self__, id=None, org=None, rules=None, vdc_group_id=None):
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if org and not isinstance(org, str):
            raise TypeError("Expected argument 'org' to be a str")
        pulumi.set(__self__, "org", org)
        if rules and not isinstance(rules, list):
            raise TypeError("Expected argument 'rules' to be a list")
        pulumi.set(__self__, "rules", rules)
        if vdc_group_id and not isinstance(vdc_group_id, str):
            raise TypeError("Expected argument 'vdc_group_id' to be a str")
        pulumi.set(__self__, "vdc_group_id", vdc_group_id)

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
    def rules(self) -> Sequence['outputs.GetNsxtDistributedFirewallRuleResult']:
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter(name="vdcGroupId")
    def vdc_group_id(self) -> str:
        return pulumi.get(self, "vdc_group_id")


class AwaitableGetNsxtDistributedFirewallResult(GetNsxtDistributedFirewallResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNsxtDistributedFirewallResult(
            id=self.id,
            org=self.org,
            rules=self.rules,
            vdc_group_id=self.vdc_group_id)


def get_nsxt_distributed_firewall(org: Optional[str] = None,
                                  vdc_group_id: Optional[str] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNsxtDistributedFirewallResult:
    """
    The Distributed Firewall data source reads all defined rules for a particular VDC Group.

    > There is a different data source
    [`NsxtDistributedFirewallRule`](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_distributed_firewall_rule)
    resource are available that can fetch a single firewall rule by name.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vcd as vcd

    g1 = vcd.get_vdc_group(org="my-org",
        name="my-vdc-group")
    t1 = vcd.get_nsxt_distributed_firewall(org="my-org",
        vdc_group_id=g1.id)
    ```


    :param str org: The name of organization in which Distributed Firewall is located. Optional if
           defined at provider level.
    :param str vdc_group_id: The ID of a VDC Group
    """
    __args__ = dict()
    __args__['org'] = org
    __args__['vdcGroupId'] = vdc_group_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getNsxtDistributedFirewall:getNsxtDistributedFirewall', __args__, opts=opts, typ=GetNsxtDistributedFirewallResult).value

    return AwaitableGetNsxtDistributedFirewallResult(
        id=pulumi.get(__ret__, 'id'),
        org=pulumi.get(__ret__, 'org'),
        rules=pulumi.get(__ret__, 'rules'),
        vdc_group_id=pulumi.get(__ret__, 'vdc_group_id'))
def get_nsxt_distributed_firewall_output(org: Optional[pulumi.Input[Optional[str]]] = None,
                                         vdc_group_id: Optional[pulumi.Input[str]] = None,
                                         opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetNsxtDistributedFirewallResult]:
    """
    The Distributed Firewall data source reads all defined rules for a particular VDC Group.

    > There is a different data source
    [`NsxtDistributedFirewallRule`](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_distributed_firewall_rule)
    resource are available that can fetch a single firewall rule by name.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vcd as vcd

    g1 = vcd.get_vdc_group(org="my-org",
        name="my-vdc-group")
    t1 = vcd.get_nsxt_distributed_firewall(org="my-org",
        vdc_group_id=g1.id)
    ```


    :param str org: The name of organization in which Distributed Firewall is located. Optional if
           defined at provider level.
    :param str vdc_group_id: The ID of a VDC Group
    """
    __args__ = dict()
    __args__['org'] = org
    __args__['vdcGroupId'] = vdc_group_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vcd:index/getNsxtDistributedFirewall:getNsxtDistributedFirewall', __args__, opts=opts, typ=GetNsxtDistributedFirewallResult)
    return __ret__.apply(lambda __response__: GetNsxtDistributedFirewallResult(
        id=pulumi.get(__response__, 'id'),
        org=pulumi.get(__response__, 'org'),
        rules=pulumi.get(__response__, 'rules'),
        vdc_group_id=pulumi.get(__response__, 'vdc_group_id')))
