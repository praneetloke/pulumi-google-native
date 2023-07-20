// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

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
	case "google-native:artifactregistry/v1beta1:Repository":
		r = &Repository{}
	case "google-native:artifactregistry/v1beta1:RepositoryIamBinding":
		r = &RepositoryIamBinding{}
	case "google-native:artifactregistry/v1beta1:RepositoryIamMember":
		r = &RepositoryIamMember{}
	case "google-native:artifactregistry/v1beta1:RepositoryIamPolicy":
		r = &RepositoryIamPolicy{}
	case "google-native:artifactregistry/v1beta1:Tag":
		r = &Tag{}
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
		"artifactregistry/v1beta1",
		&module{version},
	)
}
