//go:build darwin

package regex

import "regexp"

var fileFolderName = regexp.MustCompile(`^[^:\x00]+$`)
