// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.SQLAdmin.V1Beta4
{
    public static class GetDatabase
    {
        /// <summary>
        /// Retrieves a resource containing information about a database inside a Cloud SQL instance.
        /// </summary>
        public static Task<GetDatabaseResult> InvokeAsync(GetDatabaseArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetDatabaseResult>("google-native:sqladmin/v1beta4:getDatabase", args ?? new GetDatabaseArgs(), options.WithDefaults());

        /// <summary>
        /// Retrieves a resource containing information about a database inside a Cloud SQL instance.
        /// </summary>
        public static Output<GetDatabaseResult> Invoke(GetDatabaseInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetDatabaseResult>("google-native:sqladmin/v1beta4:getDatabase", args ?? new GetDatabaseInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetDatabaseArgs : global::Pulumi.InvokeArgs
    {
        [Input("database", required: true)]
        public string Database { get; set; } = null!;

        [Input("instance", required: true)]
        public string Instance { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetDatabaseArgs()
        {
        }
        public static new GetDatabaseArgs Empty => new GetDatabaseArgs();
    }

    public sealed class GetDatabaseInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("database", required: true)]
        public Input<string> Database { get; set; } = null!;

        [Input("instance", required: true)]
        public Input<string> Instance { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetDatabaseInvokeArgs()
        {
        }
        public static new GetDatabaseInvokeArgs Empty => new GetDatabaseInvokeArgs();
    }


    [OutputType]
    public sealed class GetDatabaseResult
    {
        /// <summary>
        /// The Cloud SQL charset value.
        /// </summary>
        public readonly string Charset;
        /// <summary>
        /// The Cloud SQL collation value.
        /// </summary>
        public readonly string Collation;
        /// <summary>
        /// This field is deprecated and will be removed from a future version of the API.
        /// </summary>
        public readonly string Etag;
        /// <summary>
        /// The name of the Cloud SQL instance. This does not include the project ID.
        /// </summary>
        public readonly string Instance;
        /// <summary>
        /// This is always `sql#database`.
        /// </summary>
        public readonly string Kind;
        /// <summary>
        /// The name of the database in the Cloud SQL instance. This does not include the project ID or instance name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The project ID of the project containing the Cloud SQL database. The Google apps domain is prefixed if applicable.
        /// </summary>
        public readonly string Project;
        /// <summary>
        /// The URI of this resource.
        /// </summary>
        public readonly string SelfLink;
        public readonly Outputs.SqlServerDatabaseDetailsResponse SqlserverDatabaseDetails;

        [OutputConstructor]
        private GetDatabaseResult(
            string charset,

            string collation,

            string etag,

            string instance,

            string kind,

            string name,

            string project,

            string selfLink,

            Outputs.SqlServerDatabaseDetailsResponse sqlserverDatabaseDetails)
        {
            Charset = charset;
            Collation = collation;
            Etag = etag;
            Instance = instance;
            Kind = kind;
            Name = name;
            Project = project;
            SelfLink = selfLink;
            SqlserverDatabaseDetails = sqlserverDatabaseDetails;
        }
    }
}
