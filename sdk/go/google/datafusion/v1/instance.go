// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new Data Fusion instance in the specified project and location.
// Auto-naming is currently not supported for this resource.
type Instance struct {
	pulumi.CustomResourceState

	// List of accelerators enabled for this CDF instance.
	Accelerators AcceleratorResponseArrayOutput `pulumi:"accelerators"`
	// Endpoint on which the REST APIs is accessible.
	ApiEndpoint pulumi.StringOutput `pulumi:"apiEndpoint"`
	// Available versions that the instance can be upgraded to using UpdateInstanceRequest.
	AvailableVersion VersionResponseArrayOutput `pulumi:"availableVersion"`
	// The time the instance was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The crypto key configuration. This field is used by the Customer-Managed Encryption Keys (CMEK) feature.
	CryptoKeyConfig CryptoKeyConfigResponseOutput `pulumi:"cryptoKeyConfig"`
	// Optional. Reserved for future use.
	DataplexDataLineageIntegrationEnabled pulumi.BoolOutput `pulumi:"dataplexDataLineageIntegrationEnabled"`
	// User-managed service account to set on Dataproc when Cloud Data Fusion creates Dataproc to run data processing pipelines. This allows users to have fine-grained access control on Dataproc's accesses to cloud resources.
	DataprocServiceAccount pulumi.StringOutput `pulumi:"dataprocServiceAccount"`
	// A description of this instance.
	Description pulumi.StringOutput `pulumi:"description"`
	// If the instance state is DISABLED, the reason for disabling the instance.
	DisabledReason pulumi.StringArrayOutput `pulumi:"disabledReason"`
	// Display name for an instance.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Option to enable granular role-based access control.
	EnableRbac pulumi.BoolOutput `pulumi:"enableRbac"`
	// Option to enable Stackdriver Logging.
	EnableStackdriverLogging pulumi.BoolOutput `pulumi:"enableStackdriverLogging"`
	// Option to enable Stackdriver Monitoring.
	EnableStackdriverMonitoring pulumi.BoolOutput `pulumi:"enableStackdriverMonitoring"`
	// Option to enable granular zone separation.
	EnableZoneSeparation pulumi.BoolOutput `pulumi:"enableZoneSeparation"`
	// Option to enable and pass metadata for event publishing.
	EventPublishConfig EventPublishConfigResponseOutput `pulumi:"eventPublishConfig"`
	// Cloud Storage bucket generated by Data Fusion in the customer project.
	GcsBucket pulumi.StringOutput `pulumi:"gcsBucket"`
	// Required. The name of the instance to create.
	InstanceId pulumi.StringOutput `pulumi:"instanceId"`
	// The resource labels for instance to use to annotate any related underlying resources such as Compute Engine VMs. The character '=' is not allowed to be used within the labels.
	Labels   pulumi.StringMapOutput `pulumi:"labels"`
	Location pulumi.StringOutput    `pulumi:"location"`
	// The name of this instance is in the form of projects/{project}/locations/{location}/instances/{instance}.
	Name pulumi.StringOutput `pulumi:"name"`
	// Network configuration options. These are required when a private Data Fusion instance is to be created.
	NetworkConfig NetworkConfigResponseOutput `pulumi:"networkConfig"`
	// Map of additional options used to configure the behavior of Data Fusion instance.
	Options pulumi.StringMapOutput `pulumi:"options"`
	// P4 service account for the customer project.
	P4ServiceAccount pulumi.StringOutput `pulumi:"p4ServiceAccount"`
	// Optional. Current patch revision of the Data Fusion.
	PatchRevision pulumi.StringOutput `pulumi:"patchRevision"`
	// Specifies whether the Data Fusion instance should be private. If set to true, all Data Fusion nodes will have private IP addresses and will not be able to access the public internet.
	PrivateInstance pulumi.BoolOutput   `pulumi:"privateInstance"`
	Project         pulumi.StringOutput `pulumi:"project"`
	// Reserved for future use.
	SatisfiesPzs pulumi.BoolOutput `pulumi:"satisfiesPzs"`
	// Deprecated. Use tenant_project_id instead to extract the tenant project ID.
	//
	// Deprecated: Output only. Deprecated. Use tenant_project_id instead to extract the tenant project ID.
	ServiceAccount pulumi.StringOutput `pulumi:"serviceAccount"`
	// Endpoint on which the Data Fusion UI is accessible.
	ServiceEndpoint pulumi.StringOutput `pulumi:"serviceEndpoint"`
	// The current state of this Data Fusion instance.
	State pulumi.StringOutput `pulumi:"state"`
	// Additional information about the current state of this Data Fusion instance if available.
	StateMessage pulumi.StringOutput `pulumi:"stateMessage"`
	// The name of the tenant project.
	TenantProjectId pulumi.StringOutput `pulumi:"tenantProjectId"`
	// Instance type.
	Type pulumi.StringOutput `pulumi:"type"`
	// The time the instance was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Current version of the Data Fusion. Only specifiable in Update.
	Version pulumi.StringOutput `pulumi:"version"`
	// Endpoint on which the Data Fusion UI is accessible to third-party users
	WorkforceIdentityServiceEndpoint pulumi.StringOutput `pulumi:"workforceIdentityServiceEndpoint"`
	// Name of the zone in which the Data Fusion instance will be created. Only DEVELOPER instances use this field.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewInstance registers a new resource with the given unique name, arguments, and options.
func NewInstance(ctx *pulumi.Context,
	name string, args *InstanceArgs, opts ...pulumi.ResourceOption) (*Instance, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.InstanceId == nil {
		return nil, errors.New("invalid value for required argument 'InstanceId'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"instanceId",
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Instance
	err := ctx.RegisterResource("google-native:datafusion/v1:Instance", name, args, &resource, opts...)
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
	err := ctx.ReadResource("google-native:datafusion/v1:Instance", name, id, state, &resource, opts...)
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
	// The crypto key configuration. This field is used by the Customer-Managed Encryption Keys (CMEK) feature.
	CryptoKeyConfig *CryptoKeyConfig `pulumi:"cryptoKeyConfig"`
	// Optional. Reserved for future use.
	DataplexDataLineageIntegrationEnabled *bool `pulumi:"dataplexDataLineageIntegrationEnabled"`
	// User-managed service account to set on Dataproc when Cloud Data Fusion creates Dataproc to run data processing pipelines. This allows users to have fine-grained access control on Dataproc's accesses to cloud resources.
	DataprocServiceAccount *string `pulumi:"dataprocServiceAccount"`
	// A description of this instance.
	Description *string `pulumi:"description"`
	// Display name for an instance.
	DisplayName *string `pulumi:"displayName"`
	// Option to enable granular role-based access control.
	EnableRbac *bool `pulumi:"enableRbac"`
	// Option to enable Stackdriver Logging.
	EnableStackdriverLogging *bool `pulumi:"enableStackdriverLogging"`
	// Option to enable Stackdriver Monitoring.
	EnableStackdriverMonitoring *bool `pulumi:"enableStackdriverMonitoring"`
	// Option to enable granular zone separation.
	EnableZoneSeparation *bool `pulumi:"enableZoneSeparation"`
	// Option to enable and pass metadata for event publishing.
	EventPublishConfig *EventPublishConfig `pulumi:"eventPublishConfig"`
	// Required. The name of the instance to create.
	InstanceId string `pulumi:"instanceId"`
	// The resource labels for instance to use to annotate any related underlying resources such as Compute Engine VMs. The character '=' is not allowed to be used within the labels.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	// Network configuration options. These are required when a private Data Fusion instance is to be created.
	NetworkConfig *NetworkConfig `pulumi:"networkConfig"`
	// Map of additional options used to configure the behavior of Data Fusion instance.
	Options map[string]string `pulumi:"options"`
	// Optional. Current patch revision of the Data Fusion.
	PatchRevision *string `pulumi:"patchRevision"`
	// Specifies whether the Data Fusion instance should be private. If set to true, all Data Fusion nodes will have private IP addresses and will not be able to access the public internet.
	PrivateInstance *bool   `pulumi:"privateInstance"`
	Project         *string `pulumi:"project"`
	// Instance type.
	Type InstanceType `pulumi:"type"`
	// Current version of the Data Fusion. Only specifiable in Update.
	Version *string `pulumi:"version"`
	// Name of the zone in which the Data Fusion instance will be created. Only DEVELOPER instances use this field.
	Zone *string `pulumi:"zone"`
}

// The set of arguments for constructing a Instance resource.
type InstanceArgs struct {
	// The crypto key configuration. This field is used by the Customer-Managed Encryption Keys (CMEK) feature.
	CryptoKeyConfig CryptoKeyConfigPtrInput
	// Optional. Reserved for future use.
	DataplexDataLineageIntegrationEnabled pulumi.BoolPtrInput
	// User-managed service account to set on Dataproc when Cloud Data Fusion creates Dataproc to run data processing pipelines. This allows users to have fine-grained access control on Dataproc's accesses to cloud resources.
	DataprocServiceAccount pulumi.StringPtrInput
	// A description of this instance.
	Description pulumi.StringPtrInput
	// Display name for an instance.
	DisplayName pulumi.StringPtrInput
	// Option to enable granular role-based access control.
	EnableRbac pulumi.BoolPtrInput
	// Option to enable Stackdriver Logging.
	EnableStackdriverLogging pulumi.BoolPtrInput
	// Option to enable Stackdriver Monitoring.
	EnableStackdriverMonitoring pulumi.BoolPtrInput
	// Option to enable granular zone separation.
	EnableZoneSeparation pulumi.BoolPtrInput
	// Option to enable and pass metadata for event publishing.
	EventPublishConfig EventPublishConfigPtrInput
	// Required. The name of the instance to create.
	InstanceId pulumi.StringInput
	// The resource labels for instance to use to annotate any related underlying resources such as Compute Engine VMs. The character '=' is not allowed to be used within the labels.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	// Network configuration options. These are required when a private Data Fusion instance is to be created.
	NetworkConfig NetworkConfigPtrInput
	// Map of additional options used to configure the behavior of Data Fusion instance.
	Options pulumi.StringMapInput
	// Optional. Current patch revision of the Data Fusion.
	PatchRevision pulumi.StringPtrInput
	// Specifies whether the Data Fusion instance should be private. If set to true, all Data Fusion nodes will have private IP addresses and will not be able to access the public internet.
	PrivateInstance pulumi.BoolPtrInput
	Project         pulumi.StringPtrInput
	// Instance type.
	Type InstanceTypeInput
	// Current version of the Data Fusion. Only specifiable in Update.
	Version pulumi.StringPtrInput
	// Name of the zone in which the Data Fusion instance will be created. Only DEVELOPER instances use this field.
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

// List of accelerators enabled for this CDF instance.
func (o InstanceOutput) Accelerators() AcceleratorResponseArrayOutput {
	return o.ApplyT(func(v *Instance) AcceleratorResponseArrayOutput { return v.Accelerators }).(AcceleratorResponseArrayOutput)
}

// Endpoint on which the REST APIs is accessible.
func (o InstanceOutput) ApiEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ApiEndpoint }).(pulumi.StringOutput)
}

// Available versions that the instance can be upgraded to using UpdateInstanceRequest.
func (o InstanceOutput) AvailableVersion() VersionResponseArrayOutput {
	return o.ApplyT(func(v *Instance) VersionResponseArrayOutput { return v.AvailableVersion }).(VersionResponseArrayOutput)
}

// The time the instance was created.
func (o InstanceOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The crypto key configuration. This field is used by the Customer-Managed Encryption Keys (CMEK) feature.
func (o InstanceOutput) CryptoKeyConfig() CryptoKeyConfigResponseOutput {
	return o.ApplyT(func(v *Instance) CryptoKeyConfigResponseOutput { return v.CryptoKeyConfig }).(CryptoKeyConfigResponseOutput)
}

// Optional. Reserved for future use.
func (o InstanceOutput) DataplexDataLineageIntegrationEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolOutput { return v.DataplexDataLineageIntegrationEnabled }).(pulumi.BoolOutput)
}

