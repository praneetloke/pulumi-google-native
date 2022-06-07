// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ContainerAnalysis.V1.Inputs
{

    /// <summary>
    /// Identifies the entity that executed the recipe, which is trusted to have correctly performed the operation and populated this provenance.
    /// </summary>
    public sealed class GrafeasV1SlsaProvenanceZeroTwoSlsaBuilderArgs : Pulumi.ResourceArgs
    {
        [Input("id")]
        public Input<string>? Id { get; set; }

        public GrafeasV1SlsaProvenanceZeroTwoSlsaBuilderArgs()
        {
        }
    }
}
