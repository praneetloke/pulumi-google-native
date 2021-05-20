// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new Connectivity Test. After you create a test, the reachability analysis is performed as part of the long running operation, which completes when the analysis completes. If the endpoint specifications in `ConnectivityTest` are invalid (for example, containing non-existent resources in the network, or you don't have read permissions to the network configurations of listed projects), then the reachability result returns a value of `UNKNOWN`. If the endpoint specifications in `ConnectivityTest` are incomplete, the reachability result returns a value of AMBIGUOUS. For more information, see the Connectivity Test documentation.
type ConnectivityTest struct {
	pulumi.CustomResourceState

	// The time the test was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The user-supplied description of the Connectivity Test. Maximum of 512 characters.
	Description pulumi.StringOutput `pulumi:"description"`
	// Required. Destination specification of the Connectivity Test. You can use a combination of destination IP address, Compute Engine VM instance, or VPC network to uniquely identify the destination location. Even if the destination IP address is not unique, the source IP location is unique. Usually, the analysis can infer the destination endpoint from route information. If the destination you specify is a VM instance and the instance has multiple network interfaces, then you must also specify either a destination IP address or VPC network to identify the destination interface. A reachability analysis proceeds even if the destination location is ambiguous. However, the result can include endpoints that you don't intend to test.
	Destination EndpointResponseOutput `pulumi:"destination"`
	// The display name of a Connectivity Test.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Resource labels to represent user-provided metadata.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Required. Unique name of the resource using the form: `projects/{project_id}/locations/global/connectivityTests/{test}`
	Name pulumi.StringOutput `pulumi:"name"`
	// The probing details of this test from the latest run, present for applicable tests only. The details are updated when creating a new test, updating an existing test, or triggering a one-time rerun of an existing test.
	ProbingDetails ProbingDetailsResponseOutput `pulumi:"probingDetails"`
	// IP Protocol of the test. When not provided, "TCP" is assumed.
	Protocol pulumi.StringOutput `pulumi:"protocol"`
	// The reachability details of this test from the latest run. The details are updated when creating a new test, updating an existing test, or triggering a one-time rerun of an existing test.
	ReachabilityDetails ReachabilityDetailsResponseOutput `pulumi:"reachabilityDetails"`
	// Other projects that may be relevant for reachability analysis. This is applicable to scenarios where a test can cross project boundaries.
	RelatedProjects pulumi.StringArrayOutput `pulumi:"relatedProjects"`
	// Required. Source specification of the Connectivity Test. You can use a combination of source IP address, virtual machine (VM) instance, or Compute Engine network to uniquely identify the source location. Examples: If the source IP address is an internal IP address within a Google Cloud Virtual Private Cloud (VPC) network, then you must also specify the VPC network. Otherwise, specify the VM instance, which already contains its internal IP address and VPC network information. If the source of the test is within an on-premises network, then you must provide the destination VPC network. If the source endpoint is a Compute Engine VM instance with multiple network interfaces, the instance itself is not sufficient to identify the endpoint. So, you must also specify the source IP address or VPC network. A reachability analysis proceeds even if the source location is ambiguous. However, the test result may include endpoints that you don't intend to test.
	Source EndpointResponseOutput `pulumi:"source"`
	// The time the test's configuration was updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewConnectivityTest registers a new resource with the given unique name, arguments, and options.
func NewConnectivityTest(ctx *pulumi.Context,
	name string, args *ConnectivityTestArgs, opts ...pulumi.ResourceOption) (*ConnectivityTest, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectivityTestId == nil {
		return nil, errors.New("invalid value for required argument 'ConnectivityTestId'")
	}
	if args.Project == nil {
		return nil, errors.New("invalid value for required argument 'Project'")
	}
	if args.TestId == nil {
		return nil, errors.New("invalid value for required argument 'TestId'")
	}
	var resource ConnectivityTest
	err := ctx.RegisterResource("google-native:networkmanagement/v1beta1:ConnectivityTest", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnectivityTest gets an existing ConnectivityTest resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnectivityTest(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectivityTestState, opts ...pulumi.ResourceOption) (*ConnectivityTest, error) {
	var resource ConnectivityTest
	err := ctx.ReadResource("google-native:networkmanagement/v1beta1:ConnectivityTest", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConnectivityTest resources.
type connectivityTestState struct {
	// The time the test was created.
	CreateTime *string `pulumi:"createTime"`
	// The user-supplied description of the Connectivity Test. Maximum of 512 characters.
	Description *string `pulumi:"description"`
	// Required. Destination specification of the Connectivity Test. You can use a combination of destination IP address, Compute Engine VM instance, or VPC network to uniquely identify the destination location. Even if the destination IP address is not unique, the source IP location is unique. Usually, the analysis can infer the destination endpoint from route information. If the destination you specify is a VM instance and the instance has multiple network interfaces, then you must also specify either a destination IP address or VPC network to identify the destination interface. A reachability analysis proceeds even if the destination location is ambiguous. However, the result can include endpoints that you don't intend to test.
	Destination *EndpointResponse `pulumi:"destination"`
	// The display name of a Connectivity Test.
	DisplayName *string `pulumi:"displayName"`
	// Resource labels to represent user-provided metadata.
	Labels map[string]string `pulumi:"labels"`
	// Required. Unique name of the resource using the form: `projects/{project_id}/locations/global/connectivityTests/{test}`
	Name *string `pulumi:"name"`
	// The probing details of this test from the latest run, present for applicable tests only. The details are updated when creating a new test, updating an existing test, or triggering a one-time rerun of an existing test.
	ProbingDetails *ProbingDetailsResponse `pulumi:"probingDetails"`
	// IP Protocol of the test. When not provided, "TCP" is assumed.
	Protocol *string `pulumi:"protocol"`
	// The reachability details of this test from the latest run. The details are updated when creating a new test, updating an existing test, or triggering a one-time rerun of an existing test.
	ReachabilityDetails *ReachabilityDetailsResponse `pulumi:"reachabilityDetails"`
	// Other projects that may be relevant for reachability analysis. This is applicable to scenarios where a test can cross project boundaries.
	RelatedProjects []string `pulumi:"relatedProjects"`
	// Required. Source specification of the Connectivity Test. You can use a combination of source IP address, virtual machine (VM) instance, or Compute Engine network to uniquely identify the source location. Examples: If the source IP address is an internal IP address within a Google Cloud Virtual Private Cloud (VPC) network, then you must also specify the VPC network. Otherwise, specify the VM instance, which already contains its internal IP address and VPC network information. If the source of the test is within an on-premises network, then you must provide the destination VPC network. If the source endpoint is a Compute Engine VM instance with multiple network interfaces, the instance itself is not sufficient to identify the endpoint. So, you must also specify the source IP address or VPC network. A reachability analysis proceeds even if the source location is ambiguous. However, the test result may include endpoints that you don't intend to test.
	Source *EndpointResponse `pulumi:"source"`
	// The time the test's configuration was updated.
	UpdateTime *string `pulumi:"updateTime"`
}

type ConnectivityTestState struct {
	// The time the test was created.
	CreateTime pulumi.StringPtrInput
	// The user-supplied description of the Connectivity Test. Maximum of 512 characters.
	Description pulumi.StringPtrInput
	// Required. Destination specification of the Connectivity Test. You can use a combination of destination IP address, Compute Engine VM instance, or VPC network to uniquely identify the destination location. Even if the destination IP address is not unique, the source IP location is unique. Usually, the analysis can infer the destination endpoint from route information. If the destination you specify is a VM instance and the instance has multiple network interfaces, then you must also specify either a destination IP address or VPC network to identify the destination interface. A reachability analysis proceeds even if the destination location is ambiguous. However, the result can include endpoints that you don't intend to test.
	Destination EndpointResponsePtrInput
	// The display name of a Connectivity Test.
	DisplayName pulumi.StringPtrInput
	// Resource labels to represent user-provided metadata.
	Labels pulumi.StringMapInput
	// Required. Unique name of the resource using the form: `projects/{project_id}/locations/global/connectivityTests/{test}`
	Name pulumi.StringPtrInput
	// The probing details of this test from the latest run, present for applicable tests only. The details are updated when creating a new test, updating an existing test, or triggering a one-time rerun of an existing test.
	ProbingDetails ProbingDetailsResponsePtrInput
	// IP Protocol of the test. When not provided, "TCP" is assumed.
	Protocol pulumi.StringPtrInput
	// The reachability details of this test from the latest run. The details are updated when creating a new test, updating an existing test, or triggering a one-time rerun of an existing test.
	ReachabilityDetails ReachabilityDetailsResponsePtrInput
	// Other projects that may be relevant for reachability analysis. This is applicable to scenarios where a test can cross project boundaries.
	RelatedProjects pulumi.StringArrayInput
	// Required. Source specification of the Connectivity Test. You can use a combination of source IP address, virtual machine (VM) instance, or Compute Engine network to uniquely identify the source location. Examples: If the source IP address is an internal IP address within a Google Cloud Virtual Private Cloud (VPC) network, then you must also specify the VPC network. Otherwise, specify the VM instance, which already contains its internal IP address and VPC network information. If the source of the test is within an on-premises network, then you must provide the destination VPC network. If the source endpoint is a Compute Engine VM instance with multiple network interfaces, the instance itself is not sufficient to identify the endpoint. So, you must also specify the source IP address or VPC network. A reachability analysis proceeds even if the source location is ambiguous. However, the test result may include endpoints that you don't intend to test.
	Source EndpointResponsePtrInput
	// The time the test's configuration was updated.
	UpdateTime pulumi.StringPtrInput
}

func (ConnectivityTestState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectivityTestState)(nil)).Elem()
}

type connectivityTestArgs struct {
	ConnectivityTestId string `pulumi:"connectivityTestId"`
	// The user-supplied description of the Connectivity Test. Maximum of 512 characters.
	Description *string `pulumi:"description"`
	// Required. Destination specification of the Connectivity Test. You can use a combination of destination IP address, Compute Engine VM instance, or VPC network to uniquely identify the destination location. Even if the destination IP address is not unique, the source IP location is unique. Usually, the analysis can infer the destination endpoint from route information. If the destination you specify is a VM instance and the instance has multiple network interfaces, then you must also specify either a destination IP address or VPC network to identify the destination interface. A reachability analysis proceeds even if the destination location is ambiguous. However, the result can include endpoints that you don't intend to test.
	Destination *Endpoint `pulumi:"destination"`
	// Resource labels to represent user-provided metadata.
	Labels map[string]string `pulumi:"labels"`
	// Required. Unique name of the resource using the form: `projects/{project_id}/locations/global/connectivityTests/{test}`
	Name    *string `pulumi:"name"`
	Project string  `pulumi:"project"`
	// IP Protocol of the test. When not provided, "TCP" is assumed.
	Protocol *string `pulumi:"protocol"`
	// Other projects that may be relevant for reachability analysis. This is applicable to scenarios where a test can cross project boundaries.
	RelatedProjects []string `pulumi:"relatedProjects"`
	// Required. Source specification of the Connectivity Test. You can use a combination of source IP address, virtual machine (VM) instance, or Compute Engine network to uniquely identify the source location. Examples: If the source IP address is an internal IP address within a Google Cloud Virtual Private Cloud (VPC) network, then you must also specify the VPC network. Otherwise, specify the VM instance, which already contains its internal IP address and VPC network information. If the source of the test is within an on-premises network, then you must provide the destination VPC network. If the source endpoint is a Compute Engine VM instance with multiple network interfaces, the instance itself is not sufficient to identify the endpoint. So, you must also specify the source IP address or VPC network. A reachability analysis proceeds even if the source location is ambiguous. However, the test result may include endpoints that you don't intend to test.
	Source *Endpoint `pulumi:"source"`
	TestId string    `pulumi:"testId"`
}

// The set of arguments for constructing a ConnectivityTest resource.
type ConnectivityTestArgs struct {
	ConnectivityTestId pulumi.StringInput
	// The user-supplied description of the Connectivity Test. Maximum of 512 characters.
	Description pulumi.StringPtrInput
	// Required. Destination specification of the Connectivity Test. You can use a combination of destination IP address, Compute Engine VM instance, or VPC network to uniquely identify the destination location. Even if the destination IP address is not unique, the source IP location is unique. Usually, the analysis can infer the destination endpoint from route information. If the destination you specify is a VM instance and the instance has multiple network interfaces, then you must also specify either a destination IP address or VPC network to identify the destination interface. A reachability analysis proceeds even if the destination location is ambiguous. However, the result can include endpoints that you don't intend to test.
	Destination EndpointPtrInput
	// Resource labels to represent user-provided metadata.
	Labels pulumi.StringMapInput
	// Required. Unique name of the resource using the form: `projects/{project_id}/locations/global/connectivityTests/{test}`
	Name    pulumi.StringPtrInput
	Project pulumi.StringInput
	// IP Protocol of the test. When not provided, "TCP" is assumed.
	Protocol pulumi.StringPtrInput
	// Other projects that may be relevant for reachability analysis. This is applicable to scenarios where a test can cross project boundaries.
	RelatedProjects pulumi.StringArrayInput
	// Required. Source specification of the Connectivity Test. You can use a combination of source IP address, virtual machine (VM) instance, or Compute Engine network to uniquely identify the source location. Examples: If the source IP address is an internal IP address within a Google Cloud Virtual Private Cloud (VPC) network, then you must also specify the VPC network. Otherwise, specify the VM instance, which already contains its internal IP address and VPC network information. If the source of the test is within an on-premises network, then you must provide the destination VPC network. If the source endpoint is a Compute Engine VM instance with multiple network interfaces, the instance itself is not sufficient to identify the endpoint. So, you must also specify the source IP address or VPC network. A reachability analysis proceeds even if the source location is ambiguous. However, the test result may include endpoints that you don't intend to test.
	Source EndpointPtrInput
	TestId pulumi.StringInput
}

func (ConnectivityTestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectivityTestArgs)(nil)).Elem()
}

type ConnectivityTestInput interface {
	pulumi.Input

	ToConnectivityTestOutput() ConnectivityTestOutput
	ToConnectivityTestOutputWithContext(ctx context.Context) ConnectivityTestOutput
}

func (*ConnectivityTest) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectivityTest)(nil))
}

func (i *ConnectivityTest) ToConnectivityTestOutput() ConnectivityTestOutput {
	return i.ToConnectivityTestOutputWithContext(context.Background())
}

func (i *ConnectivityTest) ToConnectivityTestOutputWithContext(ctx context.Context) ConnectivityTestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectivityTestOutput)
}

type ConnectivityTestOutput struct {
	*pulumi.OutputState
}

func (ConnectivityTestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectivityTest)(nil))
}

func (o ConnectivityTestOutput) ToConnectivityTestOutput() ConnectivityTestOutput {
	return o
}

func (o ConnectivityTestOutput) ToConnectivityTestOutputWithContext(ctx context.Context) ConnectivityTestOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ConnectivityTestOutput{})
}
