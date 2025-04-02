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
    'GetNsxtIpsecVpnTunnelResult',
    'AwaitableGetNsxtIpsecVpnTunnelResult',
    'get_nsxt_ipsec_vpn_tunnel',
    'get_nsxt_ipsec_vpn_tunnel_output',
]

@pulumi.output_type
class GetNsxtIpsecVpnTunnelResult:
    """
    A collection of values returned by getNsxtIpsecVpnTunnel.
    """
    def __init__(__self__, authentication_mode=None, ca_certificate_id=None, certificate_id=None, description=None, edge_gateway_id=None, enabled=None, id=None, ike_fail_reason=None, ike_service_status=None, local_ip_address=None, local_networks=None, logging=None, name=None, org=None, pre_shared_key=None, remote_id=None, remote_ip_address=None, remote_networks=None, security_profile=None, security_profile_customizations=None, status=None, vdc=None):
        if authentication_mode and not isinstance(authentication_mode, str):
            raise TypeError("Expected argument 'authentication_mode' to be a str")
        pulumi.set(__self__, "authentication_mode", authentication_mode)
        if ca_certificate_id and not isinstance(ca_certificate_id, str):
            raise TypeError("Expected argument 'ca_certificate_id' to be a str")
        pulumi.set(__self__, "ca_certificate_id", ca_certificate_id)
        if certificate_id and not isinstance(certificate_id, str):
            raise TypeError("Expected argument 'certificate_id' to be a str")
        pulumi.set(__self__, "certificate_id", certificate_id)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if edge_gateway_id and not isinstance(edge_gateway_id, str):
            raise TypeError("Expected argument 'edge_gateway_id' to be a str")
        pulumi.set(__self__, "edge_gateway_id", edge_gateway_id)
        if enabled and not isinstance(enabled, bool):
            raise TypeError("Expected argument 'enabled' to be a bool")
        pulumi.set(__self__, "enabled", enabled)
        if id and not isinstance(id, str):
            raise TypeError("Expected argument 'id' to be a str")
        pulumi.set(__self__, "id", id)
        if ike_fail_reason and not isinstance(ike_fail_reason, str):
            raise TypeError("Expected argument 'ike_fail_reason' to be a str")
        pulumi.set(__self__, "ike_fail_reason", ike_fail_reason)
        if ike_service_status and not isinstance(ike_service_status, str):
            raise TypeError("Expected argument 'ike_service_status' to be a str")
        pulumi.set(__self__, "ike_service_status", ike_service_status)
        if local_ip_address and not isinstance(local_ip_address, str):
            raise TypeError("Expected argument 'local_ip_address' to be a str")
        pulumi.set(__self__, "local_ip_address", local_ip_address)
        if local_networks and not isinstance(local_networks, list):
            raise TypeError("Expected argument 'local_networks' to be a list")
        pulumi.set(__self__, "local_networks", local_networks)
        if logging and not isinstance(logging, bool):
            raise TypeError("Expected argument 'logging' to be a bool")
        pulumi.set(__self__, "logging", logging)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if org and not isinstance(org, str):
            raise TypeError("Expected argument 'org' to be a str")
        pulumi.set(__self__, "org", org)
        if pre_shared_key and not isinstance(pre_shared_key, str):
            raise TypeError("Expected argument 'pre_shared_key' to be a str")
        pulumi.set(__self__, "pre_shared_key", pre_shared_key)
        if remote_id and not isinstance(remote_id, str):
            raise TypeError("Expected argument 'remote_id' to be a str")
        pulumi.set(__self__, "remote_id", remote_id)
        if remote_ip_address and not isinstance(remote_ip_address, str):
            raise TypeError("Expected argument 'remote_ip_address' to be a str")
        pulumi.set(__self__, "remote_ip_address", remote_ip_address)
        if remote_networks and not isinstance(remote_networks, list):
            raise TypeError("Expected argument 'remote_networks' to be a list")
        pulumi.set(__self__, "remote_networks", remote_networks)
        if security_profile and not isinstance(security_profile, str):
            raise TypeError("Expected argument 'security_profile' to be a str")
        pulumi.set(__self__, "security_profile", security_profile)
        if security_profile_customizations and not isinstance(security_profile_customizations, list):
            raise TypeError("Expected argument 'security_profile_customizations' to be a list")
        pulumi.set(__self__, "security_profile_customizations", security_profile_customizations)
        if status and not isinstance(status, str):
            raise TypeError("Expected argument 'status' to be a str")
        pulumi.set(__self__, "status", status)
        if vdc and not isinstance(vdc, str):
            raise TypeError("Expected argument 'vdc' to be a str")
        pulumi.set(__self__, "vdc", vdc)

    @property
    @pulumi.getter(name="authenticationMode")
    def authentication_mode(self) -> str:
        return pulumi.get(self, "authentication_mode")

    @property
    @pulumi.getter(name="caCertificateId")
    def ca_certificate_id(self) -> str:
        return pulumi.get(self, "ca_certificate_id")

    @property
    @pulumi.getter(name="certificateId")
    def certificate_id(self) -> str:
        return pulumi.get(self, "certificate_id")

    @property
    @pulumi.getter
    def description(self) -> str:
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="edgeGatewayId")
    def edge_gateway_id(self) -> str:
        return pulumi.get(self, "edge_gateway_id")

    @property
    @pulumi.getter
    def enabled(self) -> bool:
        return pulumi.get(self, "enabled")

    @property
    @pulumi.getter
    def id(self) -> str:
        """
        The provider-assigned unique ID for this managed resource.
        """
        return pulumi.get(self, "id")

    @property
    @pulumi.getter(name="ikeFailReason")
    def ike_fail_reason(self) -> str:
        return pulumi.get(self, "ike_fail_reason")

    @property
    @pulumi.getter(name="ikeServiceStatus")
    def ike_service_status(self) -> str:
        return pulumi.get(self, "ike_service_status")

    @property
    @pulumi.getter(name="localIpAddress")
    def local_ip_address(self) -> str:
        return pulumi.get(self, "local_ip_address")

    @property
    @pulumi.getter(name="localNetworks")
    def local_networks(self) -> Sequence[str]:
        return pulumi.get(self, "local_networks")

    @property
    @pulumi.getter
    def logging(self) -> bool:
        return pulumi.get(self, "logging")

    @property
    @pulumi.getter
    def name(self) -> str:
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def org(self) -> Optional[str]:
        return pulumi.get(self, "org")

    @property
    @pulumi.getter(name="preSharedKey")
    def pre_shared_key(self) -> str:
        return pulumi.get(self, "pre_shared_key")

    @property
    @pulumi.getter(name="remoteId")
    def remote_id(self) -> str:
        return pulumi.get(self, "remote_id")

    @property
    @pulumi.getter(name="remoteIpAddress")
    def remote_ip_address(self) -> str:
        return pulumi.get(self, "remote_ip_address")

    @property
    @pulumi.getter(name="remoteNetworks")
    def remote_networks(self) -> Sequence[str]:
        return pulumi.get(self, "remote_networks")

    @property
    @pulumi.getter(name="securityProfile")
    def security_profile(self) -> str:
        return pulumi.get(self, "security_profile")

    @property
    @pulumi.getter(name="securityProfileCustomizations")
    def security_profile_customizations(self) -> Sequence['outputs.GetNsxtIpsecVpnTunnelSecurityProfileCustomizationResult']:
        return pulumi.get(self, "security_profile_customizations")

    @property
    @pulumi.getter
    def status(self) -> str:
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    @_utilities.deprecated("""Edge Gateway will be looked up based on 'edge_gateway_id' field""")
    def vdc(self) -> Optional[str]:
        return pulumi.get(self, "vdc")


