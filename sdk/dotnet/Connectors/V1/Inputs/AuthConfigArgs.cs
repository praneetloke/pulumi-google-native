// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Connectors.V1.Inputs
{

    /// <summary>
    /// AuthConfig defines details of a authentication type.
    /// </summary>
    public sealed class AuthConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("additionalVariables")]
        private InputList<Inputs.ConfigVariableArgs>? _additionalVariables;

        /// <summary>
        /// List containing additional auth configs.
        /// </summary>
        public InputList<Inputs.ConfigVariableArgs> AdditionalVariables
        {
            get => _additionalVariables ?? (_additionalVariables = new InputList<Inputs.ConfigVariableArgs>());
            set => _additionalVariables = value;
        }

        /// <summary>
        /// The type of authentication configured.
        /// </summary>
        [Input("authType")]
        public Input<Pulumi.GoogleNative.Connectors.V1.AuthConfigAuthType>? AuthType { get; set; }

        /// <summary>
        /// Oauth2ClientCredentials.
        /// </summary>
        [Input("oauth2ClientCredentials")]
        public Input<Inputs.Oauth2ClientCredentialsArgs>? Oauth2ClientCredentials { get; set; }

        /// <summary>
        /// Oauth2JwtBearer.
        /// </summary>
        [Input("oauth2JwtBearer")]
        public Input<Inputs.Oauth2JwtBearerArgs>? Oauth2JwtBearer { get; set; }

        /// <summary>
        /// SSH Public Key.
        /// </summary>
        [Input("sshPublicKey")]
        public Input<Inputs.SshPublicKeyArgs>? SshPublicKey { get; set; }

        /// <summary>
        /// UserPassword.
        /// </summary>
        [Input("userPassword")]
        public Input<Inputs.UserPasswordArgs>? UserPassword { get; set; }

        public AuthConfigArgs()
        {
        }
        public static new AuthConfigArgs Empty => new AuthConfigArgs();
    }
}
