// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

// Export members:
export * from "./organization";
export * from "./organizationAnalyticDatastore";
export * from "./organizationApi";
export * from "./organizationApiproduct";
export * from "./organizationApiproductRateplan";
export * from "./organizationDatacollector";
export * from "./organizationDeveloper";
export * from "./organizationDeveloperApp";
export * from "./organizationDeveloperAppKey";
export * from "./organizationDeveloperSubscription";
export * from "./organizationEnvgroup";
export * from "./organizationEnvgroupAttachment";
export * from "./organizationEnvironment";
export * from "./organizationEnvironmentAnalyticExport";
export * from "./organizationEnvironmentApiRevisionDebugsession";
export * from "./organizationEnvironmentIamPolicy";
export * from "./organizationEnvironmentKeystore";
export * from "./organizationEnvironmentKeystoreAlias";
export * from "./organizationEnvironmentQuery";
export * from "./organizationEnvironmentReference";
export * from "./organizationEnvironmentResourcefile";
export * from "./organizationEnvironmentTargetserver";
export * from "./organizationEnvironmentTraceConfigOverride";
export * from "./organizationHostQuery";
export * from "./organizationInstance";
export * from "./organizationInstanceAttachment";
export * from "./organizationInstanceCanaryevaluation";
export * from "./organizationInstanceNatAddress";
export * from "./organizationReport";
export * from "./organizationSharedflow";
export * from "./organizationSiteApicategory";