// User-managed service account to set on Dataproc when Cloud Data Fusion creates Dataproc to run data processing pipelines. This allows users to have fine-grained access control on Dataproc's accesses to cloud resources.
func (o InstanceOutput) DataprocServiceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.DataprocServiceAccount }).(pulumi.StringOutput)
}

// A description of this instance.
func (o InstanceOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// If the instance state is DISABLED, the reason for disabling the instance.
func (o InstanceOutput) DisabledReason() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringArrayOutput { return v.DisabledReason }).(pulumi.StringArrayOutput)
}

// Display name for an instance.
func (o InstanceOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Option to enable granular role-based access control.
func (o InstanceOutput) EnableRbac() pulumi.BoolOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolOutput { return v.EnableRbac }).(pulumi.BoolOutput)
}

// Option to enable Stackdriver Logging.
func (o InstanceOutput) EnableStackdriverLogging() pulumi.BoolOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolOutput { return v.EnableStackdriverLogging }).(pulumi.BoolOutput)
}

// Option to enable Stackdriver Monitoring.
func (o InstanceOutput) EnableStackdriverMonitoring() pulumi.BoolOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolOutput { return v.EnableStackdriverMonitoring }).(pulumi.BoolOutput)
}

// Option to enable granular zone separation.
func (o InstanceOutput) EnableZoneSeparation() pulumi.BoolOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolOutput { return v.EnableZoneSeparation }).(pulumi.BoolOutput)
}

