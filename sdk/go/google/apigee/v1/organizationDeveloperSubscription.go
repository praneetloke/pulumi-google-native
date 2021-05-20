// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a subscription to an API product.
type OrganizationDeveloperSubscription struct {
	pulumi.CustomResourceState

	// Name of the API product for which the developer is purchasing a subscription.
	Apiproduct pulumi.StringOutput `pulumi:"apiproduct"`
	// Time when the API product subscription was created in milliseconds since epoch.
	CreatedAt pulumi.StringOutput `pulumi:"createdAt"`
	// Time when the API product subscription ends in milliseconds since epoch.
	EndTime pulumi.StringOutput `pulumi:"endTime"`
	// Time when the API product subscription was last modified in milliseconds since epoch.
	LastModifiedAt pulumi.StringOutput `pulumi:"lastModifiedAt"`
	// Name of the API product subscription.
	Name pulumi.StringOutput `pulumi:"name"`
	// Time when the API product subscription starts in milliseconds since epoch.
	StartTime pulumi.StringOutput `pulumi:"startTime"`
}

// NewOrganizationDeveloperSubscription registers a new resource with the given unique name, arguments, and options.
func NewOrganizationDeveloperSubscription(ctx *pulumi.Context,
	name string, args *OrganizationDeveloperSubscriptionArgs, opts ...pulumi.ResourceOption) (*OrganizationDeveloperSubscription, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeveloperId == nil {
		return nil, errors.New("invalid value for required argument 'DeveloperId'")
	}
	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	if args.SubscriptionId == nil {
		return nil, errors.New("invalid value for required argument 'SubscriptionId'")
	}
	var resource OrganizationDeveloperSubscription
	err := ctx.RegisterResource("google-native:apigee/v1:OrganizationDeveloperSubscription", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationDeveloperSubscription gets an existing OrganizationDeveloperSubscription resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationDeveloperSubscription(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationDeveloperSubscriptionState, opts ...pulumi.ResourceOption) (*OrganizationDeveloperSubscription, error) {
	var resource OrganizationDeveloperSubscription
	err := ctx.ReadResource("google-native:apigee/v1:OrganizationDeveloperSubscription", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationDeveloperSubscription resources.
type organizationDeveloperSubscriptionState struct {
	// Name of the API product for which the developer is purchasing a subscription.
	Apiproduct *string `pulumi:"apiproduct"`
	// Time when the API product subscription was created in milliseconds since epoch.
	CreatedAt *string `pulumi:"createdAt"`
	// Time when the API product subscription ends in milliseconds since epoch.
	EndTime *string `pulumi:"endTime"`
	// Time when the API product subscription was last modified in milliseconds since epoch.
	LastModifiedAt *string `pulumi:"lastModifiedAt"`
	// Name of the API product subscription.
	Name *string `pulumi:"name"`
	// Time when the API product subscription starts in milliseconds since epoch.
	StartTime *string `pulumi:"startTime"`
}

type OrganizationDeveloperSubscriptionState struct {
	// Name of the API product for which the developer is purchasing a subscription.
	Apiproduct pulumi.StringPtrInput
	// Time when the API product subscription was created in milliseconds since epoch.
	CreatedAt pulumi.StringPtrInput
	// Time when the API product subscription ends in milliseconds since epoch.
	EndTime pulumi.StringPtrInput
	// Time when the API product subscription was last modified in milliseconds since epoch.
	LastModifiedAt pulumi.StringPtrInput
	// Name of the API product subscription.
	Name pulumi.StringPtrInput
	// Time when the API product subscription starts in milliseconds since epoch.
	StartTime pulumi.StringPtrInput
}

func (OrganizationDeveloperSubscriptionState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationDeveloperSubscriptionState)(nil)).Elem()
}

type organizationDeveloperSubscriptionArgs struct {
	// Name of the API product for which the developer is purchasing a subscription.
	Apiproduct  *string `pulumi:"apiproduct"`
	DeveloperId string  `pulumi:"developerId"`
	// Time when the API product subscription ends in milliseconds since epoch.
	EndTime        *string `pulumi:"endTime"`
	OrganizationId string  `pulumi:"organizationId"`
	// Time when the API product subscription starts in milliseconds since epoch.
	StartTime      *string `pulumi:"startTime"`
	SubscriptionId string  `pulumi:"subscriptionId"`
}

// The set of arguments for constructing a OrganizationDeveloperSubscription resource.
type OrganizationDeveloperSubscriptionArgs struct {
	// Name of the API product for which the developer is purchasing a subscription.
	Apiproduct  pulumi.StringPtrInput
	DeveloperId pulumi.StringInput
	// Time when the API product subscription ends in milliseconds since epoch.
	EndTime        pulumi.StringPtrInput
	OrganizationId pulumi.StringInput
	// Time when the API product subscription starts in milliseconds since epoch.
	StartTime      pulumi.StringPtrInput
	SubscriptionId pulumi.StringInput
}

func (OrganizationDeveloperSubscriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationDeveloperSubscriptionArgs)(nil)).Elem()
}

type OrganizationDeveloperSubscriptionInput interface {
	pulumi.Input

	ToOrganizationDeveloperSubscriptionOutput() OrganizationDeveloperSubscriptionOutput
	ToOrganizationDeveloperSubscriptionOutputWithContext(ctx context.Context) OrganizationDeveloperSubscriptionOutput
}

func (*OrganizationDeveloperSubscription) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationDeveloperSubscription)(nil))
}

func (i *OrganizationDeveloperSubscription) ToOrganizationDeveloperSubscriptionOutput() OrganizationDeveloperSubscriptionOutput {
	return i.ToOrganizationDeveloperSubscriptionOutputWithContext(context.Background())
}

func (i *OrganizationDeveloperSubscription) ToOrganizationDeveloperSubscriptionOutputWithContext(ctx context.Context) OrganizationDeveloperSubscriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationDeveloperSubscriptionOutput)
}

type OrganizationDeveloperSubscriptionOutput struct {
	*pulumi.OutputState
}

func (OrganizationDeveloperSubscriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationDeveloperSubscription)(nil))
}

func (o OrganizationDeveloperSubscriptionOutput) ToOrganizationDeveloperSubscriptionOutput() OrganizationDeveloperSubscriptionOutput {
	return o
}

func (o OrganizationDeveloperSubscriptionOutput) ToOrganizationDeveloperSubscriptionOutputWithContext(ctx context.Context) OrganizationDeveloperSubscriptionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(OrganizationDeveloperSubscriptionOutput{})
}
