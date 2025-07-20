package validation

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidCEP(t *testing.T) {
	cep := "12345678"
	valid := IsValidCEP(cep)
	assert.Equal(t, true, valid)
}

func TestIsValidCEPInvalid(t *testing.T) {
	cep := "1234567"
	valid := IsValidCEP(cep)
	assert.Equal(t, false, valid)
}

func TestIsValidCEPInvalidLetter(t *testing.T) {
	cep := "1234568a"
	valid := IsValidCEP(cep)
	assert.Equal(t, false, valid)
}
