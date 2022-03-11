// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a new DnsAuthorization in a given project and location.
 */
export class DnsAuthorization extends pulumi.CustomResource {
    /**
     * Get an existing DnsAuthorization resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): DnsAuthorization {
        return new DnsAuthorization(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:certificatemanager/v1:DnsAuthorization';

    /**
     * Returns true if the given object is an instance of DnsAuthorization.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DnsAuthorization {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DnsAuthorization.__pulumiType;
    }

    /**
     * The creation timestamp of a DnsAuthorization.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * One or more paragraphs of text description of a DnsAuthorization.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * DNS Resource Record that needs to be added to DNS configuration.
     */
    public /*out*/ readonly dnsResourceRecord!: pulumi.Output<outputs.certificatemanager.v1.DnsResourceRecordResponse>;
    /**
     * Immutable. A domain which is being authorized. A DnsAuthorization resource covers a single domain and its wildcard, e.g. authorization for `example.com` can be used to issue certificates for `example.com` and `*.example.com`.
     */
    public readonly domain!: pulumi.Output<string>;
    /**
     * Set of labels associated with a DnsAuthorization.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * A user-defined name of the dns authorization. DnsAuthorization names must be unique globally and match pattern `projects/*&#47;locations/*&#47;dnsAuthorizations/*`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The last update timestamp of a DnsAuthorization.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a DnsAuthorization resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DnsAuthorizationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.dnsAuthorizationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'dnsAuthorizationId'");
            }
            if ((!args || args.domain === undefined) && !opts.urn) {
                throw new Error("Missing required property 'domain'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["dnsAuthorizationId"] = args ? args.dnsAuthorizationId : undefined;
            resourceInputs["domain"] = args ? args.domain : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["dnsResourceRecord"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["dnsResourceRecord"] = undefined /*out*/;
            resourceInputs["domain"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DnsAuthorization.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a DnsAuthorization resource.
 */
export interface DnsAuthorizationArgs {
    /**
     * One or more paragraphs of text description of a DnsAuthorization.
     */
    description?: pulumi.Input<string>;
    dnsAuthorizationId: pulumi.Input<string>;
    /**
     * Immutable. A domain which is being authorized. A DnsAuthorization resource covers a single domain and its wildcard, e.g. authorization for `example.com` can be used to issue certificates for `example.com` and `*.example.com`.
     */
    domain: pulumi.Input<string>;
    /**
     * Set of labels associated with a DnsAuthorization.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * A user-defined name of the dns authorization. DnsAuthorization names must be unique globally and match pattern `projects/*&#47;locations/*&#47;dnsAuthorizations/*`.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
}
