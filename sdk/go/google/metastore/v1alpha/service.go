// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a metastore service in a project and location.
type Service struct {
	pulumi.CustomResourceState

	// A Cloud Storage URI (starting with gs://) that specifies where artifacts related to the metastore service are stored.
	ArtifactGcsUri pulumi.StringOutput `pulumi:"artifactGcsUri"`
	// The time when the metastore service was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Immutable. The database type that the Metastore service stores its data.
	DatabaseType pulumi.StringOutput `pulumi:"databaseType"`
	// Immutable. Information used to configure the Dataproc Metastore service to encrypt customer data at rest. Cannot be updated.
	EncryptionConfig EncryptionConfigResponseOutput `pulumi:"encryptionConfig"`
	// The URI of the endpoint used to access the metastore service.
	EndpointUri pulumi.StringOutput `pulumi:"endpointUri"`
	// Configuration information specific to running Hive metastore software as the metastore service.
	HiveMetastoreConfig HiveMetastoreConfigResponseOutput `pulumi:"hiveMetastoreConfig"`
	// User-defined labels for the metastore service.
	Labels   pulumi.StringMapOutput `pulumi:"labels"`
	Location pulumi.StringOutput    `pulumi:"location"`
	// The one hour maintenance window of the metastore service. This specifies when the service can be restarted for maintenance purposes in UTC time. Maintenance window is not needed for services with the SPANNER database type.
	MaintenanceWindow MaintenanceWindowResponseOutput `pulumi:"maintenanceWindow"`
	// The setting that defines how metastore metadata should be integrated with external services and systems.
	MetadataIntegration MetadataIntegrationResponseOutput `pulumi:"metadataIntegration"`
	// The metadata management activities of the metastore service.
	MetadataManagementActivity MetadataManagementActivityResponseOutput `pulumi:"metadataManagementActivity"`
	// Immutable. The relative resource name of the metastore service, in the following format:projects/{project_number}/locations/{location_id}/services/{service_id}.
	Name pulumi.StringOutput `pulumi:"name"`
	// Immutable. The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:projects/{project_number}/global/networks/{network_id}.
	Network pulumi.StringOutput `pulumi:"network"`
	// The configuration specifying the network settings for the Dataproc Metastore service.
	NetworkConfig NetworkConfigResponseOutput `pulumi:"networkConfig"`
	// The TCP port at which the metastore service is reached. Default: 9083.
	Port    pulumi.IntOutput    `pulumi:"port"`
	Project pulumi.StringOutput `pulumi:"project"`
	// Immutable. The release channel of the service. If unspecified, defaults to STABLE.
	ReleaseChannel pulumi.StringOutput `pulumi:"releaseChannel"`
	// Optional. A request ID. Specify a unique request ID to allow the server to ignore the request if it has completed. The server will ignore subsequent requests that provide a duplicate request ID for at least 60 minutes after the first request.For example, if an initial request times out, followed by another request with the same request ID, the server ignores the second request to prevent the creation of duplicate commitments.The request ID must be a valid UUID (https://en.wikipedia.org/wiki/Universally_unique_identifier#Format) A zero UUID (00000000-0000-0000-0000-000000000000) is not supported.
	RequestId pulumi.StringPtrOutput `pulumi:"requestId"`
	// Scaling configuration of the metastore service.
	ScalingConfig ScalingConfigResponseOutput `pulumi:"scalingConfig"`
	// Required. The ID of the metastore service, which is used as the final component of the metastore service's name.This value must be between 2 and 63 characters long inclusive, begin with a letter, end with a letter or number, and consist of alpha-numeric ASCII characters or hyphens.
	ServiceId pulumi.StringOutput `pulumi:"serviceId"`
	// The current state of the metastore service.
	State pulumi.StringOutput `pulumi:"state"`
	// Additional information about the current state of the metastore service, if available.
	StateMessage pulumi.StringOutput `pulumi:"stateMessage"`
	// The configuration specifying telemetry settings for the Dataproc Metastore service. If unspecified defaults to JSON.
	TelemetryConfig TelemetryConfigResponseOutput `pulumi:"telemetryConfig"`
	// The tier of the service.
	Tier pulumi.StringOutput `pulumi:"tier"`
	// The globally unique resource identifier of the metastore service.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// The time when the metastore service was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ServiceId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceId'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"project",
		"serviceId",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Service
	err := ctx.RegisterResource("google-native:metastore/v1alpha:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("google-native:metastore/v1alpha:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type serviceState struct {
}

type ServiceState struct {
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	// Immutable. The database type that the Metastore service stores its data.
	DatabaseType *ServiceDatabaseType `pulumi:"databaseType"`
	// Immutable. Information used to configure the Dataproc Metastore service to encrypt customer data at rest. Cannot be updated.
	EncryptionConfig *EncryptionConfig `pulumi:"encryptionConfig"`
	// Configuration information specific to running Hive metastore software as the metastore service.
	HiveMetastoreConfig *HiveMetastoreConfig `pulumi:"hiveMetastoreConfig"`
	// User-defined labels for the metastore service.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	// The one hour maintenance window of the metastore service. This specifies when the service can be restarted for maintenance purposes in UTC time. Maintenance window is not needed for services with the SPANNER database type.
	MaintenanceWindow *MaintenanceWindow `pulumi:"maintenanceWindow"`
	// The setting that defines how metastore metadata should be integrated with external services and systems.
	MetadataIntegration *MetadataIntegration `pulumi:"metadataIntegration"`
	// Immutable. The relative resource name of the metastore service, in the following format:projects/{project_number}/locations/{location_id}/services/{service_id}.
	Name *string `pulumi:"name"`
	// Immutable. The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:projects/{project_number}/global/networks/{network_id}.
	Network *string `pulumi:"network"`
	// The configuration specifying the network settings for the Dataproc Metastore service.
	NetworkConfig *NetworkConfig `pulumi:"networkConfig"`
	// The TCP port at which the metastore service is reached. Default: 9083.
	Port    *int    `pulumi:"port"`
	Project *string `pulumi:"project"`
	// Immutable. The release channel of the service. If unspecified, defaults to STABLE.
	ReleaseChannel *ServiceReleaseChannel `pulumi:"releaseChannel"`
	// Optional. A request ID. Specify a unique request ID to allow the server to ignore the request if it has completed. The server will ignore subsequent requests that provide a duplicate request ID for at least 60 minutes after the first request.For example, if an initial request times out, followed by another request with the same request ID, the server ignores the second request to prevent the creation of duplicate commitments.The request ID must be a valid UUID (https://en.wikipedia.org/wiki/Universally_unique_identifier#Format) A zero UUID (00000000-0000-0000-0000-000000000000) is not supported.
	RequestId *string `pulumi:"requestId"`
	// Scaling configuration of the metastore service.
	ScalingConfig *ScalingConfig `pulumi:"scalingConfig"`
	// Required. The ID of the metastore service, which is used as the final component of the metastore service's name.This value must be between 2 and 63 characters long inclusive, begin with a letter, end with a letter or number, and consist of alpha-numeric ASCII characters or hyphens.
	ServiceId string `pulumi:"serviceId"`
	// The configuration specifying telemetry settings for the Dataproc Metastore service. If unspecified defaults to JSON.
	TelemetryConfig *TelemetryConfig `pulumi:"telemetryConfig"`
	// The tier of the service.
	Tier *ServiceTier `pulumi:"tier"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// Immutable. The database type that the Metastore service stores its data.
	DatabaseType ServiceDatabaseTypePtrInput
	// Immutable. Information used to configure the Dataproc Metastore service to encrypt customer data at rest. Cannot be updated.
	EncryptionConfig EncryptionConfigPtrInput
	// Configuration information specific to running Hive metastore software as the metastore service.
	HiveMetastoreConfig HiveMetastoreConfigPtrInput
	// User-defined labels for the metastore service.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	// The one hour maintenance window of the metastore service. This specifies when the service can be restarted for maintenance purposes in UTC time. Maintenance window is not needed for services with the SPANNER database type.
	MaintenanceWindow MaintenanceWindowPtrInput
	// The setting that defines how metastore metadata should be integrated with external services and systems.
	MetadataIntegration MetadataIntegrationPtrInput
	// Immutable. The relative resource name of the metastore service, in the following format:projects/{project_number}/locations/{location_id}/services/{service_id}.
	Name pulumi.StringPtrInput
	// Immutable. The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:projects/{project_number}/global/networks/{network_id}.
	Network pulumi.StringPtrInput
	// The configuration specifying the network settings for the Dataproc Metastore service.
	NetworkConfig NetworkConfigPtrInput
	// The TCP port at which the metastore service is reached. Default: 9083.
	Port    pulumi.IntPtrInput
	Project pulumi.StringPtrInput
	// Immutable. The release channel of the service. If unspecified, defaults to STABLE.
	ReleaseChannel ServiceReleaseChannelPtrInput
	// Optional. A request ID. Specify a unique request ID to allow the server to ignore the request if it has completed. The server will ignore subsequent requests that provide a duplicate request ID for at least 60 minutes after the first request.For example, if an initial request times out, followed by another request with the same request ID, the server ignores the second request to prevent the creation of duplicate commitments.The request ID must be a valid UUID (https://en.wikipedia.org/wiki/Universally_unique_identifier#Format) A zero UUID (00000000-0000-0000-0000-000000000000) is not supported.
	RequestId pulumi.StringPtrInput
	// Scaling configuration of the metastore service.
	ScalingConfig ScalingConfigPtrInput
	// Required. The ID of the metastore service, which is used as the final component of the metastore service's name.This value must be between 2 and 63 characters long inclusive, begin with a letter, end with a letter or number, and consist of alpha-numeric ASCII characters or hyphens.
	ServiceId pulumi.StringInput
	// The configuration specifying telemetry settings for the Dataproc Metastore service. If unspecified defaults to JSON.
	TelemetryConfig TelemetryConfigPtrInput
	// The tier of the service.
	Tier ServiceTierPtrInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}

type ServiceInput interface {
	pulumi.Input

	ToServiceOutput() ServiceOutput
	ToServiceOutputWithContext(ctx context.Context) ServiceOutput
}

func (*Service) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (i *Service) ToServiceOutput() ServiceOutput {
	return i.ToServiceOutputWithContext(context.Background())
}

func (i *Service) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceOutput)
}

type ServiceOutput struct{ *pulumi.OutputState }

func (ServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (o ServiceOutput) ToServiceOutput() ServiceOutput {
	return o
}

func (o ServiceOutput) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return o
}

// A Cloud Storage URI (starting with gs://) that specifies where artifacts related to the metastore service are stored.
func (o ServiceOutput) ArtifactGcsUri() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.ArtifactGcsUri }).(pulumi.StringOutput)
}

// The time when the metastore service was created.
func (o ServiceOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Immutable. The database type that the Metastore service stores its data.
func (o ServiceOutput) DatabaseType() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.DatabaseType }).(pulumi.StringOutput)
}

// Immutable. Information used to configure the Dataproc Metastore service to encrypt customer data at rest. Cannot be updated.
func (o ServiceOutput) EncryptionConfig() EncryptionConfigResponseOutput {
	return o.ApplyT(func(v *Service) EncryptionConfigResponseOutput { return v.EncryptionConfig }).(EncryptionConfigResponseOutput)
}

// The URI of the endpoint used to access the metastore service.
func (o ServiceOutput) EndpointUri() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.EndpointUri }).(pulumi.StringOutput)
}

// Configuration information specific to running Hive metastore software as the metastore service.
func (o ServiceOutput) HiveMetastoreConfig() HiveMetastoreConfigResponseOutput {
	return o.ApplyT(func(v *Service) HiveMetastoreConfigResponseOutput { return v.HiveMetastoreConfig }).(HiveMetastoreConfigResponseOutput)
}

// User-defined labels for the metastore service.
func (o ServiceOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Service) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

func (o ServiceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The one hour maintenance window of the metastore service. This specifies when the service can be restarted for maintenance purposes in UTC time. Maintenance window is not needed for services with the SPANNER database type.
func (o ServiceOutput) MaintenanceWindow() MaintenanceWindowResponseOutput {
	return o.ApplyT(func(v *Service) MaintenanceWindowResponseOutput { return v.MaintenanceWindow }).(MaintenanceWindowResponseOutput)
}

// The setting that defines how metastore metadata should be integrated with external services and systems.
func (o ServiceOutput) MetadataIntegration() MetadataIntegrationResponseOutput {
	return o.ApplyT(func(v *Service) MetadataIntegrationResponseOutput { return v.MetadataIntegration }).(MetadataIntegrationResponseOutput)
}

// The metadata management activities of the metastore service.
func (o ServiceOutput) MetadataManagementActivity() MetadataManagementActivityResponseOutput {
	return o.ApplyT(func(v *Service) MetadataManagementActivityResponseOutput { return v.MetadataManagementActivity }).(MetadataManagementActivityResponseOutput)
}

// Immutable. The relative resource name of the metastore service, in the following format:projects/{project_number}/locations/{location_id}/services/{service_id}.
func (o ServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Immutable. The relative resource name of the VPC network on which the instance can be accessed. It is specified in the following form:projects/{project_number}/global/networks/{network_id}.
func (o ServiceOutput) Network() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Network }).(pulumi.StringOutput)
}

// The configuration specifying the network settings for the Dataproc Metastore service.
func (o ServiceOutput) NetworkConfig() NetworkConfigResponseOutput {
	return o.ApplyT(func(v *Service) NetworkConfigResponseOutput { return v.NetworkConfig }).(NetworkConfigResponseOutput)
}

// The TCP port at which the metastore service is reached. Default: 9083.
func (o ServiceOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *Service) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

func (o ServiceOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Immutable. The release channel of the service. If unspecified, defaults to STABLE.
func (o ServiceOutput) ReleaseChannel() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.ReleaseChannel }).(pulumi.StringOutput)
}

// Optional. A request ID. Specify a unique request ID to allow the server to ignore the request if it has completed. The server will ignore subsequent requests that provide a duplicate request ID for at least 60 minutes after the first request.For example, if an initial request times out, followed by another request with the same request ID, the server ignores the second request to prevent the creation of duplicate commitments.The request ID must be a valid UUID (https://en.wikipedia.org/wiki/Universally_unique_identifier#Format) A zero UUID (00000000-0000-0000-0000-000000000000) is not supported.
func (o ServiceOutput) RequestId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.RequestId }).(pulumi.StringPtrOutput)
}

// Scaling configuration of the metastore service.
func (o ServiceOutput) ScalingConfig() ScalingConfigResponseOutput {
	return o.ApplyT(func(v *Service) ScalingConfigResponseOutput { return v.ScalingConfig }).(ScalingConfigResponseOutput)
}

// Required. The ID of the metastore service, which is used as the final component of the metastore service's name.This value must be between 2 and 63 characters long inclusive, begin with a letter, end with a letter or number, and consist of alpha-numeric ASCII characters or hyphens.
func (o ServiceOutput) ServiceId() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.ServiceId }).(pulumi.StringOutput)
}

// The current state of the metastore service.
func (o ServiceOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Additional information about the current state of the metastore service, if available.
func (o ServiceOutput) StateMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.StateMessage }).(pulumi.StringOutput)
}

// The configuration specifying telemetry settings for the Dataproc Metastore service. If unspecified defaults to JSON.
func (o ServiceOutput) TelemetryConfig() TelemetryConfigResponseOutput {
	return o.ApplyT(func(v *Service) TelemetryConfigResponseOutput { return v.TelemetryConfig }).(TelemetryConfigResponseOutput)
}

// The tier of the service.
func (o ServiceOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Tier }).(pulumi.StringOutput)
}

// The globally unique resource identifier of the metastore service.
func (o ServiceOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// The time when the metastore service was last updated.
func (o ServiceOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceInput)(nil)).Elem(), &Service{})
	pulumi.RegisterOutputType(ServiceOutput{})
}
