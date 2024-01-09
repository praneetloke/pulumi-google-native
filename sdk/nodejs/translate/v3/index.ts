// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export { AdaptiveMtDatasetArgs } from "./adaptiveMtDataset";
export type AdaptiveMtDataset = import("./adaptiveMtDataset").AdaptiveMtDataset;
export const AdaptiveMtDataset: typeof import("./adaptiveMtDataset").AdaptiveMtDataset = null as any;
utilities.lazyLoad(exports, ["AdaptiveMtDataset"], () => require("./adaptiveMtDataset"));

export { DatasetArgs } from "./dataset";
export type Dataset = import("./dataset").Dataset;
export const Dataset: typeof import("./dataset").Dataset = null as any;
utilities.lazyLoad(exports, ["Dataset"], () => require("./dataset"));

export { GetAdaptiveMtDatasetArgs, GetAdaptiveMtDatasetResult, GetAdaptiveMtDatasetOutputArgs } from "./getAdaptiveMtDataset";
export const getAdaptiveMtDataset: typeof import("./getAdaptiveMtDataset").getAdaptiveMtDataset = null as any;
export const getAdaptiveMtDatasetOutput: typeof import("./getAdaptiveMtDataset").getAdaptiveMtDatasetOutput = null as any;
utilities.lazyLoad(exports, ["getAdaptiveMtDataset","getAdaptiveMtDatasetOutput"], () => require("./getAdaptiveMtDataset"));

export { GetDatasetArgs, GetDatasetResult, GetDatasetOutputArgs } from "./getDataset";
export const getDataset: typeof import("./getDataset").getDataset = null as any;
export const getDatasetOutput: typeof import("./getDataset").getDatasetOutput = null as any;
utilities.lazyLoad(exports, ["getDataset","getDatasetOutput"], () => require("./getDataset"));

export { GetGlossaryArgs, GetGlossaryResult, GetGlossaryOutputArgs } from "./getGlossary";
export const getGlossary: typeof import("./getGlossary").getGlossary = null as any;
export const getGlossaryOutput: typeof import("./getGlossary").getGlossaryOutput = null as any;
utilities.lazyLoad(exports, ["getGlossary","getGlossaryOutput"], () => require("./getGlossary"));

export { GetGlossaryEntryArgs, GetGlossaryEntryResult, GetGlossaryEntryOutputArgs } from "./getGlossaryEntry";
export const getGlossaryEntry: typeof import("./getGlossaryEntry").getGlossaryEntry = null as any;
export const getGlossaryEntryOutput: typeof import("./getGlossaryEntry").getGlossaryEntryOutput = null as any;
utilities.lazyLoad(exports, ["getGlossaryEntry","getGlossaryEntryOutput"], () => require("./getGlossaryEntry"));

export { GetModelArgs, GetModelResult, GetModelOutputArgs } from "./getModel";
export const getModel: typeof import("./getModel").getModel = null as any;
export const getModelOutput: typeof import("./getModel").getModelOutput = null as any;
utilities.lazyLoad(exports, ["getModel","getModelOutput"], () => require("./getModel"));

export { GlossaryArgs } from "./glossary";
export type Glossary = import("./glossary").Glossary;
export const Glossary: typeof import("./glossary").Glossary = null as any;
utilities.lazyLoad(exports, ["Glossary"], () => require("./glossary"));

export { GlossaryEntryArgs } from "./glossaryEntry";
export type GlossaryEntry = import("./glossaryEntry").GlossaryEntry;
export const GlossaryEntry: typeof import("./glossaryEntry").GlossaryEntry = null as any;
utilities.lazyLoad(exports, ["GlossaryEntry"], () => require("./glossaryEntry"));

export { ModelArgs } from "./model";
export type Model = import("./model").Model;
export const Model: typeof import("./model").Model = null as any;
utilities.lazyLoad(exports, ["Model"], () => require("./model"));


const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "google-native:translate/v3:AdaptiveMtDataset":
                return new AdaptiveMtDataset(name, <any>undefined, { urn })
            case "google-native:translate/v3:Dataset":
                return new Dataset(name, <any>undefined, { urn })
            case "google-native:translate/v3:Glossary":
                return new Glossary(name, <any>undefined, { urn })
            case "google-native:translate/v3:GlossaryEntry":
                return new GlossaryEntry(name, <any>undefined, { urn })
            case "google-native:translate/v3:Model":
                return new Model(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("google-native", "translate/v3", _module)
