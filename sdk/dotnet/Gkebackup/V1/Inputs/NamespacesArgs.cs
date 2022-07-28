// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Gkebackup.V1.Inputs
{

    /// <summary>
    /// A list of Kubernetes Namespaces
    /// </summary>
    public sealed class NamespacesArgs : global::Pulumi.ResourceArgs
    {
        [Input("namespaces")]
        private InputList<string>? _namespaces;

        /// <summary>
        /// A list of Kubernetes Namespaces
        /// </summary>
        public InputList<string> Namespaces
        {
            get => _namespaces ?? (_namespaces = new InputList<string>());
            set => _namespaces = value;
        }

        public NamespacesArgs()
        {
        }
        public static new NamespacesArgs Empty => new NamespacesArgs();
    }
}
