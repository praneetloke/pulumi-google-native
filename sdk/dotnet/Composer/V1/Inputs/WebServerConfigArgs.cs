// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Composer.V1.Inputs
{

    /// <summary>
    /// The configuration settings for the Airflow web server App Engine instance. Supported for Cloud Composer environments in versions composer-1.*.*-airflow-*.*.*
    /// </summary>
    public sealed class WebServerConfigArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional. Machine type on which Airflow web server is running. It has to be one of: composer-n1-webserver-2, composer-n1-webserver-4 or composer-n1-webserver-8. If not specified, composer-n1-webserver-2 will be used. Value custom is returned only in response, if Airflow web server parameters were manually changed to a non-standard values.
        /// </summary>
        [Input("machineType")]
        public Input<string>? MachineType { get; set; }

        public WebServerConfigArgs()
        {
        }
        public static new WebServerConfigArgs Empty => new WebServerConfigArgs();
    }
}
