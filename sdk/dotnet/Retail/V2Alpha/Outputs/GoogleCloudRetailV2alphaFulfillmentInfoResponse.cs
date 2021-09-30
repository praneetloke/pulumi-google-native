// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Retail.V2Alpha.Outputs
{

    /// <summary>
    /// Fulfillment information, such as the store IDs for in-store pickup or region IDs for different shipping methods.
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudRetailV2alphaFulfillmentInfoResponse
    {
        /// <summary>
        /// The IDs for this type, such as the store IDs for FulfillmentInfo.type.pickup-in-store or the region IDs for FulfillmentInfo.type.same-day-delivery. A maximum of 2000 values are allowed. Each value must be a string with a length limit of 10 characters, matching the pattern [a-zA-Z0-9_-]+, such as "store1" or "REGION-2". Otherwise, an INVALID_ARGUMENT error is returned.
        /// </summary>
        public readonly ImmutableArray<string> PlaceIds;
        /// <summary>
        /// The fulfillment type, including commonly used types (such as pickup in store and same day delivery), and custom types. Customers have to map custom types to their display names before rendering UI. Supported values: * "pickup-in-store" * "ship-to-store" * "same-day-delivery" * "next-day-delivery" * "custom-type-1" * "custom-type-2" * "custom-type-3" * "custom-type-4" * "custom-type-5" If this field is set to an invalid value other than these, an INVALID_ARGUMENT error is returned.
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private GoogleCloudRetailV2alphaFulfillmentInfoResponse(
            ImmutableArray<string> placeIds,

            string type)
        {
            PlaceIds = placeIds;
            Type = type;
        }
    }
}