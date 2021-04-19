// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new budget. See Quotas and limits for more information on the limits of the number of budgets you can create.
type BillingAccountBudget struct {
	pulumi.CustomResourceState

	// Optional. Rules to apply to notifications sent based on budget spend and thresholds.
	AllUpdatesRule GoogleCloudBillingBudgetsV1beta1AllUpdatesRuleResponseOutput `pulumi:"allUpdatesRule"`
	// Required. Budgeted amount.
	Amount GoogleCloudBillingBudgetsV1beta1BudgetAmountResponseOutput `pulumi:"amount"`
	// Optional. Filters that define which resources are used to compute the actual spend against the budget amount, such as projects, services, and the budget's time period, as well as other filters.
	BudgetFilter GoogleCloudBillingBudgetsV1beta1FilterResponseOutput `pulumi:"budgetFilter"`
	// User data for display name in UI. Validation: <= 60 chars.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Optional. Etag to validate that the object is unchanged for a read-modify-write operation. An empty etag will cause an update to overwrite other changes.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Resource name of the budget. The resource name implies the scope of a budget. Values are of the form `billingAccounts/{billingAccountId}/budgets/{budgetId}`.
	Name pulumi.StringOutput `pulumi:"name"`
	// Optional. Rules that trigger alerts (notifications of thresholds being crossed) when spend exceeds the specified percentages of the budget.
	ThresholdRules GoogleCloudBillingBudgetsV1beta1ThresholdRuleResponseArrayOutput `pulumi:"thresholdRules"`
}

