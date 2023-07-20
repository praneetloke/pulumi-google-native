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
    'GetDatabaseResult',
    'AwaitableGetDatabaseResult',
    'get_database',
    'get_database_output',
]

@pulumi.output_type
class GetDatabaseResult:
    def __init__(__self__, charset=None, collation=None, etag=None, instance=None, kind=None, name=None, project=None, self_link=None, sqlserver_database_details=None):
        if charset and not isinstance(charset, str):
            raise TypeError("Expected argument 'charset' to be a str")
        pulumi.set(__self__, "charset", charset)
        if collation and not isinstance(collation, str):
            raise TypeError("Expected argument 'collation' to be a str")
        pulumi.set(__self__, "collation", collation)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if instance and not isinstance(instance, str):
            raise TypeError("Expected argument 'instance' to be a str")
        pulumi.set(__self__, "instance", instance)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if project and not isinstance(project, str):
            raise TypeError("Expected argument 'project' to be a str")
        pulumi.set(__self__, "project", project)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if sqlserver_database_details and not isinstance(sqlserver_database_details, dict):
            raise TypeError("Expected argument 'sqlserver_database_details' to be a dict")
        pulumi.set(__self__, "sqlserver_database_details", sqlserver_database_details)

    @property
    @pulumi.getter
    def charset(self) -> str:
        """
        The Cloud SQL charset value.
        """
        return pulumi.get(self, "charset")

    @property
    @pulumi.getter
    def collation(self) -> str:
        """
        The Cloud SQL collation value.
        """
        return pulumi.get(self, "collation")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        This field is deprecated and will be removed from a future version of the API.
        """
        warnings.warn("""This field is deprecated and will be removed from a future version of the API.""", DeprecationWarning)
        pulumi.log.warn("""etag is deprecated: This field is deprecated and will be removed from a future version of the API.""")

        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def instance(self) -> str:
        """
        The name of the Cloud SQL instance. This does not include the project ID.
        """
        return pulumi.get(self, "instance")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        This is always `sql#database`.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the database in the Cloud SQL instance. This does not include the project ID or instance name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> str:
        """
        The project ID of the project containing the Cloud SQL database. The Google apps domain is prefixed if applicable.
        """
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        The URI of this resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="sqlserverDatabaseDetails")
    def sqlserver_database_details(self) -> 'outputs.SqlServerDatabaseDetailsResponse':
        return pulumi.get(self, "sqlserver_database_details")


class AwaitableGetDatabaseResult(GetDatabaseResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetDatabaseResult(
            charset=self.charset,
            collation=self.collation,
            etag=self.etag,
            instance=self.instance,
            kind=self.kind,
            name=self.name,
            project=self.project,
            self_link=self.self_link,
            sqlserver_database_details=self.sqlserver_database_details)


def get_database(database: Optional[str] = None,
                 instance: Optional[str] = None,
                 project: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetDatabaseResult:
    """
    Retrieves a resource containing information about a database inside a Cloud SQL instance.
    """
    __args__ = dict()
    __args__['database'] = database
    __args__['instance'] = instance
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:sqladmin/v1:getDatabase', __args__, opts=opts, typ=GetDatabaseResult).value

    return AwaitableGetDatabaseResult(
        charset=pulumi.get(__ret__, 'charset'),
        collation=pulumi.get(__ret__, 'collation'),
        etag=pulumi.get(__ret__, 'etag'),
        instance=pulumi.get(__ret__, 'instance'),
        kind=pulumi.get(__ret__, 'kind'),
        name=pulumi.get(__ret__, 'name'),
        project=pulumi.get(__ret__, 'project'),
        self_link=pulumi.get(__ret__, 'self_link'),
        sqlserver_database_details=pulumi.get(__ret__, 'sqlserver_database_details'))


@_utilities.lift_output_func(get_database)
def get_database_output(database: Optional[pulumi.Input[str]] = None,
                        instance: Optional[pulumi.Input[str]] = None,
                        project: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetDatabaseResult]:
    """
    Retrieves a resource containing information about a database inside a Cloud SQL instance.
    """
    ...
