// Code generated by the Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package v2

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi-google-native/sdk/go/google/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Creates a new routine in the dataset.
// Auto-naming is currently not supported for this resource.
type Routine struct {
	pulumi.CustomResourceState

	// Optional.
	Arguments ArgumentResponseArrayOutput `pulumi:"arguments"`
	// The time when this routine was created, in milliseconds since the epoch.
	CreationTime pulumi.StringOutput `pulumi:"creationTime"`
	DatasetId    pulumi.StringOutput `pulumi:"datasetId"`
	// The body of the routine. For functions, this is the expression in the AS clause. If language=SQL, it is the substring inside (but excluding) the parentheses. For example, for the function created with the following statement: `CREATE FUNCTION JoinLines(x string, y string) as (concat(x, "\n", y))` The definition_body is `concat(x, "\n", y)` (\n is not replaced with linebreak). If language=JAVASCRIPT, it is the evaluated string in the AS clause. For example, for the function created with the following statement: `CREATE FUNCTION f() RETURNS STRING LANGUAGE js AS 'return "\n";\n'` The definition_body is `return "\n";\n` Note that both \n are replaced with linebreaks.
	DefinitionBody pulumi.StringOutput `pulumi:"definitionBody"`
	// Optional. The description of the routine, if defined.
	Description pulumi.StringOutput `pulumi:"description"`
	// Optional. The determinism level of the JavaScript UDF, if defined.
	DeterminismLevel pulumi.StringOutput `pulumi:"determinismLevel"`
	// A hash of this resource.
	Etag pulumi.StringOutput `pulumi:"etag"`
	// Optional. If language = "JAVASCRIPT", this field stores the path of the imported JAVASCRIPT libraries.
	ImportedLibraries pulumi.StringArrayOutput `pulumi:"importedLibraries"`
	// Optional. Defaults to "SQL" if remote_function_options field is absent, not set otherwise.
	Language pulumi.StringOutput `pulumi:"language"`
	// The time when this routine was last modified, in milliseconds since the epoch.
	LastModifiedTime pulumi.StringOutput `pulumi:"lastModifiedTime"`
	Project          pulumi.StringOutput `pulumi:"project"`
	// Optional. Remote function specific options.
	RemoteFunctionOptions RemoteFunctionOptionsResponseOutput `pulumi:"remoteFunctionOptions"`
	// Optional. Can be set only if routine_type = "TABLE_VALUED_FUNCTION". If absent, the return table type is inferred from definition_body at query time in each query that references this routine. If present, then the columns in the evaluated table result will be cast to match the column types specified in return table type, at query time.
	ReturnTableType StandardSqlTableTypeResponseOutput `pulumi:"returnTableType"`
	// Optional if language = "SQL"; required otherwise. Cannot be set if routine_type = "TABLE_VALUED_FUNCTION". If absent, the return type is inferred from definition_body at query time in each query that references this routine. If present, then the evaluated result will be cast to the specified returned type at query time. For example, for the functions created with the following statements: * `CREATE FUNCTION Add(x FLOAT64, y FLOAT64) RETURNS FLOAT64 AS (x + y);` * `CREATE FUNCTION Increment(x FLOAT64) AS (Add(x, 1));` * `CREATE FUNCTION Decrement(x FLOAT64) RETURNS FLOAT64 AS (Add(x, -1));` The return_type is `{type_kind: "FLOAT64"}` for `Add` and `Decrement`, and is absent for `Increment` (inferred as FLOAT64 at query time). Suppose the function `Add` is replaced by `CREATE OR REPLACE FUNCTION Add(x INT64, y INT64) AS (x + y);` Then the inferred return type of `Increment` is automatically changed to INT64 at query time, while the return type of `Decrement` remains FLOAT64.
	ReturnType StandardSqlDataTypeResponseOutput `pulumi:"returnType"`
	// Reference describing the ID of this routine.
	RoutineReference RoutineReferenceResponseOutput `pulumi:"routineReference"`
	// The type of routine.
	RoutineType pulumi.StringOutput `pulumi:"routineType"`
	// Optional. Spark specific options.
	SparkOptions SparkOptionsResponseOutput `pulumi:"sparkOptions"`
	// Optional. Can be set for procedures only. If true (default), the definition body will be validated in the creation and the updates of the procedure. For procedures with an argument of ANY TYPE, the definition body validtion is not supported at creation/update time, and thus this field must be set to false explicitly.
	StrictMode pulumi.BoolOutput `pulumi:"strictMode"`
}

