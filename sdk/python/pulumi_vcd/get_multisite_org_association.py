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
    'GetMultisiteOrgAssociationResult',
    'AwaitableGetMultisiteOrgAssociationResult',
    'get_multisite_org_association',
    'get_multisite_org_association_output',
]

@pulumi.output_type
class GetMultisiteOrgAssociationResult:
    """
    A collection of values returned by getMultisiteOrgAssociation.
    """
    def __init__(__self__, associated_org_id=None, associated_org_name=None, associated_site_id=None, association_data_file=None, id=None, org_id=None, status=None):
        if associated_org_id and not isinstance(associated_org_id, str):
            raise TypeError("Expected argument 'associated_org_id' to be a str")
        pulumi.set(__self__, "associated_org_id", associated_org_id)
        if associated_org_name and not isinstance(associated_org_name, str):
            raise TypeError("Expected argument 'associated_org_name' to be a str")
        pulumi.set(__self__, "associated_org_name", associated_org_name)
        if associated_site_id and not isinstance(associated_site_id, str):
            raise TypeError("Expected argument 'associated_site_id' to be a str")
        pulumi.set(__self__, "associated_site_id", associated_site_id)
        if association_data_file and not isinstance(association_data_file, str):
            raise TypeError("Expected argument 'association_data_file' to be a str")
        pulumi.set(__self__, "association_data_file", association_data_file)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if org_id and not isinstance(org_id, str):
            raise TypeError("Expected argument 'org_id' to be a str")
        pulumi.set(__self__, "org_id", org_id)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)

    @property
    @pulumi.getter(name="associatedOrgId")
    def associated_org_id(self) -> Optional[str]:
        return pulumi.get(self, "associated_org_id")

    @property
    @pulumi.getter(name="associatedOrgName")
    def associated_org_name(self) -> str:
        """
        The name of the associated Org.
        """
        return pulumi.get(self, "associated_org_name")

    @property
    @pulumi.getter(name="associatedSiteId")
    def associated_site_id(self) -> str:
        """
        The ID of the associated site.
        """
        return pulumi.get(self, "associated_site_id")

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
    @pulumi.getter(name="orgId")
    def org_id(self) -> str:
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter
    def status(self) -> str:
        """
        The status of the association (one of `ASYMMETRIC`, `ACTIVE`, `UNREACHABLE`, `ERROR`)
        """
        return pulumi.get(self, "status")


class AwaitableGetMultisiteOrgAssociationResult(GetMultisiteOrgAssociationResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMultisiteOrgAssociationResult(
            associated_org_id=self.associated_org_id,
            associated_org_name=self.associated_org_name,
            associated_site_id=self.associated_site_id,
            association_data_file=self.association_data_file,
            id=self.id,
            org_id=self.org_id,
            status=self.status)


def get_multisite_org_association(associated_org_id: Optional[str] = None,
                                  association_data_file: Optional[str] = None,
                                  org_id: Optional[str] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMultisiteOrgAssociationResult:
    """
    Provides a data source to read a VMware Cloud Director Org association information.

    Supported in provider *v3.13+*

    ## Example Usage

    ### 1

    Retrieving an Org association using the associated Org ID.

    ```python
    import pulumi
    import pulumi_vcd as vcd

    my_org = vcd.get_org(name="my-org")
    org1_org2 = vcd.get_multisite_org_association(org_id=my_org.id,
        associated_org_id="urn:vcloud:org:3901d87d-1596-4a5a-a74b-57a7313737cf")
    ```

    ### 2

    Retrieving an Org association using the association data file.

    ```python
    import pulumi
    import pulumi_vcd as vcd

    my_org = vcd.get_org(name="my-org")
    org1_org2 = vcd.get_multisite_org_association(org_id=my_org.id,
        association_data_file="remote-org.xml")
    ```

    ## More information

    See [Site and Org association](https://www.terraform.io/providers/vmware/vcd/latest/docs/guides/site_org_association) for a broader description
    of association workflows.


    :param str associated_org_id: ID of the remote organization associated with the current one. (Used in alternative to
           `associated_data_file`)
    :param str association_data_file: Name of the file containing the data used to associate this Org to another one.
           (Used when `associated_org_id` is not known)
    :param str org_id: The ID of the organization for which we need to collect the data.
    """
    __args__ = dict()
    __args__['associatedOrgId'] = associated_org_id
    __args__['associationDataFile'] = association_data_file
    __args__['orgId'] = org_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getMultisiteOrgAssociation:getMultisiteOrgAssociation', __args__, opts=opts, typ=GetMultisiteOrgAssociationResult).value

    return AwaitableGetMultisiteOrgAssociationResult(
        associated_org_id=pulumi.get(__ret__, 'associated_org_id'),
        associated_org_name=pulumi.get(__ret__, 'associated_org_name'),
        associated_site_id=pulumi.get(__ret__, 'associated_site_id'),
        association_data_file=pulumi.get(__ret__, 'association_data_file'),
        id=pulumi.get(__ret__, 'id'),
        org_id=pulumi.get(__ret__, 'org_id'),
        status=pulumi.get(__ret__, 'status'))
def get_multisite_org_association_output(associated_org_id: Optional[pulumi.Input[Optional[str]]] = None,
                                         association_data_file: Optional[pulumi.Input[Optional[str]]] = None,
                                         org_id: Optional[pulumi.Input[str]] = None,
                                         opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetMultisiteOrgAssociationResult]:
    """
    Provides a data source to read a VMware Cloud Director Org association information.

    Supported in provider *v3.13+*

    ## Example Usage

    ### 1

    Retrieving an Org association using the associated Org ID.

    ```python
    import pulumi
    import pulumi_vcd as vcd

    my_org = vcd.get_org(name="my-org")
    org1_org2 = vcd.get_multisite_org_association(org_id=my_org.id,
        associated_org_id="urn:vcloud:org:3901d87d-1596-4a5a-a74b-57a7313737cf")
    ```

    ### 2

    Retrieving an Org association using the association data file.

    ```python
    import pulumi
    import pulumi_vcd as vcd

    my_org = vcd.get_org(name="my-org")
    org1_org2 = vcd.get_multisite_org_association(org_id=my_org.id,
        association_data_file="remote-org.xml")
    ```

    ## More information

    See [Site and Org association](https://www.terraform.io/providers/vmware/vcd/latest/docs/guides/site_org_association) for a broader description
    of association workflows.


    :param str associated_org_id: ID of the remote organization associated with the current one. (Used in alternative to
           `associated_data_file`)
    :param str association_data_file: Name of the file containing the data used to associate this Org to another one.
           (Used when `associated_org_id` is not known)
    :param str org_id: The ID of the organization for which we need to collect the data.
    """
    __args__ = dict()
    __args__['associatedOrgId'] = associated_org_id
    __args__['associationDataFile'] = association_data_file
    __args__['orgId'] = org_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vcd:index/getMultisiteOrgAssociation:getMultisiteOrgAssociation', __args__, opts=opts, typ=GetMultisiteOrgAssociationResult)
    return __ret__.apply(lambda __response__: GetMultisiteOrgAssociationResult(
        associated_org_id=pulumi.get(__response__, 'associated_org_id'),
        associated_org_name=pulumi.get(__response__, 'associated_org_name'),
        associated_site_id=pulumi.get(__response__, 'associated_site_id'),
        association_data_file=pulumi.get(__response__, 'association_data_file'),
        id=pulumi.get(__response__, 'id'),
        org_id=pulumi.get(__response__, 'org_id'),
        status=pulumi.get(__response__, 'status')))
