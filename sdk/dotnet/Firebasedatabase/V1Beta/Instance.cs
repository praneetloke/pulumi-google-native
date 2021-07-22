// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Firebasedatabase.V1Beta
{
    /// <summary>
    /// Requests that a new DatabaseInstance be created. The state of a successfully created DatabaseInstance is ACTIVE. Only available for projects on the Blaze plan. Projects can be upgraded using the Cloud Billing API https://cloud.google.com/billing/reference/rest/v1/projects/updateBillingInfo. Note that it might take a few minutes for billing enablement state to propagate to Firebase systems.
    /// </summary>
    [GoogleNativeResourceType("google-native:firebasedatabase/v1beta:Instance")]
    public partial class Instance : Pulumi.CustomResource
    {
        /// <summary>
        /// Immutable. The globally unique hostname of the database.
        /// </summary>
        [Output("databaseUrl")]
        public Output<string> DatabaseUrl { get; private set; } = null!;

        /// <summary>
        /// The fully qualified resource name of the database instance, in the form: `projects/{project-number}/locations/{location-id}/instances/{database-id}`.
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The resource name of the project this instance belongs to. For example: `projects/{project-number}`.
        /// </summary>
        [Output("project")]
        public Output<string> Project { get; private set; } = null!;

        /// <summary>
        /// The database's lifecycle state. Read-only.
        /// </summary>
        [Output("state")]
        public Output<string> State { get; private set; } = null!;

        /// <summary>
        /// The database instance type. On creation only USER_DATABASE is allowed, which is also the default when omitted.
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a Instance resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Instance(string name, InstanceArgs? args = null, CustomResourceOptions? options = null)
            : base("google-native:firebasedatabase/v1beta:Instance", name, args ?? new InstanceArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Instance(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:firebasedatabase/v1beta:Instance", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Instance resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Instance Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Instance(name, id, options);
        }
    }

    public sealed class InstanceArgs : Pulumi.ResourceArgs
    {
        [Input("databaseId")]
        public Input<string>? DatabaseId { get; set; }

        /// <summary>
        /// Immutable. The globally unique hostname of the database.
        /// </summary>
        [Input("databaseUrl")]
        public Input<string>? DatabaseUrl { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// The fully qualified resource name of the database instance, in the form: `projects/{project-number}/locations/{location-id}/instances/{database-id}`.
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// The resource name of the project this instance belongs to. For example: `projects/{project-number}`.
        /// </summary>
        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// The database's lifecycle state. Read-only.
        /// </summary>
        [Input("state")]
        public Input<Pulumi.GoogleNative.Firebasedatabase.V1Beta.InstanceState>? State { get; set; }

        /// <summary>
        /// The database instance type. On creation only USER_DATABASE is allowed, which is also the default when omitted.
        /// </summary>
        [Input("type")]
        public Input<Pulumi.GoogleNative.Firebasedatabase.V1Beta.InstanceType>? Type { get; set; }

        [Input("validateOnly")]
        public Input<string>? ValidateOnly { get; set; }

        public InstanceArgs()
        {
        }
    }
}
