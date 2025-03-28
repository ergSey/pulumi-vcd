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
    'GetResourcePoolResult',
    'AwaitableGetResourcePoolResult',
    'get_resource_pool',
    'get_resource_pool_output',
]

@pulumi.output_type
class GetResourcePoolResult:
    """
    A collection of values returned by getResourcePool.
    """
    def __init__(__self__, cluster_moref=None, hardware_version=None, id=None, name=None, vcenter_id=None):
        if cluster_moref and not isinstance(cluster_moref, str):
            raise TypeError("Expected argument 'cluster_moref' to be a str")
        pulumi.set(__self__, "cluster_moref", cluster_moref)
        if hardware_version and not isinstance(hardware_version, str):
            raise TypeError("Expected argument 'hardware_version' to be a str")
        pulumi.set(__self__, "hardware_version", hardware_version)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if vcenter_id and not isinstance(vcenter_id, str):
            raise TypeError("Expected argument 'vcenter_id' to be a str")
        pulumi.set(__self__, "vcenter_id", vcenter_id)

    @property
    @pulumi.getter(name="clusterMoref")
    def cluster_moref(self) -> str:
        """
        managed object reference of the vCenter cluster that this resource pool is hosted on.
        """
        return pulumi.get(self, "cluster_moref")

    @property
    @pulumi.getter(name="hardwareVersion")
    def hardware_version(self) -> str:
        """
        default hardware version available to this resource pool.
        """
        return pulumi.get(self, "hardware_version")

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
    @pulumi.getter(name="vcenterId")
    def vcenter_id(self) -> str:
        return pulumi.get(self, "vcenter_id")


class AwaitableGetResourcePoolResult(GetResourcePoolResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetResourcePoolResult(
            cluster_moref=self.cluster_moref,
            hardware_version=self.hardware_version,
            id=self.id,
            name=self.name,
            vcenter_id=self.vcenter_id)


def get_resource_pool(name: Optional[str] = None,
                      vcenter_id: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetResourcePoolResult:
    """
    Provides a data source for a resource pool attached to a vCenter. A resource pool is an essential component of a Provider VDC.

    > Note 1: this data source requires System Administrator privileges

    > Note 2: you can create or modify a resource pool using vSphere provider

    Supported in provider *v3.10+*

    ## Example Usage

    ### 1

    ```python
    import pulumi
    import pulumi_vcd as vcd

    vcenter1 = vcd.get_vcenter(name="vc1")
    rp1 = vcd.get_resource_pool(name="resource-pool-for-vcd-01",
        vcenter_id=vcenter1.id)
    ```

    ### 2

    ```python
    import pulumi
    import pulumi_vcd as vcd

    rp1 = vcd.get_resource_pool(name="common-name",
        vcenter_id=vcenter1["id"])
    ```

    When you receive such error, you can run the script again, but using the resource pool ID instead of the name.

    ```python
    import pulumi
    import pulumi_vcd as vcd

    rp1 = vcd.get_resource_pool(name="resgroup-241",
        vcenter_id=vcenter1["id"])
    ```


    :param str name: resource pool name. The name may not be unique within the vCenter. If that happens, you will get an
           error message with the list of IDs for the pools with the same name, and can subsequently enter the resource pool ID instead of the name.
           (See Example Usage 2)
    :param str vcenter_id: ID of the vCenter to which this resource pool belongs.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vcenterId'] = vcenter_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getResourcePool:getResourcePool', __args__, opts=opts, typ=GetResourcePoolResult).value

    return AwaitableGetResourcePoolResult(
        cluster_moref=pulumi.get(__ret__, 'cluster_moref'),
        hardware_version=pulumi.get(__ret__, 'hardware_version'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        vcenter_id=pulumi.get(__ret__, 'vcenter_id'))
def get_resource_pool_output(name: Optional[pulumi.Input[str]] = None,
                             vcenter_id: Optional[pulumi.Input[str]] = None,
                             opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetResourcePoolResult]:
    """
    Provides a data source for a resource pool attached to a vCenter. A resource pool is an essential component of a Provider VDC.

    > Note 1: this data source requires System Administrator privileges

    > Note 2: you can create or modify a resource pool using vSphere provider

    Supported in provider *v3.10+*

    ## Example Usage

    ### 1

    ```python
    import pulumi
    import pulumi_vcd as vcd

    vcenter1 = vcd.get_vcenter(name="vc1")
    rp1 = vcd.get_resource_pool(name="resource-pool-for-vcd-01",
        vcenter_id=vcenter1.id)
    ```

    ### 2

    ```python
    import pulumi
    import pulumi_vcd as vcd

    rp1 = vcd.get_resource_pool(name="common-name",
        vcenter_id=vcenter1["id"])
    ```

    When you receive such error, you can run the script again, but using the resource pool ID instead of the name.

    ```python
    import pulumi
    import pulumi_vcd as vcd

    rp1 = vcd.get_resource_pool(name="resgroup-241",
        vcenter_id=vcenter1["id"])
    ```


    :param str name: resource pool name. The name may not be unique within the vCenter. If that happens, you will get an
           error message with the list of IDs for the pools with the same name, and can subsequently enter the resource pool ID instead of the name.
           (See Example Usage 2)
    :param str vcenter_id: ID of the vCenter to which this resource pool belongs.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vcenterId'] = vcenter_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vcd:index/getResourcePool:getResourcePool', __args__, opts=opts, typ=GetResourcePoolResult)
    return __ret__.apply(lambda __response__: GetResourcePoolResult(
        cluster_moref=pulumi.get(__response__, 'cluster_moref'),
        hardware_version=pulumi.get(__response__, 'hardware_version'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        vcenter_id=pulumi.get(__response__, 'vcenter_id')))
