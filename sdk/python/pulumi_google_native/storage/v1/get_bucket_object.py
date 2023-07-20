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
    'GetBucketObjectResult',
    'AwaitableGetBucketObjectResult',
    'get_bucket_object',
    'get_bucket_object_output',
]

@pulumi.output_type
class GetBucketObjectResult:
    def __init__(__self__, acl=None, bucket=None, cache_control=None, component_count=None, content_disposition=None, content_encoding=None, content_language=None, content_type=None, crc32c=None, custom_time=None, customer_encryption=None, etag=None, event_based_hold=None, generation=None, kind=None, kms_key_name=None, md5_hash=None, media_link=None, metadata=None, metageneration=None, name=None, owner=None, retention_expiration_time=None, self_link=None, size=None, storage_class=None, temporary_hold=None, time_created=None, time_deleted=None, time_storage_class_updated=None, updated=None):
        if acl and not isinstance(acl, list):
            raise TypeError("Expected argument 'acl' to be a list")
        pulumi.set(__self__, "acl", acl)
        if bucket and not isinstance(bucket, str):
            raise TypeError("Expected argument 'bucket' to be a str")
        pulumi.set(__self__, "bucket", bucket)
        if cache_control and not isinstance(cache_control, str):
            raise TypeError("Expected argument 'cache_control' to be a str")
        pulumi.set(__self__, "cache_control", cache_control)
        if component_count and not isinstance(component_count, int):
            raise TypeError("Expected argument 'component_count' to be a int")
        pulumi.set(__self__, "component_count", component_count)
        if content_disposition and not isinstance(content_disposition, str):
            raise TypeError("Expected argument 'content_disposition' to be a str")
        pulumi.set(__self__, "content_disposition", content_disposition)
        if content_encoding and not isinstance(content_encoding, str):
            raise TypeError("Expected argument 'content_encoding' to be a str")
        pulumi.set(__self__, "content_encoding", content_encoding)
        if content_language and not isinstance(content_language, str):
            raise TypeError("Expected argument 'content_language' to be a str")
        pulumi.set(__self__, "content_language", content_language)
        if content_type and not isinstance(content_type, str):
            raise TypeError("Expected argument 'content_type' to be a str")
        pulumi.set(__self__, "content_type", content_type)
        if crc32c and not isinstance(crc32c, str):
            raise TypeError("Expected argument 'crc32c' to be a str")
        pulumi.set(__self__, "crc32c", crc32c)
        if custom_time and not isinstance(custom_time, str):
            raise TypeError("Expected argument 'custom_time' to be a str")
        pulumi.set(__self__, "custom_time", custom_time)
        if customer_encryption and not isinstance(customer_encryption, dict):
            raise TypeError("Expected argument 'customer_encryption' to be a dict")
        pulumi.set(__self__, "customer_encryption", customer_encryption)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if event_based_hold and not isinstance(event_based_hold, bool):
            raise TypeError("Expected argument 'event_based_hold' to be a bool")
        pulumi.set(__self__, "event_based_hold", event_based_hold)
        if generation and not isinstance(generation, str):
            raise TypeError("Expected argument 'generation' to be a str")
        pulumi.set(__self__, "generation", generation)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if kms_key_name and not isinstance(kms_key_name, str):
            raise TypeError("Expected argument 'kms_key_name' to be a str")
        pulumi.set(__self__, "kms_key_name", kms_key_name)
        if md5_hash and not isinstance(md5_hash, str):
            raise TypeError("Expected argument 'md5_hash' to be a str")
        pulumi.set(__self__, "md5_hash", md5_hash)
        if media_link and not isinstance(media_link, str):
            raise TypeError("Expected argument 'media_link' to be a str")
        pulumi.set(__self__, "media_link", media_link)
        if metadata and not isinstance(metadata, dict):
            raise TypeError("Expected argument 'metadata' to be a dict")
        pulumi.set(__self__, "metadata", metadata)
        if metageneration and not isinstance(metageneration, str):
            raise TypeError("Expected argument 'metageneration' to be a str")
        pulumi.set(__self__, "metageneration", metageneration)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if owner and not isinstance(owner, dict):
            raise TypeError("Expected argument 'owner' to be a dict")
        pulumi.set(__self__, "owner", owner)
        if retention_expiration_time and not isinstance(retention_expiration_time, str):
            raise TypeError("Expected argument 'retention_expiration_time' to be a str")
        pulumi.set(__self__, "retention_expiration_time", retention_expiration_time)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if size and not isinstance(size, str):
            raise TypeError("Expected argument 'size' to be a str")
        pulumi.set(__self__, "size", size)
        if storage_class and not isinstance(storage_class, str):
            raise TypeError("Expected argument 'storage_class' to be a str")
        pulumi.set(__self__, "storage_class", storage_class)
        if temporary_hold and not isinstance(temporary_hold, bool):
            raise TypeError("Expected argument 'temporary_hold' to be a bool")
        pulumi.set(__self__, "temporary_hold", temporary_hold)
        if time_created and not isinstance(time_created, str):
            raise TypeError("Expected argument 'time_created' to be a str")
        pulumi.set(__self__, "time_created", time_created)
        if time_deleted and not isinstance(time_deleted, str):
            raise TypeError("Expected argument 'time_deleted' to be a str")
        pulumi.set(__self__, "time_deleted", time_deleted)
        if time_storage_class_updated and not isinstance(time_storage_class_updated, str):
            raise TypeError("Expected argument 'time_storage_class_updated' to be a str")
        pulumi.set(__self__, "time_storage_class_updated", time_storage_class_updated)
        if updated and not isinstance(updated, str):
            raise TypeError("Expected argument 'updated' to be a str")
        pulumi.set(__self__, "updated", updated)

    @property
    @pulumi.getter
    def acl(self) -> Sequence['outputs.ObjectAccessControlResponse']:
        """
        Access controls on the object.
        """
        return pulumi.get(self, "acl")

    @property
    @pulumi.getter
    def bucket(self) -> str:
        """
        The name of the bucket containing this object.
        """
        return pulumi.get(self, "bucket")

    @property
    @pulumi.getter(name="cacheControl")
    def cache_control(self) -> str:
        """
        Cache-Control directive for the object data. If omitted, and the object is accessible to all anonymous users, the default will be public, max-age=3600.
        """
        return pulumi.get(self, "cache_control")

    @property
    @pulumi.getter(name="componentCount")
    def component_count(self) -> int:
        """
        Number of underlying components that make up this object. Components are accumulated by compose operations.
        """
        return pulumi.get(self, "component_count")

    @property
    @pulumi.getter(name="contentDisposition")
    def content_disposition(self) -> str:
        """
        Content-Disposition of the object data.
        """
        return pulumi.get(self, "content_disposition")

    @property
    @pulumi.getter(name="contentEncoding")
    def content_encoding(self) -> str:
        """
        Content-Encoding of the object data.
        """
        return pulumi.get(self, "content_encoding")

    @property
    @pulumi.getter(name="contentLanguage")
    def content_language(self) -> str:
        """
        Content-Language of the object data.
        """
        return pulumi.get(self, "content_language")

    @property
    @pulumi.getter(name="contentType")
    def content_type(self) -> str:
        """
        Content-Type of the object data. If an object is stored without a Content-Type, it is served as application/octet-stream.
        """
        return pulumi.get(self, "content_type")

    @property
    @pulumi.getter
    def crc32c(self) -> str:
        """
        CRC32c checksum, as described in RFC 4960, Appendix B; encoded using base64 in big-endian byte order. For more information about using the CRC32c checksum, see Hashes and ETags: Best Practices.
        """
        return pulumi.get(self, "crc32c")

    @property
    @pulumi.getter(name="customTime")
    def custom_time(self) -> str:
        """
        A timestamp in RFC 3339 format specified by the user for an object.
        """
        return pulumi.get(self, "custom_time")

    @property
    @pulumi.getter(name="customerEncryption")
    def customer_encryption(self) -> 'outputs.BucketObjectCustomerEncryptionResponse':
        """
        Metadata of customer-supplied encryption key, if the object is encrypted by such a key.
        """
        return pulumi.get(self, "customer_encryption")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        HTTP 1.1 Entity tag for the object.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter(name="eventBasedHold")
    def event_based_hold(self) -> bool:
        """
        Whether an object is under event-based hold. Event-based hold is a way to retain objects until an event occurs, which is signified by the hold's release (i.e. this value is set to false). After being released (set to false), such objects will be subject to bucket-level retention (if any). One sample use case of this flag is for banks to hold loan documents for at least 3 years after loan is paid in full. Here, bucket-level retention is 3 years and the event is the loan being paid in full. In this example, these objects will be held intact for any number of years until the event has occurred (event-based hold on the object is released) and then 3 more years after that. That means retention duration of the objects begins from the moment event-based hold transitioned from true to false.
        """
        return pulumi.get(self, "event_based_hold")

    @property
    @pulumi.getter
    def generation(self) -> str:
        """
        The content generation of this object. Used for object versioning.
        """
        return pulumi.get(self, "generation")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        The kind of item this is. For objects, this is always storage#object.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="kmsKeyName")
    def kms_key_name(self) -> str:
        """
        Not currently supported. Specifying the parameter causes the request to fail with status code 400 - Bad Request.
        """
        return pulumi.get(self, "kms_key_name")

    @property
    @pulumi.getter(name="md5Hash")
    def md5_hash(self) -> str:
        """
        MD5 hash of the data; encoded using base64. For more information about using the MD5 hash, see Hashes and ETags: Best Practices.
        """
        return pulumi.get(self, "md5_hash")

    @property
    @pulumi.getter(name="mediaLink")
    def media_link(self) -> str:
        """
        Media download link.
        """
        return pulumi.get(self, "media_link")

    @property
    @pulumi.getter
    def metadata(self) -> Mapping[str, str]:
        """
        User-provided metadata, in key/value pairs.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter
    def metageneration(self) -> str:
        """
        The version of the metadata for this object at this generation. Used for preconditions and for detecting changes in metadata. A metageneration number is only meaningful in the context of a particular generation of a particular object.
        """
        return pulumi.get(self, "metageneration")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the object. Required if not specified by URL parameter.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def owner(self) -> 'outputs.BucketObjectOwnerResponse':
        """
        The owner of the object. This will always be the uploader of the object.
        """
        return pulumi.get(self, "owner")

    @property
    @pulumi.getter(name="retentionExpirationTime")
    def retention_expiration_time(self) -> str:
        """
        A server-determined value that specifies the earliest time that the object's retention period expires. This value is in RFC 3339 format. Note 1: This field is not provided for objects with an active event-based hold, since retention expiration is unknown until the hold is removed. Note 2: This value can be provided even when temporary hold is set (so that the user can reason about policy without having to first unset the temporary hold).
        """
        return pulumi.get(self, "retention_expiration_time")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        The link to this object.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter
    def size(self) -> str:
        """
        Content-Length of the data in bytes.
        """
        return pulumi.get(self, "size")

    @property
    @pulumi.getter(name="storageClass")
    def storage_class(self) -> str:
        """
        Storage class of the object.
        """
        return pulumi.get(self, "storage_class")

    @property
    @pulumi.getter(name="temporaryHold")
    def temporary_hold(self) -> bool:
        """
        Whether an object is under temporary hold. While this flag is set to true, the object is protected against deletion and overwrites. A common use case of this flag is regulatory investigations where objects need to be retained while the investigation is ongoing. Note that unlike event-based hold, temporary hold does not impact retention expiration time of an object.
        """
        return pulumi.get(self, "temporary_hold")

    @property
    @pulumi.getter(name="timeCreated")
    def time_created(self) -> str:
        """
        The creation time of the object in RFC 3339 format.
        """
        return pulumi.get(self, "time_created")

    @property
    @pulumi.getter(name="timeDeleted")
    def time_deleted(self) -> str:
        """
        The deletion time of the object in RFC 3339 format. Will be returned if and only if this version of the object has been deleted.
        """
        return pulumi.get(self, "time_deleted")

    @property
    @pulumi.getter(name="timeStorageClassUpdated")
    def time_storage_class_updated(self) -> str:
        """
        The time at which the object's storage class was last changed. When the object is initially created, it will be set to timeCreated.
        """
        return pulumi.get(self, "time_storage_class_updated")

    @property
    @pulumi.getter
    def updated(self) -> str:
        """
        The modification time of the object metadata in RFC 3339 format. Set initially to object creation time and then updated whenever any metadata of the object changes. This includes changes made by a requester, such as modifying custom metadata, as well as changes made by Cloud Storage on behalf of a requester, such as changing the storage class based on an Object Lifecycle Configuration.
        """
        return pulumi.get(self, "updated")


class AwaitableGetBucketObjectResult(GetBucketObjectResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetBucketObjectResult(
            acl=self.acl,
            bucket=self.bucket,
            cache_control=self.cache_control,
            component_count=self.component_count,
            content_disposition=self.content_disposition,
            content_encoding=self.content_encoding,
            content_language=self.content_language,
            content_type=self.content_type,
            crc32c=self.crc32c,
            custom_time=self.custom_time,
            customer_encryption=self.customer_encryption,
            etag=self.etag,
            event_based_hold=self.event_based_hold,
            generation=self.generation,
            kind=self.kind,
            kms_key_name=self.kms_key_name,
            md5_hash=self.md5_hash,
            media_link=self.media_link,
            metadata=self.metadata,
            metageneration=self.metageneration,
            name=self.name,
            owner=self.owner,
            retention_expiration_time=self.retention_expiration_time,
            self_link=self.self_link,
            size=self.size,
            storage_class=self.storage_class,
            temporary_hold=self.temporary_hold,
            time_created=self.time_created,
            time_deleted=self.time_deleted,
            time_storage_class_updated=self.time_storage_class_updated,
            updated=self.updated)


def get_bucket_object(bucket: Optional[str] = None,
                      generation: Optional[str] = None,
                      if_generation_match: Optional[str] = None,
                      if_generation_not_match: Optional[str] = None,
                      if_metageneration_match: Optional[str] = None,
                      if_metageneration_not_match: Optional[str] = None,
                      object: Optional[str] = None,
                      projection: Optional[str] = None,
                      user_project: Optional[str] = None,
                      opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetBucketObjectResult:
    """
    Retrieves an object or its metadata.
    """
    __args__ = dict()
    __args__['bucket'] = bucket
    __args__['generation'] = generation
    __args__['ifGenerationMatch'] = if_generation_match
    __args__['ifGenerationNotMatch'] = if_generation_not_match
    __args__['ifMetagenerationMatch'] = if_metageneration_match
    __args__['ifMetagenerationNotMatch'] = if_metageneration_not_match
    __args__['object'] = object
    __args__['projection'] = projection
    __args__['userProject'] = user_project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:storage/v1:getBucketObject', __args__, opts=opts, typ=GetBucketObjectResult).value

    return AwaitableGetBucketObjectResult(
        acl=pulumi.get(__ret__, 'acl'),
        bucket=pulumi.get(__ret__, 'bucket'),
        cache_control=pulumi.get(__ret__, 'cache_control'),
        component_count=pulumi.get(__ret__, 'component_count'),
        content_disposition=pulumi.get(__ret__, 'content_disposition'),
        content_encoding=pulumi.get(__ret__, 'content_encoding'),
        content_language=pulumi.get(__ret__, 'content_language'),
        content_type=pulumi.get(__ret__, 'content_type'),
        crc32c=pulumi.get(__ret__, 'crc32c'),
        custom_time=pulumi.get(__ret__, 'custom_time'),
        customer_encryption=pulumi.get(__ret__, 'customer_encryption'),
        etag=pulumi.get(__ret__, 'etag'),
        event_based_hold=pulumi.get(__ret__, 'event_based_hold'),
        generation=pulumi.get(__ret__, 'generation'),
        kind=pulumi.get(__ret__, 'kind'),
        kms_key_name=pulumi.get(__ret__, 'kms_key_name'),
        md5_hash=pulumi.get(__ret__, 'md5_hash'),
        media_link=pulumi.get(__ret__, 'media_link'),
        metadata=pulumi.get(__ret__, 'metadata'),
        metageneration=pulumi.get(__ret__, 'metageneration'),
        name=pulumi.get(__ret__, 'name'),
        owner=pulumi.get(__ret__, 'owner'),
        retention_expiration_time=pulumi.get(__ret__, 'retention_expiration_time'),
        self_link=pulumi.get(__ret__, 'self_link'),
        size=pulumi.get(__ret__, 'size'),
        storage_class=pulumi.get(__ret__, 'storage_class'),
        temporary_hold=pulumi.get(__ret__, 'temporary_hold'),
        time_created=pulumi.get(__ret__, 'time_created'),
        time_deleted=pulumi.get(__ret__, 'time_deleted'),
        time_storage_class_updated=pulumi.get(__ret__, 'time_storage_class_updated'),
        updated=pulumi.get(__ret__, 'updated'))


@_utilities.lift_output_func(get_bucket_object)
def get_bucket_object_output(bucket: Optional[pulumi.Input[str]] = None,
                             generation: Optional[pulumi.Input[Optional[str]]] = None,
                             if_generation_match: Optional[pulumi.Input[Optional[str]]] = None,
                             if_generation_not_match: Optional[pulumi.Input[Optional[str]]] = None,
                             if_metageneration_match: Optional[pulumi.Input[Optional[str]]] = None,
                             if_metageneration_not_match: Optional[pulumi.Input[Optional[str]]] = None,
                             object: Optional[pulumi.Input[str]] = None,
                             projection: Optional[pulumi.Input[Optional[str]]] = None,
                             user_project: Optional[pulumi.Input[Optional[str]]] = None,
                             opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetBucketObjectResult]:
    """
    Retrieves an object or its metadata.
    """
    ...
