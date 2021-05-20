# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs
from ._inputs import *

__all__ = ['DatabaseCollectionGroupIndexArgs', 'DatabaseCollectionGroupIndex']

@pulumi.input_type
class DatabaseCollectionGroupIndexArgs:
    def __init__(__self__, *,
                 collection_group_id: pulumi.Input[str],
                 database_id: pulumi.Input[str],
                 index_id: pulumi.Input[str],
                 project: pulumi.Input[str],
                 fields: Optional[pulumi.Input[Sequence[pulumi.Input['GoogleFirestoreAdminV1beta2IndexFieldArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 query_scope: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DatabaseCollectionGroupIndex resource.
        :param pulumi.Input[Sequence[pulumi.Input['GoogleFirestoreAdminV1beta2IndexFieldArgs']]] fields: The fields supported by this index. For composite indexes, this is always 2 or more fields. The last field entry is always for the field path `__name__`. If, on creation, `__name__` was not specified as the last field, it will be added automatically with the same direction as that of the last field defined. If the final field in a composite index is not directional, the `__name__` will be ordered ASCENDING (unless explicitly specified). For single field indexes, this will always be exactly one entry with a field path equal to the field path of the associated field.
        :param pulumi.Input[str] name: A server defined name for this index. The form of this name for composite indexes will be: `projects/{project_id}/databases/{database_id}/collectionGroups/{collection_id}/indexes/{composite_index_id}` For single field indexes, this field will be empty.
        :param pulumi.Input[str] query_scope: Indexes with a collection query scope specified allow queries against a collection that is the child of a specific document, specified at query time, and that has the same collection id. Indexes with a collection group query scope specified allow queries against all collections descended from a specific document, specified at query time, and that have the same collection id as this index.
        :param pulumi.Input[str] state: The serving state of the index.
        """
        pulumi.set(__self__, "collection_group_id", collection_group_id)
        pulumi.set(__self__, "database_id", database_id)
        pulumi.set(__self__, "index_id", index_id)
        pulumi.set(__self__, "project", project)
        if fields is not None:
            pulumi.set(__self__, "fields", fields)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if query_scope is not None:
            pulumi.set(__self__, "query_scope", query_scope)
        if state is not None:
            pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter(name="collectionGroupId")
    def collection_group_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "collection_group_id")

    @collection_group_id.setter
    def collection_group_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "collection_group_id", value)

    @property
    @pulumi.getter(name="databaseId")
    def database_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "database_id")

    @database_id.setter
    def database_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "database_id", value)

    @property
    @pulumi.getter(name="indexId")
    def index_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "index_id")

    @index_id.setter
    def index_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "index_id", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def fields(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GoogleFirestoreAdminV1beta2IndexFieldArgs']]]]:
        """
        The fields supported by this index. For composite indexes, this is always 2 or more fields. The last field entry is always for the field path `__name__`. If, on creation, `__name__` was not specified as the last field, it will be added automatically with the same direction as that of the last field defined. If the final field in a composite index is not directional, the `__name__` will be ordered ASCENDING (unless explicitly specified). For single field indexes, this will always be exactly one entry with a field path equal to the field path of the associated field.
        """
        return pulumi.get(self, "fields")

    @fields.setter
    def fields(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GoogleFirestoreAdminV1beta2IndexFieldArgs']]]]):
        pulumi.set(self, "fields", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A server defined name for this index. The form of this name for composite indexes will be: `projects/{project_id}/databases/{database_id}/collectionGroups/{collection_id}/indexes/{composite_index_id}` For single field indexes, this field will be empty.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="queryScope")
    def query_scope(self) -> Optional[pulumi.Input[str]]:
        """
        Indexes with a collection query scope specified allow queries against a collection that is the child of a specific document, specified at query time, and that has the same collection id. Indexes with a collection group query scope specified allow queries against all collections descended from a specific document, specified at query time, and that have the same collection id as this index.
        """
        return pulumi.get(self, "query_scope")

    @query_scope.setter
    def query_scope(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "query_scope", value)

    @property
    @pulumi.getter
    def state(self) -> Optional[pulumi.Input[str]]:
        """
        The serving state of the index.
        """
        return pulumi.get(self, "state")

    @state.setter
    def state(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "state", value)


class DatabaseCollectionGroupIndex(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 collection_group_id: Optional[pulumi.Input[str]] = None,
                 database_id: Optional[pulumi.Input[str]] = None,
                 fields: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleFirestoreAdminV1beta2IndexFieldArgs']]]]] = None,
                 index_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 query_scope: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a composite index. This returns a google.longrunning.Operation which may be used to track the status of the creation. The metadata for the operation will be the type IndexOperationMetadata.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleFirestoreAdminV1beta2IndexFieldArgs']]]] fields: The fields supported by this index. For composite indexes, this is always 2 or more fields. The last field entry is always for the field path `__name__`. If, on creation, `__name__` was not specified as the last field, it will be added automatically with the same direction as that of the last field defined. If the final field in a composite index is not directional, the `__name__` will be ordered ASCENDING (unless explicitly specified). For single field indexes, this will always be exactly one entry with a field path equal to the field path of the associated field.
        :param pulumi.Input[str] name: A server defined name for this index. The form of this name for composite indexes will be: `projects/{project_id}/databases/{database_id}/collectionGroups/{collection_id}/indexes/{composite_index_id}` For single field indexes, this field will be empty.
        :param pulumi.Input[str] query_scope: Indexes with a collection query scope specified allow queries against a collection that is the child of a specific document, specified at query time, and that has the same collection id. Indexes with a collection group query scope specified allow queries against all collections descended from a specific document, specified at query time, and that have the same collection id as this index.
        :param pulumi.Input[str] state: The serving state of the index.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DatabaseCollectionGroupIndexArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a composite index. This returns a google.longrunning.Operation which may be used to track the status of the creation. The metadata for the operation will be the type IndexOperationMetadata.

        :param str resource_name: The name of the resource.
        :param DatabaseCollectionGroupIndexArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DatabaseCollectionGroupIndexArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 collection_group_id: Optional[pulumi.Input[str]] = None,
                 database_id: Optional[pulumi.Input[str]] = None,
                 fields: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleFirestoreAdminV1beta2IndexFieldArgs']]]]] = None,
                 index_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 query_scope: Optional[pulumi.Input[str]] = None,
                 state: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DatabaseCollectionGroupIndexArgs.__new__(DatabaseCollectionGroupIndexArgs)

            if collection_group_id is None and not opts.urn:
                raise TypeError("Missing required property 'collection_group_id'")
            __props__.__dict__["collection_group_id"] = collection_group_id
            if database_id is None and not opts.urn:
                raise TypeError("Missing required property 'database_id'")
            __props__.__dict__["database_id"] = database_id
            __props__.__dict__["fields"] = fields
            if index_id is None and not opts.urn:
                raise TypeError("Missing required property 'index_id'")
            __props__.__dict__["index_id"] = index_id
            __props__.__dict__["name"] = name
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            __props__.__dict__["query_scope"] = query_scope
            __props__.__dict__["state"] = state
        super(DatabaseCollectionGroupIndex, __self__).__init__(
            'google-native:firestore/v1beta2:DatabaseCollectionGroupIndex',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DatabaseCollectionGroupIndex':
        """
        Get an existing DatabaseCollectionGroupIndex resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DatabaseCollectionGroupIndexArgs.__new__(DatabaseCollectionGroupIndexArgs)

        __props__.__dict__["fields"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["query_scope"] = None
        __props__.__dict__["state"] = None
        return DatabaseCollectionGroupIndex(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def fields(self) -> pulumi.Output[Sequence['outputs.GoogleFirestoreAdminV1beta2IndexFieldResponse']]:
        """
        The fields supported by this index. For composite indexes, this is always 2 or more fields. The last field entry is always for the field path `__name__`. If, on creation, `__name__` was not specified as the last field, it will be added automatically with the same direction as that of the last field defined. If the final field in a composite index is not directional, the `__name__` will be ordered ASCENDING (unless explicitly specified). For single field indexes, this will always be exactly one entry with a field path equal to the field path of the associated field.
        """
        return pulumi.get(self, "fields")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A server defined name for this index. The form of this name for composite indexes will be: `projects/{project_id}/databases/{database_id}/collectionGroups/{collection_id}/indexes/{composite_index_id}` For single field indexes, this field will be empty.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="queryScope")
    def query_scope(self) -> pulumi.Output[str]:
        """
        Indexes with a collection query scope specified allow queries against a collection that is the child of a specific document, specified at query time, and that has the same collection id. Indexes with a collection group query scope specified allow queries against all collections descended from a specific document, specified at query time, and that have the same collection id as this index.
        """
        return pulumi.get(self, "query_scope")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The serving state of the index.
        """
        return pulumi.get(self, "state")

