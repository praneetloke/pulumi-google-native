// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Dataproc.V1.Outputs
{

    /// <summary>
    /// Historical state information.
    /// </summary>
    [OutputType]
    public sealed class SessionStateHistoryResponse
    {
        /// <summary>
        /// The state of the session at this point in the session history.
        /// </summary>
        public readonly string State;
        /// <summary>
        /// Details about the state at this point in the session history.
        /// </summary>
        public readonly string StateMessage;
        /// <summary>
        /// The time when the session entered the historical state.
        /// </summary>
        public readonly string StateStartTime;

        [OutputConstructor]
        private SessionStateHistoryResponse(
            string state,

            string stateMessage,

            string stateStartTime)
        {
            State = state;
            StateMessage = stateMessage;
            StateStartTime = stateStartTime;
        }
    }
}
