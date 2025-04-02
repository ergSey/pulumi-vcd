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

__all__ = ['OrgVdcTemplateInstanceArgs', 'OrgVdcTemplateInstance']

@pulumi.input_type
class OrgVdcTemplateInstanceArgs:
    def __init__(__self__, *,
                 delete_instantiated_vdc_on_removal: pulumi.Input[bool],
                 org_id: pulumi.Input[str],
                 org_vdc_template_id: pulumi.Input[str],
                 delete_force: Optional[pulumi.Input[bool]] = None,
                 delete_recursive: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a OrgVdcTemplateInstance resource.
        :param pulumi.Input[bool] delete_instantiated_vdc_on_removal: If this flag is set to `true`, removing this resource will attempt to delete the instantiated VDC
        :param pulumi.Input[str] org_id: ID of the Organization where the VDC will be instantiated
        :param pulumi.Input[str] org_vdc_template_id: The ID of the VDC Template to instantiate
        :param pulumi.Input[bool] delete_force: Defaults to `false`. If this flag is set to `true`, it forcefully deletes the VDC, only when `delete_instantiated_vdc_on_removal=true`
        :param pulumi.Input[bool] delete_recursive: Defaults to `false`. If this flag is set to `true`, it recursively deletes the VDC, only when `delete_instantiated_vdc_on_removal=true`
        :param pulumi.Input[str] description: Description of the instantiated Organization VDC
        :param pulumi.Input[str] name: Name to give to the instantiated Organization VDC
        """
        pulumi.set(__self__, "delete_instantiated_vdc_on_removal", delete_instantiated_vdc_on_removal)
        pulumi.set(__self__, "org_id", org_id)
        pulumi.set(__self__, "org_vdc_template_id", org_vdc_template_id)
        if delete_force is not None:
            pulumi.set(__self__, "delete_force", delete_force)
        if delete_recursive is not None:
            pulumi.set(__self__, "delete_recursive", delete_recursive)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="deleteInstantiatedVdcOnRemoval")
    def delete_instantiated_vdc_on_removal(self) -> pulumi.Input[bool]:
        """
        If this flag is set to `true`, removing this resource will attempt to delete the instantiated VDC
        """
        return pulumi.get(self, "delete_instantiated_vdc_on_removal")

    @delete_instantiated_vdc_on_removal.setter
    def delete_instantiated_vdc_on_removal(self, value: pulumi.Input[bool]):
        pulumi.set(self, "delete_instantiated_vdc_on_removal", value)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Input[str]:
        """
        ID of the Organization where the VDC will be instantiated
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "org_id", value)

    @property
    @pulumi.getter(name="orgVdcTemplateId")
    def org_vdc_template_id(self) -> pulumi.Input[str]:
        """
        The ID of the VDC Template to instantiate
        """
        return pulumi.get(self, "org_vdc_template_id")

    @org_vdc_template_id.setter
    def org_vdc_template_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "org_vdc_template_id", value)

    @property
    @pulumi.getter(name="deleteForce")
    def delete_force(self) -> Optional[pulumi.Input[bool]]:
        """
        Defaults to `false`. If this flag is set to `true`, it forcefully deletes the VDC, only when `delete_instantiated_vdc_on_removal=true`
        """
        return pulumi.get(self, "delete_force")

    @delete_force.setter
    def delete_force(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "delete_force", value)

    @property
    @pulumi.getter(name="deleteRecursive")
    def delete_recursive(self) -> Optional[pulumi.Input[bool]]:
        """
        Defaults to `false`. If this flag is set to `true`, it recursively deletes the VDC, only when `delete_instantiated_vdc_on_removal=true`
        """
        return pulumi.get(self, "delete_recursive")

    @delete_recursive.setter
    def delete_recursive(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "delete_recursive", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the instantiated Organization VDC
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name to give to the instantiated Organization VDC
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _OrgVdcTemplateInstanceState:
    def __init__(__self__, *,
                 delete_force: Optional[pulumi.Input[bool]] = None,
                 delete_instantiated_vdc_on_removal: Optional[pulumi.Input[bool]] = None,
                 delete_recursive: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 org_vdc_template_id: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering OrgVdcTemplateInstance resources.
        :param pulumi.Input[bool] delete_force: Defaults to `false`. If this flag is set to `true`, it forcefully deletes the VDC, only when `delete_instantiated_vdc_on_removal=true`
        :param pulumi.Input[bool] delete_instantiated_vdc_on_removal: If this flag is set to `true`, removing this resource will attempt to delete the instantiated VDC
        :param pulumi.Input[bool] delete_recursive: Defaults to `false`. If this flag is set to `true`, it recursively deletes the VDC, only when `delete_instantiated_vdc_on_removal=true`
        :param pulumi.Input[str] description: Description of the instantiated Organization VDC
        :param pulumi.Input[str] name: Name to give to the instantiated Organization VDC
        :param pulumi.Input[str] org_id: ID of the Organization where the VDC will be instantiated
        :param pulumi.Input[str] org_vdc_template_id: The ID of the VDC Template to instantiate
        """
        if delete_force is not None:
            pulumi.set(__self__, "delete_force", delete_force)
        if delete_instantiated_vdc_on_removal is not None:
            pulumi.set(__self__, "delete_instantiated_vdc_on_removal", delete_instantiated_vdc_on_removal)
        if delete_recursive is not None:
            pulumi.set(__self__, "delete_recursive", delete_recursive)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org_id is not None:
            pulumi.set(__self__, "org_id", org_id)
        if org_vdc_template_id is not None:
            pulumi.set(__self__, "org_vdc_template_id", org_vdc_template_id)

    @property
    @pulumi.getter(name="deleteForce")
    def delete_force(self) -> Optional[pulumi.Input[bool]]:
        """
        Defaults to `false`. If this flag is set to `true`, it forcefully deletes the VDC, only when `delete_instantiated_vdc_on_removal=true`
        """
        return pulumi.get(self, "delete_force")

    @delete_force.setter
    def delete_force(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "delete_force", value)

    @property
    @pulumi.getter(name="deleteInstantiatedVdcOnRemoval")
    def delete_instantiated_vdc_on_removal(self) -> Optional[pulumi.Input[bool]]:
        """
        If this flag is set to `true`, removing this resource will attempt to delete the instantiated VDC
        """
        return pulumi.get(self, "delete_instantiated_vdc_on_removal")

    @delete_instantiated_vdc_on_removal.setter
    def delete_instantiated_vdc_on_removal(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "delete_instantiated_vdc_on_removal", value)

    @property
    @pulumi.getter(name="deleteRecursive")
    def delete_recursive(self) -> Optional[pulumi.Input[bool]]:
        """
        Defaults to `false`. If this flag is set to `true`, it recursively deletes the VDC, only when `delete_instantiated_vdc_on_removal=true`
        """
        return pulumi.get(self, "delete_recursive")

    @delete_recursive.setter
    def delete_recursive(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "delete_recursive", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description of the instantiated Organization VDC
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name to give to the instantiated Organization VDC
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> Optional[pulumi.Input[str]]:
        """
        ID of the Organization where the VDC will be instantiated
        """
        return pulumi.get(self, "org_id")

    @org_id.setter
    def org_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org_id", value)

    @property
    @pulumi.getter(name="orgVdcTemplateId")
    def org_vdc_template_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the VDC Template to instantiate
        """
        return pulumi.get(self, "org_vdc_template_id")

    @org_vdc_template_id.setter
    def org_vdc_template_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org_vdc_template_id", value)


class OrgVdcTemplateInstance(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 delete_force: Optional[pulumi.Input[bool]] = None,
                 delete_instantiated_vdc_on_removal: Optional[pulumi.Input[bool]] = None,
                 delete_recursive: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 org_vdc_template_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a OrgVdcTemplateInstance resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] delete_force: Defaults to `false`. If this flag is set to `true`, it forcefully deletes the VDC, only when `delete_instantiated_vdc_on_removal=true`
        :param pulumi.Input[bool] delete_instantiated_vdc_on_removal: If this flag is set to `true`, removing this resource will attempt to delete the instantiated VDC
        :param pulumi.Input[bool] delete_recursive: Defaults to `false`. If this flag is set to `true`, it recursively deletes the VDC, only when `delete_instantiated_vdc_on_removal=true`
        :param pulumi.Input[str] description: Description of the instantiated Organization VDC
        :param pulumi.Input[str] name: Name to give to the instantiated Organization VDC
        :param pulumi.Input[str] org_id: ID of the Organization where the VDC will be instantiated
        :param pulumi.Input[str] org_vdc_template_id: The ID of the VDC Template to instantiate
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OrgVdcTemplateInstanceArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a OrgVdcTemplateInstance resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param OrgVdcTemplateInstanceArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OrgVdcTemplateInstanceArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 delete_force: Optional[pulumi.Input[bool]] = None,
                 delete_instantiated_vdc_on_removal: Optional[pulumi.Input[bool]] = None,
                 delete_recursive: Optional[pulumi.Input[bool]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org_id: Optional[pulumi.Input[str]] = None,
                 org_vdc_template_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = OrgVdcTemplateInstanceArgs.__new__(OrgVdcTemplateInstanceArgs)

            __props__.__dict__["delete_force"] = delete_force
            if delete_instantiated_vdc_on_removal is None and not opts.urn:
                raise TypeError("Missing required property 'delete_instantiated_vdc_on_removal'")
            __props__.__dict__["delete_instantiated_vdc_on_removal"] = delete_instantiated_vdc_on_removal
            __props__.__dict__["delete_recursive"] = delete_recursive
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            if org_id is None and not opts.urn:
                raise TypeError("Missing required property 'org_id'")
            __props__.__dict__["org_id"] = org_id
            if org_vdc_template_id is None and not opts.urn:
                raise TypeError("Missing required property 'org_vdc_template_id'")
            __props__.__dict__["org_vdc_template_id"] = org_vdc_template_id
        super(OrgVdcTemplateInstance, __self__).__init__(
            'vcd:index/orgVdcTemplateInstance:OrgVdcTemplateInstance',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            delete_force: Optional[pulumi.Input[bool]] = None,
            delete_instantiated_vdc_on_removal: Optional[pulumi.Input[bool]] = None,
            delete_recursive: Optional[pulumi.Input[bool]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            org_id: Optional[pulumi.Input[str]] = None,
            org_vdc_template_id: Optional[pulumi.Input[str]] = None) -> 'OrgVdcTemplateInstance':
        """
        Get an existing OrgVdcTemplateInstance resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] delete_force: Defaults to `false`. If this flag is set to `true`, it forcefully deletes the VDC, only when `delete_instantiated_vdc_on_removal=true`
        :param pulumi.Input[bool] delete_instantiated_vdc_on_removal: If this flag is set to `true`, removing this resource will attempt to delete the instantiated VDC
        :param pulumi.Input[bool] delete_recursive: Defaults to `false`. If this flag is set to `true`, it recursively deletes the VDC, only when `delete_instantiated_vdc_on_removal=true`
        :param pulumi.Input[str] description: Description of the instantiated Organization VDC
        :param pulumi.Input[str] name: Name to give to the instantiated Organization VDC
        :param pulumi.Input[str] org_id: ID of the Organization where the VDC will be instantiated
        :param pulumi.Input[str] org_vdc_template_id: The ID of the VDC Template to instantiate
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _OrgVdcTemplateInstanceState.__new__(_OrgVdcTemplateInstanceState)

        __props__.__dict__["delete_force"] = delete_force
        __props__.__dict__["delete_instantiated_vdc_on_removal"] = delete_instantiated_vdc_on_removal
        __props__.__dict__["delete_recursive"] = delete_recursive
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["org_id"] = org_id
        __props__.__dict__["org_vdc_template_id"] = org_vdc_template_id
        return OrgVdcTemplateInstance(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="deleteForce")
    def delete_force(self) -> pulumi.Output[Optional[bool]]:
        """
        Defaults to `false`. If this flag is set to `true`, it forcefully deletes the VDC, only when `delete_instantiated_vdc_on_removal=true`
        """
        return pulumi.get(self, "delete_force")

    @property
    @pulumi.getter(name="deleteInstantiatedVdcOnRemoval")
    def delete_instantiated_vdc_on_removal(self) -> pulumi.Output[bool]:
        """
        If this flag is set to `true`, removing this resource will attempt to delete the instantiated VDC
        """
        return pulumi.get(self, "delete_instantiated_vdc_on_removal")

    @property
    @pulumi.getter(name="deleteRecursive")
    def delete_recursive(self) -> pulumi.Output[Optional[bool]]:
        """
        Defaults to `false`. If this flag is set to `true`, it recursively deletes the VDC, only when `delete_instantiated_vdc_on_removal=true`
        """
        return pulumi.get(self, "delete_recursive")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description of the instantiated Organization VDC
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name to give to the instantiated Organization VDC
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="orgId")
    def org_id(self) -> pulumi.Output[str]:
        """
        ID of the Organization where the VDC will be instantiated
        """
        return pulumi.get(self, "org_id")

    @property
    @pulumi.getter(name="orgVdcTemplateId")
    def org_vdc_template_id(self) -> pulumi.Output[str]:
        """
        The ID of the VDC Template to instantiate
        """
        return pulumi.get(self, "org_vdc_template_id")

