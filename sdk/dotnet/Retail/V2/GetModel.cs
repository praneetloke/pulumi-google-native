// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.GoogleNative.Retail.V2
{
    public static class GetModel
    {
        /// <summary>
        /// Gets a model.
        /// </summary>
        public static Task<GetModelResult> InvokeAsync(GetModelArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetModelResult>("google-native:retail/v2:getModel", args ?? new GetModelArgs(), options.WithDefaults());

        /// <summary>
        /// Gets a model.
        /// </summary>
        public static Output<GetModelResult> Invoke(GetModelInvokeArgs args, InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetModelResult>("google-native:retail/v2:getModel", args ?? new GetModelInvokeArgs(), options.WithDefaults());
    }


    public sealed class GetModelArgs : global::Pulumi.InvokeArgs
    {
        [Input("catalogId", required: true)]
        public string CatalogId { get; set; } = null!;

        [Input("location", required: true)]
        public string Location { get; set; } = null!;

        [Input("modelId", required: true)]
        public string ModelId { get; set; } = null!;

        [Input("project")]
        public string? Project { get; set; }

        public GetModelArgs()
        {
        }
        public static new GetModelArgs Empty => new GetModelArgs();
    }

    public sealed class GetModelInvokeArgs : global::Pulumi.InvokeArgs
    {
        [Input("catalogId", required: true)]
        public Input<string> CatalogId { get; set; } = null!;

        [Input("location", required: true)]
        public Input<string> Location { get; set; } = null!;

        [Input("modelId", required: true)]
        public Input<string> ModelId { get; set; } = null!;

        [Input("project")]
        public Input<string>? Project { get; set; }

        public GetModelInvokeArgs()
        {
        }
        public static new GetModelInvokeArgs Empty => new GetModelInvokeArgs();
    }


    [OutputType]
    public sealed class GetModelResult
    {
        /// <summary>
        /// Timestamp the Recommendation Model was created at.
        /// </summary>
        public readonly string CreateTime;
        /// <summary>
        /// The state of data requirements for this model: `DATA_OK` and `DATA_ERROR`. Recommendation model cannot be trained if the data is in `DATA_ERROR` state. Recommendation model can have `DATA_ERROR` state even if serving state is `ACTIVE`: models were trained successfully before, but cannot be refreshed because model no longer has sufficient data for training.
        /// </summary>
        public readonly string DataState;
        /// <summary>
        /// The display name of the model. Should be human readable, used to display Recommendation Models in the Retail Cloud Console Dashboard. UTF-8 encoded string with limit of 1024 characters.
        /// </summary>
        public readonly string DisplayName;
        /// <summary>
        /// Optional. If `RECOMMENDATIONS_FILTERING_ENABLED`, recommendation filtering by attributes is enabled for the model.
        /// </summary>
        public readonly string FilteringOption;
        /// <summary>
        /// The timestamp when the latest successful tune finished.
        /// </summary>
        public readonly string LastTuneTime;
        /// <summary>
        /// The fully qualified resource name of the model. Format: `projects/{project_number}/locations/{location_id}/catalogs/{catalog_id}/models/{model_id}` catalog_id has char limit of 50. recommendation_model_id has char limit of 40.
        /// </summary>
        public readonly string Name;
        /// <summary>
        /// Optional. The optimization objective e.g. `cvr`. Currently supported values: `ctr`, `cvr`, `revenue-per-order`. If not specified, we choose default based on model type. Default depends on type of recommendation: `recommended-for-you` =&gt; `ctr` `others-you-may-like` =&gt; `ctr` `frequently-bought-together` =&gt; `revenue_per_order` This field together with optimization_objective describe model metadata to use to control model training and serving. See https://cloud.google.com/retail/docs/models for more details on what the model metadata control and which combination of parameters are valid. For invalid combinations of parameters (e.g. type = `frequently-bought-together` and optimization_objective = `ctr`), you receive an error 400 if you try to create/update a recommendation with this set of knobs.
        /// </summary>
        public readonly string OptimizationObjective;
        /// <summary>
        /// Optional. The state of periodic tuning. The period we use is 3 months - to do a one-off tune earlier use the `TuneModel` method. Default value is `PERIODIC_TUNING_ENABLED`.
        /// </summary>
        public readonly string PeriodicTuningState;
        /// <summary>
        /// The list of valid serving configs associated with the PageOptimizationConfig.
        /// </summary>
        public readonly ImmutableArray<Outputs.GoogleCloudRetailV2ModelServingConfigListResponse> ServingConfigLists;
        /// <summary>
        /// The serving state of the model: `ACTIVE`, `NOT_ACTIVE`.
        /// </summary>
        public readonly string ServingState;
        /// <summary>
        /// Optional. The training state that the model is in (e.g. `TRAINING` or `PAUSED`). Since part of the cost of running the service is frequency of training - this can be used to determine when to train model in order to control cost. If not specified: the default value for `CreateModel` method is `TRAINING`. The default value for `UpdateModel` method is to keep the state the same as before.
        /// </summary>
        public readonly string TrainingState;
        /// <summary>
        /// The tune operation associated with the model. Can be used to determine if there is an ongoing tune for this recommendation. Empty field implies no tune is goig on.
        /// </summary>
        public readonly string TuningOperation;
        /// <summary>
        /// The type of model e.g. `home-page`. Currently supported values: `recommended-for-you`, `others-you-may-like`, `frequently-bought-together`, `page-optimization`, `similar-items`, `buy-it-again`, `on-sale-items`, and `recently-viewed`(readonly value). This field together with optimization_objective describe model metadata to use to control model training and serving. See https://cloud.google.com/retail/docs/models for more details on what the model metadata control and which combination of parameters are valid. For invalid combinations of parameters (e.g. type = `frequently-bought-together` and optimization_objective = `ctr`), you receive an error 400 if you try to create/update a recommendation with this set of knobs.
        /// </summary>
        public readonly string Type;
        /// <summary>
        /// Timestamp the Recommendation Model was last updated. E.g. if a Recommendation Model was paused - this would be the time the pause was initiated.
        /// </summary>
        public readonly string UpdateTime;

        [OutputConstructor]
        private GetModelResult(
            string createTime,

            string dataState,

            string displayName,

            string filteringOption,

            string lastTuneTime,

            string name,

            string optimizationObjective,

            string periodicTuningState,

            ImmutableArray<Outputs.GoogleCloudRetailV2ModelServingConfigListResponse> servingConfigLists,

            string servingState,

            string trainingState,

            string tuningOperation,

            string type,

            string updateTime)
        {
            CreateTime = createTime;
            DataState = dataState;
            DisplayName = displayName;
            FilteringOption = filteringOption;
            LastTuneTime = lastTuneTime;
            Name = name;
            OptimizationObjective = optimizationObjective;
            PeriodicTuningState = periodicTuningState;
            ServingConfigLists = servingConfigLists;
            ServingState = servingState;
            TrainingState = trainingState;
            TuningOperation = tuningOperation;
            Type = type;
            UpdateTime = updateTime;
        }
    }
}
