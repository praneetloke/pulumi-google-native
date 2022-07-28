// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AppEngine.V1
{
    public static class GetApp
    {
        /// <summary>
        /// Gets information about an application.
        /// </summary>
        public static Task<GetAppResult> InvokeAsync(GetAppArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAppResult>("google-native:appengine/v1:getApp", args ?? new GetAppArgs(), options.WithDefaults());

        /// <summary>
        /// Gets information about an application.
        /// </summary>
        public static Output<GetAppResult> Invoke(GetAppInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetAppResult>("google-native:appengine/v1:getApp", args ?? new GetAppInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetAppArgs : global::Pulumi.InvokeArgs
    {
        [Input("appId", required: true)]
        public string AppId { get; set; } = null!;

        public GetAppArgs()
        {
        }
        public static new GetAppArgs Empty => new GetAppArgs();
    }

    public sealed class GetAppInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("appId", required: true)]
        public Input<string> AppId { get; set; } = null!;

        public GetAppInvokeArgs()
        {
        }
        public static new GetAppInvokeArgs Empty => new GetAppInvokeArgs();
    }


    [OutputType]
    public sealed class GetAppResult
    {
        /// <summary>
        /// Google Apps authentication domain that controls which users can access this application.Defaults to open access for any Google Account.
        /// </summary>
        public readonly string AuthDomain;
        /// <summary>
        /// Google Cloud Storage bucket that can be used for storing files associated with this application. This bucket is associated with the application and can be used by the gcloud deployment commands.
        /// </summary>
        public readonly string CodeBucket;
        /// <summary>
        /// The type of the Cloud Firestore or Cloud Datastore database associated with this application.
        /// </summary>
        public readonly string DatabaseType;
        /// <summary>
        /// Google Cloud Storage bucket that can be used by this application to store content.
        /// </summary>
        public readonly string DefaultBucket;
        /// <summary>
        /// Cookie expiration policy for this application.
        /// </summary>
        public readonly string DefaultCookieExpiration;
        /// <summary>
        /// Hostname used to reach this application, as resolved by App Engine.
        /// </summary>
        public readonly string DefaultHostname;
        /// <summary>
        /// HTTP path dispatch rules for requests to the application that do not explicitly target a service or version. Rules are order-dependent. Up to 20 dispatch rules can be supported.
        /// </summary>
        public readonly ImmutableArray<Outputs.UrlDispatchRuleResponse> DispatchRules;
        /// <summary>
        /// The feature specific settings to be used in the application.
        /// </summary>
        public readonly Outputs.FeatureSettingsResponse FeatureSettings;
        /// <summary>
        /// The Google Container Registry domain used for storing managed build docker images for this application.
        /// </summary>
        public readonly string GcrDomain;
        public readonly Outputs.IdentityAwareProxyResponse Iap;
        /// <summary>
        /// Location from which this application runs. Application instances run out of the data centers in the specified location, which is also where all of the application's end user content is stored.Defaults to us-central.View the list of supported locations (https://cloud.google.com/appengine/docs/locations).
        /// </summary>
        public readonly string Location;
        /// <summary>
        /// Full path to the Application resource in the API. Example: apps/myapp.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The service account associated with the application. This is the app-level default identity. If no identity provided during create version, Admin API will fallback to this one.
        /// </summary>
        public readonly string ServiceAccount;
        /// <summary>
        /// Serving status of this application.
        /// </summary>
        public readonly string ServingStatus;

        [OutputConstructor]
        private GetAppResult(
            string authDomain,

            string codeBucket,

            string databaseType,

            string defaultBucket,

            string defaultCookieExpiration,

            string defaultHostname,

            ImmutableArray<Outputs.UrlDispatchRuleResponse> dispatchRules,

            Outputs.FeatureSettingsResponse featureSettings,

            string gcrDomain,

            Outputs.IdentityAwareProxyResponse iap,

            string location,

            string name,

            string serviceAccount,

            string servingStatus)
        {
            AuthDomain = authDomain;
            CodeBucket = codeBucket;
            DatabaseType = databaseType;
            DefaultBucket = defaultBucket;
            DefaultCookieExpiration = defaultCookieExpiration;
            DefaultHostname = defaultHostname;
            DispatchRules = dispatchRules;
            FeatureSettings = featureSettings;
            GcrDomain = gcrDomain;
            Iap = iap;
            Location = location;
            Name = name;
            ServiceAccount = serviceAccount;
            ServingStatus = servingStatus;
        }
    }
}
