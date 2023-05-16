// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single bare metal admin cluster.
func LookupBareMetalAdminCluster(ctx *pulumi.Context, args *LookupBareMetalAdminClusterArgs, opts ...pulumi.InvokeOption) (*LookupBareMetalAdminClusterResult, error) {
	var rv LookupBareMetalAdminClusterResult
	err := ctx.Invoke("google-native:gkeonprem/v1:getBareMetalAdminCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBareMetalAdminClusterArgs struct {
	BareMetalAdminClusterId string  `pulumi:"bareMetalAdminClusterId"`
	Location                string  `pulumi:"location"`
	Project                 *string `pulumi:"project"`
}

type LookupBareMetalAdminClusterResult struct {
	// Annotations on the bare metal admin cluster. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
	Annotations map[string]string `pulumi:"annotations"`
	// The Anthos clusters on bare metal version for the bare metal admin cluster.
	BareMetalVersion string `pulumi:"bareMetalVersion"`
	// Cluster operations configuration.
	ClusterOperations BareMetalAdminClusterOperationsConfigResponse `pulumi:"clusterOperations"`
	// Control plane configuration.
	ControlPlane BareMetalAdminControlPlaneConfigResponse `pulumi:"controlPlane"`
	// The time at which this bare metal admin cluster was created.
	CreateTime string `pulumi:"createTime"`
	// The time at which this bare metal admin cluster was deleted. If the resource is not deleted, this must be empty
	DeleteTime string `pulumi:"deleteTime"`
	// A human readable description of this bare metal admin cluster.
	Description string `pulumi:"description"`
	// The IP address name of bare metal admin cluster's API server.
	Endpoint string `pulumi:"endpoint"`
	// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding. Allows clients to perform consistent read-modify-writes through optimistic concurrency control.
	Etag string `pulumi:"etag"`
	// Fleet configuration for the cluster.
	Fleet FleetResponse `pulumi:"fleet"`
	// Load balancer configuration.
	LoadBalancer BareMetalAdminLoadBalancerConfigResponse `pulumi:"loadBalancer"`
	// The object name of the bare metal cluster custom resource. This field is used to support conflicting names when enrolling existing clusters to the API. When used as a part of cluster enrollment, this field will differ from the ID in the resource name. For new clusters, this field will match the user provided cluster name and be visible in the last component of the resource name. It is not modifiable. All users should use this name to access their cluster using gkectl or kubectl and should expect to see the local name when viewing admin cluster controller logs.
	LocalName string `pulumi:"localName"`
	// Maintenance configuration.
	MaintenanceConfig BareMetalAdminMaintenanceConfigResponse `pulumi:"maintenanceConfig"`
	// MaintenanceStatus representing state of maintenance.
	MaintenanceStatus BareMetalAdminMaintenanceStatusResponse `pulumi:"maintenanceStatus"`
	// Immutable. The bare metal admin cluster resource name.
	Name string `pulumi:"name"`
	// Network configuration.
	NetworkConfig BareMetalAdminNetworkConfigResponse `pulumi:"networkConfig"`
	// Node access related configurations.
	NodeAccessConfig BareMetalAdminNodeAccessConfigResponse `pulumi:"nodeAccessConfig"`
	// Workload node configuration.
	NodeConfig BareMetalAdminWorkloadNodeConfigResponse `pulumi:"nodeConfig"`
	// OS environment related configurations.
	OsEnvironmentConfig BareMetalAdminOsEnvironmentConfigResponse `pulumi:"osEnvironmentConfig"`
	// Proxy configuration.
	Proxy BareMetalAdminProxyConfigResponse `pulumi:"proxy"`
	// If set, there are currently changes in flight to the bare metal Admin Cluster.
	Reconciling bool `pulumi:"reconciling"`
	// Security related configuration.
	SecurityConfig BareMetalAdminSecurityConfigResponse `pulumi:"securityConfig"`
	// The current state of the bare metal admin cluster.
	State string `pulumi:"state"`
	// ResourceStatus representing detailed cluster status.
	Status ResourceStatusResponse `pulumi:"status"`
	// Storage configuration.
	Storage BareMetalAdminStorageConfigResponse `pulumi:"storage"`
	// The unique identifier of the bare metal admin cluster.
	Uid string `pulumi:"uid"`
	// The time at which this bare metal admin cluster was last updated.
	UpdateTime string `pulumi:"updateTime"`
	// ValidationCheck representing the result of the preflight check.
	ValidationCheck ValidationCheckResponse `pulumi:"validationCheck"`
}

func LookupBareMetalAdminClusterOutput(ctx *pulumi.Context, args LookupBareMetalAdminClusterOutputArgs, opts ...pulumi.InvokeOption) LookupBareMetalAdminClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBareMetalAdminClusterResult, error) {
			args := v.(LookupBareMetalAdminClusterArgs)
			r, err := LookupBareMetalAdminCluster(ctx, &args, opts...)
			var s LookupBareMetalAdminClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBareMetalAdminClusterResultOutput)
}

