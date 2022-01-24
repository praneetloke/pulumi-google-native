// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the batch workload resource representation.
func LookupBatch(ctx *pulumi.Context, args *LookupBatchArgs, opts ...pulumi.InvokeOption) (*LookupBatchResult, error) {
	var rv LookupBatchResult
	err := ctx.Invoke("google-native:dataproc/v1:getBatch", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupBatchArgs struct {
	BatchId  string  `pulumi:"batchId"`
	Location string  `pulumi:"location"`
	Project  *string `pulumi:"project"`
}

type LookupBatchResult struct {
	// The time when the batch was created.
	CreateTime string `pulumi:"createTime"`
	// The email address of the user who created the batch.
	Creator string `pulumi:"creator"`
	// Optional. Environment configuration for the batch execution.
	EnvironmentConfig EnvironmentConfigResponse `pulumi:"environmentConfig"`
	// Optional. The labels to associate with this batch. Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a batch.
	Labels map[string]string `pulumi:"labels"`
	// The resource name of the batch.
	Name string `pulumi:"name"`
	// The resource name of the operation associated with this batch.
	Operation string `pulumi:"operation"`
	// Optional. PySpark batch config.
	PysparkBatch PySparkBatchResponse `pulumi:"pysparkBatch"`
	// Optional. Runtime configuration for the batch execution.
	RuntimeConfig RuntimeConfigResponse `pulumi:"runtimeConfig"`
	// Runtime information about batch execution.
	RuntimeInfo RuntimeInfoResponse `pulumi:"runtimeInfo"`
	// Optional. Spark batch config.
	SparkBatch SparkBatchResponse `pulumi:"sparkBatch"`
	// Optional. SparkR batch config.
	SparkRBatch SparkRBatchResponse `pulumi:"sparkRBatch"`
	// Optional. SparkSql batch config.
	SparkSqlBatch SparkSqlBatchResponse `pulumi:"sparkSqlBatch"`
	// The state of the batch.
	State string `pulumi:"state"`
	// Historical state information for the batch.
	StateHistory []StateHistoryResponse `pulumi:"stateHistory"`
	// Batch state details, such as a failure description if the state is FAILED.
	StateMessage string `pulumi:"stateMessage"`
	// The time when the batch entered a current state.
	StateTime string `pulumi:"stateTime"`
	// A batch UUID (Unique Universal Identifier). The service generates this value when it creates the batch.
	Uuid string `pulumi:"uuid"`
}

func LookupBatchOutput(ctx *pulumi.Context, args LookupBatchOutputArgs, opts ...pulumi.InvokeOption) LookupBatchResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupBatchResult, error) {
			args := v.(LookupBatchArgs)
			r, err := LookupBatch(ctx, &args, opts...)
			return *r, err
		}).(LookupBatchResultOutput)
}

type LookupBatchOutputArgs struct {
	BatchId  pulumi.StringInput    `pulumi:"batchId"`
	Location pulumi.StringInput    `pulumi:"location"`
	Project  pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupBatchOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBatchArgs)(nil)).Elem()
}

type LookupBatchResultOutput struct{ *pulumi.OutputState }

func (LookupBatchResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupBatchResult)(nil)).Elem()
}

func (o LookupBatchResultOutput) ToLookupBatchResultOutput() LookupBatchResultOutput {
	return o
}

func (o LookupBatchResultOutput) ToLookupBatchResultOutputWithContext(ctx context.Context) LookupBatchResultOutput {
	return o
}

// The time when the batch was created.
func (o LookupBatchResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The email address of the user who created the batch.
func (o LookupBatchResultOutput) Creator() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchResult) string { return v.Creator }).(pulumi.StringOutput)
}

// Optional. Environment configuration for the batch execution.
func (o LookupBatchResultOutput) EnvironmentConfig() EnvironmentConfigResponseOutput {
	return o.ApplyT(func(v LookupBatchResult) EnvironmentConfigResponse { return v.EnvironmentConfig }).(EnvironmentConfigResponseOutput)
}

// Optional. The labels to associate with this batch. Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt). No more than 32 labels can be associated with a batch.
func (o LookupBatchResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupBatchResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// The resource name of the batch.
func (o LookupBatchResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchResult) string { return v.Name }).(pulumi.StringOutput)
}

// The resource name of the operation associated with this batch.
func (o LookupBatchResultOutput) Operation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchResult) string { return v.Operation }).(pulumi.StringOutput)
}

// Optional. PySpark batch config.
func (o LookupBatchResultOutput) PysparkBatch() PySparkBatchResponseOutput {
	return o.ApplyT(func(v LookupBatchResult) PySparkBatchResponse { return v.PysparkBatch }).(PySparkBatchResponseOutput)
}

// Optional. Runtime configuration for the batch execution.
func (o LookupBatchResultOutput) RuntimeConfig() RuntimeConfigResponseOutput {
	return o.ApplyT(func(v LookupBatchResult) RuntimeConfigResponse { return v.RuntimeConfig }).(RuntimeConfigResponseOutput)
}

// Runtime information about batch execution.
func (o LookupBatchResultOutput) RuntimeInfo() RuntimeInfoResponseOutput {
	return o.ApplyT(func(v LookupBatchResult) RuntimeInfoResponse { return v.RuntimeInfo }).(RuntimeInfoResponseOutput)
}

// Optional. Spark batch config.
func (o LookupBatchResultOutput) SparkBatch() SparkBatchResponseOutput {
	return o.ApplyT(func(v LookupBatchResult) SparkBatchResponse { return v.SparkBatch }).(SparkBatchResponseOutput)
}

// Optional. SparkR batch config.
func (o LookupBatchResultOutput) SparkRBatch() SparkRBatchResponseOutput {
	return o.ApplyT(func(v LookupBatchResult) SparkRBatchResponse { return v.SparkRBatch }).(SparkRBatchResponseOutput)
}

// Optional. SparkSql batch config.
func (o LookupBatchResultOutput) SparkSqlBatch() SparkSqlBatchResponseOutput {
	return o.ApplyT(func(v LookupBatchResult) SparkSqlBatchResponse { return v.SparkSqlBatch }).(SparkSqlBatchResponseOutput)
}

// The state of the batch.
func (o LookupBatchResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchResult) string { return v.State }).(pulumi.StringOutput)
}

// Historical state information for the batch.
func (o LookupBatchResultOutput) StateHistory() StateHistoryResponseArrayOutput {
	return o.ApplyT(func(v LookupBatchResult) []StateHistoryResponse { return v.StateHistory }).(StateHistoryResponseArrayOutput)
}

// Batch state details, such as a failure description if the state is FAILED.
func (o LookupBatchResultOutput) StateMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchResult) string { return v.StateMessage }).(pulumi.StringOutput)
}

// The time when the batch entered a current state.
func (o LookupBatchResultOutput) StateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchResult) string { return v.StateTime }).(pulumi.StringOutput)
}

// A batch UUID (Unique Universal Identifier). The service generates this value when it creates the batch.
func (o LookupBatchResultOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupBatchResult) string { return v.Uuid }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupBatchResultOutput{})
}
