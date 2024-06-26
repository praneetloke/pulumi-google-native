// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.BlockchainNodeEngine.V1.Outputs
{

    /// <summary>
    /// Options for the Geth execution client. See [Command-line Options](https://geth.ethereum.org/docs/fundamentals/command-line-options) for more details.
    /// </summary>
    [OutputType]
    public sealed class GethDetailsResponse
    {
        /// <summary>
        /// Immutable. Blockchain garbage collection mode.
        /// </summary>
        public readonly string GarbageCollectionMode;

        [OutputConstructor]
        private GethDetailsResponse(string garbageCollectionMode)
        {
            GarbageCollectionMode = garbageCollectionMode;
        }
    }
}
