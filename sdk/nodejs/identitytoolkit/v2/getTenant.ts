// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Get a tenant. Requires read permission on the Tenant resource.
 */
export function getTenant(args: GetTenantArgs, opts?: pulumi.InvokeOptions): Promise<GetTenantResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("google-native:identitytoolkit/v2:getTenant", {
        "project": args.project,
        "tenantId": args.tenantId,
    }, opts);
}

export interface GetTenantArgs {
    project?: string;
    tenantId: string;
}

export interface GetTenantResult {
    /**
     * Whether to allow email/password user authentication.
     */
    readonly allowPasswordSignup: boolean;
    /**
     * Whether anonymous users will be auto-deleted after a period of 30 days.
     */
    readonly autodeleteAnonymousUsers: boolean;
    /**
     * Options related to how clients making requests on behalf of a project should be configured.
     */
    readonly client: outputs.identitytoolkit.v2.GoogleCloudIdentitytoolkitAdminV2ClientPermissionConfigResponse;
    /**
     * Whether authentication is disabled for the tenant. If true, the users under the disabled tenant are not allowed to sign-in. Admins of the disabled tenant are not able to manage its users.
     */
    readonly disableAuth: boolean;
    /**
     * Display name of the tenant.
     */
    readonly displayName: string;
    /**
     * Whether to enable anonymous user authentication.
     */
    readonly enableAnonymousUser: boolean;
    /**
     * Whether to enable email link user authentication.
     */
    readonly enableEmailLinkSignin: boolean;
    /**
     * Hash config information of a tenant for display on Pantheon. This can only be displayed on Pantheon to avoid the sensitive information to get accidentally leaked. Only returned in GetTenant response to restrict reading of this information. Requires firebaseauth.configs.getHashConfig permission on the agent project for returning this field.
     */
    readonly hashConfig: outputs.identitytoolkit.v2.GoogleCloudIdentitytoolkitAdminV2HashConfigResponse;
    /**
     * Specify the settings that the tenant could inherit.
     */
    readonly inheritance: outputs.identitytoolkit.v2.GoogleCloudIdentitytoolkitAdminV2InheritanceResponse;
    /**
     * The tenant-level configuration of MFA options.
     */
    readonly mfaConfig: outputs.identitytoolkit.v2.GoogleCloudIdentitytoolkitAdminV2MultiFactorAuthConfigResponse;
    /**
     * Configuration related to monitoring project activity.
     */
    readonly monitoring: outputs.identitytoolkit.v2.GoogleCloudIdentitytoolkitAdminV2MonitoringConfigResponse;
    /**
     * Resource name of a tenant. For example: "projects/{project-id}/tenants/{tenant-id}"
     */
    readonly name: string;
    /**
     * Configures which regions are enabled for SMS verification code sending.
     */
    readonly smsRegionConfig: outputs.identitytoolkit.v2.GoogleCloudIdentitytoolkitAdminV2SmsRegionConfigResponse;
    /**
     * A map of pairs that can be used for MFA. The phone number should be in E.164 format (https://www.itu.int/rec/T-REC-E.164/) and a maximum of 10 pairs can be added (error will be thrown once exceeded).
     */
    readonly testPhoneNumbers: {[key: string]: string};
}

export function getTenantOutput(args: GetTenantOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTenantResult> {
    return pulumi.output(args).apply(a => getTenant(a, opts))
}

export interface GetTenantOutputArgs {
    project?: pulumi.Input<string>;
    tenantId: pulumi.Input<string>;
}
