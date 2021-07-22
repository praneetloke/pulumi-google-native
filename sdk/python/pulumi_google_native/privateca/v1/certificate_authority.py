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

__all__ = ['CertificateAuthorityArgs', 'CertificateAuthority']

@pulumi.input_type
class CertificateAuthorityArgs:
    def __init__(__self__, *,
                 ca_pool_id: pulumi.Input[str],
                 certificate_authority_id: pulumi.Input[str],
                 config: pulumi.Input['CertificateConfigArgs'],
                 key_spec: pulumi.Input['KeyVersionSpecArgs'],
                 lifetime: pulumi.Input[str],
                 type: pulumi.Input['CertificateAuthorityType'],
                 gcs_bucket: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 subordinate_config: Optional[pulumi.Input['SubordinateConfigArgs']] = None):
        """
        The set of arguments for constructing a CertificateAuthority resource.
        :param pulumi.Input['CertificateConfigArgs'] config: Immutable. The config used to create a self-signed X.509 certificate or CSR.
        :param pulumi.Input['KeyVersionSpecArgs'] key_spec: Immutable. Used when issuing certificates for this CertificateAuthority. If this CertificateAuthority is a self-signed CertificateAuthority, this key is also used to sign the self-signed CA certificate. Otherwise, it is used to sign a CSR.
        :param pulumi.Input[str] lifetime: The desired lifetime of the CA certificate. Used to create the "not_before_time" and "not_after_time" fields inside an X.509 certificate.
        :param pulumi.Input['CertificateAuthorityType'] type: Immutable. The Type of this CertificateAuthority.
        :param pulumi.Input[str] gcs_bucket: Immutable. The name of a Cloud Storage bucket where this CertificateAuthority will publish content, such as the CA certificate and CRLs. This must be a bucket name, without any prefixes (such as `gs://`) or suffixes (such as `.googleapis.com`). For example, to use a bucket named `my-bucket`, you would simply specify `my-bucket`. If not specified, a managed bucket will be created.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Optional. Labels with user-defined metadata.
        :param pulumi.Input['SubordinateConfigArgs'] subordinate_config: Optional. If this is a subordinate CertificateAuthority, this field will be set with the subordinate configuration, which describes its issuers. This may be updated, but this CertificateAuthority must continue to validate.
        """
        pulumi.set(__self__, "ca_pool_id", ca_pool_id)
        pulumi.set(__self__, "certificate_authority_id", certificate_authority_id)
        pulumi.set(__self__, "config", config)
        pulumi.set(__self__, "key_spec", key_spec)
        pulumi.set(__self__, "lifetime", lifetime)
        pulumi.set(__self__, "type", type)
        if gcs_bucket is not None:
            pulumi.set(__self__, "gcs_bucket", gcs_bucket)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if request_id is not None:
            pulumi.set(__self__, "request_id", request_id)
        if subordinate_config is not None:
            pulumi.set(__self__, "subordinate_config", subordinate_config)

    @property
    @pulumi.getter(name="caPoolId")
    def ca_pool_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "ca_pool_id")

    @ca_pool_id.setter
    def ca_pool_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "ca_pool_id", value)

    @property
    @pulumi.getter(name="certificateAuthorityId")
    def certificate_authority_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "certificate_authority_id")

    @certificate_authority_id.setter
    def certificate_authority_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "certificate_authority_id", value)

    @property
    @pulumi.getter
    def config(self) -> pulumi.Input['CertificateConfigArgs']:
        """
        Immutable. The config used to create a self-signed X.509 certificate or CSR.
        """
        return pulumi.get(self, "config")

    @config.setter
    def config(self, value: pulumi.Input['CertificateConfigArgs']):
        pulumi.set(self, "config", value)

    @property
    @pulumi.getter(name="keySpec")
    def key_spec(self) -> pulumi.Input['KeyVersionSpecArgs']:
        """
        Immutable. Used when issuing certificates for this CertificateAuthority. If this CertificateAuthority is a self-signed CertificateAuthority, this key is also used to sign the self-signed CA certificate. Otherwise, it is used to sign a CSR.
        """
        return pulumi.get(self, "key_spec")

    @key_spec.setter
    def key_spec(self, value: pulumi.Input['KeyVersionSpecArgs']):
        pulumi.set(self, "key_spec", value)

    @property
    @pulumi.getter
    def lifetime(self) -> pulumi.Input[str]:
        """
        The desired lifetime of the CA certificate. Used to create the "not_before_time" and "not_after_time" fields inside an X.509 certificate.
        """
        return pulumi.get(self, "lifetime")

    @lifetime.setter
    def lifetime(self, value: pulumi.Input[str]):
        pulumi.set(self, "lifetime", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input['CertificateAuthorityType']:
        """
        Immutable. The Type of this CertificateAuthority.
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input['CertificateAuthorityType']):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="gcsBucket")
    def gcs_bucket(self) -> Optional[pulumi.Input[str]]:
        """
        Immutable. The name of a Cloud Storage bucket where this CertificateAuthority will publish content, such as the CA certificate and CRLs. This must be a bucket name, without any prefixes (such as `gs://`) or suffixes (such as `.googleapis.com`). For example, to use a bucket named `my-bucket`, you would simply specify `my-bucket`. If not specified, a managed bucket will be created.
        """
        return pulumi.get(self, "gcs_bucket")

    @gcs_bucket.setter
    def gcs_bucket(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "gcs_bucket", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Optional. Labels with user-defined metadata.
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

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "request_id")

    @request_id.setter
    def request_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_id", value)

    @property
    @pulumi.getter(name="subordinateConfig")
    def subordinate_config(self) -> Optional[pulumi.Input['SubordinateConfigArgs']]:
        """
        Optional. If this is a subordinate CertificateAuthority, this field will be set with the subordinate configuration, which describes its issuers. This may be updated, but this CertificateAuthority must continue to validate.
        """
        return pulumi.get(self, "subordinate_config")

    @subordinate_config.setter
    def subordinate_config(self, value: Optional[pulumi.Input['SubordinateConfigArgs']]):
        pulumi.set(self, "subordinate_config", value)


class CertificateAuthority(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ca_pool_id: Optional[pulumi.Input[str]] = None,
                 certificate_authority_id: Optional[pulumi.Input[str]] = None,
                 config: Optional[pulumi.Input[pulumi.InputType['CertificateConfigArgs']]] = None,
                 gcs_bucket: Optional[pulumi.Input[str]] = None,
                 key_spec: Optional[pulumi.Input[pulumi.InputType['KeyVersionSpecArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 lifetime: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 subordinate_config: Optional[pulumi.Input[pulumi.InputType['SubordinateConfigArgs']]] = None,
                 type: Optional[pulumi.Input['CertificateAuthorityType']] = None,
                 __props__=None):
        """
        Create a new CertificateAuthority in a given Project and Location.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['CertificateConfigArgs']] config: Immutable. The config used to create a self-signed X.509 certificate or CSR.
        :param pulumi.Input[str] gcs_bucket: Immutable. The name of a Cloud Storage bucket where this CertificateAuthority will publish content, such as the CA certificate and CRLs. This must be a bucket name, without any prefixes (such as `gs://`) or suffixes (such as `.googleapis.com`). For example, to use a bucket named `my-bucket`, you would simply specify `my-bucket`. If not specified, a managed bucket will be created.
        :param pulumi.Input[pulumi.InputType['KeyVersionSpecArgs']] key_spec: Immutable. Used when issuing certificates for this CertificateAuthority. If this CertificateAuthority is a self-signed CertificateAuthority, this key is also used to sign the self-signed CA certificate. Otherwise, it is used to sign a CSR.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Optional. Labels with user-defined metadata.
        :param pulumi.Input[str] lifetime: The desired lifetime of the CA certificate. Used to create the "not_before_time" and "not_after_time" fields inside an X.509 certificate.
        :param pulumi.Input[pulumi.InputType['SubordinateConfigArgs']] subordinate_config: Optional. If this is a subordinate CertificateAuthority, this field will be set with the subordinate configuration, which describes its issuers. This may be updated, but this CertificateAuthority must continue to validate.
        :param pulumi.Input['CertificateAuthorityType'] type: Immutable. The Type of this CertificateAuthority.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CertificateAuthorityArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a new CertificateAuthority in a given Project and Location.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param CertificateAuthorityArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CertificateAuthorityArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 ca_pool_id: Optional[pulumi.Input[str]] = None,
                 certificate_authority_id: Optional[pulumi.Input[str]] = None,
                 config: Optional[pulumi.Input[pulumi.InputType['CertificateConfigArgs']]] = None,
                 gcs_bucket: Optional[pulumi.Input[str]] = None,
                 key_spec: Optional[pulumi.Input[pulumi.InputType['KeyVersionSpecArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 lifetime: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 subordinate_config: Optional[pulumi.Input[pulumi.InputType['SubordinateConfigArgs']]] = None,
                 type: Optional[pulumi.Input['CertificateAuthorityType']] = None,
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
            __props__ = CertificateAuthorityArgs.__new__(CertificateAuthorityArgs)

            if ca_pool_id is None and not opts.urn:
                raise TypeError("Missing required property 'ca_pool_id'")
            __props__.__dict__["ca_pool_id"] = ca_pool_id
            if certificate_authority_id is None and not opts.urn:
                raise TypeError("Missing required property 'certificate_authority_id'")
            __props__.__dict__["certificate_authority_id"] = certificate_authority_id
            if config is None and not opts.urn:
                raise TypeError("Missing required property 'config'")
            __props__.__dict__["config"] = config
            __props__.__dict__["gcs_bucket"] = gcs_bucket
            if key_spec is None and not opts.urn:
                raise TypeError("Missing required property 'key_spec'")
            __props__.__dict__["key_spec"] = key_spec
            __props__.__dict__["labels"] = labels
            if lifetime is None and not opts.urn:
                raise TypeError("Missing required property 'lifetime'")
            __props__.__dict__["lifetime"] = lifetime
            __props__.__dict__["location"] = location
            __props__.__dict__["project"] = project
            __props__.__dict__["request_id"] = request_id
            __props__.__dict__["subordinate_config"] = subordinate_config
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
            __props__.__dict__["access_urls"] = None
            __props__.__dict__["ca_certificate_descriptions"] = None
            __props__.__dict__["create_time"] = None
            __props__.__dict__["delete_time"] = None
            __props__.__dict__["expire_time"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["pem_ca_certificates"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["tier"] = None
            __props__.__dict__["update_time"] = None
        super(CertificateAuthority, __self__).__init__(
            'google-native:privateca/v1:CertificateAuthority',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'CertificateAuthority':
        """
        Get an existing CertificateAuthority resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = CertificateAuthorityArgs.__new__(CertificateAuthorityArgs)

        __props__.__dict__["access_urls"] = None
        __props__.__dict__["ca_certificate_descriptions"] = None
        __props__.__dict__["config"] = None
        __props__.__dict__["create_time"] = None
        __props__.__dict__["delete_time"] = None
        __props__.__dict__["expire_time"] = None
        __props__.__dict__["gcs_bucket"] = None
        __props__.__dict__["key_spec"] = None
        __props__.__dict__["labels"] = None
        __props__.__dict__["lifetime"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["pem_ca_certificates"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["subordinate_config"] = None
        __props__.__dict__["tier"] = None
        __props__.__dict__["type"] = None
        __props__.__dict__["update_time"] = None
        return CertificateAuthority(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accessUrls")
    def access_urls(self) -> pulumi.Output['outputs.AccessUrlsResponse']:
        """
        URLs for accessing content published by this CA, such as the CA certificate and CRLs.
        """
        return pulumi.get(self, "access_urls")

    @property
    @pulumi.getter(name="caCertificateDescriptions")
    def ca_certificate_descriptions(self) -> pulumi.Output[Sequence['outputs.CertificateDescriptionResponse']]:
        """
        A structured description of this CertificateAuthority's CA certificate and its issuers. Ordered as self-to-root.
        """
        return pulumi.get(self, "ca_certificate_descriptions")

    @property
    @pulumi.getter
    def config(self) -> pulumi.Output['outputs.CertificateConfigResponse']:
        """
        Immutable. The config used to create a self-signed X.509 certificate or CSR.
        """
        return pulumi.get(self, "config")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The time at which this CertificateAuthority was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="deleteTime")
    def delete_time(self) -> pulumi.Output[str]:
        """
        The time at which this CertificateAuthority was soft deleted, if it is in the DELETED state.
        """
        return pulumi.get(self, "delete_time")

    @property
    @pulumi.getter(name="expireTime")
    def expire_time(self) -> pulumi.Output[str]:
        """
        The time at which this CertificateAuthority will be permanently purged, if it is in the DELETED state.
        """
        return pulumi.get(self, "expire_time")

    @property
    @pulumi.getter(name="gcsBucket")
    def gcs_bucket(self) -> pulumi.Output[str]:
        """
        Immutable. The name of a Cloud Storage bucket where this CertificateAuthority will publish content, such as the CA certificate and CRLs. This must be a bucket name, without any prefixes (such as `gs://`) or suffixes (such as `.googleapis.com`). For example, to use a bucket named `my-bucket`, you would simply specify `my-bucket`. If not specified, a managed bucket will be created.
        """
        return pulumi.get(self, "gcs_bucket")

    @property
    @pulumi.getter(name="keySpec")
    def key_spec(self) -> pulumi.Output['outputs.KeyVersionSpecResponse']:
        """
        Immutable. Used when issuing certificates for this CertificateAuthority. If this CertificateAuthority is a self-signed CertificateAuthority, this key is also used to sign the self-signed CA certificate. Otherwise, it is used to sign a CSR.
        """
        return pulumi.get(self, "key_spec")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Optional. Labels with user-defined metadata.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def lifetime(self) -> pulumi.Output[str]:
        """
        The desired lifetime of the CA certificate. Used to create the "not_before_time" and "not_after_time" fields inside an X.509 certificate.
        """
        return pulumi.get(self, "lifetime")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name for this CertificateAuthority in the format `projects/*/locations/*/caPools/*/certificateAuthorities/*`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="pemCaCertificates")
    def pem_ca_certificates(self) -> pulumi.Output[Sequence[str]]:
        """
        This CertificateAuthority's certificate chain, including the current CertificateAuthority's certificate. Ordered such that the root issuer is the final element (consistent with RFC 5246). For a self-signed CA, this will only list the current CertificateAuthority's certificate.
        """
        return pulumi.get(self, "pem_ca_certificates")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output[str]:
        """
        The State for this CertificateAuthority.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="subordinateConfig")
    def subordinate_config(self) -> pulumi.Output['outputs.SubordinateConfigResponse']:
        """
        Optional. If this is a subordinate CertificateAuthority, this field will be set with the subordinate configuration, which describes its issuers. This may be updated, but this CertificateAuthority must continue to validate.
        """
        return pulumi.get(self, "subordinate_config")

    @property
    @pulumi.getter
    def tier(self) -> pulumi.Output[str]:
        """
        The CaPool.Tier of the CaPool that includes this CertificateAuthority.
        """
        return pulumi.get(self, "tier")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Immutable. The Type of this CertificateAuthority.
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The time at which this CertificateAuthority was last updated.
        """
        return pulumi.get(self, "update_time")

