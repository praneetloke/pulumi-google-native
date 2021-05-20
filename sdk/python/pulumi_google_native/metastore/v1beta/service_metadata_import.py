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

__all__ = ['ServiceMetadataImportArgs', 'ServiceMetadataImport']

@pulumi.input_type
class ServiceMetadataImportArgs:
    def __init__(__self__, *,
                 location: pulumi.Input[str],
                 metadata_import_id: pulumi.Input[str],
                 project: pulumi.Input[str],
                 service_id: pulumi.Input[str],
                 database_dump: Optional[pulumi.Input['DatabaseDumpArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a ServiceMetadataImport resource.
        :param pulumi.Input['DatabaseDumpArgs'] database_dump: Immutable. A database dump from a pre-existing metastore's database.
        :param pulumi.Input[str] description: The description of the metadata import.
        :param pulumi.Input[str] name: Immutable. The relative resource name of the metadata import, of the form:projects/{project_number}/locations/{location_id}/services/{service_id}/metadataImports/{metadata_import_id}.
        """
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "metadata_import_id", metadata_import_id)
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "service_id", service_id)
        if database_dump is not None:
            pulumi.set(__self__, "database_dump", database_dump)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if request_id is not None:
            pulumi.set(__self__, "request_id", request_id)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Input[str]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: pulumi.Input[str]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="metadataImportId")
    def metadata_import_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "metadata_import_id")

    @metadata_import_id.setter
    def metadata_import_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "metadata_import_id", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="serviceId")
    def service_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "service_id")

    @service_id.setter
    def service_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "service_id", value)

    @property
    @pulumi.getter(name="databaseDump")
    def database_dump(self) -> Optional[pulumi.Input['DatabaseDumpArgs']]:
        """
        Immutable. A database dump from a pre-existing metastore's database.
        """
        return pulumi.get(self, "database_dump")

    @database_dump.setter
    def database_dump(self, value: Optional[pulumi.Input['DatabaseDumpArgs']]):
        pulumi.set(self, "database_dump", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the metadata import.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Immutable. The relative resource name of the metadata import, of the form:projects/{project_number}/locations/{location_id}/services/{service_id}/metadataImports/{metadata_import_id}.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "request_id")

    @request_id.setter
    def request_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_id", value)


class ServiceMetadataImport(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database_dump: Optional[pulumi.Input[pulumi.InputType['DatabaseDumpArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 metadata_import_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a new MetadataImport in a given project and location.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['DatabaseDumpArgs']] database_dump: Immutable. A database dump from a pre-existing metastore's database.
        :param pulumi.Input[str] description: The description of the metadata import.
        :param pulumi.Input[str] name: Immutable. The relative resource name of the metadata import, of the form:projects/{project_number}/locations/{location_id}/services/{service_id}/metadataImports/{metadata_import_id}.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ServiceMetadataImportArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new MetadataImport in a given project and location.

        :param str resource_name: The name of the resource.
        :param ServiceMetadataImportArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ServiceMetadataImportArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 database_dump: Optional[pulumi.Input[pulumi.InputType['DatabaseDumpArgs']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 metadata_import_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 service_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = ServiceMetadataImportArgs.__new__(ServiceMetadataImportArgs)

            __props__.__dict__["database_dump"] = database_dump
            __props__.__dict__["description"] = description
            if location is None and not opts.urn:
                raise TypeError("Missing required property 'location'")
            __props__.__dict__["location"] = location
            if metadata_import_id is None and not opts.urn:
                raise TypeError("Missing required property 'metadata_import_id'")
            __props__.__dict__["metadata_import_id"] = metadata_import_id
            __props__.__dict__["name"] = name
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            __props__.__dict__["request_id"] = request_id
            if service_id is None and not opts.urn:
                raise TypeError("Missing required property 'service_id'")
            __props__.__dict__["service_id"] = service_id
            __props__.__dict__["create_time"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["update_time"] = None
        super(ServiceMetadataImport, __self__).__init__(
            'google-native:metastore/v1beta:ServiceMetadataImport',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'ServiceMetadataImport':
        """
        Get an existing ServiceMetadataImport resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ServiceMetadataImportArgs.__new__(ServiceMetadataImportArgs)

        __props__.__dict__["create_time"] = None
        __props__.__dict__["database_dump"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["update_time"] = None
        return ServiceMetadataImport(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The time when the metadata import was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="databaseDump")
    def database_dump(self) -> pulumi.Output['outputs.DatabaseDumpResponse']:
        """
        Immutable. A database dump from a pre-existing metastore's database.
        """
        return pulumi.get(self, "database_dump")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The description of the metadata import.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Immutable. The relative resource name of the metadata import, of the form:projects/{project_number}/locations/{location_id}/services/{service_id}/metadataImports/{metadata_import_id}.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The current state of the metadata import.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The time when the metadata import was last updated.
        """
        return pulumi.get(self, "update_time")

