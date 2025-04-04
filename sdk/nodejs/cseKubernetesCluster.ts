// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as enums from "./types/enums";
import * as utilities from "./utilities";

export class CseKubernetesCluster extends pulumi.CustomResource {
    /**
     * Get an existing CseKubernetesCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: CseKubernetesClusterState, opts?: pulumi.CustomResourceOptions): CseKubernetesCluster {
        return new CseKubernetesCluster(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'vcd:index/cseKubernetesCluster:CseKubernetesCluster';

    /**
     * Returns true if the given object is an instance of CseKubernetesCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is CseKubernetesCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === CseKubernetesCluster.__pulumiType;
    }

    /**
     * Must be a file generated by [`vcd.ApiToken` resource](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/api_token),
     * or a file that follows the same formatting, that stores the API token used to create and manage the cluster,
     * owned by the user specified in `owner`. Be careful about this file, as it contains sensitive information
     */
    public readonly apiTokenFile!: pulumi.Output<string | undefined>;
    /**
     * If errors occur before the Kubernetes cluster becomes available, and this argument is `true`,
     * CSE Server will automatically attempt to repair the cluster. Defaults to `false`.
     * Since CSE 4.1.1, when the cluster is available/provisioned, this flag is set automatically to false.
     */
    public readonly autoRepairOnErrors!: pulumi.Output<boolean>;
    /**
     * The version of CAPVCD used by this cluster
     */
    public /*out*/ readonly capvcdVersion!: pulumi.Output<string>;
    /**
     * The cluster resource set bindings of this cluster
     */
    public /*out*/ readonly clusterResourceSetBindings!: pulumi.Output<string[]>;
    /**
     * See **Control Plane**
     */
    public readonly controlPlane!: pulumi.Output<outputs.CseKubernetesClusterControlPlane>;
    /**
     * The version of the Cloud Provider Interface used by this cluster
     */
    public /*out*/ readonly cpiVersion!: pulumi.Output<string>;
    /**
     * Specifies the CSE version to use. Accepted versions: `4.1.0`, `4.1.1` (also for *4.1.1a*),
     * `4.2.0`, `4.2.1`, `4.2.2` (VCD Provider *v3.14.1+*) and `4.2.3` (VCD Provider *v3.14.1+*)
     */
    public readonly cseVersion!: pulumi.Output<string>;
    /**
     * The version of the Container Storage Interface used by this cluster
     */
    public /*out*/ readonly csiVersion!: pulumi.Output<string>;
    /**
     * See **Default Storage Class**
     */
    public readonly defaultStorageClass!: pulumi.Output<outputs.CseKubernetesClusterDefaultStorageClass | undefined>;
    /**
     * A set of events that happened during the Kubernetes cluster lifecycle. They're ordered from most recent to least. Each event has:
     */
    public /*out*/ readonly events!: pulumi.Output<outputs.CseKubernetesClusterEvent[]>;
    /**
     * The ready-to-use Kubeconfig file **contents** as a raw string. Only available when `state=provisioned`
     */
    public /*out*/ readonly kubeconfig!: pulumi.Output<string>;
    /**
     * The ID of the vApp Template that corresponds to a Kubernetes template OVA
     */
    public readonly kubernetesTemplateId!: pulumi.Output<string>;
    /**
     * The version of Kubernetes installed in this cluster
     */
    public /*out*/ readonly kubernetesVersion!: pulumi.Output<string>;
    /**
     * The name of the Kubernetes cluster. It must contain only lowercase alphanumeric characters or "-",
     * start with an alphabetic character, end with an alphanumeric, and contain at most 31 characters
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The ID of the network that the Kubernetes cluster will use
     */
    public readonly networkId!: pulumi.Output<string>;
    /**
     * After the Kubernetes cluster becomes available, nodes that become unhealthy will be
     * remediated according to unhealthy node conditions and remediation rules. Defaults to `false`
     */
    public readonly nodeHealthCheck!: pulumi.Output<boolean | undefined>;
    /**
     * The time, in minutes, to wait for the cluster operations to be successfully completed. For example, during cluster
     * creation, it should be in `provisioned`state before the timeout is reached, otherwise the operation will return an
     * error. For cluster deletion, this timeoutspecifies the time to wait until the cluster is completely deleted. Setting
     * this argument to `0` means to wait indefinitely
     */
    public readonly operationsTimeoutMinutes!: pulumi.Output<number | undefined>;
    /**
     * The name of organization that will host the Kubernetes cluster, optional if defined in the provider configuration
     */
    public readonly org!: pulumi.Output<string | undefined>;
    /**
     * The user that creates the cluster and owns the API token specified in `apiToken`.
     * It must have the `Kubernetes Cluster Author` role that was created during CSE installation.
     * If not specified, it assumes it's the user from the provider configuration
     *
     * > Versions 4.2.2 and 4.2.3 should not use the System administrator for the `owner` nor `apiTokenFile`, as stated in their
     * [release notes](https://docs.vmware.com/en/VMware-Cloud-Director-Container-Service-Extension/4.2.2/rn/vmware-cloud-director-container-service-extension-422-release-notes/index.html#Known%20Issues),
     * there is an existing issue that prevents the cluster to be created.
     */
    public readonly owner!: pulumi.Output<string | undefined>;
    /**
     * A CIDR block for the pods to use. Defaults to `100.96.0.0/11`
     */
    public readonly podsCidr!: pulumi.Output<string | undefined>;
    /**
     * Specifies the Kubernetes runtime to use. Defaults to `tkg` (Tanzu Kubernetes Grid)
     */
    public readonly runtime!: pulumi.Output<string | undefined>;
    /**
     * A CIDR block for the services to use. Defaults to `100.64.0.0/13`
     */
    public readonly servicesCidr!: pulumi.Output<string | undefined>;
    /**
     * The SSH public key used to log in into the cluster nodes
     */
    public readonly sshPublicKey!: pulumi.Output<string | undefined>;
    /**
     * The Kubernetes cluster status, can be `provisioning` when it is being created, `provisioned` when it was successfully
     * created and ready to use, or `error` when an error occurred. `provisioning` can only be obtained when a timeout happens during
     * cluster creation. `error` can only be obtained either with a timeout or when `auto_repair_on_errors=false`.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * A set of vApp Template names that can be fetched with a
     * [`vcd.CatalogVappTemplate` data source](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/catalog_vapp_template) to upgrade the cluster.
     */
    public /*out*/ readonly supportedUpgrades!: pulumi.Output<string[]>;
    /**
     * The version of TKG installed in this cluster
     */
    public /*out*/ readonly tkgProductVersion!: pulumi.Output<string>;
    /**
     * The ID of the VDC that hosts the Kubernetes cluster
     */
    public readonly vdcId!: pulumi.Output<string>;
    /**
     * A virtual IP subnet for the cluster
     */
    public readonly virtualIpSubnet!: pulumi.Output<string | undefined>;
    /**
     * See **Worker Pools**
     */
    public readonly workerPools!: pulumi.Output<outputs.CseKubernetesClusterWorkerPool[]>;

