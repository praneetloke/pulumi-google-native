// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Metastore.V1Alpha.Inputs
{

    /// <summary>
    /// Specifies configuration information specific to running Hive metastore software as the metastore service.
    /// </summary>
    public sealed class HiveMetastoreConfigArgs : Pulumi.ResourceArgs
    {
        [Input("auxiliaryVersions")]
        private InputMap<string>? _auxiliaryVersions;

        /// <summary>
        /// A mapping of Hive metastore version to the auxiliary version configuration. When specified, a secondary Hive metastore service is created along with the primary service. All auxiliary versions must be less than the service's primary version. The key is the auxiliary service name and it must match the regular expression a-z?. This means that the first character must be a lowercase letter, and all the following characters must be hyphens, lowercase letters, or digits, except the last character, which cannot be a hyphen.
        /// </summary>
        public InputMap<string> AuxiliaryVersions
        {
            get => _auxiliaryVersions ?? (_auxiliaryVersions = new InputMap<string>());
            set => _auxiliaryVersions = value;
        }

        [Input("configOverrides")]
        private InputMap<string>? _configOverrides;

        /// <summary>
        /// A mapping of Hive metastore configuration key-value pairs to apply to the Hive metastore (configured in hive-site.xml). The mappings override system defaults (some keys cannot be overridden). These overrides are also applied to auxiliary versions and can be further customized in the auxiliary version's AuxiliaryVersionConfig.
        /// </summary>
        public InputMap<string> ConfigOverrides
        {
            get => _configOverrides ?? (_configOverrides = new InputMap<string>());
            set => _configOverrides = value;
        }

        /// <summary>
        /// The protocol to use for the metastore service endpoint. If unspecified, defaults to THRIFT.
        /// </summary>
        [Input("endpointProtocol")]
        public Input<Pulumi.GoogleNative.Metastore.V1Alpha.HiveMetastoreConfigEndpointProtocol>? EndpointProtocol { get; set; }

        /// <summary>
        /// Information used to configure the Hive metastore service as a service principal in a Kerberos realm. To disable Kerberos, use the UpdateService method and specify this field's path (hive_metastore_config.kerberos_config) in the request's update_mask while omitting this field from the request's service.
        /// </summary>
        [Input("kerberosConfig")]
        public Input<Inputs.KerberosConfigArgs>? KerberosConfig { get; set; }

        /// <summary>
        /// Immutable. The Hive metastore schema version.
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public HiveMetastoreConfigArgs()
        {
        }
    }
}
