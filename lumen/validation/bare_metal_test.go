package validation

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestValidateBareMetalServerName_ValidInputs(t *testing.T) {
	validInputs := []string{
		"mybaremetal",
		"abcdefghijk-01235",
		"67890l-mnopqrstuvwxyz",
		"multiple.domain.levels",
		"m",
		strings.Repeat("a", 63),
		strings.Repeat(strings.Repeat("a", 49)+".", 5) + "com",
	}

	for _, input := range validInputs {
		err := ValidateBareMetalServerName(input)
		assert.Nil(t, err)
	}
}

func TestValidateBareMetalServerName_InvalidInputs(t *testing.T) {
	invalidInputs := []string{
		"",
		strings.Repeat("a", 64),
		strings.Repeat(strings.Repeat("a", 49)+".", 5) + "bike",
		".leading-period",
		"training-period.",
		"repeated..periods",
		"-leading-hyphen",
	}

	for _, invalidInput := range invalidInputs {
		err := ValidateBareMetalServerName(invalidInput)
		assert.NotNil(t, err)
	}
}

func TestValidationBareMetalUsername_ValidInputs(t *testing.T) {
	validInputs := []string{
		"username",
		"user-name",
		"user_name",
		"User-_name",
		"NotRoot",
	}

	for _, input := range validInputs {
		err := ValidateBareMetalUsername(input)
		assert.Nil(t, err)
	}
}

func TestValidateBareMetalUsername_InvalidInputs(t *testing.T) {
	invalidInputs := []string{
		"",
		"root",
		"ROOT",
		"RooT",
		"user*name",
		"#username",
		"@username",
	}

	for _, input := range invalidInputs {
		err := ValidateBareMetalUsername(input)
		assert.NotNil(t, err)
	}
}
