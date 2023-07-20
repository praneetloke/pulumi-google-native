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
	case "google-native:certificatemanager/v1:Certificate":
		r = &Certificate{}
	case "google-native:certificatemanager/v1:CertificateIssuanceConfig":
		r = &CertificateIssuanceConfig{}
	case "google-native:certificatemanager/v1:CertificateMap":
		r = &CertificateMap{}
	case "google-native:certificatemanager/v1:CertificateMapEntry":
		r = &CertificateMapEntry{}
	case "google-native:certificatemanager/v1:DnsAuthorization":
		r = &DnsAuthorization{}
	case "google-native:certificatemanager/v1:TrustConfig":
		r = &TrustConfig{}
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
		"certificatemanager/v1",
		&module{version},
	)
}
