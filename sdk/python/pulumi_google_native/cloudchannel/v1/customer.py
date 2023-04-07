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
from ._inputs import *

__all__ = ['CustomerArgs', 'Customer']

@pulumi.input_type
class CustomerArgs:
    def __init__(__self__, *,
                 account_id: pulumi.Input[str],
                 channel_partner_link_id: pulumi.Input[str],
                 domain: pulumi.Input[str],
                 org_display_name: pulumi.Input[str],
                 org_postal_address: pulumi.Input['GoogleTypePostalAddressArgs'],
                 alternate_email: Optional[pulumi.Input[str]] = None,
                 channel_partner_id: Optional[pulumi.Input[str]] = None,
                 correlation_id: Optional[pulumi.Input[str]] = None,
                 language_code: Optional[pulumi.Input[str]] = None,
                 primary_contact_info: Optional[pulumi.Input['GoogleCloudChannelV1ContactInfoArgs']] = None):
        """
        The set of arguments for constructing a Customer resource.
        :param pulumi.Input[str] domain: The customer's primary domain. Must match the primary contact email's domain.
        :param pulumi.Input[str] org_display_name: Name of the organization that the customer entity represents.
        :param pulumi.Input['GoogleTypePostalAddressArgs'] org_postal_address: The organization address for the customer. To enforce US laws and embargoes, we require a region and zip code. You must provide valid addresses for every customer. To set the customer's language, use the Customer-level language code.
        :param pulumi.Input[str] alternate_email: Secondary contact email. You need to provide an alternate email to create different domains if a primary contact email already exists. Users will receive a notification with credentials when you create an admin.google.com account. Secondary emails are also recovery email addresses. Alternate emails are optional when you create Team customers.
        :param pulumi.Input[str] channel_partner_id: Cloud Identity ID of the customer's channel partner. Populated only if a channel partner exists for this customer.
        :param pulumi.Input[str] correlation_id: Optional. External CRM ID for the customer. Populated only if a CRM ID exists for this customer.
        :param pulumi.Input[str] language_code: Optional. The BCP-47 language code, such as "en-US" or "sr-Latn". For more information, see https://www.unicode.org/reports/tr35/#Unicode_locale_identifier.
        :param pulumi.Input['GoogleCloudChannelV1ContactInfoArgs'] primary_contact_info: Primary contact info.
        """
        pulumi.set(__self__, "account_id", account_id)
        pulumi.set(__self__, "channel_partner_link_id", channel_partner_link_id)
        pulumi.set(__self__, "domain", domain)
        pulumi.set(__self__, "org_display_name", org_display_name)
        pulumi.set(__self__, "org_postal_address", org_postal_address)
        if alternate_email is not None:
            pulumi.set(__self__, "alternate_email", alternate_email)
        if channel_partner_id is not None:
            pulumi.set(__self__, "channel_partner_id", channel_partner_id)
        if correlation_id is not None:
            pulumi.set(__self__, "correlation_id", correlation_id)
        if language_code is not None:
            pulumi.set(__self__, "language_code", language_code)
        if primary_contact_info is not None:
            pulumi.set(__self__, "primary_contact_info", primary_contact_info)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "account_id")

    @account_id.setter
    def account_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "account_id", value)

    @property
    @pulumi.getter(name="channelPartnerLinkId")
    def channel_partner_link_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "channel_partner_link_id")

    @channel_partner_link_id.setter
    def channel_partner_link_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "channel_partner_link_id", value)

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Input[str]:
        """
        The customer's primary domain. Must match the primary contact email's domain.
        """
        return pulumi.get(self, "domain")

    @domain.setter
    def domain(self, value: pulumi.Input[str]):
        pulumi.set(self, "domain", value)

    @property
    @pulumi.getter(name="orgDisplayName")
    def org_display_name(self) -> pulumi.Input[str]:
        """
        Name of the organization that the customer entity represents.
        """
        return pulumi.get(self, "org_display_name")

    @org_display_name.setter
    def org_display_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "org_display_name", value)

    @property
    @pulumi.getter(name="orgPostalAddress")
    def org_postal_address(self) -> pulumi.Input['GoogleTypePostalAddressArgs']:
        """
        The organization address for the customer. To enforce US laws and embargoes, we require a region and zip code. You must provide valid addresses for every customer. To set the customer's language, use the Customer-level language code.
        """
        return pulumi.get(self, "org_postal_address")

    @org_postal_address.setter
    def org_postal_address(self, value: pulumi.Input['GoogleTypePostalAddressArgs']):
        pulumi.set(self, "org_postal_address", value)

    @property
    @pulumi.getter(name="alternateEmail")
    def alternate_email(self) -> Optional[pulumi.Input[str]]:
        """
        Secondary contact email. You need to provide an alternate email to create different domains if a primary contact email already exists. Users will receive a notification with credentials when you create an admin.google.com account. Secondary emails are also recovery email addresses. Alternate emails are optional when you create Team customers.
        """
        return pulumi.get(self, "alternate_email")

    @alternate_email.setter
    def alternate_email(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "alternate_email", value)

    @property
    @pulumi.getter(name="channelPartnerId")
    def channel_partner_id(self) -> Optional[pulumi.Input[str]]:
        """
        Cloud Identity ID of the customer's channel partner. Populated only if a channel partner exists for this customer.
        """
        return pulumi.get(self, "channel_partner_id")

    @channel_partner_id.setter
    def channel_partner_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "channel_partner_id", value)

    @property
    @pulumi.getter(name="correlationId")
    def correlation_id(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. External CRM ID for the customer. Populated only if a CRM ID exists for this customer.
        """
        return pulumi.get(self, "correlation_id")

    @correlation_id.setter
    def correlation_id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "correlation_id", value)

    @property
    @pulumi.getter(name="languageCode")
    def language_code(self) -> Optional[pulumi.Input[str]]:
        """
        Optional. The BCP-47 language code, such as "en-US" or "sr-Latn". For more information, see https://www.unicode.org/reports/tr35/#Unicode_locale_identifier.
        """
        return pulumi.get(self, "language_code")

    @language_code.setter
    def language_code(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "language_code", value)

    @property
    @pulumi.getter(name="primaryContactInfo")
    def primary_contact_info(self) -> Optional[pulumi.Input['GoogleCloudChannelV1ContactInfoArgs']]:
        """
        Primary contact info.
        """
        return pulumi.get(self, "primary_contact_info")

    @primary_contact_info.setter
    def primary_contact_info(self, value: Optional[pulumi.Input['GoogleCloudChannelV1ContactInfoArgs']]):
        pulumi.set(self, "primary_contact_info", value)


class Customer(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_id: Optional[pulumi.Input[str]] = None,
                 alternate_email: Optional[pulumi.Input[str]] = None,
                 channel_partner_id: Optional[pulumi.Input[str]] = None,
                 channel_partner_link_id: Optional[pulumi.Input[str]] = None,
                 correlation_id: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 language_code: Optional[pulumi.Input[str]] = None,
                 org_display_name: Optional[pulumi.Input[str]] = None,
                 org_postal_address: Optional[pulumi.Input[pulumi.InputType['GoogleTypePostalAddressArgs']]] = None,
                 primary_contact_info: Optional[pulumi.Input[pulumi.InputType['GoogleCloudChannelV1ContactInfoArgs']]] = None,
                 __props__=None):
        """
        Creates a new Customer resource under the reseller or distributor account. Possible error codes: * PERMISSION_DENIED: The reseller account making the request is different from the reseller account in the API request. * INVALID_ARGUMENT: * Required request parameters are missing or invalid. * Domain field value doesn't match the primary email domain. Return value: The newly created Customer resource.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] alternate_email: Secondary contact email. You need to provide an alternate email to create different domains if a primary contact email already exists. Users will receive a notification with credentials when you create an admin.google.com account. Secondary emails are also recovery email addresses. Alternate emails are optional when you create Team customers.
        :param pulumi.Input[str] channel_partner_id: Cloud Identity ID of the customer's channel partner. Populated only if a channel partner exists for this customer.
        :param pulumi.Input[str] correlation_id: Optional. External CRM ID for the customer. Populated only if a CRM ID exists for this customer.
        :param pulumi.Input[str] domain: The customer's primary domain. Must match the primary contact email's domain.
        :param pulumi.Input[str] language_code: Optional. The BCP-47 language code, such as "en-US" or "sr-Latn". For more information, see https://www.unicode.org/reports/tr35/#Unicode_locale_identifier.
        :param pulumi.Input[str] org_display_name: Name of the organization that the customer entity represents.
        :param pulumi.Input[pulumi.InputType['GoogleTypePostalAddressArgs']] org_postal_address: The organization address for the customer. To enforce US laws and embargoes, we require a region and zip code. You must provide valid addresses for every customer. To set the customer's language, use the Customer-level language code.
        :param pulumi.Input[pulumi.InputType['GoogleCloudChannelV1ContactInfoArgs']] primary_contact_info: Primary contact info.
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CustomerArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Creates a new Customer resource under the reseller or distributor account. Possible error codes: * PERMISSION_DENIED: The reseller account making the request is different from the reseller account in the API request. * INVALID_ARGUMENT: * Required request parameters are missing or invalid. * Domain field value doesn't match the primary email domain. Return value: The newly created Customer resource.
        Auto-naming is currently not supported for this resource.

        :param str resource_name: The name of the resource.
        :param CustomerArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CustomerArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 account_id: Optional[pulumi.Input[str]] = None,
                 alternate_email: Optional[pulumi.Input[str]] = None,
                 channel_partner_id: Optional[pulumi.Input[str]] = None,
                 channel_partner_link_id: Optional[pulumi.Input[str]] = None,
                 correlation_id: Optional[pulumi.Input[str]] = None,
                 domain: Optional[pulumi.Input[str]] = None,
                 language_code: Optional[pulumi.Input[str]] = None,
                 org_display_name: Optional[pulumi.Input[str]] = None,
                 org_postal_address: Optional[pulumi.Input[pulumi.InputType['GoogleTypePostalAddressArgs']]] = None,
                 primary_contact_info: Optional[pulumi.Input[pulumi.InputType['GoogleCloudChannelV1ContactInfoArgs']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CustomerArgs.__new__(CustomerArgs)

            if account_id is None and not opts.urn:
                raise TypeError("Missing required property 'account_id'")
            __props__.__dict__["account_id"] = account_id
            __props__.__dict__["alternate_email"] = alternate_email
            __props__.__dict__["channel_partner_id"] = channel_partner_id
            if channel_partner_link_id is None and not opts.urn:
                raise TypeError("Missing required property 'channel_partner_link_id'")
            __props__.__dict__["channel_partner_link_id"] = channel_partner_link_id
            __props__.__dict__["correlation_id"] = correlation_id
            if domain is None and not opts.urn:
                raise TypeError("Missing required property 'domain'")
            __props__.__dict__["domain"] = domain
            __props__.__dict__["language_code"] = language_code
            if org_display_name is None and not opts.urn:
                raise TypeError("Missing required property 'org_display_name'")
            __props__.__dict__["org_display_name"] = org_display_name
            if org_postal_address is None and not opts.urn:
                raise TypeError("Missing required property 'org_postal_address'")
            __props__.__dict__["org_postal_address"] = org_postal_address
            __props__.__dict__["primary_contact_info"] = primary_contact_info
            __props__.__dict__["cloud_identity_id"] = None
            __props__.__dict__["cloud_identity_info"] = None
            __props__.__dict__["create_time"] = None
            __props__.__dict__["name"] = None
            __props__.__dict__["update_time"] = None
        replace_on_changes = pulumi.ResourceOptions(replace_on_changes=["account_id", "channel_partner_link_id"])
        opts = pulumi.ResourceOptions.merge(opts, replace_on_changes)
        super(Customer, __self__).__init__(
            'google-native:cloudchannel/v1:Customer',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None) -> 'Customer':
        """
        Get an existing Customer resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = CustomerArgs.__new__(CustomerArgs)

        __props__.__dict__["account_id"] = None
        __props__.__dict__["alternate_email"] = None
        __props__.__dict__["channel_partner_id"] = None
        __props__.__dict__["channel_partner_link_id"] = None
        __props__.__dict__["cloud_identity_id"] = None
        __props__.__dict__["cloud_identity_info"] = None
        __props__.__dict__["correlation_id"] = None
        __props__.__dict__["create_time"] = None
        __props__.__dict__["domain"] = None
        __props__.__dict__["language_code"] = None
        __props__.__dict__["name"] = None
        __props__.__dict__["org_display_name"] = None
        __props__.__dict__["org_postal_address"] = None
        __props__.__dict__["primary_contact_info"] = None
        __props__.__dict__["update_time"] = None
        return Customer(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accountId")
    def account_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "account_id")

    @property
    @pulumi.getter(name="alternateEmail")
    def alternate_email(self) -> pulumi.Output[str]:
        """
        Secondary contact email. You need to provide an alternate email to create different domains if a primary contact email already exists. Users will receive a notification with credentials when you create an admin.google.com account. Secondary emails are also recovery email addresses. Alternate emails are optional when you create Team customers.
        """
        return pulumi.get(self, "alternate_email")

    @property
    @pulumi.getter(name="channelPartnerId")
    def channel_partner_id(self) -> pulumi.Output[str]:
        """
        Cloud Identity ID of the customer's channel partner. Populated only if a channel partner exists for this customer.
        """
        return pulumi.get(self, "channel_partner_id")

    @property
    @pulumi.getter(name="channelPartnerLinkId")
    def channel_partner_link_id(self) -> pulumi.Output[str]:
        return pulumi.get(self, "channel_partner_link_id")

    @property
    @pulumi.getter(name="cloudIdentityId")
    def cloud_identity_id(self) -> pulumi.Output[str]:
        """
        The customer's Cloud Identity ID if the customer has a Cloud Identity resource.
        """
        return pulumi.get(self, "cloud_identity_id")

    @property
    @pulumi.getter(name="cloudIdentityInfo")
    def cloud_identity_info(self) -> pulumi.Output['outputs.GoogleCloudChannelV1CloudIdentityInfoResponse']:
        """
        Cloud Identity information for the customer. Populated only if a Cloud Identity account exists for this customer.
        """
        return pulumi.get(self, "cloud_identity_info")

    @property
    @pulumi.getter(name="correlationId")
    def correlation_id(self) -> pulumi.Output[str]:
        """
        Optional. External CRM ID for the customer. Populated only if a CRM ID exists for this customer.
        """
        return pulumi.get(self, "correlation_id")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> pulumi.Output[str]:
        """
        Time when the customer was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def domain(self) -> pulumi.Output[str]:
        """
        The customer's primary domain. Must match the primary contact email's domain.
        """
        return pulumi.get(self, "domain")

    @property
    @pulumi.getter(name="languageCode")
    def language_code(self) -> pulumi.Output[str]:
        """
        Optional. The BCP-47 language code, such as "en-US" or "sr-Latn". For more information, see https://www.unicode.org/reports/tr35/#Unicode_locale_identifier.
        """
        return pulumi.get(self, "language_code")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        Resource name of the customer. Format: accounts/{account_id}/customers/{customer_id}
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="orgDisplayName")
    def org_display_name(self) -> pulumi.Output[str]:
        """
        Name of the organization that the customer entity represents.
        """
        return pulumi.get(self, "org_display_name")

    @property
    @pulumi.getter(name="orgPostalAddress")
    def org_postal_address(self) -> pulumi.Output['outputs.GoogleTypePostalAddressResponse']:
        """
        The organization address for the customer. To enforce US laws and embargoes, we require a region and zip code. You must provide valid addresses for every customer. To set the customer's language, use the Customer-level language code.
        """
        return pulumi.get(self, "org_postal_address")

    @property
    @pulumi.getter(name="primaryContactInfo")
    def primary_contact_info(self) -> pulumi.Output['outputs.GoogleCloudChannelV1ContactInfoResponse']:
        """
        Primary contact info.
        """
        return pulumi.get(self, "primary_contact_info")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> pulumi.Output[str]:
        """
        Time when the customer was updated.
        """
        return pulumi.get(self, "update_time")

