// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1alpha1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new `Note`.
type Note struct {
	pulumi.CustomResourceState

	// A note describing an attestation role.
	AttestationAuthority AttestationAuthorityResponseOutput `pulumi:"attestationAuthority"`
	// A note describing a base image.
	BaseImage BasisResponseOutput `pulumi:"baseImage"`
	// Build provenance type for a verifiable build.
	BuildType BuildTypeResponseOutput `pulumi:"buildType"`
	// A note describing a compliance check.
	Compliance ComplianceNoteResponseOutput `pulumi:"compliance"`
	// The time this note was created. This field can be used as a filter in list requests.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// A note describing something that can be deployed.
	Deployable DeployableResponseOutput `pulumi:"deployable"`
	// A note describing a provider/analysis type.
	Discovery DiscoveryResponseOutput `pulumi:"discovery"`
	// A note describing a dsse attestation note.
	DsseAttestation DSSEAttestationNoteResponseOutput `pulumi:"dsseAttestation"`
	// Time of expiration for this note, null if note does not expire.
	ExpirationTime pulumi.StringOutput `pulumi:"expirationTime"`
	// This explicitly denotes which kind of note is specified. This field can be used as a filter in list requests.
	Kind pulumi.StringOutput `pulumi:"kind"`
	// A detailed description of this `Note`.
	LongDescription pulumi.StringOutput `pulumi:"longDescription"`
	// The name of the project. Should be of the form "providers/{provider_id}". @Deprecated
	Name pulumi.StringOutput `pulumi:"name"`
	// The ID to use for this note.
	NoteId pulumi.StringPtrOutput `pulumi:"noteId"`
	// A note describing a package hosted by various package managers.
	Package PackageResponseOutput `pulumi:"package"`
	Project pulumi.StringOutput   `pulumi:"project"`
	// URLs associated with this note
	RelatedUrl RelatedUrlResponseArrayOutput `pulumi:"relatedUrl"`
	// A note describing a software bill of materials.
	Sbom DocumentNoteResponseOutput `pulumi:"sbom"`
	// A note describing a reference to an SBOM.
	SbomReference SBOMReferenceNoteResponseOutput `pulumi:"sbomReference"`
	// A one sentence description of this `Note`.
	ShortDescription pulumi.StringOutput `pulumi:"shortDescription"`
	// A note describing an SPDX File.
	SpdxFile FileNoteResponseOutput `pulumi:"spdxFile"`
	// A note describing an SPDX Package.
	SpdxPackage PackageInfoNoteResponseOutput `pulumi:"spdxPackage"`
	// A note describing a relationship between SPDX elements.
	SpdxRelationship RelationshipNoteResponseOutput `pulumi:"spdxRelationship"`
	// The time this note was last updated. This field can be used as a filter in list requests.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// A note describing an upgrade.
	Upgrade UpgradeNoteResponseOutput `pulumi:"upgrade"`
	// A note describing a vulnerability assessment.
	VulnerabilityAssessment VulnerabilityAssessmentNoteResponseOutput `pulumi:"vulnerabilityAssessment"`
	// A package vulnerability type of note.
	VulnerabilityType VulnerabilityTypeResponseOutput `pulumi:"vulnerabilityType"`
}

