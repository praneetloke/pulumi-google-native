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
    'GetWorkflowResult',
    'AwaitableGetWorkflowResult',
    'get_workflow',
    'get_workflow_output',
]

@pulumi.output_type
class GetWorkflowResult:
    def __init__(__self__, call_log_level=None, create_time=None, crypto_key_name=None, description=None, labels=None, name=None, revision_create_time=None, revision_id=None, service_account=None, source_contents=None, state=None, state_error=None, update_time=None):
        if call_log_level and not isinstance(call_log_level, str):
            raise TypeError("Expected argument 'call_log_level' to be a str")
        pulumi.set(__self__, "call_log_level", call_log_level)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if crypto_key_name and not isinstance(crypto_key_name, str):
            raise TypeError("Expected argument 'crypto_key_name' to be a str")
        pulumi.set(__self__, "crypto_key_name", crypto_key_name)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if revision_create_time and not isinstance(revision_create_time, str):
            raise TypeError("Expected argument 'revision_create_time' to be a str")
        pulumi.set(__self__, "revision_create_time", revision_create_time)
        if revision_id and not isinstance(revision_id, str):
            raise TypeError("Expected argument 'revision_id' to be a str")
        pulumi.set(__self__, "revision_id", revision_id)
        if service_account and not isinstance(service_account, str):
            raise TypeError("Expected argument 'service_account' to be a str")
        pulumi.set(__self__, "service_account", service_account)
        if source_contents and not isinstance(source_contents, str):
            raise TypeError("Expected argument 'source_contents' to be a str")
        pulumi.set(__self__, "source_contents", source_contents)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if state_error and not isinstance(state_error, dict):
            raise TypeError("Expected argument 'state_error' to be a dict")
        pulumi.set(__self__, "state_error", state_error)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="callLogLevel")
    def call_log_level(self) -> str:
        """
        Optional. Describes the level of platform logging to apply to calls and call responses during executions of this workflow. If both the workflow and the execution specify a logging level, the execution level takes precedence.
        """
        return pulumi.get(self, "call_log_level")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The timestamp for when the workflow was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="cryptoKeyName")
    def crypto_key_name(self) -> str:
        """
        Optional. The resource name of a KMS crypto key used to encrypt or decrypt the data associated with the workflow. Format: projects/{project}/locations/{location}/keyRings/{keyRing}/cryptoKeys/{cryptoKey} Using `-` as a wildcard for the `{project}` or not providing one at all will infer the project from the account. If not provided, data associated with the workflow will not be CMEK-encrypted.
        """
        return pulumi.get(self, "crypto_key_name")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        Description of the workflow provided by the user. Must be at most 1000 unicode characters long.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Labels associated with this workflow. Labels can contain at most 64 entries. Keys and values can be no longer than 63 characters and can only contain lowercase letters, numeric characters, underscores, and dashes. Label keys must start with a letter. International characters are allowed.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        The resource name of the workflow. Format: projects/{project}/locations/{location}/workflows/{workflow}
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="revisionCreateTime")
    def revision_create_time(self) -> str:
        """
        The timestamp for the latest revision of the workflow's creation.
        """
        return pulumi.get(self, "revision_create_time")

    @property
    @pulumi.getter(name="revisionId")
    def revision_id(self) -> str:
        """
        The revision of the workflow. A new revision of a workflow is created as a result of updating the following properties of a workflow: - Service account - Workflow code to be executed The format is "000001-a4d", where the first 6 characters define the zero-padded revision ordinal number. They are followed by a hyphen and 3 hexadecimal random characters.
        """
        return pulumi.get(self, "revision_id")

    @property
    @pulumi.getter(name="serviceAccount")
    def service_account(self) -> str:
        """
        The service account associated with the latest workflow version. This service account represents the identity of the workflow and determines what permissions the workflow has. Format: projects/{project}/serviceAccounts/{account} or {account} Using `-` as a wildcard for the `{project}` or not providing one at all will infer the project from the account. The `{account}` value can be the `email` address or the `unique_id` of the service account. If not provided, workflow will use the project's default service account. Modifying this field for an existing workflow results in a new workflow revision.
        """
        return pulumi.get(self, "service_account")

    @property
    @pulumi.getter(name="sourceContents")
    def source_contents(self) -> str:
        """
        Workflow code to be executed. The size limit is 128KB.
        """
        return pulumi.get(self, "source_contents")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        State of the workflow deployment.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter(name="stateError")
    def state_error(self) -> 'outputs.StateErrorResponse':
        """
        Error regarding the state of the workflow. For example, this field will have error details if the execution data is unavailable due to revoked KMS key permissions.
        """
        return pulumi.get(self, "state_error")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The timestamp for when the workflow was last updated.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetWorkflowResult(GetWorkflowResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetWorkflowResult(
            call_log_level=self.call_log_level,
            create_time=self.create_time,
            crypto_key_name=self.crypto_key_name,
            description=self.description,
            labels=self.labels,
            name=self.name,
            revision_create_time=self.revision_create_time,
            revision_id=self.revision_id,
            service_account=self.service_account,
            source_contents=self.source_contents,
            state=self.state,
            state_error=self.state_error,
            update_time=self.update_time)


def get_workflow(location: Optional[str] = None,
                 project: Optional[str] = None,
                 revision_id: Optional[str] = None,
                 workflow_id: Optional[str] = None,
                 opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetWorkflowResult:
    """
    Gets details of a single workflow.
    """
    __args__ = dict()
    __args__['location'] = location
    __args__['project'] = project
    __args__['revisionId'] = revision_id
    __args__['workflowId'] = workflow_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:workflows/v1:getWorkflow', __args__, opts=opts, typ=GetWorkflowResult).value

    return AwaitableGetWorkflowResult(
        call_log_level=__ret__.call_log_level,
        create_time=__ret__.create_time,
        crypto_key_name=__ret__.crypto_key_name,
        description=__ret__.description,
        labels=__ret__.labels,
        name=__ret__.name,
        revision_create_time=__ret__.revision_create_time,
        revision_id=__ret__.revision_id,
        service_account=__ret__.service_account,
        source_contents=__ret__.source_contents,
        state=__ret__.state,
        state_error=__ret__.state_error,
        update_time=__ret__.update_time)


@_utilities.lift_output_func(get_workflow)
def get_workflow_output(location: Optional[pulumi.Input[str]] = None,
                        project: Optional[pulumi.Input[Optional[str]]] = None,
                        revision_id: Optional[pulumi.Input[Optional[str]]] = None,
                        workflow_id: Optional[pulumi.Input[str]] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetWorkflowResult]:
    """
    Gets details of a single workflow.
    """
    ...
