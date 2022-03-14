// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates an Apigee organization. See [Create an Apigee organization](https://cloud.google.com/apigee/docs/api-platform/get-started/create-org).
 * Auto-naming is currently not supported for this resource.
 */
export class Organization extends pulumi.CustomResource {
    /**
     * Get an existing Organization resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Organization {
        return new Organization(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:apigee/v1:Organization';

    /**
     * Returns true if the given object is an instance of Organization.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Organization {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Organization.__pulumiType;
    }

    /**
     * Addon configurations of the Apigee organization.
     */
    public readonly addonsConfig!: pulumi.Output<outputs.apigee.v1.GoogleCloudApigeeV1AddonsConfigResponse>;
    /**
     * Not used by Apigee.
     */
    public readonly attributes!: pulumi.Output<string[]>;
    /**
     * Compute Engine network used for Service Networking to be peered with Apigee runtime instances. See [Getting started with the Service Networking API](https://cloud.google.com/service-infrastructure/docs/service-networking/getting-started). Valid only when [RuntimeType](#RuntimeType) is set to `CLOUD`. The value must be set before the creation of a runtime instance and can be updated only when there are no runtime instances. For example: `default`. Apigee also supports shared VPC (that is, the host network project is not the same as the one that is peering with Apigee). See [Shared VPC overview](https://cloud.google.com/vpc/docs/shared-vpc). To use a shared VPC network, use the following format: `projects/{host-project-id}/{region}/networks/{network-name}`. For example: `projects/my-sharedvpc-host/global/networks/mynetwork` **Note:** Not supported for Apigee hybrid.
     */
    public readonly authorizedNetwork!: pulumi.Output<string>;
    /**
     * Billing type of the Apigee organization. See [Apigee pricing](https://cloud.google.com/apigee/pricing).
     */
    public readonly billingType!: pulumi.Output<string>;
    /**
     * Base64-encoded public certificate for the root CA of the Apigee organization. Valid only when [RuntimeType](#RuntimeType) is `CLOUD`.
     */
    public /*out*/ readonly caCertificate!: pulumi.Output<string>;
    /**
     * Time that the Apigee organization was created in milliseconds since epoch.
     */
    public /*out*/ readonly createdAt!: pulumi.Output<string>;
    /**
     * Not used by Apigee.
     */
    public readonly customerName!: pulumi.Output<string>;
    /**
     * Description of the Apigee organization.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Display name for the Apigee organization. Unused, but reserved for future use.
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * List of environments in the Apigee organization.
     */
    public /*out*/ readonly environments!: pulumi.Output<string[]>;
    /**
     * Time that the Apigee organization is scheduled for deletion.
     */
    public /*out*/ readonly expiresAt!: pulumi.Output<string>;
    /**
     * Time that the Apigee organization was last modified in milliseconds since epoch.
     */
    public /*out*/ readonly lastModifiedAt!: pulumi.Output<string>;
    /**
     * Name of the Apigee organization.
     */
    public /*out*/ readonly name!: pulumi.Output<string>;
    /**
     * Configuration for the Portals settings.
     */
    public readonly portalDisabled!: pulumi.Output<boolean>;
    /**
     * Project ID associated with the Apigee organization.
     */
    public /*out*/ readonly project!: pulumi.Output<string>;
    /**
     * Properties defined in the Apigee organization profile.
     */
    public readonly properties!: pulumi.Output<outputs.apigee.v1.GoogleCloudApigeeV1PropertiesResponse>;
    /**
     * Cloud KMS key name used for encrypting the data that is stored and replicated across runtime instances. Update is not allowed after the organization is created. Required when [RuntimeType](#RuntimeType) is `CLOUD`. If not specified when [RuntimeType](#RuntimeType) is `TRIAL`, a Google-Managed encryption key will be used. For example: "projects/foo/locations/us/keyRings/bar/cryptoKeys/baz". **Note:** Not supported for Apigee hybrid.
     */
    public readonly runtimeDatabaseEncryptionKeyName!: pulumi.Output<string>;
    /**
     * Runtime type of the Apigee organization based on the Apigee subscription purchased.
     */
    public readonly runtimeType!: pulumi.Output<string>;
    /**
     * State of the organization. Values other than ACTIVE means the resource is not ready to use.
     */
    public /*out*/ readonly state!: pulumi.Output<string>;
    /**
     * Not used by Apigee.
     */
    public readonly type!: pulumi.Output<string>;

