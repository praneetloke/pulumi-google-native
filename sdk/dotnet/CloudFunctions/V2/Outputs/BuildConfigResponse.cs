// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.CloudFunctions.V2.Outputs
{

    /// <summary>
    /// Describes the Build step of the function that builds a container from the given source.
    /// </summary>
    [OutputType]
    public sealed class BuildConfigResponse
    {
        /// <summary>
        /// The Cloud Build name of the latest successful deployment of the function.
        /// </summary>
        public readonly string Build;
        /// <summary>
        /// Optional. User managed repository created in Artifact Registry optionally with a customer managed encryption key. This is the repository to which the function docker image will be pushed after it is built by Cloud Build. If unspecified, GCF will create and use a repository named 'gcf-artifacts' for every deployed region. It must match the pattern `projects/{project}/locations/{location}/repositories/{repository}`. Cross-project repositories are not supported. Cross-location repositories are not supported. Repository format must be 'DOCKER'.
        /// </summary>
        public readonly string DockerRepository;
        /// <summary>
        /// The name of the function (as defined in source code) that will be executed. Defaults to the resource name suffix, if not specified. For backward compatibility, if function with given name is not found, then the system will try to use function named "function". For Node.js this is name of a function exported by the module specified in `source_location`.
        /// </summary>
        public readonly string EntryPoint;
        /// <summary>
        /// User-provided build-time environment variables for the function
        /// </summary>
        public readonly ImmutableDictionary<string, string> EnvironmentVariables;
        /// <summary>
        /// The runtime in which to run the function. Required when deploying a new function, optional when updating an existing function. For a complete list of possible choices, see the [`gcloud` command reference](https://cloud.google.com/sdk/gcloud/reference/functions/deploy#--runtime).
        /// </summary>
        public readonly string Runtime;
        /// <summary>
        /// The location of the function source code.
        /// </summary>
        public readonly Outputs.SourceResponse Source;
        /// <summary>
        /// A permanent fixed identifier for source.
        /// </summary>
        public readonly Outputs.SourceProvenanceResponse SourceProvenance;
        /// <summary>
        /// Name of the Cloud Build Custom Worker Pool that should be used to build the function. The format of this field is `projects/{project}/locations/{region}/workerPools/{workerPool}` where {project} and {region} are the project id and region respectively where the worker pool is defined and {workerPool} is the short name of the worker pool. If the project id is not the same as the function, then the Cloud Functions Service Agent (service-@gcf-admin-robot.iam.gserviceaccount.com) must be granted the role Cloud Build Custom Workers Builder (roles/cloudbuild.customworkers.builder) in the project.
        /// </summary>
        public readonly string WorkerPool;

        [OutputConstructor]
        private BuildConfigResponse(
            string build,

            string dockerRepository,

            string entryPoint,

            ImmutableDictionary<string, string> environmentVariables,

            string runtime,

            Outputs.SourceResponse source,

            Outputs.SourceProvenanceResponse sourceProvenance,

            string workerPool)
        {
            Build = build;
            DockerRepository = dockerRepository;
            EntryPoint = entryPoint;
            EnvironmentVariables = environmentVariables;
            Runtime = runtime;
            Source = source;
            SourceProvenance = sourceProvenance;
            WorkerPool = workerPool;
        }
    }
}
