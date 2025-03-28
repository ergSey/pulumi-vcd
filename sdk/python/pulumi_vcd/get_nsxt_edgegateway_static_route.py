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
    'GetNsxtEdgegatewayStaticRouteResult',
    'AwaitableGetNsxtEdgegatewayStaticRouteResult',
    'get_nsxt_edgegateway_static_route',
    'get_nsxt_edgegateway_static_route_output',
]

@pulumi.output_type
class GetNsxtEdgegatewayStaticRouteResult:
    """
    A collection of values returned by getNsxtEdgegatewayStaticRoute.
    """
    def __init__(__self__, description=None, edge_gateway_id=None, id=None, name=None, network_cidr=None, next_hops=None, org=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if edge_gateway_id and not isinstance(edge_gateway_id, str):
            raise TypeError("Expected argument 'edge_gateway_id' to be a str")
        pulumi.set(__self__, "edge_gateway_id", edge_gateway_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network_cidr and not isinstance(network_cidr, str):
            raise TypeError("Expected argument 'network_cidr' to be a str")
        pulumi.set(__self__, "network_cidr", network_cidr)
        if next_hops and not isinstance(next_hops, list):
            raise TypeError("Expected argument 'next_hops' to be a list")
        pulumi.set(__self__, "next_hops", next_hops)
        if org and not isinstance(org, str):
            raise TypeError("Expected argument 'org' to be a str")
        pulumi.set(__self__, "org", org)

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
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkCidr")
    def network_cidr(self) -> str:
        return pulumi.get(self, "network_cidr")

    @property
    @pulumi.getter(name="nextHops")
    def next_hops(self) -> Sequence['outputs.GetNsxtEdgegatewayStaticRouteNextHopResult']:
        return pulumi.get(self, "next_hops")

    @property
    @pulumi.getter
    def org(self) -> Optional[str]:
        return pulumi.get(self, "org")


class AwaitableGetNsxtEdgegatewayStaticRouteResult(GetNsxtEdgegatewayStaticRouteResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNsxtEdgegatewayStaticRouteResult(
            description=self.description,
            edge_gateway_id=self.edge_gateway_id,
            id=self.id,
            name=self.name,
            network_cidr=self.network_cidr,
            next_hops=self.next_hops,
            org=self.org)


def get_nsxt_edgegateway_static_route(edge_gateway_id: Optional[str] = None,
                                      name: Optional[str] = None,
                                      network_cidr: Optional[str] = None,
                                      org: Optional[str] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNsxtEdgegatewayStaticRouteResult:
    """
    Supported in provider *v3.10+* and VCD 10.4.0+ with NSX-T.

    Provides a data source to read NSX-T Edge Gateway Static Routes.

    ## Example Usage

    ### By Name Only)

    ```python
    import pulumi
    import pulumi_vcd as vcd

    by_name = vcd.get_nsxt_edgegateway_static_route(edge_gateway_id=existing["id"],
        name="existing-static-route")
    ```

    ### By Name And Network CIDR )

    ```python
    import pulumi
    import pulumi_vcd as vcd

    by_name_and_cidr = vcd.get_nsxt_edgegateway_static_route(edge_gateway_id=existing["id"],
        name="duplicate-name-sr",
        network_cidr="10.10.11.0/24")
    ```


    :param str edge_gateway_id: NSX-T Edge Gateway ID
    :param str name: Name of Static Route. **Note** names *can be duplicate* and one can use
           `network_cidr` to make filtering more precise
    :param str network_cidr: Network CIDR for Static Route
           
           > It may happen that there are multiple NSX-T Static Routes with the same `name`. In such a case, a
           data source will return an error as it expects to find only one entity. If this happens, one can
           make the filtering more precise by supplying `network_cidr` in addition to `name`.
    """
    __args__ = dict()
    __args__['edgeGatewayId'] = edge_gateway_id
    __args__['name'] = name
    __args__['networkCidr'] = network_cidr
    __args__['org'] = org
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getNsxtEdgegatewayStaticRoute:getNsxtEdgegatewayStaticRoute', __args__, opts=opts, typ=GetNsxtEdgegatewayStaticRouteResult).value

    return AwaitableGetNsxtEdgegatewayStaticRouteResult(
        description=pulumi.get(__ret__, 'description'),
        edge_gateway_id=pulumi.get(__ret__, 'edge_gateway_id'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        network_cidr=pulumi.get(__ret__, 'network_cidr'),
        next_hops=pulumi.get(__ret__, 'next_hops'),
        org=pulumi.get(__ret__, 'org'))
def get_nsxt_edgegateway_static_route_output(edge_gateway_id: Optional[pulumi.Input[str]] = None,
                                             name: Optional[pulumi.Input[str]] = None,
                                             network_cidr: Optional[pulumi.Input[Optional[str]]] = None,
                                             org: Optional[pulumi.Input[Optional[str]]] = None,
                                             opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetNsxtEdgegatewayStaticRouteResult]:
    """
    Supported in provider *v3.10+* and VCD 10.4.0+ with NSX-T.

    Provides a data source to read NSX-T Edge Gateway Static Routes.

    ## Example Usage

    ### By Name Only)

    ```python
    import pulumi
    import pulumi_vcd as vcd

    by_name = vcd.get_nsxt_edgegateway_static_route(edge_gateway_id=existing["id"],
        name="existing-static-route")
    ```

    ### By Name And Network CIDR )

    ```python
    import pulumi
    import pulumi_vcd as vcd

    by_name_and_cidr = vcd.get_nsxt_edgegateway_static_route(edge_gateway_id=existing["id"],
        name="duplicate-name-sr",
        network_cidr="10.10.11.0/24")
    ```


    :param str edge_gateway_id: NSX-T Edge Gateway ID
    :param str name: Name of Static Route. **Note** names *can be duplicate* and one can use
           `network_cidr` to make filtering more precise
    :param str network_cidr: Network CIDR for Static Route
           
           > It may happen that there are multiple NSX-T Static Routes with the same `name`. In such a case, a
           data source will return an error as it expects to find only one entity. If this happens, one can
           make the filtering more precise by supplying `network_cidr` in addition to `name`.
    """
    __args__ = dict()
    __args__['edgeGatewayId'] = edge_gateway_id
    __args__['name'] = name
    __args__['networkCidr'] = network_cidr
    __args__['org'] = org
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vcd:index/getNsxtEdgegatewayStaticRoute:getNsxtEdgegatewayStaticRoute', __args__, opts=opts, typ=GetNsxtEdgegatewayStaticRouteResult)
    return __ret__.apply(lambda __response__: GetNsxtEdgegatewayStaticRouteResult(
        description=pulumi.get(__response__, 'description'),
        edge_gateway_id=pulumi.get(__response__, 'edge_gateway_id'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        network_cidr=pulumi.get(__response__, 'network_cidr'),
        next_hops=pulumi.get(__response__, 'next_hops'),
        org=pulumi.get(__response__, 'org')))
