// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DeploymentManager.Alpha.Outputs
{

    [OutputType]
    public sealed class InstancesBulkInsertOperationMetadataResponse
    {
        /// <summary>
        /// Status information per location (location name is key). Example key: zones/us-central1-a
        /// </summary>
        public readonly Outputs.BulkInsertOperationStatusResponse PerLocationStatus;

        [OutputConstructor]
        private InstancesBulkInsertOperationMetadataResponse(Outputs.BulkInsertOperationStatusResponse perLocationStatus)
        {
            PerLocationStatus = perLocationStatus;
        }
    }
}
