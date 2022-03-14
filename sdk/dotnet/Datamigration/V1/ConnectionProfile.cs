// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datamigration.V1
{
    /// <summary>
    /// Creates a new connection profile in a given project and location.
    /// </summary>
    [GoogleNativeResourceType("google-native:datamigration/v1:ConnectionProfile")]
    public partial class ConnectionProfile : Pulumi.CustomResource
    {
        /// <summary>
        /// A CloudSQL database connection profile.
        /// </summary>
        [Output("cloudsql")]
        public Output<Outputs.CloudSqlConnectionProfileResponse> Cloudsql { get; private set; } = null!;

        /// <summary>
        /// The timestamp when the resource was created. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Output("createTime")]
        public Output<string> CreateTime { get; private set; } = null!;

        /// <summary>
        /// The connection profile display name.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// The error details in case of state FAILED.
        /// </summary>
        [Output("error")]
        public Output<Outputs.StatusResponse> Error { get; private set; } = null!;

        /// <summary>
        /// The resource labels for connection profile to use to annotate any related underlying resources such as Compute Engine VMs. An object containing a list of "key": "value" pairs. Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
        /// </summary>
        [Output("labels")]
        public Output<ImmutableDictionary<string, string>> Labels { get; private set; } = null!;

        /// <summary>
        /// A MySQL database connection profile.
        /// </summary>
        [Output("mysql")]
        public Output<Outputs.MySqlConnectionProfileResponse> Mysql { get; private set; } = null!;

        /// <summary>
        /// The name of this connection profile resource in the form of projects/{project}/locations/{location}/connectionProfiles/{connectionProfile}.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A PostgreSQL database connection profile.
        /// </summary>
        [Output("postgresql")]
        public Output<Outputs.PostgreSqlConnectionProfileResponse> Postgresql { get; private set; } = null!;

        /// <summary>
        /// The database provider.
        /// </summary>
        [Output("provider")]
        public Output<string> Provider { get; private set; } = null!;

        /// <summary>
        /// The current connection profile state (e.g. DRAFT, READY, or FAILED).
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The timestamp when the resource was last updated. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
        /// </summary>
        [Output("updateTime")]
        public Output<string> UpdateTime { get; private set; } = null!;


        /// <summary>
        /// Create a ConnectionProfile resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ConnectionProfile(string name, ConnectionProfileArgs args, CustomResourceOptions? options = null)
            : base("google-native:datamigration/v1:ConnectionProfile", name, args ?? new ConnectionProfileArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ConnectionProfile(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:datamigration/v1:ConnectionProfile", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing ConnectionProfile resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ConnectionProfile Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new ConnectionProfile(name, id, options);
        }
    }

    public sealed class ConnectionProfileArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// A CloudSQL database connection profile.
        /// </summary>
        [Input("cloudsql")]
        public Input<Inputs.CloudSqlConnectionProfileArgs>? Cloudsql { get; set; }

        /// <summary>
        /// Required. The connection profile identifier.
        /// </summary>
        [Input("connectionProfileId", required: true)]
        public Input<string> ConnectionProfileId { get; set; } = null!;

        /// <summary>
        /// The connection profile display name.
        /// </summary>
        [Input("displayName")]
        public Input<string>? DisplayName { get; set; }

        [Input("labels")]
        private InputMap<string>? _labels;

        /// <summary>
        /// The resource labels for connection profile to use to annotate any related underlying resources such as Compute Engine VMs. An object containing a list of "key": "value" pairs. Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
        /// </summary>
        public InputMap<string> Labels
        {
            get => _labels ?? (_labels = new InputMap<string>());
            set => _labels = value;
        }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// A MySQL database connection profile.
        /// </summary>
        [Input("mysql")]
        public Input<Inputs.MySqlConnectionProfileArgs>? Mysql { get; set; }

        /// <summary>
        /// The name of this connection profile resource in the form of projects/{project}/locations/{location}/connectionProfiles/{connectionProfile}.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// A PostgreSQL database connection profile.
        /// </summary>
        [Input("postgresql")]
        public Input<Inputs.PostgreSqlConnectionProfileArgs>? Postgresql { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The database provider.
        /// </summary>
        [Input("provider")]
        public Input<Pulumi.GoogleNative.Datamigration.V1.ConnectionProfileProvider>? Provider { get; set; }

        /// <summary>
        /// A unique id used to identify the request. If the server receives two requests with the same id, then the second request will be ignored. It is recommended to always set this value to a UUID. The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). The maximum length is 40 characters.
        /// </summary>
        [Input("requestId")]
        public Input<string>? RequestId { get; set; }

        /// <summary>
        /// The current connection profile state (e.g. DRAFT, READY, or FAILED).
        /// </summary>
        [Input("state")]
        public Input<Pulumi.GoogleNative.Datamigration.V1.ConnectionProfileState>? State { get; set; }

        public ConnectionProfileArgs()
        {
        }
    }
}
