// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class LbVirtualServer extends pulumi.CustomResource {
    /**
     * Get an existing LbVirtualServer resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: LbVirtualServerState, opts?: pulumi.CustomResourceOptions): LbVirtualServer {
        return new LbVirtualServer(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/lbVirtualServer:LbVirtualServer';

    /**
     * Returns true if the given object is an instance of LbVirtualServer.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is LbVirtualServer {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === LbVirtualServer.__pulumiType;
    }

    /**
     * Application profile ID to be associated with the virtual server
     */
    public readonly appProfileId!: pulumi.Output<string | undefined>;
    /**
     * List of attached application rule IDs
     */
    public readonly appRuleIds!: pulumi.Output<string[] | undefined>;
    /**
     * Maximum concurrent connections that the virtual server can process
     */
    public readonly connectionLimit!: pulumi.Output<number | undefined>;
    /**
     * Maximum incoming new connection requests per second
     */
    public readonly connectionRateLimit!: pulumi.Output<number | undefined>;
    /**
     * Virtual server description
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * The name of the edge gateway on which the virtual server is to be
     * created
     */
    public readonly edgeGateway!: pulumi.Output<string>;
    /**
     * Defines if the virtual server uses acceleration. Default
     * `false`
     */
    public readonly enableAcceleration!: pulumi.Output<boolean | undefined>;
    /**
     * Defines if the virtual server is enabled. Default `true`
     */
    public readonly enabled!: pulumi.Output<boolean | undefined>;
    /**
     * Set the IP address that the load balancer listens on
     */
    public readonly ipAddress!: pulumi.Output<string>;
    /**
     * Virtual server name
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful
     * when connected as sysadmin working across different organisations
     */
    public readonly org!: pulumi.Output<string | undefined>;
    /**
     * The port number that the load balancer listens on
     */
    public readonly port!: pulumi.Output<number>;
    /**
     * Select the protocol that the virtual server accepts. One of `tcp`, `udp`,
     * `http`, or `https` **Note**: You must select the same protocol used by the selected
     * **Application Profile**
     */
    public readonly protocol!: pulumi.Output<string>;
    /**
     * The server pool that the load balancer will use
     */
    public readonly serverPoolId!: pulumi.Output<string | undefined>;
    /**
     * The name of VDC to use, optional if defined at provider level
     */
    public readonly vdc!: pulumi.Output<string | undefined>;

    /**
     * Create a LbVirtualServer resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: LbVirtualServerArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: LbVirtualServerArgs | LbVirtualServerState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as LbVirtualServerState | undefined;
            resourceInputs["appProfileId"] = state ? state.appProfileId : undefined;
            resourceInputs["appRuleIds"] = state ? state.appRuleIds : undefined;
            resourceInputs["connectionLimit"] = state ? state.connectionLimit : undefined;
            resourceInputs["connectionRateLimit"] = state ? state.connectionRateLimit : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["edgeGateway"] = state ? state.edgeGateway : undefined;
            resourceInputs["enableAcceleration"] = state ? state.enableAcceleration : undefined;
            resourceInputs["enabled"] = state ? state.enabled : undefined;
            resourceInputs["ipAddress"] = state ? state.ipAddress : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["org"] = state ? state.org : undefined;
            resourceInputs["port"] = state ? state.port : undefined;
            resourceInputs["protocol"] = state ? state.protocol : undefined;
            resourceInputs["serverPoolId"] = state ? state.serverPoolId : undefined;
            resourceInputs["vdc"] = state ? state.vdc : undefined;
        } else {
            const args = argsOrState as LbVirtualServerArgs | undefined;
            if ((!args || args.edgeGateway === undefined) && !opts.urn) {
                throw new Error("Missing required property 'edgeGateway'");
            }
            if ((!args || args.ipAddress === undefined) && !opts.urn) {
                throw new Error("Missing required property 'ipAddress'");
            }
            if ((!args || args.port === undefined) && !opts.urn) {
                throw new Error("Missing required property 'port'");
            }
            if ((!args || args.protocol === undefined) && !opts.urn) {
                throw new Error("Missing required property 'protocol'");
            }
            resourceInputs["appProfileId"] = args ? args.appProfileId : undefined;
            resourceInputs["appRuleIds"] = args ? args.appRuleIds : undefined;
            resourceInputs["connectionLimit"] = args ? args.connectionLimit : undefined;
            resourceInputs["connectionRateLimit"] = args ? args.connectionRateLimit : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["edgeGateway"] = args ? args.edgeGateway : undefined;
            resourceInputs["enableAcceleration"] = args ? args.enableAcceleration : undefined;
            resourceInputs["enabled"] = args ? args.enabled : undefined;
            resourceInputs["ipAddress"] = args ? args.ipAddress : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["org"] = args ? args.org : undefined;
            resourceInputs["port"] = args ? args.port : undefined;
            resourceInputs["protocol"] = args ? args.protocol : undefined;
            resourceInputs["serverPoolId"] = args ? args.serverPoolId : undefined;
            resourceInputs["vdc"] = args ? args.vdc : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(LbVirtualServer.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering LbVirtualServer resources.
 */
