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
    'GetNsxvApplicationResult',
    'AwaitableGetNsxvApplicationResult',
    'get_nsxv_application',
    'get_nsxv_application_output',
]

@pulumi.output_type
class GetNsxvApplicationResult:
    """
    A collection of values returned by getNsxvApplication.
    """
    def __init__(__self__, app_guid=None, id=None, name=None, ports=None, protocol=None, source_port=None, vdc_id=None):
        if app_guid and not isinstance(app_guid, str):
            raise TypeError("Expected argument 'app_guid' to be a str")
        pulumi.set(__self__, "app_guid", app_guid)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if ports and not isinstance(ports, str):
            raise TypeError("Expected argument 'ports' to be a str")
        pulumi.set(__self__, "ports", ports)
        if protocol and not isinstance(protocol, str):
            raise TypeError("Expected argument 'protocol' to be a str")
        pulumi.set(__self__, "protocol", protocol)
        if source_port and not isinstance(source_port, str):
            raise TypeError("Expected argument 'source_port' to be a str")
        pulumi.set(__self__, "source_port", source_port)
        if vdc_id and not isinstance(vdc_id, str):
            raise TypeError("Expected argument 'vdc_id' to be a str")
        pulumi.set(__self__, "vdc_id", vdc_id)

    @property
    @pulumi.getter(name="appGuid")
    def app_guid(self) -> str:
        """
        The application identifier, when available
        """
        return pulumi.get(self, "app_guid")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The identifier of the application
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def ports(self) -> str:
        """
        The ports used by the application. Could be a number, a list of numbers, or a range
        """
        return pulumi.get(self, "ports")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        """
        The protocol used by the application
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="sourcePort")
    def source_port(self) -> str:
        """
        The source port used by this application. Not all applications provide one
        """
        return pulumi.get(self, "source_port")

    @property
    @pulumi.getter(name="vdcId")
    def vdc_id(self) -> str:
        return pulumi.get(self, "vdc_id")


class AwaitableGetNsxvApplicationResult(GetNsxvApplicationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNsxvApplicationResult(
            app_guid=self.app_guid,
            id=self.id,
            name=self.name,
            ports=self.ports,
            protocol=self.protocol,
            source_port=self.source_port,
            vdc_id=self.vdc_id)


def get_nsxv_application(name: Optional[str] = None,
                         vdc_id: Optional[str] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNsxvApplicationResult:
    """
    Provides a VMware Cloud Director data source for reading NSX-V distributed firewall applications.

    Supported in provider *v3.9+*


    :param str name: The name of the application
    :param str vdc_id: The ID of VDC to use
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdcId'] = vdc_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getNsxvApplication:getNsxvApplication', __args__, opts=opts, typ=GetNsxvApplicationResult).value

    return AwaitableGetNsxvApplicationResult(
        app_guid=pulumi.get(__ret__, 'app_guid'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        ports=pulumi.get(__ret__, 'ports'),
        protocol=pulumi.get(__ret__, 'protocol'),
        source_port=pulumi.get(__ret__, 'source_port'),
        vdc_id=pulumi.get(__ret__, 'vdc_id'))
def get_nsxv_application_output(name: Optional[pulumi.Input[str]] = None,
                                vdc_id: Optional[pulumi.Input[str]] = None,
                                opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetNsxvApplicationResult]:
    """
    Provides a VMware Cloud Director data source for reading NSX-V distributed firewall applications.

    Supported in provider *v3.9+*


    :param str name: The name of the application
    :param str vdc_id: The ID of VDC to use
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['vdcId'] = vdc_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vcd:index/getNsxvApplication:getNsxvApplication', __args__, opts=opts, typ=GetNsxvApplicationResult)
    return __ret__.apply(lambda __response__: GetNsxvApplicationResult(
        app_guid=pulumi.get(__response__, 'app_guid'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        ports=pulumi.get(__response__, 'ports'),
        protocol=pulumi.get(__response__, 'protocol'),
        source_port=pulumi.get(__response__, 'source_port'),
        vdc_id=pulumi.get(__response__, 'vdc_id')))
