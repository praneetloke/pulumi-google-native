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
    'GetVmwareClusterResult',
    'AwaitableGetVmwareClusterResult',
    'get_vmware_cluster',
    'get_vmware_cluster_output',
]

@pulumi.output_type
class GetVmwareClusterResult:
    def __init__(__self__, admin_cluster_membership=None, admin_cluster_name=None, annotations=None, anti_affinity_groups=None, authorization=None, auto_repair_config=None, control_plane_node=None, create_time=None, dataplane_v2=None, delete_time=None, description=None, enable_control_plane_v2=None, endpoint=None, etag=None, fleet=None, load_balancer=None, local_name=None, name=None, network_config=None, on_prem_version=None, reconciling=None, state=None, status=None, storage=None, uid=None, update_time=None, validation_check=None, vcenter=None, vm_tracking_enabled=None):
        if admin_cluster_membership and not isinstance(admin_cluster_membership, str):
            raise TypeError("Expected argument 'admin_cluster_membership' to be a str")
        pulumi.set(__self__, "admin_cluster_membership", admin_cluster_membership)
        if admin_cluster_name and not isinstance(admin_cluster_name, str):
            raise TypeError("Expected argument 'admin_cluster_name' to be a str")
        pulumi.set(__self__, "admin_cluster_name", admin_cluster_name)
        if annotations and not isinstance(annotations, dict):
            raise TypeError("Expected argument 'annotations' to be a dict")
        pulumi.set(__self__, "annotations", annotations)
        if anti_affinity_groups and not isinstance(anti_affinity_groups, dict):
            raise TypeError("Expected argument 'anti_affinity_groups' to be a dict")
        pulumi.set(__self__, "anti_affinity_groups", anti_affinity_groups)
        if authorization and not isinstance(authorization, dict):
            raise TypeError("Expected argument 'authorization' to be a dict")
        pulumi.set(__self__, "authorization", authorization)
        if auto_repair_config and not isinstance(auto_repair_config, dict):
            raise TypeError("Expected argument 'auto_repair_config' to be a dict")
        pulumi.set(__self__, "auto_repair_config", auto_repair_config)
        if control_plane_node and not isinstance(control_plane_node, dict):
            raise TypeError("Expected argument 'control_plane_node' to be a dict")
        pulumi.set(__self__, "control_plane_node", control_plane_node)
        if create_time and not isinstance(create_time, str):
            raise TypeError("Expected argument 'create_time' to be a str")
        pulumi.set(__self__, "create_time", create_time)
        if dataplane_v2 and not isinstance(dataplane_v2, dict):
            raise TypeError("Expected argument 'dataplane_v2' to be a dict")
        pulumi.set(__self__, "dataplane_v2", dataplane_v2)
        if delete_time and not isinstance(delete_time, str):
            raise TypeError("Expected argument 'delete_time' to be a str")
        pulumi.set(__self__, "delete_time", delete_time)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if enable_control_plane_v2 and not isinstance(enable_control_plane_v2, bool):
            raise TypeError("Expected argument 'enable_control_plane_v2' to be a bool")
        pulumi.set(__self__, "enable_control_plane_v2", enable_control_plane_v2)
        if endpoint and not isinstance(endpoint, str):
            raise TypeError("Expected argument 'endpoint' to be a str")
        pulumi.set(__self__, "endpoint", endpoint)
        if etag and not isinstance(etag, str):
            raise TypeError("Expected argument 'etag' to be a str")
        pulumi.set(__self__, "etag", etag)
        if fleet and not isinstance(fleet, dict):
            raise TypeError("Expected argument 'fleet' to be a dict")
        pulumi.set(__self__, "fleet", fleet)
        if load_balancer and not isinstance(load_balancer, dict):
            raise TypeError("Expected argument 'load_balancer' to be a dict")
        pulumi.set(__self__, "load_balancer", load_balancer)
        if local_name and not isinstance(local_name, str):
            raise TypeError("Expected argument 'local_name' to be a str")
        pulumi.set(__self__, "local_name", local_name)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if network_config and not isinstance(network_config, dict):
            raise TypeError("Expected argument 'network_config' to be a dict")
        pulumi.set(__self__, "network_config", network_config)
        if on_prem_version and not isinstance(on_prem_version, str):
            raise TypeError("Expected argument 'on_prem_version' to be a str")
        pulumi.set(__self__, "on_prem_version", on_prem_version)
        if reconciling and not isinstance(reconciling, bool):
            raise TypeError("Expected argument 'reconciling' to be a bool")
        pulumi.set(__self__, "reconciling", reconciling)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if status and not isinstance(status, dict):
            raise TypeError("Expected argument 'status' to be a dict")
        pulumi.set(__self__, "status", status)
        if storage and not isinstance(storage, dict):
            raise TypeError("Expected argument 'storage' to be a dict")
        pulumi.set(__self__, "storage", storage)
        if uid and not isinstance(uid, str):
            raise TypeError("Expected argument 'uid' to be a str")
        pulumi.set(__self__, "uid", uid)
        if update_time and not isinstance(update_time, str):
            raise TypeError("Expected argument 'update_time' to be a str")
        pulumi.set(__self__, "update_time", update_time)
        if validation_check and not isinstance(validation_check, dict):
            raise TypeError("Expected argument 'validation_check' to be a dict")
        pulumi.set(__self__, "validation_check", validation_check)
        if vcenter and not isinstance(vcenter, dict):
            raise TypeError("Expected argument 'vcenter' to be a dict")
        pulumi.set(__self__, "vcenter", vcenter)
        if vm_tracking_enabled and not isinstance(vm_tracking_enabled, bool):
            raise TypeError("Expected argument 'vm_tracking_enabled' to be a bool")
        pulumi.set(__self__, "vm_tracking_enabled", vm_tracking_enabled)

    @property
    @pulumi.getter(name="adminClusterMembership")
    def admin_cluster_membership(self) -> str:
        """
        The admin cluster this VMware user cluster belongs to. This is the full resource name of the admin cluster's fleet membership. In the future, references to other resource types might be allowed if admin clusters are modeled as their own resources.
        """
        return pulumi.get(self, "admin_cluster_membership")

    @property
    @pulumi.getter(name="adminClusterName")
    def admin_cluster_name(self) -> str:
        """
        The resource name of the VMware admin cluster hosting this user cluster.
        """
        return pulumi.get(self, "admin_cluster_name")

    @property
    @pulumi.getter
    def annotations(self) -> Mapping[str, str]:
        """
        Annotations on the VMware user cluster. This field has the same restrictions as Kubernetes annotations. The total size of all keys and values combined is limited to 256k. Key can have 2 segments: prefix (optional) and name (required), separated by a slash (/). Prefix must be a DNS subdomain. Name must be 63 characters or less, begin and end with alphanumerics, with dashes (-), underscores (_), dots (.), and alphanumerics between.
        """
        return pulumi.get(self, "annotations")

    @property
    @pulumi.getter(name="antiAffinityGroups")
    def anti_affinity_groups(self) -> 'outputs.VmwareAAGConfigResponse':
        """
        AAGConfig specifies whether to spread VMware user cluster nodes across at least three physical hosts in the datacenter.
        """
        return pulumi.get(self, "anti_affinity_groups")

    @property
    @pulumi.getter
    def authorization(self) -> 'outputs.AuthorizationResponse':
        """
        RBAC policy that will be applied and managed by the Anthos On-Prem API.
        """
        return pulumi.get(self, "authorization")

    @property
    @pulumi.getter(name="autoRepairConfig")
    def auto_repair_config(self) -> 'outputs.VmwareAutoRepairConfigResponse':
        """
        Configuration for auto repairing.
        """
        return pulumi.get(self, "auto_repair_config")

    @property
    @pulumi.getter(name="controlPlaneNode")
    def control_plane_node(self) -> 'outputs.VmwareControlPlaneNodeConfigResponse':
        """
        VMware user cluster control plane nodes must have either 1 or 3 replicas.
        """
        return pulumi.get(self, "control_plane_node")

    @property
    @pulumi.getter(name="createTime")
    def create_time(self) -> str:
        """
        The time at which VMware user cluster was created.
        """
        return pulumi.get(self, "create_time")

    @property
    @pulumi.getter(name="dataplaneV2")
    def dataplane_v2(self) -> 'outputs.VmwareDataplaneV2ConfigResponse':
        """
        VmwareDataplaneV2Config specifies configuration for Dataplane V2.
        """
        return pulumi.get(self, "dataplane_v2")

    @property
    @pulumi.getter(name="deleteTime")
    def delete_time(self) -> str:
        """
        The time at which VMware user cluster was deleted.
        """
        return pulumi.get(self, "delete_time")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        A human readable description of this VMware user cluster.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="enableControlPlaneV2")
    def enable_control_plane_v2(self) -> bool:
        """
        Enable control plane V2. Default to false.
        """
        return pulumi.get(self, "enable_control_plane_v2")

    @property
    @pulumi.getter
    def endpoint(self) -> str:
        """
        The DNS name of VMware user cluster's API server.
        """
        return pulumi.get(self, "endpoint")

    @property
    @pulumi.getter
    def etag(self) -> str:
        """
        This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding. Allows clients to perform consistent read-modify-writes through optimistic concurrency control.
        """
        return pulumi.get(self, "etag")

    @property
    @pulumi.getter
    def fleet(self) -> 'outputs.FleetResponse':
        """
        Fleet configuration for the cluster.
        """
        return pulumi.get(self, "fleet")

    @property
    @pulumi.getter(name="loadBalancer")
    def load_balancer(self) -> 'outputs.VmwareLoadBalancerConfigResponse':
        """
        Load balancer configuration.
        """
        return pulumi.get(self, "load_balancer")

    @property
    @pulumi.getter(name="localName")
    def local_name(self) -> str:
        """
        The object name of the VMware OnPremUserCluster custom resource on the associated admin cluster. This field is used to support conflicting names when enrolling existing clusters to the API. When used as a part of cluster enrollment, this field will differ from the ID in the resource name. For new clusters, this field will match the user provided cluster name and be visible in the last component of the resource name. It is not modifiable. All users should use this name to access their cluster using gkectl or kubectl and should expect to see the local name when viewing admin cluster controller logs.
        """
        return pulumi.get(self, "local_name")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Immutable. The VMware user cluster resource name.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="networkConfig")
    def network_config(self) -> 'outputs.VmwareNetworkConfigResponse':
        """
        The VMware user cluster network configuration.
        """
        return pulumi.get(self, "network_config")

    @property
    @pulumi.getter(name="onPremVersion")
    def on_prem_version(self) -> str:
        """
        The Anthos clusters on the VMware version for your user cluster. Defaults to the admin cluster version.
        """
        return pulumi.get(self, "on_prem_version")

    @property
    @pulumi.getter
    def reconciling(self) -> bool:
        """
        If set, there are currently changes in flight to the VMware user cluster.
        """
        return pulumi.get(self, "reconciling")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The current state of VMware user cluster.
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def status(self) -> 'outputs.ResourceStatusResponse':
        """
        ResourceStatus representing detailed cluster state.
        """
        return pulumi.get(self, "status")

    @property
    @pulumi.getter
    def storage(self) -> 'outputs.VmwareStorageConfigResponse':
        """
        Storage configuration.
        """
        return pulumi.get(self, "storage")

    @property
    @pulumi.getter
    def uid(self) -> str:
        """
        The unique identifier of the VMware user cluster.
        """
        return pulumi.get(self, "uid")

    @property
    @pulumi.getter(name="updateTime")
    def update_time(self) -> str:
        """
        The time at which VMware user cluster was last updated.
        """
        return pulumi.get(self, "update_time")

    @property
    @pulumi.getter(name="validationCheck")
    def validation_check(self) -> 'outputs.ValidationCheckResponse':
        """
        ValidationCheck represents the result of the preflight check job.
        """
        return pulumi.get(self, "validation_check")

    @property
    @pulumi.getter
    def vcenter(self) -> 'outputs.VmwareVCenterConfigResponse':
        """
        VmwareVCenterConfig specifies vCenter config for the user cluster. Inherited from the admin cluster.
        """
        return pulumi.get(self, "vcenter")

    @property
    @pulumi.getter(name="vmTrackingEnabled")
    def vm_tracking_enabled(self) -> bool:
        """
        Enable VM tracking.
        """
        return pulumi.get(self, "vm_tracking_enabled")


