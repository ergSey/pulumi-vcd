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
from ._inputs import *

__all__ = ['SolutionLandingZoneArgs', 'SolutionLandingZone']

@pulumi.input_type
class SolutionLandingZoneArgs:
    def __init__(__self__, *,
                 catalog: pulumi.Input['SolutionLandingZoneCatalogArgs'],
                 vdcs: pulumi.Input[Sequence[pulumi.Input['SolutionLandingZoneVdcArgs']]],
                 org: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a SolutionLandingZone resource.
        :param pulumi.Input['SolutionLandingZoneCatalogArgs'] catalog: This catalog stores all executable .ISO files for solution add-ons. There
               can be a single `catalog` element and the required field is `id`.
        :param pulumi.Input[Sequence[pulumi.Input['SolutionLandingZoneVdcArgs']]] vdcs: A single vdc block that defines landing VDC configuration
               
               <a id="vdc"></a>
        :param pulumi.Input[str] org: Destination Organization name for Solution Add-ons
        """
        pulumi.set(__self__, "catalog", catalog)
        pulumi.set(__self__, "vdcs", vdcs)
        if org is not None:
            pulumi.set(__self__, "org", org)

    @property
    @pulumi.getter
    def catalog(self) -> pulumi.Input['SolutionLandingZoneCatalogArgs']:
        """
        This catalog stores all executable .ISO files for solution add-ons. There
        can be a single `catalog` element and the required field is `id`.
        """
        return pulumi.get(self, "catalog")

    @catalog.setter
    def catalog(self, value: pulumi.Input['SolutionLandingZoneCatalogArgs']):
        pulumi.set(self, "catalog", value)

    @property
    @pulumi.getter
    def vdcs(self) -> pulumi.Input[Sequence[pulumi.Input['SolutionLandingZoneVdcArgs']]]:
        """
        A single vdc block that defines landing VDC configuration

        <a id="vdc"></a>
        """
        return pulumi.get(self, "vdcs")

    @vdcs.setter
    def vdcs(self, value: pulumi.Input[Sequence[pulumi.Input['SolutionLandingZoneVdcArgs']]]):
        pulumi.set(self, "vdcs", value)

    @property
    @pulumi.getter
    def org(self) -> Optional[pulumi.Input[str]]:
        """
        Destination Organization name for Solution Add-ons
        """
        return pulumi.get(self, "org")

    @org.setter
    def org(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org", value)


@pulumi.input_type
class _SolutionLandingZoneState:
    def __init__(__self__, *,
                 catalog: Optional[pulumi.Input['SolutionLandingZoneCatalogArgs']] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 vdcs: Optional[pulumi.Input[Sequence[pulumi.Input['SolutionLandingZoneVdcArgs']]]] = None):
        """
        Input properties used for looking up and filtering SolutionLandingZone resources.
        :param pulumi.Input['SolutionLandingZoneCatalogArgs'] catalog: This catalog stores all executable .ISO files for solution add-ons. There
               can be a single `catalog` element and the required field is `id`.
        :param pulumi.Input[str] org: Destination Organization name for Solution Add-ons
        :param pulumi.Input[str] state: reports the state of parent [Runtime Defined
               Entity](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/rde)
        :param pulumi.Input[Sequence[pulumi.Input['SolutionLandingZoneVdcArgs']]] vdcs: A single vdc block that defines landing VDC configuration
               
               <a id="vdc"></a>
        """
        if catalog is not None:
            pulumi.set(__self__, "catalog", catalog)
        if org is not None:
            pulumi.set(__self__, "org", org)
        if state is not None:
            pulumi.set(__self__, "state", state)
        if vdcs is not None:
            pulumi.set(__self__, "vdcs", vdcs)

    @property
    @pulumi.getter
    def catalog(self) -> Optional[pulumi.Input['SolutionLandingZoneCatalogArgs']]:
        """
        This catalog stores all executable .ISO files for solution add-ons. There
        can be a single `catalog` element and the required field is `id`.
        """
        return pulumi.get(self, "catalog")

    @catalog.setter
    def catalog(self, value: Optional[pulumi.Input['SolutionLandingZoneCatalogArgs']]):
        pulumi.set(self, "catalog", value)

    @property
    @pulumi.getter
    def org(self) -> Optional[pulumi.Input[str]]:
        """
        Destination Organization name for Solution Add-ons
        """
        return pulumi.get(self, "org")

    @org.setter
    def org(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        reports the state of parent [Runtime Defined
        Entity](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/rde)
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)

    @property
    @pulumi.getter
    def vdcs(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SolutionLandingZoneVdcArgs']]]]:
        """
        A single vdc block that defines landing VDC configuration

        <a id="vdc"></a>
        """
        return pulumi.get(self, "vdcs")

    @vdcs.setter
    def vdcs(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SolutionLandingZoneVdcArgs']]]]):
        pulumi.set(self, "vdcs", value)


class SolutionLandingZone(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 catalog: Optional[pulumi.Input[Union['SolutionLandingZoneCatalogArgs', 'SolutionLandingZoneCatalogArgsDict']]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 vdcs: Optional[pulumi.Input[Sequence[pulumi.Input[Union['SolutionLandingZoneVdcArgs', 'SolutionLandingZoneVdcArgsDict']]]]] = None,
                 __props__=None):
        """
        Create a SolutionLandingZone resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['SolutionLandingZoneCatalogArgs', 'SolutionLandingZoneCatalogArgsDict']] catalog: This catalog stores all executable .ISO files for solution add-ons. There
               can be a single `catalog` element and the required field is `id`.
        :param pulumi.Input[str] org: Destination Organization name for Solution Add-ons
        :param pulumi.Input[Sequence[pulumi.Input[Union['SolutionLandingZoneVdcArgs', 'SolutionLandingZoneVdcArgsDict']]]] vdcs: A single vdc block that defines landing VDC configuration
               
               <a id="vdc"></a>
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SolutionLandingZoneArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a SolutionLandingZone resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param SolutionLandingZoneArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SolutionLandingZoneArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 catalog: Optional[pulumi.Input[Union['SolutionLandingZoneCatalogArgs', 'SolutionLandingZoneCatalogArgsDict']]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 vdcs: Optional[pulumi.Input[Sequence[pulumi.Input[Union['SolutionLandingZoneVdcArgs', 'SolutionLandingZoneVdcArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SolutionLandingZoneArgs.__new__(SolutionLandingZoneArgs)

            if catalog is None and not opts.urn:
                raise TypeError("Missing required property 'catalog'")
            __props__.__dict__["catalog"] = catalog
            __props__.__dict__["org"] = org
            if vdcs is None and not opts.urn:
                raise TypeError("Missing required property 'vdcs'")
            __props__.__dict__["vdcs"] = vdcs
            __props__.__dict__["state"] = None
        super(SolutionLandingZone, __self__).__init__(
            'vcd:index/solutionLandingZone:SolutionLandingZone',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            catalog: Optional[pulumi.Input[Union['SolutionLandingZoneCatalogArgs', 'SolutionLandingZoneCatalogArgsDict']]] = None,
            org: Optional[pulumi.Input[str]] = None,
            state: Optional[pulumi.Input[str]] = None,
            vdcs: Optional[pulumi.Input[Sequence[pulumi.Input[Union['SolutionLandingZoneVdcArgs', 'SolutionLandingZoneVdcArgsDict']]]]] = None) -> 'SolutionLandingZone':
        """
        Get an existing SolutionLandingZone resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['SolutionLandingZoneCatalogArgs', 'SolutionLandingZoneCatalogArgsDict']] catalog: This catalog stores all executable .ISO files for solution add-ons. There
               can be a single `catalog` element and the required field is `id`.
        :param pulumi.Input[str] org: Destination Organization name for Solution Add-ons
        :param pulumi.Input[str] state: reports the state of parent [Runtime Defined
               Entity](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/rde)
        :param pulumi.Input[Sequence[pulumi.Input[Union['SolutionLandingZoneVdcArgs', 'SolutionLandingZoneVdcArgsDict']]]] vdcs: A single vdc block that defines landing VDC configuration
               
               <a id="vdc"></a>
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SolutionLandingZoneState.__new__(_SolutionLandingZoneState)

        __props__.__dict__["catalog"] = catalog
        __props__.__dict__["org"] = org
        __props__.__dict__["state"] = state
        __props__.__dict__["vdcs"] = vdcs
        return SolutionLandingZone(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def catalog(self) -> pulumi.Output['outputs.SolutionLandingZoneCatalog']:
        """
        This catalog stores all executable .ISO files for solution add-ons. There
        can be a single `catalog` element and the required field is `id`.
        """
        return pulumi.get(self, "catalog")

    @property
    @pulumi.getter
    def org(self) -> pulumi.Output[Optional[str]]:
        """
        Destination Organization name for Solution Add-ons
        """
        return pulumi.get(self, "org")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        reports the state of parent [Runtime Defined
        Entity](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/rde)
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def vdcs(self) -> pulumi.Output[Sequence['outputs.SolutionLandingZoneVdc']]:
        """
        A single vdc block that defines landing VDC configuration

        <a id="vdc"></a>
        """
        return pulumi.get(self, "vdcs")

