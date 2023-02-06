// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v1beta1

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Associates `members`, or principals, with a `role`.
type Binding struct {
	// The condition that is associated with this binding. If the condition evaluates to `true`, then this binding applies to the current request. If the condition evaluates to `false`, then this binding does not apply to the current request. However, a different role binding might grant the same role to one or more of the principals in this binding. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Condition *Expr `pulumi:"condition"`
	// Specifies the principals requesting access for a Google Cloud resource. `members` can have the following values: * `allUsers`: A special identifier that represents anyone who is on the internet; with or without a Google account. * `allAuthenticatedUsers`: A special identifier that represents anyone who is authenticated with a Google account or a service account. Does not include identities that come from external identity providers (IdPs) through identity federation. * `user:{emailid}`: An email address that represents a specific Google account. For example, `alice@example.com` . * `serviceAccount:{emailid}`: An email address that represents a Google service account. For example, `my-other-app@appspot.gserviceaccount.com`. * `serviceAccount:{projectid}.svc.id.goog[{namespace}/{kubernetes-sa}]`: An identifier for a [Kubernetes service account](https://cloud.google.com/kubernetes-engine/docs/how-to/kubernetes-service-accounts). For example, `my-project.svc.id.goog[my-namespace/my-kubernetes-sa]`. * `group:{emailid}`: An email address that represents a Google group. For example, `admins@example.com`. * `domain:{domain}`: The G Suite domain (primary) that represents all the users of that domain. For example, `google.com` or `example.com`. * `deleted:user:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a user that has been recently deleted. For example, `alice@example.com?uid=123456789012345678901`. If the user is recovered, this value reverts to `user:{emailid}` and the recovered user retains the role in the binding. * `deleted:serviceAccount:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, `my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901`. If the service account is undeleted, this value reverts to `serviceAccount:{emailid}` and the undeleted service account retains the role in the binding. * `deleted:group:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, `admins@example.com?uid=123456789012345678901`. If the group is recovered, this value reverts to `group:{emailid}` and the recovered group retains the role in the binding.
	Members []string `pulumi:"members"`
	// Role that is assigned to the list of `members`, or principals. For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
	Role *string `pulumi:"role"`
}

// BindingInput is an input type that accepts BindingArgs and BindingOutput values.
// You can construct a concrete instance of `BindingInput` via:
//
//	BindingArgs{...}
type BindingInput interface {
	pulumi.Input

	ToBindingOutput() BindingOutput
	ToBindingOutputWithContext(context.Context) BindingOutput
}

// Associates `members`, or principals, with a `role`.
type BindingArgs struct {
	// The condition that is associated with this binding. If the condition evaluates to `true`, then this binding applies to the current request. If the condition evaluates to `false`, then this binding does not apply to the current request. However, a different role binding might grant the same role to one or more of the principals in this binding. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Condition ExprPtrInput `pulumi:"condition"`
	// Specifies the principals requesting access for a Google Cloud resource. `members` can have the following values: * `allUsers`: A special identifier that represents anyone who is on the internet; with or without a Google account. * `allAuthenticatedUsers`: A special identifier that represents anyone who is authenticated with a Google account or a service account. Does not include identities that come from external identity providers (IdPs) through identity federation. * `user:{emailid}`: An email address that represents a specific Google account. For example, `alice@example.com` . * `serviceAccount:{emailid}`: An email address that represents a Google service account. For example, `my-other-app@appspot.gserviceaccount.com`. * `serviceAccount:{projectid}.svc.id.goog[{namespace}/{kubernetes-sa}]`: An identifier for a [Kubernetes service account](https://cloud.google.com/kubernetes-engine/docs/how-to/kubernetes-service-accounts). For example, `my-project.svc.id.goog[my-namespace/my-kubernetes-sa]`. * `group:{emailid}`: An email address that represents a Google group. For example, `admins@example.com`. * `domain:{domain}`: The G Suite domain (primary) that represents all the users of that domain. For example, `google.com` or `example.com`. * `deleted:user:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a user that has been recently deleted. For example, `alice@example.com?uid=123456789012345678901`. If the user is recovered, this value reverts to `user:{emailid}` and the recovered user retains the role in the binding. * `deleted:serviceAccount:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, `my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901`. If the service account is undeleted, this value reverts to `serviceAccount:{emailid}` and the undeleted service account retains the role in the binding. * `deleted:group:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, `admins@example.com?uid=123456789012345678901`. If the group is recovered, this value reverts to `group:{emailid}` and the recovered group retains the role in the binding.
	Members pulumi.StringArrayInput `pulumi:"members"`
	// Role that is assigned to the list of `members`, or principals. For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
	Role pulumi.StringPtrInput `pulumi:"role"`
}

