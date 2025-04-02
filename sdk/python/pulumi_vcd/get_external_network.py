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
    'GetExternalNetworkResult',
    'AwaitableGetExternalNetworkResult',
    'get_external_network',
    'get_external_network_output',
]

@pulumi.output_type
class GetExternalNetworkResult:
    """
    A collection of values returned by getExternalNetwork.
    """
    def __init__(__self__, description=None, id=None, ip_scopes=None, name=None, retain_net_info_across_deployments=None, vsphere_networks=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ip_scopes and not isinstance(ip_scopes, list):
            raise TypeError("Expected argument 'ip_scopes' to be a list")
        pulumi.set(__self__, "ip_scopes", ip_scopes)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if retain_net_info_across_deployments and not isinstance(retain_net_info_across_deployments, bool):
            raise TypeError("Expected argument 'retain_net_info_across_deployments' to be a bool")
        pulumi.set(__self__, "retain_net_info_across_deployments", retain_net_info_across_deployments)
        if vsphere_networks and not isinstance(vsphere_networks, list):
            raise TypeError("Expected argument 'vsphere_networks' to be a list")
        pulumi.set(__self__, "vsphere_networks", vsphere_networks)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Network friendly description
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
    @pulumi.getter(name="ipScopes")
    def ip_scopes(self) -> Sequence['outputs.GetExternalNetworkIpScopeResult']:
        """
        A list of IP scopes for the network. See [IP Scope](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/external_network#ipscope)
        for details.
        """
        return pulumi.get(self, "ip_scopes")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="retainNetInfoAcrossDeployments")
    def retain_net_info_across_deployments(self) -> bool:
        """
        Specifies whether the network resources such as IP/MAC of router will be 
        retained across deployments.
        """
        return pulumi.get(self, "retain_net_info_across_deployments")

    @property
    @pulumi.getter(name="vsphereNetworks")
    def vsphere_networks(self) -> Sequence['outputs.GetExternalNetworkVsphereNetworkResult']:
        """
        A list of DV_PORTGROUP or NETWORK objects names that back this network. Each referenced 
        DV_PORTGROUP or NETWORK must exist on a vCenter server registered with the system.
        See [vSphere Network](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/external_network#vspherenetwork) for details.
        """
        return pulumi.get(self, "vsphere_networks")


class AwaitableGetExternalNetworkResult(GetExternalNetworkResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetExternalNetworkResult(
            description=self.description,
            id=self.id,
            ip_scopes=self.ip_scopes,
            name=self.name,
            retain_net_info_across_deployments=self.retain_net_info_across_deployments,
            vsphere_networks=self.vsphere_networks)


def get_external_network(name: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetExternalNetworkResult:
    """
    Provides a VMware Cloud Director external network data source. This can be used to reference external networks and their properties.

    Supported in provider *v2.5+*

    > This resource is deprecated in favor of [`ExternalNetworkV2`](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/external_network_v2)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vcd as vcd

    tf_external_network = vcd.get_external_network(name="my-extnet")
    tf_nat_rule = vcd.index.Dnat("tf-nat-rule",
        org=tf-org,
        vdc=tf-vdc,
        network_name=tf_external_network.name,
        network_type=ext,
        edge_gateway=tf-gw,
        external_ip=extnet_datacloud.ip_scope[0].static_ip_pool[0].start_address,
        port=7777,
        protocol=tcp,
        internal_ip=10.10.102.60,
        translated_port=77,
        description=test run)
    ```


    :param str name: external network name
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getExternalNetwork:getExternalNetwork', __args__, opts=opts, typ=GetExternalNetworkResult).value

    return AwaitableGetExternalNetworkResult(
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        ip_scopes=pulumi.get(__ret__, 'ip_scopes'),
        name=pulumi.get(__ret__, 'name'),
        retain_net_info_across_deployments=pulumi.get(__ret__, 'retain_net_info_across_deployments'),
        vsphere_networks=pulumi.get(__ret__, 'vsphere_networks'))
def get_external_network_output(name: Optional[pulumi.Input[str]] = None,
                                opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetExternalNetworkResult]:
    """
    Provides a VMware Cloud Director external network data source. This can be used to reference external networks and their properties.

    Supported in provider *v2.5+*

    > This resource is deprecated in favor of [`ExternalNetworkV2`](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/external_network_v2)

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vcd as vcd

    tf_external_network = vcd.get_external_network(name="my-extnet")
    tf_nat_rule = vcd.index.Dnat("tf-nat-rule",
        org=tf-org,
        vdc=tf-vdc,
        network_name=tf_external_network.name,
        network_type=ext,
        edge_gateway=tf-gw,
        external_ip=extnet_datacloud.ip_scope[0].static_ip_pool[0].start_address,
        port=7777,
        protocol=tcp,
        internal_ip=10.10.102.60,
        translated_port=77,
        description=test run)
    ```


    :param str name: external network name
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vcd:index/getExternalNetwork:getExternalNetwork', __args__, opts=opts, typ=GetExternalNetworkResult)
    return __ret__.apply(lambda __response__: GetExternalNetworkResult(
        description=pulumi.get(__response__, 'description'),
        id=pulumi.get(__response__, 'id'),
        ip_scopes=pulumi.get(__response__, 'ip_scopes'),
        name=pulumi.get(__response__, 'name'),
        retain_net_info_across_deployments=pulumi.get(__response__, 'retain_net_info_across_deployments'),
        vsphere_networks=pulumi.get(__response__, 'vsphere_networks')))
