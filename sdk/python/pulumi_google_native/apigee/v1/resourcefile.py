# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = ['ResourcefileArgs', 'Resourcefile']

@pulumi.input_type
class ResourcefileArgs:
    def __init__(__self__, *,
                 environment_id: pulumi.Input[str],
                 name: pulumi.Input[str],
                 organization_id: pulumi.Input[str],
                 type: pulumi.Input[str],
                 content_type: Optional[pulumi.Input[str]] = None,
                 data: Optional[pulumi.Input[str]] = None,
                 extensions: Optional[pulumi.Input[Sequence[pulumi.Input[Mapping[str, Any]]]]] = None,
                 file: Optional[pulumi.Input[Union[pulumi.Asset, pulumi.Archive]]] = None):
        """
        The set of arguments for constructing a Resourcefile resource.
        :param pulumi.Input[str] name: Required. Name of the resource file. Must match the regular expression: [a-zA-Z0-9:/\\\\!@#$%^&{}\\[\\]()+\\-=,.~'` ]{1,255}
        :param pulumi.Input[str] type: Required. Resource file type. {{ resource_file_type }}
        :param pulumi.Input[str] content_type: The HTTP Content-Type header value specifying the content type of the body.
        :param pulumi.Input[str] data: The HTTP request/response body as raw binary.
        :param pulumi.Input[Sequence[pulumi.Input[Mapping[str, Any]]]] extensions: Application specific response metadata. Must be set in the first response for streaming APIs.
        :param pulumi.Input[Union[pulumi.Asset, pulumi.Archive]] file: File to upload.
        """
        pulumi.set(__self__, "environment_id", environment_id)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "organization_id", organization_id)
        pulumi.set(__self__, "type", type)
        if content_type is not None:
            pulumi.set(__self__, "content_type", content_type)
        if data is not None:
            pulumi.set(__self__, "data", data)
        if extensions is not None:
            pulumi.set(__self__, "extensions", extensions)
        if file is not None:
            pulumi.set(__self__, "file", file)

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "environment_id")

    @environment_id.setter
    def environment_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "environment_id", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Required. Name of the resource file. Must match the regular expression: [a-zA-Z0-9:/\\\\!@#$%^&{}\\[\\]()+\\-=,.~'` ]{1,255}
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "organization_id")

    @organization_id.setter
    def organization_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "organization_id", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        Required. Resource file type. {{ resource_file_type }}
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> Optional[pulumi.Input[str]]:
        """
        The HTTP Content-Type header value specifying the content type of the body.
        """
        return pulumi.get(self, "content_type")

    @content_type.setter
    def content_type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_type", value)

    @property
    @pulumi.getter
    def data(self) -> Optional[pulumi.Input[str]]:
        """
        The HTTP request/response body as raw binary.
        """
        return pulumi.get(self, "data")

    @data.setter
    def data(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "data", value)

    @property
    @pulumi.getter
    def extensions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[Mapping[str, Any]]]]]:
        """
        Application specific response metadata. Must be set in the first response for streaming APIs.
        """
        return pulumi.get(self, "extensions")

    @extensions.setter
    def extensions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[Mapping[str, Any]]]]]):
        pulumi.set(self, "extensions", value)

    @property
    @pulumi.getter
    def file(self) -> Optional[pulumi.Input[Union[pulumi.Asset, pulumi.Archive]]]:
        """
        File to upload.
        """
        return pulumi.get(self, "file")

    @file.setter
    def file(self, value: Optional[pulumi.Input[Union[pulumi.Asset, pulumi.Archive]]]):
        pulumi.set(self, "file", value)


class Resourcefile(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 content_type: Optional[pulumi.Input[str]] = None,
                 data: Optional[pulumi.Input[str]] = None,
                 environment_id: Optional[pulumi.Input[str]] = None,
                 extensions: Optional[pulumi.Input[Sequence[pulumi.Input[Mapping[str, Any]]]]] = None,
                 file: Optional[pulumi.Input[Union[pulumi.Asset, pulumi.Archive]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a resource file. Specify the `Content-Type` as `application/octet-stream` or `multipart/form-data`. For more information about resource files, see [Resource files](https://cloud.google.com/apigee/docs/api-platform/develop/resource-files).
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] content_type: The HTTP Content-Type header value specifying the content type of the body.
        :param pulumi.Input[str] data: The HTTP request/response body as raw binary.
        :param pulumi.Input[Sequence[pulumi.Input[Mapping[str, Any]]]] extensions: Application specific response metadata. Must be set in the first response for streaming APIs.
        :param pulumi.Input[Union[pulumi.Asset, pulumi.Archive]] file: File to upload.
        :param pulumi.Input[str] name: Required. Name of the resource file. Must match the regular expression: [a-zA-Z0-9:/\\\\!@#$%^&{}\\[\\]()+\\-=,.~'` ]{1,255}
        :param pulumi.Input[str] type: Required. Resource file type. {{ resource_file_type }}
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ResourcefileArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a resource file. Specify the `Content-Type` as `application/octet-stream` or `multipart/form-data`. For more information about resource files, see [Resource files](https://cloud.google.com/apigee/docs/api-platform/develop/resource-files).
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param ResourcefileArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ResourcefileArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 content_type: Optional[pulumi.Input[str]] = None,
                 data: Optional[pulumi.Input[str]] = None,
                 environment_id: Optional[pulumi.Input[str]] = None,
                 extensions: Optional[pulumi.Input[Sequence[pulumi.Input[Mapping[str, Any]]]]] = None,
                 file: Optional[pulumi.Input[Union[pulumi.Asset, pulumi.Archive]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ResourcefileArgs.__new__(ResourcefileArgs)

            __props__.__dict__["content_type"] = content_type
            __props__.__dict__["data"] = data
            if environment_id is None and not opts.urn:
                raise TypeError("Missing required property 'environment_id'")
            __props__.__dict__["environment_id"] = environment_id
            __props__.__dict__["extensions"] = extensions
            __props__.__dict__["file"] = file
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            if organization_id is None and not opts.urn:
                raise TypeError("Missing required property 'organization_id'")
            __props__.__dict__["organization_id"] = organization_id
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["environmentId", "name", "organizationId", "type"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Resourcefile, __self__).__init__(
            'google-native:apigee/v1:Resourcefile',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Resourcefile':
        """
        Get an existing Resourcefile resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ResourcefileArgs.__new__(ResourcefileArgs)

        __props__.__dict__["content_type"] = None
        __props__.__dict__["data"] = None
        __props__.__dict__["environment_id"] = None
        __props__.__dict__["extensions"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["organization_id"] = None
        __props__.__dict__["type"] = None
        return Resourcefile(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> pulumi.Output[str]:
        """
        The HTTP Content-Type header value specifying the content type of the body.
        """
        return pulumi.get(self, "content_type")

    @property
    @pulumi.getter
    def data(self) -> pulumi.Output[str]:
        """
        The HTTP request/response body as raw binary.
        """
        return pulumi.get(self, "data")

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "environment_id")

    @property
    @pulumi.getter
    def extensions(self) -> pulumi.Output[Sequence[Mapping[str, Any]]]:
        """
        Application specific response metadata. Must be set in the first response for streaming APIs.
        """
        return pulumi.get(self, "extensions")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Required. Name of the resource file. Must match the regular expression: [a-zA-Z0-9:/\\\\!@#$%^&{}\\[\\]()+\\-=,.~'` ]{1,255}
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="organizationId")
    def organization_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "organization_id")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        Required. Resource file type. {{ resource_file_type }}
        """
        return pulumi.get(self, "type")

