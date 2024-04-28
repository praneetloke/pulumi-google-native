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
from ._enums import *
from ._inputs import *

__all__ = ['FeatureArgs', 'Feature']

@pulumi.input_type
class FeatureArgs:
    def __init__(__self__, *,
                 feature_id: Optional[pulumi.Input[str]] = None,
                 fleet_default_member_config: Optional[pulumi.Input['CommonFleetDefaultMemberConfigSpecArgs']] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 membership_specs: Optional[pulumi.Input['MembershipFeatureSpecArgs']] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 scope_specs: Optional[pulumi.Input['ScopeFeatureSpecArgs']] = None,
                 spec: Optional[pulumi.Input['CommonFeatureSpecArgs']] = None):
        """
        The set of arguments for constructing a Feature resource.
        :param pulumi.Input[str] feature_id: The ID of the feature to create.
        :param pulumi.Input['CommonFleetDefaultMemberConfigSpecArgs'] fleet_default_member_config: Optional. Feature configuration applicable to all memberships of the fleet.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels for this Feature.
        :param pulumi.Input['MembershipFeatureSpecArgs'] membership_specs: Optional. Membership-specific configuration for this Feature. If this Feature does not support any per-Membership configuration, this field may be unused. The keys indicate which Membership the configuration is for, in the form: `projects/{p}/locations/{l}/memberships/{m}` Where {p} is the project, {l} is a valid location and {m} is a valid Membership in this project at that location. {p} WILL match the Feature's project. {p} will always be returned as the project number, but the project ID is also accepted during input. If the same Membership is specified in the map twice (using the project ID form, and the project number form), exactly ONE of the entries will be saved, with no guarantees as to which. For this reason, it is recommended the same format be used for all entries when mutating a Feature.
        :param pulumi.Input[str] request_id: A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes after the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        :param pulumi.Input['ScopeFeatureSpecArgs'] scope_specs: Optional. Scope-specific configuration for this Feature. If this Feature does not support any per-Scope configuration, this field may be unused. The keys indicate which Scope the configuration is for, in the form: `projects/{p}/locations/global/scopes/{s}` Where {p} is the project, {s} is a valid Scope in this project. {p} WILL match the Feature's project. {p} will always be returned as the project number, but the project ID is also accepted during input. If the same Scope is specified in the map twice (using the project ID form, and the project number form), exactly ONE of the entries will be saved, with no guarantees as to which. For this reason, it is recommended the same format be used for all entries when mutating a Feature.
        :param pulumi.Input['CommonFeatureSpecArgs'] spec: Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
        """
        if feature_id is not None:
            pulumi.set(__self__, "feature_id", feature_id)
        if fleet_default_member_config is not None:
            pulumi.set(__self__, "fleet_default_member_config", fleet_default_member_config)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if membership_specs is not None:
            pulumi.set(__self__, "membership_specs", membership_specs)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if request_id is not None:
            pulumi.set(__self__, "request_id", request_id)
        if scope_specs is not None:
            pulumi.set(__self__, "scope_specs", scope_specs)
        if spec is not None:
            pulumi.set(__self__, "spec", spec)

    @property
    @pulumi.getter(name="featureId")
    def feature_id(self) -> Optional[pulumi.Input[str]]:
        """
        The ID of the feature to create.
        """
        return pulumi.get(self, "feature_id")

    @feature_id.setter
    def feature_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "feature_id", value)

    @property
    @pulumi.getter(name="fleetDefaultMemberConfig")
    def fleet_default_member_config(self) -> Optional[pulumi.Input['CommonFleetDefaultMemberConfigSpecArgs']]:
        """
        Optional. Feature configuration applicable to all memberships of the fleet.
        """
        return pulumi.get(self, "fleet_default_member_config")

    @fleet_default_member_config.setter
    def fleet_default_member_config(self, value: Optional[pulumi.Input['CommonFleetDefaultMemberConfigSpecArgs']]):
        pulumi.set(self, "fleet_default_member_config", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        Labels for this Feature.
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
    @pulumi.getter(name="membershipSpecs")
    def membership_specs(self) -> Optional[pulumi.Input['MembershipFeatureSpecArgs']]:
        """
        Optional. Membership-specific configuration for this Feature. If this Feature does not support any per-Membership configuration, this field may be unused. The keys indicate which Membership the configuration is for, in the form: `projects/{p}/locations/{l}/memberships/{m}` Where {p} is the project, {l} is a valid location and {m} is a valid Membership in this project at that location. {p} WILL match the Feature's project. {p} will always be returned as the project number, but the project ID is also accepted during input. If the same Membership is specified in the map twice (using the project ID form, and the project number form), exactly ONE of the entries will be saved, with no guarantees as to which. For this reason, it is recommended the same format be used for all entries when mutating a Feature.
        """
        return pulumi.get(self, "membership_specs")

    @membership_specs.setter
    def membership_specs(self, value: Optional[pulumi.Input['MembershipFeatureSpecArgs']]):
        pulumi.set(self, "membership_specs", value)

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
        """
        A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes after the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        """
        return pulumi.get(self, "request_id")

    @request_id.setter
    def request_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "request_id", value)

    @property
    @pulumi.getter(name="scopeSpecs")
    def scope_specs(self) -> Optional[pulumi.Input['ScopeFeatureSpecArgs']]:
        """
        Optional. Scope-specific configuration for this Feature. If this Feature does not support any per-Scope configuration, this field may be unused. The keys indicate which Scope the configuration is for, in the form: `projects/{p}/locations/global/scopes/{s}` Where {p} is the project, {s} is a valid Scope in this project. {p} WILL match the Feature's project. {p} will always be returned as the project number, but the project ID is also accepted during input. If the same Scope is specified in the map twice (using the project ID form, and the project number form), exactly ONE of the entries will be saved, with no guarantees as to which. For this reason, it is recommended the same format be used for all entries when mutating a Feature.
        """
        return pulumi.get(self, "scope_specs")

    @scope_specs.setter
    def scope_specs(self, value: Optional[pulumi.Input['ScopeFeatureSpecArgs']]):
        pulumi.set(self, "scope_specs", value)

    @property
    @pulumi.getter
    def spec(self) -> Optional[pulumi.Input['CommonFeatureSpecArgs']]:
        """
        Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
        """
        return pulumi.get(self, "spec")

    @spec.setter
    def spec(self, value: Optional[pulumi.Input['CommonFeatureSpecArgs']]):
        pulumi.set(self, "spec", value)


