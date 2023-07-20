// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates new workflow template.
// Auto-naming is currently not supported for this resource.
type WorkflowTemplate struct {
	pulumi.CustomResourceState

	// The time template was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// Optional. Timeout duration for the DAG of jobs, expressed in seconds (see JSON representation of duration (https://developers.google.com/protocol-buffers/docs/proto3#json)). The timeout duration must be from 10 minutes ("600s") to 24 hours ("86400s"). The timer begins when the first job is submitted. If the workflow is running at the end of the timeout period, any remaining jobs are cancelled, the workflow is ended, and if the workflow was running on a managed cluster, the cluster is deleted.
	DagTimeout pulumi.StringOutput `pulumi:"dagTimeout"`
	// The Directed Acyclic Graph of Jobs to submit.
	Jobs OrderedJobResponseArrayOutput `pulumi:"jobs"`
	// Optional. The labels to associate with this template. These labels will be propagated to all jobs and clusters created by the workflow instance.Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt).Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt).No more than 32 labels can be associated with a template.
	Labels   pulumi.StringMapOutput `pulumi:"labels"`
	Location pulumi.StringOutput    `pulumi:"location"`
	// The resource name of the workflow template, as described in https://cloud.google.com/apis/design/resource_names. For projects.regions.workflowTemplates, the resource name of the template has the following format: projects/{project_id}/regions/{region}/workflowTemplates/{template_id} For projects.locations.workflowTemplates, the resource name of the template has the following format: projects/{project_id}/locations/{location}/workflowTemplates/{template_id}
	Name pulumi.StringOutput `pulumi:"name"`
	// Optional. Template parameters whose values are substituted into the template. Values for parameters must be provided when the template is instantiated.
	Parameters TemplateParameterResponseArrayOutput `pulumi:"parameters"`
	// WorkflowTemplate scheduling information.
	Placement WorkflowTemplatePlacementResponseOutput `pulumi:"placement"`
	Project   pulumi.StringOutput                     `pulumi:"project"`
	// The time template was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
	// Optional. Used to perform a consistent read-modify-write.This field should be left blank for a CreateWorkflowTemplate request. It is required for an UpdateWorkflowTemplate request, and must match the current server version. A typical update template flow would fetch the current template with a GetWorkflowTemplate request, which will return the current template with the version field filled in with the current server version. The user updates other fields in the template, then returns it as part of the UpdateWorkflowTemplate request.
	Version pulumi.IntOutput `pulumi:"version"`
}