// NewNote registers a new resource with the given unique name, arguments, and options.
func NewNote(ctx *pulumi.Context,
	name string, args *NoteArgs, opts ...pulumi.ResourceOption) (*Note, error) {
	if args == nil {
		args = &NoteArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Note
	err := ctx.RegisterResource("google-native:containeranalysis/v1alpha1:Note", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNote gets an existing Note resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNote(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NoteState, opts ...pulumi.ResourceOption) (*Note, error) {
	var resource Note
	err := ctx.ReadResource("google-native:containeranalysis/v1alpha1:Note", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Note resources.
type noteState struct {
}

type NoteState struct {
}

func (NoteState) ElementType() reflect.Type {
	return reflect.TypeOf((*noteState)(nil)).Elem()
}

type noteArgs struct {
	// A note describing an attestation role.
	AttestationAuthority *AttestationAuthority `pulumi:"attestationAuthority"`
	// A note describing a base image.
	BaseImage *Basis `pulumi:"baseImage"`
	// Build provenance type for a verifiable build.
	BuildType *BuildType `pulumi:"buildType"`
	// A note describing a compliance check.
	Compliance *ComplianceNote `pulumi:"compliance"`
	// A note describing something that can be deployed.
	Deployable *Deployable `pulumi:"deployable"`
	// A note describing a provider/analysis type.
	Discovery *Discovery `pulumi:"discovery"`
	// A note describing a dsse attestation note.
	DsseAttestation *DSSEAttestationNote `pulumi:"dsseAttestation"`
	// Time of expiration for this note, null if note does not expire.
	ExpirationTime *string `pulumi:"expirationTime"`
	// A detailed description of this `Note`.
	LongDescription *string `pulumi:"longDescription"`
	// The name of the note in the form "projects/{provider_project_id}/notes/{NOTE_ID}"
	Name *string `pulumi:"name"`
	// The ID to use for this note.
	NoteId *string `pulumi:"noteId"`
	// A note describing a package hosted by various package managers.
	Package *Package `pulumi:"package"`
	Project *string  `pulumi:"project"`
	// URLs associated with this note
	RelatedUrl []RelatedUrl `pulumi:"relatedUrl"`
	// A note describing a software bill of materials.
	Sbom *DocumentNote `pulumi:"sbom"`
	// A note describing a reference to an SBOM.
	SbomReference *SBOMReferenceNote `pulumi:"sbomReference"`
	// A one sentence description of this `Note`.
	ShortDescription *string `pulumi:"shortDescription"`
	// A note describing an SPDX File.
	SpdxFile *FileNote `pulumi:"spdxFile"`
	// A note describing an SPDX Package.
	SpdxPackage *PackageInfoNote `pulumi:"spdxPackage"`
	// A note describing a relationship between SPDX elements.
	SpdxRelationship *RelationshipNote `pulumi:"spdxRelationship"`
	// A note describing an upgrade.
	Upgrade *UpgradeNote `pulumi:"upgrade"`
	// A note describing a vulnerability assessment.
	VulnerabilityAssessment *VulnerabilityAssessmentNote `pulumi:"vulnerabilityAssessment"`
	// A package vulnerability type of note.
	VulnerabilityType *VulnerabilityType `pulumi:"vulnerabilityType"`
}

// The set of arguments for constructing a Note resource.
type NoteArgs struct {
	// A note describing an attestation role.
	AttestationAuthority AttestationAuthorityPtrInput
	// A note describing a base image.
	BaseImage BasisPtrInput
	// Build provenance type for a verifiable build.
	BuildType BuildTypePtrInput
	// A note describing a compliance check.
	Compliance ComplianceNotePtrInput
	// A note describing something that can be deployed.
	Deployable DeployablePtrInput
	// A note describing a provider/analysis type.
	Discovery DiscoveryPtrInput
	// A note describing a dsse attestation note.
	DsseAttestation DSSEAttestationNotePtrInput
	// Time of expiration for this note, null if note does not expire.
	ExpirationTime pulumi.StringPtrInput
	// A detailed description of this `Note`.
	LongDescription pulumi.StringPtrInput
	// The name of the note in the form "projects/{provider_project_id}/notes/{NOTE_ID}"
	Name pulumi.StringPtrInput
	// The ID to use for this note.
	NoteId pulumi.StringPtrInput
	// A note describing a package hosted by various package managers.
	Package PackagePtrInput
	Project pulumi.StringPtrInput
	// URLs associated with this note
	RelatedUrl RelatedUrlArrayInput
	// A note describing a software bill of materials.
	Sbom DocumentNotePtrInput
	// A note describing a reference to an SBOM.
	SbomReference SBOMReferenceNotePtrInput
	// A one sentence description of this `Note`.
	ShortDescription pulumi.StringPtrInput
	// A note describing an SPDX File.
	SpdxFile FileNotePtrInput
	// A note describing an SPDX Package.
	SpdxPackage PackageInfoNotePtrInput
	// A note describing a relationship between SPDX elements.
	SpdxRelationship RelationshipNotePtrInput
	// A note describing an upgrade.
	Upgrade UpgradeNotePtrInput
	// A note describing a vulnerability assessment.
	VulnerabilityAssessment VulnerabilityAssessmentNotePtrInput
	// A package vulnerability type of note.
	VulnerabilityType VulnerabilityTypePtrInput
}

func (NoteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*noteArgs)(nil)).Elem()
}

type NoteInput interface {
	pulumi.Input

	ToNoteOutput() NoteOutput
	ToNoteOutputWithContext(ctx context.Context) NoteOutput
}

func (*Note) ElementType() reflect.Type {
	return reflect.TypeOf((**Note)(nil)).Elem()
}

func (i *Note) ToNoteOutput() NoteOutput {
	return i.ToNoteOutputWithContext(context.Background())
}

func (i *Note) ToNoteOutputWithContext(ctx context.Context) NoteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NoteOutput)
}

type NoteOutput struct{ *pulumi.OutputState }

func (NoteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Note)(nil)).Elem()
}

