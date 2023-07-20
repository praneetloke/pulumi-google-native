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
    'GetProjectResult',
    'AwaitableGetProjectResult',
    'get_project',
    'get_project_output',
]

@pulumi.output_type
class GetProjectResult:
    def __init__(__self__, create_time=None, labels=None, lifecycle_state=None, name=None, parent=None, project_id=None, project_number=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if lifecycle_state and not isinstance(lifecycle_state, str):
            raise TypeError("Expected argument 'lifecycle_state' to be a str")
        pulumi.set(__self__, "lifecycle_state", lifecycle_state)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if parent and not isinstance(parent, dict):
            raise TypeError("Expected argument 'parent' to be a dict")
        pulumi.set(__self__, "parent", parent)
        if project_id and not isinstance(project_id, str):
            raise TypeError("Expected argument 'project_id' to be a str")
        pulumi.set(__self__, "project_id", project_id)
        if project_number and not isinstance(project_number, str):
            raise TypeError("Expected argument 'project_number' to be a str")
        pulumi.set(__self__, "project_number", project_number)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Creation time. Read-only.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        The labels associated with this Project. Label keys must be between 1 and 63 characters long and must conform to the following regular expression: a-z{0,62}. Label values must be between 0 and 63 characters long and must conform to the regular expression [a-z0-9_-]{0,63}. A label value can be empty. No more than 256 labels can be associated with a given resource. Clients should store labels in a representation such as JSON that does not depend on specific characters being disallowed. Example: `"environment" : "dev"` Read-write.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="lifecycleState")
    def lifecycle_state(self) -> str:
        """
        The Project lifecycle state. Read-only.
        """
        return pulumi.get(self, "lifecycle_state")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The optional user-assigned display name of the Project. When present it must be between 4 to 30 characters. Allowed characters are: lowercase and uppercase letters, numbers, hyphen, single-quote, double-quote, space, and exclamation point. Example: `My Project` Read-write.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parent(self) -> 'outputs.ResourceIdResponse':
        """
        An optional reference to a parent Resource. Supported parent types include "organization" and "folder". Once set, the parent cannot be cleared. The `parent` can be set on creation or using the `UpdateProject` method; the end user must have the `resourcemanager.projects.create` permission on the parent. Read-write.
        """
        return pulumi.get(self, "parent")

    @property
    @pulumi.getter(name="projectId")
    def project_id(self) -> str:
        """
        The unique, user-assigned ID of the Project. It must be 6 to 30 lowercase letters, digits, or hyphens. It must start with a letter. Trailing hyphens are prohibited. Example: `tokyo-rain-123` Read-only after creation.
        """
        return pulumi.get(self, "project_id")

    @property
    @pulumi.getter(name="projectNumber")
    def project_number(self) -> str:
        """
        The number uniquely identifying the project. Example: `415104041262` Read-only.
        """
        return pulumi.get(self, "project_number")


class AwaitableGetProjectResult(GetProjectResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProjectResult(
            create_time=self.create_time,
            labels=self.labels,
            lifecycle_state=self.lifecycle_state,
            name=self.name,
            parent=self.parent,
            project_id=self.project_id,
            project_number=self.project_number)


def get_project(project: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProjectResult:
    """
    Retrieves the Project identified by the specified `project_id` (for example, `my-project-123`). The caller must have read permissions for this Project.
    """
    __args__ = dict()
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:cloudresourcemanager/v1beta1:getProject', __args__, opts=opts, typ=GetProjectResult).value

    return AwaitableGetProjectResult(
        create_time=pulumi.get(__ret__, 'create_time'),
        labels=pulumi.get(__ret__, 'labels'),
        lifecycle_state=pulumi.get(__ret__, 'lifecycle_state'),
        name=pulumi.get(__ret__, 'name'),
        parent=pulumi.get(__ret__, 'parent'),
        project_id=pulumi.get(__ret__, 'project_id'),
        project_number=pulumi.get(__ret__, 'project_number'))


@_utilities.lift_output_func(get_project)
def get_project_output(project: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetProjectResult]:
    """
    Retrieves the Project identified by the specified `project_id` (for example, `my-project-123`). The caller must have read permissions for this Project.
    """
    ...
