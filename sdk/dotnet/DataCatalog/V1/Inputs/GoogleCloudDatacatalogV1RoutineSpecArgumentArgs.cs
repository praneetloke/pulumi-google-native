// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DataCatalog.V1.Inputs
{

    /// <summary>
    /// Input or output argument of a function or stored procedure.
    /// </summary>
    public sealed class GoogleCloudDatacatalogV1RoutineSpecArgumentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Specifies whether the argument is input or output.
        /// </summary>
        [Input("mode")]
        public Input<Pulumi.GoogleNative.DataCatalog.V1.GoogleCloudDatacatalogV1RoutineSpecArgumentMode>? Mode { get; set; }

        /// <summary>
        /// The name of the argument. A return argument of a function might not have a name.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// Type of the argument. The exact value depends on the source system and the language.
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public GoogleCloudDatacatalogV1RoutineSpecArgumentArgs()
        {
        }
        public static new GoogleCloudDatacatalogV1RoutineSpecArgumentArgs Empty => new GoogleCloudDatacatalogV1RoutineSpecArgumentArgs();
    }
}
