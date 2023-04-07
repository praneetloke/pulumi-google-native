// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2alpha

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new model.
type Model struct {
	pulumi.CustomResourceState

	CatalogId pulumi.StringOutput `pulumi:"catalogId"`
	// Timestamp the Recommendation Model was created at.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The state of data requirements for this model: `DATA_OK` and `DATA_ERROR`. Recommendation model cannot be trained if the data is in `DATA_ERROR` state. Recommendation model can have `DATA_ERROR` state even if serving state is `ACTIVE`: models were trained successfully before, but cannot be refreshed because model no longer has sufficient data for training.
	DataState pulumi.StringOutput `pulumi:"dataState"`
	// The display name of the model. Should be human readable, used to display Recommendation Models in the Retail Cloud Console Dashboard. UTF-8 encoded string with limit of 1024 characters.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// Optional. Whether to run a dry run to validate the request (without actually creating the model).
	DryRun pulumi.BoolPtrOutput `pulumi:"dryRun"`
	// Optional. If `RECOMMENDATIONS_FILTERING_ENABLED`, recommendation filtering by attributes is enabled for the model.
	FilteringOption pulumi.StringOutput `pulumi:"filteringOption"`
	// The timestamp when the latest successful tune finished.
	LastTuneTime pulumi.StringOutput `pulumi:"lastTuneTime"`
	Location     pulumi.StringOutput `pulumi:"location"`
	// The fully qualified resource name of the model. Format: `projects/{project_number}/locations/{location_id}/catalogs/{catalog_id}/models/{model_id}` catalog_id has char limit of 50. recommendation_model_id has char limit of 40.
	Name pulumi.StringOutput `pulumi:"name"`
	// Optional. The optimization objective e.g. `cvr`. Currently supported values: `ctr`, `cvr`, `revenue-per-order`. If not specified, we choose default based on model type. Default depends on type of recommendation: `recommended-for-you` => `ctr` `others-you-may-like` => `ctr` `frequently-bought-together` => `revenue_per_order` This field together with optimization_objective describe model metadata to use to control model training and serving. See https://cloud.google.com/retail/docs/models for more details on what the model metadata control and which combination of parameters are valid. For invalid combinations of parameters (e.g. type = `frequently-bought-together` and optimization_objective = `ctr`), you receive an error 400 if you try to create/update a recommendation with this set of knobs.
	OptimizationObjective pulumi.StringOutput `pulumi:"optimizationObjective"`
	// Optional. The page optimization config.
	PageOptimizationConfig GoogleCloudRetailV2alphaModelPageOptimizationConfigResponseOutput `pulumi:"pageOptimizationConfig"`
	// Optional. The state of periodic tuning. The period we use is 3 months - to do a one-off tune earlier use the `TuneModel` method. Default value is `PERIODIC_TUNING_ENABLED`.
	PeriodicTuningState pulumi.StringOutput `pulumi:"periodicTuningState"`
	Project             pulumi.StringOutput `pulumi:"project"`
	// The list of valid serving configs associated with the PageOptimizationConfig.
	ServingConfigLists GoogleCloudRetailV2alphaModelServingConfigListResponseArrayOutput `pulumi:"servingConfigLists"`
	// The serving state of the model: `ACTIVE`, `NOT_ACTIVE`.
	ServingState pulumi.StringOutput `pulumi:"servingState"`
	// Optional. The training state that the model is in (e.g. `TRAINING` or `PAUSED`). Since part of the cost of running the service is frequency of training - this can be used to determine when to train model in order to control cost. If not specified: the default value for `CreateModel` method is `TRAINING`. The default value for `UpdateModel` method is to keep the state the same as before.
	TrainingState pulumi.StringOutput `pulumi:"trainingState"`
	// The tune operation associated with the model. Can be used to determine if there is an ongoing tune for this recommendation. Empty field implies no tune is goig on.
	TuningOperation pulumi.StringOutput `pulumi:"tuningOperation"`
	// The type of model e.g. `home-page`. Currently supported values: `recommended-for-you`, `others-you-may-like`, `frequently-bought-together`, `page-optimization`, `similar-items`, `buy-it-again`, `on-sale-items`, and `recently-viewed`(readonly value). This field together with optimization_objective describe model metadata to use to control model training and serving. See https://cloud.google.com/retail/docs/models for more details on what the model metadata control and which combination of parameters are valid. For invalid combinations of parameters (e.g. type = `frequently-bought-together` and optimization_objective = `ctr`), you receive an error 400 if you try to create/update a recommendation with this set of knobs.
	Type pulumi.StringOutput `pulumi:"type"`
	// Timestamp the Recommendation Model was last updated. E.g. if a Recommendation Model was paused - this would be the time the pause was initiated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewModel registers a new resource with the given unique name, arguments, and options.
func NewModel(ctx *pulumi.Context,
	name string, args *ModelArgs, opts ...pulumi.ResourceOption) (*Model, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CatalogId == nil {
		return nil, errors.New("invalid value for required argument 'CatalogId'")
	}
	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"catalogId",
		"location",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	var resource Model
	err := ctx.RegisterResource("google-native:retail/v2alpha:Model", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetModel gets an existing Model resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetModel(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ModelState, opts ...pulumi.ResourceOption) (*Model, error) {
	var resource Model
	err := ctx.ReadResource("google-native:retail/v2alpha:Model", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Model resources.
type modelState struct {
}

type ModelState struct {
}

func (ModelState) ElementType() reflect.Type {
	return reflect.TypeOf((*modelState)(nil)).Elem()
}

type modelArgs struct {
	CatalogId string `pulumi:"catalogId"`
	// The display name of the model. Should be human readable, used to display Recommendation Models in the Retail Cloud Console Dashboard. UTF-8 encoded string with limit of 1024 characters.
	DisplayName string `pulumi:"displayName"`
	// Optional. Whether to run a dry run to validate the request (without actually creating the model).
	DryRun *bool `pulumi:"dryRun"`
	// Optional. If `RECOMMENDATIONS_FILTERING_ENABLED`, recommendation filtering by attributes is enabled for the model.
	FilteringOption *ModelFilteringOption `pulumi:"filteringOption"`
	Location        *string               `pulumi:"location"`
	// The fully qualified resource name of the model. Format: `projects/{project_number}/locations/{location_id}/catalogs/{catalog_id}/models/{model_id}` catalog_id has char limit of 50. recommendation_model_id has char limit of 40.
	Name *string `pulumi:"name"`
	// Optional. The optimization objective e.g. `cvr`. Currently supported values: `ctr`, `cvr`, `revenue-per-order`. If not specified, we choose default based on model type. Default depends on type of recommendation: `recommended-for-you` => `ctr` `others-you-may-like` => `ctr` `frequently-bought-together` => `revenue_per_order` This field together with optimization_objective describe model metadata to use to control model training and serving. See https://cloud.google.com/retail/docs/models for more details on what the model metadata control and which combination of parameters are valid. For invalid combinations of parameters (e.g. type = `frequently-bought-together` and optimization_objective = `ctr`), you receive an error 400 if you try to create/update a recommendation with this set of knobs.
	OptimizationObjective *string `pulumi:"optimizationObjective"`
	// Optional. The page optimization config.
	PageOptimizationConfig *GoogleCloudRetailV2alphaModelPageOptimizationConfig `pulumi:"pageOptimizationConfig"`
	// Optional. The state of periodic tuning. The period we use is 3 months - to do a one-off tune earlier use the `TuneModel` method. Default value is `PERIODIC_TUNING_ENABLED`.
	PeriodicTuningState *ModelPeriodicTuningState `pulumi:"periodicTuningState"`
	Project             *string                   `pulumi:"project"`
	// Optional. The training state that the model is in (e.g. `TRAINING` or `PAUSED`). Since part of the cost of running the service is frequency of training - this can be used to determine when to train model in order to control cost. If not specified: the default value for `CreateModel` method is `TRAINING`. The default value for `UpdateModel` method is to keep the state the same as before.
	TrainingState *ModelTrainingState `pulumi:"trainingState"`
	// The type of model e.g. `home-page`. Currently supported values: `recommended-for-you`, `others-you-may-like`, `frequently-bought-together`, `page-optimization`, `similar-items`, `buy-it-again`, `on-sale-items`, and `recently-viewed`(readonly value). This field together with optimization_objective describe model metadata to use to control model training and serving. See https://cloud.google.com/retail/docs/models for more details on what the model metadata control and which combination of parameters are valid. For invalid combinations of parameters (e.g. type = `frequently-bought-together` and optimization_objective = `ctr`), you receive an error 400 if you try to create/update a recommendation with this set of knobs.
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a Model resource.
type ModelArgs struct {
	CatalogId pulumi.StringInput
	// The display name of the model. Should be human readable, used to display Recommendation Models in the Retail Cloud Console Dashboard. UTF-8 encoded string with limit of 1024 characters.
	DisplayName pulumi.StringInput
	// Optional. Whether to run a dry run to validate the request (without actually creating the model).
	DryRun pulumi.BoolPtrInput
	// Optional. If `RECOMMENDATIONS_FILTERING_ENABLED`, recommendation filtering by attributes is enabled for the model.
	FilteringOption ModelFilteringOptionPtrInput
	Location        pulumi.StringPtrInput
	// The fully qualified resource name of the model. Format: `projects/{project_number}/locations/{location_id}/catalogs/{catalog_id}/models/{model_id}` catalog_id has char limit of 50. recommendation_model_id has char limit of 40.
	Name pulumi.StringPtrInput
	// Optional. The optimization objective e.g. `cvr`. Currently supported values: `ctr`, `cvr`, `revenue-per-order`. If not specified, we choose default based on model type. Default depends on type of recommendation: `recommended-for-you` => `ctr` `others-you-may-like` => `ctr` `frequently-bought-together` => `revenue_per_order` This field together with optimization_objective describe model metadata to use to control model training and serving. See https://cloud.google.com/retail/docs/models for more details on what the model metadata control and which combination of parameters are valid. For invalid combinations of parameters (e.g. type = `frequently-bought-together` and optimization_objective = `ctr`), you receive an error 400 if you try to create/update a recommendation with this set of knobs.
	OptimizationObjective pulumi.StringPtrInput
	// Optional. The page optimization config.
	PageOptimizationConfig GoogleCloudRetailV2alphaModelPageOptimizationConfigPtrInput
	// Optional. The state of periodic tuning. The period we use is 3 months - to do a one-off tune earlier use the `TuneModel` method. Default value is `PERIODIC_TUNING_ENABLED`.
	PeriodicTuningState ModelPeriodicTuningStatePtrInput
	Project             pulumi.StringPtrInput
	// Optional. The training state that the model is in (e.g. `TRAINING` or `PAUSED`). Since part of the cost of running the service is frequency of training - this can be used to determine when to train model in order to control cost. If not specified: the default value for `CreateModel` method is `TRAINING`. The default value for `UpdateModel` method is to keep the state the same as before.
	TrainingState ModelTrainingStatePtrInput
	// The type of model e.g. `home-page`. Currently supported values: `recommended-for-you`, `others-you-may-like`, `frequently-bought-together`, `page-optimization`, `similar-items`, `buy-it-again`, `on-sale-items`, and `recently-viewed`(readonly value). This field together with optimization_objective describe model metadata to use to control model training and serving. See https://cloud.google.com/retail/docs/models for more details on what the model metadata control and which combination of parameters are valid. For invalid combinations of parameters (e.g. type = `frequently-bought-together` and optimization_objective = `ctr`), you receive an error 400 if you try to create/update a recommendation with this set of knobs.
	Type pulumi.StringInput
}

func (ModelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*modelArgs)(nil)).Elem()
}

type ModelInput interface {
	pulumi.Input

	ToModelOutput() ModelOutput
	ToModelOutputWithContext(ctx context.Context) ModelOutput
}

func (*Model) ElementType() reflect.Type {
	return reflect.TypeOf((**Model)(nil)).Elem()
}

func (i *Model) ToModelOutput() ModelOutput {
	return i.ToModelOutputWithContext(context.Background())
}

func (i *Model) ToModelOutputWithContext(ctx context.Context) ModelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModelOutput)
}

type ModelOutput struct{ *pulumi.OutputState }

func (ModelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Model)(nil)).Elem()
}

