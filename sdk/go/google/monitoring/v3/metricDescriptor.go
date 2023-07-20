// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v3

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new metric descriptor. The creation is executed asynchronously. User-created metric descriptors define custom metrics (https://cloud.google.com/monitoring/custom-metrics). The metric descriptor is updated if it already exists, except that metric labels are never removed.
type MetricDescriptor struct {
	pulumi.CustomResourceState

	// A detailed description of the metric, which can be used in documentation.
	Description pulumi.StringOutput `pulumi:"description"`
	// A concise name for the metric, which can be displayed in user interfaces. Use sentence case without an ending period, for example "Request count". This field is optional but it is recommended to be set for any metrics associated with user-visible concepts, such as Quota.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The set of labels that can be used to describe a specific instance of this metric type. For example, the appengine.googleapis.com/http/server/response_latencies metric type has a label for the HTTP response code, response_code, so you can look at latencies for successful responses or just for responses that failed.
	Labels LabelDescriptorResponseArrayOutput `pulumi:"labels"`
	// Optional. The launch stage of the metric definition.
	LaunchStage pulumi.StringOutput `pulumi:"launchStage"`
	// Optional. Metadata which can be used to guide usage of the metric.
	Metadata MetricDescriptorMetadataResponseOutput `pulumi:"metadata"`
	// Whether the metric records instantaneous values, changes to a value, etc. Some combinations of metric_kind and value_type might not be supported.
	MetricKind pulumi.StringOutput `pulumi:"metricKind"`
	// Read-only. If present, then a time series, which is identified partially by a metric type and a MonitoredResourceDescriptor, that is associated with this metric type can only be associated with one of the monitored resource types listed here.
	MonitoredResourceTypes pulumi.StringArrayOutput `pulumi:"monitoredResourceTypes"`
	// The resource name of the metric descriptor.
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// The metric type, including its DNS name prefix. The type is not URL-encoded. All user-defined metric types have the DNS name custom.googleapis.com or external.googleapis.com. Metric types should use a natural hierarchical grouping. For example: "custom.googleapis.com/invoice/paid/amount" "external.googleapis.com/prometheus/up" "appengine.googleapis.com/http/server/response_latencies"
	Type pulumi.StringOutput `pulumi:"type"`
	// The units in which the metric value is reported. It is only applicable if the value_type is INT64, DOUBLE, or DISTRIBUTION. The unit defines the representation of the stored metric values.Different systems might scale the values to be more easily displayed (so a value of 0.02kBy might be displayed as 20By, and a value of 3523kBy might be displayed as 3.5MBy). However, if the unit is kBy, then the value of the metric is always in thousands of bytes, no matter how it might be displayed.If you want a custom metric to record the exact number of CPU-seconds used by a job, you can create an INT64 CUMULATIVE metric whose unit is s{CPU} (or equivalently 1s{CPU} or just s). If the job uses 12,005 CPU-seconds, then the value is written as 12005.Alternatively, if you want a custom metric to record data in a more granular way, you can create a DOUBLE CUMULATIVE metric whose unit is ks{CPU}, and then write the value 12.005 (which is 12005/1000), or use Kis{CPU} and write 11.723 (which is 12005/1024).The supported units are a subset of The Unified Code for Units of Measure (https://unitsofmeasure.org/ucum.html) standard:Basic units (UNIT) bit bit By byte s second min minute h hour d day 1 dimensionlessPrefixes (PREFIX) k kilo (10^3) M mega (10^6) G giga (10^9) T tera (10^12) P peta (10^15) E exa (10^18) Z zetta (10^21) Y yotta (10^24) m milli (10^-3) u micro (10^-6) n nano (10^-9) p pico (10^-12) f femto (10^-15) a atto (10^-18) z zepto (10^-21) y yocto (10^-24) Ki kibi (2^10) Mi mebi (2^20) Gi gibi (2^30) Ti tebi (2^40) Pi pebi (2^50)GrammarThe grammar also includes these connectors: / division or ratio (as an infix operator). For examples, kBy/{email} or MiBy/10ms (although you should almost never have /s in a metric unit; rates should always be computed at query time from the underlying cumulative or delta value). . multiplication or composition (as an infix operator). For examples, GBy.d or k{watt}.h.The grammar for a unit is as follows: Expression = Component { "." Component } { "/" Component } ; Component = ( [ PREFIX ] UNIT | "%" ) [ Annotation ] | Annotation | "1" ; Annotation = "{" NAME "}" ; Notes: Annotation is just a comment if it follows a UNIT. If the annotation is used alone, then the unit is equivalent to 1. For examples, {request}/s == 1/s, By{transmitted}/s == By/s. NAME is a sequence of non-blank printable ASCII characters not containing { or }. 1 represents a unitary dimensionless unit (https://en.wikipedia.org/wiki/Dimensionless_quantity) of 1, such as in 1/s. It is typically used when none of the basic units are appropriate. For example, "new users per day" can be represented as 1/d or {new-users}/d (and a metric value 5 would mean "5 new users). Alternatively, "thousands of page views per day" would be represented as 1000/d or k1/d or k{page_views}/d (and a metric value of 5.3 would mean "5300 page views per day"). % represents dimensionless value of 1/100, and annotates values giving a percentage (so the metric values are typically in the range of 0..100, and a metric value 3 means "3 percent"). 10^2.% indicates a metric contains a ratio, typically in the range 0..1, that will be multiplied by 100 and displayed as a percentage (so a metric value 0.03 means "3 percent").
	Unit pulumi.StringOutput `pulumi:"unit"`
	// Whether the measurement is an integer, a floating-point number, etc. Some combinations of metric_kind and value_type might not be supported.
	ValueType pulumi.StringOutput `pulumi:"valueType"`
}

// NewMetricDescriptor registers a new resource with the given unique name, arguments, and options.
func NewMetricDescriptor(ctx *pulumi.Context,
	name string, args *MetricDescriptorArgs, opts ...pulumi.ResourceOption) (*MetricDescriptor, error) {
	if args == nil {
		args = &MetricDescriptorArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource MetricDescriptor
	err := ctx.RegisterResource("google-native:monitoring/v3:MetricDescriptor", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetMetricDescriptor gets an existing MetricDescriptor resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetMetricDescriptor(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MetricDescriptorState, opts ...pulumi.ResourceOption) (*MetricDescriptor, error) {
	var resource MetricDescriptor
	err := ctx.ReadResource("google-native:monitoring/v3:MetricDescriptor", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering MetricDescriptor resources.
type metricDescriptorState struct {
}

type MetricDescriptorState struct {
}

func (MetricDescriptorState) ElementType() reflect.Type {
	return reflect.TypeOf((*metricDescriptorState)(nil)).Elem()
}

type metricDescriptorArgs struct {
	// A detailed description of the metric, which can be used in documentation.
	Description *string `pulumi:"description"`
	// A concise name for the metric, which can be displayed in user interfaces. Use sentence case without an ending period, for example "Request count". This field is optional but it is recommended to be set for any metrics associated with user-visible concepts, such as Quota.
	DisplayName *string `pulumi:"displayName"`
	// The set of labels that can be used to describe a specific instance of this metric type. For example, the appengine.googleapis.com/http/server/response_latencies metric type has a label for the HTTP response code, response_code, so you can look at latencies for successful responses or just for responses that failed.
	Labels []LabelDescriptor `pulumi:"labels"`
	// Optional. The launch stage of the metric definition.
	LaunchStage *MetricDescriptorLaunchStage `pulumi:"launchStage"`
	// Optional. Metadata which can be used to guide usage of the metric.
	Metadata *MetricDescriptorMetadata `pulumi:"metadata"`
	// Whether the metric records instantaneous values, changes to a value, etc. Some combinations of metric_kind and value_type might not be supported.
	MetricKind *MetricDescriptorMetricKind `pulumi:"metricKind"`
	// Read-only. If present, then a time series, which is identified partially by a metric type and a MonitoredResourceDescriptor, that is associated with this metric type can only be associated with one of the monitored resource types listed here.
	MonitoredResourceTypes []string `pulumi:"monitoredResourceTypes"`
	// The resource name of the metric descriptor.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The metric type, including its DNS name prefix. The type is not URL-encoded. All user-defined metric types have the DNS name custom.googleapis.com or external.googleapis.com. Metric types should use a natural hierarchical grouping. For example: "custom.googleapis.com/invoice/paid/amount" "external.googleapis.com/prometheus/up" "appengine.googleapis.com/http/server/response_latencies"
	Type *string `pulumi:"type"`
	// The units in which the metric value is reported. It is only applicable if the value_type is INT64, DOUBLE, or DISTRIBUTION. The unit defines the representation of the stored metric values.Different systems might scale the values to be more easily displayed (so a value of 0.02kBy might be displayed as 20By, and a value of 3523kBy might be displayed as 3.5MBy). However, if the unit is kBy, then the value of the metric is always in thousands of bytes, no matter how it might be displayed.If you want a custom metric to record the exact number of CPU-seconds used by a job, you can create an INT64 CUMULATIVE metric whose unit is s{CPU} (or equivalently 1s{CPU} or just s). If the job uses 12,005 CPU-seconds, then the value is written as 12005.Alternatively, if you want a custom metric to record data in a more granular way, you can create a DOUBLE CUMULATIVE metric whose unit is ks{CPU}, and then write the value 12.005 (which is 12005/1000), or use Kis{CPU} and write 11.723 (which is 12005/1024).The supported units are a subset of The Unified Code for Units of Measure (https://unitsofmeasure.org/ucum.html) standard:Basic units (UNIT) bit bit By byte s second min minute h hour d day 1 dimensionlessPrefixes (PREFIX) k kilo (10^3) M mega (10^6) G giga (10^9) T tera (10^12) P peta (10^15) E exa (10^18) Z zetta (10^21) Y yotta (10^24) m milli (10^-3) u micro (10^-6) n nano (10^-9) p pico (10^-12) f femto (10^-15) a atto (10^-18) z zepto (10^-21) y yocto (10^-24) Ki kibi (2^10) Mi mebi (2^20) Gi gibi (2^30) Ti tebi (2^40) Pi pebi (2^50)GrammarThe grammar also includes these connectors: / division or ratio (as an infix operator). For examples, kBy/{email} or MiBy/10ms (although you should almost never have /s in a metric unit; rates should always be computed at query time from the underlying cumulative or delta value). . multiplication or composition (as an infix operator). For examples, GBy.d or k{watt}.h.The grammar for a unit is as follows: Expression = Component { "." Component } { "/" Component } ; Component = ( [ PREFIX ] UNIT | "%" ) [ Annotation ] | Annotation | "1" ; Annotation = "{" NAME "}" ; Notes: Annotation is just a comment if it follows a UNIT. If the annotation is used alone, then the unit is equivalent to 1. For examples, {request}/s == 1/s, By{transmitted}/s == By/s. NAME is a sequence of non-blank printable ASCII characters not containing { or }. 1 represents a unitary dimensionless unit (https://en.wikipedia.org/wiki/Dimensionless_quantity) of 1, such as in 1/s. It is typically used when none of the basic units are appropriate. For example, "new users per day" can be represented as 1/d or {new-users}/d (and a metric value 5 would mean "5 new users). Alternatively, "thousands of page views per day" would be represented as 1000/d or k1/d or k{page_views}/d (and a metric value of 5.3 would mean "5300 page views per day"). % represents dimensionless value of 1/100, and annotates values giving a percentage (so the metric values are typically in the range of 0..100, and a metric value 3 means "3 percent"). 10^2.% indicates a metric contains a ratio, typically in the range 0..1, that will be multiplied by 100 and displayed as a percentage (so a metric value 0.03 means "3 percent").
	Unit *string `pulumi:"unit"`
	// Whether the measurement is an integer, a floating-point number, etc. Some combinations of metric_kind and value_type might not be supported.
	ValueType *MetricDescriptorValueType `pulumi:"valueType"`
}

// The set of arguments for constructing a MetricDescriptor resource.
type MetricDescriptorArgs struct {
	// A detailed description of the metric, which can be used in documentation.
	Description pulumi.StringPtrInput
	// A concise name for the metric, which can be displayed in user interfaces. Use sentence case without an ending period, for example "Request count". This field is optional but it is recommended to be set for any metrics associated with user-visible concepts, such as Quota.
	DisplayName pulumi.StringPtrInput
	// The set of labels that can be used to describe a specific instance of this metric type. For example, the appengine.googleapis.com/http/server/response_latencies metric type has a label for the HTTP response code, response_code, so you can look at latencies for successful responses or just for responses that failed.
	Labels LabelDescriptorArrayInput
	// Optional. The launch stage of the metric definition.
	LaunchStage MetricDescriptorLaunchStagePtrInput
	// Optional. Metadata which can be used to guide usage of the metric.
	Metadata MetricDescriptorMetadataPtrInput
	// Whether the metric records instantaneous values, changes to a value, etc. Some combinations of metric_kind and value_type might not be supported.
	MetricKind MetricDescriptorMetricKindPtrInput
	// Read-only. If present, then a time series, which is identified partially by a metric type and a MonitoredResourceDescriptor, that is associated with this metric type can only be associated with one of the monitored resource types listed here.
	MonitoredResourceTypes pulumi.StringArrayInput
	// The resource name of the metric descriptor.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The metric type, including its DNS name prefix. The type is not URL-encoded. All user-defined metric types have the DNS name custom.googleapis.com or external.googleapis.com. Metric types should use a natural hierarchical grouping. For example: "custom.googleapis.com/invoice/paid/amount" "external.googleapis.com/prometheus/up" "appengine.googleapis.com/http/server/response_latencies"
	Type pulumi.StringPtrInput
	// The units in which the metric value is reported. It is only applicable if the value_type is INT64, DOUBLE, or DISTRIBUTION. The unit defines the representation of the stored metric values.Different systems might scale the values to be more easily displayed (so a value of 0.02kBy might be displayed as 20By, and a value of 3523kBy might be displayed as 3.5MBy). However, if the unit is kBy, then the value of the metric is always in thousands of bytes, no matter how it might be displayed.If you want a custom metric to record the exact number of CPU-seconds used by a job, you can create an INT64 CUMULATIVE metric whose unit is s{CPU} (or equivalently 1s{CPU} or just s). If the job uses 12,005 CPU-seconds, then the value is written as 12005.Alternatively, if you want a custom metric to record data in a more granular way, you can create a DOUBLE CUMULATIVE metric whose unit is ks{CPU}, and then write the value 12.005 (which is 12005/1000), or use Kis{CPU} and write 11.723 (which is 12005/1024).The supported units are a subset of The Unified Code for Units of Measure (https://unitsofmeasure.org/ucum.html) standard:Basic units (UNIT) bit bit By byte s second min minute h hour d day 1 dimensionlessPrefixes (PREFIX) k kilo (10^3) M mega (10^6) G giga (10^9) T tera (10^12) P peta (10^15) E exa (10^18) Z zetta (10^21) Y yotta (10^24) m milli (10^-3) u micro (10^-6) n nano (10^-9) p pico (10^-12) f femto (10^-15) a atto (10^-18) z zepto (10^-21) y yocto (10^-24) Ki kibi (2^10) Mi mebi (2^20) Gi gibi (2^30) Ti tebi (2^40) Pi pebi (2^50)GrammarThe grammar also includes these connectors: / division or ratio (as an infix operator). For examples, kBy/{email} or MiBy/10ms (although you should almost never have /s in a metric unit; rates should always be computed at query time from the underlying cumulative or delta value). . multiplication or composition (as an infix operator). For examples, GBy.d or k{watt}.h.The grammar for a unit is as follows: Expression = Component { "." Component } { "/" Component } ; Component = ( [ PREFIX ] UNIT | "%" ) [ Annotation ] | Annotation | "1" ; Annotation = "{" NAME "}" ; Notes: Annotation is just a comment if it follows a UNIT. If the annotation is used alone, then the unit is equivalent to 1. For examples, {request}/s == 1/s, By{transmitted}/s == By/s. NAME is a sequence of non-blank printable ASCII characters not containing { or }. 1 represents a unitary dimensionless unit (https://en.wikipedia.org/wiki/Dimensionless_quantity) of 1, such as in 1/s. It is typically used when none of the basic units are appropriate. For example, "new users per day" can be represented as 1/d or {new-users}/d (and a metric value 5 would mean "5 new users). Alternatively, "thousands of page views per day" would be represented as 1000/d or k1/d or k{page_views}/d (and a metric value of 5.3 would mean "5300 page views per day"). % represents dimensionless value of 1/100, and annotates values giving a percentage (so the metric values are typically in the range of 0..100, and a metric value 3 means "3 percent"). 10^2.% indicates a metric contains a ratio, typically in the range 0..1, that will be multiplied by 100 and displayed as a percentage (so a metric value 0.03 means "3 percent").
	Unit pulumi.StringPtrInput
	// Whether the measurement is an integer, a floating-point number, etc. Some combinations of metric_kind and value_type might not be supported.
	ValueType MetricDescriptorValueTypePtrInput
}

func (MetricDescriptorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*metricDescriptorArgs)(nil)).Elem()
}

type MetricDescriptorInput interface {
	pulumi.Input

	ToMetricDescriptorOutput() MetricDescriptorOutput
	ToMetricDescriptorOutputWithContext(ctx context.Context) MetricDescriptorOutput
}

func (*MetricDescriptor) ElementType() reflect.Type {
	return reflect.TypeOf((**MetricDescriptor)(nil)).Elem()
}

func (i *MetricDescriptor) ToMetricDescriptorOutput() MetricDescriptorOutput {
	return i.ToMetricDescriptorOutputWithContext(context.Background())
}

func (i *MetricDescriptor) ToMetricDescriptorOutputWithContext(ctx context.Context) MetricDescriptorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricDescriptorOutput)
}

type MetricDescriptorOutput struct{ *pulumi.OutputState }

func (MetricDescriptorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetricDescriptor)(nil)).Elem()
}

