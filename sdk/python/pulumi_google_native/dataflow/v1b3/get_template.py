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
    'GetTemplateResult',
    'AwaitableGetTemplateResult',
    'get_template',
    'get_template_output',
]

@pulumi.output_type
class GetTemplateResult:
    def __init__(__self__, metadata=None, runtime_metadata=None, status=None, template_type=None):
        if metadata and not isinstance(metadata, dict):
            raise TypeError("Expected argument 'metadata' to be a dict")
        pulumi.set(__self__, "metadata", metadata)
        if runtime_metadata and not isinstance(runtime_metadata, dict):
            raise TypeError("Expected argument 'runtime_metadata' to be a dict")
        pulumi.set(__self__, "runtime_metadata", runtime_metadata)
        if status and not isinstance(status, dict):
            raise TypeError("Expected argument 'status' to be a dict")
        pulumi.set(__self__, "status", status)
        if template_type and not isinstance(template_type, str):
            raise TypeError("Expected argument 'template_type' to be a str")
        pulumi.set(__self__, "template_type", template_type)

    @property
    @pulumi.getter
    def metadata(self) -> 'outputs.TemplateMetadataResponse':
        """
        The template metadata describing the template name, available parameters, etc.
        """
        return pulumi.get(self, "metadata")

    @property
    @pulumi.getter(name="runtimeMetadata")
    def runtime_metadata(self) -> 'outputs.RuntimeMetadataResponse':
        """
        Describes the runtime metadata with SDKInfo and available parameters.
        """
        return pulumi.get(self, "runtime_metadata")

    @property
    @pulumi.getter
    def status(self) -> 'outputs.StatusResponse':
        """
        The status of the get template request. Any problems with the request will be indicated in the error_details.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter(name="templateType")
    def template_type(self) -> str:
        """
        Template Type.
        """
        return pulumi.get(self, "template_type")


class AwaitableGetTemplateResult(GetTemplateResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetTemplateResult(
            metadata=self.metadata,
            runtime_metadata=self.runtime_metadata,
            status=self.status,
            template_type=self.template_type)


def get_template(gcs_path: Optional[str] = None,
                 location: Optional[str] = None,
                 project: Optional[str] = None,
                 view: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetTemplateResult:
    """
    Get the template associated with a template.
    """
    __args__ = dict()
    __args__['gcsPath'] = gcs_path
    __args__['location'] = location
    __args__['project'] = project
    __args__['view'] = view
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:dataflow/v1b3:getTemplate', __args__, opts=opts, typ=GetTemplateResult).value

    return AwaitableGetTemplateResult(
        metadata=pulumi.get(__ret__, 'metadata'),
        runtime_metadata=pulumi.get(__ret__, 'runtime_metadata'),
        status=pulumi.get(__ret__, 'status'),
        template_type=pulumi.get(__ret__, 'template_type'))


@_utilities.lift_output_func(get_template)
def get_template_output(gcs_path: Optional[pulumi.Input[str]] = None,
                        location: Optional[pulumi.Input[str]] = None,
                        project: Optional[pulumi.Input[Optional[str]]] = None,
                        view: Optional[pulumi.Input[Optional[str]]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetTemplateResult]:
    """
    Get the template associated with a template.
    """
    ...
