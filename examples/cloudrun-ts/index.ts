// Copyright 2016-2021, Pulumi Corporation.

import * as pulumi from "@pulumi/pulumi";
import * as google from "@pulumi/google-native";
import * as random from "@pulumi/random"

const config = new pulumi.Config("google-native");
const project = config.require("project");
const region = config.require("region");

const randomString = new random.RandomString("service-name", {
    upper: false,
    number: false,
    special: false,
    length: 8,
});
const serviceName = pulumi.interpolate`run-${randomString.result}`;

const service = new google.run.v1.Service("service", {
    project,
    location: region,
    serviceId: serviceName,
    apiVersion: "serving.knative.dev/v1",
    kind: "Service",
    metadata: {
        name: serviceName,
    },
    spec: {
        template: {
            spec: {
                containers: [{image: "gcr.io/cloudrun/hello"}],
            },
        },
    },
});

const iamHello = new google.run.v1.ServiceIamPolicy("allow-all", {
    project,
    location: region,
    serviceId: service.metadata.name,
    bindings: [{
        members: ["allUsers"],
        role: "roles/run.invoker",
    }],
});

export const serviceUrl = service.status.url;