func (o MetricDescriptorOutput) ToMetricDescriptorOutput() MetricDescriptorOutput {
	return o
}

func (o MetricDescriptorOutput) ToMetricDescriptorOutputWithContext(ctx context.Context) MetricDescriptorOutput {
	return o
}

// A detailed description of the metric, which can be used in documentation.
func (o MetricDescriptorOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricDescriptor) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// A concise name for the metric, which can be displayed in user interfaces. Use sentence case without an ending period, for example "Request count". This field is optional but it is recommended to be set for any metrics associated with user-visible concepts, such as Quota.
func (o MetricDescriptorOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricDescriptor) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// The set of labels that can be used to describe a specific instance of this metric type. For example, the appengine.googleapis.com/http/server/response_latencies metric type has a label for the HTTP response code, response_code, so you can look at latencies for successful responses or just for responses that failed.
func (o MetricDescriptorOutput) Labels() LabelDescriptorResponseArrayOutput {
	return o.ApplyT(func(v *MetricDescriptor) LabelDescriptorResponseArrayOutput { return v.Labels }).(LabelDescriptorResponseArrayOutput)
}

// Optional. The launch stage of the metric definition.
func (o MetricDescriptorOutput) LaunchStage() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricDescriptor) pulumi.StringOutput { return v.LaunchStage }).(pulumi.StringOutput)
}

