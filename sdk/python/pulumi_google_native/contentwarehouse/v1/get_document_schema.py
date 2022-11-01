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
    'GetDocumentSchemaResult',
    'AwaitableGetDocumentSchemaResult',
    'get_document_schema',
    'get_document_schema_output',
]

@pulumi.output_type
class GetDocumentSchemaResult:
    def __init__(__self__, create_time=None, description=None, display_name=None, document_is_folder=None, name=None, property_definitions=None, update_time=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if document_is_folder and not isinstance(document_is_folder, bool):
            raise TypeError("Expected argument 'document_is_folder' to be a bool")
        pulumi.set(__self__, "document_is_folder", document_is_folder)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if property_definitions and not isinstance(property_definitions, list):
            raise TypeError("Expected argument 'property_definitions' to be a list")
        pulumi.set(__self__, "property_definitions", property_definitions)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The time when the document schema is created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Schema description.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Name of the schema given by the user. Must be unique per customer.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="documentIsFolder")
    def document_is_folder(self) -> bool:
        """
        Document Type, true refers the document is a folder, otherwise it is a typical document.
        """
        return pulumi.get(self, "document_is_folder")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name of the document schema. Format: projects/{project_number}/locations/{location}/documentSchemas/{document_schema_id}. The name is ignored when creating a document schema.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="propertyDefinitions")
    def property_definitions(self) -> Sequence['outputs.GoogleCloudContentwarehouseV1PropertyDefinitionResponse']:
        """
        Document details.
        """
        return pulumi.get(self, "property_definitions")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The time when the document schema is last updated.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetDocumentSchemaResult(GetDocumentSchemaResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDocumentSchemaResult(
            create_time=self.create_time,
            description=self.description,
            display_name=self.display_name,
            document_is_folder=self.document_is_folder,
            name=self.name,
            property_definitions=self.property_definitions,
            update_time=self.update_time)


def get_document_schema(document_schema_id: Optional[str] = None,
                        location: Optional[str] = None,
                        project: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDocumentSchemaResult:
    """
    Gets a document schema. Returns NOT_FOUND if the document schema does not exist.
    """
    __args__ = dict()
    __args__['documentSchemaId'] = document_schema_id
    __args__['location'] = location
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:contentwarehouse/v1:getDocumentSchema', __args__, opts=opts, typ=GetDocumentSchemaResult).value

    return AwaitableGetDocumentSchemaResult(
        create_time=__ret__.create_time,
        description=__ret__.description,
        display_name=__ret__.display_name,
        document_is_folder=__ret__.document_is_folder,
        name=__ret__.name,
        property_definitions=__ret__.property_definitions,
        update_time=__ret__.update_time)


@_utilities.lift_output_func(get_document_schema)
def get_document_schema_output(document_schema_id: Optional[pulumi.Input[str]] = None,
                               location: Optional[pulumi.Input[str]] = None,
                               project: Optional[pulumi.Input[Optional[str]]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDocumentSchemaResult]:
    """
    Gets a document schema. Returns NOT_FOUND if the document schema does not exist.
    """
    ...