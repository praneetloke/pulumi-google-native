// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a Microsoft AD Domain in a given project. Operation
// Auto-naming is currently not supported for this resource.
type Domain struct {
	pulumi.CustomResourceState

	// Optional. Configuration for audit logs. True if audit logs are enabled, else false. Default is audit logs disabled.
	AuditLogsEnabled pulumi.BoolOutput `pulumi:"auditLogsEnabled"`
	// Optional. The full names of the Google Compute Engine [networks](/compute/docs/networks-and-firewalls#networks) to which the instance is connected. Network can be added using UpdateDomain later. Domain is only available on network part of authorized_networks. Caller needs to make sure that CIDR subnets do not overlap between networks, else domain creation will fail.
	AuthorizedNetworks pulumi.StringArrayOutput `pulumi:"authorizedNetworks"`
	// The time the instance was created. Synthetic field is populated automatically by CCFE. go/ccfe-synthetic-field-user-guide
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The fully qualified domain name. e.g. mydomain.myorganization.com, with the following restrictions: * Must contain only lowercase letters, numbers, periods and hyphens. * Must start with a letter. * Must contain between 2-64 characters. * Must end with a number or a letter. * Must not start with period. * Must be unique within the project. * First segment length (mydomain form example above) shouldn't exceed 15 chars. * The last segment cannot be fully numeric.
	DomainName pulumi.StringPtrOutput `pulumi:"domainName"`
	// Fully-qualified domain name of the exposed domain used by clients to connect to the service. Similar to what would be chosen for an Active Directory that is set up on an internal network.
	Fqdn pulumi.StringOutput `pulumi:"fqdn"`
	// Optional. Resource labels to represent user provided metadata
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Locations where domain needs to be provisioned. regions e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
	Locations pulumi.StringArrayOutput `pulumi:"locations"`
	// Optional. Name of customer-visible admin used to perform Active Directory operations. If not specified `setupadmin` would be used.
	ManagedIdentitiesAdminName pulumi.StringOutput `pulumi:"managedIdentitiesAdminName"`
	// Unique name of the domain in this scope including projects and location using the form: `projects/{project_id}/locations/global/domains/{domain_name}`.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger. Ranges must be unique and non-overlapping with existing subnets in [Domain].[authorized_networks].
	ReservedIpRange pulumi.StringOutput `pulumi:"reservedIpRange"`
	// The current state of this domain.
	State pulumi.StringOutput `pulumi:"state"`
	// Additional information about the current status of this domain, if available.
	StatusMessage pulumi.StringOutput `pulumi:"statusMessage"`
	// The current trusts associated with the domain.
	Trusts TrustResponseArrayOutput `pulumi:"trusts"`
	// Last update time. Synthetic field is populated automatically by CCFE.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewDomain registers a new resource with the given unique name, arguments, and options.
func NewDomain(ctx *pulumi.Context,
	name string, args *DomainArgs, opts ...pulumi.ResourceOption) (*Domain, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Locations == nil {
		return nil, errors.New("invalid value for required argument 'Locations'")
	}
	if args.ReservedIpRange == nil {
		return nil, errors.New("invalid value for required argument 'ReservedIpRange'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Domain
	err := ctx.RegisterResource("google-native:managedidentities/v1alpha1:Domain", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDomain gets an existing Domain resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDomain(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DomainState, opts ...pulumi.ResourceOption) (*Domain, error) {
	var resource Domain
	err := ctx.ReadResource("google-native:managedidentities/v1alpha1:Domain", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Domain resources.
type domainState struct {
}

type DomainState struct {
}

func (DomainState) ElementType() reflect.Type {
	return reflect.TypeOf((*domainState)(nil)).Elem()
}

type domainArgs struct {
	// Optional. Configuration for audit logs. True if audit logs are enabled, else false. Default is audit logs disabled.
	AuditLogsEnabled *bool `pulumi:"auditLogsEnabled"`
	// Optional. The full names of the Google Compute Engine [networks](/compute/docs/networks-and-firewalls#networks) to which the instance is connected. Network can be added using UpdateDomain later. Domain is only available on network part of authorized_networks. Caller needs to make sure that CIDR subnets do not overlap between networks, else domain creation will fail.
	AuthorizedNetworks []string `pulumi:"authorizedNetworks"`
	// The fully qualified domain name. e.g. mydomain.myorganization.com, with the following restrictions: * Must contain only lowercase letters, numbers, periods and hyphens. * Must start with a letter. * Must contain between 2-64 characters. * Must end with a number or a letter. * Must not start with period. * Must be unique within the project. * First segment length (mydomain form example above) shouldn't exceed 15 chars. * The last segment cannot be fully numeric.
	DomainName *string `pulumi:"domainName"`
	// Optional. Resource labels to represent user provided metadata
	Labels map[string]string `pulumi:"labels"`
	// Locations where domain needs to be provisioned. regions e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
	Locations []string `pulumi:"locations"`
	// Optional. Name of customer-visible admin used to perform Active Directory operations. If not specified `setupadmin` would be used.
	ManagedIdentitiesAdminName *string `pulumi:"managedIdentitiesAdminName"`
	Project                    *string `pulumi:"project"`
	// The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger. Ranges must be unique and non-overlapping with existing subnets in [Domain].[authorized_networks].
	ReservedIpRange string `pulumi:"reservedIpRange"`
}

// The set of arguments for constructing a Domain resource.
type DomainArgs struct {
	// Optional. Configuration for audit logs. True if audit logs are enabled, else false. Default is audit logs disabled.
	AuditLogsEnabled pulumi.BoolPtrInput
	// Optional. The full names of the Google Compute Engine [networks](/compute/docs/networks-and-firewalls#networks) to which the instance is connected. Network can be added using UpdateDomain later. Domain is only available on network part of authorized_networks. Caller needs to make sure that CIDR subnets do not overlap between networks, else domain creation will fail.
	AuthorizedNetworks pulumi.StringArrayInput
	// The fully qualified domain name. e.g. mydomain.myorganization.com, with the following restrictions: * Must contain only lowercase letters, numbers, periods and hyphens. * Must start with a letter. * Must contain between 2-64 characters. * Must end with a number or a letter. * Must not start with period. * Must be unique within the project. * First segment length (mydomain form example above) shouldn't exceed 15 chars. * The last segment cannot be fully numeric.
	DomainName pulumi.StringPtrInput
	// Optional. Resource labels to represent user provided metadata
	Labels pulumi.StringMapInput
	// Locations where domain needs to be provisioned. regions e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
	Locations pulumi.StringArrayInput
	// Optional. Name of customer-visible admin used to perform Active Directory operations. If not specified `setupadmin` would be used.
	ManagedIdentitiesAdminName pulumi.StringPtrInput
	Project                    pulumi.StringPtrInput
	// The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger. Ranges must be unique and non-overlapping with existing subnets in [Domain].[authorized_networks].
	ReservedIpRange pulumi.StringInput
}

func (DomainArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*domainArgs)(nil)).Elem()
}

type DomainInput interface {
	pulumi.Input

	ToDomainOutput() DomainOutput
	ToDomainOutputWithContext(ctx context.Context) DomainOutput
}

func (*Domain) ElementType() reflect.Type {
	return reflect.TypeOf((**Domain)(nil)).Elem()
}

func (i *Domain) ToDomainOutput() DomainOutput {
	return i.ToDomainOutputWithContext(context.Background())
}

func (i *Domain) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DomainOutput)
}

type DomainOutput struct{ *pulumi.OutputState }

func (DomainOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Domain)(nil)).Elem()
}

func (o DomainOutput) ToDomainOutput() DomainOutput {
	return o
}

func (o DomainOutput) ToDomainOutputWithContext(ctx context.Context) DomainOutput {
	return o
}

// Optional. Configuration for audit logs. True if audit logs are enabled, else false. Default is audit logs disabled.
func (o DomainOutput) AuditLogsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *Domain) pulumi.BoolOutput { return v.AuditLogsEnabled }).(pulumi.BoolOutput)
}

