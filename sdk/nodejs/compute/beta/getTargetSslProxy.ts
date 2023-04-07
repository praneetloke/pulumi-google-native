// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Returns the specified TargetSslProxy resource.
 */
export function getTargetSslProxy(args: GetTargetSslProxyArgs, opts?: pulumi.InvokeOptions): Promise<GetTargetSslProxyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:compute/beta:getTargetSslProxy", {
        "project": args.project,
        "targetSslProxy": args.targetSslProxy,
    }, opts);
}

export interface GetTargetSslProxyArgs {
    project?: string;
    targetSslProxy: string;
}

export interface GetTargetSslProxyResult {
    /**
     * URL of a certificate map that identifies a certificate map associated with the given target proxy. This field can only be set for global target proxies. If set, sslCertificates will be ignored.
     */
    readonly certificateMap: string;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    /**
     * Type of the resource. Always compute#targetSslProxy for target SSL proxies.
     */
    readonly kind: string;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name: string;
    /**
     * Specifies the type of proxy header to append before sending data to the backend, either NONE or PROXY_V1. The default is NONE.
     */
    readonly proxyHeader: string;
    /**
     * Server-defined URL for the resource.
     */
    readonly selfLink: string;
    /**
     * URL to the BackendService resource.
     */
    readonly service: string;
    /**
     * URLs to SslCertificate resources that are used to authenticate connections to Backends. At least one SSL certificate must be specified. Currently, you may specify up to 15 SSL certificates. sslCertificates do not apply when the load balancing scheme is set to INTERNAL_SELF_MANAGED.
     */
    readonly sslCertificates: string[];
    /**
     * URL of SslPolicy resource that will be associated with the TargetSslProxy resource. If not set, the TargetSslProxy resource will not have any SSL policy configured.
     */
    readonly sslPolicy: string;
}
/**
 * Returns the specified TargetSslProxy resource.
 */
export function getTargetSslProxyOutput(args: GetTargetSslProxyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetTargetSslProxyResult> {
    return pulumi.output(args).apply((a: any) => getTargetSslProxy(a, opts))
}

export interface GetTargetSslProxyOutputArgs {
    project?: pulumi.Input<string>;
    targetSslProxy: pulumi.Input<string>;
}
