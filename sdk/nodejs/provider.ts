// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

/**
 * The provider type for the vcd package. By default, resources use package-wide configuration
 * settings, however an explicit `Provider` instance may be created and passed during resource
 * construction to achieve fine-grained programmatic control over provider settings. See the
 * [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
 */
export class Provider extends pulumi.ProviderResource {
    /** @internal */
    public static readonly __pulumiType = 'vcd';

    /**
     * Returns true if the given object is an instance of Provider.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Provider {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === "pulumi:providers:" + Provider.__pulumiType;
    }

    /**
     * The API token used instead of username/password for VCD API operations. (Requires VCD 10.3.1+)
     */
    public readonly apiToken!: pulumi.Output<string | undefined>;
    /**
     * The API token file instead of username/password for VCD API operations. (Requires VCD 10.3.1+)
     */
    public readonly apiTokenFile!: pulumi.Output<string | undefined>;
    /**
     * 'integrated', 'saml_adfs', 'token', 'api_token', 'api_token_file' and 'service_account_token_file' are supported.
     * 'integrated' is default.
     */
    public readonly authType!: pulumi.Output<string | undefined>;
    public readonly importSeparator!: pulumi.Output<string | undefined>;
    /**
     * Defines the full name of the logging file for API calls (requires 'logging')
     */
    public readonly loggingFile!: pulumi.Output<string | undefined>;
    /**
     * The VCD Org for API operations
     */
    public readonly org!: pulumi.Output<string>;
    /**
     * The user password for VCD API operations.
     */
    public readonly password!: pulumi.Output<string | undefined>;
    /**
     * Allows to specify custom cookie for ADFS server lookup. '{{.Org}}' is replaced by real Org - e.g. 'sso-preferred=yes;
     * sso_redirect_org={{.Org}}'
     */
    public readonly samlAdfsCookie!: pulumi.Output<string | undefined>;
    /**
     * Allows to specify custom Relaying Party Trust Identifier for auth_type=saml_adfs
     */
    public readonly samlAdfsRptId!: pulumi.Output<string | undefined>;
    /**
     * The Service Account API token file instead of username/password for VCD API operations. (Requires VCD 10.4.0+)
     */
    public readonly serviceAccountTokenFile!: pulumi.Output<string | undefined>;
    /**
     * The VCD Org for user authentication
     */
    public readonly sysorg!: pulumi.Output<string | undefined>;
    /**
     * The token used instead of username/password for VCD API operations.
     */
    public readonly token!: pulumi.Output<string | undefined>;
    /**
     * The VCD url for VCD API operations.
     */
    public readonly url!: pulumi.Output<string>;
    /**
     * The user name for VCD API operations.
     */
    public readonly user!: pulumi.Output<string | undefined>;
    /**
     * The VDC for API operations
     */
    public readonly vdc!: pulumi.Output<string | undefined>;

