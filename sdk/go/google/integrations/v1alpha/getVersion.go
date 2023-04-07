// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Get a integration in the specified project.
func LookupVersion(ctx *pulumi.Context, args *LookupVersionArgs, opts ...pulumi.InvokeOption) (*LookupVersionResult, error) {
	var rv LookupVersionResult
	err := ctx.Invoke("google-native:integrations/v1alpha:getVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVersionArgs struct {
	IntegrationId string  `pulumi:"integrationId"`
	Location      string  `pulumi:"location"`
	ProductId     string  `pulumi:"productId"`
	Project       *string `pulumi:"project"`
	VersionId     string  `pulumi:"versionId"`
}

type LookupVersionResult struct {
	// Auto-generated.
	CreateTime string `pulumi:"createTime"`
	// Optional. Flag to disable database persistence for execution data, including event execution info, execution export info, execution metadata index and execution param index.
	DatabasePersistencePolicy string `pulumi:"databasePersistencePolicy"`
	// Optional. The integration description.
	Description string `pulumi:"description"`
	// Optional. Error Catch Task configuration for the integration. It's optional.
	ErrorCatcherConfigs []GoogleCloudIntegrationsV1alphaErrorCatcherConfigResponse `pulumi:"errorCatcherConfigs"`
	// Optional. Parameters that are expected to be passed to the integration when an event is triggered. This consists of all the parameters that are expected in the integration execution. This gives the user the ability to provide default values, add information like PII and also provide data types of each parameter.
	IntegrationParameters []GoogleCloudIntegrationsV1alphaIntegrationParameterResponse `pulumi:"integrationParameters"`
	// Optional. Parameters that are expected to be passed to the integration when an event is triggered. This consists of all the parameters that are expected in the integration execution. This gives the user the ability to provide default values, add information like PII and also provide data types of each parameter.
	IntegrationParametersInternal EnterpriseCrmFrontendsEventbusProtoWorkflowParametersResponse `pulumi:"integrationParametersInternal"`
	// Optional. The last modifier's email address. Generated based on the End User Credentials/LOAS role of the user making the call.
	LastModifierEmail string `pulumi:"lastModifierEmail"`
	// Optional. The edit lock holder's email address. Generated based on the End User Credentials/LOAS role of the user making the call.
	LockHolder string `pulumi:"lockHolder"`
	// Auto-generated primary key.
	Name string `pulumi:"name"`
	// Optional. The origin that indicates where this integration is coming from.
	Origin string `pulumi:"origin"`
	// Optional. The id of the template which was used to create this integration_version.
	ParentTemplateId string `pulumi:"parentTemplateId"`
	// Optional. The run-as service account email, if set and auth config is not configured, that will be used to generate auth token to be used in Connector task, Rest caller task and Cloud function task.
	RunAsServiceAccount string `pulumi:"runAsServiceAccount"`
	// Optional. An increasing sequence that is set when a new snapshot is created. The last created snapshot can be identified by [workflow_name, org_id latest(snapshot_number)]. However, last created snapshot need not be same as the HEAD. So users should always use "HEAD" tag to identify the head.
	SnapshotNumber string `pulumi:"snapshotNumber"`
	// User should not set it as an input.
	State string `pulumi:"state"`
	// Generated by eventbus. User should not set it as an input.
	Status string `pulumi:"status"`
	// Optional. Task configuration for the integration. It's optional, but the integration doesn't do anything without task_configs.
	TaskConfigs []GoogleCloudIntegrationsV1alphaTaskConfigResponse `pulumi:"taskConfigs"`
	// Optional. Task configuration for the integration. It's optional, but the integration doesn't do anything without task_configs.
	TaskConfigsInternal []EnterpriseCrmFrontendsEventbusProtoTaskConfigResponse `pulumi:"taskConfigsInternal"`
	// Optional. Contains a graph of tasks that will be executed before putting the event in a terminal state (SUCCEEDED/FAILED/FATAL), regardless of success or failure, similar to "finally" in code.
	Teardown EnterpriseCrmEventbusProtoTeardownResponse `pulumi:"teardown"`
	// Optional. Trigger configurations.
	TriggerConfigs []GoogleCloudIntegrationsV1alphaTriggerConfigResponse `pulumi:"triggerConfigs"`
	// Optional. Trigger configurations.
	TriggerConfigsInternal []EnterpriseCrmFrontendsEventbusProtoTriggerConfigResponse `pulumi:"triggerConfigsInternal"`
	// Auto-generated.
	UpdateTime string `pulumi:"updateTime"`
	// Optional. A user-defined label that annotates an integration version. Typically, this is only set when the integration version is created.
	UserLabel string `pulumi:"userLabel"`
}

func LookupVersionOutput(ctx *pulumi.Context, args LookupVersionOutputArgs, opts ...pulumi.InvokeOption) LookupVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVersionResult, error) {
			args := v.(LookupVersionArgs)
			r, err := LookupVersion(ctx, &args, opts...)
			var s LookupVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVersionResultOutput)
}

