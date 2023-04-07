// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the resource representation for a cluster in a project.
func LookupCluster(ctx *pulumi.Context, args *LookupClusterArgs, opts ...pulumi.InvokeOption) (*LookupClusterResult, error) {
	var rv LookupClusterResult
	err := ctx.Invoke("google-native:dataproc/v1:getCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupClusterArgs struct {
	ClusterName string  `pulumi:"clusterName"`
	Project     *string `pulumi:"project"`
	Region      string  `pulumi:"region"`
}

type LookupClusterResult struct {
	// The cluster name, which must be unique within a project. The name must start with a lowercase letter, and can contain up to 51 lowercase letters, numbers, and hyphens. It cannot end with a hyphen. The name of a deleted cluster can be reused.
	ClusterName string `pulumi:"clusterName"`
	// A cluster UUID (Unique Universal Identifier). Dataproc generates this value when it creates the cluster.
	ClusterUuid string `pulumi:"clusterUuid"`
	// Optional. The cluster config for a cluster of Compute Engine Instances. Note that Dataproc may set default values, and values may change when clusters are updated.Exactly one of ClusterConfig or VirtualClusterConfig must be specified.
	Config ClusterConfigResponse `pulumi:"config"`
	// Optional. The labels to associate with this cluster. Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a cluster.
	Labels map[string]string `pulumi:"labels"`
	// Contains cluster daemon metrics such as HDFS and YARN stats.Beta Feature: This report is available for testing purposes only. It may be changed before final release.
	Metrics ClusterMetricsResponse `pulumi:"metrics"`
	// The Google Cloud Platform project ID that the cluster belongs to.
	Project string `pulumi:"project"`
	// Cluster status.
	Status ClusterStatusResponse `pulumi:"status"`
	// The previous cluster status.
	StatusHistory []ClusterStatusResponse `pulumi:"statusHistory"`
	// Optional. The virtual cluster config is used when creating a Dataproc cluster that does not directly control the underlying compute resources, for example, when creating a Dataproc-on-GKE cluster (https://cloud.google.com/dataproc/docs/guides/dpgke/dataproc-gke-overview). Dataproc may set default values, and values may change when clusters are updated. Exactly one of config or virtual_cluster_config must be specified.
	VirtualClusterConfig VirtualClusterConfigResponse `pulumi:"virtualClusterConfig"`
}

func LookupClusterOutput(ctx *pulumi.Context, args LookupClusterOutputArgs, opts ...pulumi.InvokeOption) LookupClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupClusterResult, error) {
			args := v.(LookupClusterArgs)
			r, err := LookupCluster(ctx, &args, opts...)
			var s LookupClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupClusterResultOutput)
}

type LookupClusterOutputArgs struct {
	ClusterName pulumi.StringInput    `pulumi:"clusterName"`
	Project     pulumi.StringPtrInput `pulumi:"project"`
	Region      pulumi.StringInput    `pulumi:"region"`
}

func (LookupClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterArgs)(nil)).Elem()
}

type LookupClusterResultOutput struct{ *pulumi.OutputState }

func (LookupClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupClusterResult)(nil)).Elem()
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutput() LookupClusterResultOutput {
	return o
}

func (o LookupClusterResultOutput) ToLookupClusterResultOutputWithContext(ctx context.Context) LookupClusterResultOutput {
	return o
}

// The cluster name, which must be unique within a project. The name must start with a lowercase letter, and can contain up to 51 lowercase letters, numbers, and hyphens. It cannot end with a hyphen. The name of a deleted cluster can be reused.
func (o LookupClusterResultOutput) ClusterName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ClusterName }).(pulumi.StringOutput)
}

// A cluster UUID (Unique Universal Identifier). Dataproc generates this value when it creates the cluster.
func (o LookupClusterResultOutput) ClusterUuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.ClusterUuid }).(pulumi.StringOutput)
}

// Optional. The cluster config for a cluster of Compute Engine Instances. Note that Dataproc may set default values, and values may change when clusters are updated.Exactly one of ClusterConfig or VirtualClusterConfig must be specified.
func (o LookupClusterResultOutput) Config() ClusterConfigResponseOutput {
	return o.ApplyT(func(v LookupClusterResult) ClusterConfigResponse { return v.Config }).(ClusterConfigResponseOutput)
}

// Optional. The labels to associate with this cluster. Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a cluster.
func (o LookupClusterResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupClusterResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Contains cluster daemon metrics such as HDFS and YARN stats.Beta Feature: This report is available for testing purposes only. It may be changed before final release.
func (o LookupClusterResultOutput) Metrics() ClusterMetricsResponseOutput {
	return o.ApplyT(func(v LookupClusterResult) ClusterMetricsResponse { return v.Metrics }).(ClusterMetricsResponseOutput)
}

// The Google Cloud Platform project ID that the cluster belongs to.
func (o LookupClusterResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupClusterResult) string { return v.Project }).(pulumi.StringOutput)
}

// Cluster status.
func (o LookupClusterResultOutput) Status() ClusterStatusResponseOutput {
	return o.ApplyT(func(v LookupClusterResult) ClusterStatusResponse { return v.Status }).(ClusterStatusResponseOutput)
}

// The previous cluster status.
func (o LookupClusterResultOutput) StatusHistory() ClusterStatusResponseArrayOutput {
	return o.ApplyT(func(v LookupClusterResult) []ClusterStatusResponse { return v.StatusHistory }).(ClusterStatusResponseArrayOutput)
}

// Optional. The virtual cluster config is used when creating a Dataproc cluster that does not directly control the underlying compute resources, for example, when creating a Dataproc-on-GKE cluster (https://cloud.google.com/dataproc/docs/guides/dpgke/dataproc-gke-overview). Dataproc may set default values, and values may change when clusters are updated. Exactly one of config or virtual_cluster_config must be specified.
func (o LookupClusterResultOutput) VirtualClusterConfig() VirtualClusterConfigResponseOutput {
	return o.ApplyT(func(v LookupClusterResult) VirtualClusterConfigResponse { return v.VirtualClusterConfig }).(VirtualClusterConfigResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupClusterResultOutput{})
}
