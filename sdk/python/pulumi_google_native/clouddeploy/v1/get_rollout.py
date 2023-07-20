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
    'GetRolloutResult',
    'AwaitableGetRolloutResult',
    'get_rollout',
    'get_rollout_output',
]

@pulumi.output_type
class GetRolloutResult:
    def __init__(__self__, annotations=None, approval_state=None, approve_time=None, controller_rollout=None, create_time=None, deploy_end_time=None, deploy_failure_cause=None, deploy_start_time=None, deploying_build=None, description=None, enqueue_time=None, etag=None, failure_reason=None, labels=None, metadata=None, name=None, phases=None, state=None, target_id=None, uid=None):
        if annotations and not isinstance(annotations, dict):
            raise TypeError("Expected argument 'annotations' to be a dict")
        pulumi.set(__self__, "annotations", annotations)
        if approval_state and not isinstance(approval_state, str):
            raise TypeError("Expected argument 'approval_state' to be a str")
        pulumi.set(__self__, "approval_state", approval_state)
        if approve_time and not isinstance(approve_time, str):
            raise TypeError("Expected argument 'approve_time' to be a str")
        pulumi.set(__self__, "approve_time", approve_time)
        if controller_rollout and not isinstance(controller_rollout, str):
            raise TypeError("Expected argument 'controller_rollout' to be a str")
        pulumi.set(__self__, "controller_rollout", controller_rollout)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if deploy_end_time and not isinstance(deploy_end_time, str):
            raise TypeError("Expected argument 'deploy_end_time' to be a str")
        pulumi.set(__self__, "deploy_end_time", deploy_end_time)
        if deploy_failure_cause and not isinstance(deploy_failure_cause, str):
            raise TypeError("Expected argument 'deploy_failure_cause' to be a str")
        pulumi.set(__self__, "deploy_failure_cause", deploy_failure_cause)
        if deploy_start_time and not isinstance(deploy_start_time, str):
            raise TypeError("Expected argument 'deploy_start_time' to be a str")
        pulumi.set(__self__, "deploy_start_time", deploy_start_time)
        if deploying_build and not isinstance(deploying_build, str):
            raise TypeError("Expected argument 'deploying_build' to be a str")
        pulumi.set(__self__, "deploying_build", deploying_build)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if enqueue_time and not isinstance(enqueue_time, str):
            raise TypeError("Expected argument 'enqueue_time' to be a str")
        pulumi.set(__self__, "enqueue_time", enqueue_time)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if failure_reason and not isinstance(failure_reason, str):
            raise TypeError("Expected argument 'failure_reason' to be a str")
        pulumi.set(__self__, "failure_reason", failure_reason)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if metadata and not isinstance(metadata, dict):
            raise TypeError("Expected argument 'metadata' to be a dict")
        pulumi.set(__self__, "metadata", metadata)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if phases and not isinstance(phases, list):
            raise TypeError("Expected argument 'phases' to be a list")
        pulumi.set(__self__, "phases", phases)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if target_id and not isinstance(target_id, str):
            raise TypeError("Expected argument 'target_id' to be a str")
        pulumi.set(__self__, "target_id", target_id)
        if uid and not isinstance(uid, str):
            raise TypeError("Expected argument 'uid' to be a str")
        pulumi.set(__self__, "uid", uid)

    @property
    @pulumi.getter
    def annotations(self) -> Mapping[str, str]:
        """
        User annotations. These attributes can only be set and used by the user, and not by Google Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
        """
        return pulumi.get(self, "annotations")

    @property
    @pulumi.getter(name="approvalState")
    def approval_state(self) -> str:
        """
        Approval state of the `Rollout`.
        """
        return pulumi.get(self, "approval_state")

    @property
    @pulumi.getter(name="approveTime")
    def approve_time(self) -> str:
        """
        Time at which the `Rollout` was approved.
        """
        return pulumi.get(self, "approve_time")

    @property
    @pulumi.getter(name="controllerRollout")
    def controller_rollout(self) -> str:
        """
        Name of the `ControllerRollout`. Format is projects/{project}/ locations/{location}/deliveryPipelines/{deliveryPipeline}/ releases/{release}/rollouts/a-z{0,62}.
        """
        return pulumi.get(self, "controller_rollout")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Time at which the `Rollout` was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="deployEndTime")
    def deploy_end_time(self) -> str:
        """
        Time at which the `Rollout` finished deploying.
        """
        return pulumi.get(self, "deploy_end_time")

    @property
    @pulumi.getter(name="deployFailureCause")
    def deploy_failure_cause(self) -> str:
        """
        The reason this rollout failed. This will always be unspecified while the rollout is in progress.
        """
        return pulumi.get(self, "deploy_failure_cause")

    @property
    @pulumi.getter(name="deployStartTime")
    def deploy_start_time(self) -> str:
        """
        Time at which the `Rollout` started deploying.
        """
        return pulumi.get(self, "deploy_start_time")

    @property
    @pulumi.getter(name="deployingBuild")
    def deploying_build(self) -> str:
        """
        The resource name of the Cloud Build `Build` object that is used to deploy the Rollout. Format is `projects/{project}/locations/{location}/builds/{build}`.
        """
        return pulumi.get(self, "deploying_build")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Description of the `Rollout` for user purposes. Max length is 255 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="enqueueTime")
    def enqueue_time(self) -> str:
        """
        Time at which the `Rollout` was enqueued.
        """
        return pulumi.get(self, "enqueue_time")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="failureReason")
    def failure_reason(self) -> str:
        """
        Additional information about the rollout failure, if available.
        """
        return pulumi.get(self, "failure_reason")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Labels are attributes that can be set and used by both the user and by Google Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be <= 128 bytes.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def metadata(self) -> 'outputs.MetadataResponse':
        """
        Metadata contains information about the rollout.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Optional. Name of the `Rollout`. Format is projects/{project}/ locations/{location}/deliveryPipelines/{deliveryPipeline}/ releases/{release}/rollouts/a-z{0,62}.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def phases(self) -> Sequence['outputs.PhaseResponse']:
        """
        The phases that represent the workflows of this `Rollout`.
        """
        return pulumi.get(self, "phases")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        Current state of the `Rollout`.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="targetId")
    def target_id(self) -> str:
        """
        The ID of Target to which this `Rollout` is deploying.
        """
        return pulumi.get(self, "target_id")

    @property
    @pulumi.getter
    def uid(self) -> str:
        """
        Unique identifier of the `Rollout`.
        """
        return pulumi.get(self, "uid")


class AwaitableGetRolloutResult(GetRolloutResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRolloutResult(
            annotations=self.annotations,
            approval_state=self.approval_state,
            approve_time=self.approve_time,
            controller_rollout=self.controller_rollout,
            create_time=self.create_time,
            deploy_end_time=self.deploy_end_time,
            deploy_failure_cause=self.deploy_failure_cause,
            deploy_start_time=self.deploy_start_time,
            deploying_build=self.deploying_build,
            description=self.description,
            enqueue_time=self.enqueue_time,
            etag=self.etag,
            failure_reason=self.failure_reason,
            labels=self.labels,
            metadata=self.metadata,
            name=self.name,
            phases=self.phases,
            state=self.state,
            target_id=self.target_id,
            uid=self.uid)


def get_rollout(delivery_pipeline_id: Optional[str] = None,
                location: Optional[str] = None,
                project: Optional[str] = None,
                release_id: Optional[str] = None,
                rollout_id: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRolloutResult:
    """
    Gets details of a single Rollout.
    """
    __args__ = dict()
    __args__['deliveryPipelineId'] = delivery_pipeline_id
    __args__['location'] = location
    __args__['project'] = project
    __args__['releaseId'] = release_id
    __args__['rolloutId'] = rollout_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:clouddeploy/v1:getRollout', __args__, opts=opts, typ=GetRolloutResult).value

    return AwaitableGetRolloutResult(
        annotations=pulumi.get(__ret__, 'annotations'),
        approval_state=pulumi.get(__ret__, 'approval_state'),
        approve_time=pulumi.get(__ret__, 'approve_time'),
        controller_rollout=pulumi.get(__ret__, 'controller_rollout'),
        create_time=pulumi.get(__ret__, 'create_time'),
        deploy_end_time=pulumi.get(__ret__, 'deploy_end_time'),
        deploy_failure_cause=pulumi.get(__ret__, 'deploy_failure_cause'),
        deploy_start_time=pulumi.get(__ret__, 'deploy_start_time'),
        deploying_build=pulumi.get(__ret__, 'deploying_build'),
        description=pulumi.get(__ret__, 'description'),
        enqueue_time=pulumi.get(__ret__, 'enqueue_time'),
        etag=pulumi.get(__ret__, 'etag'),
        failure_reason=pulumi.get(__ret__, 'failure_reason'),
        labels=pulumi.get(__ret__, 'labels'),
        metadata=pulumi.get(__ret__, 'metadata'),
        name=pulumi.get(__ret__, 'name'),
        phases=pulumi.get(__ret__, 'phases'),
        state=pulumi.get(__ret__, 'state'),
        target_id=pulumi.get(__ret__, 'target_id'),
        uid=pulumi.get(__ret__, 'uid'))


@_utilities.lift_output_func(get_rollout)
def get_rollout_output(delivery_pipeline_id: Optional[pulumi.Input[str]] = None,
                       location: Optional[pulumi.Input[str]] = None,
                       project: Optional[pulumi.Input[Optional[str]]] = None,
                       release_id: Optional[pulumi.Input[str]] = None,
                       rollout_id: Optional[pulumi.Input[str]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRolloutResult]:
    """
    Gets details of a single Rollout.
    """
    ...
