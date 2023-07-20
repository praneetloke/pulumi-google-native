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
    'GetDocumentResult',
    'AwaitableGetDocumentResult',
    'get_document',
    'get_document_output',
]

@pulumi.output_type
class GetDocumentResult:
    def __init__(__self__, json_data=None, name=None, parent_document_id=None, schema_id=None, struct_data=None):
        if json_data and not isinstance(json_data, str):
            raise TypeError("Expected argument 'json_data' to be a str")
        pulumi.set(__self__, "json_data", json_data)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if parent_document_id and not isinstance(parent_document_id, str):
            raise TypeError("Expected argument 'parent_document_id' to be a str")
        pulumi.set(__self__, "parent_document_id", parent_document_id)
        if schema_id and not isinstance(schema_id, str):
            raise TypeError("Expected argument 'schema_id' to be a str")
        pulumi.set(__self__, "schema_id", schema_id)
        if struct_data and not isinstance(struct_data, dict):
            raise TypeError("Expected argument 'struct_data' to be a dict")
        pulumi.set(__self__, "struct_data", struct_data)

    @property
    @pulumi.getter(name="jsonData")
    def json_data(self) -> str:
        """
        The JSON string representation of the document. It should conform to the registered Schema.schema or an `INVALID_ARGUMENT` error is thrown.
        """
        return pulumi.get(self, "json_data")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Immutable. The full resource name of the document. Format: `projects/{project}/locations/{location}/collections/{collection}/dataStores/{data_store}/branches/{branch}/documents/{document_id}`. This field must be a UTF-8 encoded string with a length limit of 1024 characters.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="parentDocumentId")
    def parent_document_id(self) -> str:
        """
        The identifier of the parent document. Currently supports at most two level document hierarchy. Id should conform to [RFC-1034](https://tools.ietf.org/html/rfc1034) standard with a length limit of 63 characters.
        """
        return pulumi.get(self, "parent_document_id")

    @property
    @pulumi.getter(name="schemaId")
    def schema_id(self) -> str:
        """
        The identifier of the schema located in the same data store.
        """
        return pulumi.get(self, "schema_id")

    @property
    @pulumi.getter(name="structData")
    def struct_data(self) -> Mapping[str, str]:
        """
        The structured JSON data for the document. It should conform to the registered Schema.schema or an `INVALID_ARGUMENT` error is thrown.
        """
        return pulumi.get(self, "struct_data")


class AwaitableGetDocumentResult(GetDocumentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDocumentResult(
            json_data=self.json_data,
            name=self.name,
            parent_document_id=self.parent_document_id,
            schema_id=self.schema_id,
            struct_data=self.struct_data)


def get_document(branch_id: Optional[str] = None,
                 collection_id: Optional[str] = None,
                 data_store_id: Optional[str] = None,
                 document_id: Optional[str] = None,
                 location: Optional[str] = None,
                 project: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDocumentResult:
    """
    Gets a Document.
    """
    __args__ = dict()
    __args__['branchId'] = branch_id
    __args__['collectionId'] = collection_id
    __args__['dataStoreId'] = data_store_id
    __args__['documentId'] = document_id
    __args__['location'] = location
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:discoveryengine/v1beta:getDocument', __args__, opts=opts, typ=GetDocumentResult).value

    return AwaitableGetDocumentResult(
        json_data=pulumi.get(__ret__, 'json_data'),
        name=pulumi.get(__ret__, 'name'),
        parent_document_id=pulumi.get(__ret__, 'parent_document_id'),
        schema_id=pulumi.get(__ret__, 'schema_id'),
        struct_data=pulumi.get(__ret__, 'struct_data'))


@_utilities.lift_output_func(get_document)
def get_document_output(branch_id: Optional[pulumi.Input[str]] = None,
                        collection_id: Optional[pulumi.Input[str]] = None,
                        data_store_id: Optional[pulumi.Input[str]] = None,
                        document_id: Optional[pulumi.Input[str]] = None,
                        location: Optional[pulumi.Input[str]] = None,
                        project: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDocumentResult]:
    """
    Gets a Document.
    """
    ...
