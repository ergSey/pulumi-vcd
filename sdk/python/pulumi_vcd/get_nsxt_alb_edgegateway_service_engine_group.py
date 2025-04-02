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
    'GetNsxtAlbEdgegatewayServiceEngineGroupResult',
    'AwaitableGetNsxtAlbEdgegatewayServiceEngineGroupResult',
    'get_nsxt_alb_edgegateway_service_engine_group',
    'get_nsxt_alb_edgegateway_service_engine_group_output',
]

@pulumi.output_type
class GetNsxtAlbEdgegatewayServiceEngineGroupResult:
    """
    A collection of values returned by getNsxtAlbEdgegatewayServiceEngineGroup.
    """
    def __init__(__self__, deployed_virtual_services=None, edge_gateway_id=None, id=None, max_virtual_services=None, org=None, reserved_virtual_services=None, service_engine_group_id=None, service_engine_group_name=None, vdc=None):
        if deployed_virtual_services and not isinstance(deployed_virtual_services, int):
            raise TypeError("Expected argument 'deployed_virtual_services' to be a int")
        pulumi.set(__self__, "deployed_virtual_services", deployed_virtual_services)
        if edge_gateway_id and not isinstance(edge_gateway_id, str):
            raise TypeError("Expected argument 'edge_gateway_id' to be a str")
        pulumi.set(__self__, "edge_gateway_id", edge_gateway_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if max_virtual_services and not isinstance(max_virtual_services, int):
            raise TypeError("Expected argument 'max_virtual_services' to be a int")
        pulumi.set(__self__, "max_virtual_services", max_virtual_services)
        if org and not isinstance(org, str):
            raise TypeError("Expected argument 'org' to be a str")
        pulumi.set(__self__, "org", org)
        if reserved_virtual_services and not isinstance(reserved_virtual_services, str):
            raise TypeError("Expected argument 'reserved_virtual_services' to be a str")
        pulumi.set(__self__, "reserved_virtual_services", reserved_virtual_services)
        if service_engine_group_id and not isinstance(service_engine_group_id, str):
            raise TypeError("Expected argument 'service_engine_group_id' to be a str")
        pulumi.set(__self__, "service_engine_group_id", service_engine_group_id)
        if service_engine_group_name and not isinstance(service_engine_group_name, str):
            raise TypeError("Expected argument 'service_engine_group_name' to be a str")
        pulumi.set(__self__, "service_engine_group_name", service_engine_group_name)
        if vdc and not isinstance(vdc, str):
            raise TypeError("Expected argument 'vdc' to be a str")
        pulumi.set(__self__, "vdc", vdc)

    @property
    @pulumi.getter(name="deployedVirtualServices")
    def deployed_virtual_services(self) -> int:
        return pulumi.get(self, "deployed_virtual_services")

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
    @pulumi.getter(name="maxVirtualServices")
    def max_virtual_services(self) -> int:
        return pulumi.get(self, "max_virtual_services")

    @property
    @pulumi.getter
    def org(self) -> Optional[str]:
        return pulumi.get(self, "org")

    @property
    @pulumi.getter(name="reservedVirtualServices")
    def reserved_virtual_services(self) -> str:
        return pulumi.get(self, "reserved_virtual_services")

    @property
    @pulumi.getter(name="serviceEngineGroupId")
    def service_engine_group_id(self) -> str:
        return pulumi.get(self, "service_engine_group_id")

    @property
    @pulumi.getter(name="serviceEngineGroupName")
    def service_engine_group_name(self) -> str:
        return pulumi.get(self, "service_engine_group_name")

    @property
    @pulumi.getter
    @_utilities.deprecated("""Edge Gateway will be looked up based on 'edge_gateway_id' field""")
    def vdc(self) -> str:
        return pulumi.get(self, "vdc")


class AwaitableGetNsxtAlbEdgegatewayServiceEngineGroupResult(GetNsxtAlbEdgegatewayServiceEngineGroupResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNsxtAlbEdgegatewayServiceEngineGroupResult(
            deployed_virtual_services=self.deployed_virtual_services,
            edge_gateway_id=self.edge_gateway_id,
            id=self.id,
            max_virtual_services=self.max_virtual_services,
            org=self.org,
            reserved_virtual_services=self.reserved_virtual_services,
            service_engine_group_id=self.service_engine_group_id,
            service_engine_group_name=self.service_engine_group_name,
            vdc=self.vdc)


def get_nsxt_alb_edgegateway_service_engine_group(edge_gateway_id: Optional[str] = None,
                                                  org: Optional[str] = None,
                                                  service_engine_group_id: Optional[str] = None,
                                                  service_engine_group_name: Optional[str] = None,
                                                  vdc: Optional[str] = None,
                                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNsxtAlbEdgegatewayServiceEngineGroupResult:
    """
    Supported in provider *v3.5+* and VCD 10.2+ with NSX-T and ALB.

    Provides a datasource to read ALB Service Engine Group assignment to NSX-T Edge Gateway.

    ## Example Usage

    ### Referencing Service Engine Group By ID)

    ```python
    import pulumi
    import pulumi_vcd as vcd

    existing = vcd.get_nsxt_edgegateway(org="my-org",
        vdc="nsxt-vdc",
        name="nsxt-gw")
    first = vcd.get_nsxt_alb_service_engine_group(name="first-se")
    test = vcd.get_nsxt_alb_edgegateway_service_engine_group(edge_gateway_id=existing.id,
        service_engine_group_id=first.id)
    ```

    ### Referencing Service Engine Group By Name)

    ```python
    import pulumi
    import pulumi_vcd as vcd

    existing = vcd.get_nsxt_edgegateway(org="my-org",
        vdc="nsxt-vdc",
        name="nsxt-gw")
    test = vcd.get_nsxt_alb_edgegateway_service_engine_group(edge_gateway_id=existing.id,
        service_engine_group_name="known-service-engine-group-name")
    ```


    :param str edge_gateway_id: An ID of NSX-T Edge Gateway. Can be looked up using
           [NsxtEdgegateway](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_edgegateway) data source
    :param str org: The name of organization to which the edge gateway belongs. Optional if defined at provider level.
    :param str service_engine_group_id: An ID of NSX-T Service Engine Group. Can be looked up using
           [NsxtAlbServiceEngineGroup](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_alb_service_engine_group) data
           source. **Note** Either `service_engine_group_name` or `service_engine_group_id` require it.
    :param str service_engine_group_name: A Name of NSX-T Service Engine Group. **Note** Either
           `service_engine_group_name` or `service_engine_group_id` require it.
    """
    __args__ = dict()
    __args__['edgeGatewayId'] = edge_gateway_id
    __args__['org'] = org
    __args__['serviceEngineGroupId'] = service_engine_group_id
    __args__['serviceEngineGroupName'] = service_engine_group_name
    __args__['vdc'] = vdc
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getNsxtAlbEdgegatewayServiceEngineGroup:getNsxtAlbEdgegatewayServiceEngineGroup', __args__, opts=opts, typ=GetNsxtAlbEdgegatewayServiceEngineGroupResult).value

    return AwaitableGetNsxtAlbEdgegatewayServiceEngineGroupResult(
        deployed_virtual_services=pulumi.get(__ret__, 'deployed_virtual_services'),
        edge_gateway_id=pulumi.get(__ret__, 'edge_gateway_id'),
        id=pulumi.get(__ret__, 'id'),
        max_virtual_services=pulumi.get(__ret__, 'max_virtual_services'),
        org=pulumi.get(__ret__, 'org'),
        reserved_virtual_services=pulumi.get(__ret__, 'reserved_virtual_services'),
        service_engine_group_id=pulumi.get(__ret__, 'service_engine_group_id'),
        service_engine_group_name=pulumi.get(__ret__, 'service_engine_group_name'),
        vdc=pulumi.get(__ret__, 'vdc'))
def get_nsxt_alb_edgegateway_service_engine_group_output(edge_gateway_id: Optional[pulumi.Input[str]] = None,
                                                         org: Optional[pulumi.Input[Optional[str]]] = None,
                                                         service_engine_group_id: Optional[pulumi.Input[Optional[str]]] = None,
                                                         service_engine_group_name: Optional[pulumi.Input[Optional[str]]] = None,
                                                         vdc: Optional[pulumi.Input[Optional[str]]] = None,
                                                         opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetNsxtAlbEdgegatewayServiceEngineGroupResult]:
    """
    Supported in provider *v3.5+* and VCD 10.2+ with NSX-T and ALB.

    Provides a datasource to read ALB Service Engine Group assignment to NSX-T Edge Gateway.

    ## Example Usage

    ### Referencing Service Engine Group By ID)

    ```python
    import pulumi
    import pulumi_vcd as vcd

    existing = vcd.get_nsxt_edgegateway(org="my-org",
        vdc="nsxt-vdc",
        name="nsxt-gw")
    first = vcd.get_nsxt_alb_service_engine_group(name="first-se")
    test = vcd.get_nsxt_alb_edgegateway_service_engine_group(edge_gateway_id=existing.id,
        service_engine_group_id=first.id)
    ```

    ### Referencing Service Engine Group By Name)

    ```python
    import pulumi
    import pulumi_vcd as vcd

    existing = vcd.get_nsxt_edgegateway(org="my-org",
        vdc="nsxt-vdc",
        name="nsxt-gw")
    test = vcd.get_nsxt_alb_edgegateway_service_engine_group(edge_gateway_id=existing.id,
        service_engine_group_name="known-service-engine-group-name")
    ```


    :param str edge_gateway_id: An ID of NSX-T Edge Gateway. Can be looked up using
           [NsxtEdgegateway](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_edgegateway) data source
    :param str org: The name of organization to which the edge gateway belongs. Optional if defined at provider level.
    :param str service_engine_group_id: An ID of NSX-T Service Engine Group. Can be looked up using
           [NsxtAlbServiceEngineGroup](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_alb_service_engine_group) data
           source. **Note** Either `service_engine_group_name` or `service_engine_group_id` require it.
    :param str service_engine_group_name: A Name of NSX-T Service Engine Group. **Note** Either
           `service_engine_group_name` or `service_engine_group_id` require it.
    """
    __args__ = dict()
    __args__['edgeGatewayId'] = edge_gateway_id
    __args__['org'] = org
    __args__['serviceEngineGroupId'] = service_engine_group_id
    __args__['serviceEngineGroupName'] = service_engine_group_name
    __args__['vdc'] = vdc
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vcd:index/getNsxtAlbEdgegatewayServiceEngineGroup:getNsxtAlbEdgegatewayServiceEngineGroup', __args__, opts=opts, typ=GetNsxtAlbEdgegatewayServiceEngineGroupResult)
    return __ret__.apply(lambda __response__: GetNsxtAlbEdgegatewayServiceEngineGroupResult(
        deployed_virtual_services=pulumi.get(__response__, 'deployed_virtual_services'),
        edge_gateway_id=pulumi.get(__response__, 'edge_gateway_id'),
        id=pulumi.get(__response__, 'id'),
        max_virtual_services=pulumi.get(__response__, 'max_virtual_services'),
        org=pulumi.get(__response__, 'org'),
        reserved_virtual_services=pulumi.get(__response__, 'reserved_virtual_services'),
        service_engine_group_id=pulumi.get(__response__, 'service_engine_group_id'),
        service_engine_group_name=pulumi.get(__response__, 'service_engine_group_name'),
        vdc=pulumi.get(__response__, 'vdc')))