func (BindingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Binding)(nil)).Elem()
}

func (i BindingArgs) ToBindingOutput() BindingOutput {
	return i.ToBindingOutputWithContext(context.Background())
}

func (i BindingArgs) ToBindingOutputWithContext(ctx context.Context) BindingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BindingOutput)
}

// BindingArrayInput is an input type that accepts BindingArray and BindingArrayOutput values.
// You can construct a concrete instance of `BindingArrayInput` via:
//
//	BindingArray{ BindingArgs{...} }
type BindingArrayInput interface {
	pulumi.Input

	ToBindingArrayOutput() BindingArrayOutput
	ToBindingArrayOutputWithContext(context.Context) BindingArrayOutput
}

type BindingArray []BindingInput

func (BindingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Binding)(nil)).Elem()
}

func (i BindingArray) ToBindingArrayOutput() BindingArrayOutput {
	return i.ToBindingArrayOutputWithContext(context.Background())
}

func (i BindingArray) ToBindingArrayOutputWithContext(ctx context.Context) BindingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BindingArrayOutput)
}

// Associates `members`, or principals, with a `role`.
type BindingOutput struct{ *pulumi.OutputState }

func (BindingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Binding)(nil)).Elem()
}

func (o BindingOutput) ToBindingOutput() BindingOutput {
	return o
}

func (o BindingOutput) ToBindingOutputWithContext(ctx context.Context) BindingOutput {
	return o
}

// The condition that is associated with this binding. If the condition evaluates to `true`, then this binding applies to the current request. If the condition evaluates to `false`, then this binding does not apply to the current request. However, a different role binding might grant the same role to one or more of the principals in this binding. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
func (o BindingOutput) Condition() ExprPtrOutput {
	return o.ApplyT(func(v Binding) *Expr { return v.Condition }).(ExprPtrOutput)
}

// Specifies the principals requesting access for a Google Cloud resource. `members` can have the following values: * `allUsers`: A special identifier that represents anyone who is on the internet; with or without a Google account. * `allAuthenticatedUsers`: A special identifier that represents anyone who is authenticated with a Google account or a service account. Does not include identities that come from external identity providers (IdPs) through identity federation. * `user:{emailid}`: An email address that represents a specific Google account. For example, `alice@example.com` . * `serviceAccount:{emailid}`: An email address that represents a Google service account. For example, `my-other-app@appspot.gserviceaccount.com`. * `serviceAccount:{projectid}.svc.id.goog[{namespace}/{kubernetes-sa}]`: An identifier for a [Kubernetes service account](https://cloud.google.com/kubernetes-engine/docs/how-to/kubernetes-service-accounts). For example, `my-project.svc.id.goog[my-namespace/my-kubernetes-sa]`. * `group:{emailid}`: An email address that represents a Google group. For example, `admins@example.com`. * `domain:{domain}`: The G Suite domain (primary) that represents all the users of that domain. For example, `google.com` or `example.com`. * `deleted:user:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a user that has been recently deleted. For example, `alice@example.com?uid=123456789012345678901`. If the user is recovered, this value reverts to `user:{emailid}` and the recovered user retains the role in the binding. * `deleted:serviceAccount:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, `my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901`. If the service account is undeleted, this value reverts to `serviceAccount:{emailid}` and the undeleted service account retains the role in the binding. * `deleted:group:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, `admins@example.com?uid=123456789012345678901`. If the group is recovered, this value reverts to `group:{emailid}` and the recovered group retains the role in the binding.
func (o BindingOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Binding) []string { return v.Members }).(pulumi.StringArrayOutput)
}

// Role that is assigned to the list of `members`, or principals. For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
func (o BindingOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Binding) *string { return v.Role }).(pulumi.StringPtrOutput)
}

type BindingArrayOutput struct{ *pulumi.OutputState }

func (BindingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Binding)(nil)).Elem()
}

