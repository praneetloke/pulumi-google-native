// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Sets the access control policy on the specified resource. Replaces any existing policy.
 */
export class InstanceIamPolicy extends pulumi.CustomResource {
    /**
     * Get an existing InstanceIamPolicy resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): InstanceIamPolicy {
        return new InstanceIamPolicy(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/alpha:InstanceIamPolicy';

    /**
     * Returns true if the given object is an instance of InstanceIamPolicy.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is InstanceIamPolicy {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === InstanceIamPolicy.__pulumiType;
    }

    /**
     * Specifies cloud audit logging configuration for this policy.
     */
    public readonly auditConfigs!: pulumi.Output<outputs.compute.alpha.AuditConfigResponse[]>;
    /**
     * Associates a list of `members` to a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one member.
     */
    public readonly bindings!: pulumi.Output<outputs.compute.alpha.BindingResponse[]>;
    /**
     * `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
     */
    public readonly etag!: pulumi.Output<string>;
    /**
     * This is deprecated and has no effect. Do not use.
     */
    public readonly iamOwned!: pulumi.Output<boolean>;
    /**
     * This is deprecated and has no effect. Do not use.
     */
    public readonly rules!: pulumi.Output<outputs.compute.alpha.RuleResponse[]>;
    /**
     * Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
     */
    public readonly version!: pulumi.Output<number>;

    /**
     * Create a InstanceIamPolicy resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: InstanceIamPolicyArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.resource === undefined) && !opts.urn) {
                throw new Error("Missing required property 'resource'");
            }
            inputs["auditConfigs"] = args ? args.auditConfigs : undefined;
            inputs["bindings"] = args ? args.bindings : undefined;
            inputs["etag"] = args ? args.etag : undefined;
            inputs["iamOwned"] = args ? args.iamOwned : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["resource"] = args ? args.resource : undefined;
            inputs["rules"] = args ? args.rules : undefined;
            inputs["version"] = args ? args.version : undefined;
            inputs["zone"] = args ? args.zone : undefined;
        } else {
            inputs["auditConfigs"] = undefined /*out*/;
            inputs["bindings"] = undefined /*out*/;
            inputs["etag"] = undefined /*out*/;
            inputs["iamOwned"] = undefined /*out*/;
            inputs["rules"] = undefined /*out*/;
            inputs["version"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(InstanceIamPolicy.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a InstanceIamPolicy resource.
 */
export interface InstanceIamPolicyArgs {
    /**
     * Specifies cloud audit logging configuration for this policy.
     */
    auditConfigs?: pulumi.Input<pulumi.Input<inputs.compute.alpha.AuditConfigArgs>[]>;
    /**
     * Associates a list of `members` to a `role`. Optionally, may specify a `condition` that determines how and when the `bindings` are applied. Each of the `bindings` must contain at least one member.
     */
    bindings?: pulumi.Input<pulumi.Input<inputs.compute.alpha.BindingArgs>[]>;
    /**
     * `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a policy from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform policy updates in order to avoid race conditions: An `etag` is returned in the response to `getIamPolicy`, and systems are expected to put that etag in the request to `setIamPolicy` to ensure that their change will be applied to the same version of the policy. **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost.
     */
    etag?: pulumi.Input<string>;
    /**
     * This is deprecated and has no effect. Do not use.
     */
    iamOwned?: pulumi.Input<boolean>;
    project?: pulumi.Input<string>;
    resource: pulumi.Input<string>;
    /**
     * This is deprecated and has no effect. Do not use.
     */
    rules?: pulumi.Input<pulumi.Input<inputs.compute.alpha.RuleArgs>[]>;
    /**
     * Specifies the format of the policy. Valid values are `0`, `1`, and `3`. Requests that specify an invalid value are rejected. Any operation that affects conditional role bindings must specify version `3`. This requirement applies to the following operations: * Getting a policy that includes a conditional role binding * Adding a conditional role binding to a policy * Changing a conditional role binding in a policy * Removing any role binding, with or without a condition, from a policy that includes conditions **Important:** If you use IAM Conditions, you must include the `etag` field whenever you call `setIamPolicy`. If you omit this field, then IAM allows you to overwrite a version `3` policy with a version `1` policy, and all of the conditions in the version `3` policy are lost. If a policy does not include any conditions, operations on that policy may specify any valid version or leave the field unset. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
     */
    version?: pulumi.Input<number>;
    zone?: pulumi.Input<string>;
}