// NewRoutine registers a new resource with the given unique name, arguments, and options.
func NewRoutine(ctx *pulumi.Context,
	name string, args *RoutineArgs, opts ...pulumi.ResourceOption) (*Routine, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DatasetId == nil {
		return nil, errors.New("invalid value for required argument 'DatasetId'")
	}
	if args.DefinitionBody == nil {
		return nil, errors.New("invalid value for required argument 'DefinitionBody'")
	}
	if args.RoutineReference == nil {
		return nil, errors.New("invalid value for required argument 'RoutineReference'")
	}
	if args.RoutineType == nil {
		return nil, errors.New("invalid value for required argument 'RoutineType'")
	}
	replaceOnChanges := pulumi.ReplaceOnChanges([]string{
		"datasetId",
		"project",
	})
	opts = append(opts, replaceOnChanges)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Routine
	err := ctx.RegisterResource("google-native:bigquery/v2:Routine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoutine gets an existing Routine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoutine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RoutineState, opts ...pulumi.ResourceOption) (*Routine, error) {
	var resource Routine
	err := ctx.ReadResource("google-native:bigquery/v2:Routine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Routine resources.
type routineState struct {
}

type RoutineState struct {
}

func (RoutineState) ElementType() reflect.Type {
	return reflect.TypeOf((*routineState)(nil)).Elem()
}

type routineArgs struct {
	// Optional.
	Arguments []Argument `pulumi:"arguments"`
	DatasetId string     `pulumi:"datasetId"`
	// The body of the routine. For functions, this is the expression in the AS clause. If language=SQL, it is the substring inside (but excluding) the parentheses. For example, for the function created with the following statement: `CREATE FUNCTION JoinLines(x string, y string) as (concat(x, "\n", y))` The definition_body is `concat(x, "\n", y)` (\n is not replaced with linebreak). If language=JAVASCRIPT, it is the evaluated string in the AS clause. For example, for the function created with the following statement: `CREATE FUNCTION f() RETURNS STRING LANGUAGE js AS 'return "\n";\n'` The definition_body is `return "\n";\n` Note that both \n are replaced with linebreaks.
	DefinitionBody string `pulumi:"definitionBody"`
	// Optional. The description of the routine, if defined.
	Description *string `pulumi:"description"`
	// Optional. The determinism level of the JavaScript UDF, if defined.
	DeterminismLevel *RoutineDeterminismLevel `pulumi:"determinismLevel"`
	// Optional. If language = "JAVASCRIPT", this field stores the path of the imported JAVASCRIPT libraries.
	ImportedLibraries []string `pulumi:"importedLibraries"`
	// Optional. Defaults to "SQL" if remote_function_options field is absent, not set otherwise.
	Language *RoutineLanguage `pulumi:"language"`
	Project  *string          `pulumi:"project"`
	// Optional. Remote function specific options.
	RemoteFunctionOptions *RemoteFunctionOptions `pulumi:"remoteFunctionOptions"`
	// Optional. Can be set only if routine_type = "TABLE_VALUED_FUNCTION". If absent, the return table type is inferred from definition_body at query time in each query that references this routine. If present, then the columns in the evaluated table result will be cast to match the column types specified in return table type, at query time.
	ReturnTableType *StandardSqlTableType `pulumi:"returnTableType"`
	// Optional if language = "SQL"; required otherwise. Cannot be set if routine_type = "TABLE_VALUED_FUNCTION". If absent, the return type is inferred from definition_body at query time in each query that references this routine. If present, then the evaluated result will be cast to the specified returned type at query time. For example, for the functions created with the following statements: * `CREATE FUNCTION Add(x FLOAT64, y FLOAT64) RETURNS FLOAT64 AS (x + y);` * `CREATE FUNCTION Increment(x FLOAT64) AS (Add(x, 1));` * `CREATE FUNCTION Decrement(x FLOAT64) RETURNS FLOAT64 AS (Add(x, -1));` The return_type is `{type_kind: "FLOAT64"}` for `Add` and `Decrement`, and is absent for `Increment` (inferred as FLOAT64 at query time). Suppose the function `Add` is replaced by `CREATE OR REPLACE FUNCTION Add(x INT64, y INT64) AS (x + y);` Then the inferred return type of `Increment` is automatically changed to INT64 at query time, while the return type of `Decrement` remains FLOAT64.
	ReturnType *StandardSqlDataType `pulumi:"returnType"`
	// Reference describing the ID of this routine.
	RoutineReference RoutineReference `pulumi:"routineReference"`
	// The type of routine.
	RoutineType RoutineRoutineType `pulumi:"routineType"`
	// Optional. Spark specific options.
	SparkOptions *SparkOptions `pulumi:"sparkOptions"`
	// Optional. Can be set for procedures only. If true (default), the definition body will be validated in the creation and the updates of the procedure. For procedures with an argument of ANY TYPE, the definition body validtion is not supported at creation/update time, and thus this field must be set to false explicitly.
	StrictMode *bool `pulumi:"strictMode"`
}

// The set of arguments for constructing a Routine resource.
type RoutineArgs struct {
	// Optional.
	Arguments ArgumentArrayInput
	DatasetId pulumi.StringInput
	// The body of the routine. For functions, this is the expression in the AS clause. If language=SQL, it is the substring inside (but excluding) the parentheses. For example, for the function created with the following statement: `CREATE FUNCTION JoinLines(x string, y string) as (concat(x, "\n", y))` The definition_body is `concat(x, "\n", y)` (\n is not replaced with linebreak). If language=JAVASCRIPT, it is the evaluated string in the AS clause. For example, for the function created with the following statement: `CREATE FUNCTION f() RETURNS STRING LANGUAGE js AS 'return "\n";\n'` The definition_body is `return "\n";\n` Note that both \n are replaced with linebreaks.
	DefinitionBody pulumi.StringInput
	// Optional. The description of the routine, if defined.
	Description pulumi.StringPtrInput
	// Optional. The determinism level of the JavaScript UDF, if defined.
	DeterminismLevel RoutineDeterminismLevelPtrInput
	// Optional. If language = "JAVASCRIPT", this field stores the path of the imported JAVASCRIPT libraries.
	ImportedLibraries pulumi.StringArrayInput
	// Optional. Defaults to "SQL" if remote_function_options field is absent, not set otherwise.
	Language RoutineLanguagePtrInput
	Project  pulumi.StringPtrInput
	// Optional. Remote function specific options.
	RemoteFunctionOptions RemoteFunctionOptionsPtrInput
	// Optional. Can be set only if routine_type = "TABLE_VALUED_FUNCTION". If absent, the return table type is inferred from definition_body at query time in each query that references this routine. If present, then the columns in the evaluated table result will be cast to match the column types specified in return table type, at query time.
	ReturnTableType StandardSqlTableTypePtrInput
	// Optional if language = "SQL"; required otherwise. Cannot be set if routine_type = "TABLE_VALUED_FUNCTION". If absent, the return type is inferred from definition_body at query time in each query that references this routine. If present, then the evaluated result will be cast to the specified returned type at query time. For example, for the functions created with the following statements: * `CREATE FUNCTION Add(x FLOAT64, y FLOAT64) RETURNS FLOAT64 AS (x + y);` * `CREATE FUNCTION Increment(x FLOAT64) AS (Add(x, 1));` * `CREATE FUNCTION Decrement(x FLOAT64) RETURNS FLOAT64 AS (Add(x, -1));` The return_type is `{type_kind: "FLOAT64"}` for `Add` and `Decrement`, and is absent for `Increment` (inferred as FLOAT64 at query time). Suppose the function `Add` is replaced by `CREATE OR REPLACE FUNCTION Add(x INT64, y INT64) AS (x + y);` Then the inferred return type of `Increment` is automatically changed to INT64 at query time, while the return type of `Decrement` remains FLOAT64.
	ReturnType StandardSqlDataTypePtrInput
	// Reference describing the ID of this routine.
	RoutineReference RoutineReferenceInput
	// The type of routine.
	RoutineType RoutineRoutineTypeInput
	// Optional. Spark specific options.
	SparkOptions SparkOptionsPtrInput
	// Optional. Can be set for procedures only. If true (default), the definition body will be validated in the creation and the updates of the procedure. For procedures with an argument of ANY TYPE, the definition body validtion is not supported at creation/update time, and thus this field must be set to false explicitly.
	StrictMode pulumi.BoolPtrInput
}

func (RoutineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routineArgs)(nil)).Elem()
}

type RoutineInput interface {
	pulumi.Input

	ToRoutineOutput() RoutineOutput
	ToRoutineOutputWithContext(ctx context.Context) RoutineOutput
}

func (*Routine) ElementType() reflect.Type {
	return reflect.TypeOf((**Routine)(nil)).Elem()
}

func (i *Routine) ToRoutineOutput() RoutineOutput {
	return i.ToRoutineOutputWithContext(context.Background())
}

func (i *Routine) ToRoutineOutputWithContext(ctx context.Context) RoutineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutineOutput)
}

