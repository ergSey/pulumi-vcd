// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

export class NsxtIpsecVpnTunnel extends pulumi.CustomResource {
    /**
     * Get an existing NsxtIpsecVpnTunnel resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: NsxtIpsecVpnTunnelState, opts?: pulumi.CustomResourceOptions): NsxtIpsecVpnTunnel {
        return new NsxtIpsecVpnTunnel(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/nsxtIpsecVpnTunnel:NsxtIpsecVpnTunnel';

    /**
     * Returns true if the given object is an instance of NsxtIpsecVpnTunnel.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is NsxtIpsecVpnTunnel {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === NsxtIpsecVpnTunnel.__pulumiType;
    }

    /**
     * `PSK` (pre-shared key) or `CERTIFICATE` (default -
     * `PSK`)
     */
    public readonly authenticationMode!: pulumi.Output<string | undefined>;
    /**
     * CA Certificate ID (can be handled by
     * `vcd.LibraryCertificate` resource or datasource) *Note* `authenticationMode` must be set to
     * `CERTIFICATE`
     */
    public readonly caCertificateId!: pulumi.Output<string | undefined>;
    /**
     * Certificate ID (can be handled by `vcd.LibraryCertificate`
     * resource or datasource). *Note* `authenticationMode` must be set to `CERTIFICATE`
     */
    public readonly certificateId!: pulumi.Output<string | undefined>;
    /**
     * An optional description of the NSX-T IPsec VPN Tunnel
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The ID of the Edge Gateway (NSX-T only). Can be looked up using
     * `vcd.NsxtEdgegateway` data source
     */
    public readonly edgeGatewayId!: pulumi.Output<string>;
    /**
     * Enables or disables IPsec VPN Tunnel (default `true`)
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Provides more details of failure if the IKE service is not UP
     */
    public /*out*/ readonly ikeFailReason!: pulumi.Output<string>;
    /**
     * Status for the actual IKE Session for the given tunnel
     */
    public /*out*/ readonly ikeServiceStatus!: pulumi.Output<string>;
    /**
     * IPv4 Address for the endpoint. This has to be a suballocated IP on the Edge Gateway.
     */
    public readonly localIpAddress!: pulumi.Output<string>;
    /**
     * A set of local networks in CIDR format. At least one value required
     */
    public readonly localNetworks!: pulumi.Output<string[]>;
    /**
     * Sets whether logging for the tunnel is enabled or not. (default - `false`)
     */
    public readonly logging!: pulumi.Output<boolean | undefined>;
    /**
     * A name for NSX-T IPsec VPN Tunnel
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful
     * when connected as sysadmin working across different organisations.
     */
    public readonly org!: pulumi.Output<string | undefined>;
    /**
     * Pre-shared key for negotiation. **Note** the pre-shared key must be
     * the same on the other end of the IPSec VPN tunnel and `authenticationMode` must be `PSK`
     */
    public readonly preSharedKey!: pulumi.Output<string>;
    /**
     * Remote ID uniquely identifies the peer site. If the remote ID is
     * not set, it will default to the remote IP address
     */
    public readonly remoteId!: pulumi.Output<string>;
    /**
     * Public IPv4 Address of the remote device terminating the VPN connection
     */
    public readonly remoteIpAddress!: pulumi.Output<string>;
    /**
     * Set of remote networks in CIDR format. Leaving it empty is interpreted as 0.0.0.0/0
     */
    public readonly remoteNetworks!: pulumi.Output<string[] | undefined>;
    /**
     * `DEFAULT` for system provided configuration or `CUSTOM` if `securityProfileCustomization` is set
     */
    public /*out*/ readonly securityProfile!: pulumi.Output<string>;
    /**
     * a block allowing to
     * customize default security profile parameters
     *
     * <a id="security-profile"></a>
     */
    public readonly securityProfileCustomization!: pulumi.Output<outputs.NsxtIpsecVpnTunnelSecurityProfileCustomization | undefined>;
    /**
     * Overall IPsec VPN Tunnel Status
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The name of VDC to use, optional if defined at provider level
     *
     * @deprecated Edge Gateway will be looked up based on 'edge_gateway_id' field
     */
    public readonly vdc!: pulumi.Output<string>;

