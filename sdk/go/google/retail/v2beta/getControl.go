// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a Control.
func LookupControl(ctx *pulumi.Context, args *LookupControlArgs, opts ...pulumi.InvokeOption) (*LookupControlResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupControlResult
	err := ctx.Invoke("google-native:retail/v2beta:getControl", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupControlArgs struct {
	CatalogId string  `pulumi:"catalogId"`
	ControlId string  `pulumi:"controlId"`
	Location  string  `pulumi:"location"`
	Project   *string `pulumi:"project"`
}

type LookupControlResult struct {
	// List of serving config ids that are associated with this control in the same Catalog. Note the association is managed via the ServingConfig, this is an output only denormalized view.
	AssociatedServingConfigIds []string `pulumi:"associatedServingConfigIds"`
	// The human readable control display name. Used in Retail UI. This field must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is thrown.
	DisplayName string `pulumi:"displayName"`
	// A facet specification to perform faceted search. Note that this field is deprecated and will throw NOT_IMPLEMENTED if used for creating a control.
	FacetSpec GoogleCloudRetailV2betaSearchRequestFacetSpecResponse `pulumi:"facetSpec"`
	// Immutable. Fully qualified name `projects/*/locations/global/catalogs/*/controls/*`
	Name string `pulumi:"name"`
	// A rule control - a condition-action pair. Enacts a set action when the condition is triggered. For example: Boost "gShoe" when query full matches "Running Shoes".
	Rule GoogleCloudRetailV2betaRuleResponse `pulumi:"rule"`
	// Specifies the use case for the control. Affects what condition fields can be set. Only settable by search controls. Will default to SEARCH_SOLUTION_USE_CASE_SEARCH if not specified. Currently only allow one search_solution_use_case per control.
	SearchSolutionUseCase []string `pulumi:"searchSolutionUseCase"`
	// Immutable. The solution types that the control is used for. Currently we support setting only one type of solution at creation time. Only `SOLUTION_TYPE_SEARCH` value is supported at the moment. If no solution type is provided at creation time, will default to SOLUTION_TYPE_SEARCH.
	SolutionTypes []string `pulumi:"solutionTypes"`
}

func LookupControlOutput(ctx *pulumi.Context, args LookupControlOutputArgs, opts ...pulumi.InvokeOption) LookupControlResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupControlResult, error) {
			args := v.(LookupControlArgs)
			r, err := LookupControl(ctx, &args, opts...)
			var s LookupControlResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupControlResultOutput)
}

type LookupControlOutputArgs struct {
	CatalogId pulumi.StringInput    `pulumi:"catalogId"`
	ControlId pulumi.StringInput    `pulumi:"controlId"`
	Location  pulumi.StringInput    `pulumi:"location"`
	Project   pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupControlOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupControlArgs)(nil)).Elem()
}

type LookupControlResultOutput struct{ *pulumi.OutputState }

func (LookupControlResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupControlResult)(nil)).Elem()
}

func (o LookupControlResultOutput) ToLookupControlResultOutput() LookupControlResultOutput {
	return o
}

func (o LookupControlResultOutput) ToLookupControlResultOutputWithContext(ctx context.Context) LookupControlResultOutput {
	return o
}

// List of serving config ids that are associated with this control in the same Catalog. Note the association is managed via the ServingConfig, this is an output only denormalized view.
func (o LookupControlResultOutput) AssociatedServingConfigIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupControlResult) []string { return v.AssociatedServingConfigIds }).(pulumi.StringArrayOutput)
}

// The human readable control display name. Used in Retail UI. This field must be a UTF-8 encoded string with a length limit of 128 characters. Otherwise, an INVALID_ARGUMENT error is thrown.
func (o LookupControlResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControlResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// A facet specification to perform faceted search. Note that this field is deprecated and will throw NOT_IMPLEMENTED if used for creating a control.
func (o LookupControlResultOutput) FacetSpec() GoogleCloudRetailV2betaSearchRequestFacetSpecResponseOutput {
	return o.ApplyT(func(v LookupControlResult) GoogleCloudRetailV2betaSearchRequestFacetSpecResponse { return v.FacetSpec }).(GoogleCloudRetailV2betaSearchRequestFacetSpecResponseOutput)
}

// Immutable. Fully qualified name `projects/*/locations/global/catalogs/*/controls/*`
func (o LookupControlResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupControlResult) string { return v.Name }).(pulumi.StringOutput)
}

// A rule control - a condition-action pair. Enacts a set action when the condition is triggered. For example: Boost "gShoe" when query full matches "Running Shoes".
func (o LookupControlResultOutput) Rule() GoogleCloudRetailV2betaRuleResponseOutput {
	return o.ApplyT(func(v LookupControlResult) GoogleCloudRetailV2betaRuleResponse { return v.Rule }).(GoogleCloudRetailV2betaRuleResponseOutput)
}

// Specifies the use case for the control. Affects what condition fields can be set. Only settable by search controls. Will default to SEARCH_SOLUTION_USE_CASE_SEARCH if not specified. Currently only allow one search_solution_use_case per control.
func (o LookupControlResultOutput) SearchSolutionUseCase() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupControlResult) []string { return v.SearchSolutionUseCase }).(pulumi.StringArrayOutput)
}

// Immutable. The solution types that the control is used for. Currently we support setting only one type of solution at creation time. Only `SOLUTION_TYPE_SEARCH` value is supported at the moment. If no solution type is provided at creation time, will default to SOLUTION_TYPE_SEARCH.
func (o LookupControlResultOutput) SolutionTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupControlResult) []string { return v.SolutionTypes }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupControlResultOutput{})
}
