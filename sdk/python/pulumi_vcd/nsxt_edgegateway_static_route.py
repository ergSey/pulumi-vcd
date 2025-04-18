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

__all__ = ['NsxtEdgegatewayStaticRouteArgs', 'NsxtEdgegatewayStaticRoute']

@pulumi.input_type
class NsxtEdgegatewayStaticRouteArgs:
    def __init__(__self__, *,
                 edge_gateway_id: pulumi.Input[str],
                 network_cidr: pulumi.Input[str],
                 next_hops: pulumi.Input[Sequence[pulumi.Input['NsxtEdgegatewayStaticRouteNextHopArgs']]],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a NsxtEdgegatewayStaticRoute resource.
        :param pulumi.Input[str] edge_gateway_id: NSX-T Edge Gateway ID
        :param pulumi.Input[str] network_cidr: Specifies network prefix in CIDR format. Both IPv4 and IPv6 formats
               are supported.
        :param pulumi.Input[Sequence[pulumi.Input['NsxtEdgegatewayStaticRouteNextHopArgs']]] next_hops: A set of next hops to use within the static route. At least one is
               required. See Next Hop for definition structure.
               
               <a id="next-hop"></a>
        :param pulumi.Input[str] description: Description for NSX-T Edge Gateway Static Route
        :param pulumi.Input[str] name: Name for NSX-T Edge Gateway Static Route
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
               different organizations
        """
        pulumi.set(__self__, "edge_gateway_id", edge_gateway_id)
        pulumi.set(__self__, "network_cidr", network_cidr)
        pulumi.set(__self__, "next_hops", next_hops)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if org is not None:
            pulumi.set(__self__, "org", org)

    @property
    @pulumi.getter(name="edgeGatewayId")
    def edge_gateway_id(self) -> pulumi.Input[str]:
        """
        NSX-T Edge Gateway ID
        """
        return pulumi.get(self, "edge_gateway_id")

    @edge_gateway_id.setter
    def edge_gateway_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "edge_gateway_id", value)

    @property
    @pulumi.getter(name="networkCidr")
    def network_cidr(self) -> pulumi.Input[str]:
        """
        Specifies network prefix in CIDR format. Both IPv4 and IPv6 formats
        are supported.
        """
        return pulumi.get(self, "network_cidr")

    @network_cidr.setter
    def network_cidr(self, value: pulumi.Input[str]):
        pulumi.set(self, "network_cidr", value)

    @property
    @pulumi.getter(name="nextHops")
    def next_hops(self) -> pulumi.Input[Sequence[pulumi.Input['NsxtEdgegatewayStaticRouteNextHopArgs']]]:
        """
        A set of next hops to use within the static route. At least one is
        required. See Next Hop for definition structure.

        <a id="next-hop"></a>
        """
        return pulumi.get(self, "next_hops")

    @next_hops.setter
    def next_hops(self, value: pulumi.Input[Sequence[pulumi.Input['NsxtEdgegatewayStaticRouteNextHopArgs']]]):
        pulumi.set(self, "next_hops", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description for NSX-T Edge Gateway Static Route
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name for NSX-T Edge Gateway Static Route
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def org(self) -> Optional[pulumi.Input[str]]:
        """
        The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        different organizations
        """
        return pulumi.get(self, "org")

    @org.setter
    def org(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org", value)


@pulumi.input_type
class _NsxtEdgegatewayStaticRouteState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 edge_gateway_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_cidr: Optional[pulumi.Input[str]] = None,
                 next_hops: Optional[pulumi.Input[Sequence[pulumi.Input['NsxtEdgegatewayStaticRouteNextHopArgs']]]] = None,
                 org: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering NsxtEdgegatewayStaticRoute resources.
        :param pulumi.Input[str] description: Description for NSX-T Edge Gateway Static Route
        :param pulumi.Input[str] edge_gateway_id: NSX-T Edge Gateway ID
        :param pulumi.Input[str] name: Name for NSX-T Edge Gateway Static Route
        :param pulumi.Input[str] network_cidr: Specifies network prefix in CIDR format. Both IPv4 and IPv6 formats
               are supported.
        :param pulumi.Input[Sequence[pulumi.Input['NsxtEdgegatewayStaticRouteNextHopArgs']]] next_hops: A set of next hops to use within the static route. At least one is
               required. See Next Hop for definition structure.
               
               <a id="next-hop"></a>
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
               different organizations
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if edge_gateway_id is not None:
            pulumi.set(__self__, "edge_gateway_id", edge_gateway_id)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if network_cidr is not None:
            pulumi.set(__self__, "network_cidr", network_cidr)
        if next_hops is not None:
            pulumi.set(__self__, "next_hops", next_hops)
        if org is not None:
            pulumi.set(__self__, "org", org)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Description for NSX-T Edge Gateway Static Route
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="edgeGatewayId")
    def edge_gateway_id(self) -> Optional[pulumi.Input[str]]:
        """
        NSX-T Edge Gateway ID
        """
        return pulumi.get(self, "edge_gateway_id")

    @edge_gateway_id.setter
    def edge_gateway_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "edge_gateway_id", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Name for NSX-T Edge Gateway Static Route
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="networkCidr")
    def network_cidr(self) -> Optional[pulumi.Input[str]]:
        """
        Specifies network prefix in CIDR format. Both IPv4 and IPv6 formats
        are supported.
        """
        return pulumi.get(self, "network_cidr")

    @network_cidr.setter
    def network_cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_cidr", value)

    @property
    @pulumi.getter(name="nextHops")
    def next_hops(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['NsxtEdgegatewayStaticRouteNextHopArgs']]]]:
        """
        A set of next hops to use within the static route. At least one is
        required. See Next Hop for definition structure.

        <a id="next-hop"></a>
        """
        return pulumi.get(self, "next_hops")

    @next_hops.setter
    def next_hops(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['NsxtEdgegatewayStaticRouteNextHopArgs']]]]):
        pulumi.set(self, "next_hops", value)

    @property
    @pulumi.getter
    def org(self) -> Optional[pulumi.Input[str]]:
        """
        The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        different organizations
        """
        return pulumi.get(self, "org")

    @org.setter
    def org(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org", value)


class NsxtEdgegatewayStaticRoute(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 edge_gateway_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_cidr: Optional[pulumi.Input[str]] = None,
                 next_hops: Optional[pulumi.Input[Sequence[pulumi.Input[Union['NsxtEdgegatewayStaticRouteNextHopArgs', 'NsxtEdgegatewayStaticRouteNextHopArgsDict']]]]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a NsxtEdgegatewayStaticRoute resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description for NSX-T Edge Gateway Static Route
        :param pulumi.Input[str] edge_gateway_id: NSX-T Edge Gateway ID
        :param pulumi.Input[str] name: Name for NSX-T Edge Gateway Static Route
        :param pulumi.Input[str] network_cidr: Specifies network prefix in CIDR format. Both IPv4 and IPv6 formats
               are supported.
        :param pulumi.Input[Sequence[pulumi.Input[Union['NsxtEdgegatewayStaticRouteNextHopArgs', 'NsxtEdgegatewayStaticRouteNextHopArgsDict']]]] next_hops: A set of next hops to use within the static route. At least one is
               required. See Next Hop for definition structure.
               
               <a id="next-hop"></a>
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
               different organizations
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: NsxtEdgegatewayStaticRouteArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a NsxtEdgegatewayStaticRoute resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param NsxtEdgegatewayStaticRouteArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(NsxtEdgegatewayStaticRouteArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 edge_gateway_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 network_cidr: Optional[pulumi.Input[str]] = None,
                 next_hops: Optional[pulumi.Input[Sequence[pulumi.Input[Union['NsxtEdgegatewayStaticRouteNextHopArgs', 'NsxtEdgegatewayStaticRouteNextHopArgsDict']]]]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = NsxtEdgegatewayStaticRouteArgs.__new__(NsxtEdgegatewayStaticRouteArgs)

            __props__.__dict__["description"] = description
            if edge_gateway_id is None and not opts.urn:
                raise TypeError("Missing required property 'edge_gateway_id'")
            __props__.__dict__["edge_gateway_id"] = edge_gateway_id
            __props__.__dict__["name"] = name
            if network_cidr is None and not opts.urn:
                raise TypeError("Missing required property 'network_cidr'")
            __props__.__dict__["network_cidr"] = network_cidr
            if next_hops is None and not opts.urn:
                raise TypeError("Missing required property 'next_hops'")
            __props__.__dict__["next_hops"] = next_hops
            __props__.__dict__["org"] = org
        super(NsxtEdgegatewayStaticRoute, __self__).__init__(
            'vcd:index/nsxtEdgegatewayStaticRoute:NsxtEdgegatewayStaticRoute',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            edge_gateway_id: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            network_cidr: Optional[pulumi.Input[str]] = None,
            next_hops: Optional[pulumi.Input[Sequence[pulumi.Input[Union['NsxtEdgegatewayStaticRouteNextHopArgs', 'NsxtEdgegatewayStaticRouteNextHopArgsDict']]]]] = None,
            org: Optional[pulumi.Input[str]] = None) -> 'NsxtEdgegatewayStaticRoute':
        """
        Get an existing NsxtEdgegatewayStaticRoute resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Description for NSX-T Edge Gateway Static Route
        :param pulumi.Input[str] edge_gateway_id: NSX-T Edge Gateway ID
        :param pulumi.Input[str] name: Name for NSX-T Edge Gateway Static Route
        :param pulumi.Input[str] network_cidr: Specifies network prefix in CIDR format. Both IPv4 and IPv6 formats
               are supported.
        :param pulumi.Input[Sequence[pulumi.Input[Union['NsxtEdgegatewayStaticRouteNextHopArgs', 'NsxtEdgegatewayStaticRouteNextHopArgsDict']]]] next_hops: A set of next hops to use within the static route. At least one is
               required. See Next Hop for definition structure.
               
               <a id="next-hop"></a>
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
               different organizations
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _NsxtEdgegatewayStaticRouteState.__new__(_NsxtEdgegatewayStaticRouteState)

        __props__.__dict__["description"] = description
        __props__.__dict__["edge_gateway_id"] = edge_gateway_id
        __props__.__dict__["name"] = name
        __props__.__dict__["network_cidr"] = network_cidr
        __props__.__dict__["next_hops"] = next_hops
        __props__.__dict__["org"] = org
        return NsxtEdgegatewayStaticRoute(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        Description for NSX-T Edge Gateway Static Route
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="edgeGatewayId")
    def edge_gateway_id(self) -> pulumi.Output[str]:
        """
        NSX-T Edge Gateway ID
        """
        return pulumi.get(self, "edge_gateway_id")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Name for NSX-T Edge Gateway Static Route
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkCidr")
    def network_cidr(self) -> pulumi.Output[str]:
        """
        Specifies network prefix in CIDR format. Both IPv4 and IPv6 formats
        are supported.
        """
        return pulumi.get(self, "network_cidr")

    @property
    @pulumi.getter(name="nextHops")
    def next_hops(self) -> pulumi.Output[Sequence['outputs.NsxtEdgegatewayStaticRouteNextHop']]:
        """
        A set of next hops to use within the static route. At least one is
        required. See Next Hop for definition structure.

        <a id="next-hop"></a>
        """
        return pulumi.get(self, "next_hops")

    @property
    @pulumi.getter
    def org(self) -> pulumi.Output[Optional[str]]:
        """
        The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across
        different organizations
        """
        return pulumi.get(self, "org")

