// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudDeploy.V1.Inputs
{

    /// <summary>
    /// Information specifying a multiTarget.
    /// </summary>
    public sealed class MultiTargetArgs : global::Pulumi.ResourceArgs
    {
        [Input("targetIds", required: true)]
        private InputList<string>? _targetIds;

        /// <summary>
        /// The target_ids of this multiTarget.
        /// </summary>
        public InputList<string> TargetIds
        {
            get => _targetIds ?? (_targetIds = new InputList<string>());
            set => _targetIds = value;
        }

        public MultiTargetArgs()
        {
        }
        public static new MultiTargetArgs Empty => new MultiTargetArgs();
    }
}
