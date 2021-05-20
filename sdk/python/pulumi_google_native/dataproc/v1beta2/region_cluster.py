# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs
from ._inputs import *

__all__ = ['RegionClusterArgs', 'RegionCluster']

@pulumi.input_type
class RegionClusterArgs:
    def __init__(__self__, *,
                 cluster_name: pulumi.Input[str],
                 project: pulumi.Input[str],
                 region: pulumi.Input[str],
                 config: Optional[pulumi.Input['ClusterConfigArgs']] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 request_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RegionCluster resource.
        :param pulumi.Input[str] cluster_name: Required. The cluster name. Cluster names within a project must be unique. Names of deleted clusters can be reused.
        :param pulumi.Input[str] project: Required. The Google Cloud Platform project ID that the cluster belongs to.
        :param pulumi.Input['ClusterConfigArgs'] config: Required. The cluster config. Note that Dataproc may set default values, and values may change when clusters are updated.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Optional. The labels to associate with this cluster. Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a cluster.
        """
        pulumi.set(__self__, "cluster_name", cluster_name)
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "region", region)
        if config is not None:
            pulumi.set(__self__, "config", config)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if request_id is not None:
            pulumi.set(__self__, "request_id", request_id)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Input[str]:
        """
        Required. The cluster name. Cluster names within a project must be unique. Names of deleted clusters can be reused.
        """
        return pulumi.get(self, "cluster_name")

    @cluster_name.setter
    def cluster_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster_name", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        """
        Required. The Google Cloud Platform project ID that the cluster belongs to.
        """
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def region(self) -> pulumi.Input[str]:
        return pulumi.get(self, "region")

    @region.setter
    def region(self, value: pulumi.Input[str]):
        pulumi.set(self, "region", value)

    @property
    @pulumi.getter
    def config(self) -> Optional[pulumi.Input['ClusterConfigArgs']]:
        """
        Required. The cluster config. Note that Dataproc may set default values, and values may change when clusters are updated.
        """
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: Optional[pulumi.Input['ClusterConfigArgs']]):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Optional. The labels to associate with this cluster. Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a cluster.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "request_id")

    @request_id.setter
    def request_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_id", value)


class RegionCluster(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 config: Optional[pulumi.Input[pulumi.InputType['ClusterConfigArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a cluster in a project. The returned Operation.metadata will be ClusterOperationMetadata (https://cloud.google.com/dataproc/docs/reference/rpc/google.cloud.dataproc.v1beta2#clusteroperationmetadata).

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] cluster_name: Required. The cluster name. Cluster names within a project must be unique. Names of deleted clusters can be reused.
        :param pulumi.Input[pulumi.InputType['ClusterConfigArgs']] config: Required. The cluster config. Note that Dataproc may set default values, and values may change when clusters are updated.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Optional. The labels to associate with this cluster. Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a cluster.
        :param pulumi.Input[str] project: Required. The Google Cloud Platform project ID that the cluster belongs to.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RegionClusterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a cluster in a project. The returned Operation.metadata will be ClusterOperationMetadata (https://cloud.google.com/dataproc/docs/reference/rpc/google.cloud.dataproc.v1beta2#clusteroperationmetadata).

        :param str resource_name: The name of the resource.
        :param RegionClusterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RegionClusterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cluster_name: Optional[pulumi.Input[str]] = None,
                 config: Optional[pulumi.Input[pulumi.InputType['ClusterConfigArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 region: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RegionClusterArgs.__new__(RegionClusterArgs)

            if cluster_name is None and not opts.urn:
                raise TypeError("Missing required property 'cluster_name'")
            __props__.__dict__["cluster_name"] = cluster_name
            __props__.__dict__["config"] = config
            __props__.__dict__["labels"] = labels
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if region is None and not opts.urn:
                raise TypeError("Missing required property 'region'")
            __props__.__dict__["region"] = region
            __props__.__dict__["request_id"] = request_id
            __props__.__dict__["cluster_uuid"] = None
            __props__.__dict__["metrics"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["status_history"] = None
        super(RegionCluster, __self__).__init__(
            'google-native:dataproc/v1beta2:RegionCluster',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'RegionCluster':
        """
        Get an existing RegionCluster resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = RegionClusterArgs.__new__(RegionClusterArgs)

        __props__.__dict__["cluster_name"] = None
        __props__.__dict__["cluster_uuid"] = None
        __props__.__dict__["config"] = None
        __props__.__dict__["labels"] = None
        __props__.__dict__["metrics"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["status_history"] = None
        return RegionCluster(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="clusterName")
    def cluster_name(self) -> pulumi.Output[str]:
        """
        Required. The cluster name. Cluster names within a project must be unique. Names of deleted clusters can be reused.
        """
        return pulumi.get(self, "cluster_name")

    @property
    @pulumi.getter(name="clusterUuid")
    def cluster_uuid(self) -> pulumi.Output[str]:
        """
        A cluster UUID (Unique Universal Identifier). Dataproc generates this value when it creates the cluster.
        """
        return pulumi.get(self, "cluster_uuid")

    @property
    @pulumi.getter
    def config(self) -> pulumi.Output['outputs.ClusterConfigResponse']:
        """
        Required. The cluster config. Note that Dataproc may set default values, and values may change when clusters are updated.
        """
        return pulumi.get(self, "config")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Optional. The labels to associate with this cluster. Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a cluster.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def metrics(self) -> pulumi.Output['outputs.ClusterMetricsResponse']:
        """
        Contains cluster daemon metrics such as HDFS and YARN stats.Beta Feature: This report is available for testing purposes only. It may be changed before final release.
        """
        return pulumi.get(self, "metrics")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        """
        Required. The Google Cloud Platform project ID that the cluster belongs to.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output['outputs.ClusterStatusResponse']:
        """
        Cluster status.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="statusHistory")
    def status_history(self) -> pulumi.Output[Sequence['outputs.ClusterStatusResponse']]:
        """
        The previous cluster status.
        """
        return pulumi.get(self, "status_history")

