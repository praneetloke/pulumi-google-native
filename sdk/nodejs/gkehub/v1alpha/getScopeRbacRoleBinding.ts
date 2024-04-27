// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Returns the details of a Scope RBACRoleBinding.
 */
export function getScopeRbacRoleBinding(args: GetScopeRbacRoleBindingArgs, opts?: pulumi.InvokeOptions): Promise<GetScopeRbacRoleBindingResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:gkehub/v1alpha:getScopeRbacRoleBinding", {
        "location": args.location,
        "project": args.project,
        "rbacrolebindingId": args.rbacrolebindingId,
        "scopeId": args.scopeId,
    }, opts);
}

export interface GetScopeRbacRoleBindingArgs {
    location: string;
    project?: string;
    rbacrolebindingId: string;
    scopeId: string;
}

export interface GetScopeRbacRoleBindingResult {
    /**
     * When the rbacrolebinding was created.
     */
    readonly createTime: string;
    /**
     * When the rbacrolebinding was deleted.
     */
    readonly deleteTime: string;
    /**
     * group is the group, as seen by the kubernetes cluster.
     */
    readonly group: string;
    /**
     * Optional. Labels for this RBACRolebinding.
     */
    readonly labels: {[key: string]: string};
    /**
     * The resource name for the rbacrolebinding `projects/{project}/locations/{location}/scopes/{scope}/rbacrolebindings/{rbacrolebinding}` or `projects/{project}/locations/{location}/memberships/{membership}/rbacrolebindings/{rbacrolebinding}`
     */
    readonly name: string;
    /**
     * Role to bind to the principal
     */
    readonly role: outputs.gkehub.v1alpha.RoleResponse;
    /**
     * State of the rbacrolebinding resource.
     */
    readonly state: outputs.gkehub.v1alpha.RBACRoleBindingLifecycleStateResponse;
    /**
     * Google-generated UUID for this resource. This is unique across all rbacrolebinding resources. If a rbacrolebinding resource is deleted and another resource with the same name is created, it gets a different uid.
     */
    readonly uid: string;
    /**
     * When the rbacrolebinding was last updated.
     */
    readonly updateTime: string;
    /**
     * user is the name of the user as seen by the kubernetes cluster, example "alice" or "alice@domain.tld"
     */
    readonly user: string;
}
/**
 * Returns the details of a Scope RBACRoleBinding.
 */
export function getScopeRbacRoleBindingOutput(args: GetScopeRbacRoleBindingOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetScopeRbacRoleBindingResult> {
    return pulumi.output(args).apply((a: any) => getScopeRbacRoleBinding(a, opts))
}

export interface GetScopeRbacRoleBindingOutputArgs {
    location: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    rbacrolebindingId: pulumi.Input<string>;
    scopeId: pulumi.Input<string>;
}