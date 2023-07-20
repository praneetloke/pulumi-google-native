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
    'GetRegionInstanceGroupManagerResult',
    'AwaitableGetRegionInstanceGroupManagerResult',
    'get_region_instance_group_manager',
    'get_region_instance_group_manager_output',
]

@pulumi.output_type
class GetRegionInstanceGroupManagerResult:
    def __init__(__self__, auto_healing_policies=None, base_instance_name=None, creation_timestamp=None, current_actions=None, description=None, distribution_policy=None, fingerprint=None, instance_group=None, instance_template=None, kind=None, list_managed_instances_results=None, name=None, named_ports=None, region=None, self_link=None, stateful_policy=None, status=None, target_pools=None, target_size=None, update_policy=None, versions=None, zone=None):
        if auto_healing_policies and not isinstance(auto_healing_policies, list):
            raise TypeError("Expected argument 'auto_healing_policies' to be a list")
        pulumi.set(__self__, "auto_healing_policies", auto_healing_policies)
        if base_instance_name and not isinstance(base_instance_name, str):
            raise TypeError("Expected argument 'base_instance_name' to be a str")
        pulumi.set(__self__, "base_instance_name", base_instance_name)
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError("Expected argument 'creation_timestamp' to be a str")
        pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if current_actions and not isinstance(current_actions, dict):
            raise TypeError("Expected argument 'current_actions' to be a dict")
        pulumi.set(__self__, "current_actions", current_actions)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if distribution_policy and not isinstance(distribution_policy, dict):
            raise TypeError("Expected argument 'distribution_policy' to be a dict")
        pulumi.set(__self__, "distribution_policy", distribution_policy)
        if fingerprint and not isinstance(fingerprint, str):
            raise TypeError("Expected argument 'fingerprint' to be a str")
        pulumi.set(__self__, "fingerprint", fingerprint)
        if instance_group and not isinstance(instance_group, str):
            raise TypeError("Expected argument 'instance_group' to be a str")
        pulumi.set(__self__, "instance_group", instance_group)
        if instance_template and not isinstance(instance_template, str):
            raise TypeError("Expected argument 'instance_template' to be a str")
        pulumi.set(__self__, "instance_template", instance_template)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if list_managed_instances_results and not isinstance(list_managed_instances_results, str):
            raise TypeError("Expected argument 'list_managed_instances_results' to be a str")
        pulumi.set(__self__, "list_managed_instances_results", list_managed_instances_results)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if named_ports and not isinstance(named_ports, list):
            raise TypeError("Expected argument 'named_ports' to be a list")
        pulumi.set(__self__, "named_ports", named_ports)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if stateful_policy and not isinstance(stateful_policy, dict):
            raise TypeError("Expected argument 'stateful_policy' to be a dict")
        pulumi.set(__self__, "stateful_policy", stateful_policy)
        if status and not isinstance(status, dict):
            raise TypeError("Expected argument 'status' to be a dict")
        pulumi.set(__self__, "status", status)
        if target_pools and not isinstance(target_pools, list):
            raise TypeError("Expected argument 'target_pools' to be a list")
        pulumi.set(__self__, "target_pools", target_pools)
        if target_size and not isinstance(target_size, int):
            raise TypeError("Expected argument 'target_size' to be a int")
        pulumi.set(__self__, "target_size", target_size)
        if update_policy and not isinstance(update_policy, dict):
            raise TypeError("Expected argument 'update_policy' to be a dict")
        pulumi.set(__self__, "update_policy", update_policy)
        if versions and not isinstance(versions, list):
            raise TypeError("Expected argument 'versions' to be a list")
        pulumi.set(__self__, "versions", versions)
        if zone and not isinstance(zone, str):
            raise TypeError("Expected argument 'zone' to be a str")
        pulumi.set(__self__, "zone", zone)

    @property
    @pulumi.getter(name="autoHealingPolicies")
    def auto_healing_policies(self) -> Sequence['outputs.InstanceGroupManagerAutoHealingPolicyResponse']:
        """
        The autohealing policy for this managed instance group. You can specify only one value.
        """
        return pulumi.get(self, "auto_healing_policies")

    @property
    @pulumi.getter(name="baseInstanceName")
    def base_instance_name(self) -> str:
        """
        The base instance name to use for instances in this group. The value must be 1-58 characters long. Instances are named by appending a hyphen and a random four-character string to the base instance name. The base instance name must comply with RFC1035.
        """
        return pulumi.get(self, "base_instance_name")

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> str:
        """
        The creation timestamp for this managed instance group in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter(name="currentActions")
    def current_actions(self) -> 'outputs.InstanceGroupManagerActionsSummaryResponse':
        """
        The list of instance actions and the number of instances in this managed instance group that are scheduled for each of those actions.
        """
        return pulumi.get(self, "current_actions")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        An optional description of this resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="distributionPolicy")
    def distribution_policy(self) -> 'outputs.DistributionPolicyResponse':
        """
        Policy specifying the intended distribution of managed instances across zones in a regional managed instance group.
        """
        return pulumi.get(self, "distribution_policy")

    @property
    @pulumi.getter
    def fingerprint(self) -> str:
        """
        Fingerprint of this resource. This field may be used in optimistic locking. It will be ignored when inserting an InstanceGroupManager. An up-to-date fingerprint must be provided in order to update the InstanceGroupManager, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve an InstanceGroupManager.
        """
        return pulumi.get(self, "fingerprint")

    @property
    @pulumi.getter(name="instanceGroup")
    def instance_group(self) -> str:
        """
        The URL of the Instance Group resource.
        """
        return pulumi.get(self, "instance_group")

    @property
    @pulumi.getter(name="instanceTemplate")
    def instance_template(self) -> str:
        """
        The URL of the instance template that is specified for this managed instance group. The group uses this template to create all new instances in the managed instance group. The templates for existing instances in the group do not change unless you run recreateInstances, run applyUpdatesToInstances, or set the group's updatePolicy.type to PROACTIVE.
        """
        return pulumi.get(self, "instance_template")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        The resource type, which is always compute#instanceGroupManager for managed instance groups.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="listManagedInstancesResults")
    def list_managed_instances_results(self) -> str:
        """
        Pagination behavior of the listManagedInstances API method for this managed instance group.
        """
        return pulumi.get(self, "list_managed_instances_results")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the managed instance group. The name must be 1-63 characters long, and comply with RFC1035.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="namedPorts")
    def named_ports(self) -> Sequence['outputs.NamedPortResponse']:
        """
        Named ports configured for the Instance Groups complementary to this Instance Group Manager.
        """
        return pulumi.get(self, "named_ports")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        The URL of the region where the managed instance group resides (for regional resources).
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        The URL for this managed instance group. The server defines this URL.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter(name="statefulPolicy")
    def stateful_policy(self) -> 'outputs.StatefulPolicyResponse':
        """
        Stateful configuration for this Instanced Group Manager
        """
        return pulumi.get(self, "stateful_policy")

    @property
    @pulumi.getter
    def status(self) -> 'outputs.InstanceGroupManagerStatusResponse':
        """
        The status of this managed instance group.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="targetPools")
    def target_pools(self) -> Sequence[str]:
        """
        The URLs for all TargetPool resources to which instances in the instanceGroup field are added. The target pools automatically apply to all of the instances in the managed instance group.
        """
        return pulumi.get(self, "target_pools")

    @property
    @pulumi.getter(name="targetSize")
    def target_size(self) -> int:
        """
        The target number of running instances for this managed instance group. You can reduce this number by using the instanceGroupManager deleteInstances or abandonInstances methods. Resizing the group also changes this number.
        """
        return pulumi.get(self, "target_size")

    @property
    @pulumi.getter(name="updatePolicy")
    def update_policy(self) -> 'outputs.InstanceGroupManagerUpdatePolicyResponse':
        """
        The update policy for this managed instance group.
        """
        return pulumi.get(self, "update_policy")

    @property
    @pulumi.getter
    def versions(self) -> Sequence['outputs.InstanceGroupManagerVersionResponse']:
        """
        Specifies the instance templates used by this managed instance group to create instances. Each version is defined by an instanceTemplate and a name. Every version can appear at most once per instance group. This field overrides the top-level instanceTemplate field. Read more about the relationships between these fields. Exactly one version must leave the targetSize field unset. That version will be applied to all remaining instances. For more information, read about canary updates.
        """
        return pulumi.get(self, "versions")

    @property
    @pulumi.getter
    def zone(self) -> str:
        """
        The URL of a zone where the managed instance group is located (for zonal resources).
        """
        return pulumi.get(self, "zone")


