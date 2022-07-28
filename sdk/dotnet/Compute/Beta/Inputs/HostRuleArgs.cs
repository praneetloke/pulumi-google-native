// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Inputs
{

    /// <summary>
    /// UrlMaps A host-matching rule for a URL. If matched, will use the named PathMatcher to select the BackendService.
    /// </summary>
    public sealed class HostRuleArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// An optional description of this resource. Provide this property when you create the resource.
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("hosts")]
        private InputList<string>? _hosts;

        /// <summary>
        /// The list of host patterns to match. They must be valid hostnames with optional port numbers in the format host:port. * matches any string of ([a-z0-9-.]*). In that case, * must be the first character, and if followed by anything, the immediate following character must be either - or .. * based matching is not supported when the URL map is bound to a target gRPC proxy that has the validateForProxyless field set to true.
        /// </summary>
        public InputList<string> Hosts
        {
            get => _hosts ?? (_hosts = new InputList<string>());
            set => _hosts = value;
        }

        /// <summary>
        /// The name of the PathMatcher to use to match the path portion of the URL if the hostRule matches the URL's host portion.
        /// </summary>
        [Input("pathMatcher")]
        public Input<string>? PathMatcher { get; set; }

        public HostRuleArgs()
        {
        }
        public static new HostRuleArgs Empty => new HostRuleArgs();
    }
}
