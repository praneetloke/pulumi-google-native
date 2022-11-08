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
    'GetSfdcInstanceResult',
    'AwaitableGetSfdcInstanceResult',
    'get_sfdc_instance',
    'get_sfdc_instance_output',
]

@pulumi.output_type
class GetSfdcInstanceResult:
    def __init__(__self__, auth_config_id=None, create_time=None, delete_time=None, description=None, display_name=None, name=None, service_authority=None, sfdc_org_id=None, update_time=None):
        if auth_config_id and not isinstance(auth_config_id, list):
            raise TypeError("Expected argument 'auth_config_id' to be a list")
        pulumi.set(__self__, "auth_config_id", auth_config_id)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if delete_time and not isinstance(delete_time, str):
            raise TypeError("Expected argument 'delete_time' to be a str")
        pulumi.set(__self__, "delete_time", delete_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if service_authority and not isinstance(service_authority, str):
            raise TypeError("Expected argument 'service_authority' to be a str")
        pulumi.set(__self__, "service_authority", service_authority)
        if sfdc_org_id and not isinstance(sfdc_org_id, str):
            raise TypeError("Expected argument 'sfdc_org_id' to be a str")
        pulumi.set(__self__, "sfdc_org_id", sfdc_org_id)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="authConfigId")
    def auth_config_id(self) -> Sequence[str]:
        """
        A list of AuthConfigs that can be tried to open the channel to SFDC
        """
        return pulumi.get(self, "auth_config_id")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        Time when the instance is created
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="deleteTime")
    def delete_time(self) -> str:
        """
        Time when the instance was deleted. Empty if not deleted.
        """
        return pulumi.get(self, "delete_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        A description of the sfdc instance.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        User selected unique name/alias to easily reference an instance.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name of the SFDC instance projects/{project}/locations/{location}/sfdcInstances/{sfdcInstance}.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="serviceAuthority")
    def service_authority(self) -> str:
        """
        URL used for API calls after authentication (the login authority is configured within the referenced AuthConfig).
        """
        return pulumi.get(self, "service_authority")

    @property
    @pulumi.getter(name="sfdcOrgId")
    def sfdc_org_id(self) -> str:
        """
        The SFDC Org Id. This is defined in salesforce.
        """
        return pulumi.get(self, "sfdc_org_id")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        Time when the instance was last updated
        """
        return pulumi.get(self, "update_time")


class AwaitableGetSfdcInstanceResult(GetSfdcInstanceResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSfdcInstanceResult(
            auth_config_id=self.auth_config_id,
            create_time=self.create_time,
            delete_time=self.delete_time,
            description=self.description,
            display_name=self.display_name,
            name=self.name,
            service_authority=self.service_authority,
            sfdc_org_id=self.sfdc_org_id,
            update_time=self.update_time)


def get_sfdc_instance(location: Optional[str] = None,
                      product_id: Optional[str] = None,
                      project: Optional[str] = None,
                      sfdc_instance_id: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSfdcInstanceResult:
    """
    Gets an sfdc instance. If the instance doesn't exist, Code.NOT_FOUND exception will be thrown.
    """
    __args__ = dict()
    __args__['location'] = location
    __args__['productId'] = product_id
    __args__['project'] = project
    __args__['sfdcInstanceId'] = sfdc_instance_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:integrations/v1alpha:getSfdcInstance', __args__, opts=opts, typ=GetSfdcInstanceResult).value

    return AwaitableGetSfdcInstanceResult(
        auth_config_id=__ret__.auth_config_id,
        create_time=__ret__.create_time,
        delete_time=__ret__.delete_time,
        description=__ret__.description,
        display_name=__ret__.display_name,
        name=__ret__.name,
        service_authority=__ret__.service_authority,
        sfdc_org_id=__ret__.sfdc_org_id,
        update_time=__ret__.update_time)


@_utilities.lift_output_func(get_sfdc_instance)
def get_sfdc_instance_output(location: Optional[pulumi.Input[str]] = None,
                             product_id: Optional[pulumi.Input[str]] = None,
                             project: Optional[pulumi.Input[Optional[str]]] = None,
                             sfdc_instance_id: Optional[pulumi.Input[str]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSfdcInstanceResult]:
    """
    Gets an sfdc instance. If the instance doesn't exist, Code.NOT_FOUND exception will be thrown.
    """
    ...