func (o BindingArrayOutput) ToBindingArrayOutput() BindingArrayOutput {
	return o
}

func (o BindingArrayOutput) ToBindingArrayOutputWithContext(ctx context.Context) BindingArrayOutput {
	return o
}

func (o BindingArrayOutput) Index(i pulumi.IntInput) BindingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Binding {
		return vs[0].([]Binding)[vs[1].(int)]
	}).(BindingOutput)
}

// Associates `members`, or principals, with a `role`.
type BindingResponse struct {
	// The condition that is associated with this binding. If the condition evaluates to `true`, then this binding applies to the current request. If the condition evaluates to `false`, then this binding does not apply to the current request. However, a different role binding might grant the same role to one or more of the principals in this binding. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
	Condition ExprResponse `pulumi:"condition"`
	// Specifies the principals requesting access for a Google Cloud resource. `members` can have the following values: * `allUsers`: A special identifier that represents anyone who is on the internet; with or without a Google account. * `allAuthenticatedUsers`: A special identifier that represents anyone who is authenticated with a Google account or a service account. Does not include identities that come from external identity providers (IdPs) through identity federation. * `user:{emailid}`: An email address that represents a specific Google account. For example, `alice@example.com` . * `serviceAccount:{emailid}`: An email address that represents a Google service account. For example, `my-other-app@appspot.gserviceaccount.com`. * `serviceAccount:{projectid}.svc.id.goog[{namespace}/{kubernetes-sa}]`: An identifier for a [Kubernetes service account](https://cloud.google.com/kubernetes-engine/docs/how-to/kubernetes-service-accounts). For example, `my-project.svc.id.goog[my-namespace/my-kubernetes-sa]`. * `group:{emailid}`: An email address that represents a Google group. For example, `admins@example.com`. * `domain:{domain}`: The G Suite domain (primary) that represents all the users of that domain. For example, `google.com` or `example.com`. * `deleted:user:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a user that has been recently deleted. For example, `alice@example.com?uid=123456789012345678901`. If the user is recovered, this value reverts to `user:{emailid}` and the recovered user retains the role in the binding. * `deleted:serviceAccount:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, `my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901`. If the service account is undeleted, this value reverts to `serviceAccount:{emailid}` and the undeleted service account retains the role in the binding. * `deleted:group:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, `admins@example.com?uid=123456789012345678901`. If the group is recovered, this value reverts to `group:{emailid}` and the recovered group retains the role in the binding.
	Members []string `pulumi:"members"`
	// Role that is assigned to the list of `members`, or principals. For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
	Role string `pulumi:"role"`
}

// Associates `members`, or principals, with a `role`.
type BindingResponseOutput struct{ *pulumi.OutputState }

func (BindingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BindingResponse)(nil)).Elem()
}

func (o BindingResponseOutput) ToBindingResponseOutput() BindingResponseOutput {
	return o
}

func (o BindingResponseOutput) ToBindingResponseOutputWithContext(ctx context.Context) BindingResponseOutput {
	return o
}

// The condition that is associated with this binding. If the condition evaluates to `true`, then this binding applies to the current request. If the condition evaluates to `false`, then this binding does not apply to the current request. However, a different role binding might grant the same role to one or more of the principals in this binding. To learn which resources support conditions in their IAM policies, see the [IAM documentation](https://cloud.google.com/iam/help/conditions/resource-policies).
func (o BindingResponseOutput) Condition() ExprResponseOutput {
	return o.ApplyT(func(v BindingResponse) ExprResponse { return v.Condition }).(ExprResponseOutput)
}

