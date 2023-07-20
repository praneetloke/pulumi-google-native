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

// Creates a new bare metal cluster in a given project and location.
type BareMetalCluster struct {
	pulumi.CustomResourceState

	// The admin cluster this bare metal user cluster belongs to. This is the full resource name of the admin cluster's fleet membership.
	AdminClusterMembership pulumi.StringOutput `pulumi:"adminClusterMembership"`
	// The resource name of the bare metal admin cluster managing this user cluster.
	AdminClusterName pulumi.StringOutput `pulumi:"adminClusterName"`
	// Annotations on the bare metal user cluster. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
	Annotations pulumi.StringMapOutput `pulumi:"annotations"`
	// Required. User provided identifier that is used as part of the resource name; must conform to RFC-1034 and additionally restrict to lower-cased letters. This comes out roughly to: /^a-z+[a-z0-9]$/
	BareMetalClusterId pulumi.StringOutput `pulumi:"bareMetalClusterId"`
	// The Anthos clusters on bare metal version for your user cluster.
	BareMetalVersion pulumi.StringOutput `pulumi:"bareMetalVersion"`
	// Cluster operations configuration.
	ClusterOperations BareMetalClusterOperationsConfigResponseOutput `pulumi:"clusterOperations"`
	// Control plane configuration.
	ControlPlane BareMetalControlPlaneConfigResponseOutput `pulumi:"controlPlane"`
	// The time when the bare metal user cluster was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The time when the bare metal user cluster was deleted. If the resource is not deleted, this must be empty
	DeleteTime pulumi.StringOutput `pulumi:"deleteTime"`
	// A human readable description of this bare metal user cluster.
	Description pulumi.StringOutput `pulumi:"description"`
	// The IP address of the bare metal user cluster's API server.
	Endpoint pulumi.StringOutput `pulumi:"endpoint"`
	// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding. Allows clients to perform consistent read-modify-writes through optimistic concurrency control.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Fleet configuration for the cluster.
	Fleet FleetResponseOutput `pulumi:"fleet"`
	// Load balancer configuration.
	LoadBalancer BareMetalLoadBalancerConfigResponseOutput `pulumi:"loadBalancer"`
	// The object name of the bare metal user cluster custom resource on the associated admin cluster. This field is used to support conflicting names when enrolling existing clusters to the API. When used as a part of cluster enrollment, this field will differ from the name in the resource name. For new clusters, this field will match the user provided cluster name and be visible in the last component of the resource name. It is not modifiable. When the local name and cluster name differ, the local name is used in the admin cluster controller logs. You use the cluster name when accessing the cluster using bmctl and kubectl.
	LocalName pulumi.StringOutput `pulumi:"localName"`
	Location  pulumi.StringOutput `pulumi:"location"`
	// Maintenance configuration.
	MaintenanceConfig BareMetalMaintenanceConfigResponseOutput `pulumi:"maintenanceConfig"`
	// Status of on-going maintenance tasks.
	MaintenanceStatus BareMetalMaintenanceStatusResponseOutput `pulumi:"maintenanceStatus"`
	// Immutable. The bare metal user cluster resource name.
	Name pulumi.StringOutput `pulumi:"name"`
	// Network configuration.
	NetworkConfig BareMetalNetworkConfigResponseOutput `pulumi:"networkConfig"`
	// Node access related configurations.
	NodeAccessConfig BareMetalNodeAccessConfigResponseOutput `pulumi:"nodeAccessConfig"`
	// Workload node configuration.
	NodeConfig BareMetalWorkloadNodeConfigResponseOutput `pulumi:"nodeConfig"`
	// OS environment related configurations.
	OsEnvironmentConfig BareMetalOsEnvironmentConfigResponseOutput `pulumi:"osEnvironmentConfig"`
	Project             pulumi.StringOutput                        `pulumi:"project"`
	// Proxy configuration.
	Proxy BareMetalProxyConfigResponseOutput `pulumi:"proxy"`
	// If set, there are currently changes in flight to the bare metal user cluster.
	Reconciling pulumi.BoolOutput `pulumi:"reconciling"`
	// Security related setting configuration.
	SecurityConfig BareMetalSecurityConfigResponseOutput `pulumi:"securityConfig"`
	// The current state of the bare metal user cluster.
	State pulumi.StringOutput `pulumi:"state"`
	// Detailed cluster status.
	Status ResourceStatusResponseOutput `pulumi:"status"`
	// Storage configuration.
	Storage BareMetalStorageConfigResponseOutput `pulumi:"storage"`
	// The unique identifier of the bare metal user cluster.
	Uid pulumi.StringOutput `pulumi:"uid"`
	// The time when the bare metal user cluster was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// The result of the preflight check.
	ValidationCheck ValidationCheckResponseOutput `pulumi:"validationCheck"`
}