class Feature(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 feature_id: Optional[pulumi.Input[str]] = None,
                 fleet_default_member_config: Optional[pulumi.Input[pulumi.InputType['CommonFleetDefaultMemberConfigSpecArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 membership_specs: Optional[pulumi.Input[pulumi.InputType['MembershipFeatureSpecArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 scope_specs: Optional[pulumi.Input[pulumi.InputType['ScopeFeatureSpecArgs']]] = None,
                 spec: Optional[pulumi.Input[pulumi.InputType['CommonFeatureSpecArgs']]] = None,
                 __props__=None):
        """
        Adds a new Feature.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] feature_id: The ID of the feature to create.
        :param pulumi.Input[pulumi.InputType['CommonFleetDefaultMemberConfigSpecArgs']] fleet_default_member_config: Optional. Feature configuration applicable to all memberships of the fleet.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: Labels for this Feature.
        :param pulumi.Input[pulumi.InputType['MembershipFeatureSpecArgs']] membership_specs: Optional. Membership-specific configuration for this Feature. If this Feature does not support any per-Membership configuration, this field may be unused. The keys indicate which Membership the configuration is for, in the form: `projects/{p}/locations/{l}/memberships/{m}` Where {p} is the project, {l} is a valid location and {m} is a valid Membership in this project at that location. {p} WILL match the Feature's project. {p} will always be returned as the project number, but the project ID is also accepted during input. If the same Membership is specified in the map twice (using the project ID form, and the project number form), exactly ONE of the entries will be saved, with no guarantees as to which. For this reason, it is recommended the same format be used for all entries when mutating a Feature.
        :param pulumi.Input[str] request_id: A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes after the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        :param pulumi.Input[pulumi.InputType['ScopeFeatureSpecArgs']] scope_specs: Optional. Scope-specific configuration for this Feature. If this Feature does not support any per-Scope configuration, this field may be unused. The keys indicate which Scope the configuration is for, in the form: `projects/{p}/locations/global/scopes/{s}` Where {p} is the project, {s} is a valid Scope in this project. {p} WILL match the Feature's project. {p} will always be returned as the project number, but the project ID is also accepted during input. If the same Scope is specified in the map twice (using the project ID form, and the project number form), exactly ONE of the entries will be saved, with no guarantees as to which. For this reason, it is recommended the same format be used for all entries when mutating a Feature.
        :param pulumi.Input[pulumi.InputType['CommonFeatureSpecArgs']] spec: Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[FeatureArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Adds a new Feature.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param FeatureArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FeatureArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 feature_id: Optional[pulumi.Input[str]] = None,
                 fleet_default_member_config: Optional[pulumi.Input[pulumi.InputType['CommonFleetDefaultMemberConfigSpecArgs']]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 membership_specs: Optional[pulumi.Input[pulumi.InputType['MembershipFeatureSpecArgs']]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 request_id: Optional[pulumi.Input[str]] = None,
                 scope_specs: Optional[pulumi.Input[pulumi.InputType['ScopeFeatureSpecArgs']]] = None,
                 spec: Optional[pulumi.Input[pulumi.InputType['CommonFeatureSpecArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FeatureArgs.__new__(FeatureArgs)

            __props__.__dict__["feature_id"] = feature_id
            __props__.__dict__["fleet_default_member_config"] = fleet_default_member_config
            __props__.__dict__["labels"] = labels
            __props__.__dict__["location"] = location
            __props__.__dict__["membership_specs"] = membership_specs
            __props__.__dict__["project"] = project
            __props__.__dict__["request_id"] = request_id
            __props__.__dict__["scope_specs"] = scope_specs
            __props__.__dict__["spec"] = spec
            __props__.__dict__["create_time"] = None
            __props__.__dict__["delete_time"] = None
            __props__.__dict__["membership_states"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["resource_state"] = None
            __props__.__dict__["scope_states"] = None
            __props__.__dict__["state"] = None
            __props__.__dict__["update_time"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["location", "project"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Feature, __self__).__init__(
            'google-native:gkehub/v1:Feature',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Feature':
        """
        Get an existing Feature resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = FeatureArgs.__new__(FeatureArgs)

        __props__.__dict__["create_time"] = None
        __props__.__dict__["delete_time"] = None
        __props__.__dict__["feature_id"] = None
        __props__.__dict__["fleet_default_member_config"] = None
        __props__.__dict__["labels"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["membership_specs"] = None
        __props__.__dict__["membership_states"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["request_id"] = None
        __props__.__dict__["resource_state"] = None
        __props__.__dict__["scope_specs"] = None
        __props__.__dict__["scope_states"] = None
        __props__.__dict__["spec"] = None
        __props__.__dict__["state"] = None
        __props__.__dict__["update_time"] = None
        return Feature(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        When the Feature resource was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="deleteTime")
    def delete_time(self) -> pulumi.Output[str]:
        """
        When the Feature resource was deleted.
        """
        return pulumi.get(self, "delete_time")

    @property
    @pulumi.getter(name="featureId")
    def feature_id(self) -> pulumi.Output[Optional[str]]:
        """
        The ID of the feature to create.
        """
        return pulumi.get(self, "feature_id")

    @property
    @pulumi.getter(name="fleetDefaultMemberConfig")
    def fleet_default_member_config(self) -> pulumi.Output['outputs.CommonFleetDefaultMemberConfigSpecResponse']:
        """
        Optional. Feature configuration applicable to all memberships of the fleet.
        """
        return pulumi.get(self, "fleet_default_member_config")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        Labels for this Feature.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="membershipSpecs")
    def membership_specs(self) -> pulumi.Output['outputs.MembershipFeatureSpecResponse']:
        """
        Optional. Membership-specific configuration for this Feature. If this Feature does not support any per-Membership configuration, this field may be unused. The keys indicate which Membership the configuration is for, in the form: `projects/{p}/locations/{l}/memberships/{m}` Where {p} is the project, {l} is a valid location and {m} is a valid Membership in this project at that location. {p} WILL match the Feature's project. {p} will always be returned as the project number, but the project ID is also accepted during input. If the same Membership is specified in the map twice (using the project ID form, and the project number form), exactly ONE of the entries will be saved, with no guarantees as to which. For this reason, it is recommended the same format be used for all entries when mutating a Feature.
        """
        return pulumi.get(self, "membership_specs")

    @property
    @pulumi.getter(name="membershipStates")
    def membership_states(self) -> pulumi.Output['outputs.MembershipFeatureStateResponse']:
        """
        Membership-specific Feature status. If this Feature does report any per-Membership status, this field may be unused. The keys indicate which Membership the state is for, in the form: `projects/{p}/locations/{l}/memberships/{m}` Where {p} is the project number, {l} is a valid location and {m} is a valid Membership in this project at that location. {p} MUST match the Feature's project number.
        """
        return pulumi.get(self, "membership_states")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The full, unique name of this Feature resource in the format `projects/*/locations/*/features/*`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="requestId")
    def request_id(self) -> pulumi.Output[Optional[str]]:
        """
        A request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. The server will guarantee that for at least 60 minutes after the first request. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported (00000000-0000-0000-0000-000000000000).
        """
        return pulumi.get(self, "request_id")

    @property
    @pulumi.getter(name="resourceState")
    def resource_state(self) -> pulumi.Output['outputs.FeatureResourceStateResponse']:
        """
        State of the Feature resource itself.
        """
        return pulumi.get(self, "resource_state")

    @property
    @pulumi.getter(name="scopeSpecs")
    def scope_specs(self) -> pulumi.Output['outputs.ScopeFeatureSpecResponse']:
        """
        Optional. Scope-specific configuration for this Feature. If this Feature does not support any per-Scope configuration, this field may be unused. The keys indicate which Scope the configuration is for, in the form: `projects/{p}/locations/global/scopes/{s}` Where {p} is the project, {s} is a valid Scope in this project. {p} WILL match the Feature's project. {p} will always be returned as the project number, but the project ID is also accepted during input. If the same Scope is specified in the map twice (using the project ID form, and the project number form), exactly ONE of the entries will be saved, with no guarantees as to which. For this reason, it is recommended the same format be used for all entries when mutating a Feature.
        """
        return pulumi.get(self, "scope_specs")

    @property
    @pulumi.getter(name="scopeStates")
    def scope_states(self) -> pulumi.Output['outputs.ScopeFeatureStateResponse']:
        """
        Scope-specific Feature status. If this Feature does report any per-Scope status, this field may be unused. The keys indicate which Scope the state is for, in the form: `projects/{p}/locations/global/scopes/{s}` Where {p} is the project, {s} is a valid Scope in this project. {p} WILL match the Feature's project.
        """
        return pulumi.get(self, "scope_states")

    @property
    @pulumi.getter
    def spec(self) -> pulumi.Output['outputs.CommonFeatureSpecResponse']:
        """
        Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
        """
        return pulumi.get(self, "spec")

    @property
    @pulumi.getter
    def state(self) -> pulumi.Output['outputs.CommonFeatureStateResponse']:
        """
        The Hub-wide Feature state.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        When the Feature resource was last updated.
        """
        return pulumi.get(self, "update_time")

