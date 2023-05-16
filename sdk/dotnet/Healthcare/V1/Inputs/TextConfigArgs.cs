// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1.Inputs
{

    public sealed class TextConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("additionalTransformations")]
        private InputList<Inputs.InfoTypeTransformationArgs>? _additionalTransformations;

        /// <summary>
        /// Transformations to apply to the detected data, overridden by `exclude_info_types`.
        /// </summary>
        public InputList<Inputs.InfoTypeTransformationArgs> AdditionalTransformations
        {
            get => _additionalTransformations ?? (_additionalTransformations = new InputList<Inputs.InfoTypeTransformationArgs>());
            set => _additionalTransformations = value;
        }

        [Input("excludeInfoTypes")]
        private InputList<string>? _excludeInfoTypes;

        /// <summary>
        /// InfoTypes to skip transforming, overriding `additional_transformations`.
        /// </summary>
        public InputList<string> ExcludeInfoTypes
        {
            get => _excludeInfoTypes ?? (_excludeInfoTypes = new InputList<string>());
            set => _excludeInfoTypes = value;
        }

        [Input("transformations")]
        private InputList<Inputs.InfoTypeTransformationArgs>? _transformations;

        /// <summary>
        /// The transformations to apply to the detected data. Deprecated. Use `additional_transformations` instead.
        /// </summary>
        [Obsolete(@"The transformations to apply to the detected data. Deprecated. Use `additional_transformations` instead.")]
        public InputList<Inputs.InfoTypeTransformationArgs> Transformations
        {
            get => _transformations ?? (_transformations = new InputList<Inputs.InfoTypeTransformationArgs>());
            set => _transformations = value;
        }

        public TextConfigArgs()
        {
        }
        public static new TextConfigArgs Empty => new TextConfigArgs();
    }
}
