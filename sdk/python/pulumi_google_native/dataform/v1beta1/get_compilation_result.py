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
    'GetCompilationResultResult',
    'AwaitableGetCompilationResultResult',
    'get_compilation_result',
    'get_compilation_result_output',
]

@pulumi.output_type
class GetCompilationResultResult:
    def __init__(__self__, code_compilation_config=None, compilation_errors=None, dataform_core_version=None, git_commitish=None, name=None, release_config=None, resolved_git_commit_sha=None, workspace=None):
        if code_compilation_config and not isinstance(code_compilation_config, dict):
            raise TypeError("Expected argument 'code_compilation_config' to be a dict")
        pulumi.set(__self__, "code_compilation_config", code_compilation_config)
        if compilation_errors and not isinstance(compilation_errors, list):
            raise TypeError("Expected argument 'compilation_errors' to be a list")
        pulumi.set(__self__, "compilation_errors", compilation_errors)
        if dataform_core_version and not isinstance(dataform_core_version, str):
            raise TypeError("Expected argument 'dataform_core_version' to be a str")
        pulumi.set(__self__, "dataform_core_version", dataform_core_version)
        if git_commitish and not isinstance(git_commitish, str):
            raise TypeError("Expected argument 'git_commitish' to be a str")
        pulumi.set(__self__, "git_commitish", git_commitish)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if release_config and not isinstance(release_config, str):
            raise TypeError("Expected argument 'release_config' to be a str")
        pulumi.set(__self__, "release_config", release_config)
        if resolved_git_commit_sha and not isinstance(resolved_git_commit_sha, str):
            raise TypeError("Expected argument 'resolved_git_commit_sha' to be a str")
        pulumi.set(__self__, "resolved_git_commit_sha", resolved_git_commit_sha)
        if workspace and not isinstance(workspace, str):
            raise TypeError("Expected argument 'workspace' to be a str")
        pulumi.set(__self__, "workspace", workspace)

    @property
    @pulumi.getter(name="codeCompilationConfig")
    def code_compilation_config(self) -> 'outputs.CodeCompilationConfigResponse':
        """
        Immutable. If set, fields of `code_compilation_config` override the default compilation settings that are specified in dataform.json.
        """
        return pulumi.get(self, "code_compilation_config")

    @property
    @pulumi.getter(name="compilationErrors")
    def compilation_errors(self) -> Sequence['outputs.CompilationErrorResponse']:
        """
        Errors encountered during project compilation.
        """
        return pulumi.get(self, "compilation_errors")

    @property
    @pulumi.getter(name="dataformCoreVersion")
    def dataform_core_version(self) -> str:
        """
        The version of `@dataform/core` that was used for compilation.
        """
        return pulumi.get(self, "dataform_core_version")

    @property
    @pulumi.getter(name="gitCommitish")
    def git_commitish(self) -> str:
        """
        Immutable. Git commit/tag/branch name at which the repository should be compiled. Must exist in the remote repository. Examples: - a commit SHA: `12ade345` - a tag: `tag1` - a branch name: `branch1`
        """
        return pulumi.get(self, "git_commitish")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The compilation result's name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="releaseConfig")
    def release_config(self) -> str:
        """
        Immutable. The name of the release config to compile. The release config's 'current_compilation_result' field will be updated to this compilation result. Must be in the format `projects/*/locations/*/repositories/*/releaseConfigs/*`.
        """
        return pulumi.get(self, "release_config")

    @property
    @pulumi.getter(name="resolvedGitCommitSha")
    def resolved_git_commit_sha(self) -> str:
        """
        The fully resolved Git commit SHA of the code that was compiled. Not set for compilation results whose source is a workspace.
        """
        return pulumi.get(self, "resolved_git_commit_sha")

    @property
    @pulumi.getter
    def workspace(self) -> str:
        """
        Immutable. The name of the workspace to compile. Must be in the format `projects/*/locations/*/repositories/*/workspaces/*`.
        """
        return pulumi.get(self, "workspace")


class AwaitableGetCompilationResultResult(GetCompilationResultResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetCompilationResultResult(
            code_compilation_config=self.code_compilation_config,
            compilation_errors=self.compilation_errors,
            dataform_core_version=self.dataform_core_version,
            git_commitish=self.git_commitish,
            name=self.name,
            release_config=self.release_config,
            resolved_git_commit_sha=self.resolved_git_commit_sha,
            workspace=self.workspace)


def get_compilation_result(compilation_result_id: Optional[str] = None,
                           location: Optional[str] = None,
                           project: Optional[str] = None,
                           repository_id: Optional[str] = None,
                           opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetCompilationResultResult:
    """
    Fetches a single CompilationResult.
    """
    __args__ = dict()
    __args__['compilationResultId'] = compilation_result_id
    __args__['location'] = location
    __args__['project'] = project
    __args__['repositoryId'] = repository_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:dataform/v1beta1:getCompilationResult', __args__, opts=opts, typ=GetCompilationResultResult).value

    return AwaitableGetCompilationResultResult(
        code_compilation_config=pulumi.get(__ret__, 'code_compilation_config'),
        compilation_errors=pulumi.get(__ret__, 'compilation_errors'),
        dataform_core_version=pulumi.get(__ret__, 'dataform_core_version'),
        git_commitish=pulumi.get(__ret__, 'git_commitish'),
        name=pulumi.get(__ret__, 'name'),
        release_config=pulumi.get(__ret__, 'release_config'),
        resolved_git_commit_sha=pulumi.get(__ret__, 'resolved_git_commit_sha'),
        workspace=pulumi.get(__ret__, 'workspace'))


@_utilities.lift_output_func(get_compilation_result)
def get_compilation_result_output(compilation_result_id: Optional[pulumi.Input[str]] = None,
                                  location: Optional[pulumi.Input[str]] = None,
                                  project: Optional[pulumi.Input[Optional[str]]] = None,
                                  repository_id: Optional[pulumi.Input[str]] = None,
                                  opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetCompilationResultResult]:
    """
    Fetches a single CompilationResult.
    """
    ...
