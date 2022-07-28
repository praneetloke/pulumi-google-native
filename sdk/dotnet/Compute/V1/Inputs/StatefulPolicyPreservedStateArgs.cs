// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.V1.Inputs
{

    /// <summary>
    /// Configuration of preserved resources.
    /// </summary>
    public sealed class StatefulPolicyPreservedStateArgs : global::Pulumi.ResourceArgs
    {
        [Input("disks")]
        private InputMap<string>? _disks;

        /// <summary>
        /// Disks created on the instances that will be preserved on instance delete, update, etc. This map is keyed with the device names of the disks.
        /// </summary>
        public InputMap<string> Disks
        {
            get => _disks ?? (_disks = new InputMap<string>());
            set => _disks = value;
        }

        public StatefulPolicyPreservedStateArgs()
        {
        }
        public static new StatefulPolicyPreservedStateArgs Empty => new StatefulPolicyPreservedStateArgs();
    }
}
