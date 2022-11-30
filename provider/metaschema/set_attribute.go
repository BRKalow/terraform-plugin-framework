package metaschema

import (
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/internal/fwschema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// Ensure the implementation satisifies the desired interfaces.
var (
	_ Attribute = SetAttribute{}
)

// SetAttribute represents a schema attribute that is a set with a single
// element type. When retrieving the value for this attribute, use types.Set
// as the value type unless the CustomType field is set. The ElementType field
// must be set.
//
// Use SetNestedAttribute if the underlying elements should be objects and
// require definition beyond type information.
//
// Terraform configurations configure this attribute using expressions that
// return a set or directly via square brace syntax.
//
//	# set of strings
//	example_attribute = ["first", "second"]
//
// Terraform configurations reference this attribute using expressions that
// accept a set. Sets cannot be indexed in Terraform, therefore an expression
// is required to access an explicit element.
type SetAttribute struct {
	// ElementType is the type for all elements of the set. This field must be
	// set.
	ElementType attr.Type

	// CustomType enables the use of a custom attribute type in place of the
	// default types.SetType. When retrieving data, the types.SetValuable
	// associated with this custom type must be used in place of types.Set.
	CustomType types.SetTypable

	// Required indicates whether the practitioner must enter a value for
	// this attribute or not. Required and Optional cannot both be true,
	// and Required and Computed cannot both be true.
	Required bool

	// Optional indicates whether the practitioner can choose to enter a value
	// for this attribute or not. Optional and Required cannot both be true.
	Optional bool

	// Description is used in various tooling, like the language server, to
	// give practitioners more information about what this attribute is,
	// what it's for, and how it should be used. It should be written as
	// plain text, with no special formatting.
	Description string

	// MarkdownDescription is used in various tooling, like the
	// documentation generator, to give practitioners more information
	// about what this attribute is, what it's for, and how it should be
	// used. It should be formatted using Markdown.
	MarkdownDescription string
}

// ApplyTerraform5AttributePathStep returns the result of stepping into a set
// index or an error.
func (a SetAttribute) ApplyTerraform5AttributePathStep(step tftypes.AttributePathStep) (interface{}, error) {
	return a.GetType().ApplyTerraform5AttributePathStep(step)
}

// Equal returns true if the given Attribute is a SetAttribute
// and all fields are equal.
func (a SetAttribute) Equal(o fwschema.Attribute) bool {
	if _, ok := o.(SetAttribute); !ok {
		return false
	}

	return fwschema.AttributesEqual(a, o)
}

// GetDeprecationMessage always returns an empty string as there is no
// deprecation validation support for provider meta schemas.
func (a SetAttribute) GetDeprecationMessage() string {
	return ""
}

// GetDescription returns the Description field value.
func (a SetAttribute) GetDescription() string {
	return a.Description
}

// GetMarkdownDescription returns the MarkdownDescription field value.
func (a SetAttribute) GetMarkdownDescription() string {
	return a.MarkdownDescription
}

// GetType returns types.SetType or the CustomType field value if defined.
func (a SetAttribute) GetType() attr.Type {
	if a.CustomType != nil {
		return a.CustomType
	}

	return types.SetType{
		ElemType: a.ElementType,
	}
}

// IsComputed always returns false as provider schemas cannot be Computed.
func (a SetAttribute) IsComputed() bool {
	return false
}

// IsOptional returns the Optional field value.
func (a SetAttribute) IsOptional() bool {
	return a.Optional
}

// IsRequired returns the Required field value.
func (a SetAttribute) IsRequired() bool {
	return a.Required
}

// IsSensitive always returns false as there is no plan for provider meta
// schema data.
func (a SetAttribute) IsSensitive() bool {
	return false
}