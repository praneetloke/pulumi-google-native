# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = [
    'GetNoteResult',
    'AwaitableGetNoteResult',
    'get_note',
    'get_note_output',
]

@pulumi.output_type
class GetNoteResult:
    def __init__(__self__, attestation=None, build=None, compliance=None, create_time=None, deployment=None, discovery=None, dsse_attestation=None, expiration_time=None, image=None, kind=None, long_description=None, name=None, package=None, related_note_names=None, related_url=None, sbom_reference=None, short_description=None, update_time=None, upgrade=None, vulnerability=None, vulnerability_assessment=None):
        if attestation and not isinstance(attestation, dict):
            raise TypeError("Expected argument 'attestation' to be a dict")
        pulumi.set(__self__, "attestation", attestation)
        if build and not isinstance(build, dict):
            raise TypeError("Expected argument 'build' to be a dict")
        pulumi.set(__self__, "build", build)
        if compliance and not isinstance(compliance, dict):
            raise TypeError("Expected argument 'compliance' to be a dict")
        pulumi.set(__self__, "compliance", compliance)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if deployment and not isinstance(deployment, dict):
            raise TypeError("Expected argument 'deployment' to be a dict")
        pulumi.set(__self__, "deployment", deployment)
        if discovery and not isinstance(discovery, dict):
            raise TypeError("Expected argument 'discovery' to be a dict")
        pulumi.set(__self__, "discovery", discovery)
        if dsse_attestation and not isinstance(dsse_attestation, dict):
            raise TypeError("Expected argument 'dsse_attestation' to be a dict")
        pulumi.set(__self__, "dsse_attestation", dsse_attestation)
        if expiration_time and not isinstance(expiration_time, str):
            raise TypeError("Expected argument 'expiration_time' to be a str")
        pulumi.set(__self__, "expiration_time", expiration_time)
        if image and not isinstance(image, dict):
            raise TypeError("Expected argument 'image' to be a dict")
        pulumi.set(__self__, "image", image)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if long_description and not isinstance(long_description, str):
            raise TypeError("Expected argument 'long_description' to be a str")
        pulumi.set(__self__, "long_description", long_description)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if package and not isinstance(package, dict):
            raise TypeError("Expected argument 'package' to be a dict")
        pulumi.set(__self__, "package", package)
        if related_note_names and not isinstance(related_note_names, list):
            raise TypeError("Expected argument 'related_note_names' to be a list")
        pulumi.set(__self__, "related_note_names", related_note_names)
        if related_url and not isinstance(related_url, list):
            raise TypeError("Expected argument 'related_url' to be a list")
        pulumi.set(__self__, "related_url", related_url)
        if sbom_reference and not isinstance(sbom_reference, dict):
            raise TypeError("Expected argument 'sbom_reference' to be a dict")
        pulumi.set(__self__, "sbom_reference", sbom_reference)
        if short_description and not isinstance(short_description, str):
            raise TypeError("Expected argument 'short_description' to be a str")
        pulumi.set(__self__, "short_description", short_description)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)
        if upgrade and not isinstance(upgrade, dict):
            raise TypeError("Expected argument 'upgrade' to be a dict")
        pulumi.set(__self__, "upgrade", upgrade)
        if vulnerability and not isinstance(vulnerability, dict):
            raise TypeError("Expected argument 'vulnerability' to be a dict")
        pulumi.set(__self__, "vulnerability", vulnerability)
        if vulnerability_assessment and not isinstance(vulnerability_assessment, dict):
            raise TypeError("Expected argument 'vulnerability_assessment' to be a dict")
        pulumi.set(__self__, "vulnerability_assessment", vulnerability_assessment)

    @property
    @pulumi.getter
    def attestation(self) -> 'outputs.AttestationNoteResponse':
        """
        A note describing an attestation role.
        """
        return pulumi.get(self, "attestation")

    @property
    @pulumi.getter
    def build(self) -> 'outputs.BuildNoteResponse':
        """
        A note describing build provenance for a verifiable build.
        """
        return pulumi.get(self, "build")

    @property
    @pulumi.getter
    def compliance(self) -> 'outputs.ComplianceNoteResponse':
        """
        A note describing a compliance check.
        """
        return pulumi.get(self, "compliance")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The time this note was created. This field can be used as a filter in list requests.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def deployment(self) -> 'outputs.DeploymentNoteResponse':
        """
        A note describing something that can be deployed.
        """
        return pulumi.get(self, "deployment")

    @property
    @pulumi.getter
    def discovery(self) -> 'outputs.DiscoveryNoteResponse':
        """
        A note describing the initial analysis of a resource.
        """
        return pulumi.get(self, "discovery")

    @property
    @pulumi.getter(name="dsseAttestation")
    def dsse_attestation(self) -> 'outputs.DSSEAttestationNoteResponse':
        """
        A note describing a dsse attestation note.
        """
        return pulumi.get(self, "dsse_attestation")

    @property
    @pulumi.getter(name="expirationTime")
    def expiration_time(self) -> str:
        """
        Time of expiration for this note. Empty if note does not expire.
        """
        return pulumi.get(self, "expiration_time")

    @property
    @pulumi.getter
    def image(self) -> 'outputs.ImageNoteResponse':
        """
        A note describing a base image.
        """
        return pulumi.get(self, "image")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        The type of analysis. This field can be used as a filter in list requests.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="longDescription")
    def long_description(self) -> str:
        """
        A detailed description of this note.
        """
        return pulumi.get(self, "long_description")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The name of the note in the form of `projects/[PROVIDER_ID]/notes/[NOTE_ID]`.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def package(self) -> 'outputs.PackageNoteResponse':
        """
        A note describing a package hosted by various package managers.
        """
        return pulumi.get(self, "package")

    @property
    @pulumi.getter(name="relatedNoteNames")
    def related_note_names(self) -> Sequence[str]:
        """
        Other notes related to this note.
        """
        return pulumi.get(self, "related_note_names")

    @property
    @pulumi.getter(name="relatedUrl")
    def related_url(self) -> Sequence['outputs.RelatedUrlResponse']:
        """
        URLs associated with this note.
        """
        return pulumi.get(self, "related_url")

    @property
    @pulumi.getter(name="sbomReference")
    def sbom_reference(self) -> 'outputs.SBOMReferenceNoteResponse':
        """
        A note describing an SBOM reference.
        """
        return pulumi.get(self, "sbom_reference")

    @property
    @pulumi.getter(name="shortDescription")
    def short_description(self) -> str:
        """
        A one sentence description of this note.
        """
        return pulumi.get(self, "short_description")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The time this note was last updated. This field can be used as a filter in list requests.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter
    def upgrade(self) -> 'outputs.UpgradeNoteResponse':
        """
        A note describing available package upgrades.
        """
        return pulumi.get(self, "upgrade")

    @property
    @pulumi.getter
    def vulnerability(self) -> 'outputs.VulnerabilityNoteResponse':
        """
        A note describing a package vulnerability.
        """
        return pulumi.get(self, "vulnerability")

    @property
    @pulumi.getter(name="vulnerabilityAssessment")
    def vulnerability_assessment(self) -> 'outputs.VulnerabilityAssessmentNoteResponse':
        """
        A note describing a vulnerability assessment.
        """
        return pulumi.get(self, "vulnerability_assessment")