type RoutineOutput struct{ *pulumi.OutputState }

func (RoutineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Routine)(nil)).Elem()
}

func (o RoutineOutput) ToRoutineOutput() RoutineOutput {
	return o
}

func (o RoutineOutput) ToRoutineOutputWithContext(ctx context.Context) RoutineOutput {
	return o
}

// Optional.
func (o RoutineOutput) Arguments() ArgumentResponseArrayOutput {
	return o.ApplyT(func(v *Routine) ArgumentResponseArrayOutput { return v.Arguments }).(ArgumentResponseArrayOutput)
}

// The time when this routine was created, in milliseconds since the epoch.
func (o RoutineOutput) CreationTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Routine) pulumi.StringOutput { return v.CreationTime }).(pulumi.StringOutput)
}

func (o RoutineOutput) DatasetId() pulumi.StringOutput {
	return o.ApplyT(func(v *Routine) pulumi.StringOutput { return v.DatasetId }).(pulumi.StringOutput)
}

// The body of the routine. For functions, this is the expression in the AS clause. If language=SQL, it is the substring inside (but excluding) the parentheses. For example, for the function created with the following statement: `CREATE FUNCTION JoinLines(x string, y string) as (concat(x, "\n", y))` The definition_body is `concat(x, "\n", y)` (\n is not replaced with linebreak). If language=JAVASCRIPT, it is the evaluated string in the AS clause. For example, for the function created with the following statement: `CREATE FUNCTION f() RETURNS STRING LANGUAGE js AS 'return "\n";\n'` The definition_body is `return "\n";\n` Note that both \n are replaced with linebreaks.
func (o RoutineOutput) DefinitionBody() pulumi.StringOutput {
	return o.ApplyT(func(v *Routine) pulumi.StringOutput { return v.DefinitionBody }).(pulumi.StringOutput)
}