type LookupVersionOutputArgs struct {
	IntegrationId pulumi.StringInput    `pulumi:"integrationId"`
	Location      pulumi.StringInput    `pulumi:"location"`
	ProductId     pulumi.StringInput    `pulumi:"productId"`
	Project       pulumi.StringPtrInput `pulumi:"project"`
	VersionId     pulumi.StringInput    `pulumi:"versionId"`
}

func (LookupVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVersionArgs)(nil)).Elem()
}

type LookupVersionResultOutput struct{ *pulumi.OutputState }

func (LookupVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVersionResult)(nil)).Elem()
}

func (o LookupVersionResultOutput) ToLookupVersionResultOutput() LookupVersionResultOutput {
	return o
}

func (o LookupVersionResultOutput) ToLookupVersionResultOutputWithContext(ctx context.Context) LookupVersionResultOutput {
	return o
}

// Auto-generated.
func (o LookupVersionResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. Flag to disable database persistence for execution data, including event execution info, execution export info, execution metadata index and execution param index.
func (o LookupVersionResultOutput) DatabasePersistencePolicy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.DatabasePersistencePolicy }).(pulumi.StringOutput)
}

// Optional. The integration description.
func (o LookupVersionResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.Description }).(pulumi.StringOutput)
}

// Optional. Error Catch Task configuration for the integration. It's optional.
func (o LookupVersionResultOutput) ErrorCatcherConfigs() GoogleCloudIntegrationsV1alphaErrorCatcherConfigResponseArrayOutput {
	return o.ApplyT(func(v LookupVersionResult) []GoogleCloudIntegrationsV1alphaErrorCatcherConfigResponse {
		return v.ErrorCatcherConfigs
	}).(GoogleCloudIntegrationsV1alphaErrorCatcherConfigResponseArrayOutput)
}

// Optional. Parameters that are expected to be passed to the integration when an event is triggered. This consists of all the parameters that are expected in the integration execution. This gives the user the ability to provide default values, add information like PII and also provide data types of each parameter.
func (o LookupVersionResultOutput) IntegrationParameters() GoogleCloudIntegrationsV1alphaIntegrationParameterResponseArrayOutput {
	return o.ApplyT(func(v LookupVersionResult) []GoogleCloudIntegrationsV1alphaIntegrationParameterResponse {
		return v.IntegrationParameters
	}).(GoogleCloudIntegrationsV1alphaIntegrationParameterResponseArrayOutput)
}

// Optional. Parameters that are expected to be passed to the integration when an event is triggered. This consists of all the parameters that are expected in the integration execution. This gives the user the ability to provide default values, add information like PII and also provide data types of each parameter.
func (o LookupVersionResultOutput) IntegrationParametersInternal() EnterpriseCrmFrontendsEventbusProtoWorkflowParametersResponseOutput {
	return o.ApplyT(func(v LookupVersionResult) EnterpriseCrmFrontendsEventbusProtoWorkflowParametersResponse {
		return v.IntegrationParametersInternal
	}).(EnterpriseCrmFrontendsEventbusProtoWorkflowParametersResponseOutput)
}

// Optional. The last modifier's email address. Generated based on the End User Credentials/LOAS role of the user making the call.
func (o LookupVersionResultOutput) LastModifierEmail() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.LastModifierEmail }).(pulumi.StringOutput)
}