// NewBareMetalCluster registers a new resource with the given unique name, arguments, and options.
func NewBareMetalCluster(ctx *pulumi.Context,
	name string, args *BareMetalClusterArgs, opts ...pulumi.ResourceOption) (*BareMetalCluster, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AdminClusterMembership == nil {
		return nil, errors.New("invalid value for required argument 'AdminClusterMembership'")
	}
	if args.BareMetalClusterId == nil {
		return nil, errors.New("invalid value for required argument 'BareMetalClusterId'")
	}
	if args.BareMetalVersion == nil {
		return nil, errors.New("invalid value for required argument 'BareMetalVersion'")
	}
	if args.ControlPlane == nil {
		return nil, errors.New("invalid value for required argument 'ControlPlane'")
	}
	if args.LoadBalancer == nil {
		return nil, errors.New("invalid value for required argument 'LoadBalancer'")
	}
	if args.NetworkConfig == nil {
		return nil, errors.New("invalid value for required argument 'NetworkConfig'")
	}
	if args.Storage == nil {
		return nil, errors.New("invalid value for required argument 'Storage'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"bareMetalClusterId",
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource BareMetalCluster
	err := ctx.RegisterResource("google-native:gkeonprem/v1:BareMetalCluster", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBareMetalCluster gets an existing BareMetalCluster resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBareMetalCluster(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BareMetalClusterState, opts ...pulumi.ResourceOption) (*BareMetalCluster, error) {
	var resource BareMetalCluster
	err := ctx.ReadResource("google-native:gkeonprem/v1:BareMetalCluster", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BareMetalCluster resources.
type bareMetalClusterState struct {
}

type BareMetalClusterState struct {
}

func (BareMetalClusterState) ElementType() reflect.Type {
	return reflect.TypeOf((*bareMetalClusterState)(nil)).Elem()
}

type bareMetalClusterArgs struct {
	// The admin cluster this bare metal user cluster belongs to. This is the full resource name of the admin cluster's fleet membership.
	AdminClusterMembership string `pulumi:"adminClusterMembership"`
	// Annotations on the bare metal user cluster. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
	Annotations map[string]string `pulumi:"annotations"`
	// Required. User provided identifier that is used as part of the resource name; must conform to RFC-1034 and additionally restrict to lower-cased letters. This comes out roughly to: /^a-z+[a-z0-9]$/
	BareMetalClusterId string `pulumi:"bareMetalClusterId"`
	// The Anthos clusters on bare metal version for your user cluster.
	BareMetalVersion string `pulumi:"bareMetalVersion"`
	// Cluster operations configuration.
	ClusterOperations *BareMetalClusterOperationsConfig `pulumi:"clusterOperations"`
	// Control plane configuration.
	ControlPlane BareMetalControlPlaneConfig `pulumi:"controlPlane"`
	// A human readable description of this bare metal user cluster.
	Description *string `pulumi:"description"`
	// Load balancer configuration.
	LoadBalancer BareMetalLoadBalancerConfig `pulumi:"loadBalancer"`
	Location     *string                     `pulumi:"location"`
	// Maintenance configuration.
	MaintenanceConfig *BareMetalMaintenanceConfig `pulumi:"maintenanceConfig"`
	// Immutable. The bare metal user cluster resource name.
	Name *string `pulumi:"name"`
	// Network configuration.
	NetworkConfig BareMetalNetworkConfig `pulumi:"networkConfig"`
	// Node access related configurations.
	NodeAccessConfig *BareMetalNodeAccessConfig `pulumi:"nodeAccessConfig"`
	// Workload node configuration.
	NodeConfig *BareMetalWorkloadNodeConfig `pulumi:"nodeConfig"`
	// OS environment related configurations.
	OsEnvironmentConfig *BareMetalOsEnvironmentConfig `pulumi:"osEnvironmentConfig"`
	Project             *string                       `pulumi:"project"`
	// Proxy configuration.
	Proxy *BareMetalProxyConfig `pulumi:"proxy"`
	// Security related setting configuration.
	SecurityConfig *BareMetalSecurityConfig `pulumi:"securityConfig"`
	// Storage configuration.
	Storage BareMetalStorageConfig `pulumi:"storage"`
}

// The set of arguments for constructing a BareMetalCluster resource.
type BareMetalClusterArgs struct {
	// The admin cluster this bare metal user cluster belongs to. This is the full resource name of the admin cluster's fleet membership.
	AdminClusterMembership pulumi.StringInput
	// Annotations on the bare metal user cluster. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
	Annotations pulumi.StringMapInput
	// Required. User provided identifier that is used as part of the resource name; must conform to RFC-1034 and additionally restrict to lower-cased letters. This comes out roughly to: /^a-z+[a-z0-9]$/
	BareMetalClusterId pulumi.StringInput
	// The Anthos clusters on bare metal version for your user cluster.
	BareMetalVersion pulumi.StringInput
	// Cluster operations configuration.
	ClusterOperations BareMetalClusterOperationsConfigPtrInput
	// Control plane configuration.
	ControlPlane BareMetalControlPlaneConfigInput
	// A human readable description of this bare metal user cluster.
	Description pulumi.StringPtrInput
	// Load balancer configuration.
	LoadBalancer BareMetalLoadBalancerConfigInput
	Location     pulumi.StringPtrInput
	// Maintenance configuration.
	MaintenanceConfig BareMetalMaintenanceConfigPtrInput
	// Immutable. The bare metal user cluster resource name.
	Name pulumi.StringPtrInput
	// Network configuration.
	NetworkConfig BareMetalNetworkConfigInput
	// Node access related configurations.
	NodeAccessConfig BareMetalNodeAccessConfigPtrInput
	// Workload node configuration.
	NodeConfig BareMetalWorkloadNodeConfigPtrInput
	// OS environment related configurations.
	OsEnvironmentConfig BareMetalOsEnvironmentConfigPtrInput
	Project             pulumi.StringPtrInput
	// Proxy configuration.
	Proxy BareMetalProxyConfigPtrInput
	// Security related setting configuration.
	SecurityConfig BareMetalSecurityConfigPtrInput
	// Storage configuration.
	Storage BareMetalStorageConfigInput
}

func (BareMetalClusterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*bareMetalClusterArgs)(nil)).Elem()
}

type BareMetalClusterInput interface {
	pulumi.Input

	ToBareMetalClusterOutput() BareMetalClusterOutput
	ToBareMetalClusterOutputWithContext(ctx context.Context) BareMetalClusterOutput
}

func (*BareMetalCluster) ElementType() reflect.Type {
	return reflect.TypeOf((**BareMetalCluster)(nil)).Elem()
}

func (i *BareMetalCluster) ToBareMetalClusterOutput() BareMetalClusterOutput {
	return i.ToBareMetalClusterOutputWithContext(context.Background())
}

func (i *BareMetalCluster) ToBareMetalClusterOutputWithContext(ctx context.Context) BareMetalClusterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BareMetalClusterOutput)
}

type BareMetalClusterOutput struct{ *pulumi.OutputState }

func (BareMetalClusterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BareMetalCluster)(nil)).Elem()
}

