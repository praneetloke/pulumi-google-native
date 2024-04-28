# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'GetEntityResult',
    'AwaitableGetEntityResult',
    'get_entity',
    'get_entity_output',
]

@pulumi.output_type
class GetEntityResult:
    def __init__(__self__, fields=None, name=None):
        if fields and not isinstance(fields, dict):
            raise TypeError("Expected argument 'fields' to be a dict")
        pulumi.set(__self__, "fields", fields)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def fields(self) -> Mapping[str, Any]:
        """
        Fields of the entity. The key is name of the field and the value contains the applicable `google.protobuf.Value` entry for this field.
        """
        return pulumi.get(self, "fields")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name of the Entity. Format: projects/{project}/locations/{location}/connections/{connection}/entityTypes/{type}/entities/{id}
        """
        return pulumi.get(self, "name")


class AwaitableGetEntityResult(GetEntityResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetEntityResult(
            fields=self.fields,
            name=self.name)


def get_entity(connection_id: Optional[str] = None,
               entity_id: Optional[str] = None,
               entity_type_id: Optional[str] = None,
               location: Optional[str] = None,
               project: Optional[str] = None,
               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetEntityResult:
    """
    Gets a single entity row matching the entity type and entity id specified in the request.
    """
    __args__ = dict()
    __args__['connectionId'] = connection_id
    __args__['entityId'] = entity_id
    __args__['entityTypeId'] = entity_type_id
    __args__['location'] = location
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:connectors/v2:getEntity', __args__, opts=opts, typ=GetEntityResult).value

    return AwaitableGetEntityResult(
        fields=pulumi.get(__ret__, 'fields'),
        name=pulumi.get(__ret__, 'name'))


@_utilities.lift_output_func(get_entity)
def get_entity_output(connection_id: Optional[pulumi.Input[str]] = None,
                      entity_id: Optional[pulumi.Input[str]] = None,
                      entity_type_id: Optional[pulumi.Input[str]] = None,
                      location: Optional[pulumi.Input[str]] = None,
                      project: Optional[pulumi.Input[Optional[str]]] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetEntityResult]:
    """
    Gets a single entity row matching the entity type and entity id specified in the request.
    """
    ...
