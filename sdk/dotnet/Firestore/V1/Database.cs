// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Firestore.V1
{
    /// <summary>
    /// Create a database.
    /// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
    /// on Google Cloud even though it will be deleted from Pulumi state.
    /// </summary>
    [GoogleNativeResourceType("google-native:firestore/v1:Database")]
    public partial class Database : global::Pulumi.CustomResource
    {
        /// <summary>
        /// The App Engine integration mode to use for this database.
        /// </summary>
        [Output("appEngineIntegrationMode")]
        public Output<string> AppEngineIntegrationMode { get; private set; } = null!;

        /// <summary>
        /// The concurrency control mode to use for this database.
        /// </summary>
        [Output("concurrencyMode")]
        public Output<string> ConcurrencyMode { get; private set; } = null!;

        /// <summary>
        /// Required. The ID to use for the database, which will become the final component of the database's resource name. This value should be 4-63 characters. Valid characters are /a-z-/ with first character a letter and the last a letter or a number. Must not be UUID-like /[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}/. "(default)" database id is also valid.
        /// </summary>
        [Output("databaseId")]
        public Output<string> DatabaseId { get; private set; } = null!;

        /// <summary>
        /// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// The key_prefix for this database. This key_prefix is used, in combination with the project id ("~") to construct the application id that is returned from the Cloud Datastore APIs in Google App Engine first generation runtimes. This value may be empty in which case the appid to use for URL-encoded keys is the project_id (eg: foo instead of v~foo).
        /// </summary>
        [Output("keyPrefix")]
        public Output<string> KeyPrefix { get; private set; } = null!;

        /// <summary>
        /// The location of the database. Available databases are listed at https://cloud.google.com/firestore/docs/locations.
        /// </summary>
        [Output("location")]
        public Output<string> Location { get; private set; } = null!;

        /// <summary>
        /// The resource name of the Database. Format: `projects/{project}/databases/{database}`
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The type of the database. See https://cloud.google.com/datastore/docs/firestore-or-datastore for information about how to choose.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Database resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Database(string name, DatabaseArgs args, CustomResourceOptions? options = null)
            : base("google-native:firestore/v1:Database", name, args ?? new DatabaseArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Database(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:firestore/v1:Database", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "databaseId",
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Database resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Database Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Database(name, id, options);
        }
    }

    public sealed class DatabaseArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// The App Engine integration mode to use for this database.
        /// </summary>
        [Input("appEngineIntegrationMode")]
        public Input<Pulumi.GoogleNative.Firestore.V1.DatabaseAppEngineIntegrationMode>? AppEngineIntegrationMode { get; set; }

        /// <summary>
        /// The concurrency control mode to use for this database.
        /// </summary>
        [Input("concurrencyMode")]
        public Input<Pulumi.GoogleNative.Firestore.V1.DatabaseConcurrencyMode>? ConcurrencyMode { get; set; }

        /// <summary>
        /// Required. The ID to use for the database, which will become the final component of the database's resource name. This value should be 4-63 characters. Valid characters are /a-z-/ with first character a letter and the last a letter or a number. Must not be UUID-like /[0-9a-f]{8}(-[0-9a-f]{4}){3}-[0-9a-f]{12}/. "(default)" database id is also valid.
        /// </summary>
        [Input("databaseId", required: true)]
        public Input<string> DatabaseId { get; set; } = null!;

        /// <summary>
        /// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// The location of the database. Available databases are listed at https://cloud.google.com/firestore/docs/locations.
        /// </summary>
        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The resource name of the Database. Format: `projects/{project}/databases/{database}`
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The type of the database. See https://cloud.google.com/datastore/docs/firestore-or-datastore for information about how to choose.
        /// </summary>
        [Input("type")]
        public Input<Pulumi.GoogleNative.Firestore.V1.DatabaseType>? Type { get; set; }

        public DatabaseArgs()
        {
        }
        public static new DatabaseArgs Empty => new DatabaseArgs();
    }
}
