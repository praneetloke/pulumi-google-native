// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "google-native:apigateway/v1:Api":
		r = &Api{}
	case "google-native:apigateway/v1:ApiConfigIamBinding":
		r = &ApiConfigIamBinding{}
	case "google-native:apigateway/v1:ApiConfigIamMember":
		r = &ApiConfigIamMember{}
	case "google-native:apigateway/v1:ApiConfigIamPolicy":
		r = &ApiConfigIamPolicy{}
	case "google-native:apigateway/v1:ApiIamBinding":
		r = &ApiIamBinding{}
	case "google-native:apigateway/v1:ApiIamMember":
		r = &ApiIamMember{}
	case "google-native:apigateway/v1:ApiIamPolicy":
		r = &ApiIamPolicy{}
	case "google-native:apigateway/v1:Config":
		r = &Config{}
	case "google-native:apigateway/v1:Gateway":
		r = &Gateway{}
	case "google-native:apigateway/v1:GatewayIamBinding":
		r = &GatewayIamBinding{}
	case "google-native:apigateway/v1:GatewayIamMember":
		r = &GatewayIamMember{}
	case "google-native:apigateway/v1:GatewayIamPolicy":
		r = &GatewayIamPolicy{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"google-native",
		"apigateway/v1",
		&module{version},
	)
}