type LookupBareMetalAdminClusterOutputArgs struct {
	BareMetalAdminClusterId pulumi.StringInput    `pulumi:"bareMetalAdminClusterId"`
	Location                pulumi.StringInput    `pulumi:"location"`
	Project                 pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupBareMetalAdminClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBareMetalAdminClusterArgs)(nil)).Elem()
}

type LookupBareMetalAdminClusterResultOutput struct{ *pulumi.OutputState }

func (LookupBareMetalAdminClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBareMetalAdminClusterResult)(nil)).Elem()
}

func (o LookupBareMetalAdminClusterResultOutput) ToLookupBareMetalAdminClusterResultOutput() LookupBareMetalAdminClusterResultOutput {
	return o
}

func (o LookupBareMetalAdminClusterResultOutput) ToLookupBareMetalAdminClusterResultOutputWithContext(ctx context.Context) LookupBareMetalAdminClusterResultOutput {
	return o
}

// Annotations on the bare metal admin cluster. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
func (o LookupBareMetalAdminClusterResultOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupBareMetalAdminClusterResult) map[string]string { return v.Annotations }).(pulumi.StringMapOutput)
}

// The Anthos clusters on bare metal version for the bare metal admin cluster.
func (o LookupBareMetalAdminClusterResultOutput) BareMetalVersion() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBareMetalAdminClusterResult) string { return v.BareMetalVersion }).(pulumi.StringOutput)
}

// Cluster operations configuration.
func (o LookupBareMetalAdminClusterResultOutput) ClusterOperations() BareMetalAdminClusterOperationsConfigResponseOutput {
	return o.ApplyT(func(v LookupBareMetalAdminClusterResult) BareMetalAdminClusterOperationsConfigResponse {
		return v.ClusterOperations
	}).(BareMetalAdminClusterOperationsConfigResponseOutput)
}

// Control plane configuration.
func (o LookupBareMetalAdminClusterResultOutput) ControlPlane() BareMetalAdminControlPlaneConfigResponseOutput {
	return o.ApplyT(func(v LookupBareMetalAdminClusterResult) BareMetalAdminControlPlaneConfigResponse {
		return v.ControlPlane
	}).(BareMetalAdminControlPlaneConfigResponseOutput)
}

// The time at which this bare metal admin cluster was created.
func (o LookupBareMetalAdminClusterResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBareMetalAdminClusterResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The time at which this bare metal admin cluster was deleted. If the resource is not deleted, this must be empty
func (o LookupBareMetalAdminClusterResultOutput) DeleteTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBareMetalAdminClusterResult) string { return v.DeleteTime }).(pulumi.StringOutput)
}

// A human readable description of this bare metal admin cluster.
func (o LookupBareMetalAdminClusterResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBareMetalAdminClusterResult) string { return v.Description }).(pulumi.StringOutput)
}

// The IP address name of bare metal admin cluster's API server.
func (o LookupBareMetalAdminClusterResultOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBareMetalAdminClusterResult) string { return v.Endpoint }).(pulumi.StringOutput)
}

// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding. Allows clients to perform consistent read-modify-writes through optimistic concurrency control.
func (o LookupBareMetalAdminClusterResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBareMetalAdminClusterResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Fleet configuration for the cluster.
func (o LookupBareMetalAdminClusterResultOutput) Fleet() FleetResponseOutput {
	return o.ApplyT(func(v LookupBareMetalAdminClusterResult) FleetResponse { return v.Fleet }).(FleetResponseOutput)
}

// Load balancer configuration.
func (o LookupBareMetalAdminClusterResultOutput) LoadBalancer() BareMetalAdminLoadBalancerConfigResponseOutput {
	return o.ApplyT(func(v LookupBareMetalAdminClusterResult) BareMetalAdminLoadBalancerConfigResponse {
		return v.LoadBalancer
	}).(BareMetalAdminLoadBalancerConfigResponseOutput)
}

// The object name of the bare metal cluster custom resource. This field is used to support conflicting names when enrolling existing clusters to the API. When used as a part of cluster enrollment, this field will differ from the ID in the resource name. For new clusters, this field will match the user provided cluster name and be visible in the last component of the resource name. It is not modifiable. All users should use this name to access their cluster using gkectl or kubectl and should expect to see the local name when viewing admin cluster controller logs.
func (o LookupBareMetalAdminClusterResultOutput) LocalName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBareMetalAdminClusterResult) string { return v.LocalName }).(pulumi.StringOutput)
}

// Maintenance configuration.
func (o LookupBareMetalAdminClusterResultOutput) MaintenanceConfig() BareMetalAdminMaintenanceConfigResponseOutput {
	return o.ApplyT(func(v LookupBareMetalAdminClusterResult) BareMetalAdminMaintenanceConfigResponse {
		return v.MaintenanceConfig
	}).(BareMetalAdminMaintenanceConfigResponseOutput)
}

