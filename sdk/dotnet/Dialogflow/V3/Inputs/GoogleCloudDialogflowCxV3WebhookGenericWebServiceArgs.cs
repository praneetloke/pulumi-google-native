// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dialogflow.V3.Inputs
{

    /// <summary>
    /// Represents configuration for a generic web service.
    /// </summary>
    public sealed class GoogleCloudDialogflowCxV3WebhookGenericWebServiceArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowedCaCerts")]
        private InputList<string>? _allowedCaCerts;

        /// <summary>
        /// Optional. Specifies a list of allowed custom CA certificates (in DER format) for HTTPS verification. This overrides the default SSL trust store. If this is empty or unspecified, Dialogflow will use Google's default trust store to verify certificates. N.B. Make sure the HTTPS server certificates are signed with "subject alt name". For instance a certificate can be self-signed using the following command, ``` openssl x509 -req -days 200 -in example.com.csr \ -signkey example.com.key \ -out example.com.crt \ -extfile &lt;(printf "\nsubjectAltName='DNS:www.example.com'") ```
        /// </summary>
        public InputList<string> AllowedCaCerts
        {
            get => _allowedCaCerts ?? (_allowedCaCerts = new InputList<string>());
            set => _allowedCaCerts = value;
        }

        /// <summary>
        /// The password for HTTP Basic authentication.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        [Input("requestHeaders")]
        private InputMap<string>? _requestHeaders;

        /// <summary>
        /// The HTTP request headers to send together with webhook requests.
        /// </summary>
        public InputMap<string> RequestHeaders
        {
            get => _requestHeaders ?? (_requestHeaders = new InputMap<string>());
            set => _requestHeaders = value;
        }

        /// <summary>
        /// The webhook URI for receiving POST requests. It must use https protocol.
        /// </summary>
        [Input("uri", required: true)]
        public Input<string> Uri { get; set; } = null!;

        /// <summary>
        /// The user name for HTTP Basic authentication.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public GoogleCloudDialogflowCxV3WebhookGenericWebServiceArgs()
        {
        }
        public static new GoogleCloudDialogflowCxV3WebhookGenericWebServiceArgs Empty => new GoogleCloudDialogflowCxV3WebhookGenericWebServiceArgs();
    }
}
