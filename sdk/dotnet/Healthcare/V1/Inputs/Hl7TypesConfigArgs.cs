// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Healthcare.V1.Inputs
{

    /// <summary>
    /// Root config for HL7v2 datatype definitions for a specific HL7v2 version.
    /// </summary>
    public sealed class Hl7TypesConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("type")]
        private InputList<Inputs.TypeArgs>? _type;

        /// <summary>
        /// The HL7v2 type definitions.
        /// </summary>
        public InputList<Inputs.TypeArgs> Type
        {
            get => _type ?? (_type = new InputList<Inputs.TypeArgs>());
            set => _type = value;
        }

        [Input("version")]
        private InputList<Inputs.VersionSourceArgs>? _version;

        /// <summary>
        /// The version selectors that this config applies to. A message must match ALL version sources to apply.
        /// </summary>
        public InputList<Inputs.VersionSourceArgs> Version
        {
            get => _version ?? (_version = new InputList<Inputs.VersionSourceArgs>());
            set => _version = value;
        }

        public Hl7TypesConfigArgs()
        {
        }
        public static new Hl7TypesConfigArgs Empty => new Hl7TypesConfigArgs();
    }
}
