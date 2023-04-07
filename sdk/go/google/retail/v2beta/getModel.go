// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2beta

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a model.
func LookupModel(ctx *pulumi.Context, args *LookupModelArgs, opts ...pulumi.InvokeOption) (*LookupModelResult, error) {
	var rv LookupModelResult
	err := ctx.Invoke("google-native:retail/v2beta:getModel", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupModelArgs struct {
	CatalogId string  `pulumi:"catalogId"`
	Location  string  `pulumi:"location"`
	ModelId   string  `pulumi:"modelId"`
	Project   *string `pulumi:"project"`
}

type LookupModelResult struct {
	// Timestamp the Recommendation Model was created at.
	CreateTime string `pulumi:"createTime"`
	// The state of data requirements for this model: `DATA_OK` and `DATA_ERROR`. Recommendation model cannot be trained if the data is in `DATA_ERROR` state. Recommendation model can have `DATA_ERROR` state even if serving state is `ACTIVE`: models were trained successfully before, but cannot be refreshed because model no longer has sufficient data for training.
	DataState string `pulumi:"dataState"`
	// The display name of the model. Should be human readable, used to display Recommendation Models in the Retail Cloud Console Dashboard. UTF-8 encoded string with limit of 1024 characters.
	DisplayName string `pulumi:"displayName"`
	// Optional. If `RECOMMENDATIONS_FILTERING_ENABLED`, recommendation filtering by attributes is enabled for the model.
	FilteringOption string `pulumi:"filteringOption"`
	// The timestamp when the latest successful tune finished.
	LastTuneTime string `pulumi:"lastTuneTime"`
	// The fully qualified resource name of the model. Format: `projects/{project_number}/locations/{location_id}/catalogs/{catalog_id}/models/{model_id}` catalog_id has char limit of 50. recommendation_model_id has char limit of 40.
	Name string `pulumi:"name"`
	// Optional. The optimization objective e.g. `cvr`. Currently supported values: `ctr`, `cvr`, `revenue-per-order`. If not specified, we choose default based on model type. Default depends on type of recommendation: `recommended-for-you` => `ctr` `others-you-may-like` => `ctr` `frequently-bought-together` => `revenue_per_order` This field together with optimization_objective describe model metadata to use to control model training and serving. See https://cloud.google.com/retail/docs/models for more details on what the model metadata control and which combination of parameters are valid. For invalid combinations of parameters (e.g. type = `frequently-bought-together` and optimization_objective = `ctr`), you receive an error 400 if you try to create/update a recommendation with this set of knobs.
	OptimizationObjective string `pulumi:"optimizationObjective"`
	// Optional. The state of periodic tuning. The period we use is 3 months - to do a one-off tune earlier use the `TuneModel` method. Default value is `PERIODIC_TUNING_ENABLED`.
	PeriodicTuningState string `pulumi:"periodicTuningState"`
	// The list of valid serving configs associated with the PageOptimizationConfig.
	ServingConfigLists []GoogleCloudRetailV2betaModelServingConfigListResponse `pulumi:"servingConfigLists"`
	// The serving state of the model: `ACTIVE`, `NOT_ACTIVE`.
	ServingState string `pulumi:"servingState"`
	// Optional. The training state that the model is in (e.g. `TRAINING` or `PAUSED`). Since part of the cost of running the service is frequency of training - this can be used to determine when to train model in order to control cost. If not specified: the default value for `CreateModel` method is `TRAINING`. The default value for `UpdateModel` method is to keep the state the same as before.
	TrainingState string `pulumi:"trainingState"`
	// The tune operation associated with the model. Can be used to determine if there is an ongoing tune for this recommendation. Empty field implies no tune is goig on.
	TuningOperation string `pulumi:"tuningOperation"`
	// The type of model e.g. `home-page`. Currently supported values: `recommended-for-you`, `others-you-may-like`, `frequently-bought-together`, `page-optimization`, `similar-items`, `buy-it-again`, `on-sale-items`, and `recently-viewed`(readonly value). This field together with optimization_objective describe model metadata to use to control model training and serving. See https://cloud.google.com/retail/docs/models for more details on what the model metadata control and which combination of parameters are valid. For invalid combinations of parameters (e.g. type = `frequently-bought-together` and optimization_objective = `ctr`), you receive an error 400 if you try to create/update a recommendation with this set of knobs.
	Type string `pulumi:"type"`
	// Timestamp the Recommendation Model was last updated. E.g. if a Recommendation Model was paused - this would be the time the pause was initiated.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupModelOutput(ctx *pulumi.Context, args LookupModelOutputArgs, opts ...pulumi.InvokeOption) LookupModelResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupModelResult, error) {
			args := v.(LookupModelArgs)
			r, err := LookupModel(ctx, &args, opts...)
			var s LookupModelResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupModelResultOutput)
}

type LookupModelOutputArgs struct {
	CatalogId pulumi.StringInput    `pulumi:"catalogId"`
	Location  pulumi.StringInput    `pulumi:"location"`
	ModelId   pulumi.StringInput    `pulumi:"modelId"`
	Project   pulumi.StringPtrInput `pulumi:"project"`
}

func (LookupModelOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupModelArgs)(nil)).Elem()
}

type LookupModelResultOutput struct{ *pulumi.OutputState }

