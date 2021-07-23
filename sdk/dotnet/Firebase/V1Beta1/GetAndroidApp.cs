// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Firebase.V1Beta1
{
    public static class GetAndroidApp
    {
        /// <summary>
        /// Gets the specified AndroidApp.
        /// </summary>
        public static Task<GetAndroidAppResult> InvokeAsync(GetAndroidAppArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetAndroidAppResult>("google-native:firebase/v1beta1:getAndroidApp", args ?? new GetAndroidAppArgs(), options.WithVersion());
    }


    public sealed class GetAndroidAppArgs : Pulumi.InvokeArgs
    {
        [Input("androidAppId", required: true)]
        public string AndroidAppId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetAndroidAppArgs()
        {
        }
    }


    [OutputType]
    public sealed class GetAndroidAppResult
    {
        /// <summary>
        /// Immutable. The globally unique, Firebase-assigned identifier for the `AndroidApp`. This identifier should be treated as an opaque token, as the data format is not specified.
        /// </summary>
        public readonly string AppId;
        /// <summary>
        /// The user-assigned display name for the `AndroidApp`.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// The resource name of the AndroidApp, in the format: projects/ PROJECT_IDENTIFIER/androidApps/APP_ID * PROJECT_IDENTIFIER: the parent Project's [`ProjectNumber`](../projects#FirebaseProject.FIELDS.project_number) ***(recommended)*** or its [`ProjectId`](../projects#FirebaseProject.FIELDS.project_id). Learn more about using project identifiers in Google's [AIP 2510 standard](https://google.aip.dev/cloud/2510). Note that the value for PROJECT_IDENTIFIER in any response body will be the `ProjectId`. * APP_ID: the globally unique, Firebase-assigned identifier for the App (see [`appId`](../projects.androidApps#AndroidApp.FIELDS.app_id)).
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Immutable. The canonical package name of the Android app as would appear in the Google Play Developer Console.
        /// </summary>
        public readonly string PackageName;
        /// <summary>
        /// Immutable. A user-assigned unique identifier of the parent FirebaseProject for the `AndroidApp`.
        /// </summary>
        public readonly string Project;

        [OutputConstructor]
        private GetAndroidAppResult(
            string appId,

            string displayName,

            string name,

            string packageName,

            string project)
        {
            AppId = appId;
            DisplayName = displayName;
            Name = name;
            PackageName = packageName;
            Project = project;
        }
    }
}