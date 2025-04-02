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
    'GetNsxtAlbSettingsResult',
    'AwaitableGetNsxtAlbSettingsResult',
    'get_nsxt_alb_settings',
    'get_nsxt_alb_settings_output',
]

@pulumi.output_type
class GetNsxtAlbSettingsResult:
    """
    A collection of values returned by getNsxtAlbSettings.
    """
    def __init__(__self__, edge_gateway_id=None, id=None, ipv6_service_network_specification=None, is_active=None, is_transparent_mode_enabled=None, org=None, service_network_specification=None, supported_feature_set=None, vdc=None):
        if edge_gateway_id and not isinstance(edge_gateway_id, str):
            raise TypeError("Expected argument 'edge_gateway_id' to be a str")
        pulumi.set(__self__, "edge_gateway_id", edge_gateway_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ipv6_service_network_specification and not isinstance(ipv6_service_network_specification, str):
            raise TypeError("Expected argument 'ipv6_service_network_specification' to be a str")
        pulumi.set(__self__, "ipv6_service_network_specification", ipv6_service_network_specification)
        if is_active and not isinstance(is_active, bool):
            raise TypeError("Expected argument 'is_active' to be a bool")
        pulumi.set(__self__, "is_active", is_active)
        if is_transparent_mode_enabled and not isinstance(is_transparent_mode_enabled, bool):
            raise TypeError("Expected argument 'is_transparent_mode_enabled' to be a bool")
        pulumi.set(__self__, "is_transparent_mode_enabled", is_transparent_mode_enabled)
        if org and not isinstance(org, str):
            raise TypeError("Expected argument 'org' to be a str")
        pulumi.set(__self__, "org", org)
        if service_network_specification and not isinstance(service_network_specification, str):
            raise TypeError("Expected argument 'service_network_specification' to be a str")
        pulumi.set(__self__, "service_network_specification", service_network_specification)
        if supported_feature_set and not isinstance(supported_feature_set, str):
            raise TypeError("Expected argument 'supported_feature_set' to be a str")
        pulumi.set(__self__, "supported_feature_set", supported_feature_set)
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
    @pulumi.getter(name="ipv6ServiceNetworkSpecification")
    def ipv6_service_network_specification(self) -> str:
        return pulumi.get(self, "ipv6_service_network_specification")

    @property
    @pulumi.getter(name="isActive")
    def is_active(self) -> bool:
        return pulumi.get(self, "is_active")

    @property
    @pulumi.getter(name="isTransparentModeEnabled")
    def is_transparent_mode_enabled(self) -> bool:
        return pulumi.get(self, "is_transparent_mode_enabled")

    @property
    @pulumi.getter
    def org(self) -> Optional[str]:
        return pulumi.get(self, "org")

    @property
    @pulumi.getter(name="serviceNetworkSpecification")
    def service_network_specification(self) -> str:
        return pulumi.get(self, "service_network_specification")

    @property
    @pulumi.getter(name="supportedFeatureSet")
    def supported_feature_set(self) -> str:
        return pulumi.get(self, "supported_feature_set")

    @property
    @pulumi.getter
    @_utilities.deprecated("""Edge Gateway will be looked up based on 'edge_gateway_id' field""")
    def vdc(self) -> Optional[str]:
        return pulumi.get(self, "vdc")


class AwaitableGetNsxtAlbSettingsResult(GetNsxtAlbSettingsResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNsxtAlbSettingsResult(
            edge_gateway_id=self.edge_gateway_id,
            id=self.id,
            ipv6_service_network_specification=self.ipv6_service_network_specification,
            is_active=self.is_active,
            is_transparent_mode_enabled=self.is_transparent_mode_enabled,
            org=self.org,
            service_network_specification=self.service_network_specification,
            supported_feature_set=self.supported_feature_set,
            vdc=self.vdc)


def get_nsxt_alb_settings(edge_gateway_id: Optional[str] = None,
                          org: Optional[str] = None,
                          service_network_specification: Optional[str] = None,
                          vdc: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNsxtAlbSettingsResult:
    """
    Supported in provider *v3.5+* and VCD 10.2+ with NSX-T and ALB.

    Provides a data source to read ALB General Settings for particular NSX-T Edge Gateway.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vcd as vcd

    existing = vcd.get_nsxt_edgegateway(org="my-org",
        vdc="nsxt-vdc",
        name="nsxt-gw")
    test = vcd.get_nsxt_alb_settings(org="my-org",
        edge_gateway_id=existing.id)
    ```


    :param str edge_gateway_id: An ID of NSX-T Edge Gateway. Can be looked up using
           [NsxtEdgegateway](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_edgegateway) data source
    :param str org: The name of organization to which the edge gateway belongs. Optional if defined at provider level.
    """
    __args__ = dict()
    __args__['edgeGatewayId'] = edge_gateway_id
    __args__['org'] = org
    __args__['serviceNetworkSpecification'] = service_network_specification
    __args__['vdc'] = vdc
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getNsxtAlbSettings:getNsxtAlbSettings', __args__, opts=opts, typ=GetNsxtAlbSettingsResult).value

    return AwaitableGetNsxtAlbSettingsResult(
        edge_gateway_id=pulumi.get(__ret__, 'edge_gateway_id'),
        id=pulumi.get(__ret__, 'id'),
        ipv6_service_network_specification=pulumi.get(__ret__, 'ipv6_service_network_specification'),
        is_active=pulumi.get(__ret__, 'is_active'),
        is_transparent_mode_enabled=pulumi.get(__ret__, 'is_transparent_mode_enabled'),
        org=pulumi.get(__ret__, 'org'),
        service_network_specification=pulumi.get(__ret__, 'service_network_specification'),
        supported_feature_set=pulumi.get(__ret__, 'supported_feature_set'),
        vdc=pulumi.get(__ret__, 'vdc'))
def get_nsxt_alb_settings_output(edge_gateway_id: Optional[pulumi.Input[str]] = None,
                                 org: Optional[pulumi.Input[Optional[str]]] = None,
                                 service_network_specification: Optional[pulumi.Input[Optional[str]]] = None,
                                 vdc: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetNsxtAlbSettingsResult]:
    """
    Supported in provider *v3.5+* and VCD 10.2+ with NSX-T and ALB.

    Provides a data source to read ALB General Settings for particular NSX-T Edge Gateway.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vcd as vcd

    existing = vcd.get_nsxt_edgegateway(org="my-org",
        vdc="nsxt-vdc",
        name="nsxt-gw")
    test = vcd.get_nsxt_alb_settings(org="my-org",
        edge_gateway_id=existing.id)
    ```


    :param str edge_gateway_id: An ID of NSX-T Edge Gateway. Can be looked up using
           [NsxtEdgegateway](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/nsxt_edgegateway) data source
    :param str org: The name of organization to which the edge gateway belongs. Optional if defined at provider level.
    """
    __args__ = dict()
    __args__['edgeGatewayId'] = edge_gateway_id
    __args__['org'] = org
    __args__['serviceNetworkSpecification'] = service_network_specification
    __args__['vdc'] = vdc
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vcd:index/getNsxtAlbSettings:getNsxtAlbSettings', __args__, opts=opts, typ=GetNsxtAlbSettingsResult)
    return __ret__.apply(lambda __response__: GetNsxtAlbSettingsResult(
        edge_gateway_id=pulumi.get(__response__, 'edge_gateway_id'),
        id=pulumi.get(__response__, 'id'),
        ipv6_service_network_specification=pulumi.get(__response__, 'ipv6_service_network_specification'),
        is_active=pulumi.get(__response__, 'is_active'),
        is_transparent_mode_enabled=pulumi.get(__response__, 'is_transparent_mode_enabled'),
        org=pulumi.get(__response__, 'org'),
        service_network_specification=pulumi.get(__response__, 'service_network_specification'),
        supported_feature_set=pulumi.get(__response__, 'supported_feature_set'),
        vdc=pulumi.get(__response__, 'vdc')))
