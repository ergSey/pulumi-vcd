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
    'GetNsxtManagerResult',
    'AwaitableGetNsxtManagerResult',
    'get_nsxt_manager',
    'get_nsxt_manager_output',
]

@pulumi.output_type
class GetNsxtManagerResult:
    """
    A collection of values returned by getNsxtManager.
    """
    def __init__(__self__, href=None, id=None, name=None):
        if href and not isinstance(href, str):
            raise TypeError("Expected argument 'href' to be a str")
        pulumi.set(__self__, "href", href)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def href(self) -> str:
        """
        Full URL of the manager
        """
        return pulumi.get(self, "href")

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


class AwaitableGetNsxtManagerResult(GetNsxtManagerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNsxtManagerResult(
            href=self.href,
            id=self.id,
            name=self.name)


def get_nsxt_manager(name: Optional[str] = None,
                     opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNsxtManagerResult:
    """
    Provides a data source for NSX-T manager.

    Supported in provider *v3.0+*

    > **Note:** This resource uses new VMware Cloud Director
    [OpenAPI](https://code.vmware.com/docs/11982/getting-started-with-vmware-cloud-director-openapi) and
    requires at least VCD *10.1.1+* and NSX-T *3.0+*.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vcd as vcd

    main = vcd.get_nsxt_manager(name="nsxt-manager-one")
    ```


    :param str name: NSX-T manager name
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getNsxtManager:getNsxtManager', __args__, opts=opts, typ=GetNsxtManagerResult).value

    return AwaitableGetNsxtManagerResult(
        href=pulumi.get(__ret__, 'href'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'))
def get_nsxt_manager_output(name: Optional[pulumi.Input[str]] = None,
                            opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetNsxtManagerResult]:
    """
    Provides a data source for NSX-T manager.

    Supported in provider *v3.0+*

    > **Note:** This resource uses new VMware Cloud Director
    [OpenAPI](https://code.vmware.com/docs/11982/getting-started-with-vmware-cloud-director-openapi) and
    requires at least VCD *10.1.1+* and NSX-T *3.0+*.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vcd as vcd

    main = vcd.get_nsxt_manager(name="nsxt-manager-one")
    ```


    :param str name: NSX-T manager name
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vcd:index/getNsxtManager:getNsxtManager', __args__, opts=opts, typ=GetNsxtManagerResult)
    return __ret__.apply(lambda __response__: GetNsxtManagerResult(
        href=pulumi.get(__response__, 'href'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name')))
