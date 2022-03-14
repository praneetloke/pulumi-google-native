// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Requests that a new DatabaseInstance be created. The state of a successfully created DatabaseInstance is ACTIVE. Only available for projects on the Blaze plan. Projects can be upgraded using the Cloud Billing API https://cloud.google.com/billing/reference/rest/v1/projects/updateBillingInfo. Note that it might take a few minutes for billing enablement state to propagate to Firebase systems.
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:firebasedatabase/v1beta:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * Immutable. The globally unique hostname of the database.
     */
    public readonly databaseUrl!: pulumi.Output<string>;
    /**
     * The fully qualified resource name of the database instance, in the form: `projects/{project-number}/locations/{location-id}/instances/{database-id}`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The resource name of the project this instance belongs to. For example: `projects/{project-number}`.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * The database's lifecycle state. Read-only.
     */
    public readonly state!: pulumi.Output<string>;
    /**
     * The database instance type. On creation only USER_DATABASE is allowed, which is also the default when omitted.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: InstanceArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["databaseId"] = args ? args.databaseId : undefined;
            resourceInputs["databaseUrl"] = args ? args.databaseUrl : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["state"] = args ? args.state : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["validateOnly"] = args ? args.validateOnly : undefined;
        } else {
            resourceInputs["databaseUrl"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Instance.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * The globally unique identifier of the database instance.
     */
    databaseId?: pulumi.Input<string>;
    /**
     * Immutable. The globally unique hostname of the database.
     */
    databaseUrl?: pulumi.Input<string>;
    location?: pulumi.Input<string>;
    /**
     * The fully qualified resource name of the database instance, in the form: `projects/{project-number}/locations/{location-id}/instances/{database-id}`.
     */
    name?: pulumi.Input<string>;
    /**
     * The resource name of the project this instance belongs to. For example: `projects/{project-number}`.
     */
    project?: pulumi.Input<string>;
    /**
     * The database's lifecycle state. Read-only.
     */
    state?: pulumi.Input<enums.firebasedatabase.v1beta.InstanceState>;
    /**
     * The database instance type. On creation only USER_DATABASE is allowed, which is also the default when omitted.
     */
    type?: pulumi.Input<enums.firebasedatabase.v1beta.InstanceType>;
    /**
     * When set to true, the request will be validated but not submitted.
     */
    validateOnly?: pulumi.Input<string>;
}
