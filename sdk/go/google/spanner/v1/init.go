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
	case "google-native:spanner/v1:Backup":
		r = &Backup{}
	case "google-native:spanner/v1:Database":
		r = &Database{}
	case "google-native:spanner/v1:Instance":
		r = &Instance{}
	case "google-native:spanner/v1:InstanceBackupIamBinding":
		r = &InstanceBackupIamBinding{}
	case "google-native:spanner/v1:InstanceBackupIamMember":
		r = &InstanceBackupIamMember{}
	case "google-native:spanner/v1:InstanceBackupIamPolicy":
		r = &InstanceBackupIamPolicy{}
	case "google-native:spanner/v1:InstanceConfig":
		r = &InstanceConfig{}
	case "google-native:spanner/v1:InstanceDatabaseIamBinding":
		r = &InstanceDatabaseIamBinding{}
	case "google-native:spanner/v1:InstanceDatabaseIamMember":
		r = &InstanceDatabaseIamMember{}
	case "google-native:spanner/v1:InstanceDatabaseIamPolicy":
		r = &InstanceDatabaseIamPolicy{}
	case "google-native:spanner/v1:InstanceIamBinding":
		r = &InstanceIamBinding{}
	case "google-native:spanner/v1:InstanceIamMember":
		r = &InstanceIamMember{}
	case "google-native:spanner/v1:InstanceIamPolicy":
		r = &InstanceIamPolicy{}
	case "google-native:spanner/v1:Session":
		r = &Session{}
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
		"spanner/v1",
		&module{version},
	)
}
