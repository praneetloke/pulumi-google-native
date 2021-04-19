// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BigQuery.V2.Inputs
{

    public sealed class TableFieldSchemaArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// [Optional] The categories attached to this field, used for field-level access control.
        /// </summary>
        [Input("categories")]
        public Input<Inputs.JobCategoriesArgs>? Categories { get; set; }

        /// <summary>
        /// [Optional] The field description. The maximum length is 1,024 characters.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("fields")]
        private InputList<Inputs.TableFieldSchemaArgs>? _fields;

        /// <summary>
        /// [Optional] Describes the nested schema fields if the type property is set to RECORD.
        /// </summary>
        public InputList<Inputs.TableFieldSchemaArgs> Fields
        {
            get => _fields ?? (_fields = new InputList<Inputs.TableFieldSchemaArgs>());
            set => _fields = value;
        }

        /// <summary>
        /// [Optional] Maximum length of values of this field for STRINGS or BYTES. If max_length is not specified, no maximum length constraint is imposed on this field. If type = "STRING", then max_length represents the maximum UTF-8 length of strings in this field. If type = "BYTES", then max_length represents the maximum number of bytes in this field. It is invalid to set this field if type ≠ "STRING" and ≠ "BYTES".
        /// </summary>
        [Input("maxLength")]
        public Input<string>? MaxLength { get; set; }

        /// <summary>
        /// [Optional] The field mode. Possible values include NULLABLE, REQUIRED and REPEATED. The default value is NULLABLE.
        /// </summary>
        [Input("mode")]
        public Input<string>? Mode { get; set; }

        /// <summary>
        /// [Required] The field name. The name must contain only letters (a-z, A-Z), numbers (0-9), or underscores (_), and must start with a letter or underscore. The maximum length is 300 characters.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("policyTags")]
        public Input<Inputs.JobPolicyTagsArgs>? PolicyTags { get; set; }

        /// <summary>
        /// [Optional] Precision (maximum number of total digits in base 10) and scale (maximum number of digits in the fractional part in base 10) constraints for values of this field for NUMERIC or BIGNUMERIC. It is invalid to set precision or scale if type ≠ "NUMERIC" and ≠ "BIGNUMERIC". If precision and scale are not specified, no value range constraint is imposed on this field insofar as values are permitted by the type. Values of this NUMERIC or BIGNUMERIC field must be in this range when: - Precision (P) and scale (S) are specified: [-10P-S + 10-S, 10P-S - 10-S] - Precision (P) is specified but not scale (and thus scale is interpreted to be equal to zero): [-10P + 1, 10P - 1]. Acceptable values for precision and scale if both are specified: - If type = "NUMERIC": 1 ≤ precision - scale ≤ 29 and 0 ≤ scale ≤ 9. - If type = "BIGNUMERIC": 1 ≤ precision - scale ≤ 38 and 0 ≤ scale ≤ 38. Acceptable values for precision if only precision is specified but not scale (and thus scale is interpreted to be equal to zero): - If type = "NUMERIC": 1 ≤ precision ≤ 29. - If type = "BIGNUMERIC": 1 ≤ precision ≤ 38. If scale is specified but not precision, then it is invalid.
        /// </summary>
        [Input("precision")]
        public Input<string>? Precision { get; set; }

        /// <summary>
        /// [Optional] See documentation for precision.
        /// </summary>
        [Input("scale")]
        public Input<string>? Scale { get; set; }

        /// <summary>
        /// [Required] The field data type. Possible values include STRING, BYTES, INTEGER, INT64 (same as INTEGER), FLOAT, FLOAT64 (same as FLOAT), NUMERIC, BIGNUMERIC, BOOLEAN, BOOL (same as BOOLEAN), TIMESTAMP, DATE, TIME, DATETIME, INTERVAL, RECORD (where RECORD indicates that the field contains a nested schema) or STRUCT (same as RECORD).
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public TableFieldSchemaArgs()
        {
        }
    }
}
