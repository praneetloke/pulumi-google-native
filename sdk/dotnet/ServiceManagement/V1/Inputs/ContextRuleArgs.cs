// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ServiceManagement.V1.Inputs
{

    /// <summary>
    /// A context rule provides information about the context for an individual API element.
    /// </summary>
    public sealed class ContextRuleArgs : global::Pulumi.ResourceArgs
    {
        [Input("allowedRequestExtensions")]
        private InputList<string>? _allowedRequestExtensions;

        /// <summary>
        /// A list of full type names or extension IDs of extensions allowed in grpc side channel from client to backend.
        /// </summary>
        public InputList<string> AllowedRequestExtensions
        {
            get => _allowedRequestExtensions ?? (_allowedRequestExtensions = new InputList<string>());
            set => _allowedRequestExtensions = value;
        }

        [Input("allowedResponseExtensions")]
        private InputList<string>? _allowedResponseExtensions;

        /// <summary>
        /// A list of full type names or extension IDs of extensions allowed in grpc side channel from backend to client.
        /// </summary>
        public InputList<string> AllowedResponseExtensions
        {
            get => _allowedResponseExtensions ?? (_allowedResponseExtensions = new InputList<string>());
            set => _allowedResponseExtensions = value;
        }

        [Input("provided")]
        private InputList<string>? _provided;

        /// <summary>
        /// A list of full type names of provided contexts.
        /// </summary>
        public InputList<string> Provided
        {
            get => _provided ?? (_provided = new InputList<string>());
            set => _provided = value;
        }

        [Input("requested")]
        private InputList<string>? _requested;

        /// <summary>
        /// A list of full type names of requested contexts.
        /// </summary>
        public InputList<string> Requested
        {
            get => _requested ?? (_requested = new InputList<string>());
            set => _requested = value;
        }

        /// <summary>
        /// Selects the methods to which this rule applies. Refer to selector for syntax details.
        /// </summary>
        [Input("selector")]
        public Input<string>? Selector { get; set; }

        public ContextRuleArgs()
        {
        }
        public static new ContextRuleArgs Empty => new ContextRuleArgs();
    }
}
