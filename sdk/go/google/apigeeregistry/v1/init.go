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
	case "google-native:apigeeregistry/v1:Api":
		r = &Api{}
	case "google-native:apigeeregistry/v1:ApiArtifactIamBinding":
		r = &ApiArtifactIamBinding{}
	case "google-native:apigeeregistry/v1:ApiArtifactIamMember":
		r = &ApiArtifactIamMember{}
	case "google-native:apigeeregistry/v1:ApiArtifactIamPolicy":
		r = &ApiArtifactIamPolicy{}
	case "google-native:apigeeregistry/v1:ApiDeploymentIamBinding":
		r = &ApiDeploymentIamBinding{}
	case "google-native:apigeeregistry/v1:ApiDeploymentIamMember":
		r = &ApiDeploymentIamMember{}
	case "google-native:apigeeregistry/v1:ApiDeploymentIamPolicy":
		r = &ApiDeploymentIamPolicy{}
	case "google-native:apigeeregistry/v1:ApiIamBinding":
		r = &ApiIamBinding{}
	case "google-native:apigeeregistry/v1:ApiIamMember":
		r = &ApiIamMember{}
	case "google-native:apigeeregistry/v1:ApiIamPolicy":
		r = &ApiIamPolicy{}
	case "google-native:apigeeregistry/v1:ApiVersionArtifactIamBinding":
		r = &ApiVersionArtifactIamBinding{}
	case "google-native:apigeeregistry/v1:ApiVersionArtifactIamMember":
		r = &ApiVersionArtifactIamMember{}
	case "google-native:apigeeregistry/v1:ApiVersionArtifactIamPolicy":
		r = &ApiVersionArtifactIamPolicy{}
	case "google-native:apigeeregistry/v1:ApiVersionIamBinding":
		r = &ApiVersionIamBinding{}
	case "google-native:apigeeregistry/v1:ApiVersionIamMember":
		r = &ApiVersionIamMember{}
	case "google-native:apigeeregistry/v1:ApiVersionIamPolicy":
		r = &ApiVersionIamPolicy{}
	case "google-native:apigeeregistry/v1:ApiVersionSpecArtifactIamBinding":
		r = &ApiVersionSpecArtifactIamBinding{}
	case "google-native:apigeeregistry/v1:ApiVersionSpecArtifactIamMember":
		r = &ApiVersionSpecArtifactIamMember{}
	case "google-native:apigeeregistry/v1:ApiVersionSpecArtifactIamPolicy":
		r = &ApiVersionSpecArtifactIamPolicy{}
	case "google-native:apigeeregistry/v1:ApiVersionSpecIamBinding":
		r = &ApiVersionSpecIamBinding{}
	case "google-native:apigeeregistry/v1:ApiVersionSpecIamMember":
		r = &ApiVersionSpecIamMember{}
	case "google-native:apigeeregistry/v1:ApiVersionSpecIamPolicy":
		r = &ApiVersionSpecIamPolicy{}
	case "google-native:apigeeregistry/v1:Artifact":
		r = &Artifact{}
	case "google-native:apigeeregistry/v1:ArtifactIamBinding":
		r = &ArtifactIamBinding{}
	case "google-native:apigeeregistry/v1:ArtifactIamMember":
		r = &ArtifactIamMember{}
	case "google-native:apigeeregistry/v1:ArtifactIamPolicy":
		r = &ArtifactIamPolicy{}
	case "google-native:apigeeregistry/v1:Deployment":
		r = &Deployment{}
	case "google-native:apigeeregistry/v1:DeploymentArtifact":
		r = &DeploymentArtifact{}
	case "google-native:apigeeregistry/v1:Instance":
		r = &Instance{}
	case "google-native:apigeeregistry/v1:InstanceIamBinding":
		r = &InstanceIamBinding{}
	case "google-native:apigeeregistry/v1:InstanceIamMember":
		r = &InstanceIamMember{}
	case "google-native:apigeeregistry/v1:InstanceIamPolicy":
		r = &InstanceIamPolicy{}
	case "google-native:apigeeregistry/v1:RuntimeIamBinding":
		r = &RuntimeIamBinding{}
	case "google-native:apigeeregistry/v1:RuntimeIamMember":
		r = &RuntimeIamMember{}
	case "google-native:apigeeregistry/v1:RuntimeIamPolicy":
		r = &RuntimeIamPolicy{}
	case "google-native:apigeeregistry/v1:Spec":
		r = &Spec{}
	case "google-native:apigeeregistry/v1:Version":
		r = &Version{}
	case "google-native:apigeeregistry/v1:VersionArtifact":
		r = &VersionArtifact{}
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
		"apigeeregistry/v1",
		&module{version},
	)
}
