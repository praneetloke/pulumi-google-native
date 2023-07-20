// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

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
	case "google-native:networkconnectivity/v1alpha1:Hub":
		r = &Hub{}
	case "google-native:networkconnectivity/v1alpha1:HubIamBinding":
		r = &HubIamBinding{}
	case "google-native:networkconnectivity/v1alpha1:HubIamMember":
		r = &HubIamMember{}
	case "google-native:networkconnectivity/v1alpha1:HubIamPolicy":
		r = &HubIamPolicy{}
	case "google-native:networkconnectivity/v1alpha1:InternalRange":
		r = &InternalRange{}
	case "google-native:networkconnectivity/v1alpha1:InternalRangeIamBinding":
		r = &InternalRangeIamBinding{}
	case "google-native:networkconnectivity/v1alpha1:InternalRangeIamMember":
		r = &InternalRangeIamMember{}
	case "google-native:networkconnectivity/v1alpha1:InternalRangeIamPolicy":
		r = &InternalRangeIamPolicy{}
	case "google-native:networkconnectivity/v1alpha1:Spoke":
		r = &Spoke{}
	case "google-native:networkconnectivity/v1alpha1:SpokeIamBinding":
		r = &SpokeIamBinding{}
	case "google-native:networkconnectivity/v1alpha1:SpokeIamMember":
		r = &SpokeIamMember{}
	case "google-native:networkconnectivity/v1alpha1:SpokeIamPolicy":
		r = &SpokeIamPolicy{}
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
		"networkconnectivity/v1alpha1",
		&module{version},
	)
}
