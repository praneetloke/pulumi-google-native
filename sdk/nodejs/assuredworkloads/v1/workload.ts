// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates Assured Workload.
 */
export class Workload extends pulumi.CustomResource {
    /**
     * Get an existing Workload resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Workload {
        return new Workload(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:assuredworkloads/v1:Workload';

    /**
     * Returns true if the given object is an instance of Workload.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Workload {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Workload.__pulumiType;
    }

    /**
     * Optional. The billing account used for the resources which are direct children of workload. This billing account is initially associated with the resources created as part of Workload creation. After the initial creation of these resources, the customer can change the assigned billing account. The resource name has the form `billingAccounts/{billing_account_id}`. For example, `billingAccounts/012345-567890-ABCDEF`.
     */
    public readonly billingAccount!: pulumi.Output<string>;
    /**
     * Immutable. Compliance Regime associated with this workload.
     */
    public readonly complianceRegime!: pulumi.Output<string>;
    /**
     * Count of active Violations in the Workload.
     */
    public /*out*/ readonly complianceStatus!: pulumi.Output<outputs.assuredworkloads.v1.GoogleCloudAssuredworkloadsV1WorkloadComplianceStatusResponse>;
    /**
     * Urls for services which are compliant for this Assured Workload, but which are currently disallowed by the ResourceUsageRestriction org policy. Invoke RestrictAllowedResources endpoint to allow your project developers to use these services in their environment."
     */
    public /*out*/ readonly compliantButDisallowedServices!: pulumi.Output<string[]>;
    /**
     * Immutable. The Workload creation timestamp.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * The user-assigned display name of the Workload. When present it must be between 4 to 30 characters. Allowed characters are: lowercase and uppercase letters, numbers, hyphen, and spaces. Example: My Workload
     */
    public readonly displayName!: pulumi.Output<string>;
    /**
     * Optional. Indicates the sovereignty status of the given workload. Currently meant to be used by Europe/Canada customers.
     */
    public readonly enableSovereignControls!: pulumi.Output<boolean>;
    /**
     * Optional. ETag of the workload, it is calculated on the basis of the Workload contents. It will be used in Update & Delete operations.
     */
    public readonly etag!: pulumi.Output<string>;
    /**
     * Optional. A identifier associated with the workload and underlying projects which allows for the break down of billing costs for a workload. The value provided for the identifier will add a label to the workload and contained projects with the identifier as the value.
     */
    public readonly externalId!: pulumi.Output<string | undefined>;
    /**
     * Represents the KAJ enrollment state of the given workload.
     */
    public /*out*/ readonly kajEnrollmentState!: pulumi.Output<string>;
    /**
     * Input only. Settings used to create a CMEK crypto key. When set, a project with a KMS CMEK key is provisioned. This field is deprecated as of Feb 28, 2022. In order to create a Keyring, callers should specify, ENCRYPTION_KEYS_PROJECT or KEYRING in ResourceSettings.resource_type field.
     *
     * @deprecated Input only. Settings used to create a CMEK crypto key. When set, a project with a KMS CMEK key is provisioned. This field is deprecated as of Feb 28, 2022. In order to create a Keyring, callers should specify, ENCRYPTION_KEYS_PROJECT or KEYRING in ResourceSettings.resource_type field.
     */
    public readonly kmsSettings!: pulumi.Output<outputs.assuredworkloads.v1.GoogleCloudAssuredworkloadsV1WorkloadKMSSettingsResponse>;
    /**
     * Optional. Labels applied to the workload.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    public readonly location!: pulumi.Output<string>;
    /**
     * Optional. The resource name of the workload. Format: organizations/{organization}/locations/{location}/workloads/{workload} Read-only.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly organizationId!: pulumi.Output<string>;
    /**
     * Optional. Compliance Regime associated with this workload.
     */
    public readonly partner!: pulumi.Output<string>;
    /**
     * Input only. The parent resource for the resources managed by this Assured Workload. May be either empty or a folder resource which is a child of the Workload parent. If not specified all resources are created under the parent organization. Format: folders/{folder_id}
     */
    public readonly provisionedResourcesParent!: pulumi.Output<string>;
    /**
     * Input only. Resource properties that are used to customize workload resources. These properties (such as custom project id) will be used to create workload resources if possible. This field is optional.
     */
    public readonly resourceSettings!: pulumi.Output<outputs.assuredworkloads.v1.GoogleCloudAssuredworkloadsV1WorkloadResourceSettingsResponse[]>;
    /**
     * The resources associated with this workload. These resources will be created when creating the workload. If any of the projects already exist, the workload creation will fail. Always read only.
     */
    public /*out*/ readonly resources!: pulumi.Output<outputs.assuredworkloads.v1.GoogleCloudAssuredworkloadsV1WorkloadResourceInfoResponse[]>;
    /**
     * Represents the SAA enrollment response of the given workload. SAA enrollment response is queried during GetWorkload call. In failure cases, user friendly error message is shown in SAA details page.
     */
    public /*out*/ readonly saaEnrollmentResponse!: pulumi.Output<outputs.assuredworkloads.v1.GoogleCloudAssuredworkloadsV1WorkloadSaaEnrollmentResponseResponse>;

