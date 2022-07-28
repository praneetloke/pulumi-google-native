// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Ml.V1.Inputs
{

    /// <summary>
    /// Specifies the audit configuration for a service. The configuration determines which permission types are logged, and what identities, if any, are exempted from logging. An AuditConfig must have one or more AuditLogConfigs. If there are AuditConfigs for both `allServices` and a specific service, the union of the two AuditConfigs is used for that service: the log_types specified in each AuditConfig are enabled, and the exempted_members in each AuditLogConfig are exempted. Example Policy with multiple AuditConfigs: { "audit_configs": [ { "service": "allServices", "audit_log_configs": [ { "log_type": "DATA_READ", "exempted_members": [ "user:jose@example.com" ] }, { "log_type": "DATA_WRITE" }, { "log_type": "ADMIN_READ" } ] }, { "service": "sampleservice.googleapis.com", "audit_log_configs": [ { "log_type": "DATA_READ" }, { "log_type": "DATA_WRITE", "exempted_members": [ "user:aliya@example.com" ] } ] } ] } For sampleservice, this policy enables DATA_READ, DATA_WRITE and ADMIN_READ logging. It also exempts `jose@example.com` from DATA_READ logging, and `aliya@example.com` from DATA_WRITE logging.
    /// </summary>
    public sealed class GoogleIamV1__AuditConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("auditLogConfigs")]
        private InputList<Inputs.GoogleIamV1__AuditLogConfigArgs>? _auditLogConfigs;

        /// <summary>
        /// The configuration for logging of each type of permission.
        /// </summary>
        public InputList<Inputs.GoogleIamV1__AuditLogConfigArgs> AuditLogConfigs
        {
            get => _auditLogConfigs ?? (_auditLogConfigs = new InputList<Inputs.GoogleIamV1__AuditLogConfigArgs>());
            set => _auditLogConfigs = value;
        }

        /// <summary>
        /// Specifies a service that will be enabled for audit logging. For example, `storage.googleapis.com`, `cloudsql.googleapis.com`. `allServices` is a special value that covers all services.
        /// </summary>
        [Input("service")]
        public Input<string>? Service { get; set; }

        public GoogleIamV1__AuditConfigArgs()
        {
        }
        public static new GoogleIamV1__AuditConfigArgs Empty => new GoogleIamV1__AuditConfigArgs();
    }
}
