// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a source.
type OrganizationSource struct {
	pulumi.CustomResourceState

	// The description of the source (max of 1024 characters). Example: "Web Security Scanner is a web security scanner for common vulnerabilities in App Engine applications. It can automatically scan and detect four common vulnerabilities, including cross-site-scripting (XSS), Flash injection, mixed content (HTTP in HTTPS), and outdated/insecure libraries."
	Description pulumi.StringOutput `pulumi:"description"`
	// The source's display name. A source's display name must be unique amongst its siblings, for example, two sources with the same parent can't share the same display name. The display name must have a length between 1 and 64 characters (inclusive).
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The relative resource name of this source. See: https://cloud.google.com/apis/design/resource_names#relative_resource_name Example: "organizations/{organization_id}/sources/{source_id}"
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewOrganizationSource registers a new resource with the given unique name, arguments, and options.
func NewOrganizationSource(ctx *pulumi.Context,
	name string, args *OrganizationSourceArgs, opts ...pulumi.ResourceOption) (*OrganizationSource, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.OrganizationId == nil {
		return nil, errors.New("invalid value for required argument 'OrganizationId'")
	}
	if args.SourceId == nil {
		return nil, errors.New("invalid value for required argument 'SourceId'")
	}
	var resource OrganizationSource
	err := ctx.RegisterResource("google-native:securitycenter/v1beta1:OrganizationSource", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetOrganizationSource gets an existing OrganizationSource resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetOrganizationSource(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *OrganizationSourceState, opts ...pulumi.ResourceOption) (*OrganizationSource, error) {
	var resource OrganizationSource
	err := ctx.ReadResource("google-native:securitycenter/v1beta1:OrganizationSource", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering OrganizationSource resources.
type organizationSourceState struct {
	// The description of the source (max of 1024 characters). Example: "Web Security Scanner is a web security scanner for common vulnerabilities in App Engine applications. It can automatically scan and detect four common vulnerabilities, including cross-site-scripting (XSS), Flash injection, mixed content (HTTP in HTTPS), and outdated/insecure libraries."
	Description *string `pulumi:"description"`
	// The source's display name. A source's display name must be unique amongst its siblings, for example, two sources with the same parent can't share the same display name. The display name must have a length between 1 and 64 characters (inclusive).
	DisplayName *string `pulumi:"displayName"`
	// The relative resource name of this source. See: https://cloud.google.com/apis/design/resource_names#relative_resource_name Example: "organizations/{organization_id}/sources/{source_id}"
	Name *string `pulumi:"name"`
}

type OrganizationSourceState struct {
	// The description of the source (max of 1024 characters). Example: "Web Security Scanner is a web security scanner for common vulnerabilities in App Engine applications. It can automatically scan and detect four common vulnerabilities, including cross-site-scripting (XSS), Flash injection, mixed content (HTTP in HTTPS), and outdated/insecure libraries."
	Description pulumi.StringPtrInput
	// The source's display name. A source's display name must be unique amongst its siblings, for example, two sources with the same parent can't share the same display name. The display name must have a length between 1 and 64 characters (inclusive).
	DisplayName pulumi.StringPtrInput
	// The relative resource name of this source. See: https://cloud.google.com/apis/design/resource_names#relative_resource_name Example: "organizations/{organization_id}/sources/{source_id}"
	Name pulumi.StringPtrInput
}

func (OrganizationSourceState) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationSourceState)(nil)).Elem()
}

type organizationSourceArgs struct {
	// The description of the source (max of 1024 characters). Example: "Web Security Scanner is a web security scanner for common vulnerabilities in App Engine applications. It can automatically scan and detect four common vulnerabilities, including cross-site-scripting (XSS), Flash injection, mixed content (HTTP in HTTPS), and outdated/insecure libraries."
	Description *string `pulumi:"description"`
	// The source's display name. A source's display name must be unique amongst its siblings, for example, two sources with the same parent can't share the same display name. The display name must have a length between 1 and 64 characters (inclusive).
	DisplayName *string `pulumi:"displayName"`
	// The relative resource name of this source. See: https://cloud.google.com/apis/design/resource_names#relative_resource_name Example: "organizations/{organization_id}/sources/{source_id}"
	Name           *string `pulumi:"name"`
	OrganizationId string  `pulumi:"organizationId"`
	SourceId       string  `pulumi:"sourceId"`
}

// The set of arguments for constructing a OrganizationSource resource.
type OrganizationSourceArgs struct {
	// The description of the source (max of 1024 characters). Example: "Web Security Scanner is a web security scanner for common vulnerabilities in App Engine applications. It can automatically scan and detect four common vulnerabilities, including cross-site-scripting (XSS), Flash injection, mixed content (HTTP in HTTPS), and outdated/insecure libraries."
	Description pulumi.StringPtrInput
	// The source's display name. A source's display name must be unique amongst its siblings, for example, two sources with the same parent can't share the same display name. The display name must have a length between 1 and 64 characters (inclusive).
	DisplayName pulumi.StringPtrInput
	// The relative resource name of this source. See: https://cloud.google.com/apis/design/resource_names#relative_resource_name Example: "organizations/{organization_id}/sources/{source_id}"
	Name           pulumi.StringPtrInput
	OrganizationId pulumi.StringInput
	SourceId       pulumi.StringInput
}

func (OrganizationSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*organizationSourceArgs)(nil)).Elem()
}

type OrganizationSourceInput interface {
	pulumi.Input

	ToOrganizationSourceOutput() OrganizationSourceOutput
	ToOrganizationSourceOutputWithContext(ctx context.Context) OrganizationSourceOutput
}

func (*OrganizationSource) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationSource)(nil))
}

func (i *OrganizationSource) ToOrganizationSourceOutput() OrganizationSourceOutput {
	return i.ToOrganizationSourceOutputWithContext(context.Background())
}

func (i *OrganizationSource) ToOrganizationSourceOutputWithContext(ctx context.Context) OrganizationSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OrganizationSourceOutput)
}

type OrganizationSourceOutput struct {
	*pulumi.OutputState
}

func (OrganizationSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OrganizationSource)(nil))
}

func (o OrganizationSourceOutput) ToOrganizationSourceOutput() OrganizationSourceOutput {
	return o
}

func (o OrganizationSourceOutput) ToOrganizationSourceOutputWithContext(ctx context.Context) OrganizationSourceOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(OrganizationSourceOutput{})
}