func (o NoteOutput) ToNoteOutput() NoteOutput {
	return o
}

func (o NoteOutput) ToNoteOutputWithContext(ctx context.Context) NoteOutput {
	return o
}

// A note describing an attestation role.
func (o NoteOutput) AttestationAuthority() AttestationAuthorityResponseOutput {
	return o.ApplyT(func(v *Note) AttestationAuthorityResponseOutput { return v.AttestationAuthority }).(AttestationAuthorityResponseOutput)
}

// A note describing a base image.
func (o NoteOutput) BaseImage() BasisResponseOutput {
	return o.ApplyT(func(v *Note) BasisResponseOutput { return v.BaseImage }).(BasisResponseOutput)
}

// Build provenance type for a verifiable build.
func (o NoteOutput) BuildType() BuildTypeResponseOutput {
	return o.ApplyT(func(v *Note) BuildTypeResponseOutput { return v.BuildType }).(BuildTypeResponseOutput)
}

// A note describing a compliance check.
func (o NoteOutput) Compliance() ComplianceNoteResponseOutput {
	return o.ApplyT(func(v *Note) ComplianceNoteResponseOutput { return v.Compliance }).(ComplianceNoteResponseOutput)
}

// The time this note was created. This field can be used as a filter in list requests.
func (o NoteOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Note) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// A note describing something that can be deployed.
func (o NoteOutput) Deployable() DeployableResponseOutput {
	return o.ApplyT(func(v *Note) DeployableResponseOutput { return v.Deployable }).(DeployableResponseOutput)
}

// A note describing a provider/analysis type.
func (o NoteOutput) Discovery() DiscoveryResponseOutput {
	return o.ApplyT(func(v *Note) DiscoveryResponseOutput { return v.Discovery }).(DiscoveryResponseOutput)
}

// A note describing a dsse attestation note.
func (o NoteOutput) DsseAttestation() DSSEAttestationNoteResponseOutput {
	return o.ApplyT(func(v *Note) DSSEAttestationNoteResponseOutput { return v.DsseAttestation }).(DSSEAttestationNoteResponseOutput)
}

// Time of expiration for this note, null if note does not expire.
func (o NoteOutput) ExpirationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Note) pulumi.StringOutput { return v.ExpirationTime }).(pulumi.StringOutput)
}

// This explicitly denotes which kind of note is specified. This field can be used as a filter in list requests.
func (o NoteOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Note) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

