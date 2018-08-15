package builder

import (
	"github.com/turbonomic/turbo-go-sdk/pkg/proto"
)

type AccountValueConstraint string

const (
	Mandatory AccountValueConstraint = "MANDATORY"
	Optional  AccountValueConstraint = "OPTIONAL"
)

type VerificationRegex string

const (
	DEFAULT_STRING_REGEXP VerificationRegex = ".*"
)

type AccountFieldMeta interface {
	TargetId() bool
	DisplayName() string
	Description() string
	VerificationRegexp() string
	Secret() bool
	ScopeType() proto.EntityDTO_EntityType
	DefaultValue() string
	Constraint() AccountValueConstraint
}

type AccountField struct {
	// Metadata that will be converted to AccountValueEntry
	displayName       string
	description       string
	targetId          bool
	secret            bool
	verificationRegex string
	constraint        AccountValueConstraint
	defaultValue      string

	// Actual Value
	value interface{} //field type determined using reflection
}

func (userfield *AccountField) TargetId() bool {
	return userfield.targetId
}

func (userfield *AccountField) DisplayName() string {
	return userfield.displayName
}
func (userfield *AccountField) Description() string {
	return userfield.description
}

func (userfield *AccountField) VerificationRegexp() string {
	if userfield.verificationRegex == "" {
		return string(DEFAULT_STRING_REGEXP)
	} else {
		return userfield.verificationRegex
	}
}

func (userfield *AccountField) Secret() bool {
	return userfield.secret //default false
}

func (userfield *AccountField) ScopeType() proto.EntityDTO_EntityType {
	return proto.EntityDTO_VIRTUAL_MACHINE
}

func (userfield *AccountField) DefaultValue() string {
	return ""
}

func (userfield *AccountField) Constraint() AccountValueConstraint {
	if userfield.constraint != "" {
		return userfield.constraint
	}
	return Mandatory
}

type PredefinedAccountDefinition struct {
	//constraint AccountValueConstraint
	//defaultValue string
	displayName string
	description string
	secret      bool
	fieldType   FieldType
}

type FieldType string

var (
	STRING  FieldType = "String"
	BOOLEAN FieldType = "Boolean"
	Numeric FieldType = "Numeric"
)

var (
	Username PredefinedAccountDefinition = PredefinedAccountDefinition{
		displayName: "User name", description: "User name to connect to target with",
		secret: false, fieldType: STRING}
	Password PredefinedAccountDefinition = PredefinedAccountDefinition{
		displayName: "Password", description: "Password to use to connect to target",
		secret: true, fieldType: STRING}
	Address PredefinedAccountDefinition = PredefinedAccountDefinition{
		displayName: "TargetId", description: "Address of a target",
		secret: true, fieldType: STRING}
)

type CustomAccountDefinition struct {
	displayName       string
	description       string
	secret            bool
	fieldType         string
	verificationRegex string
}