class AwaitableGetNoteResult(GetNoteResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetNoteResult(
            attestation=self.attestation,
            build=self.build,
            compliance=self.compliance,
            create_time=self.create_time,
            deployment=self.deployment,
            discovery=self.discovery,
            dsse_attestation=self.dsse_attestation,
            expiration_time=self.expiration_time,
            image=self.image,
            kind=self.kind,
            long_description=self.long_description,
            name=self.name,
            package=self.package,
            related_note_names=self.related_note_names,
            related_url=self.related_url,
            sbom_reference=self.sbom_reference,
            short_description=self.short_description,
            update_time=self.update_time,
            upgrade=self.upgrade,
            vulnerability=self.vulnerability,
            vulnerability_assessment=self.vulnerability_assessment)


def get_note(note_id: Optional[str] = None,
             project: Optional[str] = None,
             opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetNoteResult:
    """
    Gets the specified note.
    """
    __args__ = dict()
    __args__['noteId'] = note_id
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:containeranalysis/v1:getNote', __args__, opts=opts, typ=GetNoteResult).value

    return AwaitableGetNoteResult(
        attestation=pulumi.get(__ret__, 'attestation'),
        build=pulumi.get(__ret__, 'build'),
        compliance=pulumi.get(__ret__, 'compliance'),
        create_time=pulumi.get(__ret__, 'create_time'),
        deployment=pulumi.get(__ret__, 'deployment'),
        discovery=pulumi.get(__ret__, 'discovery'),
        dsse_attestation=pulumi.get(__ret__, 'dsse_attestation'),
        expiration_time=pulumi.get(__ret__, 'expiration_time'),
        image=pulumi.get(__ret__, 'image'),
        kind=pulumi.get(__ret__, 'kind'),
        long_description=pulumi.get(__ret__, 'long_description'),
        name=pulumi.get(__ret__, 'name'),
        package=pulumi.get(__ret__, 'package'),
        related_note_names=pulumi.get(__ret__, 'related_note_names'),
        related_url=pulumi.get(__ret__, 'related_url'),
        sbom_reference=pulumi.get(__ret__, 'sbom_reference'),
        short_description=pulumi.get(__ret__, 'short_description'),
        update_time=pulumi.get(__ret__, 'update_time'),
        upgrade=pulumi.get(__ret__, 'upgrade'),
        vulnerability=pulumi.get(__ret__, 'vulnerability'),
        vulnerability_assessment=pulumi.get(__ret__, 'vulnerability_assessment'))


@_utilities.lift_output_func(get_note)
def get_note_output(note_id: Optional[pulumi.Input[str]] = None,
                    project: Optional[pulumi.Input[Optional[str]]] = None,
                    opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetNoteResult]:
    """
    Gets the specified note.
    """
    ...
