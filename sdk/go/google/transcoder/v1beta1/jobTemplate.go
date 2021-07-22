// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a job template in the specified region.
type JobTemplate struct {
	pulumi.CustomResourceState

	// The configuration for this template.
	Config JobConfigResponseOutput `pulumi:"config"`
	// The resource name of the job template. Format: `projects/{project}/locations/{location}/jobTemplates/{job_template}`
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewJobTemplate registers a new resource with the given unique name, arguments, and options.
func NewJobTemplate(ctx *pulumi.Context,
	name string, args *JobTemplateArgs, opts ...pulumi.ResourceOption) (*JobTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.JobTemplateId == nil {
		return nil, errors.New("invalid value for required argument 'JobTemplateId'")
	}
	var resource JobTemplate
	err := ctx.RegisterResource("google-native:transcoder/v1beta1:JobTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetJobTemplate gets an existing JobTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetJobTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobTemplateState, opts ...pulumi.ResourceOption) (*JobTemplate, error) {
	var resource JobTemplate
	err := ctx.ReadResource("google-native:transcoder/v1beta1:JobTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering JobTemplate resources.
type jobTemplateState struct {
}

type JobTemplateState struct {
}

func (JobTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobTemplateState)(nil)).Elem()
}

type jobTemplateArgs struct {
	// The configuration for this template.
	Config        *JobConfig `pulumi:"config"`
	JobTemplateId string     `pulumi:"jobTemplateId"`
	Location      *string    `pulumi:"location"`
	// The resource name of the job template. Format: `projects/{project}/locations/{location}/jobTemplates/{job_template}`
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
}

// The set of arguments for constructing a JobTemplate resource.
type JobTemplateArgs struct {
	// The configuration for this template.
	Config        JobConfigPtrInput
	JobTemplateId pulumi.StringInput
	Location      pulumi.StringPtrInput
	// The resource name of the job template. Format: `projects/{project}/locations/{location}/jobTemplates/{job_template}`
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
}

func (JobTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobTemplateArgs)(nil)).Elem()
}

type JobTemplateInput interface {
	pulumi.Input

	ToJobTemplateOutput() JobTemplateOutput
	ToJobTemplateOutputWithContext(ctx context.Context) JobTemplateOutput
}

func (*JobTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((*JobTemplate)(nil))
}

func (i *JobTemplate) ToJobTemplateOutput() JobTemplateOutput {
	return i.ToJobTemplateOutputWithContext(context.Background())
}

func (i *JobTemplate) ToJobTemplateOutputWithContext(ctx context.Context) JobTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobTemplateOutput)
}

type JobTemplateOutput struct {
	*pulumi.OutputState
}

func (JobTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobTemplate)(nil))
}

func (o JobTemplateOutput) ToJobTemplateOutput() JobTemplateOutput {
	return o
}

func (o JobTemplateOutput) ToJobTemplateOutputWithContext(ctx context.Context) JobTemplateOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(JobTemplateOutput{})
}
