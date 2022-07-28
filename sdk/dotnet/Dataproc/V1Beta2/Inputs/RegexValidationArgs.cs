// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1Beta2.Inputs
{

    /// <summary>
    /// Validation based on regular expressions.
    /// </summary>
    public sealed class RegexValidationArgs : global::Pulumi.ResourceArgs
    {
        [Input("regexes", required: true)]
        private InputList<string>? _regexes;

        /// <summary>
        /// RE2 regular expressions used to validate the parameter's value. The value must match the regex in its entirety (substring matches are not sufficient).
        /// </summary>
        public InputList<string> Regexes
        {
            get => _regexes ?? (_regexes = new InputList<string>());
            set => _regexes = value;
        }

        public RegexValidationArgs()
        {
        }
        public static new RegexValidationArgs Empty => new RegexValidationArgs();
    }
}
