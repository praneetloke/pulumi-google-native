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
    'GetResponsePolicyResult',
    'AwaitableGetResponsePolicyResult',
    'get_response_policy',
    'get_response_policy_output',
]

@pulumi.output_type
class GetResponsePolicyResult:
    def __init__(__self__, description=None, gke_clusters=None, kind=None, labels=None, networks=None, response_policy_name=None):
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if gke_clusters and not isinstance(gke_clusters, list):
            raise TypeError("Expected argument 'gke_clusters' to be a list")
        pulumi.set(__self__, "gke_clusters", gke_clusters)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if networks and not isinstance(networks, list):
            raise TypeError("Expected argument 'networks' to be a list")
        pulumi.set(__self__, "networks", networks)
        if response_policy_name and not isinstance(response_policy_name, str):
            raise TypeError("Expected argument 'response_policy_name' to be a str")
        pulumi.set(__self__, "response_policy_name", response_policy_name)

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        User-provided description for this Response Policy.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="gkeClusters")
    def gke_clusters(self) -> Sequence['outputs.ResponsePolicyGKEClusterResponse']:
        """
        The list of Google Kubernetes Engine clusters to which this response policy is applied.
        """
        return pulumi.get(self, "gke_clusters")

    @property
    @pulumi.getter
    def kind(self) -> str:
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        User labels.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def networks(self) -> Sequence['outputs.ResponsePolicyNetworkResponse']:
        """
        List of network names specifying networks to which this policy is applied.
        """
        return pulumi.get(self, "networks")

    @property
    @pulumi.getter(name="responsePolicyName")
    def response_policy_name(self) -> str:
        """
        User assigned name for this Response Policy.
        """
        return pulumi.get(self, "response_policy_name")


class AwaitableGetResponsePolicyResult(GetResponsePolicyResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetResponsePolicyResult(
            description=self.description,
            gke_clusters=self.gke_clusters,
            kind=self.kind,
            labels=self.labels,
            networks=self.networks,
            response_policy_name=self.response_policy_name)


def get_response_policy(client_operation_id: Optional[str] = None,
                        project: Optional[str] = None,
                        response_policy: Optional[str] = None,
                        opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetResponsePolicyResult:
    """
    Fetches the representation of an existing Response Policy.
    """
    __args__ = dict()
    __args__['clientOperationId'] = client_operation_id
    __args__['project'] = project
    __args__['responsePolicy'] = response_policy
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:dns/v1beta2:getResponsePolicy', __args__, opts=opts, typ=GetResponsePolicyResult).value

    return AwaitableGetResponsePolicyResult(
        description=pulumi.get(__ret__, 'description'),
        gke_clusters=pulumi.get(__ret__, 'gke_clusters'),
        kind=pulumi.get(__ret__, 'kind'),
        labels=pulumi.get(__ret__, 'labels'),
        networks=pulumi.get(__ret__, 'networks'),
        response_policy_name=pulumi.get(__ret__, 'response_policy_name'))


@_utilities.lift_output_func(get_response_policy)
def get_response_policy_output(client_operation_id: Optional[pulumi.Input[Optional[str]]] = None,
                               project: Optional[pulumi.Input[Optional[str]]] = None,
                               response_policy: Optional[pulumi.Input[str]] = None,
                               opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetResponsePolicyResult]:
    """
    Fetches the representation of an existing Response Policy.
    """
    ...
