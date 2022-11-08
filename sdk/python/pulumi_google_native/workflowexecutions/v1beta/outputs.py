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

__all__ = [
    'ErrorResponse',
    'PositionResponse',
    'StackTraceElementResponse',
    'StackTraceResponse',
    'StatusResponse',
    'StepResponse',
]

@pulumi.output_type
class ErrorResponse(dict):
    """
    Error describes why the execution was abnormally terminated.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "stackTrace":
            suggest = "stack_trace"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in ErrorResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        ErrorResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        ErrorResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 context: str,
                 payload: str,
                 stack_trace: 'outputs.StackTraceResponse'):
        """
        Error describes why the execution was abnormally terminated.
        :param str context: Human-readable stack trace string.
        :param str payload: Error message and data returned represented as a JSON string.
        :param 'StackTraceResponse' stack_trace: Stack trace with detailed information of where error was generated.
        """
        pulumi.set(__self__, "context", context)
        pulumi.set(__self__, "payload", payload)
        pulumi.set(__self__, "stack_trace", stack_trace)

    @property
    @pulumi.getter
    def context(self) -> str:
        """
        Human-readable stack trace string.
        """
        return pulumi.get(self, "context")

    @property
    @pulumi.getter
    def payload(self) -> str:
        """
        Error message and data returned represented as a JSON string.
        """
        return pulumi.get(self, "payload")

    @property
    @pulumi.getter(name="stackTrace")
    def stack_trace(self) -> 'outputs.StackTraceResponse':
        """
        Stack trace with detailed information of where error was generated.
        """
        return pulumi.get(self, "stack_trace")


@pulumi.output_type
class PositionResponse(dict):
    """
    Position contains source position information about the stack trace element such as line number, column number and length of the code block in bytes.
    """
    def __init__(__self__, *,
                 column: str,
                 length: str,
                 line: str):
        """
        Position contains source position information about the stack trace element such as line number, column number and length of the code block in bytes.
        :param str column: The source code column position (of the line) the current instruction was generated from.
        :param str length: The number of bytes of source code making up this stack trace element.
        :param str line: The source code line number the current instruction was generated from.
        """
        pulumi.set(__self__, "column", column)
        pulumi.set(__self__, "length", length)
        pulumi.set(__self__, "line", line)

    @property
    @pulumi.getter
    def column(self) -> str:
        """
        The source code column position (of the line) the current instruction was generated from.
        """
        return pulumi.get(self, "column")

    @property
    @pulumi.getter
    def length(self) -> str:
        """
        The number of bytes of source code making up this stack trace element.
        """
        return pulumi.get(self, "length")

    @property
    @pulumi.getter
    def line(self) -> str:
        """
        The source code line number the current instruction was generated from.
        """
        return pulumi.get(self, "line")


@pulumi.output_type
class StackTraceElementResponse(dict):
    """
    A single stack element (frame) where an error occurred.
    """
    def __init__(__self__, *,
                 position: 'outputs.PositionResponse',
                 routine: str,
                 step: str):
        """
        A single stack element (frame) where an error occurred.
        :param 'PositionResponse' position: The source position information of the stack trace element.
        :param str routine: The routine where the error occurred.
        :param str step: The step the error occurred at.
        """
        pulumi.set(__self__, "position", position)
        pulumi.set(__self__, "routine", routine)
        pulumi.set(__self__, "step", step)

    @property
    @pulumi.getter
    def position(self) -> 'outputs.PositionResponse':
        """
        The source position information of the stack trace element.
        """
        return pulumi.get(self, "position")

    @property
    @pulumi.getter
    def routine(self) -> str:
        """
        The routine where the error occurred.
        """
        return pulumi.get(self, "routine")

    @property
    @pulumi.getter
    def step(self) -> str:
        """
        The step the error occurred at.
        """
        return pulumi.get(self, "step")


@pulumi.output_type
class StackTraceResponse(dict):
    """
    A collection of stack elements (frames) where an error occurred.
    """
    def __init__(__self__, *,
                 elements: Sequence['outputs.StackTraceElementResponse']):
        """
        A collection of stack elements (frames) where an error occurred.
        :param Sequence['StackTraceElementResponse'] elements: An array of stack elements.
        """
        pulumi.set(__self__, "elements", elements)

    @property
    @pulumi.getter
    def elements(self) -> Sequence['outputs.StackTraceElementResponse']:
        """
        An array of stack elements.
        """
        return pulumi.get(self, "elements")


@pulumi.output_type
class StatusResponse(dict):
    """
    Represents the current status of this execution.
    """
    @staticmethod
    def __key_warning(key: str):
        suggest = None
        if key == "currentSteps":
            suggest = "current_steps"

        if suggest:
            pulumi.log.warn(f"Key '{key}' not found in StatusResponse. Access the value via the '{suggest}' property getter instead.")

    def __getitem__(self, key: str) -> Any:
        StatusResponse.__key_warning(key)
        return super().__getitem__(key)

    def get(self, key: str, default = None) -> Any:
        StatusResponse.__key_warning(key)
        return super().get(key, default)

    def __init__(__self__, *,
                 current_steps: Sequence['outputs.StepResponse']):
        """
        Represents the current status of this execution.
        :param Sequence['StepResponse'] current_steps: A list of currently executing or last executed step names for the workflow execution currently running. If the workflow has succeeded or failed, this is the last attempted or executed step. Presently, if the current step is inside a subworkflow, the list only includes that step. In the future, the list will contain items for each step in the call stack, starting with the outermost step in the `main` subworkflow, and ending with the most deeply nested step.
        """
        pulumi.set(__self__, "current_steps", current_steps)

    @property
    @pulumi.getter(name="currentSteps")
    def current_steps(self) -> Sequence['outputs.StepResponse']:
        """
        A list of currently executing or last executed step names for the workflow execution currently running. If the workflow has succeeded or failed, this is the last attempted or executed step. Presently, if the current step is inside a subworkflow, the list only includes that step. In the future, the list will contain items for each step in the call stack, starting with the outermost step in the `main` subworkflow, and ending with the most deeply nested step.
        """
        return pulumi.get(self, "current_steps")


@pulumi.output_type
class StepResponse(dict):
    """
    Represents a step of the workflow this execution is running.
    """
    def __init__(__self__, *,
                 routine: str,
                 step: str):
        """
        Represents a step of the workflow this execution is running.
        :param str routine: Name of a routine within the workflow.
        :param str step: Name of a step within the routine.
        """
        pulumi.set(__self__, "routine", routine)
        pulumi.set(__self__, "step", step)

    @property
    @pulumi.getter
    def routine(self) -> str:
        """
        Name of a routine within the workflow.
        """
        return pulumi.get(self, "routine")

    @property
    @pulumi.getter
    def step(self) -> str:
        """
        Name of a step within the routine.
        """
        return pulumi.get(self, "step")