func (o ModelOutput) ToModelOutput() ModelOutput {
	return o
}

func (o ModelOutput) ToModelOutputWithContext(ctx context.Context) ModelOutput {
	return o
}

func (o ModelOutput) CatalogId() pulumi.StringOutput {
	return o.ApplyT(func(v *Model) pulumi.StringOutput { return v.CatalogId }).(pulumi.StringOutput)
}

// Timestamp the Recommendation Model was created at.
func (o ModelOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Model) pulumi.StringOutput { return v.CreateTime }).(pulumi.StringOutput)
}

// The state of data requirements for this model: `DATA_OK` and `DATA_ERROR`. Recommendation model cannot be trained if the data is in `DATA_ERROR` state. Recommendation model can have `DATA_ERROR` state even if serving state is `ACTIVE`: models were trained successfully before, but cannot be refreshed because model no longer has sufficient data for training.
func (o ModelOutput) DataState() pulumi.StringOutput {
	return o.ApplyT(func(v *Model) pulumi.StringOutput { return v.DataState }).(pulumi.StringOutput)
}

// The display name of the model. Should be human readable, used to display Recommendation Models in the Retail Cloud Console Dashboard. UTF-8 encoded string with limit of 1024 characters.
func (o ModelOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *Model) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

// Optional. Whether to run a dry run to validate the request (without actually creating the model).
func (o ModelOutput) DryRun() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Model) pulumi.BoolPtrOutput { return v.DryRun }).(pulumi.BoolPtrOutput)
}

