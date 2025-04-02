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
    'GetNsxtIpSetResult',
    'AwaitableGetNsxtIpSetResult',
    'get_nsxt_ip_set',
    'get_nsxt_ip_set_output',
]

@pulumi.output_type
class GetNsxtIpSetResult:
    """
    A collection of values returned by getNsxtIpSet.
    """
    def __init__(__self__, description=None, edge_gateway_id=None, id=None, ip_addresses=None, name=None, org=None, owner_id=None, vdc=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if edge_gateway_id and not isinstance(edge_gateway_id, str):
            raise TypeError("Expected argument 'edge_gateway_id' to be a str")
        pulumi.set(__self__, "edge_gateway_id", edge_gateway_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_addresses and not isinstance(ip_addresses, list):
            raise TypeError("Expected argument 'ip_addresses' to be a list")
        pulumi.set(__self__, "ip_addresses", ip_addresses)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if org and not isinstance(org, str):
            raise TypeError("Expected argument 'org' to be a str")
        pulumi.set(__self__, "org", org)
        if owner_id and not isinstance(owner_id, str):
            raise TypeError("Expected argument 'owner_id' to be a str")
        pulumi.set(__self__, "owner_id", owner_id)
        if vdc and not isinstance(vdc, str):
            raise TypeError("Expected argument 'vdc' to be a str")
        pulumi.set(__self__, "vdc", vdc)

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

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
    @pulumi.getter(name="ipAddresses")
    def ip_addresses(self) -> Sequence[str]:
        return pulumi.get(self, "ip_addresses")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def org(self) -> Optional[str]:
        return pulumi.get(self, "org")

    @property
    @pulumi.getter(name="ownerId")
    def owner_id(self) -> str:
        """
        Parent VDC or VDC Group ID.
        """
        return pulumi.get(self, "owner_id")

    @property
    @pulumi.getter
    @_utilities.deprecated("""Deprecated in favor of `edge_gateway_id`. IP Set will inherit VDC from parent Edge Gateway.""")
    def vdc(self) -> Optional[str]:
        return pulumi.get(self, "vdc")


class AwaitableGetNsxtIpSetResult(GetNsxtIpSetResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNsxtIpSetResult(
            description=self.description,
            edge_gateway_id=self.edge_gateway_id,
            id=self.id,
            ip_addresses=self.ip_addresses,
            name=self.name,
            org=self.org,
            owner_id=self.owner_id,
            vdc=self.vdc)


def get_nsxt_ip_set(edge_gateway_id: Optional[str] = None,
                    name: Optional[str] = None,
                    org: Optional[str] = None,
                    vdc: Optional[str] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNsxtIpSetResult:
    """
    Supported in provider *v3.3+* and VCD 10.1+ with NSX-T backed VDCs.

    Provides a data source to read NSX-T IP Set. IP Sets are groups of objects to which the firewall rules apply. Combining
    multiple objects into IP Sets helps reduce the total number of firewall rules to be created.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vcd as vcd

    main = vcd.get_nsxt_edgegateway(org="my-org",
        name="main-edge")
    my_set_1 = vcd.get_nsxt_ip_set(org="my-org",
        edge_gateway_id=main.id,
        name="frontend-servers")
    ```


    :param str edge_gateway_id: The ID of the Edge Gateway (NSX-T only). Can be looked up using
    :param str name: Unique name of existing IP Set.
    :param str org: The name of organization to use, optional if defined at provider level. Useful
           when connected as sysadmin working across different organisations.
    :param str vdc: The name of VDC to use, optional if defined at provider level. **Deprecated**
           in favor of `edge_gateway_id` field.
    """
    __args__ = dict()
    __args__['edgeGatewayId'] = edge_gateway_id
    __args__['name'] = name
    __args__['org'] = org
    __args__['vdc'] = vdc
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getNsxtIpSet:getNsxtIpSet', __args__, opts=opts, typ=GetNsxtIpSetResult).value

    return AwaitableGetNsxtIpSetResult(
        description=pulumi.get(__ret__, 'description'),
        edge_gateway_id=pulumi.get(__ret__, 'edge_gateway_id'),
        id=pulumi.get(__ret__, 'id'),
        ip_addresses=pulumi.get(__ret__, 'ip_addresses'),
        name=pulumi.get(__ret__, 'name'),
        org=pulumi.get(__ret__, 'org'),
        owner_id=pulumi.get(__ret__, 'owner_id'),
        vdc=pulumi.get(__ret__, 'vdc'))
def get_nsxt_ip_set_output(edge_gateway_id: Optional[pulumi.Input[str]] = None,
                           name: Optional[pulumi.Input[str]] = None,
                           org: Optional[pulumi.Input[Optional[str]]] = None,
                           vdc: Optional[pulumi.Input[Optional[str]]] = None,
                           opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetNsxtIpSetResult]:
    """
    Supported in provider *v3.3+* and VCD 10.1+ with NSX-T backed VDCs.

    Provides a data source to read NSX-T IP Set. IP Sets are groups of objects to which the firewall rules apply. Combining
    multiple objects into IP Sets helps reduce the total number of firewall rules to be created.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vcd as vcd

    main = vcd.get_nsxt_edgegateway(org="my-org",
        name="main-edge")
    my_set_1 = vcd.get_nsxt_ip_set(org="my-org",
        edge_gateway_id=main.id,
        name="frontend-servers")
    ```


    :param str edge_gateway_id: The ID of the Edge Gateway (NSX-T only). Can be looked up using
    :param str name: Unique name of existing IP Set.
    :param str org: The name of organization to use, optional if defined at provider level. Useful
           when connected as sysadmin working across different organisations.
    :param str vdc: The name of VDC to use, optional if defined at provider level. **Deprecated**
           in favor of `edge_gateway_id` field.
    """
    __args__ = dict()
    __args__['edgeGatewayId'] = edge_gateway_id
    __args__['name'] = name
    __args__['org'] = org
    __args__['vdc'] = vdc
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vcd:index/getNsxtIpSet:getNsxtIpSet', __args__, opts=opts, typ=GetNsxtIpSetResult)
    return __ret__.apply(lambda __response__: GetNsxtIpSetResult(
        description=pulumi.get(__response__, 'description'),
        edge_gateway_id=pulumi.get(__response__, 'edge_gateway_id'),
        id=pulumi.get(__response__, 'id'),
        ip_addresses=pulumi.get(__response__, 'ip_addresses'),
        name=pulumi.get(__response__, 'name'),
        org=pulumi.get(__response__, 'org'),
        owner_id=pulumi.get(__response__, 'owner_id'),
        vdc=pulumi.get(__response__, 'vdc')))
