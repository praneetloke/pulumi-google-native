// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a new workstation.
 */
export class Workstation extends pulumi.CustomResource {
    /**
     * Get an existing Workstation resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Workstation {
        return new Workstation(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:workstations/v1beta:Workstation';

    /**
     * Returns true if the given object is an instance of Workstation.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Workstation {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Workstation.__pulumiType;
    }

    /**
     * Client-specified annotations.
     */
    public readonly annotations!: pulumi.Output<{[key: string]: string}>;
    /**
     * Time when this resource was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Time when this resource was soft-deleted.
     */
    public /*out*/ readonly deleteTime!: pulumi.Output<string>;
    /**
     * Human-readable name for this resource.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Checksum computed by the server. May be sent on update and delete requests to ensure that the client has an up-to-date value before proceeding.
     */
    public readonly etag!: pulumi.Output<string>;
    /**
     * Host to which clients can send HTTPS traffic that will be received by the workstation. Authorized traffic will be received to the workstation as HTTP on port 80. To send traffic to a different port, clients may prefix the host with the destination port in the format `{port}-{host}`.
     */
    public /*out*/ readonly host!: pulumi.Output<string>;
    /**
     * Client-specified labels that are applied to the resource and that are also propagated to the underlying Compute Engine resources.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    public readonly location!: pulumi.Output<string>;
    /**
     * Full name of this resource.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * Indicates whether this resource is currently being updated to match its intended state.
     */
    public /*out*/ readonly reconciling!: pulumi.Output<boolean>;
    /**
     * Current state of the workstation.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * A system-assigned unique identified for this resource.
     */
    public /*out*/ readonly uid!: pulumi.Output<string>;
    /**
     * Time when this resource was most recently updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;
    public readonly workstationClusterId!: pulumi.Output<string>;
    public readonly workstationConfigId!: pulumi.Output<string>;
    /**
     * Required. ID to use for the workstation.
     */
    public readonly workstationId!: pulumi.Output<string>;

    /**
     * Create a Workstation resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkstationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.workstationClusterId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workstationClusterId'");
            }
            if ((!args || args.workstationConfigId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workstationConfigId'");
            }
            if ((!args || args.workstationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'workstationId'");
            }
            resourceInputs["annotations"] = args ? args.annotations : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["etag"] = args ? args.etag : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["workstationClusterId"] = args ? args.workstationClusterId : undefined;
            resourceInputs["workstationConfigId"] = args ? args.workstationConfigId : undefined;
            resourceInputs["workstationId"] = args ? args.workstationId : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["deleteTime"] = undefined /*out*/;
            resourceInputs["host"] = undefined /*out*/;
            resourceInputs["reconciling"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["annotations"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["deleteTime"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["host"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["reconciling"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["uid"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
            resourceInputs["workstationClusterId"] = undefined /*out*/;
            resourceInputs["workstationConfigId"] = undefined /*out*/;
            resourceInputs["workstationId"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["location", "project", "workstationClusterId", "workstationConfigId", "workstationId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Workstation.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Workstation resource.
 */
export interface WorkstationArgs {
    /**
     * Client-specified annotations.
     */
    annotations?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Human-readable name for this resource.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Checksum computed by the server. May be sent on update and delete requests to ensure that the client has an up-to-date value before proceeding.
     */
    etag?: pulumi.Input<string>;
    /**
     * Client-specified labels that are applied to the resource and that are also propagated to the underlying Compute Engine resources.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * Full name of this resource.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    workstationClusterId: pulumi.Input<string>;
    workstationConfigId: pulumi.Input<string>;
    /**
     * Required. ID to use for the workstation.
     */
    workstationId: pulumi.Input<string>;
}
