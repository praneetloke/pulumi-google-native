// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a trace configuration override. The response contains a system-generated UUID, that can be used to view, update, or delete the configuration override. Use the List API to view the existing trace configuration overrides.
 */
export class OrganizationEnvironmentTraceConfigOverride extends pulumi.CustomResource {
    /**
     * Get an existing OrganizationEnvironmentTraceConfigOverride resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): OrganizationEnvironmentTraceConfigOverride {
        return new OrganizationEnvironmentTraceConfigOverride(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:apigee/v1:OrganizationEnvironmentTraceConfigOverride';

    /**
     * Returns true if the given object is an instance of OrganizationEnvironmentTraceConfigOverride.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is OrganizationEnvironmentTraceConfigOverride {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === OrganizationEnvironmentTraceConfigOverride.__pulumiType;
    }

    /**
     * ID of the API proxy that will have its trace configuration overridden.
     */
    public readonly apiProxy!: pulumi.Output<string>;
    /**
     * ID of the trace configuration override specified as a system-generated UUID.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * Trace configuration to override.
     */
    public readonly samplingConfig!: pulumi.Output<outputs.apigee.v1.GoogleCloudApigeeV1TraceSamplingConfigResponse>;

    /**
     * Create a OrganizationEnvironmentTraceConfigOverride resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrganizationEnvironmentTraceConfigOverrideArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.environmentId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'environmentId'");
            }
            if ((!args || args.organizationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organizationId'");
            }
            if ((!args || args.overrideId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'overrideId'");
            }
            inputs["apiProxy"] = args ? args.apiProxy : undefined;
            inputs["environmentId"] = args ? args.environmentId : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["organizationId"] = args ? args.organizationId : undefined;
            inputs["overrideId"] = args ? args.overrideId : undefined;
            inputs["samplingConfig"] = args ? args.samplingConfig : undefined;
        } else {
            inputs["apiProxy"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["samplingConfig"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(OrganizationEnvironmentTraceConfigOverride.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a OrganizationEnvironmentTraceConfigOverride resource.
 */
export interface OrganizationEnvironmentTraceConfigOverrideArgs {
    /**
     * ID of the API proxy that will have its trace configuration overridden.
     */
    readonly apiProxy?: pulumi.Input<string>;
    readonly environmentId: pulumi.Input<string>;
    /**
     * ID of the trace configuration override specified as a system-generated UUID.
     */
    readonly name?: pulumi.Input<string>;
    readonly organizationId: pulumi.Input<string>;
    readonly overrideId: pulumi.Input<string>;
    /**
     * Trace configuration to override.
     */
    readonly samplingConfig?: pulumi.Input<inputs.apigee.v1.GoogleCloudApigeeV1TraceSamplingConfigArgs>;
}
