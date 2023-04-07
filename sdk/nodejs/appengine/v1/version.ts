// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Deploys code and resource files to a new version.
 * Auto-naming is currently not supported for this resource.
 */
export class Version extends pulumi.CustomResource {
    /**
     * Get an existing Version resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Version {
        return new Version(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:appengine/v1:Version';

    /**
     * Returns true if the given object is an instance of Version.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Version {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Version.__pulumiType;
    }

    /**
     * Serving configuration for Google Cloud Endpoints (https://cloud.google.com/endpoints).Only returned in GET requests if view=FULL is set.
     */
    public readonly apiConfig!: pulumi.Output<outputs.appengine.v1.ApiConfigHandlerResponse>;
    /**
     * Allows App Engine second generation runtimes to access the legacy bundled services.
     */
    public readonly appEngineApis!: pulumi.Output<boolean>;
    public readonly appId!: pulumi.Output<string>;
    /**
     * Automatic scaling is based on request rate, response latencies, and other application metrics. Instances are dynamically created and destroyed as needed in order to handle traffic.
     */
    public readonly automaticScaling!: pulumi.Output<outputs.appengine.v1.AutomaticScalingResponse>;
    /**
     * A service with basic scaling will create an instance when the application receives a request. The instance will be turned down when the app becomes idle. Basic scaling is ideal for work that is intermittent or driven by user activity.
     */
    public readonly basicScaling!: pulumi.Output<outputs.appengine.v1.BasicScalingResponse>;
    /**
     * Metadata settings that are supplied to this version to enable beta runtime features.
     */
    public readonly betaSettings!: pulumi.Output<{[key: string]: string}>;
    /**
     * Environment variables available to the build environment.Only returned in GET requests if view=FULL is set.
     */
    public readonly buildEnvVariables!: pulumi.Output<{[key: string]: string}>;
    /**
     * Time that this version was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Email address of the user who created this version.
     */
    public /*out*/ readonly createdBy!: pulumi.Output<string>;
    /**
     * Duration that static files should be cached by web proxies and browsers. Only applicable if the corresponding StaticFilesHandler (https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions#StaticFilesHandler) does not specify its own expiration time.Only returned in GET requests if view=FULL is set.
     */
    public readonly defaultExpiration!: pulumi.Output<string>;
    /**
     * Code and application artifacts that make up this version.Only returned in GET requests if view=FULL is set.
     */
    public readonly deployment!: pulumi.Output<outputs.appengine.v1.DeploymentResponse>;
    /**
     * Total size in bytes of all the files that are included in this version and currently hosted on the App Engine disk.
     */
    public /*out*/ readonly diskUsageBytes!: pulumi.Output<string>;
    /**
     * Cloud Endpoints configuration.If endpoints_api_service is set, the Cloud Endpoints Extensible Service Proxy will be provided to serve the API implemented by the app.
     */
    public readonly endpointsApiService!: pulumi.Output<outputs.appengine.v1.EndpointsApiServiceResponse>;
    /**
     * The entrypoint for the application.
     */
    public readonly entrypoint!: pulumi.Output<outputs.appengine.v1.EntrypointResponse>;
    /**
     * App Engine execution environment for this version.Defaults to standard.
     */
    public readonly env!: pulumi.Output<string>;
    /**
     * Environment variables available to the application.Only returned in GET requests if view=FULL is set.
     */
    public readonly envVariables!: pulumi.Output<{[key: string]: string}>;
    /**
     * Custom static error pages. Limited to 10KB per page.Only returned in GET requests if view=FULL is set.
     */
    public readonly errorHandlers!: pulumi.Output<outputs.appengine.v1.ErrorHandlerResponse[]>;
    /**
     * Settings for App Engine flexible runtimes.
     */
    public readonly flexibleRuntimeSettings!: pulumi.Output<outputs.appengine.v1.FlexibleRuntimeSettingsResponse>;
    /**
     * An ordered list of URL-matching patterns that should be applied to incoming requests. The first matching URL handles the request and other request handlers are not attempted.Only returned in GET requests if view=FULL is set.
     */
    public readonly handlers!: pulumi.Output<outputs.appengine.v1.UrlMapResponse[]>;
    /**
     * Configures health checking for instances. Unhealthy instances are stopped and replaced with new instances. Only applicable in the App Engine flexible environment.Only returned in GET requests if view=FULL is set.
     */
    public readonly healthCheck!: pulumi.Output<outputs.appengine.v1.HealthCheckResponse>;
    /**
     * Before an application can receive email or XMPP messages, the application must be configured to enable the service.
     */
    public readonly inboundServices!: pulumi.Output<string[]>;
    /**
     * Instance class that is used to run this version. Valid values are: AutomaticScaling: F1, F2, F4, F4_1G ManualScaling or BasicScaling: B1, B2, B4, B8, B4_1GDefaults to F1 for AutomaticScaling and B1 for ManualScaling or BasicScaling.
     */
    public readonly instanceClass!: pulumi.Output<string>;
    /**
     * Configuration for third-party Python runtime libraries that are required by the application.Only returned in GET requests if view=FULL is set.
     */
    public readonly libraries!: pulumi.Output<outputs.appengine.v1.LibraryResponse[]>;
    /**
     * Configures liveness health checking for instances. Unhealthy instances are stopped and replaced with new instancesOnly returned in GET requests if view=FULL is set.
     */
    public readonly livenessCheck!: pulumi.Output<outputs.appengine.v1.LivenessCheckResponse>;
    /**
     * A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time. Manually scaled versions are sometimes referred to as "backends".
     */
    public readonly manualScaling!: pulumi.Output<outputs.appengine.v1.ManualScalingResponse>;
    /**
     * Full path to the Version resource in the API. Example: apps/myapp/services/default/versions/v1.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Extra network settings. Only applicable in the App Engine flexible environment.
     */
    public readonly network!: pulumi.Output<outputs.appengine.v1.NetworkResponse>;
    /**
     * Files that match this pattern will not be built into this version. Only applicable for Go runtimes.Only returned in GET requests if view=FULL is set.
     */
    public readonly nobuildFilesRegex!: pulumi.Output<string>;
    /**
     * Configures readiness health checking for instances. Unhealthy instances are not put into the backend traffic rotation.Only returned in GET requests if view=FULL is set.
     */
    public readonly readinessCheck!: pulumi.Output<outputs.appengine.v1.ReadinessCheckResponse>;
    /**
     * Machine resources for this version. Only applicable in the App Engine flexible environment.
     */
    public readonly resources!: pulumi.Output<outputs.appengine.v1.ResourcesResponse>;
    /**
     * Desired runtime. Example: python27.
     */
    public readonly runtime!: pulumi.Output<string>;
    /**
     * The version of the API in the given runtime environment. Please see the app.yaml reference for valid values at https://cloud.google.com/appengine/docs/standard//config/appref
     */
    public readonly runtimeApiVersion!: pulumi.Output<string>;
    /**
     * The channel of the runtime to use. Only available for some runtimes. Defaults to the default channel.
     */
    public readonly runtimeChannel!: pulumi.Output<string>;
    /**
     * The path or name of the app's main executable.
     */
    public readonly runtimeMainExecutablePath!: pulumi.Output<string>;
    /**
     * The identity that the deployed version will run as. Admin API will use the App Engine Appspot service account as default if this field is neither provided in app.yaml file nor through CLI flag.
     */
    public readonly serviceAccount!: pulumi.Output<string>;
    public readonly serviceId!: pulumi.Output<string>;
    /**
     * Current serving status of this version. Only the versions with a SERVING status create instances and can be billed.SERVING_STATUS_UNSPECIFIED is an invalid value. Defaults to SERVING.
     */
    public readonly servingStatus!: pulumi.Output<string>;
    /**
     * Whether multiple requests can be dispatched to this version at once.
     */
    public readonly threadsafe!: pulumi.Output<boolean>;
    /**
     * Serving URL for this version. Example: "https://myversion-dot-myservice-dot-myapp.appspot.com"
     */
    public /*out*/ readonly versionUrl!: pulumi.Output<string>;
    /**
     * Whether to deploy this version in a container on a virtual machine.
     */
    public readonly vm!: pulumi.Output<boolean>;
    /**
     * Enables VPC connectivity for standard apps.
     */
    public readonly vpcAccessConnector!: pulumi.Output<outputs.appengine.v1.VpcAccessConnectorResponse>;
    /**
     * The Google Compute Engine zones that are supported by this version in the App Engine flexible environment. Deprecated.
     *
     * @deprecated The Google Compute Engine zones that are supported by this version in the App Engine flexible environment. Deprecated.
     */
    public readonly zones!: pulumi.Output<string[]>;

