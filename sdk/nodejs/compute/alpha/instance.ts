// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs, enums } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates an instance resource in the specified project using the data included in the request.
 */
export class Instance extends pulumi.CustomResource {
    /**
     * Get an existing Instance resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): Instance {
        return new Instance(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/alpha:Instance';

    /**
     * Returns true if the given object is an instance of Instance.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Instance {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Instance.__pulumiType;
    }

    /**
     * Controls for advanced machine-related behavior features.
     */
    public readonly advancedMachineFeatures!: pulumi.Output<outputs.compute.alpha.AdvancedMachineFeaturesResponse>;
    /**
     * Allows this instance to send and receive packets with non-matching destination or source IPs. This is required if you plan to use this instance to forward routes. For more information, see Enabling IP Forwarding .
     */
    public readonly canIpForward!: pulumi.Output<boolean>;
    public readonly confidentialInstanceConfig!: pulumi.Output<outputs.compute.alpha.ConfidentialInstanceConfigResponse>;
    /**
     * The CPU platform used by this instance.
     */
    public /*out*/ readonly cpuPlatform!: pulumi.Output<string>;
    /**
     * Creation timestamp in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * Whether the resource should be protected against deletion.
     */
    public readonly deletionProtection!: pulumi.Output<boolean>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Array of disks associated with this instance. Persistent disks must be created before you can assign them.
     */
    public readonly disks!: pulumi.Output<outputs.compute.alpha.AttachedDiskResponse[]>;
    /**
     * Enables display device for the instance.
     */
    public readonly displayDevice!: pulumi.Output<outputs.compute.alpha.DisplayDeviceResponse>;
    /**
     * Specifies whether the disks restored from source snapshots or source machine image should erase Windows specific VSS signature.
     */
    public readonly eraseWindowsVssSignature!: pulumi.Output<boolean>;
    /**
     * Specifies a fingerprint for this resource, which is essentially a hash of the instance's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update the instance. You must always provide an up-to-date fingerprint hash in order to update the instance. To see the latest fingerprint, make get() request to the instance.
     */
    public /*out*/ readonly fingerprint!: pulumi.Output<string>;
    /**
     * A list of the type and count of accelerator cards attached to the instance.
     */
    public readonly guestAccelerators!: pulumi.Output<outputs.compute.alpha.AcceleratorConfigResponse[]>;
    /**
     * Specifies the hostname of the instance. The specified hostname must be RFC1035 compliant. If hostname is not specified, the default hostname is [INSTANCE_NAME].c.[PROJECT_ID].internal when using the global DNS, and [INSTANCE_NAME].[ZONE].c.[PROJECT_ID].internal when using zonal DNS.
     */
    public readonly hostname!: pulumi.Output<string>;
    /**
     * Encrypts or decrypts data for an instance with a customer-supplied encryption key. If you are creating a new instance, this field encrypts the local SSD and in-memory contents of the instance using a key that you provide. If you are restarting an instance protected with a customer-supplied encryption key, you must provide the correct key in order to successfully restart the instance. If you do not provide an encryption key when creating the instance, then the local SSD and in-memory contents will be encrypted using an automatically generated key and you do not need to provide a key to start the instance later. Instance templates do not store customer-supplied encryption keys, so you cannot use your own keys to encrypt local SSDs and in-memory content in a managed instance group.
     */
    public readonly instanceEncryptionKey!: pulumi.Output<outputs.compute.alpha.CustomerEncryptionKeyResponse>;
    /**
     * Type of the resource. Always compute#instance for instances.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * A fingerprint for this request, which is essentially a hash of the label's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels. To see the latest fingerprint, make get() request to the instance.
     */
    public /*out*/ readonly labelFingerprint!: pulumi.Output<string>;
    /**
     * Labels to apply to this instance. These can be later modified by the setLabels method.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Last start timestamp in RFC3339 text format.
     */
    public /*out*/ readonly lastStartTimestamp!: pulumi.Output<string>;
    /**
     * Last stop timestamp in RFC3339 text format.
     */
    public /*out*/ readonly lastStopTimestamp!: pulumi.Output<string>;
    /**
     * Last suspended timestamp in RFC3339 text format.
     */
    public /*out*/ readonly lastSuspendedTimestamp!: pulumi.Output<string>;
    /**
     * Full or partial URL of the machine type resource to use for this instance, in the format: zones/zone/machineTypes/machine-type. This is provided by the client when the instance is created. For example, the following is a valid partial url to a predefined machine type: zones/us-central1-f/machineTypes/n1-standard-1 To create a custom machine type, provide a URL to a machine type in the following format, where CPUS is 1 or an even number up to 32 (2, 4, 6, ... 24, etc), and MEMORY is the total memory for this instance. Memory must be a multiple of 256 MB and must be supplied in MB (e.g. 5 GB of memory is 5120 MB): zones/zone/machineTypes/custom-CPUS-MEMORY For example: zones/us-central1-f/machineTypes/custom-4-5120 For a full list of restrictions, read the Specifications for custom machine types.
     */
    public readonly machineType!: pulumi.Output<string>;
    /**
     * The metadata key/value pairs assigned to this instance. This includes custom metadata and predefined keys.
     */
    public readonly metadata!: pulumi.Output<outputs.compute.alpha.MetadataResponse>;
    /**
     * Specifies a minimum CPU platform for the VM instance. Applicable values are the friendly names of CPU platforms, such as minCpuPlatform: "Intel Haswell" or minCpuPlatform: "Intel Sandy Bridge".
     */
    public readonly minCpuPlatform!: pulumi.Output<string>;
    /**
     * The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * An array of network configurations for this instance. These specify how interfaces are configured to interact with other network services, such as connecting to the internet. Multiple interfaces are supported per instance.
     */
    public readonly networkInterfaces!: pulumi.Output<outputs.compute.alpha.NetworkInterfaceResponse[]>;
    public readonly networkPerformanceConfig!: pulumi.Output<outputs.compute.alpha.NetworkPerformanceConfigResponse>;
    /**
     * PostKeyRevocationActionType of the instance.
     */
    public readonly postKeyRevocationActionType!: pulumi.Output<string>;
    /**
     * Total amount of preserved state for SUSPENDED instances. Read-only in the api.
     */
    public readonly preservedStateSizeGb!: pulumi.Output<string>;
    /**
     * The private IPv6 google access type for the VM. If not specified, use INHERIT_FROM_SUBNETWORK as default.
     */
    public readonly privateIpv6GoogleAccess!: pulumi.Output<string>;
    /**
     * Specifies the reservations that this instance can consume from.
     */
    public readonly reservationAffinity!: pulumi.Output<outputs.compute.alpha.ReservationAffinityResponse>;
    /**
     * Resource policies applied to this instance.
     */
    public readonly resourcePolicies!: pulumi.Output<string[]>;
    /**
     * Specifies values set for instance attributes as compared to the values requested by user in the corresponding input only field.
     */
    public /*out*/ readonly resourceStatus!: pulumi.Output<outputs.compute.alpha.ResourceStatusResponse>;
    /**
     * Reserved for future use.
     */
    public /*out*/ readonly satisfiesPzs!: pulumi.Output<boolean>;
    /**
     * Sets the scheduling options for this instance.
     */
    public readonly scheduling!: pulumi.Output<outputs.compute.alpha.SchedulingResponse>;
    /**
     * [Input Only] Secure tags to apply to this instance. These can be later modified by the update method. Maximum number of secure tags allowed is 50.
     */
    public readonly secureTags!: pulumi.Output<string[]>;
    /**
     * Server-defined URL for this resource.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * Server-defined URL for this resource with the resource id.
     */
    public /*out*/ readonly selfLinkWithId!: pulumi.Output<string>;
    /**
     * A list of service accounts, with their specified scopes, authorized for this instance. Only one service account per VM instance is supported. Service accounts generate access tokens that can be accessed through the metadata server and used to authenticate applications on the instance. See Service Accounts for more information.
     */
    public readonly serviceAccounts!: pulumi.Output<outputs.compute.alpha.ServiceAccountResponse[]>;
    public readonly shieldedInstanceConfig!: pulumi.Output<outputs.compute.alpha.ShieldedInstanceConfigResponse>;
    public readonly shieldedInstanceIntegrityPolicy!: pulumi.Output<outputs.compute.alpha.ShieldedInstanceIntegrityPolicyResponse>;
    /**
     * Deprecating, please use shielded_instance_config.
     */
    public readonly shieldedVmConfig!: pulumi.Output<outputs.compute.alpha.ShieldedVmConfigResponse>;
    /**
     * Deprecating, please use shielded_instance_integrity_policy.
     */
    public readonly shieldedVmIntegrityPolicy!: pulumi.Output<outputs.compute.alpha.ShieldedVmIntegrityPolicyResponse>;
    /**
     * Source machine image
     */
    public readonly sourceMachineImage!: pulumi.Output<string>;
    /**
     * Source machine image encryption key when creating an instance from a machine image.
     */
    public readonly sourceMachineImageEncryptionKey!: pulumi.Output<outputs.compute.alpha.CustomerEncryptionKeyResponse>;
    /**
     * Whether a VM has been restricted for start because Compute Engine has detected suspicious activity.
     */
    public /*out*/ readonly startRestricted!: pulumi.Output<boolean>;
    /**
     * The status of the instance. One of the following values: PROVISIONING, STAGING, RUNNING, STOPPING, SUSPENDING, SUSPENDED, REPAIRING, and TERMINATED. For more information about the status of the instance, see Instance life cycle.
     */
    public /*out*/ readonly status!: pulumi.Output<string>;
    /**
     * An optional, human-readable explanation of the status.
     */
    public /*out*/ readonly statusMessage!: pulumi.Output<string>;
    /**
     * Tags to apply to this instance. Tags are used to identify valid sources or targets for network firewalls and are specified by the client during instance creation. The tags can be later modified by the setTags method. Each tag within the list must comply with RFC1035. Multiple tags can be specified via the 'tags.items' field.
     */
    public readonly tags!: pulumi.Output<outputs.compute.alpha.TagsResponse>;
    /**
     * Specifies upcoming maintenance for the instance.
     */
    public /*out*/ readonly upcomingMaintenance!: pulumi.Output<outputs.compute.alpha.UpcomingMaintenanceResponse>;
    /**
     * URL of the zone where the instance resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
     */
    public readonly zone!: pulumi.Output<string>;

