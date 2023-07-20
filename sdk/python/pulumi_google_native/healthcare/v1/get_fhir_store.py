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
    'GetFhirStoreResult',
    'AwaitableGetFhirStoreResult',
    'get_fhir_store',
    'get_fhir_store_output',
]

@pulumi.output_type
class GetFhirStoreResult:
    def __init__(__self__, complex_data_type_reference_parsing=None, default_search_handling_strict=None, disable_referential_integrity=None, disable_resource_versioning=None, enable_update_create=None, labels=None, name=None, notification_config=None, notification_configs=None, stream_configs=None, validation_config=None, version=None):
        if complex_data_type_reference_parsing and not isinstance(complex_data_type_reference_parsing, str):
            raise TypeError("Expected argument 'complex_data_type_reference_parsing' to be a str")
        pulumi.set(__self__, "complex_data_type_reference_parsing", complex_data_type_reference_parsing)
        if default_search_handling_strict and not isinstance(default_search_handling_strict, bool):
            raise TypeError("Expected argument 'default_search_handling_strict' to be a bool")
        pulumi.set(__self__, "default_search_handling_strict", default_search_handling_strict)
        if disable_referential_integrity and not isinstance(disable_referential_integrity, bool):
            raise TypeError("Expected argument 'disable_referential_integrity' to be a bool")
        pulumi.set(__self__, "disable_referential_integrity", disable_referential_integrity)
        if disable_resource_versioning and not isinstance(disable_resource_versioning, bool):
            raise TypeError("Expected argument 'disable_resource_versioning' to be a bool")
        pulumi.set(__self__, "disable_resource_versioning", disable_resource_versioning)
        if enable_update_create and not isinstance(enable_update_create, bool):
            raise TypeError("Expected argument 'enable_update_create' to be a bool")
        pulumi.set(__self__, "enable_update_create", enable_update_create)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if notification_config and not isinstance(notification_config, dict):
            raise TypeError("Expected argument 'notification_config' to be a dict")
        pulumi.set(__self__, "notification_config", notification_config)
        if notification_configs and not isinstance(notification_configs, list):
            raise TypeError("Expected argument 'notification_configs' to be a list")
        pulumi.set(__self__, "notification_configs", notification_configs)
        if stream_configs and not isinstance(stream_configs, list):
            raise TypeError("Expected argument 'stream_configs' to be a list")
        pulumi.set(__self__, "stream_configs", stream_configs)
        if validation_config and not isinstance(validation_config, dict):
            raise TypeError("Expected argument 'validation_config' to be a dict")
        pulumi.set(__self__, "validation_config", validation_config)
        if version and not isinstance(version, str):
            raise TypeError("Expected argument 'version' to be a str")
        pulumi.set(__self__, "version", version)

    @property
    @pulumi.getter(name="complexDataTypeReferenceParsing")
    def complex_data_type_reference_parsing(self) -> str:
        """
        Enable parsing of references within complex FHIR data types such as Extensions. If this value is set to ENABLED, then features like referential integrity and Bundle reference rewriting apply to all references. If this flag has not been specified the behavior of the FHIR store will not change, references in complex data types will not be parsed. New stores will have this value set to ENABLED after a notification period. Warning: turning on this flag causes processing existing resources to fail if they contain references to non-existent resources.
        """
        return pulumi.get(self, "complex_data_type_reference_parsing")

    @property
    @pulumi.getter(name="defaultSearchHandlingStrict")
    def default_search_handling_strict(self) -> bool:
        """
        If true, overrides the default search behavior for this FHIR store to `handling=strict` which returns an error for unrecognized search parameters. If false, uses the FHIR specification default `handling=lenient` which ignores unrecognized search parameters. The handling can always be changed from the default on an individual API call by setting the HTTP header `Prefer: handling=strict` or `Prefer: handling=lenient`.
        """
        return pulumi.get(self, "default_search_handling_strict")

    @property
    @pulumi.getter(name="disableReferentialIntegrity")
    def disable_referential_integrity(self) -> bool:
        """
        Immutable. Whether to disable referential integrity in this FHIR store. This field is immutable after FHIR store creation. The default value is false, meaning that the API enforces referential integrity and fails the requests that result in inconsistent state in the FHIR store. When this field is set to true, the API skips referential integrity checks. Consequently, operations that rely on references, such as GetPatientEverything, do not return all the results if broken references exist.
        """
        return pulumi.get(self, "disable_referential_integrity")

    @property
    @pulumi.getter(name="disableResourceVersioning")
    def disable_resource_versioning(self) -> bool:
        """
        Immutable. Whether to disable resource versioning for this FHIR store. This field can not be changed after the creation of FHIR store. If set to false, which is the default behavior, all write operations cause historical versions to be recorded automatically. The historical versions can be fetched through the history APIs, but cannot be updated. If set to true, no historical versions are kept. The server sends errors for attempts to read the historical versions.
        """
        return pulumi.get(self, "disable_resource_versioning")

    @property
    @pulumi.getter(name="enableUpdateCreate")
    def enable_update_create(self) -> bool:
        """
        Whether this FHIR store has the [updateCreate capability](https://www.hl7.org/fhir/capabilitystatement-definitions.html#CapabilityStatement.rest.resource.updateCreate). This determines if the client can use an Update operation to create a new resource with a client-specified ID. If false, all IDs are server-assigned through the Create operation and attempts to update a non-existent resource return errors. It is strongly advised not to include or encode any sensitive data such as patient identifiers in client-specified resource IDs. Those IDs are part of the FHIR resource path recorded in Cloud audit logs and Pub/Sub notifications. Those IDs can also be contained in reference fields within other resources.
        """
        return pulumi.get(self, "enable_update_create")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        User-supplied key-value pairs used to organize FHIR stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \\p{Ll}\\p{Lo}{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\\p{Ll}\\p{Lo}\\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Resource name of the FHIR store, of the form `projects/{project_id}/datasets/{dataset_id}/fhirStores/{fhir_store_id}`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="notificationConfig")
    def notification_config(self) -> 'outputs.NotificationConfigResponse':
        """
        Deprecated. Use `notification_configs` instead. If non-empty, publish all resource modifications of this FHIR store to this destination. The Pub/Sub message attributes contain a map with a string describing the action that has triggered the notification. For example, "action":"CreateResource".
        """
        warnings.warn("""Deprecated. Use `notification_configs` instead. If non-empty, publish all resource modifications of this FHIR store to this destination. The Pub/Sub message attributes contain a map with a string describing the action that has triggered the notification. For example, \"action\":\"CreateResource\".""", DeprecationWarning)
        pulumi.log.warn("""notification_config is deprecated: Deprecated. Use `notification_configs` instead. If non-empty, publish all resource modifications of this FHIR store to this destination. The Pub/Sub message attributes contain a map with a string describing the action that has triggered the notification. For example, \"action\":\"CreateResource\".""")

        return pulumi.get(self, "notification_config")

    @property
    @pulumi.getter(name="notificationConfigs")
    def notification_configs(self) -> Sequence['outputs.FhirNotificationConfigResponse']:
        """
        Specifies where and whether to send notifications upon changes to a FHIR store.
        """
        return pulumi.get(self, "notification_configs")

    @property
    @pulumi.getter(name="streamConfigs")
    def stream_configs(self) -> Sequence['outputs.StreamConfigResponse']:
        """
        A list of streaming configs that configure the destinations of streaming export for every resource mutation in this FHIR store. Each store is allowed to have up to 10 streaming configs. After a new config is added, the next resource mutation is streamed to the new location in addition to the existing ones. When a location is removed from the list, the server stops streaming to that location. Before adding a new config, you must add the required [`bigquery.dataEditor`](https://cloud.google.com/bigquery/docs/access-control#bigquery.dataEditor) role to your project's **Cloud Healthcare Service Agent** [service account](https://cloud.google.com/iam/docs/service-accounts). Some lag (typically on the order of dozens of seconds) is expected before the results show up in the streaming destination.
        """
        return pulumi.get(self, "stream_configs")

    @property
    @pulumi.getter(name="validationConfig")
    def validation_config(self) -> 'outputs.ValidationConfigResponse':
        """
        Configuration for how to validate incoming FHIR resources against configured profiles.
        """
        return pulumi.get(self, "validation_config")

    @property
    @pulumi.getter
    def version(self) -> str:
        """
        Immutable. The FHIR specification version that this FHIR store supports natively. This field is immutable after store creation. Requests are rejected if they contain FHIR resources of a different version. Version is required for every FHIR store.
        """
        return pulumi.get(self, "version")


