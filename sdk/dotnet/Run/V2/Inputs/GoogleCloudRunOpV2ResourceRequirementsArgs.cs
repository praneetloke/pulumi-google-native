// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Run.V2.Inputs
{

    /// <summary>
    /// ResourceRequirements describes the compute resource requirements.
    /// </summary>
    public sealed class GoogleCloudRunOpV2ResourceRequirementsArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Determines whether CPU should be throttled or not outside of requests.
        /// </summary>
        [Input("cpuIdle")]
        public Input<bool>? CpuIdle { get; set; }

        [Input("limits")]
        private InputMap<string>? _limits;

        /// <summary>
        /// Only memory and CPU are supported. Note: The only supported values for CPU are '1', '2', and '4'. Setting 4 CPU requires at least 2Gi of memory. The values of the map is string form of the 'quantity' k8s type: https://github.com/kubernetes/kubernetes/blob/master/staging/src/k8s.io/apimachinery/pkg/api/resource/quantity.go
        /// </summary>
        public InputMap<string> Limits
        {
            get => _limits ?? (_limits = new InputMap<string>());
            set => _limits = value;
        }

        public GoogleCloudRunOpV2ResourceRequirementsArgs()
        {
        }
    }
}
