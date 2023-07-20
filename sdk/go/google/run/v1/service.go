// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new Service. Service creation will trigger a new deployment. Use GetService, and check service.status to determine if the Service is ready.
// Auto-naming is currently not supported for this resource.
type Service struct {
	pulumi.CustomResourceState

	// The API version for this call. It must be "serving.knative.dev/v1".
	ApiVersion pulumi.StringOutput `pulumi:"apiVersion"`
	// Indicates that the server should validate the request and populate default values without persisting the request. Supported values: `all`
	DryRun pulumi.StringPtrOutput `pulumi:"dryRun"`
	// The kind of resource. It must be "Service".
	Kind     pulumi.StringOutput `pulumi:"kind"`
	Location pulumi.StringOutput `pulumi:"location"`
	// Metadata associated with this Service, including name, namespace, labels, and annotations. In Cloud Run, annotations with 'run.googleapis.com/' and 'autoscaling.knative.dev' are restricted, and the accepted annotations will be different depending on the resource type. The following Cloud Run-specific annotations are accepted in Service.metadata.annotations. * `run.googleapis.com/binary-authorization-breakglass` * `run.googleapis.com/binary-authorization` * `run.googleapis.com/client-name` * `run.googleapis.com/custom-audiences` * `run.googleapis.com/description` * `run.googleapis.com/gc-traffic-tags` * `run.googleapis.com/ingress` * `run.googleapis.com/ingress` sets the ingress settings for the Service. See [the ingress settings documentation](/run/docs/securing/ingress) for details on configuring ingress settings. * `run.googleapis.com/ingress-status` is output-only and contains the currently active ingress settings for the Service. `run.googleapis.com/ingress-status` may differ from `run.googleapis.com/ingress` while the system is processing a change to `run.googleapis.com/ingress` or if the system failed to process a change to `run.googleapis.com/ingress`. When the system has processed all changes successfully `run.googleapis.com/ingress-status` and `run.googleapis.com/ingress` are equal.
	Metadata ObjectMetaResponseOutput `pulumi:"metadata"`
	Project  pulumi.StringOutput      `pulumi:"project"`
	// Holds the desired state of the Service (from the client).
	Spec ServiceSpecResponseOutput `pulumi:"spec"`
	// Communicates the system-controlled state of the Service.
	Status ServiceStatusResponseOutput `pulumi:"status"`
}

// NewService registers a new resource with the given unique name, arguments, and options.
func NewService(ctx *pulumi.Context,
	name string, args *ServiceArgs, opts ...pulumi.ResourceOption) (*Service, error) {
	if args == nil {
		args = &ServiceArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Service
	err := ctx.RegisterResource("google-native:run/v1:Service", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetService gets an existing Service resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceState, opts ...pulumi.ResourceOption) (*Service, error) {
	var resource Service
	err := ctx.ReadResource("google-native:run/v1:Service", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Service resources.
type serviceState struct {
}

type ServiceState struct {
}

func (ServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceState)(nil)).Elem()
}

type serviceArgs struct {
	// The API version for this call. It must be "serving.knative.dev/v1".
	ApiVersion *string `pulumi:"apiVersion"`
	// Indicates that the server should validate the request and populate default values without persisting the request. Supported values: `all`
	DryRun *string `pulumi:"dryRun"`
	// The kind of resource. It must be "Service".
	Kind     *string `pulumi:"kind"`
	Location *string `pulumi:"location"`
	// Metadata associated with this Service, including name, namespace, labels, and annotations. In Cloud Run, annotations with 'run.googleapis.com/' and 'autoscaling.knative.dev' are restricted, and the accepted annotations will be different depending on the resource type. The following Cloud Run-specific annotations are accepted in Service.metadata.annotations. * `run.googleapis.com/binary-authorization-breakglass` * `run.googleapis.com/binary-authorization` * `run.googleapis.com/client-name` * `run.googleapis.com/custom-audiences` * `run.googleapis.com/description` * `run.googleapis.com/gc-traffic-tags` * `run.googleapis.com/ingress` * `run.googleapis.com/ingress` sets the ingress settings for the Service. See [the ingress settings documentation](/run/docs/securing/ingress) for details on configuring ingress settings. * `run.googleapis.com/ingress-status` is output-only and contains the currently active ingress settings for the Service. `run.googleapis.com/ingress-status` may differ from `run.googleapis.com/ingress` while the system is processing a change to `run.googleapis.com/ingress` or if the system failed to process a change to `run.googleapis.com/ingress`. When the system has processed all changes successfully `run.googleapis.com/ingress-status` and `run.googleapis.com/ingress` are equal.
	Metadata *ObjectMeta `pulumi:"metadata"`
	Project  *string     `pulumi:"project"`
	// Holds the desired state of the Service (from the client).
	Spec *ServiceSpec `pulumi:"spec"`
}

// The set of arguments for constructing a Service resource.
type ServiceArgs struct {
	// The API version for this call. It must be "serving.knative.dev/v1".
	ApiVersion pulumi.StringPtrInput
	// Indicates that the server should validate the request and populate default values without persisting the request. Supported values: `all`
	DryRun pulumi.StringPtrInput
	// The kind of resource. It must be "Service".
	Kind     pulumi.StringPtrInput
	Location pulumi.StringPtrInput
	// Metadata associated with this Service, including name, namespace, labels, and annotations. In Cloud Run, annotations with 'run.googleapis.com/' and 'autoscaling.knative.dev' are restricted, and the accepted annotations will be different depending on the resource type. The following Cloud Run-specific annotations are accepted in Service.metadata.annotations. * `run.googleapis.com/binary-authorization-breakglass` * `run.googleapis.com/binary-authorization` * `run.googleapis.com/client-name` * `run.googleapis.com/custom-audiences` * `run.googleapis.com/description` * `run.googleapis.com/gc-traffic-tags` * `run.googleapis.com/ingress` * `run.googleapis.com/ingress` sets the ingress settings for the Service. See [the ingress settings documentation](/run/docs/securing/ingress) for details on configuring ingress settings. * `run.googleapis.com/ingress-status` is output-only and contains the currently active ingress settings for the Service. `run.googleapis.com/ingress-status` may differ from `run.googleapis.com/ingress` while the system is processing a change to `run.googleapis.com/ingress` or if the system failed to process a change to `run.googleapis.com/ingress`. When the system has processed all changes successfully `run.googleapis.com/ingress-status` and `run.googleapis.com/ingress` are equal.
	Metadata ObjectMetaPtrInput
	Project  pulumi.StringPtrInput
	// Holds the desired state of the Service (from the client).
	Spec ServiceSpecPtrInput
}

func (ServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceArgs)(nil)).Elem()
}

