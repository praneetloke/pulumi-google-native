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
from ._inputs import *

__all__ = ['DataScanArgs', 'DataScan']

@pulumi.input_type
class DataScanArgs:
    def __init__(__self__, *,
                 data: pulumi.Input['GoogleCloudDataplexV1DataSourceArgs'],
                 data_scan_id: pulumi.Input[str],
                 data_profile_spec: Optional[pulumi.Input['GoogleCloudDataplexV1DataProfileSpecArgs']] = None,
                 data_quality_spec: Optional[pulumi.Input['GoogleCloudDataplexV1DataQualitySpecArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 execution_spec: Optional[pulumi.Input['GoogleCloudDataplexV1DataScanExecutionSpecArgs']] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a DataScan resource.
        :param pulumi.Input['GoogleCloudDataplexV1DataSourceArgs'] data: The data source for DataScan.
        :param pulumi.Input[str] data_scan_id: Required. DataScan identifier. Must contain only lowercase letters, numbers and hyphens. Must start with a letter. Must end with a number or a letter. Must be between 1-63 characters. Must be unique within the customer project / location.
        :param pulumi.Input['GoogleCloudDataplexV1DataProfileSpecArgs'] data_profile_spec: DataProfileScan related setting.
        :param pulumi.Input['GoogleCloudDataplexV1DataQualitySpecArgs'] data_quality_spec: DataQualityScan related setting.
        :param pulumi.Input[str] description: Optional. Description of the scan. Must be between 1-1024 characters.
        :param pulumi.Input[str] display_name: Optional. User friendly display name. Must be between 1-256 characters.
        :param pulumi.Input['GoogleCloudDataplexV1DataScanExecutionSpecArgs'] execution_spec: Optional. DataScan execution settings.If not specified, the fields in it will use their default values.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Optional. User-defined labels for the scan.
        """
        pulumi.set(__self__, "data", data)
        pulumi.set(__self__, "data_scan_id", data_scan_id)
        if data_profile_spec is not None:
            pulumi.set(__self__, "data_profile_spec", data_profile_spec)
        if data_quality_spec is not None:
            pulumi.set(__self__, "data_quality_spec", data_quality_spec)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if execution_spec is not None:
            pulumi.set(__self__, "execution_spec", execution_spec)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter
    def data(self) -> pulumi.Input['GoogleCloudDataplexV1DataSourceArgs']:
        """
        The data source for DataScan.
        """
        return pulumi.get(self, "data")

    @data.setter
    def data(self, value: pulumi.Input['GoogleCloudDataplexV1DataSourceArgs']):
        pulumi.set(self, "data", value)

    @property
    @pulumi.getter(name="dataScanId")
    def data_scan_id(self) -> pulumi.Input[str]:
        """
        Required. DataScan identifier. Must contain only lowercase letters, numbers and hyphens. Must start with a letter. Must end with a number or a letter. Must be between 1-63 characters. Must be unique within the customer project / location.
        """
        return pulumi.get(self, "data_scan_id")

    @data_scan_id.setter
    def data_scan_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "data_scan_id", value)

    @property
    @pulumi.getter(name="dataProfileSpec")
    def data_profile_spec(self) -> Optional[pulumi.Input['GoogleCloudDataplexV1DataProfileSpecArgs']]:
        """
        DataProfileScan related setting.
        """
        return pulumi.get(self, "data_profile_spec")

    @data_profile_spec.setter
    def data_profile_spec(self, value: Optional[pulumi.Input['GoogleCloudDataplexV1DataProfileSpecArgs']]):
        pulumi.set(self, "data_profile_spec", value)

    @property
    @pulumi.getter(name="dataQualitySpec")
    def data_quality_spec(self) -> Optional[pulumi.Input['GoogleCloudDataplexV1DataQualitySpecArgs']]:
        """
        DataQualityScan related setting.
        """
        return pulumi.get(self, "data_quality_spec")

    @data_quality_spec.setter
    def data_quality_spec(self, value: Optional[pulumi.Input['GoogleCloudDataplexV1DataQualitySpecArgs']]):
        pulumi.set(self, "data_quality_spec", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Description of the scan. Must be between 1-1024 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. User friendly display name. Must be between 1-256 characters.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="executionSpec")
    def execution_spec(self) -> Optional[pulumi.Input['GoogleCloudDataplexV1DataScanExecutionSpecArgs']]:
        """
        Optional. DataScan execution settings.If not specified, the fields in it will use their default values.
        """
        return pulumi.get(self, "execution_spec")

    @execution_spec.setter
    def execution_spec(self, value: Optional[pulumi.Input['GoogleCloudDataplexV1DataScanExecutionSpecArgs']]):
        pulumi.set(self, "execution_spec", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Optional. User-defined labels for the scan.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)


class DataScan(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDataplexV1DataSourceArgs']]] = None,
                 data_profile_spec: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDataplexV1DataProfileSpecArgs']]] = None,
                 data_quality_spec: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDataplexV1DataQualitySpecArgs']]] = None,
                 data_scan_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 execution_spec: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDataplexV1DataScanExecutionSpecArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a DataScan resource.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['GoogleCloudDataplexV1DataSourceArgs']] data: The data source for DataScan.
        :param pulumi.Input[pulumi.InputType['GoogleCloudDataplexV1DataProfileSpecArgs']] data_profile_spec: DataProfileScan related setting.
        :param pulumi.Input[pulumi.InputType['GoogleCloudDataplexV1DataQualitySpecArgs']] data_quality_spec: DataQualityScan related setting.
        :param pulumi.Input[str] data_scan_id: Required. DataScan identifier. Must contain only lowercase letters, numbers and hyphens. Must start with a letter. Must end with a number or a letter. Must be between 1-63 characters. Must be unique within the customer project / location.
        :param pulumi.Input[str] description: Optional. Description of the scan. Must be between 1-1024 characters.
        :param pulumi.Input[str] display_name: Optional. User friendly display name. Must be between 1-256 characters.
        :param pulumi.Input[pulumi.InputType['GoogleCloudDataplexV1DataScanExecutionSpecArgs']] execution_spec: Optional. DataScan execution settings.If not specified, the fields in it will use their default values.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Optional. User-defined labels for the scan.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DataScanArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a DataScan resource.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param DataScanArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DataScanArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 data: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDataplexV1DataSourceArgs']]] = None,
                 data_profile_spec: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDataplexV1DataProfileSpecArgs']]] = None,
                 data_quality_spec: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDataplexV1DataQualitySpecArgs']]] = None,
                 data_scan_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 execution_spec: Optional[pulumi.Input[pulumi.InputType['GoogleCloudDataplexV1DataScanExecutionSpecArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = DataScanArgs.__new__(DataScanArgs)

            if data is None and not opts.urn:
                raise TypeError("Missing required property 'data'")
            __props__.__dict__["data"] = data
            __props__.__dict__["data_profile_spec"] = data_profile_spec
            __props__.__dict__["data_quality_spec"] = data_quality_spec
            if data_scan_id is None and not opts.urn:
                raise TypeError("Missing required property 'data_scan_id'")
            __props__.__dict__["data_scan_id"] = data_scan_id
            __props__.__dict__["description"] = description
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["execution_spec"] = execution_spec
            __props__.__dict__["labels"] = labels
            __props__.__dict__["location"] = location
            __props__.__dict__["project"] = project
            __props__.__dict__["create_time"] = None
            __props__.__dict__["data_profile_result"] = None
            __props__.__dict__["data_quality_result"] = None
            __props__.__dict__["execution_status"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["type"] = None
            __props__.__dict__["uid"] = None
            __props__.__dict__["update_time"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["data_scan_id", "location", "project"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(DataScan, __self__).__init__(
            'google-native:dataplex/v1:DataScan',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'DataScan':
        """
        Get an existing DataScan resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DataScanArgs.__new__(DataScanArgs)

        __props__.__dict__["create_time"] = None
        __props__.__dict__["data"] = None
        __props__.__dict__["data_profile_result"] = None
        __props__.__dict__["data_profile_spec"] = None
        __props__.__dict__["data_quality_result"] = None
        __props__.__dict__["data_quality_spec"] = None
        __props__.__dict__["data_scan_id"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["execution_spec"] = None
        __props__.__dict__["execution_status"] = None
        __props__.__dict__["labels"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["uid"] = None
        __props__.__dict__["update_time"] = None
        return DataScan(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The time when the scan was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def data(self) -> pulumi.Output['outputs.GoogleCloudDataplexV1DataSourceResponse']:
        """
        The data source for DataScan.
        """
        return pulumi.get(self, "data")

    @property
    @pulumi.getter(name="dataProfileResult")
    def data_profile_result(self) -> pulumi.Output['outputs.GoogleCloudDataplexV1DataProfileResultResponse']:
        """
        The result of the data profile scan.
        """
        return pulumi.get(self, "data_profile_result")

    @property
    @pulumi.getter(name="dataProfileSpec")
    def data_profile_spec(self) -> pulumi.Output['outputs.GoogleCloudDataplexV1DataProfileSpecResponse']:
        """
        DataProfileScan related setting.
        """
        return pulumi.get(self, "data_profile_spec")

    @property
    @pulumi.getter(name="dataQualityResult")
    def data_quality_result(self) -> pulumi.Output['outputs.GoogleCloudDataplexV1DataQualityResultResponse']:
        """
        The result of the data quality scan.
        """
        return pulumi.get(self, "data_quality_result")

    @property
    @pulumi.getter(name="dataQualitySpec")
    def data_quality_spec(self) -> pulumi.Output['outputs.GoogleCloudDataplexV1DataQualitySpecResponse']:
        """
        DataQualityScan related setting.
        """
        return pulumi.get(self, "data_quality_spec")

    @property
    @pulumi.getter(name="dataScanId")
    def data_scan_id(self) -> pulumi.Output[str]:
        """
        Required. DataScan identifier. Must contain only lowercase letters, numbers and hyphens. Must start with a letter. Must end with a number or a letter. Must be between 1-63 characters. Must be unique within the customer project / location.
        """
        return pulumi.get(self, "data_scan_id")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Optional. Description of the scan. Must be between 1-1024 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Optional. User friendly display name. Must be between 1-256 characters.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="executionSpec")
    def execution_spec(self) -> pulumi.Output['outputs.GoogleCloudDataplexV1DataScanExecutionSpecResponse']:
        """
        Optional. DataScan execution settings.If not specified, the fields in it will use their default values.
        """
        return pulumi.get(self, "execution_spec")

    @property
    @pulumi.getter(name="executionStatus")
    def execution_status(self) -> pulumi.Output['outputs.GoogleCloudDataplexV1DataScanExecutionStatusResponse']:
        """
        Status of the data scan execution.
        """
        return pulumi.get(self, "execution_status")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Optional. User-defined labels for the scan.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The relative resource name of the scan, of the form: projects/{project}/locations/{location_id}/dataScans/{datascan_id}, where project refers to a project_id or project_number and location_id refers to a GCP region.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        Current state of the DataScan.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        The type of DataScan.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def uid(self) -> pulumi.Output[str]:
        """
        System generated globally unique ID for the scan. This ID will be different if the scan is deleted and re-created with the same name.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The time when the scan was last updated.
        """
        return pulumi.get(self, "update_time")

