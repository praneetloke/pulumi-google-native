// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.TPU.V2Alpha1.Inputs
{

    /// <summary>
    /// A service account.
    /// </summary>
    public sealed class ServiceAccountArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Email address of the service account. If empty, default Compute service account will be used.
        /// </summary>
        [Input("email")]
        public Input<string>? Email { get; set; }

        [Input("scope")]
        private InputList<string>? _scope;

        /// <summary>
        /// The list of scopes to be made available for this service account. If empty, access to all Cloud APIs will be allowed.
        /// </summary>
        public InputList<string> Scope
        {
            get => _scope ?? (_scope = new InputList<string>());
            set => _scope = value;
        }

        public ServiceAccountArgs()
        {
        }
        public static new ServiceAccountArgs Empty => new ServiceAccountArgs();
    }
}
