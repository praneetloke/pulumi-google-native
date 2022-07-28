// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.ServiceManagement.V1
{
    public static class GetConfig
    {
        /// <summary>
        /// Gets a service configuration (version) for a managed service.
        /// </summary>
        public static Task<GetConfigResult> InvokeAsync(GetConfigArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.InvokeAsync<GetConfigResult>("google-native:servicemanagement/v1:getConfig", args ?? new GetConfigArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a service configuration (version) for a managed service.
        /// </summary>
        public static Output<GetConfigResult> Invoke(GetConfigInvokeArgs args, InvokeOptions? options = null)
            => Pulumi.Deployment.Instance.Invoke<GetConfigResult>("google-native:servicemanagement/v1:getConfig", args ?? new GetConfigInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetConfigArgs : global::Pulumi.InvokeArgs
    {
        [Input("configId", required: true)]
        public string ConfigId { get; set; } = null!;

        [Input("serviceName", required: true)]
        public string ServiceName { get; set; } = null!;

        [Input("view")]
        public string? View { get; set; }

        public GetConfigArgs()
        {
        }
        public static new GetConfigArgs Empty => new GetConfigArgs();
    }

    public sealed class GetConfigInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("configId", required: true)]
        public Input<string> ConfigId { get; set; } = null!;

        [Input("serviceName", required: true)]
        public Input<string> ServiceName { get; set; } = null!;

        [Input("view")]
        public Input<string>? View { get; set; }

        public GetConfigInvokeArgs()
        {
        }
        public static new GetConfigInvokeArgs Empty => new GetConfigInvokeArgs();
    }


    [OutputType]
    public sealed class GetConfigResult
    {
        /// <summary>
        /// A list of API interfaces exported by this service. Only the `name` field of the google.protobuf.Api needs to be provided by the configuration author, as the remaining fields will be derived from the IDL during the normalization process. It is an error to specify an API interface here which cannot be resolved against the associated IDL files.
        /// </summary>
        public readonly ImmutableArray<Outputs.ApiResponse> Apis;
        /// <summary>
        /// Auth configuration.
        /// </summary>
        public readonly Outputs.AuthenticationResponse Authentication;
        /// <summary>
        /// API backend configuration.
        /// </summary>
        public readonly Outputs.BackendResponse Backend;
        /// <summary>
        /// Billing configuration.
        /// </summary>
        public readonly Outputs.BillingResponse Billing;
        /// <summary>
        /// Obsolete. Do not use. This field has no semantic meaning. The service config compiler always sets this field to `3`.
        /// </summary>
        public readonly int ConfigVersion;
        /// <summary>
        /// Context configuration.
        /// </summary>
        public readonly Outputs.ContextResponse Context;
        /// <summary>
        /// Configuration for the service control plane.
        /// </summary>
        public readonly Outputs.ControlResponse Control;
        /// <summary>
        /// Custom error configuration.
        /// </summary>
        public readonly Outputs.CustomErrorResponse CustomError;
        /// <summary>
        /// Additional API documentation.
        /// </summary>
        public readonly Outputs.DocumentationResponse Documentation;
        /// <summary>
        /// Configuration for network endpoints. If this is empty, then an endpoint with the same name as the service is automatically generated to service all defined APIs.
        /// </summary>
        public readonly ImmutableArray<Outputs.EndpointResponse> Endpoints;
        /// <summary>
        /// A list of all enum types included in this API service. Enums referenced directly or indirectly by the `apis` are automatically included. Enums which are not referenced but shall be included should be listed here by name by the configuration author. Example: enums: - name: google.someapi.v1.SomeEnum
        /// </summary>
        public readonly ImmutableArray<Outputs.EnumResponse> Enums;
        /// <summary>
        /// HTTP configuration.
        /// </summary>
        public readonly Outputs.HttpResponse Http;
        /// <summary>
        /// Logging configuration.
        /// </summary>
        public readonly Outputs.LoggingResponse Logging;
        /// <summary>
        /// Defines the logs used by this service.
        /// </summary>
        public readonly ImmutableArray<Outputs.LogDescriptorResponse> Logs;
        /// <summary>
        /// Defines the metrics used by this service.
        /// </summary>
        public readonly ImmutableArray<Outputs.MetricDescriptorResponse> Metrics;
        /// <summary>
        /// Defines the monitored resources used by this service. This is required by the Service.monitoring and Service.logging configurations.
        /// </summary>
        public readonly ImmutableArray<Outputs.MonitoredResourceDescriptorResponse> MonitoredResources;
        /// <summary>
        /// Monitoring configuration.
        /// </summary>
        public readonly Outputs.MonitoringResponse Monitoring;
        /// <summary>
        /// The service name, which is a DNS-like logical identifier for the service, such as `calendar.googleapis.com`. The service name typically goes through DNS verification to make sure the owner of the service also owns the DNS name.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// The Google project that owns this service.
        /// </summary>
        public readonly string ProducerProjectId;
        /// <summary>
        /// Quota configuration.
        /// </summary>
        public readonly Outputs.QuotaResponse Quota;
        /// <summary>
        /// The source information for this configuration if available.
        /// </summary>
        public readonly Outputs.SourceInfoResponse SourceInfo;
        /// <summary>
        /// System parameter configuration.
        /// </summary>
        public readonly Outputs.SystemParametersResponse SystemParameters;
        /// <summary>
        /// A list of all proto message types included in this API service. It serves similar purpose as [google.api.Service.types], except that these types are not needed by user-defined APIs. Therefore, they will not show up in the generated discovery doc. This field should only be used to define system APIs in ESF.
        /// </summary>
        public readonly ImmutableArray<Outputs.TypeResponse> SystemTypes;
        /// <summary>
        /// The product title for this service, it is the name displayed in Google Cloud Console.
        /// </summary>
        public readonly string Title;
        /// <summary>
        /// A list of all proto message types included in this API service. Types referenced directly or indirectly by the `apis` are automatically included. Messages which are not referenced but shall be included, such as types used by the `google.protobuf.Any` type, should be listed here by name by the configuration author. Example: types: - name: google.protobuf.Int32
        /// </summary>
        public readonly ImmutableArray<Outputs.TypeResponse> Types;
        /// <summary>
        /// Configuration controlling usage of this service.
        /// </summary>
        public readonly Outputs.UsageResponse Usage;

        [OutputConstructor]
        private GetConfigResult(
            ImmutableArray<Outputs.ApiResponse> apis,

            Outputs.AuthenticationResponse authentication,

            Outputs.BackendResponse backend,

            Outputs.BillingResponse billing,

            int configVersion,

            Outputs.ContextResponse context,

            Outputs.ControlResponse control,

            Outputs.CustomErrorResponse customError,

            Outputs.DocumentationResponse documentation,

            ImmutableArray<Outputs.EndpointResponse> endpoints,

            ImmutableArray<Outputs.EnumResponse> enums,

            Outputs.HttpResponse http,

            Outputs.LoggingResponse logging,

            ImmutableArray<Outputs.LogDescriptorResponse> logs,

            ImmutableArray<Outputs.MetricDescriptorResponse> metrics,

            ImmutableArray<Outputs.MonitoredResourceDescriptorResponse> monitoredResources,

            Outputs.MonitoringResponse monitoring,

            string name,

            string producerProjectId,

            Outputs.QuotaResponse quota,

            Outputs.SourceInfoResponse sourceInfo,

            Outputs.SystemParametersResponse systemParameters,

            ImmutableArray<Outputs.TypeResponse> systemTypes,

            string title,

            ImmutableArray<Outputs.TypeResponse> types,

            Outputs.UsageResponse usage)
        {
            Apis = apis;
            Authentication = authentication;
            Backend = backend;
            Billing = billing;
            ConfigVersion = configVersion;
            Context = context;
            Control = control;
            CustomError = customError;
            Documentation = documentation;
            Endpoints = endpoints;
            Enums = enums;
            Http = http;
            Logging = logging;
            Logs = logs;
            Metrics = metrics;
            MonitoredResources = monitoredResources;
            Monitoring = monitoring;
            Name = name;
            ProducerProjectId = producerProjectId;
            Quota = quota;
            SourceInfo = sourceInfo;
            SystemParameters = systemParameters;
            SystemTypes = systemTypes;
            Title = title;
            Types = types;
            Usage = usage;
        }
    }
}
