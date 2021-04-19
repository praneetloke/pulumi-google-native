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

__all__ = ['AgentIntentArgs', 'AgentIntent']

@pulumi.input_type
class AgentIntentArgs:
    def __init__(__self__, *,
                 agents_id: pulumi.Input[str],
                 intents_id: pulumi.Input[str],
                 locations_id: pulumi.Input[str],
                 projects_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 is_fallback: Optional[pulumi.Input[bool]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudDialogflowCxV3IntentParameterArgs']]]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 training_phrases: Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudDialogflowCxV3IntentTrainingPhraseArgs']]]] = None):
        """
        The set of arguments for constructing a AgentIntent resource.
        :param pulumi.Input[str] description: Human readable description for better understanding an intent like its scope, content, result etc. Maximum character limit: 140 characters.
        :param pulumi.Input[str] display_name: Required. The human-readable name of the intent, unique within the agent.
        :param pulumi.Input[bool] is_fallback: Indicates whether this is a fallback intent. Currently only default fallback intent is allowed in the agent, which is added upon agent creation. Adding training phrases to fallback intent is useful in the case of requests that are mistakenly matched, since training phrases assigned to fallback intents act as negative examples that triggers no-match event.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: The key/value metadata to label an intent. Labels can contain lowercase letters, digits and the symbols '-' and '_'. International characters are allowed, including letters from unicase alphabets. Keys must start with a letter. Keys and values can be no longer than 63 characters and no more than 128 bytes. Prefix "sys." is reserved for Dialogflow defined labels. Currently allowed Dialogflow defined labels include: * sys.head * sys.contextual The above labels do not require value. "sys.head" means the intent is a head intent. "sys.contextual" means the intent is a contextual intent.
        :param pulumi.Input[str] name: The unique identifier of the intent. Required for the Intents.UpdateIntent method. Intents.CreateIntent populates the name automatically. Format: `projects//locations//agents//intents/`.
        :param pulumi.Input[Sequence[pulumi.Input['GoogleCloudDialogflowCxV3IntentParameterArgs']]] parameters: The collection of parameters associated with the intent.
        :param pulumi.Input[int] priority: The priority of this intent. Higher numbers represent higher priorities. - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the `Normal` priority in the console. - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
        :param pulumi.Input[Sequence[pulumi.Input['GoogleCloudDialogflowCxV3IntentTrainingPhraseArgs']]] training_phrases: The collection of training phrases the agent is trained on to identify the intent.
        """
        pulumi.set(__self__, "agents_id", agents_id)
        pulumi.set(__self__, "intents_id", intents_id)
        pulumi.set(__self__, "locations_id", locations_id)
        pulumi.set(__self__, "projects_id", projects_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if display_name is not None:
            pulumi.set(__self__, "display_name", display_name)
        if is_fallback is not None:
            pulumi.set(__self__, "is_fallback", is_fallback)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if parameters is not None:
            pulumi.set(__self__, "parameters", parameters)
        if priority is not None:
            pulumi.set(__self__, "priority", priority)
        if training_phrases is not None:
            pulumi.set(__self__, "training_phrases", training_phrases)

    @property
    @pulumi.getter(name="agentsId")
    def agents_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "agents_id")

    @agents_id.setter
    def agents_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "agents_id", value)

    @property
    @pulumi.getter(name="intentsId")
    def intents_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "intents_id")

    @intents_id.setter
    def intents_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "intents_id", value)

    @property
    @pulumi.getter(name="locationsId")
    def locations_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "locations_id")

    @locations_id.setter
    def locations_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "locations_id", value)

    @property
    @pulumi.getter(name="projectsId")
    def projects_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "projects_id")

    @projects_id.setter
    def projects_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "projects_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Human readable description for better understanding an intent like its scope, content, result etc. Maximum character limit: 140 characters.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> Optional[pulumi.Input[str]]:
        """
        Required. The human-readable name of the intent, unique within the agent.
        """
        return pulumi.get(self, "display_name")

    @display_name.setter
    def display_name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "display_name", value)

    @property
    @pulumi.getter(name="isFallback")
    def is_fallback(self) -> Optional[pulumi.Input[bool]]:
        """
        Indicates whether this is a fallback intent. Currently only default fallback intent is allowed in the agent, which is added upon agent creation. Adding training phrases to fallback intent is useful in the case of requests that are mistakenly matched, since training phrases assigned to fallback intents act as negative examples that triggers no-match event.
        """
        return pulumi.get(self, "is_fallback")

    @is_fallback.setter
    def is_fallback(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "is_fallback", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        The key/value metadata to label an intent. Labels can contain lowercase letters, digits and the symbols '-' and '_'. International characters are allowed, including letters from unicase alphabets. Keys must start with a letter. Keys and values can be no longer than 63 characters and no more than 128 bytes. Prefix "sys." is reserved for Dialogflow defined labels. Currently allowed Dialogflow defined labels include: * sys.head * sys.contextual The above labels do not require value. "sys.head" means the intent is a head intent. "sys.contextual" means the intent is a contextual intent.
        """
        return pulumi.get(self, "labels")

    @labels.setter
    def labels(self, value: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]):
        pulumi.set(self, "labels", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The unique identifier of the intent. Required for the Intents.UpdateIntent method. Intents.CreateIntent populates the name automatically. Format: `projects//locations//agents//intents/`.
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def parameters(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudDialogflowCxV3IntentParameterArgs']]]]:
        """
        The collection of parameters associated with the intent.
        """
        return pulumi.get(self, "parameters")

    @parameters.setter
    def parameters(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudDialogflowCxV3IntentParameterArgs']]]]):
        pulumi.set(self, "parameters", value)

    @property
    @pulumi.getter
    def priority(self) -> Optional[pulumi.Input[int]]:
        """
        The priority of this intent. Higher numbers represent higher priorities. - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the `Normal` priority in the console. - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
        """
        return pulumi.get(self, "priority")

    @priority.setter
    def priority(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "priority", value)

    @property
    @pulumi.getter(name="trainingPhrases")
    def training_phrases(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudDialogflowCxV3IntentTrainingPhraseArgs']]]]:
        """
        The collection of training phrases the agent is trained on to identify the intent.
        """
        return pulumi.get(self, "training_phrases")

    @training_phrases.setter
    def training_phrases(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GoogleCloudDialogflowCxV3IntentTrainingPhraseArgs']]]]):
        pulumi.set(self, "training_phrases", value)


class AgentIntent(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agents_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 intents_id: Optional[pulumi.Input[str]] = None,
                 is_fallback: Optional[pulumi.Input[bool]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 locations_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3IntentParameterArgs']]]]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 training_phrases: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3IntentTrainingPhraseArgs']]]]] = None,
                 __props__=None):
        """
        Creates an intent in the specified agent.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: Human readable description for better understanding an intent like its scope, content, result etc. Maximum character limit: 140 characters.
        :param pulumi.Input[str] display_name: Required. The human-readable name of the intent, unique within the agent.
        :param pulumi.Input[bool] is_fallback: Indicates whether this is a fallback intent. Currently only default fallback intent is allowed in the agent, which is added upon agent creation. Adding training phrases to fallback intent is useful in the case of requests that are mistakenly matched, since training phrases assigned to fallback intents act as negative examples that triggers no-match event.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: The key/value metadata to label an intent. Labels can contain lowercase letters, digits and the symbols '-' and '_'. International characters are allowed, including letters from unicase alphabets. Keys must start with a letter. Keys and values can be no longer than 63 characters and no more than 128 bytes. Prefix "sys." is reserved for Dialogflow defined labels. Currently allowed Dialogflow defined labels include: * sys.head * sys.contextual The above labels do not require value. "sys.head" means the intent is a head intent. "sys.contextual" means the intent is a contextual intent.
        :param pulumi.Input[str] name: The unique identifier of the intent. Required for the Intents.UpdateIntent method. Intents.CreateIntent populates the name automatically. Format: `projects//locations//agents//intents/`.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3IntentParameterArgs']]]] parameters: The collection of parameters associated with the intent.
        :param pulumi.Input[int] priority: The priority of this intent. Higher numbers represent higher priorities. - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the `Normal` priority in the console. - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
        :param pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3IntentTrainingPhraseArgs']]]] training_phrases: The collection of training phrases the agent is trained on to identify the intent.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AgentIntentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates an intent in the specified agent.

        :param str resource_name: The name of the resource.
        :param AgentIntentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AgentIntentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 agents_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 display_name: Optional[pulumi.Input[str]] = None,
                 intents_id: Optional[pulumi.Input[str]] = None,
                 is_fallback: Optional[pulumi.Input[bool]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 locations_id: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 parameters: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3IntentParameterArgs']]]]] = None,
                 priority: Optional[pulumi.Input[int]] = None,
                 projects_id: Optional[pulumi.Input[str]] = None,
                 training_phrases: Optional[pulumi.Input[Sequence[pulumi.Input[pulumi.InputType['GoogleCloudDialogflowCxV3IntentTrainingPhraseArgs']]]]] = None,
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
            __props__ = AgentIntentArgs.__new__(AgentIntentArgs)

            if agents_id is None and not opts.urn:
                raise TypeError("Missing required property 'agents_id'")
            __props__.__dict__["agents_id"] = agents_id
            __props__.__dict__["description"] = description
            __props__.__dict__["display_name"] = display_name
            if intents_id is None and not opts.urn:
                raise TypeError("Missing required property 'intents_id'")
            __props__.__dict__["intents_id"] = intents_id
            __props__.__dict__["is_fallback"] = is_fallback
            __props__.__dict__["labels"] = labels
            if locations_id is None and not opts.urn:
                raise TypeError("Missing required property 'locations_id'")
            __props__.__dict__["locations_id"] = locations_id
            __props__.__dict__["name"] = name
            __props__.__dict__["parameters"] = parameters
            __props__.__dict__["priority"] = priority
            if projects_id is None and not opts.urn:
                raise TypeError("Missing required property 'projects_id'")
            __props__.__dict__["projects_id"] = projects_id
            __props__.__dict__["training_phrases"] = training_phrases
        super(AgentIntent, __self__).__init__(
            'google-native:dialogflow/v3:AgentIntent',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'AgentIntent':
        """
        Get an existing AgentIntent resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AgentIntentArgs.__new__(AgentIntentArgs)

        __props__.__dict__["description"] = None
        __props__.__dict__["display_name"] = None
        __props__.__dict__["is_fallback"] = None
        __props__.__dict__["labels"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["parameters"] = None
        __props__.__dict__["priority"] = None
        __props__.__dict__["training_phrases"] = None
        return AgentIntent(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Human readable description for better understanding an intent like its scope, content, result etc. Maximum character limit: 140 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> pulumi.Output[str]:
        """
        Required. The human-readable name of the intent, unique within the agent.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter(name="isFallback")
    def is_fallback(self) -> pulumi.Output[bool]:
        """
        Indicates whether this is a fallback intent. Currently only default fallback intent is allowed in the agent, which is added upon agent creation. Adding training phrases to fallback intent is useful in the case of requests that are mistakenly matched, since training phrases assigned to fallback intents act as negative examples that triggers no-match event.
        """
        return pulumi.get(self, "is_fallback")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        The key/value metadata to label an intent. Labels can contain lowercase letters, digits and the symbols '-' and '_'. International characters are allowed, including letters from unicase alphabets. Keys must start with a letter. Keys and values can be no longer than 63 characters and no more than 128 bytes. Prefix "sys." is reserved for Dialogflow defined labels. Currently allowed Dialogflow defined labels include: * sys.head * sys.contextual The above labels do not require value. "sys.head" means the intent is a head intent. "sys.contextual" means the intent is a contextual intent.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The unique identifier of the intent. Required for the Intents.UpdateIntent method. Intents.CreateIntent populates the name automatically. Format: `projects//locations//agents//intents/`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def parameters(self) -> pulumi.Output[Sequence['outputs.GoogleCloudDialogflowCxV3IntentParameterResponse']]:
        """
        The collection of parameters associated with the intent.
        """
        return pulumi.get(self, "parameters")

    @property
    @pulumi.getter
    def priority(self) -> pulumi.Output[int]:
        """
        The priority of this intent. Higher numbers represent higher priorities. - If the supplied value is unspecified or 0, the service translates the value to 500,000, which corresponds to the `Normal` priority in the console. - If the supplied value is negative, the intent is ignored in runtime detect intent requests.
        """
        return pulumi.get(self, "priority")

    @property
    @pulumi.getter(name="trainingPhrases")
    def training_phrases(self) -> pulumi.Output[Sequence['outputs.GoogleCloudDialogflowCxV3IntentTrainingPhraseResponse']]:
        """
        The collection of training phrases the agent is trained on to identify the intent.
        """
        return pulumi.get(self, "training_phrases")

