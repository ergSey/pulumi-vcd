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

__all__ = ['NsxtAlbServiceEngineGroupArgs', 'NsxtAlbServiceEngineGroup']

@pulumi.input_type
class NsxtAlbServiceEngineGroupArgs:
    def __init__(__self__, *,
                 alb_cloud_id: pulumi.Input[str],
                 importable_service_engine_group_name: pulumi.Input[str],
                 reservation_model: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 overallocated: Optional[pulumi.Input[bool]] = None,
                 supported_feature_set: Optional[pulumi.Input[str]] = None,
                 sync_on_refresh: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a NsxtAlbServiceEngineGroup resource.
        :param pulumi.Input[str] alb_cloud_id: A reference ALB Cloud. Can be looked up using `NsxtAlbCloud` resource or data
               source
        :param pulumi.Input[str] importable_service_engine_group_name: Name of available Service Engine Group in ALB
        :param pulumi.Input[str] reservation_model: Definition if the Service Engine Group is `DEDICATED` or `SHARED`
        :param pulumi.Input[str] description: An optional description ALB Service Engine Group
        :param pulumi.Input[str] name: A name for ALB Service Engine Group
        :param pulumi.Input[bool] overallocated: Boolean value stating if there are more deployed virtual services than allocated ones
        :param pulumi.Input[str] supported_feature_set: Feature set of this ALB Service Engine Group (`STANDARD` or `PREMIUM`)
               
               > The attribute `supported_feature_set` must not be used in VCD versions lower than 10.4. Starting with 10.4, it replaces `license_type` field in [nsxt_alb_controller](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_controller).
        :param pulumi.Input[bool] sync_on_refresh: Boolean value that shows if sync should be performed on every refresh
        """
        pulumi.set(__self__, "alb_cloud_id", alb_cloud_id)
        pulumi.set(__self__, "importable_service_engine_group_name", importable_service_engine_group_name)
        pulumi.set(__self__, "reservation_model", reservation_model)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if overallocated is not None:
            pulumi.set(__self__, "overallocated", overallocated)
        if supported_feature_set is not None:
            pulumi.set(__self__, "supported_feature_set", supported_feature_set)
        if sync_on_refresh is not None:
            pulumi.set(__self__, "sync_on_refresh", sync_on_refresh)

    @property
    @pulumi.getter(name="albCloudId")
    def alb_cloud_id(self) -> pulumi.Input[str]:
        """
        A reference ALB Cloud. Can be looked up using `NsxtAlbCloud` resource or data
        source
        """
        return pulumi.get(self, "alb_cloud_id")

    @alb_cloud_id.setter
    def alb_cloud_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "alb_cloud_id", value)

    @property
    @pulumi.getter(name="importableServiceEngineGroupName")
    def importable_service_engine_group_name(self) -> pulumi.Input[str]:
        """
        Name of available Service Engine Group in ALB
        """
        return pulumi.get(self, "importable_service_engine_group_name")

    @importable_service_engine_group_name.setter
    def importable_service_engine_group_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "importable_service_engine_group_name", value)

    @property
    @pulumi.getter(name="reservationModel")
    def reservation_model(self) -> pulumi.Input[str]:
        """
        Definition if the Service Engine Group is `DEDICATED` or `SHARED`
        """
        return pulumi.get(self, "reservation_model")

    @reservation_model.setter
    def reservation_model(self, value: pulumi.Input[str]):
        pulumi.set(self, "reservation_model", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description ALB Service Engine Group
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A name for ALB Service Engine Group
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def overallocated(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean value stating if there are more deployed virtual services than allocated ones
        """
        return pulumi.get(self, "overallocated")

    @overallocated.setter
    def overallocated(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "overallocated", value)

    @property
    @pulumi.getter(name="supportedFeatureSet")
    def supported_feature_set(self) -> Optional[pulumi.Input[str]]:
        """
        Feature set of this ALB Service Engine Group (`STANDARD` or `PREMIUM`)

        > The attribute `supported_feature_set` must not be used in VCD versions lower than 10.4. Starting with 10.4, it replaces `license_type` field in [nsxt_alb_controller](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_controller).
        """
        return pulumi.get(self, "supported_feature_set")

    @supported_feature_set.setter
    def supported_feature_set(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "supported_feature_set", value)

    @property
    @pulumi.getter(name="syncOnRefresh")
    def sync_on_refresh(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean value that shows if sync should be performed on every refresh
        """
        return pulumi.get(self, "sync_on_refresh")

    @sync_on_refresh.setter
    def sync_on_refresh(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "sync_on_refresh", value)


@pulumi.input_type
class _NsxtAlbServiceEngineGroupState:
    def __init__(__self__, *,
                 alb_cloud_id: Optional[pulumi.Input[str]] = None,
                 deployed_virtual_services: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 ha_mode: Optional[pulumi.Input[str]] = None,
                 importable_service_engine_group_name: Optional[pulumi.Input[str]] = None,
                 max_virtual_services: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 overallocated: Optional[pulumi.Input[bool]] = None,
                 reservation_model: Optional[pulumi.Input[str]] = None,
                 reserved_virtual_services: Optional[pulumi.Input[int]] = None,
                 supported_feature_set: Optional[pulumi.Input[str]] = None,
                 sync_on_refresh: Optional[pulumi.Input[bool]] = None):
        """
        Input properties used for looking up and filtering NsxtAlbServiceEngineGroup resources.
        :param pulumi.Input[str] alb_cloud_id: A reference ALB Cloud. Can be looked up using `NsxtAlbCloud` resource or data
               source
        :param pulumi.Input[int] deployed_virtual_services: Number of deployed virtual services
        :param pulumi.Input[str] description: An optional description ALB Service Engine Group
        :param pulumi.Input[str] ha_mode: defines High Availability Mode for Service Engine Group. One off:
               * ELASTIC_N_PLUS_M_BUFFER - Service Engines will scale out to N active nodes with M nodes as buffer.
               * ELASTIC_ACTIVE_ACTIVE - Active-Active with scale out.
               * LEGACY_ACTIVE_STANDBY - Traditional single Active-Standby configuration
        :param pulumi.Input[str] importable_service_engine_group_name: Name of available Service Engine Group in ALB
        :param pulumi.Input[int] max_virtual_services: Maximum number of virtual services this ALB Service Engine Group can run
        :param pulumi.Input[str] name: A name for ALB Service Engine Group
        :param pulumi.Input[bool] overallocated: Boolean value stating if there are more deployed virtual services than allocated ones
        :param pulumi.Input[str] reservation_model: Definition if the Service Engine Group is `DEDICATED` or `SHARED`
        :param pulumi.Input[int] reserved_virtual_services: Number of reserved virtual services
        :param pulumi.Input[str] supported_feature_set: Feature set of this ALB Service Engine Group (`STANDARD` or `PREMIUM`)
               
               > The attribute `supported_feature_set` must not be used in VCD versions lower than 10.4. Starting with 10.4, it replaces `license_type` field in [nsxt_alb_controller](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_controller).
        :param pulumi.Input[bool] sync_on_refresh: Boolean value that shows if sync should be performed on every refresh
        """
        if alb_cloud_id is not None:
            pulumi.set(__self__, "alb_cloud_id", alb_cloud_id)
        if deployed_virtual_services is not None:
            pulumi.set(__self__, "deployed_virtual_services", deployed_virtual_services)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if ha_mode is not None:
            pulumi.set(__self__, "ha_mode", ha_mode)
        if importable_service_engine_group_name is not None:
            pulumi.set(__self__, "importable_service_engine_group_name", importable_service_engine_group_name)
        if max_virtual_services is not None:
            pulumi.set(__self__, "max_virtual_services", max_virtual_services)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if overallocated is not None:
            pulumi.set(__self__, "overallocated", overallocated)
        if reservation_model is not None:
            pulumi.set(__self__, "reservation_model", reservation_model)
        if reserved_virtual_services is not None:
            pulumi.set(__self__, "reserved_virtual_services", reserved_virtual_services)
        if supported_feature_set is not None:
            pulumi.set(__self__, "supported_feature_set", supported_feature_set)
        if sync_on_refresh is not None:
            pulumi.set(__self__, "sync_on_refresh", sync_on_refresh)

    @property
    @pulumi.getter(name="albCloudId")
    def alb_cloud_id(self) -> Optional[pulumi.Input[str]]:
        """
        A reference ALB Cloud. Can be looked up using `NsxtAlbCloud` resource or data
        source
        """
        return pulumi.get(self, "alb_cloud_id")

    @alb_cloud_id.setter
    def alb_cloud_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alb_cloud_id", value)

    @property
    @pulumi.getter(name="deployedVirtualServices")
    def deployed_virtual_services(self) -> Optional[pulumi.Input[int]]:
        """
        Number of deployed virtual services
        """
        return pulumi.get(self, "deployed_virtual_services")

    @deployed_virtual_services.setter
    def deployed_virtual_services(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "deployed_virtual_services", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        An optional description ALB Service Engine Group
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="haMode")
    def ha_mode(self) -> Optional[pulumi.Input[str]]:
        """
        defines High Availability Mode for Service Engine Group. One off:
        * ELASTIC_N_PLUS_M_BUFFER - Service Engines will scale out to N active nodes with M nodes as buffer.
        * ELASTIC_ACTIVE_ACTIVE - Active-Active with scale out.
        * LEGACY_ACTIVE_STANDBY - Traditional single Active-Standby configuration
        """
        return pulumi.get(self, "ha_mode")

    @ha_mode.setter
    def ha_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "ha_mode", value)

    @property
    @pulumi.getter(name="importableServiceEngineGroupName")
    def importable_service_engine_group_name(self) -> Optional[pulumi.Input[str]]:
        """
        Name of available Service Engine Group in ALB
        """
        return pulumi.get(self, "importable_service_engine_group_name")

    @importable_service_engine_group_name.setter
    def importable_service_engine_group_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "importable_service_engine_group_name", value)

    @property
    @pulumi.getter(name="maxVirtualServices")
    def max_virtual_services(self) -> Optional[pulumi.Input[int]]:
        """
        Maximum number of virtual services this ALB Service Engine Group can run
        """
        return pulumi.get(self, "max_virtual_services")

    @max_virtual_services.setter
    def max_virtual_services(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "max_virtual_services", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A name for ALB Service Engine Group
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def overallocated(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean value stating if there are more deployed virtual services than allocated ones
        """
        return pulumi.get(self, "overallocated")

    @overallocated.setter
    def overallocated(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "overallocated", value)

    @property
    @pulumi.getter(name="reservationModel")
    def reservation_model(self) -> Optional[pulumi.Input[str]]:
        """
        Definition if the Service Engine Group is `DEDICATED` or `SHARED`
        """
        return pulumi.get(self, "reservation_model")

    @reservation_model.setter
    def reservation_model(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "reservation_model", value)

    @property
    @pulumi.getter(name="reservedVirtualServices")
    def reserved_virtual_services(self) -> Optional[pulumi.Input[int]]:
        """
        Number of reserved virtual services
        """
        return pulumi.get(self, "reserved_virtual_services")

    @reserved_virtual_services.setter
    def reserved_virtual_services(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "reserved_virtual_services", value)

    @property
    @pulumi.getter(name="supportedFeatureSet")
    def supported_feature_set(self) -> Optional[pulumi.Input[str]]:
        """
        Feature set of this ALB Service Engine Group (`STANDARD` or `PREMIUM`)

        > The attribute `supported_feature_set` must not be used in VCD versions lower than 10.4. Starting with 10.4, it replaces `license_type` field in [nsxt_alb_controller](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_controller).
        """
        return pulumi.get(self, "supported_feature_set")

    @supported_feature_set.setter
    def supported_feature_set(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "supported_feature_set", value)

    @property
    @pulumi.getter(name="syncOnRefresh")
    def sync_on_refresh(self) -> Optional[pulumi.Input[bool]]:
        """
        Boolean value that shows if sync should be performed on every refresh
        """
        return pulumi.get(self, "sync_on_refresh")

    @sync_on_refresh.setter
    def sync_on_refresh(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "sync_on_refresh", value)


class NsxtAlbServiceEngineGroup(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alb_cloud_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 importable_service_engine_group_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 overallocated: Optional[pulumi.Input[bool]] = None,
                 reservation_model: Optional[pulumi.Input[str]] = None,
                 supported_feature_set: Optional[pulumi.Input[str]] = None,
                 sync_on_refresh: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Create a NsxtAlbServiceEngineGroup resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alb_cloud_id: A reference ALB Cloud. Can be looked up using `NsxtAlbCloud` resource or data
               source
        :param pulumi.Input[str] description: An optional description ALB Service Engine Group
        :param pulumi.Input[str] importable_service_engine_group_name: Name of available Service Engine Group in ALB
        :param pulumi.Input[str] name: A name for ALB Service Engine Group
        :param pulumi.Input[bool] overallocated: Boolean value stating if there are more deployed virtual services than allocated ones
        :param pulumi.Input[str] reservation_model: Definition if the Service Engine Group is `DEDICATED` or `SHARED`
        :param pulumi.Input[str] supported_feature_set: Feature set of this ALB Service Engine Group (`STANDARD` or `PREMIUM`)
               
               > The attribute `supported_feature_set` must not be used in VCD versions lower than 10.4. Starting with 10.4, it replaces `license_type` field in [nsxt_alb_controller](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_controller).
        :param pulumi.Input[bool] sync_on_refresh: Boolean value that shows if sync should be performed on every refresh
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NsxtAlbServiceEngineGroupArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a NsxtAlbServiceEngineGroup resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param NsxtAlbServiceEngineGroupArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NsxtAlbServiceEngineGroupArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 alb_cloud_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 importable_service_engine_group_name: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 overallocated: Optional[pulumi.Input[bool]] = None,
                 reservation_model: Optional[pulumi.Input[str]] = None,
                 supported_feature_set: Optional[pulumi.Input[str]] = None,
                 sync_on_refresh: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NsxtAlbServiceEngineGroupArgs.__new__(NsxtAlbServiceEngineGroupArgs)

            if alb_cloud_id is None and not opts.urn:
                raise TypeError("Missing required property 'alb_cloud_id'")
            __props__.__dict__["alb_cloud_id"] = alb_cloud_id
            __props__.__dict__["description"] = description
            if importable_service_engine_group_name is None and not opts.urn:
                raise TypeError("Missing required property 'importable_service_engine_group_name'")
            __props__.__dict__["importable_service_engine_group_name"] = importable_service_engine_group_name
            __props__.__dict__["name"] = name
            __props__.__dict__["overallocated"] = overallocated
            if reservation_model is None and not opts.urn:
                raise TypeError("Missing required property 'reservation_model'")
            __props__.__dict__["reservation_model"] = reservation_model
            __props__.__dict__["supported_feature_set"] = supported_feature_set
            __props__.__dict__["sync_on_refresh"] = sync_on_refresh
            __props__.__dict__["deployed_virtual_services"] = None
            __props__.__dict__["ha_mode"] = None
            __props__.__dict__["max_virtual_services"] = None
            __props__.__dict__["reserved_virtual_services"] = None
        super(NsxtAlbServiceEngineGroup, __self__).__init__(
            'vcd:index/nsxtAlbServiceEngineGroup:NsxtAlbServiceEngineGroup',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            alb_cloud_id: Optional[pulumi.Input[str]] = None,
            deployed_virtual_services: Optional[pulumi.Input[int]] = None,
            description: Optional[pulumi.Input[str]] = None,
            ha_mode: Optional[pulumi.Input[str]] = None,
            importable_service_engine_group_name: Optional[pulumi.Input[str]] = None,
            max_virtual_services: Optional[pulumi.Input[int]] = None,
            name: Optional[pulumi.Input[str]] = None,
            overallocated: Optional[pulumi.Input[bool]] = None,
            reservation_model: Optional[pulumi.Input[str]] = None,
            reserved_virtual_services: Optional[pulumi.Input[int]] = None,
            supported_feature_set: Optional[pulumi.Input[str]] = None,
            sync_on_refresh: Optional[pulumi.Input[bool]] = None) -> 'NsxtAlbServiceEngineGroup':
        """
        Get an existing NsxtAlbServiceEngineGroup resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alb_cloud_id: A reference ALB Cloud. Can be looked up using `NsxtAlbCloud` resource or data
               source
        :param pulumi.Input[int] deployed_virtual_services: Number of deployed virtual services
        :param pulumi.Input[str] description: An optional description ALB Service Engine Group
        :param pulumi.Input[str] ha_mode: defines High Availability Mode for Service Engine Group. One off:
               * ELASTIC_N_PLUS_M_BUFFER - Service Engines will scale out to N active nodes with M nodes as buffer.
               * ELASTIC_ACTIVE_ACTIVE - Active-Active with scale out.
               * LEGACY_ACTIVE_STANDBY - Traditional single Active-Standby configuration
        :param pulumi.Input[str] importable_service_engine_group_name: Name of available Service Engine Group in ALB
        :param pulumi.Input[int] max_virtual_services: Maximum number of virtual services this ALB Service Engine Group can run
        :param pulumi.Input[str] name: A name for ALB Service Engine Group
        :param pulumi.Input[bool] overallocated: Boolean value stating if there are more deployed virtual services than allocated ones
        :param pulumi.Input[str] reservation_model: Definition if the Service Engine Group is `DEDICATED` or `SHARED`
        :param pulumi.Input[int] reserved_virtual_services: Number of reserved virtual services
        :param pulumi.Input[str] supported_feature_set: Feature set of this ALB Service Engine Group (`STANDARD` or `PREMIUM`)
               
               > The attribute `supported_feature_set` must not be used in VCD versions lower than 10.4. Starting with 10.4, it replaces `license_type` field in [nsxt_alb_controller](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_controller).
        :param pulumi.Input[bool] sync_on_refresh: Boolean value that shows if sync should be performed on every refresh
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NsxtAlbServiceEngineGroupState.__new__(_NsxtAlbServiceEngineGroupState)

        __props__.__dict__["alb_cloud_id"] = alb_cloud_id
        __props__.__dict__["deployed_virtual_services"] = deployed_virtual_services
        __props__.__dict__["description"] = description
        __props__.__dict__["ha_mode"] = ha_mode
        __props__.__dict__["importable_service_engine_group_name"] = importable_service_engine_group_name
        __props__.__dict__["max_virtual_services"] = max_virtual_services
        __props__.__dict__["name"] = name
        __props__.__dict__["overallocated"] = overallocated
        __props__.__dict__["reservation_model"] = reservation_model
        __props__.__dict__["reserved_virtual_services"] = reserved_virtual_services
        __props__.__dict__["supported_feature_set"] = supported_feature_set
        __props__.__dict__["sync_on_refresh"] = sync_on_refresh
        return NsxtAlbServiceEngineGroup(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="albCloudId")
    def alb_cloud_id(self) -> pulumi.Output[str]:
        """
        A reference ALB Cloud. Can be looked up using `NsxtAlbCloud` resource or data
        source
        """
        return pulumi.get(self, "alb_cloud_id")

    @property
    @pulumi.getter(name="deployedVirtualServices")
    def deployed_virtual_services(self) -> pulumi.Output[int]:
        """
        Number of deployed virtual services
        """
        return pulumi.get(self, "deployed_virtual_services")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        An optional description ALB Service Engine Group
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="haMode")
    def ha_mode(self) -> pulumi.Output[str]:
        """
        defines High Availability Mode for Service Engine Group. One off:
        * ELASTIC_N_PLUS_M_BUFFER - Service Engines will scale out to N active nodes with M nodes as buffer.
        * ELASTIC_ACTIVE_ACTIVE - Active-Active with scale out.
        * LEGACY_ACTIVE_STANDBY - Traditional single Active-Standby configuration
        """
        return pulumi.get(self, "ha_mode")

    @property
    @pulumi.getter(name="importableServiceEngineGroupName")
    def importable_service_engine_group_name(self) -> pulumi.Output[str]:
        """
        Name of available Service Engine Group in ALB
        """
        return pulumi.get(self, "importable_service_engine_group_name")

    @property
    @pulumi.getter(name="maxVirtualServices")
    def max_virtual_services(self) -> pulumi.Output[int]:
        """
        Maximum number of virtual services this ALB Service Engine Group can run
        """
        return pulumi.get(self, "max_virtual_services")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A name for ALB Service Engine Group
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def overallocated(self) -> pulumi.Output[Optional[bool]]:
        """
        Boolean value stating if there are more deployed virtual services than allocated ones
        """
        return pulumi.get(self, "overallocated")

    @property
    @pulumi.getter(name="reservationModel")
    def reservation_model(self) -> pulumi.Output[str]:
        """
        Definition if the Service Engine Group is `DEDICATED` or `SHARED`
        """
        return pulumi.get(self, "reservation_model")

    @property
    @pulumi.getter(name="reservedVirtualServices")
    def reserved_virtual_services(self) -> pulumi.Output[int]:
        """
        Number of reserved virtual services
        """
        return pulumi.get(self, "reserved_virtual_services")

    @property
    @pulumi.getter(name="supportedFeatureSet")
    def supported_feature_set(self) -> pulumi.Output[str]:
        """
        Feature set of this ALB Service Engine Group (`STANDARD` or `PREMIUM`)

        > The attribute `supported_feature_set` must not be used in VCD versions lower than 10.4. Starting with 10.4, it replaces `license_type` field in [nsxt_alb_controller](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/nsxt_alb_controller).
        """
        return pulumi.get(self, "supported_feature_set")

    @property
    @pulumi.getter(name="syncOnRefresh")
    def sync_on_refresh(self) -> pulumi.Output[Optional[bool]]:
        """
        Boolean value that shows if sync should be performed on every refresh
        """
        return pulumi.get(self, "sync_on_refresh")

