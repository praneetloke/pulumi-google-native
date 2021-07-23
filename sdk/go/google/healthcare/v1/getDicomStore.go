// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified DICOM store.
func LookupDicomStore(ctx *pulumi.Context, args *LookupDicomStoreArgs, opts ...pulumi.InvokeOption) (*LookupDicomStoreResult, error) {
	var rv LookupDicomStoreResult
	err := ctx.Invoke("google-native:healthcare/v1:getDicomStore", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDicomStoreArgs struct {
	DatasetId    string  `pulumi:"datasetId"`
	DicomStoreId string  `pulumi:"dicomStoreId"`
	Location     string  `pulumi:"location"`
	Project      *string `pulumi:"project"`
}

type LookupDicomStoreResult struct {
	// User-supplied key-value pairs used to organize DICOM stores. Label keys must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: \p{Ll}\p{Lo}{0,62} Label values are optional, must be between 1 and 63 characters long, have a UTF-8 encoding of maximum 128 bytes, and must conform to the following PCRE regular expression: [\p{Ll}\p{Lo}\p{N}_-]{0,63} No more than 64 labels can be associated with a given store.
	Labels map[string]string `pulumi:"labels"`
	// Resource name of the DICOM store, of the form `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}/dicomStores/{dicom_store_id}`.
	Name string `pulumi:"name"`
	// Notification destination for new DICOM instances. Supplied by the client.
	NotificationConfig NotificationConfigResponse `pulumi:"notificationConfig"`
}