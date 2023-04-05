// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Create a integration with a draft version in the specified project.
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
    public static readonly __pulumiType = 'google-native:integrations/v1alpha:Version';

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
     * Auto-generated.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Optional. Flag to disable database persistence for execution data, including event execution info, execution export info, execution metadata index and execution param index.
     */
    public readonly databasePersistencePolicy!: pulumi.Output<string>;
    /**
     * Optional. The integration description.
     */
    public readonly description!: pulumi.Output<string>;
    public readonly integrationId!: pulumi.Output<string>;
    /**
     * Optional. Parameters that are expected to be passed to the integration when an event is triggered. This consists of all the parameters that are expected in the integration execution. This gives the user the ability to provide default values, add information like PII and also provide data types of each parameter.
     */
    public readonly integrationParameters!: pulumi.Output<outputs.integrations.v1alpha.GoogleCloudIntegrationsV1alphaIntegrationParameterResponse[]>;
    /**
     * Optional. Parameters that are expected to be passed to the integration when an event is triggered. This consists of all the parameters that are expected in the integration execution. This gives the user the ability to provide default values, add information like PII and also provide data types of each parameter.
     */
    public readonly integrationParametersInternal!: pulumi.Output<outputs.integrations.v1alpha.EnterpriseCrmFrontendsEventbusProtoWorkflowParametersResponse>;
    /**
     * Optional. The last modifier's email address. Generated based on the End User Credentials/LOAS role of the user making the call.
     */
    public readonly lastModifierEmail!: pulumi.Output<string>;
    public readonly location!: pulumi.Output<string>;
    /**
     * Optional. The edit lock holder's email address. Generated based on the End User Credentials/LOAS role of the user making the call.
     */
    public readonly lockHolder!: pulumi.Output<string>;
    /**
     * Auto-generated primary key.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Set this flag to true, if draft version is to be created for a brand new integration. False, if the request is for an existing integration. For backward compatibility reasons, even if this flag is set to `false` and no existing integration is found, a new draft integration will still be created.
     */
    public readonly newIntegration!: pulumi.Output<boolean | undefined>;
    /**
     * Optional. The origin that indicates where this integration is coming from.
     */
    public readonly origin!: pulumi.Output<string>;
    /**
     * Optional. The id of the template which was used to create this integration_version.
     */
    public readonly parentTemplateId!: pulumi.Output<string>;
    public readonly productId!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * Optional. An increasing sequence that is set when a new snapshot is created. The last created snapshot can be identified by [workflow_name, org_id latest(snapshot_number)]. However, last created snapshot need not be same as the HEAD. So users should always use "HEAD" tag to identify the head.
     */
    public readonly snapshotNumber!: pulumi.Output<string>;
    /**
     * User should not set it as an input.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Generated by eventbus. User should not set it as an input.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * Optional. Task configuration for the integration. It's optional, but the integration doesn't do anything without task_configs.
     */
    public readonly taskConfigs!: pulumi.Output<outputs.integrations.v1alpha.GoogleCloudIntegrationsV1alphaTaskConfigResponse[]>;
    /**
     * Optional. Task configuration for the integration. It's optional, but the integration doesn't do anything without task_configs.
     */
    public readonly taskConfigsInternal!: pulumi.Output<outputs.integrations.v1alpha.EnterpriseCrmFrontendsEventbusProtoTaskConfigResponse[]>;
    /**
     * Optional. Contains a graph of tasks that will be executed before putting the event in a terminal state (SUCCEEDED/FAILED/FATAL), regardless of success or failure, similar to "finally" in code.
     */
    public readonly teardown!: pulumi.Output<outputs.integrations.v1alpha.EnterpriseCrmEventbusProtoTeardownResponse>;
    /**
     * Optional. Trigger configurations.
     */
    public readonly triggerConfigs!: pulumi.Output<outputs.integrations.v1alpha.GoogleCloudIntegrationsV1alphaTriggerConfigResponse[]>;
    /**
     * Optional. Trigger configurations.
     */
    public readonly triggerConfigsInternal!: pulumi.Output<outputs.integrations.v1alpha.EnterpriseCrmFrontendsEventbusProtoTriggerConfigResponse[]>;
    /**
     * Auto-generated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    /**
     * Optional. A user-defined label that annotates an integration version. Typically, this is only set when the integration version is created.
     */
    public readonly userLabel!: pulumi.Output<string>;

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
            if ((!args || args.integrationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'integrationId'");
            }
            if ((!args || args.productId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'productId'");
            }
            resourceInputs["databasePersistencePolicy"] = args ? args.databasePersistencePolicy : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["integrationId"] = args ? args.integrationId : undefined;
            resourceInputs["integrationParameters"] = args ? args.integrationParameters : undefined;
            resourceInputs["integrationParametersInternal"] = args ? args.integrationParametersInternal : undefined;
            resourceInputs["lastModifierEmail"] = args ? args.lastModifierEmail : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["lockHolder"] = args ? args.lockHolder : undefined;
            resourceInputs["newIntegration"] = args ? args.newIntegration : undefined;
            resourceInputs["origin"] = args ? args.origin : undefined;
            resourceInputs["parentTemplateId"] = args ? args.parentTemplateId : undefined;
            resourceInputs["productId"] = args ? args.productId : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["snapshotNumber"] = args ? args.snapshotNumber : undefined;
            resourceInputs["taskConfigs"] = args ? args.taskConfigs : undefined;
            resourceInputs["taskConfigsInternal"] = args ? args.taskConfigsInternal : undefined;
            resourceInputs["teardown"] = args ? args.teardown : undefined;
            resourceInputs["triggerConfigs"] = args ? args.triggerConfigs : undefined;
            resourceInputs["triggerConfigsInternal"] = args ? args.triggerConfigsInternal : undefined;
            resourceInputs["userLabel"] = args ? args.userLabel : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["databasePersistencePolicy"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["integrationId"] = undefined /*out*/;
            resourceInputs["integrationParameters"] = undefined /*out*/;
            resourceInputs["integrationParametersInternal"] = undefined /*out*/;
            resourceInputs["lastModifierEmail"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["lockHolder"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["newIntegration"] = undefined /*out*/;
            resourceInputs["origin"] = undefined /*out*/;
            resourceInputs["parentTemplateId"] = undefined /*out*/;
            resourceInputs["productId"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["snapshotNumber"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["status"] = undefined /*out*/;
            resourceInputs["taskConfigs"] = undefined /*out*/;
            resourceInputs["taskConfigsInternal"] = undefined /*out*/;
            resourceInputs["teardown"] = undefined /*out*/;
            resourceInputs["triggerConfigs"] = undefined /*out*/;
            resourceInputs["triggerConfigsInternal"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["userLabel"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["integrationId", "location", "productId", "project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Version.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Version resource.
 */
export interface VersionArgs {
    /**
     * Optional. Flag to disable database persistence for execution data, including event execution info, execution export info, execution metadata index and execution param index.
     */
    databasePersistencePolicy?: pulumi.Input<enums.integrations.v1alpha.VersionDatabasePersistencePolicy>;
    /**
     * Optional. The integration description.
     */
    description?: pulumi.Input<string>;
    integrationId: pulumi.Input<string>;
    /**
     * Optional. Parameters that are expected to be passed to the integration when an event is triggered. This consists of all the parameters that are expected in the integration execution. This gives the user the ability to provide default values, add information like PII and also provide data types of each parameter.
     */
    integrationParameters?: pulumi.Input<pulumi.Input<inputs.integrations.v1alpha.GoogleCloudIntegrationsV1alphaIntegrationParameterArgs>[]>;
    /**
     * Optional. Parameters that are expected to be passed to the integration when an event is triggered. This consists of all the parameters that are expected in the integration execution. This gives the user the ability to provide default values, add information like PII and also provide data types of each parameter.
     */
    integrationParametersInternal?: pulumi.Input<inputs.integrations.v1alpha.EnterpriseCrmFrontendsEventbusProtoWorkflowParametersArgs>;
    /**
     * Optional. The last modifier's email address. Generated based on the End User Credentials/LOAS role of the user making the call.
     */
    lastModifierEmail?: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    /**
     * Optional. The edit lock holder's email address. Generated based on the End User Credentials/LOAS role of the user making the call.
     */
    lockHolder?: pulumi.Input<string>;
    /**
     * Set this flag to true, if draft version is to be created for a brand new integration. False, if the request is for an existing integration. For backward compatibility reasons, even if this flag is set to `false` and no existing integration is found, a new draft integration will still be created.
     */
    newIntegration?: pulumi.Input<boolean>;
    /**
     * Optional. The origin that indicates where this integration is coming from.
     */
    origin?: pulumi.Input<enums.integrations.v1alpha.VersionOrigin>;
    /**
     * Optional. The id of the template which was used to create this integration_version.
     */
    parentTemplateId?: pulumi.Input<string>;
    productId: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * Optional. An increasing sequence that is set when a new snapshot is created. The last created snapshot can be identified by [workflow_name, org_id latest(snapshot_number)]. However, last created snapshot need not be same as the HEAD. So users should always use "HEAD" tag to identify the head.
     */
    snapshotNumber?: pulumi.Input<string>;
    /**
     * Optional. Task configuration for the integration. It's optional, but the integration doesn't do anything without task_configs.
     */
    taskConfigs?: pulumi.Input<pulumi.Input<inputs.integrations.v1alpha.GoogleCloudIntegrationsV1alphaTaskConfigArgs>[]>;
    /**
     * Optional. Task configuration for the integration. It's optional, but the integration doesn't do anything without task_configs.
     */
    taskConfigsInternal?: pulumi.Input<pulumi.Input<inputs.integrations.v1alpha.EnterpriseCrmFrontendsEventbusProtoTaskConfigArgs>[]>;
    /**
     * Optional. Contains a graph of tasks that will be executed before putting the event in a terminal state (SUCCEEDED/FAILED/FATAL), regardless of success or failure, similar to "finally" in code.
     */
    teardown?: pulumi.Input<inputs.integrations.v1alpha.EnterpriseCrmEventbusProtoTeardownArgs>;
    /**
     * Optional. Trigger configurations.
     */
    triggerConfigs?: pulumi.Input<pulumi.Input<inputs.integrations.v1alpha.GoogleCloudIntegrationsV1alphaTriggerConfigArgs>[]>;
    /**
     * Optional. Trigger configurations.
     */
    triggerConfigsInternal?: pulumi.Input<pulumi.Input<inputs.integrations.v1alpha.EnterpriseCrmFrontendsEventbusProtoTriggerConfigArgs>[]>;
    /**
     * Optional. A user-defined label that annotates an integration version. Typically, this is only set when the integration version is created.
     */
    userLabel?: pulumi.Input<string>;
}