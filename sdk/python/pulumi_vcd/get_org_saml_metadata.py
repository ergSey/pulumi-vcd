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
    'GetOrgSamlMetadataResult',
    'AwaitableGetOrgSamlMetadataResult',
    'get_org_saml_metadata',
    'get_org_saml_metadata_output',
]

@pulumi.output_type
class GetOrgSamlMetadataResult:
    """
    A collection of values returned by getOrgSamlMetadata.
    """
    def __init__(__self__, file_name=None, id=None, metadata_text=None, org_id=None):
        if file_name and not isinstance(file_name, str):
            raise TypeError("Expected argument 'file_name' to be a str")
        pulumi.set(__self__, "file_name", file_name)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if metadata_text and not isinstance(metadata_text, str):
            raise TypeError("Expected argument 'metadata_text' to be a str")
        pulumi.set(__self__, "metadata_text", metadata_text)
        if org_id and not isinstance(org_id, str):
            raise TypeError("Expected argument 'org_id' to be a str")
        pulumi.set(__self__, "org_id", org_id)

    @property
    @pulumi.getter(name="fileName")
    def file_name(self) -> Optional[str]:
        return pulumi.get(self, "file_name")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="metadataText")
    def metadata_text(self) -> str:
        """
        the text of the metadata for this organization.
        """
        return pulumi.get(self, "metadata_text")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> str:
        return pulumi.get(self, "org_id")


class AwaitableGetOrgSamlMetadataResult(GetOrgSamlMetadataResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOrgSamlMetadataResult(
            file_name=self.file_name,
            id=self.id,
            metadata_text=self.metadata_text,
            org_id=self.org_id)


def get_org_saml_metadata(file_name: Optional[str] = None,
                          org_id: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOrgSamlMetadataResult:
    """
    Supported in provider *v3.10+*.

    Provides a data source to read service provider SAML metadata for an organization.
    This service provider metadata is used to configure the identity provider.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vcd as vcd

    my_org = vcd.get_org(name="my-org")
    first = vcd.get_org_saml_metadata(org_id=my_org.id,
        file_name="vcd-metadata.txt")
    ```


    :param str file_name: name of the file where to store the metadata.
    :param str org_id: ID of the organization containing the SAML metadata
    """
    __args__ = dict()
    __args__['fileName'] = file_name
    __args__['orgId'] = org_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getOrgSamlMetadata:getOrgSamlMetadata', __args__, opts=opts, typ=GetOrgSamlMetadataResult).value

    return AwaitableGetOrgSamlMetadataResult(
        file_name=pulumi.get(__ret__, 'file_name'),
        id=pulumi.get(__ret__, 'id'),
        metadata_text=pulumi.get(__ret__, 'metadata_text'),
        org_id=pulumi.get(__ret__, 'org_id'))
def get_org_saml_metadata_output(file_name: Optional[pulumi.Input[Optional[str]]] = None,
                                 org_id: Optional[pulumi.Input[str]] = None,
                                 opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetOrgSamlMetadataResult]:
    """
    Supported in provider *v3.10+*.

    Provides a data source to read service provider SAML metadata for an organization.
    This service provider metadata is used to configure the identity provider.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vcd as vcd

    my_org = vcd.get_org(name="my-org")
    first = vcd.get_org_saml_metadata(org_id=my_org.id,
        file_name="vcd-metadata.txt")
    ```


    :param str file_name: name of the file where to store the metadata.
    :param str org_id: ID of the organization containing the SAML metadata
    """
    __args__ = dict()
    __args__['fileName'] = file_name
    __args__['orgId'] = org_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vcd:index/getOrgSamlMetadata:getOrgSamlMetadata', __args__, opts=opts, typ=GetOrgSamlMetadataResult)
    return __ret__.apply(lambda __response__: GetOrgSamlMetadataResult(
        file_name=pulumi.get(__response__, 'file_name'),
        id=pulumi.get(__response__, 'id'),
        metadata_text=pulumi.get(__response__, 'metadata_text'),
        org_id=pulumi.get(__response__, 'org_id')))
