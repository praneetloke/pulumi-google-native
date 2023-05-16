// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets the specified note.
func LookupNote(ctx *pulumi.Context, args *LookupNoteArgs, opts ...pulumi.InvokeOption) (*LookupNoteResult, error) {
	var rv LookupNoteResult
	err := ctx.Invoke("google-native:containeranalysis/v1:getNote", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNoteArgs struct {
	NoteId  string  `pulumi:"noteId"`
	Project *string `pulumi:"project"`
}

type LookupNoteResult struct {
	// A note describing an attestation role.
	Attestation AttestationNoteResponse `pulumi:"attestation"`
	// A note describing build provenance for a verifiable build.
	Build BuildNoteResponse `pulumi:"build"`
	// A note describing a compliance check.
	Compliance ComplianceNoteResponse `pulumi:"compliance"`
	// The time this note was created. This field can be used as a filter in list requests.
	CreateTime string `pulumi:"createTime"`
	// A note describing something that can be deployed.
	Deployment DeploymentNoteResponse `pulumi:"deployment"`
	// A note describing the initial analysis of a resource.
	Discovery DiscoveryNoteResponse `pulumi:"discovery"`
	// A note describing a dsse attestation note.
	DsseAttestation DSSEAttestationNoteResponse `pulumi:"dsseAttestation"`
	// Time of expiration for this note. Empty if note does not expire.
	ExpirationTime string `pulumi:"expirationTime"`
	// A note describing a base image.
	Image ImageNoteResponse `pulumi:"image"`
	// The type of analysis. This field can be used as a filter in list requests.
	Kind string `pulumi:"kind"`
	// A detailed description of this note.
	LongDescription string `pulumi:"longDescription"`
	// The name of the note in the form of `projects/[PROVIDER_ID]/notes/[NOTE_ID]`.
	Name string `pulumi:"name"`
	// A note describing a package hosted by various package managers.
	Package PackageNoteResponse `pulumi:"package"`
	// Other notes related to this note.
	RelatedNoteNames []string `pulumi:"relatedNoteNames"`
	// URLs associated with this note.
	RelatedUrl []RelatedUrlResponse `pulumi:"relatedUrl"`
	// A note describing an SBOM reference.
	SbomReference SBOMReferenceNoteResponse `pulumi:"sbomReference"`
	// A one sentence description of this note.
	ShortDescription string `pulumi:"shortDescription"`
	// The time this note was last updated. This field can be used as a filter in list requests.
	UpdateTime string `pulumi:"updateTime"`
	// A note describing available package upgrades.
	Upgrade UpgradeNoteResponse `pulumi:"upgrade"`
	// A note describing a package vulnerability.
	Vulnerability VulnerabilityNoteResponse `pulumi:"vulnerability"`
	// A note describing a vulnerability assessment.
	VulnerabilityAssessment VulnerabilityAssessmentNoteResponse `pulumi:"vulnerabilityAssessment"`
}

func LookupNoteOutput(ctx *pulumi.Context, args LookupNoteOutputArgs, opts ...pulumi.InvokeOption) LookupNoteResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNoteResult, error) {
			args := v.(LookupNoteArgs)
			r, err := LookupNote(ctx, &args, opts...)
			var s LookupNoteResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNoteResultOutput)
}

type LookupNoteOutputArgs struct {
	NoteId  pulumi.StringInput    `pulumi:"noteId"`
	Project pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupNoteOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNoteArgs)(nil)).Elem()
}

type LookupNoteResultOutput struct{ *pulumi.OutputState }

func (LookupNoteResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNoteResult)(nil)).Elem()
}

func (o LookupNoteResultOutput) ToLookupNoteResultOutput() LookupNoteResultOutput {
	return o
}

func (o LookupNoteResultOutput) ToLookupNoteResultOutputWithContext(ctx context.Context) LookupNoteResultOutput {
	return o
}

// A note describing an attestation role.
func (o LookupNoteResultOutput) Attestation() AttestationNoteResponseOutput {
	return o.ApplyT(func(v LookupNoteResult) AttestationNoteResponse { return v.Attestation }).(AttestationNoteResponseOutput)
}

// A note describing build provenance for a verifiable build.
func (o LookupNoteResultOutput) Build() BuildNoteResponseOutput {
	return o.ApplyT(func(v LookupNoteResult) BuildNoteResponse { return v.Build }).(BuildNoteResponseOutput)
}