// Optional. The description of the routine, if defined.
func (o RoutineOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Routine) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// Optional. The determinism level of the JavaScript UDF, if defined.
func (o RoutineOutput) DeterminismLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *Routine) pulumi.StringOutput { return v.DeterminismLevel }).(pulumi.StringOutput)
}

// A hash of this resource.
func (o RoutineOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Routine) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

// Optional. If language = "JAVASCRIPT", this field stores the path of the imported JAVASCRIPT libraries.
func (o RoutineOutput) ImportedLibraries() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Routine) pulumi.StringArrayOutput { return v.ImportedLibraries }).(pulumi.StringArrayOutput)
}

// Optional. Defaults to "SQL" if remote_function_options field is absent, not set otherwise.
func (o RoutineOutput) Language() pulumi.StringOutput {
	return o.ApplyT(func(v *Routine) pulumi.StringOutput { return v.Language }).(pulumi.StringOutput)
}

// The time when this routine was last modified, in milliseconds since the epoch.
func (o RoutineOutput) LastModifiedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Routine) pulumi.StringOutput { return v.LastModifiedTime }).(pulumi.StringOutput)
}

func (o RoutineOutput) Project() pulumi.StringOutput {
	return o.ApplyT(func(v *Routine) pulumi.StringOutput { return v.Project }).(pulumi.StringOutput)
}

