package input

import (
	"regexp"
	"strings"
)

// Produces a commandline safe and valid version of the given string
// Does not allow ./ or ../ to be part of a filename
func SanitizeFileName(str string) string {
	result := str

	illegal := []string{
		// Prevent directory manipulation
		".", "/", "\\",

		// Prevent injection of commands
		"\"", "'", "`", "´",

		// Replace special characters
		",", ":", "*", "?", "<", ">", "|", "~", "´",
	}

	for _, c := range illegal {
		result = strings.ReplaceAll(result, c, "")
	}

	return result
}

// Only allows subdirectories no ./ ../
func SanitizeDirectory(dir string) string {
	result := dir

	pattern := regexp.MustCompile("\\.+/")
	result = pattern.ReplaceAllString(result, "")

	return result
}