// A note describing a compliance check.
func (o LookupNoteResultOutput) Compliance() ComplianceNoteResponseOutput {
	return o.ApplyT(func(v LookupNoteResult) ComplianceNoteResponse { return v.Compliance }).(ComplianceNoteResponseOutput)
}

// The time this note was created. This field can be used as a filter in list requests.
func (o LookupNoteResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNoteResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// A note describing something that can be deployed.
func (o LookupNoteResultOutput) Deployment() DeploymentNoteResponseOutput {
	return o.ApplyT(func(v LookupNoteResult) DeploymentNoteResponse { return v.Deployment }).(DeploymentNoteResponseOutput)
}

// A note describing the initial analysis of a resource.
func (o LookupNoteResultOutput) Discovery() DiscoveryNoteResponseOutput {
	return o.ApplyT(func(v LookupNoteResult) DiscoveryNoteResponse { return v.Discovery }).(DiscoveryNoteResponseOutput)
}

// A note describing a dsse attestation note.
func (o LookupNoteResultOutput) DsseAttestation() DSSEAttestationNoteResponseOutput {
	return o.ApplyT(func(v LookupNoteResult) DSSEAttestationNoteResponse { return v.DsseAttestation }).(DSSEAttestationNoteResponseOutput)
}

// Time of expiration for this note. Empty if note does not expire.
func (o LookupNoteResultOutput) ExpirationTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNoteResult) string { return v.ExpirationTime }).(pulumi.StringOutput)
}

// A note describing a base image.
func (o LookupNoteResultOutput) Image() ImageNoteResponseOutput {
	return o.ApplyT(func(v LookupNoteResult) ImageNoteResponse { return v.Image }).(ImageNoteResponseOutput)
}

// The type of analysis. This field can be used as a filter in list requests.
func (o LookupNoteResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNoteResult) string { return v.Kind }).(pulumi.StringOutput)
}

// A detailed description of this note.
func (o LookupNoteResultOutput) LongDescription() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNoteResult) string { return v.LongDescription }).(pulumi.StringOutput)
}

// The name of the note in the form of `projects/[PROVIDER_ID]/notes/[NOTE_ID]`.
func (o LookupNoteResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNoteResult) string { return v.Name }).(pulumi.StringOutput)
}

// A note describing a package hosted by various package managers.
func (o LookupNoteResultOutput) Package() PackageNoteResponseOutput {
	return o.ApplyT(func(v LookupNoteResult) PackageNoteResponse { return v.Package }).(PackageNoteResponseOutput)
}

// Other notes related to this note.
func (o LookupNoteResultOutput) RelatedNoteNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupNoteResult) []string { return v.RelatedNoteNames }).(pulumi.StringArrayOutput)
}

// URLs associated with this note.
func (o LookupNoteResultOutput) RelatedUrl() RelatedUrlResponseArrayOutput {
	return o.ApplyT(func(v LookupNoteResult) []RelatedUrlResponse { return v.RelatedUrl }).(RelatedUrlResponseArrayOutput)
}

// A note describing an SBOM reference.
func (o LookupNoteResultOutput) SbomReference() SBOMReferenceNoteResponseOutput {
	return o.ApplyT(func(v LookupNoteResult) SBOMReferenceNoteResponse { return v.SbomReference }).(SBOMReferenceNoteResponseOutput)
}

// A one sentence description of this note.
func (o LookupNoteResultOutput) ShortDescription() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNoteResult) string { return v.ShortDescription }).(pulumi.StringOutput)
}

// The time this note was last updated. This field can be used as a filter in list requests.
func (o LookupNoteResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNoteResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

// A note describing available package upgrades.
func (o LookupNoteResultOutput) Upgrade() UpgradeNoteResponseOutput {
	return o.ApplyT(func(v LookupNoteResult) UpgradeNoteResponse { return v.Upgrade }).(UpgradeNoteResponseOutput)
}

// A note describing a package vulnerability.
func (o LookupNoteResultOutput) Vulnerability() VulnerabilityNoteResponseOutput {
	return o.ApplyT(func(v LookupNoteResult) VulnerabilityNoteResponse { return v.Vulnerability }).(VulnerabilityNoteResponseOutput)
}

// A note describing a vulnerability assessment.
func (o LookupNoteResultOutput) VulnerabilityAssessment() VulnerabilityAssessmentNoteResponseOutput {
	return o.ApplyT(func(v LookupNoteResult) VulnerabilityAssessmentNoteResponse { return v.VulnerabilityAssessment }).(VulnerabilityAssessmentNoteResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNoteResultOutput{})
}