    /**
     * Create a Provider resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ProviderArgs, opts?: pulumi.ResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        {
            if ((!args || args.org === undefined) && !opts.urn) {
                throw new Error("Missing required property 'org'");
            }
            if ((!args || args.url === undefined) && !opts.urn) {
                throw new Error("Missing required property 'url'");
            }
            resourceInputs["allowApiTokenFile"] = pulumi.output(args ? args.allowApiTokenFile : undefined).apply(JSON.stringify);
            resourceInputs["allowServiceAccountTokenFile"] = pulumi.output(args ? args.allowServiceAccountTokenFile : undefined).apply(JSON.stringify);
            resourceInputs["allowUnverifiedSsl"] = pulumi.output(args ? args.allowUnverifiedSsl : undefined).apply(JSON.stringify);
            resourceInputs["apiToken"] = args ? args.apiToken : undefined;
            resourceInputs["apiTokenFile"] = args ? args.apiTokenFile : undefined;
            resourceInputs["authType"] = args ? args.authType : undefined;
            resourceInputs["ignoreMetadataChanges"] = pulumi.output(args ? args.ignoreMetadataChanges : undefined).apply(JSON.stringify);
            resourceInputs["importSeparator"] = args ? args.importSeparator : undefined;
            resourceInputs["logging"] = pulumi.output(args ? args.logging : undefined).apply(JSON.stringify);
            resourceInputs["loggingFile"] = args ? args.loggingFile : undefined;
            resourceInputs["maxRetryTimeout"] = pulumi.output(args ? args.maxRetryTimeout : undefined).apply(JSON.stringify);
            resourceInputs["org"] = args ? args.org : undefined;
            resourceInputs["password"] = args ? args.password : undefined;
            resourceInputs["samlAdfsCookie"] = args ? args.samlAdfsCookie : undefined;
            resourceInputs["samlAdfsRptId"] = args ? args.samlAdfsRptId : undefined;
            resourceInputs["serviceAccountTokenFile"] = args ? args.serviceAccountTokenFile : undefined;
            resourceInputs["sysorg"] = args ? args.sysorg : undefined;
            resourceInputs["token"] = args ? args.token : undefined;
            resourceInputs["url"] = args ? args.url : undefined;
            resourceInputs["user"] = args ? args.user : undefined;
            resourceInputs["vdc"] = args ? args.vdc : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Provider.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Provider resource.
 */
export interface ProviderArgs {
    /**
     * Set this to true if you understand the security risks of using API token files and would like to suppress the warnings
     */
    allowApiTokenFile?: pulumi.Input<boolean>;
    /**
     * Set this to true if you understand the security risks of using Service Account token files and would like to suppress
     * the warnings
     */
    allowServiceAccountTokenFile?: pulumi.Input<boolean>;
    /**
     * If set, VCDClient will permit unverifiable SSL certificates.
     */
    allowUnverifiedSsl?: pulumi.Input<boolean>;
    /**
     * The API token used instead of username/password for VCD API operations. (Requires VCD 10.3.1+)
     */
    apiToken?: pulumi.Input<string>;
    /**
     * The API token file instead of username/password for VCD API operations. (Requires VCD 10.3.1+)
     */
    apiTokenFile?: pulumi.Input<string>;
    /**
     * 'integrated', 'saml_adfs', 'token', 'api_token', 'api_token_file' and 'service_account_token_file' are supported.
     * 'integrated' is default.
     */
    authType?: pulumi.Input<string>;
    /**
     * Defines a set of `metadataEntry` that need to be ignored by this provider. All filters on this attribute are computed
     * with a logical AND
     */
    ignoreMetadataChanges?: pulumi.Input<pulumi.Input<inputs.ProviderIgnoreMetadataChange>[]>;
    importSeparator?: pulumi.Input<string>;
    /**
     * If set, it will enable logging of API requests and responses
     */
    logging?: pulumi.Input<boolean>;
    /**
     * Defines the full name of the logging file for API calls (requires 'logging')
     */
    loggingFile?: pulumi.Input<string>;
    /**
     * Max num seconds to wait for successful response when operating on resources within vCloud (defaults to 60)
     */
    maxRetryTimeout?: pulumi.Input<number>;
    /**
     * The VCD Org for API operations
     */
    org: pulumi.Input<string>;
    /**
     * The user password for VCD API operations.
     */
    password?: pulumi.Input<string>;
    /**
     * Allows to specify custom cookie for ADFS server lookup. '{{.Org}}' is replaced by real Org - e.g. 'sso-preferred=yes;
     * sso_redirect_org={{.Org}}'
     */
    samlAdfsCookie?: pulumi.Input<string>;
    /**
     * Allows to specify custom Relaying Party Trust Identifier for auth_type=saml_adfs
     */
    samlAdfsRptId?: pulumi.Input<string>;
    /**
     * The Service Account API token file instead of username/password for VCD API operations. (Requires VCD 10.4.0+)
     */
    serviceAccountTokenFile?: pulumi.Input<string>;
    /**
     * The VCD Org for user authentication
     */
    sysorg?: pulumi.Input<string>;
    /**
     * The token used instead of username/password for VCD API operations.
     */
    token?: pulumi.Input<string>;
    /**
     * The VCD url for VCD API operations.
     */
    url: pulumi.Input<string>;
    /**
     * The user name for VCD API operations.
     */
    user?: pulumi.Input<string>;
    /**
     * The VDC for API operations
     */
    vdc?: pulumi.Input<string>;
}
