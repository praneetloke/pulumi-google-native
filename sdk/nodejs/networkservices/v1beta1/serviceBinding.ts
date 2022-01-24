// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "../../utilities";

/**
 * Creates a new ServiceBinding in a given project and location.
 */
export class ServiceBinding extends pulumi.CustomResource {
    /**
     * Get an existing ServiceBinding resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): ServiceBinding {
        return new ServiceBinding(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:networkservices/v1beta1:ServiceBinding';

    /**
     * Returns true if the given object is an instance of ServiceBinding.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ServiceBinding {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ServiceBinding.__pulumiType;
    }

    /**
     * The timestamp when the resource was created.
     */
    public /*out*/ readonly createTime!: pulumi.Output<string>;
    /**
     * Optional. A free-text description of the resource. Max length 1024 characters.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * Optional. The endpoint filter associated with the Service Binding. The syntax is described in http://cloud/service-directory/docs/reference/rpc/google.cloud.servicedirectory.v1#google.cloud.servicedirectory.v1.ResolveServiceRequest
     */
    public readonly endpointFilter!: pulumi.Output<string>;
    /**
     * Optional. Set of label tags associated with the ServiceBinding resource.
     */
    public readonly labels!: pulumi.Output<{[key: string]: string}>;
    /**
     * Name of the ServiceBinding resource. It matches pattern `projects/*&#47;locations/global/serviceBindings/service_binding_name>`.
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * The full service directory service name of the format /projects/*&#47;locations/*&#47;namespaces/*&#47;services/*
     */
    public readonly service!: pulumi.Output<string>;
    /**
     * The timestamp when the resource was updated.
     */
    public /*out*/ readonly updateTime!: pulumi.Output<string>;

    /**
     * Create a ServiceBinding resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ServiceBindingArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.service === undefined) && !opts.urn) {
                throw new Error("Missing required property 'service'");
            }
            if ((!args || args.serviceBindingId === undefined) && !opts.urn) {
                throw new Error("Missing required property 'serviceBindingId'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["endpointFilter"] = args ? args.endpointFilter : undefined;
            resourceInputs["labels"] = args ? args.labels : undefined;
            resourceInputs["location"] = args ? args.location : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["service"] = args ? args.service : undefined;
            resourceInputs["serviceBindingId"] = args ? args.serviceBindingId : undefined;
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        } else {
            resourceInputs["createTime"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["endpointFilter"] = undefined /*out*/;
            resourceInputs["labels"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["service"] = undefined /*out*/;
            resourceInputs["updateTime"] = undefined /*out*/;
        }
        if (!opts.version) {
            opts = pulumi.mergeOptions(opts, { version: utilities.getVersion()});
        }
        super(ServiceBinding.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a ServiceBinding resource.
 */
export interface ServiceBindingArgs {
    /**
     * Optional. A free-text description of the resource. Max length 1024 characters.
     */
    description?: pulumi.Input<string>;
    /**
     * Optional. The endpoint filter associated with the Service Binding. The syntax is described in http://cloud/service-directory/docs/reference/rpc/google.cloud.servicedirectory.v1#google.cloud.servicedirectory.v1.ResolveServiceRequest
     */
    endpointFilter?: pulumi.Input<string>;
    /**
     * Optional. Set of label tags associated with the ServiceBinding resource.
     */
    labels?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
    location?: pulumi.Input<string>;
    /**
     * Name of the ServiceBinding resource. It matches pattern `projects/*&#47;locations/global/serviceBindings/service_binding_name>`.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * The full service directory service name of the format /projects/*&#47;locations/*&#47;namespaces/*&#47;services/*
     */
    service: pulumi.Input<string>;
    serviceBindingId: pulumi.Input<string>;
}
