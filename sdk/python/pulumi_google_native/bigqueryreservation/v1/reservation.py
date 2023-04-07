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

__all__ = ['ReservationArgs', 'Reservation']

@pulumi.input_type
class ReservationArgs:
    def __init__(__self__, *,
                 autoscale: Optional[pulumi.Input['AutoscaleArgs']] = None,
                 concurrency: Optional[pulumi.Input[str]] = None,
                 edition: Optional[pulumi.Input['ReservationEdition']] = None,
                 ignore_idle_slots: Optional[pulumi.Input[bool]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 multi_region_auxiliary: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 reservation_id: Optional[pulumi.Input[str]] = None,
                 slot_capacity: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Reservation resource.
        :param pulumi.Input['AutoscaleArgs'] autoscale: The configuration parameters for the auto scaling feature. Note this is an alpha feature.
        :param pulumi.Input[str] concurrency: Job concurrency target which sets a soft upper bound on the number of jobs that can run concurrently in this reservation. This is a soft target due to asynchronous nature of the system and various optimizations for small queries. Default value is 0 which means that concurrency target will be automatically computed by the system. NOTE: this field is exposed as `target_job_concurrency` in the Information Schema, DDL and BQ CLI.
        :param pulumi.Input['ReservationEdition'] edition: Edition of the reservation.
        :param pulumi.Input[bool] ignore_idle_slots: If false, any query or pipeline job using this reservation will use idle slots from other reservations within the same admin project. If true, a query or pipeline job using this reservation will execute with the slot capacity specified in the slot_capacity field at most.
        :param pulumi.Input[bool] multi_region_auxiliary: Applicable only for reservations located within one of the BigQuery multi-regions (US or EU). If set to true, this reservation is placed in the organization's secondary region which is designated for disaster recovery purposes. If false, this reservation is placed in the organization's default region. NOTE: this is a preview feature. Project must be allow-listed in order to set this field.
        :param pulumi.Input[str] name: The resource name of the reservation, e.g., `projects/*/locations/*/reservations/team1-prod`. The reservation_id must only contain lower case alphanumeric characters or dashes. It must start with a letter and must not end with a dash. Its maximum length is 64 characters.
        :param pulumi.Input[str] reservation_id: The reservation ID. It must only contain lower case alphanumeric characters or dashes. It must start with a letter and must not end with a dash. Its maximum length is 64 characters.
        :param pulumi.Input[str] slot_capacity: Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the unit of parallelism. Queries using this reservation might use more slots during runtime if ignore_idle_slots is set to false. If total slot_capacity of the reservation and its siblings exceeds the total slot_count of all capacity commitments, the request will fail with `google.rpc.Code.RESOURCE_EXHAUSTED`. NOTE: for reservations in US or EU multi-regions, slot capacity constraints are checked separately for default and auxiliary regions. See multi_region_auxiliary flag for more details.
        """
        if autoscale is not None:
            pulumi.set(__self__, "autoscale", autoscale)
        if concurrency is not None:
            pulumi.set(__self__, "concurrency", concurrency)
        if edition is not None:
            pulumi.set(__self__, "edition", edition)
        if ignore_idle_slots is not None:
            pulumi.set(__self__, "ignore_idle_slots", ignore_idle_slots)
        if location is not None:
            pulumi.set(__self__, "location", location)
        if multi_region_auxiliary is not None:
            pulumi.set(__self__, "multi_region_auxiliary", multi_region_auxiliary)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if reservation_id is not None:
            pulumi.set(__self__, "reservation_id", reservation_id)
        if slot_capacity is not None:
            pulumi.set(__self__, "slot_capacity", slot_capacity)

    @property
    @pulumi.getter
    def autoscale(self) -> Optional[pulumi.Input['AutoscaleArgs']]:
        """
        The configuration parameters for the auto scaling feature. Note this is an alpha feature.
        """
        return pulumi.get(self, "autoscale")

    @autoscale.setter
    def autoscale(self, value: Optional[pulumi.Input['AutoscaleArgs']]):
        pulumi.set(self, "autoscale", value)

    @property
    @pulumi.getter
    def concurrency(self) -> Optional[pulumi.Input[str]]:
        """
        Job concurrency target which sets a soft upper bound on the number of jobs that can run concurrently in this reservation. This is a soft target due to asynchronous nature of the system and various optimizations for small queries. Default value is 0 which means that concurrency target will be automatically computed by the system. NOTE: this field is exposed as `target_job_concurrency` in the Information Schema, DDL and BQ CLI.
        """
        return pulumi.get(self, "concurrency")

    @concurrency.setter
    def concurrency(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "concurrency", value)

    @property
    @pulumi.getter
    def edition(self) -> Optional[pulumi.Input['ReservationEdition']]:
        """
        Edition of the reservation.
        """
        return pulumi.get(self, "edition")

    @edition.setter
    def edition(self, value: Optional[pulumi.Input['ReservationEdition']]):
        pulumi.set(self, "edition", value)

    @property
    @pulumi.getter(name="ignoreIdleSlots")
    def ignore_idle_slots(self) -> Optional[pulumi.Input[bool]]:
        """
        If false, any query or pipeline job using this reservation will use idle slots from other reservations within the same admin project. If true, a query or pipeline job using this reservation will execute with the slot capacity specified in the slot_capacity field at most.
        """
        return pulumi.get(self, "ignore_idle_slots")

    @ignore_idle_slots.setter
    def ignore_idle_slots(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "ignore_idle_slots", value)

    @property
    @pulumi.getter
    def location(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "location")

    @location.setter
    def location(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "location", value)

    @property
    @pulumi.getter(name="multiRegionAuxiliary")
    def multi_region_auxiliary(self) -> Optional[pulumi.Input[bool]]:
        """
        Applicable only for reservations located within one of the BigQuery multi-regions (US or EU). If set to true, this reservation is placed in the organization's secondary region which is designated for disaster recovery purposes. If false, this reservation is placed in the organization's default region. NOTE: this is a preview feature. Project must be allow-listed in order to set this field.
        """
        return pulumi.get(self, "multi_region_auxiliary")

    @multi_region_auxiliary.setter
    def multi_region_auxiliary(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "multi_region_auxiliary", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The resource name of the reservation, e.g., `projects/*/locations/*/reservations/team1-prod`. The reservation_id must only contain lower case alphanumeric characters or dashes. It must start with a letter and must not end with a dash. Its maximum length is 64 characters.
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
    @pulumi.getter(name="reservationId")
    def reservation_id(self) -> Optional[pulumi.Input[str]]:
        """
        The reservation ID. It must only contain lower case alphanumeric characters or dashes. It must start with a letter and must not end with a dash. Its maximum length is 64 characters.
        """
        return pulumi.get(self, "reservation_id")

    @reservation_id.setter
    def reservation_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "reservation_id", value)

    @property
    @pulumi.getter(name="slotCapacity")
    def slot_capacity(self) -> Optional[pulumi.Input[str]]:
        """
        Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the unit of parallelism. Queries using this reservation might use more slots during runtime if ignore_idle_slots is set to false. If total slot_capacity of the reservation and its siblings exceeds the total slot_count of all capacity commitments, the request will fail with `google.rpc.Code.RESOURCE_EXHAUSTED`. NOTE: for reservations in US or EU multi-regions, slot capacity constraints are checked separately for default and auxiliary regions. See multi_region_auxiliary flag for more details.
        """
        return pulumi.get(self, "slot_capacity")

    @slot_capacity.setter
    def slot_capacity(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "slot_capacity", value)


class Reservation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 autoscale: Optional[pulumi.Input[pulumi.InputType['AutoscaleArgs']]] = None,
                 concurrency: Optional[pulumi.Input[str]] = None,
                 edition: Optional[pulumi.Input['ReservationEdition']] = None,
                 ignore_idle_slots: Optional[pulumi.Input[bool]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 multi_region_auxiliary: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 reservation_id: Optional[pulumi.Input[str]] = None,
                 slot_capacity: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        Creates a new reservation resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[pulumi.InputType['AutoscaleArgs']] autoscale: The configuration parameters for the auto scaling feature. Note this is an alpha feature.
        :param pulumi.Input[str] concurrency: Job concurrency target which sets a soft upper bound on the number of jobs that can run concurrently in this reservation. This is a soft target due to asynchronous nature of the system and various optimizations for small queries. Default value is 0 which means that concurrency target will be automatically computed by the system. NOTE: this field is exposed as `target_job_concurrency` in the Information Schema, DDL and BQ CLI.
        :param pulumi.Input['ReservationEdition'] edition: Edition of the reservation.
        :param pulumi.Input[bool] ignore_idle_slots: If false, any query or pipeline job using this reservation will use idle slots from other reservations within the same admin project. If true, a query or pipeline job using this reservation will execute with the slot capacity specified in the slot_capacity field at most.
        :param pulumi.Input[bool] multi_region_auxiliary: Applicable only for reservations located within one of the BigQuery multi-regions (US or EU). If set to true, this reservation is placed in the organization's secondary region which is designated for disaster recovery purposes. If false, this reservation is placed in the organization's default region. NOTE: this is a preview feature. Project must be allow-listed in order to set this field.
        :param pulumi.Input[str] name: The resource name of the reservation, e.g., `projects/*/locations/*/reservations/team1-prod`. The reservation_id must only contain lower case alphanumeric characters or dashes. It must start with a letter and must not end with a dash. Its maximum length is 64 characters.
        :param pulumi.Input[str] reservation_id: The reservation ID. It must only contain lower case alphanumeric characters or dashes. It must start with a letter and must not end with a dash. Its maximum length is 64 characters.
        :param pulumi.Input[str] slot_capacity: Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the unit of parallelism. Queries using this reservation might use more slots during runtime if ignore_idle_slots is set to false. If total slot_capacity of the reservation and its siblings exceeds the total slot_count of all capacity commitments, the request will fail with `google.rpc.Code.RESOURCE_EXHAUSTED`. NOTE: for reservations in US or EU multi-regions, slot capacity constraints are checked separately for default and auxiliary regions. See multi_region_auxiliary flag for more details.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[ReservationArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new reservation resource.

        :param str resource_name: The name of the resource.
        :param ReservationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ReservationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 autoscale: Optional[pulumi.Input[pulumi.InputType['AutoscaleArgs']]] = None,
                 concurrency: Optional[pulumi.Input[str]] = None,
                 edition: Optional[pulumi.Input['ReservationEdition']] = None,
                 ignore_idle_slots: Optional[pulumi.Input[bool]] = None,
                 location: Optional[pulumi.Input[str]] = None,
                 multi_region_auxiliary: Optional[pulumi.Input[bool]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 reservation_id: Optional[pulumi.Input[str]] = None,
                 slot_capacity: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ReservationArgs.__new__(ReservationArgs)

            __props__.__dict__["autoscale"] = autoscale
            __props__.__dict__["concurrency"] = concurrency
            __props__.__dict__["edition"] = edition
            __props__.__dict__["ignore_idle_slots"] = ignore_idle_slots
            __props__.__dict__["location"] = location
            __props__.__dict__["multi_region_auxiliary"] = multi_region_auxiliary
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            __props__.__dict__["reservation_id"] = reservation_id
            __props__.__dict__["slot_capacity"] = slot_capacity
            __props__.__dict__["creation_time"] = None
            __props__.__dict__["update_time"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["location", "project"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Reservation, __self__).__init__(
            'google-native:bigqueryreservation/v1:Reservation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Reservation':
        """
        Get an existing Reservation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = ReservationArgs.__new__(ReservationArgs)

        __props__.__dict__["autoscale"] = None
        __props__.__dict__["concurrency"] = None
        __props__.__dict__["creation_time"] = None
        __props__.__dict__["edition"] = None
        __props__.__dict__["ignore_idle_slots"] = None
        __props__.__dict__["location"] = None
        __props__.__dict__["multi_region_auxiliary"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["project"] = None
        __props__.__dict__["reservation_id"] = None
        __props__.__dict__["slot_capacity"] = None
        __props__.__dict__["update_time"] = None
        return Reservation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def autoscale(self) -> pulumi.Output['outputs.AutoscaleResponse']:
        """
        The configuration parameters for the auto scaling feature. Note this is an alpha feature.
        """
        return pulumi.get(self, "autoscale")

    @property
    @pulumi.getter
    def concurrency(self) -> pulumi.Output[str]:
        """
        Job concurrency target which sets a soft upper bound on the number of jobs that can run concurrently in this reservation. This is a soft target due to asynchronous nature of the system and various optimizations for small queries. Default value is 0 which means that concurrency target will be automatically computed by the system. NOTE: this field is exposed as `target_job_concurrency` in the Information Schema, DDL and BQ CLI.
        """
        return pulumi.get(self, "concurrency")

    @property
    @pulumi.getter(name="creationTime")
    def creation_time(self) -> pulumi.Output[str]:
        """
        Creation time of the reservation.
        """
        return pulumi.get(self, "creation_time")

    @property
    @pulumi.getter
    def edition(self) -> pulumi.Output[str]:
        """
        Edition of the reservation.
        """
        return pulumi.get(self, "edition")

    @property
    @pulumi.getter(name="ignoreIdleSlots")
    def ignore_idle_slots(self) -> pulumi.Output[bool]:
        """
        If false, any query or pipeline job using this reservation will use idle slots from other reservations within the same admin project. If true, a query or pipeline job using this reservation will execute with the slot capacity specified in the slot_capacity field at most.
        """
        return pulumi.get(self, "ignore_idle_slots")

    @property
    @pulumi.getter
    def location(self) -> pulumi.Output[str]:
        return pulumi.get(self, "location")

    @property
    @pulumi.getter(name="multiRegionAuxiliary")
    def multi_region_auxiliary(self) -> pulumi.Output[bool]:
        """
        Applicable only for reservations located within one of the BigQuery multi-regions (US or EU). If set to true, this reservation is placed in the organization's secondary region which is designated for disaster recovery purposes. If false, this reservation is placed in the organization's default region. NOTE: this is a preview feature. Project must be allow-listed in order to set this field.
        """
        return pulumi.get(self, "multi_region_auxiliary")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name of the reservation, e.g., `projects/*/locations/*/reservations/team1-prod`. The reservation_id must only contain lower case alphanumeric characters or dashes. It must start with a letter and must not end with a dash. Its maximum length is 64 characters.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def project(self) -> pulumi.Output[str]:
        return pulumi.get(self, "project")

    @property
    @pulumi.getter(name="reservationId")
    def reservation_id(self) -> pulumi.Output[Optional[str]]:
        """
        The reservation ID. It must only contain lower case alphanumeric characters or dashes. It must start with a letter and must not end with a dash. Its maximum length is 64 characters.
        """
        return pulumi.get(self, "reservation_id")

    @property
    @pulumi.getter(name="slotCapacity")
    def slot_capacity(self) -> pulumi.Output[str]:
        """
        Minimum slots available to this reservation. A slot is a unit of computational power in BigQuery, and serves as the unit of parallelism. Queries using this reservation might use more slots during runtime if ignore_idle_slots is set to false. If total slot_capacity of the reservation and its siblings exceeds the total slot_count of all capacity commitments, the request will fail with `google.rpc.Code.RESOURCE_EXHAUSTED`. NOTE: for reservations in US or EU multi-regions, slot capacity constraints are checked separately for default and auxiliary regions. See multi_region_auxiliary flag for more details.
        """
        return pulumi.get(self, "slot_capacity")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        Last update time of the reservation.
        """
        return pulumi.get(self, "update_time")

