// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Checks the status of a test matrix and the executions once they are created. The test matrix will contain the list of test executions to run if and only if the resultStorage.toolResultsExecution fields have been populated. Note: Flaky test executions may still be added to the matrix at a later stage. May return any of the following canonical error codes: - PERMISSION_DENIED - if the user is not authorized to read project - INVALID_ARGUMENT - if the request is malformed - NOT_FOUND - if the Test Matrix does not exist
func LookupTestMatrix(ctx *pulumi.Context, args *LookupTestMatrixArgs, opts ...pulumi.InvokeOption) (*LookupTestMatrixResult, error) {
	var rv LookupTestMatrixResult
	err := ctx.Invoke("google-native:testing/v1:getTestMatrix", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTestMatrixArgs struct {
	Project      *string `pulumi:"project"`
	TestMatrixId string  `pulumi:"testMatrixId"`
}

type LookupTestMatrixResult struct {
	// Information about the client which invoked the test.
	ClientInfo ClientInfoResponse `pulumi:"clientInfo"`
	// The devices the tests are being executed on.
	EnvironmentMatrix EnvironmentMatrixResponse `pulumi:"environmentMatrix"`
	// If true, only a single attempt at most will be made to run each execution/shard in the matrix. Flaky test attempts are not affected. Normally, 2 or more attempts are made if a potential infrastructure issue is detected. This feature is for latency sensitive workloads. The incidence of execution failures may be significantly greater for fail-fast matrices and support is more limited because of that expectation.
	FailFast bool `pulumi:"failFast"`
	// The number of times a TestExecution should be re-attempted if one or more of its test cases fail for any reason. The maximum number of reruns allowed is 10. Default is 0, which implies no reruns.
	FlakyTestAttempts int `pulumi:"flakyTestAttempts"`
	// Describes why the matrix is considered invalid. Only useful for matrices in the INVALID state.
	InvalidMatrixDetails string `pulumi:"invalidMatrixDetails"`
	// Output Only. The overall outcome of the test. Only set when the test matrix state is FINISHED.
	OutcomeSummary string `pulumi:"outcomeSummary"`
	// The cloud project that owns the test matrix.
	Project string `pulumi:"project"`
	// Where the results for the matrix are written.
	ResultStorage ResultStorageResponse `pulumi:"resultStorage"`
	// Indicates the current progress of the test matrix.
	State string `pulumi:"state"`
	// The list of test executions that the service creates for this matrix.
	TestExecutions []TestExecutionResponse `pulumi:"testExecutions"`
	// Unique id set by the service.
	TestMatrixId string `pulumi:"testMatrixId"`
	// How to run the test.
	TestSpecification TestSpecificationResponse `pulumi:"testSpecification"`
	// The time this test matrix was initially created.
	Timestamp string `pulumi:"timestamp"`
}

func LookupTestMatrixOutput(ctx *pulumi.Context, args LookupTestMatrixOutputArgs, opts ...pulumi.InvokeOption) LookupTestMatrixResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTestMatrixResult, error) {
			args := v.(LookupTestMatrixArgs)
			r, err := LookupTestMatrix(ctx, &args, opts...)
			var s LookupTestMatrixResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTestMatrixResultOutput)
}

type LookupTestMatrixOutputArgs struct {
	Project      pulumi.StringPtrInput `pulumi:"project"`
	TestMatrixId pulumi.StringInput    `pulumi:"testMatrixId"`
}

func (LookupTestMatrixOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTestMatrixArgs)(nil)).Elem()
}

type LookupTestMatrixResultOutput struct{ *pulumi.OutputState }

func (LookupTestMatrixResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTestMatrixResult)(nil)).Elem()
}

func (o LookupTestMatrixResultOutput) ToLookupTestMatrixResultOutput() LookupTestMatrixResultOutput {
	return o
}

func (o LookupTestMatrixResultOutput) ToLookupTestMatrixResultOutputWithContext(ctx context.Context) LookupTestMatrixResultOutput {
	return o
}

// Information about the client which invoked the test.
func (o LookupTestMatrixResultOutput) ClientInfo() ClientInfoResponseOutput {
	return o.ApplyT(func(v LookupTestMatrixResult) ClientInfoResponse { return v.ClientInfo }).(ClientInfoResponseOutput)
}

// The devices the tests are being executed on.
func (o LookupTestMatrixResultOutput) EnvironmentMatrix() EnvironmentMatrixResponseOutput {
	return o.ApplyT(func(v LookupTestMatrixResult) EnvironmentMatrixResponse { return v.EnvironmentMatrix }).(EnvironmentMatrixResponseOutput)
}

// If true, only a single attempt at most will be made to run each execution/shard in the matrix. Flaky test attempts are not affected. Normally, 2 or more attempts are made if a potential infrastructure issue is detected. This feature is for latency sensitive workloads. The incidence of execution failures may be significantly greater for fail-fast matrices and support is more limited because of that expectation.
func (o LookupTestMatrixResultOutput) FailFast() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupTestMatrixResult) bool { return v.FailFast }).(pulumi.BoolOutput)
}

// The number of times a TestExecution should be re-attempted if one or more of its test cases fail for any reason. The maximum number of reruns allowed is 10. Default is 0, which implies no reruns.
func (o LookupTestMatrixResultOutput) FlakyTestAttempts() pulumi.IntOutput {
	return o.ApplyT(func(v LookupTestMatrixResult) int { return v.FlakyTestAttempts }).(pulumi.IntOutput)
}

// Describes why the matrix is considered invalid. Only useful for matrices in the INVALID state.
func (o LookupTestMatrixResultOutput) InvalidMatrixDetails() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestMatrixResult) string { return v.InvalidMatrixDetails }).(pulumi.StringOutput)
}

// Output Only. The overall outcome of the test. Only set when the test matrix state is FINISHED.
func (o LookupTestMatrixResultOutput) OutcomeSummary() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestMatrixResult) string { return v.OutcomeSummary }).(pulumi.StringOutput)
}

// The cloud project that owns the test matrix.
func (o LookupTestMatrixResultOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestMatrixResult) string { return v.Project }).(pulumi.StringOutput)
}

// Where the results for the matrix are written.
func (o LookupTestMatrixResultOutput) ResultStorage() ResultStorageResponseOutput {
	return o.ApplyT(func(v LookupTestMatrixResult) ResultStorageResponse { return v.ResultStorage }).(ResultStorageResponseOutput)
}

// Indicates the current progress of the test matrix.
func (o LookupTestMatrixResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestMatrixResult) string { return v.State }).(pulumi.StringOutput)
}

// The list of test executions that the service creates for this matrix.
func (o LookupTestMatrixResultOutput) TestExecutions() TestExecutionResponseArrayOutput {
	return o.ApplyT(func(v LookupTestMatrixResult) []TestExecutionResponse { return v.TestExecutions }).(TestExecutionResponseArrayOutput)
}

// Unique id set by the service.
func (o LookupTestMatrixResultOutput) TestMatrixId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestMatrixResult) string { return v.TestMatrixId }).(pulumi.StringOutput)
}

// How to run the test.
func (o LookupTestMatrixResultOutput) TestSpecification() TestSpecificationResponseOutput {
	return o.ApplyT(func(v LookupTestMatrixResult) TestSpecificationResponse { return v.TestSpecification }).(TestSpecificationResponseOutput)
}

// The time this test matrix was initially created.
func (o LookupTestMatrixResultOutput) Timestamp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTestMatrixResult) string { return v.Timestamp }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTestMatrixResultOutput{})
}
