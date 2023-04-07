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
    'GetDatabaseResult',
    'AwaitableGetDatabaseResult',
    'get_database',
    'get_database_output',
]

@pulumi.output_type
class GetDatabaseResult:
    def __init__(__self__, app_engine_integration_mode=None, concurrency_mode=None, create_time=None, etag=None, key_prefix=None, location=None, name=None, type=None, uid=None, update_time=None):
        if app_engine_integration_mode and not isinstance(app_engine_integration_mode, str):
            raise TypeError("Expected argument 'app_engine_integration_mode' to be a str")
        pulumi.set(__self__, "app_engine_integration_mode", app_engine_integration_mode)
        if concurrency_mode and not isinstance(concurrency_mode, str):
            raise TypeError("Expected argument 'concurrency_mode' to be a str")
        pulumi.set(__self__, "concurrency_mode", concurrency_mode)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if key_prefix and not isinstance(key_prefix, str):
            raise TypeError("Expected argument 'key_prefix' to be a str")
        pulumi.set(__self__, "key_prefix", key_prefix)
        if location and not isinstance(location, str):
            raise TypeError("Expected argument 'location' to be a str")
        pulumi.set(__self__, "location", location)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
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
    @pulumi.getter(name="appEngineIntegrationMode")
    def app_engine_integration_mode(self) -> str:
        """
        The App Engine integration mode to use for this database.
        """
        return pulumi.get(self, "app_engine_integration_mode")

    @property
    @pulumi.getter(name="concurrencyMode")
    def concurrency_mode(self) -> str:
        """
        The concurrency control mode to use for this database.
        """
        return pulumi.get(self, "concurrency_mode")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The timestamp at which this database was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="keyPrefix")
    def key_prefix(self) -> str:
        """
        The key_prefix for this database. This key_prefix is used, in combination with the project id ("~") to construct the application id that is returned from the Cloud Datastore APIs in Google App Engine first generation runtimes. This value may be empty in which case the appid to use for URL-encoded keys is the project_id (eg: foo instead of v~foo).
        """
        return pulumi.get(self, "key_prefix")

    @property
    @pulumi.getter
    def location(self) -> str:
        """
        The location of the database. Available databases are listed at https://cloud.google.com/firestore/docs/locations.
        """
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name of the Database. Format: `projects/{project}/databases/{database}`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of the database. See https://cloud.google.com/datastore/docs/firestore-or-datastore for information about how to choose.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter
    def uid(self) -> str:
        """
        The system-generated UUID4 for this Database.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The timestamp at which this database was most recently updated. Note this only includes updates to the database resource and not data contained by the database.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetDatabaseResult(GetDatabaseResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabaseResult(
            app_engine_integration_mode=self.app_engine_integration_mode,
            concurrency_mode=self.concurrency_mode,
            create_time=self.create_time,
            etag=self.etag,
            key_prefix=self.key_prefix,
            location=self.location,
            name=self.name,
            type=self.type,
            uid=self.uid,
            update_time=self.update_time)


def get_database(database_id: Optional[str] = None,
                 project: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabaseResult:
    """
    Gets information about a database.
    """
    __args__ = dict()
    __args__['databaseId'] = database_id
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:firestore/v1:getDatabase', __args__, opts=opts, typ=GetDatabaseResult).value

    return AwaitableGetDatabaseResult(
        app_engine_integration_mode=__ret__.app_engine_integration_mode,
        concurrency_mode=__ret__.concurrency_mode,
        create_time=__ret__.create_time,
        etag=__ret__.etag,
        key_prefix=__ret__.key_prefix,
        location=__ret__.location,
        name=__ret__.name,
        type=__ret__.type,
        uid=__ret__.uid,
        update_time=__ret__.update_time)


@_utilities.lift_output_func(get_database)
def get_database_output(database_id: Optional[pulumi.Input[str]] = None,
                        project: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDatabaseResult]:
    """
    Gets information about a database.
    """
    ...