// Optional. Remote function specific options.
func (o RoutineOutput) RemoteFunctionOptions() RemoteFunctionOptionsResponseOutput {
	return o.ApplyT(func(v *Routine) RemoteFunctionOptionsResponseOutput { return v.RemoteFunctionOptions }).(RemoteFunctionOptionsResponseOutput)
}

// Optional. Can be set only if routine_type = "TABLE_VALUED_FUNCTION". If absent, the return table type is inferred from definition_body at query time in each query that references this routine. If present, then the columns in the evaluated table result will be cast to match the column types specified in return table type, at query time.
func (o RoutineOutput) ReturnTableType() StandardSqlTableTypeResponseOutput {
	return o.ApplyT(func(v *Routine) StandardSqlTableTypeResponseOutput { return v.ReturnTableType }).(StandardSqlTableTypeResponseOutput)
}

// Optional if language = "SQL"; required otherwise. Cannot be set if routine_type = "TABLE_VALUED_FUNCTION". If absent, the return type is inferred from definition_body at query time in each query that references this routine. If present, then the evaluated result will be cast to the specified returned type at query time. For example, for the functions created with the following statements: * `CREATE FUNCTION Add(x FLOAT64, y FLOAT64) RETURNS FLOAT64 AS (x + y);` * `CREATE FUNCTION Increment(x FLOAT64) AS (Add(x, 1));` * `CREATE FUNCTION Decrement(x FLOAT64) RETURNS FLOAT64 AS (Add(x, -1));` The return_type is `{type_kind: "FLOAT64"}` for `Add` and `Decrement`, and is absent for `Increment` (inferred as FLOAT64 at query time). Suppose the function `Add` is replaced by `CREATE OR REPLACE FUNCTION Add(x INT64, y INT64) AS (x + y);` Then the inferred return type of `Increment` is automatically changed to INT64 at query time, while the return type of `Decrement` remains FLOAT64.
func (o RoutineOutput) ReturnType() StandardSqlDataTypeResponseOutput {
	return o.ApplyT(func(v *Routine) StandardSqlDataTypeResponseOutput { return v.ReturnType }).(StandardSqlDataTypeResponseOutput)
}

// Reference describing the ID of this routine.
func (o RoutineOutput) RoutineReference() RoutineReferenceResponseOutput {
	return o.ApplyT(func(v *Routine) RoutineReferenceResponseOutput { return v.RoutineReference }).(RoutineReferenceResponseOutput)
}

// The type of routine.
func (o RoutineOutput) RoutineType() pulumi.StringOutput {
	return o.ApplyT(func(v *Routine) pulumi.StringOutput { return v.RoutineType }).(pulumi.StringOutput)
}

// Optional. Spark specific options.
func (o RoutineOutput) SparkOptions() SparkOptionsResponseOutput {
	return o.ApplyT(func(v *Routine) SparkOptionsResponseOutput { return v.SparkOptions }).(SparkOptionsResponseOutput)
}

// Optional. Can be set for procedures only. If true (default), the definition body will be validated in the creation and the updates of the procedure. For procedures with an argument of ANY TYPE, the definition body validtion is not supported at creation/update time, and thus this field must be set to false explicitly.
func (o RoutineOutput) StrictMode() pulumi.BoolOutput {
	return o.ApplyT(func(v *Routine) pulumi.BoolOutput { return v.StrictMode }).(pulumi.BoolOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*RoutineInput)(nil)).Elem(), &Routine{})
	pulumi.RegisterOutputType(RoutineOutput{})
}
