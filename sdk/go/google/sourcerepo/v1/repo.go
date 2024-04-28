// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a repo in the given project with the given name. If the named repository already exists, `CreateRepo` returns `ALREADY_EXISTS`.
type Repo struct {
	pulumi.CustomResourceState

	// How this repository mirrors a repository managed by another service. Read-only field.
	MirrorConfig MirrorConfigResponseOutput `pulumi:"mirrorConfig"`
	// Resource name of the repository, of the form `projects//repos/`. The repo name may contain slashes. eg, `projects/myproject/repos/name/with/slash`
	Name    pulumi.StringOutput `pulumi:"name"`
	Project pulumi.StringOutput `pulumi:"project"`
	// How this repository publishes a change in the repository through Cloud Pub/Sub. Keyed by the topic names.
	PubsubConfigs PubsubConfigResponseOutput `pulumi:"pubsubConfigs"`
	// The disk usage of the repo, in bytes. Read-only field. Size is only returned by GetRepo.
	Size pulumi.StringOutput `pulumi:"size"`
	// URL to clone the repository from Google Cloud Source Repositories. Read-only field.
	Url pulumi.StringOutput `pulumi:"url"`
}

// NewRepo registers a new resource with the given unique name, arguments, and options.
func NewRepo(ctx *pulumi.Context,
	name string, args *RepoArgs, opts ...pulumi.ResourceOption) (*Repo, error) {
	if args == nil {
		args = &RepoArgs{}
	}

	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Repo
	err := ctx.RegisterResource("google-native:sourcerepo/v1:Repo", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepo gets an existing Repo resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepo(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepoState, opts ...pulumi.ResourceOption) (*Repo, error) {
	var resource Repo
	err := ctx.ReadResource("google-native:sourcerepo/v1:Repo", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Repo resources.
type repoState struct {
}

type RepoState struct {
}

func (RepoState) ElementType() reflect.Type {
	return reflect.TypeOf((*repoState)(nil)).Elem()
}

type repoArgs struct {
	// How this repository mirrors a repository managed by another service. Read-only field.
	MirrorConfig *MirrorConfig `pulumi:"mirrorConfig"`
	// Resource name of the repository, of the form `projects//repos/`. The repo name may contain slashes. eg, `projects/myproject/repos/name/with/slash`
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// How this repository publishes a change in the repository through Cloud Pub/Sub. Keyed by the topic names.
	PubsubConfigs *PubsubConfig `pulumi:"pubsubConfigs"`
	// The disk usage of the repo, in bytes. Read-only field. Size is only returned by GetRepo.
	Size *string `pulumi:"size"`
	// URL to clone the repository from Google Cloud Source Repositories. Read-only field.
	Url *string `pulumi:"url"`
}

// The set of arguments for constructing a Repo resource.
type RepoArgs struct {
	// How this repository mirrors a repository managed by another service. Read-only field.
	MirrorConfig MirrorConfigPtrInput
	// Resource name of the repository, of the form `projects//repos/`. The repo name may contain slashes. eg, `projects/myproject/repos/name/with/slash`
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// How this repository publishes a change in the repository through Cloud Pub/Sub. Keyed by the topic names.
	PubsubConfigs PubsubConfigPtrInput
	// The disk usage of the repo, in bytes. Read-only field. Size is only returned by GetRepo.
	Size pulumi.StringPtrInput
	// URL to clone the repository from Google Cloud Source Repositories. Read-only field.
	Url pulumi.StringPtrInput
}

func (RepoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repoArgs)(nil)).Elem()
}

type RepoInput interface {
	pulumi.Input

	ToRepoOutput() RepoOutput
	ToRepoOutputWithContext(ctx context.Context) RepoOutput
}

func (*Repo) ElementType() reflect.Type {
	return reflect.TypeOf((**Repo)(nil)).Elem()
}

func (i *Repo) ToRepoOutput() RepoOutput {
	return i.ToRepoOutputWithContext(context.Background())
}

func (i *Repo) ToRepoOutputWithContext(ctx context.Context) RepoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepoOutput)
}

type RepoOutput struct{ *pulumi.OutputState }

func (RepoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Repo)(nil)).Elem()
}

func (o RepoOutput) ToRepoOutput() RepoOutput {
	return o
}

func (o RepoOutput) ToRepoOutputWithContext(ctx context.Context) RepoOutput {
	return o
}

// How this repository mirrors a repository managed by another service. Read-only field.
func (o RepoOutput) MirrorConfig() MirrorConfigResponseOutput {
	return o.ApplyT(func(v *Repo) MirrorConfigResponseOutput { return v.MirrorConfig }).(MirrorConfigResponseOutput)
}

// Resource name of the repository, of the form `projects//repos/`. The repo name may contain slashes. eg, `projects/myproject/repos/name/with/slash`
func (o RepoOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Repo) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RepoOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Repo) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// How this repository publishes a change in the repository through Cloud Pub/Sub. Keyed by the topic names.
func (o RepoOutput) PubsubConfigs() PubsubConfigResponseOutput {
	return o.ApplyT(func(v *Repo) PubsubConfigResponseOutput { return v.PubsubConfigs }).(PubsubConfigResponseOutput)
}

// The disk usage of the repo, in bytes. Read-only field. Size is only returned by GetRepo.
func (o RepoOutput) Size() pulumi.StringOutput {
	return o.ApplyT(func(v *Repo) pulumi.StringOutput { return v.Size }).(pulumi.StringOutput)
}

// URL to clone the repository from Google Cloud Source Repositories. Read-only field.
func (o RepoOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v *Repo) pulumi.StringOutput { return v.Url }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RepoInput)(nil)).Elem(), &Repo{})
	pulumi.RegisterOutputType(RepoOutput{})
}
