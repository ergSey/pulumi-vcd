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

__all__ = ['VappFirewallRulesArgs', 'VappFirewallRules']

@pulumi.input_type
class VappFirewallRulesArgs:
    def __init__(__self__, *,
                 default_action: pulumi.Input[str],
                 network_id: pulumi.Input[str],
                 vapp_id: pulumi.Input[str],
                 enabled: Optional[pulumi.Input[bool]] = None,
                 log_default_action: Optional[pulumi.Input[bool]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['VappFirewallRulesRuleArgs']]]] = None,
                 vdc: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a VappFirewallRules resource.
        :param pulumi.Input[str] default_action: Either 'allow' or 'drop'. Specifies what to do should none of the rules match.
        :param pulumi.Input[str] network_id: The identifier of [vApp network](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp_network).
        :param pulumi.Input[str] vapp_id: The identifier of [vApp](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp).
        :param pulumi.Input[bool] enabled: Enable or disable firewall. Default is `true`.
        :param pulumi.Input[bool] log_default_action: Flag to enable logging for default action. Default value is `false`.
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.
        :param pulumi.Input[Sequence[pulumi.Input['VappFirewallRulesRuleArgs']]] rules: Configures a firewall rule; see Rules below for details.
               
               <a id="rules"></a>
        :param pulumi.Input[str] vdc: The name of VDC to use, optional if defined at provider level.
        """
        pulumi.set(__self__, "default_action", default_action)
        pulumi.set(__self__, "network_id", network_id)
        pulumi.set(__self__, "vapp_id", vapp_id)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if log_default_action is not None:
            pulumi.set(__self__, "log_default_action", log_default_action)
        if org is not None:
            pulumi.set(__self__, "org", org)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)
        if vdc is not None:
            pulumi.set(__self__, "vdc", vdc)

    @property
    @pulumi.getter(name="defaultAction")
    def default_action(self) -> pulumi.Input[str]:
        """
        Either 'allow' or 'drop'. Specifies what to do should none of the rules match.
        """
        return pulumi.get(self, "default_action")

    @default_action.setter
    def default_action(self, value: pulumi.Input[str]):
        pulumi.set(self, "default_action", value)

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> pulumi.Input[str]:
        """
        The identifier of [vApp network](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp_network).
        """
        return pulumi.get(self, "network_id")

    @network_id.setter
    def network_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "network_id", value)

    @property
    @pulumi.getter(name="vappId")
    def vapp_id(self) -> pulumi.Input[str]:
        """
        The identifier of [vApp](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp).
        """
        return pulumi.get(self, "vapp_id")

    @vapp_id.setter
    def vapp_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "vapp_id", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable or disable firewall. Default is `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="logDefaultAction")
    def log_default_action(self) -> Optional[pulumi.Input[bool]]:
        """
        Flag to enable logging for default action. Default value is `false`.
        """
        return pulumi.get(self, "log_default_action")

    @log_default_action.setter
    def log_default_action(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "log_default_action", value)

    @property
    @pulumi.getter
    def org(self) -> Optional[pulumi.Input[str]]:
        """
        The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.
        """
        return pulumi.get(self, "org")

    @org.setter
    def org(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VappFirewallRulesRuleArgs']]]]:
        """
        Configures a firewall rule; see Rules below for details.

        <a id="rules"></a>
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VappFirewallRulesRuleArgs']]]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter
    def vdc(self) -> Optional[pulumi.Input[str]]:
        """
        The name of VDC to use, optional if defined at provider level.
        """
        return pulumi.get(self, "vdc")

    @vdc.setter
    def vdc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdc", value)


