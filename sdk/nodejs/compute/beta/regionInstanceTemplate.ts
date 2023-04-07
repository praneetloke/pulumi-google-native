// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "../../types/input";
import * as outputs from "../../types/output";
import * as enums from "../../types/enums";
import * as utilities from "../../utilities";

/**
 * Creates an instance template in the specified project and region using the global instance template whose URL is included in the request.
 */
export class RegionInstanceTemplate extends pulumi.CustomResource {
    /**
     * Get an existing RegionInstanceTemplate resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, opts?: pulumi.CustomResourceOptions): RegionInstanceTemplate {
        return new RegionInstanceTemplate(name, undefined as any, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'google-native:compute/beta:RegionInstanceTemplate';

    /**
     * Returns true if the given object is an instance of RegionInstanceTemplate.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is RegionInstanceTemplate {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === RegionInstanceTemplate.__pulumiType;
    }

    /**
     * The creation timestamp for this instance template in RFC3339 text format.
     */
    public /*out*/ readonly creationTimestamp!: pulumi.Output<string>;
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    public readonly description!: pulumi.Output<string>;
    /**
     * The resource type, which is always compute#instanceTemplate for instance templates.
     */
    public /*out*/ readonly kind!: pulumi.Output<string>;
    /**
     * Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    public readonly name!: pulumi.Output<string>;
    public readonly project!: pulumi.Output<string>;
    /**
     * The instance properties for this instance template.
     */
    public readonly properties!: pulumi.Output<outputs.compute.beta.InstancePropertiesResponse>;
    public readonly region!: pulumi.Output<string>;
    /**
     * An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
     */
    public readonly requestId!: pulumi.Output<string | undefined>;
    /**
     * The URL for this instance template. The server defines this URL.
     */
    public /*out*/ readonly selfLink!: pulumi.Output<string>;
    /**
     * The source instance used to create the template. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instances/instance - projects/project/zones/zone/instances/instance 
     */
    public readonly sourceInstance!: pulumi.Output<string>;
    /**
     * The source instance params to use to create this instance template.
     */
    public readonly sourceInstanceParams!: pulumi.Output<outputs.compute.beta.SourceInstanceParamsResponse>;

    /**
     * Create a RegionInstanceTemplate resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: RegionInstanceTemplateArgs, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.region === undefined) && !opts.urn) {
                throw new Error("Missing required property 'region'");
            }
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["project"] = args ? args.project : undefined;
            resourceInputs["properties"] = args ? args.properties : undefined;
            resourceInputs["region"] = args ? args.region : undefined;
            resourceInputs["requestId"] = args ? args.requestId : undefined;
            resourceInputs["sourceInstance"] = args ? args.sourceInstance : undefined;
            resourceInputs["sourceInstanceParams"] = args ? args.sourceInstanceParams : undefined;
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
        } else {
            resourceInputs["creationTimestamp"] = undefined /*out*/;
            resourceInputs["description"] = undefined /*out*/;
            resourceInputs["kind"] = undefined /*out*/;
            resourceInputs["name"] = undefined /*out*/;
            resourceInputs["project"] = undefined /*out*/;
            resourceInputs["properties"] = undefined /*out*/;
            resourceInputs["region"] = undefined /*out*/;
            resourceInputs["requestId"] = undefined /*out*/;
            resourceInputs["selfLink"] = undefined /*out*/;
            resourceInputs["sourceInstance"] = undefined /*out*/;
            resourceInputs["sourceInstanceParams"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        const replaceOnChanges = { replaceOnChanges: ["project", "region"] };
        opts = pulumi.mergeOptions(opts, replaceOnChanges);
        super(RegionInstanceTemplate.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * The set of arguments for constructing a RegionInstanceTemplate resource.
 */
export interface RegionInstanceTemplateArgs {
    /**
     * An optional description of this resource. Provide this property when you create the resource.
     */
    description?: pulumi.Input<string>;
    /**
     * Name of the resource; provided by the client when the resource is created. The name must be 1-63 characters long, and comply with RFC1035. Specifically, the name must be 1-63 characters long and match the regular expression `[a-z]([-a-z0-9]*[a-z0-9])?` which means the first character must be a lowercase letter, and all following characters must be a dash, lowercase letter, or digit, except the last character, which cannot be a dash.
     */
    name?: pulumi.Input<string>;
    project?: pulumi.Input<string>;
    /**
     * The instance properties for this instance template.
     */
    properties?: pulumi.Input<inputs.compute.beta.InstancePropertiesArgs>;
    region: pulumi.Input<string>;
    /**
     * An optional request ID to identify requests. Specify a unique request ID so that if you must retry your request, the server will know to ignore the request if it has already been completed. For example, consider a situation where you make an initial request and the request times out. If you make the request again with the same request ID, the server can check if original operation with the same request ID was received, and if so, will ignore the second request. This prevents clients from accidentally creating duplicate commitments. The request ID must be a valid UUID with the exception that zero UUID is not supported ( 00000000-0000-0000-0000-000000000000).
     */
    requestId?: pulumi.Input<string>;
    /**
     * The source instance used to create the template. You can provide this as a partial or full URL to the resource. For example, the following are valid values: - https://www.googleapis.com/compute/v1/projects/project/zones/zone /instances/instance - projects/project/zones/zone/instances/instance 
     */
    sourceInstance?: pulumi.Input<string>;
    /**
     * The source instance params to use to create this instance template.
     */
    sourceInstanceParams?: pulumi.Input<inputs.compute.beta.SourceInstanceParamsArgs>;
}
