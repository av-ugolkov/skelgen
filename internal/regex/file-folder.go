package regex

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	ErrInvalidFileName   = fmt.Errorf("invalid file name")
	ErrInvalidFolderName = fmt.Errorf("invalid folder name")
)

var (
	linuxFileFolderName = regexp.MustCompile(`^[^/\x00]+$`)
	macosFileFolderName = regexp.MustCompile(`^[^:\x00]+$`)
)

var linuxIsValidate = func(s string) bool {
	return linuxFileFolderName.MatchString(s)
}

var macosIsValidate = func(s string) bool {
	return macosFileFolderName.MatchString(s)
}

var windowsIsValidate = func(s string) bool {
	if s == "" {
		return false
	}

	reserved := map[string]bool{
		"CON": true, "PRN": true, "AUX": true, "NUL": true,
		"COM1": true, "COM2": true, "COM3": true, "COM4": true, "COM5": true, "COM6": true, "COM7": true, "COM8": true, "COM9": true,
		"LPT1": true, "LPT2": true, "LPT3": true, "LPT4": true, "LPT5": true, "LPT6": true, "LPT7": true, "LPT8": true, "LPT9": true,
	}

	for _, r := range s {
		if r < 32 || strings.ContainsRune(`<>:"/\|?*`, r) {
			return false
		}
	}

	upper := strings.ToUpper(s)
	if dot := strings.IndexRune(upper, '.'); dot != -1 {
		upper = upper[:dot]
	}
	if reserved[upper] {
		return false
	}

	if s == "." || s == ".." ||
		strings.HasSuffix(s, " ") || strings.HasSuffix(s, ".") {
		return false
	}

	return true
}

func IsValidate(s string) bool {
	switch platform {
	case "linux":
		return linuxIsValidate(s)
	case "windows":
		return windowsIsValidate(s)
	case "macos":
		return macosIsValidate(s)
	default:
		return false
	}
}
