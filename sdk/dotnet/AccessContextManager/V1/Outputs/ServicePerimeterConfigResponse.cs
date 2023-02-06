// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.AccessContextManager.V1.Outputs
{

    /// <summary>
    /// `ServicePerimeterConfig` specifies a set of Google Cloud resources that describe specific Service Perimeter configuration.
    /// </summary>
    [OutputType]
    public sealed class ServicePerimeterConfigResponse
    {
        /// <summary>
        /// A list of `AccessLevel` resource names that allow resources within the `ServicePerimeter` to be accessed from the internet. `AccessLevels` listed must be in the same policy as this `ServicePerimeter`. Referencing a nonexistent `AccessLevel` is a syntax error. If no `AccessLevel` names are listed, resources within the perimeter can only be accessed via Google Cloud calls with request origins within the perimeter. Example: `"accessPolicies/MY_POLICY/accessLevels/MY_LEVEL"`. For Service Perimeter Bridge, must be empty.
        /// </summary>
        public readonly ImmutableArray<string> AccessLevels;
        /// <summary>
        /// List of EgressPolicies to apply to the perimeter. A perimeter may have multiple EgressPolicies, each of which is evaluated separately. Access is granted if any EgressPolicy grants it. Must be empty for a perimeter bridge.
        /// </summary>
        public readonly ImmutableArray<Outputs.EgressPolicyResponse> EgressPolicies;
        /// <summary>
        /// List of IngressPolicies to apply to the perimeter. A perimeter may have multiple IngressPolicies, each of which is evaluated separately. Access is granted if any Ingress Policy grants it. Must be empty for a perimeter bridge.
        /// </summary>
        public readonly ImmutableArray<Outputs.IngressPolicyResponse> IngressPolicies;
        /// <summary>
        /// A list of Google Cloud resources that are inside of the service perimeter. Currently only projects and VPCs are allowed. Project format: `projects/{project_number}` VPC network format: `//compute.googleapis.com/projects/{PROJECT_ID}/global/networks/{NAME}`.
        /// </summary>
        public readonly ImmutableArray<string> Resources;
        /// <summary>
        /// Google Cloud services that are subject to the Service Perimeter restrictions. For example, if `storage.googleapis.com` is specified, access to the storage buckets inside the perimeter must meet the perimeter's access restrictions.
        /// </summary>
        public readonly ImmutableArray<string> RestrictedServices;
        /// <summary>
        /// Configuration for APIs allowed within Perimeter.
        /// </summary>
        public readonly Outputs.VpcAccessibleServicesResponse VpcAccessibleServices;

        [OutputConstructor]
        private ServicePerimeterConfigResponse(
            ImmutableArray<string> accessLevels,

            ImmutableArray<Outputs.EgressPolicyResponse> egressPolicies,

            ImmutableArray<Outputs.IngressPolicyResponse> ingressPolicies,

            ImmutableArray<string> resources,

            ImmutableArray<string> restrictedServices,

            Outputs.VpcAccessibleServicesResponse vpcAccessibleServices)
        {
            AccessLevels = accessLevels;
            EgressPolicies = egressPolicies;
            IngressPolicies = ingressPolicies;
            Resources = resources;
            RestrictedServices = restrictedServices;
            VpcAccessibleServices = vpcAccessibleServices;
        }
    }
}
