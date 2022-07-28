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
    /// A test of an Android application that can control an Android component independently of its normal lifecycle. Android instrumentation tests run an application APK and test APK inside the same process on a virtual or physical AndroidDevice. They also specify a test runner class, such as com.google.GoogleTestRunner, which can vary on the specific instrumentation framework chosen. See for more information on types of Android tests.
    /// </summary>
    public sealed class AndroidInstrumentationTestArgs : global::Pulumi.ResourceArgs
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
        /// The java package for the application under test. The default value is determined by examining the application's manifest.
        /// </summary>
        [Input("appPackageId")]
        public Input<string>? AppPackageId { get; set; }

        /// <summary>
        /// The option of whether running each test within its own invocation of instrumentation with Android Test Orchestrator or not. ** Orchestrator is only compatible with AndroidJUnitRunner version 1.1 or higher! ** Orchestrator offers the following benefits: - No shared state - Crashes are isolated - Logs are scoped per test See for more information about Android Test Orchestrator. If not set, the test will be run without the orchestrator.
        /// </summary>
        [Input("orchestratorOption")]
        public Input<Pulumi.GoogleNative.Testing.V1.AndroidInstrumentationTestOrchestratorOption>? OrchestratorOption { get; set; }

        /// <summary>
        /// The option to run tests in multiple shards in parallel.
        /// </summary>
        [Input("shardingOption")]
        public Input<Inputs.ShardingOptionArgs>? ShardingOption { get; set; }

        /// <summary>
        /// The APK containing the test code to be executed.
        /// </summary>
        [Input("testApk", required: true)]
        public Input<Inputs.FileReferenceArgs> TestApk { get; set; } = null!;

        /// <summary>
        /// The java package for the test to be executed. The default value is determined by examining the application's manifest.
        /// </summary>
        [Input("testPackageId")]
        public Input<string>? TestPackageId { get; set; }

        /// <summary>
        /// The InstrumentationTestRunner class. The default value is determined by examining the application's manifest.
        /// </summary>
        [Input("testRunnerClass")]
        public Input<string>? TestRunnerClass { get; set; }

        [Input("testTargets")]
        private InputList<string>? _testTargets;

        /// <summary>
        /// Each target must be fully qualified with the package name or class name, in one of these formats: - "package package_name" - "class package_name.class_name" - "class package_name.class_name#method_name" If empty, all targets in the module will be run.
        /// </summary>
        public InputList<string> TestTargets
        {
            get => _testTargets ?? (_testTargets = new InputList<string>());
            set => _testTargets = value;
        }

        public AndroidInstrumentationTestArgs()
        {
        }
        public static new AndroidInstrumentationTestArgs Empty => new AndroidInstrumentationTestArgs();
    }
}
