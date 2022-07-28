// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ServiceManagement.V1.Inputs
{

    /// <summary>
    /// An object that describes the schema of a MonitoredResource object using a type name and a set of labels. For example, the monitored resource descriptor for Google Compute Engine VM instances has a type of `"gce_instance"` and specifies the use of the labels `"instance_id"` and `"zone"` to identify particular VM instances. Different APIs can support different monitored resource types. APIs generally provide a `list` method that returns the monitored resource descriptors used by the API. 
    /// </summary>
    public sealed class MonitoredResourceDescriptorArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. A detailed description of the monitored resource type that might be used in documentation.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// Optional. A concise name for the monitored resource type that might be displayed in user interfaces. It should be a Title Cased Noun Phrase, without any article or other determiners. For example, `"Google Cloud SQL Database"`.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("labels", required: true)]
        private InputList<Inputs.LabelDescriptorArgs>? _labels;

        /// <summary>
        /// A set of labels used to describe instances of this monitored resource type. For example, an individual Google Cloud SQL database is identified by values for the labels `"database_id"` and `"zone"`.
        /// </summary>
        public InputList<Inputs.LabelDescriptorArgs> Labels
        {
            get => _labels ?? (_labels = new InputList<Inputs.LabelDescriptorArgs>());
            set => _labels = value;
        }

        /// <summary>
        /// Optional. The launch stage of the monitored resource definition.
        /// </summary>
        [Input("launchStage")]
        public Input<Pulumi.GoogleNative.ServiceManagement.V1.MonitoredResourceDescriptorLaunchStage>? LaunchStage { get; set; }

        /// <summary>
        /// Optional. The resource name of the monitored resource descriptor: `"projects/{project_id}/monitoredResourceDescriptors/{type}"` where {type} is the value of the `type` field in this object and {project_id} is a project ID that provides API-specific context for accessing the type. APIs that do not use project information can use the resource name format `"monitoredResourceDescriptors/{type}"`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The monitored resource type. For example, the type `"cloudsql_database"` represents databases in Google Cloud SQL. For a list of types, see [Monitoring resource types](https://cloud.google.com/monitoring/api/resources) and [Logging resource types](https://cloud.google.com/logging/docs/api/v2/resource-list).
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public MonitoredResourceDescriptorArgs()
        {
        }
        public static new MonitoredResourceDescriptorArgs Empty => new MonitoredResourceDescriptorArgs();
    }
}
