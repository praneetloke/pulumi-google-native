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
    'GetPartitionResult',
    'AwaitableGetPartitionResult',
    'get_partition',
    'get_partition_output',
]

@pulumi.output_type
class GetPartitionResult:
    def __init__(__self__, etag=None, location=None, name=None, values=None):
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if values and not isinstance(values, list):
            raise TypeError("Expected argument 'values' to be a list")
        pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        Optional. The etag for this partition.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        Immutable. The location of the entity data within the partition, for example, gs://bucket/path/to/entity/key1=value1/key2=value2. Or projects//datasets//tables/
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Partition values used in the HTTP URL must be double encoded. For example, url_encode(url_encode(value)) can be used to encode "US:CA/CA#Sunnyvale so that the request URL ends with "/partitions/US%253ACA/CA%2523Sunnyvale". The name field in the response retains the encoded format.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def values(self) -> Sequence[str]:
        """
        Immutable. The set of values representing the partition, which correspond to the partition schema defined in the parent entity.
        """
        return pulumi.get(self, "values")


class AwaitableGetPartitionResult(GetPartitionResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetPartitionResult(
            etag=self.etag,
            location=self.location,
            name=self.name,
            values=self.values)


def get_partition(entity_id: Optional[str] = None,
                  lake_id: Optional[str] = None,
                  location: Optional[str] = None,
                  partition_id: Optional[str] = None,
                  project: Optional[str] = None,
                  zone: Optional[str] = None,
                  opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetPartitionResult:
    """
    Get a metadata partition of an entity.
    """
    __args__ = dict()
    __args__['entityId'] = entity_id
    __args__['lakeId'] = lake_id
    __args__['location'] = location
    __args__['partitionId'] = partition_id
    __args__['project'] = project
    __args__['zone'] = zone
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:dataplex/v1:getPartition', __args__, opts=opts, typ=GetPartitionResult).value

    return AwaitableGetPartitionResult(
        etag=pulumi.get(__ret__, 'etag'),
        location=pulumi.get(__ret__, 'location'),
        name=pulumi.get(__ret__, 'name'),
        values=pulumi.get(__ret__, 'values'))


@_utilities.lift_output_func(get_partition)
def get_partition_output(entity_id: Optional[pulumi.Input[str]] = None,
                         lake_id: Optional[pulumi.Input[str]] = None,
                         location: Optional[pulumi.Input[str]] = None,
                         partition_id: Optional[pulumi.Input[str]] = None,
                         project: Optional[pulumi.Input[Optional[str]]] = None,
                         zone: Optional[pulumi.Input[str]] = None,
                         opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetPartitionResult]:
    """
    Get a metadata partition of an entity.
    """
    ...
