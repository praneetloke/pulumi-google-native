// *** WARNING: this file was generated by the Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta2

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Gets a repository.
func LookupRepository(ctx *pulumi.Context, args *LookupRepositoryArgs, opts ...pulumi.InvokeOption) (*LookupRepositoryResult, error) {
	var rv LookupRepositoryResult
	err := ctx.Invoke("google-native:artifactregistry/v1beta2:getRepository", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRepositoryArgs struct {
	Location     string  `pulumi:"location"`
	Project      *string `pulumi:"project"`
	RepositoryId string  `pulumi:"repositoryId"`
}

type LookupRepositoryResult struct {
	// The time when the repository was created.
	CreateTime string `pulumi:"createTime"`
	// The user-provided description of the repository.
	Description string `pulumi:"description"`
	// The format of packages that are stored in the repository.
	Format string `pulumi:"format"`
	// The Cloud KMS resource name of the customer managed encryption key that's used to encrypt the contents of the Repository. Has the form: `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`. This value may not be changed after the Repository has been created.
	KmsKeyName string `pulumi:"kmsKeyName"`
	// Labels with user-defined metadata. This field may contain up to 64 entries. Label keys and values may be no longer than 63 characters. Label keys must begin with a lowercase letter and may only contain lowercase letters, numeric characters, underscores, and dashes.
	Labels map[string]string `pulumi:"labels"`
	// Maven repository config contains repository level configuration for the repositories of maven type.
	MavenConfig MavenRepositoryConfigResponse `pulumi:"mavenConfig"`
	// The name of the repository, for example: "projects/p1/locations/us-central1/repositories/repo1".
	Name string `pulumi:"name"`
	// The size, in bytes, of all artifact storage in this repository. Repositories that are generally available or in public preview use this to calculate storage costs.
	SizeBytes string `pulumi:"sizeBytes"`
	// The time when the repository was last updated.
	UpdateTime string `pulumi:"updateTime"`
}

func LookupRepositoryOutput(ctx *pulumi.Context, args LookupRepositoryOutputArgs, opts ...pulumi.InvokeOption) LookupRepositoryResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRepositoryResult, error) {
			args := v.(LookupRepositoryArgs)
			r, err := LookupRepository(ctx, &args, opts...)
			return *r, err
		}).(LookupRepositoryResultOutput)
}

type LookupRepositoryOutputArgs struct {
	Location     pulumi.StringInput    `pulumi:"location"`
	Project      pulumi.StringPtrInput `pulumi:"project"`
	RepositoryId pulumi.StringInput    `pulumi:"repositoryId"`
}

func (LookupRepositoryOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRepositoryArgs)(nil)).Elem()
}

type LookupRepositoryResultOutput struct{ *pulumi.OutputState }

func (LookupRepositoryResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRepositoryResult)(nil)).Elem()
}

func (o LookupRepositoryResultOutput) ToLookupRepositoryResultOutput() LookupRepositoryResultOutput {
	return o
}

func (o LookupRepositoryResultOutput) ToLookupRepositoryResultOutputWithContext(ctx context.Context) LookupRepositoryResultOutput {
	return o
}

// The time when the repository was created.
func (o LookupRepositoryResultOutput) CreateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryResult) string { return v.CreateTime }).(pulumi.StringOutput)
}

// The user-provided description of the repository.
func (o LookupRepositoryResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryResult) string { return v.Description }).(pulumi.StringOutput)
}

// The format of packages that are stored in the repository.
func (o LookupRepositoryResultOutput) Format() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryResult) string { return v.Format }).(pulumi.StringOutput)
}

// The Cloud KMS resource name of the customer managed encryption key that's used to encrypt the contents of the Repository. Has the form: `projects/my-project/locations/my-region/keyRings/my-kr/cryptoKeys/my-key`. This value may not be changed after the Repository has been created.
func (o LookupRepositoryResultOutput) KmsKeyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryResult) string { return v.KmsKeyName }).(pulumi.StringOutput)
}

// Labels with user-defined metadata. This field may contain up to 64 entries. Label keys and values may be no longer than 63 characters. Label keys must begin with a lowercase letter and may only contain lowercase letters, numeric characters, underscores, and dashes.
func (o LookupRepositoryResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupRepositoryResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Maven repository config contains repository level configuration for the repositories of maven type.
func (o LookupRepositoryResultOutput) MavenConfig() MavenRepositoryConfigResponseOutput {
	return o.ApplyT(func(v LookupRepositoryResult) MavenRepositoryConfigResponse { return v.MavenConfig }).(MavenRepositoryConfigResponseOutput)
}

// The name of the repository, for example: "projects/p1/locations/us-central1/repositories/repo1".
func (o LookupRepositoryResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryResult) string { return v.Name }).(pulumi.StringOutput)
}

// The size, in bytes, of all artifact storage in this repository. Repositories that are generally available or in public preview use this to calculate storage costs.
func (o LookupRepositoryResultOutput) SizeBytes() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryResult) string { return v.SizeBytes }).(pulumi.StringOutput)
}

// The time when the repository was last updated.
func (o LookupRepositoryResultOutput) UpdateTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRepositoryResult) string { return v.UpdateTime }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRepositoryResultOutput{})
}
