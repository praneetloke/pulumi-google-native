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
    'GetKeyResult',
    'AwaitableGetKeyResult',
    'get_key',
    'get_key_output',
]

@pulumi.output_type
class GetKeyResult:
    def __init__(__self__, android_settings=None, create_time=None, display_name=None, ios_settings=None, labels=None, name=None, testing_options=None, waf_settings=None, web_settings=None):
        if android_settings and not isinstance(android_settings, dict):
            raise TypeError("Expected argument 'android_settings' to be a dict")
        pulumi.set(__self__, "android_settings", android_settings)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if ios_settings and not isinstance(ios_settings, dict):
            raise TypeError("Expected argument 'ios_settings' to be a dict")
        pulumi.set(__self__, "ios_settings", ios_settings)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if testing_options and not isinstance(testing_options, dict):
            raise TypeError("Expected argument 'testing_options' to be a dict")
        pulumi.set(__self__, "testing_options", testing_options)
        if waf_settings and not isinstance(waf_settings, dict):
            raise TypeError("Expected argument 'waf_settings' to be a dict")
        pulumi.set(__self__, "waf_settings", waf_settings)
        if web_settings and not isinstance(web_settings, dict):
            raise TypeError("Expected argument 'web_settings' to be a dict")
        pulumi.set(__self__, "web_settings", web_settings)

    @property
    @pulumi.getter(name="androidSettings")
    def android_settings(self) -> 'outputs.GoogleCloudRecaptchaenterpriseV1AndroidKeySettingsResponse':
        """
        Settings for keys that can be used by Android apps.
        """
        return pulumi.get(self, "android_settings")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The timestamp corresponding to the creation of this key.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Human-readable display name of this key. Modifiable by user.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="iosSettings")
    def ios_settings(self) -> 'outputs.GoogleCloudRecaptchaenterpriseV1IOSKeySettingsResponse':
        """
        Settings for keys that can be used by iOS apps.
        """
        return pulumi.get(self, "ios_settings")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        See Creating and managing labels.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name for the Key in the format "projects/{project}/keys/{key}".
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="testingOptions")
    def testing_options(self) -> 'outputs.GoogleCloudRecaptchaenterpriseV1TestingOptionsResponse':
        """
        Options for user acceptance testing.
        """
        return pulumi.get(self, "testing_options")

    @property
    @pulumi.getter(name="wafSettings")
    def waf_settings(self) -> 'outputs.GoogleCloudRecaptchaenterpriseV1WafSettingsResponse':
        """
        Settings for WAF
        """
        return pulumi.get(self, "waf_settings")

    @property
    @pulumi.getter(name="webSettings")
    def web_settings(self) -> 'outputs.GoogleCloudRecaptchaenterpriseV1WebKeySettingsResponse':
        """
        Settings for keys that can be used by websites.
        """
        return pulumi.get(self, "web_settings")


class AwaitableGetKeyResult(GetKeyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetKeyResult(
            android_settings=self.android_settings,
            create_time=self.create_time,
            display_name=self.display_name,
            ios_settings=self.ios_settings,
            labels=self.labels,
            name=self.name,
            testing_options=self.testing_options,
            waf_settings=self.waf_settings,
            web_settings=self.web_settings)


def get_key(key_id: Optional[str] = None,
            project: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetKeyResult:
    """
    Returns the specified key.
    """
    __args__ = dict()
    __args__['keyId'] = key_id
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:recaptchaenterprise/v1:getKey', __args__, opts=opts, typ=GetKeyResult).value

    return AwaitableGetKeyResult(
        android_settings=pulumi.get(__ret__, 'android_settings'),
        create_time=pulumi.get(__ret__, 'create_time'),
        display_name=pulumi.get(__ret__, 'display_name'),
        ios_settings=pulumi.get(__ret__, 'ios_settings'),
        labels=pulumi.get(__ret__, 'labels'),
        name=pulumi.get(__ret__, 'name'),
        testing_options=pulumi.get(__ret__, 'testing_options'),
        waf_settings=pulumi.get(__ret__, 'waf_settings'),
        web_settings=pulumi.get(__ret__, 'web_settings'))


@_utilities.lift_output_func(get_key)
def get_key_output(key_id: Optional[pulumi.Input[str]] = None,
                   project: Optional[pulumi.Input[Optional[str]]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetKeyResult]:
    """
    Returns the specified key.
    """
    ...
