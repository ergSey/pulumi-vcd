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
    'GetRoleResult',
    'AwaitableGetRoleResult',
    'get_role',
    'get_role_output',
]

@pulumi.output_type
class GetRoleResult:
    """
    A collection of values returned by getRole.
    """
    def __init__(__self__, bundle_key=None, description=None, id=None, name=None, org=None, read_only=None, rights=None):
        if bundle_key and not isinstance(bundle_key, str):
            raise TypeError("Expected argument 'bundle_key' to be a str")
        pulumi.set(__self__, "bundle_key", bundle_key)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if org and not isinstance(org, str):
            raise TypeError("Expected argument 'org' to be a str")
        pulumi.set(__self__, "org", org)
        if read_only and not isinstance(read_only, bool):
            raise TypeError("Expected argument 'read_only' to be a bool")
        pulumi.set(__self__, "read_only", read_only)
        if rights and not isinstance(rights, list):
            raise TypeError("Expected argument 'rights' to be a list")
        pulumi.set(__self__, "rights", rights)

    @property
    @pulumi.getter(name="bundleKey")
    def bundle_key(self) -> str:
        """
        Key used for internationalization.
        """
        return pulumi.get(self, "bundle_key")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        A description of the role
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
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def org(self) -> Optional[str]:
        return pulumi.get(self, "org")

    @property
    @pulumi.getter(name="readOnly")
    def read_only(self) -> bool:
        """
        Whether this role is read-only
        """
        return pulumi.get(self, "read_only")

    @property
    @pulumi.getter
    def rights(self) -> Sequence[str]:
        """
        Set of rights assigned to this role
        """
        return pulumi.get(self, "rights")


class AwaitableGetRoleResult(GetRoleResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRoleResult(
            bundle_key=self.bundle_key,
            description=self.description,
            id=self.id,
            name=self.name,
            org=self.org,
            read_only=self.read_only,
            rights=self.rights)


def get_role(name: Optional[str] = None,
             org: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRoleResult:
    """
    Provides a VMware Cloud Director role data source. This can be used to read roles.

    Supported in provider *v3.3+*

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vcd as vcd

    vapp_author = vcd.get_role(org="my-org",
        name="vApp Author")
    ```

    ## More information

    See [Roles management](https://www.terraform.io/providers/vmware/vcd/latest/docs/guides/roles_management) for a broader description of how roles and
    rights work together.


    :param str name: The name of the role.
    :param str org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['org'] = org
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getRole:getRole', __args__, opts=opts, typ=GetRoleResult).value

    return AwaitableGetRoleResult(
        bundle_key=pulumi.get(__ret__, 'bundle_key'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        org=pulumi.get(__ret__, 'org'),
        read_only=pulumi.get(__ret__, 'read_only'),
        rights=pulumi.get(__ret__, 'rights'))
def get_role_output(name: Optional[pulumi.Input[str]] = None,
                    org: Optional[pulumi.Input[Optional[str]]] = None,
                    opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetRoleResult]:
    """
    Provides a VMware Cloud Director role data source. This can be used to read roles.

    Supported in provider *v3.3+*

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vcd as vcd

    vapp_author = vcd.get_role(org="my-org",
        name="vApp Author")
    ```

    ## More information

    See [Roles management](https://www.terraform.io/providers/vmware/vcd/latest/docs/guides/roles_management) for a broader description of how roles and
    rights work together.


    :param str name: The name of the role.
    :param str org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['org'] = org
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vcd:index/getRole:getRole', __args__, opts=opts, typ=GetRoleResult)
    return __ret__.apply(lambda __response__: GetRoleResult(
        bundle_key=pulumi.get(__response__, 'bundle_key'),
        description=pulumi.get(__response__, 'description'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        org=pulumi.get(__response__, 'org'),
        read_only=pulumi.get(__response__, 'read_only'),
        rights=pulumi.get(__response__, 'rights')))