class AwaitableGetNsxtIpsecVpnTunnelResult(GetNsxtIpsecVpnTunnelResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNsxtIpsecVpnTunnelResult(
            authentication_mode=self.authentication_mode,
            ca_certificate_id=self.ca_certificate_id,
            certificate_id=self.certificate_id,
            description=self.description,
            edge_gateway_id=self.edge_gateway_id,
            enabled=self.enabled,
            id=self.id,
            ike_fail_reason=self.ike_fail_reason,
            ike_service_status=self.ike_service_status,
            local_ip_address=self.local_ip_address,
            local_networks=self.local_networks,
            logging=self.logging,
            name=self.name,
            org=self.org,
            pre_shared_key=self.pre_shared_key,
            remote_id=self.remote_id,
            remote_ip_address=self.remote_ip_address,
            remote_networks=self.remote_networks,
            security_profile=self.security_profile,
            security_profile_customizations=self.security_profile_customizations,
            status=self.status,
            vdc=self.vdc)


def get_nsxt_ipsec_vpn_tunnel(edge_gateway_id: Optional[str] = None,
                              name: Optional[str] = None,
                              org: Optional[str] = None,
                              vdc: Optional[str] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNsxtIpsecVpnTunnelResult:
    """
    Supported in provider *v3.3+* and VCD 10.1+ with NSX-T backed VDCs.

    Provides a data source to read NSX-T IPsec VPN Tunnel. You can configure site-to-site connectivity between an NSX-T Data
    Center Edge Gateway and remote sites. The remote sites must use NSX-T Data Center, have third-party hardware routers,
    or VPN gateways that support IPSec.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vcd as vcd

    tunnel1 = vcd.get_nsxt_ipsec_vpn_tunnel(org="my-org",
        edge_gateway_id=existing["id"],
        name="tunnel-1")
    ```


    :param str edge_gateway_id: The ID of the Edge Gateway (NSX-T only). Can be looked up using `NsxtEdgegateway`
           data source
    :param str name: Name of existing IPsec VPN Tunnel
    :param str org: The name of organization to use, optional if defined at provider level. Useful
           when connected as sysadmin working across different organisations.
    """
    __args__ = dict()
    __args__['edgeGatewayId'] = edge_gateway_id
    __args__['name'] = name
    __args__['org'] = org
    __args__['vdc'] = vdc
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('vcd:index/getNsxtIpsecVpnTunnel:getNsxtIpsecVpnTunnel', __args__, opts=opts, typ=GetNsxtIpsecVpnTunnelResult).value

    return AwaitableGetNsxtIpsecVpnTunnelResult(
        authentication_mode=pulumi.get(__ret__, 'authentication_mode'),
        ca_certificate_id=pulumi.get(__ret__, 'ca_certificate_id'),
        certificate_id=pulumi.get(__ret__, 'certificate_id'),
        description=pulumi.get(__ret__, 'description'),
        edge_gateway_id=pulumi.get(__ret__, 'edge_gateway_id'),
        enabled=pulumi.get(__ret__, 'enabled'),
        id=pulumi.get(__ret__, 'id'),
        ike_fail_reason=pulumi.get(__ret__, 'ike_fail_reason'),
        ike_service_status=pulumi.get(__ret__, 'ike_service_status'),
        local_ip_address=pulumi.get(__ret__, 'local_ip_address'),
        local_networks=pulumi.get(__ret__, 'local_networks'),
        logging=pulumi.get(__ret__, 'logging'),
        name=pulumi.get(__ret__, 'name'),
        org=pulumi.get(__ret__, 'org'),
        pre_shared_key=pulumi.get(__ret__, 'pre_shared_key'),
        remote_id=pulumi.get(__ret__, 'remote_id'),
        remote_ip_address=pulumi.get(__ret__, 'remote_ip_address'),
        remote_networks=pulumi.get(__ret__, 'remote_networks'),
        security_profile=pulumi.get(__ret__, 'security_profile'),
        security_profile_customizations=pulumi.get(__ret__, 'security_profile_customizations'),
        status=pulumi.get(__ret__, 'status'),
        vdc=pulumi.get(__ret__, 'vdc'))
def get_nsxt_ipsec_vpn_tunnel_output(edge_gateway_id: Optional[pulumi.Input[str]] = None,
                                     name: Optional[pulumi.Input[str]] = None,
                                     org: Optional[pulumi.Input[Optional[str]]] = None,
                                     vdc: Optional[pulumi.Input[Optional[str]]] = None,
                                     opts: Optional[Union[pulumi.InvokeOptions, pulumi.InvokeOutputOptions]] = None) -> pulumi.Output[GetNsxtIpsecVpnTunnelResult]:
    """
    Supported in provider *v3.3+* and VCD 10.1+ with NSX-T backed VDCs.

    Provides a data source to read NSX-T IPsec VPN Tunnel. You can configure site-to-site connectivity between an NSX-T Data
    Center Edge Gateway and remote sites. The remote sites must use NSX-T Data Center, have third-party hardware routers,
    or VPN gateways that support IPSec.

    ## Example Usage

    ```python
    import pulumi
    import pulumi_vcd as vcd

    tunnel1 = vcd.get_nsxt_ipsec_vpn_tunnel(org="my-org",
        edge_gateway_id=existing["id"],
        name="tunnel-1")
    ```


    :param str edge_gateway_id: The ID of the Edge Gateway (NSX-T only). Can be looked up using `NsxtEdgegateway`
           data source
    :param str name: Name of existing IPsec VPN Tunnel
    :param str org: The name of organization to use, optional if defined at provider level. Useful
           when connected as sysadmin working across different organisations.
    """
    __args__ = dict()
    __args__['edgeGatewayId'] = edge_gateway_id
    __args__['name'] = name
    __args__['org'] = org
    __args__['vdc'] = vdc
    opts = pulumi.InvokeOutputOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke_output('vcd:index/getNsxtIpsecVpnTunnel:getNsxtIpsecVpnTunnel', __args__, opts=opts, typ=GetNsxtIpsecVpnTunnelResult)
    return __ret__.apply(lambda __response__: GetNsxtIpsecVpnTunnelResult(
        authentication_mode=pulumi.get(__response__, 'authentication_mode'),
        ca_certificate_id=pulumi.get(__response__, 'ca_certificate_id'),
        certificate_id=pulumi.get(__response__, 'certificate_id'),
        description=pulumi.get(__response__, 'description'),
        edge_gateway_id=pulumi.get(__response__, 'edge_gateway_id'),
        enabled=pulumi.get(__response__, 'enabled'),
        id=pulumi.get(__response__, 'id'),
        ike_fail_reason=pulumi.get(__response__, 'ike_fail_reason'),
        ike_service_status=pulumi.get(__response__, 'ike_service_status'),
        local_ip_address=pulumi.get(__response__, 'local_ip_address'),
        local_networks=pulumi.get(__response__, 'local_networks'),
        logging=pulumi.get(__response__, 'logging'),
        name=pulumi.get(__response__, 'name'),
        org=pulumi.get(__response__, 'org'),
        pre_shared_key=pulumi.get(__response__, 'pre_shared_key'),
        remote_id=pulumi.get(__response__, 'remote_id'),
        remote_ip_address=pulumi.get(__response__, 'remote_ip_address'),
        remote_networks=pulumi.get(__response__, 'remote_networks'),
        security_profile=pulumi.get(__response__, 'security_profile'),
        security_profile_customizations=pulumi.get(__response__, 'security_profile_customizations'),
        status=pulumi.get(__response__, 'status'),
        vdc=pulumi.get(__response__, 'vdc')))