    /**
     * Create a Organization resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: OrganizationArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.parent === undefined) && !opts.urn) {
                throw new Error("Missing required property 'parent'");
            }
            if ((!args || args.runtimeType === undefined) && !opts.urn) {
                throw new Error("Missing required property 'runtimeType'");
            }
            resourceInputs["addonsConfig"] = args ? args.addonsConfig : undefined;
            resourceInputs["attributes"] = args ? args.attributes : undefined;
            resourceInputs["authorizedNetwork"] = args ? args.authorizedNetwork : undefined;
            resourceInputs["billingType"] = args ? args.billingType : undefined;
            resourceInputs["customerName"] = args ? args.customerName : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["parent"] = args ? args.parent : undefined;
            resourceInputs["portalDisabled"] = args ? args.portalDisabled : undefined;
            resourceInputs["properties"] = args ? args.properties : undefined;
            resourceInputs["runtimeDatabaseEncryptionKeyName"] = args ? args.runtimeDatabaseEncryptionKeyName : undefined;
            resourceInputs["runtimeType"] = args ? args.runtimeType : undefined;
            resourceInputs["type"] = args ? args.type : undefined;
            resourceInputs["caCertificate"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["environments"] = undefined /*out*/;
            resourceInputs["expiresAt"] = undefined /*out*/;
            resourceInputs["lastModifiedAt"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
        } else {
            resourceInputs["addonsConfig"] = undefined /*out*/;
            resourceInputs["attributes"] = undefined /*out*/;
            resourceInputs["authorizedNetwork"] = undefined /*out*/;
            resourceInputs["billingType"] = undefined /*out*/;
            resourceInputs["caCertificate"] = undefined /*out*/;
            resourceInputs["createdAt"] = undefined /*out*/;
            resourceInputs["customerName"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["environments"] = undefined /*out*/;
            resourceInputs["expiresAt"] = undefined /*out*/;
            resourceInputs["lastModifiedAt"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["portalDisabled"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["properties"] = undefined /*out*/;
            resourceInputs["runtimeDatabaseEncryptionKeyName"] = undefined /*out*/;
            resourceInputs["runtimeType"] = undefined /*out*/;
            resourceInputs["state"] = undefined /*out*/;
            resourceInputs["type"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Organization.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Organization resource.
 */
export interface OrganizationArgs {
    /**
     * Addon configurations of the Apigee organization.
     */
    addonsConfig?: pulumi.Input<inputs.apigee.v1.GoogleCloudApigeeV1AddonsConfigArgs>;
    /**
     * Not used by Apigee.
     */
    attributes?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Compute Engine network used for Service Networking to be peered with Apigee runtime instances. See [Getting started with the Service Networking API](https://cloud.google.com/service-infrastructure/docs/service-networking/getting-started). Valid only when [RuntimeType](#RuntimeType) is set to `CLOUD`. The value must be set before the creation of a runtime instance and can be updated only when there are no runtime instances. For example: `default`. Apigee also supports shared VPC (that is, the host network project is not the same as the one that is peering with Apigee). See [Shared VPC overview](https://cloud.google.com/vpc/docs/shared-vpc). To use a shared VPC network, use the following format: `projects/{host-project-id}/{region}/networks/{network-name}`. For example: `projects/my-sharedvpc-host/global/networks/mynetwork` **Note:** Not supported for Apigee hybrid.
     */
    authorizedNetwork?: pulumi.Input<string>;
    /**
     * Billing type of the Apigee organization. See [Apigee pricing](https://cloud.google.com/apigee/pricing).
     */
    billingType?: pulumi.Input<enums.apigee.v1.OrganizationBillingType>;
    /**
     * Not used by Apigee.
     */
    customerName?: pulumi.Input<string>;
    /**
     * Description of the Apigee organization.
     */
    description?: pulumi.Input<string>;
    /**
     * Display name for the Apigee organization. Unused, but reserved for future use.
     */
    displayName?: pulumi.Input<string>;
    /**
     * Required. Name of the GCP project in which to associate the Apigee organization. Pass the information as a query parameter using the following structure in your request: `projects/`
     */
    parent: pulumi.Input<string>;
    /**
     * Configuration for the Portals settings.
     */
    portalDisabled?: pulumi.Input<boolean>;
    /**
     * Properties defined in the Apigee organization profile.
     */
    properties?: pulumi.Input<inputs.apigee.v1.GoogleCloudApigeeV1PropertiesArgs>;
    /**
     * Cloud KMS key name used for encrypting the data that is stored and replicated across runtime instances. Update is not allowed after the organization is created. Required when [RuntimeType](#RuntimeType) is `CLOUD`. If not specified when [RuntimeType](#RuntimeType) is `TRIAL`, a Google-Managed encryption key will be used. For example: "projects/foo/locations/us/keyRings/bar/cryptoKeys/baz". **Note:** Not supported for Apigee hybrid.
     */
    runtimeDatabaseEncryptionKeyName?: pulumi.Input<string>;
    /**
     * Runtime type of the Apigee organization based on the Apigee subscription purchased.
     */
    runtimeType: pulumi.Input<enums.apigee.v1.OrganizationRuntimeType>;
    /**
     * Not used by Apigee.
     */
    type?: pulumi.Input<enums.apigee.v1.OrganizationType>;
}
