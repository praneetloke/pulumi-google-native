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
    'GetProcessResult',
    'AwaitableGetProcessResult',
    'get_process',
    'get_process_output',
]

@pulumi.output_type
class GetProcessResult:
    def __init__(__self__, attributes=None, display_name=None, name=None, origin=None):
        if attributes and not isinstance(attributes, dict):
            raise TypeError("Expected argument 'attributes' to be a dict")
        pulumi.set(__self__, "attributes", attributes)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if origin and not isinstance(origin, dict):
            raise TypeError("Expected argument 'origin' to be a dict")
        pulumi.set(__self__, "origin", origin)

    @property
    @pulumi.getter
    def attributes(self) -> Mapping[str, str]:
        """
        Optional. The attributes of the process. Should only be used for the purpose of non-semantic management (classifying, describing or labeling the process). Up to 100 attributes are allowed.
        """
        return pulumi.get(self, "attributes")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Optional. A human-readable name you can set to display in a user interface. Must be not longer than 200 characters and only contain UTF-8 letters or numbers, spaces or characters like `_-:&.`
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Immutable. The resource name of the lineage process. Format: `projects/{project}/locations/{location}/processes/{process}`. Can be specified or auto-assigned. {process} must be not longer than 200 characters and only contain characters in a set: `a-zA-Z0-9_-:.`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def origin(self) -> 'outputs.GoogleCloudDatacatalogLineageV1OriginResponse':
        """
        Optional. The origin of this process and its runs and lineage events.
        """
        return pulumi.get(self, "origin")


class AwaitableGetProcessResult(GetProcessResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetProcessResult(
            attributes=self.attributes,
            display_name=self.display_name,
            name=self.name,
            origin=self.origin)


def get_process(location: Optional[str] = None,
                process_id: Optional[str] = None,
                project: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetProcessResult:
    """
    Gets the details of the specified process.
    """
    __args__ = dict()
    __args__['location'] = location
    __args__['processId'] = process_id
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:datalineage/v1:getProcess', __args__, opts=opts, typ=GetProcessResult).value

    return AwaitableGetProcessResult(
        attributes=pulumi.get(__ret__, 'attributes'),
        display_name=pulumi.get(__ret__, 'display_name'),
        name=pulumi.get(__ret__, 'name'),
        origin=pulumi.get(__ret__, 'origin'))


@_utilities.lift_output_func(get_process)
def get_process_output(location: Optional[pulumi.Input[str]] = None,
                       process_id: Optional[pulumi.Input[str]] = None,
                       project: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetProcessResult]:
    """
    Gets the details of the specified process.
    """
    ...