// Specifies the principals requesting access for a Google Cloud resource. `members` can have the following values: * `allUsers`: A special identifier that represents anyone who is on the internet; with or without a Google account. * `allAuthenticatedUsers`: A special identifier that represents anyone who is authenticated with a Google account or a service account. Does not include identities that come from external identity providers (IdPs) through identity federation. * `user:{emailid}`: An email address that represents a specific Google account. For example, `alice@example.com` . * `serviceAccount:{emailid}`: An email address that represents a Google service account. For example, `my-other-app@appspot.gserviceaccount.com`. * `serviceAccount:{projectid}.svc.id.goog[{namespace}/{kubernetes-sa}]`: An identifier for a [Kubernetes service account](https://cloud.google.com/kubernetes-engine/docs/how-to/kubernetes-service-accounts). For example, `my-project.svc.id.goog[my-namespace/my-kubernetes-sa]`. * `group:{emailid}`: An email address that represents a Google group. For example, `admins@example.com`. * `domain:{domain}`: The G Suite domain (primary) that represents all the users of that domain. For example, `google.com` or `example.com`. * `deleted:user:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a user that has been recently deleted. For example, `alice@example.com?uid=123456789012345678901`. If the user is recovered, this value reverts to `user:{emailid}` and the recovered user retains the role in the binding. * `deleted:serviceAccount:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a service account that has been recently deleted. For example, `my-other-app@appspot.gserviceaccount.com?uid=123456789012345678901`. If the service account is undeleted, this value reverts to `serviceAccount:{emailid}` and the undeleted service account retains the role in the binding. * `deleted:group:{emailid}?uid={uniqueid}`: An email address (plus unique identifier) representing a Google group that has been recently deleted. For example, `admins@example.com?uid=123456789012345678901`. If the group is recovered, this value reverts to `group:{emailid}` and the recovered group retains the role in the binding.
func (o BindingResponseOutput) Members() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BindingResponse) []string { return v.Members }).(pulumi.StringArrayOutput)
}

// Role that is assigned to the list of `members`, or principals. For example, `roles/viewer`, `roles/editor`, or `roles/owner`.
func (o BindingResponseOutput) Role() pulumi.StringOutput {
	return o.ApplyT(func(v BindingResponse) string { return v.Role }).(pulumi.StringOutput)
}

type BindingResponseArrayOutput struct{ *pulumi.OutputState }

func (BindingResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BindingResponse)(nil)).Elem()
}

func (o BindingResponseArrayOutput) ToBindingResponseArrayOutput() BindingResponseArrayOutput {
	return o
}

func (o BindingResponseArrayOutput) ToBindingResponseArrayOutputWithContext(ctx context.Context) BindingResponseArrayOutput {
	return o
}

func (o BindingResponseArrayOutput) Index(i pulumi.IntInput) BindingResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BindingResponse {
		return vs[0].([]BindingResponse)[vs[1].(int)]
	}).(BindingResponseOutput)
}

// Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language. The syntax and semantics of CEL are documented at https://github.com/google/cel-spec. Example (Comparison): title: "Summary size limit" description: "Determines if a summary is less than 100 chars" expression: "document.summary.size() < 100" Example (Equality): title: "Requestor is owner" description: "Determines if requestor is the document owner" expression: "document.owner == request.auth.claims.email" Example (Logic): title: "Public documents" description: "Determine whether the document should be publicly visible" expression: "document.type != 'private' && document.type != 'internal'" Example (Data Manipulation): title: "Notification string" description: "Create a notification string with a timestamp." expression: "'New message received at ' + string(document.create_time)" The exact variables and functions that may be referenced within an expression are determined by the service that evaluates it. See the service documentation for additional information.
type Expr struct {
	// Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
	Description *string `pulumi:"description"`
	// Textual representation of an expression in Common Expression Language syntax.
	Expression *string `pulumi:"expression"`
	// Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
	Location *string `pulumi:"location"`
	// Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
	Title *string `pulumi:"title"`
}

// ExprInput is an input type that accepts ExprArgs and ExprOutput values.
// You can construct a concrete instance of `ExprInput` via:
//
//	ExprArgs{...}
type ExprInput interface {
	pulumi.Input

	ToExprOutput() ExprOutput
	ToExprOutputWithContext(context.Context) ExprOutput
}

// Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language. The syntax and semantics of CEL are documented at https://github.com/google/cel-spec. Example (Comparison): title: "Summary size limit" description: "Determines if a summary is less than 100 chars" expression: "document.summary.size() < 100" Example (Equality): title: "Requestor is owner" description: "Determines if requestor is the document owner" expression: "document.owner == request.auth.claims.email" Example (Logic): title: "Public documents" description: "Determine whether the document should be publicly visible" expression: "document.type != 'private' && document.type != 'internal'" Example (Data Manipulation): title: "Notification string" description: "Create a notification string with a timestamp." expression: "'New message received at ' + string(document.create_time)" The exact variables and functions that may be referenced within an expression are determined by the service that evaluates it. See the service documentation for additional information.
type ExprArgs struct {
	// Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// Textual representation of an expression in Common Expression Language syntax.
	Expression pulumi.StringPtrInput `pulumi:"expression"`
	// Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
	Location pulumi.StringPtrInput `pulumi:"location"`
	// Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
	Title pulumi.StringPtrInput `pulumi:"title"`
}

