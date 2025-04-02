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
    'GetResourceSchemaResult',
    'AwaitableGetResourceSchemaResult',
    'get_resource_schema',
    'get_resource_schema_output',
]

@pulumi.output_type
class GetResourceSchemaResult:
    """
    A collection of values returned by getResourceSchema.
    """
    def __init__(__self__, attributes=None, block_attributes=None, id=None, name=None, resource_type=None):
        if attributes and not isinstance(attributes, list):
            raise TypeError("Expected argument 'attributes' to be a list")
        pulumi.set(__self__, "attributes", attributes)
        if block_attributes and not isinstance(block_attributes, list):
            raise TypeError("Expected argument 'block_attributes' to be a list")
        pulumi.set(__self__, "block_attributes", block_attributes)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if resource_type and not isinstance(resource_type, str):
            raise TypeError("Expected argument 'resource_type' to be a str")
        pulumi.set(__self__, "resource_type", resource_type)

    @property
    @pulumi.getter
    def attributes(self) -> Sequence['outputs.GetResourceSchemaAttributeResult']:
        """
        (Computed) Same composition of the simple `attributes` above.
        """
        return pulumi.get(self, "attributes")

    @property
    @pulumi.getter(name="blockAttributes")
    def block_attributes(self) -> Sequence['outputs.GetResourceSchemaBlockAttributeResult']:
        """
        (Computed) The list of compound attributes
        Each bock attribute is made of:
        """
        return pulumi.get(self, "block_attributes")

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
        the attribute name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceType")
    def resource_type(self) -> str:
        return pulumi.get(self, "resource_type")


class AwaitableGetResourceSchemaResult(GetResourceSchemaResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetResourceSchemaResult(
            attributes=self.attributes,
            block_attributes=self.block_attributes,
            id=self.id,
            name=self.name,
            resource_type=self.resource_type)


def get_resource_schema(name: Optional[str] = None,
                        resource_type: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetResourceSchemaResult:
    """
    Provides a VMware Cloud Director generic structure data source. It shows the structure of any VCD resource.

    Supported in provider *v3.1+*

    ## Example Usage

    ### 1

    Showing a structure with simple attributes only

    ```python
    import pulumi
    import pulumi_vcd as vcd

    org_struct = vcd.get_resource_schema(name="org_struct",
        resource_type="vcd_org")
    pulumi.export("orgStruct", org_struct.attributes)
    ```

    ### 2

    Showing a structure with both simple and compound attributes

    ```python
    import pulumi
    import pulumi_vcd as vcd

    network_isolated_struct = vcd.get_resource_schema(name="net_struct",
        resource_type="vcd_network_isolated")
    pulumi.export("netStruct", net_struct_vcd_resource_schema)
    ```


    :param str name: An unique name to identify the data source
    :param str resource_type: Which resource we want to list. It needs to use the full name of the resource (i.e. "Org",
           not simply "org")
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceType'] = resource_type
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getResourceSchema:getResourceSchema', __args__, opts=opts, typ=GetResourceSchemaResult).value

    return AwaitableGetResourceSchemaResult(
        attributes=pulumi.get(__ret__, 'attributes'),
        block_attributes=pulumi.get(__ret__, 'block_attributes'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        resource_type=pulumi.get(__ret__, 'resource_type'))
def get_resource_schema_output(name: Optional[pulumi.Input[str]] = None,
                               resource_type: Optional[pulumi.Input[str]] = None,
                               opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetResourceSchemaResult]:
    """
    Provides a VMware Cloud Director generic structure data source. It shows the structure of any VCD resource.

    Supported in provider *v3.1+*

    ## Example Usage

    ### 1

    Showing a structure with simple attributes only

    ```python
    import pulumi
    import pulumi_vcd as vcd

    org_struct = vcd.get_resource_schema(name="org_struct",
        resource_type="vcd_org")
    pulumi.export("orgStruct", org_struct.attributes)
    ```

    ### 2

    Showing a structure with both simple and compound attributes

    ```python
    import pulumi
    import pulumi_vcd as vcd

    network_isolated_struct = vcd.get_resource_schema(name="net_struct",
        resource_type="vcd_network_isolated")
    pulumi.export("netStruct", net_struct_vcd_resource_schema)
    ```


    :param str name: An unique name to identify the data source
    :param str resource_type: Which resource we want to list. It needs to use the full name of the resource (i.e. "Org",
           not simply "org")
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['resourceType'] = resource_type
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vcd:index/getResourceSchema:getResourceSchema', __args__, opts=opts, typ=GetResourceSchemaResult)
    return __ret__.apply(lambda __response__: GetResourceSchemaResult(
        attributes=pulumi.get(__response__, 'attributes'),
        block_attributes=pulumi.get(__response__, 'block_attributes'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        resource_type=pulumi.get(__response__, 'resource_type')))
