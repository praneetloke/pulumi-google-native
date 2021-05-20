// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new function. If a function with the given name already exists in the specified project, the long running operation will return `ALREADY_EXISTS` error.
 */
export class Function extends pulumi.CustomResource {
    /**
     * Get an existing Function resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Function {
        return new Function(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:cloudfunctions/v1:Function';

    /**
     * Returns true if the given object is an instance of Function.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Function {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Function.__pulumiType;
    }

    /**
     * The amount of memory in MB available for a function. Defaults to 256MB.
     */
    public readonly availableMemoryMb!: pulumi.Output<number>;
    /**
     * Build environment variables that shall be available during build time.
     */
    public readonly buildEnvironmentVariables!: pulumi.Output<{[key: string]: string}>;
    /**
     * The Cloud Build ID of the latest successful deployment of the function.
     */
    public /*out*/ readonly buildId!: pulumi.Output<string>;
    /**
     * Name of the Cloud Build Custom Worker Pool that should be used to build the function. The format of this field is `projects/{project}/locations/{region}/workerPools/{workerPool}` where {project} and {region} are the project id and region respectively where the worker pool is defined and {workerPool} is the short name of the worker pool. If the project id is not the same as the function, then the Cloud Functions Service Agent (service-@gcf-admin-robot.iam.gserviceaccount.com) must be granted the role Cloud Build Custom Workers Builder (roles/cloudbuild.customworkers.builder) in the project.
     */
    public readonly buildWorkerPool!: pulumi.Output<string>;
    /**
     * User-provided description of a function.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The name of the function (as defined in source code) that will be executed. Defaults to the resource name suffix, if not specified. For backward compatibility, if function with given name is not found, then the system will try to use function named "function". For Node.js this is name of a function exported by the module specified in `source_location`.
     */
    public readonly entryPoint!: pulumi.Output<string>;
    /**
     * Environment variables that shall be available during function execution.
     */
    public readonly environmentVariables!: pulumi.Output<{[key: string]: string}>;
    /**
     * A source that fires events in response to a condition in another service.
     */
    public readonly eventTrigger!: pulumi.Output<outputs.cloudfunctions.v1.EventTriggerResponse>;
    /**
     * An HTTPS endpoint type of source that can be triggered via URL.
     */
    public readonly httpsTrigger!: pulumi.Output<outputs.cloudfunctions.v1.HttpsTriggerResponse>;
    /**
     * The ingress settings for the function, controlling what traffic can reach it.
     */
    public readonly ingressSettings!: pulumi.Output<string>;
    /**
     * Labels associated with this Cloud Function.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * The limit on the maximum number of function instances that may coexist at a given time. In some cases, such as rapid traffic surges, Cloud Functions may, for a short period of time, create more instances than the specified max instances limit. If your function cannot tolerate this temporary behavior, you may want to factor in a safety margin and set a lower max instances value than your function can tolerate. See the [Max Instances](https://cloud.google.com/functions/docs/max-instances) Guide for more details.
     */
    public readonly maxInstances!: pulumi.Output<number>;
    /**
     * A user-defined name of the function. Function names must be unique globally and match pattern `projects/*&#47;locations/*&#47;functions/*`
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The VPC Network that this cloud function can connect to. It can be either the fully-qualified URI, or the short name of the network resource. If the short network name is used, the network must belong to the same project. Otherwise, it must belong to a project within the same organization. The format of this field is either `projects/{project}/global/networks/{network}` or `{network}`, where {project} is a project id where the network is defined, and {network} is the short name of the network. This field is mutually exclusive with `vpc_connector` and will be replaced by it. See [the VPC documentation](https://cloud.google.com/compute/docs/vpc) for more information on connecting Cloud projects.
     */
    public readonly network!: pulumi.Output<string>;
    /**
     * The runtime in which to run the function. Required when deploying a new function, optional when updating an existing function. For a complete list of possible choices, see the [`gcloud` command reference](/sdk/gcloud/reference/functions/deploy#--runtime).
     */
    public readonly runtime!: pulumi.Output<string>;
    /**
     * The email of the function's service account. If empty, defaults to `{project_id}@appspot.gserviceaccount.com`.
     */
    public readonly serviceAccountEmail!: pulumi.Output<string>;
    /**
     * The Google Cloud Storage URL, starting with gs://, pointing to the zip archive which contains the function.
     */
    public readonly sourceArchiveUrl!: pulumi.Output<string>;
    /**
     * **Beta Feature** The source repository where a function is hosted.
     */
    public readonly sourceRepository!: pulumi.Output<outputs.cloudfunctions.v1.SourceRepositoryResponse>;
    /**
     * Input only. An identifier for Firebase function sources. Disclaimer: This field is only supported for Firebase function deployments.
     */
    public readonly sourceToken!: pulumi.Output<string>;
    /**
     * The Google Cloud Storage signed URL used for source uploading, generated by google.cloud.functions.v1.GenerateUploadUrl
     */
    public readonly sourceUploadUrl!: pulumi.Output<string>;
    /**
     * Status of the function deployment.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * The function execution timeout. Execution is considered failed and can be terminated if the function is not completed at the end of the timeout period. Defaults to 60 seconds.
     */
    public readonly timeout!: pulumi.Output<string>;
    /**
     * The last update timestamp of a Cloud Function.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * The version identifier of the Cloud Function. Each deployment attempt results in a new version of a function being created.
     */
    public /*out*/ readonly versionId!: pulumi.Output<string>;
    /**
     * The VPC Network Connector that this cloud function can connect to. It can be either the fully-qualified URI, or the short name of the network connector resource. The format of this field is `projects/*&#47;locations/*&#47;connectors/*` This field is mutually exclusive with `network` field and will eventually replace it. See [the VPC documentation](https://cloud.google.com/compute/docs/vpc) for more information on connecting Cloud projects.
     */
    public readonly vpcConnector!: pulumi.Output<string>;
    /**
     * The egress settings for the connector, controlling what traffic is diverted through it.
     */
    public readonly vpcConnectorEgressSettings!: pulumi.Output<string>;

