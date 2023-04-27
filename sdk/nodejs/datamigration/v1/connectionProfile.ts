// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates a new connection profile in a given project and location.
 */
export class ConnectionProfile extends pulumi.CustomResource {
    /**
     * Get an existing ConnectionProfile resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ConnectionProfile {
        return new ConnectionProfile(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:datamigration/v1:ConnectionProfile';

    /**
     * Returns true if the given object is an instance of ConnectionProfile.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ConnectionProfile {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ConnectionProfile.__pulumiType;
    }

    /**
     * An AlloyDB cluster connection profile.
     */
    public readonly alloydb!: pulumi.Output<outputs.datamigration.v1.AlloyDbConnectionProfileResponse>;
    /**
     * A CloudSQL database connection profile.
     */
    public readonly cloudsql!: pulumi.Output<outputs.datamigration.v1.CloudSqlConnectionProfileResponse>;
    /**
     * Required. The connection profile identifier.
     */
    public readonly connectionProfileId!: pulumi.Output<string>;
    /**
     * The timestamp when the resource was created. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The connection profile display name.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * The error details in case of state FAILED.
     */
    public /*out*/ readonly error!: pulumi.Output<outputs.datamigration.v1.StatusResponse>;
    /**
     * The resource labels for connection profile to use to annotate any related underlying resources such as Compute Engine VMs. An object containing a list of "key": "value" pairs. Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    public readonly location!: pulumi.Output<string>;
    /**
     * A MySQL database connection profile.
     */
    public readonly mysql!: pulumi.Output<outputs.datamigration.v1.MySqlConnectionProfileResponse>;
    /**
     * The name of this connection profile resource in the form of projects/{project}/locations/{location}/connectionProfiles/{connectionProfile}.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * An Oracle database connection profile.
     */
    public readonly oracle!: pulumi.Output<outputs.datamigration.v1.OracleConnectionProfileResponse>;
    /**
     * A PostgreSQL database connection profile.
     */
    public readonly postgresql!: pulumi.Output<outputs.datamigration.v1.PostgreSqlConnectionProfileResponse>;
    public readonly project!: pulumi.Output<string>;
    /**
     * The database provider.
     */
    public readonly provider!: pulumi.Output<string>;
    /**
     * Optional. A unique ID used to identify the request. If the server receives two requests with the same ID, then the second request is ignored. It is recommended to always set this value to a UUID. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). The maximum length is 40 characters.
     */
    public readonly requestId!: pulumi.Output<string | undefined>;
    /**
     * Optional. Create the connection profile without validating it. The default is false. Only supported for Oracle connection profiles.
     */
    public readonly skipValidation!: pulumi.Output<boolean | undefined>;
    /**
     * The current connection profile state (e.g. DRAFT, READY, or FAILED).
     */
    public readonly state!: pulumi.Output<string>;
    /**
     * The timestamp when the resource was last updated. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a ConnectionProfile resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ConnectionProfileArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.connectionProfileId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'connectionProfileId'");
            }
            resourceInputs["alloydb"] = args ? args.alloydb : undefined;
            resourceInputs["cloudsql"] = args ? args.cloudsql : undefined;
            resourceInputs["connectionProfileId"] = args ? args.connectionProfileId : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["mysql"] = args ? args.mysql : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["oracle"] = args ? args.oracle : undefined;
            resourceInputs["postgresql"] = args ? args.postgresql : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["provider"] = args ? args.provider : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["skipValidation"] = args ? args.skipValidation : undefined;
            resourceInputs["state"] = args ? args.state : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["error"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["alloydb"] = undefined /*out*/;
            resourceInputs["cloudsql"] = undefined /*out*/;
            resourceInputs["connectionProfileId"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["error"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["mysql"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["oracle"] = undefined /*out*/;
            resourceInputs["postgresql"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["provider"] = undefined /*out*/;
            resourceInputs["requestId"] = undefined /*out*/;
            resourceInputs["skipValidation"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["connectionProfileId", "location", "project"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(ConnectionProfile.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ConnectionProfile resource.
 */
export interface ConnectionProfileArgs {
    /**
     * An AlloyDB cluster connection profile.
     */
    alloydb?: pulumi.Input<inputs.datamigration.v1.AlloyDbConnectionProfileArgs>;
    /**
     * A CloudSQL database connection profile.
     */
    cloudsql?: pulumi.Input<inputs.datamigration.v1.CloudSqlConnectionProfileArgs>;
    /**
     * Required. The connection profile identifier.
     */
    connectionProfileId: pulumi.Input<string>;
    /**
     * The connection profile display name.
     */
    displayName?: pulumi.Input<string>;
    /**
     * The resource labels for connection profile to use to annotate any related underlying resources such as Compute Engine VMs. An object containing a list of "key": "value" pairs. Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * A MySQL database connection profile.
     */
    mysql?: pulumi.Input<inputs.datamigration.v1.MySqlConnectionProfileArgs>;
    /**
     * The name of this connection profile resource in the form of projects/{project}/locations/{location}/connectionProfiles/{connectionProfile}.
     */
    name?: pulumi.Input<string>;
    /**
     * An Oracle database connection profile.
     */
    oracle?: pulumi.Input<inputs.datamigration.v1.OracleConnectionProfileArgs>;
    /**
     * A PostgreSQL database connection profile.
     */
    postgresql?: pulumi.Input<inputs.datamigration.v1.PostgreSqlConnectionProfileArgs>;
    project?: pulumi.Input<string>;
    /**
     * The database provider.
     */
    provider?: pulumi.Input<enums.datamigration.v1.ConnectionProfileProvider>;
    /**
     * Optional. A unique ID used to identify the request. If the server receives two requests with the same ID, then the second request is ignored. It is recommended to always set this value to a UUID. The ID must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). The maximum length is 40 characters.
     */
    requestId?: pulumi.Input<string>;
    /**
     * Optional. Create the connection profile without validating it. The default is false. Only supported for Oracle connection profiles.
     */
    skipValidation?: pulumi.Input<boolean>;
    /**
     * The current connection profile state (e.g. DRAFT, READY, or FAILED).
     */
    state?: pulumi.Input<enums.datamigration.v1.ConnectionProfileState>;
}
