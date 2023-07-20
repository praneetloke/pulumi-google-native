// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3beta1

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
	case "google-native:dialogflow/v3beta1:Agent":
		r = &Agent{}
	case "google-native:dialogflow/v3beta1:EntityType":
		r = &EntityType{}
	case "google-native:dialogflow/v3beta1:Environment":
		r = &Environment{}
	case "google-native:dialogflow/v3beta1:Experiment":
		r = &Experiment{}
	case "google-native:dialogflow/v3beta1:Flow":
		r = &Flow{}
	case "google-native:dialogflow/v3beta1:Intent":
		r = &Intent{}
	case "google-native:dialogflow/v3beta1:Page":
		r = &Page{}
	case "google-native:dialogflow/v3beta1:SecuritySetting":
		r = &SecuritySetting{}
	case "google-native:dialogflow/v3beta1:SessionEntityType":
		r = &SessionEntityType{}
	case "google-native:dialogflow/v3beta1:TestCase":
		r = &TestCase{}
	case "google-native:dialogflow/v3beta1:TransitionRouteGroup":
		r = &TransitionRouteGroup{}
	case "google-native:dialogflow/v3beta1:Version":
		r = &Version{}
	case "google-native:dialogflow/v3beta1:Webhook":
		r = &Webhook{}
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
		"dialogflow/v3beta1",
		&module{version},
	)
}
