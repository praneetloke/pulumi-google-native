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
    'GetDataAttributeBindingResult',
    'AwaitableGetDataAttributeBindingResult',
    'get_data_attribute_binding',
    'get_data_attribute_binding_output',
]

@pulumi.output_type
class GetDataAttributeBindingResult:
    def __init__(__self__, attributes=None, create_time=None, description=None, display_name=None, etag=None, labels=None, name=None, paths=None, resource=None, uid=None, update_time=None):
        if attributes and not isinstance(attributes, list):
            raise TypeError("Expected argument 'attributes' to be a list")
        pulumi.set(__self__, "attributes", attributes)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if paths and not isinstance(paths, list):
            raise TypeError("Expected argument 'paths' to be a list")
        pulumi.set(__self__, "paths", paths)
        if resource and not isinstance(resource, str):
            raise TypeError("Expected argument 'resource' to be a str")
        pulumi.set(__self__, "resource", resource)
        if uid and not isinstance(uid, str):
            raise TypeError("Expected argument 'uid' to be a str")
        pulumi.set(__self__, "uid", uid)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter
    def attributes(self) -> Sequence[str]:
        """
        Optional. List of attributes to be associated with the resource, provided in the form: projects/{project}/locations/{location}/dataTaxonomies/{dataTaxonomy}/attributes/{data_attribute_id}
        """
        return pulumi.get(self, "attributes")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The time when the DataAttributeBinding was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Optional. Description of the DataAttributeBinding.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Optional. User friendly display name.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding. Etags must be used when calling the DeleteDataAttributeBinding and the UpdateDataAttributeBinding method.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Optional. User-defined labels for the DataAttributeBinding.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The relative resource name of the Data Attribute Binding, of the form: projects/{project_number}/locations/{location}/dataAttributeBindings/{data_attribute_binding_id}
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def paths(self) -> Sequence['outputs.GoogleCloudDataplexV1DataAttributeBindingPathResponse']:
        """
        Optional. The list of paths for items within the associated resource (eg. columns and partitions within a table) along with attribute bindings.
        """
        return pulumi.get(self, "paths")

    @property
    @pulumi.getter
    def resource(self) -> str:
        """
        Optional. Immutable. The resource name of the resource that is associated to attributes. Presently, only entity resource is supported in the form: projects/{project}/locations/{location}/lakes/{lake}/zones/{zone}/entities/{entity_id} Must belong in the same project and region as the attribute binding, and there can only exist one active binding for a resource.
        """
        return pulumi.get(self, "resource")

    @property
    @pulumi.getter
    def uid(self) -> str:
        """
        System generated globally unique ID for the DataAttributeBinding. This ID will be different if the DataAttributeBinding is deleted and re-created with the same name.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The time when the DataAttributeBinding was last updated.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetDataAttributeBindingResult(GetDataAttributeBindingResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDataAttributeBindingResult(
            attributes=self.attributes,
            create_time=self.create_time,
            description=self.description,
            display_name=self.display_name,
            etag=self.etag,
            labels=self.labels,
            name=self.name,
            paths=self.paths,
            resource=self.resource,
            uid=self.uid,
            update_time=self.update_time)


def get_data_attribute_binding(data_attribute_binding_id: Optional[str] = None,
                               location: Optional[str] = None,
                               project: Optional[str] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDataAttributeBindingResult:
    """
    Retrieves a DataAttributeBinding resource.
    """
    __args__ = dict()
    __args__['dataAttributeBindingId'] = data_attribute_binding_id
    __args__['location'] = location
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:dataplex/v1:getDataAttributeBinding', __args__, opts=opts, typ=GetDataAttributeBindingResult).value

    return AwaitableGetDataAttributeBindingResult(
        attributes=__ret__.attributes,
        create_time=__ret__.create_time,
        description=__ret__.description,
        display_name=__ret__.display_name,
        etag=__ret__.etag,
        labels=__ret__.labels,
        name=__ret__.name,
        paths=__ret__.paths,
        resource=__ret__.resource,
        uid=__ret__.uid,
        update_time=__ret__.update_time)


@_utilities.lift_output_func(get_data_attribute_binding)
def get_data_attribute_binding_output(data_attribute_binding_id: Optional[pulumi.Input[str]] = None,
                                      location: Optional[pulumi.Input[str]] = None,
                                      project: Optional[pulumi.Input[Optional[str]]] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDataAttributeBindingResult]:
    """
    Retrieves a DataAttributeBinding resource.
    """
    ...