func (o BareMetalClusterOutput) ToBareMetalClusterOutput() BareMetalClusterOutput {
	return o
}

func (o BareMetalClusterOutput) ToBareMetalClusterOutputWithContext(ctx context.Context) BareMetalClusterOutput {
	return o
}

// The admin cluster this bare metal user cluster belongs to. This is the full resource name of the admin cluster's fleet membership.
func (o BareMetalClusterOutput) AdminClusterMembership() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalCluster) pulumi.StringOutput { return v.AdminClusterMembership }).(pulumi.StringOutput)
}

// The resource name of the bare metal admin cluster managing this user cluster.
func (o BareMetalClusterOutput) AdminClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalCluster) pulumi.StringOutput { return v.AdminClusterName }).(pulumi.StringOutput)
}

// Annotations on the bare metal user cluster. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
func (o BareMetalClusterOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BareMetalCluster) pulumi.StringMapOutput { return v.Annotations }).(pulumi.StringMapOutput)
}

// Required. User provided identifier that is used as part of the resource name; must conform to RFC-1034 and additionally restrict to lower-cased letters. This comes out roughly to: /^a-z+[a-z0-9]$/
func (o BareMetalClusterOutput) BareMetalClusterId() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalCluster) pulumi.StringOutput { return v.BareMetalClusterId }).(pulumi.StringOutput)
}

