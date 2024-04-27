// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1.Inputs
{

    /// <summary>
    /// Instance flexibility Policy allowing a mixture of VM shapes and provisioning models.
    /// </summary>
    public sealed class InstanceFlexibilityPolicyArgs : global::Pulumi.ResourceArgs
    {
        [Input("instanceSelectionList")]
        private InputList<Inputs.InstanceSelectionArgs>? _instanceSelectionList;

        /// <summary>
        /// Optional. List of instance selection options that the group will use when creating new VMs.
        /// </summary>
        public InputList<Inputs.InstanceSelectionArgs> InstanceSelectionList
        {
            get => _instanceSelectionList ?? (_instanceSelectionList = new InputList<Inputs.InstanceSelectionArgs>());
            set => _instanceSelectionList = value;
        }

        public InstanceFlexibilityPolicyArgs()
        {
        }
        public static new InstanceFlexibilityPolicyArgs Empty => new InstanceFlexibilityPolicyArgs();
    }
}