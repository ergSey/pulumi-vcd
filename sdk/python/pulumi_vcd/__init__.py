# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

from . import _utilities
import typing
# Export this package's modules as members:
from .api_filter import *
from .api_token import *
from .catalog import *
from .catalog_access_control import *
from .catalog_item import *
from .catalog_media import *
from .catalog_vapp_template import *
from .cloned_vapp import *
from .cse_kubernetes_cluster import *
from .dse_registry_configuration import *
from .dse_solution_publish import *
from .edgegateway import *
from .edgegateway_settings import *
from .edgegateway_vpn import *
from .external_endpoint import *
from .external_network import *
from .external_network_v2 import *
from .get_api_filter import *
from .get_catalog import *
from .get_catalog_access_control import *
from .get_catalog_item import *
from .get_catalog_media import *
from .get_catalog_vapp_template import *
from .get_cse_kubernetes_cluster import *
from .get_dse_registry_configuration import *
from .get_dse_solution_publish import *
from .get_edgegateway import *
from .get_external_endpoint import *
from .get_external_network import *
from .get_external_network_v2 import *
from .get_global_role import *
from .get_independent_disk import *
from .get_ip_space import *
from .get_ip_space_custom_quota import *
from .get_ip_space_ip_allocation import *
from .get_ip_space_uplink import *
from .get_lb_app_profile import *
from .get_lb_app_rule import *
from .get_lb_server_pool import *
from .get_lb_service_monitor import *
from .get_lb_virtual_server import *
from .get_library_certificate import *
from .get_multisite_org_association import *
from .get_multisite_org_data import *
from .get_multisite_site import *
from .get_multisite_site_association import *
from .get_multisite_site_data import *
from .get_network_direct import *
from .get_network_isolated import *
from .get_network_isolated_v2 import *
from .get_network_pool import *
from .get_network_routed import *
from .get_network_routed_v2 import *
from .get_nsxt_alb_cloud import *
from .get_nsxt_alb_controller import *
from .get_nsxt_alb_edgegateway_service_engine_group import *
from .get_nsxt_alb_importable_cloud import *
from .get_nsxt_alb_pool import *
from .get_nsxt_alb_service_engine_group import *
from .get_nsxt_alb_settings import *
from .get_nsxt_alb_virtual_service import *
from .get_nsxt_alb_virtual_service_http_req_rules import *
from .get_nsxt_alb_virtual_service_http_resp_rules import *
from .get_nsxt_alb_virtual_service_http_sec_rules import *
from .get_nsxt_app_port_profile import *
from .get_nsxt_distributed_firewall import *
from .get_nsxt_distributed_firewall_rule import *
from .get_nsxt_dynamic_security_group import *
from .get_nsxt_edge_cluster import *
from .get_nsxt_edgegateway import *
from .get_nsxt_edgegateway_bgp_configuration import *
from .get_nsxt_edgegateway_bgp_ip_prefix_list import *
from .get_nsxt_edgegateway_bgp_neighbor import *
from .get_nsxt_edgegateway_dhcp_forwarding import *
from .get_nsxt_edgegateway_dhcpv6 import *
from .get_nsxt_edgegateway_dns import *
from .get_nsxt_edgegateway_l2_vpn_tunnel import *
from .get_nsxt_edgegateway_qos_profile import *
from .get_nsxt_edgegateway_rate_limiting import *
from .get_nsxt_edgegateway_static_route import *
from .get_nsxt_firewall import *
from .get_nsxt_global_default_segment_profile_template import *
from .get_nsxt_ip_set import *
from .get_nsxt_ipsec_vpn_tunnel import *
from .get_nsxt_manager import *
from .get_nsxt_nat_rule import *
from .get_nsxt_network_context_profile import *
from .get_nsxt_network_dhcp import *
from .get_nsxt_network_dhcp_binding import *
from .get_nsxt_network_imported import *
from .get_nsxt_network_segment_profile import *
from .get_nsxt_route_advertisement import *
from .get_nsxt_security_group import *
from .get_nsxt_segment_ip_discovery_profile import *
from .get_nsxt_segment_mac_discovery_profile import *
from .get_nsxt_segment_profile_template import *
from .get_nsxt_segment_qos_profile import *
from .get_nsxt_segment_security_profile import *
from .get_nsxt_segment_spoof_guard_profile import *
from .get_nsxt_tier0_router import *
from .get_nsxt_tier0_router_interface import *
from .get_nsxv_application import *
from .get_nsxv_application_finder import *
from .get_nsxv_application_group import *
from .get_nsxv_dhcp_relay import *
from .get_nsxv_distributed_firewall import *
from .get_nsxv_dnat import *
from .get_nsxv_firewall_rule import *
from .get_nsxv_ip_set import *
from .get_nsxv_snat import *
from .get_org import *
from .get_org_group import *
from .get_org_ldap import *
from .get_org_oidc import *
from .get_org_saml import *
from .get_org_saml_metadata import *
from .get_org_user import *
from .get_org_vdc import *
from .get_org_vdc_nsxt_network_profile import *
from .get_org_vdc_template import *
from .get_portgroup import *
from .get_provider_vdc import *
from .get_rde import *
from .get_rde_behavior_invocation import *
from .get_rde_interface import *
from .get_rde_interface_behavior import *
from .get_rde_type import *
from .get_rde_type_behavior import *
from .get_rde_type_behavior_acl import *
from .get_resource_list import *
from .get_resource_pool import *
from .get_resource_schema import *
from .get_right import *
from .get_rights_bundle import *
from .get_role import *
from .get_service_account import *
from .get_solution_add_on import *
from .get_solution_add_on_instance import *
from .get_solution_add_on_instance_publish import *
from .get_solution_landing_zone import *
from .get_storage_profile import *
from .get_subscribed_catalog import *
from .get_task import *
from .get_ui_plugin import *
from .get_vapp import *
from .get_vapp_network import *
from .get_vapp_org_network import *
from .get_vapp_vm import *
from .get_vcenter import *
from .get_vdc_group import *
from .get_version import *
from .get_vgpu_profile import *
from .get_vm import *
from .get_vm_affinity_rule import *
from .get_vm_group import *
from .get_vm_placement_policy import *
from .get_vm_sizing_policy import *
from .get_vm_vgpu_policy import *
from .global_role import *
from .independent_disk import *
from .inserted_media import *
from .ip_space import *
from .ip_space_custom_quota import *
from .ip_space_ip_allocation import *
from .ip_space_uplink import *
from .lb_app_profile import *
from .lb_app_rule import *
from .lb_server_pool import *
from .lb_service_monitor import *
from .lb_virtual_server import *
from .library_certificate import *
from .multisite_org_association import *
from .multisite_site_association import *
from .network_direct import *
from .network_isolated import *
from .network_isolated_v2 import *
from .network_pool import *
from .network_routed import *
from .network_routed_v2 import *
from .nsxt_alb_cloud import *
from .nsxt_alb_controller import *
from .nsxt_alb_edgegateway_service_engine_group import *
from .nsxt_alb_pool import *
from .nsxt_alb_service_engine_group import *
from .nsxt_alb_settings import *
from .nsxt_alb_virtual_service import *
from .nsxt_alb_virtual_service_http_req_rules import *
from .nsxt_alb_virtual_service_http_resp_rules import *
from .nsxt_alb_virtual_service_http_sec_rules import *
from .nsxt_app_port_profile import *
from .nsxt_distributed_firewall import *
from .nsxt_distributed_firewall_rule import *
from .nsxt_dynamic_security_group import *
from .nsxt_edgegateway import *
from .nsxt_edgegateway_bgp_configuration import *
from .nsxt_edgegateway_bgp_ip_prefix_list import *
from .nsxt_edgegateway_bgp_neighbor import *
from .nsxt_edgegateway_dhcp_forwarding import *
from .nsxt_edgegateway_dhcpv6 import *
from .nsxt_edgegateway_dns import *
from .nsxt_edgegateway_l2_vpn_tunnel import *
from .nsxt_edgegateway_rate_limiting import *
from .nsxt_edgegateway_static_route import *
from .nsxt_firewall import *
from .nsxt_global_default_segment_profile_template import *
from .nsxt_ip_set import *
from .nsxt_ipsec_vpn_tunnel import *
from .nsxt_nat_rule import *
from .nsxt_network_dhcp import *
from .nsxt_network_dhcp_binding import *
from .nsxt_network_imported import *
from .nsxt_network_segment_profile import *
from .nsxt_route_advertisement import *
from .nsxt_security_group import *
from .nsxt_segment_profile_template import *
from .nsxv_dhcp_relay import *
from .nsxv_distributed_firewall import *
from .nsxv_dnat import *
from .nsxv_firewall_rule import *
from .nsxv_ip_set import *
from .nsxv_snat import *
from .org import *
from .org_group import *
from .org_ldap import *
from .org_oidc import *
from .org_saml import *
from .org_user import *
from .org_vdc import *
from .org_vdc_access_control import *
from .org_vdc_nsxt_network_profile import *
from .org_vdc_template import *
from .org_vdc_template_instance import *
from .provider import *
from .provider_vdc import *
from .rde import *
from .rde_interface import *
from .rde_interface_behavior import *
from .rde_type import *
from .rde_type_behavior import *
from .rde_type_behavior_acl import *
from .rights_bundle import *
from .role import *
from .security_tag import *
from .service_account import *
from .solution_add_on import *
from .solution_add_on_instance import *
from .solution_add_on_instance_publish import *
from .solution_landing_zone import *
from .subscribed_catalog import *
from .ui_plugin import *
from .vapp import *
from .vapp_access_control import *
from .vapp_firewall_rules import *
from .vapp_nat_rules import *
from .vapp_network import *
from .vapp_org_network import *
from .vapp_static_routing import *
from .vapp_vm import *
from .vdc_group import *
from .vm import *
from .vm_affinity_rule import *
from .vm_internal_disk import *
from .vm_placement_policy import *
from .vm_sizing_policy import *
from .vm_vgpu_policy import *
from ._inputs import *
from . import outputs

