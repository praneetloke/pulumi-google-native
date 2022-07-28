// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Beta.Inputs
{

    /// <summary>
    /// Initial State for shielded instance, these are public keys which are safe to store in public
    /// </summary>
    public sealed class InitialStateConfigArgs : global::Pulumi.ResourceArgs
    {
        [Input("dbs")]
        private InputList<Inputs.FileContentBufferArgs>? _dbs;

        /// <summary>
        /// The Key Database (db).
        /// </summary>
        public InputList<Inputs.FileContentBufferArgs> Dbs
        {
            get => _dbs ?? (_dbs = new InputList<Inputs.FileContentBufferArgs>());
            set => _dbs = value;
        }

        [Input("dbxs")]
        private InputList<Inputs.FileContentBufferArgs>? _dbxs;

        /// <summary>
        /// The forbidden key database (dbx).
        /// </summary>
        public InputList<Inputs.FileContentBufferArgs> Dbxs
        {
            get => _dbxs ?? (_dbxs = new InputList<Inputs.FileContentBufferArgs>());
            set => _dbxs = value;
        }

        [Input("keks")]
        private InputList<Inputs.FileContentBufferArgs>? _keks;

        /// <summary>
        /// The Key Exchange Key (KEK).
        /// </summary>
        public InputList<Inputs.FileContentBufferArgs> Keks
        {
            get => _keks ?? (_keks = new InputList<Inputs.FileContentBufferArgs>());
            set => _keks = value;
        }

        /// <summary>
        /// The Platform Key (PK).
        /// </summary>
        [Input("pk")]
        public Input<Inputs.FileContentBufferArgs>? Pk { get; set; }

        public InitialStateConfigArgs()
        {
        }
        public static new InitialStateConfigArgs Empty => new InitialStateConfigArgs();
    }
}
