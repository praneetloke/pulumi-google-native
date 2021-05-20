// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../../types";
import * as utilities from "../../utilities";

/**
 * Creates a cluster in a project. The returned Operation.metadata will be ClusterOperationMetadata (https://cloud.google.com/dataproc/docs/reference/rpc/google.cloud.dataproc.v1#clusteroperationmetadata).
 */
export class RegionCluster extends pulumi.CustomResource {
    /**
     * Get an existing RegionCluster resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RegionCluster {
        return new RegionCluster(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:dataproc/v1:RegionCluster';

    /**
     * Returns true if the given object is an instance of RegionCluster.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegionCluster {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegionCluster.__pulumiType;
    }

    /**
     * Required. The cluster name. Cluster names within a project must be unique. Names of deleted clusters can be reused.
     */
    public readonly clusterName!: pulumi.Output<string>;
    /**
     * A cluster UUID (Unique Universal Identifier). Dataproc generates this value when it creates the cluster.
     */
    public /*out*/ readonly clusterUuid!: pulumi.Output<string>;
    /**
     * Required. The cluster config. Note that Dataproc may set default values, and values may change when clusters are updated.
     */
    public readonly config!: pulumi.Output<outputs.dataproc.v1.ClusterConfigResponse>;
    /**
     * Optional. The labels to associate with this cluster. Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a cluster.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Contains cluster daemon metrics such as HDFS and YARN stats.Beta Feature: This report is available for testing purposes only. It may be changed before final release.
     */
    public /*out*/ readonly metrics!: pulumi.Output<outputs.dataproc.v1.ClusterMetricsResponse>;
    /**
     * Required. The Google Cloud Platform project ID that the cluster belongs to.
     */
    public readonly project!: pulumi.Output<string>;
    /**
     * Cluster status.
     */
    public /*out*/ readonly status!: pulumi.Output<outputs.dataproc.v1.ClusterStatusResponse>;
    /**
     * The previous cluster status.
     */
    public /*out*/ readonly statusHistory!: pulumi.Output<outputs.dataproc.v1.ClusterStatusResponse[]>;

    /**
     * Create a RegionCluster resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegionClusterArgs, opts?: pulumi.CustomResourceOptions) {
        let inputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.clusterName === undefined) && !opts.urn) {
                throw new Error("Missing required property 'clusterName'");
            }
            if ((!args || args.project === undefined) && !opts.urn) {
                throw new Error("Missing required property 'project'");
            }
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            inputs["clusterName"] = args ? args.clusterName : undefined;
            inputs["config"] = args ? args.config : undefined;
            inputs["labels"] = args ? args.labels : undefined;
            inputs["project"] = args ? args.project : undefined;
            inputs["region"] = args ? args.region : undefined;
            inputs["requestId"] = args ? args.requestId : undefined;
            inputs["clusterUuid"] = undefined /*out*/;
            inputs["metrics"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["statusHistory"] = undefined /*out*/;
        } else {
            inputs["clusterName"] = undefined /*out*/;
            inputs["clusterUuid"] = undefined /*out*/;
            inputs["config"] = undefined /*out*/;
            inputs["labels"] = undefined /*out*/;
            inputs["metrics"] = undefined /*out*/;
            inputs["project"] = undefined /*out*/;
            inputs["status"] = undefined /*out*/;
            inputs["statusHistory"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(RegionCluster.__pulumiType, name, inputs, opts);
    }
}

/**
 * The set of arguments for constructing a RegionCluster resource.
 */
export interface RegionClusterArgs {
    /**
     * Required. The cluster name. Cluster names within a project must be unique. Names of deleted clusters can be reused.
     */
    readonly clusterName: pulumi.Input<string>;
    /**
     * Required. The cluster config. Note that Dataproc may set default values, and values may change when clusters are updated.
     */
    readonly config?: pulumi.Input<inputs.dataproc.v1.ClusterConfigArgs>;
    /**
     * Optional. The labels to associate with this cluster. Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a cluster.
     */
    readonly labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    /**
     * Required. The Google Cloud Platform project ID that the cluster belongs to.
     */
    readonly project: pulumi.Input<string>;
    readonly region: pulumi.Input<string>;
    readonly requestId?: pulumi.Input<string>;
}
