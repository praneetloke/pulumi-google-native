// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Returns the specified TargetHttpProxy resource.
 */
export function getTargetHttpProxy(args: GetTargetHttpProxyArgs, opts?: pulumi.InvokeOptions): Promise<GetTargetHttpProxyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:compute/v1:getTargetHttpProxy", {
        "project": args.project,
        "targetHttpProxy": args.targetHttpProxy,
    }, opts);
}

export interface GetTargetHttpProxyArgs {
    project?: string;
    targetHttpProxy: string;
}

export interface GetTargetHttpProxyResult {
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    /**
     * Fingerprint of this resource. A hash of the contents stored in this object. This field is used in optimistic locking. This field will be ignored when inserting a TargetHttpProxy. An up-to-date fingerprint must be provided in order to patch/update the TargetHttpProxy; otherwise, the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve the TargetHttpProxy.
     */
    readonly fingerprint: string;
    /**
     * Type of resource. Always compute#targetHttpProxy for target HTTP proxies.
     */
    readonly kind: string;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name: string;
    /**
     * This field only applies when the forwarding rule that references this target proxy has a loadBalancingScheme set to INTERNAL_SELF_MANAGED. When this field is set to true, Envoy proxies set up inbound traffic interception and bind to the IP address and port specified in the forwarding rule. This is generally useful when using Traffic Director to configure Envoy as a gateway or middle proxy (in other words, not a sidecar proxy). The Envoy proxy listens for inbound requests and handles requests when it receives them. The default is false.
     */
    readonly proxyBind: boolean;
    /**
     * URL of the region where the regional Target HTTP Proxy resides. This field is not applicable to global Target HTTP Proxies.
     */
    readonly region: string;
    /**
     * Server-defined URL for the resource.
     */
    readonly selfLink: string;
    /**
     * URL to the UrlMap resource that defines the mapping from URL to the BackendService.
     */
    readonly urlMap: string;
}
/**
 * Returns the specified TargetHttpProxy resource.
 */
export function getTargetHttpProxyOutput(args: GetTargetHttpProxyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTargetHttpProxyResult> {
    return pulumi.output(args).apply((a: any) => getTargetHttpProxy(a, opts))
}

export interface GetTargetHttpProxyOutputArgs {
    project?: pulumi.Input<string>;
    targetHttpProxy: pulumi.Input<string>;
}