// Import resources to register:
import { Organization } from "./organization";
import { OrganizationAnalyticDatastore } from "./organizationAnalyticDatastore";
import { OrganizationApi } from "./organizationApi";
import { OrganizationApiproduct } from "./organizationApiproduct";
import { OrganizationApiproductRateplan } from "./organizationApiproductRateplan";
import { OrganizationDatacollector } from "./organizationDatacollector";
import { OrganizationDeveloper } from "./organizationDeveloper";
import { OrganizationDeveloperApp } from "./organizationDeveloperApp";
import { OrganizationDeveloperAppKey } from "./organizationDeveloperAppKey";
import { OrganizationDeveloperSubscription } from "./organizationDeveloperSubscription";
import { OrganizationEnvgroup } from "./organizationEnvgroup";
import { OrganizationEnvgroupAttachment } from "./organizationEnvgroupAttachment";
import { OrganizationEnvironment } from "./organizationEnvironment";
import { OrganizationEnvironmentAnalyticExport } from "./organizationEnvironmentAnalyticExport";
import { OrganizationEnvironmentApiRevisionDebugsession } from "./organizationEnvironmentApiRevisionDebugsession";
import { OrganizationEnvironmentIamPolicy } from "./organizationEnvironmentIamPolicy";
import { OrganizationEnvironmentKeystore } from "./organizationEnvironmentKeystore";
import { OrganizationEnvironmentKeystoreAlias } from "./organizationEnvironmentKeystoreAlias";
import { OrganizationEnvironmentQuery } from "./organizationEnvironmentQuery";
import { OrganizationEnvironmentReference } from "./organizationEnvironmentReference";
import { OrganizationEnvironmentResourcefile } from "./organizationEnvironmentResourcefile";
import { OrganizationEnvironmentTargetserver } from "./organizationEnvironmentTargetserver";
import { OrganizationEnvironmentTraceConfigOverride } from "./organizationEnvironmentTraceConfigOverride";
import { OrganizationHostQuery } from "./organizationHostQuery";
import { OrganizationInstance } from "./organizationInstance";
import { OrganizationInstanceAttachment } from "./organizationInstanceAttachment";
import { OrganizationInstanceCanaryevaluation } from "./organizationInstanceCanaryevaluation";
import { OrganizationInstanceNatAddress } from "./organizationInstanceNatAddress";
import { OrganizationReport } from "./organizationReport";
import { OrganizationSharedflow } from "./organizationSharedflow";
import { OrganizationSiteApicategory } from "./organizationSiteApicategory";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "google-native:apigee/v1:Organization":
                return new Organization(name, <any>undefined, { urn })
            case "google-native:apigee/v1:OrganizationAnalyticDatastore":
                return new OrganizationAnalyticDatastore(name, <any>undefined, { urn })
            case "google-native:apigee/v1:OrganizationApi":
                return new OrganizationApi(name, <any>undefined, { urn })
            case "google-native:apigee/v1:OrganizationApiproduct":
                return new OrganizationApiproduct(name, <any>undefined, { urn })
            case "google-native:apigee/v1:OrganizationApiproductRateplan":
                return new OrganizationApiproductRateplan(name, <any>undefined, { urn })
            case "google-native:apigee/v1:OrganizationDatacollector":
                return new OrganizationDatacollector(name, <any>undefined, { urn })
            case "google-native:apigee/v1:OrganizationDeveloper":
                return new OrganizationDeveloper(name, <any>undefined, { urn })
            case "google-native:apigee/v1:OrganizationDeveloperApp":
                return new OrganizationDeveloperApp(name, <any>undefined, { urn })
            case "google-native:apigee/v1:OrganizationDeveloperAppKey":
                return new OrganizationDeveloperAppKey(name, <any>undefined, { urn })
            case "google-native:apigee/v1:OrganizationDeveloperSubscription":
                return new OrganizationDeveloperSubscription(name, <any>undefined, { urn })
            case "google-native:apigee/v1:OrganizationEnvgroup":
                return new OrganizationEnvgroup(name, <any>undefined, { urn })
            case "google-native:apigee/v1:OrganizationEnvgroupAttachment":
                return new OrganizationEnvgroupAttachment(name, <any>undefined, { urn })
            case "google-native:apigee/v1:OrganizationEnvironment":
                return new OrganizationEnvironment(name, <any>undefined, { urn })
            case "google-native:apigee/v1:OrganizationEnvironmentAnalyticExport":
                return new OrganizationEnvironmentAnalyticExport(name, <any>undefined, { urn })
            case "google-native:apigee/v1:OrganizationEnvironmentApiRevisionDebugsession":
                return new OrganizationEnvironmentApiRevisionDebugsession(name, <any>undefined, { urn })
            case "google-native:apigee/v1:OrganizationEnvironmentIamPolicy":
                return new OrganizationEnvironmentIamPolicy(name, <any>undefined, { urn })
            case "google-native:apigee/v1:OrganizationEnvironmentKeystore":
                return new OrganizationEnvironmentKeystore(name, <any>undefined, { urn })
            case "google-native:apigee/v1:OrganizationEnvironmentKeystoreAlias":
                return new OrganizationEnvironmentKeystoreAlias(name, <any>undefined, { urn })
            case "google-native:apigee/v1:OrganizationEnvironmentQuery":
                return new OrganizationEnvironmentQuery(name, <any>undefined, { urn })
            case "google-native:apigee/v1:OrganizationEnvironmentReference":
                return new OrganizationEnvironmentReference(name, <any>undefined, { urn })
            case "google-native:apigee/v1:OrganizationEnvironmentResourcefile":
                return new OrganizationEnvironmentResourcefile(name, <any>undefined, { urn })
            case "google-native:apigee/v1:OrganizationEnvironmentTargetserver":
                return new OrganizationEnvironmentTargetserver(name, <any>undefined, { urn })
            case "google-native:apigee/v1:OrganizationEnvironmentTraceConfigOverride":
                return new OrganizationEnvironmentTraceConfigOverride(name, <any>undefined, { urn })
            case "google-native:apigee/v1:OrganizationHostQuery":
                return new OrganizationHostQuery(name, <any>undefined, { urn })
            case "google-native:apigee/v1:OrganizationInstance":
                return new OrganizationInstance(name, <any>undefined, { urn })
            case "google-native:apigee/v1:OrganizationInstanceAttachment":
                return new OrganizationInstanceAttachment(name, <any>undefined, { urn })
            case "google-native:apigee/v1:OrganizationInstanceCanaryevaluation":
                return new OrganizationInstanceCanaryevaluation(name, <any>undefined, { urn })
            case "google-native:apigee/v1:OrganizationInstanceNatAddress":
                return new OrganizationInstanceNatAddress(name, <any>undefined, { urn })
            case "google-native:apigee/v1:OrganizationReport":
                return new OrganizationReport(name, <any>undefined, { urn })
            case "google-native:apigee/v1:OrganizationSharedflow":
                return new OrganizationSharedflow(name, <any>undefined, { urn })
            case "google-native:apigee/v1:OrganizationSiteApicategory":
                return new OrganizationSiteApicategory(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("google-native", "apigee/v1", _module)
