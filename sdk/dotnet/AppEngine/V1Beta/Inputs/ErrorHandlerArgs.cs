// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AppEngine.V1Beta.Inputs
{

    /// <summary>
    /// Custom static error page to be served when an error occurs.
    /// </summary>
    public sealed class ErrorHandlerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Error condition this handler applies to.
        /// </summary>
        [Input("errorCode")]
        public Input<Pulumi.GoogleNative.AppEngine.V1Beta.ErrorHandlerErrorCode>? ErrorCode { get; set; }

        /// <summary>
        /// MIME type of file. Defaults to text/html.
        /// </summary>
        [Input("mimeType")]
        public Input<string>? MimeType { get; set; }

        /// <summary>
        /// Static file content to be served for this error.
        /// </summary>
        [Input("staticFile")]
        public Input<string>? StaticFile { get; set; }

        public ErrorHandlerArgs()
        {
        }
        public static new ErrorHandlerArgs Empty => new ErrorHandlerArgs();
    }
}
