# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = [
    'GetServiceResult',
    'AwaitableGetServiceResult',
    'get_service',
    'get_service_output',
]

@pulumi.output_type
class GetServiceResult:
    def __init__(__self__, annotations=None, binary_authorization=None, client=None, client_version=None, conditions=None, create_time=None, creator=None, delete_time=None, description=None, etag=None, expire_time=None, generation=None, ingress=None, labels=None, last_modifier=None, latest_created_revision=None, latest_ready_revision=None, launch_stage=None, name=None, observed_generation=None, reconciling=None, satisfies_pzs=None, template=None, terminal_condition=None, traffic=None, traffic_statuses=None, uid=None, update_time=None, uri=None):
        if annotations and not isinstance(annotations, dict):
            raise TypeError("Expected argument 'annotations' to be a dict")
        pulumi.set(__self__, "annotations", annotations)
        if binary_authorization and not isinstance(binary_authorization, dict):
            raise TypeError("Expected argument 'binary_authorization' to be a dict")
        pulumi.set(__self__, "binary_authorization", binary_authorization)
        if client and not isinstance(client, str):
            raise TypeError("Expected argument 'client' to be a str")
        pulumi.set(__self__, "client", client)
        if client_version and not isinstance(client_version, str):
            raise TypeError("Expected argument 'client_version' to be a str")
        pulumi.set(__self__, "client_version", client_version)
        if conditions and not isinstance(conditions, list):
            raise TypeError("Expected argument 'conditions' to be a list")
        pulumi.set(__self__, "conditions", conditions)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if creator and not isinstance(creator, str):
            raise TypeError("Expected argument 'creator' to be a str")
        pulumi.set(__self__, "creator", creator)
        if delete_time and not isinstance(delete_time, str):
            raise TypeError("Expected argument 'delete_time' to be a str")
        pulumi.set(__self__, "delete_time", delete_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if expire_time and not isinstance(expire_time, str):
            raise TypeError("Expected argument 'expire_time' to be a str")
        pulumi.set(__self__, "expire_time", expire_time)
        if generation and not isinstance(generation, str):
            raise TypeError("Expected argument 'generation' to be a str")
        pulumi.set(__self__, "generation", generation)
        if ingress and not isinstance(ingress, str):
            raise TypeError("Expected argument 'ingress' to be a str")
        pulumi.set(__self__, "ingress", ingress)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if last_modifier and not isinstance(last_modifier, str):
            raise TypeError("Expected argument 'last_modifier' to be a str")
        pulumi.set(__self__, "last_modifier", last_modifier)
        if latest_created_revision and not isinstance(latest_created_revision, str):
            raise TypeError("Expected argument 'latest_created_revision' to be a str")
        pulumi.set(__self__, "latest_created_revision", latest_created_revision)
        if latest_ready_revision and not isinstance(latest_ready_revision, str):
            raise TypeError("Expected argument 'latest_ready_revision' to be a str")
        pulumi.set(__self__, "latest_ready_revision", latest_ready_revision)
        if launch_stage and not isinstance(launch_stage, str):
            raise TypeError("Expected argument 'launch_stage' to be a str")
        pulumi.set(__self__, "launch_stage", launch_stage)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if observed_generation and not isinstance(observed_generation, str):
            raise TypeError("Expected argument 'observed_generation' to be a str")
        pulumi.set(__self__, "observed_generation", observed_generation)
        if reconciling and not isinstance(reconciling, bool):
            raise TypeError("Expected argument 'reconciling' to be a bool")
        pulumi.set(__self__, "reconciling", reconciling)
        if satisfies_pzs and not isinstance(satisfies_pzs, bool):
            raise TypeError("Expected argument 'satisfies_pzs' to be a bool")
        pulumi.set(__self__, "satisfies_pzs", satisfies_pzs)
        if template and not isinstance(template, dict):
            raise TypeError("Expected argument 'template' to be a dict")
        pulumi.set(__self__, "template", template)
        if terminal_condition and not isinstance(terminal_condition, dict):
            raise TypeError("Expected argument 'terminal_condition' to be a dict")
        pulumi.set(__self__, "terminal_condition", terminal_condition)
        if traffic and not isinstance(traffic, list):
            raise TypeError("Expected argument 'traffic' to be a list")
        pulumi.set(__self__, "traffic", traffic)
        if traffic_statuses and not isinstance(traffic_statuses, list):
            raise TypeError("Expected argument 'traffic_statuses' to be a list")
        pulumi.set(__self__, "traffic_statuses", traffic_statuses)
        if uid and not isinstance(uid, str):
            raise TypeError("Expected argument 'uid' to be a str")
        pulumi.set(__self__, "uid", uid)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)
        if uri and not isinstance(uri, str):
            raise TypeError("Expected argument 'uri' to be a str")
        pulumi.set(__self__, "uri", uri)

    @property
    @pulumi.getter
    def annotations(self) -> Mapping[str, str]:
        """
        Unstructured key value map that may be set by external tools to store and arbitrary metadata. They are not queryable and should be preserved when modifying objects. Cloud Run API v2 does not support annotations with `run.googleapis.com`, `cloud.googleapis.com`, `serving.knative.dev`, or `autoscaling.knative.dev` namespaces, and they will be rejected. All system annotations in v1 now have a corresponding field in v2 Service. This field follows Kubernetes annotations' namespacing, limits, and rules. More info: https://kubernetes.io/docs/user-guide/annotations
        """
        return pulumi.get(self, "annotations")

    @property
    @pulumi.getter(name="binaryAuthorization")
    def binary_authorization(self) -> 'outputs.GoogleCloudRunV2BinaryAuthorizationResponse':
        """
        Settings for the Binary Authorization feature.
        """
        return pulumi.get(self, "binary_authorization")

    @property
    @pulumi.getter
    def client(self) -> str:
        """
        Arbitrary identifier for the API client.
        """
        return pulumi.get(self, "client")

    @property
    @pulumi.getter(name="clientVersion")
    def client_version(self) -> str:
        """
        Arbitrary version identifier for the API client.
        """
        return pulumi.get(self, "client_version")

    @property
    @pulumi.getter
    def conditions(self) -> Sequence['outputs.GoogleCloudRunV2ConditionResponse']:
        """
        The Conditions of all other associated sub-resources. They contain additional diagnostics information in case the Service does not reach its Serving state. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
        """
        return pulumi.get(self, "conditions")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The creation time.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def creator(self) -> str:
        """
        Email address of the authenticated creator.
        """
        return pulumi.get(self, "creator")

    @property
    @pulumi.getter(name="deleteTime")
    def delete_time(self) -> str:
        """
        The deletion time.
        """
        return pulumi.get(self, "delete_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        User-provided description of the Service. This field currently has a 512-character limit.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        A system-generated fingerprint for this version of the resource. May be used to detect modification conflict during updates.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> str:
        """
        For a deleted resource, the time after which it will be permamently deleted.
        """
        return pulumi.get(self, "expire_time")

    @property
    @pulumi.getter
    def generation(self) -> str:
        """
        A number that monotonically increases every time the user modifies the desired state. Please note that unlike v1, this is an int64 value. As with most Google APIs, its JSON representation will be a `string` instead of an `integer`.
        """
        return pulumi.get(self, "generation")

    @property
    @pulumi.getter
    def ingress(self) -> str:
        """
        Provides the ingress settings for this Service. On output, returns the currently observed ingress settings, or INGRESS_TRAFFIC_UNSPECIFIED if no revision is active.
        """
        return pulumi.get(self, "ingress")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Map of string keys and values that can be used to organize and categorize objects. User-provided labels are shared with Google's billing system, so they can be used to filter, or break down billing charges by team, component, environment, state, etc. For more information, visit https://cloud.google.com/resource-manager/docs/creating-managing-labels or https://cloud.google.com/run/docs/configuring/labels Cloud Run API v2 does not support labels with `run.googleapis.com`, `cloud.googleapis.com`, `serving.knative.dev`, or `autoscaling.knative.dev` namespaces, and they will be rejected. All system labels in v1 now have a corresponding field in v2 Service.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="lastModifier")
    def last_modifier(self) -> str:
        """
        Email address of the last authenticated modifier.
        """
        return pulumi.get(self, "last_modifier")

    @property
    @pulumi.getter(name="latestCreatedRevision")
    def latest_created_revision(self) -> str:
        """
        Name of the last created revision. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
        """
        return pulumi.get(self, "latest_created_revision")

    @property
    @pulumi.getter(name="latestReadyRevision")
    def latest_ready_revision(self) -> str:
        """
        Name of the latest revision that is serving traffic. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
        """
        return pulumi.get(self, "latest_ready_revision")

    @property
    @pulumi.getter(name="launchStage")
    def launch_stage(self) -> str:
        """
        The launch stage as defined by [Google Cloud Platform Launch Stages](https://cloud.google.com/terms/launch-stages). Cloud Run supports `ALPHA`, `BETA`, and `GA`. If no value is specified, GA is assumed. Set the launch stage to a preview stage on input to allow use of preview features in that stage. On read (or output), describes whether the resource uses preview features. For example, if ALPHA is provided as input, but only BETA and GA-level features are used, this field will be BETA on output.
        """
        return pulumi.get(self, "launch_stage")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The fully qualified name of this Service. In CreateServiceRequest, this field is ignored, and instead composed from CreateServiceRequest.parent and CreateServiceRequest.service_id. Format: projects/{project}/locations/{location}/services/{service_id}
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="observedGeneration")
    def observed_generation(self) -> str:
        """
        The generation of this Service currently serving traffic. See comments in `reconciling` for additional information on reconciliation process in Cloud Run. Please note that unlike v1, this is an int64 value. As with most Google APIs, its JSON representation will be a `string` instead of an `integer`.
        """
        return pulumi.get(self, "observed_generation")

    @property
    @pulumi.getter
    def reconciling(self) -> bool:
        """
        Returns true if the Service is currently being acted upon by the system to bring it into the desired state. When a new Service is created, or an existing one is updated, Cloud Run will asynchronously perform all necessary steps to bring the Service to the desired serving state. This process is called reconciliation. While reconciliation is in process, `observed_generation`, `latest_ready_revison`, `traffic_statuses`, and `uri` will have transient values that might mismatch the intended state: Once reconciliation is over (and this field is false), there are two possible outcomes: reconciliation succeeded and the serving state matches the Service, or there was an error, and reconciliation failed. This state can be found in `terminal_condition.state`. If reconciliation succeeded, the following fields will match: `traffic` and `traffic_statuses`, `observed_generation` and `generation`, `latest_ready_revision` and `latest_created_revision`. If reconciliation failed, `traffic_statuses`, `observed_generation`, and `latest_ready_revision` will have the state of the last serving revision, or empty for newly created Services. Additional information on the failure can be found in `terminal_condition` and `conditions`.
        """
        return pulumi.get(self, "reconciling")

    @property
    @pulumi.getter(name="satisfiesPzs")
    def satisfies_pzs(self) -> bool:
        """
        Reserved for future use.
        """
        return pulumi.get(self, "satisfies_pzs")

    @property
    @pulumi.getter
    def template(self) -> 'outputs.GoogleCloudRunV2RevisionTemplateResponse':
        """
        The template used to create revisions for this Service.
        """
        return pulumi.get(self, "template")

    @property
    @pulumi.getter(name="terminalCondition")
    def terminal_condition(self) -> 'outputs.GoogleCloudRunV2ConditionResponse':
        """
        The Condition of this Service, containing its readiness status, and detailed error information in case it did not reach a serving state. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
        """
        return pulumi.get(self, "terminal_condition")

    @property
    @pulumi.getter
    def traffic(self) -> Sequence['outputs.GoogleCloudRunV2TrafficTargetResponse']:
        """
        Specifies how to distribute traffic over a collection of Revisions belonging to the Service. If traffic is empty or not provided, defaults to 100% traffic to the latest `Ready` Revision.
        """
        return pulumi.get(self, "traffic")

    @property
    @pulumi.getter(name="trafficStatuses")
    def traffic_statuses(self) -> Sequence['outputs.GoogleCloudRunV2TrafficTargetStatusResponse']:
        """
        Detailed status information for corresponding traffic targets. See comments in `reconciling` for additional information on reconciliation process in Cloud Run.
        """
        return pulumi.get(self, "traffic_statuses")

    @property
    @pulumi.getter
    def uid(self) -> str:
        """
        Server assigned unique identifier for the trigger. The value is a UUID4 string and guaranteed to remain unchanged until the resource is deleted.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The last-modified time.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter
    def uri(self) -> str:
        """
        The main URI in which this Service is serving traffic.
        """
        return pulumi.get(self, "uri")


class AwaitableGetServiceResult(GetServiceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetServiceResult(
            annotations=self.annotations,
            binary_authorization=self.binary_authorization,
            client=self.client,
            client_version=self.client_version,
            conditions=self.conditions,
            create_time=self.create_time,
            creator=self.creator,
            delete_time=self.delete_time,
            description=self.description,
            etag=self.etag,
            expire_time=self.expire_time,
            generation=self.generation,
            ingress=self.ingress,
            labels=self.labels,
            last_modifier=self.last_modifier,
            latest_created_revision=self.latest_created_revision,
            latest_ready_revision=self.latest_ready_revision,
            launch_stage=self.launch_stage,
            name=self.name,
            observed_generation=self.observed_generation,
            reconciling=self.reconciling,
            satisfies_pzs=self.satisfies_pzs,
            template=self.template,
            terminal_condition=self.terminal_condition,
            traffic=self.traffic,
            traffic_statuses=self.traffic_statuses,
            uid=self.uid,
            update_time=self.update_time,
            uri=self.uri)


def get_service(location: Optional[str] = None,
                project: Optional[str] = None,
                service_id: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetServiceResult:
    """
    Gets information about a Service.
    """
    __args__ = dict()
    __args__['location'] = location
    __args__['project'] = project
    __args__['serviceId'] = service_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:run/v2:getService', __args__, opts=opts, typ=GetServiceResult).value

    return AwaitableGetServiceResult(
        annotations=__ret__.annotations,
        binary_authorization=__ret__.binary_authorization,
        client=__ret__.client,
        client_version=__ret__.client_version,
        conditions=__ret__.conditions,
        create_time=__ret__.create_time,
        creator=__ret__.creator,
        delete_time=__ret__.delete_time,
        description=__ret__.description,
        etag=__ret__.etag,
        expire_time=__ret__.expire_time,
        generation=__ret__.generation,
        ingress=__ret__.ingress,
        labels=__ret__.labels,
        last_modifier=__ret__.last_modifier,
        latest_created_revision=__ret__.latest_created_revision,
        latest_ready_revision=__ret__.latest_ready_revision,
        launch_stage=__ret__.launch_stage,
        name=__ret__.name,
        observed_generation=__ret__.observed_generation,
        reconciling=__ret__.reconciling,
        satisfies_pzs=__ret__.satisfies_pzs,
        template=__ret__.template,
        terminal_condition=__ret__.terminal_condition,
        traffic=__ret__.traffic,
        traffic_statuses=__ret__.traffic_statuses,
        uid=__ret__.uid,
        update_time=__ret__.update_time,
        uri=__ret__.uri)


@_utilities.lift_output_func(get_service)
def get_service_output(location: Optional[pulumi.Input[str]] = None,
                       project: Optional[pulumi.Input[Optional[str]]] = None,
                       service_id: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetServiceResult]:
    """
    Gets information about a Service.
    """
    ...
