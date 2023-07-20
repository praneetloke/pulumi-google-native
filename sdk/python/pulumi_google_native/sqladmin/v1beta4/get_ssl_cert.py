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
    'GetSslCertResult',
    'AwaitableGetSslCertResult',
    'get_ssl_cert',
    'get_ssl_cert_output',
]

@pulumi.output_type
class GetSslCertResult:
    def __init__(__self__, cert=None, cert_serial_number=None, common_name=None, create_time=None, expiration_time=None, instance=None, kind=None, self_link=None, sha1_fingerprint=None):
        if cert and not isinstance(cert, str):
            raise TypeError("Expected argument 'cert' to be a str")
        pulumi.set(__self__, "cert", cert)
        if cert_serial_number and not isinstance(cert_serial_number, str):
            raise TypeError("Expected argument 'cert_serial_number' to be a str")
        pulumi.set(__self__, "cert_serial_number", cert_serial_number)
        if common_name and not isinstance(common_name, str):
            raise TypeError("Expected argument 'common_name' to be a str")
        pulumi.set(__self__, "common_name", common_name)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if expiration_time and not isinstance(expiration_time, str):
            raise TypeError("Expected argument 'expiration_time' to be a str")
        pulumi.set(__self__, "expiration_time", expiration_time)
        if instance and not isinstance(instance, str):
            raise TypeError("Expected argument 'instance' to be a str")
        pulumi.set(__self__, "instance", instance)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if sha1_fingerprint and not isinstance(sha1_fingerprint, str):
            raise TypeError("Expected argument 'sha1_fingerprint' to be a str")
        pulumi.set(__self__, "sha1_fingerprint", sha1_fingerprint)

    @property
    @pulumi.getter
    def cert(self) -> str:
        """
        PEM representation.
        """
        return pulumi.get(self, "cert")

    @property
    @pulumi.getter(name="certSerialNumber")
    def cert_serial_number(self) -> str:
        """
        Serial number, as extracted from the certificate.
        """
        return pulumi.get(self, "cert_serial_number")

    @property
    @pulumi.getter(name="commonName")
    def common_name(self) -> str:
        """
        User supplied name. Constrained to [a-zA-Z.-_ ]+.
        """
        return pulumi.get(self, "common_name")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The time when the certificate was created in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="expirationTime")
    def expiration_time(self) -> str:
        """
        The time when the certificate expires in [RFC 3339](https://tools.ietf.org/html/rfc3339) format, for example `2012-11-15T16:19:00.094Z`.
        """
        return pulumi.get(self, "expiration_time")

    @property
    @pulumi.getter
    def instance(self) -> str:
        """
        Name of the database instance.
        """
        return pulumi.get(self, "instance")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        This is always `sql#sslCert`.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        The URI of this resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="sha1Fingerprint")
    def sha1_fingerprint(self) -> str:
        """
        Sha1 Fingerprint.
        """
        return pulumi.get(self, "sha1_fingerprint")


class AwaitableGetSslCertResult(GetSslCertResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetSslCertResult(
            cert=self.cert,
            cert_serial_number=self.cert_serial_number,
            common_name=self.common_name,
            create_time=self.create_time,
            expiration_time=self.expiration_time,
            instance=self.instance,
            kind=self.kind,
            self_link=self.self_link,
            sha1_fingerprint=self.sha1_fingerprint)


def get_ssl_cert(instance: Optional[str] = None,
                 project: Optional[str] = None,
                 sha1_fingerprint: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetSslCertResult:
    """
    Retrieves a particular SSL certificate. Does not include the private key (required for usage). The private key must be saved from the response to initial creation.
    """
    __args__ = dict()
    __args__['instance'] = instance
    __args__['project'] = project
    __args__['sha1Fingerprint'] = sha1_fingerprint
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:sqladmin/v1beta4:getSslCert', __args__, opts=opts, typ=GetSslCertResult).value

    return AwaitableGetSslCertResult(
        cert=pulumi.get(__ret__, 'cert'),
        cert_serial_number=pulumi.get(__ret__, 'cert_serial_number'),
        common_name=pulumi.get(__ret__, 'common_name'),
        create_time=pulumi.get(__ret__, 'create_time'),
        expiration_time=pulumi.get(__ret__, 'expiration_time'),
        instance=pulumi.get(__ret__, 'instance'),
        kind=pulumi.get(__ret__, 'kind'),
        self_link=pulumi.get(__ret__, 'self_link'),
        sha1_fingerprint=pulumi.get(__ret__, 'sha1_fingerprint'))


@_utilities.lift_output_func(get_ssl_cert)
def get_ssl_cert_output(instance: Optional[pulumi.Input[str]] = None,
                        project: Optional[pulumi.Input[Optional[str]]] = None,
                        sha1_fingerprint: Optional[pulumi.Input[str]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetSslCertResult]:
    """
    Retrieves a particular SSL certificate. Does not include the private key (required for usage). The private key must be saved from the response to initial creation.
    """
    ...
