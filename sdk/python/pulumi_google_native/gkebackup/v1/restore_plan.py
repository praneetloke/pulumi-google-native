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

__all__ = ['RestorePlanArgs', 'RestorePlan']

@pulumi.input_type
class RestorePlanArgs:
    def __init__(__self__, *,
                 backup_plan: pulumi.Input[str],
                 cluster: pulumi.Input[str],
                 restore_config: pulumi.Input['RestoreConfigArgs'],
                 restore_plan_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a RestorePlan resource.
        :param pulumi.Input[str] backup_plan: Immutable. A reference to the BackupPlan from which Backups may be used as the source for Restores created via this RestorePlan. Format: `projects/*/locations/*/backupPlans/*`.
        :param pulumi.Input[str] cluster: Immutable. The target cluster into which Restores created via this RestorePlan will restore data. NOTE: the cluster's region must be the same as the RestorePlan. Valid formats: - `projects/*/locations/*/clusters/*` - `projects/*/zones/*/clusters/*`
        :param pulumi.Input['RestoreConfigArgs'] restore_config: Configuration of Restores created via this RestorePlan.
        :param pulumi.Input[str] restore_plan_id: Required. The client-provided short name for the RestorePlan resource. This name must: - be between 1 and 63 characters long (inclusive) - consist of only lower-case ASCII letters, numbers, and dashes - start with a lower-case letter - end with a lower-case letter or number - be unique within the set of RestorePlans in this location
        :param pulumi.Input[str] description: User specified descriptive string for this RestorePlan.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of custom labels supplied by user.
        """
        pulumi.set(__self__, "backup_plan", backup_plan)
        pulumi.set(__self__, "cluster", cluster)
        pulumi.set(__self__, "restore_config", restore_config)
        pulumi.set(__self__, "restore_plan_id", restore_plan_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if labels is not None:
            pulumi.set(__self__, "labels", labels)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if project is not None:
            pulumi.set(__self__, "project", project)

    @property
    @pulumi.getter(name="backupPlan")
    def backup_plan(self) -> pulumi.Input[str]:
        """
        Immutable. A reference to the BackupPlan from which Backups may be used as the source for Restores created via this RestorePlan. Format: `projects/*/locations/*/backupPlans/*`.
        """
        return pulumi.get(self, "backup_plan")

    @backup_plan.setter
    def backup_plan(self, value: pulumi.Input[str]):
        pulumi.set(self, "backup_plan", value)

    @property
    @pulumi.getter
    def cluster(self) -> pulumi.Input[str]:
        """
        Immutable. The target cluster into which Restores created via this RestorePlan will restore data. NOTE: the cluster's region must be the same as the RestorePlan. Valid formats: - `projects/*/locations/*/clusters/*` - `projects/*/zones/*/clusters/*`
        """
        return pulumi.get(self, "cluster")

    @cluster.setter
    def cluster(self, value: pulumi.Input[str]):
        pulumi.set(self, "cluster", value)

    @property
    @pulumi.getter(name="restoreConfig")
    def restore_config(self) -> pulumi.Input['RestoreConfigArgs']:
        """
        Configuration of Restores created via this RestorePlan.
        """
        return pulumi.get(self, "restore_config")

    @restore_config.setter
    def restore_config(self, value: pulumi.Input['RestoreConfigArgs']):
        pulumi.set(self, "restore_config", value)

    @property
    @pulumi.getter(name="restorePlanId")
    def restore_plan_id(self) -> pulumi.Input[str]:
        """
        Required. The client-provided short name for the RestorePlan resource. This name must: - be between 1 and 63 characters long (inclusive) - consist of only lower-case ASCII letters, numbers, and dashes - start with a lower-case letter - end with a lower-case letter or number - be unique within the set of RestorePlans in this location
        """
        return pulumi.get(self, "restore_plan_id")

    @restore_plan_id.setter
    def restore_plan_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "restore_plan_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        User specified descriptive string for this RestorePlan.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def labels(self) -> Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]]:
        """
        A set of custom labels supplied by user.
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


class RestorePlan(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_plan: Optional[pulumi.Input[str]] = None,
                 cluster: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 restore_config: Optional[pulumi.Input[pulumi.InputType['RestoreConfigArgs']]] = None,
                 restore_plan_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a new RestorePlan in a given location.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] backup_plan: Immutable. A reference to the BackupPlan from which Backups may be used as the source for Restores created via this RestorePlan. Format: `projects/*/locations/*/backupPlans/*`.
        :param pulumi.Input[str] cluster: Immutable. The target cluster into which Restores created via this RestorePlan will restore data. NOTE: the cluster's region must be the same as the RestorePlan. Valid formats: - `projects/*/locations/*/clusters/*` - `projects/*/zones/*/clusters/*`
        :param pulumi.Input[str] description: User specified descriptive string for this RestorePlan.
        :param pulumi.Input[Mapping[str, pulumi.Input[str]]] labels: A set of custom labels supplied by user.
        :param pulumi.Input[pulumi.InputType['RestoreConfigArgs']] restore_config: Configuration of Restores created via this RestorePlan.
        :param pulumi.Input[str] restore_plan_id: Required. The client-provided short name for the RestorePlan resource. This name must: - be between 1 and 63 characters long (inclusive) - consist of only lower-case ASCII letters, numbers, and dashes - start with a lower-case letter - end with a lower-case letter or number - be unique within the set of RestorePlans in this location
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: RestorePlanArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new RestorePlan in a given location.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param RestorePlanArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(RestorePlanArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 backup_plan: Optional[pulumi.Input[str]] = None,
                 cluster: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 labels: Optional[pulumi.Input[Mapping[str, pulumi.Input[str]]]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 restore_config: Optional[pulumi.Input[pulumi.InputType['RestoreConfigArgs']]] = None,
                 restore_plan_id: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = RestorePlanArgs.__new__(RestorePlanArgs)

            if backup_plan is None and not opts.urn:
                raise TypeError("Missing required property 'backup_plan'")
            __props__.__dict__["backup_plan"] = backup_plan
            if cluster is None and not opts.urn:
                raise TypeError("Missing required property 'cluster'")
            __props__.__dict__["cluster"] = cluster
            __props__.__dict__["description"] = description
            __props__.__dict__["labels"] = labels
            __props__.__dict__["location"] = location
            __props__.__dict__["project"] = project
            if restore_config is None and not opts.urn:
                raise TypeError("Missing required property 'restore_config'")
            __props__.__dict__["restore_config"] = restore_config
            if restore_plan_id is None and not opts.urn:
                raise TypeError("Missing required property 'restore_plan_id'")
            __props__.__dict__["restore_plan_id"] = restore_plan_id
            __props__.__dict__["create_time"] = None
            __props__.__dict__["etag"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["uid"] = None
            __props__.__dict__["update_time"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["location", "project", "restore_plan_id"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(RestorePlan, __self__).__init__(
            'google-native:gkebackup/v1:RestorePlan',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'RestorePlan':
        """
        Get an existing RestorePlan resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = RestorePlanArgs.__new__(RestorePlanArgs)

        __props__.__dict__["backup_plan"] = None
        __props__.__dict__["cluster"] = None
        __props__.__dict__["create_time"] = None
        __props__.__dict__["description"] = None
        __props__.__dict__["etag"] = None
        __props__.__dict__["labels"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["restore_config"] = None
        __props__.__dict__["restore_plan_id"] = None
        __props__.__dict__["uid"] = None
        __props__.__dict__["update_time"] = None
        return RestorePlan(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="backupPlan")
    def backup_plan(self) -> pulumi.Output[str]:
        """
        Immutable. A reference to the BackupPlan from which Backups may be used as the source for Restores created via this RestorePlan. Format: `projects/*/locations/*/backupPlans/*`.
        """
        return pulumi.get(self, "backup_plan")

    @property
    @pulumi.getter
    def cluster(self) -> pulumi.Output[str]:
        """
        Immutable. The target cluster into which Restores created via this RestorePlan will restore data. NOTE: the cluster's region must be the same as the RestorePlan. Valid formats: - `projects/*/locations/*/clusters/*` - `projects/*/zones/*/clusters/*`
        """
        return pulumi.get(self, "cluster")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        The timestamp when this RestorePlan resource was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        User specified descriptive string for this RestorePlan.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        `etag` is used for optimistic concurrency control as a way to help prevent simultaneous updates of a restore from overwriting each other. It is strongly suggested that systems make use of the `etag` in the read-modify-write cycle to perform restore updates in order to avoid race conditions: An `etag` is returned in the response to `GetRestorePlan`, and systems are expected to put that etag in the request to `UpdateRestorePlan` or `DeleteRestorePlan` to ensure that their change will be applied to the same version of the resource.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def labels(self) -> pulumi.Output[Mapping[str, str]]:
        """
        A set of custom labels supplied by user.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The full name of the RestorePlan resource. Format: `projects/*/locations/*/restorePlans/*`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="restoreConfig")
    def restore_config(self) -> pulumi.Output['outputs.RestoreConfigResponse']:
        """
        Configuration of Restores created via this RestorePlan.
        """
        return pulumi.get(self, "restore_config")

    @property
    @pulumi.getter(name="restorePlanId")
    def restore_plan_id(self) -> pulumi.Output[str]:
        """
        Required. The client-provided short name for the RestorePlan resource. This name must: - be between 1 and 63 characters long (inclusive) - consist of only lower-case ASCII letters, numbers, and dashes - start with a lower-case letter - end with a lower-case letter or number - be unique within the set of RestorePlans in this location
        """
        return pulumi.get(self, "restore_plan_id")

    @property
    @pulumi.getter
    def uid(self) -> pulumi.Output[str]:
        """
        Server generated global unique identifier of [UUID](https://en.wikipedia.org/wiki/Universally_unique_identifier) format.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        The timestamp when this RestorePlan resource was last updated.
        """
        return pulumi.get(self, "update_time")

