// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DataCatalog.V1.Inputs
{

    public sealed class GoogleCloudDatacatalogV1FieldTypeEnumTypeArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowedValues")]
        private InputList<Inputs.GoogleCloudDatacatalogV1FieldTypeEnumTypeEnumValueArgs>? _allowedValues;

        /// <summary>
        /// The set of allowed values for this enum. This set must not be empty and can include up to 100 allowed values. The display names of the values in this set must not be empty and must be case-insensitively unique within this set. The order of items in this set is preserved. This field can be used to create, remove, and reorder enum values. To rename enum values, use the `RenameTagTemplateFieldEnumValue` method.
        /// </summary>
        public InputList<Inputs.GoogleCloudDatacatalogV1FieldTypeEnumTypeEnumValueArgs> AllowedValues
        {
            get => _allowedValues ?? (_allowedValues = new InputList<Inputs.GoogleCloudDatacatalogV1FieldTypeEnumTypeEnumValueArgs>());
            set => _allowedValues = value;
        }

        public GoogleCloudDatacatalogV1FieldTypeEnumTypeArgs()
        {
        }
        public static new GoogleCloudDatacatalogV1FieldTypeEnumTypeArgs Empty => new GoogleCloudDatacatalogV1FieldTypeEnumTypeArgs();
    }
}