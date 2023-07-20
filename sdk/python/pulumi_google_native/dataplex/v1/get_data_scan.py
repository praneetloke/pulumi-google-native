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
    'GetDataScanResult',
    'AwaitableGetDataScanResult',
    'get_data_scan',
    'get_data_scan_output',
]

@pulumi.output_type
class GetDataScanResult:
    def __init__(__self__, create_time=None, data=None, data_profile_result=None, data_profile_spec=None, data_quality_result=None, data_quality_spec=None, description=None, display_name=None, execution_spec=None, execution_status=None, labels=None, name=None, state=None, type=None, uid=None, update_time=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if data and not isinstance(data, dict):
            raise TypeError("Expected argument 'data' to be a dict")
        pulumi.set(__self__, "data", data)
        if data_profile_result and not isinstance(data_profile_result, dict):
            raise TypeError("Expected argument 'data_profile_result' to be a dict")
        pulumi.set(__self__, "data_profile_result", data_profile_result)
        if data_profile_spec and not isinstance(data_profile_spec, dict):
            raise TypeError("Expected argument 'data_profile_spec' to be a dict")
        pulumi.set(__self__, "data_profile_spec", data_profile_spec)
        if data_quality_result and not isinstance(data_quality_result, dict):
            raise TypeError("Expected argument 'data_quality_result' to be a dict")
        pulumi.set(__self__, "data_quality_result", data_quality_result)
        if data_quality_spec and not isinstance(data_quality_spec, dict):
            raise TypeError("Expected argument 'data_quality_spec' to be a dict")
        pulumi.set(__self__, "data_quality_spec", data_quality_spec)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if execution_spec and not isinstance(execution_spec, dict):
            raise TypeError("Expected argument 'execution_spec' to be a dict")
        pulumi.set(__self__, "execution_spec", execution_spec)
        if execution_status and not isinstance(execution_status, dict):
            raise TypeError("Expected argument 'execution_status' to be a dict")
        pulumi.set(__self__, "execution_status", execution_status)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if uid and not isinstance(uid, str):
            raise TypeError("Expected argument 'uid' to be a str")
        pulumi.set(__self__, "uid", uid)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The time when the scan was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def data(self) -> 'outputs.GoogleCloudDataplexV1DataSourceResponse':
        """
        The data source for DataScan.
        """
        return pulumi.get(self, "data")

    @property
    @pulumi.getter(name="dataProfileResult")
    def data_profile_result(self) -> 'outputs.GoogleCloudDataplexV1DataProfileResultResponse':
        """
        The result of the data profile scan.
        """
        return pulumi.get(self, "data_profile_result")

    @property
    @pulumi.getter(name="dataProfileSpec")
    def data_profile_spec(self) -> 'outputs.GoogleCloudDataplexV1DataProfileSpecResponse':
        """
        DataProfileScan related setting.
        """
        return pulumi.get(self, "data_profile_spec")

    @property
    @pulumi.getter(name="dataQualityResult")
    def data_quality_result(self) -> 'outputs.GoogleCloudDataplexV1DataQualityResultResponse':
        """
        The result of the data quality scan.
        """
        return pulumi.get(self, "data_quality_result")

    @property
    @pulumi.getter(name="dataQualitySpec")
    def data_quality_spec(self) -> 'outputs.GoogleCloudDataplexV1DataQualitySpecResponse':
        """
        DataQualityScan related setting.
        """
        return pulumi.get(self, "data_quality_spec")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Optional. Description of the scan. Must be between 1-1024 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Optional. User friendly display name. Must be between 1-256 characters.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="executionSpec")
    def execution_spec(self) -> 'outputs.GoogleCloudDataplexV1DataScanExecutionSpecResponse':
        """
        Optional. DataScan execution settings.If not specified, the fields in it will use their default values.
        """
        return pulumi.get(self, "execution_spec")

    @property
    @pulumi.getter(name="executionStatus")
    def execution_status(self) -> 'outputs.GoogleCloudDataplexV1DataScanExecutionStatusResponse':
        """
        Status of the data scan execution.
        """
        return pulumi.get(self, "execution_status")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Optional. User-defined labels for the scan.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The relative resource name of the scan, of the form: projects/{project}/locations/{location_id}/dataScans/{datascan_id}, where project refers to a project_id or project_number and location_id refers to a GCP region.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        Current state of the DataScan.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of DataScan.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def uid(self) -> str:
        """
        System generated globally unique ID for the scan. This ID will be different if the scan is deleted and re-created with the same name.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The time when the scan was last updated.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetDataScanResult(GetDataScanResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDataScanResult(
            create_time=self.create_time,
            data=self.data,
            data_profile_result=self.data_profile_result,
            data_profile_spec=self.data_profile_spec,
            data_quality_result=self.data_quality_result,
            data_quality_spec=self.data_quality_spec,
            description=self.description,
            display_name=self.display_name,
            execution_spec=self.execution_spec,
            execution_status=self.execution_status,
            labels=self.labels,
            name=self.name,
            state=self.state,
            type=self.type,
            uid=self.uid,
            update_time=self.update_time)


def get_data_scan(data_scan_id: Optional[str] = None,
                  location: Optional[str] = None,
                  project: Optional[str] = None,
                  view: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDataScanResult:
    """
    Gets a DataScan resource.
    """
    __args__ = dict()
    __args__['dataScanId'] = data_scan_id
    __args__['location'] = location
    __args__['project'] = project
    __args__['view'] = view
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:dataplex/v1:getDataScan', __args__, opts=opts, typ=GetDataScanResult).value

    return AwaitableGetDataScanResult(
        create_time=pulumi.get(__ret__, 'create_time'),
        data=pulumi.get(__ret__, 'data'),
        data_profile_result=pulumi.get(__ret__, 'data_profile_result'),
        data_profile_spec=pulumi.get(__ret__, 'data_profile_spec'),
        data_quality_result=pulumi.get(__ret__, 'data_quality_result'),
        data_quality_spec=pulumi.get(__ret__, 'data_quality_spec'),
        description=pulumi.get(__ret__, 'description'),
        display_name=pulumi.get(__ret__, 'display_name'),
        execution_spec=pulumi.get(__ret__, 'execution_spec'),
        execution_status=pulumi.get(__ret__, 'execution_status'),
        labels=pulumi.get(__ret__, 'labels'),
        name=pulumi.get(__ret__, 'name'),
        state=pulumi.get(__ret__, 'state'),
        type=pulumi.get(__ret__, 'type'),
        uid=pulumi.get(__ret__, 'uid'),
        update_time=pulumi.get(__ret__, 'update_time'))


@_utilities.lift_output_func(get_data_scan)
def get_data_scan_output(data_scan_id: Optional[pulumi.Input[str]] = None,
                         location: Optional[pulumi.Input[str]] = None,
                         project: Optional[pulumi.Input[Optional[str]]] = None,
                         view: Optional[pulumi.Input[Optional[str]]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDataScanResult]:
    """
    Gets a DataScan resource.
    """
    ...
