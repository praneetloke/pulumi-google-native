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

__all__ = ['ExecutionArgs', 'Execution']

@pulumi.input_type
class ExecutionArgs:
    def __init__(__self__, *,
                 workflow_id: pulumi.Input[str],
                 argument: Optional[pulumi.Input[str]] = None,
                 call_log_level: Optional[pulumi.Input['ExecutionCallLogLevel']] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Execution resource.
        :param pulumi.Input[str] argument: Input parameters of the execution represented as a JSON string. The size limit is 32KB. *Note*: If you are using the REST API directly to run your workflow, you must escape any JSON string value of `argument`. Example: `'{"argument":"{\\"firstName\\":\\"FIRST\\",\\"lastName\\":\\"LAST\\"}"}'`
        :param pulumi.Input['ExecutionCallLogLevel'] call_log_level: The call logging level associated to this execution.
        """
        pulumi.set(__self__, "workflow_id", workflow_id)
        if argument is not None:
            pulumi.set(__self__, "argument", argument)
        if call_log_level is not None:
            pulumi.set(__self__, "call_log_level", call_log_level)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="workflowId")
    def workflow_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "workflow_id")

    @workflow_id.setter
    def workflow_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "workflow_id", value)

    @property
    @pulumi.getter
    def argument(self) -> Optional[pulumi.Input[str]]:
        """
        Input parameters of the execution represented as a JSON string. The size limit is 32KB. *Note*: If you are using the REST API directly to run your workflow, you must escape any JSON string value of `argument`. Example: `'{"argument":"{\\"firstName\\":\\"FIRST\\",\\"lastName\\":\\"LAST\\"}"}'`
        """
        return pulumi.get(self, "argument")

    @argument.setter
    def argument(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "argument", value)

    @property
    @pulumi.getter(name="callLogLevel")
    def call_log_level(self) -> Optional[pulumi.Input['ExecutionCallLogLevel']]:
        """
        The call logging level associated to this execution.
        """
        return pulumi.get(self, "call_log_level")

    @call_log_level.setter
    def call_log_level(self, value: Optional[pulumi.Input['ExecutionCallLogLevel']]):
        pulumi.set(self, "call_log_level", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


class Execution(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 argument: Optional[pulumi.Input[str]] = None,
                 call_log_level: Optional[pulumi.Input['ExecutionCallLogLevel']] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 workflow_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a new execution using the latest revision of the given workflow.
        Auto-naming is currently not supported for this resource.
        Note - this resource's API doesn't support deletion. When deleted, the resource will persist
        on Google Cloud even though it will be deleted from Pulumi state.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] argument: Input parameters of the execution represented as a JSON string. The size limit is 32KB. *Note*: If you are using the REST API directly to run your workflow, you must escape any JSON string value of `argument`. Example: `'{"argument":"{\\"firstName\\":\\"FIRST\\",\\"lastName\\":\\"LAST\\"}"}'`
        :param pulumi.Input['ExecutionCallLogLevel'] call_log_level: The call logging level associated to this execution.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ExecutionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new execution using the latest revision of the given workflow.
        Auto-naming is currently not supported for this resource.
        Note - this resource's API doesn't support deletion. When deleted, the resource will persist
        on Google Cloud even though it will be deleted from Pulumi state.

        :param str resource_name: The name of the resource.
        :param ExecutionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ExecutionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 argument: Optional[pulumi.Input[str]] = None,
                 call_log_level: Optional[pulumi.Input['ExecutionCallLogLevel']] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 workflow_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ExecutionArgs.__new__(ExecutionArgs)

            __props__.__dict__["argument"] = argument
            __props__.__dict__["call_log_level"] = call_log_level
            __props__.__dict__["location"] = location
            __props__.__dict__["project"] = project
            if workflow_id is None and not opts.urn:
                raise TypeError("Missing required property 'workflow_id'")
            __props__.__dict__["workflow_id"] = workflow_id
            __props__.__dict__["end_time"] = None
            __props__.__dict__["error"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["result"] = None
            __props__.__dict__["start_time"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["status"] = None
            __props__.__dict__["workflow_revision_id"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["location", "project", "workflow_id"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Execution, __self__).__init__(
            'google-native:workflowexecutions/v1:Execution',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Execution':
        """
        Get an existing Execution resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ExecutionArgs.__new__(ExecutionArgs)

        __props__.__dict__["argument"] = None
        __props__.__dict__["call_log_level"] = None
        __props__.__dict__["end_time"] = None
        __props__.__dict__["error"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["result"] = None
        __props__.__dict__["start_time"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["status"] = None
        __props__.__dict__["workflow_id"] = None
        __props__.__dict__["workflow_revision_id"] = None
        return Execution(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def argument(self) -> pulumi.Output[str]:
        """
        Input parameters of the execution represented as a JSON string. The size limit is 32KB. *Note*: If you are using the REST API directly to run your workflow, you must escape any JSON string value of `argument`. Example: `'{"argument":"{\\"firstName\\":\\"FIRST\\",\\"lastName\\":\\"LAST\\"}"}'`
        """
        return pulumi.get(self, "argument")

    @property
    @pulumi.getter(name="callLogLevel")
    def call_log_level(self) -> pulumi.Output[str]:
        """
        The call logging level associated to this execution.
        """
        return pulumi.get(self, "call_log_level")

    @property
    @pulumi.getter(name="endTime")
    def end_time(self) -> pulumi.Output[str]:
        """
        Marks the end of execution, successful or not.
        """
        return pulumi.get(self, "end_time")

    @property
    @pulumi.getter
    def error(self) -> pulumi.Output['outputs.ErrorResponse']:
        """
        The error which caused the execution to finish prematurely. The value is only present if the execution's state is `FAILED` or `CANCELLED`.
        """
        return pulumi.get(self, "error")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name of the execution. Format: projects/{project}/locations/{location}/workflows/{workflow}/executions/{execution}
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def result(self) -> pulumi.Output[str]:
        """
        Output of the execution represented as a JSON string. The value can only be present if the execution's state is `SUCCEEDED`.
        """
        return pulumi.get(self, "result")

    @property
    @pulumi.getter(name="startTime")
    def start_time(self) -> pulumi.Output[str]:
        """
        Marks the beginning of execution.
        """
        return pulumi.get(self, "start_time")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        Current state of the execution.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def status(self) -> pulumi.Output['outputs.StatusResponse']:
        """
        Status tracks the current steps and progress data of this execution.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="workflowId")
    def workflow_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "workflow_id")

    @property
    @pulumi.getter(name="workflowRevisionId")
    def workflow_revision_id(self) -> pulumi.Output[str]:
        """
        Revision of the workflow this execution is using.
        """
        return pulumi.get(self, "workflow_revision_id")