func (LookupModelResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupModelResult)(nil)).Elem()
}

func (o LookupModelResultOutput) ToLookupModelResultOutput() LookupModelResultOutput {
	return o
}

func (o LookupModelResultOutput) ToLookupModelResultOutputWithContext(ctx context.Context) LookupModelResultOutput {
	return o
}

// Timestamp the Recommendation Model was created at.
func (o LookupModelResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModelResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The state of data requirements for this model: `DATA_OK` and `DATA_ERROR`. Recommendation model cannot be trained if the data is in `DATA_ERROR` state. Recommendation model can have `DATA_ERROR` state even if serving state is `ACTIVE`: models were trained successfully before, but cannot be refreshed because model no longer has sufficient data for training.
func (o LookupModelResultOutput) DataState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModelResult) string { return v.DataState }).(pulumi.StringOutput)
}

// The display name of the model. Should be human readable, used to display Recommendation Models in the Retail Cloud Console Dashboard. UTF-8 encoded string with limit of 1024 characters.
func (o LookupModelResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModelResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

// Optional. If `RECOMMENDATIONS_FILTERING_ENABLED`, recommendation filtering by attributes is enabled for the model.
func (o LookupModelResultOutput) FilteringOption() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModelResult) string { return v.FilteringOption }).(pulumi.StringOutput)
}

// The timestamp when the latest successful tune finished.
func (o LookupModelResultOutput) LastTuneTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModelResult) string { return v.LastTuneTime }).(pulumi.StringOutput)
}

// The fully qualified resource name of the model. Format: `projects/{project_number}/locations/{location_id}/catalogs/{catalog_id}/models/{model_id}` catalog_id has char limit of 50. recommendation_model_id has char limit of 40.
func (o LookupModelResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModelResult) string { return v.Name }).(pulumi.StringOutput)
}

// Optional. The optimization objective e.g. `cvr`. Currently supported values: `ctr`, `cvr`, `revenue-per-order`. If not specified, we choose default based on model type. Default depends on type of recommendation: `recommended-for-you` => `ctr` `others-you-may-like` => `ctr` `frequently-bought-together` => `revenue_per_order` This field together with optimization_objective describe model metadata to use to control model training and serving. See https://cloud.google.com/retail/docs/models for more details on what the model metadata control and which combination of parameters are valid. For invalid combinations of parameters (e.g. type = `frequently-bought-together` and optimization_objective = `ctr`), you receive an error 400 if you try to create/update a recommendation with this set of knobs.
func (o LookupModelResultOutput) OptimizationObjective() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModelResult) string { return v.OptimizationObjective }).(pulumi.StringOutput)
}

// Optional. The state of periodic tuning. The period we use is 3 months - to do a one-off tune earlier use the `TuneModel` method. Default value is `PERIODIC_TUNING_ENABLED`.
func (o LookupModelResultOutput) PeriodicTuningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModelResult) string { return v.PeriodicTuningState }).(pulumi.StringOutput)
}

// The list of valid serving configs associated with the PageOptimizationConfig.
func (o LookupModelResultOutput) ServingConfigLists() GoogleCloudRetailV2betaModelServingConfigListResponseArrayOutput {
	return o.ApplyT(func(v LookupModelResult) []GoogleCloudRetailV2betaModelServingConfigListResponse {
		return v.ServingConfigLists
	}).(GoogleCloudRetailV2betaModelServingConfigListResponseArrayOutput)
}

// The serving state of the model: `ACTIVE`, `NOT_ACTIVE`.
func (o LookupModelResultOutput) ServingState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModelResult) string { return v.ServingState }).(pulumi.StringOutput)
}

// Optional. The training state that the model is in (e.g. `TRAINING` or `PAUSED`). Since part of the cost of running the service is frequency of training - this can be used to determine when to train model in order to control cost. If not specified: the default value for `CreateModel` method is `TRAINING`. The default value for `UpdateModel` method is to keep the state the same as before.
func (o LookupModelResultOutput) TrainingState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModelResult) string { return v.TrainingState }).(pulumi.StringOutput)
}

// The tune operation associated with the model. Can be used to determine if there is an ongoing tune for this recommendation. Empty field implies no tune is goig on.
func (o LookupModelResultOutput) TuningOperation() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModelResult) string { return v.TuningOperation }).(pulumi.StringOutput)
}

// The type of model e.g. `home-page`. Currently supported values: `recommended-for-you`, `others-you-may-like`, `frequently-bought-together`, `page-optimization`, `similar-items`, `buy-it-again`, `on-sale-items`, and `recently-viewed`(readonly value). This field together with optimization_objective describe model metadata to use to control model training and serving. See https://cloud.google.com/retail/docs/models for more details on what the model metadata control and which combination of parameters are valid. For invalid combinations of parameters (e.g. type = `frequently-bought-together` and optimization_objective = `ctr`), you receive an error 400 if you try to create/update a recommendation with this set of knobs.
func (o LookupModelResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModelResult) string { return v.Type }).(pulumi.StringOutput)
}

// Timestamp the Recommendation Model was last updated. E.g. if a Recommendation Model was paused - this would be the time the pause was initiated.
func (o LookupModelResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupModelResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupModelResultOutput{})
}
