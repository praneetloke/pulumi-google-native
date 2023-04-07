// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * List all of the ordered rules present in a single specified policy.
 */
export function getOrganizationSecurityPolicy(args: GetOrganizationSecurityPolicyArgs, opts?: pulumi.InvokeOptions): Promise<GetOrganizationSecurityPolicyResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:compute/alpha:getOrganizationSecurityPolicy", {
        "securityPolicy": args.securityPolicy,
    }, opts);
}

export interface GetOrganizationSecurityPolicyArgs {
    securityPolicy: string;
}

export interface GetOrganizationSecurityPolicyResult {
    readonly adaptiveProtectionConfig: outputs.compute.alpha.SecurityPolicyAdaptiveProtectionConfigResponse;
    readonly advancedOptionsConfig: outputs.compute.alpha.SecurityPolicyAdvancedOptionsConfigResponse;
    /**
     * A list of associations that belong to this policy.
     */
    readonly associations: outputs.compute.alpha.SecurityPolicyAssociationResponse[];
    readonly cloudArmorConfig: outputs.compute.alpha.SecurityPolicyCloudArmorConfigResponse;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    readonly creationTimestamp: string;
    readonly ddosProtectionConfig: outputs.compute.alpha.SecurityPolicyDdosProtectionConfigResponse;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    readonly description: string;
    /**
     * User-provided name of the Organization security plicy. The name should be unique in the organization in which the security policy is created. This should only be used when SecurityPolicyType is FIREWALL. The name must be 1-63 characters long, and comply with https://www.ietf.org/rfc/rfc1035.txt. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly displayName: string;
    /**
     * Specifies a fingerprint for this resource, which is essentially a hash of the metadata's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update metadata. You must always provide an up-to-date fingerprint hash in order to update or change metadata, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make get() request to the security policy.
     */
    readonly fingerprint: string;
    /**
     * [Output only] Type of the resource. Always compute#securityPolicyfor security policies
     */
    readonly kind: string;
    /**
     * A fingerprint for the labels being applied to this security policy, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels. To see the latest fingerprint, make get() request to the security policy.
     */
    readonly labelFingerprint: string;
    /**
     * Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
     */
    readonly labels: {[key: string]: string};
    /**
     * Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    readonly name: string;
    /**
     * The parent of the security policy.
     */
    readonly parent: string;
    readonly recaptchaOptionsConfig: outputs.compute.alpha.SecurityPolicyRecaptchaOptionsConfigResponse;
    /**
     * URL of the region where the regional security policy resides. This field is not applicable to global security policies.
     */
    readonly region: string;
    /**
     * Total count of all security policy rule tuples. A security policy can not exceed a set number of tuples.
     */
    readonly ruleTupleCount: number;
    /**
     * A list of rules that belong to this policy. There must always be a default rule which is a rule with priority 2147483647 and match all condition (for the match condition this means match "*" for srcIpRanges and for the networkMatch condition every field must be either match "*" or not set). If no rules are provided when creating a security policy, a default rule with action "allow" will be added.
     */
    readonly rules: outputs.compute.alpha.SecurityPolicyRuleResponse[];
    /**
     * Server-defined URL for the resource.
     */
    readonly selfLink: string;
    /**
     * Server-defined URL for this resource with the resource id.
     */
    readonly selfLinkWithId: string;
    /**
     * The type indicates the intended use of the security policy. - CLOUD_ARMOR: Cloud Armor backend security policies can be configured to filter incoming HTTP requests targeting backend services. They filter requests before they hit the origin servers. - CLOUD_ARMOR_EDGE: Cloud Armor edge security policies can be configured to filter incoming HTTP requests targeting backend services (including Cloud CDN-enabled) as well as backend buckets (Cloud Storage). They filter requests before the request is served from Google's cache. - CLOUD_ARMOR_INTERNAL_SERVICE: Cloud Armor internal service policies can be configured to filter HTTP requests targeting services managed by Traffic Director in a service mesh. They filter requests before the request is served from the application. - CLOUD_ARMOR_NETWORK: Cloud Armor network policies can be configured to filter packets targeting network load balancing resources such as backend services, target pools, target instances, and instances with external IPs. They filter requests before the request is served from the application. This field can be set only at resource creation time.
     */
    readonly type: string;
    /**
     * Definitions of user-defined fields for CLOUD_ARMOR_NETWORK policies. A user-defined field consists of up to 4 bytes extracted from a fixed offset in the packet, relative to the IPv4, IPv6, TCP, or UDP header, with an optional mask to select certain bits. Rules may then specify matching values for these fields. Example: userDefinedFields: - name: "ipv4_fragment_offset" base: IPV4 offset: 6 size: 2 mask: "0x1fff"
     */
    readonly userDefinedFields: outputs.compute.alpha.SecurityPolicyUserDefinedFieldResponse[];
}
/**
 * List all of the ordered rules present in a single specified policy.
 */
export function getOrganizationSecurityPolicyOutput(args: GetOrganizationSecurityPolicyOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetOrganizationSecurityPolicyResult> {
    return pulumi.output(args).apply((a: any) => getOrganizationSecurityPolicy(a, opts))
}

export interface GetOrganizationSecurityPolicyOutputArgs {
    securityPolicy: pulumi.Input<string>;
}
