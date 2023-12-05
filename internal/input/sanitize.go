package input

import (
	"regexp"
	"strings"
)

// Produces a commandline safe version of the given filename
// Does not allow ./ or ../ to be part of a filename
func SanitizeFileName(fileName string) string {
	result := fileName

	// Prevent directory manipulation
	result = strings.ReplaceAll(result, ".", "")
	result = strings.ReplaceAll(result, "/", "")
	result = strings.ReplaceAll(result, "\\", "")

	// Prevent injection of commands
	result = strings.ReplaceAll(result, "\"", "")
	result = strings.ReplaceAll(result, "'", "")
	result = strings.ReplaceAll(result, "`", "")
	result = strings.ReplaceAll(result, "´", "")

	return result
}

// Only allows subdirectories no ./ ../
func SanitizeDirectory(dir string) string {
	result := dir

	pattern := regexp.MustCompile("\\.+/")
	result = pattern.ReplaceAllString(result, "")

	return result
}
