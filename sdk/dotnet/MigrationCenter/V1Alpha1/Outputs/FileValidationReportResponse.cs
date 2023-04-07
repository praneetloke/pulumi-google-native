// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.MigrationCenter.V1Alpha1.Outputs
{

    /// <summary>
    /// A resource that aggregates the validation errors found in an import job file.
    /// </summary>
    [OutputType]
    public sealed class FileValidationReportResponse
    {
        /// <summary>
        /// List of file level errors.
        /// </summary>
        public readonly ImmutableArray<Outputs.ImportErrorResponse> FileErrors;
        /// <summary>
        /// The name of the file.
        /// </summary>
        public readonly string FileName;
        /// <summary>
        /// Flag indicating that processing was aborted due to maximum number of errors.
        /// </summary>
        public readonly bool PartialReport;
        /// <summary>
        /// Partial list of rows that encountered validation error.
        /// </summary>
        public readonly ImmutableArray<Outputs.ImportRowErrorResponse> RowErrors;

        [OutputConstructor]
        private FileValidationReportResponse(
            ImmutableArray<Outputs.ImportErrorResponse> fileErrors,

            string fileName,

            bool partialReport,

            ImmutableArray<Outputs.ImportRowErrorResponse> rowErrors)
        {
            FileErrors = fileErrors;
            FileName = fileName;
            PartialReport = partialReport;
            RowErrors = rowErrors;
        }
    }
}
