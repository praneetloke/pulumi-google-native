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

__all__ = [
    'GetFeatureResult',
    'AwaitableGetFeatureResult',
    'get_feature',
    'get_feature_output',
]

@pulumi.output_type
class GetFeatureResult:
    def __init__(__self__, create_time=None, delete_time=None, fleet_default_member_config=None, labels=None, membership_specs=None, membership_states=None, name=None, resource_state=None, scope_specs=None, scope_states=None, spec=None, state=None, update_time=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if delete_time and not isinstance(delete_time, str):
            raise TypeError("Expected argument 'delete_time' to be a str")
        pulumi.set(__self__, "delete_time", delete_time)
        if fleet_default_member_config and not isinstance(fleet_default_member_config, dict):
            raise TypeError("Expected argument 'fleet_default_member_config' to be a dict")
        pulumi.set(__self__, "fleet_default_member_config", fleet_default_member_config)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if membership_specs and not isinstance(membership_specs, dict):
            raise TypeError("Expected argument 'membership_specs' to be a dict")
        pulumi.set(__self__, "membership_specs", membership_specs)
        if membership_states and not isinstance(membership_states, dict):
            raise TypeError("Expected argument 'membership_states' to be a dict")
        pulumi.set(__self__, "membership_states", membership_states)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if resource_state and not isinstance(resource_state, dict):
            raise TypeError("Expected argument 'resource_state' to be a dict")
        pulumi.set(__self__, "resource_state", resource_state)
        if scope_specs and not isinstance(scope_specs, dict):
            raise TypeError("Expected argument 'scope_specs' to be a dict")
        pulumi.set(__self__, "scope_specs", scope_specs)
        if scope_states and not isinstance(scope_states, dict):
            raise TypeError("Expected argument 'scope_states' to be a dict")
        pulumi.set(__self__, "scope_states", scope_states)
        if spec and not isinstance(spec, dict):
            raise TypeError("Expected argument 'spec' to be a dict")
        pulumi.set(__self__, "spec", spec)
        if state and not isinstance(state, dict):
            raise TypeError("Expected argument 'state' to be a dict")
        pulumi.set(__self__, "state", state)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        When the Feature resource was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="deleteTime")
    def delete_time(self) -> str:
        """
        When the Feature resource was deleted.
        """
        return pulumi.get(self, "delete_time")

    @property
    @pulumi.getter(name="fleetDefaultMemberConfig")
    def fleet_default_member_config(self) -> 'outputs.CommonFleetDefaultMemberConfigSpecResponse':
        """
        Optional. Feature configuration applicable to all memberships of the fleet.
        """
        return pulumi.get(self, "fleet_default_member_config")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Labels for this Feature.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter(name="membershipSpecs")
    def membership_specs(self) -> Mapping[str, str]:
        """
        Optional. Membership-specific configuration for this Feature. If this Feature does not support any per-Membership configuration, this field may be unused. The keys indicate which Membership the configuration is for, in the form: `projects/{p}/locations/{l}/memberships/{m}` Where {p} is the project, {l} is a valid location and {m} is a valid Membership in this project at that location. {p} WILL match the Feature's project. {p} will always be returned as the project number, but the project ID is also accepted during input. If the same Membership is specified in the map twice (using the project ID form, and the project number form), exactly ONE of the entries will be saved, with no guarantees as to which. For this reason, it is recommended the same format be used for all entries when mutating a Feature.
        """
        return pulumi.get(self, "membership_specs")

    @property
    @pulumi.getter(name="membershipStates")
    def membership_states(self) -> Mapping[str, str]:
        """
        Membership-specific Feature status. If this Feature does report any per-Membership status, this field may be unused. The keys indicate which Membership the state is for, in the form: `projects/{p}/locations/{l}/memberships/{m}` Where {p} is the project number, {l} is a valid location and {m} is a valid Membership in this project at that location. {p} MUST match the Feature's project number.
        """
        return pulumi.get(self, "membership_states")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The full, unique name of this Feature resource in the format `projects/*/locations/*/features/*`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resourceState")
    def resource_state(self) -> 'outputs.FeatureResourceStateResponse':
        """
        State of the Feature resource itself.
        """
        return pulumi.get(self, "resource_state")

    @property
    @pulumi.getter(name="scopeSpecs")
    def scope_specs(self) -> Mapping[str, str]:
        """
        Optional. Scope-specific configuration for this Feature. If this Feature does not support any per-Scope configuration, this field may be unused. The keys indicate which Scope the configuration is for, in the form: `projects/{p}/locations/global/scopes/{s}` Where {p} is the project, {s} is a valid Scope in this project. {p} WILL match the Feature's project. {p} will always be returned as the project number, but the project ID is also accepted during input. If the same Scope is specified in the map twice (using the project ID form, and the project number form), exactly ONE of the entries will be saved, with no guarantees as to which. For this reason, it is recommended the same format be used for all entries when mutating a Feature.
        """
        return pulumi.get(self, "scope_specs")

    @property
    @pulumi.getter(name="scopeStates")
    def scope_states(self) -> Mapping[str, str]:
        """
        Scope-specific Feature status. If this Feature does report any per-Scope status, this field may be unused. The keys indicate which Scope the state is for, in the form: `projects/{p}/locations/global/scopes/{s}` Where {p} is the project, {s} is a valid Scope in this project. {p} WILL match the Feature's project.
        """
        return pulumi.get(self, "scope_states")

    @property
    @pulumi.getter
    def spec(self) -> 'outputs.CommonFeatureSpecResponse':
        """
        Optional. Hub-wide Feature configuration. If this Feature does not support any Hub-wide configuration, this field may be unused.
        """
        return pulumi.get(self, "spec")

    @property
    @pulumi.getter
    def state(self) -> 'outputs.CommonFeatureStateResponse':
        """
        The Hub-wide Feature state.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        When the Feature resource was last updated.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetFeatureResult(GetFeatureResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFeatureResult(
            create_time=self.create_time,
            delete_time=self.delete_time,
            fleet_default_member_config=self.fleet_default_member_config,
            labels=self.labels,
            membership_specs=self.membership_specs,
            membership_states=self.membership_states,
            name=self.name,
            resource_state=self.resource_state,
            scope_specs=self.scope_specs,
            scope_states=self.scope_states,
            spec=self.spec,
            state=self.state,
            update_time=self.update_time)


def get_feature(feature_id: Optional[str] = None,
                location: Optional[str] = None,
                project: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFeatureResult:
    """
    Gets details of a single Feature.
    """
    __args__ = dict()
    __args__['featureId'] = feature_id
    __args__['location'] = location
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:gkehub/v1:getFeature', __args__, opts=opts, typ=GetFeatureResult).value

    return AwaitableGetFeatureResult(
        create_time=pulumi.get(__ret__, 'create_time'),
        delete_time=pulumi.get(__ret__, 'delete_time'),
        fleet_default_member_config=pulumi.get(__ret__, 'fleet_default_member_config'),
        labels=pulumi.get(__ret__, 'labels'),
        membership_specs=pulumi.get(__ret__, 'membership_specs'),
        membership_states=pulumi.get(__ret__, 'membership_states'),
        name=pulumi.get(__ret__, 'name'),
        resource_state=pulumi.get(__ret__, 'resource_state'),
        scope_specs=pulumi.get(__ret__, 'scope_specs'),
        scope_states=pulumi.get(__ret__, 'scope_states'),
        spec=pulumi.get(__ret__, 'spec'),
        state=pulumi.get(__ret__, 'state'),
        update_time=pulumi.get(__ret__, 'update_time'))


@_utilities.lift_output_func(get_feature)
def get_feature_output(feature_id: Optional[pulumi.Input[str]] = None,
                       location: Optional[pulumi.Input[str]] = None,
                       project: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFeatureResult]:
    """
    Gets details of a single Feature.
    """
    ...
