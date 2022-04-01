// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a repository. The returned Operation will finish once the repository has been created. Its response will be the created Repository.
type Repository struct {
	pulumi.CustomResourceState

	// The time when the repository was created.
	CreateTime pulumi.StringOutput `pulumi:"createTime"`
	// The user-provided description of the repository.
	Description pulumi.StringOutput `pulumi:"description"`
	// The format of packages that are stored in the repository.
	Format pulumi.StringOutput `pulumi:"format"`
	// The Cloud KMS resource name of the customer managed encryption key that's used to encrypt the contents of the Repository. Has the form: `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`. This value may not be changed after the Repository has been created.
	KmsKeyName pulumi.StringOutput `pulumi:"kmsKeyName"`
	// Labels with user-defined metadata. This field may contain up to 64 entries. Label keys and values may be no longer than 63 characters. Label keys must begin with a lowercase letter and may only contain lowercase letters, numeric characters, underscores, and dashes.
	Labels pulumi.StringMapOutput `pulumi:"labels"`
	// Maven repository config contains repository level configuration for the repositories of maven type.
	MavenConfig MavenRepositoryConfigResponseOutput `pulumi:"mavenConfig"`
	// The name of the repository, for example: "projects/p1/locations/us-central1/repositories/repo1".
	Name pulumi.StringOutput `pulumi:"name"`
	// The size, in bytes, of all artifact storage in this repository. Repositories that are generally available or in public preview use this to calculate storage costs.
	SizeBytes pulumi.StringOutput `pulumi:"sizeBytes"`
	// The time when the repository was last updated.
	UpdateTime pulumi.StringOutput `pulumi:"updateTime"`
}

// NewRepository registers a new resource with the given unique name, arguments, and options.
func NewRepository(ctx *pulumi.Context,
	name string, args *RepositoryArgs, opts ...pulumi.ResourceOption) (*Repository, error) {
	if args == nil {
		args = &RepositoryArgs{}
	}

	var resource Repository
	err := ctx.RegisterResource("google-native:artifactregistry/v1:Repository", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRepository gets an existing Repository resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRepository(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RepositoryState, opts ...pulumi.ResourceOption) (*Repository, error) {
	var resource Repository
	err := ctx.ReadResource("google-native:artifactregistry/v1:Repository", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Repository resources.
type repositoryState struct {
}

type RepositoryState struct {
}

func (RepositoryState) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryState)(nil)).Elem()
}

type repositoryArgs struct {
	// The time when the repository was created.
	CreateTime *string `pulumi:"createTime"`
	// The user-provided description of the repository.
	Description *string `pulumi:"description"`
	// The format of packages that are stored in the repository.
	Format *RepositoryFormat `pulumi:"format"`
	// The Cloud KMS resource name of the customer managed encryption key that's used to encrypt the contents of the Repository. Has the form: `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`. This value may not be changed after the Repository has been created.
	KmsKeyName *string `pulumi:"kmsKeyName"`
	// Labels with user-defined metadata. This field may contain up to 64 entries. Label keys and values may be no longer than 63 characters. Label keys must begin with a lowercase letter and may only contain lowercase letters, numeric characters, underscores, and dashes.
	Labels   map[string]string `pulumi:"labels"`
	Location *string           `pulumi:"location"`
	// Maven repository config contains repository level configuration for the repositories of maven type.
	MavenConfig *MavenRepositoryConfig `pulumi:"mavenConfig"`
	// The name of the repository, for example: "projects/p1/locations/us-central1/repositories/repo1".
	Name    *string `pulumi:"name"`
	Project *string `pulumi:"project"`
	// The repository id to use for this repository.
	RepositoryId *string `pulumi:"repositoryId"`
	// The time when the repository was last updated.
	UpdateTime *string `pulumi:"updateTime"`
}

// The set of arguments for constructing a Repository resource.
type RepositoryArgs struct {
	// The time when the repository was created.
	CreateTime pulumi.StringPtrInput
	// The user-provided description of the repository.
	Description pulumi.StringPtrInput
	// The format of packages that are stored in the repository.
	Format RepositoryFormatPtrInput
	// The Cloud KMS resource name of the customer managed encryption key that's used to encrypt the contents of the Repository. Has the form: `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`. This value may not be changed after the Repository has been created.
	KmsKeyName pulumi.StringPtrInput
	// Labels with user-defined metadata. This field may contain up to 64 entries. Label keys and values may be no longer than 63 characters. Label keys must begin with a lowercase letter and may only contain lowercase letters, numeric characters, underscores, and dashes.
	Labels   pulumi.StringMapInput
	Location pulumi.StringPtrInput
	// Maven repository config contains repository level configuration for the repositories of maven type.
	MavenConfig MavenRepositoryConfigPtrInput
	// The name of the repository, for example: "projects/p1/locations/us-central1/repositories/repo1".
	Name    pulumi.StringPtrInput
	Project pulumi.StringPtrInput
	// The repository id to use for this repository.
	RepositoryId pulumi.StringPtrInput
	// The time when the repository was last updated.
	UpdateTime pulumi.StringPtrInput
}

func (RepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*repositoryArgs)(nil)).Elem()
}

type RepositoryInput interface {
	pulumi.Input

	ToRepositoryOutput() RepositoryOutput
	ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput
}

func (*Repository) ElementType() reflect.Type {
	return reflect.TypeOf((**Repository)(nil)).Elem()
}

func (i *Repository) ToRepositoryOutput() RepositoryOutput {
	return i.ToRepositoryOutputWithContext(context.Background())
}

func (i *Repository) ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryOutput)
}

type RepositoryOutput struct{ *pulumi.OutputState }

func (RepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Repository)(nil)).Elem()
}

func (o RepositoryOutput) ToRepositoryOutput() RepositoryOutput {
	return o
}

func (o RepositoryOutput) ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RepositoryInput)(nil)).Elem(), &Repository{})
	pulumi.RegisterOutputType(RepositoryOutput{})
}