class AwaitableGetFhirStoreResult(GetFhirStoreResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetFhirStoreResult(
            complex_data_type_reference_parsing=self.complex_data_type_reference_parsing,
            default_search_handling_strict=self.default_search_handling_strict,
            disable_referential_integrity=self.disable_referential_integrity,
            disable_resource_versioning=self.disable_resource_versioning,
            enable_update_create=self.enable_update_create,
            labels=self.labels,
            name=self.name,
            notification_config=self.notification_config,
            notification_configs=self.notification_configs,
            stream_configs=self.stream_configs,
            validation_config=self.validation_config,
            version=self.version)


def get_fhir_store(dataset_id: Optional[str] = None,
                   fhir_store_id: Optional[str] = None,
                   location: Optional[str] = None,
                   project: Optional[str] = None,
                   opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetFhirStoreResult:
    """
    Gets the configuration of the specified FHIR store.
    """
    __args__ = dict()
    __args__['datasetId'] = dataset_id
    __args__['fhirStoreId'] = fhir_store_id
    __args__['location'] = location
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:healthcare/v1:getFhirStore', __args__, opts=opts, typ=GetFhirStoreResult).value

    return AwaitableGetFhirStoreResult(
        complex_data_type_reference_parsing=pulumi.get(__ret__, 'complex_data_type_reference_parsing'),
        default_search_handling_strict=pulumi.get(__ret__, 'default_search_handling_strict'),
        disable_referential_integrity=pulumi.get(__ret__, 'disable_referential_integrity'),
        disable_resource_versioning=pulumi.get(__ret__, 'disable_resource_versioning'),
        enable_update_create=pulumi.get(__ret__, 'enable_update_create'),
        labels=pulumi.get(__ret__, 'labels'),
        name=pulumi.get(__ret__, 'name'),
        notification_config=pulumi.get(__ret__, 'notification_config'),
        notification_configs=pulumi.get(__ret__, 'notification_configs'),
        stream_configs=pulumi.get(__ret__, 'stream_configs'),
        validation_config=pulumi.get(__ret__, 'validation_config'),
        version=pulumi.get(__ret__, 'version'))


@_utilities.lift_output_func(get_fhir_store)
def get_fhir_store_output(dataset_id: Optional[pulumi.Input[str]] = None,
                          fhir_store_id: Optional[pulumi.Input[str]] = None,
                          location: Optional[pulumi.Input[str]] = None,
                          project: Optional[pulumi.Input[Optional[str]]] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetFhirStoreResult]:
    """
    Gets the configuration of the specified FHIR store.
    """
    ...
