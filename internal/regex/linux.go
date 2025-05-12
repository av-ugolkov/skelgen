//go:build linux

package regex

import "regexp"

var fileFolderName = regexp.MustCompile(`^[^/\x00]+$`)
