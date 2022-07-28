// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1.Inputs
{

    /// <summary>
    /// Identity related configuration, including service account based secure multi-tenancy user mappings.
    /// </summary>
    public sealed class IdentityConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("userServiceAccountMapping", required: true)]
        private InputMap<string>? _userServiceAccountMapping;

        /// <summary>
        /// Map of user to service account.
        /// </summary>
        public InputMap<string> UserServiceAccountMapping
        {
            get => _userServiceAccountMapping ?? (_userServiceAccountMapping = new InputMap<string>());
            set => _userServiceAccountMapping = value;
        }

        public IdentityConfigArgs()
        {
        }
        public static new IdentityConfigArgs Empty => new IdentityConfigArgs();
    }
}
