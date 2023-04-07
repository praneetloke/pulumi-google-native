// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Get a integration in the specified project.
 */
export function getVersion(args: GetVersionArgs, opts?: pulumi.InvokeOptions): Promise<GetVersionResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:integrations/v1alpha:getVersion", {
        "integrationId": args.integrationId,
        "location": args.location,
        "productId": args.productId,
        "project": args.project,
        "versionId": args.versionId,
    }, opts);
}

export interface GetVersionArgs {
    integrationId: string;
    location: string;
    productId: string;
    project?: string;
    versionId: string;
}

export interface GetVersionResult {
    /**
     * Auto-generated.
     */
    readonly createTime: string;
    /**
     * Optional. Flag to disable database persistence for execution data, including event execution info, execution export info, execution metadata index and execution param index.
     */
    readonly databasePersistencePolicy: string;
    /**
     * Optional. The integration description.
     */
    readonly description: string;
    /**
     * Optional. Error Catch Task configuration for the integration. It's optional.
     */
    readonly errorCatcherConfigs: outputs.integrations.v1alpha.GoogleCloudIntegrationsV1alphaErrorCatcherConfigResponse[];
    /**
     * Optional. Parameters that are expected to be passed to the integration when an event is triggered. This consists of all the parameters that are expected in the integration execution. This gives the user the ability to provide default values, add information like PII and also provide data types of each parameter.
     */
    readonly integrationParameters: outputs.integrations.v1alpha.GoogleCloudIntegrationsV1alphaIntegrationParameterResponse[];
    /**
     * Optional. Parameters that are expected to be passed to the integration when an event is triggered. This consists of all the parameters that are expected in the integration execution. This gives the user the ability to provide default values, add information like PII and also provide data types of each parameter.
     */
    readonly integrationParametersInternal: outputs.integrations.v1alpha.EnterpriseCrmFrontendsEventbusProtoWorkflowParametersResponse;
    /**
     * Optional. The last modifier's email address. Generated based on the End User Credentials/LOAS role of the user making the call.
     */
    readonly lastModifierEmail: string;
    /**
     * Optional. The edit lock holder's email address. Generated based on the End User Credentials/LOAS role of the user making the call.
     */
    readonly lockHolder: string;
    /**
     * Auto-generated primary key.
     */
    readonly name: string;
    /**
     * Optional. The origin that indicates where this integration is coming from.
     */
    readonly origin: string;
    /**
     * Optional. The id of the template which was used to create this integration_version.
     */
    readonly parentTemplateId: string;
    /**
     * Optional. The run-as service account email, if set and auth config is not configured, that will be used to generate auth token to be used in Connector task, Rest caller task and Cloud function task.
     */
    readonly runAsServiceAccount: string;
    /**
     * Optional. An increasing sequence that is set when a new snapshot is created. The last created snapshot can be identified by [workflow_name, org_id latest(snapshot_number)]. However, last created snapshot need not be same as the HEAD. So users should always use "HEAD" tag to identify the head.
     */
    readonly snapshotNumber: string;
    /**
     * User should not set it as an input.
     */
    readonly state: string;
    /**
     * Generated by eventbus. User should not set it as an input.
     */
    readonly status: string;
    /**
     * Optional. Task configuration for the integration. It's optional, but the integration doesn't do anything without task_configs.
     */
    readonly taskConfigs: outputs.integrations.v1alpha.GoogleCloudIntegrationsV1alphaTaskConfigResponse[];
    /**
     * Optional. Task configuration for the integration. It's optional, but the integration doesn't do anything without task_configs.
     */
    readonly taskConfigsInternal: outputs.integrations.v1alpha.EnterpriseCrmFrontendsEventbusProtoTaskConfigResponse[];
    /**
     * Optional. Contains a graph of tasks that will be executed before putting the event in a terminal state (SUCCEEDED/FAILED/FATAL), regardless of success or failure, similar to "finally" in code.
     */
    readonly teardown: outputs.integrations.v1alpha.EnterpriseCrmEventbusProtoTeardownResponse;
    /**
     * Optional. Trigger configurations.
     */
    readonly triggerConfigs: outputs.integrations.v1alpha.GoogleCloudIntegrationsV1alphaTriggerConfigResponse[];
    /**
     * Optional. Trigger configurations.
     */
    readonly triggerConfigsInternal: outputs.integrations.v1alpha.EnterpriseCrmFrontendsEventbusProtoTriggerConfigResponse[];
    /**
     * Auto-generated.
     */
    readonly updateTime: string;
    /**
     * Optional. A user-defined label that annotates an integration version. Typically, this is only set when the integration version is created.
     */
    readonly userLabel: string;
}
/**
 * Get a integration in the specified project.
 */
export function getVersionOutput(args: GetVersionOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetVersionResult> {
    return pulumi.output(args).apply((a: any) => getVersion(a, opts))
}

export interface GetVersionOutputArgs {
    integrationId: pulumi.Input<string>;
    location: pulumi.Input<string>;
    productId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    versionId: pulumi.Input<string>;
}
