// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Integrations.V1Alpha.Outputs
{

    /// <summary>
    /// LINT.IfChange This is the frontend version of WorkflowParameters. It's exactly like the backend version except that instead of flattening protobuf parameters and treating every field and subfield of a protobuf parameter as a separate parameter, the fields/subfields of a protobuf parameter will be nested as "children" (see 'children' field below) parameters of the parent parameter. Please refer to enterprise/crm/eventbus/proto/workflow_parameters.proto for more information about WorkflowParameters.
    /// </summary>
    [OutputType]
    public sealed class EnterpriseCrmFrontendsEventbusProtoWorkflowParametersResponse
    {
        /// <summary>
        /// Parameters are a part of Event and can be used to communiticate between different tasks that are part of the same workflow execution.
        /// </summary>
        public readonly ImmutableArray<Outputs.EnterpriseCrmFrontendsEventbusProtoWorkflowParameterEntryResponse> Parameters;

        [OutputConstructor]
        private EnterpriseCrmFrontendsEventbusProtoWorkflowParametersResponse(ImmutableArray<Outputs.EnterpriseCrmFrontendsEventbusProtoWorkflowParameterEntryResponse> parameters)
        {
            Parameters = parameters;
        }
    }
}
