// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Adds a new Membership.
type Membership struct {
	pulumi.CustomResourceState

	// Optional. How to identify workloads from this Membership. See the documentation on Workload Identity for more details: https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
	Authority AuthorityResponseOutput `pulumi:"authority"`
	// When the Membership was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// When the Membership was deleted.
	DeleteTime pulumi.StringOutput `pulumi:"deleteTime"`
	// Optional. Description of this membership, limited to 63 characters. Must match the regex: `a-zA-Z0-9*`
	Description pulumi.StringOutput `pulumi:"description"`
	// Optional. Endpoint information to reach this member.
	Endpoint MembershipEndpointResponseOutput `pulumi:"endpoint"`
	// Optional. An externally-generated and managed ID for this Membership. This ID may be modified after creation, but this is not recommended. For GKE clusters, external_id is managed by the Hub API and updates will be ignored. The ID must match the regex: `a-zA-Z0-9*` If this Membership represents a Kubernetes cluster, this value should be set to the UID of the `kube-system` namespace object.
	ExternalId pulumi.StringOutput `pulumi:"externalId"`
	// Optional. The infrastructure type this Membership is running on.
	InfrastructureType pulumi.StringOutput `pulumi:"infrastructureType"`
	// Optional. GCP labels for this membership.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// For clusters using Connect, the timestamp of the most recent connection established with Google Cloud. This time is updated every several minutes, not continuously. For clusters that do not use GKE Connect, or that have never connected successfully, this field will be unset.
	LastConnectionTime pulumi.StringOutput `pulumi:"lastConnectionTime"`
	// The full, unique name of this Membership resource in the format `projects/*/locations/*/memberships/{membership_id}`, set during creation. `membership_id` must be a valid RFC 1123 compliant DNS label: 1. At most 63 characters in length 2. It must consist of lower case alphanumeric characters or `-` 3. It must start and end with an alphanumeric character Which can be expressed as the regex: `[a-z0-9]([-a-z0-9]*[a-z0-9])?`, with a maximum length of 63 characters.
	Name pulumi.StringOutput `pulumi:"name"`
	// State of the Membership resource.
	State MembershipStateResponseOutput `pulumi:"state"`
	// Google-generated UUID for this resource. This is unique across all Membership resources. If a Membership resource is deleted and another resource with the same name is created, it gets a different unique_id.
	UniqueId pulumi.StringOutput `pulumi:"uniqueId"`
	// When the Membership was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewMembership registers a new resource with the given unique name, arguments, and options.
func NewMembership(ctx *pulumi.Context,
	name string, args *MembershipArgs, opts ...pulumi.ResourceOption) (*Membership, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Location == nil {
		return nil, errors.New("invalid value for required argument 'Location'")
	}
	if args.MembershipId == nil {
		return nil, errors.New("invalid value for required argument 'MembershipId'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	var resource Membership
	err := ctx.RegisterResource("google-native:gkehub/v1beta1:Membership", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMembership gets an existing Membership resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMembership(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MembershipState, opts ...pulumi.ResourceOption) (*Membership, error) {
	var resource Membership
	err := ctx.ReadResource("google-native:gkehub/v1beta1:Membership", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Membership resources.
type membershipState struct {
	// Optional. How to identify workloads from this Membership. See the documentation on Workload Identity for more details: https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
	Authority *AuthorityResponse `pulumi:"authority"`
	// When the Membership was created.
	CreateTime *string `pulumi:"createTime"`
	// When the Membership was deleted.
	DeleteTime *string `pulumi:"deleteTime"`
	// Optional. Description of this membership, limited to 63 characters. Must match the regex: `a-zA-Z0-9*`
	Description *string `pulumi:"description"`
	// Optional. Endpoint information to reach this member.
	Endpoint *MembershipEndpointResponse `pulumi:"endpoint"`
	// Optional. An externally-generated and managed ID for this Membership. This ID may be modified after creation, but this is not recommended. For GKE clusters, external_id is managed by the Hub API and updates will be ignored. The ID must match the regex: `a-zA-Z0-9*` If this Membership represents a Kubernetes cluster, this value should be set to the UID of the `kube-system` namespace object.
	ExternalId *string `pulumi:"externalId"`
	// Optional. The infrastructure type this Membership is running on.
	InfrastructureType *string `pulumi:"infrastructureType"`
	// Optional. GCP labels for this membership.
	Labels map[string]string `pulumi:"labels"`
	// For clusters using Connect, the timestamp of the most recent connection established with Google Cloud. This time is updated every several minutes, not continuously. For clusters that do not use GKE Connect, or that have never connected successfully, this field will be unset.
	LastConnectionTime *string `pulumi:"lastConnectionTime"`
	// The full, unique name of this Membership resource in the format `projects/*/locations/*/memberships/{membership_id}`, set during creation. `membership_id` must be a valid RFC 1123 compliant DNS label: 1. At most 63 characters in length 2. It must consist of lower case alphanumeric characters or `-` 3. It must start and end with an alphanumeric character Which can be expressed as the regex: `[a-z0-9]([-a-z0-9]*[a-z0-9])?`, with a maximum length of 63 characters.
	Name *string `pulumi:"name"`
	// State of the Membership resource.
	State *MembershipStateResponse `pulumi:"state"`
	// Google-generated UUID for this resource. This is unique across all Membership resources. If a Membership resource is deleted and another resource with the same name is created, it gets a different unique_id.
	UniqueId *string `pulumi:"uniqueId"`
	// When the Membership was last updated.
	UpdateTime *string `pulumi:"updateTime"`
}

type MembershipState struct {
	// Optional. How to identify workloads from this Membership. See the documentation on Workload Identity for more details: https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
	Authority AuthorityResponsePtrInput
	// When the Membership was created.
	CreateTime pulumi.StringPtrInput
	// When the Membership was deleted.
	DeleteTime pulumi.StringPtrInput
	// Optional. Description of this membership, limited to 63 characters. Must match the regex: `a-zA-Z0-9*`
	Description pulumi.StringPtrInput
	// Optional. Endpoint information to reach this member.
	Endpoint MembershipEndpointResponsePtrInput
	// Optional. An externally-generated and managed ID for this Membership. This ID may be modified after creation, but this is not recommended. For GKE clusters, external_id is managed by the Hub API and updates will be ignored. The ID must match the regex: `a-zA-Z0-9*` If this Membership represents a Kubernetes cluster, this value should be set to the UID of the `kube-system` namespace object.
	ExternalId pulumi.StringPtrInput
	// Optional. The infrastructure type this Membership is running on.
	InfrastructureType pulumi.StringPtrInput
	// Optional. GCP labels for this membership.
	Labels pulumi.StringMapInput
	// For clusters using Connect, the timestamp of the most recent connection established with Google Cloud. This time is updated every several minutes, not continuously. For clusters that do not use GKE Connect, or that have never connected successfully, this field will be unset.
	LastConnectionTime pulumi.StringPtrInput
	// The full, unique name of this Membership resource in the format `projects/*/locations/*/memberships/{membership_id}`, set during creation. `membership_id` must be a valid RFC 1123 compliant DNS label: 1. At most 63 characters in length 2. It must consist of lower case alphanumeric characters or `-` 3. It must start and end with an alphanumeric character Which can be expressed as the regex: `[a-z0-9]([-a-z0-9]*[a-z0-9])?`, with a maximum length of 63 characters.
	Name pulumi.StringPtrInput
	// State of the Membership resource.
	State MembershipStateResponsePtrInput
	// Google-generated UUID for this resource. This is unique across all Membership resources. If a Membership resource is deleted and another resource with the same name is created, it gets a different unique_id.
	UniqueId pulumi.StringPtrInput
	// When the Membership was last updated.
	UpdateTime pulumi.StringPtrInput
}

func (MembershipState) ElementType() reflect.Type {
	return reflect.TypeOf((*membershipState)(nil)).Elem()
}

type membershipArgs struct {
	// Optional. How to identify workloads from this Membership. See the documentation on Workload Identity for more details: https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
	Authority *Authority `pulumi:"authority"`
	// Optional. Description of this membership, limited to 63 characters. Must match the regex: `a-zA-Z0-9*`
	Description *string `pulumi:"description"`
	// Optional. Endpoint information to reach this member.
	Endpoint *MembershipEndpoint `pulumi:"endpoint"`
	// Optional. An externally-generated and managed ID for this Membership. This ID may be modified after creation, but this is not recommended. For GKE clusters, external_id is managed by the Hub API and updates will be ignored. The ID must match the regex: `a-zA-Z0-9*` If this Membership represents a Kubernetes cluster, this value should be set to the UID of the `kube-system` namespace object.
	ExternalId *string `pulumi:"externalId"`
	// Optional. The infrastructure type this Membership is running on.
	InfrastructureType *string `pulumi:"infrastructureType"`
	// Optional. GCP labels for this membership.
	Labels       map[string]string `pulumi:"labels"`
	Location     string            `pulumi:"location"`
	MembershipId string            `pulumi:"membershipId"`
	Project      string            `pulumi:"project"`
	RequestId    *string           `pulumi:"requestId"`
}

// The set of arguments for constructing a Membership resource.
type MembershipArgs struct {
	// Optional. How to identify workloads from this Membership. See the documentation on Workload Identity for more details: https://cloud.google.com/kubernetes-engine/docs/how-to/workload-identity
	Authority AuthorityPtrInput
	// Optional. Description of this membership, limited to 63 characters. Must match the regex: `a-zA-Z0-9*`
	Description pulumi.StringPtrInput
	// Optional. Endpoint information to reach this member.
	Endpoint MembershipEndpointPtrInput
	// Optional. An externally-generated and managed ID for this Membership. This ID may be modified after creation, but this is not recommended. For GKE clusters, external_id is managed by the Hub API and updates will be ignored. The ID must match the regex: `a-zA-Z0-9*` If this Membership represents a Kubernetes cluster, this value should be set to the UID of the `kube-system` namespace object.
	ExternalId pulumi.StringPtrInput
	// Optional. The infrastructure type this Membership is running on.
	InfrastructureType pulumi.StringPtrInput
	// Optional. GCP labels for this membership.
	Labels       pulumi.StringMapInput
	Location     pulumi.StringInput
	MembershipId pulumi.StringInput
	Project      pulumi.StringInput
	RequestId    pulumi.StringPtrInput
}

func (MembershipArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*membershipArgs)(nil)).Elem()
}

type MembershipInput interface {
	pulumi.Input

	ToMembershipOutput() MembershipOutput
	ToMembershipOutputWithContext(ctx context.Context) MembershipOutput
}

func (*Membership) ElementType() reflect.Type {
	return reflect.TypeOf((*Membership)(nil))
}

func (i *Membership) ToMembershipOutput() MembershipOutput {
	return i.ToMembershipOutputWithContext(context.Background())
}

func (i *Membership) ToMembershipOutputWithContext(ctx context.Context) MembershipOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MembershipOutput)
}

type MembershipOutput struct {
	*pulumi.OutputState
}

func (MembershipOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Membership)(nil))
}

func (o MembershipOutput) ToMembershipOutput() MembershipOutput {
	return o
}

func (o MembershipOutput) ToMembershipOutputWithContext(ctx context.Context) MembershipOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(MembershipOutput{})
}
