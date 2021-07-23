// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Gets a stored infoType. See https://cloud.google.com/dlp/docs/creating-stored-infotypes to learn more.
 */
export function getStoredInfoType(args: GetStoredInfoTypeArgs, opts?: pulumi.InvokeOptions): Promise<GetStoredInfoTypeResult> {
    if (!opts) {
        opts = {}
    }

    if (!opts.version) {
        opts.version = utilities.getVersion();
    }
    return pulumi.runtime.invoke("google-native:dlp/v2:getStoredInfoType", {
        "location": args.location,
        "project": args.project,
        "storedInfoTypeId": args.storedInfoTypeId,
    }, opts);
}

export interface GetStoredInfoTypeArgs {
    location: string;
    project?: string;
    storedInfoTypeId: string;
}

export interface GetStoredInfoTypeResult {
    /**
     * Current version of the stored info type.
     */
    readonly currentVersion: outputs.dlp.v2.GooglePrivacyDlpV2StoredInfoTypeVersionResponse;
    /**
     * Resource name.
     */
    readonly name: string;
    /**
     * Pending versions of the stored info type. Empty if no versions are pending.
     */
    readonly pendingVersions: outputs.dlp.v2.GooglePrivacyDlpV2StoredInfoTypeVersionResponse[];
}