// The Anthos clusters on bare metal version for your user cluster.
func (o BareMetalClusterOutput) BareMetalVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalCluster) pulumi.StringOutput { return v.BareMetalVersion }).(pulumi.StringOutput)
}

// Cluster operations configuration.
func (o BareMetalClusterOutput) ClusterOperations() BareMetalClusterOperationsConfigResponseOutput {
	return o.ApplyT(func(v *BareMetalCluster) BareMetalClusterOperationsConfigResponseOutput { return v.ClusterOperations }).(BareMetalClusterOperationsConfigResponseOutput)
}

// Control plane configuration.
func (o BareMetalClusterOutput) ControlPlane() BareMetalControlPlaneConfigResponseOutput {
	return o.ApplyT(func(v *BareMetalCluster) BareMetalControlPlaneConfigResponseOutput { return v.ControlPlane }).(BareMetalControlPlaneConfigResponseOutput)
}

// The time when the bare metal user cluster was created.
func (o BareMetalClusterOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalCluster) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The time when the bare metal user cluster was deleted. If the resource is not deleted, this must be empty
func (o BareMetalClusterOutput) DeleteTime() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalCluster) pulumi.StringOutput { return v.DeleteTime }).(pulumi.StringOutput)
}

// A human readable description of this bare metal user cluster.
func (o BareMetalClusterOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalCluster) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The IP address of the bare metal user cluster's API server.
func (o BareMetalClusterOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalCluster) pulumi.StringOutput { return v.Endpoint }).(pulumi.StringOutput)
}

// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding. Allows clients to perform consistent read-modify-writes through optimistic concurrency control.
func (o BareMetalClusterOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalCluster) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Fleet configuration for the cluster.
func (o BareMetalClusterOutput) Fleet() FleetResponseOutput {
	return o.ApplyT(func(v *BareMetalCluster) FleetResponseOutput { return v.Fleet }).(FleetResponseOutput)
}

// Load balancer configuration.
func (o BareMetalClusterOutput) LoadBalancer() BareMetalLoadBalancerConfigResponseOutput {
	return o.ApplyT(func(v *BareMetalCluster) BareMetalLoadBalancerConfigResponseOutput { return v.LoadBalancer }).(BareMetalLoadBalancerConfigResponseOutput)
}

// The object name of the bare metal user cluster custom resource on the associated admin cluster. This field is used to support conflicting names when enrolling existing clusters to the API. When used as a part of cluster enrollment, this field will differ from the name in the resource name. For new clusters, this field will match the user provided cluster name and be visible in the last component of the resource name. It is not modifiable. When the local name and cluster name differ, the local name is used in the admin cluster controller logs. You use the cluster name when accessing the cluster using bmctl and kubectl.
func (o BareMetalClusterOutput) LocalName() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalCluster) pulumi.StringOutput { return v.LocalName }).(pulumi.StringOutput)
}

func (o BareMetalClusterOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalCluster) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Maintenance configuration.
func (o BareMetalClusterOutput) MaintenanceConfig() BareMetalMaintenanceConfigResponseOutput {
	return o.ApplyT(func(v *BareMetalCluster) BareMetalMaintenanceConfigResponseOutput { return v.MaintenanceConfig }).(BareMetalMaintenanceConfigResponseOutput)
}

