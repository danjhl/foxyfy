package input_test

import (
	"foxyfy/internal/input"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitizeFileName(t *testing.T) {
	// Should sanitize directory manipulation
	assert.Equal(t, "notright", input.SanitizeFileName("../../not/right"))
	assert.Equal(t, "not_right", input.SanitizeFileName("./not_right"))

	// Should sanitize injection
	assert.Equal(t, " malicious ", input.SanitizeFileName("\" malicious \""))
	assert.Equal(t, " malicious ", input.SanitizeFileName("' malicious '"))

	// Should sanitize special characters
	assert.Equal(t, "special ", input.SanitizeFileName("special ,:*?<>|~´"))
}

func TestSanitizeDirectory(t *testing.T) {
	assert.Equal(t, "a/directory", input.SanitizeDirectory("a/./../directory"))
	assert.Equal(t, "a/directory", input.SanitizeDirectory("a/.../directory"))
}
