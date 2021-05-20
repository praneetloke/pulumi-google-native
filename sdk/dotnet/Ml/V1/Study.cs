// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Ml.V1
{
    /// <summary>
    /// Creates a study.
    /// </summary>
    [GoogleNativeResourceType("google-native:ml/v1:Study")]
    public partial class Study : Pulumi.CustomResource
    {
        /// <summary>
        /// Time at which the study was created.
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// A human readable reason why the Study is inactive. This should be empty if a study is ACTIVE or COMPLETED.
        /// </summary>
        [Output("inactiveReason")]
        public Output<string> InactiveReason { get; private set; } = null!;

        /// <summary>
        /// The name of a study.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The detailed state of a study.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// Required. Configuration of the study.
        /// </summary>
        [Output("studyConfig")]
        public Output<Outputs.GoogleCloudMlV1__StudyConfigResponse> StudyConfig { get; private set; } = null!;


        /// <summary>
        /// Create a Study resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Study(string name, StudyArgs args, CustomResourceOptions? options = null)
            : base("google-native:ml/v1:Study", name, args ?? new StudyArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Study(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:ml/v1:Study", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Study resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Study Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Study(name, id, options);
        }
    }

    public sealed class StudyArgs : Pulumi.ResourceArgs
    {
        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("project", required: true)]
        public Input<string> Project { get; set; } = null!;

        /// <summary>
        /// Required. Configuration of the study.
        /// </summary>
        [Input("studyConfig")]
        public Input<Inputs.GoogleCloudMlV1__StudyConfigArgs>? StudyConfig { get; set; }

        [Input("studyId", required: true)]
        public Input<string> StudyId { get; set; } = null!;

        public StudyArgs()
        {
        }
    }
}
