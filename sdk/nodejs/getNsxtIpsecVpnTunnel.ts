// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

/**
 * Supported in provider *v3.3+* and VCD 10.1+ with NSX-T backed VDCs.
 *
 * Provides a data source to read NSX-T IPsec VPN Tunnel. You can configure site-to-site connectivity between an NSX-T Data
 * Center Edge Gateway and remote sites. The remote sites must use NSX-T Data Center, have third-party hardware routers,
 * or VPN gateways that support IPSec.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const tunnel1 = vcd.getNsxtIpsecVpnTunnel({
 *     org: "my-org",
 *     edgeGatewayId: existing.id,
 *     name: "tunnel-1",
 * });
 * ```
 */
export function getNsxtIpsecVpnTunnel(args: GetNsxtIpsecVpnTunnelArgs, opts?: pulumi.InvokeOptions): Promise<GetNsxtIpsecVpnTunnelResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("vcd:index/getNsxtIpsecVpnTunnel:getNsxtIpsecVpnTunnel", {
        "edgeGatewayId": args.edgeGatewayId,
        "name": args.name,
        "org": args.org,
        "vdc": args.vdc,
    }, opts);
}

/**
 * A collection of arguments for invoking getNsxtIpsecVpnTunnel.
 */
export interface GetNsxtIpsecVpnTunnelArgs {
    /**
     * The ID of the Edge Gateway (NSX-T only). Can be looked up using `vcd.NsxtEdgegateway`
     * data source
     */
    edgeGatewayId: string;
    /**
     * Name of existing IPsec VPN Tunnel
     */
    name: string;
    /**
     * The name of organization to use, optional if defined at provider level. Useful
     * when connected as sysadmin working across different organisations.
     */
    org?: string;
    /**
     * @deprecated Edge Gateway will be looked up based on 'edge_gateway_id' field
     */
    vdc?: string;
}

/**
 * A collection of values returned by getNsxtIpsecVpnTunnel.
 */
export interface GetNsxtIpsecVpnTunnelResult {
    readonly authenticationMode: string;
    readonly caCertificateId: string;
    readonly certificateId: string;
    readonly description: string;
    readonly edgeGatewayId: string;
    readonly enabled: boolean;
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly ikeFailReason: string;
    readonly ikeServiceStatus: string;
    readonly localIpAddress: string;
    readonly localNetworks: string[];
    readonly logging: boolean;
    readonly name: string;
    readonly org?: string;
    readonly preSharedKey: string;
    readonly remoteId: string;
    readonly remoteIpAddress: string;
    readonly remoteNetworks: string[];
    readonly securityProfile: string;
    readonly securityProfileCustomizations: outputs.GetNsxtIpsecVpnTunnelSecurityProfileCustomization[];
    readonly status: string;
    /**
     * @deprecated Edge Gateway will be looked up based on 'edge_gateway_id' field
     */
    readonly vdc?: string;
}
/**
 * Supported in provider *v3.3+* and VCD 10.1+ with NSX-T backed VDCs.
 *
 * Provides a data source to read NSX-T IPsec VPN Tunnel. You can configure site-to-site connectivity between an NSX-T Data
 * Center Edge Gateway and remote sites. The remote sites must use NSX-T Data Center, have third-party hardware routers,
 * or VPN gateways that support IPSec.
 *
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as vcd from "@pulumi/vcd";
 *
 * const tunnel1 = vcd.getNsxtIpsecVpnTunnel({
 *     org: "my-org",
 *     edgeGatewayId: existing.id,
 *     name: "tunnel-1",
 * });
 * ```
 */
export function getNsxtIpsecVpnTunnelOutput(args: GetNsxtIpsecVpnTunnelOutputArgs, opts?: pulumi.InvokeOutputOptions): pulumi.Output<GetNsxtIpsecVpnTunnelResult> {
    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invokeOutput("vcd:index/getNsxtIpsecVpnTunnel:getNsxtIpsecVpnTunnel", {
        "edgeGatewayId": args.edgeGatewayId,
        "name": args.name,
        "org": args.org,
        "vdc": args.vdc,
    }, opts);
}

/**
 * A collection of arguments for invoking getNsxtIpsecVpnTunnel.
 */
export interface GetNsxtIpsecVpnTunnelOutputArgs {
    /**
     * The ID of the Edge Gateway (NSX-T only). Can be looked up using `vcd.NsxtEdgegateway`
     * data source
     */
    edgeGatewayId: pulumi.Input<string>;
    /**
     * Name of existing IPsec VPN Tunnel
     */
    name: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful
     * when connected as sysadmin working across different organisations.
     */
    org?: pulumi.Input<string>;
    /**
     * @deprecated Edge Gateway will be looked up based on 'edge_gateway_id' field
     */
    vdc?: pulumi.Input<string>;
}
