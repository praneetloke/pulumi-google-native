// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a Version in the specified Flow.
type Version struct {
	pulumi.CustomResourceState

	// Create time of the version.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The description of the version. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description pulumi.StringOutput `pulumi:"description"`
	// The human-readable name of the version. Limit of 64 characters.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Format: projects//locations//agents//flows//versions/. Version ID is a self-increasing number generated by Dialogflow upon version creation.
	Name pulumi.StringOutput `pulumi:"name"`
	// The NLU settings of the flow at version creation.
	NluSettings GoogleCloudDialogflowCxV3NluSettingsResponseOutput `pulumi:"nluSettings"`
	// The state of this version. This field is read-only and cannot be set by create and update methods.
	State pulumi.StringOutput `pulumi:"state"`
}

// NewVersion registers a new resource with the given unique name, arguments, and options.
func NewVersion(ctx *pulumi.Context,
	name string, args *VersionArgs, opts ...pulumi.ResourceOption) (*Version, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AgentId == nil {
		return nil, errors.New("invalid value for required argument 'AgentId'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.FlowId == nil {
		return nil, errors.New("invalid value for required argument 'FlowId'")
	}
	var resource Version
	err := ctx.RegisterResource("google-native:dialogflow/v3:Version", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetVersion gets an existing Version resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetVersion(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VersionState, opts ...pulumi.ResourceOption) (*Version, error) {
	var resource Version
	err := ctx.ReadResource("google-native:dialogflow/v3:Version", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Version resources.
type versionState struct {
}

type VersionState struct {
}

func (VersionState) ElementType() reflect.Type {
	return reflect.TypeOf((*versionState)(nil)).Elem()
}

type versionArgs struct {
	AgentId string `pulumi:"agentId"`
	// The description of the version. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description *string `pulumi:"description"`
	// The human-readable name of the version. Limit of 64 characters.
	DisplayName string  `pulumi:"displayName"`
	FlowId      string  `pulumi:"flowId"`
	Location    *string `pulumi:"location"`
	// Format: projects//locations//agents//flows//versions/. Version ID is a self-increasing number generated by Dialogflow upon version creation.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a Version resource.
type VersionArgs struct {
	AgentId pulumi.StringInput
	// The description of the version. The maximum length is 500 characters. If exceeded, the request is rejected.
	Description pulumi.StringPtrInput
	// The human-readable name of the version. Limit of 64 characters.
	DisplayName pulumi.StringInput
	FlowId      pulumi.StringInput
	Location    pulumi.StringPtrInput
	// Format: projects//locations//agents//flows//versions/. Version ID is a self-increasing number generated by Dialogflow upon version creation.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
}

func (VersionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*versionArgs)(nil)).Elem()
}

type VersionInput interface {
	pulumi.Input

	ToVersionOutput() VersionOutput
	ToVersionOutputWithContext(ctx context.Context) VersionOutput
}

func (*Version) ElementType() reflect.Type {
	return reflect.TypeOf((*Version)(nil))
}

func (i *Version) ToVersionOutput() VersionOutput {
	return i.ToVersionOutputWithContext(context.Background())
}

func (i *Version) ToVersionOutputWithContext(ctx context.Context) VersionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VersionOutput)
}

type VersionOutput struct {
	*pulumi.OutputState
}

func (VersionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Version)(nil))
}

func (o VersionOutput) ToVersionOutput() VersionOutput {
	return o
}

func (o VersionOutput) ToVersionOutputWithContext(ctx context.Context) VersionOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VersionOutput{})
}
