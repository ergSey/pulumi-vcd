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

__all__ = ['VmInternalDiskInitArgs', 'VmInternalDisk']

@pulumi.input_type
class VmInternalDiskInitArgs:
    def __init__(__self__, *,
                 bus_number: pulumi.Input[int],
                 bus_type: pulumi.Input[str],
                 size_in_mb: pulumi.Input[int],
                 unit_number: pulumi.Input[int],
                 vapp_name: pulumi.Input[str],
                 vm_name: pulumi.Input[str],
                 allow_vm_reboot: Optional[pulumi.Input[bool]] = None,
                 iops: Optional[pulumi.Input[int]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 storage_profile: Optional[pulumi.Input[str]] = None,
                 vdc: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a VmInternalDisk resource.
        :param pulumi.Input[int] bus_number: The number of the SCSI or IDE controller itself.
        :param pulumi.Input[str] bus_type: The type of disk controller. Possible values: `ide`, `parallel`( LSI Logic Parallel SCSI),
               `sas`(LSI Logic SAS (SCSI)), `paravirtual`(Paravirtual (SCSI)), `sata`, `nvme`. **Note** `nvme` requires *v3.4.0+* and
               VCD *10.2.1+*
        :param pulumi.Input[int] size_in_mb: The size of the disk in MB.
        :param pulumi.Input[int] unit_number: The device number on the SCSI or IDE controller of the disk.
        :param pulumi.Input[str] vapp_name: The vAPP this VM internal disk belongs to.
        :param pulumi.Input[str] vm_name: VM in vAPP in which internal disk is created.
        :param pulumi.Input[bool] allow_vm_reboot: Powers off VM when changing any attribute of an IDE disk or unit/bus number of other disk types, after the change is complete VM is powered back on. Without this setting enabled, such changes on a powered-on VM would fail. Defaults to false.
        :param pulumi.Input[int] iops: Specifies the IOPS for the disk. Default is 0.
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
        :param pulumi.Input[str] storage_profile: Storage profile which overrides the VM default one.
        :param pulumi.Input[str] vdc: The name of VDC to use, optional if defined at provider level
        """
        pulumi.set(__self__, "bus_number", bus_number)
        pulumi.set(__self__, "bus_type", bus_type)
        pulumi.set(__self__, "size_in_mb", size_in_mb)
        pulumi.set(__self__, "unit_number", unit_number)
        pulumi.set(__self__, "vapp_name", vapp_name)
        pulumi.set(__self__, "vm_name", vm_name)
        if allow_vm_reboot is not None:
            pulumi.set(__self__, "allow_vm_reboot", allow_vm_reboot)
        if iops is not None:
            pulumi.set(__self__, "iops", iops)
        if org is not None:
            pulumi.set(__self__, "org", org)
        if storage_profile is not None:
            pulumi.set(__self__, "storage_profile", storage_profile)
        if vdc is not None:
            pulumi.set(__self__, "vdc", vdc)

    @property
    @pulumi.getter(name="busNumber")
    def bus_number(self) -> pulumi.Input[int]:
        """
        The number of the SCSI or IDE controller itself.
        """
        return pulumi.get(self, "bus_number")

    @bus_number.setter
    def bus_number(self, value: pulumi.Input[int]):
        pulumi.set(self, "bus_number", value)

    @property
    @pulumi.getter(name="busType")
    def bus_type(self) -> pulumi.Input[str]:
        """
        The type of disk controller. Possible values: `ide`, `parallel`( LSI Logic Parallel SCSI),
        `sas`(LSI Logic SAS (SCSI)), `paravirtual`(Paravirtual (SCSI)), `sata`, `nvme`. **Note** `nvme` requires *v3.4.0+* and
        VCD *10.2.1+*
        """
        return pulumi.get(self, "bus_type")

    @bus_type.setter
    def bus_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "bus_type", value)

    @property
    @pulumi.getter(name="sizeInMb")
    def size_in_mb(self) -> pulumi.Input[int]:
        """
        The size of the disk in MB.
        """
        return pulumi.get(self, "size_in_mb")

    @size_in_mb.setter
    def size_in_mb(self, value: pulumi.Input[int]):
        pulumi.set(self, "size_in_mb", value)

    @property
    @pulumi.getter(name="unitNumber")
    def unit_number(self) -> pulumi.Input[int]:
        """
        The device number on the SCSI or IDE controller of the disk.
        """
        return pulumi.get(self, "unit_number")

    @unit_number.setter
    def unit_number(self, value: pulumi.Input[int]):
        pulumi.set(self, "unit_number", value)

    @property
    @pulumi.getter(name="vappName")
    def vapp_name(self) -> pulumi.Input[str]:
        """
        The vAPP this VM internal disk belongs to.
        """
        return pulumi.get(self, "vapp_name")

    @vapp_name.setter
    def vapp_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "vapp_name", value)

    @property
    @pulumi.getter(name="vmName")
    def vm_name(self) -> pulumi.Input[str]:
        """
        VM in vAPP in which internal disk is created.
        """
        return pulumi.get(self, "vm_name")

    @vm_name.setter
    def vm_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "vm_name", value)

    @property
    @pulumi.getter(name="allowVmReboot")
    def allow_vm_reboot(self) -> Optional[pulumi.Input[bool]]:
        """
        Powers off VM when changing any attribute of an IDE disk or unit/bus number of other disk types, after the change is complete VM is powered back on. Without this setting enabled, such changes on a powered-on VM would fail. Defaults to false.
        """
        return pulumi.get(self, "allow_vm_reboot")

    @allow_vm_reboot.setter
    def allow_vm_reboot(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_vm_reboot", value)

    @property
    @pulumi.getter
    def iops(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the IOPS for the disk. Default is 0.
        """
        return pulumi.get(self, "iops")

    @iops.setter
    def iops(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "iops", value)

    @property
    @pulumi.getter
    def org(self) -> Optional[pulumi.Input[str]]:
        """
        The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
        """
        return pulumi.get(self, "org")

    @org.setter
    def org(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org", value)

    @property
    @pulumi.getter(name="storageProfile")
    def storage_profile(self) -> Optional[pulumi.Input[str]]:
        """
        Storage profile which overrides the VM default one.
        """
        return pulumi.get(self, "storage_profile")

    @storage_profile.setter
    def storage_profile(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_profile", value)

    @property
    @pulumi.getter
    def vdc(self) -> Optional[pulumi.Input[str]]:
        """
        The name of VDC to use, optional if defined at provider level
        """
        return pulumi.get(self, "vdc")

    @vdc.setter
    def vdc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdc", value)


@pulumi.input_type
class _VmInternalDiskState:
    def __init__(__self__, *,
                 allow_vm_reboot: Optional[pulumi.Input[bool]] = None,
                 bus_number: Optional[pulumi.Input[int]] = None,
                 bus_type: Optional[pulumi.Input[str]] = None,
                 iops: Optional[pulumi.Input[int]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 size_in_mb: Optional[pulumi.Input[int]] = None,
                 storage_profile: Optional[pulumi.Input[str]] = None,
                 thin_provisioned: Optional[pulumi.Input[bool]] = None,
                 unit_number: Optional[pulumi.Input[int]] = None,
                 vapp_name: Optional[pulumi.Input[str]] = None,
                 vdc: Optional[pulumi.Input[str]] = None,
                 vm_name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VmInternalDisk resources.
        :param pulumi.Input[bool] allow_vm_reboot: Powers off VM when changing any attribute of an IDE disk or unit/bus number of other disk types, after the change is complete VM is powered back on. Without this setting enabled, such changes on a powered-on VM would fail. Defaults to false.
        :param pulumi.Input[int] bus_number: The number of the SCSI or IDE controller itself.
        :param pulumi.Input[str] bus_type: The type of disk controller. Possible values: `ide`, `parallel`( LSI Logic Parallel SCSI),
               `sas`(LSI Logic SAS (SCSI)), `paravirtual`(Paravirtual (SCSI)), `sata`, `nvme`. **Note** `nvme` requires *v3.4.0+* and
               VCD *10.2.1+*
        :param pulumi.Input[int] iops: Specifies the IOPS for the disk. Default is 0.
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
        :param pulumi.Input[int] size_in_mb: The size of the disk in MB.
        :param pulumi.Input[str] storage_profile: Storage profile which overrides the VM default one.
        :param pulumi.Input[bool] thin_provisioned: Specifies whether the disk storage is pre-allocated or allocated on demand.
        :param pulumi.Input[int] unit_number: The device number on the SCSI or IDE controller of the disk.
        :param pulumi.Input[str] vapp_name: The vAPP this VM internal disk belongs to.
        :param pulumi.Input[str] vdc: The name of VDC to use, optional if defined at provider level
        :param pulumi.Input[str] vm_name: VM in vAPP in which internal disk is created.
        """
        if allow_vm_reboot is not None:
            pulumi.set(__self__, "allow_vm_reboot", allow_vm_reboot)
        if bus_number is not None:
            pulumi.set(__self__, "bus_number", bus_number)
        if bus_type is not None:
            pulumi.set(__self__, "bus_type", bus_type)
        if iops is not None:
            pulumi.set(__self__, "iops", iops)
        if org is not None:
            pulumi.set(__self__, "org", org)
        if size_in_mb is not None:
            pulumi.set(__self__, "size_in_mb", size_in_mb)
        if storage_profile is not None:
            pulumi.set(__self__, "storage_profile", storage_profile)
        if thin_provisioned is not None:
            pulumi.set(__self__, "thin_provisioned", thin_provisioned)
        if unit_number is not None:
            pulumi.set(__self__, "unit_number", unit_number)
        if vapp_name is not None:
            pulumi.set(__self__, "vapp_name", vapp_name)
        if vdc is not None:
            pulumi.set(__self__, "vdc", vdc)
        if vm_name is not None:
            pulumi.set(__self__, "vm_name", vm_name)

    @property
    @pulumi.getter(name="allowVmReboot")
    def allow_vm_reboot(self) -> Optional[pulumi.Input[bool]]:
        """
        Powers off VM when changing any attribute of an IDE disk or unit/bus number of other disk types, after the change is complete VM is powered back on. Without this setting enabled, such changes on a powered-on VM would fail. Defaults to false.
        """
        return pulumi.get(self, "allow_vm_reboot")

    @allow_vm_reboot.setter
    def allow_vm_reboot(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "allow_vm_reboot", value)

    @property
    @pulumi.getter(name="busNumber")
    def bus_number(self) -> Optional[pulumi.Input[int]]:
        """
        The number of the SCSI or IDE controller itself.
        """
        return pulumi.get(self, "bus_number")

    @bus_number.setter
    def bus_number(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "bus_number", value)

    @property
    @pulumi.getter(name="busType")
    def bus_type(self) -> Optional[pulumi.Input[str]]:
        """
        The type of disk controller. Possible values: `ide`, `parallel`( LSI Logic Parallel SCSI),
        `sas`(LSI Logic SAS (SCSI)), `paravirtual`(Paravirtual (SCSI)), `sata`, `nvme`. **Note** `nvme` requires *v3.4.0+* and
        VCD *10.2.1+*
        """
        return pulumi.get(self, "bus_type")

    @bus_type.setter
    def bus_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "bus_type", value)

    @property
    @pulumi.getter
    def iops(self) -> Optional[pulumi.Input[int]]:
        """
        Specifies the IOPS for the disk. Default is 0.
        """
        return pulumi.get(self, "iops")

    @iops.setter
    def iops(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "iops", value)

    @property
    @pulumi.getter
    def org(self) -> Optional[pulumi.Input[str]]:
        """
        The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
        """
        return pulumi.get(self, "org")

    @org.setter
    def org(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org", value)

    @property
    @pulumi.getter(name="sizeInMb")
    def size_in_mb(self) -> Optional[pulumi.Input[int]]:
        """
        The size of the disk in MB.
        """
        return pulumi.get(self, "size_in_mb")

    @size_in_mb.setter
    def size_in_mb(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "size_in_mb", value)

    @property
    @pulumi.getter(name="storageProfile")
    def storage_profile(self) -> Optional[pulumi.Input[str]]:
        """
        Storage profile which overrides the VM default one.
        """
        return pulumi.get(self, "storage_profile")

    @storage_profile.setter
    def storage_profile(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "storage_profile", value)

    @property
    @pulumi.getter(name="thinProvisioned")
    def thin_provisioned(self) -> Optional[pulumi.Input[bool]]:
        """
        Specifies whether the disk storage is pre-allocated or allocated on demand.
        """
        return pulumi.get(self, "thin_provisioned")

    @thin_provisioned.setter
    def thin_provisioned(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "thin_provisioned", value)

    @property
    @pulumi.getter(name="unitNumber")
    def unit_number(self) -> Optional[pulumi.Input[int]]:
        """
        The device number on the SCSI or IDE controller of the disk.
        """
        return pulumi.get(self, "unit_number")

    @unit_number.setter
    def unit_number(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "unit_number", value)

    @property
    @pulumi.getter(name="vappName")
    def vapp_name(self) -> Optional[pulumi.Input[str]]:
        """
        The vAPP this VM internal disk belongs to.
        """
        return pulumi.get(self, "vapp_name")

    @vapp_name.setter
    def vapp_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vapp_name", value)

    @property
    @pulumi.getter
    def vdc(self) -> Optional[pulumi.Input[str]]:
        """
        The name of VDC to use, optional if defined at provider level
        """
        return pulumi.get(self, "vdc")

    @vdc.setter
    def vdc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdc", value)

    @property
    @pulumi.getter(name="vmName")
    def vm_name(self) -> Optional[pulumi.Input[str]]:
        """
        VM in vAPP in which internal disk is created.
        """
        return pulumi.get(self, "vm_name")

    @vm_name.setter
    def vm_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vm_name", value)


class VmInternalDisk(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_vm_reboot: Optional[pulumi.Input[bool]] = None,
                 bus_number: Optional[pulumi.Input[int]] = None,
                 bus_type: Optional[pulumi.Input[str]] = None,
                 iops: Optional[pulumi.Input[int]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 size_in_mb: Optional[pulumi.Input[int]] = None,
                 storage_profile: Optional[pulumi.Input[str]] = None,
                 unit_number: Optional[pulumi.Input[int]] = None,
                 vapp_name: Optional[pulumi.Input[str]] = None,
                 vdc: Optional[pulumi.Input[str]] = None,
                 vm_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a VmInternalDisk resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_vm_reboot: Powers off VM when changing any attribute of an IDE disk or unit/bus number of other disk types, after the change is complete VM is powered back on. Without this setting enabled, such changes on a powered-on VM would fail. Defaults to false.
        :param pulumi.Input[int] bus_number: The number of the SCSI or IDE controller itself.
        :param pulumi.Input[str] bus_type: The type of disk controller. Possible values: `ide`, `parallel`( LSI Logic Parallel SCSI),
               `sas`(LSI Logic SAS (SCSI)), `paravirtual`(Paravirtual (SCSI)), `sata`, `nvme`. **Note** `nvme` requires *v3.4.0+* and
               VCD *10.2.1+*
        :param pulumi.Input[int] iops: Specifies the IOPS for the disk. Default is 0.
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
        :param pulumi.Input[int] size_in_mb: The size of the disk in MB.
        :param pulumi.Input[str] storage_profile: Storage profile which overrides the VM default one.
        :param pulumi.Input[int] unit_number: The device number on the SCSI or IDE controller of the disk.
        :param pulumi.Input[str] vapp_name: The vAPP this VM internal disk belongs to.
        :param pulumi.Input[str] vdc: The name of VDC to use, optional if defined at provider level
        :param pulumi.Input[str] vm_name: VM in vAPP in which internal disk is created.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VmInternalDiskInitArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a VmInternalDisk resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param VmInternalDiskInitArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VmInternalDiskInitArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 allow_vm_reboot: Optional[pulumi.Input[bool]] = None,
                 bus_number: Optional[pulumi.Input[int]] = None,
                 bus_type: Optional[pulumi.Input[str]] = None,
                 iops: Optional[pulumi.Input[int]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 size_in_mb: Optional[pulumi.Input[int]] = None,
                 storage_profile: Optional[pulumi.Input[str]] = None,
                 unit_number: Optional[pulumi.Input[int]] = None,
                 vapp_name: Optional[pulumi.Input[str]] = None,
                 vdc: Optional[pulumi.Input[str]] = None,
                 vm_name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VmInternalDiskInitArgs.__new__(VmInternalDiskInitArgs)

            __props__.__dict__["allow_vm_reboot"] = allow_vm_reboot
            if bus_number is None and not opts.urn:
                raise TypeError("Missing required property 'bus_number'")
            __props__.__dict__["bus_number"] = bus_number
            if bus_type is None and not opts.urn:
                raise TypeError("Missing required property 'bus_type'")
            __props__.__dict__["bus_type"] = bus_type
            __props__.__dict__["iops"] = iops
            __props__.__dict__["org"] = org
            if size_in_mb is None and not opts.urn:
                raise TypeError("Missing required property 'size_in_mb'")
            __props__.__dict__["size_in_mb"] = size_in_mb
            __props__.__dict__["storage_profile"] = storage_profile
            if unit_number is None and not opts.urn:
                raise TypeError("Missing required property 'unit_number'")
            __props__.__dict__["unit_number"] = unit_number
            if vapp_name is None and not opts.urn:
                raise TypeError("Missing required property 'vapp_name'")
            __props__.__dict__["vapp_name"] = vapp_name
            __props__.__dict__["vdc"] = vdc
            if vm_name is None and not opts.urn:
                raise TypeError("Missing required property 'vm_name'")
            __props__.__dict__["vm_name"] = vm_name
            __props__.__dict__["thin_provisioned"] = None
        super(VmInternalDisk, __self__).__init__(
            'vcd:index/vmInternalDisk:VmInternalDisk',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            allow_vm_reboot: Optional[pulumi.Input[bool]] = None,
            bus_number: Optional[pulumi.Input[int]] = None,
            bus_type: Optional[pulumi.Input[str]] = None,
            iops: Optional[pulumi.Input[int]] = None,
            org: Optional[pulumi.Input[str]] = None,
            size_in_mb: Optional[pulumi.Input[int]] = None,
            storage_profile: Optional[pulumi.Input[str]] = None,
            thin_provisioned: Optional[pulumi.Input[bool]] = None,
            unit_number: Optional[pulumi.Input[int]] = None,
            vapp_name: Optional[pulumi.Input[str]] = None,
            vdc: Optional[pulumi.Input[str]] = None,
            vm_name: Optional[pulumi.Input[str]] = None) -> 'VmInternalDisk':
        """
        Get an existing VmInternalDisk resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[bool] allow_vm_reboot: Powers off VM when changing any attribute of an IDE disk or unit/bus number of other disk types, after the change is complete VM is powered back on. Without this setting enabled, such changes on a powered-on VM would fail. Defaults to false.
        :param pulumi.Input[int] bus_number: The number of the SCSI or IDE controller itself.
        :param pulumi.Input[str] bus_type: The type of disk controller. Possible values: `ide`, `parallel`( LSI Logic Parallel SCSI),
               `sas`(LSI Logic SAS (SCSI)), `paravirtual`(Paravirtual (SCSI)), `sata`, `nvme`. **Note** `nvme` requires *v3.4.0+* and
               VCD *10.2.1+*
        :param pulumi.Input[int] iops: Specifies the IOPS for the disk. Default is 0.
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
        :param pulumi.Input[int] size_in_mb: The size of the disk in MB.
        :param pulumi.Input[str] storage_profile: Storage profile which overrides the VM default one.
        :param pulumi.Input[bool] thin_provisioned: Specifies whether the disk storage is pre-allocated or allocated on demand.
        :param pulumi.Input[int] unit_number: The device number on the SCSI or IDE controller of the disk.
        :param pulumi.Input[str] vapp_name: The vAPP this VM internal disk belongs to.
        :param pulumi.Input[str] vdc: The name of VDC to use, optional if defined at provider level
        :param pulumi.Input[str] vm_name: VM in vAPP in which internal disk is created.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VmInternalDiskState.__new__(_VmInternalDiskState)

        __props__.__dict__["allow_vm_reboot"] = allow_vm_reboot
        __props__.__dict__["bus_number"] = bus_number
        __props__.__dict__["bus_type"] = bus_type
        __props__.__dict__["iops"] = iops
        __props__.__dict__["org"] = org
        __props__.__dict__["size_in_mb"] = size_in_mb
        __props__.__dict__["storage_profile"] = storage_profile
        __props__.__dict__["thin_provisioned"] = thin_provisioned
        __props__.__dict__["unit_number"] = unit_number
        __props__.__dict__["vapp_name"] = vapp_name
        __props__.__dict__["vdc"] = vdc
        __props__.__dict__["vm_name"] = vm_name
        return VmInternalDisk(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="allowVmReboot")
    def allow_vm_reboot(self) -> pulumi.Output[Optional[bool]]:
        """
        Powers off VM when changing any attribute of an IDE disk or unit/bus number of other disk types, after the change is complete VM is powered back on. Without this setting enabled, such changes on a powered-on VM would fail. Defaults to false.
        """
        return pulumi.get(self, "allow_vm_reboot")

    @property
    @pulumi.getter(name="busNumber")
    def bus_number(self) -> pulumi.Output[int]:
        """
        The number of the SCSI or IDE controller itself.
        """
        return pulumi.get(self, "bus_number")

    @property
    @pulumi.getter(name="busType")
    def bus_type(self) -> pulumi.Output[str]:
        """
        The type of disk controller. Possible values: `ide`, `parallel`( LSI Logic Parallel SCSI),
        `sas`(LSI Logic SAS (SCSI)), `paravirtual`(Paravirtual (SCSI)), `sata`, `nvme`. **Note** `nvme` requires *v3.4.0+* and
        VCD *10.2.1+*
        """
        return pulumi.get(self, "bus_type")

    @property
    @pulumi.getter
    def iops(self) -> pulumi.Output[int]:
        """
        Specifies the IOPS for the disk. Default is 0.
        """
        return pulumi.get(self, "iops")

    @property
    @pulumi.getter
    def org(self) -> pulumi.Output[Optional[str]]:
        """
        The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations
        """
        return pulumi.get(self, "org")

    @property
    @pulumi.getter(name="sizeInMb")
    def size_in_mb(self) -> pulumi.Output[int]:
        """
        The size of the disk in MB.
        """
        return pulumi.get(self, "size_in_mb")

    @property
    @pulumi.getter(name="storageProfile")
    def storage_profile(self) -> pulumi.Output[str]:
        """
        Storage profile which overrides the VM default one.
        """
        return pulumi.get(self, "storage_profile")

    @property
    @pulumi.getter(name="thinProvisioned")
    def thin_provisioned(self) -> pulumi.Output[bool]:
        """
        Specifies whether the disk storage is pre-allocated or allocated on demand.
        """
        return pulumi.get(self, "thin_provisioned")

    @property
    @pulumi.getter(name="unitNumber")
    def unit_number(self) -> pulumi.Output[int]:
        """
        The device number on the SCSI or IDE controller of the disk.
        """
        return pulumi.get(self, "unit_number")

    @property
    @pulumi.getter(name="vappName")
    def vapp_name(self) -> pulumi.Output[str]:
        """
        The vAPP this VM internal disk belongs to.
        """
        return pulumi.get(self, "vapp_name")

    @property
    @pulumi.getter
    def vdc(self) -> pulumi.Output[Optional[str]]:
        """
        The name of VDC to use, optional if defined at provider level
        """
        return pulumi.get(self, "vdc")

    @property
    @pulumi.getter(name="vmName")
    def vm_name(self) -> pulumi.Output[str]:
        """
        VM in vAPP in which internal disk is created.
        """
        return pulumi.get(self, "vm_name")

