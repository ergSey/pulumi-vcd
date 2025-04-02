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
from . import outputs

__all__ = [
    'GetRightResult',
    'AwaitableGetRightResult',
    'get_right',
    'get_right_output',
]

@pulumi.output_type
class GetRightResult:
    """
    A collection of values returned by getRight.
    """
    def __init__(__self__, bundle_key=None, category_id=None, description=None, id=None, implied_rights=None, name=None, right_type=None):
        if bundle_key and not isinstance(bundle_key, str):
            raise TypeError("Expected argument 'bundle_key' to be a str")
        pulumi.set(__self__, "bundle_key", bundle_key)
        if category_id and not isinstance(category_id, str):
            raise TypeError("Expected argument 'category_id' to be a str")
        pulumi.set(__self__, "category_id", category_id)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if implied_rights and not isinstance(implied_rights, list):
            raise TypeError("Expected argument 'implied_rights' to be a list")
        pulumi.set(__self__, "implied_rights", implied_rights)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if right_type and not isinstance(right_type, str):
            raise TypeError("Expected argument 'right_type' to be a str")
        pulumi.set(__self__, "right_type", right_type)

    @property
    @pulumi.getter(name="bundleKey")
    def bundle_key(self) -> str:
        """
        Key used for internationalization
        * `right type` - Type of the right (VIEW or MODIFY)
        """
        return pulumi.get(self, "bundle_key")

    @property
    @pulumi.getter(name="categoryId")
    def category_id(self) -> str:
        """
        The ID of the category for this right
        """
        return pulumi.get(self, "category_id")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        A description of the right
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
    @pulumi.getter(name="impliedRights")
    def implied_rights(self) -> Sequence['outputs.GetRightImpliedRightResult']:
        """
        List of rights that are implied with this one
        """
        return pulumi.get(self, "implied_rights")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="rightType")
    def right_type(self) -> str:
        return pulumi.get(self, "right_type")


class AwaitableGetRightResult(GetRightResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRightResult(
            bundle_key=self.bundle_key,
            category_id=self.category_id,
            description=self.description,
            id=self.id,
            implied_rights=self.implied_rights,
            name=self.name,
            right_type=self.right_type)


def get_right(name: Optional[str] = None,
              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRightResult:
    """
    Provides a data source for available rights.

    Supported in provider *v3.3+*

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vcd as vcd

    some_right = vcd.get_right(name="Catalog: Add vApp from My Cloud")
    pulumi.export("some-right", some_right)
    ```

    ## More information

    See [Roles management](https://www.terraform.io/providers/vmware/vcd/latest/docs/guides/roles_management) for a broader description of how roles and
    rights work together.


    :param str name: The name of the right.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getRight:getRight', __args__, opts=opts, typ=GetRightResult).value

    return AwaitableGetRightResult(
        bundle_key=pulumi.get(__ret__, 'bundle_key'),
        category_id=pulumi.get(__ret__, 'category_id'),
        description=pulumi.get(__ret__, 'description'),
        id=pulumi.get(__ret__, 'id'),
        implied_rights=pulumi.get(__ret__, 'implied_rights'),
        name=pulumi.get(__ret__, 'name'),
        right_type=pulumi.get(__ret__, 'right_type'))
def get_right_output(name: Optional[pulumi.Input[str]] = None,
                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetRightResult]:
    """
    Provides a data source for available rights.

    Supported in provider *v3.3+*

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vcd as vcd

    some_right = vcd.get_right(name="Catalog: Add vApp from My Cloud")
    pulumi.export("some-right", some_right)
    ```

    ## More information

    See [Roles management](https://www.terraform.io/providers/vmware/vcd/latest/docs/guides/roles_management) for a broader description of how roles and
    rights work together.


    :param str name: The name of the right.
    """
    __args__ = dict()
    __args__['name'] = name
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vcd:index/getRight:getRight', __args__, opts=opts, typ=GetRightResult)
    return __ret__.apply(lambda __response__: GetRightResult(
        bundle_key=pulumi.get(__response__, 'bundle_key'),
        category_id=pulumi.get(__response__, 'category_id'),
        description=pulumi.get(__response__, 'description'),
        id=pulumi.get(__response__, 'id'),
        implied_rights=pulumi.get(__response__, 'implied_rights'),
        name=pulumi.get(__response__, 'name'),
        right_type=pulumi.get(__response__, 'right_type')))
