# coding=utf-8
# *** WARNING: this file was generated by the Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from ... import _utilities
from . import outputs

__all__ = [
    'GetInterconnectAttachmentResult',
    'AwaitableGetInterconnectAttachmentResult',
    'get_interconnect_attachment',
]

@pulumi.output_type
class GetInterconnectAttachmentResult:
    def __init__(__self__, admin_enabled=None, bandwidth=None, candidate_subnets=None, cloud_router_ip_address=None, creation_timestamp=None, customer_router_ip_address=None, dataplane_version=None, description=None, edge_availability_domain=None, encryption=None, interconnect=None, ipsec_internal_addresses=None, kind=None, label_fingerprint=None, labels=None, mtu=None, name=None, operational_status=None, pairing_key=None, partner_asn=None, partner_metadata=None, private_interconnect_info=None, region=None, router=None, self_link=None, state=None, type=None, vlan_tag8021q=None):
        if admin_enabled and not isinstance(admin_enabled, bool):
            raise TypeError("Expected argument 'admin_enabled' to be a bool")
        pulumi.set(__self__, "admin_enabled", admin_enabled)
        if bandwidth and not isinstance(bandwidth, str):
            raise TypeError("Expected argument 'bandwidth' to be a str")
        pulumi.set(__self__, "bandwidth", bandwidth)
        if candidate_subnets and not isinstance(candidate_subnets, list):
            raise TypeError("Expected argument 'candidate_subnets' to be a list")
        pulumi.set(__self__, "candidate_subnets", candidate_subnets)
        if cloud_router_ip_address and not isinstance(cloud_router_ip_address, str):
            raise TypeError("Expected argument 'cloud_router_ip_address' to be a str")
        pulumi.set(__self__, "cloud_router_ip_address", cloud_router_ip_address)
        if creation_timestamp and not isinstance(creation_timestamp, str):
            raise TypeError("Expected argument 'creation_timestamp' to be a str")
        pulumi.set(__self__, "creation_timestamp", creation_timestamp)
        if customer_router_ip_address and not isinstance(customer_router_ip_address, str):
            raise TypeError("Expected argument 'customer_router_ip_address' to be a str")
        pulumi.set(__self__, "customer_router_ip_address", customer_router_ip_address)
        if dataplane_version and not isinstance(dataplane_version, int):
            raise TypeError("Expected argument 'dataplane_version' to be a int")
        pulumi.set(__self__, "dataplane_version", dataplane_version)
        if description and not isinstance(description, str):
            raise TypeError("Expected argument 'description' to be a str")
        pulumi.set(__self__, "description", description)
        if edge_availability_domain and not isinstance(edge_availability_domain, str):
            raise TypeError("Expected argument 'edge_availability_domain' to be a str")
        pulumi.set(__self__, "edge_availability_domain", edge_availability_domain)
        if encryption and not isinstance(encryption, str):
            raise TypeError("Expected argument 'encryption' to be a str")
        pulumi.set(__self__, "encryption", encryption)
        if interconnect and not isinstance(interconnect, str):
            raise TypeError("Expected argument 'interconnect' to be a str")
        pulumi.set(__self__, "interconnect", interconnect)
        if ipsec_internal_addresses and not isinstance(ipsec_internal_addresses, list):
            raise TypeError("Expected argument 'ipsec_internal_addresses' to be a list")
        pulumi.set(__self__, "ipsec_internal_addresses", ipsec_internal_addresses)
        if kind and not isinstance(kind, str):
            raise TypeError("Expected argument 'kind' to be a str")
        pulumi.set(__self__, "kind", kind)
        if label_fingerprint and not isinstance(label_fingerprint, str):
            raise TypeError("Expected argument 'label_fingerprint' to be a str")
        pulumi.set(__self__, "label_fingerprint", label_fingerprint)
        if labels and not isinstance(labels, dict):
            raise TypeError("Expected argument 'labels' to be a dict")
        pulumi.set(__self__, "labels", labels)
        if mtu and not isinstance(mtu, int):
            raise TypeError("Expected argument 'mtu' to be a int")
        pulumi.set(__self__, "mtu", mtu)
        if name and not isinstance(name, str):
            raise TypeError("Expected argument 'name' to be a str")
        pulumi.set(__self__, "name", name)
        if operational_status and not isinstance(operational_status, str):
            raise TypeError("Expected argument 'operational_status' to be a str")
        pulumi.set(__self__, "operational_status", operational_status)
        if pairing_key and not isinstance(pairing_key, str):
            raise TypeError("Expected argument 'pairing_key' to be a str")
        pulumi.set(__self__, "pairing_key", pairing_key)
        if partner_asn and not isinstance(partner_asn, str):
            raise TypeError("Expected argument 'partner_asn' to be a str")
        pulumi.set(__self__, "partner_asn", partner_asn)
        if partner_metadata and not isinstance(partner_metadata, dict):
            raise TypeError("Expected argument 'partner_metadata' to be a dict")
        pulumi.set(__self__, "partner_metadata", partner_metadata)
        if private_interconnect_info and not isinstance(private_interconnect_info, dict):
            raise TypeError("Expected argument 'private_interconnect_info' to be a dict")
        pulumi.set(__self__, "private_interconnect_info", private_interconnect_info)
        if region and not isinstance(region, str):
            raise TypeError("Expected argument 'region' to be a str")
        pulumi.set(__self__, "region", region)
        if router and not isinstance(router, str):
            raise TypeError("Expected argument 'router' to be a str")
        pulumi.set(__self__, "router", router)
        if self_link and not isinstance(self_link, str):
            raise TypeError("Expected argument 'self_link' to be a str")
        pulumi.set(__self__, "self_link", self_link)
        if state and not isinstance(state, str):
            raise TypeError("Expected argument 'state' to be a str")
        pulumi.set(__self__, "state", state)
        if type and not isinstance(type, str):
            raise TypeError("Expected argument 'type' to be a str")
        pulumi.set(__self__, "type", type)
        if vlan_tag8021q and not isinstance(vlan_tag8021q, int):
            raise TypeError("Expected argument 'vlan_tag8021q' to be a int")
        pulumi.set(__self__, "vlan_tag8021q", vlan_tag8021q)

    @property
    @pulumi.getter(name="adminEnabled")
    def admin_enabled(self) -> bool:
        """
        Determines whether this Attachment will carry packets. Not present for PARTNER_PROVIDER.
        """
        return pulumi.get(self, "admin_enabled")

    @property
    @pulumi.getter
    def bandwidth(self) -> str:
        """
        Provisioned bandwidth capacity for the interconnect attachment. For attachments of type DEDICATED, the user can set the bandwidth. For attachments of type PARTNER, the Google Partner that is operating the interconnect must set the bandwidth. Output only for PARTNER type, mutable for PARTNER_PROVIDER and DEDICATED, and can take one of the following values: - BPS_50M: 50 Mbit/s - BPS_100M: 100 Mbit/s - BPS_200M: 200 Mbit/s - BPS_300M: 300 Mbit/s - BPS_400M: 400 Mbit/s - BPS_500M: 500 Mbit/s - BPS_1G: 1 Gbit/s - BPS_2G: 2 Gbit/s - BPS_5G: 5 Gbit/s - BPS_10G: 10 Gbit/s - BPS_20G: 20 Gbit/s - BPS_50G: 50 Gbit/s 
        """
        return pulumi.get(self, "bandwidth")

    @property
    @pulumi.getter(name="candidateSubnets")
    def candidate_subnets(self) -> Sequence[str]:
        """
        Up to 16 candidate prefixes that can be used to restrict the allocation of cloudRouterIpAddress and customerRouterIpAddress for this attachment. All prefixes must be within link-local address space (169.254.0.0/16) and must be /29 or shorter (/28, /27, etc). Google will attempt to select an unused /29 from the supplied candidate prefix(es). The request will fail if all possible /29s are in use on Google's edge. If not supplied, Google will randomly select an unused /29 from all of link-local space.
        """
        return pulumi.get(self, "candidate_subnets")

    @property
    @pulumi.getter(name="cloudRouterIpAddress")
    def cloud_router_ip_address(self) -> str:
        """
        IPv4 address + prefix length to be configured on Cloud Router Interface for this interconnect attachment.
        """
        return pulumi.get(self, "cloud_router_ip_address")

    @property
    @pulumi.getter(name="creationTimestamp")
    def creation_timestamp(self) -> str:
        """
        Creation timestamp in RFC3339 text format.
        """
        return pulumi.get(self, "creation_timestamp")

    @property
    @pulumi.getter(name="customerRouterIpAddress")
    def customer_router_ip_address(self) -> str:
        """
        IPv4 address + prefix length to be configured on the customer router subinterface for this interconnect attachment.
        """
        return pulumi.get(self, "customer_router_ip_address")

    @property
    @pulumi.getter(name="dataplaneVersion")
    def dataplane_version(self) -> int:
        """
        Dataplane version for this InterconnectAttachment.
        """
        return pulumi.get(self, "dataplane_version")

    @property
    @pulumi.getter
    def description(self) -> str:
        """
        An optional description of this resource.
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="edgeAvailabilityDomain")
    def edge_availability_domain(self) -> str:
        """
        Desired availability domain for the attachment. Only available for type PARTNER, at creation time, and can take one of the following values: - AVAILABILITY_DOMAIN_ANY - AVAILABILITY_DOMAIN_1 - AVAILABILITY_DOMAIN_2 For improved reliability, customers should configure a pair of attachments, one per availability domain. The selected availability domain will be provided to the Partner via the pairing key, so that the provisioned circuit will lie in the specified domain. If not specified, the value will default to AVAILABILITY_DOMAIN_ANY.
        """
        return pulumi.get(self, "edge_availability_domain")

    @property
    @pulumi.getter
    def encryption(self) -> str:
        """
        Indicates the user-supplied encryption option of this VLAN attachment (interconnectAttachment). Can only be specified at attachment creation for PARTNER or DEDICATED attachments. Possible values are: - NONE - This is the default value, which means that the VLAN attachment carries unencrypted traffic. VMs are able to send traffic to, or receive traffic from, such a VLAN attachment. - IPSEC - The VLAN attachment carries only encrypted traffic that is encrypted by an IPsec device, such as an HA VPN gateway or third-party IPsec VPN. VMs cannot directly send traffic to, or receive traffic from, such a VLAN attachment. To use *IPsec-encrypted Cloud Interconnect*, the VLAN attachment must be created with this option. Not currently available publicly. 
        """
        return pulumi.get(self, "encryption")

    @property
    @pulumi.getter
    def interconnect(self) -> str:
        """
        URL of the underlying Interconnect object that this attachment's traffic will traverse through.
        """
        return pulumi.get(self, "interconnect")

    @property
    @pulumi.getter(name="ipsecInternalAddresses")
    def ipsec_internal_addresses(self) -> Sequence[str]:
        """
        A list of URLs of addresses that have been reserved for the VLAN attachment. Used only for the VLAN attachment that has the encryption option as IPSEC. The addresses must be regional internal IP address ranges. When creating an HA VPN gateway over the VLAN attachment, if the attachment is configured to use a regional internal IP address, then the VPN gateway's IP address is allocated from the IP address range specified here. For example, if the HA VPN gateway's interface 0 is paired to this VLAN attachment, then a regional internal IP address for the VPN gateway interface 0 will be allocated from the IP address specified for this VLAN attachment. If this field is not specified when creating the VLAN attachment, then later on when creating an HA VPN gateway on this VLAN attachment, the HA VPN gateway's IP address is allocated from the regional external IP address pool. Not currently available publicly. 
        """
        return pulumi.get(self, "ipsec_internal_addresses")

    @property
    @pulumi.getter
    def kind(self) -> str:
        """
        Type of the resource. Always compute#interconnectAttachment for interconnect attachments.
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter(name="labelFingerprint")
    def label_fingerprint(self) -> str:
        """
        A fingerprint for the labels being applied to this InterconnectAttachment, which is essentially a hash of the labels set used for optimistic locking. The fingerprint is initially generated by Compute Engine and changes after every request to modify or update labels. You must always provide an up-to-date fingerprint hash in order to update or change labels, otherwise the request will fail with error 412 conditionNotMet. To see the latest fingerprint, make a get() request to retrieve an InterconnectAttachment.
        """
        return pulumi.get(self, "label_fingerprint")

    @property
    @pulumi.getter
    def labels(self) -> Mapping[str, str]:
        """
        Labels for this resource. These can only be added or modified by the setLabels method. Each label key/value pair must comply with RFC1035. Label values may be empty.
        """
        return pulumi.get(self, "labels")

    @property
    @pulumi.getter
    def mtu(self) -> int:
        """
        Maximum Transmission Unit (MTU), in bytes, of packets passing through this interconnect attachment. Only 1440 and 1500 are allowed. If not specified, the value will default to 1440.
        """
        return pulumi.get(self, "mtu")

    @property
    @pulumi.getter
    def name(self) -> str:
        """
        Name of the resource. Provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="operationalStatus")
    def operational_status(self) -> str:
        """
        The current status of whether or not this interconnect attachment is functional, which can take one of the following values: - OS_ACTIVE: The attachment has been turned up and is ready to use. - OS_UNPROVISIONED: The attachment is not ready to use yet, because turnup is not complete. 
        """
        return pulumi.get(self, "operational_status")

    @property
    @pulumi.getter(name="pairingKey")
    def pairing_key(self) -> str:
        """
        [Output only for type PARTNER. Input only for PARTNER_PROVIDER. Not present for DEDICATED]. The opaque identifier of an PARTNER attachment used to initiate provisioning with a selected partner. Of the form "XXXXX/region/domain"
        """
        return pulumi.get(self, "pairing_key")

    @property
    @pulumi.getter(name="partnerAsn")
    def partner_asn(self) -> str:
        """
        Optional BGP ASN for the router supplied by a Layer 3 Partner if they configured BGP on behalf of the customer. Output only for PARTNER type, input only for PARTNER_PROVIDER, not available for DEDICATED.
        """
        return pulumi.get(self, "partner_asn")

    @property
    @pulumi.getter(name="partnerMetadata")
    def partner_metadata(self) -> 'outputs.InterconnectAttachmentPartnerMetadataResponse':
        """
        Informational metadata about Partner attachments from Partners to display to customers. Output only for for PARTNER type, mutable for PARTNER_PROVIDER, not available for DEDICATED.
        """
        return pulumi.get(self, "partner_metadata")

    @property
    @pulumi.getter(name="privateInterconnectInfo")
    def private_interconnect_info(self) -> 'outputs.InterconnectAttachmentPrivateInfoResponse':
        """
        Information specific to an InterconnectAttachment. This property is populated if the interconnect that this is attached to is of type DEDICATED.
        """
        return pulumi.get(self, "private_interconnect_info")

    @property
    @pulumi.getter
    def region(self) -> str:
        """
        URL of the region where the regional interconnect attachment resides. You must specify this field as part of the HTTP request URL. It is not settable as a field in the request body.
        """
        return pulumi.get(self, "region")

    @property
    @pulumi.getter
    def router(self) -> str:
        """
        URL of the Cloud Router to be used for dynamic routing. This router must be in the same region as this InterconnectAttachment. The InterconnectAttachment will automatically connect the Interconnect to the network & region within which the Cloud Router is configured.
        """
        return pulumi.get(self, "router")

    @property
    @pulumi.getter(name="selfLink")
    def self_link(self) -> str:
        """
        Server-defined URL for the resource.
        """
        return pulumi.get(self, "self_link")

    @property
    @pulumi.getter
    def state(self) -> str:
        """
        The current state of this attachment's functionality. Enum values ACTIVE and UNPROVISIONED are shared by DEDICATED/PRIVATE, PARTNER, and PARTNER_PROVIDER interconnect attachments, while enum values PENDING_PARTNER, PARTNER_REQUEST_RECEIVED, and PENDING_CUSTOMER are used for only PARTNER and PARTNER_PROVIDER interconnect attachments. This state can take one of the following values: - ACTIVE: The attachment has been turned up and is ready to use. - UNPROVISIONED: The attachment is not ready to use yet, because turnup is not complete. - PENDING_PARTNER: A newly-created PARTNER attachment that has not yet been configured on the Partner side. - PARTNER_REQUEST_RECEIVED: A PARTNER attachment is in the process of provisioning after a PARTNER_PROVIDER attachment was created that references it. - PENDING_CUSTOMER: A PARTNER or PARTNER_PROVIDER attachment that is waiting for a customer to activate it. - DEFUNCT: The attachment was deleted externally and is no longer functional. This could be because the associated Interconnect was removed, or because the other side of a Partner attachment was deleted. 
        """
        return pulumi.get(self, "state")

    @property
    @pulumi.getter
    def type(self) -> str:
        """
        The type of interconnect attachment this is, which can take one of the following values: - DEDICATED: an attachment to a Dedicated Interconnect. - PARTNER: an attachment to a Partner Interconnect, created by the customer. - PARTNER_PROVIDER: an attachment to a Partner Interconnect, created by the partner. 
        """
        return pulumi.get(self, "type")

    @property
    @pulumi.getter(name="vlanTag8021q")
    def vlan_tag8021q(self) -> int:
        """
        The IEEE 802.1Q VLAN tag for this attachment, in the range 2-4094. Only specified at creation time.
        """
        return pulumi.get(self, "vlan_tag8021q")


class AwaitableGetInterconnectAttachmentResult(GetInterconnectAttachmentResult):
    # pylint: disable=using-constant-test
    def __await__(self):
        if False:
            yield self
        return GetInterconnectAttachmentResult(
            admin_enabled=self.admin_enabled,
            bandwidth=self.bandwidth,
            candidate_subnets=self.candidate_subnets,
            cloud_router_ip_address=self.cloud_router_ip_address,
            creation_timestamp=self.creation_timestamp,
            customer_router_ip_address=self.customer_router_ip_address,
            dataplane_version=self.dataplane_version,
            description=self.description,
            edge_availability_domain=self.edge_availability_domain,
            encryption=self.encryption,
            interconnect=self.interconnect,
            ipsec_internal_addresses=self.ipsec_internal_addresses,
            kind=self.kind,
            label_fingerprint=self.label_fingerprint,
            labels=self.labels,
            mtu=self.mtu,
            name=self.name,
            operational_status=self.operational_status,
            pairing_key=self.pairing_key,
            partner_asn=self.partner_asn,
            partner_metadata=self.partner_metadata,
            private_interconnect_info=self.private_interconnect_info,
            region=self.region,
            router=self.router,
            self_link=self.self_link,
            state=self.state,
            type=self.type,
            vlan_tag8021q=self.vlan_tag8021q)


def get_interconnect_attachment(interconnect_attachment: Optional[str] = None,
                                project: Optional[str] = None,
                                region: Optional[str] = None,
                                opts: Optional[pulumi.InvokeOptions] = None) -> AwaitableGetInterconnectAttachmentResult:
    """
    Returns the specified interconnect attachment.
    """
    __args__ = dict()
    __args__['interconnectAttachment'] = interconnect_attachment
    __args__['project'] = project
    __args__['region'] = region
    if opts is None:
        opts = pulumi.InvokeOptions()
    if opts.version is None:
        opts.version = _utilities.get_version()
    __ret__ = pulumi.runtime.invoke('google-native:compute/beta:getInterconnectAttachment', __args__, opts=opts, typ=GetInterconnectAttachmentResult).value

    return AwaitableGetInterconnectAttachmentResult(
        admin_enabled=__ret__.admin_enabled,
        bandwidth=__ret__.bandwidth,
        candidate_subnets=__ret__.candidate_subnets,
        cloud_router_ip_address=__ret__.cloud_router_ip_address,
        creation_timestamp=__ret__.creation_timestamp,
        customer_router_ip_address=__ret__.customer_router_ip_address,
        dataplane_version=__ret__.dataplane_version,
        description=__ret__.description,
        edge_availability_domain=__ret__.edge_availability_domain,
        encryption=__ret__.encryption,
        interconnect=__ret__.interconnect,
        ipsec_internal_addresses=__ret__.ipsec_internal_addresses,
        kind=__ret__.kind,
        label_fingerprint=__ret__.label_fingerprint,
        labels=__ret__.labels,
        mtu=__ret__.mtu,
        name=__ret__.name,
        operational_status=__ret__.operational_status,
        pairing_key=__ret__.pairing_key,
        partner_asn=__ret__.partner_asn,
        partner_metadata=__ret__.partner_metadata,
        private_interconnect_info=__ret__.private_interconnect_info,
        region=__ret__.region,
        router=__ret__.router,
        self_link=__ret__.self_link,
        state=__ret__.state,
        type=__ret__.type,
        vlan_tag8021q=__ret__.vlan_tag8021q)