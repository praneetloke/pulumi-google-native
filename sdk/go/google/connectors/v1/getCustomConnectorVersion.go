// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets details of a single CustomConnectorVersion.
func LookupCustomConnectorVersion(ctx *pulumi.Context, args *LookupCustomConnectorVersionArgs, opts ...pulumi.InvokeOption) (*LookupCustomConnectorVersionResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupCustomConnectorVersionResult
	err := ctx.Invoke("google-native:connectors/v1:getCustomConnectorVersion", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCustomConnectorVersionArgs struct {
	CustomConnectorId        string  `pulumi:"customConnectorId"`
	CustomConnectorVersionId string  `pulumi:"customConnectorVersionId"`
	Project                  *string `pulumi:"project"`
}

type LookupCustomConnectorVersionResult struct {
	// Configuration for establishing the authentication to the connector destination.
	AuthConfig AuthConfigResponse `pulumi:"authConfig"`
	// Created time.
	CreateTime string `pulumi:"createTime"`
	// Configuration of the customConnector's destination.
	DestinationConfig DestinationConfigResponse `pulumi:"destinationConfig"`
	// Optional. Whether to enable backend destination config. This is the backend server that the connector connects to.
	EnableBackendDestinationConfig bool `pulumi:"enableBackendDestinationConfig"`
	// Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
	Labels map[string]string `pulumi:"labels"`
	// Identifier. Resource name of the Version. Format: projects/{project}/locations/{location}/customConnectors/{custom_connector}/customConnectorVersions/{custom_connector_version}
	Name string `pulumi:"name"`
	// Service account needed for runtime plane to access Custom Connector secrets.
	ServiceAccount string `pulumi:"serviceAccount"`
	// Optional. Location of the custom connector spec.
	SpecLocation string `pulumi:"specLocation"`
	// Updated time.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupCustomConnectorVersionOutput(ctx *pulumi.Context, args LookupCustomConnectorVersionOutputArgs, opts ...pulumi.InvokeOption) LookupCustomConnectorVersionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCustomConnectorVersionResult, error) {
			args := v.(LookupCustomConnectorVersionArgs)
			r, err := LookupCustomConnectorVersion(ctx, &args, opts...)
			var s LookupCustomConnectorVersionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCustomConnectorVersionResultOutput)
}

type LookupCustomConnectorVersionOutputArgs struct {
	CustomConnectorId        pulumi.StringInput    `pulumi:"customConnectorId"`
	CustomConnectorVersionId pulumi.StringInput    `pulumi:"customConnectorVersionId"`
	Project                  pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupCustomConnectorVersionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomConnectorVersionArgs)(nil)).Elem()
}

type LookupCustomConnectorVersionResultOutput struct{ *pulumi.OutputState }

func (LookupCustomConnectorVersionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCustomConnectorVersionResult)(nil)).Elem()
}

func (o LookupCustomConnectorVersionResultOutput) ToLookupCustomConnectorVersionResultOutput() LookupCustomConnectorVersionResultOutput {
	return o
}

func (o LookupCustomConnectorVersionResultOutput) ToLookupCustomConnectorVersionResultOutputWithContext(ctx context.Context) LookupCustomConnectorVersionResultOutput {
	return o
}

// Configuration for establishing the authentication to the connector destination.
func (o LookupCustomConnectorVersionResultOutput) AuthConfig() AuthConfigResponseOutput {
	return o.ApplyT(func(v LookupCustomConnectorVersionResult) AuthConfigResponse { return v.AuthConfig }).(AuthConfigResponseOutput)
}

// Created time.
func (o LookupCustomConnectorVersionResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomConnectorVersionResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// Configuration of the customConnector's destination.
func (o LookupCustomConnectorVersionResultOutput) DestinationConfig() DestinationConfigResponseOutput {
	return o.ApplyT(func(v LookupCustomConnectorVersionResult) DestinationConfigResponse { return v.DestinationConfig }).(DestinationConfigResponseOutput)
}

// Optional. Whether to enable backend destination config. This is the backend server that the connector connects to.
func (o LookupCustomConnectorVersionResultOutput) EnableBackendDestinationConfig() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupCustomConnectorVersionResult) bool { return v.EnableBackendDestinationConfig }).(pulumi.BoolOutput)
}

// Optional. Resource labels to represent user-provided metadata. Refer to cloud documentation on labels for more details. https://cloud.google.com/compute/docs/labeling-resources
func (o LookupCustomConnectorVersionResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupCustomConnectorVersionResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Identifier. Resource name of the Version. Format: projects/{project}/locations/{location}/customConnectors/{custom_connector}/customConnectorVersions/{custom_connector_version}
func (o LookupCustomConnectorVersionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomConnectorVersionResult) string { return v.Name }).(pulumi.StringOutput)
}

// Service account needed for runtime plane to access Custom Connector secrets.
func (o LookupCustomConnectorVersionResultOutput) ServiceAccount() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomConnectorVersionResult) string { return v.ServiceAccount }).(pulumi.StringOutput)
}

// Optional. Location of the custom connector spec.
func (o LookupCustomConnectorVersionResultOutput) SpecLocation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomConnectorVersionResult) string { return v.SpecLocation }).(pulumi.StringOutput)
}

// Updated time.
func (o LookupCustomConnectorVersionResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCustomConnectorVersionResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCustomConnectorVersionResultOutput{})
}