# Make subpackages available:
if typing.TYPE_CHECKING:
    import pulumi_vcd.config as __config
    config = __config
    import pulumi_vcd.region as __region
    region = __region
else:
    config = _utilities.lazy_import('pulumi_vcd.config')
    region = _utilities.lazy_import('pulumi_vcd.region')

_utilities.register(
    resource_modules="""
[
 {
  "pkg": "vcd",
  "mod": "index/apiFilter",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/apiFilter:ApiFilter": "ApiFilter"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/apiToken",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/apiToken:ApiToken": "ApiToken"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/catalog",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/catalog:Catalog": "Catalog"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/catalogAccessControl",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/catalogAccessControl:CatalogAccessControl": "CatalogAccessControl"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/catalogItem",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/catalogItem:CatalogItem": "CatalogItem"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/catalogMedia",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/catalogMedia:CatalogMedia": "CatalogMedia"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/catalogVappTemplate",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/catalogVappTemplate:CatalogVappTemplate": "CatalogVappTemplate"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/clonedVapp",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/clonedVapp:ClonedVapp": "ClonedVapp"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/cseKubernetesCluster",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/cseKubernetesCluster:CseKubernetesCluster": "CseKubernetesCluster"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/dseRegistryConfiguration",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/dseRegistryConfiguration:DseRegistryConfiguration": "DseRegistryConfiguration"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/dseSolutionPublish",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/dseSolutionPublish:DseSolutionPublish": "DseSolutionPublish"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/edgegateway",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/edgegateway:Edgegateway": "Edgegateway"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/edgegatewaySettings",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/edgegatewaySettings:EdgegatewaySettings": "EdgegatewaySettings"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/edgegatewayVpn",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/edgegatewayVpn:EdgegatewayVpn": "EdgegatewayVpn"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/externalEndpoint",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/externalEndpoint:ExternalEndpoint": "ExternalEndpoint"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/externalNetwork",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/externalNetwork:ExternalNetwork": "ExternalNetwork"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/externalNetworkV2",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/externalNetworkV2:ExternalNetworkV2": "ExternalNetworkV2"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/globalRole",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/globalRole:GlobalRole": "GlobalRole"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/independentDisk",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/independentDisk:IndependentDisk": "IndependentDisk"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/insertedMedia",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/insertedMedia:InsertedMedia": "InsertedMedia"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/ipSpace",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/ipSpace:IpSpace": "IpSpace"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/ipSpaceCustomQuota",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/ipSpaceCustomQuota:IpSpaceCustomQuota": "IpSpaceCustomQuota"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/ipSpaceIpAllocation",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/ipSpaceIpAllocation:IpSpaceIpAllocation": "IpSpaceIpAllocation"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/ipSpaceUplink",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/ipSpaceUplink:IpSpaceUplink": "IpSpaceUplink"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/lbAppProfile",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/lbAppProfile:LbAppProfile": "LbAppProfile"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/lbAppRule",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/lbAppRule:LbAppRule": "LbAppRule"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/lbServerPool",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/lbServerPool:LbServerPool": "LbServerPool"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/lbServiceMonitor",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/lbServiceMonitor:LbServiceMonitor": "LbServiceMonitor"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/lbVirtualServer",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/lbVirtualServer:LbVirtualServer": "LbVirtualServer"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/libraryCertificate",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/libraryCertificate:LibraryCertificate": "LibraryCertificate"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/multisiteOrgAssociation",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/multisiteOrgAssociation:MultisiteOrgAssociation": "MultisiteOrgAssociation"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/multisiteSiteAssociation",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/multisiteSiteAssociation:MultisiteSiteAssociation": "MultisiteSiteAssociation"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/networkDirect",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/networkDirect:NetworkDirect": "NetworkDirect"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/networkIsolated",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/networkIsolated:NetworkIsolated": "NetworkIsolated"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/networkIsolatedV2",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/networkIsolatedV2:NetworkIsolatedV2": "NetworkIsolatedV2"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/networkPool",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/networkPool:NetworkPool": "NetworkPool"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/networkRouted",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/networkRouted:NetworkRouted": "NetworkRouted"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/networkRoutedV2",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/networkRoutedV2:NetworkRoutedV2": "NetworkRoutedV2"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxtAlbCloud",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxtAlbCloud:NsxtAlbCloud": "NsxtAlbCloud"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxtAlbController",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxtAlbController:NsxtAlbController": "NsxtAlbController"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxtAlbEdgegatewayServiceEngineGroup",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxtAlbEdgegatewayServiceEngineGroup:NsxtAlbEdgegatewayServiceEngineGroup": "NsxtAlbEdgegatewayServiceEngineGroup"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxtAlbPool",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxtAlbPool:NsxtAlbPool": "NsxtAlbPool"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxtAlbServiceEngineGroup",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxtAlbServiceEngineGroup:NsxtAlbServiceEngineGroup": "NsxtAlbServiceEngineGroup"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxtAlbSettings",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxtAlbSettings:NsxtAlbSettings": "NsxtAlbSettings"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxtAlbVirtualService",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxtAlbVirtualService:NsxtAlbVirtualService": "NsxtAlbVirtualService"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxtAlbVirtualServiceHttpReqRules",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxtAlbVirtualServiceHttpReqRules:NsxtAlbVirtualServiceHttpReqRules": "NsxtAlbVirtualServiceHttpReqRules"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxtAlbVirtualServiceHttpRespRules",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxtAlbVirtualServiceHttpRespRules:NsxtAlbVirtualServiceHttpRespRules": "NsxtAlbVirtualServiceHttpRespRules"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxtAlbVirtualServiceHttpSecRules",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxtAlbVirtualServiceHttpSecRules:NsxtAlbVirtualServiceHttpSecRules": "NsxtAlbVirtualServiceHttpSecRules"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxtAppPortProfile",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxtAppPortProfile:NsxtAppPortProfile": "NsxtAppPortProfile"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxtDistributedFirewall",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxtDistributedFirewall:NsxtDistributedFirewall": "NsxtDistributedFirewall"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxtDistributedFirewallRule",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxtDistributedFirewallRule:NsxtDistributedFirewallRule": "NsxtDistributedFirewallRule"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxtDynamicSecurityGroup",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxtDynamicSecurityGroup:NsxtDynamicSecurityGroup": "NsxtDynamicSecurityGroup"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxtEdgegateway",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxtEdgegateway:NsxtEdgegateway": "NsxtEdgegateway"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxtEdgegatewayBgpConfiguration",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxtEdgegatewayBgpConfiguration:NsxtEdgegatewayBgpConfiguration": "NsxtEdgegatewayBgpConfiguration"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxtEdgegatewayBgpIpPrefixList",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxtEdgegatewayBgpIpPrefixList:NsxtEdgegatewayBgpIpPrefixList": "NsxtEdgegatewayBgpIpPrefixList"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxtEdgegatewayBgpNeighbor",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxtEdgegatewayBgpNeighbor:NsxtEdgegatewayBgpNeighbor": "NsxtEdgegatewayBgpNeighbor"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxtEdgegatewayDhcpForwarding",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxtEdgegatewayDhcpForwarding:NsxtEdgegatewayDhcpForwarding": "NsxtEdgegatewayDhcpForwarding"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxtEdgegatewayDhcpv6",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxtEdgegatewayDhcpv6:NsxtEdgegatewayDhcpv6": "NsxtEdgegatewayDhcpv6"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxtEdgegatewayDns",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxtEdgegatewayDns:NsxtEdgegatewayDns": "NsxtEdgegatewayDns"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxtEdgegatewayL2VpnTunnel",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxtEdgegatewayL2VpnTunnel:NsxtEdgegatewayL2VpnTunnel": "NsxtEdgegatewayL2VpnTunnel"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxtEdgegatewayRateLimiting",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxtEdgegatewayRateLimiting:NsxtEdgegatewayRateLimiting": "NsxtEdgegatewayRateLimiting"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxtEdgegatewayStaticRoute",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxtEdgegatewayStaticRoute:NsxtEdgegatewayStaticRoute": "NsxtEdgegatewayStaticRoute"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxtFirewall",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxtFirewall:NsxtFirewall": "NsxtFirewall"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxtGlobalDefaultSegmentProfileTemplate",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxtGlobalDefaultSegmentProfileTemplate:NsxtGlobalDefaultSegmentProfileTemplate": "NsxtGlobalDefaultSegmentProfileTemplate"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxtIpSet",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxtIpSet:NsxtIpSet": "NsxtIpSet"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxtIpsecVpnTunnel",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxtIpsecVpnTunnel:NsxtIpsecVpnTunnel": "NsxtIpsecVpnTunnel"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxtNatRule",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxtNatRule:NsxtNatRule": "NsxtNatRule"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxtNetworkDhcp",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxtNetworkDhcp:NsxtNetworkDhcp": "NsxtNetworkDhcp"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxtNetworkDhcpBinding",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxtNetworkDhcpBinding:NsxtNetworkDhcpBinding": "NsxtNetworkDhcpBinding"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxtNetworkImported",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxtNetworkImported:NsxtNetworkImported": "NsxtNetworkImported"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxtNetworkSegmentProfile",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxtNetworkSegmentProfile:NsxtNetworkSegmentProfile": "NsxtNetworkSegmentProfile"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxtRouteAdvertisement",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxtRouteAdvertisement:NsxtRouteAdvertisement": "NsxtRouteAdvertisement"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxtSecurityGroup",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxtSecurityGroup:NsxtSecurityGroup": "NsxtSecurityGroup"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxtSegmentProfileTemplate",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxtSegmentProfileTemplate:NsxtSegmentProfileTemplate": "NsxtSegmentProfileTemplate"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxvDhcpRelay",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxvDhcpRelay:NsxvDhcpRelay": "NsxvDhcpRelay"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxvDistributedFirewall",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxvDistributedFirewall:NsxvDistributedFirewall": "NsxvDistributedFirewall"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxvDnat",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxvDnat:NsxvDnat": "NsxvDnat"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxvFirewallRule",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxvFirewallRule:NsxvFirewallRule": "NsxvFirewallRule"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxvIpSet",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxvIpSet:NsxvIpSet": "NsxvIpSet"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/nsxvSnat",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/nsxvSnat:NsxvSnat": "NsxvSnat"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/org",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/org:Org": "Org"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/orgGroup",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/orgGroup:OrgGroup": "OrgGroup"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/orgLdap",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/orgLdap:OrgLdap": "OrgLdap"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/orgOidc",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/orgOidc:OrgOidc": "OrgOidc"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/orgSaml",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/orgSaml:OrgSaml": "OrgSaml"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/orgUser",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/orgUser:OrgUser": "OrgUser"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/orgVdc",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/orgVdc:OrgVdc": "OrgVdc"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/orgVdcAccessControl",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/orgVdcAccessControl:OrgVdcAccessControl": "OrgVdcAccessControl"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/orgVdcNsxtNetworkProfile",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/orgVdcNsxtNetworkProfile:OrgVdcNsxtNetworkProfile": "OrgVdcNsxtNetworkProfile"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/orgVdcTemplate",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/orgVdcTemplate:OrgVdcTemplate": "OrgVdcTemplate"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/orgVdcTemplateInstance",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/orgVdcTemplateInstance:OrgVdcTemplateInstance": "OrgVdcTemplateInstance"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/providerVdc",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/providerVdc:ProviderVdc": "ProviderVdc"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/rde",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/rde:Rde": "Rde"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/rdeInterface",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/rdeInterface:RdeInterface": "RdeInterface"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/rdeInterfaceBehavior",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/rdeInterfaceBehavior:RdeInterfaceBehavior": "RdeInterfaceBehavior"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/rdeType",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/rdeType:RdeType": "RdeType"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/rdeTypeBehavior",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/rdeTypeBehavior:RdeTypeBehavior": "RdeTypeBehavior"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/rdeTypeBehaviorAcl",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/rdeTypeBehaviorAcl:RdeTypeBehaviorAcl": "RdeTypeBehaviorAcl"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/rightsBundle",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/rightsBundle:RightsBundle": "RightsBundle"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/role",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/role:Role": "Role"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/securityTag",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/securityTag:SecurityTag": "SecurityTag"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/serviceAccount",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/serviceAccount:ServiceAccount": "ServiceAccount"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/solutionAddOn",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/solutionAddOn:SolutionAddOn": "SolutionAddOn"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/solutionAddOnInstance",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/solutionAddOnInstance:SolutionAddOnInstance": "SolutionAddOnInstance"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/solutionAddOnInstancePublish",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/solutionAddOnInstancePublish:SolutionAddOnInstancePublish": "SolutionAddOnInstancePublish"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/solutionLandingZone",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/solutionLandingZone:SolutionLandingZone": "SolutionLandingZone"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/subscribedCatalog",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/subscribedCatalog:SubscribedCatalog": "SubscribedCatalog"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/uiPlugin",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/uiPlugin:UiPlugin": "UiPlugin"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/vapp",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/vapp:Vapp": "Vapp"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/vappAccessControl",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/vappAccessControl:VappAccessControl": "VappAccessControl"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/vappFirewallRules",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/vappFirewallRules:VappFirewallRules": "VappFirewallRules"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/vappNatRules",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/vappNatRules:VappNatRules": "VappNatRules"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/vappNetwork",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/vappNetwork:VappNetwork": "VappNetwork"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/vappOrgNetwork",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/vappOrgNetwork:VappOrgNetwork": "VappOrgNetwork"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/vappStaticRouting",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/vappStaticRouting:VappStaticRouting": "VappStaticRouting"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/vappVm",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/vappVm:VappVm": "VappVm"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/vdcGroup",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/vdcGroup:VdcGroup": "VdcGroup"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/vm",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/vm:Vm": "Vm"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/vmAffinityRule",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/vmAffinityRule:VmAffinityRule": "VmAffinityRule"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/vmInternalDisk",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/vmInternalDisk:VmInternalDisk": "VmInternalDisk"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/vmPlacementPolicy",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/vmPlacementPolicy:VmPlacementPolicy": "VmPlacementPolicy"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/vmSizingPolicy",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/vmSizingPolicy:VmSizingPolicy": "VmSizingPolicy"
  }
 },
 {
  "pkg": "vcd",
  "mod": "index/vmVgpuPolicy",
  "fqn": "pulumi_vcd",
  "classes": {
   "vcd:index/vmVgpuPolicy:VmVgpuPolicy": "VmVgpuPolicy"
  }
 }
]
""",
    resource_packages="""
[
 {
  "pkg": "vcd",
  "token": "pulumi:providers:vcd",
  "fqn": "pulumi_vcd",
  "class": "Provider"
 }
]
"""
)