    /**
     * Create a Function resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: FunctionArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.functionId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'functionId'");
            }
            if ((!args || args.location === undefined) && !opts.urn) {
                throw new Error("Missing required property 'location'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            inputs["availableMemoryMb"] = args ? args.availableMemoryMb : undefined;
            inputs["buildEnvironmentVariables"] = args ? args.buildEnvironmentVariables : undefined;
            inputs["buildWorkerPool"] = args ? args.buildWorkerPool : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["entryPoint"] = args ? args.entryPoint : undefined;
            inputs["environmentVariables"] = args ? args.environmentVariables : undefined;
            inputs["eventTrigger"] = args ? args.eventTrigger : undefined;
            inputs["functionId"] = args ? args.functionId : undefined;
            inputs["httpsTrigger"] = args ? args.httpsTrigger : undefined;
            inputs["ingressSettings"] = args ? args.ingressSettings : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["location"] = args ? args.location : undefined;
            inputs["maxInstances"] = args ? args.maxInstances : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["network"] = args ? args.network : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["runtime"] = args ? args.runtime : undefined;
            inputs["serviceAccountEmail"] = args ? args.serviceAccountEmail : undefined;
            inputs["sourceArchiveUrl"] = args ? args.sourceArchiveUrl : undefined;
            inputs["sourceRepository"] = args ? args.sourceRepository : undefined;
            inputs["sourceToken"] = args ? args.sourceToken : undefined;
            inputs["sourceUploadUrl"] = args ? args.sourceUploadUrl : undefined;
            inputs["timeout"] = args ? args.timeout : undefined;
            inputs["vpcConnector"] = args ? args.vpcConnector : undefined;
            inputs["vpcConnectorEgressSettings"] = args ? args.vpcConnectorEgressSettings : undefined;
            inputs["buildId"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
            inputs["versionId"] = undefined /*out*/;
        } else {
            inputs["availableMemoryMb"] = undefined /*out*/;
            inputs["buildEnvironmentVariables"] = undefined /*out*/;
            inputs["buildId"] = undefined /*out*/;
            inputs["buildWorkerPool"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["entryPoint"] = undefined /*out*/;
            inputs["environmentVariables"] = undefined /*out*/;
            inputs["eventTrigger"] = undefined /*out*/;
            inputs["httpsTrigger"] = undefined /*out*/;
            inputs["ingressSettings"] = undefined /*out*/;
            inputs["labels"] = undefined /*out*/;
            inputs["maxInstances"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["network"] = undefined /*out*/;
            inputs["runtime"] = undefined /*out*/;
            inputs["serviceAccountEmail"] = undefined /*out*/;
            inputs["sourceArchiveUrl"] = undefined /*out*/;
            inputs["sourceRepository"] = undefined /*out*/;
            inputs["sourceToken"] = undefined /*out*/;
            inputs["sourceUploadUrl"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["timeout"] = undefined /*out*/;
            inputs["updateTime"] = undefined /*out*/;
            inputs["versionId"] = undefined /*out*/;
            inputs["vpcConnector"] = undefined /*out*/;
            inputs["vpcConnectorEgressSettings"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Function.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Function resource.
 */
export interface FunctionArgs {
    /**
     * The amount of memory in MB available for a function. Defaults to 256MB.
     */
    readonly availableMemoryMb?: pulumi.Input<number>;
    /**
     * Build environment variables that shall be available during build time.
     */
    readonly buildEnvironmentVariables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Name of the Cloud Build Custom Worker Pool that should be used to build the function. The format of this field is `projects/{project}/locations/{region}/workerPools/{workerPool}` where {project} and {region} are the project id and region respectively where the worker pool is defined and {workerPool} is the short name of the worker pool. If the project id is not the same as the function, then the Cloud Functions Service Agent (service-@gcf-admin-robot.iam.gserviceaccount.com) must be granted the role Cloud Build Custom Workers Builder (roles/cloudbuild.customworkers.builder) in the project.
     */
    readonly buildWorkerPool?: pulumi.Input<string>;
    /**
     * User-provided description of a function.
     */
    readonly description?: pulumi.Input<string>;
    /**
     * The name of the function (as defined in source code) that will be executed. Defaults to the resource name suffix, if not specified. For backward compatibility, if function with given name is not found, then the system will try to use function named "function". For Node.js this is name of a function exported by the module specified in `source_location`.
     */
    readonly entryPoint?: pulumi.Input<string>;
    /**
     * Environment variables that shall be available during function execution.
     */
    readonly environmentVariables?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * A source that fires events in response to a condition in another service.
     */
    readonly eventTrigger?: pulumi.Input<inputs.cloudfunctions.v1.EventTriggerArgs>;
    readonly functionId: pulumi.Input<string>;
    /**
     * An HTTPS endpoint type of source that can be triggered via URL.
     */
    readonly httpsTrigger?: pulumi.Input<inputs.cloudfunctions.v1.HttpsTriggerArgs>;
    /**
     * The ingress settings for the function, controlling what traffic can reach it.
     */
    readonly ingressSettings?: pulumi.Input<string>;
    /**
     * Labels associated with this Cloud Function.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    readonly location: pulumi.Input<string>;
    /**
     * The limit on the maximum number of function instances that may coexist at a given time. In some cases, such as rapid traffic surges, Cloud Functions may, for a short period of time, create more instances than the specified max instances limit. If your function cannot tolerate this temporary behavior, you may want to factor in a safety margin and set a lower max instances value than your function can tolerate. See the [Max Instances](https://cloud.google.com/functions/docs/max-instances) Guide for more details.
     */
    readonly maxInstances?: pulumi.Input<number>;
    /**
     * A user-defined name of the function. Function names must be unique globally and match pattern `projects/*&#47;locations/*&#47;functions/*`
     */
    readonly name?: pulumi.Input<string>;
    /**
     * The VPC Network that this cloud function can connect to. It can be either the fully-qualified URI, or the short name of the network resource. If the short network name is used, the network must belong to the same project. Otherwise, it must belong to a project within the same organization. The format of this field is either `projects/{project}/global/networks/{network}` or `{network}`, where {project} is a project id where the network is defined, and {network} is the short name of the network. This field is mutually exclusive with `vpc_connector` and will be replaced by it. See [the VPC documentation](https://cloud.google.com/compute/docs/vpc) for more information on connecting Cloud projects.
     */
    readonly network?: pulumi.Input<string>;
    readonly project: pulumi.Input<string>;
    /**
     * The runtime in which to run the function. Required when deploying a new function, optional when updating an existing function. For a complete list of possible choices, see the [`gcloud` command reference](/sdk/gcloud/reference/functions/deploy#--runtime).
     */
    readonly runtime?: pulumi.Input<string>;
    /**
     * The email of the function's service account. If empty, defaults to `{project_id}@appspot.gserviceaccount.com`.
     */
    readonly serviceAccountEmail?: pulumi.Input<string>;
    /**
     * The Google Cloud Storage URL, starting with gs://, pointing to the zip archive which contains the function.
     */
    readonly sourceArchiveUrl?: pulumi.Input<string>;
    /**
     * **Beta Feature** The source repository where a function is hosted.
     */
    readonly sourceRepository?: pulumi.Input<inputs.cloudfunctions.v1.SourceRepositoryArgs>;
    /**
     * Input only. An identifier for Firebase function sources. Disclaimer: This field is only supported for Firebase function deployments.
     */
    readonly sourceToken?: pulumi.Input<string>;
    /**
     * The Google Cloud Storage signed URL used for source uploading, generated by google.cloud.functions.v1.GenerateUploadUrl
     */
    readonly sourceUploadUrl?: pulumi.Input<string>;
    /**
     * The function execution timeout. Execution is considered failed and can be terminated if the function is not completed at the end of the timeout period. Defaults to 60 seconds.
     */
    readonly timeout?: pulumi.Input<string>;
    /**
     * The VPC Network Connector that this cloud function can connect to. It can be either the fully-qualified URI, or the short name of the network connector resource. The format of this field is `projects/*&#47;locations/*&#47;connectors/*` This field is mutually exclusive with `network` field and will eventually replace it. See [the VPC documentation](https://cloud.google.com/compute/docs/vpc) for more information on connecting Cloud projects.
     */
    readonly vpcConnector?: pulumi.Input<string>;
    /**
     * The egress settings for the connector, controlling what traffic is diverted through it.
     */
    readonly vpcConnectorEgressSettings?: pulumi.Input<string>;
}
