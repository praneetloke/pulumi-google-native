# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs
from ._inputs import *

__all__ = ['AgentEnvironmentUserSessionEntityTypeArgs', 'AgentEnvironmentUserSessionEntityType']

@pulumi.input_type
class AgentEnvironmentUserSessionEntityTypeArgs:
    def __init__(__self__, *,
                 entity_type_id: pulumi.Input[str],
                 environment_id: pulumi.Input[str],
                 location: pulumi.Input[str],
                 project: pulumi.Input[str],
                 session_id: pulumi.Input[str],
                 user_id: pulumi.Input[str],
                 entities: Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudDialogflowV2EntityTypeEntityArgs']]]] = None,
                 entity_override_mode: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AgentEnvironmentUserSessionEntityType resource.
        :param pulumi.Input[Sequence[pulumi.Input['GoogleCloudDialogflowV2EntityTypeEntityArgs']]] entities: Required. The collection of entities associated with this session entity type.
        :param pulumi.Input[str] entity_override_mode: Required. Indicates whether the additional data should override or supplement the custom entity type definition.
        :param pulumi.Input[str] name: Required. The unique identifier of this session entity type. Format: `projects//agent/sessions//entityTypes/`, or `projects//agent/environments//users//sessions//entityTypes/`. If `Environment ID` is not specified, we assume default 'draft' environment. If `User ID` is not specified, we assume default '-' user. `` must be the display name of an existing entity type in the same agent that will be overridden or supplemented.
        """
        pulumi.set(__self__, "entity_type_id", entity_type_id)
        pulumi.set(__self__, "environment_id", environment_id)
        pulumi.set(__self__, "location", location)
        pulumi.set(__self__, "project", project)
        pulumi.set(__self__, "session_id", session_id)
        pulumi.set(__self__, "user_id", user_id)
        if entities is not None:
            pulumi.set(__self__, "entities", entities)
        if entity_override_mode is not None:
            pulumi.set(__self__, "entity_override_mode", entity_override_mode)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter(name="entityTypeId")
    def entity_type_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "entity_type_id")

    @entity_type_id.setter
    def entity_type_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "entity_type_id", value)

    @property
    @pulumi.getter(name="environmentId")
    def environment_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "environment_id")

    @environment_id.setter
    def environment_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "environment_id", value)

    @property
    @pulumi.getter
    def location(self) -> pulumi.Input[str]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: pulumi.Input[str]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter
    def project(self) -> pulumi.Input[str]:
        return pulumi.get(self, "project")

    @project.setter
    def project(self, value: pulumi.Input[str]):
        pulumi.set(self, "project", value)

    @property
    @pulumi.getter(name="sessionId")
    def session_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "session_id")

    @session_id.setter
    def session_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "session_id", value)

    @property
    @pulumi.getter(name="userId")
    def user_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "user_id")

    @user_id.setter
    def user_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "user_id", value)

    @property
    @pulumi.getter
    def entities(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudDialogflowV2EntityTypeEntityArgs']]]]:
        """
        Required. The collection of entities associated with this session entity type.
        """
        return pulumi.get(self, "entities")

    @entities.setter
    def entities(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudDialogflowV2EntityTypeEntityArgs']]]]):
        pulumi.set(self, "entities", value)

    @property
    @pulumi.getter(name="entityOverrideMode")
    def entity_override_mode(self) -> Optional[pulumi.Input[str]]:
        """
        Required. Indicates whether the additional data should override or supplement the custom entity type definition.
        """
        return pulumi.get(self, "entity_override_mode")

    @entity_override_mode.setter
    def entity_override_mode(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "entity_override_mode", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        Required. The unique identifier of this session entity type. Format: `projects//agent/sessions//entityTypes/`, or `projects//agent/environments//users//sessions//entityTypes/`. If `Environment ID` is not specified, we assume default 'draft' environment. If `User ID` is not specified, we assume default '-' user. `` must be the display name of an existing entity type in the same agent that will be overridden or supplemented.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class AgentEnvironmentUserSessionEntityType(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 entities: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowV2EntityTypeEntityArgs']]]]] = None,
                 entity_override_mode: Optional[pulumi.Input[str]] = None,
                 entity_type_id: Optional[pulumi.Input[str]] = None,
                 environment_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 session_id: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a session entity type. If the specified session entity type already exists, overrides the session entity type. This method doesn't work with Google Assistant integration. Contact Dialogflow support if you need to use session entities with Google Assistant integration.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowV2EntityTypeEntityArgs']]]] entities: Required. The collection of entities associated with this session entity type.
        :param pulumi.Input[str] entity_override_mode: Required. Indicates whether the additional data should override or supplement the custom entity type definition.
        :param pulumi.Input[str] name: Required. The unique identifier of this session entity type. Format: `projects//agent/sessions//entityTypes/`, or `projects//agent/environments//users//sessions//entityTypes/`. If `Environment ID` is not specified, we assume default 'draft' environment. If `User ID` is not specified, we assume default '-' user. `` must be the display name of an existing entity type in the same agent that will be overridden or supplemented.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AgentEnvironmentUserSessionEntityTypeArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a session entity type. If the specified session entity type already exists, overrides the session entity type. This method doesn't work with Google Assistant integration. Contact Dialogflow support if you need to use session entities with Google Assistant integration.

        :param str resource_name: The name of the resource.
        :param AgentEnvironmentUserSessionEntityTypeArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AgentEnvironmentUserSessionEntityTypeArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 entities: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowV2EntityTypeEntityArgs']]]]] = None,
                 entity_override_mode: Optional[pulumi.Input[str]] = None,
                 entity_type_id: Optional[pulumi.Input[str]] = None,
                 environment_id: Optional[pulumi.Input[str]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 session_id: Optional[pulumi.Input[str]] = None,
                 user_id: Optional[pulumi.Input[str]] = None,
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
            __props__ = AgentEnvironmentUserSessionEntityTypeArgs.__new__(AgentEnvironmentUserSessionEntityTypeArgs)

            __props__.__dict__["entities"] = entities
            __props__.__dict__["entity_override_mode"] = entity_override_mode
            if entity_type_id is None and not opts.urn:
                raise TypeError("Missing required property 'entity_type_id'")
            __props__.__dict__["entity_type_id"] = entity_type_id
            if environment_id is None and not opts.urn:
                raise TypeError("Missing required property 'environment_id'")
            __props__.__dict__["environment_id"] = environment_id
            if location is None and not opts.urn:
                raise TypeError("Missing required property 'location'")
            __props__.__dict__["location"] = location
            __props__.__dict__["name"] = name
            if project is None and not opts.urn:
                raise TypeError("Missing required property 'project'")
            __props__.__dict__["project"] = project
            if session_id is None and not opts.urn:
                raise TypeError("Missing required property 'session_id'")
            __props__.__dict__["session_id"] = session_id
            if user_id is None and not opts.urn:
                raise TypeError("Missing required property 'user_id'")
            __props__.__dict__["user_id"] = user_id
        super(AgentEnvironmentUserSessionEntityType, __self__).__init__(
            'google-native:dialogflow/v2:AgentEnvironmentUserSessionEntityType',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AgentEnvironmentUserSessionEntityType':
        """
        Get an existing AgentEnvironmentUserSessionEntityType resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AgentEnvironmentUserSessionEntityTypeArgs.__new__(AgentEnvironmentUserSessionEntityTypeArgs)

        __props__.__dict__["entities"] = None
        __props__.__dict__["entity_override_mode"] = None
        __props__.__dict__["name"] = None
        return AgentEnvironmentUserSessionEntityType(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def entities(self) -> pulumi.Output[Sequence['outputs.GoogleCloudDialogflowV2EntityTypeEntityResponse']]:
        """
        Required. The collection of entities associated with this session entity type.
        """
        return pulumi.get(self, "entities")

    @property
    @pulumi.getter(name="entityOverrideMode")
    def entity_override_mode(self) -> pulumi.Output[str]:
        """
        Required. Indicates whether the additional data should override or supplement the custom entity type definition.
        """
        return pulumi.get(self, "entity_override_mode")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Required. The unique identifier of this session entity type. Format: `projects//agent/sessions//entityTypes/`, or `projects//agent/environments//users//sessions//entityTypes/`. If `Environment ID` is not specified, we assume default 'draft' environment. If `User ID` is not specified, we assume default '-' user. `` must be the display name of an existing entity type in the same agent that will be overridden or supplemented.
        """
        return pulumi.get(self, "name")

