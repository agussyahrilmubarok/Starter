package validator_test

import (
	"testing"

	customValidator "agussyahrilmubarok.github.io/web/pkg/validator"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
)

type testStruct struct {
	Name  string `json:"name"  validate:"required,min=2,max=100"`
	Email string `json:"email" validate:"required,email"`
	Age   string `json:"age"   validate:"numeric"`
	Bio   string `json:"bio"   validate:"max=200"`
}

type testStructDefault struct {
	Code string `json:"code" validate:"alphanum"`
}

type testStructUnique struct {
	Email string `json:"email" validate:"unique"`
}

var validate = validator.New()

func init() {
	_ = validate.RegisterValidation("unique", func(fl validator.FieldLevel) bool {
		return false
	})
}

func getValidationError(s any) error {
	return validate.Struct(s)
}

func TestParseError_Required(t *testing.T) {
	err := getValidationError(testStruct{
		Age: "25",
	})

	result := customValidator.ParseError(err)

	assert.Equal(t, "Name is required", result["name"])
	assert.Equal(t, "Email is required", result["email"])
}

func TestParseError_Email(t *testing.T) {
	err := getValidationError(testStruct{
		Name:  "Alice",
		Email: "not-an-email",
		Age:   "25",
	})

	result := customValidator.ParseError(err)

	assert.Equal(t, "Invalid email format", result["email"])
}

func TestParseError_Min(t *testing.T) {
	err := getValidationError(testStruct{
		Name:  "A",
		Email: "alice@mail.com",
		Age:   "25",
	})

	result := customValidator.ParseError(err)

	assert.Equal(t, "Name must be at least 2 characters", result["name"])
}

func TestParseError_Max(t *testing.T) {
	err := getValidationError(testStruct{
		Name:  "Alice",
		Email: "alice@mail.com",
		Age:   "25",
		Bio:   string(make([]byte, 201)),
	})

	result := customValidator.ParseError(err)

	assert.Equal(t, "Bio must be at most 200 characters", result["bio"])
}

func TestParseError_Numeric(t *testing.T) {
	err := getValidationError(testStruct{
		Name:  "Alice",
		Email: "alice@mail.com",
		Age:   "abc",
	})

	result := customValidator.ParseError(err)

	assert.Equal(t, "Age must be a number", result["age"])
}

func TestParseError_Unique(t *testing.T) {
	err := getValidationError(testStructUnique{Email: "alice@mail.com"})

	result := customValidator.ParseError(err)

	assert.Equal(t, "Email already exists", result["email"])
}

func TestParseError_Default(t *testing.T) {
	err := getValidationError(testStructDefault{Code: "!!!"})

	result := customValidator.ParseError(err)

	assert.Equal(t, "Invalid value", result["code"])
}

func TestParseError_NonValidationError(t *testing.T) {
	result := customValidator.ParseError(assert.AnError)

	assert.Empty(t, result)
}