// Status of on-going maintenance tasks.
func (o BareMetalClusterOutput) MaintenanceStatus() BareMetalMaintenanceStatusResponseOutput {
	return o.ApplyT(func(v *BareMetalCluster) BareMetalMaintenanceStatusResponseOutput { return v.MaintenanceStatus }).(BareMetalMaintenanceStatusResponseOutput)
}

// Immutable. The bare metal user cluster resource name.
func (o BareMetalClusterOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalCluster) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Network configuration.
func (o BareMetalClusterOutput) NetworkConfig() BareMetalNetworkConfigResponseOutput {
	return o.ApplyT(func(v *BareMetalCluster) BareMetalNetworkConfigResponseOutput { return v.NetworkConfig }).(BareMetalNetworkConfigResponseOutput)
}

// Node access related configurations.
func (o BareMetalClusterOutput) NodeAccessConfig() BareMetalNodeAccessConfigResponseOutput {
	return o.ApplyT(func(v *BareMetalCluster) BareMetalNodeAccessConfigResponseOutput { return v.NodeAccessConfig }).(BareMetalNodeAccessConfigResponseOutput)
}

// Workload node configuration.
func (o BareMetalClusterOutput) NodeConfig() BareMetalWorkloadNodeConfigResponseOutput {
	return o.ApplyT(func(v *BareMetalCluster) BareMetalWorkloadNodeConfigResponseOutput { return v.NodeConfig }).(BareMetalWorkloadNodeConfigResponseOutput)
}

// OS environment related configurations.
func (o BareMetalClusterOutput) OsEnvironmentConfig() BareMetalOsEnvironmentConfigResponseOutput {
	return o.ApplyT(func(v *BareMetalCluster) BareMetalOsEnvironmentConfigResponseOutput { return v.OsEnvironmentConfig }).(BareMetalOsEnvironmentConfigResponseOutput)
}

func (o BareMetalClusterOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalCluster) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Proxy configuration.
func (o BareMetalClusterOutput) Proxy() BareMetalProxyConfigResponseOutput {
	return o.ApplyT(func(v *BareMetalCluster) BareMetalProxyConfigResponseOutput { return v.Proxy }).(BareMetalProxyConfigResponseOutput)
}

// If set, there are currently changes in flight to the bare metal user cluster.
func (o BareMetalClusterOutput) Reconciling() pulumi.BoolOutput {
	return o.ApplyT(func(v *BareMetalCluster) pulumi.BoolOutput { return v.Reconciling }).(pulumi.BoolOutput)
}

// Security related setting configuration.
func (o BareMetalClusterOutput) SecurityConfig() BareMetalSecurityConfigResponseOutput {
	return o.ApplyT(func(v *BareMetalCluster) BareMetalSecurityConfigResponseOutput { return v.SecurityConfig }).(BareMetalSecurityConfigResponseOutput)
}

// The current state of the bare metal user cluster.
func (o BareMetalClusterOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalCluster) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Detailed cluster status.
func (o BareMetalClusterOutput) Status() ResourceStatusResponseOutput {
	return o.ApplyT(func(v *BareMetalCluster) ResourceStatusResponseOutput { return v.Status }).(ResourceStatusResponseOutput)
}

// Storage configuration.
func (o BareMetalClusterOutput) Storage() BareMetalStorageConfigResponseOutput {
	return o.ApplyT(func(v *BareMetalCluster) BareMetalStorageConfigResponseOutput { return v.Storage }).(BareMetalStorageConfigResponseOutput)
}

// The unique identifier of the bare metal user cluster.
func (o BareMetalClusterOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalCluster) pulumi.StringOutput { return v.Uid }).(pulumi.StringOutput)
}

// The time when the bare metal user cluster was last updated.
func (o BareMetalClusterOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *BareMetalCluster) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// The result of the preflight check.
func (o BareMetalClusterOutput) ValidationCheck() ValidationCheckResponseOutput {
	return o.ApplyT(func(v *BareMetalCluster) ValidationCheckResponseOutput { return v.ValidationCheck }).(ValidationCheckResponseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BareMetalClusterInput)(nil)).Elem(), &BareMetalCluster{})
	pulumi.RegisterOutputType(BareMetalClusterOutput{})
}
