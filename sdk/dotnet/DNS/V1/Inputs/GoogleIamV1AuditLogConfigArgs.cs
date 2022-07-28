// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.DNS.V1.Inputs
{

    /// <summary>
    /// Provides the configuration for logging a type of permissions. Example: { "audit_log_configs": [ { "log_type": "DATA_READ", "exempted_members": [ "user:jose@example.com" ] }, { "log_type": "DATA_WRITE" } ] } This enables 'DATA_READ' and 'DATA_WRITE' logging, while exempting jose@example.com from DATA_READ logging.
    /// </summary>
    public sealed class GoogleIamV1AuditLogConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("exemptedMembers")]
        private InputList<string>? _exemptedMembers;

        /// <summary>
        /// Specifies the identities that do not cause logging for this type of permission. Follows the same format of Binding.members.
        /// </summary>
        public InputList<string> ExemptedMembers
        {
            get => _exemptedMembers ?? (_exemptedMembers = new InputList<string>());
            set => _exemptedMembers = value;
        }

        /// <summary>
        /// The log type that this config enables.
        /// </summary>
        [Input("logType")]
        public Input<Pulumi.GoogleNative.DNS.V1.GoogleIamV1AuditLogConfigLogType>? LogType { get; set; }

        public GoogleIamV1AuditLogConfigArgs()
        {
        }
        public static new GoogleIamV1AuditLogConfigArgs Empty => new GoogleIamV1AuditLogConfigArgs();
    }
}
