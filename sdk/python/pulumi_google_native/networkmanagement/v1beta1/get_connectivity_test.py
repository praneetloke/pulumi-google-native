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
    'GetConnectivityTestResult',
    'AwaitableGetConnectivityTestResult',
    'get_connectivity_test',
    'get_connectivity_test_output',
]

@pulumi.output_type
class GetConnectivityTestResult:
    def __init__(__self__, create_time=None, description=None, destination=None, display_name=None, labels=None, name=None, probing_details=None, protocol=None, reachability_details=None, related_projects=None, source=None, update_time=None):
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if destination and not isinstance(destination, dict):
            raise TypeError("Expected argument 'destination' to be a dict")
        pulumi.set(__self__, "destination", destination)
        if display_name and not isinstance(display_name, str):
            raise TypeError("Expected argument 'display_name' to be a str")
        pulumi.set(__self__, "display_name", display_name)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if probing_details and not isinstance(probing_details, dict):
            raise TypeError("Expected argument 'probing_details' to be a dict")
        pulumi.set(__self__, "probing_details", probing_details)
        if protocol and not isinstance(protocol, str):
            raise TypeError("Expected argument 'protocol' to be a str")
        pulumi.set(__self__, "protocol", protocol)
        if reachability_details and not isinstance(reachability_details, dict):
            raise TypeError("Expected argument 'reachability_details' to be a dict")
        pulumi.set(__self__, "reachability_details", reachability_details)
        if related_projects and not isinstance(related_projects, list):
            raise TypeError("Expected argument 'related_projects' to be a list")
        pulumi.set(__self__, "related_projects", related_projects)
        if source and not isinstance(source, dict):
            raise TypeError("Expected argument 'source' to be a dict")
        pulumi.set(__self__, "source", source)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The time the test was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        The user-supplied description of the Connectivity Test. Maximum of 512 characters.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def destination(self) -> 'outputs.EndpointResponse':
        """
        Destination specification of the Connectivity Test. You can use a combination of destination IP address, Compute Engine VM instance, or VPC network to uniquely identify the destination location. Even if the destination IP address is not unique, the source IP location is unique. Usually, the analysis can infer the destination endpoint from route information. If the destination you specify is a VM instance and the instance has multiple network interfaces, then you must also specify either a destination IP address or VPC network to identify the destination interface. A reachability analysis proceeds even if the destination location is ambiguous. However, the result can include endpoints that you don't intend to test.
        """
        return pulumi.get(self, "destination")

    @property
    @pulumi.getter(name="displayName")
    def display_name(self) -> str:
        """
        The display name of a Connectivity Test.
        """
        return pulumi.get(self, "display_name")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Resource labels to represent user-provided metadata.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Unique name of the resource using the form: `projects/{project_id}/locations/global/connectivityTests/{test}`
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="probingDetails")
    def probing_details(self) -> 'outputs.ProbingDetailsResponse':
        """
        The probing details of this test from the latest run, present for applicable tests only. The details are updated when creating a new test, updating an existing test, or triggering a one-time rerun of an existing test.
        """
        return pulumi.get(self, "probing_details")

    @property
    @pulumi.getter
    def protocol(self) -> str:
        """
        IP Protocol of the test. When not provided, "TCP" is assumed.
        """
        return pulumi.get(self, "protocol")

    @property
    @pulumi.getter(name="reachabilityDetails")
    def reachability_details(self) -> 'outputs.ReachabilityDetailsResponse':
        """
        The reachability details of this test from the latest run. The details are updated when creating a new test, updating an existing test, or triggering a one-time rerun of an existing test.
        """
        return pulumi.get(self, "reachability_details")

    @property
    @pulumi.getter(name="relatedProjects")
    def related_projects(self) -> Sequence[str]:
        """
        Other projects that may be relevant for reachability analysis. This is applicable to scenarios where a test can cross project boundaries.
        """
        return pulumi.get(self, "related_projects")

    @property
    @pulumi.getter
    def source(self) -> 'outputs.EndpointResponse':
        """
        Source specification of the Connectivity Test. You can use a combination of source IP address, virtual machine (VM) instance, or Compute Engine network to uniquely identify the source location. Examples: If the source IP address is an internal IP address within a Google Cloud Virtual Private Cloud (VPC) network, then you must also specify the VPC network. Otherwise, specify the VM instance, which already contains its internal IP address and VPC network information. If the source of the test is within an on-premises network, then you must provide the destination VPC network. If the source endpoint is a Compute Engine VM instance with multiple network interfaces, the instance itself is not sufficient to identify the endpoint. So, you must also specify the source IP address or VPC network. A reachability analysis proceeds even if the source location is ambiguous. However, the test result may include endpoints that you don't intend to test.
        """
        return pulumi.get(self, "source")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The time the test's configuration was updated.
        """
        return pulumi.get(self, "update_time")


class AwaitableGetConnectivityTestResult(GetConnectivityTestResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetConnectivityTestResult(
            create_time=self.create_time,
            description=self.description,
            destination=self.destination,
            display_name=self.display_name,
            labels=self.labels,
            name=self.name,
            probing_details=self.probing_details,
            protocol=self.protocol,
            reachability_details=self.reachability_details,
            related_projects=self.related_projects,
            source=self.source,
            update_time=self.update_time)


def get_connectivity_test(connectivity_test_id: Optional[str] = None,
                          project: Optional[str] = None,
                          opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetConnectivityTestResult:
    """
    Gets the details of a specific Connectivity Test.
    """
    __args__ = dict()
    __args__['connectivityTestId'] = connectivity_test_id
    __args__['project'] = project
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:networkmanagement/v1beta1:getConnectivityTest', __args__, opts=opts, typ=GetConnectivityTestResult).value

    return AwaitableGetConnectivityTestResult(
        create_time=pulumi.get(__ret__, 'create_time'),
        description=pulumi.get(__ret__, 'description'),
        destination=pulumi.get(__ret__, 'destination'),
        display_name=pulumi.get(__ret__, 'display_name'),
        labels=pulumi.get(__ret__, 'labels'),
        name=pulumi.get(__ret__, 'name'),
        probing_details=pulumi.get(__ret__, 'probing_details'),
        protocol=pulumi.get(__ret__, 'protocol'),
        reachability_details=pulumi.get(__ret__, 'reachability_details'),
        related_projects=pulumi.get(__ret__, 'related_projects'),
        source=pulumi.get(__ret__, 'source'),
        update_time=pulumi.get(__ret__, 'update_time'))


@_utilities.lift_output_func(get_connectivity_test)
def get_connectivity_test_output(connectivity_test_id: Optional[pulumi.Input[str]] = None,
                                 project: Optional[pulumi.Input[Optional[str]]] = None,
                                 opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetConnectivityTestResult]:
    """
    Gets the details of a specific Connectivity Test.
    """
    ...