// MaintenanceStatus representing state of maintenance.
func (o LookupBareMetalAdminClusterResultOutput) MaintenanceStatus() BareMetalAdminMaintenanceStatusResponseOutput {
	return o.ApplyT(func(v LookupBareMetalAdminClusterResult) BareMetalAdminMaintenanceStatusResponse {
		return v.MaintenanceStatus
	}).(BareMetalAdminMaintenanceStatusResponseOutput)
}

// Immutable. The bare metal admin cluster resource name.
func (o LookupBareMetalAdminClusterResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBareMetalAdminClusterResult) string { return v.Name }).(pulumi.StringOutput)
}

// Network configuration.
func (o LookupBareMetalAdminClusterResultOutput) NetworkConfig() BareMetalAdminNetworkConfigResponseOutput {
	return o.ApplyT(func(v LookupBareMetalAdminClusterResult) BareMetalAdminNetworkConfigResponse { return v.NetworkConfig }).(BareMetalAdminNetworkConfigResponseOutput)
}

// Node access related configurations.
func (o LookupBareMetalAdminClusterResultOutput) NodeAccessConfig() BareMetalAdminNodeAccessConfigResponseOutput {
	return o.ApplyT(func(v LookupBareMetalAdminClusterResult) BareMetalAdminNodeAccessConfigResponse {
		return v.NodeAccessConfig
	}).(BareMetalAdminNodeAccessConfigResponseOutput)
}

// Workload node configuration.
func (o LookupBareMetalAdminClusterResultOutput) NodeConfig() BareMetalAdminWorkloadNodeConfigResponseOutput {
	return o.ApplyT(func(v LookupBareMetalAdminClusterResult) BareMetalAdminWorkloadNodeConfigResponse {
		return v.NodeConfig
	}).(BareMetalAdminWorkloadNodeConfigResponseOutput)
}

// OS environment related configurations.
func (o LookupBareMetalAdminClusterResultOutput) OsEnvironmentConfig() BareMetalAdminOsEnvironmentConfigResponseOutput {
	return o.ApplyT(func(v LookupBareMetalAdminClusterResult) BareMetalAdminOsEnvironmentConfigResponse {
		return v.OsEnvironmentConfig
	}).(BareMetalAdminOsEnvironmentConfigResponseOutput)
}

// Proxy configuration.
func (o LookupBareMetalAdminClusterResultOutput) Proxy() BareMetalAdminProxyConfigResponseOutput {
	return o.ApplyT(func(v LookupBareMetalAdminClusterResult) BareMetalAdminProxyConfigResponse { return v.Proxy }).(BareMetalAdminProxyConfigResponseOutput)
}

// If set, there are currently changes in flight to the bare metal Admin Cluster.
func (o LookupBareMetalAdminClusterResultOutput) Reconciling() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupBareMetalAdminClusterResult) bool { return v.Reconciling }).(pulumi.BoolOutput)
}

// Security related configuration.
func (o LookupBareMetalAdminClusterResultOutput) SecurityConfig() BareMetalAdminSecurityConfigResponseOutput {
	return o.ApplyT(func(v LookupBareMetalAdminClusterResult) BareMetalAdminSecurityConfigResponse {
		return v.SecurityConfig
	}).(BareMetalAdminSecurityConfigResponseOutput)
}

// The current state of the bare metal admin cluster.
func (o LookupBareMetalAdminClusterResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBareMetalAdminClusterResult) string { return v.State }).(pulumi.StringOutput)
}

// ResourceStatus representing detailed cluster status.
func (o LookupBareMetalAdminClusterResultOutput) Status() ResourceStatusResponseOutput {
	return o.ApplyT(func(v LookupBareMetalAdminClusterResult) ResourceStatusResponse { return v.Status }).(ResourceStatusResponseOutput)
}

// Storage configuration.
func (o LookupBareMetalAdminClusterResultOutput) Storage() BareMetalAdminStorageConfigResponseOutput {
	return o.ApplyT(func(v LookupBareMetalAdminClusterResult) BareMetalAdminStorageConfigResponse { return v.Storage }).(BareMetalAdminStorageConfigResponseOutput)
}

// The unique identifier of the bare metal admin cluster.
func (o LookupBareMetalAdminClusterResultOutput) Uid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBareMetalAdminClusterResult) string { return v.Uid }).(pulumi.StringOutput)
}

// The time at which this bare metal admin cluster was last updated.
func (o LookupBareMetalAdminClusterResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBareMetalAdminClusterResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

// ValidationCheck representing the result of the preflight check.
func (o LookupBareMetalAdminClusterResultOutput) ValidationCheck() ValidationCheckResponseOutput {
	return o.ApplyT(func(v LookupBareMetalAdminClusterResult) ValidationCheckResponse { return v.ValidationCheck }).(ValidationCheckResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBareMetalAdminClusterResultOutput{})
}