    /**
     * Create a Version resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VersionArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.appId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'appId'");
            }
            if ((!args || args.serviceId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceId'");
            }
            resourceInputs["apiConfig"] = args ? args.apiConfig : undefined;
            resourceInputs["appEngineApis"] = args ? args.appEngineApis : undefined;
            resourceInputs["appId"] = args ? args.appId : undefined;
            resourceInputs["automaticScaling"] = args ? args.automaticScaling : undefined;
            resourceInputs["basicScaling"] = args ? args.basicScaling : undefined;
            resourceInputs["betaSettings"] = args ? args.betaSettings : undefined;
            resourceInputs["buildEnvVariables"] = args ? args.buildEnvVariables : undefined;
            resourceInputs["defaultExpiration"] = args ? args.defaultExpiration : undefined;
            resourceInputs["deployment"] = args ? args.deployment : undefined;
            resourceInputs["endpointsApiService"] = args ? args.endpointsApiService : undefined;
            resourceInputs["entrypoint"] = args ? args.entrypoint : undefined;
            resourceInputs["env"] = args ? args.env : undefined;
            resourceInputs["envVariables"] = args ? args.envVariables : undefined;
            resourceInputs["errorHandlers"] = args ? args.errorHandlers : undefined;
            resourceInputs["flexibleRuntimeSettings"] = args ? args.flexibleRuntimeSettings : undefined;
            resourceInputs["handlers"] = args ? args.handlers : undefined;
            resourceInputs["healthCheck"] = args ? args.healthCheck : undefined;
            resourceInputs["id"] = args ? args.id : undefined;
            resourceInputs["inboundServices"] = args ? args.inboundServices : undefined;
            resourceInputs["instanceClass"] = args ? args.instanceClass : undefined;
            resourceInputs["libraries"] = args ? args.libraries : undefined;
            resourceInputs["livenessCheck"] = args ? args.livenessCheck : undefined;
            resourceInputs["manualScaling"] = args ? args.manualScaling : undefined;
            resourceInputs["network"] = args ? args.network : undefined;
            resourceInputs["nobuildFilesRegex"] = args ? args.nobuildFilesRegex : undefined;
            resourceInputs["readinessCheck"] = args ? args.readinessCheck : undefined;
            resourceInputs["resources"] = args ? args.resources : undefined;
            resourceInputs["runtime"] = args ? args.runtime : undefined;
            resourceInputs["runtimeApiVersion"] = args ? args.runtimeApiVersion : undefined;
            resourceInputs["runtimeChannel"] = args ? args.runtimeChannel : undefined;
            resourceInputs["runtimeMainExecutablePath"] = args ? args.runtimeMainExecutablePath : undefined;
            resourceInputs["serviceAccount"] = args ? args.serviceAccount : undefined;
            resourceInputs["serviceId"] = args ? args.serviceId : undefined;
            resourceInputs["servingStatus"] = args ? args.servingStatus : undefined;
            resourceInputs["threadsafe"] = args ? args.threadsafe : undefined;
            resourceInputs["vm"] = args ? args.vm : undefined;
            resourceInputs["vpcAccessConnector"] = args ? args.vpcAccessConnector : undefined;
            resourceInputs["zones"] = args ? args.zones : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["createdBy"] = undefined /*out*/;
            resourceInputs["diskUsageBytes"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["versionUrl"] = undefined /*out*/;
        } else {
            resourceInputs["apiConfig"] = undefined /*out*/;
            resourceInputs["appEngineApis"] = undefined /*out*/;
            resourceInputs["appId"] = undefined /*out*/;
            resourceInputs["automaticScaling"] = undefined /*out*/;
            resourceInputs["basicScaling"] = undefined /*out*/;
            resourceInputs["betaSettings"] = undefined /*out*/;
            resourceInputs["buildEnvVariables"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["createdBy"] = undefined /*out*/;
            resourceInputs["defaultExpiration"] = undefined /*out*/;
            resourceInputs["deployment"] = undefined /*out*/;
            resourceInputs["diskUsageBytes"] = undefined /*out*/;
            resourceInputs["endpointsApiService"] = undefined /*out*/;
            resourceInputs["entrypoint"] = undefined /*out*/;
            resourceInputs["env"] = undefined /*out*/;
            resourceInputs["envVariables"] = undefined /*out*/;
            resourceInputs["errorHandlers"] = undefined /*out*/;
            resourceInputs["flexibleRuntimeSettings"] = undefined /*out*/;
            resourceInputs["handlers"] = undefined /*out*/;
            resourceInputs["healthCheck"] = undefined /*out*/;
            resourceInputs["inboundServices"] = undefined /*out*/;
            resourceInputs["instanceClass"] = undefined /*out*/;
            resourceInputs["libraries"] = undefined /*out*/;
            resourceInputs["livenessCheck"] = undefined /*out*/;
            resourceInputs["manualScaling"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["network"] = undefined /*out*/;
            resourceInputs["nobuildFilesRegex"] = undefined /*out*/;
            resourceInputs["readinessCheck"] = undefined /*out*/;
            resourceInputs["resources"] = undefined /*out*/;
            resourceInputs["runtime"] = undefined /*out*/;
            resourceInputs["runtimeApiVersion"] = undefined /*out*/;
            resourceInputs["runtimeChannel"] = undefined /*out*/;
            resourceInputs["runtimeMainExecutablePath"] = undefined /*out*/;
            resourceInputs["serviceAccount"] = undefined /*out*/;
            resourceInputs["serviceId"] = undefined /*out*/;
            resourceInputs["servingStatus"] = undefined /*out*/;
            resourceInputs["threadsafe"] = undefined /*out*/;
            resourceInputs["versionUrl"] = undefined /*out*/;
            resourceInputs["vm"] = undefined /*out*/;
            resourceInputs["vpcAccessConnector"] = undefined /*out*/;
            resourceInputs["zones"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["appId", "serviceId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Version.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Version resource.
 */
export interface VersionArgs {
    /**
     * Serving configuration for Google Cloud Endpoints (https://cloud.google.com/endpoints).Only returned in GET requests if view=FULL is set.
     */
    apiConfig?: pulumi.Input<inputs.appengine.v1.ApiConfigHandlerArgs>;
    /**
     * Allows App Engine second generation runtimes to access the legacy bundled services.
     */
    appEngineApis?: pulumi.Input<boolean>;
    appId: pulumi.Input<string>;
    /**
     * Automatic scaling is based on request rate, response latencies, and other application metrics. Instances are dynamically created and destroyed as needed in order to handle traffic.
     */
    automaticScaling?: pulumi.Input<inputs.appengine.v1.AutomaticScalingArgs>;
    /**
     * A service with basic scaling will create an instance when the application receives a request. The instance will be turned down when the app becomes idle. Basic scaling is ideal for work that is intermittent or driven by user activity.
     */
    basicScaling?: pulumi.Input<inputs.appengine.v1.BasicScalingArgs>;
    /**
     * Metadata settings that are supplied to this version to enable beta runtime features.
     */
    betaSettings?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Environment variables available to the build environment.Only returned in GET requests if view=FULL is set.
     */
    buildEnvVariables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Duration that static files should be cached by web proxies and browsers. Only applicable if the corresponding StaticFilesHandler (https://cloud.google.com/appengine/docs/admin-api/reference/rest/v1/apps.services.versions#StaticFilesHandler) does not specify its own expiration time.Only returned in GET requests if view=FULL is set.
     */
    defaultExpiration?: pulumi.Input<string>;
    /**
     * Code and application artifacts that make up this version.Only returned in GET requests if view=FULL is set.
     */
    deployment?: pulumi.Input<inputs.appengine.v1.DeploymentArgs>;
    /**
     * Cloud Endpoints configuration.If endpoints_api_service is set, the Cloud Endpoints Extensible Service Proxy will be provided to serve the API implemented by the app.
     */
    endpointsApiService?: pulumi.Input<inputs.appengine.v1.EndpointsApiServiceArgs>;
    /**
     * The entrypoint for the application.
     */
    entrypoint?: pulumi.Input<inputs.appengine.v1.EntrypointArgs>;
    /**
     * App Engine execution environment for this version.Defaults to standard.
     */
    env?: pulumi.Input<string>;
    /**
     * Environment variables available to the application.Only returned in GET requests if view=FULL is set.
     */
    envVariables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Custom static error pages. Limited to 10KB per page.Only returned in GET requests if view=FULL is set.
     */
    errorHandlers?: pulumi.Input<pulumi.Input<inputs.appengine.v1.ErrorHandlerArgs>[]>;
    /**
     * Settings for App Engine flexible runtimes.
     */
    flexibleRuntimeSettings?: pulumi.Input<inputs.appengine.v1.FlexibleRuntimeSettingsArgs>;
    /**
     * An ordered list of URL-matching patterns that should be applied to incoming requests. The first matching URL handles the request and other request handlers are not attempted.Only returned in GET requests if view=FULL is set.
     */
    handlers?: pulumi.Input<pulumi.Input<inputs.appengine.v1.UrlMapArgs>[]>;
    /**
     * Configures health checking for instances. Unhealthy instances are stopped and replaced with new instances. Only applicable in the App Engine flexible environment.Only returned in GET requests if view=FULL is set.
     */
    healthCheck?: pulumi.Input<inputs.appengine.v1.HealthCheckArgs>;
    /**
     * Relative name of the version within the service. Example: v1. Version names can contain only lowercase letters, numbers, or hyphens. Reserved names: "default", "latest", and any name with the prefix "ah-".
     */
    id?: pulumi.Input<string>;
    /**
     * Before an application can receive email or XMPP messages, the application must be configured to enable the service.
     */
    inboundServices?: pulumi.Input<pulumi.Input<enums.appengine.v1.VersionInboundServicesItem>[]>;
    /**
     * Instance class that is used to run this version. Valid values are: AutomaticScaling: F1, F2, F4, F4_1G ManualScaling or BasicScaling: B1, B2, B4, B8, B4_1GDefaults to F1 for AutomaticScaling and B1 for ManualScaling or BasicScaling.
     */
    instanceClass?: pulumi.Input<string>;
    /**
     * Configuration for third-party Python runtime libraries that are required by the application.Only returned in GET requests if view=FULL is set.
     */
    libraries?: pulumi.Input<pulumi.Input<inputs.appengine.v1.LibraryArgs>[]>;
    /**
     * Configures liveness health checking for instances. Unhealthy instances are stopped and replaced with new instancesOnly returned in GET requests if view=FULL is set.
     */
    livenessCheck?: pulumi.Input<inputs.appengine.v1.LivenessCheckArgs>;
    /**
     * A service with manual scaling runs continuously, allowing you to perform complex initialization and rely on the state of its memory over time. Manually scaled versions are sometimes referred to as "backends".
     */
    manualScaling?: pulumi.Input<inputs.appengine.v1.ManualScalingArgs>;
    /**
     * Extra network settings. Only applicable in the App Engine flexible environment.
     */
    network?: pulumi.Input<inputs.appengine.v1.NetworkArgs>;
    /**
     * Files that match this pattern will not be built into this version. Only applicable for Go runtimes.Only returned in GET requests if view=FULL is set.
     */
    nobuildFilesRegex?: pulumi.Input<string>;
    /**
     * Configures readiness health checking for instances. Unhealthy instances are not put into the backend traffic rotation.Only returned in GET requests if view=FULL is set.
     */
    readinessCheck?: pulumi.Input<inputs.appengine.v1.ReadinessCheckArgs>;
    /**
     * Machine resources for this version. Only applicable in the App Engine flexible environment.
     */
    resources?: pulumi.Input<inputs.appengine.v1.ResourcesArgs>;
    /**
     * Desired runtime. Example: python27.
     */
    runtime?: pulumi.Input<string>;
    /**
     * The version of the API in the given runtime environment. Please see the app.yaml reference for valid values at https://cloud.google.com/appengine/docs/standard//config/appref
     */
    runtimeApiVersion?: pulumi.Input<string>;
    /**
     * The channel of the runtime to use. Only available for some runtimes. Defaults to the default channel.
     */
    runtimeChannel?: pulumi.Input<string>;
    /**
     * The path or name of the app's main executable.
     */
    runtimeMainExecutablePath?: pulumi.Input<string>;
    /**
     * The identity that the deployed version will run as. Admin API will use the App Engine Appspot service account as default if this field is neither provided in app.yaml file nor through CLI flag.
     */
    serviceAccount?: pulumi.Input<string>;
    serviceId: pulumi.Input<string>;
    /**
     * Current serving status of this version. Only the versions with a SERVING status create instances and can be billed.SERVING_STATUS_UNSPECIFIED is an invalid value. Defaults to SERVING.
     */
    servingStatus?: pulumi.Input<enums.appengine.v1.VersionServingStatus>;
    /**
     * Whether multiple requests can be dispatched to this version at once.
     */
    threadsafe?: pulumi.Input<boolean>;
    /**
     * Whether to deploy this version in a container on a virtual machine.
     */
    vm?: pulumi.Input<boolean>;
    /**
     * Enables VPC connectivity for standard apps.
     */
    vpcAccessConnector?: pulumi.Input<inputs.appengine.v1.VpcAccessConnectorArgs>;
    /**
     * The Google Compute Engine zones that are supported by this version in the App Engine flexible environment. Deprecated.
     *
     * @deprecated The Google Compute Engine zones that are supported by this version in the App Engine flexible environment. Deprecated.
     */
    zones?: pulumi.Input<pulumi.Input<string>[]>;
}
