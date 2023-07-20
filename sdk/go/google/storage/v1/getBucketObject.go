// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Retrieves an object or its metadata.
func LookupBucketObject(ctx *pulumi.Context, args *LookupBucketObjectArgs, opts ...pulumi.InvokeOption) (*LookupBucketObjectResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupBucketObjectResult
	err := ctx.Invoke("google-native:storage/v1:getBucketObject", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBucketObjectArgs struct {
	Bucket                   string  `pulumi:"bucket"`
	Generation               *string `pulumi:"generation"`
	IfGenerationMatch        *string `pulumi:"ifGenerationMatch"`
	IfGenerationNotMatch     *string `pulumi:"ifGenerationNotMatch"`
	IfMetagenerationMatch    *string `pulumi:"ifMetagenerationMatch"`
	IfMetagenerationNotMatch *string `pulumi:"ifMetagenerationNotMatch"`
	Object                   string  `pulumi:"object"`
	Projection               *string `pulumi:"projection"`
	UserProject              *string `pulumi:"userProject"`
}

type LookupBucketObjectResult struct {
	// Access controls on the object.
	Acl []ObjectAccessControlResponse `pulumi:"acl"`
	// The name of the bucket containing this object.
	Bucket string `pulumi:"bucket"`
	// Cache-Control directive for the object data. If omitted, and the object is accessible to all anonymous users, the default will be public, max-age=3600.
	CacheControl string `pulumi:"cacheControl"`
	// Number of underlying components that make up this object. Components are accumulated by compose operations.
	ComponentCount int `pulumi:"componentCount"`
	// Content-Disposition of the object data.
	ContentDisposition string `pulumi:"contentDisposition"`
	// Content-Encoding of the object data.
	ContentEncoding string `pulumi:"contentEncoding"`
	// Content-Language of the object data.
	ContentLanguage string `pulumi:"contentLanguage"`
	// Content-Type of the object data. If an object is stored without a Content-Type, it is served as application/octet-stream.
	ContentType string `pulumi:"contentType"`
	// CRC32c checksum, as described in RFC 4960, Appendix B; encoded using base64 in big-endian byte order. For more information about using the CRC32c checksum, see Hashes and ETags: Best Practices.
	Crc32c string `pulumi:"crc32c"`
	// A timestamp in RFC 3339 format specified by the user for an object.
	CustomTime string `pulumi:"customTime"`
	// Metadata of customer-supplied encryption key, if the object is encrypted by such a key.
	CustomerEncryption BucketObjectCustomerEncryptionResponse `pulumi:"customerEncryption"`
	// HTTP 1.1 Entity tag for the object.
	Etag string `pulumi:"etag"`
	// Whether an object is under event-based hold. Event-based hold is a way to retain objects until an event occurs, which is signified by the hold's release (i.e. this value is set to false). After being released (set to false), such objects will be subject to bucket-level retention (if any). One sample use case of this flag is for banks to hold loan documents for at least 3 years after loan is paid in full. Here, bucket-level retention is 3 years and the event is the loan being paid in full. In this example, these objects will be held intact for any number of years until the event has occurred (event-based hold on the object is released) and then 3 more years after that. That means retention duration of the objects begins from the moment event-based hold transitioned from true to false.
	EventBasedHold bool `pulumi:"eventBasedHold"`
	// The content generation of this object. Used for object versioning.
	Generation string `pulumi:"generation"`
	// The kind of item this is. For objects, this is always storage#object.
	Kind string `pulumi:"kind"`
	// Not currently supported. Specifying the parameter causes the request to fail with status code 400 - Bad Request.
	KmsKeyName string `pulumi:"kmsKeyName"`
	// MD5 hash of the data; encoded using base64. For more information about using the MD5 hash, see Hashes and ETags: Best Practices.
	Md5Hash string `pulumi:"md5Hash"`
	// Media download link.
	MediaLink string `pulumi:"mediaLink"`
	// User-provided metadata, in key/value pairs.
	Metadata map[string]string `pulumi:"metadata"`
	// The version of the metadata for this object at this generation. Used for preconditions and for detecting changes in metadata. A metageneration number is only meaningful in the context of a particular generation of a particular object.
	Metageneration string `pulumi:"metageneration"`
	// The name of the object. Required if not specified by URL parameter.
	Name string `pulumi:"name"`
	// The owner of the object. This will always be the uploader of the object.
	Owner BucketObjectOwnerResponse `pulumi:"owner"`
	// A server-determined value that specifies the earliest time that the object's retention period expires. This value is in RFC 3339 format. Note 1: This field is not provided for objects with an active event-based hold, since retention expiration is unknown until the hold is removed. Note 2: This value can be provided even when temporary hold is set (so that the user can reason about policy without having to first unset the temporary hold).
	RetentionExpirationTime string `pulumi:"retentionExpirationTime"`
	// The link to this object.
	SelfLink string `pulumi:"selfLink"`
	// Content-Length of the data in bytes.
	Size string `pulumi:"size"`
	// Storage class of the object.
	StorageClass string `pulumi:"storageClass"`
	// Whether an object is under temporary hold. While this flag is set to true, the object is protected against deletion and overwrites. A common use case of this flag is regulatory investigations where objects need to be retained while the investigation is ongoing. Note that unlike event-based hold, temporary hold does not impact retention expiration time of an object.
	TemporaryHold bool `pulumi:"temporaryHold"`
	// The creation time of the object in RFC 3339 format.
	TimeCreated string `pulumi:"timeCreated"`
	// The deletion time of the object in RFC 3339 format. Will be returned if and only if this version of the object has been deleted.
	TimeDeleted string `pulumi:"timeDeleted"`
	// The time at which the object's storage class was last changed. When the object is initially created, it will be set to timeCreated.
	TimeStorageClassUpdated string `pulumi:"timeStorageClassUpdated"`
	// The modification time of the object metadata in RFC 3339 format. Set initially to object creation time and then updated whenever any metadata of the object changes. This includes changes made by a requester, such as modifying custom metadata, as well as changes made by Cloud Storage on behalf of a requester, such as changing the storage class based on an Object Lifecycle Configuration.
	Updated string `pulumi:"updated"`
}

func LookupBucketObjectOutput(ctx *pulumi.Context, args LookupBucketObjectOutputArgs, opts ...pulumi.InvokeOption) LookupBucketObjectResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBucketObjectResult, error) {
			args := v.(LookupBucketObjectArgs)
			r, err := LookupBucketObject(ctx, &args, opts...)
			var s LookupBucketObjectResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupBucketObjectResultOutput)
}