// Option to enable and pass metadata for event publishing.
func (o InstanceOutput) EventPublishConfig() EventPublishConfigResponseOutput {
	return o.ApplyT(func(v *Instance) EventPublishConfigResponseOutput { return v.EventPublishConfig }).(EventPublishConfigResponseOutput)
}

// Cloud Storage bucket generated by Data Fusion in the customer project.
func (o InstanceOutput) GcsBucket() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.GcsBucket }).(pulumi.StringOutput)
}

// Required. The name of the instance to create.
func (o InstanceOutput) InstanceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.InstanceId }).(pulumi.StringOutput)
}

// The resource labels for instance to use to annotate any related underlying resources such as Compute Engine VMs. The character '=' is not allowed to be used within the labels.
func (o InstanceOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

func (o InstanceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The name of this instance is in the form of projects/{project}/locations/{location}/instances/{instance}.
func (o InstanceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Network configuration options. These are required when a private Data Fusion instance is to be created.
func (o InstanceOutput) NetworkConfig() NetworkConfigResponseOutput {
	return o.ApplyT(func(v *Instance) NetworkConfigResponseOutput { return v.NetworkConfig }).(NetworkConfigResponseOutput)
}

// Map of additional options used to configure the behavior of Data Fusion instance.
func (o InstanceOutput) Options() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringMapOutput { return v.Options }).(pulumi.StringMapOutput)
}

// P4 service account for the customer project.
func (o InstanceOutput) P4ServiceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.P4ServiceAccount }).(pulumi.StringOutput)
}

