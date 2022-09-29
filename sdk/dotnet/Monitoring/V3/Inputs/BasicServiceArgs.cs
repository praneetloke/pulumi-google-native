// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Monitoring.V3.Inputs
{

    /// <summary>
    /// A well-known service type, defined by its service type and service labels. Documentation and examples here (https://cloud.google.com/stackdriver/docs/solutions/slo-monitoring/api/api-structures#basic-svc-w-basic-sli).
    /// </summary>
    public sealed class BasicServiceArgs : global::Pulumi.ResourceArgs
    {
        [Input("serviceLabels")]
        private InputMap<string>? _serviceLabels;

        /// <summary>
        /// Labels that specify the resource that emits the monitoring data which is used for SLO reporting of this Service. Documentation and valid values for given service types here (https://cloud.google.com/stackdriver/docs/solutions/slo-monitoring/api/api-structures#basic-svc-w-basic-sli).
        /// </summary>
        public InputMap<string> ServiceLabels
        {
            get => _serviceLabels ?? (_serviceLabels = new InputMap<string>());
            set => _serviceLabels = value;
        }

        /// <summary>
        /// The type of service that this basic service defines, e.g. APP_ENGINE service type. Documentation and valid values here (https://cloud.google.com/stackdriver/docs/solutions/slo-monitoring/api/api-structures#basic-svc-w-basic-sli).
        /// </summary>
        [Input("serviceType")]
        public Input<string>? ServiceType { get; set; }

        public BasicServiceArgs()
        {
        }
        public static new BasicServiceArgs Empty => new BasicServiceArgs();
    }
}
