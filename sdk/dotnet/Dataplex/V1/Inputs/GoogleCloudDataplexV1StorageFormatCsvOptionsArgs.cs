// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataplex.V1.Inputs
{

    /// <summary>
    /// Describes CSV and similar semi-structured data formats.
    /// </summary>
    public sealed class GoogleCloudDataplexV1StorageFormatCsvOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. The delimiter used to separate values. Defaults to ','.
        /// </summary>
        [Input("delimiter")]
        public Input<string>? Delimiter { get; set; }

        /// <summary>
        /// Optional. The character encoding of the data. Accepts "US-ASCII", "UTF-8", and "ISO-8859-1". Defaults to UTF-8 if unspecified.
        /// </summary>
        [Input("encoding")]
        public Input<string>? Encoding { get; set; }

        /// <summary>
        /// Optional. The number of rows to interpret as header rows that should be skipped when reading data rows. Defaults to 0.
        /// </summary>
        [Input("headerRows")]
        public Input<int>? HeaderRows { get; set; }

        /// <summary>
        /// Optional. The character used to quote column values. Accepts '"' (double quotation mark) or ''' (single quotation mark). Defaults to '"' (double quotation mark) if unspecified.
        /// </summary>
        [Input("quote")]
        public Input<string>? Quote { get; set; }

        public GoogleCloudDataplexV1StorageFormatCsvOptionsArgs()
        {
        }
        public static new GoogleCloudDataplexV1StorageFormatCsvOptionsArgs Empty => new GoogleCloudDataplexV1StorageFormatCsvOptionsArgs();
    }
}
