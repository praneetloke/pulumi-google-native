// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi-google-native/sdk/go/google"
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
	case "google-native:servicemanagement/v1:Service":
		r = &Service{}
	case "google-native:servicemanagement/v1:ServiceConfig":
		r = &ServiceConfig{}
	case "google-native:servicemanagement/v1:ServiceConsumerIamPolicy":
		r = &ServiceConsumerIamPolicy{}
	case "google-native:servicemanagement/v1:ServiceIamPolicy":
		r = &ServiceIamPolicy{}
	case "google-native:servicemanagement/v1:ServiceRollout":
		r = &ServiceRollout{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

func init() {
	version, err := google.PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"google-native",
		"servicemanagement/v1",
		&module{version},
	)
}
