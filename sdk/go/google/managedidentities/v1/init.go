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
	case "google-native:managedidentities/v1:Backup":
		r = &Backup{}
	case "google-native:managedidentities/v1:Domain":
		r = &Domain{}
	case "google-native:managedidentities/v1:DomainBackupIamBinding":
		r = &DomainBackupIamBinding{}
	case "google-native:managedidentities/v1:DomainBackupIamMember":
		r = &DomainBackupIamMember{}
	case "google-native:managedidentities/v1:DomainBackupIamPolicy":
		r = &DomainBackupIamPolicy{}
	case "google-native:managedidentities/v1:DomainIamBinding":
		r = &DomainIamBinding{}
	case "google-native:managedidentities/v1:DomainIamMember":
		r = &DomainIamMember{}
	case "google-native:managedidentities/v1:DomainIamPolicy":
		r = &DomainIamPolicy{}
	case "google-native:managedidentities/v1:Peering":
		r = &Peering{}
	case "google-native:managedidentities/v1:PeeringIamBinding":
		r = &PeeringIamBinding{}
	case "google-native:managedidentities/v1:PeeringIamMember":
		r = &PeeringIamMember{}
	case "google-native:managedidentities/v1:PeeringIamPolicy":
		r = &PeeringIamPolicy{}
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
		"managedidentities/v1",
		&module{version},
	)
}
