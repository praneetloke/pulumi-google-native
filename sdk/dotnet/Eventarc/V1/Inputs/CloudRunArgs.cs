// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Eventarc.V1.Inputs
{

    /// <summary>
    /// Represents a Cloud Run destination.
    /// </summary>
    public sealed class CloudRunArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. The relative path on the Cloud Run service the events should be sent to. The value must conform to the definition of a URI path segment (section 3.3 of RFC2396). Examples: "/route", "route", "route/subroute".
        /// </summary>
        [Input("path")]
        public Input<string>? Path { get; set; }

        /// <summary>
        /// The region the Cloud Run service is deployed in.
        /// </summary>
        [Input("region", required: true)]
        public Input<string> Region { get; set; } = null!;

        /// <summary>
        /// The name of the Cloud Run service being addressed. See https://cloud.google.com/run/docs/reference/rest/v1/namespaces.services. Only services located in the same project as the trigger object can be addressed.
        /// </summary>
        [Input("service", required: true)]
        public Input<string> Service { get; set; } = null!;

        public CloudRunArgs()
        {
        }
        public static new CloudRunArgs Empty => new CloudRunArgs();
    }
}
