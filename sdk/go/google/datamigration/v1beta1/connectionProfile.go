// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new connection profile in a given project and location.
type ConnectionProfile struct {
	pulumi.CustomResourceState

	// A CloudSQL database connection profile.
	Cloudsql CloudSqlConnectionProfileResponseOutput `pulumi:"cloudsql"`
	// The timestamp when the resource was created. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The connection profile display name.
	DisplayName pulumi.StringOutput `pulumi:"displayName"`
	// The error details in case of state FAILED.
	Error StatusResponseOutput `pulumi:"error"`
	// The resource labels for connection profile to use to annotate any related underlying resources such as Compute Engine VMs. An object containing a list of "key": "value" pairs. Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// A MySQL database connection profile.
	Mysql MySqlConnectionProfileResponseOutput `pulumi:"mysql"`
	// The name of this connection profile resource in the form of projects/{project}/locations/{location}/instances/{instance}.
	Name pulumi.StringOutput `pulumi:"name"`
	// The database provider.
	Provider pulumi.StringOutput `pulumi:"provider"`
	// The current connection profile state (e.g. DRAFT, READY, or FAILED).
	State pulumi.StringOutput `pulumi:"state"`
	// The timestamp when the resource was last updated. A timestamp in RFC3339 UTC "Zulu" format, accurate to nanoseconds. Example: "2014-10-02T15:01:23.045123456Z".
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewConnectionProfile registers a new resource with the given unique name, arguments, and options.
func NewConnectionProfile(ctx *pulumi.Context,
	name string, args *ConnectionProfileArgs, opts ...pulumi.ResourceOption) (*ConnectionProfile, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConnectionProfileId == nil {
		return nil, errors.New("invalid value for required argument 'ConnectionProfileId'")
	}
	var resource ConnectionProfile
	err := ctx.RegisterResource("google-native:datamigration/v1beta1:ConnectionProfile", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConnectionProfile gets an existing ConnectionProfile resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConnectionProfile(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConnectionProfileState, opts ...pulumi.ResourceOption) (*ConnectionProfile, error) {
	var resource ConnectionProfile
	err := ctx.ReadResource("google-native:datamigration/v1beta1:ConnectionProfile", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConnectionProfile resources.
type connectionProfileState struct {
}

type ConnectionProfileState struct {
}

func (ConnectionProfileState) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionProfileState)(nil)).Elem()
}

type connectionProfileArgs struct {
	// A CloudSQL database connection profile.
	Cloudsql            *CloudSqlConnectionProfile `pulumi:"cloudsql"`
	ConnectionProfileId string                     `pulumi:"connectionProfileId"`
	// The connection profile display name.
	DisplayName *string `pulumi:"displayName"`
	// The resource labels for connection profile to use to annotate any related underlying resources such as Compute Engine VMs. An object containing a list of "key": "value" pairs. Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	// A MySQL database connection profile.
	Mysql *MySqlConnectionProfile `pulumi:"mysql"`
	// The name of this connection profile resource in the form of projects/{project}/locations/{location}/instances/{instance}.
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The database provider.
	Provider  *ConnectionProfileProvider `pulumi:"provider"`
	RequestId *string                    `pulumi:"requestId"`
	// The current connection profile state (e.g. DRAFT, READY, or FAILED).
	State *ConnectionProfileStateEnum `pulumi:"state"`
}

// The set of arguments for constructing a ConnectionProfile resource.
type ConnectionProfileArgs struct {
	// A CloudSQL database connection profile.
	Cloudsql            CloudSqlConnectionProfilePtrInput
	ConnectionProfileId pulumi.StringInput
	// The connection profile display name.
	DisplayName pulumi.StringPtrInput
	// The resource labels for connection profile to use to annotate any related underlying resources such as Compute Engine VMs. An object containing a list of "key": "value" pairs. Example: `{ "name": "wrench", "mass": "1.3kg", "count": "3" }`.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	// A MySQL database connection profile.
	Mysql MySqlConnectionProfilePtrInput
	// The name of this connection profile resource in the form of projects/{project}/locations/{location}/instances/{instance}.
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The database provider.
	Provider  ConnectionProfileProviderPtrInput
	RequestId pulumi.StringPtrInput
	// The current connection profile state (e.g. DRAFT, READY, or FAILED).
	State ConnectionProfileStateEnumPtrInput
}

func (ConnectionProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*connectionProfileArgs)(nil)).Elem()
}

type ConnectionProfileInput interface {
	pulumi.Input

	ToConnectionProfileOutput() ConnectionProfileOutput
	ToConnectionProfileOutputWithContext(ctx context.Context) ConnectionProfileOutput
}

func (*ConnectionProfile) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionProfile)(nil))
}

func (i *ConnectionProfile) ToConnectionProfileOutput() ConnectionProfileOutput {
	return i.ToConnectionProfileOutputWithContext(context.Background())
}

func (i *ConnectionProfile) ToConnectionProfileOutputWithContext(ctx context.Context) ConnectionProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionProfileOutput)
}

type ConnectionProfileOutput struct {
	*pulumi.OutputState
}

func (ConnectionProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionProfile)(nil))
}

func (o ConnectionProfileOutput) ToConnectionProfileOutput() ConnectionProfileOutput {
	return o
}

func (o ConnectionProfileOutput) ToConnectionProfileOutputWithContext(ctx context.Context) ConnectionProfileOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ConnectionProfileOutput{})
}
