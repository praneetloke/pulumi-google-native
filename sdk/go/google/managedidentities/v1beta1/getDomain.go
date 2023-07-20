// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets information about a domain.
func LookupDomain(ctx *pulumi.Context, args *LookupDomainArgs, opts ...pulumi.InvokeOption) (*LookupDomainResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupDomainResult
	err := ctx.Invoke("google-native:managedidentities/v1beta1:getDomain", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDomainArgs struct {
	DomainId string  `pulumi:"domainId"`
	Project  *string `pulumi:"project"`
}

type LookupDomainResult struct {
	// Optional. The name of delegated administrator account used to perform Active Directory operations. If not specified, `setupadmin` will be used.
	Admin string `pulumi:"admin"`
	// Optional. Configuration for audit logs. True if audit logs are enabled, else false. Default is audit logs disabled.
	AuditLogsEnabled bool `pulumi:"auditLogsEnabled"`
	// Optional. The full names of the Google Compute Engine [networks](/compute/docs/networks-and-firewalls#networks) the domain instance is connected to. Networks can be added using UpdateDomain. The domain is only available on networks listed in `authorized_networks`. If CIDR subnets overlap between networks, domain creation will fail.
	AuthorizedNetworks []string `pulumi:"authorizedNetworks"`
	// The time the instance was created.
	CreateTime string `pulumi:"createTime"`
	// The fully-qualified domain name of the exposed domain used by clients to connect to the service. Similar to what would be chosen for an Active Directory set up on an internal network.
	Fqdn string `pulumi:"fqdn"`
	// Optional. Resource labels that can contain user-provided metadata.
	Labels map[string]string `pulumi:"labels"`
	// Locations where domain needs to be provisioned. regions e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
	Locations []string `pulumi:"locations"`
	// The unique name of the domain using the form: `projects/{project_id}/locations/global/domains/{domain_name}`.
	Name string `pulumi:"name"`
	// The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger. Ranges must be unique and non-overlapping with existing subnets in [Domain].[authorized_networks].
	ReservedIpRange string `pulumi:"reservedIpRange"`
	// The current state of this domain.
	State string `pulumi:"state"`
	// Additional information about the current status of this domain, if available.
	StatusMessage string `pulumi:"statusMessage"`
	// The current trusts associated with the domain.
	Trusts []TrustResponse `pulumi:"trusts"`
	// The last update time.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupDomainOutput(ctx *pulumi.Context, args LookupDomainOutputArgs, opts ...pulumi.InvokeOption) LookupDomainResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDomainResult, error) {
			args := v.(LookupDomainArgs)
			r, err := LookupDomain(ctx, &args, opts...)
			var s LookupDomainResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDomainResultOutput)
}

type LookupDomainOutputArgs struct {
	DomainId pulumi.StringInput    `pulumi:"domainId"`
	Project  pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupDomainOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainArgs)(nil)).Elem()
}

type LookupDomainResultOutput struct{ *pulumi.OutputState }

func (LookupDomainResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDomainResult)(nil)).Elem()
}

func (o LookupDomainResultOutput) ToLookupDomainResultOutput() LookupDomainResultOutput {
	return o
}

func (o LookupDomainResultOutput) ToLookupDomainResultOutputWithContext(ctx context.Context) LookupDomainResultOutput {
	return o
}

// Optional. The name of delegated administrator account used to perform Active Directory operations. If not specified, `setupadmin` will be used.
func (o LookupDomainResultOutput) Admin() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.Admin }).(pulumi.StringOutput)
}

// Optional. Configuration for audit logs. True if audit logs are enabled, else false. Default is audit logs disabled.
func (o LookupDomainResultOutput) AuditLogsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDomainResult) bool { return v.AuditLogsEnabled }).(pulumi.BoolOutput)
}

// Optional. The full names of the Google Compute Engine [networks](/compute/docs/networks-and-firewalls#networks) the domain instance is connected to. Networks can be added using UpdateDomain. The domain is only available on networks listed in `authorized_networks`. If CIDR subnets overlap between networks, domain creation will fail.
func (o LookupDomainResultOutput) AuthorizedNetworks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDomainResult) []string { return v.AuthorizedNetworks }).(pulumi.StringArrayOutput)
}

// The time the instance was created.
func (o LookupDomainResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The fully-qualified domain name of the exposed domain used by clients to connect to the service. Similar to what would be chosen for an Active Directory set up on an internal network.
func (o LookupDomainResultOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.Fqdn }).(pulumi.StringOutput)
}

// Optional. Resource labels that can contain user-provided metadata.
func (o LookupDomainResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupDomainResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Locations where domain needs to be provisioned. regions e.g. us-west1 or us-east4 Service supports up to 4 locations at once. Each location will use a /26 block.
func (o LookupDomainResultOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupDomainResult) []string { return v.Locations }).(pulumi.StringArrayOutput)
}

// The unique name of the domain using the form: `projects/{project_id}/locations/global/domains/{domain_name}`.
func (o LookupDomainResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.Name }).(pulumi.StringOutput)
}

// The CIDR range of internal addresses that are reserved for this domain. Reserved networks must be /24 or larger. Ranges must be unique and non-overlapping with existing subnets in [Domain].[authorized_networks].
func (o LookupDomainResultOutput) ReservedIpRange() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.ReservedIpRange }).(pulumi.StringOutput)
}

// The current state of this domain.
func (o LookupDomainResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.State }).(pulumi.StringOutput)
}

// Additional information about the current status of this domain, if available.
func (o LookupDomainResultOutput) StatusMessage() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.StatusMessage }).(pulumi.StringOutput)
}

// The current trusts associated with the domain.
func (o LookupDomainResultOutput) Trusts() TrustResponseArrayOutput {
	return o.ApplyT(func(v LookupDomainResult) []TrustResponse { return v.Trusts }).(TrustResponseArrayOutput)
}

// The last update time.
func (o LookupDomainResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDomainResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDomainResultOutput{})
}
