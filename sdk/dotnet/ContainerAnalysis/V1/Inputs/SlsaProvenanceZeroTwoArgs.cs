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
    /// See full explanation of fields at slsa.dev/provenance/v0.2.
    /// </summary>
    public sealed class SlsaProvenanceZeroTwoArgs : global::Pulumi.ResourceArgs
    {
        [Input("buildConfig")]
        private InputMap<object>? _buildConfig;
        public InputMap<object> BuildConfig
        {
            get => _buildConfig ?? (_buildConfig = new InputMap<object>());
            set => _buildConfig = value;
        }

        [Input("buildType")]
        public Input<string>? BuildType { get; set; }

        [Input("builder")]
        public Input<Inputs.GrafeasV1SlsaProvenanceZeroTwoSlsaBuilderArgs>? Builder { get; set; }

        [Input("invocation")]
        public Input<Inputs.GrafeasV1SlsaProvenanceZeroTwoSlsaInvocationArgs>? Invocation { get; set; }

        [Input("materials")]
        private InputList<Inputs.GrafeasV1SlsaProvenanceZeroTwoSlsaMaterialArgs>? _materials;
        public InputList<Inputs.GrafeasV1SlsaProvenanceZeroTwoSlsaMaterialArgs> Materials
        {
            get => _materials ?? (_materials = new InputList<Inputs.GrafeasV1SlsaProvenanceZeroTwoSlsaMaterialArgs>());
            set => _materials = value;
        }

        [Input("metadata")]
        public Input<Inputs.GrafeasV1SlsaProvenanceZeroTwoSlsaMetadataArgs>? Metadata { get; set; }

        public SlsaProvenanceZeroTwoArgs()
        {
        }
        public static new SlsaProvenanceZeroTwoArgs Empty => new SlsaProvenanceZeroTwoArgs();
    }
}