class AwaitableGetVmwareClusterResult(GetVmwareClusterResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetVmwareClusterResult(
            admin_cluster_membership=self.admin_cluster_membership,
            admin_cluster_name=self.admin_cluster_name,
            annotations=self.annotations,
            anti_affinity_groups=self.anti_affinity_groups,
            authorization=self.authorization,
            auto_repair_config=self.auto_repair_config,
            control_plane_node=self.control_plane_node,
            create_time=self.create_time,
            dataplane_v2=self.dataplane_v2,
            delete_time=self.delete_time,
            description=self.description,
            enable_control_plane_v2=self.enable_control_plane_v2,
            endpoint=self.endpoint,
            etag=self.etag,
            fleet=self.fleet,
            load_balancer=self.load_balancer,
            local_name=self.local_name,
            name=self.name,
            network_config=self.network_config,
            on_prem_version=self.on_prem_version,
            reconciling=self.reconciling,
            state=self.state,
            status=self.status,
            storage=self.storage,
            uid=self.uid,
            update_time=self.update_time,
            validation_check=self.validation_check,
            vcenter=self.vcenter,
            vm_tracking_enabled=self.vm_tracking_enabled)


def get_vmware_cluster(location: Optional[str] = None,
                       project: Optional[str] = None,
                       view: Optional[str] = None,
                       vmware_cluster_id: Optional[str] = None,
                       opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetVmwareClusterResult:
    """
    Gets details of a single VMware Cluster.
    """
    __args__ = dict()
    __args__['location'] = location
    __args__['project'] = project
    __args__['view'] = view
    __args__['vmwareClusterId'] = vmware_cluster_id
    opts = pulumi.InvokeOptions.merge(_utilities.get_invoke_opts_defaults(), opts)
    __ret__ = pulumi.runtime.invoke('google-native:gkeonprem/v1:getVmwareCluster', __args__, opts=opts, typ=GetVmwareClusterResult).value

    return AwaitableGetVmwareClusterResult(
        admin_cluster_membership=pulumi.get(__ret__, 'admin_cluster_membership'),
        admin_cluster_name=pulumi.get(__ret__, 'admin_cluster_name'),
        annotations=pulumi.get(__ret__, 'annotations'),
        anti_affinity_groups=pulumi.get(__ret__, 'anti_affinity_groups'),
        authorization=pulumi.get(__ret__, 'authorization'),
        auto_repair_config=pulumi.get(__ret__, 'auto_repair_config'),
        control_plane_node=pulumi.get(__ret__, 'control_plane_node'),
        create_time=pulumi.get(__ret__, 'create_time'),
        dataplane_v2=pulumi.get(__ret__, 'dataplane_v2'),
        delete_time=pulumi.get(__ret__, 'delete_time'),
        description=pulumi.get(__ret__, 'description'),
        enable_control_plane_v2=pulumi.get(__ret__, 'enable_control_plane_v2'),
        endpoint=pulumi.get(__ret__, 'endpoint'),
        etag=pulumi.get(__ret__, 'etag'),
        fleet=pulumi.get(__ret__, 'fleet'),
        load_balancer=pulumi.get(__ret__, 'load_balancer'),
        local_name=pulumi.get(__ret__, 'local_name'),
        name=pulumi.get(__ret__, 'name'),
        network_config=pulumi.get(__ret__, 'network_config'),
        on_prem_version=pulumi.get(__ret__, 'on_prem_version'),
        reconciling=pulumi.get(__ret__, 'reconciling'),
        state=pulumi.get(__ret__, 'state'),
        status=pulumi.get(__ret__, 'status'),
        storage=pulumi.get(__ret__, 'storage'),
        uid=pulumi.get(__ret__, 'uid'),
        update_time=pulumi.get(__ret__, 'update_time'),
        validation_check=pulumi.get(__ret__, 'validation_check'),
        vcenter=pulumi.get(__ret__, 'vcenter'),
        vm_tracking_enabled=pulumi.get(__ret__, 'vm_tracking_enabled'))


@_utilities.lift_output_func(get_vmware_cluster)
def get_vmware_cluster_output(location: Optional[pulumi.Input[str]] = None,
                              project: Optional[pulumi.Input[Optional[str]]] = None,
                              view: Optional[pulumi.Input[Optional[str]]] = None,
                              vmware_cluster_id: Optional[pulumi.Input[str]] = None,
                              opts: Optional[pulumi.InvokeOptions] = None) -> pulumi.Output[GetVmwareClusterResult]:
    """
    Gets details of a single VMware Cluster.
    """
    ...