type LookupBucketObjectOutputArgs struct {
	Bucket                   pulumi.StringInput    `pulumi:"bucket"`
	Generation               pulumi.StringPtrInput `pulumi:"generation"`
	IfGenerationMatch        pulumi.StringPtrInput `pulumi:"ifGenerationMatch"`
	IfGenerationNotMatch     pulumi.StringPtrInput `pulumi:"ifGenerationNotMatch"`
	IfMetagenerationMatch    pulumi.StringPtrInput `pulumi:"ifMetagenerationMatch"`
	IfMetagenerationNotMatch pulumi.StringPtrInput `pulumi:"ifMetagenerationNotMatch"`
	Object                   pulumi.StringInput    `pulumi:"object"`
	Projection               pulumi.StringPtrInput `pulumi:"projection"`
	UserProject              pulumi.StringPtrInput `pulumi:"userProject"`
}

func (LookupBucketObjectOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBucketObjectArgs)(nil)).Elem()
}

type LookupBucketObjectResultOutput struct{ *pulumi.OutputState }

func (LookupBucketObjectResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBucketObjectResult)(nil)).Elem()
}

func (o LookupBucketObjectResultOutput) ToLookupBucketObjectResultOutput() LookupBucketObjectResultOutput {
	return o
}

func (o LookupBucketObjectResultOutput) ToLookupBucketObjectResultOutputWithContext(ctx context.Context) LookupBucketObjectResultOutput {
	return o
}

// Access controls on the object.
func (o LookupBucketObjectResultOutput) Acl() ObjectAccessControlResponseArrayOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) []ObjectAccessControlResponse { return v.Acl }).(ObjectAccessControlResponseArrayOutput)
}

// The name of the bucket containing this object.
func (o LookupBucketObjectResultOutput) Bucket() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.Bucket }).(pulumi.StringOutput)
}

// Cache-Control directive for the object data. If omitted, and the object is accessible to all anonymous users, the default will be public, max-age=3600.
func (o LookupBucketObjectResultOutput) CacheControl() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.CacheControl }).(pulumi.StringOutput)
}

// Number of underlying components that make up this object. Components are accumulated by compose operations.
func (o LookupBucketObjectResultOutput) ComponentCount() pulumi.IntOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) int { return v.ComponentCount }).(pulumi.IntOutput)
}

// Content-Disposition of the object data.
func (o LookupBucketObjectResultOutput) ContentDisposition() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.ContentDisposition }).(pulumi.StringOutput)
}

// Content-Encoding of the object data.
func (o LookupBucketObjectResultOutput) ContentEncoding() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.ContentEncoding }).(pulumi.StringOutput)
}

// Content-Language of the object data.
func (o LookupBucketObjectResultOutput) ContentLanguage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.ContentLanguage }).(pulumi.StringOutput)
}

// Content-Type of the object data. If an object is stored without a Content-Type, it is served as application/octet-stream.
func (o LookupBucketObjectResultOutput) ContentType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.ContentType }).(pulumi.StringOutput)
}

// CRC32c checksum, as described in RFC 4960, Appendix B; encoded using base64 in big-endian byte order. For more information about using the CRC32c checksum, see Hashes and ETags: Best Practices.
func (o LookupBucketObjectResultOutput) Crc32c() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.Crc32c }).(pulumi.StringOutput)
}

// A timestamp in RFC 3339 format specified by the user for an object.
func (o LookupBucketObjectResultOutput) CustomTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.CustomTime }).(pulumi.StringOutput)
}

