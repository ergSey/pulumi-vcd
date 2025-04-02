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
    'GetRdeInterfaceBehaviorResult',
    'AwaitableGetRdeInterfaceBehaviorResult',
    'get_rde_interface_behavior',
    'get_rde_interface_behavior_output',
]

@pulumi.output_type
class GetRdeInterfaceBehaviorResult:
    """
    A collection of values returned by getRdeInterfaceBehavior.
    """
    def __init__(__self__, description=None, execution=None, execution_json=None, id=None, name=None, rde_interface_id=None, ref=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if execution and not isinstance(execution, dict):
            raise TypeError("Expected argument 'execution' to be a dict")
        pulumi.set(__self__, "execution", execution)
        if execution_json and not isinstance(execution_json, str):
            raise TypeError("Expected argument 'execution_json' to be a str")
        pulumi.set(__self__, "execution_json", execution_json)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if rde_interface_id and not isinstance(rde_interface_id, str):
            raise TypeError("Expected argument 'rde_interface_id' to be a str")
        pulumi.set(__self__, "rde_interface_id", rde_interface_id)
        if ref and not isinstance(ref, str):
            raise TypeError("Expected argument 'ref' to be a str")
        pulumi.set(__self__, "ref", ref)

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    @_utilities.deprecated("""This argument cannot consider complex execution structures. For that purpose, use 'execution_json' instead""")
    def execution(self) -> Mapping[str, str]:
        return pulumi.get(self, "execution")

    @property
    @pulumi.getter(name="executionJson")
    def execution_json(self) -> str:
        return pulumi.get(self, "execution_json")

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
    @pulumi.getter(name="rdeInterfaceId")
    def rde_interface_id(self) -> str:
        return pulumi.get(self, "rde_interface_id")

    @property
    @pulumi.getter
    def ref(self) -> str:
        return pulumi.get(self, "ref")


class AwaitableGetRdeInterfaceBehaviorResult(GetRdeInterfaceBehaviorResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRdeInterfaceBehaviorResult(
            description=self.description,
            execution=self.execution,
            execution_json=self.execution_json,
            id=self.id,
            name=self.name,
            rde_interface_id=self.rde_interface_id,
            ref=self.ref)


def get_rde_interface_behavior(name: Optional[str] = None,
                               rde_interface_id: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRdeInterfaceBehaviorResult:
    """
    Provides the capability of fetching an existing RDE Interface Behavior from VMware Cloud Director.

    Supported in provider *v3.10+*. Requires System Administrator privileges.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vcd as vcd

    my_interface = vcd.get_rde_interface(vendor="vmware",
        nss="k8s",
        version="1.0.0")
    my_behavior = vcd.get_rde_interface_behavior(rde_interface_id=my_interface.id,
        name="createKubeConfig")
    pulumi.export("executionId", my_behavior.execution["id"])
    pulumi.export("executionType", my_behavior.execution["type"])
    ```


    :param str name: The name of the RDE Interface Behavior to fetch
    :param str rde_interface_id: The ID of the RDE Interface that owns the Behavior to fetch
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['rdeInterfaceId'] = rde_interface_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getRdeInterfaceBehavior:getRdeInterfaceBehavior', __args__, opts=opts, typ=GetRdeInterfaceBehaviorResult).value

    return AwaitableGetRdeInterfaceBehaviorResult(
        description=pulumi.get(__ret__, 'description'),
        execution=pulumi.get(__ret__, 'execution'),
        execution_json=pulumi.get(__ret__, 'execution_json'),
        id=pulumi.get(__ret__, 'id'),
        name=pulumi.get(__ret__, 'name'),
        rde_interface_id=pulumi.get(__ret__, 'rde_interface_id'),
        ref=pulumi.get(__ret__, 'ref'))
def get_rde_interface_behavior_output(name: Optional[pulumi.Input[str]] = None,
                                      rde_interface_id: Optional[pulumi.Input[str]] = None,
                                      opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetRdeInterfaceBehaviorResult]:
    """
    Provides the capability of fetching an existing RDE Interface Behavior from VMware Cloud Director.

    Supported in provider *v3.10+*. Requires System Administrator privileges.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vcd as vcd

    my_interface = vcd.get_rde_interface(vendor="vmware",
        nss="k8s",
        version="1.0.0")
    my_behavior = vcd.get_rde_interface_behavior(rde_interface_id=my_interface.id,
        name="createKubeConfig")
    pulumi.export("executionId", my_behavior.execution["id"])
    pulumi.export("executionType", my_behavior.execution["type"])
    ```


    :param str name: The name of the RDE Interface Behavior to fetch
    :param str rde_interface_id: The ID of the RDE Interface that owns the Behavior to fetch
    """
    __args__ = dict()
    __args__['name'] = name
    __args__['rdeInterfaceId'] = rde_interface_id
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vcd:index/getRdeInterfaceBehavior:getRdeInterfaceBehavior', __args__, opts=opts, typ=GetRdeInterfaceBehaviorResult)
    return __ret__.apply(lambda __response__: GetRdeInterfaceBehaviorResult(
        description=pulumi.get(__response__, 'description'),
        execution=pulumi.get(__response__, 'execution'),
        execution_json=pulumi.get(__response__, 'execution_json'),
        id=pulumi.get(__response__, 'id'),
        name=pulumi.get(__response__, 'name'),
        rde_interface_id=pulumi.get(__response__, 'rde_interface_id'),
        ref=pulumi.get(__response__, 'ref')))