// Optional. Current patch revision of the Data Fusion.
func (o InstanceOutput) PatchRevision() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.PatchRevision }).(pulumi.StringOutput)
}

// Specifies whether the Data Fusion instance should be private. If set to true, all Data Fusion nodes will have private IP addresses and will not be able to access the public internet.
func (o InstanceOutput) PrivateInstance() pulumi.BoolOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolOutput { return v.PrivateInstance }).(pulumi.BoolOutput)
}

func (o InstanceOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Reserved for future use.
func (o InstanceOutput) SatisfiesPzs() pulumi.BoolOutput {
	return o.ApplyT(func(v *Instance) pulumi.BoolOutput { return v.SatisfiesPzs }).(pulumi.BoolOutput)
}

// Deprecated. Use tenant_project_id instead to extract the tenant project ID.
//
// Deprecated: Output only. Deprecated. Use tenant_project_id instead to extract the tenant project ID.
func (o InstanceOutput) ServiceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ServiceAccount }).(pulumi.StringOutput)
}

// Endpoint on which the Data Fusion UI is accessible.
func (o InstanceOutput) ServiceEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.ServiceEndpoint }).(pulumi.StringOutput)
}

// The current state of this Data Fusion instance.
func (o InstanceOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Additional information about the current state of this Data Fusion instance if available.
func (o InstanceOutput) StateMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.StateMessage }).(pulumi.StringOutput)
}

// The name of the tenant project.
func (o InstanceOutput) TenantProjectId() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.TenantProjectId }).(pulumi.StringOutput)
}

// Instance type.
func (o InstanceOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The time the instance was last updated.
func (o InstanceOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// Current version of the Data Fusion. Only specifiable in Update.
func (o InstanceOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

// Endpoint on which the Data Fusion UI is accessible to third-party users
func (o InstanceOutput) WorkforceIdentityServiceEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.WorkforceIdentityServiceEndpoint }).(pulumi.StringOutput)
}

// Name of the zone in which the Data Fusion instance will be created. Only DEVELOPER instances use this field.
func (o InstanceOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *Instance) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*InstanceInput)(nil)).Elem(), &Instance{})
	pulumi.RegisterOutputType(InstanceOutput{})
}
