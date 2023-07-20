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
    'GetKeyResult',
    'AwaitableGetKeyResult',
    'get_key',
    'get_key_output',
]

@pulumi.output_type
class GetKeyResult:
    def __init__(__self__, disabled=None, key_algorithm=None, key_origin=None, key_type=None, name=None, private_key_data=None, private_key_type=None, public_key_data=None, valid_after_time=None, valid_before_time=None):
        if disabled and not isinstance(disabled, bool):
            raise TypeError("Expected argument 'disabled' to be a bool")
        pulumi.set(__self__, "disabled", disabled)
        if key_algorithm and not isinstance(key_algorithm, str):
            raise TypeError("Expected argument 'key_algorithm' to be a str")
        pulumi.set(__self__, "key_algorithm", key_algorithm)
        if key_origin and not isinstance(key_origin, str):
            raise TypeError("Expected argument 'key_origin' to be a str")
        pulumi.set(__self__, "key_origin", key_origin)
        if key_type and not isinstance(key_type, str):
            raise TypeError("Expected argument 'key_type' to be a str")
        pulumi.set(__self__, "key_type", key_type)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if private_key_data and not isinstance(private_key_data, str):
            raise TypeError("Expected argument 'private_key_data' to be a str")
        pulumi.set(__self__, "private_key_data", private_key_data)
        if private_key_type and not isinstance(private_key_type, str):
            raise TypeError("Expected argument 'private_key_type' to be a str")
        pulumi.set(__self__, "private_key_type", private_key_type)
        if public_key_data and not isinstance(public_key_data, str):
            raise TypeError("Expected argument 'public_key_data' to be a str")
        pulumi.set(__self__, "public_key_data", public_key_data)
        if valid_after_time and not isinstance(valid_after_time, str):
            raise TypeError("Expected argument 'valid_after_time' to be a str")
        pulumi.set(__self__, "valid_after_time", valid_after_time)
        if valid_before_time and not isinstance(valid_before_time, str):
            raise TypeError("Expected argument 'valid_before_time' to be a str")
        pulumi.set(__self__, "valid_before_time", valid_before_time)

    @property
    @pulumi.getter
    def disabled(self) -> bool:
        """
        The key status.
        """
        return pulumi.get(self, "disabled")

    @property
    @pulumi.getter(name="keyAlgorithm")
    def key_algorithm(self) -> str:
        """
        Specifies the algorithm (and possibly key size) for the key.
        """
        return pulumi.get(self, "key_algorithm")

    @property
    @pulumi.getter(name="keyOrigin")
    def key_origin(self) -> str:
        """
        The key origin.
        """
        return pulumi.get(self, "key_origin")

    @property
    @pulumi.getter(name="keyType")
    def key_type(self) -> str:
        """
        The key type.
        """
        return pulumi.get(self, "key_type")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name of the service account key in the following format `projects/{PROJECT_ID}/serviceAccounts/{ACCOUNT}/keys/{key}`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="privateKeyData")
    def private_key_data(self) -> str:
        """
        The private key data. Only provided in `CreateServiceAccountKey` responses. Make sure to keep the private key data secure because it allows for the assertion of the service account identity. When base64 decoded, the private key data can be used to authenticate with Google API client libraries and with gcloud auth activate-service-account.
        """
        return pulumi.get(self, "private_key_data")

    @property
    @pulumi.getter(name="privateKeyType")
    def private_key_type(self) -> str:
        """
        The output format for the private key. Only provided in `CreateServiceAccountKey` responses, not in `GetServiceAccountKey` or `ListServiceAccountKey` responses. Google never exposes system-managed private keys, and never retains user-managed private keys.
        """
        return pulumi.get(self, "private_key_type")

    @property
    @pulumi.getter(name="publicKeyData")
    def public_key_data(self) -> str:
        """
        The public key data. Only provided in `GetServiceAccountKey` responses.
        """
        return pulumi.get(self, "public_key_data")

    @property
    @pulumi.getter(name="validAfterTime")
    def valid_after_time(self) -> str:
        """
        The key can be used after this timestamp.
        """
        return pulumi.get(self, "valid_after_time")

    @property
    @pulumi.getter(name="validBeforeTime")
    def valid_before_time(self) -> str:
        """
        The key can be used before this timestamp. For system-managed key pairs, this timestamp is the end time for the private key signing operation. The public key could still be used for verification for a few hours after this time.
        """
        return pulumi.get(self, "valid_before_time")


class AwaitableGetKeyResult(GetKeyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetKeyResult(
            disabled=self.disabled,
            key_algorithm=self.key_algorithm,
            key_origin=self.key_origin,
            key_type=self.key_type,
            name=self.name,
            private_key_data=self.private_key_data,
            private_key_type=self.private_key_type,
            public_key_data=self.public_key_data,
            valid_after_time=self.valid_after_time,
            valid_before_time=self.valid_before_time)


def get_key(key_id: Optional[str] = None,
            project: Optional[str] = None,
            public_key_type: Optional[str] = None,
            service_account_id: Optional[str] = None,
            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetKeyResult:
    """
    Gets a ServiceAccountKey.
    """
    __args__ = dict()
    __args__['keyId'] = key_id
    __args__['project'] = project
    __args__['publicKeyType'] = public_key_type
    __args__['serviceAccountId'] = service_account_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:iam/v1:getKey', __args__, opts=opts, typ=GetKeyResult).value

    return AwaitableGetKeyResult(
        disabled=pulumi.get(__ret__, 'disabled'),
        key_algorithm=pulumi.get(__ret__, 'key_algorithm'),
        key_origin=pulumi.get(__ret__, 'key_origin'),
        key_type=pulumi.get(__ret__, 'key_type'),
        name=pulumi.get(__ret__, 'name'),
        private_key_data=pulumi.get(__ret__, 'private_key_data'),
        private_key_type=pulumi.get(__ret__, 'private_key_type'),
        public_key_data=pulumi.get(__ret__, 'public_key_data'),
        valid_after_time=pulumi.get(__ret__, 'valid_after_time'),
        valid_before_time=pulumi.get(__ret__, 'valid_before_time'))


@_utilities.lift_output_func(get_key)
def get_key_output(key_id: Optional[pulumi.Input[str]] = None,
                   project: Optional[pulumi.Input[Optional[str]]] = None,
                   public_key_type: Optional[pulumi.Input[Optional[str]]] = None,
                   service_account_id: Optional[pulumi.Input[str]] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetKeyResult]:
    """
    Gets a ServiceAccountKey.
    """
    ...
