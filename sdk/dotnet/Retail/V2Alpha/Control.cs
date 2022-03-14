// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Retail.V2Alpha
{
    /// <summary>
    /// Creates a Control. If the Control to create already exists, an ALREADY_EXISTS error is returned.
    /// Auto-naming is currently not supported for this resource.
    /// </summary>
    [GoogleNativeResourceType("google-native:retail/v2alpha:Control")]
    public partial class Control : Pulumi.CustomResource
    {
        /// <summary>
        /// List of serving configuration ids that that are associated with this control. Note the association is managed via the ServingConfig, this is an output only denormalizeed view. Assumed to be in the same catalog.
        /// </summary>
        [Output("associatedServingConfigIds")]
        public Output<ImmutableArray<string>> AssociatedServingConfigIds { get; private set; } = null!;

        /// <summary>
        /// The human readable control display name. Used in Retail UI. This field must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is thrown.
        /// </summary>
        [Output("displayName")]
        public Output<string> DisplayName { get; private set; } = null!;

        /// <summary>
        /// A facet specification to perform faceted search.
        /// </summary>
        [Output("facetSpec")]
        public Output<Outputs.GoogleCloudRetailV2alphaSearchRequestFacetSpecResponse> FacetSpec { get; private set; } = null!;

        /// <summary>
        /// Immutable. Fully qualified name projects/*/locations/global/catalogs/*/controls/*
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// A rule control - a condition-action pair. Enacts a set action when the condition is triggered. For example: Boost "gShoe" when query full matches "Running Shoes".
        /// </summary>
        [Output("rule")]
        public Output<Outputs.GoogleCloudRetailV2alphaRuleResponse> Rule { get; private set; } = null!;

        /// <summary>
        /// Immutable. The solution types that the serving config is used for. Currently we support setting only one type of solution at creation time. Only `SOLUTION_TYPE_SEARCH` value is supported at the moment. If no solution type is provided at creation time, will default to SOLUTION_TYPE_SEARCH.
        /// </summary>
        [Output("solutionTypes")]
        public Output<ImmutableArray<string>> SolutionTypes { get; private set; } = null!;


        /// <summary>
        /// Create a Control resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Control(string name, ControlArgs args, CustomResourceOptions? options = null)
            : base("google-native:retail/v2alpha:Control", name, args ?? new ControlArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Control(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:retail/v2alpha:Control", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing Control resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Control Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new Control(name, id, options);
        }
    }

    public sealed class ControlArgs : Pulumi.ResourceArgs
    {
        [Input("catalogId", required: true)]
        public Input<string> CatalogId { get; set; } = null!;

        /// <summary>
        /// Required. The ID to use for the Control, which will become the final component of the Control's resource name. This value should be 4-63 characters, and valid characters are /a-z-_/.
        /// </summary>
        [Input("controlId", required: true)]
        public Input<string> ControlId { get; set; } = null!;

        /// <summary>
        /// The human readable control display name. Used in Retail UI. This field must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is thrown.
        /// </summary>
        [Input("displayName", required: true)]
        public Input<string> DisplayName { get; set; } = null!;

        /// <summary>
        /// A facet specification to perform faceted search.
        /// </summary>
        [Input("facetSpec")]
        public Input<Inputs.GoogleCloudRetailV2alphaSearchRequestFacetSpecArgs>? FacetSpec { get; set; }

        [Input("location")]
        public Input<string>? Location { get; set; }

        /// <summary>
        /// Immutable. Fully qualified name projects/*/locations/global/catalogs/*/controls/*
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("project")]
        public Input<string>? Project { get; set; }

        /// <summary>
        /// A rule control - a condition-action pair. Enacts a set action when the condition is triggered. For example: Boost "gShoe" when query full matches "Running Shoes".
        /// </summary>
        [Input("rule")]
        public Input<Inputs.GoogleCloudRetailV2alphaRuleArgs>? Rule { get; set; }

        [Input("solutionTypes", required: true)]
        private InputList<Pulumi.GoogleNative.Retail.V2Alpha.ControlSolutionTypesItem>? _solutionTypes;

        /// <summary>
        /// Immutable. The solution types that the serving config is used for. Currently we support setting only one type of solution at creation time. Only `SOLUTION_TYPE_SEARCH` value is supported at the moment. If no solution type is provided at creation time, will default to SOLUTION_TYPE_SEARCH.
        /// </summary>
        public InputList<Pulumi.GoogleNative.Retail.V2Alpha.ControlSolutionTypesItem> SolutionTypes
        {
            get => _solutionTypes ?? (_solutionTypes = new InputList<Pulumi.GoogleNative.Retail.V2Alpha.ControlSolutionTypesItem>());
            set => _solutionTypes = value;
        }

        public ControlArgs()
        {
        }
    }
}
