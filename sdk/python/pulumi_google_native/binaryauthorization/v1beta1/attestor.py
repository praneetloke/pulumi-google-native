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
from ._inputs import *

__all__ = ['AttestorArgs', 'Attestor']

@pulumi.input_type
class AttestorArgs:
    def __init__(__self__, *,
                 attestor_id: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 user_owned_drydock_note: Optional[pulumi.Input['UserOwnedDrydockNoteArgs']] = None):
        """
        The set of arguments for constructing a Attestor resource.
        :param pulumi.Input[str] attestor_id: Required. The attestors ID.
        :param pulumi.Input[str] description: Optional. A descriptive comment. This field may be updated. The field may be displayed in chooser dialogs.
        :param pulumi.Input[str] etag: Optional. Used to prevent updating the attestor when another request has updated it since it was retrieved.
        :param pulumi.Input[str] name: The resource name, in the format: `projects/*/attestors/*`. This field may not be updated.
        :param pulumi.Input['UserOwnedDrydockNoteArgs'] user_owned_drydock_note: A Drydock ATTESTATION_AUTHORITY Note, created by the user.
        """
        pulumi.set(__self__, "attestor_id", attestor_id)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if etag is not None:
            pulumi.set(__self__, "etag", etag)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if project is not None:
            pulumi.set(__self__, "project", project)
        if user_owned_drydock_note is not None:
            pulumi.set(__self__, "user_owned_drydock_note", user_owned_drydock_note)

    @property
    @pulumi.getter(name="attestorId")
    def attestor_id(self) -> pulumi.Input[str]:
        """
        Required. The attestors ID.
        """
        return pulumi.get(self, "attestor_id")

    @attestor_id.setter
    def attestor_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "attestor_id", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. A descriptive comment. This field may be updated. The field may be displayed in chooser dialogs.
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def etag(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. Used to prevent updating the attestor when another request has updated it since it was retrieved.
        """
        return pulumi.get(self, "etag")

    @etag.setter
    def etag(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "etag", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The resource name, in the format: `projects/*/attestors/*`. This field may not be updated.
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
    @pulumi.getter(name="userOwnedDrydockNote")
    def user_owned_drydock_note(self) -> Optional[pulumi.Input['UserOwnedDrydockNoteArgs']]:
        """
        A Drydock ATTESTATION_AUTHORITY Note, created by the user.
        """
        return pulumi.get(self, "user_owned_drydock_note")

    @user_owned_drydock_note.setter
    def user_owned_drydock_note(self, value: Optional[pulumi.Input['UserOwnedDrydockNoteArgs']]):
        pulumi.set(self, "user_owned_drydock_note", value)


class Attestor(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attestor_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 user_owned_drydock_note: Optional[pulumi.Input[pulumi.InputType['UserOwnedDrydockNoteArgs']]] = None,
                 __props__=None):
        """
        Creates an attestor, and returns a copy of the new attestor. Returns NOT_FOUND if the project does not exist, INVALID_ARGUMENT if the request is malformed, ALREADY_EXISTS if the attestor already exists.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] attestor_id: Required. The attestors ID.
        :param pulumi.Input[str] description: Optional. A descriptive comment. This field may be updated. The field may be displayed in chooser dialogs.
        :param pulumi.Input[str] etag: Optional. Used to prevent updating the attestor when another request has updated it since it was retrieved.
        :param pulumi.Input[str] name: The resource name, in the format: `projects/*/attestors/*`. This field may not be updated.
        :param pulumi.Input[pulumi.InputType['UserOwnedDrydockNoteArgs']] user_owned_drydock_note: A Drydock ATTESTATION_AUTHORITY Note, created by the user.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AttestorArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates an attestor, and returns a copy of the new attestor. Returns NOT_FOUND if the project does not exist, INVALID_ARGUMENT if the request is malformed, ALREADY_EXISTS if the attestor already exists.

        :param str resource_name: The name of the resource.
        :param AttestorArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AttestorArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 attestor_id: Optional[pulumi.Input[str]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 etag: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 project: Optional[pulumi.Input[str]] = None,
                 user_owned_drydock_note: Optional[pulumi.Input[pulumi.InputType['UserOwnedDrydockNoteArgs']]] = None,
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
            __props__ = AttestorArgs.__new__(AttestorArgs)

            if attestor_id is None and not opts.urn:
                raise TypeError("Missing required property 'attestor_id'")
            __props__.__dict__["attestor_id"] = attestor_id
            __props__.__dict__["description"] = description
            __props__.__dict__["etag"] = etag
            __props__.__dict__["name"] = name
            __props__.__dict__["project"] = project
            __props__.__dict__["user_owned_drydock_note"] = user_owned_drydock_note
            __props__.__dict__["update_time"] = None
        super(Attestor, __self__).__init__(
            'google-native:binaryauthorization/v1beta1:Attestor',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Attestor':
        """
        Get an existing Attestor resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = AttestorArgs.__new__(AttestorArgs)

        __props__.__dict__["description"] = None
        __props__.__dict__["etag"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["update_time"] = None
        __props__.__dict__["user_owned_drydock_note"] = None
        return Attestor(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        Optional. A descriptive comment. This field may be updated. The field may be displayed in chooser dialogs.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def etag(self) -> pulumi.Output[str]:
        """
        Optional. Used to prevent updating the attestor when another request has updated it since it was retrieved.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The resource name, in the format: `projects/*/attestors/*`. This field may not be updated.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        Time when the attestor was last updated.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter(name="userOwnedDrydockNote")
    def user_owned_drydock_note(self) -> pulumi.Output['outputs.UserOwnedDrydockNoteResponse']:
        """
        A Drydock ATTESTATION_AUTHORITY Note, created by the user.
        """
        return pulumi.get(self, "user_owned_drydock_note")

