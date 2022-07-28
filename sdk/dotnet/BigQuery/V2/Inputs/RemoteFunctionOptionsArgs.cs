// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BigQuery.V2.Inputs
{

    /// <summary>
    /// Options for a remote user-defined function.
    /// </summary>
    public sealed class RemoteFunctionOptionsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Fully qualified name of the user-provided connection object which holds the authentication information to send requests to the remote service. projects/{project_id}/locations/{location_id}/connections/{connection_id}
        /// </summary>
        [Input("connection")]
        public Input<string>? Connection { get; set; }

        /// <summary>
        /// Endpoint of the user-provided remote service (e.g. a function url in Google Cloud Functions).
        /// </summary>
        [Input("endpoint")]
        public Input<string>? Endpoint { get; set; }

        /// <summary>
        /// Max number of rows in each batch sent to the remote service. If absent or if 0, it means no limit.
        /// </summary>
        [Input("maxBatchingRows")]
        public Input<string>? MaxBatchingRows { get; set; }

        [Input("userDefinedContext")]
        private InputMap<string>? _userDefinedContext;

        /// <summary>
        /// User-defined context as a set of key/value pairs, which will be sent as function invocation context together with batched arguments in the requests to the remote service. The total number of bytes of keys and values must be less than 8KB.
        /// </summary>
        public InputMap<string> UserDefinedContext
        {
            get => _userDefinedContext ?? (_userDefinedContext = new InputMap<string>());
            set => _userDefinedContext = value;
        }

        public RemoteFunctionOptionsArgs()
        {
        }
        public static new RemoteFunctionOptionsArgs Empty => new RemoteFunctionOptionsArgs();
    }
}
