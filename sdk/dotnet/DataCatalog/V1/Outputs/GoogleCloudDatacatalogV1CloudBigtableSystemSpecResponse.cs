// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DataCatalog.V1.Outputs
{

    /// <summary>
    /// Specification that applies to all entries that are part of `CLOUD_BIGTABLE` system (user_specified_type)
    /// </summary>
    [OutputType]
    public sealed class GoogleCloudDatacatalogV1CloudBigtableSystemSpecResponse
    {
        /// <summary>
        /// Display name of the Instance. This is user specified and different from the resource name.
        /// </summary>
        public readonly string InstanceDisplayName;

        [OutputConstructor]
        private GoogleCloudDatacatalogV1CloudBigtableSystemSpecResponse(string instanceDisplayName)
        {
            InstanceDisplayName = instanceDisplayName;
        }
    }
}
