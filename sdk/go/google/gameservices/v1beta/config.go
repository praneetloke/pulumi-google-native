// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new game server config in a given project, location, and game server deployment. Game server configs are immutable, and are not applied until referenced in the game server deployment rollout resource.
type Config struct {
	pulumi.CustomResourceState

	// The creation time.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The description of the game server config.
	Description pulumi.StringOutput `pulumi:"description"`
	// FleetConfig contains a list of Agones fleet specs. Only one FleetConfig is allowed.
	FleetConfigs FleetConfigResponseArrayOutput `pulumi:"fleetConfigs"`
	// The labels associated with this game server config. Each label is a key-value pair.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// The resource name of the game server config, in the following form: `projects/{project}/locations/{location}/gameServerDeployments/{deployment}/configs/{config}`. For example, `projects/my-project/locations/global/gameServerDeployments/my-game/configs/my-config`.
	Name pulumi.StringOutput `pulumi:"name"`
	// The autoscaling settings.
	ScalingConfigs ScalingConfigResponseArrayOutput `pulumi:"scalingConfigs"`
	// The last-modified time.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewConfig registers a new resource with the given unique name, arguments, and options.
func NewConfig(ctx *pulumi.Context,
	name string, args *ConfigArgs, opts ...pulumi.ResourceOption) (*Config, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConfigId == nil {
		return nil, errors.New("invalid value for required argument 'ConfigId'")
	}
	if args.GameServerDeploymentId == nil {
		return nil, errors.New("invalid value for required argument 'GameServerDeploymentId'")
	}
	var resource Config
	err := ctx.RegisterResource("google-native:gameservices/v1beta:Config", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConfig gets an existing Config resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConfig(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConfigState, opts ...pulumi.ResourceOption) (*Config, error) {
	var resource Config
	err := ctx.ReadResource("google-native:gameservices/v1beta:Config", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Config resources.
type configState struct {
}

type ConfigState struct {
}

func (ConfigState) ElementType() reflect.Type {
	return reflect.TypeOf((*configState)(nil)).Elem()
}

type configArgs struct {
	ConfigId string `pulumi:"configId"`
	// The description of the game server config.
	Description *string `pulumi:"description"`
	// FleetConfig contains a list of Agones fleet specs. Only one FleetConfig is allowed.
	FleetConfigs           []FleetConfig `pulumi:"fleetConfigs"`
	GameServerDeploymentId string        `pulumi:"gameServerDeploymentId"`
	// The labels associated with this game server config. Each label is a key-value pair.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	// The resource name of the game server config, in the following form: `projects/{project}/locations/{location}/gameServerDeployments/{deployment}/configs/{config}`. For example, `projects/my-project/locations/global/gameServerDeployments/my-game/configs/my-config`.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The autoscaling settings.
	ScalingConfigs []ScalingConfig `pulumi:"scalingConfigs"`
}

// The set of arguments for constructing a Config resource.
type ConfigArgs struct {
	ConfigId pulumi.StringInput
	// The description of the game server config.
	Description pulumi.StringPtrInput
	// FleetConfig contains a list of Agones fleet specs. Only one FleetConfig is allowed.
	FleetConfigs           FleetConfigArrayInput
	GameServerDeploymentId pulumi.StringInput
	// The labels associated with this game server config. Each label is a key-value pair.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	// The resource name of the game server config, in the following form: `projects/{project}/locations/{location}/gameServerDeployments/{deployment}/configs/{config}`. For example, `projects/my-project/locations/global/gameServerDeployments/my-game/configs/my-config`.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The autoscaling settings.
	ScalingConfigs ScalingConfigArrayInput
}

func (ConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*configArgs)(nil)).Elem()
}

type ConfigInput interface {
	pulumi.Input

	ToConfigOutput() ConfigOutput
	ToConfigOutputWithContext(ctx context.Context) ConfigOutput
}

func (*Config) ElementType() reflect.Type {
	return reflect.TypeOf((*Config)(nil))
}

func (i *Config) ToConfigOutput() ConfigOutput {
	return i.ToConfigOutputWithContext(context.Background())
}

func (i *Config) ToConfigOutputWithContext(ctx context.Context) ConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigOutput)
}

type ConfigOutput struct {
	*pulumi.OutputState
}

func (ConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Config)(nil))
}

func (o ConfigOutput) ToConfigOutput() ConfigOutput {
	return o
}

func (o ConfigOutput) ToConfigOutputWithContext(ctx context.Context) ConfigOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ConfigOutput{})
}
