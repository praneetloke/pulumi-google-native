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
    'GetFunctionResult',
    'AwaitableGetFunctionResult',
    'get_function',
    'get_function_output',
]

@pulumi.output_type
class GetFunctionResult:
    def __init__(__self__, build_config=None, description=None, environment=None, event_trigger=None, kms_key_name=None, labels=None, name=None, service_config=None, state=None, state_messages=None, update_time=None):
        if build_config and not isinstance(build_config, dict):
            raise TypeError("Expected argument 'build_config' to be a dict")
        pulumi.set(__self__, "build_config", build_config)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if environment and not isinstance(environment, str):
            raise TypeError("Expected argument 'environment' to be a str")
        pulumi.set(__self__, "environment", environment)
        if event_trigger and not isinstance(event_trigger, dict):
            raise TypeError("Expected argument 'event_trigger' to be a dict")
        pulumi.set(__self__, "event_trigger", event_trigger)
        if kms_key_name and not isinstance(kms_key_name, str):
            raise TypeError("Expected argument 'kms_key_name' to be a str")
        pulumi.set(__self__, "kms_key_name", kms_key_name)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if service_config and not isinstance(service_config, dict):
            raise TypeError("Expected argument 'service_config' to be a dict")
        pulumi.set(__self__, "service_config", service_config)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if state_messages and not isinstance(state_messages, list):
            raise TypeError("Expected argument 'state_messages' to be a list")
        pulumi.set(__self__, "state_messages", state_messages)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="buildConfig")
    def build_config(self) -> 'outputs.BuildConfigResponse':
        """
        Describes the Build step of the function that builds a container from the given source.
        """
        return pulumi.get(self, "build_config")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        User-provided description of a function.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def environment(self) -> str:
        """
        Describe whether the function is 1st Gen or 2nd Gen.
        """
        return pulumi.get(self, "environment")

    @property
    @pulumi.getter(name="eventTrigger")
    def event_trigger(self) -> 'outputs.EventTriggerResponse':
        """
        An Eventarc trigger managed by Google Cloud Functions that fires events in response to a condition in another service.
        """
        return pulumi.get(self, "event_trigger")

    @property
    @pulumi.getter(name="kmsKeyName")
    def kms_key_name(self) -> str:
        """
        [Preview] Resource name of a KMS crypto key (managed by the user) used to encrypt/decrypt function resources. It must match the pattern `projects/{project}/locations/{location}/keyRings/{key_ring}/cryptoKeys/{crypto_key}`.
        """
        return pulumi.get(self, "kms_key_name")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Labels associated with this Cloud Function.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        A user-defined name of the function. Function names must be unique globally and match pattern `projects/*/locations/*/functions/*`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="serviceConfig")
    def service_config(self) -> 'outputs.ServiceConfigResponse':
        """
        Describes the Service being deployed. Currently deploys services to Cloud Run (fully managed).
        """
        return pulumi.get(self, "service_config")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        State of the function.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="stateMessages")
    def state_messages(self) -> Sequence['outputs.GoogleCloudFunctionsV2betaStateMessageResponse']:
        """
        State Messages for this Cloud Function.
        """
        return pulumi.get(self, "state_messages")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The last update timestamp of a Cloud Function.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetFunctionResult(GetFunctionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFunctionResult(
            build_config=self.build_config,
            description=self.description,
            environment=self.environment,
            event_trigger=self.event_trigger,
            kms_key_name=self.kms_key_name,
            labels=self.labels,
            name=self.name,
            service_config=self.service_config,
            state=self.state,
            state_messages=self.state_messages,
            update_time=self.update_time)


def get_function(function_id: Optional[str] = None,
                 location: Optional[str] = None,
                 project: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFunctionResult:
    """
    Returns a function with the given name from the requested project.
    """
    __args__ = dict()
    __args__['functionId'] = function_id
    __args__['location'] = location
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:cloudfunctions/v2beta:getFunction', __args__, opts=opts, typ=GetFunctionResult).value

    return AwaitableGetFunctionResult(
        build_config=pulumi.get(__ret__, 'build_config'),
        description=pulumi.get(__ret__, 'description'),
        environment=pulumi.get(__ret__, 'environment'),
        event_trigger=pulumi.get(__ret__, 'event_trigger'),
        kms_key_name=pulumi.get(__ret__, 'kms_key_name'),
        labels=pulumi.get(__ret__, 'labels'),
        name=pulumi.get(__ret__, 'name'),
        service_config=pulumi.get(__ret__, 'service_config'),
        state=pulumi.get(__ret__, 'state'),
        state_messages=pulumi.get(__ret__, 'state_messages'),
        update_time=pulumi.get(__ret__, 'update_time'))


@_utilities.lift_output_func(get_function)
def get_function_output(function_id: Optional[pulumi.Input[str]] = None,
                        location: Optional[pulumi.Input[str]] = None,
                        project: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFunctionResult]:
    """
    Returns a function with the given name from the requested project.
    """
    ...
