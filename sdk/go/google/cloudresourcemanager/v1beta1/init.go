// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

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
	case "google-native:cloudresourcemanager/v1beta1:OrganizationIamPolicy":
		r = &OrganizationIamPolicy{}
	case "google-native:cloudresourcemanager/v1beta1:Project":
		r = &Project{}
	case "google-native:cloudresourcemanager/v1beta1:ProjectIamPolicy":
		r = &ProjectIamPolicy{}
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
		"cloudresourcemanager/v1beta1",
		&module{version},
	)
}
