// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new Release in a given project and location.
// Auto-naming is currently not supported for this resource.
// Note - this resource's API doesn't support deletion. When deleted, the resource will persist
// on Google Cloud even though it will be deleted from Pulumi state.
type Release struct {
	pulumi.CustomResourceState

	// User annotations. These attributes can only be set and used by the user, and not by Google Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
	Annotations pulumi.StringMapOutput `pulumi:"annotations"`
	// List of artifacts to pass through to Skaffold command.
	BuildArtifacts BuildArtifactResponseArrayOutput `pulumi:"buildArtifacts"`
	// Time at which the `Release` was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Snapshot of the parent pipeline taken at release creation time.
	DeliveryPipelineSnapshot DeliveryPipelineResponseOutput `pulumi:"deliveryPipelineSnapshot"`
	// Description of the `Release`. Max length is 255 characters.
	Description pulumi.StringOutput `pulumi:"description"`
	// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Labels are attributes that can be set and used by both the user and by Google Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be <= 128 bytes.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Optional. Name of the `Release`. Format is projects/{project}/ locations/{location}/deliveryPipelines/{deliveryPipeline}/ releases/a-z{0,62}.
	Name pulumi.StringOutput `pulumi:"name"`
	// Time at which the render completed.
	RenderEndTime pulumi.StringOutput `pulumi:"renderEndTime"`
	// Time at which the render began.
	RenderStartTime pulumi.StringOutput `pulumi:"renderStartTime"`
	// Current state of the render operation.
	RenderState pulumi.StringOutput `pulumi:"renderState"`
	// Filepath of the Skaffold config inside of the config URI.
	SkaffoldConfigPath pulumi.StringOutput `pulumi:"skaffoldConfigPath"`
	// Cloud Storage URI of tar.gz archive containing Skaffold configuration.
	SkaffoldConfigUri pulumi.StringOutput `pulumi:"skaffoldConfigUri"`
	// The Skaffold version to use when operating on this release, such as "1.20.0". Not all versions are valid; Google Cloud Deploy supports a specific set of versions. If unset, the most recent supported Skaffold version will be used.
	SkaffoldVersion pulumi.StringOutput `pulumi:"skaffoldVersion"`
	// Map from target ID to the target artifacts created during the render operation.
	TargetArtifacts pulumi.StringMapOutput `pulumi:"targetArtifacts"`
	// Map from target ID to details of the render operation for that target.
	TargetRenders pulumi.StringMapOutput `pulumi:"targetRenders"`
	// Snapshot of the targets taken at release creation time.
	TargetSnapshots TargetResponseArrayOutput `pulumi:"targetSnapshots"`
	// Unique identifier of the `Release`.
	Uid pulumi.StringOutput `pulumi:"uid"`
}