// Metadata of customer-supplied encryption key, if the object is encrypted by such a key.
func (o LookupBucketObjectResultOutput) CustomerEncryption() BucketObjectCustomerEncryptionResponseOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) BucketObjectCustomerEncryptionResponse { return v.CustomerEncryption }).(BucketObjectCustomerEncryptionResponseOutput)
}

// HTTP 1.1 Entity tag for the object.
func (o LookupBucketObjectResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.Etag }).(pulumi.StringOutput)
}

// Whether an object is under event-based hold. Event-based hold is a way to retain objects until an event occurs, which is signified by the hold's release (i.e. this value is set to false). After being released (set to false), such objects will be subject to bucket-level retention (if any). One sample use case of this flag is for banks to hold loan documents for at least 3 years after loan is paid in full. Here, bucket-level retention is 3 years and the event is the loan being paid in full. In this example, these objects will be held intact for any number of years until the event has occurred (event-based hold on the object is released) and then 3 more years after that. That means retention duration of the objects begins from the moment event-based hold transitioned from true to false.
func (o LookupBucketObjectResultOutput) EventBasedHold() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) bool { return v.EventBasedHold }).(pulumi.BoolOutput)
}

// The content generation of this object. Used for object versioning.
func (o LookupBucketObjectResultOutput) Generation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.Generation }).(pulumi.StringOutput)
}

// The kind of item this is. For objects, this is always storage#object.
func (o LookupBucketObjectResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.Kind }).(pulumi.StringOutput)
}

// Not currently supported. Specifying the parameter causes the request to fail with status code 400 - Bad Request.
func (o LookupBucketObjectResultOutput) KmsKeyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.KmsKeyName }).(pulumi.StringOutput)
}

// MD5 hash of the data; encoded using base64. For more information about using the MD5 hash, see Hashes and ETags: Best Practices.
func (o LookupBucketObjectResultOutput) Md5Hash() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.Md5Hash }).(pulumi.StringOutput)
}

// Media download link.
func (o LookupBucketObjectResultOutput) MediaLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.MediaLink }).(pulumi.StringOutput)
}

// User-provided metadata, in key/value pairs.
func (o LookupBucketObjectResultOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

// The version of the metadata for this object at this generation. Used for preconditions and for detecting changes in metadata. A metageneration number is only meaningful in the context of a particular generation of a particular object.
func (o LookupBucketObjectResultOutput) Metageneration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.Metageneration }).(pulumi.StringOutput)
}

// The name of the object. Required if not specified by URL parameter.
func (o LookupBucketObjectResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.Name }).(pulumi.StringOutput)
}

// The owner of the object. This will always be the uploader of the object.
func (o LookupBucketObjectResultOutput) Owner() BucketObjectOwnerResponseOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) BucketObjectOwnerResponse { return v.Owner }).(BucketObjectOwnerResponseOutput)
}

// A server-determined value that specifies the earliest time that the object's retention period expires. This value is in RFC 3339 format. Note 1: This field is not provided for objects with an active event-based hold, since retention expiration is unknown until the hold is removed. Note 2: This value can be provided even when temporary hold is set (so that the user can reason about policy without having to first unset the temporary hold).
func (o LookupBucketObjectResultOutput) RetentionExpirationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.RetentionExpirationTime }).(pulumi.StringOutput)
}

// The link to this object.
func (o LookupBucketObjectResultOutput) SelfLink() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.SelfLink }).(pulumi.StringOutput)
}

// Content-Length of the data in bytes.
func (o LookupBucketObjectResultOutput) Size() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.Size }).(pulumi.StringOutput)
}

// Storage class of the object.
func (o LookupBucketObjectResultOutput) StorageClass() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.StorageClass }).(pulumi.StringOutput)
}

// Whether an object is under temporary hold. While this flag is set to true, the object is protected against deletion and overwrites. A common use case of this flag is regulatory investigations where objects need to be retained while the investigation is ongoing. Note that unlike event-based hold, temporary hold does not impact retention expiration time of an object.
func (o LookupBucketObjectResultOutput) TemporaryHold() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) bool { return v.TemporaryHold }).(pulumi.BoolOutput)
}

// The creation time of the object in RFC 3339 format.
func (o LookupBucketObjectResultOutput) TimeCreated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.TimeCreated }).(pulumi.StringOutput)
}

// The deletion time of the object in RFC 3339 format. Will be returned if and only if this version of the object has been deleted.
func (o LookupBucketObjectResultOutput) TimeDeleted() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.TimeDeleted }).(pulumi.StringOutput)
}

// The time at which the object's storage class was last changed. When the object is initially created, it will be set to timeCreated.
func (o LookupBucketObjectResultOutput) TimeStorageClassUpdated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.TimeStorageClassUpdated }).(pulumi.StringOutput)
}

// The modification time of the object metadata in RFC 3339 format. Set initially to object creation time and then updated whenever any metadata of the object changes. This includes changes made by a requester, such as modifying custom metadata, as well as changes made by Cloud Storage on behalf of a requester, such as changing the storage class based on an Object Lifecycle Configuration.
func (o LookupBucketObjectResultOutput) Updated() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBucketObjectResult) string { return v.Updated }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBucketObjectResultOutput{})
}
