// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Fetches the representation of an existing ResourceRecordSet.
func LookupResourceRecordSet(ctx *pulumi.Context, args *LookupResourceRecordSetArgs, opts ...pulumi.InvokeOption) (*LookupResourceRecordSetResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupResourceRecordSetResult
	err := ctx.Invoke("google-native:dns/v1:getResourceRecordSet", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupResourceRecordSetArgs struct {
	ClientOperationId *string `pulumi:"clientOperationId"`
	ManagedZone       string  `pulumi:"managedZone"`
	Name              string  `pulumi:"name"`
	Project           *string `pulumi:"project"`
	Type              string  `pulumi:"type"`
}

type LookupResourceRecordSetResult struct {
	Kind string `pulumi:"kind"`
	// For example, www.example.com.
	Name string `pulumi:"name"`
	// Configures dynamic query responses based on geo location of querying user or a weighted round robin based routing policy. A ResourceRecordSet should only have either rrdata (static) or routing_policy (dynamic). An error is returned otherwise.
	RoutingPolicy RRSetRoutingPolicyResponse `pulumi:"routingPolicy"`
	// As defined in RFC 1035 (section 5) and RFC 1034 (section 3.6.1) -- see examples.
	Rrdatas []string `pulumi:"rrdatas"`
	// As defined in RFC 4034 (section 3.2).
	SignatureRrdatas []string `pulumi:"signatureRrdatas"`
	// Number of seconds that this ResourceRecordSet can be cached by resolvers.
	Ttl int `pulumi:"ttl"`
	// The identifier of a supported record type. See the list of Supported DNS record types.
	Type string `pulumi:"type"`
}

func LookupResourceRecordSetOutput(ctx *pulumi.Context, args LookupResourceRecordSetOutputArgs, opts ...pulumi.InvokeOption) LookupResourceRecordSetResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupResourceRecordSetResult, error) {
			args := v.(LookupResourceRecordSetArgs)
			r, err := LookupResourceRecordSet(ctx, &args, opts...)
			var s LookupResourceRecordSetResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupResourceRecordSetResultOutput)
}

type LookupResourceRecordSetOutputArgs struct {
	ClientOperationId pulumi.StringPtrInput `pulumi:"clientOperationId"`
	ManagedZone       pulumi.StringInput    `pulumi:"managedZone"`
	Name              pulumi.StringInput    `pulumi:"name"`
	Project           pulumi.StringPtrInput `pulumi:"project"`
	Type              pulumi.StringInput    `pulumi:"type"`
}

func (LookupResourceRecordSetOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceRecordSetArgs)(nil)).Elem()
}

type LookupResourceRecordSetResultOutput struct{ *pulumi.OutputState }

func (LookupResourceRecordSetResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupResourceRecordSetResult)(nil)).Elem()
}

func (o LookupResourceRecordSetResultOutput) ToLookupResourceRecordSetResultOutput() LookupResourceRecordSetResultOutput {
	return o
}

func (o LookupResourceRecordSetResultOutput) ToLookupResourceRecordSetResultOutputWithContext(ctx context.Context) LookupResourceRecordSetResultOutput {
	return o
}

func (o LookupResourceRecordSetResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceRecordSetResult) string { return v.Kind }).(pulumi.StringOutput)
}

// For example, www.example.com.
func (o LookupResourceRecordSetResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceRecordSetResult) string { return v.Name }).(pulumi.StringOutput)
}

// Configures dynamic query responses based on geo location of querying user or a weighted round robin based routing policy. A ResourceRecordSet should only have either rrdata (static) or routing_policy (dynamic). An error is returned otherwise.
func (o LookupResourceRecordSetResultOutput) RoutingPolicy() RRSetRoutingPolicyResponseOutput {
	return o.ApplyT(func(v LookupResourceRecordSetResult) RRSetRoutingPolicyResponse { return v.RoutingPolicy }).(RRSetRoutingPolicyResponseOutput)
}

// As defined in RFC 1035 (section 5) and RFC 1034 (section 3.6.1) -- see examples.
func (o LookupResourceRecordSetResultOutput) Rrdatas() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupResourceRecordSetResult) []string { return v.Rrdatas }).(pulumi.StringArrayOutput)
}

// As defined in RFC 4034 (section 3.2).
func (o LookupResourceRecordSetResultOutput) SignatureRrdatas() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupResourceRecordSetResult) []string { return v.SignatureRrdatas }).(pulumi.StringArrayOutput)
}

// Number of seconds that this ResourceRecordSet can be cached by resolvers.
func (o LookupResourceRecordSetResultOutput) Ttl() pulumi.IntOutput {
	return o.ApplyT(func(v LookupResourceRecordSetResult) int { return v.Ttl }).(pulumi.IntOutput)
}

// The identifier of a supported record type. See the list of Supported DNS record types.
func (o LookupResourceRecordSetResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupResourceRecordSetResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupResourceRecordSetResultOutput{})
}
