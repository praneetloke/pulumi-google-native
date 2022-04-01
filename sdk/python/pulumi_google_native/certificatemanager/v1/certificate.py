# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs
from ._enums import *
from ._inputs import *

__all__ = ['CertificateArgs', 'Certificate']

@pulumi.input_type
class CertificateArgs:
    def __init__(__self__, *,
                 certificate_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 managed: Optional[pulumi.Input['ManagedCertificateArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input['CertificateScope']] = None,
                 self_managed: Optional[pulumi.Input['SelfManagedCertificateArgs']] = None):
        """
        The set of arguments for constructing a Certificate resource.
        :param pulumi.Input[str] certificate_id: Required. A user-provided name of the certificate.
        :param pulumi.Input[str] description: One or more paragraphs of text description of a certificate.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Set of labels associated with a Certificate.
        :param pulumi.Input['ManagedCertificateArgs'] managed: If set, contains configuration and state of a managed certificate.
        :param pulumi.Input[str] name: A user-defined name of the certificate. Certificate names must be unique globally and match pattern `projects/*/locations/*/certificates/*`.
        :param pulumi.Input['CertificateScope'] scope: Immutable. The scope of the certificate.
        :param pulumi.Input['SelfManagedCertificateArgs'] self_managed: If set, defines data of a self-managed certificate.
        """
        pulumi.set(__self__, "certificate_id", certificate_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if managed is not None:
            pulumi.set(__self__, "managed", managed)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if scope is not None:
            pulumi.set(__self__, "scope", scope)
        if self_managed is not None:
            pulumi.set(__self__, "self_managed", self_managed)

    @property
    @pulumi.getter(name="certificateId")
    def certificate_id(self) -> pulumi.Input[str]:
        """
        Required. A user-provided name of the certificate.
        """
        return pulumi.get(self, "certificate_id")

    @certificate_id.setter
    def certificate_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "certificate_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        One or more paragraphs of text description of a certificate.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Set of labels associated with a Certificate.
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
    def managed(self) -> Optional[pulumi.Input['ManagedCertificateArgs']]:
        """
        If set, contains configuration and state of a managed certificate.
        """
        return pulumi.get(self, "managed")

    @managed.setter
    def managed(self, value: Optional[pulumi.Input['ManagedCertificateArgs']]):
        pulumi.set(self, "managed", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        A user-defined name of the certificate. Certificate names must be unique globally and match pattern `projects/*/locations/*/certificates/*`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def project(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter
    def scope(self) -> Optional[pulumi.Input['CertificateScope']]:
        """
        Immutable. The scope of the certificate.
        """
        return pulumi.get(self, "scope")

    @scope.setter
    def scope(self, value: Optional[pulumi.Input['CertificateScope']]):
        pulumi.set(self, "scope", value)

    @property
    @pulumi.getter(name="selfManaged")
    def self_managed(self) -> Optional[pulumi.Input['SelfManagedCertificateArgs']]:
        """
        If set, defines data of a self-managed certificate.
        """
        return pulumi.get(self, "self_managed")

    @self_managed.setter
    def self_managed(self, value: Optional[pulumi.Input['SelfManagedCertificateArgs']]):
        pulumi.set(self, "self_managed", value)


class Certificate(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 managed: Optional[pulumi.Input[pulumi.InputType['ManagedCertificateArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input['CertificateScope']] = None,
                 self_managed: Optional[pulumi.Input[pulumi.InputType['SelfManagedCertificateArgs']]] = None,
                 __props__=None):
        """
        Creates a new Certificate in a given project and location.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] certificate_id: Required. A user-provided name of the certificate.
        :param pulumi.Input[str] description: One or more paragraphs of text description of a certificate.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Set of labels associated with a Certificate.
        :param pulumi.Input[pulumi.InputType['ManagedCertificateArgs']] managed: If set, contains configuration and state of a managed certificate.
        :param pulumi.Input[str] name: A user-defined name of the certificate. Certificate names must be unique globally and match pattern `projects/*/locations/*/certificates/*`.
        :param pulumi.Input['CertificateScope'] scope: Immutable. The scope of the certificate.
        :param pulumi.Input[pulumi.InputType['SelfManagedCertificateArgs']] self_managed: If set, defines data of a self-managed certificate.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CertificateArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new Certificate in a given project and location.

        :param str resource_name: The name of the resource.
        :param CertificateArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CertificateArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 certificate_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 managed: Optional[pulumi.Input[pulumi.InputType['ManagedCertificateArgs']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 scope: Optional[pulumi.Input['CertificateScope']] = None,
                 self_managed: Optional[pulumi.Input[pulumi.InputType['SelfManagedCertificateArgs']]] = None,
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
            __props__ = CertificateArgs.__new__(CertificateArgs)

            if certificate_id is None and not opts.urn:
                raise TypeError("Missing required property 'certificate_id'")
            __props__.__dict__["certificate_id"] = certificate_id
            __props__.__dict__["description"] = description
            __props__.__dict__["labels"] = labels
            __props__.__dict__["location"] = location
            __props__.__dict__["managed"] = managed
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            __props__.__dict__["scope"] = scope
            __props__.__dict__["self_managed"] = self_managed
            __props__.__dict__["create_time"] = None
            __props__.__dict__["expire_time"] = None
            __props__.__dict__["pem_certificate"] = None
            __props__.__dict__["san_dnsnames"] = None
            __props__.__dict__["update_time"] = None
        super(Certificate, __self__).__init__(
            'google-native:certificatemanager/v1:Certificate',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Certificate':
        """
        Get an existing Certificate resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = CertificateArgs.__new__(CertificateArgs)

        __props__.__dict__["create_time"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["expire_time"] = None
        __props__.__dict__["labels"] = None
        __props__.__dict__["managed"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["pem_certificate"] = None
        __props__.__dict__["san_dnsnames"] = None
        __props__.__dict__["scope"] = None
        __props__.__dict__["self_managed"] = None
        __props__.__dict__["update_time"] = None
        return Certificate(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The creation timestamp of a Certificate.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        One or more paragraphs of text description of a certificate.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> pulumi.Output[str]:
        """
        The expiry timestamp of a Certificate.
        """
        return pulumi.get(self, "expire_time")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Set of labels associated with a Certificate.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def managed(self) -> pulumi.Output['outputs.ManagedCertificateResponse']:
        """
        If set, contains configuration and state of a managed certificate.
        """
        return pulumi.get(self, "managed")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        A user-defined name of the certificate. Certificate names must be unique globally and match pattern `projects/*/locations/*/certificates/*`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="pemCertificate")
    def pem_certificate(self) -> pulumi.Output[str]:
        """
        The PEM-encoded certificate chain.
        """
        return pulumi.get(self, "pem_certificate")

    @property
    @pulumi.getter(name="sanDnsnames")
    def san_dnsnames(self) -> pulumi.Output[Sequence[str]]:
        """
        The list of Subject Alternative Names of dnsName type defined in the certificate (see RFC 5280 4.2.1.6). Managed certificates that haven't been provisioned yet have this field populated with a value of the managed.domains field.
        """
        return pulumi.get(self, "san_dnsnames")

    @property
    @pulumi.getter
    def scope(self) -> pulumi.Output[str]:
        """
        Immutable. The scope of the certificate.
        """
        return pulumi.get(self, "scope")

    @property
    @pulumi.getter(name="selfManaged")
    def self_managed(self) -> pulumi.Output['outputs.SelfManagedCertificateResponse']:
        """
        If set, defines data of a self-managed certificate.
        """
        return pulumi.get(self, "self_managed")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The last update timestamp of a Certificate.
        """
        return pulumi.get(self, "update_time")

