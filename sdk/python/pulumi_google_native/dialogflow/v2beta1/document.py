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

__all__ = ['DocumentArgs', 'Document']

@pulumi.input_type
class DocumentArgs:
    def __init__(__self__, *,
                 display_name: pulumi.Input[str],
                 knowledge_base_id: pulumi.Input[str],
                 knowledge_types: pulumi.Input[Sequence[pulumi.Input['DocumentKnowledgeTypesItem']]],
                 mime_type: pulumi.Input[str],
                 content: Optional[pulumi.Input[str]] = None,
                 content_uri: Optional[pulumi.Input[str]] = None,
                 enable_auto_reload: Optional[pulumi.Input[bool]] = None,
                 import_gcs_custom_metadata: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 raw_content: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Document resource.
        :param pulumi.Input[str] display_name: The display name of the document. The name must be 1024 bytes or less; otherwise, the creation request fails.
        :param pulumi.Input[Sequence[pulumi.Input['DocumentKnowledgeTypesItem']]] knowledge_types: The knowledge type of document content.
        :param pulumi.Input[str] mime_type: The MIME type of this document.
        :param pulumi.Input[str] content: The raw content of the document. This field is only permitted for EXTRACTIVE_QA and FAQ knowledge types. Note: This field is in the process of being deprecated, please use raw_content instead.
        :param pulumi.Input[str] content_uri: The URI where the file content is located. For documents stored in Google Cloud Storage, these URIs must have the form `gs:///`. NOTE: External URLs must correspond to public webpages, i.e., they must be indexed by Google Search. In particular, URLs for showing documents in Google Cloud Storage (i.e. the URL in your browser) are not supported. Instead use the `gs://` format URI described above.
        :param pulumi.Input[bool] enable_auto_reload: Optional. If true, we try to automatically reload the document every day (at a time picked by the system). If false or unspecified, we don't try to automatically reload the document. Currently you can only enable automatic reload for documents sourced from a public url, see `source` field for the source types. Reload status can be tracked in `latest_reload_status`. If a reload fails, we will keep the document unchanged. If a reload fails with internal errors, the system will try to reload the document on the next day. If a reload fails with non-retriable errors (e.g. PERMISION_DENIED), the system will not try to reload the document anymore. You need to manually reload the document successfully by calling `ReloadDocument` and clear the errors.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: Optional. Metadata for the document. The metadata supports arbitrary key-value pairs. Suggested use cases include storing a document's title, an external URL distinct from the document's content_uri, etc. The max size of a `key` or a `value` of the metadata is 1024 bytes.
        :param pulumi.Input[str] name: Optional. The document resource name. The name must be empty when creating a document. Format: `projects//locations//knowledgeBases//documents/`.
        :param pulumi.Input[str] raw_content: The raw content of the document. This field is only permitted for EXTRACTIVE_QA and FAQ knowledge types.
        """
        pulumi.set(__self__, "display_name", display_name)
        pulumi.set(__self__, "knowledge_base_id", knowledge_base_id)
        pulumi.set(__self__, "knowledge_types", knowledge_types)
        pulumi.set(__self__, "mime_type", mime_type)
        if content is not None:
            pulumi.set(__self__, "content", content)
        if content_uri is not None:
            pulumi.set(__self__, "content_uri", content_uri)
        if enable_auto_reload is not None:
            pulumi.set(__self__, "enable_auto_reload", enable_auto_reload)
        if import_gcs_custom_metadata is not None:
            pulumi.set(__self__, "import_gcs_custom_metadata", import_gcs_custom_metadata)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if metadata is not None:
            pulumi.set(__self__, "metadata", metadata)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if raw_content is not None:
            pulumi.set(__self__, "raw_content", raw_content)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Input[str]:
        """
        The display name of the document. The name must be 1024 bytes or less; otherwise, the creation request fails.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="knowledgeBaseId")
    def knowledge_base_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "knowledge_base_id")

    @knowledge_base_id.setter
    def knowledge_base_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "knowledge_base_id", value)

    @property
    @pulumi.getter(name="knowledgeTypes")
    def knowledge_types(self) -> pulumi.Input[Sequence[pulumi.Input['DocumentKnowledgeTypesItem']]]:
        """
        The knowledge type of document content.
        """
        return pulumi.get(self, "knowledge_types")

    @knowledge_types.setter
    def knowledge_types(self, value: pulumi.Input[Sequence[pulumi.Input['DocumentKnowledgeTypesItem']]]):
        pulumi.set(self, "knowledge_types", value)

    @property
    @pulumi.getter(name="mimeType")
    def mime_type(self) -> pulumi.Input[str]:
        """
        The MIME type of this document.
        """
        return pulumi.get(self, "mime_type")

    @mime_type.setter
    def mime_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "mime_type", value)

    @property
    @pulumi.getter
    def content(self) -> Optional[pulumi.Input[str]]:
        """
        The raw content of the document. This field is only permitted for EXTRACTIVE_QA and FAQ knowledge types. Note: This field is in the process of being deprecated, please use raw_content instead.
        """
        return pulumi.get(self, "content")

    @content.setter
    def content(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content", value)

    @property
    @pulumi.getter(name="contentUri")
    def content_uri(self) -> Optional[pulumi.Input[str]]:
        """
        The URI where the file content is located. For documents stored in Google Cloud Storage, these URIs must have the form `gs:///`. NOTE: External URLs must correspond to public webpages, i.e., they must be indexed by Google Search. In particular, URLs for showing documents in Google Cloud Storage (i.e. the URL in your browser) are not supported. Instead use the `gs://` format URI described above.
        """
        return pulumi.get(self, "content_uri")

    @content_uri.setter
    def content_uri(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "content_uri", value)

    @property
    @pulumi.getter(name="enableAutoReload")
    def enable_auto_reload(self) -> Optional[pulumi.Input[bool]]:
        """
        Optional. If true, we try to automatically reload the document every day (at a time picked by the system). If false or unspecified, we don't try to automatically reload the document. Currently you can only enable automatic reload for documents sourced from a public url, see `source` field for the source types. Reload status can be tracked in `latest_reload_status`. If a reload fails, we will keep the document unchanged. If a reload fails with internal errors, the system will try to reload the document on the next day. If a reload fails with non-retriable errors (e.g. PERMISION_DENIED), the system will not try to reload the document anymore. You need to manually reload the document successfully by calling `ReloadDocument` and clear the errors.
        """
        return pulumi.get(self, "enable_auto_reload")

    @enable_auto_reload.setter
    def enable_auto_reload(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "enable_auto_reload", value)

    @property
    @pulumi.getter(name="importGcsCustomMetadata")
    def import_gcs_custom_metadata(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "import_gcs_custom_metadata")

    @import_gcs_custom_metadata.setter
    def import_gcs_custom_metadata(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "import_gcs_custom_metadata", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def metadata(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Optional. Metadata for the document. The metadata supports arbitrary key-value pairs. Suggested use cases include storing a document's title, an external URL distinct from the document's content_uri, etc. The max size of a `key` or a `value` of the metadata is 1024 bytes.
        """
        return pulumi.get(self, "metadata")

    @metadata.setter
    def metadata(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "metadata", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. The document resource name. The name must be empty when creating a document. Format: `projects//locations//knowledgeBases//documents/`.
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
    @pulumi.getter(name="rawContent")
    def raw_content(self) -> Optional[pulumi.Input[str]]:
        """
        The raw content of the document. This field is only permitted for EXTRACTIVE_QA and FAQ knowledge types.
        """
        return pulumi.get(self, "raw_content")

    @raw_content.setter
    def raw_content(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "raw_content", value)


class Document(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 content_uri: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 enable_auto_reload: Optional[pulumi.Input[bool]] = None,
                 import_gcs_custom_metadata: Optional[pulumi.Input[str]] = None,
                 knowledge_base_id: Optional[pulumi.Input[str]] = None,
                 knowledge_types: Optional[pulumi.Input[Sequence[pulumi.Input['DocumentKnowledgeTypesItem']]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 mime_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 raw_content: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a new document. Note: The `projects.agent.knowledgeBases.documents` resource is deprecated; only use `projects.knowledgeBases.documents`.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] content: The raw content of the document. This field is only permitted for EXTRACTIVE_QA and FAQ knowledge types. Note: This field is in the process of being deprecated, please use raw_content instead.
        :param pulumi.Input[str] content_uri: The URI where the file content is located. For documents stored in Google Cloud Storage, these URIs must have the form `gs:///`. NOTE: External URLs must correspond to public webpages, i.e., they must be indexed by Google Search. In particular, URLs for showing documents in Google Cloud Storage (i.e. the URL in your browser) are not supported. Instead use the `gs://` format URI described above.
        :param pulumi.Input[str] display_name: The display name of the document. The name must be 1024 bytes or less; otherwise, the creation request fails.
        :param pulumi.Input[bool] enable_auto_reload: Optional. If true, we try to automatically reload the document every day (at a time picked by the system). If false or unspecified, we don't try to automatically reload the document. Currently you can only enable automatic reload for documents sourced from a public url, see `source` field for the source types. Reload status can be tracked in `latest_reload_status`. If a reload fails, we will keep the document unchanged. If a reload fails with internal errors, the system will try to reload the document on the next day. If a reload fails with non-retriable errors (e.g. PERMISION_DENIED), the system will not try to reload the document anymore. You need to manually reload the document successfully by calling `ReloadDocument` and clear the errors.
        :param pulumi.Input[Sequence[pulumi.Input['DocumentKnowledgeTypesItem']]] knowledge_types: The knowledge type of document content.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] metadata: Optional. Metadata for the document. The metadata supports arbitrary key-value pairs. Suggested use cases include storing a document's title, an external URL distinct from the document's content_uri, etc. The max size of a `key` or a `value` of the metadata is 1024 bytes.
        :param pulumi.Input[str] mime_type: The MIME type of this document.
        :param pulumi.Input[str] name: Optional. The document resource name. The name must be empty when creating a document. Format: `projects//locations//knowledgeBases//documents/`.
        :param pulumi.Input[str] raw_content: The raw content of the document. This field is only permitted for EXTRACTIVE_QA and FAQ knowledge types.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: DocumentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new document. Note: The `projects.agent.knowledgeBases.documents` resource is deprecated; only use `projects.knowledgeBases.documents`.

        :param str resource_name: The name of the resource.
        :param DocumentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(DocumentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 content: Optional[pulumi.Input[str]] = None,
                 content_uri: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 enable_auto_reload: Optional[pulumi.Input[bool]] = None,
                 import_gcs_custom_metadata: Optional[pulumi.Input[str]] = None,
                 knowledge_base_id: Optional[pulumi.Input[str]] = None,
                 knowledge_types: Optional[pulumi.Input[Sequence[pulumi.Input['DocumentKnowledgeTypesItem']]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 metadata: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 mime_type: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 raw_content: Optional[pulumi.Input[str]] = None,
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
            __props__ = DocumentArgs.__new__(DocumentArgs)

            __props__.__dict__["content"] = content
            __props__.__dict__["content_uri"] = content_uri
            if display_name is None and not opts.urn:
                raise TypeError("Missing required property 'display_name'")
            __props__.__dict__["display_name"] = display_name
            __props__.__dict__["enable_auto_reload"] = enable_auto_reload
            __props__.__dict__["import_gcs_custom_metadata"] = import_gcs_custom_metadata
            if knowledge_base_id is None and not opts.urn:
                raise TypeError("Missing required property 'knowledge_base_id'")
            __props__.__dict__["knowledge_base_id"] = knowledge_base_id
            if knowledge_types is None and not opts.urn:
                raise TypeError("Missing required property 'knowledge_types'")
            __props__.__dict__["knowledge_types"] = knowledge_types
            __props__.__dict__["location"] = location
            __props__.__dict__["metadata"] = metadata
            if mime_type is None and not opts.urn:
                raise TypeError("Missing required property 'mime_type'")
            __props__.__dict__["mime_type"] = mime_type
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            __props__.__dict__["raw_content"] = raw_content
            __props__.__dict__["latest_reload_status"] = None
        super(Document, __self__).__init__(
            'google-native:dialogflow/v2beta1:Document',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Document':
        """
        Get an existing Document resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = DocumentArgs.__new__(DocumentArgs)

        __props__.__dict__["content"] = None
        __props__.__dict__["content_uri"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["enable_auto_reload"] = None
        __props__.__dict__["knowledge_types"] = None
        __props__.__dict__["latest_reload_status"] = None
        __props__.__dict__["metadata"] = None
        __props__.__dict__["mime_type"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["raw_content"] = None
        return Document(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def content(self) -> pulumi.Output[str]:
        """
        The raw content of the document. This field is only permitted for EXTRACTIVE_QA and FAQ knowledge types. Note: This field is in the process of being deprecated, please use raw_content instead.
        """
        return pulumi.get(self, "content")

    @property
    @pulumi.getter(name="contentUri")
    def content_uri(self) -> pulumi.Output[str]:
        """
        The URI where the file content is located. For documents stored in Google Cloud Storage, these URIs must have the form `gs:///`. NOTE: External URLs must correspond to public webpages, i.e., they must be indexed by Google Search. In particular, URLs for showing documents in Google Cloud Storage (i.e. the URL in your browser) are not supported. Instead use the `gs://` format URI described above.
        """
        return pulumi.get(self, "content_uri")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        The display name of the document. The name must be 1024 bytes or less; otherwise, the creation request fails.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="enableAutoReload")
    def enable_auto_reload(self) -> pulumi.Output[bool]:
        """
        Optional. If true, we try to automatically reload the document every day (at a time picked by the system). If false or unspecified, we don't try to automatically reload the document. Currently you can only enable automatic reload for documents sourced from a public url, see `source` field for the source types. Reload status can be tracked in `latest_reload_status`. If a reload fails, we will keep the document unchanged. If a reload fails with internal errors, the system will try to reload the document on the next day. If a reload fails with non-retriable errors (e.g. PERMISION_DENIED), the system will not try to reload the document anymore. You need to manually reload the document successfully by calling `ReloadDocument` and clear the errors.
        """
        return pulumi.get(self, "enable_auto_reload")

    @property
    @pulumi.getter(name="knowledgeTypes")
    def knowledge_types(self) -> pulumi.Output[Sequence[str]]:
        """
        The knowledge type of document content.
        """
        return pulumi.get(self, "knowledge_types")

    @property
    @pulumi.getter(name="latestReloadStatus")
    def latest_reload_status(self) -> pulumi.Output['outputs.GoogleCloudDialogflowV2beta1DocumentReloadStatusResponse']:
        """
        The time and status of the latest reload. This reload may have been triggered automatically or manually and may not have succeeded.
        """
        return pulumi.get(self, "latest_reload_status")

    @property
    @pulumi.getter
    def metadata(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Optional. Metadata for the document. The metadata supports arbitrary key-value pairs. Suggested use cases include storing a document's title, an external URL distinct from the document's content_uri, etc. The max size of a `key` or a `value` of the metadata is 1024 bytes.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter(name="mimeType")
    def mime_type(self) -> pulumi.Output[str]:
        """
        The MIME type of this document.
        """
        return pulumi.get(self, "mime_type")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Optional. The document resource name. The name must be empty when creating a document. Format: `projects//locations//knowledgeBases//documents/`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="rawContent")
    def raw_content(self) -> pulumi.Output[str]:
        """
        The raw content of the document. This field is only permitted for EXTRACTIVE_QA and FAQ knowledge types.
        """
        return pulumi.get(self, "raw_content")

