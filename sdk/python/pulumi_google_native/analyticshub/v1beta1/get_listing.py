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
    'GetListingResult',
    'AwaitableGetListingResult',
    'get_listing',
    'get_listing_output',
]

@pulumi.output_type
class GetListingResult:
    def __init__(__self__, bigquery_dataset=None, categories=None, data_provider=None, description=None, display_name=None, documentation=None, icon=None, name=None, primary_contact=None, publisher=None, request_access=None, state=None):
        if bigquery_dataset and not isinstance(bigquery_dataset, dict):
            raise TypeError("Expected argument 'bigquery_dataset' to be a dict")
        pulumi.set(__self__, "bigquery_dataset", bigquery_dataset)
        if categories and not isinstance(categories, list):
            raise TypeError("Expected argument 'categories' to be a list")
        pulumi.set(__self__, "categories", categories)
        if data_provider and not isinstance(data_provider, dict):
            raise TypeError("Expected argument 'data_provider' to be a dict")
        pulumi.set(__self__, "data_provider", data_provider)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if documentation and not isinstance(documentation, str):
            raise TypeError("Expected argument 'documentation' to be a str")
        pulumi.set(__self__, "documentation", documentation)
        if icon and not isinstance(icon, str):
            raise TypeError("Expected argument 'icon' to be a str")
        pulumi.set(__self__, "icon", icon)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if primary_contact and not isinstance(primary_contact, str):
            raise TypeError("Expected argument 'primary_contact' to be a str")
        pulumi.set(__self__, "primary_contact", primary_contact)
        if publisher and not isinstance(publisher, dict):
            raise TypeError("Expected argument 'publisher' to be a dict")
        pulumi.set(__self__, "publisher", publisher)
        if request_access and not isinstance(request_access, str):
            raise TypeError("Expected argument 'request_access' to be a str")
        pulumi.set(__self__, "request_access", request_access)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)

    @property
    @pulumi.getter(name="bigqueryDataset")
    def bigquery_dataset(self) -> 'outputs.BigQueryDatasetSourceResponse':
        """
        Shared dataset i.e. BigQuery dataset source.
        """
        return pulumi.get(self, "bigquery_dataset")

    @property
    @pulumi.getter
    def categories(self) -> Sequence[str]:
        """
        Optional. Categories of the listing. Up to two categories are allowed.
        """
        return pulumi.get(self, "categories")

    @property
    @pulumi.getter(name="dataProvider")
    def data_provider(self) -> 'outputs.DataProviderResponse':
        """
        Optional. Details of the data provider who owns the source data.
        """
        return pulumi.get(self, "data_provider")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Optional. Short description of the listing. The description must not contain Unicode non-characters and C0 and C1 control codes except tabs (HT), new lines (LF), carriage returns (CR), and page breaks (FF). Default value is an empty string. Max length: 2000 bytes.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        Human-readable display name of the listing. The display name must contain only Unicode letters, numbers (0-9), underscores (_), dashes (-), spaces ( ), and can't start or end with spaces. Default value is an empty string. Max length: 63 bytes.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def documentation(self) -> str:
        """
        Optional. Documentation describing the listing.
        """
        return pulumi.get(self, "documentation")

    @property
    @pulumi.getter
    def icon(self) -> str:
        """
        Optional. Base64 encoded image representing the listing. Max Size: 3.0MiB Expected image dimensions are 512x512 pixels, however the API only performs validation on size of the encoded data. Note: For byte fields, the contents of the field are base64-encoded (which increases the size of the data by 33-36%) when using JSON on the wire.
        """
        return pulumi.get(self, "icon")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name of the listing. e.g. `projects/myproject/locations/US/dataExchanges/123/listings/456`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="primaryContact")
    def primary_contact(self) -> str:
        """
        Optional. Email or URL of the primary point of contact of the listing. Max Length: 1000 bytes.
        """
        return pulumi.get(self, "primary_contact")

    @property
    @pulumi.getter
    def publisher(self) -> 'outputs.PublisherResponse':
        """
        Optional. Details of the publisher who owns the listing and who can share the source data.
        """
        return pulumi.get(self, "publisher")

    @property
    @pulumi.getter(name="requestAccess")
    def request_access(self) -> str:
        """
        Optional. Email or URL of the request access of the listing. Subscribers can use this reference to request access. Max Length: 1000 bytes.
        """
        return pulumi.get(self, "request_access")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        Current state of the listing.
        """
        return pulumi.get(self, "state")


class AwaitableGetListingResult(GetListingResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetListingResult(
            bigquery_dataset=self.bigquery_dataset,
            categories=self.categories,
            data_provider=self.data_provider,
            description=self.description,
            display_name=self.display_name,
            documentation=self.documentation,
            icon=self.icon,
            name=self.name,
            primary_contact=self.primary_contact,
            publisher=self.publisher,
            request_access=self.request_access,
            state=self.state)


def get_listing(data_exchange_id: Optional[str] = None,
                listing_id: Optional[str] = None,
                location: Optional[str] = None,
                project: Optional[str] = None,
                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetListingResult:
    """
    Gets the details of a listing.
    """
    __args__ = dict()
    __args__['dataExchangeId'] = data_exchange_id
    __args__['listingId'] = listing_id
    __args__['location'] = location
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:analyticshub/v1beta1:getListing', __args__, opts=opts, typ=GetListingResult).value

    return AwaitableGetListingResult(
        bigquery_dataset=__ret__.bigquery_dataset,
        categories=__ret__.categories,
        data_provider=__ret__.data_provider,
        description=__ret__.description,
        display_name=__ret__.display_name,
        documentation=__ret__.documentation,
        icon=__ret__.icon,
        name=__ret__.name,
        primary_contact=__ret__.primary_contact,
        publisher=__ret__.publisher,
        request_access=__ret__.request_access,
        state=__ret__.state)


@_utilities.lift_output_func(get_listing)
def get_listing_output(data_exchange_id: Optional[pulumi.Input[str]] = None,
                       listing_id: Optional[pulumi.Input[str]] = None,
                       location: Optional[pulumi.Input[str]] = None,
                       project: Optional[pulumi.Input[Optional[str]]] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetListingResult]:
    """
    Gets the details of a listing.
    """
    ...