    /**
     * Create a CseKubernetesCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: CseKubernetesClusterArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: CseKubernetesClusterArgs | CseKubernetesClusterState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as CseKubernetesClusterState | undefined;
            resourceInputs["apiTokenFile"] = state ? state.apiTokenFile : undefined;
            resourceInputs["autoRepairOnErrors"] = state ? state.autoRepairOnErrors : undefined;
            resourceInputs["capvcdVersion"] = state ? state.capvcdVersion : undefined;
            resourceInputs["clusterResourceSetBindings"] = state ? state.clusterResourceSetBindings : undefined;
            resourceInputs["controlPlane"] = state ? state.controlPlane : undefined;
            resourceInputs["cpiVersion"] = state ? state.cpiVersion : undefined;
            resourceInputs["cseVersion"] = state ? state.cseVersion : undefined;
            resourceInputs["csiVersion"] = state ? state.csiVersion : undefined;
            resourceInputs["defaultStorageClass"] = state ? state.defaultStorageClass : undefined;
            resourceInputs["events"] = state ? state.events : undefined;
            resourceInputs["kubeconfig"] = state ? state.kubeconfig : undefined;
            resourceInputs["kubernetesTemplateId"] = state ? state.kubernetesTemplateId : undefined;
            resourceInputs["kubernetesVersion"] = state ? state.kubernetesVersion : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["networkId"] = state ? state.networkId : undefined;
            resourceInputs["nodeHealthCheck"] = state ? state.nodeHealthCheck : undefined;
            resourceInputs["operationsTimeoutMinutes"] = state ? state.operationsTimeoutMinutes : undefined;
            resourceInputs["org"] = state ? state.org : undefined;
            resourceInputs["owner"] = state ? state.owner : undefined;
            resourceInputs["podsCidr"] = state ? state.podsCidr : undefined;
            resourceInputs["runtime"] = state ? state.runtime : undefined;
            resourceInputs["servicesCidr"] = state ? state.servicesCidr : undefined;
            resourceInputs["sshPublicKey"] = state ? state.sshPublicKey : undefined;
            resourceInputs["state"] = state ? state.state : undefined;
            resourceInputs["supportedUpgrades"] = state ? state.supportedUpgrades : undefined;
            resourceInputs["tkgProductVersion"] = state ? state.tkgProductVersion : undefined;
            resourceInputs["vdcId"] = state ? state.vdcId : undefined;
            resourceInputs["virtualIpSubnet"] = state ? state.virtualIpSubnet : undefined;
            resourceInputs["workerPools"] = state ? state.workerPools : undefined;
        } else {
            const args = argsOrState as CseKubernetesClusterArgs | undefined;
            if ((!args || args.controlPlane === undefined) && !opts.urn) {
                throw new Error("Missing required property 'controlPlane'");
            }
            if ((!args || args.cseVersion === undefined) && !opts.urn) {
                throw new Error("Missing required property 'cseVersion'");
            }
            if ((!args || args.kubernetesTemplateId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'kubernetesTemplateId'");
            }
            if ((!args || args.networkId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'networkId'");
            }
            if ((!args || args.vdcId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'vdcId'");
            }
            if ((!args || args.workerPools === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workerPools'");
            }
            resourceInputs["apiTokenFile"] = args ? args.apiTokenFile : undefined;
            resourceInputs["autoRepairOnErrors"] = args ? args.autoRepairOnErrors : undefined;
            resourceInputs["controlPlane"] = args ? args.controlPlane : undefined;
            resourceInputs["cseVersion"] = args ? args.cseVersion : undefined;
            resourceInputs["defaultStorageClass"] = args ? args.defaultStorageClass : undefined;
            resourceInputs["kubernetesTemplateId"] = args ? args.kubernetesTemplateId : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["networkId"] = args ? args.networkId : undefined;
            resourceInputs["nodeHealthCheck"] = args ? args.nodeHealthCheck : undefined;
            resourceInputs["operationsTimeoutMinutes"] = args ? args.operationsTimeoutMinutes : undefined;
            resourceInputs["org"] = args ? args.org : undefined;
            resourceInputs["owner"] = args ? args.owner : undefined;
            resourceInputs["podsCidr"] = args ? args.podsCidr : undefined;
            resourceInputs["runtime"] = args ? args.runtime : undefined;
            resourceInputs["servicesCidr"] = args ? args.servicesCidr : undefined;
            resourceInputs["sshPublicKey"] = args ? args.sshPublicKey : undefined;
            resourceInputs["vdcId"] = args ? args.vdcId : undefined;
            resourceInputs["virtualIpSubnet"] = args ? args.virtualIpSubnet : undefined;
            resourceInputs["workerPools"] = args ? args.workerPools : undefined;
            resourceInputs["capvcdVersion"] = undefined /*out*/;
            resourceInputs["clusterResourceSetBindings"] = undefined /*out*/;
            resourceInputs["cpiVersion"] = undefined /*out*/;
            resourceInputs["csiVersion"] = undefined /*out*/;
            resourceInputs["events"] = undefined /*out*/;
            resourceInputs["kubeconfig"] = undefined /*out*/;
            resourceInputs["kubernetesVersion"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["supportedUpgrades"] = undefined /*out*/;
            resourceInputs["tkgProductVersion"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const secretOpts = { additionalSecretOutputs: ["kubeconfig"] };
        opts = pulumi.mergeOptions(opts, secretOpts);
        super(CseKubernetesCluster.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering CseKubernetesCluster resources.
 */
export interface CseKubernetesClusterState {
    /**
     * Must be a file generated by [`vcd.ApiToken` resource](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/api_token),
     * or a file that follows the same formatting, that stores the API token used to create and manage the cluster,
     * owned by the user specified in `owner`. Be careful about this file, as it contains sensitive information
     */
    apiTokenFile?: pulumi.Input<string>;
    /**
     * If errors occur before the Kubernetes cluster becomes available, and this argument is `true`,
     * CSE Server will automatically attempt to repair the cluster. Defaults to `false`.
     * Since CSE 4.1.1, when the cluster is available/provisioned, this flag is set automatically to false.
     */
    autoRepairOnErrors?: pulumi.Input<boolean>;
    /**
     * The version of CAPVCD used by this cluster
     */
    capvcdVersion?: pulumi.Input<string>;
    /**
     * The cluster resource set bindings of this cluster
     */
    clusterResourceSetBindings?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * See **Control Plane**
     */
    controlPlane?: pulumi.Input<inputs.CseKubernetesClusterControlPlane>;
    /**
     * The version of the Cloud Provider Interface used by this cluster
     */
    cpiVersion?: pulumi.Input<string>;
    /**
     * Specifies the CSE version to use. Accepted versions: `4.1.0`, `4.1.1` (also for *4.1.1a*),
     * `4.2.0`, `4.2.1`, `4.2.2` (VCD Provider *v3.14.1+*) and `4.2.3` (VCD Provider *v3.14.1+*)
     */
    cseVersion?: pulumi.Input<string>;
    /**
     * The version of the Container Storage Interface used by this cluster
     */
    csiVersion?: pulumi.Input<string>;
    /**
     * See **Default Storage Class**
     */
    defaultStorageClass?: pulumi.Input<inputs.CseKubernetesClusterDefaultStorageClass>;
    /**
     * A set of events that happened during the Kubernetes cluster lifecycle. They're ordered from most recent to least. Each event has:
     */
    events?: pulumi.Input<pulumi.Input<inputs.CseKubernetesClusterEvent>[]>;
    /**
     * The ready-to-use Kubeconfig file **contents** as a raw string. Only available when `state=provisioned`
     */
    kubeconfig?: pulumi.Input<string>;
    /**
     * The ID of the vApp Template that corresponds to a Kubernetes template OVA
     */
    kubernetesTemplateId?: pulumi.Input<string>;
    /**
     * The version of Kubernetes installed in this cluster
     */
    kubernetesVersion?: pulumi.Input<string>;
    /**
     * The name of the Kubernetes cluster. It must contain only lowercase alphanumeric characters or "-",
     * start with an alphabetic character, end with an alphanumeric, and contain at most 31 characters
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the network that the Kubernetes cluster will use
     */
    networkId?: pulumi.Input<string>;
    /**
     * After the Kubernetes cluster becomes available, nodes that become unhealthy will be
     * remediated according to unhealthy node conditions and remediation rules. Defaults to `false`
     */
    nodeHealthCheck?: pulumi.Input<boolean>;
    /**
     * The time, in minutes, to wait for the cluster operations to be successfully completed. For example, during cluster
     * creation, it should be in `provisioned`state before the timeout is reached, otherwise the operation will return an
     * error. For cluster deletion, this timeoutspecifies the time to wait until the cluster is completely deleted. Setting
     * this argument to `0` means to wait indefinitely
     */
    operationsTimeoutMinutes?: pulumi.Input<number>;
    /**
     * The name of organization that will host the Kubernetes cluster, optional if defined in the provider configuration
     */
    org?: pulumi.Input<string>;
    /**
     * The user that creates the cluster and owns the API token specified in `apiToken`.
     * It must have the `Kubernetes Cluster Author` role that was created during CSE installation.
     * If not specified, it assumes it's the user from the provider configuration
     *
     * > Versions 4.2.2 and 4.2.3 should not use the System administrator for the `owner` nor `apiTokenFile`, as stated in their
     * [release notes](https://docs.vmware.com/en/VMware-Cloud-Director-Container-Service-Extension/4.2.2/rn/vmware-cloud-director-container-service-extension-422-release-notes/index.html#Known%20Issues),
     * there is an existing issue that prevents the cluster to be created.
     */
    owner?: pulumi.Input<string>;
    /**
     * A CIDR block for the pods to use. Defaults to `100.96.0.0/11`
     */
    podsCidr?: pulumi.Input<string>;
    /**
     * Specifies the Kubernetes runtime to use. Defaults to `tkg` (Tanzu Kubernetes Grid)
     */
    runtime?: pulumi.Input<string>;
    /**
     * A CIDR block for the services to use. Defaults to `100.64.0.0/13`
     */
    servicesCidr?: pulumi.Input<string>;
    /**
     * The SSH public key used to log in into the cluster nodes
     */
    sshPublicKey?: pulumi.Input<string>;
    /**
     * The Kubernetes cluster status, can be `provisioning` when it is being created, `provisioned` when it was successfully
     * created and ready to use, or `error` when an error occurred. `provisioning` can only be obtained when a timeout happens during
     * cluster creation. `error` can only be obtained either with a timeout or when `auto_repair_on_errors=false`.
     */
    state?: pulumi.Input<string>;
    /**
     * A set of vApp Template names that can be fetched with a
     * [`vcd.CatalogVappTemplate` data source](https://www.terraform.io/providers/vmware/vcd/latest/docs/data-sources/catalog_vapp_template) to upgrade the cluster.
     */
    supportedUpgrades?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * The version of TKG installed in this cluster
     */
    tkgProductVersion?: pulumi.Input<string>;
    /**
     * The ID of the VDC that hosts the Kubernetes cluster
     */
    vdcId?: pulumi.Input<string>;
    /**
     * A virtual IP subnet for the cluster
     */
    virtualIpSubnet?: pulumi.Input<string>;
    /**
     * See **Worker Pools**
     */
    workerPools?: pulumi.Input<pulumi.Input<inputs.CseKubernetesClusterWorkerPool>[]>;
}

/**
 * The set of arguments for constructing a CseKubernetesCluster resource.
 */
export interface CseKubernetesClusterArgs {
    /**
     * Must be a file generated by [`vcd.ApiToken` resource](https://www.terraform.io/providers/vmware/vcd/latest/docs/resources/api_token),
     * or a file that follows the same formatting, that stores the API token used to create and manage the cluster,
     * owned by the user specified in `owner`. Be careful about this file, as it contains sensitive information
     */
    apiTokenFile?: pulumi.Input<string>;
    /**
     * If errors occur before the Kubernetes cluster becomes available, and this argument is `true`,
     * CSE Server will automatically attempt to repair the cluster. Defaults to `false`.
     * Since CSE 4.1.1, when the cluster is available/provisioned, this flag is set automatically to false.
     */
    autoRepairOnErrors?: pulumi.Input<boolean>;
    /**
     * See **Control Plane**
     */
    controlPlane: pulumi.Input<inputs.CseKubernetesClusterControlPlane>;
    /**
     * Specifies the CSE version to use. Accepted versions: `4.1.0`, `4.1.1` (also for *4.1.1a*),
     * `4.2.0`, `4.2.1`, `4.2.2` (VCD Provider *v3.14.1+*) and `4.2.3` (VCD Provider *v3.14.1+*)
     */
    cseVersion: pulumi.Input<string>;
    /**
     * See **Default Storage Class**
     */
    defaultStorageClass?: pulumi.Input<inputs.CseKubernetesClusterDefaultStorageClass>;
    /**
     * The ID of the vApp Template that corresponds to a Kubernetes template OVA
     */
    kubernetesTemplateId: pulumi.Input<string>;
    /**
     * The name of the Kubernetes cluster. It must contain only lowercase alphanumeric characters or "-",
     * start with an alphabetic character, end with an alphanumeric, and contain at most 31 characters
     */
    name?: pulumi.Input<string>;
    /**
     * The ID of the network that the Kubernetes cluster will use
     */
    networkId: pulumi.Input<string>;
    /**
     * After the Kubernetes cluster becomes available, nodes that become unhealthy will be
     * remediated according to unhealthy node conditions and remediation rules. Defaults to `false`
     */
    nodeHealthCheck?: pulumi.Input<boolean>;
    /**
     * The time, in minutes, to wait for the cluster operations to be successfully completed. For example, during cluster
     * creation, it should be in `provisioned`state before the timeout is reached, otherwise the operation will return an
     * error. For cluster deletion, this timeoutspecifies the time to wait until the cluster is completely deleted. Setting
     * this argument to `0` means to wait indefinitely
     */
    operationsTimeoutMinutes?: pulumi.Input<number>;
    /**
     * The name of organization that will host the Kubernetes cluster, optional if defined in the provider configuration
     */
    org?: pulumi.Input<string>;
    /**
     * The user that creates the cluster and owns the API token specified in `apiToken`.
     * It must have the `Kubernetes Cluster Author` role that was created during CSE installation.
     * If not specified, it assumes it's the user from the provider configuration
     *
     * > Versions 4.2.2 and 4.2.3 should not use the System administrator for the `owner` nor `apiTokenFile`, as stated in their
     * [release notes](https://docs.vmware.com/en/VMware-Cloud-Director-Container-Service-Extension/4.2.2/rn/vmware-cloud-director-container-service-extension-422-release-notes/index.html#Known%20Issues),
     * there is an existing issue that prevents the cluster to be created.
     */
    owner?: pulumi.Input<string>;
    /**
     * A CIDR block for the pods to use. Defaults to `100.96.0.0/11`
     */
    podsCidr?: pulumi.Input<string>;
    /**
     * Specifies the Kubernetes runtime to use. Defaults to `tkg` (Tanzu Kubernetes Grid)
     */
    runtime?: pulumi.Input<string>;
    /**
     * A CIDR block for the services to use. Defaults to `100.64.0.0/13`
     */
    servicesCidr?: pulumi.Input<string>;
    /**
     * The SSH public key used to log in into the cluster nodes
     */
    sshPublicKey?: pulumi.Input<string>;
    /**
     * The ID of the VDC that hosts the Kubernetes cluster
     */
    vdcId: pulumi.Input<string>;
    /**
     * A virtual IP subnet for the cluster
     */
    virtualIpSubnet?: pulumi.Input<string>;
    /**
     * See **Worker Pools**
     */
    workerPools: pulumi.Input<pulumi.Input<inputs.CseKubernetesClusterWorkerPool>[]>;
}
