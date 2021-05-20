# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = ['OrganizationSharedflowArgs', 'OrganizationSharedflow']

@pulumi.input_type
class OrganizationSharedflowArgs:
    def __init__(__self__, *,
                 action: pulumi.Input[str],
                 name: pulumi.Input[str],
                 organization_id: pulumi.Input[str],
                 sharedflow_id: pulumi.Input[str],
                 content_type: Optional[pulumi.Input[str]] = None,
                 data: Optional[pulumi.Input[str]] = None,
                 extensions: Optional[pulumi.Input[Sequence[pulumi.Input[Mapping[str, pulumi.Input[str]]]]]] = None):
        """
        The set of arguments for constructing a OrganizationSharedflow resource.
        :param pulumi.Input[str] content_type: The HTTP Content-Type header value specifying the content type of the body.
        :param pulumi.Input[str] data: The HTTP request/response body as raw binary.
        :param pulumi.Input[Sequence[pulumi.Input[Mapping[str, pulumi.Input[str]]]]] extensions: Application specific response metadata. Must be set in the first response for streaming APIs.
        """
        pulumi.set(__self__, "action", action)
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "organization_id", organization_id)
        pulumi.set(__self__, "sharedflow_id", sharedflow_id)
        if content_type is not None:
            pulumi.set(__self__, "content_type", content_type)
        if data is not None:
            pulumi.set(__self__, "data", data)
        if extensions is not None:
            pulumi.set(__self__, "extensions", extensions)

    @property
    @pulumi.getter
    def action(self) -> pulumi.Input[str]:
        return pulumi.get(self, "action")

    @action.setter
    def action(self, value: pulumi.Input[str]):
        pulumi.set(self, "action", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
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
    @pulumi.getter(name="sharedflowId")
    def sharedflow_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "sharedflow_id")

    @sharedflow_id.setter
    def sharedflow_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "sharedflow_id", value)

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
    def extensions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input[Mapping[str, pulumi.Input[str]]]]]]:
        """
        Application specific response metadata. Must be set in the first response for streaming APIs.
        """
        return pulumi.get(self, "extensions")

    @extensions.setter
    def extensions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input[Mapping[str, pulumi.Input[str]]]]]]):
        pulumi.set(self, "extensions", value)


class OrganizationSharedflow(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 content_type: Optional[pulumi.Input[str]] = None,
                 data: Optional[pulumi.Input[str]] = None,
                 extensions: Optional[pulumi.Input[Sequence[pulumi.Input[Mapping[str, pulumi.Input[str]]]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 sharedflow_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Uploads a ZIP-formatted shared flow configuration bundle to an organization. If the shared flow already exists, this creates a new revision of it. If the shared flow does not exist, this creates it. Once imported, the shared flow revision must be deployed before it can be accessed at runtime. The size limit of a shared flow bundle is 15 MB.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] content_type: The HTTP Content-Type header value specifying the content type of the body.
        :param pulumi.Input[str] data: The HTTP request/response body as raw binary.
        :param pulumi.Input[Sequence[pulumi.Input[Mapping[str, pulumi.Input[str]]]]] extensions: Application specific response metadata. Must be set in the first response for streaming APIs.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: OrganizationSharedflowArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Uploads a ZIP-formatted shared flow configuration bundle to an organization. If the shared flow already exists, this creates a new revision of it. If the shared flow does not exist, this creates it. Once imported, the shared flow revision must be deployed before it can be accessed at runtime. The size limit of a shared flow bundle is 15 MB.

        :param str resource_name: The name of the resource.
        :param OrganizationSharedflowArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(OrganizationSharedflowArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 action: Optional[pulumi.Input[str]] = None,
                 content_type: Optional[pulumi.Input[str]] = None,
                 data: Optional[pulumi.Input[str]] = None,
                 extensions: Optional[pulumi.Input[Sequence[pulumi.Input[Mapping[str, pulumi.Input[str]]]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 organization_id: Optional[pulumi.Input[str]] = None,
                 sharedflow_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = OrganizationSharedflowArgs.__new__(OrganizationSharedflowArgs)

            if action is None and not opts.urn:
                raise TypeError("Missing required property 'action'")
            __props__.__dict__["action"] = action
            __props__.__dict__["content_type"] = content_type
            __props__.__dict__["data"] = data
            __props__.__dict__["extensions"] = extensions
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            if organization_id is None and not opts.urn:
                raise TypeError("Missing required property 'organization_id'")
            __props__.__dict__["organization_id"] = organization_id
            if sharedflow_id is None and not opts.urn:
                raise TypeError("Missing required property 'sharedflow_id'")
            __props__.__dict__["sharedflow_id"] = sharedflow_id
            __props__.__dict__["latest_revision_id"] = None
            __props__.__dict__["meta_data"] = None
            __props__.__dict__["revision"] = None
        super(OrganizationSharedflow, __self__).__init__(
            'google-native:apigee/v1:OrganizationSharedflow',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'OrganizationSharedflow':
        """
        Get an existing OrganizationSharedflow resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = OrganizationSharedflowArgs.__new__(OrganizationSharedflowArgs)

        __props__.__dict__["latest_revision_id"] = None
        __props__.__dict__["meta_data"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["revision"] = None
        return OrganizationSharedflow(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="latestRevisionId")
    def latest_revision_id(self) -> pulumi.Output[str]:
        """
        The id of the most recently created revision for this shared flow.
        """
        return pulumi.get(self, "latest_revision_id")

    @property
    @pulumi.getter(name="metaData")
    def meta_data(self) -> pulumi.Output['outputs.GoogleCloudApigeeV1EntityMetadataResponse']:
        """
        Metadata describing the shared flow.
        """
        return pulumi.get(self, "meta_data")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The ID of the shared flow.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def revision(self) -> pulumi.Output[Sequence[str]]:
        """
        A list of revisions of this shared flow.
        """
        return pulumi.get(self, "revision")

