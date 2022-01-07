// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Testing.V1.Inputs
{

    /// <summary>
    /// A test of an android application that explores the application on a virtual or physical Android Device, finding culprits and crashes as it goes.
    /// </summary>
    public sealed class AndroidRoboTestArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// The APK for the application under test.
        /// </summary>
        [Input("appApk")]
        public Input<Inputs.FileReferenceArgs>? AppApk { get; set; }

        /// <summary>
        /// A multi-apk app bundle for the application under test.
        /// </summary>
        [Input("appBundle")]
        public Input<Inputs.AppBundleArgs>? AppBundle { get; set; }

        /// <summary>
        /// The initial activity that should be used to start the app.
        /// </summary>
        [Input("appInitialActivity")]
        public Input<string>? AppInitialActivity { get; set; }

        /// <summary>
        /// The java package for the application under test. The default value is determined by examining the application's manifest.
        /// </summary>
        [Input("appPackageId")]
        public Input<string>? AppPackageId { get; set; }

        [Input("roboDirectives")]
        private InputList<Inputs.RoboDirectiveArgs>? _roboDirectives;

        /// <summary>
        /// A set of directives Robo should apply during the crawl. This allows users to customize the crawl. For example, the username and password for a test account can be provided.
        /// </summary>
        public InputList<Inputs.RoboDirectiveArgs> RoboDirectives
        {
            get => _roboDirectives ?? (_roboDirectives = new InputList<Inputs.RoboDirectiveArgs>());
            set => _roboDirectives = value;
        }

        /// <summary>
        /// The mode in which Robo should run. Most clients should allow the server to populate this field automatically.
        /// </summary>
        [Input("roboMode")]
        public Input<Pulumi.GoogleNative.Testing.V1.AndroidRoboTestRoboMode>? RoboMode { get; set; }

        /// <summary>
        /// A JSON file with a sequence of actions Robo should perform as a prologue for the crawl.
        /// </summary>
        [Input("roboScript")]
        public Input<Inputs.FileReferenceArgs>? RoboScript { get; set; }

        [Input("startingIntents")]
        private InputList<Inputs.RoboStartingIntentArgs>? _startingIntents;

        /// <summary>
        /// The intents used to launch the app for the crawl. If none are provided, then the main launcher activity is launched. If some are provided, then only those provided are launched (the main launcher activity must be provided explicitly).
        /// </summary>
        public InputList<Inputs.RoboStartingIntentArgs> StartingIntents
        {
            get => _startingIntents ?? (_startingIntents = new InputList<Inputs.RoboStartingIntentArgs>());
            set => _startingIntents = value;
        }

        public AndroidRoboTestArgs()
        {
        }
    }
}
