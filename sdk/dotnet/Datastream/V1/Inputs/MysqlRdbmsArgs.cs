// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Datastream.V1.Inputs
{

    /// <summary>
    /// MySQL database structure
    /// </summary>
    public sealed class MysqlRdbmsArgs : global::Pulumi.ResourceArgs
    {
        [Input("mysqlDatabases")]
        private InputList<Inputs.MysqlDatabaseArgs>? _mysqlDatabases;

        /// <summary>
        /// Mysql databases on the server
        /// </summary>
        public InputList<Inputs.MysqlDatabaseArgs> MysqlDatabases
        {
            get => _mysqlDatabases ?? (_mysqlDatabases = new InputList<Inputs.MysqlDatabaseArgs>());
            set => _mysqlDatabases = value;
        }

        public MysqlRdbmsArgs()
        {
        }
        public static new MysqlRdbmsArgs Empty => new MysqlRdbmsArgs();
    }
}
