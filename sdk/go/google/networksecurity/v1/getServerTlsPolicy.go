// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single ServerTlsPolicy.
func LookupServerTlsPolicy(ctx *pulumi.Context, args *LookupServerTlsPolicyArgs, opts ...pulumi.InvokeOption) (*LookupServerTlsPolicyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupServerTlsPolicyResult
	err := ctx.Invoke("google-native:networksecurity/v1:getServerTlsPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupServerTlsPolicyArgs struct {
	Location          string  `pulumi:"location"`
	Project           *string `pulumi:"project"`
	ServerTlsPolicyId string  `pulumi:"serverTlsPolicyId"`
}

type LookupServerTlsPolicyResult struct {
	// This field applies only for Traffic Director policies. It is must be set to false for external HTTPS load balancer policies. Determines if server allows plaintext connections. If set to true, server allows plain text connections. By default, it is set to false. This setting is not exclusive of other encryption modes. For example, if `allow_open` and `mtls_policy` are set, server allows both plain text and mTLS connections. See documentation of other encryption modes to confirm compatibility. Consider using it if you wish to upgrade in place your deployment to TLS while having mixed TLS and non-TLS traffic reaching port :80.
	AllowOpen bool `pulumi:"allowOpen"`
	// The timestamp when the resource was created.
	CreateTime string `pulumi:"createTime"`
	// Free-text description of the resource.
	Description string `pulumi:"description"`
	// Set of label tags associated with the resource.
	Labels map[string]string `pulumi:"labels"`
	// This field is required if the policy is used with external HTTPS load balancers. This field can be empty for Traffic Director. Defines a mechanism to provision peer validation certificates for peer to peer authentication (Mutual TLS - mTLS). If not specified, client certificate will not be requested. The connection is treated as TLS and not mTLS. If `allow_open` and `mtls_policy` are set, server allows both plain text and mTLS connections.
	MtlsPolicy MTLSPolicyResponse `pulumi:"mtlsPolicy"`
	// Name of the ServerTlsPolicy resource. It matches the pattern `projects/*/locations/{location}/serverTlsPolicies/{server_tls_policy}`
	Name string `pulumi:"name"`
	// Optional if policy is to be used with Traffic Director. For external HTTPS load balancer must be empty. Defines a mechanism to provision server identity (public and private keys). Cannot be combined with `allow_open` as a permissive mode that allows both plain text and TLS is not supported.
	ServerCertificate GoogleCloudNetworksecurityV1CertificateProviderResponse `pulumi:"serverCertificate"`
	// The timestamp when the resource was updated.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupServerTlsPolicyOutput(ctx *pulumi.Context, args LookupServerTlsPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupServerTlsPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupServerTlsPolicyResult, error) {
			args := v.(LookupServerTlsPolicyArgs)
			r, err := LookupServerTlsPolicy(ctx, &args, opts...)
			var s LookupServerTlsPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupServerTlsPolicyResultOutput)
}

type LookupServerTlsPolicyOutputArgs struct {
	Location          pulumi.StringInput    `pulumi:"location"`
	Project           pulumi.StringPtrInput `pulumi:"project"`
	ServerTlsPolicyId pulumi.StringInput    `pulumi:"serverTlsPolicyId"`
}

func (LookupServerTlsPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerTlsPolicyArgs)(nil)).Elem()
}

type LookupServerTlsPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupServerTlsPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupServerTlsPolicyResult)(nil)).Elem()
}

func (o LookupServerTlsPolicyResultOutput) ToLookupServerTlsPolicyResultOutput() LookupServerTlsPolicyResultOutput {
	return o
}

func (o LookupServerTlsPolicyResultOutput) ToLookupServerTlsPolicyResultOutputWithContext(ctx context.Context) LookupServerTlsPolicyResultOutput {
	return o
}

// This field applies only for Traffic Director policies. It is must be set to false for external HTTPS load balancer policies. Determines if server allows plaintext connections. If set to true, server allows plain text connections. By default, it is set to false. This setting is not exclusive of other encryption modes. For example, if `allow_open` and `mtls_policy` are set, server allows both plain text and mTLS connections. See documentation of other encryption modes to confirm compatibility. Consider using it if you wish to upgrade in place your deployment to TLS while having mixed TLS and non-TLS traffic reaching port :80.
func (o LookupServerTlsPolicyResultOutput) AllowOpen() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupServerTlsPolicyResult) bool { return v.AllowOpen }).(pulumi.BoolOutput)
}

// The timestamp when the resource was created.
func (o LookupServerTlsPolicyResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerTlsPolicyResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Free-text description of the resource.
func (o LookupServerTlsPolicyResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerTlsPolicyResult) string { return v.Description }).(pulumi.StringOutput)
}

// Set of label tags associated with the resource.
func (o LookupServerTlsPolicyResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupServerTlsPolicyResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// This field is required if the policy is used with external HTTPS load balancers. This field can be empty for Traffic Director. Defines a mechanism to provision peer validation certificates for peer to peer authentication (Mutual TLS - mTLS). If not specified, client certificate will not be requested. The connection is treated as TLS and not mTLS. If `allow_open` and `mtls_policy` are set, server allows both plain text and mTLS connections.
func (o LookupServerTlsPolicyResultOutput) MtlsPolicy() MTLSPolicyResponseOutput {
	return o.ApplyT(func(v LookupServerTlsPolicyResult) MTLSPolicyResponse { return v.MtlsPolicy }).(MTLSPolicyResponseOutput)
}

// Name of the ServerTlsPolicy resource. It matches the pattern `projects/*/locations/{location}/serverTlsPolicies/{server_tls_policy}`
func (o LookupServerTlsPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerTlsPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional if policy is to be used with Traffic Director. For external HTTPS load balancer must be empty. Defines a mechanism to provision server identity (public and private keys). Cannot be combined with `allow_open` as a permissive mode that allows both plain text and TLS is not supported.
func (o LookupServerTlsPolicyResultOutput) ServerCertificate() GoogleCloudNetworksecurityV1CertificateProviderResponseOutput {
	return o.ApplyT(func(v LookupServerTlsPolicyResult) GoogleCloudNetworksecurityV1CertificateProviderResponse {
		return v.ServerCertificate
	}).(GoogleCloudNetworksecurityV1CertificateProviderResponseOutput)
}

// The timestamp when the resource was updated.
func (o LookupServerTlsPolicyResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupServerTlsPolicyResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupServerTlsPolicyResultOutput{})
}
