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
    'GetOrganizationReplayResult',
    'AwaitableGetOrganizationReplayResult',
    'get_organization_replay',
    'get_organization_replay_output',
]

@pulumi.output_type
class GetOrganizationReplayResult:
    def __init__(__self__, config=None, name=None, results_summary=None, state=None):
        if config and not isinstance(config, dict):
            raise TypeError("Expected argument 'config' to be a dict")
        pulumi.set(__self__, "config", config)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if results_summary and not isinstance(results_summary, dict):
            raise TypeError("Expected argument 'results_summary' to be a dict")
        pulumi.set(__self__, "results_summary", results_summary)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter
    def config(self) -> 'outputs.GoogleCloudPolicysimulatorV1ReplayConfigResponse':
        """
        The configuration used for the `Replay`.
        """
        return pulumi.get(self, "config")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name of the `Replay`, which has the following format: `{projects|folders|organizations}/{resource-id}/locations/global/replays/{replay-id}`, where `{resource-id}` is the ID of the project, folder, or organization that owns the Replay. Example: `projects/my-example-project/locations/global/replays/506a5f7f-38ce-4d7d-8e03-479ce1833c36`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="resultsSummary")
    def results_summary(self) -> 'outputs.GoogleCloudPolicysimulatorV1ReplayResultsSummaryResponse':
        """
        Summary statistics about the replayed log entries.
        """
        return pulumi.get(self, "results_summary")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The current state of the `Replay`.
        """
        return pulumi.get(self, "state")


class AwaitableGetOrganizationReplayResult(GetOrganizationReplayResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOrganizationReplayResult(
            config=self.config,
            name=self.name,
            results_summary=self.results_summary,
            state=self.state)


def get_organization_replay(location: Optional[str] = None,
                            organization_id: Optional[str] = None,
                            replay_id: Optional[str] = None,
                            opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOrganizationReplayResult:
    """
    Gets the specified Replay. Each `Replay` is available for at least 7 days.
    """
    __args__ = dict()
    __args__['location'] = location
    __args__['organizationId'] = organization_id
    __args__['replayId'] = replay_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:policysimulator/v1:getOrganizationReplay', __args__, opts=opts, typ=GetOrganizationReplayResult).value

    return AwaitableGetOrganizationReplayResult(
        config=pulumi.get(__ret__, 'config'),
        name=pulumi.get(__ret__, 'name'),
        results_summary=pulumi.get(__ret__, 'results_summary'),
        state=pulumi.get(__ret__, 'state'))


@_utilities.lift_output_func(get_organization_replay)
def get_organization_replay_output(location: Optional[pulumi.Input[str]] = None,
                                   organization_id: Optional[pulumi.Input[str]] = None,
                                   replay_id: Optional[pulumi.Input[str]] = None,
                                   opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetOrganizationReplayResult]:
    """
    Gets the specified Replay. Each `Replay` is available for at least 7 days.
    """
    ...
