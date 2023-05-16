// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Compute.Alpha.Outputs
{

    /// <summary>
    /// The state that the future reservation will be reverted to should the amendment be declined.
    /// </summary>
    [OutputType]
    public sealed class FutureReservationStatusLastKnownGoodStateResponse
    {
        /// <summary>
        /// The description of the FutureReservation before an amendment was requested.
        /// </summary>
        public readonly string Description;
        public readonly Outputs.FutureReservationStatusLastKnownGoodStateFutureReservationSpecsResponse FutureReservationSpecs;
        /// <summary>
        /// The name prefix of the Future Reservation before an amendment was requested.
        /// </summary>
        public readonly string NamePrefix;
        /// <summary>
        /// The status of the last known good state for the Future Reservation.
        /// </summary>
        public readonly string ProcurementStatus;

        [OutputConstructor]
        private FutureReservationStatusLastKnownGoodStateResponse(
            string description,

            Outputs.FutureReservationStatusLastKnownGoodStateFutureReservationSpecsResponse futureReservationSpecs,

            string namePrefix,

            string procurementStatus)
        {
            Description = description;
            FutureReservationSpecs = futureReservationSpecs;
            NamePrefix = namePrefix;
            ProcurementStatus = procurementStatus;
        }
    }
}