// Optional. The edit lock holder's email address. Generated based on the End User Credentials/LOAS role of the user making the call.
func (o LookupVersionResultOutput) LockHolder() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.LockHolder }).(pulumi.StringOutput)
}

// Auto-generated primary key.
func (o LookupVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional. The origin that indicates where this integration is coming from.
func (o LookupVersionResultOutput) Origin() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.Origin }).(pulumi.StringOutput)
}

// Optional. The id of the template which was used to create this integration_version.
func (o LookupVersionResultOutput) ParentTemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.ParentTemplateId }).(pulumi.StringOutput)
}

// Optional. The run-as service account email, if set and auth config is not configured, that will be used to generate auth token to be used in Connector task, Rest caller task and Cloud function task.
func (o LookupVersionResultOutput) RunAsServiceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.RunAsServiceAccount }).(pulumi.StringOutput)
}

// Optional. An increasing sequence that is set when a new snapshot is created. The last created snapshot can be identified by [workflow_name, org_id latest(snapshot_number)]. However, last created snapshot need not be same as the HEAD. So users should always use "HEAD" tag to identify the head.
func (o LookupVersionResultOutput) SnapshotNumber() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.SnapshotNumber }).(pulumi.StringOutput)
}

// User should not set it as an input.
func (o LookupVersionResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.State }).(pulumi.StringOutput)
}

// Generated by eventbus. User should not set it as an input.
func (o LookupVersionResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.Status }).(pulumi.StringOutput)
}

// Optional. Task configuration for the integration. It's optional, but the integration doesn't do anything without task_configs.
func (o LookupVersionResultOutput) TaskConfigs() GoogleCloudIntegrationsV1alphaTaskConfigResponseArrayOutput {
	return o.ApplyT(func(v LookupVersionResult) []GoogleCloudIntegrationsV1alphaTaskConfigResponse { return v.TaskConfigs }).(GoogleCloudIntegrationsV1alphaTaskConfigResponseArrayOutput)
}

// Optional. Task configuration for the integration. It's optional, but the integration doesn't do anything without task_configs.
func (o LookupVersionResultOutput) TaskConfigsInternal() EnterpriseCrmFrontendsEventbusProtoTaskConfigResponseArrayOutput {
	return o.ApplyT(func(v LookupVersionResult) []EnterpriseCrmFrontendsEventbusProtoTaskConfigResponse {
		return v.TaskConfigsInternal
	}).(EnterpriseCrmFrontendsEventbusProtoTaskConfigResponseArrayOutput)
}

// Optional. Contains a graph of tasks that will be executed before putting the event in a terminal state (SUCCEEDED/FAILED/FATAL), regardless of success or failure, similar to "finally" in code.
func (o LookupVersionResultOutput) Teardown() EnterpriseCrmEventbusProtoTeardownResponseOutput {
	return o.ApplyT(func(v LookupVersionResult) EnterpriseCrmEventbusProtoTeardownResponse { return v.Teardown }).(EnterpriseCrmEventbusProtoTeardownResponseOutput)
}

// Optional. Trigger configurations.
func (o LookupVersionResultOutput) TriggerConfigs() GoogleCloudIntegrationsV1alphaTriggerConfigResponseArrayOutput {
	return o.ApplyT(func(v LookupVersionResult) []GoogleCloudIntegrationsV1alphaTriggerConfigResponse {
		return v.TriggerConfigs
	}).(GoogleCloudIntegrationsV1alphaTriggerConfigResponseArrayOutput)
}

// Optional. Trigger configurations.
func (o LookupVersionResultOutput) TriggerConfigsInternal() EnterpriseCrmFrontendsEventbusProtoTriggerConfigResponseArrayOutput {
	return o.ApplyT(func(v LookupVersionResult) []EnterpriseCrmFrontendsEventbusProtoTriggerConfigResponse {
		return v.TriggerConfigsInternal
	}).(EnterpriseCrmFrontendsEventbusProtoTriggerConfigResponseArrayOutput)
}

// Auto-generated.
func (o LookupVersionResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

// Optional. A user-defined label that annotates an integration version. Typically, this is only set when the integration version is created.
func (o LookupVersionResultOutput) UserLabel() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVersionResult) string { return v.UserLabel }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVersionResultOutput{})
}