    /**
     * Create a NsxtIpsecVpnTunnel resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: NsxtIpsecVpnTunnelArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: NsxtIpsecVpnTunnelArgs | NsxtIpsecVpnTunnelState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as NsxtIpsecVpnTunnelState | undefined;
            resourceInputs["authenticationMode"] = state ? state.authenticationMode : undefined;
            resourceInputs["caCertificateId"] = state ? state.caCertificateId : undefined;
            resourceInputs["certificateId"] = state ? state.certificateId : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["edgeGatewayId"] = state ? state.edgeGatewayId : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["ikeFailReason"] = state ? state.ikeFailReason : undefined;
            resourceInputs["ikeServiceStatus"] = state ? state.ikeServiceStatus : undefined;
            resourceInputs["localIpAddress"] = state ? state.localIpAddress : undefined;
            resourceInputs["localNetworks"] = state ? state.localNetworks : undefined;
            resourceInputs["logging"] = state ? state.logging : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["org"] = state ? state.org : undefined;
            resourceInputs["preSharedKey"] = state ? state.preSharedKey : undefined;
            resourceInputs["remoteId"] = state ? state.remoteId : undefined;
            resourceInputs["remoteIpAddress"] = state ? state.remoteIpAddress : undefined;
            resourceInputs["remoteNetworks"] = state ? state.remoteNetworks : undefined;
            resourceInputs["securityProfile"] = state ? state.securityProfile : undefined;
            resourceInputs["securityProfileCustomization"] = state ? state.securityProfileCustomization : undefined;
            resourceInputs["status"] = state ? state.status : undefined;
            resourceInputs["vdc"] = state ? state.vdc : undefined;
        } else {
            const args = argsOrState as NsxtIpsecVpnTunnelArgs | undefined;
            if ((!args || args.edgeGatewayId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'edgeGatewayId'");
            }
            if ((!args || args.localIpAddress === undefined) && !opts.urn) {
                throw new Error("Missing required property 'localIpAddress'");
            }
            if ((!args || args.localNetworks === undefined) && !opts.urn) {
                throw new Error("Missing required property 'localNetworks'");
            }
            if ((!args || args.preSharedKey === undefined) && !opts.urn) {
                throw new Error("Missing required property 'preSharedKey'");
            }
            if ((!args || args.remoteIpAddress === undefined) && !opts.urn) {
                throw new Error("Missing required property 'remoteIpAddress'");
            }
            resourceInputs["authenticationMode"] = args ? args.authenticationMode : undefined;
            resourceInputs["caCertificateId"] = args ? args.caCertificateId : undefined;
            resourceInputs["certificateId"] = args ? args.certificateId : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["edgeGatewayId"] = args ? args.edgeGatewayId : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["localIpAddress"] = args ? args.localIpAddress : undefined;
            resourceInputs["localNetworks"] = args ? args.localNetworks : undefined;
            resourceInputs["logging"] = args ? args.logging : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["org"] = args ? args.org : undefined;
            resourceInputs["preSharedKey"] = args?.preSharedKey ? pulumi.secret(args.preSharedKey) : undefined;
            resourceInputs["remoteId"] = args ? args.remoteId : undefined;
            resourceInputs["remoteIpAddress"] = args ? args.remoteIpAddress : undefined;
            resourceInputs["remoteNetworks"] = args ? args.remoteNetworks : undefined;
            resourceInputs["securityProfileCustomization"] = args ? args.securityProfileCustomization : undefined;
            resourceInputs["vdc"] = args ? args.vdc : undefined;
            resourceInputs["ikeFailReason"] = undefined /*out*/;
            resourceInputs["ikeServiceStatus"] = undefined /*out*/;
            resourceInputs["securityProfile"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["preSharedKey"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(NsxtIpsecVpnTunnel.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering NsxtIpsecVpnTunnel resources.
 */
export interface NsxtIpsecVpnTunnelState {
    /**
     * `PSK` (pre-shared key) or `CERTIFICATE` (default -
     * `PSK`)
     */
    authenticationMode?: pulumi.Input<string>;
    /**
     * CA Certificate ID (can be handled by
     * `vcd.LibraryCertificate` resource or datasource) *Note* `authenticationMode` must be set to
     * `CERTIFICATE`
     */
    caCertificateId?: pulumi.Input<string>;
    /**
     * Certificate ID (can be handled by `vcd.LibraryCertificate`
     * resource or datasource). *Note* `authenticationMode` must be set to `CERTIFICATE`
     */
    certificateId?: pulumi.Input<string>;
    /**
     * An optional description of the NSX-T IPsec VPN Tunnel
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the Edge Gateway (NSX-T only). Can be looked up using
     * `vcd.NsxtEdgegateway` data source
     */
    edgeGatewayId?: pulumi.Input<string>;
    /**
     * Enables or disables IPsec VPN Tunnel (default `true`)
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Provides more details of failure if the IKE service is not UP
     */
    ikeFailReason?: pulumi.Input<string>;
    /**
     * Status for the actual IKE Session for the given tunnel
     */
    ikeServiceStatus?: pulumi.Input<string>;
    /**
     * IPv4 Address for the endpoint. This has to be a suballocated IP on the Edge Gateway.
     */
    localIpAddress?: pulumi.Input<string>;
    /**
     * A set of local networks in CIDR format. At least one value required
     */
    localNetworks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Sets whether logging for the tunnel is enabled or not. (default - `false`)
     */
    logging?: pulumi.Input<boolean>;
    /**
     * A name for NSX-T IPsec VPN Tunnel
     */
    name?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful
     * when connected as sysadmin working across different organisations.
     */
    org?: pulumi.Input<string>;
    /**
     * Pre-shared key for negotiation. **Note** the pre-shared key must be
     * the same on the other end of the IPSec VPN tunnel and `authenticationMode` must be `PSK`
     */
    preSharedKey?: pulumi.Input<string>;
    /**
     * Remote ID uniquely identifies the peer site. If the remote ID is
     * not set, it will default to the remote IP address
     */
    remoteId?: pulumi.Input<string>;
    /**
     * Public IPv4 Address of the remote device terminating the VPN connection
     */
    remoteIpAddress?: pulumi.Input<string>;
    /**
     * Set of remote networks in CIDR format. Leaving it empty is interpreted as 0.0.0.0/0
     */
    remoteNetworks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * `DEFAULT` for system provided configuration or `CUSTOM` if `securityProfileCustomization` is set
     */
    securityProfile?: pulumi.Input<string>;
    /**
     * a block allowing to
     * customize default security profile parameters
     *
     * <a id="security-profile"></a>
     */
    securityProfileCustomization?: pulumi.Input<inputs.NsxtIpsecVpnTunnelSecurityProfileCustomization>;
    /**
     * Overall IPsec VPN Tunnel Status
     */
    status?: pulumi.Input<string>;
    /**
     * The name of VDC to use, optional if defined at provider level
     *
     * @deprecated Edge Gateway will be looked up based on 'edge_gateway_id' field
     */
    vdc?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a NsxtIpsecVpnTunnel resource.
 */
export interface NsxtIpsecVpnTunnelArgs {
    /**
     * `PSK` (pre-shared key) or `CERTIFICATE` (default -
     * `PSK`)
     */
    authenticationMode?: pulumi.Input<string>;
    /**
     * CA Certificate ID (can be handled by
     * `vcd.LibraryCertificate` resource or datasource) *Note* `authenticationMode` must be set to
     * `CERTIFICATE`
     */
    caCertificateId?: pulumi.Input<string>;
    /**
     * Certificate ID (can be handled by `vcd.LibraryCertificate`
     * resource or datasource). *Note* `authenticationMode` must be set to `CERTIFICATE`
     */
    certificateId?: pulumi.Input<string>;
    /**
     * An optional description of the NSX-T IPsec VPN Tunnel
     */
    description?: pulumi.Input<string>;
    /**
     * The ID of the Edge Gateway (NSX-T only). Can be looked up using
     * `vcd.NsxtEdgegateway` data source
     */
    edgeGatewayId: pulumi.Input<string>;
    /**
     * Enables or disables IPsec VPN Tunnel (default `true`)
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * IPv4 Address for the endpoint. This has to be a suballocated IP on the Edge Gateway.
     */
    localIpAddress: pulumi.Input<string>;
    /**
     * A set of local networks in CIDR format. At least one value required
     */
    localNetworks: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Sets whether logging for the tunnel is enabled or not. (default - `false`)
     */
    logging?: pulumi.Input<boolean>;
    /**
     * A name for NSX-T IPsec VPN Tunnel
     */
    name?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful
     * when connected as sysadmin working across different organisations.
     */
    org?: pulumi.Input<string>;
    /**
     * Pre-shared key for negotiation. **Note** the pre-shared key must be
     * the same on the other end of the IPSec VPN tunnel and `authenticationMode` must be `PSK`
     */
    preSharedKey: pulumi.Input<string>;
    /**
     * Remote ID uniquely identifies the peer site. If the remote ID is
     * not set, it will default to the remote IP address
     */
    remoteId?: pulumi.Input<string>;
    /**
     * Public IPv4 Address of the remote device terminating the VPN connection
     */
    remoteIpAddress: pulumi.Input<string>;
    /**
     * Set of remote networks in CIDR format. Leaving it empty is interpreted as 0.0.0.0/0
     */
    remoteNetworks?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * a block allowing to
     * customize default security profile parameters
     *
     * <a id="security-profile"></a>
     */
    securityProfileCustomization?: pulumi.Input<inputs.NsxtIpsecVpnTunnelSecurityProfileCustomization>;
    /**
     * The name of VDC to use, optional if defined at provider level
     *
     * @deprecated Edge Gateway will be looked up based on 'edge_gateway_id' field
     */
    vdc?: pulumi.Input<string>;
}
