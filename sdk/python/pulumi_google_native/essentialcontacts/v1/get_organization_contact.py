# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities

__all__ = [
    'GetOrganizationContactResult',
    'AwaitableGetOrganizationContactResult',
    'get_organization_contact',
    'get_organization_contact_output',
]

@pulumi.output_type
class GetOrganizationContactResult:
    def __init__(__self__, email=None, language_tag=None, name=None, notification_category_subscriptions=None, validate_time=None, validation_state=None):
        if email and not isinstance(email, str):
            raise TypeError("Expected argument 'email' to be a str")
        pulumi.set(__self__, "email", email)
        if language_tag and not isinstance(language_tag, str):
            raise TypeError("Expected argument 'language_tag' to be a str")
        pulumi.set(__self__, "language_tag", language_tag)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if notification_category_subscriptions and not isinstance(notification_category_subscriptions, list):
            raise TypeError("Expected argument 'notification_category_subscriptions' to be a list")
        pulumi.set(__self__, "notification_category_subscriptions", notification_category_subscriptions)
        if validate_time and not isinstance(validate_time, str):
            raise TypeError("Expected argument 'validate_time' to be a str")
        pulumi.set(__self__, "validate_time", validate_time)
        if validation_state and not isinstance(validation_state, str):
            raise TypeError("Expected argument 'validation_state' to be a str")
        pulumi.set(__self__, "validation_state", validation_state)

    @property
    @pulumi.getter
    def email(self) -> str:
        """
        The email address to send notifications to. The email address does not need to be a Google Account.
        """
        return pulumi.get(self, "email")

    @property
    @pulumi.getter(name="languageTag")
    def language_tag(self) -> str:
        """
        The preferred language for notifications, as a ISO 639-1 language code. See [Supported languages](https://cloud.google.com/resource-manager/docs/managing-notification-contacts#supported-languages) for a list of supported languages.
        """
        return pulumi.get(self, "language_tag")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The identifier for the contact. Format: {resource_type}/{resource_id}/contacts/{contact_id}
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notificationCategorySubscriptions")
    def notification_category_subscriptions(self) -> Sequence[str]:
        """
        The categories of notifications that the contact will receive communications for.
        """
        return pulumi.get(self, "notification_category_subscriptions")

    @property
    @pulumi.getter(name="validateTime")
    def validate_time(self) -> str:
        """
        The last time the validation_state was updated, either manually or automatically. A contact is considered stale if its validation state was updated more than 1 year ago.
        """
        return pulumi.get(self, "validate_time")

    @property
    @pulumi.getter(name="validationState")
    def validation_state(self) -> str:
        """
        The validity of the contact. A contact is considered valid if it is the correct recipient for notifications for a particular resource.
        """
        return pulumi.get(self, "validation_state")


class AwaitableGetOrganizationContactResult(GetOrganizationContactResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetOrganizationContactResult(
            email=self.email,
            language_tag=self.language_tag,
            name=self.name,
            notification_category_subscriptions=self.notification_category_subscriptions,
            validate_time=self.validate_time,
            validation_state=self.validation_state)


def get_organization_contact(contact_id: Optional[str] = None,
                             organization_id: Optional[str] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetOrganizationContactResult:
    """
    Gets a single contact.
    """
    __args__ = dict()
    __args__['contactId'] = contact_id
    __args__['organizationId'] = organization_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:essentialcontacts/v1:getOrganizationContact', __args__, opts=opts, typ=GetOrganizationContactResult).value

    return AwaitableGetOrganizationContactResult(
        email=__ret__.email,
        language_tag=__ret__.language_tag,
        name=__ret__.name,
        notification_category_subscriptions=__ret__.notification_category_subscriptions,
        validate_time=__ret__.validate_time,
        validation_state=__ret__.validation_state)


@_utilities.lift_output_func(get_organization_contact)
def get_organization_contact_output(contact_id: Optional[pulumi.Input[str]] = None,
                                    organization_id: Optional[pulumi.Input[str]] = None,
                                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetOrganizationContactResult]:
    """
    Gets a single contact.
    """
    ...