    /**
     * Create a Workload resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: WorkloadArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.complianceRegime === undefined) && !opts.urn) {
                throw new Error("Missing required property 'complianceRegime'");
            }
            if ((!args || args.displayName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'displayName'");
            }
            if ((!args || args.organizationId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'organizationId'");
            }
            resourceInputs["billingAccount"] = args ? args.billingAccount : undefined;
            resourceInputs["complianceRegime"] = args ? args.complianceRegime : undefined;
            resourceInputs["displayName"] = args ? args.displayName : undefined;
            resourceInputs["enableSovereignControls"] = args ? args.enableSovereignControls : undefined;
            resourceInputs["etag"] = args ? args.etag : undefined;
            resourceInputs["externalId"] = args ? args.externalId : undefined;
            resourceInputs["kmsSettings"] = args ? args.kmsSettings : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["organizationId"] = args ? args.organizationId : undefined;
            resourceInputs["partner"] = args ? args.partner : undefined;
            resourceInputs["provisionedResourcesParent"] = args ? args.provisionedResourcesParent : undefined;
            resourceInputs["resourceSettings"] = args ? args.resourceSettings : undefined;
            resourceInputs["complianceStatus"] = undefined /*out*/;
            resourceInputs["compliantButDisallowedServices"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["kajEnrollmentState"] = undefined /*out*/;
            resourceInputs["resources"] = undefined /*out*/;
            resourceInputs["saaEnrollmentResponse"] = undefined /*out*/;
        } else {
            resourceInputs["billingAccount"] = undefined /*out*/;
            resourceInputs["complianceRegime"] = undefined /*out*/;
            resourceInputs["complianceStatus"] = undefined /*out*/;
            resourceInputs["compliantButDisallowedServices"] = undefined /*out*/;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["displayName"] = undefined /*out*/;
            resourceInputs["enableSovereignControls"] = undefined /*out*/;
            resourceInputs["etag"] = undefined /*out*/;
            resourceInputs["externalId"] = undefined /*out*/;
            resourceInputs["kajEnrollmentState"] = undefined /*out*/;
            resourceInputs["kmsSettings"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["location"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["organizationId"] = undefined /*out*/;
            resourceInputs["partner"] = undefined /*out*/;
            resourceInputs["provisionedResourcesParent"] = undefined /*out*/;
            resourceInputs["resourceSettings"] = undefined /*out*/;
            resourceInputs["resources"] = undefined /*out*/;
            resourceInputs["saaEnrollmentResponse"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["location", "organizationId"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(Workload.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a Workload resource.
 */
export interface WorkloadArgs {
    /**
     * Optional. The billing account used for the resources which are direct children of workload. This billing account is initially associated with the resources created as part of Workload creation. After the initial creation of these resources, the customer can change the assigned billing account. The resource name has the form `billingAccounts/{billing_account_id}`. For example, `billingAccounts/012345-567890-ABCDEF`.
     */
    billingAccount?: pulumi.Input<string>;
    /**
     * Immutable. Compliance Regime associated with this workload.
     */
    complianceRegime: pulumi.Input<enums.assuredworkloads.v1.WorkloadComplianceRegime>;
    /**
     * The user-assigned display name of the Workload. When present it must be between 4 to 30 characters. Allowed characters are: lowercase and uppercase letters, numbers, hyphen, and spaces. Example: My Workload
     */
    displayName: pulumi.Input<string>;
    /**
     * Optional. Indicates the sovereignty status of the given workload. Currently meant to be used by Europe/Canada customers.
     */
    enableSovereignControls?: pulumi.Input<boolean>;
    /**
     * Optional. ETag of the workload, it is calculated on the basis of the Workload contents. It will be used in Update & Delete operations.
     */
    etag?: pulumi.Input<string>;
    /**
     * Optional. A identifier associated with the workload and underlying projects which allows for the break down of billing costs for a workload. The value provided for the identifier will add a label to the workload and contained projects with the identifier as the value.
     */
    externalId?: pulumi.Input<string>;
    /**
     * Input only. Settings used to create a CMEK crypto key. When set, a project with a KMS CMEK key is provisioned. This field is deprecated as of Feb 28, 2022. In order to create a Keyring, callers should specify, ENCRYPTION_KEYS_PROJECT or KEYRING in ResourceSettings.resource_type field.
     *
     * @deprecated Input only. Settings used to create a CMEK crypto key. When set, a project with a KMS CMEK key is provisioned. This field is deprecated as of Feb 28, 2022. In order to create a Keyring, callers should specify, ENCRYPTION_KEYS_PROJECT or KEYRING in ResourceSettings.resource_type field.
     */
    kmsSettings?: pulumi.Input<inputs.assuredworkloads.v1.GoogleCloudAssuredworkloadsV1WorkloadKMSSettingsArgs>;
    /**
     * Optional. Labels applied to the workload.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * Optional. The resource name of the workload. Format: organizations/{organization}/locations/{location}/workloads/{workload} Read-only.
     */
    name?: pulumi.Input<string>;
    organizationId: pulumi.Input<string>;
    /**
     * Optional. Compliance Regime associated with this workload.
     */
    partner?: pulumi.Input<enums.assuredworkloads.v1.WorkloadPartner>;
    /**
     * Input only. The parent resource for the resources managed by this Assured Workload. May be either empty or a folder resource which is a child of the Workload parent. If not specified all resources are created under the parent organization. Format: folders/{folder_id}
     */
    provisionedResourcesParent?: pulumi.Input<string>;
    /**
     * Input only. Resource properties that are used to customize workload resources. These properties (such as custom project id) will be used to create workload resources if possible. This field is optional.
     */
    resourceSettings?: pulumi.Input<pulumi.Input<inputs.assuredworkloads.v1.GoogleCloudAssuredworkloadsV1WorkloadResourceSettingsArgs>[]>;
}
