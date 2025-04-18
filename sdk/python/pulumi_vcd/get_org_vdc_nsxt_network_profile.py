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
    'GetOrgVdcNsxtNetworkProfileResult',
    'AwaitableGetOrgVdcNsxtNetworkProfileResult',
    'get_org_vdc_nsxt_network_profile',
    'get_org_vdc_nsxt_network_profile_output',
]

@pulumi.output_type
class GetOrgVdcNsxtNetworkProfileResult:
    """
    A collection of values returned by getOrgVdcNsxtNetworkProfile.
    """
    def __init__(__self__, edge_cluster_id=None, id=None, org=None, vapp_networks_default_segment_profile_template_id=None, vdc=None, vdc_networks_default_segment_profile_template_id=None):
        if edge_cluster_id and not isinstance(edge_cluster_id, str):
            raise TypeError("Expected argument 'edge_cluster_id' to be a str")
        pulumi.set(__self__, "edge_cluster_id", edge_cluster_id)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if org and not isinstance(org, str):
            raise TypeError("Expected argument 'org' to be a str")
        pulumi.set(__self__, "org", org)
        if vapp_networks_default_segment_profile_template_id and not isinstance(vapp_networks_default_segment_profile_template_id, str):
            raise TypeError("Expected argument 'vapp_networks_default_segment_profile_template_id' to be a str")
        pulumi.set(__self__, "vapp_networks_default_segment_profile_template_id", vapp_networks_default_segment_profile_template_id)
        if vdc and not isinstance(vdc, str):
            raise TypeError("Expected argument 'vdc' to be a str")
        pulumi.set(__self__, "vdc", vdc)
        if vdc_networks_default_segment_profile_template_id and not isinstance(vdc_networks_default_segment_profile_template_id, str):
            raise TypeError("Expected argument 'vdc_networks_default_segment_profile_template_id' to be a str")
        pulumi.set(__self__, "vdc_networks_default_segment_profile_template_id", vdc_networks_default_segment_profile_template_id)

    @property
    @pulumi.getter(name="edgeClusterId")
    def edge_cluster_id(self) -> str:
        """
        An ID of NSX-T Edge Cluster which should provide vApp
        Networking Services or DHCP for Isolated Networks. This value **might be unavailable** in data
        source if user has insufficient rights.
        """
        return pulumi.get(self, "edge_cluster_id")

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
    @pulumi.getter(name="vappNetworksDefaultSegmentProfileTemplateId")
    def vapp_networks_default_segment_profile_template_id(self) -> str:
        """
        Default Segment Profile ID for all vApp
        networks withing this VDC
        """
        return pulumi.get(self, "vapp_networks_default_segment_profile_template_id")

    @property
    @pulumi.getter
    def vdc(self) -> Optional[str]:
        return pulumi.get(self, "vdc")

    @property
    @pulumi.getter(name="vdcNetworksDefaultSegmentProfileTemplateId")
    def vdc_networks_default_segment_profile_template_id(self) -> str:
        """
        Default Segment Profile ID for all Org VDC
        networks withing this VDC
        """
        return pulumi.get(self, "vdc_networks_default_segment_profile_template_id")


class AwaitableGetOrgVdcNsxtNetworkProfileResult(GetOrgVdcNsxtNetworkProfileResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOrgVdcNsxtNetworkProfileResult(
            edge_cluster_id=self.edge_cluster_id,
            id=self.id,
            org=self.org,
            vapp_networks_default_segment_profile_template_id=self.vapp_networks_default_segment_profile_template_id,
            vdc=self.vdc,
            vdc_networks_default_segment_profile_template_id=self.vdc_networks_default_segment_profile_template_id)


def get_org_vdc_nsxt_network_profile(org: Optional[str] = None,
                                     vdc: Optional[str] = None,
                                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOrgVdcNsxtNetworkProfileResult:
    """
    Provides a data source to read Network Profile for NSX-T VDCs.

    Supported in provider *v3.11+* and VCD 10.4.0+ with NSX-T.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vcd as vcd

    nsxt = vcd.get_org_vdc_nsxt_network_profile(org="my-org",
        vdc="my-vdc")
    ```


    :param str org: The name of organization to use, optional if defined at provider level
    :param str vdc: The name of VDC to use, optional if defined at provider level
    """
    __args__ = dict()
    __args__['org'] = org
    __args__['vdc'] = vdc
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getOrgVdcNsxtNetworkProfile:getOrgVdcNsxtNetworkProfile', __args__, opts=opts, typ=GetOrgVdcNsxtNetworkProfileResult).value

    return AwaitableGetOrgVdcNsxtNetworkProfileResult(
        edge_cluster_id=pulumi.get(__ret__, 'edge_cluster_id'),
        id=pulumi.get(__ret__, 'id'),
        org=pulumi.get(__ret__, 'org'),
        vapp_networks_default_segment_profile_template_id=pulumi.get(__ret__, 'vapp_networks_default_segment_profile_template_id'),
        vdc=pulumi.get(__ret__, 'vdc'),
        vdc_networks_default_segment_profile_template_id=pulumi.get(__ret__, 'vdc_networks_default_segment_profile_template_id'))
def get_org_vdc_nsxt_network_profile_output(org: Optional[pulumi.Input[Optional[str]]] = None,
                                            vdc: Optional[pulumi.Input[Optional[str]]] = None,
                                            opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetOrgVdcNsxtNetworkProfileResult]:
    """
    Provides a data source to read Network Profile for NSX-T VDCs.

    Supported in provider *v3.11+* and VCD 10.4.0+ with NSX-T.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vcd as vcd

    nsxt = vcd.get_org_vdc_nsxt_network_profile(org="my-org",
        vdc="my-vdc")
    ```


    :param str org: The name of organization to use, optional if defined at provider level
    :param str vdc: The name of VDC to use, optional if defined at provider level
    """
    __args__ = dict()
    __args__['org'] = org
    __args__['vdc'] = vdc
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vcd:index/getOrgVdcNsxtNetworkProfile:getOrgVdcNsxtNetworkProfile', __args__, opts=opts, typ=GetOrgVdcNsxtNetworkProfileResult)
    return __ret__.apply(lambda __response__: GetOrgVdcNsxtNetworkProfileResult(
        edge_cluster_id=pulumi.get(__response__, 'edge_cluster_id'),
        id=pulumi.get(__response__, 'id'),
        org=pulumi.get(__response__, 'org'),
        vapp_networks_default_segment_profile_template_id=pulumi.get(__response__, 'vapp_networks_default_segment_profile_template_id'),
        vdc=pulumi.get(__response__, 'vdc'),
        vdc_networks_default_segment_profile_template_id=pulumi.get(__response__, 'vdc_networks_default_segment_profile_template_id')))
