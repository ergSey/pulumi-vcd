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
    'GetMultisiteSiteResult',
    'AwaitableGetMultisiteSiteResult',
    'get_multisite_site',
    'get_multisite_site_output',
]

@pulumi.output_type
class GetMultisiteSiteResult:
    """
    A collection of values returned by getMultisiteSite.
    """
    def __init__(__self__, associations=None, description=None, id=None, name=None, number_of_associations=None):
        if associations and not isinstance(associations, list):
            raise TypeError("Expected argument 'associations' to be a list")
        pulumi.set(__self__, "associations", associations)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if number_of_associations and not isinstance(number_of_associations, int):
            raise TypeError("Expected argument 'number_of_associations' to be a int")
        pulumi.set(__self__, "number_of_associations", number_of_associations)

    @property
    @pulumi.getter
    def associations(self) -> Sequence[str]:
        """
        An alphabetically sorted list of current associations.
        """
        return pulumi.get(self, "associations")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        An optional description of the site.
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
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the site, which usually corresponds to its host name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="numberOfAssociations")
    def number_of_associations(self) -> int:
        """
        The number of current associations with other sites.
        """
        return pulumi.get(self, "number_of_associations")


class AwaitableGetMultisiteSiteResult(GetMultisiteSiteResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetMultisiteSiteResult(
            associations=self.associations,
            description=self.description,
            id=self.id,
            name=self.name,
            number_of_associations=self.number_of_associations)


def get_multisite_site(opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetMultisiteSiteResult:
    """
    Provides a data source to read a VMware Cloud Director Site in the context of multi-site operatioos

    Supported in provider *v3.13+*

    > Note: this data source requires System Administrator privileges

    ## Example Usage

    Note: there is only one site available for each VCD. No ID or name is necessary to identify it.

    ```python
    import pulumi
    import pulumi_vcd as vcd

    current_site = vcd.get_multisite_site()
    ```

    ## More information

    See [Site and Org association](https://www.terraform.io/providers/vmware/vcd/latest/docs/guides/site_org_association) for a broader description
    of association workflows.
    """
    __args__ = dict()
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getMultisiteSite:getMultisiteSite', __args__, opts=opts, typ=GetMultisiteSiteResult).value

    return AwaitableGetMultisiteSiteResult(
        associations=pulumi.get(__ret__, 'associations'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        number_of_associations=pulumi.get(__ret__, 'number_of_associations'))
def get_multisite_site_output(opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetMultisiteSiteResult]:
    """
    Provides a data source to read a VMware Cloud Director Site in the context of multi-site operatioos

    Supported in provider *v3.13+*

    > Note: this data source requires System Administrator privileges

    ## Example Usage

    Note: there is only one site available for each VCD. No ID or name is necessary to identify it.

    ```python
    import pulumi
    import pulumi_vcd as vcd

    current_site = vcd.get_multisite_site()
    ```

    ## More information

    See [Site and Org association](https://www.terraform.io/providers/vmware/vcd/latest/docs/guides/site_org_association) for a broader description
    of association workflows.
    """
    __args__ = dict()
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vcd:index/getMultisiteSite:getMultisiteSite', __args__, opts=opts, typ=GetMultisiteSiteResult)
    return __ret__.apply(lambda __response__: GetMultisiteSiteResult(
        associations=pulumi.get(__response__, 'associations'),
        description=pulumi.get(__response__, 'description'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        number_of_associations=pulumi.get(__response__, 'number_of_associations')))