// NewRelease registers a new resource with the given unique name, arguments, and options.
func NewRelease(ctx *pulumi.Context,
	name string, args *ReleaseArgs, opts ...pulumi.ResourceOption) (*Release, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DeliveryPipelineId == nil {
		return nil, errors.New("invalid value for required argument 'DeliveryPipelineId'")
	}
	if args.ReleaseId == nil {
		return nil, errors.New("invalid value for required argument 'ReleaseId'")
	}
	var resource Release
	err := ctx.RegisterResource("google-native:clouddeploy/v1:Release", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRelease gets an existing Release resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRelease(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ReleaseState, opts ...pulumi.ResourceOption) (*Release, error) {
	var resource Release
	err := ctx.ReadResource("google-native:clouddeploy/v1:Release", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Release resources.
type releaseState struct {
}

type ReleaseState struct {
}

func (ReleaseState) ElementType() reflect.Type {
	return reflect.TypeOf((*releaseState)(nil)).Elem()
}

type releaseArgs struct {
	// User annotations. These attributes can only be set and used by the user, and not by Google Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
	Annotations map[string]string `pulumi:"annotations"`
	// List of artifacts to pass through to Skaffold command.
	BuildArtifacts     []BuildArtifact `pulumi:"buildArtifacts"`
	DeliveryPipelineId string          `pulumi:"deliveryPipelineId"`
	// Description of the `Release`. Max length is 255 characters.
	Description *string `pulumi:"description"`
	// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag *string `pulumi:"etag"`
	// Labels are attributes that can be set and used by both the user and by Google Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be <= 128 bytes.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	// Optional. Name of the `Release`. Format is projects/{project}/ locations/{location}/deliveryPipelines/{deliveryPipeline}/ releases/a-z{0,62}.
	Name      *string `pulumi:"name"`
	Project   *string `pulumi:"project"`
	ReleaseId string  `pulumi:"releaseId"`
	RequestId *string `pulumi:"requestId"`
	// Filepath of the Skaffold config inside of the config URI.
	SkaffoldConfigPath *string `pulumi:"skaffoldConfigPath"`
	// Cloud Storage URI of tar.gz archive containing Skaffold configuration.
	SkaffoldConfigUri *string `pulumi:"skaffoldConfigUri"`
	// The Skaffold version to use when operating on this release, such as "1.20.0". Not all versions are valid; Google Cloud Deploy supports a specific set of versions. If unset, the most recent supported Skaffold version will be used.
	SkaffoldVersion *string `pulumi:"skaffoldVersion"`
	ValidateOnly    *string `pulumi:"validateOnly"`
}

// The set of arguments for constructing a Release resource.
type ReleaseArgs struct {
	// User annotations. These attributes can only be set and used by the user, and not by Google Cloud Deploy. See https://google.aip.dev/128#annotations for more details such as format and size limitations.
	Annotations pulumi.StringMapInput
	// List of artifacts to pass through to Skaffold command.
	BuildArtifacts     BuildArtifactArrayInput
	DeliveryPipelineId pulumi.StringInput
	// Description of the `Release`. Max length is 255 characters.
	Description pulumi.StringPtrInput
	// This checksum is computed by the server based on the value of other fields, and may be sent on update and delete requests to ensure the client has an up-to-date value before proceeding.
	Etag pulumi.StringPtrInput
	// Labels are attributes that can be set and used by both the user and by Google Cloud Deploy. Labels must meet the following constraints: * Keys and values can contain only lowercase letters, numeric characters, underscores, and dashes. * All characters must use UTF-8 encoding, and international characters are allowed. * Keys must start with a lowercase letter or international character. * Each resource is limited to a maximum of 64 labels. Both keys and values are additionally constrained to be <= 128 bytes.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	// Optional. Name of the `Release`. Format is projects/{project}/ locations/{location}/deliveryPipelines/{deliveryPipeline}/ releases/a-z{0,62}.
	Name      pulumi.StringPtrInput
	Project   pulumi.StringPtrInput
	ReleaseId pulumi.StringInput
	RequestId pulumi.StringPtrInput
	// Filepath of the Skaffold config inside of the config URI.
	SkaffoldConfigPath pulumi.StringPtrInput
	// Cloud Storage URI of tar.gz archive containing Skaffold configuration.
	SkaffoldConfigUri pulumi.StringPtrInput
	// The Skaffold version to use when operating on this release, such as "1.20.0". Not all versions are valid; Google Cloud Deploy supports a specific set of versions. If unset, the most recent supported Skaffold version will be used.
	SkaffoldVersion pulumi.StringPtrInput
	ValidateOnly    pulumi.StringPtrInput
}

func (ReleaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*releaseArgs)(nil)).Elem()
}

type ReleaseInput interface {
	pulumi.Input

	ToReleaseOutput() ReleaseOutput
	ToReleaseOutputWithContext(ctx context.Context) ReleaseOutput
}

func (*Release) ElementType() reflect.Type {
	return reflect.TypeOf((**Release)(nil)).Elem()
}

func (i *Release) ToReleaseOutput() ReleaseOutput {
	return i.ToReleaseOutputWithContext(context.Background())
}

func (i *Release) ToReleaseOutputWithContext(ctx context.Context) ReleaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ReleaseOutput)
}

type ReleaseOutput struct{ *pulumi.OutputState }

func (ReleaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Release)(nil)).Elem()
}

func (o ReleaseOutput) ToReleaseOutput() ReleaseOutput {
	return o
}

func (o ReleaseOutput) ToReleaseOutputWithContext(ctx context.Context) ReleaseOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ReleaseInput)(nil)).Elem(), &Release{})
	pulumi.RegisterOutputType(ReleaseOutput{})
}