    /**
     * Create a Instance resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: InstanceArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            inputs["advancedMachineFeatures"] = args ? args.advancedMachineFeatures : undefined;
            inputs["canIpForward"] = args ? args.canIpForward : undefined;
            inputs["confidentialInstanceConfig"] = args ? args.confidentialInstanceConfig : undefined;
            inputs["deletionProtection"] = args ? args.deletionProtection : undefined;
            inputs["description"] = args ? args.description : undefined;
            inputs["disks"] = args ? args.disks : undefined;
            inputs["displayDevice"] = args ? args.displayDevice : undefined;
            inputs["eraseWindowsVssSignature"] = args ? args.eraseWindowsVssSignature : undefined;
            inputs["guestAccelerators"] = args ? args.guestAccelerators : undefined;
            inputs["hostname"] = args ? args.hostname : undefined;
            inputs["instanceEncryptionKey"] = args ? args.instanceEncryptionKey : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["machineType"] = args ? args.machineType : undefined;
            inputs["metadata"] = args ? args.metadata : undefined;
            inputs["minCpuPlatform"] = args ? args.minCpuPlatform : undefined;
            inputs["name"] = args ? args.name : undefined;
            inputs["networkInterfaces"] = args ? args.networkInterfaces : undefined;
            inputs["networkPerformanceConfig"] = args ? args.networkPerformanceConfig : undefined;
            inputs["postKeyRevocationActionType"] = args ? args.postKeyRevocationActionType : undefined;
            inputs["preservedStateSizeGb"] = args ? args.preservedStateSizeGb : undefined;
            inputs["privateIpv6GoogleAccess"] = args ? args.privateIpv6GoogleAccess : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["requestId"] = args ? args.requestId : undefined;
            inputs["reservationAffinity"] = args ? args.reservationAffinity : undefined;
            inputs["resourcePolicies"] = args ? args.resourcePolicies : undefined;
            inputs["scheduling"] = args ? args.scheduling : undefined;
            inputs["secureTags"] = args ? args.secureTags : undefined;
            inputs["serviceAccounts"] = args ? args.serviceAccounts : undefined;
            inputs["shieldedInstanceConfig"] = args ? args.shieldedInstanceConfig : undefined;
            inputs["shieldedInstanceIntegrityPolicy"] = args ? args.shieldedInstanceIntegrityPolicy : undefined;
            inputs["shieldedVmConfig"] = args ? args.shieldedVmConfig : undefined;
            inputs["shieldedVmIntegrityPolicy"] = args ? args.shieldedVmIntegrityPolicy : undefined;
            inputs["sourceInstanceTemplate"] = args ? args.sourceInstanceTemplate : undefined;
            inputs["sourceMachineImage"] = args ? args.sourceMachineImage : undefined;
            inputs["sourceMachineImageEncryptionKey"] = args ? args.sourceMachineImageEncryptionKey : undefined;
            inputs["tags"] = args ? args.tags : undefined;
            inputs["zone"] = args ? args.zone : undefined;
            inputs["cpuPlatform"] = undefined /*out*/;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["fingerprint"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["labelFingerprint"] = undefined /*out*/;
            inputs["lastStartTimestamp"] = undefined /*out*/;
            inputs["lastStopTimestamp"] = undefined /*out*/;
            inputs["lastSuspendedTimestamp"] = undefined /*out*/;
            inputs["resourceStatus"] = undefined /*out*/;
            inputs["satisfiesPzs"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["selfLinkWithId"] = undefined /*out*/;
            inputs["startRestricted"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["statusMessage"] = undefined /*out*/;
            inputs["upcomingMaintenance"] = undefined /*out*/;
        } else {
            inputs["advancedMachineFeatures"] = undefined /*out*/;
            inputs["canIpForward"] = undefined /*out*/;
            inputs["confidentialInstanceConfig"] = undefined /*out*/;
            inputs["cpuPlatform"] = undefined /*out*/;
            inputs["creationTimestamp"] = undefined /*out*/;
            inputs["deletionProtection"] = undefined /*out*/;
            inputs["description"] = undefined /*out*/;
            inputs["disks"] = undefined /*out*/;
            inputs["displayDevice"] = undefined /*out*/;
            inputs["eraseWindowsVssSignature"] = undefined /*out*/;
            inputs["fingerprint"] = undefined /*out*/;
            inputs["guestAccelerators"] = undefined /*out*/;
            inputs["hostname"] = undefined /*out*/;
            inputs["instanceEncryptionKey"] = undefined /*out*/;
            inputs["kind"] = undefined /*out*/;
            inputs["labelFingerprint"] = undefined /*out*/;
            inputs["labels"] = undefined /*out*/;
            inputs["lastStartTimestamp"] = undefined /*out*/;
            inputs["lastStopTimestamp"] = undefined /*out*/;
            inputs["lastSuspendedTimestamp"] = undefined /*out*/;
            inputs["machineType"] = undefined /*out*/;
            inputs["metadata"] = undefined /*out*/;
            inputs["minCpuPlatform"] = undefined /*out*/;
            inputs["name"] = undefined /*out*/;
            inputs["networkInterfaces"] = undefined /*out*/;
            inputs["networkPerformanceConfig"] = undefined /*out*/;
            inputs["postKeyRevocationActionType"] = undefined /*out*/;
            inputs["preservedStateSizeGb"] = undefined /*out*/;
            inputs["privateIpv6GoogleAccess"] = undefined /*out*/;
            inputs["reservationAffinity"] = undefined /*out*/;
            inputs["resourcePolicies"] = undefined /*out*/;
            inputs["resourceStatus"] = undefined /*out*/;
            inputs["satisfiesPzs"] = undefined /*out*/;
            inputs["scheduling"] = undefined /*out*/;
            inputs["secureTags"] = undefined /*out*/;
            inputs["selfLink"] = undefined /*out*/;
            inputs["selfLinkWithId"] = undefined /*out*/;
            inputs["serviceAccounts"] = undefined /*out*/;
            inputs["shieldedInstanceConfig"] = undefined /*out*/;
            inputs["shieldedInstanceIntegrityPolicy"] = undefined /*out*/;
            inputs["shieldedVmConfig"] = undefined /*out*/;
            inputs["shieldedVmIntegrityPolicy"] = undefined /*out*/;
            inputs["sourceMachineImage"] = undefined /*out*/;
            inputs["sourceMachineImageEncryptionKey"] = undefined /*out*/;
            inputs["startRestricted"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["statusMessage"] = undefined /*out*/;
            inputs["tags"] = undefined /*out*/;
            inputs["upcomingMaintenance"] = undefined /*out*/;
            inputs["zone"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(Instance.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a Instance resource.
 */
export interface InstanceArgs {
    /**
     * Controls for advanced machine-related behavior features.
     */
    advancedMachineFeatures?: pulumi.Input<inputs.compute.alpha.AdvancedMachineFeaturesArgs>;
    /**
     * Allows this instance to send and receive packets with non-matching destination or source IPs. This is required if you plan to use this instance to forward routes. For more information, see Enabling IP Forwarding .
     */
    canIpForward?: pulumi.Input<boolean>;
    confidentialInstanceConfig?: pulumi.Input<inputs.compute.alpha.ConfidentialInstanceConfigArgs>;
    /**
     * Whether the resource should be protected against deletion.
     */
    deletionProtection?: pulumi.Input<boolean>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Array of disks associated with this instance. Persistent disks must be created before you can assign them.
     */
    disks?: pulumi.Input<pulumi.Input<inputs.compute.alpha.AttachedDiskArgs>[]>;
    /**
     * Enables display device for the instance.
     */
    displayDevice?: pulumi.Input<inputs.compute.alpha.DisplayDeviceArgs>;
    /**
     * Specifies whether the disks restored from source snapshots or source machine image should erase Windows specific VSS signature.
     */
    eraseWindowsVssSignature?: pulumi.Input<boolean>;
    /**
     * A list of the type and count of accelerator cards attached to the instance.
     */
    guestAccelerators?: pulumi.Input<pulumi.Input<inputs.compute.alpha.AcceleratorConfigArgs>[]>;
    /**
     * Specifies the hostname of the instance. The specified hostname must be RFC1035 compliant. If hostname is not specified, the default hostname is [INSTANCE_NAME].c.[PROJECT_ID].internal when using the global DNS, and [INSTANCE_NAME].[ZONE].c.[PROJECT_ID].internal when using zonal DNS.
     */
    hostname?: pulumi.Input<string>;
    /**
     * Encrypts or decrypts data for an instance with a customer-supplied encryption key. If you are creating a new instance, this field encrypts the local SSD and in-memory contents of the instance using a key that you provide. If you are restarting an instance protected with a customer-supplied encryption key, you must provide the correct key in order to successfully restart the instance. If you do not provide an encryption key when creating the instance, then the local SSD and in-memory contents will be encrypted using an automatically generated key and you do not need to provide a key to start the instance later. Instance templates do not store customer-supplied encryption keys, so you cannot use your own keys to encrypt local SSDs and in-memory content in a managed instance group.
     */
    instanceEncryptionKey?: pulumi.Input<inputs.compute.alpha.CustomerEncryptionKeyArgs>;
    /**
     * Labels to apply to this instance. These can be later modified by the setLabels method.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Full or partial URL of the machine type resource to use for this instance, in the format: zones/zone/machineTypes/machine-type. This is provided by the client when the instance is created. For example, the following is a valid partial url to a predefined machine type: zones/us-central1-f/machineTypes/n1-standard-1 To create a custom machine type, provide a URL to a machine type in the following format, where CPUS is 1 or an even number up to 32 (2, 4, 6, ... 24, etc), and MEMORY is the total memory for this instance. Memory must be a multiple of 256 MB and must be supplied in MB (e.g. 5 GB of memory is 5120 MB): zones/zone/machineTypes/custom-CPUS-MEMORY For example: zones/us-central1-f/machineTypes/custom-4-5120 For a full list of restrictions, read the Specifications for custom machine types.
     */
    machineType?: pulumi.Input<string>;
    /**
     * The metadata key/value pairs assigned to this instance. This includes custom metadata and predefined keys.
     */
    metadata?: pulumi.Input<inputs.compute.alpha.MetadataArgs>;
    /**
     * Specifies a minimum CPU platform for the VM instance. Applicable values are the friendly names of CPU platforms, such as minCpuPlatform: "Intel Haswell" or minCpuPlatform: "Intel Sandy Bridge".
     */
    minCpuPlatform?: pulumi.Input<string>;
    /**
     * The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    /**
     * An array of network configurations for this instance. These specify how interfaces are configured to interact with other network services, such as connecting to the internet. Multiple interfaces are supported per instance.
     */
    networkInterfaces?: pulumi.Input<pulumi.Input<inputs.compute.alpha.NetworkInterfaceArgs>[]>;
    networkPerformanceConfig?: pulumi.Input<inputs.compute.alpha.NetworkPerformanceConfigArgs>;
    /**
     * PostKeyRevocationActionType of the instance.
     */
    postKeyRevocationActionType?: pulumi.Input<enums.compute.alpha.InstancePostKeyRevocationActionType>;
    /**
     * Total amount of preserved state for SUSPENDED instances. Read-only in the api.
     */
    preservedStateSizeGb?: pulumi.Input<string>;
    /**
     * The private IPv6 google access type for the VM. If not specified, use INHERIT_FROM_SUBNETWORK as default.
     */
    privateIpv6GoogleAccess?: pulumi.Input<enums.compute.alpha.InstancePrivateIpv6GoogleAccess>;
    project?: pulumi.Input<string>;
    requestId?: pulumi.Input<string>;
    /**
     * Specifies the reservations that this instance can consume from.
     */
    reservationAffinity?: pulumi.Input<inputs.compute.alpha.ReservationAffinityArgs>;
    /**
     * Resource policies applied to this instance.
     */
    resourcePolicies?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * Sets the scheduling options for this instance.
     */
    scheduling?: pulumi.Input<inputs.compute.alpha.SchedulingArgs>;
    /**
     * [Input Only] Secure tags to apply to this instance. These can be later modified by the update method. Maximum number of secure tags allowed is 50.
     */
    secureTags?: pulumi.Input<pulumi.Input<string>[]>;
    /**
     * A list of service accounts, with their specified scopes, authorized for this instance. Only one service account per VM instance is supported. Service accounts generate access tokens that can be accessed through the metadata server and used to authenticate applications on the instance. See Service Accounts for more information.
     */
    serviceAccounts?: pulumi.Input<pulumi.Input<inputs.compute.alpha.ServiceAccountArgs>[]>;
    shieldedInstanceConfig?: pulumi.Input<inputs.compute.alpha.ShieldedInstanceConfigArgs>;
    shieldedInstanceIntegrityPolicy?: pulumi.Input<inputs.compute.alpha.ShieldedInstanceIntegrityPolicyArgs>;
    /**
     * Deprecating, please use shielded_instance_config.
     */
    shieldedVmConfig?: pulumi.Input<inputs.compute.alpha.ShieldedVmConfigArgs>;
    /**
     * Deprecating, please use shielded_instance_integrity_policy.
     */
    shieldedVmIntegrityPolicy?: pulumi.Input<inputs.compute.alpha.ShieldedVmIntegrityPolicyArgs>;
    sourceInstanceTemplate?: pulumi.Input<string>;
    /**
     * Source machine image
     */
    sourceMachineImage?: pulumi.Input<string>;
    /**
     * Source machine image encryption key when creating an instance from a machine image.
     */
    sourceMachineImageEncryptionKey?: pulumi.Input<inputs.compute.alpha.CustomerEncryptionKeyArgs>;
    /**
     * Tags to apply to this instance. Tags are used to identify valid sources or targets for network firewalls and are specified by the client during instance creation. The tags can be later modified by the setTags method. Each tag within the list must comply with RFC1035. Multiple tags can be specified via the 'tags.items' field.
     */
    tags?: pulumi.Input<inputs.compute.alpha.TagsArgs>;
    zone?: pulumi.Input<string>;
}
