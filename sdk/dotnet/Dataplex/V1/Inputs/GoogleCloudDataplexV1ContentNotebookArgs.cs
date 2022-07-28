// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataplex.V1.Inputs
{

    /// <summary>
    /// Configuration for Notebook content.
    /// </summary>
    public sealed class GoogleCloudDataplexV1ContentNotebookArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Kernel Type of the notebook.
        /// </summary>
        [Input("kernelType", required: true)]
        public Input<Pulumi.GoogleNative.Dataplex.V1.GoogleCloudDataplexV1ContentNotebookKernelType> KernelType { get; set; } = null!;

        public GoogleCloudDataplexV1ContentNotebookArgs()
        {
        }
        public static new GoogleCloudDataplexV1ContentNotebookArgs Empty => new GoogleCloudDataplexV1ContentNotebookArgs();
    }
}
