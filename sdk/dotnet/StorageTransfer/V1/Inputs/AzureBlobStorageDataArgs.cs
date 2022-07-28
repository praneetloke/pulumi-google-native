// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.StorageTransfer.V1.Inputs
{

    /// <summary>
    /// An AzureBlobStorageData resource can be a data source, but not a data sink. An AzureBlobStorageData resource represents one Azure container. The storage account determines the [Azure endpoint](https://docs.microsoft.com/en-us/azure/storage/common/storage-create-storage-account#storage-account-endpoints). In an AzureBlobStorageData resource, a blobs's name is the [Azure Blob Storage blob's key name](https://docs.microsoft.com/en-us/rest/api/storageservices/naming-and-referencing-containers--blobs--and-metadata#blob-names).
    /// </summary>
    public sealed class AzureBlobStorageDataArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Input only. Credentials used to authenticate API requests to Azure. For information on our data retention policy for user credentials, see [User credentials](/storage-transfer/docs/data-retention#user-credentials).
        /// </summary>
        [Input("azureCredentials", required: true)]
        public Input<Inputs.AzureCredentialsArgs> AzureCredentials { get; set; } = null!;

        /// <summary>
        /// The container to transfer from the Azure Storage account.
        /// </summary>
        [Input("container", required: true)]
        public Input<string> Container { get; set; } = null!;

        /// <summary>
        /// Root path to transfer objects. Must be an empty string or full path name that ends with a '/'. This field is treated as an object prefix. As such, it should generally not begin with a '/'.
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// The name of the Azure Storage account.
        /// </summary>
        [Input("storageAccount", required: true)]
        public Input<string> StorageAccount { get; set; } = null!;

        public AzureBlobStorageDataArgs()
        {
        }
        public static new AzureBlobStorageDataArgs Empty => new AzureBlobStorageDataArgs();
    }
}