// Optional. The full names of the Google Compute Engine [networks](/compute/docs/networks-and-firewalls#networks) to which the instance is connected. Network can be added using UpdateDomain later. Domain is only available on network part of authorized_networks. Caller needs to make sure that CIDR subnets do not overlap between networks, else domain creation will fail.
func (o DomainOutput) AuthorizedNetworks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringArrayOutput { return v.AuthorizedNetworks }).(pulumi.StringArrayOutput)
}

// The time the instance was created. Synthetic field is populated automatically by CCFE. go/ccfe-synthetic-field-user-guide
func (o DomainOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The fully qualified domain name. e.g. mydomain.myorganization.com, with the following restrictions: * Must contain only lowercase letters, numbers, periods and hyphens. * Must start with a letter. * Must contain between 2-64 characters. * Must end with a number or a letter. * Must not start with period. * Must be unique within the project. * First segment length (mydomain form example above) shouldn't exceed 15 chars. * The last segment cannot be fully numeric.
func (o DomainOutput) DomainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringPtrOutput { return v.DomainName }).(pulumi.StringPtrOutput)
}

// Fully-qualified domain name of the exposed domain used by clients to connect to the service. Similar to what would be chosen for an Active Directory that is set up on an internal network.
func (o DomainOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Fqdn }).(pulumi.StringOutput)
}

// Optional. Resource labels to represent user provided metadata
func (o DomainOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

// Locations where domain needs to be provisioned. regions e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
func (o DomainOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringArrayOutput { return v.Locations }).(pulumi.StringArrayOutput)
}

// Optional. Name of customer-visible admin used to perform Active Directory operations. If not specified `setupadmin` would be used.
func (o DomainOutput) ManagedIdentitiesAdminName() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.ManagedIdentitiesAdminName }).(pulumi.StringOutput)
}

// Unique name of the domain in this scope including projects and location using the form: `projects/{project_id}/locations/global/domains/{domain_name}`.
func (o DomainOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o DomainOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger. Ranges must be unique and non-overlapping with existing subnets in [Domain].[authorized_networks].
func (o DomainOutput) ReservedIpRange() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.ReservedIpRange }).(pulumi.StringOutput)
}

// The current state of this domain.
func (o DomainOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// Additional information about the current status of this domain, if available.
func (o DomainOutput) StatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.StatusMessage }).(pulumi.StringOutput)
}

// The current trusts associated with the domain.
func (o DomainOutput) Trusts() TrustResponseArrayOutput {
	return o.ApplyT(func(v *Domain) TrustResponseArrayOutput { return v.Trusts }).(TrustResponseArrayOutput)
}

// Last update time. Synthetic field is populated automatically by CCFE.
func (o DomainOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Domain) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DomainInput)(nil)).Elem(), &Domain{})
	pulumi.RegisterOutputType(DomainOutput{})
}
