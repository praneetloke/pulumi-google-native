// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new category on the portal.
type OrganizationSiteApicategory struct {
	pulumi.CustomResourceState

	// Details of category.
	Data GoogleCloudApigeeV1ApiCategoryDataResponseOutput `pulumi:"data"`
	// ID that can be used to find errors in the log files.
	ErrorCode pulumi.StringOutput `pulumi:"errorCode"`
	// Description of the operation.
	Message pulumi.StringOutput `pulumi:"message"`
	// ID that can be used to find request details in the log files.
	RequestId pulumi.StringOutput `pulumi:"requestId"`
	// Status of the operation.
	Status pulumi.StringOutput `pulumi:"status"`
}

// NewOrganizationSiteApicategory registers a new resource with the given unique name, arguments, and options.
func NewOrganizationSiteApicategory(ctx *pulumi.Context,
	name string, args *OrganizationSiteApicategoryArgs, opts ...pulumi.ResourceOption) (*OrganizationSiteApicategory, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ApicategoryId == nil {
		return nil, errors.New("invalid value for required argument 'ApicategoryId'")
	}
	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	if args.SiteId == nil {
		return nil, errors.New("invalid value for required argument 'SiteId'")
	}
	var resource OrganizationSiteApicategory
	err := ctx.RegisterResource("google-native:apigee/v1:OrganizationSiteApicategory", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationSiteApicategory gets an existing OrganizationSiteApicategory resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationSiteApicategory(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationSiteApicategoryState, opts ...pulumi.ResourceOption) (*OrganizationSiteApicategory, error) {
	var resource OrganizationSiteApicategory
	err := ctx.ReadResource("google-native:apigee/v1:OrganizationSiteApicategory", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationSiteApicategory resources.
type organizationSiteApicategoryState struct {
	// Details of category.
	Data *GoogleCloudApigeeV1ApiCategoryDataResponse `pulumi:"data"`
	// ID that can be used to find errors in the log files.
	ErrorCode *string `pulumi:"errorCode"`
	// Description of the operation.
	Message *string `pulumi:"message"`
	// ID that can be used to find request details in the log files.
	RequestId *string `pulumi:"requestId"`
	// Status of the operation.
	Status *string `pulumi:"status"`
}

type OrganizationSiteApicategoryState struct {
	// Details of category.
	Data GoogleCloudApigeeV1ApiCategoryDataResponsePtrInput
	// ID that can be used to find errors in the log files.
	ErrorCode pulumi.StringPtrInput
	// Description of the operation.
	Message pulumi.StringPtrInput
	// ID that can be used to find request details in the log files.
	RequestId pulumi.StringPtrInput
	// Status of the operation.
	Status pulumi.StringPtrInput
}

func (OrganizationSiteApicategoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationSiteApicategoryState)(nil)).Elem()
}

type organizationSiteApicategoryArgs struct {
	ApicategoryId string `pulumi:"apicategoryId"`
	// ID of the category (a UUID).
	Id *string `pulumi:"id"`
	// Name of the category.
	Name           *string `pulumi:"name"`
	OrganizationId string  `pulumi:"organizationId"`
	// Name of the portal.
	SiteId string `pulumi:"siteId"`
	// Time the category was last modified in milliseconds since epoch.
	UpdateTime *string `pulumi:"updateTime"`
}

// The set of arguments for constructing a OrganizationSiteApicategory resource.
type OrganizationSiteApicategoryArgs struct {
	ApicategoryId pulumi.StringInput
	// ID of the category (a UUID).
	Id pulumi.StringPtrInput
	// Name of the category.
	Name           pulumi.StringPtrInput
	OrganizationId pulumi.StringInput
	// Name of the portal.
	SiteId pulumi.StringInput
	// Time the category was last modified in milliseconds since epoch.
	UpdateTime pulumi.StringPtrInput
}

func (OrganizationSiteApicategoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationSiteApicategoryArgs)(nil)).Elem()
}

type OrganizationSiteApicategoryInput interface {
	pulumi.Input

	ToOrganizationSiteApicategoryOutput() OrganizationSiteApicategoryOutput
	ToOrganizationSiteApicategoryOutputWithContext(ctx context.Context) OrganizationSiteApicategoryOutput
}

func (*OrganizationSiteApicategory) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationSiteApicategory)(nil))
}

func (i *OrganizationSiteApicategory) ToOrganizationSiteApicategoryOutput() OrganizationSiteApicategoryOutput {
	return i.ToOrganizationSiteApicategoryOutputWithContext(context.Background())
}

func (i *OrganizationSiteApicategory) ToOrganizationSiteApicategoryOutputWithContext(ctx context.Context) OrganizationSiteApicategoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationSiteApicategoryOutput)
}

type OrganizationSiteApicategoryOutput struct {
	*pulumi.OutputState
}

func (OrganizationSiteApicategoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationSiteApicategory)(nil))
}

func (o OrganizationSiteApicategoryOutput) ToOrganizationSiteApicategoryOutput() OrganizationSiteApicategoryOutput {
	return o
}

func (o OrganizationSiteApicategoryOutput) ToOrganizationSiteApicategoryOutputWithContext(ctx context.Context) OrganizationSiteApicategoryOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(OrganizationSiteApicategoryOutput{})
}
