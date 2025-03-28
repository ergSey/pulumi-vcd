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
    'GetMultisiteSiteDataResult',
    'AwaitableGetMultisiteSiteDataResult',
    'get_multisite_site_data',
    'get_multisite_site_data_output',
]

@pulumi.output_type
class GetMultisiteSiteDataResult:
    """
    A collection of values returned by getMultisiteSiteData.
    """
    def __init__(__self__, association_data=None, download_to_file=None, id=None):
        if association_data and not isinstance(association_data, str):
            raise TypeError("Expected argument 'association_data' to be a str")
        pulumi.set(__self__, "association_data", association_data)
        if download_to_file and not isinstance(download_to_file, str):
            raise TypeError("Expected argument 'download_to_file' to be a str")
        pulumi.set(__self__, "download_to_file", download_to_file)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter(name="associationData")
    def association_data(self) -> str:
        """
        The data needed to associate this site to another one. Contains the same data that would be saved into
        the file defined in `download_to_file`.
        """
        return pulumi.get(self, "association_data")

    @property
    @pulumi.getter(name="downloadToFile")
    def download_to_file(self) -> Optional[str]:
        return pulumi.get(self, "download_to_file")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")


class AwaitableGetMultisiteSiteDataResult(GetMultisiteSiteDataResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMultisiteSiteDataResult(
            association_data=self.association_data,
            download_to_file=self.download_to_file,
            id=self.id)


def get_multisite_site_data(download_to_file: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMultisiteSiteDataResult:
    """
    Provides a data source to read a VMware Cloud Director Site association data to be used for association with another site.

    Supported in provider *v3.13+*

    ## Example Usage

    Note: there is only one site available for each VCD. No ID or name is necessary to identify it.

    > Note: this data source requires System Administrator privileges

    ```python
    import pulumi
    import pulumi_vcd as vcd

    current_site = vcd.get_multisite_site_data(download_to_file="filename.xml")
    ```


    :param str download_to_file: Name of the file that will contain the data needed to associate this site to a remote one.
           Contains the same data returned in `association_data`.
    """
    __args__ = dict()
    __args__['downloadToFile'] = download_to_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getMultisiteSiteData:getMultisiteSiteData', __args__, opts=opts, typ=GetMultisiteSiteDataResult).value

    return AwaitableGetMultisiteSiteDataResult(
        association_data=pulumi.get(__ret__, 'association_data'),
        download_to_file=pulumi.get(__ret__, 'download_to_file'),
        id=pulumi.get(__ret__, 'id'))
def get_multisite_site_data_output(download_to_file: Optional[pulumi.Input[Optional[str]]] = None,
                                   opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetMultisiteSiteDataResult]:
    """
    Provides a data source to read a VMware Cloud Director Site association data to be used for association with another site.

    Supported in provider *v3.13+*

    ## Example Usage

    Note: there is only one site available for each VCD. No ID or name is necessary to identify it.

    > Note: this data source requires System Administrator privileges

    ```python
    import pulumi
    import pulumi_vcd as vcd

    current_site = vcd.get_multisite_site_data(download_to_file="filename.xml")
    ```


    :param str download_to_file: Name of the file that will contain the data needed to associate this site to a remote one.
           Contains the same data returned in `association_data`.
    """
    __args__ = dict()
    __args__['downloadToFile'] = download_to_file
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vcd:index/getMultisiteSiteData:getMultisiteSiteData', __args__, opts=opts, typ=GetMultisiteSiteDataResult)
    return __ret__.apply(lambda __response__: GetMultisiteSiteDataResult(
        association_data=pulumi.get(__response__, 'association_data'),
        download_to_file=pulumi.get(__response__, 'download_to_file'),
        id=pulumi.get(__response__, 'id')))
