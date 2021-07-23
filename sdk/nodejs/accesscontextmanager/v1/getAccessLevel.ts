// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Get an Access Level by resource name.
 */
export function getAccessLevel(args: GetAccessLevelArgs, opts?: pulumi.InvokeOptions): Promise<GetAccessLevelResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:accesscontextmanager/v1:getAccessLevel", {
        "accessLevelFormat": args.accessLevelFormat,
        "accessLevelId": args.accessLevelId,
        "accessPolicyId": args.accessPolicyId,
    }, opts);
}

export interface GetAccessLevelArgs {
    accessLevelFormat?: string;
    accessLevelId: string;
    accessPolicyId: string;
}

export interface GetAccessLevelResult {
    /**
     * A `BasicLevel` composed of `Conditions`.
     */
    readonly basic: outputs.accesscontextmanager.v1.BasicLevelResponse;
    /**
     * A `CustomLevel` written in the Common Expression Language.
     */
    readonly custom: outputs.accesscontextmanager.v1.CustomLevelResponse;
    /**
     * Description of the `AccessLevel` and its use. Does not affect behavior.
     */
    readonly description: string;
    /**
     * Resource name for the Access Level. The `short_name` component must begin with a letter and only include alphanumeric and '_'. Format: `accessPolicies/{policy_id}/accessLevels/{short_name}`. The maximum length of the `short_name` component is 50 characters.
     */
    readonly name: string;
    /**
     * Human readable title. Must be unique within the Policy.
     */
    readonly title: string;
}