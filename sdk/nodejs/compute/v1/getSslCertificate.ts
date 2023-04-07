// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Returns the specified SslCertificate resource.
 */
export function getSslCertificate(args: GetSslCertificateArgs, opts?: pulumi.InvokeOptions): Promise<GetSslCertificateResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:compute/v1:getSslCertificate", {
        "project": args.project,
        "sslCertificate": args.sslCertificate,
    }, opts);
}

export interface GetSslCertificateArgs {
    project?: string;
    sslCertificate: string;
}

export interface GetSslCertificateResult {
    /**
     * A value read into memory from a certificate file. The certificate file must be in PEM format. The certificate chain must be no greater than 5 certs long. The chain must include at least one intermediate cert.
     */
    readonly certificate: string;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    /**
     * Expire time of the certificate. RFC3339
     */
    readonly expireTime: string;
    /**
     * Type of the resource. Always compute#sslCertificate for SSL certificates.
     */
    readonly kind: string;
    /**
     * Configuration and status of a managed SSL certificate.
     */
    readonly managed: outputs.compute.v1.SslCertificateManagedSslCertificateResponse;
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name: string;
    /**
     * A value read into memory from a write-only private key file. The private key file must be in PEM format. For security, only insert requests include this field.
     */
    readonly privateKey: string;
    /**
     * URL of the region where the regional SSL Certificate resides. This field is not applicable to global SSL Certificate.
     */
    readonly region: string;
    /**
     * [Output only] Server-defined URL for the resource.
     */
    readonly selfLink: string;
    /**
     * Configuration and status of a self-managed SSL certificate.
     */
    readonly selfManaged: outputs.compute.v1.SslCertificateSelfManagedSslCertificateResponse;
    /**
     * Domains associated with the certificate via Subject Alternative Name.
     */
    readonly subjectAlternativeNames: string[];
    /**
     * (Optional) Specifies the type of SSL certificate, either "SELF_MANAGED" or "MANAGED". If not specified, the certificate is self-managed and the fields certificate and private_key are used.
     */
    readonly type: string;
}
/**
 * Returns the specified SslCertificate resource.
 */
export function getSslCertificateOutput(args: GetSslCertificateOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetSslCertificateResult> {
    return pulumi.output(args).apply((a: any) => getSslCertificate(a, opts))
}

export interface GetSslCertificateOutputArgs {
    project?: pulumi.Input<string>;
    sslCertificate: pulumi.Input<string>;
}