// NewWorkflowTemplate registers a new resource with the given unique name, arguments, and options.
func NewWorkflowTemplate(ctx *pulumi.Context,
	name string, args *WorkflowTemplateArgs, opts ...pulumi.ResourceOption) (*WorkflowTemplate, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Id == nil {
		return nil, errors.New("invalid value for required argument 'Id'")
	}
	if args.Jobs == nil {
		return nil, errors.New("invalid value for required argument 'Jobs'")
	}
	if args.Placement == nil {
		return nil, errors.New("invalid value for required argument 'Placement'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource WorkflowTemplate
	err := ctx.RegisterResource("google-native:dataproc/v1beta2:WorkflowTemplate", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetWorkflowTemplate gets an existing WorkflowTemplate resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetWorkflowTemplate(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkflowTemplateState, opts ...pulumi.ResourceOption) (*WorkflowTemplate, error) {
	var resource WorkflowTemplate
	err := ctx.ReadResource("google-native:dataproc/v1beta2:WorkflowTemplate", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering WorkflowTemplate resources.
type workflowTemplateState struct {
}

type WorkflowTemplateState struct {
}

func (WorkflowTemplateState) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowTemplateState)(nil)).Elem()
}

type workflowTemplateArgs struct {
	// Optional. Timeout duration for the DAG of jobs, expressed in seconds (see JSON representation of duration (https://developers.google.com/protocol-buffers/docs/proto3#json)). The timeout duration must be from 10 minutes ("600s") to 24 hours ("86400s"). The timer begins when the first job is submitted. If the workflow is running at the end of the timeout period, any remaining jobs are cancelled, the workflow is ended, and if the workflow was running on a managed cluster, the cluster is deleted.
	DagTimeout *string `pulumi:"dagTimeout"`
	// The template id.The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between 3 and 50 characters..
	Id string `pulumi:"id"`
	// The Directed Acyclic Graph of Jobs to submit.
	Jobs []OrderedJob `pulumi:"jobs"`
	// Optional. The labels to associate with this template. These labels will be propagated to all jobs and clusters created by the workflow instance.Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt).Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt).No more than 32 labels can be associated with a template.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	// Optional. Template parameters whose values are substituted into the template. Values for parameters must be provided when the template is instantiated.
	Parameters []TemplateParameter `pulumi:"parameters"`
	// WorkflowTemplate scheduling information.
	Placement WorkflowTemplatePlacement `pulumi:"placement"`
	Project   *string                   `pulumi:"project"`
	// Optional. Used to perform a consistent read-modify-write.This field should be left blank for a CreateWorkflowTemplate request. It is required for an UpdateWorkflowTemplate request, and must match the current server version. A typical update template flow would fetch the current template with a GetWorkflowTemplate request, which will return the current template with the version field filled in with the current server version. The user updates other fields in the template, then returns it as part of the UpdateWorkflowTemplate request.
	Version *int `pulumi:"version"`
}

// The set of arguments for constructing a WorkflowTemplate resource.
type WorkflowTemplateArgs struct {
	// Optional. Timeout duration for the DAG of jobs, expressed in seconds (see JSON representation of duration (https://developers.google.com/protocol-buffers/docs/proto3#json)). The timeout duration must be from 10 minutes ("600s") to 24 hours ("86400s"). The timer begins when the first job is submitted. If the workflow is running at the end of the timeout period, any remaining jobs are cancelled, the workflow is ended, and if the workflow was running on a managed cluster, the cluster is deleted.
	DagTimeout pulumi.StringPtrInput
	// The template id.The id must contain only letters (a-z, A-Z), numbers (0-9), underscores (_), and hyphens (-). Cannot begin or end with underscore or hyphen. Must consist of between 3 and 50 characters..
	Id pulumi.StringInput
	// The Directed Acyclic Graph of Jobs to submit.
	Jobs OrderedJobArrayInput
	// Optional. The labels to associate with this template. These labels will be propagated to all jobs and clusters created by the workflow instance.Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt).Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt).No more than 32 labels can be associated with a template.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	// Optional. Template parameters whose values are substituted into the template. Values for parameters must be provided when the template is instantiated.
	Parameters TemplateParameterArrayInput
	// WorkflowTemplate scheduling information.
	Placement WorkflowTemplatePlacementInput
	Project   pulumi.StringPtrInput
	// Optional. Used to perform a consistent read-modify-write.This field should be left blank for a CreateWorkflowTemplate request. It is required for an UpdateWorkflowTemplate request, and must match the current server version. A typical update template flow would fetch the current template with a GetWorkflowTemplate request, which will return the current template with the version field filled in with the current server version. The user updates other fields in the template, then returns it as part of the UpdateWorkflowTemplate request.
	Version pulumi.IntPtrInput
}

func (WorkflowTemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowTemplateArgs)(nil)).Elem()
}

type WorkflowTemplateInput interface {
	pulumi.Input

	ToWorkflowTemplateOutput() WorkflowTemplateOutput
	ToWorkflowTemplateOutputWithContext(ctx context.Context) WorkflowTemplateOutput
}

func (*WorkflowTemplate) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowTemplate)(nil)).Elem()
}

func (i *WorkflowTemplate) ToWorkflowTemplateOutput() WorkflowTemplateOutput {
	return i.ToWorkflowTemplateOutputWithContext(context.Background())
}

func (i *WorkflowTemplate) ToWorkflowTemplateOutputWithContext(ctx context.Context) WorkflowTemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowTemplateOutput)
}

