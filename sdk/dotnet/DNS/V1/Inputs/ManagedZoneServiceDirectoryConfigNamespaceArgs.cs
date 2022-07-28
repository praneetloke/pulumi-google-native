// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DNS.V1.Inputs
{

    public sealed class ManagedZoneServiceDirectoryConfigNamespaceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The time that the namespace backing this zone was deleted; an empty string if it still exists. This is in RFC3339 text format. Output only.
        /// </summary>
        [Input("deletionTime")]
        public Input<string>? DeletionTime { get; set; }

        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// The fully qualified URL of the namespace associated with the zone. Format must be https://servicedirectory.googleapis.com/v1/projects/{project}/locations/{location}/namespaces/{namespace}
        /// </summary>
        [Input("namespaceUrl")]
        public Input<string>? NamespaceUrl { get; set; }

        public ManagedZoneServiceDirectoryConfigNamespaceArgs()
        {
        }
        public static new ManagedZoneServiceDirectoryConfigNamespaceArgs Empty => new ManagedZoneServiceDirectoryConfigNamespaceArgs();
    }
}
