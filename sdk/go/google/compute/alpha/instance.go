// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates an instance resource in the specified project using the data included in the request.
type Instance struct {
	pulumi.CustomResourceState

	// Controls for advanced machine-related behavior features.
	AdvancedMachineFeatures AdvancedMachineFeaturesResponseOutput `pulumi:"advancedMachineFeatures"`
	// Allows this instance to send and receive packets with non-matching destination or source IPs. This is required if you plan to use this instance to forward routes. For more information, see Enabling IP Forwarding .
	CanIpForward               pulumi.BoolOutput                        `pulumi:"canIpForward"`
	ConfidentialInstanceConfig ConfidentialInstanceConfigResponseOutput `pulumi:"confidentialInstanceConfig"`
	// The CPU platform used by this instance.
	CpuPlatform pulumi.StringOutput `pulumi:"cpuPlatform"`
	// Creation timestamp in RFC3339 text format.
	CreationTimestamp pulumi.StringOutput `pulumi:"creationTimestamp"`
	// Whether the resource should be protected against deletion.
	DeletionProtection pulumi.BoolOutput `pulumi:"deletionProtection"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringOutput `pulumi:"description"`
	// Array of disks associated with this instance. Persistent disks must be created before you can assign them.
	Disks AttachedDiskResponseArrayOutput `pulumi:"disks"`
	// Enables display device for the instance.
	DisplayDevice DisplayDeviceResponseOutput `pulumi:"displayDevice"`
	// Specifies whether the disks restored from source snapshots or source machine image should erase Windows specific VSS signature.
	EraseWindowsVssSignature pulumi.BoolOutput `pulumi:"eraseWindowsVssSignature"`
	// Specifies a fingerprint for this resource, which is essentially a hash of the instance's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update the instance. You must always provide an up-to-date fingerprint hash in order to update the instance. To see the latest fingerprint, make get() request to the instance.
	Fingerprint pulumi.StringOutput `pulumi:"fingerprint"`
	// A list of the type and count of accelerator cards attached to the instance.
	GuestAccelerators AcceleratorConfigResponseArrayOutput `pulumi:"guestAccelerators"`
	// Specifies the hostname of the instance. The specified hostname must be RFC1035 compliant. If hostname is not specified, the default hostname is [INSTANCE_NAME].c.[PROJECT_ID].internal when using the global DNS, and [INSTANCE_NAME].[ZONE].c.[PROJECT_ID].internal when using zonal DNS.
	Hostname pulumi.StringOutput `pulumi:"hostname"`
	// Encrypts or decrypts data for an instance with a customer-supplied encryption key. If you are creating a new instance, this field encrypts the local SSD and in-memory contents of the instance using a key that you provide. If you are restarting an instance protected with a customer-supplied encryption key, you must provide the correct key in order to successfully restart the instance. If you do not provide an encryption key when creating the instance, then the local SSD and in-memory contents will be encrypted using an automatically generated key and you do not need to provide a key to start the instance later. Instance templates do not store customer-supplied encryption keys, so you cannot use your own keys to encrypt local SSDs and in-memory content in a managed instance group.
	InstanceEncryptionKey CustomerEncryptionKeyResponseOutput `pulumi:"instanceEncryptionKey"`
	// KeyRevocationActionType of the instance. Supported options are "STOP" and "NONE". The default value is "NONE" if it is not specified.
	KeyRevocationActionType pulumi.StringOutput `pulumi:"keyRevocationActionType"`
	// Type of the resource. Always compute#instance for instances.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// A fingerprint for this request, which is essentially a hash of the label's contents and used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels. To see the latest fingerprint, make get() request to the instance.
	LabelFingerprint pulumi.StringOutput `pulumi:"labelFingerprint"`
	// Labels to apply to this instance. These can be later modified by the setLabels method.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Last start timestamp in RFC3339 text format.
	LastStartTimestamp pulumi.StringOutput `pulumi:"lastStartTimestamp"`
	// Last stop timestamp in RFC3339 text format.
	LastStopTimestamp pulumi.StringOutput `pulumi:"lastStopTimestamp"`
	// Last suspended timestamp in RFC3339 text format.
	LastSuspendedTimestamp pulumi.StringOutput `pulumi:"lastSuspendedTimestamp"`
	// Full or partial URL of the machine type resource to use for this instance, in the format: zones/zone/machineTypes/machine-type. This is provided by the client when the instance is created. For example, the following is a valid partial url to a predefined machine type: zones/us-central1-f/machineTypes/n1-standard-1 To create a custom machine type, provide a URL to a machine type in the following format, where CPUS is 1 or an even number up to 32 (2, 4, 6, ... 24, etc), and MEMORY is the total memory for this instance. Memory must be a multiple of 256 MB and must be supplied in MB (e.g. 5 GB of memory is 5120 MB): zones/zone/machineTypes/custom-CPUS-MEMORY For example: zones/us-central1-f/machineTypes/custom-4-5120 For a full list of restrictions, read the Specifications for custom machine types.
	MachineType pulumi.StringOutput `pulumi:"machineType"`
	// The metadata key/value pairs assigned to this instance. This includes custom metadata and predefined keys.
	Metadata MetadataResponseOutput `pulumi:"metadata"`
	// Specifies a minimum CPU platform for the VM instance. Applicable values are the friendly names of CPU platforms, such as minCpuPlatform: "Intel Haswell" or minCpuPlatform: "Intel Sandy Bridge".
	MinCpuPlatform pulumi.StringOutput `pulumi:"minCpuPlatform"`
	// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringOutput `pulumi:"name"`
	// An array of network configurations for this instance. These specify how interfaces are configured to interact with other network services, such as connecting to the internet. Multiple interfaces are supported per instance.
	NetworkInterfaces        NetworkInterfaceResponseArrayOutput    `pulumi:"networkInterfaces"`
	NetworkPerformanceConfig NetworkPerformanceConfigResponseOutput `pulumi:"networkPerformanceConfig"`
	// Input only. [Input Only] Additional params passed with the request, but not persisted as part of resource payload.
	Params InstanceParamsResponseOutput `pulumi:"params"`
	// PostKeyRevocationActionType of the instance.
	PostKeyRevocationActionType pulumi.StringOutput `pulumi:"postKeyRevocationActionType"`
	// Total amount of preserved state for SUSPENDED instances. Read-only in the api.
	PreservedStateSizeGb pulumi.StringOutput `pulumi:"preservedStateSizeGb"`
	// The private IPv6 google access type for the VM. If not specified, use INHERIT_FROM_SUBNETWORK as default.
	PrivateIpv6GoogleAccess pulumi.StringOutput `pulumi:"privateIpv6GoogleAccess"`
	// Specifies the reservations that this instance can consume from.
	ReservationAffinity ReservationAffinityResponseOutput `pulumi:"reservationAffinity"`
	// Resource policies applied to this instance.
	ResourcePolicies pulumi.StringArrayOutput `pulumi:"resourcePolicies"`
	// Specifies values set for instance attributes as compared to the values requested by user in the corresponding input only field.
	ResourceStatus ResourceStatusResponseOutput `pulumi:"resourceStatus"`
	// Reserved for future use.
	SatisfiesPzs pulumi.BoolOutput `pulumi:"satisfiesPzs"`
	// Sets the scheduling options for this instance.
	Scheduling SchedulingResponseOutput `pulumi:"scheduling"`
	// [Input Only] Secure tags to apply to this instance. These can be later modified by the update method. Maximum number of secure tags allowed is 50.
	SecureTags pulumi.StringArrayOutput `pulumi:"secureTags"`
	// Server-defined URL for this resource.
	SelfLink pulumi.StringOutput `pulumi:"selfLink"`
	// Server-defined URL for this resource with the resource id.
	SelfLinkWithId pulumi.StringOutput `pulumi:"selfLinkWithId"`
	// A list of service accounts, with their specified scopes, authorized for this instance. Only one service account per VM instance is supported. Service accounts generate access tokens that can be accessed through the metadata server and used to authenticate applications on the instance. See Service Accounts for more information.
	ServiceAccounts                 ServiceAccountResponseArrayOutput             `pulumi:"serviceAccounts"`
	ShieldedInstanceConfig          ShieldedInstanceConfigResponseOutput          `pulumi:"shieldedInstanceConfig"`
	ShieldedInstanceIntegrityPolicy ShieldedInstanceIntegrityPolicyResponseOutput `pulumi:"shieldedInstanceIntegrityPolicy"`
	// Deprecating, please use shielded_instance_config.
	ShieldedVmConfig ShieldedVmConfigResponseOutput `pulumi:"shieldedVmConfig"`
	// Deprecating, please use shielded_instance_integrity_policy.
	ShieldedVmIntegrityPolicy ShieldedVmIntegrityPolicyResponseOutput `pulumi:"shieldedVmIntegrityPolicy"`
	// Source machine image
	SourceMachineImage pulumi.StringOutput `pulumi:"sourceMachineImage"`
	// Source machine image encryption key when creating an instance from a machine image.
	SourceMachineImageEncryptionKey CustomerEncryptionKeyResponseOutput `pulumi:"sourceMachineImageEncryptionKey"`
	// Whether a VM has been restricted for start because Compute Engine has detected suspicious activity.
	StartRestricted pulumi.BoolOutput `pulumi:"startRestricted"`
	// The status of the instance. One of the following values: PROVISIONING, STAGING, RUNNING, STOPPING, SUSPENDING, SUSPENDED, REPAIRING, and TERMINATED. For more information about the status of the instance, see Instance life cycle.
	Status pulumi.StringOutput `pulumi:"status"`
	// An optional, human-readable explanation of the status.
	StatusMessage pulumi.StringOutput `pulumi:"statusMessage"`
	// Tags to apply to this instance. Tags are used to identify valid sources or targets for network firewalls and are specified by the client during instance creation. The tags can be later modified by the setTags method. Each tag within the list must comply with RFC1035. Multiple tags can be specified via the 'tags.items' field.
	Tags TagsResponseOutput `pulumi:"tags"`
	// Specifies upcoming maintenance for the instance.
	UpcomingMaintenance UpcomingMaintenanceResponseOutput `pulumi:"upcomingMaintenance"`
	// URL of the zone where the instance resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		args = &InstanceArgs{}
	}

	var resource Instance
	err := ctx.RegisterResource("google-native:compute/alpha:Instance", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetInstance gets an existing Instance resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetInstance(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *InstanceState, opts ...pulumi.ResourceOption) (*Instance, error) {
	var resource Instance
	err := ctx.ReadResource("google-native:compute/alpha:Instance", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Instance resources.
type instanceState struct {
}

type InstanceState struct {
}

func (InstanceState) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceState)(nil)).Elem()
}

type instanceArgs struct {
	// Controls for advanced machine-related behavior features.
	AdvancedMachineFeatures *AdvancedMachineFeatures `pulumi:"advancedMachineFeatures"`
	// Allows this instance to send and receive packets with non-matching destination or source IPs. This is required if you plan to use this instance to forward routes. For more information, see Enabling IP Forwarding .
	CanIpForward               *bool                       `pulumi:"canIpForward"`
	ConfidentialInstanceConfig *ConfidentialInstanceConfig `pulumi:"confidentialInstanceConfig"`
	// Whether the resource should be protected against deletion.
	DeletionProtection *bool `pulumi:"deletionProtection"`
	// An optional description of this resource. Provide this property when you create the resource.
	Description *string `pulumi:"description"`
	// Array of disks associated with this instance. Persistent disks must be created before you can assign them.
	Disks []AttachedDisk `pulumi:"disks"`
	// Enables display device for the instance.
	DisplayDevice *DisplayDevice `pulumi:"displayDevice"`
	// Specifies whether the disks restored from source snapshots or source machine image should erase Windows specific VSS signature.
	EraseWindowsVssSignature *bool `pulumi:"eraseWindowsVssSignature"`
	// A list of the type and count of accelerator cards attached to the instance.
	GuestAccelerators []AcceleratorConfig `pulumi:"guestAccelerators"`
	// Specifies the hostname of the instance. The specified hostname must be RFC1035 compliant. If hostname is not specified, the default hostname is [INSTANCE_NAME].c.[PROJECT_ID].internal when using the global DNS, and [INSTANCE_NAME].[ZONE].c.[PROJECT_ID].internal when using zonal DNS.
	Hostname *string `pulumi:"hostname"`
	// Encrypts or decrypts data for an instance with a customer-supplied encryption key. If you are creating a new instance, this field encrypts the local SSD and in-memory contents of the instance using a key that you provide. If you are restarting an instance protected with a customer-supplied encryption key, you must provide the correct key in order to successfully restart the instance. If you do not provide an encryption key when creating the instance, then the local SSD and in-memory contents will be encrypted using an automatically generated key and you do not need to provide a key to start the instance later. Instance templates do not store customer-supplied encryption keys, so you cannot use your own keys to encrypt local SSDs and in-memory content in a managed instance group.
	InstanceEncryptionKey *CustomerEncryptionKey `pulumi:"instanceEncryptionKey"`
	// KeyRevocationActionType of the instance. Supported options are "STOP" and "NONE". The default value is "NONE" if it is not specified.
	KeyRevocationActionType *InstanceKeyRevocationActionType `pulumi:"keyRevocationActionType"`
	// Labels to apply to this instance. These can be later modified by the setLabels method.
	Labels map[string]string `pulumi:"labels"`
	// Full or partial URL of the machine type resource to use for this instance, in the format: zones/zone/machineTypes/machine-type. This is provided by the client when the instance is created. For example, the following is a valid partial url to a predefined machine type: zones/us-central1-f/machineTypes/n1-standard-1 To create a custom machine type, provide a URL to a machine type in the following format, where CPUS is 1 or an even number up to 32 (2, 4, 6, ... 24, etc), and MEMORY is the total memory for this instance. Memory must be a multiple of 256 MB and must be supplied in MB (e.g. 5 GB of memory is 5120 MB): zones/zone/machineTypes/custom-CPUS-MEMORY For example: zones/us-central1-f/machineTypes/custom-4-5120 For a full list of restrictions, read the Specifications for custom machine types.
	MachineType *string `pulumi:"machineType"`
	// The metadata key/value pairs assigned to this instance. This includes custom metadata and predefined keys.
	Metadata *Metadata `pulumi:"metadata"`
	// Specifies a minimum CPU platform for the VM instance. Applicable values are the friendly names of CPU platforms, such as minCpuPlatform: "Intel Haswell" or minCpuPlatform: "Intel Sandy Bridge".
	MinCpuPlatform *string `pulumi:"minCpuPlatform"`
	// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name *string `pulumi:"name"`
	// An array of network configurations for this instance. These specify how interfaces are configured to interact with other network services, such as connecting to the internet. Multiple interfaces are supported per instance.
	NetworkInterfaces        []NetworkInterface        `pulumi:"networkInterfaces"`
	NetworkPerformanceConfig *NetworkPerformanceConfig `pulumi:"networkPerformanceConfig"`
	// Input only. [Input Only] Additional params passed with the request, but not persisted as part of resource payload.
	Params *InstanceParams `pulumi:"params"`
	// PostKeyRevocationActionType of the instance.
	PostKeyRevocationActionType *InstancePostKeyRevocationActionType `pulumi:"postKeyRevocationActionType"`
	// Total amount of preserved state for SUSPENDED instances. Read-only in the api.
	PreservedStateSizeGb *string `pulumi:"preservedStateSizeGb"`
	// The private IPv6 google access type for the VM. If not specified, use INHERIT_FROM_SUBNETWORK as default.
	PrivateIpv6GoogleAccess *InstancePrivateIpv6GoogleAccess `pulumi:"privateIpv6GoogleAccess"`
	Project                 *string                          `pulumi:"project"`
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId *string `pulumi:"requestId"`
	// Specifies the reservations that this instance can consume from.
	ReservationAffinity *ReservationAffinity `pulumi:"reservationAffinity"`
	// Resource policies applied to this instance.
	ResourcePolicies []string `pulumi:"resourcePolicies"`
	// Sets the scheduling options for this instance.
	Scheduling *Scheduling `pulumi:"scheduling"`
	// [Input Only] Secure tags to apply to this instance. These can be later modified by the update method. Maximum number of secure tags allowed is 50.
	SecureTags []string `pulumi:"secureTags"`
	// A list of service accounts, with their specified scopes, authorized for this instance. Only one service account per VM instance is supported. Service accounts generate access tokens that can be accessed through the metadata server and used to authenticate applications on the instance. See Service Accounts for more information.
	ServiceAccounts                 []ServiceAccount                 `pulumi:"serviceAccounts"`
	ShieldedInstanceConfig          *ShieldedInstanceConfig          `pulumi:"shieldedInstanceConfig"`
	ShieldedInstanceIntegrityPolicy *ShieldedInstanceIntegrityPolicy `pulumi:"shieldedInstanceIntegrityPolicy"`
	// Deprecating, please use shielded_instance_config.
	ShieldedVmConfig *ShieldedVmConfig `pulumi:"shieldedVmConfig"`
	// Deprecating, please use shielded_instance_integrity_policy.
	ShieldedVmIntegrityPolicy *ShieldedVmIntegrityPolicy `pulumi:"shieldedVmIntegrityPolicy"`
	// Specifies instance template to create the instance. This field is optional. It can be a full or partial URL. For example, the following are all valid URLs to an instance template: - https://www.googleapis.com/compute/v1/projects/project /global/instanceTemplates/instanceTemplate - projects/project/global/instanceTemplates/instanceTemplate - global/instanceTemplates/instanceTemplate
	SourceInstanceTemplate *string `pulumi:"sourceInstanceTemplate"`
	// Source machine image
	SourceMachineImage *string `pulumi:"sourceMachineImage"`
	// Source machine image encryption key when creating an instance from a machine image.
	SourceMachineImageEncryptionKey *CustomerEncryptionKey `pulumi:"sourceMachineImageEncryptionKey"`
	// Tags to apply to this instance. Tags are used to identify valid sources or targets for network firewalls and are specified by the client during instance creation. The tags can be later modified by the setTags method. Each tag within the list must comply with RFC1035. Multiple tags can be specified via the 'tags.items' field.
	Tags *Tags   `pulumi:"tags"`
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// Controls for advanced machine-related behavior features.
	AdvancedMachineFeatures AdvancedMachineFeaturesPtrInput
	// Allows this instance to send and receive packets with non-matching destination or source IPs. This is required if you plan to use this instance to forward routes. For more information, see Enabling IP Forwarding .
	CanIpForward               pulumi.BoolPtrInput
	ConfidentialInstanceConfig ConfidentialInstanceConfigPtrInput
	// Whether the resource should be protected against deletion.
	DeletionProtection pulumi.BoolPtrInput
	// An optional description of this resource. Provide this property when you create the resource.
	Description pulumi.StringPtrInput
	// Array of disks associated with this instance. Persistent disks must be created before you can assign them.
	Disks AttachedDiskArrayInput
	// Enables display device for the instance.
	DisplayDevice DisplayDevicePtrInput
	// Specifies whether the disks restored from source snapshots or source machine image should erase Windows specific VSS signature.
	EraseWindowsVssSignature pulumi.BoolPtrInput
	// A list of the type and count of accelerator cards attached to the instance.
	GuestAccelerators AcceleratorConfigArrayInput
	// Specifies the hostname of the instance. The specified hostname must be RFC1035 compliant. If hostname is not specified, the default hostname is [INSTANCE_NAME].c.[PROJECT_ID].internal when using the global DNS, and [INSTANCE_NAME].[ZONE].c.[PROJECT_ID].internal when using zonal DNS.
	Hostname pulumi.StringPtrInput
	// Encrypts or decrypts data for an instance with a customer-supplied encryption key. If you are creating a new instance, this field encrypts the local SSD and in-memory contents of the instance using a key that you provide. If you are restarting an instance protected with a customer-supplied encryption key, you must provide the correct key in order to successfully restart the instance. If you do not provide an encryption key when creating the instance, then the local SSD and in-memory contents will be encrypted using an automatically generated key and you do not need to provide a key to start the instance later. Instance templates do not store customer-supplied encryption keys, so you cannot use your own keys to encrypt local SSDs and in-memory content in a managed instance group.
	InstanceEncryptionKey CustomerEncryptionKeyPtrInput
	// KeyRevocationActionType of the instance. Supported options are "STOP" and "NONE". The default value is "NONE" if it is not specified.
	KeyRevocationActionType InstanceKeyRevocationActionTypePtrInput
	// Labels to apply to this instance. These can be later modified by the setLabels method.
	Labels pulumi.StringMapInput
	// Full or partial URL of the machine type resource to use for this instance, in the format: zones/zone/machineTypes/machine-type. This is provided by the client when the instance is created. For example, the following is a valid partial url to a predefined machine type: zones/us-central1-f/machineTypes/n1-standard-1 To create a custom machine type, provide a URL to a machine type in the following format, where CPUS is 1 or an even number up to 32 (2, 4, 6, ... 24, etc), and MEMORY is the total memory for this instance. Memory must be a multiple of 256 MB and must be supplied in MB (e.g. 5 GB of memory is 5120 MB): zones/zone/machineTypes/custom-CPUS-MEMORY For example: zones/us-central1-f/machineTypes/custom-4-5120 For a full list of restrictions, read the Specifications for custom machine types.
	MachineType pulumi.StringPtrInput
	// The metadata key/value pairs assigned to this instance. This includes custom metadata and predefined keys.
	Metadata MetadataPtrInput
	// Specifies a minimum CPU platform for the VM instance. Applicable values are the friendly names of CPU platforms, such as minCpuPlatform: "Intel Haswell" or minCpuPlatform: "Intel Sandy Bridge".
	MinCpuPlatform pulumi.StringPtrInput
	// The name of the resource, provided by the client when initially creating the resource. The resource name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
	Name pulumi.StringPtrInput
	// An array of network configurations for this instance. These specify how interfaces are configured to interact with other network services, such as connecting to the internet. Multiple interfaces are supported per instance.
	NetworkInterfaces        NetworkInterfaceArrayInput
	NetworkPerformanceConfig NetworkPerformanceConfigPtrInput
	// Input only. [Input Only] Additional params passed with the request, but not persisted as part of resource payload.
	Params InstanceParamsPtrInput
	// PostKeyRevocationActionType of the instance.
	PostKeyRevocationActionType InstancePostKeyRevocationActionTypePtrInput
	// Total amount of preserved state for SUSPENDED instances. Read-only in the api.
	PreservedStateSizeGb pulumi.StringPtrInput
	// The private IPv6 google access type for the VM. If not specified, use INHERIT_FROM_SUBNETWORK as default.
	PrivateIpv6GoogleAccess InstancePrivateIpv6GoogleAccessPtrInput
	Project                 pulumi.StringPtrInput
	// An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
	RequestId pulumi.StringPtrInput
	// Specifies the reservations that this instance can consume from.
	ReservationAffinity ReservationAffinityPtrInput
	// Resource policies applied to this instance.
	ResourcePolicies pulumi.StringArrayInput
	// Sets the scheduling options for this instance.
	Scheduling SchedulingPtrInput
	// [Input Only] Secure tags to apply to this instance. These can be later modified by the update method. Maximum number of secure tags allowed is 50.
	SecureTags pulumi.StringArrayInput
	// A list of service accounts, with their specified scopes, authorized for this instance. Only one service account per VM instance is supported. Service accounts generate access tokens that can be accessed through the metadata server and used to authenticate applications on the instance. See Service Accounts for more information.
	ServiceAccounts                 ServiceAccountArrayInput
	ShieldedInstanceConfig          ShieldedInstanceConfigPtrInput
	ShieldedInstanceIntegrityPolicy ShieldedInstanceIntegrityPolicyPtrInput
	// Deprecating, please use shielded_instance_config.
	ShieldedVmConfig ShieldedVmConfigPtrInput
	// Deprecating, please use shielded_instance_integrity_policy.
	ShieldedVmIntegrityPolicy ShieldedVmIntegrityPolicyPtrInput
	// Specifies instance template to create the instance. This field is optional. It can be a full or partial URL. For example, the following are all valid URLs to an instance template: - https://www.googleapis.com/compute/v1/projects/project /global/instanceTemplates/instanceTemplate - projects/project/global/instanceTemplates/instanceTemplate - global/instanceTemplates/instanceTemplate
	SourceInstanceTemplate pulumi.StringPtrInput
	// Source machine image
	SourceMachineImage pulumi.StringPtrInput
	// Source machine image encryption key when creating an instance from a machine image.
	SourceMachineImageEncryptionKey CustomerEncryptionKeyPtrInput
	// Tags to apply to this instance. Tags are used to identify valid sources or targets for network firewalls and are specified by the client during instance creation. The tags can be later modified by the setTags method. Each tag within the list must comply with RFC1035. Multiple tags can be specified via the 'tags.items' field.
	Tags TagsPtrInput
	Zone pulumi.StringPtrInput
}

func (InstanceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*instanceArgs)(nil)).Elem()
}

type InstanceInput interface {
	pulumi.Input

	ToInstanceOutput() InstanceOutput
	ToInstanceOutputWithContext(ctx context.Context) InstanceOutput
}

func (*Instance) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (i *Instance) ToInstanceOutput() InstanceOutput {
	return i.ToInstanceOutputWithContext(context.Background())
}

func (i *Instance) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstanceOutput)
}

type InstanceOutput struct{ *pulumi.OutputState }

func (InstanceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Instance)(nil)).Elem()
}

func (o InstanceOutput) ToInstanceOutput() InstanceOutput {
	return o
}

func (o InstanceOutput) ToInstanceOutputWithContext(ctx context.Context) InstanceOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceInput)(nil)).Elem(), &Instance{})
	pulumi.RegisterOutputType(InstanceOutput{})
}