export interface LbVirtualServerState {
    /**
     * Application profile ID to be associated with the virtual server
     */
    appProfileId?: pulumi.Input<string>;
    /**
     * List of attached application rule IDs
     */
    appRuleIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Maximum concurrent connections that the virtual server can process
     */
    connectionLimit?: pulumi.Input<number>;
    /**
     * Maximum incoming new connection requests per second
     */
    connectionRateLimit?: pulumi.Input<number>;
    /**
     * Virtual server description
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the edge gateway on which the virtual server is to be
     * created
     */
    edgeGateway?: pulumi.Input<string>;
    /**
     * Defines if the virtual server uses acceleration. Default
     * `false`
     */
    enableAcceleration?: pulumi.Input<boolean>;
    /**
     * Defines if the virtual server is enabled. Default `true`
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Set the IP address that the load balancer listens on
     */
    ipAddress?: pulumi.Input<string>;
    /**
     * Virtual server name
     */
    name?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful
     * when connected as sysadmin working across different organisations
     */
    org?: pulumi.Input<string>;
    /**
     * The port number that the load balancer listens on
     */
    port?: pulumi.Input<number>;
    /**
     * Select the protocol that the virtual server accepts. One of `tcp`, `udp`,
     * `http`, or `https` **Note**: You must select the same protocol used by the selected
     * **Application Profile**
     */
    protocol?: pulumi.Input<string>;
    /**
     * The server pool that the load balancer will use
     */
    serverPoolId?: pulumi.Input<string>;
    /**
     * The name of VDC to use, optional if defined at provider level
     */
    vdc?: pulumi.Input<string>;
}

/**
 * The set of arguments for constructing a LbVirtualServer resource.
 */
export interface LbVirtualServerArgs {
    /**
     * Application profile ID to be associated with the virtual server
     */
    appProfileId?: pulumi.Input<string>;
    /**
     * List of attached application rule IDs
     */
    appRuleIds?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Maximum concurrent connections that the virtual server can process
     */
    connectionLimit?: pulumi.Input<number>;
    /**
     * Maximum incoming new connection requests per second
     */
    connectionRateLimit?: pulumi.Input<number>;
    /**
     * Virtual server description
     */
    description?: pulumi.Input<string>;
    /**
     * The name of the edge gateway on which the virtual server is to be
     * created
     */
    edgeGateway: pulumi.Input<string>;
    /**
     * Defines if the virtual server uses acceleration. Default
     * `false`
     */
    enableAcceleration?: pulumi.Input<boolean>;
    /**
     * Defines if the virtual server is enabled. Default `true`
     */
    enabled?: pulumi.Input<boolean>;
    /**
     * Set the IP address that the load balancer listens on
     */
    ipAddress: pulumi.Input<string>;
    /**
     * Virtual server name
     */
    name?: pulumi.Input<string>;
    /**
     * The name of organization to use, optional if defined at provider level. Useful
     * when connected as sysadmin working across different organisations
     */
    org?: pulumi.Input<string>;
    /**
     * The port number that the load balancer listens on
     */
    port: pulumi.Input<number>;
    /**
     * Select the protocol that the virtual server accepts. One of `tcp`, `udp`,
     * `http`, or `https` **Note**: You must select the same protocol used by the selected
     * **Application Profile**
     */
    protocol: pulumi.Input<string>;
    /**
     * The server pool that the load balancer will use
     */
    serverPoolId?: pulumi.Input<string>;
    /**
     * The name of VDC to use, optional if defined at provider level
     */
    vdc?: pulumi.Input<string>;
}