type ServiceInput interface {
	pulumi.Input

	ToServiceOutput() ServiceOutput
	ToServiceOutputWithContext(ctx context.Context) ServiceOutput
}

func (*Service) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (i *Service) ToServiceOutput() ServiceOutput {
	return i.ToServiceOutputWithContext(context.Background())
}

func (i *Service) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceOutput)
}

type ServiceOutput struct{ *pulumi.OutputState }

func (ServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Service)(nil)).Elem()
}

func (o ServiceOutput) ToServiceOutput() ServiceOutput {
	return o
}

func (o ServiceOutput) ToServiceOutputWithContext(ctx context.Context) ServiceOutput {
	return o
}

// The API version for this call. It must be "serving.knative.dev/v1".
func (o ServiceOutput) ApiVersion() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.ApiVersion }).(pulumi.StringOutput)
}

// Indicates that the server should validate the request and populate default values without persisting the request. Supported values: `all`
func (o ServiceOutput) DryRun() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Service) pulumi.StringPtrOutput { return v.DryRun }).(pulumi.StringPtrOutput)
}

// The kind of resource. It must be "Service".
func (o ServiceOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o ServiceOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// Metadata associated with this Service, including name, namespace, labels, and annotations. In Cloud Run, annotations with 'run.googleapis.com/' and 'autoscaling.knative.dev' are restricted, and the accepted annotations will be different depending on the resource type. The following Cloud Run-specific annotations are accepted in Service.metadata.annotations. * `run.googleapis.com/binary-authorization-breakglass` * `run.googleapis.com/binary-authorization` * `run.googleapis.com/client-name` * `run.googleapis.com/custom-audiences` * `run.googleapis.com/description` * `run.googleapis.com/gc-traffic-tags` * `run.googleapis.com/ingress` * `run.googleapis.com/ingress` sets the ingress settings for the Service. See [the ingress settings documentation](/run/docs/securing/ingress) for details on configuring ingress settings. * `run.googleapis.com/ingress-status` is output-only and contains the currently active ingress settings for the Service. `run.googleapis.com/ingress-status` may differ from `run.googleapis.com/ingress` while the system is processing a change to `run.googleapis.com/ingress` or if the system failed to process a change to `run.googleapis.com/ingress`. When the system has processed all changes successfully `run.googleapis.com/ingress-status` and `run.googleapis.com/ingress` are equal.
func (o ServiceOutput) Metadata() ObjectMetaResponseOutput {
	return o.ApplyT(func(v *Service) ObjectMetaResponseOutput { return v.Metadata }).(ObjectMetaResponseOutput)
}

func (o ServiceOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Service) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Holds the desired state of the Service (from the client).
func (o ServiceOutput) Spec() ServiceSpecResponseOutput {
	return o.ApplyT(func(v *Service) ServiceSpecResponseOutput { return v.Spec }).(ServiceSpecResponseOutput)
}

// Communicates the system-controlled state of the Service.
func (o ServiceOutput) Status() ServiceStatusResponseOutput {
	return o.ApplyT(func(v *Service) ServiceStatusResponseOutput { return v.Status }).(ServiceStatusResponseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceInput)(nil)).Elem(), &Service{})
	pulumi.RegisterOutputType(ServiceOutput{})
}