// A detailed description of this `Note`.
func (o NoteOutput) LongDescription() pulumi.StringOutput {
	return o.ApplyT(func(v *Note) pulumi.StringOutput { return v.LongDescription }).(pulumi.StringOutput)
}

// The name of the project. Should be of the form "providers/{provider_id}". @Deprecated
func (o NoteOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Note) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// The ID to use for this note.
func (o NoteOutput) NoteId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Note) pulumi.StringPtrOutput { return v.NoteId }).(pulumi.StringPtrOutput)
}

// A note describing a package hosted by various package managers.
func (o NoteOutput) Package() PackageResponseOutput {
	return o.ApplyT(func(v *Note) PackageResponseOutput { return v.Package }).(PackageResponseOutput)
}

func (o NoteOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Note) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// URLs associated with this note
func (o NoteOutput) RelatedUrl() RelatedUrlResponseArrayOutput {
	return o.ApplyT(func(v *Note) RelatedUrlResponseArrayOutput { return v.RelatedUrl }).(RelatedUrlResponseArrayOutput)
}

// A note describing a software bill of materials.
func (o NoteOutput) Sbom() DocumentNoteResponseOutput {
	return o.ApplyT(func(v *Note) DocumentNoteResponseOutput { return v.Sbom }).(DocumentNoteResponseOutput)
}

// A note describing a reference to an SBOM.
func (o NoteOutput) SbomReference() SBOMReferenceNoteResponseOutput {
	return o.ApplyT(func(v *Note) SBOMReferenceNoteResponseOutput { return v.SbomReference }).(SBOMReferenceNoteResponseOutput)
}

// A one sentence description of this `Note`.
func (o NoteOutput) ShortDescription() pulumi.StringOutput {
	return o.ApplyT(func(v *Note) pulumi.StringOutput { return v.ShortDescription }).(pulumi.StringOutput)
}

// A note describing an SPDX File.
func (o NoteOutput) SpdxFile() FileNoteResponseOutput {
	return o.ApplyT(func(v *Note) FileNoteResponseOutput { return v.SpdxFile }).(FileNoteResponseOutput)
}

// A note describing an SPDX Package.
func (o NoteOutput) SpdxPackage() PackageInfoNoteResponseOutput {
	return o.ApplyT(func(v *Note) PackageInfoNoteResponseOutput { return v.SpdxPackage }).(PackageInfoNoteResponseOutput)
}

// A note describing a relationship between SPDX elements.
func (o NoteOutput) SpdxRelationship() RelationshipNoteResponseOutput {
	return o.ApplyT(func(v *Note) RelationshipNoteResponseOutput { return v.SpdxRelationship }).(RelationshipNoteResponseOutput)
}

// The time this note was last updated. This field can be used as a filter in list requests.
func (o NoteOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Note) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// A note describing an upgrade.
func (o NoteOutput) Upgrade() UpgradeNoteResponseOutput {
	return o.ApplyT(func(v *Note) UpgradeNoteResponseOutput { return v.Upgrade }).(UpgradeNoteResponseOutput)
}

// A note describing a vulnerability assessment.
func (o NoteOutput) VulnerabilityAssessment() VulnerabilityAssessmentNoteResponseOutput {
	return o.ApplyT(func(v *Note) VulnerabilityAssessmentNoteResponseOutput { return v.VulnerabilityAssessment }).(VulnerabilityAssessmentNoteResponseOutput)
}

// A package vulnerability type of note.
func (o NoteOutput) VulnerabilityType() VulnerabilityTypeResponseOutput {
	return o.ApplyT(func(v *Note) VulnerabilityTypeResponseOutput { return v.VulnerabilityType }).(VulnerabilityTypeResponseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NoteInput)(nil)).Elem(), &Note{})
	pulumi.RegisterOutputType(NoteOutput{})
}