@pulumi.input_type
class _VappFirewallRulesState:
    def __init__(__self__, *,
                 default_action: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 log_default_action: Optional[pulumi.Input[bool]] = None,
                 network_id: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input['VappFirewallRulesRuleArgs']]]] = None,
                 vapp_id: Optional[pulumi.Input[str]] = None,
                 vdc: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering VappFirewallRules resources.
        :param pulumi.Input[str] default_action: Either 'allow' or 'drop'. Specifies what to do should none of the rules match.
        :param pulumi.Input[bool] enabled: Enable or disable firewall. Default is `true`.
        :param pulumi.Input[bool] log_default_action: Flag to enable logging for default action. Default value is `false`.
        :param pulumi.Input[str] network_id: The identifier of [vApp network](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp_network).
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.
        :param pulumi.Input[Sequence[pulumi.Input['VappFirewallRulesRuleArgs']]] rules: Configures a firewall rule; see Rules below for details.
               
               <a id="rules"></a>
        :param pulumi.Input[str] vapp_id: The identifier of [vApp](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp).
        :param pulumi.Input[str] vdc: The name of VDC to use, optional if defined at provider level.
        """
        if default_action is not None:
            pulumi.set(__self__, "default_action", default_action)
        if enabled is not None:
            pulumi.set(__self__, "enabled", enabled)
        if log_default_action is not None:
            pulumi.set(__self__, "log_default_action", log_default_action)
        if network_id is not None:
            pulumi.set(__self__, "network_id", network_id)
        if org is not None:
            pulumi.set(__self__, "org", org)
        if rules is not None:
            pulumi.set(__self__, "rules", rules)
        if vapp_id is not None:
            pulumi.set(__self__, "vapp_id", vapp_id)
        if vdc is not None:
            pulumi.set(__self__, "vdc", vdc)

    @property
    @pulumi.getter(name="defaultAction")
    def default_action(self) -> Optional[pulumi.Input[str]]:
        """
        Either 'allow' or 'drop'. Specifies what to do should none of the rules match.
        """
        return pulumi.get(self, "default_action")

    @default_action.setter
    def default_action(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "default_action", value)

    @property
    @pulumi.getter
    def enabled(self) -> Optional[pulumi.Input[bool]]:
        """
        Enable or disable firewall. Default is `true`.
        """
        return pulumi.get(self, "enabled")

    @enabled.setter
    def enabled(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enabled", value)

    @property
    @pulumi.getter(name="logDefaultAction")
    def log_default_action(self) -> Optional[pulumi.Input[bool]]:
        """
        Flag to enable logging for default action. Default value is `false`.
        """
        return pulumi.get(self, "log_default_action")

    @log_default_action.setter
    def log_default_action(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "log_default_action", value)

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> Optional[pulumi.Input[str]]:
        """
        The identifier of [vApp network](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp_network).
        """
        return pulumi.get(self, "network_id")

    @network_id.setter
    def network_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "network_id", value)

    @property
    @pulumi.getter
    def org(self) -> Optional[pulumi.Input[str]]:
        """
        The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.
        """
        return pulumi.get(self, "org")

    @org.setter
    def org(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "org", value)

    @property
    @pulumi.getter
    def rules(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['VappFirewallRulesRuleArgs']]]]:
        """
        Configures a firewall rule; see Rules below for details.

        <a id="rules"></a>
        """
        return pulumi.get(self, "rules")

    @rules.setter
    def rules(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['VappFirewallRulesRuleArgs']]]]):
        pulumi.set(self, "rules", value)

    @property
    @pulumi.getter(name="vappId")
    def vapp_id(self) -> Optional[pulumi.Input[str]]:
        """
        The identifier of [vApp](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp).
        """
        return pulumi.get(self, "vapp_id")

    @vapp_id.setter
    def vapp_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vapp_id", value)

    @property
    @pulumi.getter
    def vdc(self) -> Optional[pulumi.Input[str]]:
        """
        The name of VDC to use, optional if defined at provider level.
        """
        return pulumi.get(self, "vdc")

    @vdc.setter
    def vdc(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "vdc", value)


class VappFirewallRules(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_action: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 log_default_action: Optional[pulumi.Input[bool]] = None,
                 network_id: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[Union['VappFirewallRulesRuleArgs', 'VappFirewallRulesRuleArgsDict']]]]] = None,
                 vapp_id: Optional[pulumi.Input[str]] = None,
                 vdc: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Create a VappFirewallRules resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] default_action: Either 'allow' or 'drop'. Specifies what to do should none of the rules match.
        :param pulumi.Input[bool] enabled: Enable or disable firewall. Default is `true`.
        :param pulumi.Input[bool] log_default_action: Flag to enable logging for default action. Default value is `false`.
        :param pulumi.Input[str] network_id: The identifier of [vApp network](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp_network).
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.
        :param pulumi.Input[Sequence[pulumi.Input[Union['VappFirewallRulesRuleArgs', 'VappFirewallRulesRuleArgsDict']]]] rules: Configures a firewall rule; see Rules below for details.
               
               <a id="rules"></a>
        :param pulumi.Input[str] vapp_id: The identifier of [vApp](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp).
        :param pulumi.Input[str] vdc: The name of VDC to use, optional if defined at provider level.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: VappFirewallRulesArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a VappFirewallRules resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param VappFirewallRulesArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(VappFirewallRulesArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 default_action: Optional[pulumi.Input[str]] = None,
                 enabled: Optional[pulumi.Input[bool]] = None,
                 log_default_action: Optional[pulumi.Input[bool]] = None,
                 network_id: Optional[pulumi.Input[str]] = None,
                 org: Optional[pulumi.Input[str]] = None,
                 rules: Optional[pulumi.Input[Sequence[pulumi.Input[Union['VappFirewallRulesRuleArgs', 'VappFirewallRulesRuleArgsDict']]]]] = None,
                 vapp_id: Optional[pulumi.Input[str]] = None,
                 vdc: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = VappFirewallRulesArgs.__new__(VappFirewallRulesArgs)

            if default_action is None and not opts.urn:
                raise TypeError("Missing required property 'default_action'")
            __props__.__dict__["default_action"] = default_action
            __props__.__dict__["enabled"] = enabled
            __props__.__dict__["log_default_action"] = log_default_action
            if network_id is None and not opts.urn:
                raise TypeError("Missing required property 'network_id'")
            __props__.__dict__["network_id"] = network_id
            __props__.__dict__["org"] = org
            __props__.__dict__["rules"] = rules
            if vapp_id is None and not opts.urn:
                raise TypeError("Missing required property 'vapp_id'")
            __props__.__dict__["vapp_id"] = vapp_id
            __props__.__dict__["vdc"] = vdc
        super(VappFirewallRules, __self__).__init__(
            'vcd:index/vappFirewallRules:VappFirewallRules',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            default_action: Optional[pulumi.Input[str]] = None,
            enabled: Optional[pulumi.Input[bool]] = None,
            log_default_action: Optional[pulumi.Input[bool]] = None,
            network_id: Optional[pulumi.Input[str]] = None,
            org: Optional[pulumi.Input[str]] = None,
            rules: Optional[pulumi.Input[Sequence[pulumi.Input[Union['VappFirewallRulesRuleArgs', 'VappFirewallRulesRuleArgsDict']]]]] = None,
            vapp_id: Optional[pulumi.Input[str]] = None,
            vdc: Optional[pulumi.Input[str]] = None) -> 'VappFirewallRules':
        """
        Get an existing VappFirewallRules resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] default_action: Either 'allow' or 'drop'. Specifies what to do should none of the rules match.
        :param pulumi.Input[bool] enabled: Enable or disable firewall. Default is `true`.
        :param pulumi.Input[bool] log_default_action: Flag to enable logging for default action. Default value is `false`.
        :param pulumi.Input[str] network_id: The identifier of [vApp network](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp_network).
        :param pulumi.Input[str] org: The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.
        :param pulumi.Input[Sequence[pulumi.Input[Union['VappFirewallRulesRuleArgs', 'VappFirewallRulesRuleArgsDict']]]] rules: Configures a firewall rule; see Rules below for details.
               
               <a id="rules"></a>
        :param pulumi.Input[str] vapp_id: The identifier of [vApp](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp).
        :param pulumi.Input[str] vdc: The name of VDC to use, optional if defined at provider level.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _VappFirewallRulesState.__new__(_VappFirewallRulesState)

        __props__.__dict__["default_action"] = default_action
        __props__.__dict__["enabled"] = enabled
        __props__.__dict__["log_default_action"] = log_default_action
        __props__.__dict__["network_id"] = network_id
        __props__.__dict__["org"] = org
        __props__.__dict__["rules"] = rules
        __props__.__dict__["vapp_id"] = vapp_id
        __props__.__dict__["vdc"] = vdc
        return VappFirewallRules(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="defaultAction")
    def default_action(self) -> pulumi.Output[str]:
        """
        Either 'allow' or 'drop'. Specifies what to do should none of the rules match.
        """
        return pulumi.get(self, "default_action")

    @property
    @pulumi.getter
    def enabled(self) -> pulumi.Output[Optional[bool]]:
        """
        Enable or disable firewall. Default is `true`.
        """
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter(name="logDefaultAction")
    def log_default_action(self) -> pulumi.Output[Optional[bool]]:
        """
        Flag to enable logging for default action. Default value is `false`.
        """
        return pulumi.get(self, "log_default_action")

    @property
    @pulumi.getter(name="networkId")
    def network_id(self) -> pulumi.Output[str]:
        """
        The identifier of [vApp network](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp_network).
        """
        return pulumi.get(self, "network_id")

    @property
    @pulumi.getter
    def org(self) -> pulumi.Output[Optional[str]]:
        """
        The name of organization to use, optional if defined at provider level. Useful when connected as sysadmin working across different organisations.
        """
        return pulumi.get(self, "org")

    @property
    @pulumi.getter
    def rules(self) -> pulumi.Output[Optional[Sequence['outputs.VappFirewallRulesRule']]]:
        """
        Configures a firewall rule; see Rules below for details.

        <a id="rules"></a>
        """
        return pulumi.get(self, "rules")

    @property
    @pulumi.getter(name="vappId")
    def vapp_id(self) -> pulumi.Output[str]:
        """
        The identifier of [vApp](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/vapp).
        """
        return pulumi.get(self, "vapp_id")

    @property
    @pulumi.getter
    def vdc(self) -> pulumi.Output[Optional[str]]:
        """
        The name of VDC to use, optional if defined at provider level.
        """
        return pulumi.get(self, "vdc")

