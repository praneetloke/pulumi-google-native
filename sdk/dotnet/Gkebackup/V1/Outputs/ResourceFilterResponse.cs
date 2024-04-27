// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Gkebackup.V1.Outputs
{

    /// <summary>
    /// ResourceFilter specifies matching criteria to limit the scope of a change to a specific set of kubernetes resources that are selected for restoration from a backup.
    /// </summary>
    [OutputType]
    public sealed class ResourceFilterResponse
    {
        /// <summary>
        /// Optional. (Filtering parameter) Any resource subject to transformation must belong to one of the listed "types". If this field is not provided, no type filtering will be performed (all resources of all types matching previous filtering parameters will be candidates for transformation).
        /// </summary>
        public readonly ImmutableArray<Outputs.GroupKindResponse> GroupKinds;
        /// <summary>
        /// Optional. This is a [JSONPath] (https://github.com/json-path/JsonPath/blob/master/README.md) expression that matches specific fields of candidate resources and it operates as a filtering parameter (resources that are not matched with this expression will not be candidates for transformation).
        /// </summary>
        public readonly string JsonPath;
        /// <summary>
        /// Optional. (Filtering parameter) Any resource subject to transformation must be contained within one of the listed Kubernetes Namespace in the Backup. If this field is not provided, no namespace filtering will be performed (all resources in all Namespaces, including all cluster-scoped resources, will be candidates for transformation).
        /// </summary>
        public readonly ImmutableArray<string> Namespaces;

        [OutputConstructor]
        private ResourceFilterResponse(
            ImmutableArray<Outputs.GroupKindResponse> groupKinds,

            string jsonPath,

            ImmutableArray<string> namespaces)
        {
            GroupKinds = groupKinds;
            JsonPath = jsonPath;
            Namespaces = namespaces;
        }
    }
}