func (ExprArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Expr)(nil)).Elem()
}

func (i ExprArgs) ToExprOutput() ExprOutput {
	return i.ToExprOutputWithContext(context.Background())
}

func (i ExprArgs) ToExprOutputWithContext(ctx context.Context) ExprOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExprOutput)
}

func (i ExprArgs) ToExprPtrOutput() ExprPtrOutput {
	return i.ToExprPtrOutputWithContext(context.Background())
}

func (i ExprArgs) ToExprPtrOutputWithContext(ctx context.Context) ExprPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExprOutput).ToExprPtrOutputWithContext(ctx)
}

// ExprPtrInput is an input type that accepts ExprArgs, ExprPtr and ExprPtrOutput values.
// You can construct a concrete instance of `ExprPtrInput` via:
//
//	        ExprArgs{...}
//
//	or:
//
//	        nil
type ExprPtrInput interface {
	pulumi.Input

	ToExprPtrOutput() ExprPtrOutput
	ToExprPtrOutputWithContext(context.Context) ExprPtrOutput
}

type exprPtrType ExprArgs

func ExprPtr(v *ExprArgs) ExprPtrInput {
	return (*exprPtrType)(v)
}

func (*exprPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Expr)(nil)).Elem()
}

func (i *exprPtrType) ToExprPtrOutput() ExprPtrOutput {
	return i.ToExprPtrOutputWithContext(context.Background())
}

func (i *exprPtrType) ToExprPtrOutputWithContext(ctx context.Context) ExprPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExprPtrOutput)
}

// Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language. The syntax and semantics of CEL are documented at https://github.com/google/cel-spec. Example (Comparison): title: "Summary size limit" description: "Determines if a summary is less than 100 chars" expression: "document.summary.size() < 100" Example (Equality): title: "Requestor is owner" description: "Determines if requestor is the document owner" expression: "document.owner == request.auth.claims.email" Example (Logic): title: "Public documents" description: "Determine whether the document should be publicly visible" expression: "document.type != 'private' && document.type != 'internal'" Example (Data Manipulation): title: "Notification string" description: "Create a notification string with a timestamp." expression: "'New message received at ' + string(document.create_time)" The exact variables and functions that may be referenced within an expression are determined by the service that evaluates it. See the service documentation for additional information.
type ExprOutput struct{ *pulumi.OutputState }

func (ExprOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Expr)(nil)).Elem()
}

func (o ExprOutput) ToExprOutput() ExprOutput {
	return o
}

func (o ExprOutput) ToExprOutputWithContext(ctx context.Context) ExprOutput {
	return o
}

func (o ExprOutput) ToExprPtrOutput() ExprPtrOutput {
	return o.ToExprPtrOutputWithContext(context.Background())
}

func (o ExprOutput) ToExprPtrOutputWithContext(ctx context.Context) ExprPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Expr) *Expr {
		return &v
	}).(ExprPtrOutput)
}

// Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
func (o ExprOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Expr) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Textual representation of an expression in Common Expression Language syntax.
func (o ExprOutput) Expression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Expr) *string { return v.Expression }).(pulumi.StringPtrOutput)
}

// Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
func (o ExprOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Expr) *string { return v.Location }).(pulumi.StringPtrOutput)
}

// Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
func (o ExprOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Expr) *string { return v.Title }).(pulumi.StringPtrOutput)
}

type ExprPtrOutput struct{ *pulumi.OutputState }

func (ExprPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Expr)(nil)).Elem()
}

func (o ExprPtrOutput) ToExprPtrOutput() ExprPtrOutput {
	return o
}

func (o ExprPtrOutput) ToExprPtrOutputWithContext(ctx context.Context) ExprPtrOutput {
	return o
}

func (o ExprPtrOutput) Elem() ExprOutput {
	return o.ApplyT(func(v *Expr) Expr {
		if v != nil {
			return *v
		}
		var ret Expr
		return ret
	}).(ExprOutput)
}

// Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
func (o ExprPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Expr) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

