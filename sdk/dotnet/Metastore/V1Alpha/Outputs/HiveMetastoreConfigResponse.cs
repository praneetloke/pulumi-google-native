// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Metastore.V1Alpha.Outputs
{

    /// <summary>
    /// Specifies configuration information specific to running Hive metastore software as the metastore service.
    /// </summary>
    [OutputType]
    public sealed class HiveMetastoreConfigResponse
    {
        /// <summary>
        /// A mapping of Hive metastore configuration key-value pairs to apply to the Hive metastore (configured in hive-site.xml). The mappings override system defaults (some keys cannot be overridden). These overrides are also applied to auxiliary versions and can be further customized in the auxiliary version's AuxiliaryVersionConfig.
        /// </summary>
        public readonly ImmutableDictionary<string, string> ConfigOverrides;
        /// <summary>
        /// The protocol to use for the metastore service endpoint. If unspecified, defaults to THRIFT.
        /// </summary>
        public readonly string EndpointProtocol;
        /// <summary>
        /// Information used to configure the Hive metastore service as a service principal in a Kerberos realm. To disable Kerberos, use the UpdateService method and specify this field's path (hive_metastore_config.kerberos_config) in the request's update_mask while omitting this field from the request's service.
        /// </summary>
        public readonly Outputs.KerberosConfigResponse KerberosConfig;
        /// <summary>
        /// Immutable. The Hive metastore schema version.
        /// </summary>
        public readonly string Version;

        [OutputConstructor]
        private HiveMetastoreConfigResponse(
            ImmutableDictionary<string, string> configOverrides,

            string endpointProtocol,

            Outputs.KerberosConfigResponse kerberosConfig,

            string version)
        {
            ConfigOverrides = configOverrides;
            EndpointProtocol = endpointProtocol;
            KerberosConfig = kerberosConfig;
            Version = version;
        }
    }
}
