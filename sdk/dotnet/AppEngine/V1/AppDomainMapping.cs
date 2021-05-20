// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AppEngine.V1
{
    /// <summary>
    /// Maps a domain to an application. A user must be authorized to administer a domain in order to map it to an application. For a list of available authorized domains, see AuthorizedDomains.ListAuthorizedDomains.
    /// </summary>
    [GoogleNativeResourceType("google-native:appengine/v1:AppDomainMapping")]
    public partial class AppDomainMapping : Pulumi.CustomResource
    {
        /// <summary>
        /// Full path to the DomainMapping resource in the API. Example: apps/myapp/domainMapping/example.com.@OutputOnly
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// The resource records required to configure this domain mapping. These records must be added to the domain's DNS configuration in order to serve the application via this domain mapping.@OutputOnly
        /// </summary>
        [Output("resourceRecords")]
        public Output<ImmutableArray<Outputs.ResourceRecordResponse>> ResourceRecords { get; private set; } = null!;

        /// <summary>
        /// SSL configuration for this domain. If unconfigured, this domain will not serve with SSL.
        /// </summary>
        [Output("sslSettings")]
        public Output<Outputs.SslSettingsResponse> SslSettings { get; private set; } = null!;


        /// <summary>
        /// Create a AppDomainMapping resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AppDomainMapping(string name, AppDomainMappingArgs args, CustomResourceOptions? options = null)
            : base("google-native:appengine/v1:AppDomainMapping", name, args ?? new AppDomainMappingArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AppDomainMapping(string name, Input<string> id, CustomResourceOptions? options = null)
            : base("google-native:appengine/v1:AppDomainMapping", name, null, MakeResourceOptions(options, id))
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
        /// Get an existing AppDomainMapping resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AppDomainMapping Get(string name, Input<string> id, CustomResourceOptions? options = null)
        {
            return new AppDomainMapping(name, id, options);
        }
    }

    public sealed class AppDomainMappingArgs : Pulumi.ResourceArgs
    {
        [Input("appId", required: true)]
        public Input<string> AppId { get; set; } = null!;

        [Input("domainMappingId", required: true)]
        public Input<string> DomainMappingId { get; set; } = null!;

        /// <summary>
        /// Relative name of the domain serving the application. Example: example.com.
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Full path to the DomainMapping resource in the API. Example: apps/myapp/domainMapping/example.com.@OutputOnly
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("overrideStrategy")]
        public Input<string>? OverrideStrategy { get; set; }

        [Input("resourceRecords")]
        private InputList<Inputs.ResourceRecordArgs>? _resourceRecords;

        /// <summary>
        /// The resource records required to configure this domain mapping. These records must be added to the domain's DNS configuration in order to serve the application via this domain mapping.@OutputOnly
        /// </summary>
        public InputList<Inputs.ResourceRecordArgs> ResourceRecords
        {
            get => _resourceRecords ?? (_resourceRecords = new InputList<Inputs.ResourceRecordArgs>());
            set => _resourceRecords = value;
        }

        /// <summary>
        /// SSL configuration for this domain. If unconfigured, this domain will not serve with SSL.
        /// </summary>
        [Input("sslSettings")]
        public Input<Inputs.SslSettingsArgs>? SslSettings { get; set; }

        public AppDomainMappingArgs()
        {
        }
    }
}
