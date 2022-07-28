// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.OSConfig.V1.Inputs
{

    /// <summary>
    /// A package managed by YUM. - install: `yum -y install package` - remove: `yum -y remove package`
    /// </summary>
    public sealed class OSPolicyResourcePackageResourceYUMArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Package name.
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public OSPolicyResourcePackageResourceYUMArgs()
        {
        }
        public static new OSPolicyResourcePackageResourceYUMArgs Empty => new OSPolicyResourcePackageResourceYUMArgs();
    }
}
