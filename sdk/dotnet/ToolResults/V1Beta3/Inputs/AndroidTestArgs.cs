// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ToolResults.V1Beta3.Inputs
{

    /// <summary>
    /// An Android mobile test specification.
    /// </summary>
    public sealed class AndroidTestArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Information about the application under test.
        /// </summary>
        [Input("androidAppInfo")]
        public Input<Inputs.AndroidAppInfoArgs>? AndroidAppInfo { get; set; }

        /// <summary>
        /// An Android instrumentation test.
        /// </summary>
        [Input("androidInstrumentationTest")]
        public Input<Inputs.AndroidInstrumentationTestArgs>? AndroidInstrumentationTest { get; set; }

        /// <summary>
        /// An Android robo test.
        /// </summary>
        [Input("androidRoboTest")]
        public Input<Inputs.AndroidRoboTestArgs>? AndroidRoboTest { get; set; }

        /// <summary>
        /// An Android test loop.
        /// </summary>
        [Input("androidTestLoop")]
        public Input<Inputs.AndroidTestLoopArgs>? AndroidTestLoop { get; set; }

        /// <summary>
        /// Max time a test is allowed to run before it is automatically cancelled.
        /// </summary>
        [Input("testTimeout")]
        public Input<Inputs.DurationArgs>? TestTimeout { get; set; }

        public AndroidTestArgs()
        {
        }
        public static new AndroidTestArgs Empty => new AndroidTestArgs();
    }
}
