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
    'GetDeliveryPipelineResult',
    'AwaitableGetDeliveryPipelineResult',
    'get_delivery_pipeline',
    'get_delivery_pipeline_output',
]

@pulumi.output_type
class GetDeliveryPipelineResult:
    def __init__(__self__, annotations=None, condition=None, create_time=None, description=None, etag=None, labels=None, name=None, serial_pipeline=None, suspended=None, uid=None, update_time=None):
        if annotations and not isinstance(annotations, dict):
            raise TypeError("Expected argument 'annotations' to be a dict")
        pulumi.set(__self__, "annotations", annotations)
        if condition and not isinstance(condition, dict):
            raise TypeError("Expected argument 'condition' to be a dict")
        pulumi.set(__self__, "condition", condition)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if serial_pipeline and not isinstance(serial_pipeline, dict):
            raise TypeError("Expected argument 'serial_pipeline' to be a dict")
        pulumi.set(__self__, "serial_pipeline", serial_pipeline)
        if suspended and not isinstance(suspended, bool):
            raise TypeError("Expected argument 'suspended' to be a bool")
        pulumi.set(__self__, "suspended", suspended)
        if uid and not isinstance(uid, str):
            raise TypeError("Expected argument 'uid' to be a str")
        pulumi.set(__self__, "uid", uid)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter
    def annotations(self) -> Mapping[str, str]:
        """
        User annotations. These attributes can only be set and used by the user, and not by Google Cloud Deploy.
        """
        return pulumi.get(self, "annotations")

    @property
    @pulumi.getter
    def condition(self) -> 'outputs.PipelineConditionResponse':
        """
        Information around the state of the Delivery Pipeline.
        """
        return pulumi.get(self, "condition")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Time at which the pipeline was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Description of the `DeliveryPipeline`. Max length is 255 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Labels are attributes that can be set and used by both the user and by Google Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be <= 128 bytes.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Optional. Name of the `DeliveryPipeline`. Format is projects/{project}/ locations/{location}/deliveryPipelines/a-z{0,62}.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="serialPipeline")
    def serial_pipeline(self) -> 'outputs.SerialPipelineResponse':
        """
        SerialPipeline defines a sequential set of stages for a `DeliveryPipeline`.
        """
        return pulumi.get(self, "serial_pipeline")

    @property
    @pulumi.getter
    def suspended(self) -> bool:
        """
        When suspended, no new releases or rollouts can be created, but in-progress ones will complete.
        """
        return pulumi.get(self, "suspended")

    @property
    @pulumi.getter
    def uid(self) -> str:
        """
        Unique identifier of the `DeliveryPipeline`.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        Most recent time at which the pipeline was updated.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetDeliveryPipelineResult(GetDeliveryPipelineResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDeliveryPipelineResult(
            annotations=self.annotations,
            condition=self.condition,
            create_time=self.create_time,
            description=self.description,
            etag=self.etag,
            labels=self.labels,
            name=self.name,
            serial_pipeline=self.serial_pipeline,
            suspended=self.suspended,
            uid=self.uid,
            update_time=self.update_time)


def get_delivery_pipeline(delivery_pipeline_id: Optional[str] = None,
                          location: Optional[str] = None,
                          project: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDeliveryPipelineResult:
    """
    Gets details of a single DeliveryPipeline.
    """
    __args__ = dict()
    __args__['deliveryPipelineId'] = delivery_pipeline_id
    __args__['location'] = location
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:clouddeploy/v1:getDeliveryPipeline', __args__, opts=opts, typ=GetDeliveryPipelineResult).value

    return AwaitableGetDeliveryPipelineResult(
        annotations=pulumi.get(__ret__, 'annotations'),
        condition=pulumi.get(__ret__, 'condition'),
        create_time=pulumi.get(__ret__, 'create_time'),
        description=pulumi.get(__ret__, 'description'),
        etag=pulumi.get(__ret__, 'etag'),
        labels=pulumi.get(__ret__, 'labels'),
        name=pulumi.get(__ret__, 'name'),
        serial_pipeline=pulumi.get(__ret__, 'serial_pipeline'),
        suspended=pulumi.get(__ret__, 'suspended'),
        uid=pulumi.get(__ret__, 'uid'),
        update_time=pulumi.get(__ret__, 'update_time'))


@_utilities.lift_output_func(get_delivery_pipeline)
def get_delivery_pipeline_output(delivery_pipeline_id: Optional[pulumi.Input[str]] = None,
                                 location: Optional[pulumi.Input[str]] = None,
                                 project: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDeliveryPipelineResult]:
    """
    Gets details of a single DeliveryPipeline.
    """
    ...
