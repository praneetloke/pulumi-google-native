// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Gets an Event Threat Detection custom module.
 */
export function getOrganizationEventThreatDetectionSettingCustomModule(args: GetOrganizationEventThreatDetectionSettingCustomModuleArgs, opts?: pulumi.InvokeOptions): Promise<GetOrganizationEventThreatDetectionSettingCustomModuleResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:securitycenter/v1:getOrganizationEventThreatDetectionSettingCustomModule", {
        "customModuleId": args.customModuleId,
        "organizationId": args.organizationId,
    }, opts);
}

export interface GetOrganizationEventThreatDetectionSettingCustomModuleArgs {
    customModuleId: string;
    organizationId: string;
}

export interface GetOrganizationEventThreatDetectionSettingCustomModuleResult {
    /**
     * Config for the module. For the resident module, its config value is defined at this level. For the inherited module, its config value is inherited from the ancestor module.
     */
    readonly config: {[key: string]: string};
    /**
     * The description for the module.
     */
    readonly description: string;
    /**
     * The human readable name to be displayed for the module.
     */
    readonly displayName: string;
    /**
     * The state of enablement for the module at the given level of the hierarchy.
     */
    readonly enablementState: string;
    /**
     * The editor the module was last updated by.
     */
    readonly lastEditor: string;
    /**
     * Immutable. The resource name of the Event Threat Detection custom module. Its format is: * "organizations/{organization}/eventThreatDetectionSettings/customModules/{module}". * "folders/{folder}/eventThreatDetectionSettings/customModules/{module}". * "projects/{project}/eventThreatDetectionSettings/customModules/{module}".
     */
    readonly name: string;
    /**
     * Type for the module. e.g. CONFIGURABLE_BAD_IP.
     */
    readonly type: string;
    /**
     * The time the module was last updated.
     */
    readonly updateTime: string;
}
/**
 * Gets an Event Threat Detection custom module.
 */
export function getOrganizationEventThreatDetectionSettingCustomModuleOutput(args: GetOrganizationEventThreatDetectionSettingCustomModuleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOrganizationEventThreatDetectionSettingCustomModuleResult> {
    return pulumi.output(args).apply((a: any) => getOrganizationEventThreatDetectionSettingCustomModule(a, opts))
}

export interface GetOrganizationEventThreatDetectionSettingCustomModuleOutputArgs {
    customModuleId: pulumi.Input<string>;
    organizationId: pulumi.Input<string>;
}