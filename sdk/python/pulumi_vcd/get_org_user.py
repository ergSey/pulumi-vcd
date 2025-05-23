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
    'GetOrgUserResult',
    'AwaitableGetOrgUserResult',
    'get_org_user',
    'get_org_user_output',
]

@pulumi.output_type
class GetOrgUserResult:
    """
    A collection of values returned by getOrgUser.
    """
    def __init__(__self__, deployed_vm_quota=None, description=None, email_address=None, enabled=None, full_name=None, group_names=None, id=None, instant_messaging=None, is_external=None, is_group_role=None, is_locked=None, name=None, org=None, provider_type=None, role=None, stored_vm_quota=None, telephone=None, user_id=None):
        if deployed_vm_quota and not isinstance(deployed_vm_quota, int):
            raise TypeError("Expected argument 'deployed_vm_quota' to be a int")
        pulumi.set(__self__, "deployed_vm_quota", deployed_vm_quota)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if email_address and not isinstance(email_address, str):
            raise TypeError("Expected argument 'email_address' to be a str")
        pulumi.set(__self__, "email_address", email_address)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if full_name and not isinstance(full_name, str):
            raise TypeError("Expected argument 'full_name' to be a str")
        pulumi.set(__self__, "full_name", full_name)
        if group_names and not isinstance(group_names, list):
            raise TypeError("Expected argument 'group_names' to be a list")
        pulumi.set(__self__, "group_names", group_names)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if instant_messaging and not isinstance(instant_messaging, str):
            raise TypeError("Expected argument 'instant_messaging' to be a str")
        pulumi.set(__self__, "instant_messaging", instant_messaging)
        if is_external and not isinstance(is_external, bool):
            raise TypeError("Expected argument 'is_external' to be a bool")
        pulumi.set(__self__, "is_external", is_external)
        if is_group_role and not isinstance(is_group_role, bool):
            raise TypeError("Expected argument 'is_group_role' to be a bool")
        pulumi.set(__self__, "is_group_role", is_group_role)
        if is_locked and not isinstance(is_locked, bool):
            raise TypeError("Expected argument 'is_locked' to be a bool")
        pulumi.set(__self__, "is_locked", is_locked)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if org and not isinstance(org, str):
            raise TypeError("Expected argument 'org' to be a str")
        pulumi.set(__self__, "org", org)
        if provider_type and not isinstance(provider_type, str):
            raise TypeError("Expected argument 'provider_type' to be a str")
        pulumi.set(__self__, "provider_type", provider_type)
        if role and not isinstance(role, str):
            raise TypeError("Expected argument 'role' to be a str")
        pulumi.set(__self__, "role", role)
        if stored_vm_quota and not isinstance(stored_vm_quota, int):
            raise TypeError("Expected argument 'stored_vm_quota' to be a int")
        pulumi.set(__self__, "stored_vm_quota", stored_vm_quota)
        if telephone and not isinstance(telephone, str):
            raise TypeError("Expected argument 'telephone' to be a str")
        pulumi.set(__self__, "telephone", telephone)
        if user_id and not isinstance(user_id, str):
            raise TypeError("Expected argument 'user_id' to be a str")
        pulumi.set(__self__, "user_id", user_id)

    @property
    @pulumi.getter(name="deployedVmQuota")
    def deployed_vm_quota(self) -> int:
        """
        Quota of vApps that this user can deploy. A value of 0 specifies an unlimited quota.
        """
        return pulumi.get(self, "deployed_vm_quota")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        An optional description of the user.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="emailAddress")
    def email_address(self) -> str:
        """
        The Org User email address.
        """
        return pulumi.get(self, "email_address")

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        """
        True if the user is enabled and can log in.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="fullName")
    def full_name(self) -> str:
        """
        The full name of the user.
        """
        return pulumi.get(self, "full_name")

    @property
    @pulumi.getter(name="groupNames")
    def group_names(self) -> Sequence[str]:
        """
        The set of group names to which this user belongs. It's only populated if the users
        are created after the group (with this user having a `depends_on` of the given group).
        """
        return pulumi.get(self, "group_names")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="instantMessaging")
    def instant_messaging(self) -> str:
        """
        The Org User instant messaging.
        """
        return pulumi.get(self, "instant_messaging")

    @property
    @pulumi.getter(name="isExternal")
    def is_external(self) -> bool:
        """
        If the user account was imported from an external resource, like an LDAP.
        """
        return pulumi.get(self, "is_external")

    @property
    @pulumi.getter(name="isGroupRole")
    def is_group_role(self) -> bool:
        """
        True if this user has a group role.
        """
        return pulumi.get(self, "is_group_role")

    @property
    @pulumi.getter(name="isLocked")
    def is_locked(self) -> bool:
        """
        If the user account has been locked due to too many invalid login attempts, the value will be true.
        """
        return pulumi.get(self, "is_locked")

    @property
    @pulumi.getter
    def name(self) -> Optional[str]:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def org(self) -> Optional[str]:
        return pulumi.get(self, "org")

    @property
    @pulumi.getter(name="providerType")
    def provider_type(self) -> str:
        """
        Identity provider type for this user.
        """
        return pulumi.get(self, "provider_type")

    @property
    @pulumi.getter
    def role(self) -> str:
        """
        The role of the user.
        """
        return pulumi.get(self, "role")

    @property
    @pulumi.getter(name="storedVmQuota")
    def stored_vm_quota(self) -> int:
        """
        Quota of vApps that this user can store. A value of 0 specifies an unlimited quota.
        """
        return pulumi.get(self, "stored_vm_quota")

    @property
    @pulumi.getter
    def telephone(self) -> str:
        """
        The Org User telephone number.
        """
        return pulumi.get(self, "telephone")

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> Optional[str]:
        return pulumi.get(self, "user_id")


class AwaitableGetOrgUserResult(GetOrgUserResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOrgUserResult(
            deployed_vm_quota=self.deployed_vm_quota,
            description=self.description,
            email_address=self.email_address,
            enabled=self.enabled,
            full_name=self.full_name,
            group_names=self.group_names,
            id=self.id,
            instant_messaging=self.instant_messaging,
            is_external=self.is_external,
            is_group_role=self.is_group_role,
            is_locked=self.is_locked,
            name=self.name,
            org=self.org,
            provider_type=self.provider_type,
            role=self.role,
            stored_vm_quota=self.stored_vm_quota,
            telephone=self.telephone,
            user_id=self.user_id)


def get_org_user(name: Optional[str] = None,
                 org: Optional[str] = None,
                 user_id: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOrgUserResult:
    """
    Provides a VMware Cloud Director Org User data source. This can be used to read organization users, including org administrators.

    Supported in provider *v3.0+*

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vcd as vcd

    my_org_admin = vcd.get_org_user(org="my-org",
        name="my-org-admin")
    my_vapp_creator = vcd.get_org_user(org="my-org",
        user_id="urn:vcloud:user:c311eb35-6984-4d26-3ee9-0000deadbeef")
    pulumi.export("adminUser", my_org_admin)
    pulumi.export("vappCreatorUser", my_vapp_creator)
    ```


    :param str name: The name of the user. Required if `user_id` is not set.
    :param str org: The name of organization to which the user belongs. Optional if defined at provider level.
    :param str user_id: The ID of the user. Required if `name` is not set.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['org'] = org
    __args__['userId'] = user_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getOrgUser:getOrgUser', __args__, opts=opts, typ=GetOrgUserResult).value

    return AwaitableGetOrgUserResult(
        deployed_vm_quota=pulumi.get(__ret__, 'deployed_vm_quota'),
        description=pulumi.get(__ret__, 'description'),
        email_address=pulumi.get(__ret__, 'email_address'),
        enabled=pulumi.get(__ret__, 'enabled'),
        full_name=pulumi.get(__ret__, 'full_name'),
        group_names=pulumi.get(__ret__, 'group_names'),
        id=pulumi.get(__ret__, 'id'),
        instant_messaging=pulumi.get(__ret__, 'instant_messaging'),
        is_external=pulumi.get(__ret__, 'is_external'),
        is_group_role=pulumi.get(__ret__, 'is_group_role'),
        is_locked=pulumi.get(__ret__, 'is_locked'),
        name=pulumi.get(__ret__, 'name'),
        org=pulumi.get(__ret__, 'org'),
        provider_type=pulumi.get(__ret__, 'provider_type'),
        role=pulumi.get(__ret__, 'role'),
        stored_vm_quota=pulumi.get(__ret__, 'stored_vm_quota'),
        telephone=pulumi.get(__ret__, 'telephone'),
        user_id=pulumi.get(__ret__, 'user_id'))
def get_org_user_output(name: Optional[pulumi.Input[Optional[str]]] = None,
                        org: Optional[pulumi.Input[Optional[str]]] = None,
                        user_id: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetOrgUserResult]:
    """
    Provides a VMware Cloud Director Org User data source. This can be used to read organization users, including org administrators.

    Supported in provider *v3.0+*

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vcd as vcd

    my_org_admin = vcd.get_org_user(org="my-org",
        name="my-org-admin")
    my_vapp_creator = vcd.get_org_user(org="my-org",
        user_id="urn:vcloud:user:c311eb35-6984-4d26-3ee9-0000deadbeef")
    pulumi.export("adminUser", my_org_admin)
    pulumi.export("vappCreatorUser", my_vapp_creator)
    ```


    :param str name: The name of the user. Required if `user_id` is not set.
    :param str org: The name of organization to which the user belongs. Optional if defined at provider level.
    :param str user_id: The ID of the user. Required if `name` is not set.
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['org'] = org
    __args__['userId'] = user_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vcd:index/getOrgUser:getOrgUser', __args__, opts=opts, typ=GetOrgUserResult)
    return __ret__.apply(lambda __response__: GetOrgUserResult(
        deployed_vm_quota=pulumi.get(__response__, 'deployed_vm_quota'),
        description=pulumi.get(__response__, 'description'),
        email_address=pulumi.get(__response__, 'email_address'),
        enabled=pulumi.get(__response__, 'enabled'),
        full_name=pulumi.get(__response__, 'full_name'),
        group_names=pulumi.get(__response__, 'group_names'),
        id=pulumi.get(__response__, 'id'),
        instant_messaging=pulumi.get(__response__, 'instant_messaging'),
        is_external=pulumi.get(__response__, 'is_external'),
        is_group_role=pulumi.get(__response__, 'is_group_role'),
        is_locked=pulumi.get(__response__, 'is_locked'),
        name=pulumi.get(__response__, 'name'),
        org=pulumi.get(__response__, 'org'),
        provider_type=pulumi.get(__response__, 'provider_type'),
        role=pulumi.get(__response__, 'role'),
        stored_vm_quota=pulumi.get(__response__, 'stored_vm_quota'),
        telephone=pulumi.get(__response__, 'telephone'),
        user_id=pulumi.get(__response__, 'user_id')))
