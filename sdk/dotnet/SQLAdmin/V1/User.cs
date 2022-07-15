// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.SQLAdmin.V1
{
    /// <summary>
    /// Creates a new user in a Cloud SQL instance.
    /// </summary>
    [GoogleNativeResourceType("google-native:sqladmin/v1:User")]
    public partial class User : Pulumi.CustomResource
    {
        /// <summary>
        /// Dual password status for the user.
        /// </summary>
        [Output("dualPasswordType")]
        public Output<string> DualPasswordType { get; private set; } = null!;

        /// <summary>
        /// This field is deprecated and will be removed from a future version of the API.
        /// </summary>
        [Output("etag")]
        public Output<string> Etag { get; private set; } = null!;

        /// <summary>
        /// Optional. The host from which the user can connect. For `insert` operations, host defaults to an empty string. For `update` operations, host is specified as part of the request URL. The host name cannot be updated after insertion. For a MySQL instance, it's required; for a PostgreSQL or SQL Server instance, it's optional.
        /// </summary>
        [Output("host")]
        public Output<string> Host { get; private set; } = null!;

        [Output("instance")]
        public Output<string> Instance { get; private set; } = null!;

        /// <summary>
        /// This is always `sql#user`.
        /// </summary>
        [Output("kind")]
        public Output<string> Kind { get; private set; } = null!;

        /// <summary>
        /// The name of the user in the Cloud SQL instance. Can be omitted for `update` because it is already specified in the URL.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The password for the user.
        /// </summary>
        [Output("password")]
        public Output<string> Password { get; private set; } = null!;

        /// <summary>
        /// User level password validation policy.
        /// </summary>
        [Output("passwordPolicy")]
        public Output<Outputs.UserPasswordValidationPolicyResponse> PasswordPolicy { get; private set; } = null!;

        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        [Output("sqlserverUserDetails")]
        public Output<Outputs.SqlServerUserDetailsResponse> SqlserverUserDetails { get; private set; } = null!;

        /// <summary>
        /// The user type. It determines the method to authenticate the user during login. The default is the database's built-in user type.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a User resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public User(string name, UserArgs args, CustomResourceOptions? options = null)
            : base("google-native:sqladmin/v1:User", name, args ?? new UserArgs(), MakeResourceOptions(options, ""))
        {
        }

        private User(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:sqladmin/v1:User", name, null, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                ReplaceOnChanges =
                {
                    "instance",
                    "project",
                },
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing User resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static User Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new User(name, id, options);
        }
    }

    public sealed class UserArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Dual password status for the user.
        /// </summary>
        [Input("dualPasswordType")]
        public Input<Pulumi.GoogleNative.SQLAdmin.V1.UserDualPasswordType>? DualPasswordType { get; set; }

        /// <summary>
        /// This field is deprecated and will be removed from a future version of the API.
        /// </summary>
        [Input("etag")]
        public Input<string>? Etag { get; set; }

        /// <summary>
        /// Optional. The host from which the user can connect. For `insert` operations, host defaults to an empty string. For `update` operations, host is specified as part of the request URL. The host name cannot be updated after insertion. For a MySQL instance, it's required; for a PostgreSQL or SQL Server instance, it's optional.
        /// </summary>
        [Input("host")]
        public Input<string>? Host { get; set; }

        /// <summary>
        /// The name of the Cloud SQL instance. This does not include the project ID. Can be omitted for `update` because it is already specified on the URL.
        /// </summary>
        [Input("instance", required: true)]
        public Input<string> Instance { get; set; } = null!;

        /// <summary>
        /// This is always `sql#user`.
        /// </summary>
        [Input("kind")]
        public Input<string>? Kind { get; set; }

        /// <summary>
        /// The name of the user in the Cloud SQL instance. Can be omitted for `update` because it is already specified in the URL.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The password for the user.
        /// </summary>
        [Input("password")]
        public Input<string>? Password { get; set; }

        /// <summary>
        /// User level password validation policy.
        /// </summary>
        [Input("passwordPolicy")]
        public Input<Inputs.UserPasswordValidationPolicyArgs>? PasswordPolicy { get; set; }

        /// <summary>
        /// The project ID of the project containing the Cloud SQL database. The Google apps domain is prefixed if applicable. Can be omitted for `update` because it is already specified on the URL.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        [Input("sqlserverUserDetails")]
        public Input<Inputs.SqlServerUserDetailsArgs>? SqlserverUserDetails { get; set; }

        /// <summary>
        /// The user type. It determines the method to authenticate the user during login. The default is the database's built-in user type.
        /// </summary>
        [Input("type")]
        public Input<Pulumi.GoogleNative.SQLAdmin.V1.UserType>? Type { get; set; }

        public UserArgs()
        {
        }
    }
}