class AwaitableGetRegionInstanceGroupManagerResult(GetRegionInstanceGroupManagerResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetRegionInstanceGroupManagerResult(
            auto_healing_policies=self.auto_healing_policies,
            base_instance_name=self.base_instance_name,
            creation_timestamp=self.creation_timestamp,
            current_actions=self.current_actions,
            description=self.description,
            distribution_policy=self.distribution_policy,
            fingerprint=self.fingerprint,
            instance_group=self.instance_group,
            instance_template=self.instance_template,
            kind=self.kind,
            list_managed_instances_results=self.list_managed_instances_results,
            name=self.name,
            named_ports=self.named_ports,
            region=self.region,
            self_link=self.self_link,
            stateful_policy=self.stateful_policy,
            status=self.status,
            target_pools=self.target_pools,
            target_size=self.target_size,
            update_policy=self.update_policy,
            versions=self.versions,
            zone=self.zone)


def get_region_instance_group_manager(instance_group_manager: Optional[str] = None,
                                      project: Optional[str] = None,
                                      region: Optional[str] = None,
                                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetRegionInstanceGroupManagerResult:
    """
    Returns all of the details about the specified managed instance group.
    """
    __args__ = dict()
    __args__['instanceGroupManager'] = instance_group_manager
    __args__['project'] = project
    __args__['region'] = region
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:compute/v1:getRegionInstanceGroupManager', __args__, opts=opts, typ=GetRegionInstanceGroupManagerResult).value

    return AwaitableGetRegionInstanceGroupManagerResult(
        auto_healing_policies=pulumi.get(__ret__, 'auto_healing_policies'),
        base_instance_name=pulumi.get(__ret__, 'base_instance_name'),
        creation_timestamp=pulumi.get(__ret__, 'creation_timestamp'),
        current_actions=pulumi.get(__ret__, 'current_actions'),
        description=pulumi.get(__ret__, 'description'),
        distribution_policy=pulumi.get(__ret__, 'distribution_policy'),
        fingerprint=pulumi.get(__ret__, 'fingerprint'),
        instance_group=pulumi.get(__ret__, 'instance_group'),
        instance_template=pulumi.get(__ret__, 'instance_template'),
        kind=pulumi.get(__ret__, 'kind'),
        list_managed_instances_results=pulumi.get(__ret__, 'list_managed_instances_results'),
        name=pulumi.get(__ret__, 'name'),
        named_ports=pulumi.get(__ret__, 'named_ports'),
        region=pulumi.get(__ret__, 'region'),
        self_link=pulumi.get(__ret__, 'self_link'),
        stateful_policy=pulumi.get(__ret__, 'stateful_policy'),
        status=pulumi.get(__ret__, 'status'),
        target_pools=pulumi.get(__ret__, 'target_pools'),
        target_size=pulumi.get(__ret__, 'target_size'),
        update_policy=pulumi.get(__ret__, 'update_policy'),
        versions=pulumi.get(__ret__, 'versions'),
        zone=pulumi.get(__ret__, 'zone'))


@_utilities.lift_output_func(get_region_instance_group_manager)
def get_region_instance_group_manager_output(instance_group_manager: Optional[pulumi.Input[str]] = None,
                                             project: Optional[pulumi.Input[Optional[str]]] = None,
                                             region: Optional[pulumi.Input[str]] = None,
                                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetRegionInstanceGroupManagerResult]:
    """
    Returns all of the details about the specified managed instance group.
    """
    ...
