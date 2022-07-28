// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Inputs
{

    /// <summary>
    /// Custom constraint that specifies a key and a list of allowed values for Istio attributes.
    /// </summary>
    public sealed class PermissionConstraintArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Key of the constraint.
        /// </summary>
        [Input("key")]
        public Input<string>? Key { get; set; }

        [Input("values")]
        private InputList<string>? _values;

        /// <summary>
        /// A list of allowed values.
        /// </summary>
        public InputList<string> Values
        {
            get => _values ?? (_values = new InputList<string>());
            set => _values = value;
        }

        public PermissionConstraintArgs()
        {
        }
        public static new PermissionConstraintArgs Empty => new PermissionConstraintArgs();
    }
}
