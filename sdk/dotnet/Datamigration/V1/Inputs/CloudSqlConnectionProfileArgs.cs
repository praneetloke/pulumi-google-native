// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datamigration.V1.Inputs
{

    /// <summary>
    /// Specifies required connection parameters, and, optionally, the parameters required to create a Cloud SQL destination database instance.
    /// </summary>
    public sealed class CloudSqlConnectionProfileArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Immutable. Metadata used to create the destination Cloud SQL database.
        /// </summary>
        [Input("settings")]
        public Input<Inputs.CloudSqlSettingsArgs>? Settings { get; set; }

        public CloudSqlConnectionProfileArgs()
        {
        }
        public static new CloudSqlConnectionProfileArgs Empty => new CloudSqlConnectionProfileArgs();
    }
}