type WorkflowTemplateOutput struct{ *pulumi.OutputState }

func (WorkflowTemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowTemplate)(nil)).Elem()
}

func (o WorkflowTemplateOutput) ToWorkflowTemplateOutput() WorkflowTemplateOutput {
	return o
}

func (o WorkflowTemplateOutput) ToWorkflowTemplateOutputWithContext(ctx context.Context) WorkflowTemplateOutput {
	return o
}

// The time template was created.
func (o WorkflowTemplateOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowTemplate) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// Optional. Timeout duration for the DAG of jobs, expressed in seconds (see JSON representation of duration (https://developers.google.com/protocol-buffers/docs/proto3#json)). The timeout duration must be from 10 minutes ("600s") to 24 hours ("86400s"). The timer begins when the first job is submitted. If the workflow is running at the end of the timeout period, any remaining jobs are cancelled, the workflow is ended, and if the workflow was running on a managed cluster, the cluster is deleted.
func (o WorkflowTemplateOutput) DagTimeout() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowTemplate) pulumi.StringOutput { return v.DagTimeout }).(pulumi.StringOutput)
}

// The Directed Acyclic Graph of Jobs to submit.
func (o WorkflowTemplateOutput) Jobs() OrderedJobResponseArrayOutput {
	return o.ApplyT(func(v *WorkflowTemplate) OrderedJobResponseArrayOutput { return v.Jobs }).(OrderedJobResponseArrayOutput)
}

// Optional. The labels to associate with this template. These labels will be propagated to all jobs and clusters created by the workflow instance.Label keys must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt).Label values may be empty, but, if present, must contain 1 to 63 characters, and must conform to RFC 1035 (https://www.ietf.org/rfc/rfc1035.txt).No more than 32 labels can be associated with a template.
func (o WorkflowTemplateOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v *WorkflowTemplate) pulumi.StringMapOutput { return v.Labels }).(pulumi.StringMapOutput)
}

func (o WorkflowTemplateOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowTemplate) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The resource name of the workflow template, as described in https://cloud.google.com/apis/design/resource_names. For projects.regions.workflowTemplates, the resource name of the template has the following format: projects/{project_id}/regions/{region}/workflowTemplates/{template_id} For projects.locations.workflowTemplates, the resource name of the template has the following format: projects/{project_id}/locations/{location}/workflowTemplates/{template_id}
func (o WorkflowTemplateOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowTemplate) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional. Template parameters whose values are substituted into the template. Values for parameters must be provided when the template is instantiated.
func (o WorkflowTemplateOutput) Parameters() TemplateParameterResponseArrayOutput {
	return o.ApplyT(func(v *WorkflowTemplate) TemplateParameterResponseArrayOutput { return v.Parameters }).(TemplateParameterResponseArrayOutput)
}

// WorkflowTemplate scheduling information.
func (o WorkflowTemplateOutput) Placement() WorkflowTemplatePlacementResponseOutput {
	return o.ApplyT(func(v *WorkflowTemplate) WorkflowTemplatePlacementResponseOutput { return v.Placement }).(WorkflowTemplatePlacementResponseOutput)
}

func (o WorkflowTemplateOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowTemplate) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The time template was last updated.
func (o WorkflowTemplateOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *WorkflowTemplate) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

// Optional. Used to perform a consistent read-modify-write.This field should be left blank for a CreateWorkflowTemplate request. It is required for an UpdateWorkflowTemplate request, and must match the current server version. A typical update template flow would fetch the current template with a GetWorkflowTemplate request, which will return the current template with the version field filled in with the current server version. The user updates other fields in the template, then returns it as part of the UpdateWorkflowTemplate request.
func (o WorkflowTemplateOutput) Version() pulumi.IntOutput {
	return o.ApplyT(func(v *WorkflowTemplate) pulumi.IntOutput { return v.Version }).(pulumi.IntOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*WorkflowTemplateInput)(nil)).Elem(), &WorkflowTemplate{})
	pulumi.RegisterOutputType(WorkflowTemplateOutput{})
}