// NewBillingAccountBudget registers a new resource with the given unique name, arguments, and options.
func NewBillingAccountBudget(ctx *pulumi.Context,
	name string, args *BillingAccountBudgetArgs, opts ...pulumi.ResourceOption) (*BillingAccountBudget, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BillingAccountsId == nil {
		return nil, errors.New("invalid value for required argument 'BillingAccountsId'")
	}
	if args.BudgetsId == nil {
		return nil, errors.New("invalid value for required argument 'BudgetsId'")
	}
	var resource BillingAccountBudget
	err := ctx.RegisterResource("google-native:billingbudgets/v1beta1:BillingAccountBudget", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBillingAccountBudget gets an existing BillingAccountBudget resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBillingAccountBudget(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BillingAccountBudgetState, opts ...pulumi.ResourceOption) (*BillingAccountBudget, error) {
	var resource BillingAccountBudget
	err := ctx.ReadResource("google-native:billingbudgets/v1beta1:BillingAccountBudget", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering BillingAccountBudget resources.
type billingAccountBudgetState struct {
	// Optional. Rules to apply to notifications sent based on budget spend and thresholds.
	AllUpdatesRule *GoogleCloudBillingBudgetsV1beta1AllUpdatesRuleResponse `pulumi:"allUpdatesRule"`
	// Required. Budgeted amount.
	Amount *GoogleCloudBillingBudgetsV1beta1BudgetAmountResponse `pulumi:"amount"`
	// Optional. Filters that define which resources are used to compute the actual spend against the budget amount, such as projects, services, and the budget's time period, as well as other filters.
	BudgetFilter *GoogleCloudBillingBudgetsV1beta1FilterResponse `pulumi:"budgetFilter"`
	// User data for display name in UI. Validation: <= 60 chars.
	DisplayName *string `pulumi:"displayName"`
	// Optional. Etag to validate that the object is unchanged for a read-modify-write operation. An empty etag will cause an update to overwrite other changes.
	Etag *string `pulumi:"etag"`
	// Resource name of the budget. The resource name implies the scope of a budget. Values are of the form `billingAccounts/{billingAccountId}/budgets/{budgetId}`.
	Name *string `pulumi:"name"`
	// Optional. Rules that trigger alerts (notifications of thresholds being crossed) when spend exceeds the specified percentages of the budget.
	ThresholdRules []GoogleCloudBillingBudgetsV1beta1ThresholdRuleResponse `pulumi:"thresholdRules"`
}

type BillingAccountBudgetState struct {
	// Optional. Rules to apply to notifications sent based on budget spend and thresholds.
	AllUpdatesRule GoogleCloudBillingBudgetsV1beta1AllUpdatesRuleResponsePtrInput
	// Required. Budgeted amount.
	Amount GoogleCloudBillingBudgetsV1beta1BudgetAmountResponsePtrInput
	// Optional. Filters that define which resources are used to compute the actual spend against the budget amount, such as projects, services, and the budget's time period, as well as other filters.
	BudgetFilter GoogleCloudBillingBudgetsV1beta1FilterResponsePtrInput
	// User data for display name in UI. Validation: <= 60 chars.
	DisplayName pulumi.StringPtrInput
	// Optional. Etag to validate that the object is unchanged for a read-modify-write operation. An empty etag will cause an update to overwrite other changes.
	Etag pulumi.StringPtrInput
	// Resource name of the budget. The resource name implies the scope of a budget. Values are of the form `billingAccounts/{billingAccountId}/budgets/{budgetId}`.
	Name pulumi.StringPtrInput
	// Optional. Rules that trigger alerts (notifications of thresholds being crossed) when spend exceeds the specified percentages of the budget.
	ThresholdRules GoogleCloudBillingBudgetsV1beta1ThresholdRuleResponseArrayInput
}

func (BillingAccountBudgetState) ElementType() reflect.Type {
	return reflect.TypeOf((*billingAccountBudgetState)(nil)).Elem()
}

type billingAccountBudgetArgs struct {
	// Optional. Rules to apply to notifications sent based on budget spend and thresholds.
	AllUpdatesRule *GoogleCloudBillingBudgetsV1beta1AllUpdatesRule `pulumi:"allUpdatesRule"`
	// Required. Budgeted amount.
	Amount            *GoogleCloudBillingBudgetsV1beta1BudgetAmount `pulumi:"amount"`
	BillingAccountsId string                                        `pulumi:"billingAccountsId"`
	// Optional. Filters that define which resources are used to compute the actual spend against the budget amount, such as projects, services, and the budget's time period, as well as other filters.
	BudgetFilter *GoogleCloudBillingBudgetsV1beta1Filter `pulumi:"budgetFilter"`
	BudgetsId    string                                  `pulumi:"budgetsId"`
	// User data for display name in UI. Validation: <= 60 chars.
	DisplayName *string `pulumi:"displayName"`
	// Optional. Etag to validate that the object is unchanged for a read-modify-write operation. An empty etag will cause an update to overwrite other changes.
	Etag *string `pulumi:"etag"`
	// Optional. Rules that trigger alerts (notifications of thresholds being crossed) when spend exceeds the specified percentages of the budget.
	ThresholdRules []GoogleCloudBillingBudgetsV1beta1ThresholdRule `pulumi:"thresholdRules"`
}

// The set of arguments for constructing a BillingAccountBudget resource.
type BillingAccountBudgetArgs struct {
	// Optional. Rules to apply to notifications sent based on budget spend and thresholds.
	AllUpdatesRule GoogleCloudBillingBudgetsV1beta1AllUpdatesRulePtrInput
	// Required. Budgeted amount.
	Amount            GoogleCloudBillingBudgetsV1beta1BudgetAmountPtrInput
	BillingAccountsId pulumi.StringInput
	// Optional. Filters that define which resources are used to compute the actual spend against the budget amount, such as projects, services, and the budget's time period, as well as other filters.
	BudgetFilter GoogleCloudBillingBudgetsV1beta1FilterPtrInput
	BudgetsId    pulumi.StringInput
	// User data for display name in UI. Validation: <= 60 chars.
	DisplayName pulumi.StringPtrInput
	// Optional. Etag to validate that the object is unchanged for a read-modify-write operation. An empty etag will cause an update to overwrite other changes.
	Etag pulumi.StringPtrInput
	// Optional. Rules that trigger alerts (notifications of thresholds being crossed) when spend exceeds the specified percentages of the budget.
	ThresholdRules GoogleCloudBillingBudgetsV1beta1ThresholdRuleArrayInput
}

func (BillingAccountBudgetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*billingAccountBudgetArgs)(nil)).Elem()
}

type BillingAccountBudgetInput interface {
	pulumi.Input

	ToBillingAccountBudgetOutput() BillingAccountBudgetOutput
	ToBillingAccountBudgetOutputWithContext(ctx context.Context) BillingAccountBudgetOutput
}

func (*BillingAccountBudget) ElementType() reflect.Type {
	return reflect.TypeOf((*BillingAccountBudget)(nil))
}

func (i *BillingAccountBudget) ToBillingAccountBudgetOutput() BillingAccountBudgetOutput {
	return i.ToBillingAccountBudgetOutputWithContext(context.Background())
}

func (i *BillingAccountBudget) ToBillingAccountBudgetOutputWithContext(ctx context.Context) BillingAccountBudgetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BillingAccountBudgetOutput)
}

type BillingAccountBudgetOutput struct {
	*pulumi.OutputState
}

func (BillingAccountBudgetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BillingAccountBudget)(nil))
}

func (o BillingAccountBudgetOutput) ToBillingAccountBudgetOutput() BillingAccountBudgetOutput {
	return o
}

func (o BillingAccountBudgetOutput) ToBillingAccountBudgetOutputWithContext(ctx context.Context) BillingAccountBudgetOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(BillingAccountBudgetOutput{})
}
