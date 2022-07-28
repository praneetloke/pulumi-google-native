// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudAsset.V1.Inputs
{

    /// <summary>
    /// Specifies an identity for which to determine resource access, based on roles assigned either directly to them or to the groups they belong to, directly or indirectly.
    /// </summary>
    public sealed class IdentitySelectorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The identity appear in the form of principals in [IAM policy binding](https://cloud.google.com/iam/reference/rest/v1/Binding). The examples of supported forms are: "user:mike@example.com", "group:admins@example.com", "domain:google.com", "serviceAccount:my-project-id@appspot.gserviceaccount.com". Notice that wildcard characters (such as * and ?) are not supported. You must give a specific identity.
        /// </summary>
        [Input("identity", required: true)]
        public Input<string> Identity { get; set; } = null!;

        public IdentitySelectorArgs()
        {
        }
        public static new IdentitySelectorArgs Empty => new IdentitySelectorArgs();
    }
}
