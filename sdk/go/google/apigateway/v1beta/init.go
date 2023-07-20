// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta

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
	case "google-native:apigateway/v1beta:Api":
		r = &Api{}
	case "google-native:apigateway/v1beta:ApiConfigIamBinding":
		r = &ApiConfigIamBinding{}
	case "google-native:apigateway/v1beta:ApiConfigIamMember":
		r = &ApiConfigIamMember{}
	case "google-native:apigateway/v1beta:ApiConfigIamPolicy":
		r = &ApiConfigIamPolicy{}
	case "google-native:apigateway/v1beta:ApiIamBinding":
		r = &ApiIamBinding{}
	case "google-native:apigateway/v1beta:ApiIamMember":
		r = &ApiIamMember{}
	case "google-native:apigateway/v1beta:ApiIamPolicy":
		r = &ApiIamPolicy{}
	case "google-native:apigateway/v1beta:Config":
		r = &Config{}
	case "google-native:apigateway/v1beta:Gateway":
		r = &Gateway{}
	case "google-native:apigateway/v1beta:GatewayIamBinding":
		r = &GatewayIamBinding{}
	case "google-native:apigateway/v1beta:GatewayIamMember":
		r = &GatewayIamMember{}
	case "google-native:apigateway/v1beta:GatewayIamPolicy":
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
		"apigateway/v1beta",
		&module{version},
	)
}