// Optional. If `RECOMMENDATIONS_FILTERING_ENABLED`, recommendation filtering by attributes is enabled for the model.
func (o ModelOutput) FilteringOption() pulumi.StringOutput {
	return o.ApplyT(func(v *Model) pulumi.StringOutput { return v.FilteringOption }).(pulumi.StringOutput)
}

// The timestamp when the latest successful tune finished.
func (o ModelOutput) LastTuneTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Model) pulumi.StringOutput { return v.LastTuneTime }).(pulumi.StringOutput)
}

func (o ModelOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Model) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

// The fully qualified resource name of the model. Format: `projects/{project_number}/locations/{location_id}/catalogs/{catalog_id}/models/{model_id}` catalog_id has char limit of 50. recommendation_model_id has char limit of 40.
func (o ModelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Model) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// Optional. The optimization objective e.g. `cvr`. Currently supported values: `ctr`, `cvr`, `revenue-per-order`. If not specified, we choose default based on model type. Default depends on type of recommendation: `recommended-for-you` => `ctr` `others-you-may-like` => `ctr` `frequently-bought-together` => `revenue_per_order` This field together with optimization_objective describe model metadata to use to control model training and serving. See https://cloud.google.com/retail/docs/models for more details on what the model metadata control and which combination of parameters are valid. For invalid combinations of parameters (e.g. type = `frequently-bought-together` and optimization_objective = `ctr`), you receive an error 400 if you try to create/update a recommendation with this set of knobs.
func (o ModelOutput) OptimizationObjective() pulumi.StringOutput {
	return o.ApplyT(func(v *Model) pulumi.StringOutput { return v.OptimizationObjective }).(pulumi.StringOutput)
}

