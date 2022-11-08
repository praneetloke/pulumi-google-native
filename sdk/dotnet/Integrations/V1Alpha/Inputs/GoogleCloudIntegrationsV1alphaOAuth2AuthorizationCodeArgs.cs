// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Integrations.V1Alpha.Inputs
{

    /// <summary>
    /// The OAuth Type where the client sends request with the client id and requested scopes to auth endpoint. User sees a consent screen and auth code is received at specified redirect url afterwards. The auth code is then combined with the client id and secret and sent to the token endpoint in exchange for the access and refresh token. The refresh token can be used to fetch new access tokens.
    /// </summary>
    public sealed class GoogleCloudIntegrationsV1alphaOAuth2AuthorizationCodeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The access token received from the token endpoint.
        /// </summary>
        [Input("accessToken")]
        public Input<Inputs.GoogleCloudIntegrationsV1alphaAccessTokenArgs>? AccessToken { get; set; }

        /// <summary>
        /// Indicates if the user has opted in Google Reauth Policy. If opted in, the refresh token will be valid for 20 hours, after which time users must re-authenticate in order to obtain a new one.
        /// </summary>
        [Input("applyReauthPolicy")]
        public Input<bool>? ApplyReauthPolicy { get; set; }

        /// <summary>
        /// The Auth Code that is used to initially retrieve the access token.
        /// </summary>
        [Input("authCode")]
        public Input<string>? AuthCode { get; set; }

        /// <summary>
        /// The auth url endpoint to send the auth code request to.
        /// </summary>
        [Input("authEndpoint")]
        public Input<string>? AuthEndpoint { get; set; }

        /// <summary>
        /// The auth parameters sent along with the auth code request.
        /// </summary>
        [Input("authParams")]
        public Input<Inputs.GoogleCloudIntegrationsV1alphaParameterMapArgs>? AuthParams { get; set; }

        /// <summary>
        /// The client's id.
        /// </summary>
        [Input("clientId")]
        public Input<string>? ClientId { get; set; }

        /// <summary>
        /// The client's secret.
        /// </summary>
        [Input("clientSecret")]
        public Input<string>? ClientSecret { get; set; }

        /// <summary>
        /// Represent how to pass parameters to fetch access token
        /// </summary>
        [Input("requestType")]
        public Input<Pulumi.GoogleNative.Integrations.V1Alpha.GoogleCloudIntegrationsV1alphaOAuth2AuthorizationCodeRequestType>? RequestType { get; set; }

        /// <summary>
        /// A space-delimited list of requested scope permissions.
        /// </summary>
        [Input("scope")]
        public Input<string>? Scope { get; set; }

        /// <summary>
        /// The token url endpoint to send the token request to.
        /// </summary>
        [Input("tokenEndpoint")]
        public Input<string>? TokenEndpoint { get; set; }

        /// <summary>
        /// The token parameters sent along with the token request.
        /// </summary>
        [Input("tokenParams")]
        public Input<Inputs.GoogleCloudIntegrationsV1alphaParameterMapArgs>? TokenParams { get; set; }

        public GoogleCloudIntegrationsV1alphaOAuth2AuthorizationCodeArgs()
        {
        }
        public static new GoogleCloudIntegrationsV1alphaOAuth2AuthorizationCodeArgs Empty => new GoogleCloudIntegrationsV1alphaOAuth2AuthorizationCodeArgs();
    }
}
