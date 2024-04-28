// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.GKEHub.V1.Inputs
{

    /// <summary>
    /// BundleInstallSpec is the specification configuration for a single managed bundle.
    /// </summary>
    public sealed class PolicyControllerBundleInstallSpecArgs : global::Pulumi.ResourceArgs
    {
        [Input("exemptedNamespaces")]
        private InputList<string>? _exemptedNamespaces;

        /// <summary>
        /// The set of namespaces to be exempted from the bundle.
        /// </summary>
        public InputList<string> ExemptedNamespaces
        {
            get => _exemptedNamespaces ?? (_exemptedNamespaces = new InputList<string>());
            set => _exemptedNamespaces = value;
        }

        public PolicyControllerBundleInstallSpecArgs()
        {
        }
        public static new PolicyControllerBundleInstallSpecArgs Empty => new PolicyControllerBundleInstallSpecArgs();
    }
}