// Optional. Metadata which can be used to guide usage of the metric.
func (o MetricDescriptorOutput) Metadata() MetricDescriptorMetadataResponseOutput {
	return o.ApplyT(func(v *MetricDescriptor) MetricDescriptorMetadataResponseOutput { return v.Metadata }).(MetricDescriptorMetadataResponseOutput)
}

// Whether the metric records instantaneous values, changes to a value, etc. Some combinations of metric_kind and value_type might not be supported.
func (o MetricDescriptorOutput) MetricKind() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricDescriptor) pulumi.StringOutput { return v.MetricKind }).(pulumi.StringOutput)
}

// Read-only. If present, then a time series, which is identified partially by a metric type and a MonitoredResourceDescriptor, that is associated with this metric type can only be associated with one of the monitored resource types listed here.
func (o MetricDescriptorOutput) MonitoredResourceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MetricDescriptor) pulumi.StringArrayOutput { return v.MonitoredResourceTypes }).(pulumi.StringArrayOutput)
}

// The resource name of the metric descriptor.
func (o MetricDescriptorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricDescriptor) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MetricDescriptorOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricDescriptor) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The metric type, including its DNS name prefix. The type is not URL-encoded. All user-defined metric types have the DNS name custom.googleapis.com or external.googleapis.com. Metric types should use a natural hierarchical grouping. For example: "custom.googleapis.com/invoice/paid/amount" "external.googleapis.com/prometheus/up" "appengine.googleapis.com/http/server/response_latencies"
func (o MetricDescriptorOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricDescriptor) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// The units in which the metric value is reported. It is only applicable if the value_type is INT64, DOUBLE, or DISTRIBUTION. The unit defines the representation of the stored metric values.Different systems might scale the values to be more easily displayed (so a value of 0.02kBy might be displayed as 20By, and a value of 3523kBy might be displayed as 3.5MBy). However, if the unit is kBy, then the value of the metric is always in thousands of bytes, no matter how it might be displayed.If you want a custom metric to record the exact number of CPU-seconds used by a job, you can create an INT64 CUMULATIVE metric whose unit is s{CPU} (or equivalently 1s{CPU} or just s). If the job uses 12,005 CPU-seconds, then the value is written as 12005.Alternatively, if you want a custom metric to record data in a more granular way, you can create a DOUBLE CUMULATIVE metric whose unit is ks{CPU}, and then write the value 12.005 (which is 12005/1000), or use Kis{CPU} and write 11.723 (which is 12005/1024).The supported units are a subset of The Unified Code for Units of Measure (https://unitsofmeasure.org/ucum.html) standard:Basic units (UNIT) bit bit By byte s second min minute h hour d day 1 dimensionlessPrefixes (PREFIX) k kilo (10^3) M mega (10^6) G giga (10^9) T tera (10^12) P peta (10^15) E exa (10^18) Z zetta (10^21) Y yotta (10^24) m milli (10^-3) u micro (10^-6) n nano (10^-9) p pico (10^-12) f femto (10^-15) a atto (10^-18) z zepto (10^-21) y yocto (10^-24) Ki kibi (2^10) Mi mebi (2^20) Gi gibi (2^30) Ti tebi (2^40) Pi pebi (2^50)GrammarThe grammar also includes these connectors: / division or ratio (as an infix operator). For examples, kBy/{email} or MiBy/10ms (although you should almost never have /s in a metric unit; rates should always be computed at query time from the underlying cumulative or delta value). . multiplication or composition (as an infix operator). For examples, GBy.d or k{watt}.h.The grammar for a unit is as follows: Expression = Component { "." Component } { "/" Component } ; Component = ( [ PREFIX ] UNIT | "%" ) [ Annotation ] | Annotation | "1" ; Annotation = "{" NAME "}" ; Notes: Annotation is just a comment if it follows a UNIT. If the annotation is used alone, then the unit is equivalent to 1. For examples, {request}/s == 1/s, By{transmitted}/s == By/s. NAME is a sequence of non-blank printable ASCII characters not containing { or }. 1 represents a unitary dimensionless unit (https://en.wikipedia.org/wiki/Dimensionless_quantity) of 1, such as in 1/s. It is typically used when none of the basic units are appropriate. For example, "new users per day" can be represented as 1/d or {new-users}/d (and a metric value 5 would mean "5 new users). Alternatively, "thousands of page views per day" would be represented as 1000/d or k1/d or k{page_views}/d (and a metric value of 5.3 would mean "5300 page views per day"). % represents dimensionless value of 1/100, and annotates values giving a percentage (so the metric values are typically in the range of 0..100, and a metric value 3 means "3 percent"). 10^2.% indicates a metric contains a ratio, typically in the range 0..1, that will be multiplied by 100 and displayed as a percentage (so a metric value 0.03 means "3 percent").
func (o MetricDescriptorOutput) Unit() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricDescriptor) pulumi.StringOutput { return v.Unit }).(pulumi.StringOutput)
}

// Whether the measurement is an integer, a floating-point number, etc. Some combinations of metric_kind and value_type might not be supported.
func (o MetricDescriptorOutput) ValueType() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricDescriptor) pulumi.StringOutput { return v.ValueType }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*MetricDescriptorInput)(nil)).Elem(), &MetricDescriptor{})
	pulumi.RegisterOutputType(MetricDescriptorOutput{})
}