// Textual representation of an expression in Common Expression Language syntax.
func (o ExprPtrOutput) Expression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Expr) *string {
		if v == nil {
			return nil
		}
		return v.Expression
	}).(pulumi.StringPtrOutput)
}

// Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
func (o ExprPtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Expr) *string {
		if v == nil {
			return nil
		}
		return v.Location
	}).(pulumi.StringPtrOutput)
}

// Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
func (o ExprPtrOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Expr) *string {
		if v == nil {
			return nil
		}
		return v.Title
	}).(pulumi.StringPtrOutput)
}

// Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language. The syntax and semantics of CEL are documented at https://github.com/google/cel-spec. Example (Comparison): title: "Summary size limit" description: "Determines if a summary is less than 100 chars" expression: "document.summary.size() < 100" Example (Equality): title: "Requestor is owner" description: "Determines if requestor is the document owner" expression: "document.owner == request.auth.claims.email" Example (Logic): title: "Public documents" description: "Determine whether the document should be publicly visible" expression: "document.type != 'private' && document.type != 'internal'" Example (Data Manipulation): title: "Notification string" description: "Create a notification string with a timestamp." expression: "'New message received at ' + string(document.create_time)" The exact variables and functions that may be referenced within an expression are determined by the service that evaluates it. See the service documentation for additional information.
type ExprResponse struct {
	// Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
	Description string `pulumi:"description"`
	// Textual representation of an expression in Common Expression Language syntax.
	Expression string `pulumi:"expression"`
	// Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
	Location string `pulumi:"location"`
	// Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
	Title string `pulumi:"title"`
}

// Represents a textual expression in the Common Expression Language (CEL) syntax. CEL is a C-like expression language. The syntax and semantics of CEL are documented at https://github.com/google/cel-spec. Example (Comparison): title: "Summary size limit" description: "Determines if a summary is less than 100 chars" expression: "document.summary.size() < 100" Example (Equality): title: "Requestor is owner" description: "Determines if requestor is the document owner" expression: "document.owner == request.auth.claims.email" Example (Logic): title: "Public documents" description: "Determine whether the document should be publicly visible" expression: "document.type != 'private' && document.type != 'internal'" Example (Data Manipulation): title: "Notification string" description: "Create a notification string with a timestamp." expression: "'New message received at ' + string(document.create_time)" The exact variables and functions that may be referenced within an expression are determined by the service that evaluates it. See the service documentation for additional information.
type ExprResponseOutput struct{ *pulumi.OutputState }

func (ExprResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExprResponse)(nil)).Elem()
}

func (o ExprResponseOutput) ToExprResponseOutput() ExprResponseOutput {
	return o
}

func (o ExprResponseOutput) ToExprResponseOutputWithContext(ctx context.Context) ExprResponseOutput {
	return o
}

// Optional. Description of the expression. This is a longer text which describes the expression, e.g. when hovered over it in a UI.
func (o ExprResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v ExprResponse) string { return v.Description }).(pulumi.StringOutput)
}

// Textual representation of an expression in Common Expression Language syntax.
func (o ExprResponseOutput) Expression() pulumi.StringOutput {
	return o.ApplyT(func(v ExprResponse) string { return v.Expression }).(pulumi.StringOutput)
}

// Optional. String indicating the location of the expression for error reporting, e.g. a file name and a position in the file.
func (o ExprResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v ExprResponse) string { return v.Location }).(pulumi.StringOutput)
}

// Optional. Title for the expression, i.e. a short string describing its purpose. This can be used e.g. in UIs which allow to enter the expression.
func (o ExprResponseOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v ExprResponse) string { return v.Title }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BindingInput)(nil)).Elem(), BindingArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BindingArrayInput)(nil)).Elem(), BindingArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExprInput)(nil)).Elem(), ExprArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExprPtrInput)(nil)).Elem(), ExprArgs{})
	pulumi.RegisterOutputType(BindingOutput{})
	pulumi.RegisterOutputType(BindingArrayOutput{})
	pulumi.RegisterOutputType(BindingResponseOutput{})
	pulumi.RegisterOutputType(BindingResponseArrayOutput{})
	pulumi.RegisterOutputType(ExprOutput{})
	pulumi.RegisterOutputType(ExprPtrOutput{})
	pulumi.RegisterOutputType(ExprResponseOutput{})
}
