// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.VMMigration.V1Alpha1.Inputs
{

    /// <summary>
    /// AzureSourceDetails message describes a specific source details for the Azure source type.
    /// </summary>
    public sealed class AzureSourceDetailsArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Immutable. The Azure location (region) that the source VMs will be migrated from.
        /// </summary>
        [Input("azureLocation")]
        public Input<string>? AzureLocation { get; set; }

        /// <summary>
        /// Azure Credentials using tenant ID, client ID and secret.
        /// </summary>
        [Input("clientSecretCreds")]
        public Input<Inputs.ClientSecretCredentialsArgs>? ClientSecretCreds { get; set; }

        [Input("migrationResourcesUserTags")]
        private InputMap<string>? _migrationResourcesUserTags;

        /// <summary>
        /// User specified tags to add to every M2VM generated resource in Azure. These tags will be set in addition to the default tags that are set as part of the migration process. The tags must not begin with the reserved prefix `m4ce` or `m2vm`.
        /// </summary>
        public InputMap<string> MigrationResourcesUserTags
        {
            get => _migrationResourcesUserTags ?? (_migrationResourcesUserTags = new InputMap<string>());
            set => _migrationResourcesUserTags = value;
        }

        /// <summary>
        /// Immutable. Azure subscription ID.
        /// </summary>
        [Input("subscriptionId")]
        public Input<string>? SubscriptionId { get; set; }

        public AzureSourceDetailsArgs()
        {
        }
        public static new AzureSourceDetailsArgs Empty => new AzureSourceDetailsArgs();
    }
}