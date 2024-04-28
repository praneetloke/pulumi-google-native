// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Metastore.V1Beta.Inputs
{

    /// <summary>
    /// Represents a Lake resource
    /// </summary>
    public sealed class LakeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The Lake resource name. Example: projects/{project_number}/locations/{location_id}/lakes/{lake_id}
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public LakeArgs()
        {
        }
        public static new LakeArgs Empty => new LakeArgs();
    }
}