// Optional. The page optimization config.
func (o ModelOutput) PageOptimizationConfig() GoogleCloudRetailV2alphaModelPageOptimizationConfigResponseOutput {
	return o.ApplyT(func(v *Model) GoogleCloudRetailV2alphaModelPageOptimizationConfigResponseOutput {
		return v.PageOptimizationConfig
	}).(GoogleCloudRetailV2alphaModelPageOptimizationConfigResponseOutput)
}

// Optional. The state of periodic tuning. The period we use is 3 months - to do a one-off tune earlier use the `TuneModel` method. Default value is `PERIODIC_TUNING_ENABLED`.
func (o ModelOutput) PeriodicTuningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Model) pulumi.StringOutput { return v.PeriodicTuningState }).(pulumi.StringOutput)
}

func (o ModelOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Model) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// The list of valid serving configs associated with the PageOptimizationConfig.
func (o ModelOutput) ServingConfigLists() GoogleCloudRetailV2alphaModelServingConfigListResponseArrayOutput {
	return o.ApplyT(func(v *Model) GoogleCloudRetailV2alphaModelServingConfigListResponseArrayOutput {
		return v.ServingConfigLists
	}).(GoogleCloudRetailV2alphaModelServingConfigListResponseArrayOutput)
}

// The serving state of the model: `ACTIVE`, `NOT_ACTIVE`.
func (o ModelOutput) ServingState() pulumi.StringOutput {
	return o.ApplyT(func(v *Model) pulumi.StringOutput { return v.ServingState }).(pulumi.StringOutput)
}

// Optional. The training state that the model is in (e.g. `TRAINING` or `PAUSED`). Since part of the cost of running the service is frequency of training - this can be used to determine when to train model in order to control cost. If not specified: the default value for `CreateModel` method is `TRAINING`. The default value for `UpdateModel` method is to keep the state the same as before.
func (o ModelOutput) TrainingState() pulumi.StringOutput {
	return o.ApplyT(func(v *Model) pulumi.StringOutput { return v.TrainingState }).(pulumi.StringOutput)
}

// The tune operation associated with the model. Can be used to determine if there is an ongoing tune for this recommendation. Empty field implies no tune is goig on.
func (o ModelOutput) TuningOperation() pulumi.StringOutput {
	return o.ApplyT(func(v *Model) pulumi.StringOutput { return v.TuningOperation }).(pulumi.StringOutput)
}

// The type of model e.g. `home-page`. Currently supported values: `recommended-for-you`, `others-you-may-like`, `frequently-bought-together`, `page-optimization`, `similar-items`, `buy-it-again`, `on-sale-items`, and `recently-viewed`(readonly value). This field together with optimization_objective describe model metadata to use to control model training and serving. See https://cloud.google.com/retail/docs/models for more details on what the model metadata control and which combination of parameters are valid. For invalid combinations of parameters (e.g. type = `frequently-bought-together` and optimization_objective = `ctr`), you receive an error 400 if you try to create/update a recommendation with this set of knobs.
func (o ModelOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Model) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

// Timestamp the Recommendation Model was last updated. E.g. if a Recommendation Model was paused - this would be the time the pause was initiated.
func (o ModelOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Model) pulumi.StringOutput { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ModelInput)(nil)).Elem(), &Model{})
	pulumi.RegisterOutputType(ModelOutput{})
}
