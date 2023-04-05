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
    'GetAttributeResult',
    'AwaitableGetAttributeResult',
    'get_attribute',
    'get_attribute_output',
]

@pulumi.output_type
class GetAttributeResult:
    def __init__(__self__, attribute_count=None, create_time=None, data_access_spec=None, description=None, display_name=None, etag=None, labels=None, name=None, parent_id=None, resource_access_spec=None, uid=None, update_time=None):
        if attribute_count and not isinstance(attribute_count, int):
            raise TypeError("Expected argument 'attribute_count' to be a int")
        pulumi.set(__self__, "attribute_count", attribute_count)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if data_access_spec and not isinstance(data_access_spec, dict):
            raise TypeError("Expected argument 'data_access_spec' to be a dict")
        pulumi.set(__self__, "data_access_spec", data_access_spec)
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
        if parent_id and not isinstance(parent_id, str):
            raise TypeError("Expected argument 'parent_id' to be a str")
        pulumi.set(__self__, "parent_id", parent_id)
        if resource_access_spec and not isinstance(resource_access_spec, dict):
            raise TypeError("Expected argument 'resource_access_spec' to be a dict")
        pulumi.set(__self__, "resource_access_spec", resource_access_spec)
        if uid and not isinstance(uid, str):
            raise TypeError("Expected argument 'uid' to be a str")
        pulumi.set(__self__, "uid", uid)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="attributeCount")
    def attribute_count(self) -> int:
        """
        The number of child attributes present for this attribute.
        """
        return pulumi.get(self, "attribute_count")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The time when the DataAttribute was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="dataAccessSpec")
    def data_access_spec(self) -> 'outputs.GoogleCloudDataplexV1DataAccessSpecResponse':
        """
        Optional. Specified when applied to data stored on the resource (eg: rows, columns in BigQuery Tables).
        """
        return pulumi.get(self, "data_access_spec")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Optional. Description of the DataAttribute.
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
        This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Optional. User-defined labels for the DataAttribute.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The relative resource name of the dataAttribute, of the form: projects/{project_number}/locations/{location_id}/dataTaxonomies/{dataTaxonomy}/attributes/{data_attribute_id}.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="parentId")
    def parent_id(self) -> str:
        """
        Optional. The ID of the parent DataAttribute resource, should belong to the same data taxonomy. Circular dependency in parent chain is not valid. Maximum depth of the hierarchy allowed is 4. a -> b -> c -> d -> e, depth = 4
        """
        return pulumi.get(self, "parent_id")

    @property
    @pulumi.getter(name="resourceAccessSpec")
    def resource_access_spec(self) -> 'outputs.GoogleCloudDataplexV1ResourceAccessSpecResponse':
        """
        Optional. Specified when applied to a resource (eg: Cloud Storage bucket, BigQuery dataset, BigQuery table).
        """
        return pulumi.get(self, "resource_access_spec")

    @property
    @pulumi.getter
    def uid(self) -> str:
        """
        System generated globally unique ID for the DataAttribute. This ID will be different if the DataAttribute is deleted and re-created with the same name.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The time when the DataAttribute was last updated.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetAttributeResult(GetAttributeResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetAttributeResult(
            attribute_count=self.attribute_count,
            create_time=self.create_time,
            data_access_spec=self.data_access_spec,
            description=self.description,
            display_name=self.display_name,
            etag=self.etag,
            labels=self.labels,
            name=self.name,
            parent_id=self.parent_id,
            resource_access_spec=self.resource_access_spec,
            uid=self.uid,
            update_time=self.update_time)


def get_attribute(attribute_id: Optional[str] = None,
                  data_taxonomy_id: Optional[str] = None,
                  location: Optional[str] = None,
                  project: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetAttributeResult:
    """
    Retrieves a Data Attribute resource.
    """
    __args__ = dict()
    __args__['attributeId'] = attribute_id
    __args__['dataTaxonomyId'] = data_taxonomy_id
    __args__['location'] = location
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:dataplex/v1:getAttribute', __args__, opts=opts, typ=GetAttributeResult).value

    return AwaitableGetAttributeResult(
        attribute_count=__ret__.attribute_count,
        create_time=__ret__.create_time,
        data_access_spec=__ret__.data_access_spec,
        description=__ret__.description,
        display_name=__ret__.display_name,
        etag=__ret__.etag,
        labels=__ret__.labels,
        name=__ret__.name,
        parent_id=__ret__.parent_id,
        resource_access_spec=__ret__.resource_access_spec,
        uid=__ret__.uid,
        update_time=__ret__.update_time)


@_utilities.lift_output_func(get_attribute)
def get_attribute_output(attribute_id: Optional[pulumi.Input[str]] = None,
                         data_taxonomy_id: Optional[pulumi.Input[str]] = None,
                         location: Optional[pulumi.Input[str]] = None,
                         project: Optional[pulumi.Input[Optional[str]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetAttributeResult]:
    """
    Retrieves a Data Attribute resource.
    """
    ...