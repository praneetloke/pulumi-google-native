// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Gets the resource representation for a cluster in a project.
 */
export function getCluster(args: GetClusterArgs, opts?: pulumi.InvokeOptions): Promise<GetClusterResult> {

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts || {});
    return pulumi.runtime.invoke("google-native:dataproc/v1:getCluster", {
        "clusterName": args.clusterName,
        "project": args.project,
        "region": args.region,
    }, opts);
}

export interface GetClusterArgs {
    clusterName: string;
    project?: string;
    region: string;
}

export interface GetClusterResult {
    /**
     * The cluster name, which must be unique within a project. The name must start with a lowercase letter, and can contain up to 51 lowercase letters, numbers, and hyphens. It cannot end with a hyphen. The name of a deleted cluster can be reused.
     */
    readonly clusterName: string;
    /**
     * A cluster UUID (Unique Universal Identifier). Dataproc generates this value when it creates the cluster.
     */
    readonly clusterUuid: string;
    /**
     * Optional. The cluster config for a cluster of Compute Engine Instances. Note that Dataproc may set default values, and values may change when clusters are updated.Exactly one of ClusterConfig or VirtualClusterConfig must be specified.
     */
    readonly config: outputs.dataproc.v1.ClusterConfigResponse;
    /**
     * Optional. The labels to associate with this cluster. Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a cluster.
     */
    readonly labels: {[key: string]: string};
    /**
     * Contains cluster daemon metrics such as HDFS and YARN stats.Beta Feature: This report is available for testing purposes only. It may be changed before final release.
     */
    readonly metrics: outputs.dataproc.v1.ClusterMetricsResponse;
    /**
     * The Google Cloud Platform project ID that the cluster belongs to.
     */
    readonly project: string;
    /**
     * Cluster status.
     */
    readonly status: outputs.dataproc.v1.ClusterStatusResponse;
    /**
     * The previous cluster status.
     */
    readonly statusHistory: outputs.dataproc.v1.ClusterStatusResponse[];
    /**
     * Optional. The virtual cluster config is used when creating a Dataproc cluster that does not directly control the underlying compute resources, for example, when creating a Dataproc-on-GKE cluster (https://cloud.google.com/dataproc/docs/guides/dpgke/dataproc-gke-overview). Dataproc may set default values, and values may change when clusters are updated. Exactly one of config or virtual_cluster_config must be specified.
     */
    readonly virtualClusterConfig: outputs.dataproc.v1.VirtualClusterConfigResponse;
}
/**
 * Gets the resource representation for a cluster in a project.
 */
export function getClusterOutput(args: GetClusterOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetClusterResult> {
    return pulumi.output(args).apply((a: any) => getCluster(a, opts))
}

export interface GetClusterOutputArgs {
    clusterName: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    region: pulumi.Input<string>;
}
