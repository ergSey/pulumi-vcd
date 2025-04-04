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
    'GetMultisiteSiteAssociationResult',
    'AwaitableGetMultisiteSiteAssociationResult',
    'get_multisite_site_association',
    'get_multisite_site_association_output',
]

@pulumi.output_type
class GetMultisiteSiteAssociationResult:
    """
    A collection of values returned by getMultisiteSiteAssociation.
    """
    def __init__(__self__, associated_site_href=None, associated_site_id=None, associated_site_name=None, association_data_file=None, id=None, status=None):
        if associated_site_href and not isinstance(associated_site_href, str):
            raise TypeError("Expected argument 'associated_site_href' to be a str")
        pulumi.set(__self__, "associated_site_href", associated_site_href)
        if associated_site_id and not isinstance(associated_site_id, str):
            raise TypeError("Expected argument 'associated_site_id' to be a str")
        pulumi.set(__self__, "associated_site_id", associated_site_id)
        if associated_site_name and not isinstance(associated_site_name, str):
            raise TypeError("Expected argument 'associated_site_name' to be a str")
        pulumi.set(__self__, "associated_site_name", associated_site_name)
        if association_data_file and not isinstance(association_data_file, str):
            raise TypeError("Expected argument 'association_data_file' to be a str")
        pulumi.set(__self__, "association_data_file", association_data_file)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="associatedSiteHref")
    def associated_site_href(self) -> str:
        """
        The URL of the associated site.
        """
        return pulumi.get(self, "associated_site_href")

    @property
    @pulumi.getter(name="associatedSiteId")
    def associated_site_id(self) -> Optional[str]:
        return pulumi.get(self, "associated_site_id")

    @property
    @pulumi.getter(name="associatedSiteName")
    def associated_site_name(self) -> str:
        """
        The name of the associated site.
        """
        return pulumi.get(self, "associated_site_name")

    @property
    @pulumi.getter(name="associationDataFile")
    def association_data_file(self) -> Optional[str]:
        return pulumi.get(self, "association_data_file")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of the association (one of `ASYMMETRIC`, `ACTIVE`, `UNREACHABLE`, `ERROR`)
        """
        return pulumi.get(self, "status")


class AwaitableGetMultisiteSiteAssociationResult(GetMultisiteSiteAssociationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMultisiteSiteAssociationResult(
            associated_site_href=self.associated_site_href,
            associated_site_id=self.associated_site_id,
            associated_site_name=self.associated_site_name,
            association_data_file=self.association_data_file,
            id=self.id,
            status=self.status)


def get_multisite_site_association(associated_site_id: Optional[str] = None,
                                   association_data_file: Optional[str] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMultisiteSiteAssociationResult:
    """
    Provides a data source to read a VMware Cloud Director site association information.

    > Note: this data source requires System Administrator privileges

    Supported in provider *v3.13+*

    ## Example Usage

    ### 1

    Retrieving a site association using the associated site ID.

    ```python
    import pulumi
    import pulumi_vcd as vcd

    site1_site2 = vcd.get_multisite_site_association(associated_site_id="urn:vcloud:site:dca02216-fcf3-414a-be95-a3e26cf1296b")
    ```

    ### 2

    Retrieving a site association using the association data file.

    ```python
    import pulumi
    import pulumi_vcd as vcd

    site1_site2 = vcd.get_multisite_site_association(association_data_file="remote-site.xml")
    ```

    ## More information

    See [Site and Org association](https://www.terraform.io/providers/vmware/vcd/latest/docs/guides/site_org_association) for a broader description
    of association workflows.


    :param str associated_site_id: ID of the remote site associated with the current one. (Used in alternative to
           `associated_data_file`)
    :param str association_data_file: Name of the file containing the data used to associate this site to another one.
           (Used when `associated_site_id` is not known)
    """
    __args__ = dict()
    __args__['associatedSiteId'] = associated_site_id
    __args__['associationDataFile'] = association_data_file
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getMultisiteSiteAssociation:getMultisiteSiteAssociation', __args__, opts=opts, typ=GetMultisiteSiteAssociationResult).value

    return AwaitableGetMultisiteSiteAssociationResult(
        associated_site_href=pulumi.get(__ret__, 'associated_site_href'),
        associated_site_id=pulumi.get(__ret__, 'associated_site_id'),
        associated_site_name=pulumi.get(__ret__, 'associated_site_name'),
        association_data_file=pulumi.get(__ret__, 'association_data_file'),
        id=pulumi.get(__ret__, 'id'),
        status=pulumi.get(__ret__, 'status'))
def get_multisite_site_association_output(associated_site_id: Optional[pulumi.Input[Optional[str]]] = None,
                                          association_data_file: Optional[pulumi.Input[Optional[str]]] = None,
                                          opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetMultisiteSiteAssociationResult]:
    """
    Provides a data source to read a VMware Cloud Director site association information.

    > Note: this data source requires System Administrator privileges

    Supported in provider *v3.13+*

    ## Example Usage

    ### 1

    Retrieving a site association using the associated site ID.

    ```python
    import pulumi
    import pulumi_vcd as vcd

    site1_site2 = vcd.get_multisite_site_association(associated_site_id="urn:vcloud:site:dca02216-fcf3-414a-be95-a3e26cf1296b")
    ```

    ### 2

    Retrieving a site association using the association data file.

    ```python
    import pulumi
    import pulumi_vcd as vcd

    site1_site2 = vcd.get_multisite_site_association(association_data_file="remote-site.xml")
    ```

    ## More information

    See [Site and Org association](https://www.terraform.io/providers/vmware/vcd/latest/docs/guides/site_org_association) for a broader description
    of association workflows.


    :param str associated_site_id: ID of the remote site associated with the current one. (Used in alternative to
           `associated_data_file`)
    :param str association_data_file: Name of the file containing the data used to associate this site to another one.
           (Used when `associated_site_id` is not known)
    """
    __args__ = dict()
    __args__['associatedSiteId'] = associated_site_id
    __args__['associationDataFile'] = association_data_file
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vcd:index/getMultisiteSiteAssociation:getMultisiteSiteAssociation', __args__, opts=opts, typ=GetMultisiteSiteAssociationResult)
    return __ret__.apply(lambda __response__: GetMultisiteSiteAssociationResult(
        associated_site_href=pulumi.get(__response__, 'associated_site_href'),
        associated_site_id=pulumi.get(__response__, 'associated_site_id'),
        associated_site_name=pulumi.get(__response__, 'associated_site_name'),
        association_data_file=pulumi.get(__response__, 'association_data_file'),
        id=pulumi.get(__response__, 'id'),
        status=pulumi.get(__response__, 'status')))
