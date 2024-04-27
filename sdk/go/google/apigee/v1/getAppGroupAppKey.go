// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details for a consumer key for a AppGroup app, including the key and secret value, associated API products, and other information.
func GetAppGroupAppKey(ctx *pulumi.Context, args *GetAppGroupAppKeyArgs, opts ...pulumi.InvokeOption) (*GetAppGroupAppKeyResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetAppGroupAppKeyResult
	err := ctx.Invoke("google-native:apigee/v1:getAppGroupAppKey", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetAppGroupAppKeyArgs struct {
	AppId          string `pulumi:"appId"`
	AppgroupId     string `pulumi:"appgroupId"`
	KeyId          string `pulumi:"keyId"`
	OrganizationId string `pulumi:"organizationId"`
}

type GetAppGroupAppKeyResult struct {
	// List of API products and its status for which the credential can be used. **Note**: Use UpdateAppGroupAppKeyApiProductRequest API to make the association after the consumer key and secret are created.
	ApiProducts []GoogleCloudApigeeV1APIProductAssociationResponse `pulumi:"apiProducts"`
	// List of attributes associated with the credential.
	Attributes []GoogleCloudApigeeV1AttributeResponse `pulumi:"attributes"`
	// Immutable. Consumer key.
	ConsumerKey string `pulumi:"consumerKey"`
	// Secret key.
	ConsumerSecret string `pulumi:"consumerSecret"`
	// Time the AppGroup app expires in milliseconds since epoch.
	ExpiresAt string `pulumi:"expiresAt"`
	// Immutable. Expiration time, in seconds, for the consumer key. If not set or left to the default value of `-1`, the API key never expires. The expiration time can't be updated after it is set.
	ExpiresInSeconds string `pulumi:"expiresInSeconds"`
	// Time the AppGroup app was created in milliseconds since epoch.
	IssuedAt string `pulumi:"issuedAt"`
	// Scopes to apply to the app. The specified scope names must already be defined for the API product that you associate with the app.
	Scopes []string `pulumi:"scopes"`
	// Status of the credential. Valid values include `approved` or `revoked`.
	Status string `pulumi:"status"`
}

func GetAppGroupAppKeyOutput(ctx *pulumi.Context, args GetAppGroupAppKeyOutputArgs, opts ...pulumi.InvokeOption) GetAppGroupAppKeyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetAppGroupAppKeyResult, error) {
			args := v.(GetAppGroupAppKeyArgs)
			r, err := GetAppGroupAppKey(ctx, &args, opts...)
			var s GetAppGroupAppKeyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetAppGroupAppKeyResultOutput)
}

type GetAppGroupAppKeyOutputArgs struct {
	AppId          pulumi.StringInput `pulumi:"appId"`
	AppgroupId     pulumi.StringInput `pulumi:"appgroupId"`
	KeyId          pulumi.StringInput `pulumi:"keyId"`
	OrganizationId pulumi.StringInput `pulumi:"organizationId"`
}

func (GetAppGroupAppKeyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppGroupAppKeyArgs)(nil)).Elem()
}

type GetAppGroupAppKeyResultOutput struct{ *pulumi.OutputState }

func (GetAppGroupAppKeyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetAppGroupAppKeyResult)(nil)).Elem()
}

func (o GetAppGroupAppKeyResultOutput) ToGetAppGroupAppKeyResultOutput() GetAppGroupAppKeyResultOutput {
	return o
}

func (o GetAppGroupAppKeyResultOutput) ToGetAppGroupAppKeyResultOutputWithContext(ctx context.Context) GetAppGroupAppKeyResultOutput {
	return o
}

// List of API products and its status for which the credential can be used. **Note**: Use UpdateAppGroupAppKeyApiProductRequest API to make the association after the consumer key and secret are created.
func (o GetAppGroupAppKeyResultOutput) ApiProducts() GoogleCloudApigeeV1APIProductAssociationResponseArrayOutput {
	return o.ApplyT(func(v GetAppGroupAppKeyResult) []GoogleCloudApigeeV1APIProductAssociationResponse {
		return v.ApiProducts
	}).(GoogleCloudApigeeV1APIProductAssociationResponseArrayOutput)
}

// List of attributes associated with the credential.
func (o GetAppGroupAppKeyResultOutput) Attributes() GoogleCloudApigeeV1AttributeResponseArrayOutput {
	return o.ApplyT(func(v GetAppGroupAppKeyResult) []GoogleCloudApigeeV1AttributeResponse { return v.Attributes }).(GoogleCloudApigeeV1AttributeResponseArrayOutput)
}

// Immutable. Consumer key.
func (o GetAppGroupAppKeyResultOutput) ConsumerKey() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppGroupAppKeyResult) string { return v.ConsumerKey }).(pulumi.StringOutput)
}

// Secret key.
func (o GetAppGroupAppKeyResultOutput) ConsumerSecret() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppGroupAppKeyResult) string { return v.ConsumerSecret }).(pulumi.StringOutput)
}

// Time the AppGroup app expires in milliseconds since epoch.
func (o GetAppGroupAppKeyResultOutput) ExpiresAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppGroupAppKeyResult) string { return v.ExpiresAt }).(pulumi.StringOutput)
}

// Immutable. Expiration time, in seconds, for the consumer key. If not set or left to the default value of `-1`, the API key never expires. The expiration time can't be updated after it is set.
func (o GetAppGroupAppKeyResultOutput) ExpiresInSeconds() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppGroupAppKeyResult) string { return v.ExpiresInSeconds }).(pulumi.StringOutput)
}

// Time the AppGroup app was created in milliseconds since epoch.
func (o GetAppGroupAppKeyResultOutput) IssuedAt() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppGroupAppKeyResult) string { return v.IssuedAt }).(pulumi.StringOutput)
}

// Scopes to apply to the app. The specified scope names must already be defined for the API product that you associate with the app.
func (o GetAppGroupAppKeyResultOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetAppGroupAppKeyResult) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

// Status of the credential. Valid values include `approved` or `revoked`.
func (o GetAppGroupAppKeyResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v GetAppGroupAppKeyResult) string { return v.Status }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetAppGroupAppKeyResultOutput{})
}