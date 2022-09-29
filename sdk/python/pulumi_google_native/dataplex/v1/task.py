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
from ._enums import *
from ._inputs import *

__all__ = ['TaskArgs', 'Task']

@pulumi.input_type
class TaskArgs:
    def __init__(__self__, *,
                 execution_spec: pulumi.Input['GoogleCloudDataplexV1TaskExecutionSpecArgs'],
                 lake_id: pulumi.Input[str],
                 task_id: pulumi.Input[str],
                 trigger_spec: pulumi.Input['GoogleCloudDataplexV1TaskTriggerSpecArgs'],
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 notebook: Optional[pulumi.Input['GoogleCloudDataplexV1TaskNotebookTaskConfigArgs']] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 spark: Optional[pulumi.Input['GoogleCloudDataplexV1TaskSparkTaskConfigArgs']] = None,
                 validate_only: Optional[pulumi.Input[bool]] = None):
        """
        The set of arguments for constructing a Task resource.
        :param pulumi.Input['GoogleCloudDataplexV1TaskExecutionSpecArgs'] execution_spec: Spec related to how a task is executed.
        :param pulumi.Input[str] task_id: Required. Task identifier.
        :param pulumi.Input['GoogleCloudDataplexV1TaskTriggerSpecArgs'] trigger_spec: Spec related to how often and when a task should be triggered.
        :param pulumi.Input[str] description: Optional. Description of the task.
        :param pulumi.Input[str] display_name: Optional. User friendly display name.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Optional. User-defined labels for the task.
        :param pulumi.Input['GoogleCloudDataplexV1TaskNotebookTaskConfigArgs'] notebook: Config related to running scheduled Notebooks.
        :param pulumi.Input['GoogleCloudDataplexV1TaskSparkTaskConfigArgs'] spark: Config related to running custom Spark tasks.
        :param pulumi.Input[bool] validate_only: Optional. Only validate the request, but do not perform mutations. The default is false.
        """
        pulumi.set(__self__, "execution_spec", execution_spec)
        pulumi.set(__self__, "lake_id", lake_id)
        pulumi.set(__self__, "task_id", task_id)
        pulumi.set(__self__, "trigger_spec", trigger_spec)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if notebook is not None:
            pulumi.set(__self__, "notebook", notebook)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if spark is not None:
            pulumi.set(__self__, "spark", spark)
        if validate_only is not None:
            pulumi.set(__self__, "validate_only", validate_only)

    @property
    @pulumi.getter(name="executionSpec")
    def execution_spec(self) -> pulumi.Input['GoogleCloudDataplexV1TaskExecutionSpecArgs']:
        """
        Spec related to how a task is executed.
        """
        return pulumi.get(self, "execution_spec")

    @execution_spec.setter
    def execution_spec(self, value: pulumi.Input['GoogleCloudDataplexV1TaskExecutionSpecArgs']):
        pulumi.set(self, "execution_spec", value)

    @property
    @pulumi.getter(name="lakeId")
    def lake_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "lake_id")

    @lake_id.setter
    def lake_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "lake_id", value)

    @property
    @pulumi.getter(name="taskId")
    def task_id(self) -> pulumi.Input[str]:
        """
        Required. Task identifier.
        """
        return pulumi.get(self, "task_id")

    @task_id.setter
    def task_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "task_id", value)

    @property
    @pulumi.getter(name="triggerSpec")
    def trigger_spec(self) -> pulumi.Input['GoogleCloudDataplexV1TaskTriggerSpecArgs']:
        """
        Spec related to how often and when a task should be triggered.
        """
        return pulumi.get(self, "trigger_spec")

    @trigger_spec.setter
    def trigger_spec(self, value: pulumi.Input['GoogleCloudDataplexV1TaskTriggerSpecArgs']):
        pulumi.set(self, "trigger_spec", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Description of the task.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. User friendly display name.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Optional. User-defined labels for the task.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def notebook(self) -> Optional[pulumi.Input['GoogleCloudDataplexV1TaskNotebookTaskConfigArgs']]:
        """
        Config related to running scheduled Notebooks.
        """
        return pulumi.get(self, "notebook")

    @notebook.setter
    def notebook(self, value: Optional[pulumi.Input['GoogleCloudDataplexV1TaskNotebookTaskConfigArgs']]):
        pulumi.set(self, "notebook", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def spark(self) -> Optional[pulumi.Input['GoogleCloudDataplexV1TaskSparkTaskConfigArgs']]:
        """
        Config related to running custom Spark tasks.
        """
        return pulumi.get(self, "spark")

    @spark.setter
    def spark(self, value: Optional[pulumi.Input['GoogleCloudDataplexV1TaskSparkTaskConfigArgs']]):
        pulumi.set(self, "spark", value)

    @property
    @pulumi.getter(name="validateOnly")
    def validate_only(self) -> Optional[pulumi.Input[bool]]:
        """
        Optional. Only validate the request, but do not perform mutations. The default is false.
        """
        return pulumi.get(self, "validate_only")

    @validate_only.setter
    def validate_only(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "validate_only", value)


class Task(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 execution_spec: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDataplexV1TaskExecutionSpecArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 lake_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 notebook: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDataplexV1TaskNotebookTaskConfigArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 spark: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDataplexV1TaskSparkTaskConfigArgs']]] = None,
                 task_id: Optional[pulumi.Input[str]] = None,
                 trigger_spec: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDataplexV1TaskTriggerSpecArgs']]] = None,
                 validate_only: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        """
        Creates a task resource within a lake.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Optional. Description of the task.
        :param pulumi.Input[str] display_name: Optional. User friendly display name.
        :param pulumi.Input[pulumi.InputType['GoogleCloudDataplexV1TaskExecutionSpecArgs']] execution_spec: Spec related to how a task is executed.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Optional. User-defined labels for the task.
        :param pulumi.Input[pulumi.InputType['GoogleCloudDataplexV1TaskNotebookTaskConfigArgs']] notebook: Config related to running scheduled Notebooks.
        :param pulumi.Input[pulumi.InputType['GoogleCloudDataplexV1TaskSparkTaskConfigArgs']] spark: Config related to running custom Spark tasks.
        :param pulumi.Input[str] task_id: Required. Task identifier.
        :param pulumi.Input[pulumi.InputType['GoogleCloudDataplexV1TaskTriggerSpecArgs']] trigger_spec: Spec related to how often and when a task should be triggered.
        :param pulumi.Input[bool] validate_only: Optional. Only validate the request, but do not perform mutations. The default is false.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: TaskArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a task resource within a lake.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param TaskArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(TaskArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 execution_spec: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDataplexV1TaskExecutionSpecArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 lake_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 notebook: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDataplexV1TaskNotebookTaskConfigArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 spark: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDataplexV1TaskSparkTaskConfigArgs']]] = None,
                 task_id: Optional[pulumi.Input[str]] = None,
                 trigger_spec: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDataplexV1TaskTriggerSpecArgs']]] = None,
                 validate_only: Optional[pulumi.Input[bool]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = TaskArgs.__new__(TaskArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["display_name"] = display_name
            if execution_spec is None and not opts.urn:
                raise TypeError("Missing required property 'execution_spec'")
            __props__.__dict__["execution_spec"] = execution_spec
            __props__.__dict__["labels"] = labels
            if lake_id is None and not opts.urn:
                raise TypeError("Missing required property 'lake_id'")
            __props__.__dict__["lake_id"] = lake_id
            __props__.__dict__["location"] = location
            __props__.__dict__["notebook"] = notebook
            __props__.__dict__["project"] = project
            __props__.__dict__["spark"] = spark
            if task_id is None and not opts.urn:
                raise TypeError("Missing required property 'task_id'")
            __props__.__dict__["task_id"] = task_id
            if trigger_spec is None and not opts.urn:
                raise TypeError("Missing required property 'trigger_spec'")
            __props__.__dict__["trigger_spec"] = trigger_spec
            __props__.__dict__["validate_only"] = validate_only
            __props__.__dict__["create_time"] = None
            __props__.__dict__["execution_status"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["uid"] = None
            __props__.__dict__["update_time"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["lake_id", "location", "project", "task_id"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Task, __self__).__init__(
            'google-native:dataplex/v1:Task',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Task':
        """
        Get an existing Task resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = TaskArgs.__new__(TaskArgs)

        __props__.__dict__["create_time"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["execution_spec"] = None
        __props__.__dict__["execution_status"] = None
        __props__.__dict__["labels"] = None
        __props__.__dict__["lake_id"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["notebook"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["spark"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["task_id"] = None
        __props__.__dict__["trigger_spec"] = None
        __props__.__dict__["uid"] = None
        __props__.__dict__["update_time"] = None
        __props__.__dict__["validate_only"] = None
        return Task(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The time when the task was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Optional. Description of the task.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Optional. User friendly display name.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="executionSpec")
    def execution_spec(self) -> pulumi.Output['outputs.GoogleCloudDataplexV1TaskExecutionSpecResponse']:
        """
        Spec related to how a task is executed.
        """
        return pulumi.get(self, "execution_spec")

    @property
    @pulumi.getter(name="executionStatus")
    def execution_status(self) -> pulumi.Output['outputs.GoogleCloudDataplexV1TaskExecutionStatusResponse']:
        """
        Status of the latest task executions.
        """
        return pulumi.get(self, "execution_status")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Optional. User-defined labels for the task.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="lakeId")
    def lake_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "lake_id")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The relative resource name of the task, of the form: projects/{project_number}/locations/{location_id}/lakes/{lake_id}/ tasks/{task_id}.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def notebook(self) -> pulumi.Output['outputs.GoogleCloudDataplexV1TaskNotebookTaskConfigResponse']:
        """
        Config related to running scheduled Notebooks.
        """
        return pulumi.get(self, "notebook")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def spark(self) -> pulumi.Output['outputs.GoogleCloudDataplexV1TaskSparkTaskConfigResponse']:
        """
        Config related to running custom Spark tasks.
        """
        return pulumi.get(self, "spark")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        Current state of the task.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="taskId")
    def task_id(self) -> pulumi.Output[str]:
        """
        Required. Task identifier.
        """
        return pulumi.get(self, "task_id")

    @property
    @pulumi.getter(name="triggerSpec")
    def trigger_spec(self) -> pulumi.Output['outputs.GoogleCloudDataplexV1TaskTriggerSpecResponse']:
        """
        Spec related to how often and when a task should be triggered.
        """
        return pulumi.get(self, "trigger_spec")

    @property
    @pulumi.getter
    def uid(self) -> pulumi.Output[str]:
        """
        System generated globally unique ID for the task. This ID will be different if the task is deleted and re-created with the same name.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The time when the task was last updated.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter(name="validateOnly")
    def validate_only(self) -> pulumi.Output[Optional[bool]]:
        """
        Optional. Only validate the request, but do not perform mutations. The default is false.
        """
        return pulumi.get(self, "validate_only")

